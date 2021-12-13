// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package h3yun_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type LoadBizFieldsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizFieldsHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsHeaders) SetCommonHeaders(v map[string]*string) *LoadBizFieldsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizFieldsHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizFieldsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizFieldsRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s LoadBizFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsRequest) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsRequest) SetSchemaCode(v string) *LoadBizFieldsRequest {
	s.SchemaCode = &v
	return s
}

type LoadBizFieldsResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *LoadBizFieldsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s LoadBizFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBody) SetCode(v string) *LoadBizFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizFieldsResponseBody) SetMessage(v string) *LoadBizFieldsResponseBody {
	s.Message = &v
	return s
}

func (s *LoadBizFieldsResponseBody) SetData(v *LoadBizFieldsResponseBodyData) *LoadBizFieldsResponseBody {
	s.Data = v
	return s
}

type LoadBizFieldsResponseBodyData struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 表单名称
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	// 字段、组件结构数组
	Fields []*LoadBizFieldsResponseBodyDataFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// 子表结构
	ChildForms []*LoadBizFieldsResponseBodyDataChildForms `json:"childForms,omitempty" xml:"childForms,omitempty" type:"Repeated"`
}

func (s LoadBizFieldsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyData) SetSchemaCode(v string) *LoadBizFieldsResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetFormName(v string) *LoadBizFieldsResponseBodyData {
	s.FormName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetFields(v []*LoadBizFieldsResponseBodyDataFields) *LoadBizFieldsResponseBodyData {
	s.Fields = v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetChildForms(v []*LoadBizFieldsResponseBodyDataChildForms) *LoadBizFieldsResponseBodyData {
	s.ChildForms = v
	return s
}

type LoadBizFieldsResponseBodyDataFields struct {
	// 显示名称
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 字段名称
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 字段、自定义组件的数据类型。Bool=逻辑型，DataTime=日期型、日期组件，Double=双精度数值型，Int=整形，Long=长整形，String=长文本，ShortString=短文本，ByteArray=二进制流， Image=图片类型、图片组件，File=附件类型组件，TimeSpan=时间段。Unit=参与者（单人），UnitArray=参与者（多人），Html=html类型，Xml=xml类型 BizObject=业务对象，BizObjectArray=业务对象数组、子表组件，Association=关联到其他对象、关联组件，AssociationArray=关联对象数组，Map=地图类型，Address=地址类型，Formula=公式型空间，Signature=签名控件，Plugin=文字识别Bool
	BizDataType *string `json:"bizDataType,omitempty" xml:"bizDataType,omitempty"`
}

func (s LoadBizFieldsResponseBodyDataFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataFields) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataFields) SetLabel(v string) *LoadBizFieldsResponseBodyDataFields {
	s.Label = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataFields) SetFieldName(v string) *LoadBizFieldsResponseBodyDataFields {
	s.FieldName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataFields) SetBizDataType(v string) *LoadBizFieldsResponseBodyDataFields {
	s.BizDataType = &v
	return s
}

type LoadBizFieldsResponseBodyDataChildForms struct {
	// 子表编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 子表名称
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	// 子表字段
	Fields []*LoadBizFieldsResponseBodyDataChildFormsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s LoadBizFieldsResponseBodyDataChildForms) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataChildForms) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetSchemaCode(v string) *LoadBizFieldsResponseBodyDataChildForms {
	s.SchemaCode = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetFormName(v string) *LoadBizFieldsResponseBodyDataChildForms {
	s.FormName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetFields(v []*LoadBizFieldsResponseBodyDataChildFormsFields) *LoadBizFieldsResponseBodyDataChildForms {
	s.Fields = v
	return s
}

type LoadBizFieldsResponseBodyDataChildFormsFields struct {
	// 显示名称
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 字段名或组件名
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 字段数据类型
	BizDataType *string `json:"bizDataType,omitempty" xml:"bizDataType,omitempty"`
}

func (s LoadBizFieldsResponseBodyDataChildFormsFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataChildFormsFields) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetLabel(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.Label = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetFieldName(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.FieldName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetBizDataType(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.BizDataType = &v
	return s
}

type LoadBizFieldsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoadBizFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponse) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponse) SetHeaders(v map[string]*string) *LoadBizFieldsResponse {
	s.Headers = v
	return s
}

func (s *LoadBizFieldsResponse) SetBody(v *LoadBizFieldsResponseBody) *LoadBizFieldsResponse {
	s.Body = v
	return s
}

type GetUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUsersHeaders) GoString() string {
	return s.String()
}

func (s *GetUsersHeaders) SetCommonHeaders(v map[string]*string) *GetUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUsersHeaders) SetXAcsDingtalkAccessToken(v string) *GetUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUsersRequest struct {
	// 部门id。
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 是否递归获取子级部门下的用户。默认值为false
	IsRecursive *bool `json:"isRecursive,omitempty" xml:"isRecursive,omitempty"`
}

func (s GetUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUsersRequest) GoString() string {
	return s.String()
}

func (s *GetUsersRequest) SetDepartmentId(v string) *GetUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *GetUsersRequest) SetIsRecursive(v bool) *GetUsersRequest {
	s.IsRecursive = &v
	return s
}

type GetUsersResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*GetUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBody) SetCode(v string) *GetUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBody) SetMessage(v string) *GetUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetUsersResponseBody) SetData(v []*GetUsersResponseBodyData) *GetUsersResponseBody {
	s.Data = v
	return s
}

