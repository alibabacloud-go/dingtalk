// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package yida_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SaveFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormDataRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
}

func (s SaveFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *SaveFormDataRequest) SetAppType(v string) *SaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) SetUserId(v string) *SaveFormDataRequest {
	s.UserId = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

type SaveFormDataResponseBody struct {
	// 表单实例ID
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SaveFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponseBody) SetResult(v string) *SaveFormDataResponseBody {
	s.Result = &v
	return s
}

type SaveFormDataResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponse) SetHeaders(v map[string]*string) *SaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

type SearchFormDatasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDatasHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDatasHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDatasHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDatasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDatasRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 根据表单内组件值查询
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 当前页
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 根据数据提交人工号查询
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表, 字符串格式，且为yyyy-MM-DD格式
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。字符串格式，且为yyyy-MM-DD格式。 和createFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)创建的数据。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式。 和modifiedFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)被修改的数据。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 指定排序字段
	DynamicOrder *string `json:"dynamicOrder,omitempty" xml:"dynamicOrder,omitempty"`
}

func (s SearchFormDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDatasRequest) SetAppType(v string) *SearchFormDatasRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDatasRequest) SetSystemToken(v string) *SearchFormDatasRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDatasRequest) SetUserId(v string) *SearchFormDatasRequest {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedFromTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedToTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

type SearchFormDatasResponseBody struct {
	// 当前页
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 符合条件的实例总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 实例详情列表
	Data []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s SearchFormDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBody) SetCurrentPage(v int64) *SearchFormDatasResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int64) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

type SearchFormDatasResponseBodyData struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 实例ID
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// 数据创建时间
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// 最近修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 表单id
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 模型id
	ModelUuid *string `json:"modelUuid,omitempty" xml:"modelUuid,omitempty"`
	// 发起人
	Originator *SearchFormDatasResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 修改者
	ModifyUser *SearchFormDatasResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 表单数据
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 流水号
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 表单实例原始格式值
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 数据版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 批次号
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) SetId(v int64) *SearchFormDatasResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreator(v string) *SearchFormDatasResponseBodyData {
	s.Creator = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifier(v string) *SearchFormDatasResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginator struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	Name *SearchFormDatasResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserId(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetName(v *SearchFormDatasResponseBodyDataOriginatorName) *SearchFormDatasResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetEmail(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.Email = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorName) SetType(v string) *SearchFormDatasResponseBodyDataOriginatorName {
	s.Type = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	Name *SearchFormDatasResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetName(v *SearchFormDatasResponseBodyDataModifyUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetEmail(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.Email = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserName) SetType(v string) *SearchFormDatasResponseBodyDataModifyUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponse) SetHeaders(v map[string]*string) *SearchFormDatasResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDatasResponse) SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse {
	s.Body = v
	return s
}

type LoginCodeGenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoginCodeGenHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenHeaders) GoString() string {
	return s.String()
}

func (s *LoginCodeGenHeaders) SetCommonHeaders(v map[string]*string) *LoginCodeGenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoginCodeGenHeaders) SetXAcsDingtalkAccessToken(v string) *LoginCodeGenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoginCodeGenRequest struct {
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LoginCodeGenRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenRequest) GoString() string {
	return s.String()
}

func (s *LoginCodeGenRequest) SetUserId(v string) *LoginCodeGenRequest {
	s.UserId = &v
	return s
}

type LoginCodeGenResponseBody struct {
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LoginCodeGenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponseBody) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponseBody) SetResult(v string) *LoginCodeGenResponseBody {
	s.Result = &v
	return s
}

type LoginCodeGenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoginCodeGenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginCodeGenResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponse) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponse) SetHeaders(v map[string]*string) *LoginCodeGenResponse {
	s.Headers = v
	return s
}

func (s *LoginCodeGenResponse) SetBody(v *LoginCodeGenResponseBody) *LoginCodeGenResponse {
	s.Body = v
	return s
}

type GetFormDataByIDHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormDataByIDHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormDataByIDHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormDataByIDRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetFormDataByIDRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDRequest) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDRequest) SetAppType(v string) *GetFormDataByIDRequest {
	s.AppType = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUserId(v string) *GetFormDataByIDRequest {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

type GetFormDataByIDResponseBody struct {
	// 发起人详情
	Originator *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 最后修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单实例ID
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// 表单数据详情
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormUuid(v string) *GetFormDataByIDResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	Name *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

type GetFormDataByIDResponseBodyOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInChinese(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInEnglish(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetType(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetFormDataByIDResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormDataByIDResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponse) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponse) SetHeaders(v map[string]*string) *GetFormDataByIDResponse {
	s.Headers = v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
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

func (client *Client) SaveFormData(request *SaveFormDataRequest) (_result *SaveFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormDataHeaders{}
	_result = &SaveFormDataResponse{}
	_body, _err := client.SaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, headers *SaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SaveFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDatas(request *SearchFormDatasRequest) (_result *SearchFormDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDatasHeaders{}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.SearchFormDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDatasWithOptions(request *SearchFormDatasRequest, headers *SearchFormDatasHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicOrder)) {
		body["dynamicOrder"] = request.DynamicOrder
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDatas"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginCodeGen(request *LoginCodeGenRequest) (_result *LoginCodeGenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoginCodeGenHeaders{}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.LoginCodeGenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginCodeGenWithOptions(request *LoginCodeGenRequest, headers *LoginCodeGenHeaders, runtime *util.RuntimeOptions) (_result *LoginCodeGenResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.DoROARequest(tea.String("LoginCodeGen"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/authorizations/loginCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFormDataByID(id *string, request *GetFormDataByIDRequest) (_result *GetFormDataByIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormDataByIDHeaders{}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.GetFormDataByIDWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormDataByIDWithOptions(id *string, request *GetFormDataByIDRequest, headers *GetFormDataByIDHeaders, runtime *util.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormDataByID"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/"+tea.StringValue(id)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
