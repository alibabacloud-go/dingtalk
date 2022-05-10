// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package industry_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CustomizeContactCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactCreateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactCreateRequest struct {
	// 通讯录管理员UserId
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 自定义通讯录名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 在自定义通讯录列表中的排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s CustomizeContactCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateRequest) SetManagerIdList(v []*string) *CustomizeContactCreateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactCreateRequest) SetName(v string) *CustomizeContactCreateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactCreateRequest) SetOrder(v int64) *CustomizeContactCreateRequest {
	s.Order = &v
	return s
}

type CustomizeContactCreateResponseBody struct {
	// 自定义通讯录信息
	Content *CustomizeContactCreateResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s CustomizeContactCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponseBody) SetContent(v *CustomizeContactCreateResponseBodyContent) *CustomizeContactCreateResponseBody {
	s.Content = v
	return s
}

type CustomizeContactCreateResponseBodyContent struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 自定义通讯录名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 在自定义通讯录列表中的排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 根部们Id
	RootDeptId *int64 `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s CustomizeContactCreateResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponseBodyContent) SetCode(v string) *CustomizeContactCreateResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetName(v string) *CustomizeContactCreateResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetOrder(v int64) *CustomizeContactCreateResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetRootDeptId(v int64) *CustomizeContactCreateResponseBodyContent {
	s.RootDeptId = &v
	return s
}

type CustomizeContactCreateResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponse) SetHeaders(v map[string]*string) *CustomizeContactCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactCreateResponse) SetBody(v *CustomizeContactCreateResponseBody) *CustomizeContactCreateResponse {
	s.Body = v
	return s
}

type CustomizeContactDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeleteRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CustomizeContactDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteRequest) SetCode(v string) *CustomizeContactDeleteRequest {
	s.Code = &v
	return s
}

type CustomizeContactDeleteResponseBody struct {
	// 是否操作成功
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteResponseBody) SetContent(v bool) *CustomizeContactDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeleteResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeleteResponse) SetBody(v *CustomizeContactDeleteResponseBody) *CustomizeContactDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptCreateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptCreateRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门主管列表
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 上级部门Id
	ParentDeptId *int64 `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// 引用的内部通讯录部门Id
	RefId *int64 `json:"refId,omitempty" xml:"refId,omitempty"`
	// 部门类型 0:普通部门 1:引用部门
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateRequest) SetCode(v string) *CustomizeContactDeptCreateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetManagerIdList(v []*string) *CustomizeContactDeptCreateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetName(v string) *CustomizeContactDeptCreateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetOrder(v int64) *CustomizeContactDeptCreateRequest {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetParentDeptId(v int64) *CustomizeContactDeptCreateRequest {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetRefId(v int64) *CustomizeContactDeptCreateRequest {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetType(v int64) *CustomizeContactDeptCreateRequest {
	s.Type = &v
	return s
}

type CustomizeContactDeptCreateResponseBody struct {
	// 部门Id
	Content *int64 `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateResponseBody) SetContent(v int64) *CustomizeContactDeptCreateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptCreateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeptCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptCreateResponse) SetBody(v *CustomizeContactDeptCreateResponseBody) *CustomizeContactDeptCreateResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptDeleteRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteRequest) SetCode(v string) *CustomizeContactDeptDeleteRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptDeleteRequest) SetDeptId(v int64) *CustomizeContactDeptDeleteRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptDeleteResponseBody struct {
	// 操作结果
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteResponseBody) SetContent(v bool) *CustomizeContactDeptDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptDeleteResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeptDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptDeleteResponse) SetBody(v *CustomizeContactDeptDeleteResponseBody) *CustomizeContactDeptDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptInfoRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoRequest) SetCode(v string) *CustomizeContactDeptInfoRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptInfoRequest) SetDeptId(v int64) *CustomizeContactDeptInfoRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptInfoResponseBody struct {
	// 部门信息
	Content *CustomizeContactDeptInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s CustomizeContactDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponseBody) SetContent(v *CustomizeContactDeptInfoResponseBodyContent) *CustomizeContactDeptInfoResponseBody {
	s.Content = v
	return s
}

type CustomizeContactDeptInfoResponseBodyContent struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 部门主管列表
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 上级部门Id
	ParentDeptId *int64 `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// 引用的内部通讯录部门Id
	RefId *int64 `json:"refId,omitempty" xml:"refId,omitempty"`
	// 部门类型 0:普通部门 1:引用部门
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetCode(v string) *CustomizeContactDeptInfoResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetManagerIdList(v []*string) *CustomizeContactDeptInfoResponseBodyContent {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetName(v string) *CustomizeContactDeptInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetOrder(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetParentDeptId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetRefId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetType(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Type = &v
	return s
}

type CustomizeContactDeptInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptInfoResponse) SetBody(v *CustomizeContactDeptInfoResponseBody) *CustomizeContactDeptInfoResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptListRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptListRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListRequest) SetCode(v string) *CustomizeContactDeptListRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptListRequest) SetDeptId(v int64) *CustomizeContactDeptListRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptListResponseBody struct {
	// 子部门列表
	Content []*CustomizeContactDeptListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactDeptListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponseBody) SetContent(v []*CustomizeContactDeptListResponseBodyContent) *CustomizeContactDeptListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactDeptListResponseBodyContent struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 部门主管列表
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 上级部门Id
	ParentDeptId *int64 `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// 引用的内部通讯录部门Id
	RefId *int64 `json:"refId,omitempty" xml:"refId,omitempty"`
	// 部门类型 0:普通部门 1:引用部门
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponseBodyContent) SetCode(v string) *CustomizeContactDeptListResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Id = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetManagerIdList(v []*string) *CustomizeContactDeptListResponseBodyContent {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetName(v string) *CustomizeContactDeptListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetOrder(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetParentDeptId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetRefId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetType(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Type = &v
	return s
}

type CustomizeContactDeptListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeptListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptListResponse) SetBody(v *CustomizeContactDeptListResponseBody) *CustomizeContactDeptListResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptUpdateRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门主管列表
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 上级部门Id
	ParentDeptId *int64 `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
}

func (s CustomizeContactDeptUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateRequest) SetCode(v string) *CustomizeContactDeptUpdateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetDeptId(v int64) *CustomizeContactDeptUpdateRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetManagerIdList(v []*string) *CustomizeContactDeptUpdateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetName(v string) *CustomizeContactDeptUpdateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetOrder(v int64) *CustomizeContactDeptUpdateRequest {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetParentDeptId(v int64) *CustomizeContactDeptUpdateRequest {
	s.ParentDeptId = &v
	return s
}

type CustomizeContactDeptUpdateResponseBody struct {
	// 部门Id
	Content *int64 `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateResponseBody) SetContent(v int64) *CustomizeContactDeptUpdateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptUpdateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactDeptUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptUpdateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptUpdateResponse) SetBody(v *CustomizeContactDeptUpdateResponseBody) *CustomizeContactDeptUpdateResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpAddHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpAddRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 人员Id列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpAddRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddRequest) SetCode(v string) *CustomizeContactEmpAddRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactEmpAddRequest) SetDeptId(v int64) *CustomizeContactEmpAddRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactEmpAddRequest) SetUserIdList(v []*string) *CustomizeContactEmpAddRequest {
	s.UserIdList = v
	return s
}

type CustomizeContactEmpAddResponseBody struct {
	// 操作结果
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactEmpAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddResponseBody) SetContent(v bool) *CustomizeContactEmpAddResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactEmpAddResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactEmpAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpAddResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpAddResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpAddResponse) SetBody(v *CustomizeContactEmpAddResponseBody) *CustomizeContactEmpAddResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpDeleteRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 人员Id列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteRequest) SetCode(v string) *CustomizeContactEmpDeleteRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactEmpDeleteRequest) SetDeptId(v int64) *CustomizeContactEmpDeleteRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactEmpDeleteRequest) SetUserIdList(v []*string) *CustomizeContactEmpDeleteRequest {
	s.UserIdList = v
	return s
}

type CustomizeContactEmpDeleteResponseBody struct {
	// 操作结果
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactEmpDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteResponseBody) SetContent(v bool) *CustomizeContactEmpDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactEmpDeleteResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactEmpDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpDeleteResponse) SetBody(v *CustomizeContactEmpDeleteResponseBody) *CustomizeContactEmpDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpListRequest struct {
	// 部门Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactEmpListRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListRequest) SetDeptId(v int64) *CustomizeContactEmpListRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactEmpListResponseBody struct {
	// 人员信息列表
	Content []*CustomizeContactEmpListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponseBody) SetContent(v []*CustomizeContactEmpListResponseBodyContent) *CustomizeContactEmpListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactEmpListResponseBodyContent struct {
	// 人员姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CustomizeContactEmpListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponseBodyContent) SetName(v string) *CustomizeContactEmpListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactEmpListResponseBodyContent) SetUserId(v string) *CustomizeContactEmpListResponseBodyContent {
	s.UserId = &v
	return s
}

type CustomizeContactEmpListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactEmpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpListResponse) SetBody(v *CustomizeContactEmpListResponseBody) *CustomizeContactEmpListResponse {
	s.Body = v
	return s
}

type CustomizeContactListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactListResponseBody struct {
	// content
	Content []*CustomizeContactListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponseBody) SetContent(v []*CustomizeContactListResponseBodyContent) *CustomizeContactListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactListResponseBodyContent struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 自定义通讯录名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 自定义通讯录排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 跟部门Id
	RootDeptId *int64 `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s CustomizeContactListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponseBodyContent) SetCode(v string) *CustomizeContactListResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetName(v string) *CustomizeContactListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetOrder(v int64) *CustomizeContactListResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetRootDeptId(v int64) *CustomizeContactListResponseBodyContent {
	s.RootDeptId = &v
	return s
}

type CustomizeContactListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponse) SetHeaders(v map[string]*string) *CustomizeContactListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactListResponse) SetBody(v *CustomizeContactListResponseBody) *CustomizeContactListResponse {
	s.Body = v
	return s
}

type CustomizeContactUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactUpdateRequest struct {
	// 自定义通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 通讯录管理员UserId
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	// 自定义通讯录名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 在自定义通讯录列表中的排序
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s CustomizeContactUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateRequest) SetCode(v string) *CustomizeContactUpdateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactUpdateRequest) SetManagerIdList(v []*string) *CustomizeContactUpdateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactUpdateRequest) SetName(v string) *CustomizeContactUpdateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactUpdateRequest) SetOrder(v int64) *CustomizeContactUpdateRequest {
	s.Order = &v
	return s
}

type CustomizeContactUpdateResponseBody struct {
	// 是否操作成功
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateResponseBody) SetContent(v bool) *CustomizeContactUpdateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactUpdateResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CustomizeContactUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateResponse) SetHeaders(v map[string]*string) *CustomizeContactUpdateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactUpdateResponse) SetBody(v *CustomizeContactUpdateResponseBody) *CustomizeContactUpdateResponse {
	s.Body = v
	return s
}

type DigitalStoreContactInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreContactInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreContactInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreContactInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreContactInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreContactInfoResponseBody struct {
	// 门店通通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 门店通通讯录名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 门店通通讯录根节点Id
	RootDeptId *int64 `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s DigitalStoreContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoResponseBody) SetCode(v string) *DigitalStoreContactInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DigitalStoreContactInfoResponseBody) SetName(v string) *DigitalStoreContactInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreContactInfoResponseBody) SetRootDeptId(v int64) *DigitalStoreContactInfoResponseBody {
	s.RootDeptId = &v
	return s
}

type DigitalStoreContactInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreContactInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreContactInfoResponse) SetBody(v *DigitalStoreContactInfoResponseBody) *DigitalStoreContactInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreGroupInfoRequest struct {
	// 门店分组Id
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s DigitalStoreGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoRequest) SetGroupId(v int64) *DigitalStoreGroupInfoRequest {
	s.GroupId = &v
	return s
}

type DigitalStoreGroupInfoResponseBody struct {
	// 分组Id
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 分组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 分组中门店Id列表
	StoreIdList []*int64 `json:"storeIdList,omitempty" xml:"storeIdList,omitempty" type:"Repeated"`
}

func (s DigitalStoreGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoResponseBody) SetGroupId(v int64) *DigitalStoreGroupInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DigitalStoreGroupInfoResponseBody) SetGroupName(v string) *DigitalStoreGroupInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DigitalStoreGroupInfoResponseBody) SetStoreIdList(v []*int64) *DigitalStoreGroupInfoResponseBody {
	s.StoreIdList = v
	return s
}

type DigitalStoreGroupInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreGroupInfoResponse) SetBody(v *DigitalStoreGroupInfoResponseBody) *DigitalStoreGroupInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreGroupsResponseBody struct {
	Content []*DigitalStoreGroupsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponseBody) SetContent(v []*DigitalStoreGroupsResponseBodyContent) *DigitalStoreGroupsResponseBody {
	s.Content = v
	return s
}

type DigitalStoreGroupsResponseBodyContent struct {
	// 分组Id
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 分组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s DigitalStoreGroupsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponseBodyContent) SetGroupId(v int64) *DigitalStoreGroupsResponseBodyContent {
	s.GroupId = &v
	return s
}

func (s *DigitalStoreGroupsResponseBodyContent) SetGroupName(v string) *DigitalStoreGroupsResponseBodyContent {
	s.GroupName = &v
	return s
}

type DigitalStoreGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponse) SetHeaders(v map[string]*string) *DigitalStoreGroupsResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreGroupsResponse) SetBody(v *DigitalStoreGroupsResponseBody) *DigitalStoreGroupsResponse {
	s.Body = v
	return s
}

type DigitalStoreNodeInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreNodeInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreNodeInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreNodeInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreNodeInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreNodeInfoRequest struct {
	// 门店通通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 节点Id
	NodeId *int64 `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreNodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoRequest) SetCode(v string) *DigitalStoreNodeInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreNodeInfoRequest) SetNodeId(v int64) *DigitalStoreNodeInfoRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreNodeInfoResponseBody struct {
	// 节点Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 门店名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 上级节点id
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 节点类型
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DigitalStoreNodeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoResponseBody) SetId(v int64) *DigitalStoreNodeInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetName(v string) *DigitalStoreNodeInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetParentId(v int64) *DigitalStoreNodeInfoResponseBody {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetType(v int64) *DigitalStoreNodeInfoResponseBody {
	s.Type = &v
	return s
}

type DigitalStoreNodeInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreNodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreNodeInfoResponse) SetBody(v *DigitalStoreNodeInfoResponseBody) *DigitalStoreNodeInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreRolesHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreRolesResponseBody struct {
	Content []*DigitalStoreRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponseBody) SetContent(v []*DigitalStoreRolesResponseBodyContent) *DigitalStoreRolesResponseBody {
	s.Content = v
	return s
}

type DigitalStoreRolesResponseBodyContent struct {
	// 优先级
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 角色Code
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色Id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s DigitalStoreRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponseBodyContent) SetLevel(v int64) *DigitalStoreRolesResponseBodyContent {
	s.Level = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleCode(v string) *DigitalStoreRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleId(v int64) *DigitalStoreRolesResponseBodyContent {
	s.RoleId = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleName(v string) *DigitalStoreRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

type DigitalStoreRolesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponse) SetHeaders(v map[string]*string) *DigitalStoreRolesResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreRolesResponse) SetBody(v *DigitalStoreRolesResponseBody) *DigitalStoreRolesResponse {
	s.Body = v
	return s
}

type DigitalStoreStoreInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreStoreInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreStoreInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreStoreInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreStoreInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreStoreInfoRequest struct {
	// 门店通通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 门店Id
	StoreId *int64 `json:"storeId,omitempty" xml:"storeId,omitempty"`
}

func (s DigitalStoreStoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoRequest) SetCode(v string) *DigitalStoreStoreInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreStoreInfoRequest) SetStoreId(v int64) *DigitalStoreStoreInfoRequest {
	s.StoreId = &v
	return s
}

type DigitalStoreStoreInfoResponseBody struct {
	// 门店地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 营业时间
	BusinessHours *string `json:"businessHours,omitempty" xml:"businessHours,omitempty"`
	// 纬度
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 定位地址
	LocationAddress *string `json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	// 经度
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 门店名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 上级节点id
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 门店状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 门店面积
	StoreAcreage *string `json:"storeAcreage,omitempty" xml:"storeAcreage,omitempty"`
	// 门店带宽
	StoreBandwidth *string `json:"storeBandwidth,omitempty" xml:"storeBandwidth,omitempty"`
	// 门店编号
	StoreCode *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
	// 门店Id
	StoreId *int64 `json:"storeId,omitempty" xml:"storeId,omitempty"`
	// 门店电话
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s DigitalStoreStoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoResponseBody) SetAddress(v string) *DigitalStoreStoreInfoResponseBody {
	s.Address = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetBusinessHours(v string) *DigitalStoreStoreInfoResponseBody {
	s.BusinessHours = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLatitude(v string) *DigitalStoreStoreInfoResponseBody {
	s.Latitude = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLocationAddress(v string) *DigitalStoreStoreInfoResponseBody {
	s.LocationAddress = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLongitude(v string) *DigitalStoreStoreInfoResponseBody {
	s.Longitude = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetName(v string) *DigitalStoreStoreInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetParentId(v int64) *DigitalStoreStoreInfoResponseBody {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStatus(v string) *DigitalStoreStoreInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreAcreage(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreAcreage = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreBandwidth(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreBandwidth = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreCode(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreCode = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreId(v int64) *DigitalStoreStoreInfoResponseBody {
	s.StoreId = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetTelephone(v string) *DigitalStoreStoreInfoResponseBody {
	s.Telephone = &v
	return s
}

type DigitalStoreStoreInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreStoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreStoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreStoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreStoreInfoResponse) SetBody(v *DigitalStoreStoreInfoResponseBody) *DigitalStoreStoreInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreSubNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreSubNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreSubNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreSubNodesHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreSubNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreSubNodesRequest struct {
	// 门店通通讯录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 节点Id
	NodeId *int64 `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreSubNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesRequest) SetCode(v string) *DigitalStoreSubNodesRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreSubNodesRequest) SetNodeId(v int64) *DigitalStoreSubNodesRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreSubNodesResponseBody struct {
	Content []*DigitalStoreSubNodesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreSubNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponseBody) SetContent(v []*DigitalStoreSubNodesResponseBodyContent) *DigitalStoreSubNodesResponseBody {
	s.Content = v
	return s
}

type DigitalStoreSubNodesResponseBodyContent struct {
	// 节点Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 门店名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 上级节点id
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 节点类型
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DigitalStoreSubNodesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetId(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.Id = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetName(v string) *DigitalStoreSubNodesResponseBodyContent {
	s.Name = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetParentId(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetType(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.Type = &v
	return s
}

type DigitalStoreSubNodesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreSubNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreSubNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponse) SetHeaders(v map[string]*string) *DigitalStoreSubNodesResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreSubNodesResponse) SetBody(v *DigitalStoreSubNodesResponseBody) *DigitalStoreSubNodesResponse {
	s.Body = v
	return s
}

type DigitalStoreUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreUserInfoRequest struct {
	// 门店通迅录Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoRequest) SetCode(v string) *DigitalStoreUserInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreUserInfoRequest) SetUserId(v string) *DigitalStoreUserInfoRequest {
	s.UserId = &v
	return s
}

type DigitalStoreUserInfoResponseBody struct {
	// 人员名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 管理范围
	ScopeList []*int64 `json:"scopeList,omitempty" xml:"scopeList,omitempty" type:"Repeated"`
	// 所在节点列表
	StoreList []*int64 `json:"storeList,omitempty" xml:"storeList,omitempty" type:"Repeated"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoResponseBody) SetName(v string) *DigitalStoreUserInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetScopeList(v []*int64) *DigitalStoreUserInfoResponseBody {
	s.ScopeList = v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetStoreList(v []*int64) *DigitalStoreUserInfoResponseBody {
	s.StoreList = v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetUserId(v string) *DigitalStoreUserInfoResponseBody {
	s.UserId = &v
	return s
}

type DigitalStoreUserInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreUserInfoResponse) SetBody(v *DigitalStoreUserInfoResponseBody) *DigitalStoreUserInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreUsersHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreUsersRequest struct {
	// 门店通通讯录Cod
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 节点Id
	NodeId *int64 `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersRequest) SetCode(v string) *DigitalStoreUsersRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreUsersRequest) SetNodeId(v int64) *DigitalStoreUsersRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreUsersResponseBody struct {
	// 人员列表
	Content []*DigitalStoreUsersResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponseBody) SetContent(v []*DigitalStoreUsersResponseBodyContent) *DigitalStoreUsersResponseBody {
	s.Content = v
	return s
}

type DigitalStoreUsersResponseBodyContent struct {
	// 人员姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUsersResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponseBodyContent) SetName(v string) *DigitalStoreUsersResponseBodyContent {
	s.Name = &v
	return s
}

func (s *DigitalStoreUsersResponseBodyContent) SetUserId(v string) *DigitalStoreUsersResponseBodyContent {
	s.UserId = &v
	return s
}

type DigitalStoreUsersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DigitalStoreUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponse) SetHeaders(v map[string]*string) *DigitalStoreUsersResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreUsersResponse) SetBody(v *DigitalStoreUsersResponseBody) *DigitalStoreUsersResponse {
	s.Body = v
	return s
}

type IndustryManufactureCostRecordListGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureCostRecordListGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureCostRecordListGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureCostRecordListGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureCostRecordListGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureCostRecordListGetRequest struct {
	AppId            *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds           []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName          *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor           *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime          *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId       *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId         *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId  *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrderNo          *string  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrgId            *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber       *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	StartTime        *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey         *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType   *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureCostRecordListGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppIds(v []*int64) *IndustryManufactureCostRecordListGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppName(v string) *IndustryManufactureCostRecordListGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetCorpId(v string) *IndustryManufactureCostRecordListGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetCursor(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetEndTime(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetInstanceId(v string) *IndustryManufactureCostRecordListGetRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetIsvOrgId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetMaterialNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetMicroappAgentId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetOrderNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.OrderNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetOrgId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetPageNumber(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetPageSize(v int32) *IndustryManufactureCostRecordListGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetProductionTaskNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetStartTime(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetSuiteKey(v string) *IndustryManufactureCostRecordListGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetTokenGrantType(v int32) *IndustryManufactureCostRecordListGetRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureCostRecordListGetResponseBody struct {
	HasMore    *bool                                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureCostRecordListGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                                  `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureCostRecordListGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetHasMore(v bool) *IndustryManufactureCostRecordListGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetList(v []*IndustryManufactureCostRecordListGetResponseBodyList) *IndustryManufactureCostRecordListGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetNextCursor(v int64) *IndustryManufactureCostRecordListGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetTotalCount(v int64) *IndustryManufactureCostRecordListGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureCostRecordListGetResponseBodyList struct {
	CorpId               *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count                *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext                  *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate            *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified          *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId           *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsDeleted            *string  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	MaterialCostRecordNo *string  `json:"materialCostRecordNo,omitempty" xml:"materialCostRecordNo,omitempty"`
	MaterialName         *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo           *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	Memo                 *string  `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderNo              *string  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Price                *float32 `json:"price,omitempty" xml:"price,omitempty"`
	ProcessCode          *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProductionTaskNo     *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	RealCount            *float32 `json:"realCount,omitempty" xml:"realCount,omitempty"`
	RealPrice            *float32 `json:"realPrice,omitempty" xml:"realPrice,omitempty"`
	Type                 *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit                 *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureCostRecordListGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetCorpId(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetCount(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetExt(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetGmtModified(v int64) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetInstanceId(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetIsDeleted(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.IsDeleted = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialCostRecordNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialCostRecordNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialName(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMemo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Memo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetOrderNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.OrderNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetPrice(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Price = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetProcessCode(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetProductionTaskNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetRealCount(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.RealCount = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetRealPrice(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.RealPrice = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetType(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetUnit(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureCostRecordListGetResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustryManufactureCostRecordListGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureCostRecordListGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponse) SetHeaders(v map[string]*string) *IndustryManufactureCostRecordListGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponse) SetBody(v *IndustryManufactureCostRecordListGetResponseBody) *IndustryManufactureCostRecordListGetResponse {
	s.Body = v
	return s
}

type IndustryManufactureFeeListGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureFeeListGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureFeeListGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureFeeListGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureFeeListGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureFeeListGetRequest struct {
	AppId            *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds           []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName          *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor           *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime          *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsvOrgId         *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId  *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId            *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber       *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	StartTime        *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey         *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType   *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
	Type             *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IndustryManufactureFeeListGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetRequest) SetAppId(v int64) *IndustryManufactureFeeListGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetAppIds(v []*int64) *IndustryManufactureFeeListGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetAppName(v string) *IndustryManufactureFeeListGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetCorpId(v string) *IndustryManufactureFeeListGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetCursor(v int64) *IndustryManufactureFeeListGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetEndTime(v int64) *IndustryManufactureFeeListGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetIsvOrgId(v int64) *IndustryManufactureFeeListGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetMaterialNo(v string) *IndustryManufactureFeeListGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetMicroappAgentId(v int64) *IndustryManufactureFeeListGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetOrgId(v int64) *IndustryManufactureFeeListGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetPageNumber(v int64) *IndustryManufactureFeeListGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetPageSize(v int32) *IndustryManufactureFeeListGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetProductionTaskNo(v string) *IndustryManufactureFeeListGetRequest {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetStartTime(v int64) *IndustryManufactureFeeListGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetSuiteKey(v string) *IndustryManufactureFeeListGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetTokenGrantType(v int32) *IndustryManufactureFeeListGetRequest {
	s.TokenGrantType = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetType(v string) *IndustryManufactureFeeListGetRequest {
	s.Type = &v
	return s
}

type IndustryManufactureFeeListGetResponseBody struct {
	HasMore    *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureFeeListGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureFeeListGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponseBody) SetHasMore(v bool) *IndustryManufactureFeeListGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetList(v []*IndustryManufactureFeeListGetResponseBodyList) *IndustryManufactureFeeListGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetNextCursor(v int64) *IndustryManufactureFeeListGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetTotalCount(v int64) *IndustryManufactureFeeListGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureFeeListGetResponseBodyList struct {
	Amount           *string  `json:"amount,omitempty" xml:"amount,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count            *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext              *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate        *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id               *int64   `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId       *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsDeleted        *string  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	MaterialName     *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	PerAmount        *float32 `json:"perAmount,omitempty" xml:"perAmount,omitempty"`
	ProcessCode      *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	Title            *string  `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit             *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureFeeListGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetAmount(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Amount = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetCorpId(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetCount(v float32) *IndustryManufactureFeeListGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetExt(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetGmtModified(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetId(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.Id = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetInstanceId(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetIsDeleted(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.IsDeleted = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetMaterialName(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetMaterialNo(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetPerAmount(v float32) *IndustryManufactureFeeListGetResponseBodyList {
	s.PerAmount = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetProcessCode(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetProductionTaskNo(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetTitle(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Title = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetType(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetUnit(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureFeeListGetResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustryManufactureFeeListGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureFeeListGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponse) SetHeaders(v map[string]*string) *IndustryManufactureFeeListGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureFeeListGetResponse) SetBody(v *IndustryManufactureFeeListGetResponseBody) *IndustryManufactureFeeListGetResponse {
	s.Body = v
	return s
}

type IndustryManufactureLabourCostHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureLabourCostHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureLabourCostHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureLabourCostHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureLabourCostHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureLabourCostRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsvOrgId        *string  `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber      *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProcessNo       *string  `json:"processNo,omitempty" xml:"processNo,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureLabourCostRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostRequest) SetAppId(v int64) *IndustryManufactureLabourCostRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetAppIds(v []*int64) *IndustryManufactureLabourCostRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetAppName(v string) *IndustryManufactureLabourCostRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetCorpId(v string) *IndustryManufactureLabourCostRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetCursor(v int64) *IndustryManufactureLabourCostRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetEndTime(v int64) *IndustryManufactureLabourCostRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetIsvOrgId(v string) *IndustryManufactureLabourCostRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetMaterialNo(v string) *IndustryManufactureLabourCostRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetMicroappAgentId(v int64) *IndustryManufactureLabourCostRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetOrgId(v int64) *IndustryManufactureLabourCostRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetPageNumber(v int64) *IndustryManufactureLabourCostRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetPageSize(v int32) *IndustryManufactureLabourCostRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetProcessNo(v string) *IndustryManufactureLabourCostRequest {
	s.ProcessNo = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetStartTime(v int64) *IndustryManufactureLabourCostRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetSuiteKey(v string) *IndustryManufactureLabourCostRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetTokenGrantType(v int32) *IndustryManufactureLabourCostRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureLabourCostResponseBody struct {
	HasMore    *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureLabourCostResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureLabourCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponseBody) SetHasMore(v bool) *IndustryManufactureLabourCostResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetList(v []*IndustryManufactureLabourCostResponseBodyList) *IndustryManufactureLabourCostResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetNextCursor(v int64) *IndustryManufactureLabourCostResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetTotalCount(v int64) *IndustryManufactureLabourCostResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureLabourCostResponseBodyList struct {
	CorpId             *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Ext                *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate          *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId         *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LabourCostName     *string  `json:"labourCostName,omitempty" xml:"labourCostName,omitempty"`
	LabourCostNo       *string  `json:"labourCostNo,omitempty" xml:"labourCostNo,omitempty"`
	MaterialName       *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo         *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	ProcessCode        *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessName        *string  `json:"processName,omitempty" xml:"processName,omitempty"`
	ProcessNo          *string  `json:"processNo,omitempty" xml:"processNo,omitempty"`
	QualifiedPrice     *float32 `json:"qualifiedPrice,omitempty" xml:"qualifiedPrice,omitempty"`
	UnQualifiedInfo    *string  `json:"unQualifiedInfo,omitempty" xml:"unQualifiedInfo,omitempty"`
	UnQualifiedPrice1  *float32 `json:"unQualifiedPrice1,omitempty" xml:"unQualifiedPrice1,omitempty"`
	UnQualifiedPrice2  *float32 `json:"unQualifiedPrice2,omitempty" xml:"unQualifiedPrice2,omitempty"`
	UnQualifiedReason1 *string  `json:"unQualifiedReason1,omitempty" xml:"unQualifiedReason1,omitempty"`
	UnQualifiedReason2 *string  `json:"unQualifiedReason2,omitempty" xml:"unQualifiedReason2,omitempty"`
}

func (s IndustryManufactureLabourCostResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetCorpId(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetExt(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureLabourCostResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetGmtModified(v int64) *IndustryManufactureLabourCostResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetInstanceId(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetLabourCostName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.LabourCostName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetLabourCostNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.LabourCostNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetMaterialName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetMaterialNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessCode(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetQualifiedPrice(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.QualifiedPrice = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedInfo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedInfo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedPrice1(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedPrice1 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedPrice2(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedPrice2 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedReason1(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedReason1 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedReason2(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedReason2 = &v
	return s
}

type IndustryManufactureLabourCostResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustryManufactureLabourCostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureLabourCostResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponse) SetHeaders(v map[string]*string) *IndustryManufactureLabourCostResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureLabourCostResponse) SetBody(v *IndustryManufactureLabourCostResponseBody) *IndustryManufactureLabourCostResponse {
	s.Body = v
	return s
}

type IndustryManufactureMaterialListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMaterialListHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMaterialListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMaterialListHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMaterialListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMaterialListRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CurrentPage     *int64   `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId      *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId        *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int64   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureMaterialListRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListRequest) SetAppId(v int64) *IndustryManufactureMaterialListRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetAppIds(v []*int64) *IndustryManufactureMaterialListRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetAppName(v string) *IndustryManufactureMaterialListRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCorpId(v string) *IndustryManufactureMaterialListRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCurrentPage(v int64) *IndustryManufactureMaterialListRequest {
	s.CurrentPage = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCursor(v int64) *IndustryManufactureMaterialListRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetEndTime(v int64) *IndustryManufactureMaterialListRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetInstanceId(v string) *IndustryManufactureMaterialListRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetIsvOrgId(v int64) *IndustryManufactureMaterialListRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetMaterialNo(v string) *IndustryManufactureMaterialListRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetMicroappAgentId(v int64) *IndustryManufactureMaterialListRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetOrgId(v int64) *IndustryManufactureMaterialListRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetPageSize(v int32) *IndustryManufactureMaterialListRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetStartTime(v int64) *IndustryManufactureMaterialListRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetSuiteKey(v string) *IndustryManufactureMaterialListRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetTokenGrantType(v int64) *IndustryManufactureMaterialListRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureMaterialListResponseBody struct {
	HasMore    *bool                                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureMaterialListResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                             `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureMaterialListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponseBody) SetHasMore(v bool) *IndustryManufactureMaterialListResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetList(v []*IndustryManufactureMaterialListResponseBodyList) *IndustryManufactureMaterialListResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetNextCursor(v int64) *IndustryManufactureMaterialListResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetTotalCount(v int64) *IndustryManufactureMaterialListResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureMaterialListResponseBodyList struct {
	Category      *string  `json:"category,omitempty" xml:"category,omitempty"`
	CorpId        *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Ext           *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	InstanceId    *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaterialName  *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo    *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	ProcessCode   *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Specification *string  `json:"specification,omitempty" xml:"specification,omitempty"`
	StockMaxWarn  *float32 `json:"stockMaxWarn,omitempty" xml:"stockMaxWarn,omitempty"`
	StockMinWarn  *float32 `json:"stockMinWarn,omitempty" xml:"stockMinWarn,omitempty"`
	Type          *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit          *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureMaterialListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetCategory(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Category = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetCorpId(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetExt(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetInstanceId(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetMaterialName(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetMaterialNo(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetProcessCode(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetSpecification(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Specification = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetStockMaxWarn(v float32) *IndustryManufactureMaterialListResponseBodyList {
	s.StockMaxWarn = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetStockMinWarn(v float32) *IndustryManufactureMaterialListResponseBodyList {
	s.StockMinWarn = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetType(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetUnit(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureMaterialListResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustryManufactureMaterialListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMaterialListResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponse) SetHeaders(v map[string]*string) *IndustryManufactureMaterialListResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMaterialListResponse) SetBody(v *IndustryManufactureMaterialListResponseBody) *IndustryManufactureMaterialListResponse {
	s.Body = v
	return s
}

type IndustryMmanufactureMaterialCostGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryMmanufactureMaterialCostGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryMmanufactureMaterialCostGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryMmanufactureMaterialCostGetRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId      *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId        *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber      *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppIds(v []*int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppName(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetCorpId(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetCursor(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetEndTime(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetInstanceId(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetIsvOrgId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetMaterialNo(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetMicroappAgentId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetOrgId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetPageNumber(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetPageSize(v int32) *IndustryMmanufactureMaterialCostGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetStartTime(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetSuiteKey(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetTokenGrantType(v int32) *IndustryMmanufactureMaterialCostGetRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponseBody struct {
	HasMore    *bool                                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryMmanufactureMaterialCostGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                                 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetHasMore(v bool) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetList(v []*IndustryMmanufactureMaterialCostGetResponseBodyList) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetNextCursor(v int64) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetTotalCount(v int64) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponseBodyList struct {
	ActPrice       *float32 `json:"actPrice,omitempty" xml:"actPrice,omitempty"`
	CorpId         *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count          *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext            *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate      *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId     *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaterialCostNo *string  `json:"materialCostNo,omitempty" xml:"materialCostNo,omitempty"`
	MaterialName   *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo     *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	Memo           *string  `json:"memo,omitempty" xml:"memo,omitempty"`
	Price          *float32 `json:"price,omitempty" xml:"price,omitempty"`
	ProcessCode    *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Unit           *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetActPrice(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.ActPrice = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetCorpId(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetCount(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetExt(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetGmtCreate(v int64) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetGmtModified(v int64) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetInstanceId(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialCostNo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialCostNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialName(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialNo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMemo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Memo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetPrice(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Price = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetProcessCode(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetUnit(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustryMmanufactureMaterialCostGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryMmanufactureMaterialCostGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponse) SetHeaders(v map[string]*string) *IndustryMmanufactureMaterialCostGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponse) SetBody(v *IndustryMmanufactureMaterialCostGetResponseBody) *IndustryMmanufactureMaterialCostGetResponse {
	s.Body = v
	return s
}

type QueryAllDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDepartmentRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentRequest) SetPageNumber(v int32) *QueryAllDepartmentRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllDepartmentRequest) SetPageSize(v int32) *QueryAllDepartmentRequest {
	s.PageSize = &v
	return s
}

type QueryAllDepartmentResponseBody struct {
	// 科室列表
	Content []*QueryAllDepartmentResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBody) SetContent(v []*QueryAllDepartmentResponseBodyContent) *QueryAllDepartmentResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetCurrentPage(v int32) *QueryAllDepartmentResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalCount(v int64) *QueryAllDepartmentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalPages(v int32) *QueryAllDepartmentResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllDepartmentResponseBodyContent struct {
	// 科室详情
	DeptAndExt *QueryAllDepartmentResponseBodyContentDeptAndExt `json:"deptAndExt,omitempty" xml:"deptAndExt,omitempty" type:"Struct"`
	// 医疗组列表
	GroupAndExtList []*QueryAllDepartmentResponseBodyContentGroupAndExtList `json:"groupAndExtList,omitempty" xml:"groupAndExtList,omitempty" type:"Repeated"`
	// 科室ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContent) SetDeptAndExt(v *QueryAllDepartmentResponseBodyContentDeptAndExt) *QueryAllDepartmentResponseBodyContent {
	s.DeptAndExt = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetGroupAndExtList(v []*QueryAllDepartmentResponseBodyContentGroupAndExtList) *QueryAllDepartmentResponseBodyContent {
	s.GroupAndExtList = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetId(v int64) *QueryAllDepartmentResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetName(v string) *QueryAllDepartmentResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExt struct {
	// 科室详情
	Department *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	// 科室扩展信息列表
	ExtendInfos []*QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExt) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExt) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExt) SetDepartment(v *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) *QueryAllDepartmentResponseBodyContentDeptAndExt {
	s.Department = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExt) SetExtendInfos(v []*QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) *QueryAllDepartmentResponseBodyContentDeptAndExt {
	s.ExtendInfos = v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExtDepartment struct {
	// 租户CorpID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 部门code
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	// 科室名称，同name
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// 排序
	DeptOrder *int64 `json:"deptOrder,omitempty" xml:"deptOrder,omitempty"`
	// 部门状态：0-正常，1-删除
	DeptStatus *int32 `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	// 部门类型：1-科室，2-医疗组
	DeptType *int32 `json:"deptType,omitempty" xml:"deptType,omitempty"`
	// 创建时间
	GmtCreateStr *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	// 修改时间
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	// 科室ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称，同deptName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父部门code（与dept_code来源保持一致）
	ParentDeptCode *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 病区id列表
	WardIdList []*int64 `json:"wardIdList,omitempty" xml:"wardIdList,omitempty" type:"Repeated"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetCorpId(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.CorpId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptOrder(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptOrder = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptStatus(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptStatus = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptType(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptType = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetId(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetParentDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetRemark(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Remark = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetWardIdList(v []*int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.WardIdList = v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos struct {
	// 租户CorpID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 部门code
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	// 科室扩展字段描述
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	// 科室扩展字段key
	DeptExtendKey *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	// 科室扩展字段value
	DeptExtendValue *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	// 创建时间
	GmtCreateStr *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	// 修改时间
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	// 扩展信息id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 0-有效 ，1-无效
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetCorpId(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.CorpId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendDisplayName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendKey(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendValue(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetId(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetStatus(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.Status = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtList struct {
	// 医疗组扩展信息列表
	ExtendInfos []*QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
	// 医疗组详细信息
	Group *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtList) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtList) SetExtendInfos(v []*QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) *QueryAllDepartmentResponseBodyContentGroupAndExtList {
	s.ExtendInfos = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtList) SetGroup(v *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) *QueryAllDepartmentResponseBodyContentGroupAndExtList {
	s.Group = v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos struct {
	// 租户CorpID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 部门code
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	// 医疗组扩展字段描述
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	// 医疗组扩展字段key
	DeptExtendKey *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	// 医疗组扩展字段value
	DeptExtendValue *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	// 创建时间
	GmtCreateStr *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	// 修改时间
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	// 扩展信息id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 0-有效 ，1-无效
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetCorpId(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.CorpId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendDisplayName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendKey(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendValue(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetStatus(v int32) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.Status = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListGroup struct {
	// 租户CorpID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 科室ID，同parentDeptCode，这里保留是做兼容，原来定义成Long不太好改成了String了
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门状态：0-正常，1-删除
	DeptStatus *int32 `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	// 创建时间
	GmtCreateStr *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	// 修改时间
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	// 医疗组ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组组长信息
	Leader *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父级组织id，这里医疗组的父级就是科室
	ParentDeptCode *string `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetCorpId(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.CorpId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetDeptId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.DeptId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetDeptStatus(v int32) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.DeptStatus = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetLeader(v *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Leader = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetParentDeptCode(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetRemark(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Remark = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader struct {
	// 用户工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetJobNumber(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetUserId(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.UserId = &v
	return s
}

type QueryAllDepartmentResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponse) SetHeaders(v map[string]*string) *QueryAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDepartmentResponse) SetBody(v *QueryAllDepartmentResponseBody) *QueryAllDepartmentResponse {
	s.Body = v
	return s
}

type QueryAllDoctorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDoctorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDoctorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDoctorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDoctorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDoctorsRequest struct {
	// 分页查询页码
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllDoctorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsRequest) SetPageNum(v int32) *QueryAllDoctorsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAllDoctorsRequest) SetPageSize(v int32) *QueryAllDoctorsRequest {
	s.PageSize = &v
	return s
}

type QueryAllDoctorsResponseBody struct {
	// 人员列表
	Content []*QueryAllDoctorsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllDoctorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBody) SetContent(v []*QueryAllDoctorsResponseBodyContent) *QueryAllDoctorsResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetCurrentPage(v int32) *QueryAllDoctorsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalCount(v int64) *QueryAllDoctorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalPages(v int32) *QueryAllDoctorsResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllDoctorsResponseBodyContent struct {
	// 考核医疗组id
	AssessGroupId *string `json:"assessGroupId,omitempty" xml:"assessGroupId,omitempty"`
	// 考核医疗组名称
	AssessGroupName *string `json:"assessGroupName,omitempty" xml:"assessGroupName,omitempty"`
	// 租户CorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 关联的部门id
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	// 科室医疗组标识
	DeptType *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	// 用户创建时间
	GmtCreateStr *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	// 用户最后修改时间
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	// 用户id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	// 状态0-有效，1-删除
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 租户下staffId
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 租户内staffId
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllDoctorsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBodyContent) SetAssessGroupId(v string) *QueryAllDoctorsResponseBodyContent {
	s.AssessGroupId = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetAssessGroupName(v string) *QueryAllDoctorsResponseBodyContent {
	s.AssessGroupName = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetCorpId(v string) *QueryAllDoctorsResponseBodyContent {
	s.CorpId = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetDeptCode(v string) *QueryAllDoctorsResponseBodyContent {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetDeptType(v string) *QueryAllDoctorsResponseBodyContent {
	s.DeptType = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetGmtCreateStr(v string) *QueryAllDoctorsResponseBodyContent {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetGmtModifiedStr(v string) *QueryAllDoctorsResponseBodyContent {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetId(v int64) *QueryAllDoctorsResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetJobNum(v string) *QueryAllDoctorsResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetStatus(v int32) *QueryAllDoctorsResponseBodyContent {
	s.Status = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUid(v string) *QueryAllDoctorsResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUserCode(v string) *QueryAllDoctorsResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUserName(v string) *QueryAllDoctorsResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllDoctorsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllDoctorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDoctorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponse) SetHeaders(v map[string]*string) *QueryAllDoctorsResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDoctorsResponse) SetBody(v *QueryAllDoctorsResponseBody) *QueryAllDoctorsResponse {
	s.Body = v
	return s
}

type QueryAllGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupRequest) SetPageNumber(v int32) *QueryAllGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllGroupRequest) SetPageSize(v int32) *QueryAllGroupRequest {
	s.PageSize = &v
	return s
}

type QueryAllGroupResponseBody struct {
	// Id of the request
	Content []*QueryAllGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBody) SetContent(v []*QueryAllGroupResponseBodyContent) *QueryAllGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupResponseBody) SetCurrentPage(v int64) *QueryAllGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalCount(v int64) *QueryAllGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalPages(v int64) *QueryAllGroupResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllGroupResponseBodyContent struct {
	// 所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBodyContent) SetDeptId(v int64) *QueryAllGroupResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetId(v int64) *QueryAllGroupResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetName(v string) *QueryAllGroupResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponse) SetHeaders(v map[string]*string) *QueryAllGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupResponse) SetBody(v *QueryAllGroupResponseBody) *QueryAllGroupResponse {
	s.Body = v
	return s
}

type QueryAllGroupsInDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupsInDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupsInDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupsInDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupsInDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupsInDeptRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllGroupsInDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptRequest) SetPageNumber(v int32) *QueryAllGroupsInDeptRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllGroupsInDeptRequest) SetPageSize(v int32) *QueryAllGroupsInDeptRequest {
	s.PageSize = &v
	return s
}

type QueryAllGroupsInDeptResponseBody struct {
	// Id of the request
	Content []*QueryAllGroupsInDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBody) SetContent(v []*QueryAllGroupsInDeptResponseBodyContent) *QueryAllGroupsInDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetCurrentPage(v int64) *QueryAllGroupsInDeptResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalCount(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalPages(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllGroupsInDeptResponseBodyContent struct {
	// 所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetDeptId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetName(v string) *QueryAllGroupsInDeptResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllGroupsInDeptResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllGroupsInDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupsInDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponse) SetHeaders(v map[string]*string) *QueryAllGroupsInDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupsInDeptResponse) SetBody(v *QueryAllGroupsInDeptResponseBody) *QueryAllGroupsInDeptResponse {
	s.Body = v
	return s
}

type QueryAllMemberByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByDeptRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询页容量
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllMemberByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptRequest) SetPageNumber(v int32) *QueryAllMemberByDeptRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllMemberByDeptRequest) SetPageSize(v int32) *QueryAllMemberByDeptRequest {
	s.PageSize = &v
	return s
}

type QueryAllMemberByDeptResponseBody struct {
	// 人员列表
	Content []*QueryAllMemberByDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllMemberByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBody) SetContent(v []*QueryAllMemberByDeptResponseBodyContent) *QueryAllMemberByDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetCurrentPage(v int32) *QueryAllMemberByDeptResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalCount(v int64) *QueryAllMemberByDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalPages(v int32) *QueryAllMemberByDeptResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllMemberByDeptResponseBodyContent struct {
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllMemberByDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetJobNum(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUid(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUserName(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllMemberByDeptResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllMemberByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponse) SetHeaders(v map[string]*string) *QueryAllMemberByDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByDeptResponse) SetBody(v *QueryAllMemberByDeptResponseBody) *QueryAllMemberByDeptResponse {
	s.Body = v
	return s
}

type QueryAllMemberByGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByGroupRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllMemberByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupRequest) SetPageNumber(v int32) *QueryAllMemberByGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllMemberByGroupRequest) SetPageSize(v int32) *QueryAllMemberByGroupRequest {
	s.PageSize = &v
	return s
}

type QueryAllMemberByGroupResponseBody struct {
	// 人员列表
	Content []*QueryAllMemberByGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllMemberByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBody) SetContent(v []*QueryAllMemberByGroupResponseBodyContent) *QueryAllMemberByGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetCurrentPage(v int32) *QueryAllMemberByGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalCount(v int64) *QueryAllMemberByGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalPages(v int32) *QueryAllMemberByGroupResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllMemberByGroupResponseBodyContent struct {
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllMemberByGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetJobNum(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUid(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUserName(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllMemberByGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllMemberByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponse) SetHeaders(v map[string]*string) *QueryAllMemberByGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByGroupResponse) SetBody(v *QueryAllMemberByGroupResponseBody) *QueryAllMemberByGroupResponse {
	s.Body = v
	return s
}

type QueryBizOptLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBizOptLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogHeaders) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogHeaders) SetCommonHeaders(v map[string]*string) *QueryBizOptLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBizOptLogHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBizOptLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBizOptLogRequest struct {
	// 每次拉取的数据量，最大200条
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 拉取记录的起始位置，默认从上次拉取的最后位置开始
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryBizOptLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogRequest) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogRequest) SetMaxResults(v int32) *QueryBizOptLogRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryBizOptLogRequest) SetNextToken(v int64) *QueryBizOptLogRequest {
	s.NextToken = &v
	return s
}

type QueryBizOptLogResponseBody struct {
	// content
	Content []*QueryBizOptLogResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 下次拉取数据的起始位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryBizOptLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBody) SetContent(v []*QueryBizOptLogResponseBodyContent) *QueryBizOptLogResponseBody {
	s.Content = v
	return s
}

func (s *QueryBizOptLogResponseBody) SetNextToken(v int64) *QueryBizOptLogResponseBody {
	s.NextToken = &v
	return s
}

type QueryBizOptLogResponseBodyContent struct {
	// 业务类型
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 数据类型
	DataType *int32 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// 日志ID
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 操作后对象数据快照，json格式
	OptAfterData *string `json:"optAfterData,omitempty" xml:"optAfterData,omitempty"`
	// 操作前对象数据快照，json格式
	OptBeforeData *string `json:"optBeforeData,omitempty" xml:"optBeforeData,omitempty"`
	// 操作业务类型
	OptBizType *int32 `json:"optBizType,omitempty" xml:"optBizType,omitempty"`
	// 扩展信息，map json格式
	OptExtend *string `json:"optExtend,omitempty" xml:"optExtend,omitempty"`
	// 操作者工号
	OptJobNumber *string `json:"optJobNumber,omitempty" xml:"optJobNumber,omitempty"`
	// 操作对象code，人员code，或者部门code
	OptObjectCode *string `json:"optObjectCode,omitempty" xml:"optObjectCode,omitempty"`
	// 操作对象名称
	OptObjectName *string `json:"optObjectName,omitempty" xml:"optObjectName,omitempty"`
	// 操作对象人员工号
	OptObjectUserJobNo *string `json:"optObjectUserJobNo,omitempty" xml:"optObjectUserJobNo,omitempty"`
	// 操作是否成功
	OptSuccess *int32 `json:"optSuccess,omitempty" xml:"optSuccess,omitempty"`
	// 操作时间 时间戳
	OptTime *int64 `json:"optTime,omitempty" xml:"optTime,omitempty"`
	// 操作类型
	OptType *int32 `json:"optType,omitempty" xml:"optType,omitempty"`
	// 操作用户code
	OptUserCode *string `json:"optUserCode,omitempty" xml:"optUserCode,omitempty"`
	// 操作用户名称
	OptUserName *string `json:"optUserName,omitempty" xml:"optUserName,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s QueryBizOptLogResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBodyContent) SetBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.BizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetDataType(v int32) *QueryBizOptLogResponseBodyContent {
	s.DataType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetId(v int64) *QueryBizOptLogResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptAfterData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptAfterData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBeforeData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptBeforeData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptBizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptExtend(v string) *QueryBizOptLogResponseBodyContent {
	s.OptExtend = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptJobNumber(v string) *QueryBizOptLogResponseBodyContent {
	s.OptJobNumber = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectUserJobNo(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectUserJobNo = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptSuccess(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptSuccess = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptTime(v int64) *QueryBizOptLogResponseBodyContent {
	s.OptTime = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetRemark(v string) *QueryBizOptLogResponseBodyContent {
	s.Remark = &v
	return s
}

type QueryBizOptLogResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBizOptLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizOptLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponse) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponse) SetHeaders(v map[string]*string) *QueryBizOptLogResponse {
	s.Headers = v
	return s
}

func (s *QueryBizOptLogResponse) SetBody(v *QueryBizOptLogResponseBody) *QueryBizOptLogResponse {
	s.Body = v
	return s
}

type QueryDepartmentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDepartmentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDepartmentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDepartmentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDepartmentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDepartmentInfoResponseBody struct {
	// 科室详情
	Content *QueryDepartmentInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBody) SetContent(v *QueryDepartmentInfoResponseBodyContent) *QueryDepartmentInfoResponseBody {
	s.Content = v
	return s
}

type QueryDepartmentInfoResponseBodyContent struct {
	// 科室Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室主任
	Leader *QueryDepartmentInfoResponseBodyContentLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 住院总医师
	ResidentLeader *QueryDepartmentInfoResponseBodyContentResidentLeader `json:"residentLeader,omitempty" xml:"residentLeader,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContent) SetId(v int64) *QueryDepartmentInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetLeader(v *QueryDepartmentInfoResponseBodyContentLeader) *QueryDepartmentInfoResponseBodyContent {
	s.Leader = v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetName(v string) *QueryDepartmentInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetResidentLeader(v *QueryDepartmentInfoResponseBodyContentResidentLeader) *QueryDepartmentInfoResponseBodyContent {
	s.ResidentLeader = v
	return s
}

type QueryDepartmentInfoResponseBodyContentLeader struct {
	// 工作标签
	Job *QueryDepartmentInfoResponseBodyContentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentLeader) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetJob(v *QueryDepartmentInfoResponseBodyContentLeaderJob) *QueryDepartmentInfoResponseBodyContentLeader {
	s.Job = v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetJobNumber(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetName(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeader) SetUserId(v string) *QueryDepartmentInfoResponseBodyContentLeader {
	s.UserId = &v
	return s
}

type QueryDepartmentInfoResponseBodyContentLeaderJob struct {
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetBizType(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetCategory(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetCode(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentLeaderJob) SetDisplayName(v string) *QueryDepartmentInfoResponseBodyContentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryDepartmentInfoResponseBodyContentResidentLeader struct {
	// 工作标签
	Job *QueryDepartmentInfoResponseBodyContentResidentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeader) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetJob(v *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.Job = v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetJobNumber(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetName(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeader) SetUserId(v string) *QueryDepartmentInfoResponseBodyContentResidentLeader {
	s.UserId = &v
	return s
}

type QueryDepartmentInfoResponseBodyContentResidentLeaderJob struct {
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentResidentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetBizType(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetCategory(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetCode(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentResidentLeaderJob) SetDisplayName(v string) *QueryDepartmentInfoResponseBodyContentResidentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryDepartmentInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDepartmentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDepartmentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponse) SetHeaders(v map[string]*string) *QueryDepartmentInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDepartmentInfoResponse) SetBody(v *QueryDepartmentInfoResponseBody) *QueryDepartmentInfoResponse {
	s.Body = v
	return s
}

type QueryGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupInfoResponseBody struct {
	// 医疗组详情
	Content *QueryGroupInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBody) SetContent(v *QueryGroupInfoResponseBodyContent) *QueryGroupInfoResponseBody {
	s.Content = v
	return s
}

type QueryGroupInfoResponseBodyContent struct {
	// 科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 有效期结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组组长
	Leader *QueryGroupInfoResponseBodyContentLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 有效期开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryGroupInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContent) SetDeptId(v int64) *QueryGroupInfoResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetEndTime(v int64) *QueryGroupInfoResponseBodyContent {
	s.EndTime = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetId(v int64) *QueryGroupInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetLeader(v *QueryGroupInfoResponseBodyContentLeader) *QueryGroupInfoResponseBodyContent {
	s.Leader = v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetName(v string) *QueryGroupInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetStartTime(v int64) *QueryGroupInfoResponseBodyContent {
	s.StartTime = &v
	return s
}

type QueryGroupInfoResponseBodyContentLeader struct {
	// 工作标签
	Job *QueryGroupInfoResponseBodyContentLeaderJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// 工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 人员Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentLeader) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetJob(v *QueryGroupInfoResponseBodyContentLeaderJob) *QueryGroupInfoResponseBodyContentLeader {
	s.Job = v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetJobNumber(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetName(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeader) SetUserId(v string) *QueryGroupInfoResponseBodyContentLeader {
	s.UserId = &v
	return s
}

type QueryGroupInfoResponseBodyContentLeaderJob struct {
	// 业务类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentLeaderJob) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentLeaderJob) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetBizType(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.BizType = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetCategory(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.Category = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetCode(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.Code = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentLeaderJob) SetDisplayName(v string) *QueryGroupInfoResponseBodyContentLeaderJob {
	s.DisplayName = &v
	return s
}

type QueryGroupInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponse) SetHeaders(v map[string]*string) *QueryGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoResponse) SetBody(v *QueryGroupInfoResponseBody) *QueryGroupInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalDistrictInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalDistrictInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalDistrictInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalDistrictInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalDistrictInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalDistrictInfoRequest struct {
	// 分页查询页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryHospitalDistrictInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoRequest) SetPageNumber(v int32) *QueryHospitalDistrictInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryHospitalDistrictInfoRequest) SetPageSize(v int32) *QueryHospitalDistrictInfoRequest {
	s.PageSize = &v
	return s
}

type QueryHospitalDistrictInfoResponseBody struct {
	// 院区病区详情
	Content []*QueryHospitalDistrictInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryHospitalDistrictInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponseBody) SetContent(v []*QueryHospitalDistrictInfoResponseBodyContent) *QueryHospitalDistrictInfoResponseBody {
	s.Content = v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetCurrentPage(v int64) *QueryHospitalDistrictInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetTotalCount(v int64) *QueryHospitalDistrictInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetTotalPages(v int32) *QueryHospitalDistrictInfoResponseBody {
	s.TotalPages = &v
	return s
}

type QueryHospitalDistrictInfoResponseBodyContent struct {
	// 病区对应的物理地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 删除，0:正常，其他：已删除
	Deleted *int32 `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 院区或病区名称
	DistrictName *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// 类型，1：院区；2：病区
	DistrictType *int32 `json:"districtType,omitempty" xml:"districtType,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 主键
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 院区id
	ParentDistrictId *int64 `json:"parentDistrictId,omitempty" xml:"parentDistrictId,omitempty"`
}

func (s QueryHospitalDistrictInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetAddress(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Address = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDeleted(v int32) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Deleted = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDistrictName(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.DistrictName = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDistrictType(v int32) *QueryHospitalDistrictInfoResponseBodyContent {
	s.DistrictType = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetGmtCreate(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetGmtModified(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetId(v int64) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetParentDistrictId(v int64) *QueryHospitalDistrictInfoResponseBodyContent {
	s.ParentDistrictId = &v
	return s
}

type QueryHospitalDistrictInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHospitalDistrictInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalDistrictInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponse) SetHeaders(v map[string]*string) *QueryHospitalDistrictInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalDistrictInfoResponse) SetBody(v *QueryHospitalDistrictInfoResponseBody) *QueryHospitalDistrictInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalRoleUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalRoleUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalRoleUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalRoleUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalRoleUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalRoleUserInfoRequest struct {
	// 分页查询页码
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页查询分页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryHospitalRoleUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoRequest) SetPageNumber(v int64) *QueryHospitalRoleUserInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryHospitalRoleUserInfoRequest) SetPageSize(v int64) *QueryHospitalRoleUserInfoRequest {
	s.PageSize = &v
	return s
}

type QueryHospitalRoleUserInfoResponseBody struct {
	// 角色人员信息
	Content []*QueryHospitalRoleUserInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 总页数
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryHospitalRoleUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetContent(v []*QueryHospitalRoleUserInfoResponseBodyContent) *QueryHospitalRoleUserInfoResponseBody {
	s.Content = v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetCurrentPage(v int32) *QueryHospitalRoleUserInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetTotalCount(v int64) *QueryHospitalRoleUserInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetTotalPages(v int32) *QueryHospitalRoleUserInfoResponseBody {
	s.TotalPages = &v
	return s
}

type QueryHospitalRoleUserInfoResponseBodyContent struct {
	// gmtCreate
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 用户工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 角色编码
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 用户编码
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryHospitalRoleUserInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetGmtCreate(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetGmtModified(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetJobNumber(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.JobNumber = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetRoleCode(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetRoleName(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.RoleName = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetUserCode(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetUserName(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryHospitalRoleUserInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHospitalRoleUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalRoleUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponse) SetHeaders(v map[string]*string) *QueryHospitalRoleUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalRoleUserInfoResponse) SetBody(v *QueryHospitalRoleUserInfoResponseBody) *QueryHospitalRoleUserInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalRolesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalRolesResponseBody struct {
	// 角色列表
	Content []*QueryHospitalRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryHospitalRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponseBody) SetContent(v []*QueryHospitalRolesResponseBodyContent) *QueryHospitalRolesResponseBody {
	s.Content = v
	return s
}

type QueryHospitalRolesResponseBodyContent struct {
	// 修改时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 主键
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否已删除，默认0未删除，1已删除
	IsDeleted *int64 `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 角色关联权限点是否只读
	ReadOnly *int64 `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 角色编码
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 排序，数字越小越靠前，默认0
	Sort *int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s QueryHospitalRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponseBodyContent) SetGmtCreate(v string) *QueryHospitalRolesResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetId(v int64) *QueryHospitalRolesResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetIsDeleted(v int64) *QueryHospitalRolesResponseBodyContent {
	s.IsDeleted = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetReadOnly(v int64) *QueryHospitalRolesResponseBodyContent {
	s.ReadOnly = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRemark(v string) *QueryHospitalRolesResponseBodyContent {
	s.Remark = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRoleCode(v string) *QueryHospitalRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRoleName(v string) *QueryHospitalRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetSort(v int64) *QueryHospitalRolesResponseBodyContent {
	s.Sort = &v
	return s
}

type QueryHospitalRolesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHospitalRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponse) SetHeaders(v map[string]*string) *QueryHospitalRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalRolesResponse) SetBody(v *QueryHospitalRolesResponseBody) *QueryHospitalRolesResponse {
	s.Body = v
	return s
}

type QueryJobCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryJobCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobCodeDictionaryResponseBody struct {
	// code列表
	Content []*QueryJobCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryJobCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponseBody) SetContent(v []*QueryJobCodeDictionaryResponseBodyContent) *QueryJobCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryJobCodeDictionaryResponseBodyContent struct {
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 固定字段标识
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名字
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 1:医师,0:非医师,2:待补充
	DoctorType *string `json:"doctorType,omitempty" xml:"doctorType,omitempty"`
}

func (s QueryJobCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetCode(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetDoctorType(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.DoctorType = &v
	return s
}

type QueryJobCodeDictionaryResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryJobCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryJobCodeDictionaryResponse) SetBody(v *QueryJobCodeDictionaryResponseBody) *QueryJobCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryJobStatusCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobStatusCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryJobStatusCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobStatusCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobStatusCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobStatusCodeDictionaryResponseBody struct {
	// code列表
	Content []*QueryJobStatusCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryJobStatusCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponseBody) SetContent(v []*QueryJobStatusCodeDictionaryResponseBodyContent) *QueryJobStatusCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryJobStatusCodeDictionaryResponseBodyContent struct {
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 固定字段标识
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名字
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryJobStatusCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetCode(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

type QueryJobStatusCodeDictionaryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryJobStatusCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobStatusCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryJobStatusCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponse) SetBody(v *QueryJobStatusCodeDictionaryResponseBody) *QueryJobStatusCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryMedicalEventsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMedicalEventsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsHeaders) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsHeaders) SetCommonHeaders(v map[string]*string) *QueryMedicalEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMedicalEventsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMedicalEventsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMedicalEventsResponseBody struct {
	// 事件详情列表
	Content []*QueryMedicalEventsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryMedicalEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponseBody) SetContent(v []*QueryMedicalEventsResponseBodyContent) *QueryMedicalEventsResponseBody {
	s.Content = v
	return s
}

func (s *QueryMedicalEventsResponseBody) SetSuccess(v bool) *QueryMedicalEventsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMedicalEventsResponseBody) SetTotalCount(v int64) *QueryMedicalEventsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryMedicalEventsResponseBodyContent struct {
	// 事件code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 事件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 事件id
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
}

func (s QueryMedicalEventsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponseBodyContent) SetCode(v string) *QueryMedicalEventsResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryMedicalEventsResponseBodyContent) SetContent(v string) *QueryMedicalEventsResponseBodyContent {
	s.Content = &v
	return s
}

func (s *QueryMedicalEventsResponseBodyContent) SetEventId(v int64) *QueryMedicalEventsResponseBodyContent {
	s.EventId = &v
	return s
}

type QueryMedicalEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMedicalEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMedicalEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponse) SetHeaders(v map[string]*string) *QueryMedicalEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryMedicalEventsResponse) SetBody(v *QueryMedicalEventsResponseBody) *QueryMedicalEventsResponse {
	s.Body = v
	return s
}

type QueryUserCredentialsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserCredentialsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserCredentialsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserCredentialsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserCredentialsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserCredentialsRequest struct {
	// userId列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryUserCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsRequest) SetUserIds(v []*string) *QueryUserCredentialsRequest {
	s.UserIds = v
	return s
}

type QueryUserCredentialsResponseBody struct {
	// 人员证书
	Content []*QueryUserCredentialsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBody) SetContent(v []*QueryUserCredentialsResponseBodyContent) *QueryUserCredentialsResponseBody {
	s.Content = v
	return s
}

type QueryUserCredentialsResponseBodyContent struct {
	// 证书
	CredentialList []*QueryUserCredentialsResponseBodyContentCredentialList `json:"credentialList,omitempty" xml:"credentialList,omitempty" type:"Repeated"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserCredentialsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBodyContent) SetCredentialList(v []*QueryUserCredentialsResponseBodyContentCredentialList) *QueryUserCredentialsResponseBodyContent {
	s.CredentialList = v
	return s
}

func (s *QueryUserCredentialsResponseBodyContent) SetUserId(v string) *QueryUserCredentialsResponseBodyContent {
	s.UserId = &v
	return s
}

type QueryUserCredentialsResponseBodyContentCredentialList struct {
	// 证书名称
	CredentialName *int32 `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// 证书类型
	CredentialType *int32 `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// 有效日期
	TermOfValidity *string `json:"termOfValidity,omitempty" xml:"termOfValidity,omitempty"`
}

func (s QueryUserCredentialsResponseBodyContentCredentialList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBodyContentCredentialList) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetCredentialName(v int32) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.CredentialName = &v
	return s
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetCredentialType(v int32) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.CredentialType = &v
	return s
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetTermOfValidity(v string) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.TermOfValidity = &v
	return s
}

type QueryUserCredentialsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponse) SetHeaders(v map[string]*string) *QueryUserCredentialsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserCredentialsResponse) SetBody(v *QueryUserCredentialsResponseBody) *QueryUserCredentialsResponse {
	s.Body = v
	return s
}

type QueryUserExtInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserExtInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserExtInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserExtInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserExtInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserExtInfoResponseBody struct {
	// 扩展属性
	Content []*QueryUserExtInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserExtInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBody) SetContent(v []*QueryUserExtInfoResponseBodyContent) *QueryUserExtInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserExtInfoResponseBodyContent struct {
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 组织id
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// 状态：0-有效，1-无效
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 用户标识
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// 扩展属性描述
	UserExtendDisplayName *string `json:"userExtendDisplayName,omitempty" xml:"userExtendDisplayName,omitempty"`
	// 扩展属性Key
	UserExtendKey *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	// 扩展属性值
	UserExtendValue *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s QueryUserExtInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBodyContent) SetGmtCreate(v string) *QueryUserExtInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetGmtModified(v string) *QueryUserExtInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetOrgId(v string) *QueryUserExtInfoResponseBodyContent {
	s.OrgId = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetStatus(v int32) *QueryUserExtInfoResponseBodyContent {
	s.Status = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserCode(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendDisplayName(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendDisplayName = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendKey(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendValue(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendValue = &v
	return s
}

type QueryUserExtInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserExtInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponse) SetHeaders(v map[string]*string) *QueryUserExtInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserExtInfoResponse) SetBody(v *QueryUserExtInfoResponseBody) *QueryUserExtInfoResponse {
	s.Body = v
	return s
}

type QueryUserExtendValuesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserExtendValuesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserExtendValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserExtendValuesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserExtendValuesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserExtendValuesRequest struct {
	// 用户拓展key
	UserExtendKey *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	// userId列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryUserExtendValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesRequest) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesRequest) SetUserExtendKey(v string) *QueryUserExtendValuesRequest {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtendValuesRequest) SetUserIds(v []*string) *QueryUserExtendValuesRequest {
	s.UserIds = v
	return s
}

type QueryUserExtendValuesResponseBody struct {
	// 人员列表
	Content []*QueryUserExtendValuesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryUserExtendValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponseBody) SetContent(v []*QueryUserExtendValuesResponseBodyContent) *QueryUserExtendValuesResponseBody {
	s.Content = v
	return s
}

func (s *QueryUserExtendValuesResponseBody) SetSuccess(v bool) *QueryUserExtendValuesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserExtendValuesResponseBody) SetTotalCount(v int64) *QueryUserExtendValuesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryUserExtendValuesResponseBodyContent struct {
	// 用户code
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// 扩展字段描述
	UserExtendDisplayName *string `json:"userExtendDisplayName,omitempty" xml:"userExtendDisplayName,omitempty"`
	// 扩展字段key
	UserExtendKey *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	// 扩展字段value
	UserExtendValue *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s QueryUserExtendValuesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserCode(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendDisplayName(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendDisplayName = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendKey(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendValue(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendValue = &v
	return s
}

type QueryUserExtendValuesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserExtendValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserExtendValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponse) SetHeaders(v map[string]*string) *QueryUserExtendValuesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserExtendValuesResponse) SetBody(v *QueryUserExtendValuesResponseBody) *QueryUserExtendValuesResponse {
	s.Body = v
	return s
}

type QueryUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserInfoResponseBody struct {
	// 人员详情
	Content *QueryUserInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBody) SetContent(v *QueryUserInfoResponseBodyContent) *QueryUserInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserInfoResponseBodyContent struct {
	// comments
	Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
	// 所在科室
	Dept []*QueryUserInfoResponseBodyContentDept `json:"dept,omitempty" xml:"dept,omitempty" type:"Repeated"`
	// 所在医疗组
	Group []*QueryUserInfoResponseBodyContentGroup `json:"group,omitempty" xml:"group,omitempty" type:"Repeated"`
	// 职称标签
	Job *QueryUserInfoResponseBodyContentJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// 工号
	JobNum *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	// 工作状态标签, 已废弃, 请使用jobStatusList字段
	JobStatus *QueryUserInfoResponseBodyContentJobStatus `json:"jobStatus,omitempty" xml:"jobStatus,omitempty" type:"Struct"`
	// 工作状态标签
	JobStatusList []*QueryUserInfoResponseBodyContentJobStatusList `json:"jobStatusList,omitempty" xml:"jobStatusList,omitempty" type:"Repeated"`
	// 用户Id
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 用户名称
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 人员属性标签
	UserProb *QueryUserInfoResponseBodyContentUserProb `json:"userProb,omitempty" xml:"userProb,omitempty" type:"Struct"`
}

func (s QueryUserInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContent) SetComments(v string) *QueryUserInfoResponseBodyContent {
	s.Comments = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetDept(v []*QueryUserInfoResponseBodyContentDept) *QueryUserInfoResponseBodyContent {
	s.Dept = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetGroup(v []*QueryUserInfoResponseBodyContentGroup) *QueryUserInfoResponseBodyContent {
	s.Group = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJob(v *QueryUserInfoResponseBodyContentJob) *QueryUserInfoResponseBodyContent {
	s.Job = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobNum(v string) *QueryUserInfoResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobStatus(v *QueryUserInfoResponseBodyContentJobStatus) *QueryUserInfoResponseBodyContent {
	s.JobStatus = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobStatusList(v []*QueryUserInfoResponseBodyContentJobStatusList) *QueryUserInfoResponseBodyContent {
	s.JobStatusList = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUid(v string) *QueryUserInfoResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserName(v string) *QueryUserInfoResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserProb(v *QueryUserInfoResponseBodyContentUserProb) *QueryUserInfoResponseBodyContent {
	s.UserProb = v
	return s
}

type QueryUserInfoResponseBodyContentDept struct {
	// 科室Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 科室名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryUserInfoResponseBodyContentDept) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentDept) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentDept) SetId(v int64) *QueryUserInfoResponseBodyContentDept {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetName(v string) *QueryUserInfoResponseBodyContentDept {
	s.Name = &v
	return s
}

type QueryUserInfoResponseBodyContentGroup struct {
	// 医疗组所在科室Id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 医疗组所在科室名称
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// 医疗组Id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 医疗组名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryUserInfoResponseBodyContentGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentGroup) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.DeptId = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.DeptName = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.Name = &v
	return s
}

type QueryUserInfoResponseBodyContentJob struct {
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJob) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJob) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJob) SetBizType(v string) *QueryUserInfoResponseBodyContentJob {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetCategory(v string) *QueryUserInfoResponseBodyContentJob {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetCode(v string) *QueryUserInfoResponseBodyContentJob {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJob {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentJobStatus struct {
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJobStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJobStatus) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetBizType(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCategory(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCode(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentJobStatusList struct {
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJobStatusList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJobStatusList) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetBizType(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetCategory(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetCode(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentUserProb struct {
	// 标签类型
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 标签Code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentUserProb) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentUserProb) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetBizType(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCategory(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCode(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetDisplayName(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponse) SetHeaders(v map[string]*string) *QueryUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoResponse) SetBody(v *QueryUserInfoResponseBody) *QueryUserInfoResponse {
	s.Body = v
	return s
}

type QueryUserProbCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserProbCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryUserProbCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserProbCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserProbCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserProbCodeDictionaryResponseBody struct {
	// code列表
	Content []*QueryUserProbCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserProbCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponseBody) SetContent(v []*QueryUserProbCodeDictionaryResponseBodyContent) *QueryUserProbCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryUserProbCodeDictionaryResponseBodyContent struct {
	// 分类
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 固定字段标识
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示名字
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserProbCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetCode(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

type QueryUserProbCodeDictionaryResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserProbCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserProbCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryUserProbCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryUserProbCodeDictionaryResponse) SetBody(v *QueryUserProbCodeDictionaryResponseBody) *QueryUserProbCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryUserRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserRolesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserRolesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserRolesResponseBody struct {
	// 扩展属性
	Content []*QueryUserRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBody) SetContent(v []*QueryUserRolesResponseBodyContent) *QueryUserRolesResponseBody {
	s.Content = v
	return s
}

type QueryUserRolesResponseBodyContent struct {
	// 角色编码
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s QueryUserRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBodyContent) SetRoleCode(v string) *QueryUserRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRolesResponseBodyContent) SetRoleName(v string) *QueryUserRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

type QueryUserRolesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponse) SetHeaders(v map[string]*string) *QueryUserRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRolesResponse) SetBody(v *QueryUserRolesResponseBody) *QueryUserRolesResponse {
	s.Body = v
	return s
}

type SaveUserExtendValuesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveUserExtendValuesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesHeaders) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesHeaders) SetCommonHeaders(v map[string]*string) *SaveUserExtendValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveUserExtendValuesHeaders) SetXAcsDingtalkAccessToken(v string) *SaveUserExtendValuesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveUserExtendValuesRequest struct {
	// 字段展示名称
	UserDisplayName *string `json:"userDisplayName,omitempty" xml:"userDisplayName,omitempty"`
	// 用户拓展字段key
	UserExtendKey *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	// 用户扩展字段value
	UserExtendValue *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s SaveUserExtendValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesRequest) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesRequest) SetUserDisplayName(v string) *SaveUserExtendValuesRequest {
	s.UserDisplayName = &v
	return s
}

func (s *SaveUserExtendValuesRequest) SetUserExtendKey(v string) *SaveUserExtendValuesRequest {
	s.UserExtendKey = &v
	return s
}

func (s *SaveUserExtendValuesRequest) SetUserExtendValue(v string) *SaveUserExtendValuesRequest {
	s.UserExtendValue = &v
	return s
}

type SaveUserExtendValuesResponseBody struct {
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveUserExtendValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesResponseBody) SetSuccess(v bool) *SaveUserExtendValuesResponseBody {
	s.Success = &v
	return s
}

type SaveUserExtendValuesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveUserExtendValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveUserExtendValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesResponse) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesResponse) SetHeaders(v map[string]*string) *SaveUserExtendValuesResponse {
	s.Headers = v
	return s
}

func (s *SaveUserExtendValuesResponse) SetBody(v *SaveUserExtendValuesResponseBody) *SaveUserExtendValuesResponse {
	s.Body = v
	return s
}

type UpdateUserExtendInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUserExtendInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserExtendInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserExtendInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUserExtendInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUserExtendInfoRequest struct {
	// comments
	Comments *string `json:"comments,omitempty" xml:"comments,omitempty"`
	// 职称code
	JobCode *string `json:"jobCode,omitempty" xml:"jobCode,omitempty"`
	// 工作状态code
	JobStatusCode []*string `json:"jobStatusCode,omitempty" xml:"jobStatusCode,omitempty" type:"Repeated"`
	// 用户属性code
	UserProbCode *string `json:"userProbCode,omitempty" xml:"userProbCode,omitempty"`
}

func (s UpdateUserExtendInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoRequest) SetComments(v string) *UpdateUserExtendInfoRequest {
	s.Comments = &v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetJobCode(v string) *UpdateUserExtendInfoRequest {
	s.JobCode = &v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetJobStatusCode(v []*string) *UpdateUserExtendInfoRequest {
	s.JobStatusCode = v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetUserProbCode(v string) *UpdateUserExtendInfoRequest {
	s.UserProbCode = &v
	return s
}

type UpdateUserExtendInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateUserExtendInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoResponse) SetHeaders(v map[string]*string) *UpdateUserExtendInfoResponse {
	s.Headers = v
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

func (client *Client) CustomizeContactCreate(request *CustomizeContactCreateRequest) (_result *CustomizeContactCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactCreateHeaders{}
	_result = &CustomizeContactCreateResponse{}
	_body, _err := client.CustomizeContactCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactCreateWithOptions(request *CustomizeContactCreateRequest, headers *CustomizeContactCreateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

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
	_result = &CustomizeContactCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactCreate"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/customizations/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDelete(request *CustomizeContactDeleteRequest) (_result *CustomizeContactDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeleteHeaders{}
	_result = &CustomizeContactDeleteResponse{}
	_body, _err := client.CustomizeContactDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeleteWithOptions(request *CustomizeContactDeleteRequest, headers *CustomizeContactDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeleteResponse, _err error) {
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
	_result = &CustomizeContactDeleteResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDelete"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/industry/customizations/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptCreate(request *CustomizeContactDeptCreateRequest) (_result *CustomizeContactDeptCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptCreateHeaders{}
	_result = &CustomizeContactDeptCreateResponse{}
	_body, _err := client.CustomizeContactDeptCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptCreateWithOptions(request *CustomizeContactDeptCreateRequest, headers *CustomizeContactDeptCreateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDeptId)) {
		body["parentDeptId"] = request.ParentDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.RefId)) {
		body["refId"] = request.RefId
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
	_result = &CustomizeContactDeptCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDeptCreate"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/customizations/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptDelete(request *CustomizeContactDeptDeleteRequest) (_result *CustomizeContactDeptDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptDeleteHeaders{}
	_result = &CustomizeContactDeptDeleteResponse{}
	_body, _err := client.CustomizeContactDeptDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptDeleteWithOptions(request *CustomizeContactDeptDeleteRequest, headers *CustomizeContactDeptDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

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
	_result = &CustomizeContactDeptDeleteResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDeptDelete"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/industry/customizations/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptInfo(request *CustomizeContactDeptInfoRequest) (_result *CustomizeContactDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptInfoHeaders{}
	_result = &CustomizeContactDeptInfoResponse{}
	_body, _err := client.CustomizeContactDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptInfoWithOptions(request *CustomizeContactDeptInfoRequest, headers *CustomizeContactDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

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
	_result = &CustomizeContactDeptInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDeptInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/customizations/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptList(request *CustomizeContactDeptListRequest) (_result *CustomizeContactDeptListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptListHeaders{}
	_result = &CustomizeContactDeptListResponse{}
	_body, _err := client.CustomizeContactDeptListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptListWithOptions(request *CustomizeContactDeptListRequest, headers *CustomizeContactDeptListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

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
	_result = &CustomizeContactDeptListResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDeptList"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/customizations/subsidiaryDepartments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptUpdate(request *CustomizeContactDeptUpdateRequest) (_result *CustomizeContactDeptUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptUpdateHeaders{}
	_result = &CustomizeContactDeptUpdateResponse{}
	_body, _err := client.CustomizeContactDeptUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptUpdateWithOptions(request *CustomizeContactDeptUpdateRequest, headers *CustomizeContactDeptUpdateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDeptId)) {
		body["parentDeptId"] = request.ParentDeptId
	}

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
	_result = &CustomizeContactDeptUpdateResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactDeptUpdate"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/industry/customizations/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpAdd(request *CustomizeContactEmpAddRequest) (_result *CustomizeContactEmpAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpAddHeaders{}
	_result = &CustomizeContactEmpAddResponse{}
	_body, _err := client.CustomizeContactEmpAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpAddWithOptions(request *CustomizeContactEmpAddRequest, headers *CustomizeContactEmpAddHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
	_result = &CustomizeContactEmpAddResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactEmpAdd"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/customizations/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpDelete(request *CustomizeContactEmpDeleteRequest) (_result *CustomizeContactEmpDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpDeleteHeaders{}
	_result = &CustomizeContactEmpDeleteResponse{}
	_body, _err := client.CustomizeContactEmpDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpDeleteWithOptions(request *CustomizeContactEmpDeleteRequest, headers *CustomizeContactEmpDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
	_result = &CustomizeContactEmpDeleteResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactEmpDelete"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/customizations/users/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpList(request *CustomizeContactEmpListRequest) (_result *CustomizeContactEmpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpListHeaders{}
	_result = &CustomizeContactEmpListResponse{}
	_body, _err := client.CustomizeContactEmpListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpListWithOptions(request *CustomizeContactEmpListRequest, headers *CustomizeContactEmpListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

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
	_result = &CustomizeContactEmpListResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactEmpList"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/customizations/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactList() (_result *CustomizeContactListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactListHeaders{}
	_result = &CustomizeContactListResponse{}
	_body, _err := client.CustomizeContactListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactListWithOptions(headers *CustomizeContactListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactListResponse, _err error) {
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
	_result = &CustomizeContactListResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactList"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/customizations/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactUpdate(request *CustomizeContactUpdateRequest) (_result *CustomizeContactUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactUpdateHeaders{}
	_result = &CustomizeContactUpdateResponse{}
	_body, _err := client.CustomizeContactUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactUpdateWithOptions(request *CustomizeContactUpdateRequest, headers *CustomizeContactUpdateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

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
	_result = &CustomizeContactUpdateResponse{}
	_body, _err := client.DoROARequest(tea.String("CustomizeContactUpdate"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/industry/customizations/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreContactInfo() (_result *DigitalStoreContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreContactInfoHeaders{}
	_result = &DigitalStoreContactInfoResponse{}
	_body, _err := client.DigitalStoreContactInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreContactInfoWithOptions(headers *DigitalStoreContactInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreContactInfoResponse, _err error) {
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
	_result = &DigitalStoreContactInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreContactInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/contactInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreGroupInfo(request *DigitalStoreGroupInfoRequest) (_result *DigitalStoreGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreGroupInfoHeaders{}
	_result = &DigitalStoreGroupInfoResponse{}
	_body, _err := client.DigitalStoreGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreGroupInfoWithOptions(request *DigitalStoreGroupInfoRequest, headers *DigitalStoreGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

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
	_result = &DigitalStoreGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreGroupInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/groupInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreGroups() (_result *DigitalStoreGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreGroupsHeaders{}
	_result = &DigitalStoreGroupsResponse{}
	_body, _err := client.DigitalStoreGroupsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreGroupsWithOptions(headers *DigitalStoreGroupsHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreGroupsResponse, _err error) {
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
	_result = &DigitalStoreGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreGroups"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreNodeInfo(request *DigitalStoreNodeInfoRequest) (_result *DigitalStoreNodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreNodeInfoHeaders{}
	_result = &DigitalStoreNodeInfoResponse{}
	_body, _err := client.DigitalStoreNodeInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreNodeInfoWithOptions(request *DigitalStoreNodeInfoRequest, headers *DigitalStoreNodeInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreNodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
	}

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
	_result = &DigitalStoreNodeInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreNodeInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/nodeInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreRoles() (_result *DigitalStoreRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreRolesHeaders{}
	_result = &DigitalStoreRolesResponse{}
	_body, _err := client.DigitalStoreRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreRolesWithOptions(headers *DigitalStoreRolesHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreRolesResponse, _err error) {
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
	_result = &DigitalStoreRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreRoles"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreStoreInfo(request *DigitalStoreStoreInfoRequest) (_result *DigitalStoreStoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreStoreInfoHeaders{}
	_result = &DigitalStoreStoreInfoResponse{}
	_body, _err := client.DigitalStoreStoreInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreStoreInfoWithOptions(request *DigitalStoreStoreInfoRequest, headers *DigitalStoreStoreInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreStoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["storeId"] = request.StoreId
	}

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
	_result = &DigitalStoreStoreInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreStoreInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/storeInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreSubNodes(request *DigitalStoreSubNodesRequest) (_result *DigitalStoreSubNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreSubNodesHeaders{}
	_result = &DigitalStoreSubNodesResponse{}
	_body, _err := client.DigitalStoreSubNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreSubNodesWithOptions(request *DigitalStoreSubNodesRequest, headers *DigitalStoreSubNodesHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreSubNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
	}

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
	_result = &DigitalStoreSubNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreSubNodes"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/subsidiaryNodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreUserInfo(request *DigitalStoreUserInfoRequest) (_result *DigitalStoreUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreUserInfoHeaders{}
	_result = &DigitalStoreUserInfoResponse{}
	_body, _err := client.DigitalStoreUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreUserInfoWithOptions(request *DigitalStoreUserInfoRequest, headers *DigitalStoreUserInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
	_result = &DigitalStoreUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreUserInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/userInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreUsers(request *DigitalStoreUsersRequest) (_result *DigitalStoreUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreUsersHeaders{}
	_result = &DigitalStoreUsersResponse{}
	_body, _err := client.DigitalStoreUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreUsersWithOptions(request *DigitalStoreUsersRequest, headers *DigitalStoreUsersHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
	}

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
	_result = &DigitalStoreUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("DigitalStoreUsers"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/digitalStores/nodes/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureCostRecordListGet(request *IndustryManufactureCostRecordListGetRequest) (_result *IndustryManufactureCostRecordListGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureCostRecordListGetHeaders{}
	_result = &IndustryManufactureCostRecordListGetResponse{}
	_body, _err := client.IndustryManufactureCostRecordListGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureCostRecordListGetWithOptions(request *IndustryManufactureCostRecordListGetRequest, headers *IndustryManufactureCostRecordListGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureCostRecordListGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionTaskNo)) {
		body["productionTaskNo"] = request.ProductionTaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
	}

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
	_result = &IndustryManufactureCostRecordListGetResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustryManufactureCostRecordListGet"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/manufactures/materialCostRecords/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureFeeListGet(request *IndustryManufactureFeeListGetRequest) (_result *IndustryManufactureFeeListGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureFeeListGetHeaders{}
	_result = &IndustryManufactureFeeListGetResponse{}
	_body, _err := client.IndustryManufactureFeeListGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureFeeListGetWithOptions(request *IndustryManufactureFeeListGetRequest, headers *IndustryManufactureFeeListGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureFeeListGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionTaskNo)) {
		body["productionTaskNo"] = request.ProductionTaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
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
	_result = &IndustryManufactureFeeListGetResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustryManufactureFeeListGet"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/manufactures/fees/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureLabourCost(request *IndustryManufactureLabourCostRequest) (_result *IndustryManufactureLabourCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureLabourCostHeaders{}
	_result = &IndustryManufactureLabourCostResponse{}
	_body, _err := client.IndustryManufactureLabourCostWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureLabourCostWithOptions(request *IndustryManufactureLabourCostRequest, headers *IndustryManufactureLabourCostHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureLabourCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessNo)) {
		body["processNo"] = request.ProcessNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
	}

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
	_result = &IndustryManufactureLabourCostResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustryManufactureLabourCost"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/manufactures/labourCosts/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMaterialList(request *IndustryManufactureMaterialListRequest) (_result *IndustryManufactureMaterialListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMaterialListHeaders{}
	_result = &IndustryManufactureMaterialListResponse{}
	_body, _err := client.IndustryManufactureMaterialListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMaterialListWithOptions(request *IndustryManufactureMaterialListRequest, headers *IndustryManufactureMaterialListHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMaterialListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
	}

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
	_result = &IndustryManufactureMaterialListResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustryManufactureMaterialList"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/manufactures/materials/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryMmanufactureMaterialCostGet(request *IndustryMmanufactureMaterialCostGetRequest) (_result *IndustryMmanufactureMaterialCostGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryMmanufactureMaterialCostGetHeaders{}
	_result = &IndustryMmanufactureMaterialCostGetResponse{}
	_body, _err := client.IndustryMmanufactureMaterialCostGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryMmanufactureMaterialCostGetWithOptions(request *IndustryMmanufactureMaterialCostGetRequest, headers *IndustryMmanufactureMaterialCostGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryMmanufactureMaterialCostGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
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

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
	}

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
	_result = &IndustryMmanufactureMaterialCostGetResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustryMmanufactureMaterialCostGet"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/manufactures/base/materialCosts/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDepartment(request *QueryAllDepartmentRequest) (_result *QueryAllDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDepartmentHeaders{}
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.QueryAllDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDepartmentWithOptions(request *QueryAllDepartmentRequest, headers *QueryAllDepartmentHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDepartmentResponse, _err error) {
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
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllDepartment"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDoctors(request *QueryAllDoctorsRequest) (_result *QueryAllDoctorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDoctorsHeaders{}
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.QueryAllDoctorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDoctorsWithOptions(request *QueryAllDoctorsRequest, headers *QueryAllDoctorsHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDoctorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
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
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllDoctors"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/doctors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroup(request *QueryAllGroupRequest) (_result *QueryAllGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupHeaders{}
	_result = &QueryAllGroupResponse{}
	_body, _err := client.QueryAllGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupWithOptions(request *QueryAllGroupRequest, headers *QueryAllGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupResponse, _err error) {
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
	_result = &QueryAllGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllGroup"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroupsInDept(deptId *string, request *QueryAllGroupsInDeptRequest) (_result *QueryAllGroupsInDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupsInDeptHeaders{}
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.QueryAllGroupsInDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupsInDeptWithOptions(deptId *string, request *QueryAllGroupsInDeptRequest, headers *QueryAllGroupsInDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupsInDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deptId = openapiutil.GetEncodeParam(deptId)
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
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllGroupsInDept"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)+"/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByDept(deptId *string, request *QueryAllMemberByDeptRequest) (_result *QueryAllMemberByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByDeptHeaders{}
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.QueryAllMemberByDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByDeptWithOptions(deptId *string, request *QueryAllMemberByDeptRequest, headers *QueryAllMemberByDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deptId = openapiutil.GetEncodeParam(deptId)
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
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllMemberByDept"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByGroup(groupId *string, request *QueryAllMemberByGroupRequest) (_result *QueryAllMemberByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByGroupHeaders{}
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.QueryAllMemberByGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByGroupWithOptions(groupId *string, request *QueryAllMemberByGroupRequest, headers *QueryAllMemberByGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	groupId = openapiutil.GetEncodeParam(groupId)
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
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllMemberByGroup"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups/"+tea.StringValue(groupId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizOptLog(request *QueryBizOptLogRequest) (_result *QueryBizOptLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBizOptLogHeaders{}
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.QueryBizOptLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizOptLogWithOptions(request *QueryBizOptLogRequest, headers *QueryBizOptLogHeaders, runtime *util.RuntimeOptions) (_result *QueryBizOptLogResponse, _err error) {
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
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryBizOptLog"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/bizOptLogs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDepartmentInfo(deptId *string) (_result *QueryDepartmentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDepartmentInfoHeaders{}
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.QueryDepartmentInfoWithOptions(deptId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDepartmentInfoWithOptions(deptId *string, headers *QueryDepartmentInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDepartmentInfoResponse, _err error) {
	deptId = openapiutil.GetEncodeParam(deptId)
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
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDepartmentInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/departments/"+tea.StringValue(deptId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupInfo(groupId *string) (_result *QueryGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupInfoHeaders{}
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.QueryGroupInfoWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupInfoWithOptions(groupId *string, headers *QueryGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupInfoResponse, _err error) {
	groupId = openapiutil.GetEncodeParam(groupId)
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
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryGroupInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/groups/"+tea.StringValue(groupId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalDistrictInfo(request *QueryHospitalDistrictInfoRequest) (_result *QueryHospitalDistrictInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalDistrictInfoHeaders{}
	_result = &QueryHospitalDistrictInfoResponse{}
	_body, _err := client.QueryHospitalDistrictInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalDistrictInfoWithOptions(request *QueryHospitalDistrictInfoRequest, headers *QueryHospitalDistrictInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalDistrictInfoResponse, _err error) {
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
	_result = &QueryHospitalDistrictInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHospitalDistrictInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/districts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalRoleUserInfo(request *QueryHospitalRoleUserInfoRequest) (_result *QueryHospitalRoleUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalRoleUserInfoHeaders{}
	_result = &QueryHospitalRoleUserInfoResponse{}
	_body, _err := client.QueryHospitalRoleUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalRoleUserInfoWithOptions(request *QueryHospitalRoleUserInfoRequest, headers *QueryHospitalRoleUserInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalRoleUserInfoResponse, _err error) {
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
	_result = &QueryHospitalRoleUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHospitalRoleUserInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/roles/userInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalRoles() (_result *QueryHospitalRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalRolesHeaders{}
	_result = &QueryHospitalRolesResponse{}
	_body, _err := client.QueryHospitalRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalRolesWithOptions(headers *QueryHospitalRolesHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalRolesResponse, _err error) {
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
	_result = &QueryHospitalRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryHospitalRoles"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobCodeDictionary() (_result *QueryJobCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobCodeDictionaryHeaders{}
	_result = &QueryJobCodeDictionaryResponse{}
	_body, _err := client.QueryJobCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobCodeDictionaryWithOptions(headers *QueryJobCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryJobCodeDictionaryResponse, _err error) {
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
	_result = &QueryJobCodeDictionaryResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobCodeDictionary"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/jobCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobStatusCodeDictionary() (_result *QueryJobStatusCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobStatusCodeDictionaryHeaders{}
	_result = &QueryJobStatusCodeDictionaryResponse{}
	_body, _err := client.QueryJobStatusCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobStatusCodeDictionaryWithOptions(headers *QueryJobStatusCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryJobStatusCodeDictionaryResponse, _err error) {
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
	_result = &QueryJobStatusCodeDictionaryResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryJobStatusCodeDictionary"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/jobStatusCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMedicalEvents() (_result *QueryMedicalEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMedicalEventsHeaders{}
	_result = &QueryMedicalEventsResponse{}
	_body, _err := client.QueryMedicalEventsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMedicalEventsWithOptions(headers *QueryMedicalEventsHeaders, runtime *util.RuntimeOptions) (_result *QueryMedicalEventsResponse, _err error) {
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
	_result = &QueryMedicalEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryMedicalEvents"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserCredentials(request *QueryUserCredentialsRequest) (_result *QueryUserCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserCredentialsHeaders{}
	_result = &QueryUserCredentialsResponse{}
	_body, _err := client.QueryUserCredentialsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserCredentialsWithOptions(request *QueryUserCredentialsRequest, headers *QueryUserCredentialsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_result = &QueryUserCredentialsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserCredentials"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/credentials/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserExtInfo(userId *string) (_result *QueryUserExtInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserExtInfoHeaders{}
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.QueryUserExtInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserExtInfoWithOptions(userId *string, headers *QueryUserExtInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserExtInfoResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserExtInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/extInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserExtendValues(request *QueryUserExtendValuesRequest) (_result *QueryUserExtendValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserExtendValuesHeaders{}
	_result = &QueryUserExtendValuesResponse{}
	_body, _err := client.QueryUserExtendValuesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserExtendValuesWithOptions(request *QueryUserExtendValuesRequest, headers *QueryUserExtendValuesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserExtendValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserExtendKey)) {
		body["userExtendKey"] = request.UserExtendKey
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
	_result = &QueryUserExtendValuesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserExtendValues"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/extends/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserInfo(userId *string) (_result *QueryUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserInfoHeaders{}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.QueryUserInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoWithOptions(userId *string, headers *QueryUserInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserInfoResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserProbCodeDictionary() (_result *QueryUserProbCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserProbCodeDictionaryHeaders{}
	_result = &QueryUserProbCodeDictionaryResponse{}
	_body, _err := client.QueryUserProbCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserProbCodeDictionaryWithOptions(headers *QueryUserProbCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryUserProbCodeDictionaryResponse, _err error) {
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
	_result = &QueryUserProbCodeDictionaryResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserProbCodeDictionary"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/userProbCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserRoles(userId *string) (_result *QueryUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserRolesHeaders{}
	_result = &QueryUserRolesResponse{}
	_body, _err := client.QueryUserRolesWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserRolesWithOptions(userId *string, headers *QueryUserRolesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserRolesResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserRolesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserRoles"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveUserExtendValues(userId *string, request *SaveUserExtendValuesRequest) (_result *SaveUserExtendValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveUserExtendValuesHeaders{}
	_result = &SaveUserExtendValuesResponse{}
	_body, _err := client.SaveUserExtendValuesWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveUserExtendValuesWithOptions(userId *string, request *SaveUserExtendValuesRequest, headers *SaveUserExtendValuesHeaders, runtime *util.RuntimeOptions) (_result *SaveUserExtendValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserDisplayName)) {
		query["userDisplayName"] = request.UserDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.UserExtendKey)) {
		query["userExtendKey"] = request.UserExtendKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserExtendValue)) {
		query["userExtendValue"] = request.UserExtendValue
	}

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
	_result = &SaveUserExtendValuesResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveUserExtendValues"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/extends"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserExtendInfo(userId *string, request *UpdateUserExtendInfoRequest) (_result *UpdateUserExtendInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserExtendInfoHeaders{}
	_result = &UpdateUserExtendInfoResponse{}
	_body, _err := client.UpdateUserExtendInfoWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserExtendInfoWithOptions(userId *string, request *UpdateUserExtendInfoRequest, headers *UpdateUserExtendInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserExtendInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		body["comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.JobCode)) {
		body["jobCode"] = request.JobCode
	}

	if !tea.BoolValue(util.IsUnset(request.JobStatusCode)) {
		body["jobStatusCode"] = request.JobStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserProbCode)) {
		body["userProbCode"] = request.UserProbCode
	}

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
	_result = &UpdateUserExtendInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateUserExtendInfo"), tea.String("industry_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/industry/medicals/users/"+tea.StringValue(userId)+"/extInfos"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