type GetUsersResponseBodyData struct {
	// 用户id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 用户姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 性别.None=未指定，Man=男性，Female=女性
	Sex *string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 电话
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 直属组织id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 直属组织名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 作用域类型。Unspecified=未指定、Internal=内部组织机构、External=外部组织机构
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// 兼职部门id
	PartDepartmentIds []*string `json:"partDepartmentIds,omitempty" xml:"partDepartmentIds,omitempty" type:"Repeated"`
	// 排序号
	SortKey *int64 `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
}

func (s GetUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBodyData) SetId(v string) *GetUsersResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetUsersResponseBodyData) SetName(v string) *GetUsersResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUsersResponseBodyData) SetCode(v string) *GetUsersResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBodyData) SetSex(v string) *GetUsersResponseBodyData {
	s.Sex = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDescription(v string) *GetUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUsersResponseBodyData) SetMobile(v string) *GetUsersResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetUsersResponseBodyData) SetEmail(v string) *GetUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDepartmentId(v string) *GetUsersResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDepartmentName(v string) *GetUsersResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDomainType(v string) *GetUsersResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetUsersResponseBodyData) SetPartDepartmentIds(v []*string) *GetUsersResponseBodyData {
	s.PartDepartmentIds = v
	return s
}

func (s *GetUsersResponseBodyData) SetSortKey(v int64) *GetUsersResponseBodyData {
	s.SortKey = &v
	return s
}

type GetUsersResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponse) GoString() string {
	return s.String()
}

func (s *GetUsersResponse) SetHeaders(v map[string]*string) *GetUsersResponse {
	s.Headers = v
	return s
}

func (s *GetUsersResponse) SetBody(v *GetUsersResponseBody) *GetUsersResponse {
	s.Body = v
	return s
}

type GetRoleUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRoleUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersHeaders) GoString() string {
	return s.String()
}

func (s *GetRoleUsersHeaders) SetCommonHeaders(v map[string]*string) *GetRoleUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRoleUsersHeaders) SetXAcsDingtalkAccessToken(v string) *GetRoleUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRoleUsersRequest struct {
	// 角色id
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s GetRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *GetRoleUsersRequest) SetRoleId(v string) *GetRoleUsersRequest {
	s.RoleId = &v
	return s
}

type GetRoleUsersResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*GetRoleUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponseBody) SetCode(v string) *GetRoleUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetRoleUsersResponseBody) SetMessage(v string) *GetRoleUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetRoleUsersResponseBody) SetData(v []*GetRoleUsersResponseBodyData) *GetRoleUsersResponseBody {
	s.Data = v
	return s
}

type GetRoleUsersResponseBodyData struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 性别.None=未指定，Man=男性，Female=女性
	Sex *string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 所属部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 所属部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 所属范围。Internal=内部，External=外部
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// 兼职部门id集合（含主部门id）
	PartDepartmentIds []*string `json:"partDepartmentIds,omitempty" xml:"partDepartmentIds,omitempty" type:"Repeated"`
	// 排序值
	SortKey *int64 `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
}

func (s GetRoleUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponseBodyData) SetUserId(v string) *GetRoleUsersResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetName(v string) *GetRoleUsersResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetCode(v string) *GetRoleUsersResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetSex(v string) *GetRoleUsersResponseBodyData {
	s.Sex = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDescription(v string) *GetRoleUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetMobile(v string) *GetRoleUsersResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetEmail(v string) *GetRoleUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDepartmentId(v string) *GetRoleUsersResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDepartmentName(v string) *GetRoleUsersResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDomainType(v string) *GetRoleUsersResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetPartDepartmentIds(v []*string) *GetRoleUsersResponseBodyData {
	s.PartDepartmentIds = v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetSortKey(v int64) *GetRoleUsersResponseBodyData {
	s.SortKey = &v
	return s
}

type GetRoleUsersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponse) SetHeaders(v map[string]*string) *GetRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *GetRoleUsersResponse) SetBody(v *GetRoleUsersResponseBody) *GetRoleUsersResponse {
	s.Body = v
	return s
}

type LoadBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizObjectHeaders) SetCommonHeaders(v map[string]*string) *LoadBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizObjectRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 业务数据i实例id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
}

func (s LoadBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectRequest) GoString() string {
	return s.String()
}

func (s *LoadBizObjectRequest) SetSchemaCode(v string) *LoadBizObjectRequest {
	s.SchemaCode = &v
	return s
}

func (s *LoadBizObjectRequest) SetBizObjectId(v string) *LoadBizObjectRequest {
	s.BizObjectId = &v
	return s
}

type LoadBizObjectResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s LoadBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizObjectResponseBody) SetCode(v string) *LoadBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizObjectResponseBody) SetMessage(v string) *LoadBizObjectResponseBody {
	s.Message = &v
	return s
}

func (s *LoadBizObjectResponseBody) SetData(v map[string]interface{}) *LoadBizObjectResponseBody {
	s.Data = v
	return s
}

type LoadBizObjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoadBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectResponse) GoString() string {
	return s.String()
}

func (s *LoadBizObjectResponse) SetHeaders(v map[string]*string) *LoadBizObjectResponse {
	s.Headers = v
	return s
}

func (s *LoadBizObjectResponse) SetBody(v *LoadBizObjectResponseBody) *LoadBizObjectResponse {
	s.Body = v
	return s
}

type LoadBizObjectsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizObjectsHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsHeaders) SetCommonHeaders(v map[string]*string) *LoadBizObjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizObjectsHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizObjectsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizObjectsRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 分页页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页页大小。限制在1~500
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 返回的字段.仅支持传入主表的字段
	ReturnFields []*string `json:"returnFields,omitempty" xml:"returnFields,omitempty" type:"Repeated"`
	// 排序字段结构数组
	SortByFields []*LoadBizObjectsRequestSortByFields `json:"sortByFields,omitempty" xml:"sortByFields,omitempty" type:"Repeated"`
	// json格式的动态条件过滤器参数
	MatcherJson *string `json:"matcherJson,omitempty" xml:"matcherJson,omitempty"`
}

func (s LoadBizObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsRequest) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsRequest) SetSchemaCode(v string) *LoadBizObjectsRequest {
	s.SchemaCode = &v
	return s
}

func (s *LoadBizObjectsRequest) SetPageNumber(v int32) *LoadBizObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *LoadBizObjectsRequest) SetPageSize(v int32) *LoadBizObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *LoadBizObjectsRequest) SetReturnFields(v []*string) *LoadBizObjectsRequest {
	s.ReturnFields = v
	return s
}

func (s *LoadBizObjectsRequest) SetSortByFields(v []*LoadBizObjectsRequestSortByFields) *LoadBizObjectsRequest {
	s.SortByFields = v
	return s
}

func (s *LoadBizObjectsRequest) SetMatcherJson(v string) *LoadBizObjectsRequest {
	s.MatcherJson = &v
	return s
}

type LoadBizObjectsRequestSortByFields struct {
	// 排序字段名
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 排序方向。Ascending=升序，Descending=降序
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
}

func (s LoadBizObjectsRequestSortByFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsRequestSortByFields) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsRequestSortByFields) SetFieldName(v string) *LoadBizObjectsRequestSortByFields {
	s.FieldName = &v
	return s
}

func (s *LoadBizObjectsRequestSortByFields) SetDirection(v string) *LoadBizObjectsRequestSortByFields {
	s.Direction = &v
	return s
}

type LoadBizObjectsResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *LoadBizObjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s LoadBizObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponseBody) SetCode(v string) *LoadBizObjectsResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizObjectsResponseBody) SetMessage(v string) *LoadBizObjectsResponseBody {
	s.Message = &v
	return s
}

func (s *LoadBizObjectsResponseBody) SetData(v *LoadBizObjectsResponseBodyData) *LoadBizObjectsResponseBody {
	s.Data = v
	return s
}

type LoadBizObjectsResponseBodyData struct {
	// 页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 匹配条件的结果总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 业务数据实例数组
	BizObjects []map[string]interface{} `json:"bizObjects,omitempty" xml:"bizObjects,omitempty" type:"Repeated"`
}

func (s LoadBizObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponseBodyData) SetPageNumber(v int32) *LoadBizObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetPageSize(v int32) *LoadBizObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetTotalCount(v int32) *LoadBizObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetBizObjects(v []map[string]interface{}) *LoadBizObjectsResponseBodyData {
	s.BizObjects = v
	return s
}

type LoadBizObjectsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoadBizObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponse) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponse) SetHeaders(v map[string]*string) *LoadBizObjectsResponse {
	s.Headers = v
	return s
}

func (s *LoadBizObjectsResponse) SetBody(v *LoadBizObjectsResponseBody) *LoadBizObjectsResponse {
	s.Body = v
	return s
}

type DeleteProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteProcessesInstanceRequest struct {
	// 流程实例id
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 删除成功后，是否需要更新业务表单关联的流程实例id
	IsAutoUpdateBizObject *bool `json:"isAutoUpdateBizObject,omitempty" xml:"isAutoUpdateBizObject,omitempty"`
}

func (s DeleteProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceRequest) SetProcessInstanceId(v string) *DeleteProcessesInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *DeleteProcessesInstanceRequest) SetIsAutoUpdateBizObject(v bool) *DeleteProcessesInstanceRequest {
	s.IsAutoUpdateBizObject = &v
	return s
}

type DeleteProcessesInstanceResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceResponseBody) SetCode(v string) *DeleteProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProcessesInstanceResponseBody) SetMessage(v string) *DeleteProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

type DeleteProcessesInstanceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceResponse) SetHeaders(v map[string]*string) *DeleteProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProcessesInstanceResponse) SetBody(v *DeleteProcessesInstanceResponseBody) *DeleteProcessesInstanceResponse {
	s.Body = v
	return s
}

type QueryProcessesWorkItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProcessesWorkItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsHeaders) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsHeaders) SetCommonHeaders(v map[string]*string) *QueryProcessesWorkItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProcessesWorkItemsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProcessesWorkItemsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProcessesWorkItemsRequest struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s QueryProcessesWorkItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsRequest) SetProcessInstanceId(v string) *QueryProcessesWorkItemsRequest {
	s.ProcessInstanceId = &v
	return s
}

type QueryProcessesWorkItemsResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*QueryProcessesWorkItemsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryProcessesWorkItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBody) SetCode(v string) *QueryProcessesWorkItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBody) SetMessage(v string) *QueryProcessesWorkItemsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBody) SetData(v []*QueryProcessesWorkItemsResponseBodyData) *QueryProcessesWorkItemsResponseBody {
	s.Data = v
	return s
}

