// This file is auto-generated, don't edit it. Thanks.
package yida_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateOrUpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrUpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrUpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrUpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrUpdateFormDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"countrySelectField_l0c1cwiu":[{"value":"US"}]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"key":"currentNodeName","value":"当前审批节点名称","type":"TEXT","operator":"like","componentName":"TextField"}]。详情参考 https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrUpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataRequest) SetAppType(v string) *CreateOrUpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormDataJson(v string) *CreateOrUpdateFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormUuid(v string) *CreateOrUpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetNoExecuteExpression(v bool) *CreateOrUpdateFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSearchCondition(v string) *CreateOrUpdateFormDataRequest {
	s.SearchCondition = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSystemToken(v string) *CreateOrUpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUseAlias(v bool) *CreateOrUpdateFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUserId(v string) *CreateOrUpdateFormDataRequest {
	s.UserId = &v
	return s
}

type CreateOrUpdateFormDataResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponseBody) SetResult(v []*string) *CreateOrUpdateFormDataResponseBody {
	s.Result = v
	return s
}

type CreateOrUpdateFormDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponse) SetHeaders(v map[string]*string) *CreateOrUpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetStatusCode(v int32) *CreateOrUpdateFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetBody(v *CreateOrUpdateFormDataResponseBody) *CreateOrUpdateFormDataResponse {
	s.Body = v
	return s
}

type GetFormComponentAliasListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormComponentAliasListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListHeaders) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListHeaders) SetCommonHeaders(v map[string]*string) *GetFormComponentAliasListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormComponentAliasListHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormComponentAliasListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormComponentAliasListRequest struct {
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 未知
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFormComponentAliasListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListRequest) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListRequest) SetLanguage(v string) *GetFormComponentAliasListRequest {
	s.Language = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetSystemToken(v string) *GetFormComponentAliasListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetUserId(v string) *GetFormComponentAliasListRequest {
	s.UserId = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetVersion(v int64) *GetFormComponentAliasListRequest {
	s.Version = &v
	return s
}

type GetFormComponentAliasListResponseBody struct {
	Result []*GetFormComponentAliasListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetFormComponentAliasListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponseBody) SetResult(v []*GetFormComponentAliasListResponseBodyResult) *GetFormComponentAliasListResponseBody {
	s.Result = v
	return s
}

type GetFormComponentAliasListResponseBodyResult struct {
	Alias   *string `json:"alias,omitempty" xml:"alias,omitempty"`
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s GetFormComponentAliasListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponseBodyResult) SetAlias(v string) *GetFormComponentAliasListResponseBodyResult {
	s.Alias = &v
	return s
}

func (s *GetFormComponentAliasListResponseBodyResult) SetFieldId(v string) *GetFormComponentAliasListResponseBodyResult {
	s.FieldId = &v
	return s
}

type GetFormComponentAliasListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormComponentAliasListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormComponentAliasListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponse) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponse) SetHeaders(v map[string]*string) *GetFormComponentAliasListResponse {
	s.Headers = v
	return s
}

func (s *GetFormComponentAliasListResponse) SetStatusCode(v int32) *GetFormComponentAliasListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormComponentAliasListResponse) SetBody(v *GetFormComponentAliasListResponseBody) *GetFormComponentAliasListResponse {
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
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// FORM-AA28579F69644FC19A47FE267457E664ZVR1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool   `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *GetFormDataByIDRequest) SetFormUuid(v string) *GetFormDataByIDRequest {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUseAlias(v bool) *GetFormDataByIDRequest {
	s.UseAlias = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUserId(v string) *GetFormDataByIDRequest {
	s.UserId = &v
	return s
}

type GetFormDataByIDResponseBody struct {
	// example:
	//
	// {       "numberField_jcr0069o": 1,       "multiSelectField_jcr0069s": [         "选项三",         "选项二"       ],       "textareaField_jcr0069n": "duohang",       "employeeField_jcr0069x": [         "xxxx"       ],       "departmentField_jcr0069z": "xxxx",       "cascadeDate_jcr0069u": [         "1514736000000",         "1517328000000"       ],       "cascadeSelectField_jcr006a0": [         "part",         "part_b"       ],       "tableField_jcr006a1": [         {           "departmentField_jcr006ad": "xxxx",           "cascadeDate_jcr006aa": [             "1514736000000",             "1517328000000"           ],           "selectField_jcr006a6": "选项三",           "citySelectField_jcr006ac": [             "天津",             "天津市",             "河东区"           ],           "radioField_jcr006a5": "选项二",           "employeeField_jcr006ab": [             "xxxxxx",             "yyyyyy"           ],           "dateField_jcr006a9": 1517328000000,           "textField_jcr006a2": "明细下单行",           "textareaField_jcr006a3": "明细下多行",           "cascadeSelectField_jcr006ae": [             "product",             "product_a"           ],           "numberField_jcr006a4": 2,           "checkboxField_jcr006a7": [             "选项一",             "选项三",             "选项二"           ],           "multiSelectField_jcr006a8": [             "选项一",             "选项三",             "选项二"           ]         }       ],       "selectField_jcr0069q": "选项一",       "citySelectField_jcr0069y": [         "北京",         "北京市",         "东城区"       ],       "checkboxField_jcr0069r": [         "选项三",         "选项二"       ],       "textField_jcr0069m": "danhang",       "radioField_jcr0069p": "选项一",       "dateField_jcr0069t": 1516636800000     }
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// 33f6d221-17f8-42b7-836a-682b95a046c2
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// example:
	//
	// 2018-01-24 11:22:01
	ModifiedTimeGMT *string                                `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator      *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// example:
	//
	// 开发部
	DepartmentName *string                                    `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Email          *string                                    `json:"email,omitempty" xml:"email,omitempty"`
	Name           *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId         *string                                    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

type GetFormDataByIDResponseBodyOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetFormDataByIDResponse) SetStatusCode(v int32) *GetFormDataByIDResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
	s.Body = v
	return s
}

type GetInstanceByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// FORM-ADFC8E8E5ADE4B2F8FC2316CFC42A55CJLWZ
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxyy
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 未知
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdRequest) SetAppType(v string) *GetInstanceByIdRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceByIdRequest) SetFormUuid(v string) *GetInstanceByIdRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdRequest) SetLanguage(v string) *GetInstanceByIdRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceByIdRequest) SetSystemToken(v string) *GetInstanceByIdRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceByIdRequest) SetUseAlias(v bool) *GetInstanceByIdRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstanceByIdRequest) SetUserId(v string) *GetInstanceByIdRequest {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBody struct {
	ActionExecutor    []*GetInstanceByIdResponseBodyActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	ApprovedResult    *string                                      `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	CreateTimeGMT     *string                                      `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	Data              map[string]interface{}                       `json:"data,omitempty" xml:"data,omitempty"`
	FormUuid          *string                                      `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	InstanceStatus    *string                                      `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	ModifiedTimeGMT   *string                                      `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator        *GetInstanceByIdResponseBodyOriginator       `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	ProcessCode       *string                                      `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessInstanceId *string                                      `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string                                      `json:"title,omitempty" xml:"title,omitempty"`
	Version           *int64                                       `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstanceByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBody) SetActionExecutor(v []*GetInstanceByIdResponseBodyActionExecutor) *GetInstanceByIdResponseBody {
	s.ActionExecutor = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetApprovedResult(v string) *GetInstanceByIdResponseBody {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetCreateTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetData(v map[string]interface{}) *GetInstanceByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetFormUuid(v string) *GetInstanceByIdResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetInstanceStatus(v string) *GetInstanceByIdResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetModifiedTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetOriginator(v *GetInstanceByIdResponseBodyOriginator) *GetInstanceByIdResponseBody {
	s.Originator = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessCode(v string) *GetInstanceByIdResponseBody {
	s.ProcessCode = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessInstanceId(v string) *GetInstanceByIdResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetTitle(v string) *GetInstanceByIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVersion(v int64) *GetInstanceByIdResponseBody {
	s.Version = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutor struct {
	DeptName *string                                        `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                        `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstanceByIdResponseBodyActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetDeptName(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetEmail(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetName(v *GetInstanceByIdResponseBodyActionExecutorName) *GetInstanceByIdResponseBodyActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetUserId(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetType(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponseBodyOriginator struct {
	DeptName *string                                    `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                    `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstanceByIdResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginator) SetDeptName(v string) *GetInstanceByIdResponseBodyOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetEmail(v string) *GetInstanceByIdResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetName(v *GetInstanceByIdResponseBodyOriginatorName) *GetInstanceByIdResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetUserId(v string) *GetInstanceByIdResponseBodyOriginator {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetType(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponse) SetHeaders(v map[string]*string) *GetInstanceByIdResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceByIdResponse) SetStatusCode(v int32) *GetInstanceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceByIdResponse) SetBody(v *GetInstanceByIdResponseBody) *GetInstanceByIdResponse {
	s.Body = v
	return s
}

type GetInstanceIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceIdListHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceIdListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// ding123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// {"text_field":"123"}
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// 2199132092
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding123
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstanceIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListRequest) SetAppType(v string) *GetInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceIdListRequest) SetApprovedResult(v string) *GetInstanceIdListRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateToTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetFormUuid(v string) *GetInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceIdListRequest) SetInstanceStatus(v string) *GetInstanceIdListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceIdListRequest) SetLanguage(v string) *GetInstanceIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedToTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetOriginatorId(v string) *GetInstanceIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSearchFieldJson(v string) *GetInstanceIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSystemToken(v string) *GetInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceIdListRequest) SetTaskId(v string) *GetInstanceIdListRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetUseAlias(v bool) *GetInstanceIdListRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstanceIdListRequest) SetUserId(v string) *GetInstanceIdListRequest {
	s.UserId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageNumber(v int32) *GetInstanceIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageSize(v int32) *GetInstanceIdListRequest {
	s.PageSize = &v
	return s
}

type GetInstanceIdListResponseBody struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstanceIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponseBody) SetData(v []*string) *GetInstanceIdListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceIdListResponseBody) SetPageNumber(v int64) *GetInstanceIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetTotalCount(v int64) *GetInstanceIdListResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstanceIdListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponse) SetHeaders(v map[string]*string) *GetInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIdListResponse) SetStatusCode(v int32) *GetInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIdListResponse) SetBody(v *GetInstanceIdListResponseBody) *GetInstanceIdListResponse {
	s.Body = v
	return s
}

type GetInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// 例如按照创建时间升序再按照文本组件值升序排序则填{\"gmt_create\":\"+\",\"textField_1234\":\"+\"} ，详情参考 https://www.yuque.com/yida/support/agb8im#CQro8
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// example:
	//
	// manager123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 模式1：根据组件值模糊匹配，示例：{"textField_jcr0069m":"danhang","selectField_jcr0069q":"K"}     模式2: 采用数据管理的查询过滤条件，匹配功能更强大，示例：[{"key":"currentNodeName","value":"步凡","type":"TEXT","operator":"like","componentName":"TextField”}]，详情参考  https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// 2199132092
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesRequest) SetAppType(v string) *GetInstancesRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesRequest) SetApprovedResult(v string) *GetInstancesRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesRequest) SetCreateFromTimeGMT(v string) *GetInstancesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetCreateToTimeGMT(v string) *GetInstancesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetFormUuid(v string) *GetInstancesRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesRequest) SetInstanceStatus(v string) *GetInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesRequest) SetLanguage(v string) *GetInstancesRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedFromTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedToTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetOrderConfigJson(v string) *GetInstancesRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *GetInstancesRequest) SetOriginatorId(v string) *GetInstancesRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstancesRequest) SetSearchFieldJson(v string) *GetInstancesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstancesRequest) SetSystemToken(v string) *GetInstancesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesRequest) SetTaskId(v string) *GetInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstancesRequest) SetUseAlias(v bool) *GetInstancesRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstancesRequest) SetUserId(v string) *GetInstancesRequest {
	s.UserId = &v
	return s
}

func (s *GetInstancesRequest) SetPageNumber(v int32) *GetInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesRequest) SetPageSize(v int32) *GetInstancesRequest {
	s.PageSize = &v
	return s
}

type GetInstancesResponseBody struct {
	Data []*GetInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBody) SetData(v []*GetInstancesResponseBodyData) *GetInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBody) SetPageNumber(v int64) *GetInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesResponseBody) SetTotalCount(v int64) *GetInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstancesResponseBodyData struct {
	ActionExecutor    []*GetInstancesResponseBodyDataActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	ApprovedResult    *string                                       `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	CreateTimeGMT     *string                                       `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	Data              map[string]interface{}                        `json:"data,omitempty" xml:"data,omitempty"`
	FormUuid          *string                                       `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	InstanceStatus    *string                                       `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	ModifiedTimeGMT   *string                                       `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator        *GetInstancesResponseBodyDataOriginator       `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	ProcessCode       *string                                       `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessInstanceId *string                                       `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string                                       `json:"title,omitempty" xml:"title,omitempty"`
	Version           *int64                                        `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyData) SetActionExecutor(v []*GetInstancesResponseBodyDataActionExecutor) *GetInstancesResponseBodyData {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesResponseBodyData) SetApprovedResult(v string) *GetInstancesResponseBodyData {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetCreateTimeGMT(v string) *GetInstancesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetData(v map[string]interface{}) *GetInstancesResponseBodyData {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBodyData) SetFormUuid(v string) *GetInstancesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetInstanceStatus(v string) *GetInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetModifiedTimeGMT(v string) *GetInstancesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetOriginator(v *GetInstancesResponseBodyDataOriginator) *GetInstancesResponseBodyData {
	s.Originator = v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessCode(v string) *GetInstancesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessInstanceId(v string) *GetInstancesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetTitle(v string) *GetInstancesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetVersion(v int64) *GetInstancesResponseBodyData {
	s.Version = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutor struct {
	DeptName *string                                         `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                         `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstancesResponseBodyDataActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetDeptName(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetEmail(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetName(v *GetInstancesResponseBodyDataActionExecutorName) *GetInstancesResponseBodyDataActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetUserId(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetType(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.Type = &v
	return s
}

type GetInstancesResponseBodyDataOriginator struct {
	DeptName *string                                     `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                     `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstancesResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                     `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginator) SetDeptName(v string) *GetInstancesResponseBodyDataOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetEmail(v string) *GetInstancesResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetName(v *GetInstancesResponseBodyDataOriginatorName) *GetInstancesResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetUserId(v string) *GetInstancesResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetType(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.Type = &v
	return s
}

type GetInstancesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesResponse) SetHeaders(v map[string]*string) *GetInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesResponse) SetStatusCode(v int32) *GetInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesResponse) SetBody(v *GetInstancesResponseBody) *GetInstancesResponse {
	s.Body = v
	return s
}

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
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcpm6agt": "单行","employeeField_jcos0sar": ["workno"]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-NJYJZELV8YZRDEI2N5IQ7L6VEDMR1VE9GMPCJB
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) SetUseAlias(v bool) *SaveFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *SaveFormDataRequest) SetUserId(v string) *SaveFormDataRequest {
	s.UserId = &v
	return s
}

type SaveFormDataResponseBody struct {
	// example:
	//
	// FINST-XIA66E71N7HSLK7H4KOZ388EEOP03A46YAYRK1
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SaveFormDataResponse) SetStatusCode(v int32) *SaveFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

type SearchFormDataIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataIdListHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataIdListRequest struct {
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// dign1234
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s SearchFormDataIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListRequest) SetCreateFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetCreateToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetLanguage(v string) *SearchFormDataIdListRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetOriginatorId(v string) *SearchFormDataIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSearchFieldJson(v string) *SearchFormDataIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSystemToken(v string) *SearchFormDataIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetUseAlias(v bool) *SearchFormDataIdListRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetUserId(v string) *SearchFormDataIdListRequest {
	s.UserId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageNumber(v int32) *SearchFormDataIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageSize(v int32) *SearchFormDataIdListRequest {
	s.PageSize = &v
	return s
}

type SearchFormDataIdListResponseBody struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponseBody) SetData(v []*string) *SearchFormDataIdListResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetPageNumber(v int64) *SearchFormDataIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetTotalCount(v int64) *SearchFormDataIdListResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponse) SetHeaders(v map[string]*string) *SearchFormDataIdListResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataIdListResponse) SetStatusCode(v int32) *SearchFormDataIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataIdListResponse) SetBody(v *SearchFormDataIdListResponseBody) *SearchFormDataIdListResponse {
	s.Body = v
	return s
}

type SearchFormDataSecondGenerationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataSecondGenerationHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataSecondGenerationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataSecondGenerationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-09-10
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// 例如按照创建时间升序按照文本组件值升序排序则填{\"gmt_create\":\"+\",\"textField_1234\":\"+\"}。详情参考 https://www.yuque.com/yida/support/agb8im#CQro8
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// example:
	//
	// manager123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// [{"key":"currentNodeName","value":"当前审批节点名称","type":"TEXT","operator":"like","componentName":"TextField"}]。详情参考 https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationRequest) SetAppType(v string) *SearchFormDataSecondGenerationRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetUseAlias(v bool) *SearchFormDataSecondGenerationRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetUserId(v string) *SearchFormDataSecondGenerationRequest {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBody struct {
	Data []*SearchFormDataSecondGenerationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBody) SetData(v []*SearchFormDataSecondGenerationResponseBodyData) *SearchFormDataSecondGenerationResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyData struct {
	// example:
	//
	// 2021-05-01
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// example:
	//
	// ding12345
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// {"addressField_l0c1cwiy_id":"\"海南省/469027/国营红岗农场/111\"","associationFormField_l0c1hdz4_id":"\"[{\\\"formType\\\":\\\"receipt\\\",\\\"formUuid\\\":\\\"FORM-QQ866JB1QW8YM5XZZZ64VQB61OGM1MLWE1C0LG\\\",\\\"instanceId\\\":\\\"FINST-CC666Y6198RY0LAN39XGND212MSX3TFT95S0LN31\\\",\\\"subTitle\\\":\\\"{\\\\\\\"type\\\\\\\":\\\\\\\"div\\\\\\\",\\\\\\\"props\\\\\\\":{\\\\\\\"children\\\\\\\":{\\\\\\\"type\\\\\\\":\\\\\\\"a\\\\\\\",\\\\\\\"props\\\\\\\":{\\\\\\\"children\\\\\\\":\\\\\\\"查看签名\\\\\\\",\\\\\\\"className\\\\\\\":\\\\\\\"inst-cell-item-link\\\\\\\",\\\\\\\"style\\\\\\\":{\\\\\\\"cursor\\\\\\\":\\\\\\\"pointer\\\\\\\",\\\\\\\"color\\\\\\\":\\\\\\\"#0068ff\\\\\\\"}}},\\\\\\\"className\\\\\\\":\\\\\\\"inst-cell-item\\\\\\\"}}\\\",\\\"appType\\\":\\\"APP_K6IGJJ6PFAARLPDSWKXQ\\\",\\\"title\\\":\\\"1\\\"}]\"","countrySelectField_l0c1cwiu_id":["PG"],"imageField_l0c1cwit":"[{\"previewUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80\",\"size\":2610734,\"name\":\"Bts2Z0e6oxA.jpg\",\"downloadUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\",\"url\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\"}]","rateField_l0c1cwis_value":"3","editorField_l0c1cwiz":"<div><strong>你好，这是测试</strong></div>\r\n<div><span style=\"font-family: kaiti;background-color: #ff0000;color: #ffff00;\"><em><strong>测试</strong></em></span></div>\r\n<div>&nbsp;</div>","rateField_l0c1cwis":3,"countrySelectField_l0c1cwiu":[],"attachmentField_l0ghkwv3":"[{\"downloadUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\",\"name\":\"Bts2Z0e6oxA.jpg\",\"previewUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80\",\"size\":2610734,\"url\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\"}]","addressField_l0c1cwiy":"{\"address\":\"111\",\"regionIds\":[460000,469027,469023401],\"regionText\":[{\"en_US\":\"hai+nan+sheng\",\"zh_CN\":\"海南省\"},{\"en_US\":\"cheng+mai+xian\",\"zh_CN\":\"澄迈县\"},{\"en_US\":\"guo+ying+hong+gang+nong+chang\",\"zh_CN\":\"国营红岗农场\"}]}"}
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// FINST-J8766S91O2UYN87ZX3XOF1MY8MBA2912BSV0L24
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// [{"componentName":"CountrySelectField","dateType":null,"fieldData":{"value":[{"text":{"en_US":"Papua New Guinea","zh_CN":"巴布亚新几内亚","type":"i18n"},"value":"PG"}]},"fieldDataUpdated":false,"fieldId":"countrySelectField_l0c1cwiu","format":null,"formatControls":null,"listNum":null,"options":[{"label":{"$ref":"$[0].fieldData.value[0].text"},"value":"PG"}],"rowId":null},{"componentName":"AssociationFormField","dateType":null,"fieldData":{"value":[{"formType":"receipt","formUuid":"FORM-QQ866JB1QW8YM5XZZZ64VQB61OGM1MLWE1C0LG","instanceId":"FINST-CC666Y6198RY0LAN39XGND212MSX3TFT95S0LN31","subTitle":"{\"type\":\"div\",\"props\":{\"children\":{\"type\":\"a\",\"props\":{\"children\":\"查看签名\",\"className\":\"inst-cell-item-link\",\"style\":{\"cursor\":\"pointer\",\"color\":\"#0068ff\"}}},\"className\":\"inst-cell-item\"}}","appType":"APP_K6IGJJ6PFAARLPDSWKXQ","title":"1"}]},"fieldDataUpdated":false,"fieldId":"associationFormField_l0c1hdz4","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"ImageField","dateType":null,"fieldData":{"value":[{"previewUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80","size":2610734,"name":"Bts2Z0e6oxA.jpg","downloadUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download","url":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download"}]},"fieldDataUpdated":false,"fieldId":"imageField_l0c1cwit","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"AddressField","dateType":null,"fieldData":{"value":{"address":"111","regionIds":[460000,469027,469023401],"regionText":[{"en_US":"hai+nan+sheng","zh_CN":"海南省"},{"en_US":"cheng+mai+xian","zh_CN":"澄迈县"},{"en_US":"guo+ying+hong+gang+nong+chang","zh_CN":"国营红岗农场"}]}},"fieldDataUpdated":false,"fieldId":"addressField_l0c1cwiy","format":null,"formatControls":null,"listNum":null,"options":[{"label":{"pureEn_US":"hai+nan+sheng","en_US":"hai+nan+sheng","zh_CN":"海南省","type":"i18n","key":null},"value":460000},{"label":{"pureEn_US":"cheng+mai+xian","en_US":"cheng+mai+xian","zh_CN":"澄迈县","type":"i18n","key":null},"value":469027},{"label":{"pureEn_US":"guo+ying+hong+gang+nong+chang","en_US":"guo+ying+hong+gang+nong+chang","zh_CN":"国营红岗农场","type":"i18n","key":null},"value":469023401}],"rowId":null},{"componentName":"EditorField","dateType":null,"fieldData":{"value":"<div><strong>你好，这是测试</strong></div>\r\n<div><span style=\"font-family: kaiti;background-color: #ff0000;color: #ffff00;\"><em><strong>测试</strong></em></span></div>\r\n<div>&nbsp;</div>"},"fieldDataUpdated":false,"fieldId":"editorField_l0c1cwiz","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"RateField","dateType":null,"fieldData":{"value":"3"},"fieldDataUpdated":false,"fieldId":"rateField_l0c1cwis","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"AttachmentField","dateType":null,"fieldData":{"value":[{"previewUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80","size":2610734,"name":"Bts2Z0e6oxA.jpg","downloadUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download","url":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download"}]},"fieldDataUpdated":false,"fieldId":"attachmentField_l0ghkwv3","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null}]
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// example:
	//
	// manager123
	Modifier   *string                                                   `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifyUser *SearchFormDataSecondGenerationResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	Originator *SearchFormDataSecondGenerationResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// IMPORT-388664B1BAUVB3AYZE1RIUE88TDM1QI9WIOWK2
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// example:
	//
	// YIDA909202202250027
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 李四发起的请购单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationResponseBodyDataModifyUser) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationResponseBodyDataOriginator) *SearchFormDataSecondGenerationResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUser struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginator struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataSecondGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataSecondGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetStatusCode(v int32) *SearchFormDataSecondGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetBody(v *SearchFormDataSecondGenerationResponseBody) *SearchFormDataSecondGenerationResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// {"numberField_1ac":"+"}, 表示按照字段numberField_1ac升序排列
	DynamicOrder *string `json:"dynamicOrder,omitempty" xml:"dynamicOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	OriginatorId      *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// {"textField_jcr0069m":"danhang","textareaField_jcr0069n":"duohang","numberField_jcr0069o":["1","10"],"radioField_jcr0069p":"选项一","selectField_jcr0069q":"选项一","checkboxField_jcr0069r":["选项二"],"multiSelectField_jcr0069s":["选项二","选项三"],"dateField_jcr0069t":[1514736000000,1517414399000],"cascadeDate_jcr0069u":[[1514736000000,1517414399000],[1514736000000,1517414399000]],"employeeField_jcr0069x":["xxxxx"],"citySelectField_jcr0069y":["110000","110100","110101"],"departmentField_jcr0069z":1123456,"cascadeSelectField_jcr006a0":["part","part_b"],"tableField_jcr006a1":"明细数据"}
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 173112212211
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
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

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDatasRequest) SetSystemToken(v string) *SearchFormDatasRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDatasRequest) SetUseAlias(v bool) *SearchFormDatasRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDatasRequest) SetUserId(v string) *SearchFormDatasRequest {
	s.UserId = &v
	return s
}

type SearchFormDatasResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int64                             `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int64) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDatasResponseBodyData struct {
	// example:
	//
	// 2018-01-24 11:22:01
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// example:
	//
	// 1731234567
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// 1002
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// example:
	//
	// {"numberField_jcr0069o":1,"multiSelectField_jcr0069s":["选项三","选项二"],"textareaField_jcr0069n":"duohang","employeeField_jcr0069x":["xxxx"],"departmentField_jcr0069z":"xxxx","cascadeDate_jcr0069u":["1514736000000","1517328000000"],"cascadeSelectField_jcr006a0":["part","part_b"],"tableField_jcr006a1":[{"departmentField_jcr006ad":"xxxx","cascadeDate_jcr006aa":["1514736000000","1517328000000"],"selectField_jcr006a6":"选项三","citySelectField_jcr006ac":["天津","天津市","河东区"],"radioField_jcr006a5":"选项二","employeeField_jcr006ab":["xxxxxx","yyyyyy"],"dateField_jcr006a9":1517328000000,"textField_jcr006a2":"明细下单行","textareaField_jcr006a3":"明细下多行","cascadeSelectField_jcr006ae":["product","product_a"],"numberField_jcr006a4":2,"checkboxField_jcr006a7":["选项一","选项三","选项二"],"multiSelectField_jcr006a8":["选项一","选项三","选项二"]}],"selectField_jcr0069q":"选项一","citySelectField_jcr0069y":["北京","北京市","东城区"],"checkboxField_jcr0069r":["选项三","选项二"],"textField_jcr0069m":"danhang","radioField_jcr0069p":"选项一","dateField_jcr0069t":1516636800000}
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// FINST-BNKJDRF
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-EF6Y93URN24F1SCX15VA2P918LPEIJ2H3UFORCJ1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// {"textField":"124"}
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	ModelUuid     *string `json:"modelUuid,omitempty" xml:"modelUuid,omitempty"`
	// example:
	//
	// 2018-01-24 11:22:01
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// example:
	//
	// 1731234567
	ModifierUserId *string                                    `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	ModifyUser     *SearchFormDatasResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	Originator     *SearchFormDatasResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// Squence-XXX
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// example:
	//
	// 1234
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// example:
	//
	// 张三发起的表单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 3
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatorUserId(v string) *SearchFormDatasResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetDataId(v int64) *SearchFormDatasResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstanceId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifierUserId(v string) *SearchFormDatasResponseBodyData {
	s.ModifierUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	UserId   *string                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataModifyUserUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
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

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserName(v *SearchFormDatasResponseBodyDataModifyUserUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserName = v
	return s
}

type SearchFormDatasResponseBodyDataModifyUserUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetType(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginator struct {
	UserId   *string                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataOriginatorUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
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

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserName(v *SearchFormDatasResponseBodyDataOriginatorUserName) *SearchFormDatasResponseBodyDataOriginator {
	s.UserName = v
	return s
}

type SearchFormDatasResponseBodyDataOriginatorUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetType(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SearchFormDatasResponse) SetStatusCode(v int32) *SearchFormDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDatasResponse) SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse {
	s.Body = v
	return s
}

type StartInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *StartInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 18295
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcpm6agt": "单行","employeeField_jcos0sar": ["workno"]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-NJYJZELV8YZRDEI2N5IQ7L6VEDMR1VE9GMPCJB
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// TPROC--EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ4
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// [{ 	"key": "__optionalApproval_node_ocltdztr2b1", 	"value": ["5014533041684350"] }, { 	"key": "__optionalApproval_node_ocltdztr2b3", 	"value": ["5014533041684350", "01536610064226180419"] }, { 	"key": "__optionalApproval_node_oclte07cwn1", 	"value": ["01432910392321237660"] }]
	ProcessData *string `json:"processData,omitempty" xml:"processData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731234567
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetAppType(v string) *StartInstanceRequest {
	s.AppType = &v
	return s
}

func (s *StartInstanceRequest) SetDepartmentId(v string) *StartInstanceRequest {
	s.DepartmentId = &v
	return s
}

func (s *StartInstanceRequest) SetFormDataJson(v string) *StartInstanceRequest {
	s.FormDataJson = &v
	return s
}

func (s *StartInstanceRequest) SetFormUuid(v string) *StartInstanceRequest {
	s.FormUuid = &v
	return s
}

func (s *StartInstanceRequest) SetLanguage(v string) *StartInstanceRequest {
	s.Language = &v
	return s
}

func (s *StartInstanceRequest) SetProcessCode(v string) *StartInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartInstanceRequest) SetProcessData(v string) *StartInstanceRequest {
	s.ProcessData = &v
	return s
}

func (s *StartInstanceRequest) SetSystemToken(v string) *StartInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *StartInstanceRequest) SetUseAlias(v bool) *StartInstanceRequest {
	s.UseAlias = &v
	return s
}

func (s *StartInstanceRequest) SetUserId(v string) *StartInstanceRequest {
	s.UserId = &v
	return s
}

type StartInstanceResponseBody struct {
	// example:
	//
	// f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetResult(v string) *StartInstanceResponseBody {
	s.Result = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type UpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *UpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFormDataRequest struct {
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33f6d221-17f8-42b7-836a-682b95a046c2
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-AA28579F69644FC19A47FE267457E664ZVR1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcr0069m":"danhang","textareaField_jcr0069n":"duohang","numberField_jcr0069o":1,"radioField_jcr0069p":"选项一","selectField_jcr0069q":"选项一","checkboxField_jcr0069r":["选项二","选项三"],"multiSelectField_jcr0069s":["选项二","选项三"],"dateField_jcr0069t":1516636800000,"cascadeDate_jcr0069u":["1514736000000","1517328000000"],"employeeField_jcr0069x":["xxxxx"],"citySelectField_jcr0069y":["110000","110100","110101"],"departmentField_jcr0069z":1123456,"cascadeSelectField_jcr006a0":["part","part_b"],{"attachmentField_jna1lvyb":[{"downloadUrl":"https://www.aliwork.com/fileHandle?appType=default_tianshu_app&fileName=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt&instId=&type=download","name":"test.txt","previewUrl":"https://www.aliwork.com/inst/preview?appType=default_tianshu_app&fileName=test.txt&fileSize=4&downloadUrl=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt","url":"https://www.aliwork.com/fileHandle?appType=default_tianshu_app&fileName=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt&instId=&type=download","ext":"txt"}]},"tableField_jcr006a1":[{"cascadeDate_jcr006aa":["1514736000000","1517328000000"],"cascadeSelectField_jcr006ae":["product","product_a"],"checkboxField_jcr006a7":["选项一","选项二","选项三"],"citySelectField_jcr006ac":["120000","120100","120102"],"dateField_jcr006a9":1517328000000,"departmentField_jcr006ad":1123456,"employeeField_jcr006ab":["yyyyy","xxxxx"],"multiSelectField_jcr006a8":["选项一","选项二","选项三"],"numberField_jcr006a4":2,"radioField_jcr006a5":"选项二","selectField_jcr006a6":"选项三","textField_jcr006a2":"明细下单行","textareaField_jcr006a3":"明细下多行"}]}
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// example:
	//
	// false
	UseLatestVersion *bool `json:"useLatestVersion,omitempty" xml:"useLatestVersion,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormDataRequest) SetAppType(v string) *UpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormInstanceId(v string) *UpdateFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormUuid(v string) *UpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *UpdateFormDataRequest) SetLanguage(v string) *UpdateFormDataRequest {
	s.Language = &v
	return s
}

func (s *UpdateFormDataRequest) SetSystemToken(v string) *UpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateFormDataRequest) SetUpdateFormDataJson(v string) *UpdateFormDataRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseAlias(v bool) *UpdateFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseLatestVersion(v bool) *UpdateFormDataRequest {
	s.UseLatestVersion = &v
	return s
}

func (s *UpdateFormDataRequest) SetUserId(v string) *UpdateFormDataRequest {
	s.UserId = &v
	return s
}

type UpdateFormDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormDataResponse) SetHeaders(v map[string]*string) *UpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateFormDataResponse) SetStatusCode(v int32) *UpdateFormDataResponse {
	s.StatusCode = &v
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
// 新增或更新表单实例
//
// @param request - CreateOrUpdateFormDataRequest
//
// @param headers - CreateOrUpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormDataWithOptions(request *CreateOrUpdateFormDataRequest, headers *CreateOrUpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *CreateOrUpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("CreateOrUpdateFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/insertOrUpdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或更新表单实例
//
// @param request - CreateOrUpdateFormDataRequest
//
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormData(request *CreateOrUpdateFormDataRequest) (_result *CreateOrUpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrUpdateFormDataHeaders{}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.CreateOrUpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表单组件别名列表
//
// @param request - GetFormComponentAliasListRequest
//
// @param headers - GetFormComponentAliasListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormComponentAliasListResponse
func (client *Client) GetFormComponentAliasListWithOptions(appType *string, formUuid *string, request *GetFormComponentAliasListRequest, headers *GetFormComponentAliasListHeaders, runtime *util.RuntimeOptions) (_result *GetFormComponentAliasListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
	}

	realHeaders := make(map[string]*string)
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
		Action:      tea.String("GetFormComponentAliasList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/component/alias/" + tea.StringValue(appType) + "/" + tea.StringValue(formUuid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFormComponentAliasListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单组件别名列表
//
// @param request - GetFormComponentAliasListRequest
//
// @return GetFormComponentAliasListResponse
func (client *Client) GetFormComponentAliasList(appType *string, formUuid *string, request *GetFormComponentAliasListRequest) (_result *GetFormComponentAliasListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormComponentAliasListHeaders{}
	_result = &GetFormComponentAliasListResponse{}
	_body, _err := client.GetFormComponentAliasListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据表单 ID 查询实例详情
//
// @param request - GetFormDataByIDRequest
//
// @param headers - GetFormDataByIDHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormDataByIDResponse
func (client *Client) GetFormDataByIDWithOptions(id *string, request *GetFormDataByIDRequest, headers *GetFormDataByIDHeaders, runtime *util.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		query["useAlias"] = request.UseAlias
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
		Action:      tea.String("GetFormDataByID"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据表单 ID 查询实例详情
//
// @param request - GetFormDataByIDRequest
//
// @return GetFormDataByIDResponse
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

// Summary:
//
// 根据实例 ID 获取流程实例详情
//
// @param request - GetInstanceByIdRequest
//
// @param headers - GetInstanceByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceByIdWithOptions(id *string, request *GetInstanceByIdRequest, headers *GetInstanceByIdHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		query["useAlias"] = request.UseAlias
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
		Action:      tea.String("GetInstanceById"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instancesInfos/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据实例 ID 获取流程实例详情
//
// @param request - GetInstanceByIdRequest
//
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceById(id *string, request *GetInstanceByIdRequest) (_result *GetInstanceByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceByIdHeaders{}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.GetInstanceByIdWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件搜索流程实例 ID
//
// @param request - GetInstanceIdListRequest
//
// @param headers - GetInstanceIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdListWithOptions(request *GetInstanceIdListRequest, headers *GetInstanceIdListHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceIdListResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("GetInstanceIdList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instanceIds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索流程实例 ID
//
// @param request - GetInstanceIdListRequest
//
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdList(request *GetInstanceIdListRequest) (_result *GetInstanceIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceIdListHeaders{}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.GetInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据搜索条件获取流程表单实例详情
//
// @param request - GetInstancesRequest
//
// @param headers - GetInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancesResponse
func (client *Client) GetInstancesWithOptions(request *GetInstancesRequest, headers *GetInstancesHeaders, runtime *util.RuntimeOptions) (_result *GetInstancesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("GetInstances"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据搜索条件获取流程表单实例详情
//
// @param request - GetInstancesRequest
//
// @return GetInstancesResponse
func (client *Client) GetInstances(request *GetInstancesRequest) (_result *GetInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstancesHeaders{}
	_result = &GetInstancesResponse{}
	_body, _err := client.GetInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增表单实例
//
// @param request - SaveFormDataRequest
//
// @param headers - SaveFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFormDataResponse
func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, headers *SaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SaveFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增表单实例
//
// @param request - SaveFormDataRequest
//
// @return SaveFormDataResponse
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

// Summary:
//
// 根据条件搜索表单实例 ID 列表
//
// @param request - SearchFormDataIdListRequest
//
// @param headers - SearchFormDataIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdListWithOptions(appType *string, formUuid *string, request *SearchFormDataIdListRequest, headers *SearchFormDataIdListHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataIdListResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SearchFormDataIdList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/ids/" + tea.StringValue(appType) + "/" + tea.StringValue(formUuid)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例 ID 列表
//
// @param request - SearchFormDataIdListRequest
//
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdList(appType *string, formUuid *string, request *SearchFormDataIdListRequest) (_result *SearchFormDataIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataIdListHeaders{}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.SearchFormDataIdListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过高级检索条件查询表单实例
//
// @param request - SearchFormDataSecondGenerationRequest
//
// @param headers - SearchFormDataSecondGenerationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGenerationWithOptions(request *SearchFormDataSecondGenerationRequest, headers *SearchFormDataSecondGenerationHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SearchFormDataSecondGeneration"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/advances/queryAll"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过高级检索条件查询表单实例
//
// @param request - SearchFormDataSecondGenerationRequest
//
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGeneration(request *SearchFormDataSecondGenerationRequest) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationHeaders{}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.SearchFormDataSecondGenerationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例详情列表
//
// @param request - SearchFormDatasRequest
//
// @param headers - SearchFormDatasHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDatasResponse
func (client *Client) SearchFormDatasWithOptions(request *SearchFormDatasRequest, headers *SearchFormDatasHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicOrder)) {
		body["dynamicOrder"] = request.DynamicOrder
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SearchFormDatas"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例详情列表
//
// @param request - SearchFormDatasRequest
//
// @return SearchFormDatasResponse
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

// Summary:
//
// 发起新的流程实例
//
// @param request - StartInstanceRequest
//
// @param headers - StartInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, headers *StartInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessData)) {
		body["processData"] = request.ProcessData
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("StartInstance"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instances/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起新的流程实例
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartInstanceHeaders{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新表单实例
//
// @param request - UpdateFormDataRequest
//
// @param headers - UpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFormDataResponse
func (client *Client) UpdateFormDataWithOptions(request *UpdateFormDataRequest, headers *UpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *UpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestVersion)) {
		body["useLatestVersion"] = request.UseLatestVersion
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
		Action:      tea.String("UpdateFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新表单实例
//
// @param request - UpdateFormDataRequest
//
// @return UpdateFormDataResponse
func (client *Client) UpdateFormData(request *UpdateFormDataRequest) (_result *UpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFormDataHeaders{}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.UpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