type QueryProcessesWorkItemsResponseBodyData struct {
	// 工作任务Id
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	// 工作项类型。Fill=普通工作项，Approve=审批类型工作项，Circulate=传阅
	WorkItemType *string `json:"workItemType,omitempty" xml:"workItemType,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用编码
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 工作项所关联的业务对象id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	// 工作流版本
	ProcessVersion *string `json:"processVersion,omitempty" xml:"processVersion,omitempty"`
	// 活动编码
	ActivityCode *string `json:"activityCode,omitempty" xml:"activityCode,omitempty"`
	// 当前活动名称
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// 显示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 状态。Waiting=等待的状态，Working=正在工作中的状态，Finished=处于完成状态，Canceled=已经被取消，Forwarded=已转交状态，Revoked=撤回
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 是否已完成
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// 接收时间
	ReceiveTimeGMT *string `json:"receiveTimeGMT,omitempty" xml:"receiveTimeGMT,omitempty"`
	// 开始这个任务的时间
	StartTimeGMT *string `json:"startTimeGMT,omitempty" xml:"startTimeGMT,omitempty"`
	// 完成时间
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// 对该工作项的意见
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 对该工作项是否同意
	IsApproval *bool `json:"isApproval,omitempty" xml:"isApproval,omitempty"`
	// 参与者
	Participant *QueryProcessesWorkItemsResponseBodyDataParticipant `json:"participant,omitempty" xml:"participant,omitempty" type:"Struct"`
	// 完成者
	Finisher *QueryProcessesWorkItemsResponseBodyDataFinisher `json:"finisher,omitempty" xml:"finisher,omitempty" type:"Struct"`
	// 转交工作的接收人。如无转接人，则为空
	Receiptor *QueryProcessesWorkItemsResponseBodyDataReceiptor `json:"receiptor,omitempty" xml:"receiptor,omitempty" type:"Struct"`
}

func (s QueryProcessesWorkItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetWorkItemId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.WorkItemId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetWorkItemType(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.WorkItemType = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetProcessInstanceId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetAppCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetSchemaCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetBizObjectId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetProcessVersion(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ProcessVersion = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetActivityCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ActivityCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetActivityName(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetDisplayName(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetState(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetIsFinish(v bool) *QueryProcessesWorkItemsResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetReceiveTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ReceiveTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetStartTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.StartTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetFinishTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetComment(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.Comment = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetIsApproval(v bool) *QueryProcessesWorkItemsResponseBodyData {
	s.IsApproval = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetParticipant(v *QueryProcessesWorkItemsResponseBodyDataParticipant) *QueryProcessesWorkItemsResponseBodyData {
	s.Participant = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetFinisher(v *QueryProcessesWorkItemsResponseBodyDataFinisher) *QueryProcessesWorkItemsResponseBodyData {
	s.Finisher = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetReceiptor(v *QueryProcessesWorkItemsResponseBodyDataReceiptor) *QueryProcessesWorkItemsResponseBodyData {
	s.Receiptor = v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataParticipant struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户直属的部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 用户直属的部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataParticipant) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataParticipant) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.UserId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.DepartmentName = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataFinisher struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户直属的部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 用户直属的部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataFinisher) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataFinisher) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.UserId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.DepartmentName = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataReceiptor struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户直属的部门id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 用户直属的部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataReceiptor) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataReceiptor) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.UserId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.DepartmentName = &v
	return s
}

type QueryProcessesWorkItemsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryProcessesWorkItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProcessesWorkItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponse) SetHeaders(v map[string]*string) *QueryProcessesWorkItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryProcessesWorkItemsResponse) SetBody(v *QueryProcessesWorkItemsResponseBody) *QueryProcessesWorkItemsResponse {
	s.Body = v
	return s
}

type UpdateBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectHeaders) SetCommonHeaders(v map[string]*string) *UpdateBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateBizObjectRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 业务数据id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	// 待修改的json格式业务数据
	BizObjectJson *string `json:"bizObjectJson,omitempty" xml:"bizObjectJson,omitempty"`
}

func (s UpdateBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectRequest) SetSchemaCode(v string) *UpdateBizObjectRequest {
	s.SchemaCode = &v
	return s
}

func (s *UpdateBizObjectRequest) SetBizObjectId(v string) *UpdateBizObjectRequest {
	s.BizObjectId = &v
	return s
}

func (s *UpdateBizObjectRequest) SetBizObjectJson(v string) *UpdateBizObjectRequest {
	s.BizObjectJson = &v
	return s
}

type UpdateBizObjectResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectResponseBody) SetCode(v string) *UpdateBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizObjectResponseBody) SetMessage(v string) *UpdateBizObjectResponseBody {
	s.Message = &v
	return s
}

type UpdateBizObjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectResponse) SetHeaders(v map[string]*string) *UpdateBizObjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizObjectResponse) SetBody(v *UpdateBizObjectResponseBody) *UpdateBizObjectResponse {
	s.Body = v
	return s
}

type QueryProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *QueryProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProcessesInstanceRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 业务数据id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
}

func (s QueryProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceRequest) SetSchemaCode(v string) *QueryProcessesInstanceRequest {
	s.SchemaCode = &v
	return s
}

func (s *QueryProcessesInstanceRequest) SetBizObjectId(v string) *QueryProcessesInstanceRequest {
	s.BizObjectId = &v
	return s
}

type QueryProcessesInstanceResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*QueryProcessesInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBody) SetCode(v string) *QueryProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProcessesInstanceResponseBody) SetMessage(v string) *QueryProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProcessesInstanceResponseBody) SetData(v []*QueryProcessesInstanceResponseBodyData) *QueryProcessesInstanceResponseBody {
	s.Data = v
	return s
}

type QueryProcessesInstanceResponseBodyData struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 钉钉流程Id
	DingTalkProcessId *string `json:"dingTalkProcessId,omitempty" xml:"dingTalkProcessId,omitempty"`
	// 流程名称
	ProcessDisplayName *string `json:"processDisplayName,omitempty" xml:"processDisplayName,omitempty"`
	// 工作流模板的版本
	ProcessVersion *int32 `json:"processVersion,omitempty" xml:"processVersion,omitempty"`
	// 流程所属的表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 流程关联的业务对象id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	// 流程所属的应用编码
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// 状态。Initiated=初始化完成，Starting=正在启动，Running=正在运行，Finishing=正在结束，Finished=已完成，Canceled=已取
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 流程发起人信息
	Originator *QueryProcessesInstanceResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 创建时间
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// 开始时间
	StartTimeGMT *string `json:"startTimeGMT,omitempty" xml:"startTimeGMT,omitempty"`
	// 完成时间
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
}

func (s QueryProcessesInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessInstanceId(v string) *QueryProcessesInstanceResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetDingTalkProcessId(v string) *QueryProcessesInstanceResponseBodyData {
	s.DingTalkProcessId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessDisplayName(v string) *QueryProcessesInstanceResponseBodyData {
	s.ProcessDisplayName = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessVersion(v int32) *QueryProcessesInstanceResponseBodyData {
	s.ProcessVersion = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetSchemaCode(v string) *QueryProcessesInstanceResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetBizObjectId(v string) *QueryProcessesInstanceResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetAppCode(v string) *QueryProcessesInstanceResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetState(v string) *QueryProcessesInstanceResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetOriginator(v *QueryProcessesInstanceResponseBodyDataOriginator) *QueryProcessesInstanceResponseBodyData {
	s.Originator = v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetCreatedTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetStartTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.StartTimeGMT = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetFinishTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

type QueryProcessesInstanceResponseBodyDataOriginator struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 所属组织id
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 所属组织名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
}

func (s QueryProcessesInstanceResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetUserId(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetName(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.Name = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetDepartmentId(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetDepartmentName(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

type QueryProcessesInstanceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponse) SetHeaders(v map[string]*string) *QueryProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryProcessesInstanceResponse) SetBody(v *QueryProcessesInstanceResponseBody) *QueryProcessesInstanceResponse {
	s.Body = v
	return s
}

type GetRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRolesHeaders) GoString() string {
	return s.String()
}

func (s *GetRolesHeaders) SetCommonHeaders(v map[string]*string) *GetRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRolesHeaders) SetXAcsDingtalkAccessToken(v string) *GetRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRolesResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *GetRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s GetRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBody) SetCode(v string) *GetRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRolesResponseBody) SetMessage(v string) *GetRolesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRolesResponseBody) SetData(v *GetRolesResponseBodyData) *GetRolesResponseBody {
	s.Data = v
	return s
}

type GetRolesResponseBodyData struct {
	// 角色组数组
	RoleGroups []*GetRolesResponseBodyDataRoleGroups `json:"roleGroups,omitempty" xml:"roleGroups,omitempty" type:"Repeated"`
	// 角色数组
	Roles []*GetRolesResponseBodyDataRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s GetRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyData) SetRoleGroups(v []*GetRolesResponseBodyDataRoleGroups) *GetRolesResponseBodyData {
	s.RoleGroups = v
	return s
}

func (s *GetRolesResponseBodyData) SetRoles(v []*GetRolesResponseBodyDataRoles) *GetRolesResponseBodyData {
	s.Roles = v
	return s
}

type GetRolesResponseBodyDataRoleGroups struct {
	// 组id
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 组编码
	GroupCode *string `json:"groupCode,omitempty" xml:"groupCode,omitempty"`
	// 所属企业id
	CompanyId *string `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// 可见性。All=全部可见、Normal=普通可见、OnlyAdmin=只有管理的时候可见、OnlyOrganization=本组织范围可见
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 状态。Active=激活、Inactive=未激活，Unspecified=未指定状态
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s GetRolesResponseBodyDataRoleGroups) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyDataRoleGroups) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupId(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupName(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupName = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupCode(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupCode = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetCompanyId(v string) *GetRolesResponseBodyDataRoleGroups {
	s.CompanyId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetVisibility(v string) *GetRolesResponseBodyDataRoleGroups {
	s.Visibility = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetState(v string) *GetRolesResponseBodyDataRoleGroups {
	s.State = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetDescription(v string) *GetRolesResponseBodyDataRoleGroups {
	s.Description = &v
	return s
}

type GetRolesResponseBodyDataRoles struct {
	// 角色id
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 角色编码
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 所属的角色组id
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 状态。Active=激活、Inactive=未激活，Unspecified=未指定状态
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 可见性。All=全部可见、Normal=普通可见、OnlyAdmin=只有管理的时候可见、OnlyOrganization=本组织范围可见
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 所属企业id
	CompanyId *string `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s GetRolesResponseBodyDataRoles) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyDataRoles) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyDataRoles) SetRoleId(v string) *GetRolesResponseBodyDataRoles {
	s.RoleId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetRoleName(v string) *GetRolesResponseBodyDataRoles {
	s.RoleName = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetRoleCode(v string) *GetRolesResponseBodyDataRoles {
	s.RoleCode = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetDescription(v string) *GetRolesResponseBodyDataRoles {
	s.Description = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetGroupId(v string) *GetRolesResponseBodyDataRoles {
	s.GroupId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetState(v string) *GetRolesResponseBodyDataRoles {
	s.State = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetVisibility(v string) *GetRolesResponseBodyDataRoles {
	s.Visibility = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetCompanyId(v string) *GetRolesResponseBodyDataRoles {
	s.CompanyId = &v
	return s
}

type GetRolesResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponse) GoString() string {
	return s.String()
}

func (s *GetRolesResponse) SetHeaders(v map[string]*string) *GetRolesResponse {
	s.Headers = v
	return s
}

func (s *GetRolesResponse) SetBody(v *GetRolesResponseBody) *GetRolesResponse {
	s.Body = v
	return s
}

type GetOrganizationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrganizationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationsHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrganizationsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrganizationsRequest struct {
	// 18f923a7-5a5e-426d-94ae-a55ad1a4b240
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s GetOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationsRequest) SetDepartmentId(v string) *GetOrganizationsRequest {
	s.DepartmentId = &v
	return s
}

type GetOrganizationsResponseBody struct {
	// 状态码。success=成功
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*GetOrganizationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponseBody) SetCode(v string) *GetOrganizationsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrganizationsResponseBody) SetMessage(v string) *GetOrganizationsResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrganizationsResponseBody) SetData(v []*GetOrganizationsResponseBodyData) *GetOrganizationsResponseBody {
	s.Data = v
	return s
}

type GetOrganizationsResponseBodyData struct {
	// 部门id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 父级部门id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 组织类型。Company=公司，OrganizationUnit=组织单元
	UnitType *string `json:"unitType,omitempty" xml:"unitType,omitempty"`
	// 排序值
	SortKey *int64 `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s GetOrganizationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponseBodyData) SetId(v string) *GetOrganizationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetParentId(v string) *GetOrganizationsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetName(v string) *GetOrganizationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetCode(v string) *GetOrganizationsResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetUnitType(v string) *GetOrganizationsResponseBodyData {
	s.UnitType = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetSortKey(v int64) *GetOrganizationsResponseBodyData {
	s.SortKey = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetDescription(v string) *GetOrganizationsResponseBodyData {
	s.Description = &v
	return s
}

type GetOrganizationsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponse) SetHeaders(v map[string]*string) *GetOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationsResponse) SetBody(v *GetOrganizationsResponseBody) *GetOrganizationsResponse {
	s.Body = v
	return s
}

type DeleteBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectHeaders) SetCommonHeaders(v map[string]*string) *DeleteBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteBizObjectRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 业务数据id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
}

func (s DeleteBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectRequest) SetSchemaCode(v string) *DeleteBizObjectRequest {
	s.SchemaCode = &v
	return s
}

func (s *DeleteBizObjectRequest) SetBizObjectId(v string) *DeleteBizObjectRequest {
	s.BizObjectId = &v
	return s
}

type DeleteBizObjectResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectResponseBody) SetCode(v string) *DeleteBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizObjectResponseBody) SetMessage(v string) *DeleteBizObjectResponseBody {
	s.Message = &v
	return s
}

type DeleteBizObjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectResponse) SetHeaders(v map[string]*string) *DeleteBizObjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizObjectResponse) SetBody(v *DeleteBizObjectResponseBody) *DeleteBizObjectResponse {
	s.Body = v
	return s
}

type QueryAppFunctionNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAppFunctionNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesHeaders) SetCommonHeaders(v map[string]*string) *QueryAppFunctionNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAppFunctionNodesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAppFunctionNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAppFunctionNodesRequest struct {
	// 应用编码
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
}

func (s QueryAppFunctionNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesRequest) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesRequest) SetAppCode(v string) *QueryAppFunctionNodesRequest {
	s.AppCode = &v
	return s
}

type QueryAppFunctionNodesResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*QueryAppFunctionNodesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s QueryAppFunctionNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponseBody) SetCode(v string) *QueryAppFunctionNodesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBody) SetMessage(v string) *QueryAppFunctionNodesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBody) SetData(v []*QueryAppFunctionNodesResponseBodyData) *QueryAppFunctionNodesResponseBody {
	s.Data = v
	return s
}

type QueryAppFunctionNodesResponseBodyData struct {
	// 节点编码。如果nodeType=FormModule,则为表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 节点所属的应用编码
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// 父节点的编码
	ParentCode *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	// 显示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 菜单可见类型。 Inactive=未指定 AllVisible=全部可见 PcVisible=仅pc可见 MobileVisible=仅移动端可见 InVisible=全部不可见
	NodeVisibleType *string `json:"nodeVisibleType,omitempty" xml:"nodeVisibleType,omitempty"`
	// 菜单节点类型。 AppPackage=应用程序 FormModule=列表模块(不能发起流程)、 WorkflowModule=流程列表模块(可以发起流程) ReportModule=报表模块 GroupModule=节点分组
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 菜单状态。 Inactive=未激活 Active=激活
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 排序编号
	SortKey *int64 `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	// 是否是系统节点，如果是则无法删除
	IsSystem *bool `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
}

func (s QueryAppFunctionNodesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponseBodyData) SetSchemaCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetAppCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetParentCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.ParentCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetDisplayName(v string) *QueryAppFunctionNodesResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetNodeVisibleType(v string) *QueryAppFunctionNodesResponseBodyData {
	s.NodeVisibleType = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetNodeType(v string) *QueryAppFunctionNodesResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetState(v string) *QueryAppFunctionNodesResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetSortKey(v int64) *QueryAppFunctionNodesResponseBodyData {
	s.SortKey = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetIsSystem(v bool) *QueryAppFunctionNodesResponseBodyData {
	s.IsSystem = &v
	return s
}

type QueryAppFunctionNodesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAppFunctionNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAppFunctionNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponse) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponse) SetHeaders(v map[string]*string) *QueryAppFunctionNodesResponse {
	s.Headers = v
	return s
}

func (s *QueryAppFunctionNodesResponse) SetBody(v *QueryAppFunctionNodesResponseBody) *QueryAppFunctionNodesResponse {
	s.Body = v
	return s
}

type CreateBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *CreateBizObjectHeaders) SetCommonHeaders(v map[string]*string) *CreateBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBizObjectRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 操作用户id。可从“获取用户信息”API获取
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// json格式的业务数据
	BizObjectJson *string `json:"bizObjectJson,omitempty" xml:"bizObjectJson,omitempty"`
	// 是否是草稿数据。true=草稿数据，false=生效数据
	IsDraft *bool `json:"isDraft,omitempty" xml:"isDraft,omitempty"`
}

func (s CreateBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectRequest) GoString() string {
	return s.String()
}

func (s *CreateBizObjectRequest) SetSchemaCode(v string) *CreateBizObjectRequest {
	s.SchemaCode = &v
	return s
}

func (s *CreateBizObjectRequest) SetOpUserId(v string) *CreateBizObjectRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBizObjectRequest) SetBizObjectJson(v string) *CreateBizObjectRequest {
	s.BizObjectJson = &v
	return s
}

func (s *CreateBizObjectRequest) SetIsDraft(v bool) *CreateBizObjectRequest {
	s.IsDraft = &v
	return s
}

type CreateBizObjectResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *CreateBizObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s CreateBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponseBody) SetCode(v string) *CreateBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBizObjectResponseBody) SetMessage(v string) *CreateBizObjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBizObjectResponseBody) SetData(v *CreateBizObjectResponseBodyData) *CreateBizObjectResponseBody {
	s.Data = v
	return s
}

type CreateBizObjectResponseBodyData struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 数据模型。DataList=本地存储的列表库，Workflow=工作流应用
	FormUsageType *string `json:"formUsageType,omitempty" xml:"formUsageType,omitempty"`
	// 表单业务数据id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	// 流程实例id
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s CreateBizObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponseBodyData) SetSchemaCode(v string) *CreateBizObjectResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetFormUsageType(v string) *CreateBizObjectResponseBodyData {
	s.FormUsageType = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetBizObjectId(v string) *CreateBizObjectResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetProcessInstanceId(v string) *CreateBizObjectResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

type CreateBizObjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponse) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponse) SetHeaders(v map[string]*string) *CreateBizObjectResponse {
	s.Headers = v
	return s
}

func (s *CreateBizObjectResponse) SetBody(v *CreateBizObjectResponseBody) *CreateBizObjectResponse {
	s.Body = v
	return s
}

type GetAppsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppsHeaders) GoString() string {
	return s.String()
}

func (s *GetAppsHeaders) SetCommonHeaders(v map[string]*string) *GetAppsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppsRequest struct {
	// 查询类型。All=全部，Solution=以解决方案编码为条件来查询应用，AppCode=以应用编码为条件来查询
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// 待查询的编码数组。queryType参数为All时，此值可为空
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppsRequest) GoString() string {
	return s.String()
}

func (s *GetAppsRequest) SetQueryType(v string) *GetAppsRequest {
	s.QueryType = &v
	return s
}

func (s *GetAppsRequest) SetValues(v []*string) *GetAppsRequest {
	s.Values = v
	return s
}

type GetAppsResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data []*GetAppsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppsResponseBody) SetCode(v string) *GetAppsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppsResponseBody) SetMessage(v string) *GetAppsResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppsResponseBody) SetData(v []*GetAppsResponseBodyData) *GetAppsResponseBody {
	s.Data = v
	return s
}

type GetAppsResponseBodyData struct {
	// 应用编码
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// 应用显示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 应用的来源类型。Custom=自开发的、Installed=安装的
	AppSource *string `json:"appSource,omitempty" xml:"appSource,omitempty"`
	// 应用状态。Enable=启用、Forbidden=禁用、Warring=预警
	AppState *string `json:"appState,omitempty" xml:"appState,omitempty"`
	// 应用所属的解决方案名称
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
}

func (s GetAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppsResponseBodyData) SetAppCode(v string) *GetAppsResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *GetAppsResponseBodyData) SetDisplayName(v string) *GetAppsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAppsResponseBodyData) SetAppSource(v string) *GetAppsResponseBodyData {
	s.AppSource = &v
	return s
}

func (s *GetAppsResponseBodyData) SetAppState(v string) *GetAppsResponseBodyData {
	s.AppState = &v
	return s
}

func (s *GetAppsResponseBodyData) SetSolution(v string) *GetAppsResponseBodyData {
	s.Solution = &v
	return s
}

type GetAppsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponse) GoString() string {
	return s.String()
}

func (s *GetAppsResponse) SetHeaders(v map[string]*string) *GetAppsResponse {
	s.Headers = v
	return s
}

func (s *GetAppsResponse) SetBody(v *GetAppsResponseBody) *GetAppsResponse {
	s.Body = v
	return s
}

type CreateProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProcessesInstanceRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 业务数据id
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	// 操作用户id。此为氚云的用户id，可从"获取用户数据API"获取
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreateProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceRequest) SetSchemaCode(v string) *CreateProcessesInstanceRequest {
	s.SchemaCode = &v
	return s
}

func (s *CreateProcessesInstanceRequest) SetBizObjectId(v string) *CreateProcessesInstanceRequest {
	s.BizObjectId = &v
	return s
}

func (s *CreateProcessesInstanceRequest) SetOpUserId(v string) *CreateProcessesInstanceRequest {
	s.OpUserId = &v
	return s
}

type CreateProcessesInstanceResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务响应结果
	Data *CreateProcessesInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s CreateProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponseBody) SetCode(v string) *CreateProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProcessesInstanceResponseBody) SetMessage(v string) *CreateProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProcessesInstanceResponseBody) SetData(v *CreateProcessesInstanceResponseBodyData) *CreateProcessesInstanceResponseBody {
	s.Data = v
	return s
}

type CreateProcessesInstanceResponseBodyData struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s CreateProcessesInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponseBodyData) SetProcessInstanceId(v string) *CreateProcessesInstanceResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

type CreateProcessesInstanceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponse) SetHeaders(v map[string]*string) *CreateProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessesInstanceResponse) SetBody(v *CreateProcessesInstanceResponseBody) *CreateProcessesInstanceResponse {
	s.Body = v
	return s
}

type BatchInsertBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchInsertBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectHeaders) SetCommonHeaders(v map[string]*string) *BatchInsertBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchInsertBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *BatchInsertBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchInsertBizObjectRequest struct {
	// 表单编码
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	// 操作用户id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 待新增的业对象json数组
	BizObjectJsonArray []*string `json:"bizObjectJsonArray,omitempty" xml:"bizObjectJsonArray,omitempty" type:"Repeated"`
	// 是否是草稿数据。true=创建草稿数据，false=创建生效数据
	IsDraft *bool `json:"isDraft,omitempty" xml:"isDraft,omitempty"`
}

func (s BatchInsertBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectRequest) SetSchemaCode(v string) *BatchInsertBizObjectRequest {
	s.SchemaCode = &v
	return s
}

func (s *BatchInsertBizObjectRequest) SetOpUserId(v string) *BatchInsertBizObjectRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchInsertBizObjectRequest) SetBizObjectJsonArray(v []*string) *BatchInsertBizObjectRequest {
	s.BizObjectJsonArray = v
	return s
}

func (s *BatchInsertBizObjectRequest) SetIsDraft(v bool) *BatchInsertBizObjectRequest {
	s.IsDraft = &v
	return s
}

type BatchInsertBizObjectResponseBody struct {
	// 状态码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 提示信息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *BatchInsertBizObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s BatchInsertBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponseBody) SetCode(v string) *BatchInsertBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchInsertBizObjectResponseBody) SetMessage(v string) *BatchInsertBizObjectResponseBody {
	s.Message = &v
	return s
}

func (s *BatchInsertBizObjectResponseBody) SetData(v *BatchInsertBizObjectResponseBodyData) *BatchInsertBizObjectResponseBody {
	s.Data = v
	return s
}

type BatchInsertBizObjectResponseBodyData struct {
	// 成功新增的业务对象id数组
	BizObjectIds []*string `json:"bizObjectIds,omitempty" xml:"bizObjectIds,omitempty" type:"Repeated"`
	// 新增成功的流程实例id数组
	ProcessIds []*string `json:"processIds,omitempty" xml:"processIds,omitempty" type:"Repeated"`
	// 新增失败的数据数组
	FailedDatas []*string `json:"failedDatas,omitempty" xml:"failedDatas,omitempty" type:"Repeated"`
	// 失败的提示信息数组
	FailedMessages []*string `json:"failedMessages,omitempty" xml:"failedMessages,omitempty" type:"Repeated"`
}

func (s BatchInsertBizObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponseBodyData) SetBizObjectIds(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.BizObjectIds = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetProcessIds(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.ProcessIds = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetFailedDatas(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.FailedDatas = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetFailedMessages(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.FailedMessages = v
	return s
}

type BatchInsertBizObjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchInsertBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchInsertBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponse) SetHeaders(v map[string]*string) *BatchInsertBizObjectResponse {
	s.Headers = v
	return s
}

func (s *BatchInsertBizObjectResponse) SetBody(v *BatchInsertBizObjectResponseBody) *BatchInsertBizObjectResponse {
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

func (client *Client) LoadBizFields(request *LoadBizFieldsRequest) (_result *LoadBizFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizFieldsHeaders{}
	_result = &LoadBizFieldsResponse{}
	_body, _err := client.LoadBizFieldsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizFieldsWithOptions(request *LoadBizFieldsRequest, headers *LoadBizFieldsHeaders, runtime *util.RuntimeOptions) (_result *LoadBizFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
	_result = &LoadBizFieldsResponse{}
	_body, _err := client.DoROARequest(tea.String("LoadBizFields"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/loadBizFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUsers(request *GetUsersRequest) (_result *GetUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUsersHeaders{}
	_result = &GetUsersResponse{}
	_body, _err := client.GetUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUsersWithOptions(request *GetUsersRequest, headers *GetUsersHeaders, runtime *util.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecursive)) {
		query["isRecursive"] = request.IsRecursive
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
	_result = &GetUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUsers"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoleUsers(request *GetRoleUsersRequest) (_result *GetRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRoleUsersHeaders{}
	_result = &GetRoleUsersResponse{}
	_body, _err := client.GetRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleUsersWithOptions(request *GetRoleUsersRequest, headers *GetRoleUsersHeaders, runtime *util.RuntimeOptions) (_result *GetRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["roleId"] = request.RoleId
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
	_result = &GetRoleUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRoleUsers"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/roles/roleUsers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadBizObject(request *LoadBizObjectRequest) (_result *LoadBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizObjectHeaders{}
	_result = &LoadBizObjectResponse{}
	_body, _err := client.LoadBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizObjectWithOptions(request *LoadBizObjectRequest, headers *LoadBizObjectHeaders, runtime *util.RuntimeOptions) (_result *LoadBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
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
	_result = &LoadBizObjectResponse{}
	_body, _err := client.DoROARequest(tea.String("LoadBizObject"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances/loadInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadBizObjects(request *LoadBizObjectsRequest) (_result *LoadBizObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizObjectsHeaders{}
	_result = &LoadBizObjectsResponse{}
	_body, _err := client.LoadBizObjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizObjectsWithOptions(request *LoadBizObjectsRequest, headers *LoadBizObjectsHeaders, runtime *util.RuntimeOptions) (_result *LoadBizObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFields)) {
		body["returnFields"] = request.ReturnFields
	}

	if !tea.BoolValue(util.IsUnset(request.SortByFields)) {
		body["sortByFields"] = request.SortByFields
	}

	if !tea.BoolValue(util.IsUnset(request.MatcherJson)) {
		body["matcherJson"] = request.MatcherJson
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
	_result = &LoadBizObjectsResponse{}
	_body, _err := client.DoROARequest(tea.String("LoadBizObjects"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProcessesInstance(request *DeleteProcessesInstanceRequest) (_result *DeleteProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteProcessesInstanceHeaders{}
	_result = &DeleteProcessesInstanceResponse{}
	_body, _err := client.DeleteProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProcessesInstanceWithOptions(request *DeleteProcessesInstanceRequest, headers *DeleteProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoUpdateBizObject)) {
		query["isAutoUpdateBizObject"] = request.IsAutoUpdateBizObject
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
	_result = &DeleteProcessesInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteProcessesInstance"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/h3yun/processes/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProcessesWorkItems(request *QueryProcessesWorkItemsRequest) (_result *QueryProcessesWorkItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProcessesWorkItemsHeaders{}
	_result = &QueryProcessesWorkItemsResponse{}
	_body, _err := client.QueryProcessesWorkItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProcessesWorkItemsWithOptions(request *QueryProcessesWorkItemsRequest, headers *QueryProcessesWorkItemsHeaders, runtime *util.RuntimeOptions) (_result *QueryProcessesWorkItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
	_result = &QueryProcessesWorkItemsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryProcessesWorkItems"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/processes/workItems"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizObject(request *UpdateBizObjectRequest) (_result *UpdateBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBizObjectHeaders{}
	_result = &UpdateBizObjectResponse{}
	_body, _err := client.UpdateBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizObjectWithOptions(request *UpdateBizObjectRequest, headers *UpdateBizObjectHeaders, runtime *util.RuntimeOptions) (_result *UpdateBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		body["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectJson)) {
		body["bizObjectJson"] = request.BizObjectJson
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
	_result = &UpdateBizObjectResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateBizObject"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProcessesInstance(request *QueryProcessesInstanceRequest) (_result *QueryProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProcessesInstanceHeaders{}
	_result = &QueryProcessesInstanceResponse{}
	_body, _err := client.QueryProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProcessesInstanceWithOptions(request *QueryProcessesInstanceRequest, headers *QueryProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
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
	_result = &QueryProcessesInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryProcessesInstance"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/processes/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoles() (_result *GetRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRolesHeaders{}
	_result = &GetRolesResponse{}
	_body, _err := client.GetRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRolesWithOptions(headers *GetRolesHeaders, runtime *util.RuntimeOptions) (_result *GetRolesResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRoles"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizations(request *GetOrganizationsRequest) (_result *GetOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationsHeaders{}
	_result = &GetOrganizationsResponse{}
	_body, _err := client.GetOrganizationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationsWithOptions(request *GetOrganizationsRequest, headers *GetOrganizationsHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &GetOrganizationsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOrganizations"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBizObject(request *DeleteBizObjectRequest) (_result *DeleteBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteBizObjectHeaders{}
	_result = &DeleteBizObjectResponse{}
	_body, _err := client.DeleteBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBizObjectWithOptions(request *DeleteBizObjectRequest, headers *DeleteBizObjectHeaders, runtime *util.RuntimeOptions) (_result *DeleteBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
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
	_result = &DeleteBizObjectResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteBizObject"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAppFunctionNodes(request *QueryAppFunctionNodesRequest) (_result *QueryAppFunctionNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAppFunctionNodesHeaders{}
	_result = &QueryAppFunctionNodesResponse{}
	_body, _err := client.QueryAppFunctionNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAppFunctionNodesWithOptions(request *QueryAppFunctionNodesRequest, headers *QueryAppFunctionNodesHeaders, runtime *util.RuntimeOptions) (_result *QueryAppFunctionNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["appCode"] = request.AppCode
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
	_result = &QueryAppFunctionNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAppFunctionNodes"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/h3yun/apps/functionNodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBizObject(request *CreateBizObjectRequest) (_result *CreateBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBizObjectHeaders{}
	_result = &CreateBizObjectResponse{}
	_body, _err := client.CreateBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBizObjectWithOptions(request *CreateBizObjectRequest, headers *CreateBizObjectHeaders, runtime *util.RuntimeOptions) (_result *CreateBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectJson)) {
		body["bizObjectJson"] = request.BizObjectJson
	}

	if !tea.BoolValue(util.IsUnset(request.IsDraft)) {
		body["isDraft"] = request.IsDraft
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
	_result = &CreateBizObjectResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateBizObject"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApps(request *GetAppsRequest) (_result *GetAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppsHeaders{}
	_result = &GetAppsResponse{}
	_body, _err := client.GetAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppsWithOptions(request *GetAppsRequest, headers *GetAppsHeaders, runtime *util.RuntimeOptions) (_result *GetAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.Values)) {
		body["values"] = request.Values
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
	_result = &GetAppsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApps"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/h3yun/apps/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProcessesInstance(request *CreateProcessesInstanceRequest) (_result *CreateProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProcessesInstanceHeaders{}
	_result = &CreateProcessesInstanceResponse{}
	_body, _err := client.CreateProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProcessesInstanceWithOptions(request *CreateProcessesInstanceRequest, headers *CreateProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		body["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
	_result = &CreateProcessesInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateProcessesInstance"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/h3yun/processes/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchInsertBizObject(request *BatchInsertBizObjectRequest) (_result *BatchInsertBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchInsertBizObjectHeaders{}
	_result = &BatchInsertBizObjectResponse{}
	_body, _err := client.BatchInsertBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchInsertBizObjectWithOptions(request *BatchInsertBizObjectRequest, headers *BatchInsertBizObjectHeaders, runtime *util.RuntimeOptions) (_result *BatchInsertBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectJsonArray)) {
		body["bizObjectJsonArray"] = request.BizObjectJsonArray
	}

	if !tea.BoolValue(util.IsUnset(request.IsDraft)) {
		body["isDraft"] = request.IsDraft
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
	_result = &BatchInsertBizObjectResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchInsertBizObject"), tea.String("h3yun_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/h3yun/forms/instances/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
