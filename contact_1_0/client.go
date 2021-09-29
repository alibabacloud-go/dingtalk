// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package contact_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SortUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SortUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s SortUserHeaders) GoString() string {
	return s.String()
}

func (s *SortUserHeaders) SetCommonHeaders(v map[string]*string) *SortUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SortUserHeaders) SetXAcsDingtalkAccessToken(v string) *SortUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SortUserRequest struct {
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 用户id列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// 0 根据姓名拼音升序排列 1 根据姓名拼音降序排列
	SortType *int32 `json:"sortType,omitempty" xml:"sortType,omitempty"`
}

func (s SortUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SortUserRequest) GoString() string {
	return s.String()
}

func (s *SortUserRequest) SetDingOrgId(v int64) *SortUserRequest {
	s.DingOrgId = &v
	return s
}

func (s *SortUserRequest) SetUserIdList(v []*string) *SortUserRequest {
	s.UserIdList = v
	return s
}

func (s *SortUserRequest) SetSortType(v int32) *SortUserRequest {
	s.SortType = &v
	return s
}

type SortUserResponseBody struct {
	// userIdList
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s SortUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SortUserResponseBody) GoString() string {
	return s.String()
}

func (s *SortUserResponseBody) SetUserIdList(v []*string) *SortUserResponseBody {
	s.UserIdList = v
	return s
}

type SortUserResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SortUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SortUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SortUserResponse) GoString() string {
	return s.String()
}

func (s *SortUserResponse) SetHeaders(v map[string]*string) *SortUserResponse {
	s.Headers = v
	return s
}

func (s *SortUserResponse) SetBody(v *SortUserResponseBody) *SortUserResponse {
	s.Body = v
	return s
}

type ListContactHideSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListContactHideSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsHeaders) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsHeaders) SetCommonHeaders(v map[string]*string) *ListContactHideSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListContactHideSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *ListContactHideSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListContactHideSettingsRequest struct {
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s ListContactHideSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsRequest) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsRequest) SetNextToken(v int64) *ListContactHideSettingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListContactHideSettingsRequest) SetMaxResults(v int32) *ListContactHideSettingsRequest {
	s.MaxResults = &v
	return s
}

type ListContactHideSettingsResponseBody struct {
	// 是否还有数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一次拉取数据时的offset
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 设置列表
	List []*ListContactHideSettingsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListContactHideSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponseBody) SetHasMore(v bool) *ListContactHideSettingsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListContactHideSettingsResponseBody) SetNextToken(v int64) *ListContactHideSettingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContactHideSettingsResponseBody) SetList(v []*ListContactHideSettingsResponseBodyList) *ListContactHideSettingsResponseBody {
	s.List = v
	return s
}

type ListContactHideSettingsResponseBodyList struct {
	// 设置名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 设置描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 要隐藏的员工列表
	ObjectStaffIds []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	// 要隐藏的部门列表
	ObjectDeptIds []*int64 `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	// 要影藏的角色列表
	ObjectTagIds []*int64 `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	// 白名单用户列表
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	// 白名单部门列表
	ExcludeDeptIds []*int64 `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	// 白名单角色列表
	ExcludeTagIds []*int64 `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	// 规则是否生效
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// settingId
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListContactHideSettingsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponseBodyList) SetName(v string) *ListContactHideSettingsResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetDescription(v string) *ListContactHideSettingsResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectStaffIds(v []*string) *ListContactHideSettingsResponseBodyList {
	s.ObjectStaffIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectDeptIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ObjectDeptIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetObjectTagIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ObjectTagIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeStaffIds(v []*string) *ListContactHideSettingsResponseBodyList {
	s.ExcludeStaffIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeDeptIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ExcludeDeptIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetExcludeTagIds(v []*int64) *ListContactHideSettingsResponseBodyList {
	s.ExcludeTagIds = v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetActive(v bool) *ListContactHideSettingsResponseBodyList {
	s.Active = &v
	return s
}

func (s *ListContactHideSettingsResponseBodyList) SetId(v int64) *ListContactHideSettingsResponseBodyList {
	s.Id = &v
	return s
}

type ListContactHideSettingsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContactHideSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContactHideSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactHideSettingsResponse) GoString() string {
	return s.String()
}

func (s *ListContactHideSettingsResponse) SetHeaders(v map[string]*string) *ListContactHideSettingsResponse {
	s.Headers = v
	return s
}

func (s *ListContactHideSettingsResponse) SetBody(v *ListContactHideSettingsResponseBody) *ListContactHideSettingsResponse {
	s.Body = v
	return s
}

type UpdateContactHideSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateContactHideSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateContactHideSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateContactHideSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateContactHideSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateContactHideSettingRequest struct {
	// 设置名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 设置描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 隐藏员工列表
	ObjectStaffIds []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	// 影藏部门列表
	ObjectDeptIds []*int64 `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	// 影藏角色列表
	ObjectTagIds []*int64 `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	// 白名单员工列表
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	// 白名单部门列表
	ExcludeDeptIds []*int64 `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	// 白名单角色列表
	ExcludeTagIds []*int64 `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	// 是否激活
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// settingId
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateContactHideSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingRequest) SetName(v string) *UpdateContactHideSettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetDescription(v string) *UpdateContactHideSettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectStaffIds(v []*string) *UpdateContactHideSettingRequest {
	s.ObjectStaffIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectDeptIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetObjectTagIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeStaffIds(v []*string) *UpdateContactHideSettingRequest {
	s.ExcludeStaffIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeDeptIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetExcludeTagIds(v []*int64) *UpdateContactHideSettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateContactHideSettingRequest) SetActive(v bool) *UpdateContactHideSettingRequest {
	s.Active = &v
	return s
}

func (s *UpdateContactHideSettingRequest) SetId(v int64) *UpdateContactHideSettingRequest {
	s.Id = &v
	return s
}

type UpdateContactHideSettingResponseBody struct {
	// settingId
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateContactHideSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingResponseBody) SetResult(v int64) *UpdateContactHideSettingResponseBody {
	s.Result = &v
	return s
}

type UpdateContactHideSettingResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateContactHideSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactHideSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactHideSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactHideSettingResponse) SetHeaders(v map[string]*string) *UpdateContactHideSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactHideSettingResponse) SetBody(v *UpdateContactHideSettingResponseBody) *UpdateContactHideSettingResponse {
	s.Body = v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateEmpAttrbuteVisibilitySettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateEmpAttrbuteVisibilitySettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateEmpAttrbuteVisibilitySettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingRequest struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// object员工id列表
	ObjectStaffIds []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	// object部门id列表
	ObjectDeptIds []*int64 `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	// object角色id列表
	ObjectTagIds []*int64 `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	// 隐藏字段id列表
	HideFields []*string `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	// 例外员工id列表
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	// 例外部门id列表
	ExcludeDeptIds []*int64 `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	// 例外角色id列表
	ExcludeTagIds []*int64 `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	// 是否生效
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s UpdateEmpAttrbuteVisibilitySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetId(v int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Id = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetName(v string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Name = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetDescription(v string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Description = &v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectStaffIds(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectStaffIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectDeptIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectDeptIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetObjectTagIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ObjectTagIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetHideFields(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.HideFields = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeStaffIds(v []*string) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeStaffIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeDeptIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeDeptIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetExcludeTagIds(v []*int64) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.ExcludeTagIds = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingRequest) SetActive(v bool) *UpdateEmpAttrbuteVisibilitySettingRequest {
	s.Active = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingResponseBody struct {
	// settingId
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateEmpAttrbuteVisibilitySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponseBody) SetResult(v int64) *UpdateEmpAttrbuteVisibilitySettingResponseBody {
	s.Result = &v
	return s
}

type UpdateEmpAttrbuteVisibilitySettingResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEmpAttrbuteVisibilitySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEmpAttrbuteVisibilitySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmpAttrbuteVisibilitySettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponse) SetHeaders(v map[string]*string) *UpdateEmpAttrbuteVisibilitySettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmpAttrbuteVisibilitySettingResponse) SetBody(v *UpdateEmpAttrbuteVisibilitySettingResponseBody) *UpdateEmpAttrbuteVisibilitySettingResponse {
	s.Body = v
	return s
}

type DeleteEmpAttributeVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteEmpAttributeVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeVisibilityHeaders) SetCommonHeaders(v map[string]*string) *DeleteEmpAttributeVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEmpAttributeVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteEmpAttributeVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteEmpAttributeVisibilityResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteEmpAttributeVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEmpAttributeVisibilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmpAttributeVisibilityResponse) SetHeaders(v map[string]*string) *DeleteEmpAttributeVisibilityResponse {
	s.Headers = v
	return s
}

type SearchDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *SearchDepartmentHeaders) SetCommonHeaders(v map[string]*string) *SearchDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *SearchDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchDepartmentRequest struct {
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 部门名称或者部门名称拼音
	QueryWord *string `json:"queryWord,omitempty" xml:"queryWord,omitempty"`
	// 分页查询锚点
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 分页长度
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s SearchDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentRequest) GoString() string {
	return s.String()
}

func (s *SearchDepartmentRequest) SetDingOrgId(v int64) *SearchDepartmentRequest {
	s.DingOrgId = &v
	return s
}

func (s *SearchDepartmentRequest) SetQueryWord(v string) *SearchDepartmentRequest {
	s.QueryWord = &v
	return s
}

func (s *SearchDepartmentRequest) SetOffset(v int32) *SearchDepartmentRequest {
	s.Offset = &v
	return s
}

func (s *SearchDepartmentRequest) SetSize(v int32) *SearchDepartmentRequest {
	s.Size = &v
	return s
}

type SearchDepartmentResponseBody struct {
	HasMore    *bool    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	TotalCount *int64   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	List       []*int64 `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SearchDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDepartmentResponseBody) SetHasMore(v bool) *SearchDepartmentResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchDepartmentResponseBody) SetTotalCount(v int64) *SearchDepartmentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchDepartmentResponseBody) SetList(v []*int64) *SearchDepartmentResponseBody {
	s.List = v
	return s
}

type SearchDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDepartmentResponse) GoString() string {
	return s.String()
}

func (s *SearchDepartmentResponse) SetHeaders(v map[string]*string) *SearchDepartmentResponse {
	s.Headers = v
	return s
}

func (s *SearchDepartmentResponse) SetBody(v *SearchDepartmentResponseBody) *SearchDepartmentResponse {
	s.Body = v
	return s
}

type ListManagementGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListManagementGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsHeaders) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsHeaders) SetCommonHeaders(v map[string]*string) *ListManagementGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListManagementGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *ListManagementGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListManagementGroupsRequest struct {
	// 开始读取的位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s ListManagementGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsRequest) SetNextToken(v int64) *ListManagementGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListManagementGroupsRequest) SetMaxResults(v int32) *ListManagementGroupsRequest {
	s.MaxResults = &v
	return s
}

type ListManagementGroupsResponseBody struct {
	// 下一次读取的位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有下一页
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 管理组列表
	Groups []*ListManagementGroupsResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
}

func (s ListManagementGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBody) SetNextToken(v int64) *ListManagementGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManagementGroupsResponseBody) SetHasMore(v bool) *ListManagementGroupsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListManagementGroupsResponseBody) SetGroups(v []*ListManagementGroupsResponseBodyGroups) *ListManagementGroupsResponseBody {
	s.Groups = v
	return s
}

type ListManagementGroupsResponseBodyGroups struct {
	// 管理组id
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 管理组名
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 成员
	Members []*ListManagementGroupsResponseBodyGroupsMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 管理范围
	Scope *ListManagementGroupsResponseBodyGroupsScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// 资源列表
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s ListManagementGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroups) SetGroupId(v string) *ListManagementGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetGroupName(v string) *ListManagementGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetMembers(v []*ListManagementGroupsResponseBodyGroupsMembers) *ListManagementGroupsResponseBodyGroups {
	s.Members = v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetScope(v *ListManagementGroupsResponseBodyGroupsScope) *ListManagementGroupsResponseBodyGroups {
	s.Scope = v
	return s
}

func (s *ListManagementGroupsResponseBodyGroups) SetResourceIds(v []*string) *ListManagementGroupsResponseBodyGroups {
	s.ResourceIds = v
	return s
}

type ListManagementGroupsResponseBodyGroupsMembers struct {
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
}

func (s ListManagementGroupsResponseBodyGroupsMembers) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroupsMembers) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroupsMembers) SetMemberType(v string) *ListManagementGroupsResponseBodyGroupsMembers {
	s.MemberType = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroupsMembers) SetMemberId(v string) *ListManagementGroupsResponseBodyGroupsMembers {
	s.MemberId = &v
	return s
}

type ListManagementGroupsResponseBodyGroupsScope struct {
	// 1
	ScopeType *int32 `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门列表，只在scopeType=3 生效
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
}

func (s ListManagementGroupsResponseBodyGroupsScope) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponseBodyGroupsScope) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponseBodyGroupsScope) SetScopeType(v int32) *ListManagementGroupsResponseBodyGroupsScope {
	s.ScopeType = &v
	return s
}

func (s *ListManagementGroupsResponseBodyGroupsScope) SetDeptIds(v []*int64) *ListManagementGroupsResponseBodyGroupsScope {
	s.DeptIds = v
	return s
}

type ListManagementGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListManagementGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListManagementGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManagementGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListManagementGroupsResponse) SetHeaders(v map[string]*string) *ListManagementGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListManagementGroupsResponse) SetBody(v *ListManagementGroupsResponseBody) *ListManagementGroupsResponse {
	s.Body = v
	return s
}

type ListSeniorSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSeniorSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsHeaders) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsHeaders) SetCommonHeaders(v map[string]*string) *ListSeniorSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSeniorSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSeniorSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSeniorSettingsRequest struct {
	SeniorStaffId *string `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
}

func (s ListSeniorSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsRequest) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsRequest) SetSeniorStaffId(v string) *ListSeniorSettingsRequest {
	s.SeniorStaffId = &v
	return s
}

type ListSeniorSettingsResponseBody struct {
	// Id of the request
	SeniorStaffId   *string                                          `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
	ProtectScenes   []*string                                        `json:"protectScenes,omitempty" xml:"protectScenes,omitempty" type:"Repeated"`
	SeniorWhiteList []*ListSeniorSettingsResponseBodySeniorWhiteList `json:"seniorWhiteList,omitempty" xml:"seniorWhiteList,omitempty" type:"Repeated"`
}

func (s ListSeniorSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponseBody) SetSeniorStaffId(v string) *ListSeniorSettingsResponseBody {
	s.SeniorStaffId = &v
	return s
}

func (s *ListSeniorSettingsResponseBody) SetProtectScenes(v []*string) *ListSeniorSettingsResponseBody {
	s.ProtectScenes = v
	return s
}

func (s *ListSeniorSettingsResponseBody) SetSeniorWhiteList(v []*ListSeniorSettingsResponseBodySeniorWhiteList) *ListSeniorSettingsResponseBody {
	s.SeniorWhiteList = v
	return s
}

type ListSeniorSettingsResponseBodySeniorWhiteList struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Type *int32  `json:"type,omitempty" xml:"type,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSeniorSettingsResponseBodySeniorWhiteList) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponseBodySeniorWhiteList) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetId(v string) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Id = &v
	return s
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetType(v int32) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Type = &v
	return s
}

func (s *ListSeniorSettingsResponseBodySeniorWhiteList) SetName(v string) *ListSeniorSettingsResponseBodySeniorWhiteList {
	s.Name = &v
	return s
}

type ListSeniorSettingsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSeniorSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSeniorSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSeniorSettingsResponse) GoString() string {
	return s.String()
}

func (s *ListSeniorSettingsResponse) SetHeaders(v map[string]*string) *ListSeniorSettingsResponse {
	s.Headers = v
	return s
}

func (s *ListSeniorSettingsResponse) SetBody(v *ListSeniorSettingsResponseBody) *ListSeniorSettingsResponse {
	s.Body = v
	return s
}

type ListEmpAttributeVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEmpAttributeVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityHeaders) SetCommonHeaders(v map[string]*string) *ListEmpAttributeVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEmpAttributeVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *ListEmpAttributeVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEmpAttributeVisibilityRequest struct {
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s ListEmpAttributeVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityRequest) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityRequest) SetNextToken(v int64) *ListEmpAttributeVisibilityRequest {
	s.NextToken = &v
	return s
}

func (s *ListEmpAttributeVisibilityRequest) SetMaxResults(v int32) *ListEmpAttributeVisibilityRequest {
	s.MaxResults = &v
	return s
}

type ListEmpAttributeVisibilityResponseBody struct {
	// 是否还有数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一次拉取时的offset
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 设置列表
	List []*ListEmpAttributeVisibilityResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListEmpAttributeVisibilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponseBody) SetHasMore(v bool) *ListEmpAttributeVisibilityResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBody) SetNextCursor(v int64) *ListEmpAttributeVisibilityResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBody) SetList(v []*ListEmpAttributeVisibilityResponseBodyList) *ListEmpAttributeVisibilityResponseBody {
	s.List = v
	return s
}

type ListEmpAttributeVisibilityResponseBodyList struct {
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 设置时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 设置描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 被查看的用户id列表
	ObjectStaffIds []*string `json:"objectStaffIds,omitempty" xml:"objectStaffIds,omitempty" type:"Repeated"`
	// 被查看的部门id列表
	ObjectDeptIds []*int64 `json:"objectDeptIds,omitempty" xml:"objectDeptIds,omitempty" type:"Repeated"`
	// 被查看的角色id列表
	ObjectTagIds []*int64 `json:"objectTagIds,omitempty" xml:"objectTagIds,omitempty" type:"Repeated"`
	// 隐藏的字段id列表
	HideFields []*string `json:"hideFields,omitempty" xml:"hideFields,omitempty" type:"Repeated"`
	// 白名单用户id列表
	ExcludeStaffIds []*string `json:"excludeStaffIds,omitempty" xml:"excludeStaffIds,omitempty" type:"Repeated"`
	// 白名单部门id列表
	ExcludeDeptIds []*int64 `json:"excludeDeptIds,omitempty" xml:"excludeDeptIds,omitempty" type:"Repeated"`
	// 白名单角色id列表
	ExcludeTagIds []*int64 `json:"excludeTagIds,omitempty" xml:"excludeTagIds,omitempty" type:"Repeated"`
	// 是否生效
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
}

func (s ListEmpAttributeVisibilityResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetId(v int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetGmtCreate(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetGmtModified(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetName(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetDescription(v string) *ListEmpAttributeVisibilityResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectStaffIds(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectStaffIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectDeptIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectDeptIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetObjectTagIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ObjectTagIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetHideFields(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.HideFields = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeStaffIds(v []*string) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeStaffIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeDeptIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeDeptIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetExcludeTagIds(v []*int64) *ListEmpAttributeVisibilityResponseBodyList {
	s.ExcludeTagIds = v
	return s
}

func (s *ListEmpAttributeVisibilityResponseBodyList) SetActive(v bool) *ListEmpAttributeVisibilityResponseBodyList {
	s.Active = &v
	return s
}

type ListEmpAttributeVisibilityResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEmpAttributeVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEmpAttributeVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEmpAttributeVisibilityResponse) GoString() string {
	return s.String()
}

func (s *ListEmpAttributeVisibilityResponse) SetHeaders(v map[string]*string) *ListEmpAttributeVisibilityResponse {
	s.Headers = v
	return s
}

func (s *ListEmpAttributeVisibilityResponse) SetBody(v *ListEmpAttributeVisibilityResponseBody) *ListEmpAttributeVisibilityResponse {
	s.Body = v
	return s
}

type SearchUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchUserHeaders) GoString() string {
	return s.String()
}

func (s *SearchUserHeaders) SetCommonHeaders(v map[string]*string) *SearchUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchUserHeaders) SetXAcsDingtalkAccessToken(v string) *SearchUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchUserRequest struct {
	DingOrgId *int64 `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	// 用户名称、名称拼音或英文名称
	QueryWord *string `json:"queryWord,omitempty" xml:"queryWord,omitempty"`
	// 分页查询锚点
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 分页长度
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s SearchUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchUserRequest) GoString() string {
	return s.String()
}

func (s *SearchUserRequest) SetDingOrgId(v int64) *SearchUserRequest {
	s.DingOrgId = &v
	return s
}

func (s *SearchUserRequest) SetQueryWord(v string) *SearchUserRequest {
	s.QueryWord = &v
	return s
}

func (s *SearchUserRequest) SetOffset(v int32) *SearchUserRequest {
	s.Offset = &v
	return s
}

func (s *SearchUserRequest) SetSize(v int32) *SearchUserRequest {
	s.Size = &v
	return s
}

type SearchUserResponseBody struct {
	// Id of the request
	HasMore    *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	TotalCount *int64    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	List       []*string `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SearchUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponseBody) GoString() string {
	return s.String()
}

func (s *SearchUserResponseBody) SetHasMore(v bool) *SearchUserResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchUserResponseBody) SetTotalCount(v int64) *SearchUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchUserResponseBody) SetList(v []*string) *SearchUserResponseBody {
	s.List = v
	return s
}

type SearchUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponse) GoString() string {
	return s.String()
}

func (s *SearchUserResponse) SetHeaders(v map[string]*string) *SearchUserResponse {
	s.Headers = v
	return s
}

func (s *SearchUserResponse) SetBody(v *SearchUserResponseBody) *SearchUserResponse {
	s.Body = v
	return s
}

type GetTranslateFileJobResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTranslateFileJobResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultHeaders) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultHeaders) SetCommonHeaders(v map[string]*string) *GetTranslateFileJobResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTranslateFileJobResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetTranslateFileJobResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTranslateFileJobResultRequest struct {
	// 异步转译任务id
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s GetTranslateFileJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultRequest) SetJobId(v string) *GetTranslateFileJobResultRequest {
	s.JobId = &v
	return s
}

type GetTranslateFileJobResultResponseBody struct {
	// 0 任务进行中 1 任务处理成功 2 任务处理失败
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 文件内容转译成功后的url，需要用户通过oauth2授权登录在用户侧打开
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetTranslateFileJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultResponseBody) SetStatus(v string) *GetTranslateFileJobResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetTranslateFileJobResultResponseBody) SetUrl(v string) *GetTranslateFileJobResultResponseBody {
	s.Url = &v
	return s
}

type GetTranslateFileJobResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTranslateFileJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranslateFileJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranslateFileJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetTranslateFileJobResultResponse) SetHeaders(v map[string]*string) *GetTranslateFileJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetTranslateFileJobResultResponse) SetBody(v *GetTranslateFileJobResultResponseBody) *GetTranslateFileJobResultResponse {
	s.Body = v
	return s
}

type QueryUserManagementResourcesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserManagementResourcesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserManagementResourcesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserManagementResourcesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserManagementResourcesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserManagementResourcesResponseBody struct {
	// 资源列表
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s QueryUserManagementResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesResponseBody) SetResourceIds(v []*string) *QueryUserManagementResourcesResponseBody {
	s.ResourceIds = v
	return s
}

type QueryUserManagementResourcesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserManagementResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserManagementResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserManagementResourcesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserManagementResourcesResponse) SetHeaders(v map[string]*string) *QueryUserManagementResourcesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserManagementResourcesResponse) SetBody(v *QueryUserManagementResourcesResponseBody) *QueryUserManagementResourcesResponse {
	s.Body = v
	return s
}

type DeleteContactHideSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteContactHideSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteContactHideSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteContactHideSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteContactHideSettingHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteContactHideSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteContactHideSettingResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteContactHideSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactHideSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactHideSettingResponse) SetHeaders(v map[string]*string) *DeleteContactHideSettingResponse {
	s.Headers = v
	return s
}

type UpdateUserOwnnessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUserOwnnessHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserOwnnessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserOwnnessHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUserOwnnessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUserOwnnessRequest struct {
	// 状态类型
	OwnenssType *int32 `json:"ownenssType,omitempty" xml:"ownenssType,omitempty"`
	// 业务标志id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 开始时间戳（毫秒）
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间戳（毫秒）
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 删除标记
	DeletedFlag *int32 `json:"deletedFlag,omitempty" xml:"deletedFlag,omitempty"`
}

func (s UpdateUserOwnnessRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessRequest) SetOwnenssType(v int32) *UpdateUserOwnnessRequest {
	s.OwnenssType = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetId(v int64) *UpdateUserOwnnessRequest {
	s.Id = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetStartTime(v int64) *UpdateUserOwnnessRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetEndTime(v int64) *UpdateUserOwnnessRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateUserOwnnessRequest) SetDeletedFlag(v int32) *UpdateUserOwnnessRequest {
	s.DeletedFlag = &v
	return s
}

type UpdateUserOwnnessResponseBody struct {
	// 业务标识id
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateUserOwnnessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessResponseBody) SetResult(v string) *UpdateUserOwnnessResponseBody {
	s.Result = &v
	return s
}

type UpdateUserOwnnessResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserOwnnessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserOwnnessResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserOwnnessResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserOwnnessResponse) SetHeaders(v map[string]*string) *UpdateUserOwnnessResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserOwnnessResponse) SetBody(v *UpdateUserOwnnessResponseBody) *UpdateUserOwnnessResponse {
	s.Body = v
	return s
}

type GetMigrationUnionIdByUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetMigrationUnionIdByUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMigrationUnionIdByUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetMigrationUnionIdByUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMigrationUnionIdByUnionIdRequest struct {
	// unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdRequest) SetUnionId(v string) *GetMigrationUnionIdByUnionIdRequest {
	s.UnionId = &v
	return s
}

type GetMigrationUnionIdByUnionIdResponseBody struct {
	// migrationUnionId
	MigrationUnionId *string `json:"migrationUnionId,omitempty" xml:"migrationUnionId,omitempty"`
}

func (s GetMigrationUnionIdByUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdResponseBody) SetMigrationUnionId(v string) *GetMigrationUnionIdByUnionIdResponseBody {
	s.MigrationUnionId = &v
	return s
}

type GetMigrationUnionIdByUnionIdResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMigrationUnionIdByUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMigrationUnionIdByUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationUnionIdByUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationUnionIdByUnionIdResponse) SetHeaders(v map[string]*string) *GetMigrationUnionIdByUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationUnionIdByUnionIdResponse) SetBody(v *GetMigrationUnionIdByUnionIdResponseBody) *GetMigrationUnionIdByUnionIdResponse {
	s.Body = v
	return s
}

type GetDingIdByMigrationDingIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingIdByMigrationDingIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdHeaders) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdHeaders) SetCommonHeaders(v map[string]*string) *GetDingIdByMigrationDingIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingIdByMigrationDingIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingIdByMigrationDingIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingIdByMigrationDingIdRequest struct {
	// migrationDingId
	MigrationDingId *string `json:"migrationDingId,omitempty" xml:"migrationDingId,omitempty"`
}

func (s GetDingIdByMigrationDingIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdRequest) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdRequest) SetMigrationDingId(v string) *GetDingIdByMigrationDingIdRequest {
	s.MigrationDingId = &v
	return s
}

type GetDingIdByMigrationDingIdResponseBody struct {
	// dingId
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
}

func (s GetDingIdByMigrationDingIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdResponseBody) SetDingId(v string) *GetDingIdByMigrationDingIdResponseBody {
	s.DingId = &v
	return s
}

type GetDingIdByMigrationDingIdResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDingIdByMigrationDingIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingIdByMigrationDingIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingIdByMigrationDingIdResponse) GoString() string {
	return s.String()
}

func (s *GetDingIdByMigrationDingIdResponse) SetHeaders(v map[string]*string) *GetDingIdByMigrationDingIdResponse {
	s.Headers = v
	return s
}

func (s *GetDingIdByMigrationDingIdResponse) SetBody(v *GetDingIdByMigrationDingIdResponseBody) *GetDingIdByMigrationDingIdResponse {
	s.Body = v
	return s
}

type GetCooperateOrgInviteInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCooperateOrgInviteInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoHeaders) SetCommonHeaders(v map[string]*string) *GetCooperateOrgInviteInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCooperateOrgInviteInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetCooperateOrgInviteInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCooperateOrgInviteInfoResponseBody struct {
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
}

func (s GetCooperateOrgInviteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoResponseBody) SetInviteUrl(v string) *GetCooperateOrgInviteInfoResponseBody {
	s.InviteUrl = &v
	return s
}

type GetCooperateOrgInviteInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCooperateOrgInviteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCooperateOrgInviteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCooperateOrgInviteInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCooperateOrgInviteInfoResponse) SetHeaders(v map[string]*string) *GetCooperateOrgInviteInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCooperateOrgInviteInfoResponse) SetBody(v *GetCooperateOrgInviteInfoResponseBody) *GetCooperateOrgInviteInfoResponse {
	s.Body = v
	return s
}

type UpdateManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateManagementGroupRequest struct {
	// 管理组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 管理组成员
	Members []*UpdateManagementGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Scope   *UpdateManagementGroupRequestScope     `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// 资源列表
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s UpdateManagementGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequest) SetGroupName(v string) *UpdateManagementGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateManagementGroupRequest) SetMembers(v []*UpdateManagementGroupRequestMembers) *UpdateManagementGroupRequest {
	s.Members = v
	return s
}

func (s *UpdateManagementGroupRequest) SetScope(v *UpdateManagementGroupRequestScope) *UpdateManagementGroupRequest {
	s.Scope = v
	return s
}

func (s *UpdateManagementGroupRequest) SetResourceIds(v []*string) *UpdateManagementGroupRequest {
	s.ResourceIds = v
	return s
}

type UpdateManagementGroupRequestMembers struct {
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
}

func (s UpdateManagementGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequestMembers) SetMemberType(v string) *UpdateManagementGroupRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateManagementGroupRequestMembers) SetMemberId(v string) *UpdateManagementGroupRequestMembers {
	s.MemberId = &v
	return s
}

type UpdateManagementGroupRequestScope struct {
	// 范围类型
	ScopeType *int32 `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门列表，只在scopeType=3 生效
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
}

func (s UpdateManagementGroupRequestScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupRequestScope) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupRequestScope) SetScopeType(v int32) *UpdateManagementGroupRequestScope {
	s.ScopeType = &v
	return s
}

func (s *UpdateManagementGroupRequestScope) SetDeptIds(v []*int64) *UpdateManagementGroupRequestScope {
	s.DeptIds = v
	return s
}

type UpdateManagementGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateManagementGroupResponse) SetHeaders(v map[string]*string) *UpdateManagementGroupResponse {
	s.Headers = v
	return s
}

type GetBranchAuthDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBranchAuthDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataHeaders) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataHeaders) SetCommonHeaders(v map[string]*string) *GetBranchAuthDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBranchAuthDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetBranchAuthDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBranchAuthDataRequest struct {
	// 分支组织corpId
	BranchCorpId *string `json:"branchCorpId,omitempty" xml:"branchCorpId,omitempty"`
	// 数据编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 查询条件
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBranchAuthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataRequest) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataRequest) SetBranchCorpId(v string) *GetBranchAuthDataRequest {
	s.BranchCorpId = &v
	return s
}

func (s *GetBranchAuthDataRequest) SetCode(v string) *GetBranchAuthDataRequest {
	s.Code = &v
	return s
}

func (s *GetBranchAuthDataRequest) SetBody(v map[string]*string) *GetBranchAuthDataRequest {
	s.Body = v
	return s
}

type GetBranchAuthDataResponseBody struct {
	// result
	Result []*GetBranchAuthDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetBranchAuthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponseBody) SetResult(v []*GetBranchAuthDataResponseBodyResult) *GetBranchAuthDataResponseBody {
	s.Result = v
	return s
}

type GetBranchAuthDataResponseBodyResult struct {
	// 字段code
	FieldCode *string `json:"fieldCode,omitempty" xml:"fieldCode,omitempty"`
	// 字段名称
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// 字段值
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s GetBranchAuthDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldCode(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldCode = &v
	return s
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldName(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldName = &v
	return s
}

func (s *GetBranchAuthDataResponseBodyResult) SetFieldValue(v string) *GetBranchAuthDataResponseBodyResult {
	s.FieldValue = &v
	return s
}

type GetBranchAuthDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBranchAuthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBranchAuthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBranchAuthDataResponse) GoString() string {
	return s.String()
}

func (s *GetBranchAuthDataResponse) SetHeaders(v map[string]*string) *GetBranchAuthDataResponse {
	s.Headers = v
	return s
}

func (s *GetBranchAuthDataResponse) SetBody(v *GetBranchAuthDataResponseBody) *GetBranchAuthDataResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserResponseBody struct {
	// 昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 头像url
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// openId
	OpenId *string `json:"openId,omitempty" xml:"openId,omitempty"`
	// unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 个人邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 手机号对应的国家号
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetNick(v string) *GetUserResponseBody {
	s.Nick = &v
	return s
}

func (s *GetUserResponseBody) SetAvatarUrl(v string) *GetUserResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserResponseBody) SetMobile(v string) *GetUserResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBody) SetOpenId(v string) *GetUserResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetUserResponseBody) SetUnionId(v string) *GetUserResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetStateCode(v string) *GetUserResponseBody {
	s.StateCode = &v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type QueryResourceManagementMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryResourceManagementMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersHeaders) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersHeaders) SetCommonHeaders(v map[string]*string) *QueryResourceManagementMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryResourceManagementMembersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryResourceManagementMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryResourceManagementMembersResponseBody struct {
	// 可管理资源的成员
	Members []*QueryResourceManagementMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s QueryResourceManagementMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponseBody) SetMembers(v []*QueryResourceManagementMembersResponseBodyMembers) *QueryResourceManagementMembersResponseBody {
	s.Members = v
	return s
}

type QueryResourceManagementMembersResponseBodyMembers struct {
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
}

func (s QueryResourceManagementMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponseBodyMembers) SetMemberType(v string) *QueryResourceManagementMembersResponseBodyMembers {
	s.MemberType = &v
	return s
}

func (s *QueryResourceManagementMembersResponseBodyMembers) SetMemberId(v string) *QueryResourceManagementMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

type QueryResourceManagementMembersResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryResourceManagementMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResourceManagementMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourceManagementMembersResponse) GoString() string {
	return s.String()
}

func (s *QueryResourceManagementMembersResponse) SetHeaders(v map[string]*string) *QueryResourceManagementMembersResponse {
	s.Headers = v
	return s
}

func (s *QueryResourceManagementMembersResponse) SetBody(v *QueryResourceManagementMembersResponseBody) *QueryResourceManagementMembersResponse {
	s.Body = v
	return s
}

type TranslateFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TranslateFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileHeaders) GoString() string {
	return s.String()
}

func (s *TranslateFileHeaders) SetCommonHeaders(v map[string]*string) *TranslateFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TranslateFileHeaders) SetXAcsDingtalkAccessToken(v string) *TranslateFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TranslateFileRequest struct {
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 钉盘mediaId，#号开头。可以通过单步上传api获取
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 转译后文件名（含扩展名）
	OutputFileName *string `json:"outputFileName,omitempty" xml:"outputFileName,omitempty"`
	// unionId
	UnionId         *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EagleEyeTraceId *string `json:"eagleEyeTraceId,omitempty" xml:"eagleEyeTraceId,omitempty"`
}

func (s TranslateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileRequest) GoString() string {
	return s.String()
}

func (s *TranslateFileRequest) SetDingTokenGrantType(v int64) *TranslateFileRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *TranslateFileRequest) SetDingOrgId(v int64) *TranslateFileRequest {
	s.DingOrgId = &v
	return s
}

func (s *TranslateFileRequest) SetDingIsvOrgId(v int64) *TranslateFileRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *TranslateFileRequest) SetDingSuiteKey(v string) *TranslateFileRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *TranslateFileRequest) SetMediaId(v string) *TranslateFileRequest {
	s.MediaId = &v
	return s
}

func (s *TranslateFileRequest) SetOutputFileName(v string) *TranslateFileRequest {
	s.OutputFileName = &v
	return s
}

func (s *TranslateFileRequest) SetUnionId(v string) *TranslateFileRequest {
	s.UnionId = &v
	return s
}

func (s *TranslateFileRequest) SetRequestId(v string) *TranslateFileRequest {
	s.RequestId = &v
	return s
}

func (s *TranslateFileRequest) SetEagleEyeTraceId(v string) *TranslateFileRequest {
	s.EagleEyeTraceId = &v
	return s
}

type TranslateFileResponseBody struct {
	// 异步转译任务id，最大长度为64字符
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s TranslateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateFileResponseBody) SetJobId(v string) *TranslateFileResponseBody {
	s.JobId = &v
	return s
}

type TranslateFileResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateFileResponse) GoString() string {
	return s.String()
}

func (s *TranslateFileResponse) SetHeaders(v map[string]*string) *TranslateFileResponse {
	s.Headers = v
	return s
}

func (s *TranslateFileResponse) SetBody(v *TranslateFileResponseBody) *TranslateFileResponse {
	s.Body = v
	return s
}

type UpdateSeniorSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSeniorSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateSeniorSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSeniorSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSeniorSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSeniorSettingRequest struct {
	SeniorStaffId  *string   `json:"seniorStaffId,omitempty" xml:"seniorStaffId,omitempty"`
	Open           *bool     `json:"open,omitempty" xml:"open,omitempty"`
	PermitStaffIds []*string `json:"permitStaffIds,omitempty" xml:"permitStaffIds,omitempty" type:"Repeated"`
	PermitDeptIds  []*int64  `json:"permitDeptIds,omitempty" xml:"permitDeptIds,omitempty" type:"Repeated"`
	PermitTagIds   []*int64  `json:"permitTagIds,omitempty" xml:"permitTagIds,omitempty" type:"Repeated"`
	ProtectScenes  []*string `json:"protectScenes,omitempty" xml:"protectScenes,omitempty" type:"Repeated"`
}

func (s UpdateSeniorSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingRequest) SetSeniorStaffId(v string) *UpdateSeniorSettingRequest {
	s.SeniorStaffId = &v
	return s
}

func (s *UpdateSeniorSettingRequest) SetOpen(v bool) *UpdateSeniorSettingRequest {
	s.Open = &v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitStaffIds(v []*string) *UpdateSeniorSettingRequest {
	s.PermitStaffIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitDeptIds(v []*int64) *UpdateSeniorSettingRequest {
	s.PermitDeptIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetPermitTagIds(v []*int64) *UpdateSeniorSettingRequest {
	s.PermitTagIds = v
	return s
}

func (s *UpdateSeniorSettingRequest) SetProtectScenes(v []*string) *UpdateSeniorSettingRequest {
	s.ProtectScenes = v
	return s
}

type UpdateSeniorSettingResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateSeniorSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSeniorSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSeniorSettingResponse) SetHeaders(v map[string]*string) *UpdateSeniorSettingResponse {
	s.Headers = v
	return s
}

type GetApplyInviteInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApplyInviteInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoHeaders) SetCommonHeaders(v map[string]*string) *GetApplyInviteInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplyInviteInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetApplyInviteInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApplyInviteInfoRequest struct {
	// 邀请者userId
	InviterUserId *string `json:"inviterUserId,omitempty" xml:"inviterUserId,omitempty"`
	// 获取部门邀请链接的部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s GetApplyInviteInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoRequest) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoRequest) SetInviterUserId(v string) *GetApplyInviteInfoRequest {
	s.InviterUserId = &v
	return s
}

func (s *GetApplyInviteInfoRequest) SetDeptId(v int64) *GetApplyInviteInfoRequest {
	s.DeptId = &v
	return s
}

type GetApplyInviteInfoResponseBody struct {
	// 是否开启邀请
	InviteSwitch *bool `json:"inviteSwitch,omitempty" xml:"inviteSwitch,omitempty"`
	// 是否开启通过企业名称搜索申请
	SearchNameInvite *bool `json:"searchNameInvite,omitempty" xml:"searchNameInvite,omitempty"`
	// 是否开启通过团队号申请加入
	OrgApplyCodeInvite *bool `json:"orgApplyCodeInvite,omitempty" xml:"orgApplyCodeInvite,omitempty"`
	// 是否开启通过链接邀请加入
	LinkInvite *bool `json:"linkInvite,omitempty" xml:"linkInvite,omitempty"`
	// 邀请链接
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
	// 仅部门邀请有效： 0-无需审核 1-有权限的子管理员
	AuditType *int64 `json:"auditType,omitempty" xml:"auditType,omitempty"`
	// 是否允许员工扫码加入部门，仅在无需审核的时候有效（已经在组织内的成员通过扫描部门二维码加入某个部门）
	EmpApplyJoinDept *bool `json:"empApplyJoinDept,omitempty" xml:"empApplyJoinDept,omitempty"`
}

func (s GetApplyInviteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponseBody) SetInviteSwitch(v bool) *GetApplyInviteInfoResponseBody {
	s.InviteSwitch = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetSearchNameInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.SearchNameInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetOrgApplyCodeInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.OrgApplyCodeInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetLinkInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.LinkInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetInviteUrl(v string) *GetApplyInviteInfoResponseBody {
	s.InviteUrl = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetAuditType(v int64) *GetApplyInviteInfoResponseBody {
	s.AuditType = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetEmpApplyJoinDept(v bool) *GetApplyInviteInfoResponseBody {
	s.EmpApplyJoinDept = &v
	return s
}

type GetApplyInviteInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplyInviteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplyInviteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponse) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponse) SetHeaders(v map[string]*string) *GetApplyInviteInfoResponse {
	s.Headers = v
	return s
}

func (s *GetApplyInviteInfoResponse) SetBody(v *GetApplyInviteInfoResponseBody) *GetApplyInviteInfoResponse {
	s.Body = v
	return s
}

type CreateCooperateOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCooperateOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgHeaders) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgHeaders) SetCommonHeaders(v map[string]*string) *CreateCooperateOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCooperateOrgHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCooperateOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCooperateOrgRequest struct {
	// 合作空间组织名称
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// 合作空间的logo
	LogoMediaId *string `json:"logoMediaId,omitempty" xml:"logoMediaId,omitempty"`
	// 行业code
	IndustryCode *int64 `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
}

func (s CreateCooperateOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgRequest) SetOrgName(v string) *CreateCooperateOrgRequest {
	s.OrgName = &v
	return s
}

func (s *CreateCooperateOrgRequest) SetLogoMediaId(v string) *CreateCooperateOrgRequest {
	s.LogoMediaId = &v
	return s
}

func (s *CreateCooperateOrgRequest) SetIndustryCode(v int64) *CreateCooperateOrgRequest {
	s.IndustryCode = &v
	return s
}

type CreateCooperateOrgResponseBody struct {
	// result
	CooperateCorpId *string `json:"cooperateCorpId,omitempty" xml:"cooperateCorpId,omitempty"`
}

func (s CreateCooperateOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgResponseBody) SetCooperateCorpId(v string) *CreateCooperateOrgResponseBody {
	s.CooperateCorpId = &v
	return s
}

type CreateCooperateOrgResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCooperateOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCooperateOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCooperateOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateCooperateOrgResponse) SetHeaders(v map[string]*string) *CreateCooperateOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateCooperateOrgResponse) SetBody(v *CreateCooperateOrgResponseBody) *CreateCooperateOrgResponse {
	s.Body = v
	return s
}

type GetMigrationDingIdByDingIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMigrationDingIdByDingIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdHeaders) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdHeaders) SetCommonHeaders(v map[string]*string) *GetMigrationDingIdByDingIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMigrationDingIdByDingIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetMigrationDingIdByDingIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMigrationDingIdByDingIdRequest struct {
	// dingId
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
}

func (s GetMigrationDingIdByDingIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdRequest) SetDingId(v string) *GetMigrationDingIdByDingIdRequest {
	s.DingId = &v
	return s
}

type GetMigrationDingIdByDingIdResponseBody struct {
	// migrationDingId
	MigrationDingId *string `json:"migrationDingId,omitempty" xml:"migrationDingId,omitempty"`
}

func (s GetMigrationDingIdByDingIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdResponseBody) SetMigrationDingId(v string) *GetMigrationDingIdByDingIdResponseBody {
	s.MigrationDingId = &v
	return s
}

type GetMigrationDingIdByDingIdResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMigrationDingIdByDingIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMigrationDingIdByDingIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMigrationDingIdByDingIdResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationDingIdByDingIdResponse) SetHeaders(v map[string]*string) *GetMigrationDingIdByDingIdResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationDingIdByDingIdResponse) SetBody(v *GetMigrationDingIdByDingIdResponseBody) *GetMigrationDingIdByDingIdResponse {
	s.Body = v
	return s
}

type CreateManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateManagementGroupRequest struct {
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 管理组成员
	Members []*CreateManagementGroupRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 管理范围
	Scope *CreateManagementGroupRequestScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// 资源列表
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
}

func (s CreateManagementGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequest) SetGroupName(v string) *CreateManagementGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateManagementGroupRequest) SetMembers(v []*CreateManagementGroupRequestMembers) *CreateManagementGroupRequest {
	s.Members = v
	return s
}

func (s *CreateManagementGroupRequest) SetScope(v *CreateManagementGroupRequestScope) *CreateManagementGroupRequest {
	s.Scope = v
	return s
}

func (s *CreateManagementGroupRequest) SetResourceIds(v []*string) *CreateManagementGroupRequest {
	s.ResourceIds = v
	return s
}

type CreateManagementGroupRequestMembers struct {
	// 成员类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 成员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
}

func (s CreateManagementGroupRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequestMembers) SetMemberType(v string) *CreateManagementGroupRequestMembers {
	s.MemberType = &v
	return s
}

func (s *CreateManagementGroupRequestMembers) SetMemberId(v string) *CreateManagementGroupRequestMembers {
	s.MemberId = &v
	return s
}

type CreateManagementGroupRequestScope struct {
	// 范围类型
	ScopeType *int32 `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门列表，只在scopeType=3 生效
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
}

func (s CreateManagementGroupRequestScope) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupRequestScope) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupRequestScope) SetScopeType(v int32) *CreateManagementGroupRequestScope {
	s.ScopeType = &v
	return s
}

func (s *CreateManagementGroupRequestScope) SetDeptIds(v []*int64) *CreateManagementGroupRequestScope {
	s.DeptIds = v
	return s
}

type CreateManagementGroupResponseBody struct {
	// 返回管理组groupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CreateManagementGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupResponseBody) SetGroupId(v string) *CreateManagementGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateManagementGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateManagementGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateManagementGroupResponse) SetHeaders(v map[string]*string) *CreateManagementGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateManagementGroupResponse) SetBody(v *CreateManagementGroupResponseBody) *CreateManagementGroupResponse {
	s.Body = v
	return s
}

type DeleteManagementGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteManagementGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteManagementGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteManagementGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteManagementGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteManagementGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteManagementGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteManagementGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteManagementGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteManagementGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteManagementGroupResponse) SetHeaders(v map[string]*string) *DeleteManagementGroupResponse {
	s.Headers = v
	return s
}

type TransformToExclusiveAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransformToExclusiveAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountHeaders) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountHeaders) SetCommonHeaders(v map[string]*string) *TransformToExclusiveAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransformToExclusiveAccountHeaders) SetXAcsDingtalkAccessToken(v string) *TransformToExclusiveAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransformToExclusiveAccountRequest struct {
	// transformType
	TransformType *string `json:"transformType,omitempty" xml:"transformType,omitempty"`
	// idpDingTalk
	IdpDingTalk *bool `json:"idpDingTalk,omitempty" xml:"idpDingTalk,omitempty"`
	// loginId
	LoginId *string `json:"loginId,omitempty" xml:"loginId,omitempty"`
	// initPassword
	InitPassword *string `json:"initPassword,omitempty" xml:"initPassword,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TransformToExclusiveAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountRequest) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountRequest) SetTransformType(v string) *TransformToExclusiveAccountRequest {
	s.TransformType = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetIdpDingTalk(v bool) *TransformToExclusiveAccountRequest {
	s.IdpDingTalk = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetLoginId(v string) *TransformToExclusiveAccountRequest {
	s.LoginId = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetInitPassword(v string) *TransformToExclusiveAccountRequest {
	s.InitPassword = &v
	return s
}

func (s *TransformToExclusiveAccountRequest) SetUserId(v string) *TransformToExclusiveAccountRequest {
	s.UserId = &v
	return s
}

type TransformToExclusiveAccountResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TransformToExclusiveAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformToExclusiveAccountResponse) GoString() string {
	return s.String()
}

func (s *TransformToExclusiveAccountResponse) SetHeaders(v map[string]*string) *TransformToExclusiveAccountResponse {
	s.Headers = v
	return s
}

type GetUnionIdByMigrationUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUnionIdByMigrationUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnionIdByMigrationUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUnionIdByMigrationUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUnionIdByMigrationUnionIdRequest struct {
	// migrationUnionId
	MigrationUnionId *string `json:"migrationUnionId,omitempty" xml:"migrationUnionId,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdRequest) SetMigrationUnionId(v string) *GetUnionIdByMigrationUnionIdRequest {
	s.MigrationUnionId = &v
	return s
}

type GetUnionIdByMigrationUnionIdResponseBody struct {
	// unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUnionIdByMigrationUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdResponseBody) SetUnionId(v string) *GetUnionIdByMigrationUnionIdResponseBody {
	s.UnionId = &v
	return s
}

type GetUnionIdByMigrationUnionIdResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUnionIdByMigrationUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUnionIdByMigrationUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnionIdByMigrationUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUnionIdByMigrationUnionIdResponse) SetHeaders(v map[string]*string) *GetUnionIdByMigrationUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUnionIdByMigrationUnionIdResponse) SetBody(v *GetUnionIdByMigrationUnionIdResponseBody) *GetUnionIdByMigrationUnionIdResponse {
	s.Body = v
	return s
}

type GetLatestDingIndexHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLatestDingIndexHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexHeaders) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexHeaders) SetCommonHeaders(v map[string]*string) *GetLatestDingIndexHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLatestDingIndexHeaders) SetXAcsDingtalkAccessToken(v string) *GetLatestDingIndexHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLatestDingIndexResponseBody struct {
	// 日期
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 钉钉指数
	IdxTotal *float32 `json:"idxTotal,omitempty" xml:"idxTotal,omitempty"`
	// 效率指数
	IdxEfficiency *float32 `json:"idxEfficiency,omitempty" xml:"idxEfficiency,omitempty"`
	// 绿色指数
	IdxCarbon *float32 `json:"idxCarbon,omitempty" xml:"idxCarbon,omitempty"`
	// 钉钉指数月均分
	IdxMonthlyAvg *float32 `json:"idxMonthlyAvg,omitempty" xml:"idxMonthlyAvg,omitempty"`
}

func (s GetLatestDingIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponseBody) SetStatDate(v string) *GetLatestDingIndexResponseBody {
	s.StatDate = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxTotal(v float32) *GetLatestDingIndexResponseBody {
	s.IdxTotal = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxEfficiency(v float32) *GetLatestDingIndexResponseBody {
	s.IdxEfficiency = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxCarbon(v float32) *GetLatestDingIndexResponseBody {
	s.IdxCarbon = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxMonthlyAvg(v float32) *GetLatestDingIndexResponseBody {
	s.IdxMonthlyAvg = &v
	return s
}

type GetLatestDingIndexResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLatestDingIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLatestDingIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponse) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponse) SetHeaders(v map[string]*string) *GetLatestDingIndexResponse {
	s.Headers = v
	return s
}

func (s *GetLatestDingIndexResponse) SetBody(v *GetLatestDingIndexResponseBody) *GetLatestDingIndexResponse {
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

func (client *Client) SortUser(request *SortUserRequest) (_result *SortUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SortUserHeaders{}
	_result = &SortUserResponse{}
	_body, _err := client.SortUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SortUserWithOptions(request *SortUserRequest, headers *SortUserHeaders, runtime *util.RuntimeOptions) (_result *SortUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["sortType"] = request.SortType
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
	_result = &SortUserResponse{}
	_body, _err := client.DoROARequest(tea.String("SortUser"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/users/sort"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContactHideSettings(request *ListContactHideSettingsRequest) (_result *ListContactHideSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListContactHideSettingsHeaders{}
	_result = &ListContactHideSettingsResponse{}
	_body, _err := client.ListContactHideSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactHideSettingsWithOptions(request *ListContactHideSettingsRequest, headers *ListContactHideSettingsHeaders, runtime *util.RuntimeOptions) (_result *ListContactHideSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &ListContactHideSettingsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListContactHideSettings"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/contactHideSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContactHideSetting(request *UpdateContactHideSettingRequest) (_result *UpdateContactHideSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateContactHideSettingHeaders{}
	_result = &UpdateContactHideSettingResponse{}
	_body, _err := client.UpdateContactHideSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactHideSettingWithOptions(request *UpdateContactHideSettingRequest, headers *UpdateContactHideSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateContactHideSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectStaffIds)) {
		body["objectStaffIds"] = request.ObjectStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeStaffIds)) {
		body["excludeStaffIds"] = request.ExcludeStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
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
	_result = &UpdateContactHideSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateContactHideSetting"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/contact/contactHideSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEmpAttrbuteVisibilitySetting(request *UpdateEmpAttrbuteVisibilitySettingRequest) (_result *UpdateEmpAttrbuteVisibilitySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateEmpAttrbuteVisibilitySettingHeaders{}
	_result = &UpdateEmpAttrbuteVisibilitySettingResponse{}
	_body, _err := client.UpdateEmpAttrbuteVisibilitySettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEmpAttrbuteVisibilitySettingWithOptions(request *UpdateEmpAttrbuteVisibilitySettingRequest, headers *UpdateEmpAttrbuteVisibilitySettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateEmpAttrbuteVisibilitySettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectStaffIds)) {
		body["objectStaffIds"] = request.ObjectStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectDeptIds)) {
		body["objectDeptIds"] = request.ObjectDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectTagIds)) {
		body["objectTagIds"] = request.ObjectTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideFields)) {
		body["hideFields"] = request.HideFields
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeStaffIds)) {
		body["excludeStaffIds"] = request.ExcludeStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeDeptIds)) {
		body["excludeDeptIds"] = request.ExcludeDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeTagIds)) {
		body["excludeTagIds"] = request.ExcludeTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["active"] = request.Active
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
	_result = &UpdateEmpAttrbuteVisibilitySettingResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateEmpAttrbuteVisibilitySetting"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/staffAttributes/visibilitySettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEmpAttributeVisibility(settingId *string) (_result *DeleteEmpAttributeVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteEmpAttributeVisibilityHeaders{}
	_result = &DeleteEmpAttributeVisibilityResponse{}
	_body, _err := client.DeleteEmpAttributeVisibilityWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEmpAttributeVisibilityWithOptions(settingId *string, headers *DeleteEmpAttributeVisibilityHeaders, runtime *util.RuntimeOptions) (_result *DeleteEmpAttributeVisibilityResponse, _err error) {
	settingId = openapiutil.GetEncodeParam(settingId)
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
	_result = &DeleteEmpAttributeVisibilityResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEmpAttributeVisibility"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/contact/staffAttributes/visibilitySettings/"+tea.StringValue(settingId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchDepartment(request *SearchDepartmentRequest) (_result *SearchDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchDepartmentHeaders{}
	_result = &SearchDepartmentResponse{}
	_body, _err := client.SearchDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchDepartmentWithOptions(request *SearchDepartmentRequest, headers *SearchDepartmentHeaders, runtime *util.RuntimeOptions) (_result *SearchDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryWord)) {
		body["queryWord"] = request.QueryWord
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
	_result = &SearchDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchDepartment"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/departments/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListManagementGroups(request *ListManagementGroupsRequest) (_result *ListManagementGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListManagementGroupsHeaders{}
	_result = &ListManagementGroupsResponse{}
	_body, _err := client.ListManagementGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListManagementGroupsWithOptions(request *ListManagementGroupsRequest, headers *ListManagementGroupsHeaders, runtime *util.RuntimeOptions) (_result *ListManagementGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &ListManagementGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListManagementGroups"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/managementGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSeniorSettings(request *ListSeniorSettingsRequest) (_result *ListSeniorSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSeniorSettingsHeaders{}
	_result = &ListSeniorSettingsResponse{}
	_body, _err := client.ListSeniorSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSeniorSettingsWithOptions(request *ListSeniorSettingsRequest, headers *ListSeniorSettingsHeaders, runtime *util.RuntimeOptions) (_result *ListSeniorSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SeniorStaffId)) {
		query["seniorStaffId"] = request.SeniorStaffId
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
	_result = &ListSeniorSettingsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSeniorSettings"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/seniorSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEmpAttributeVisibility(request *ListEmpAttributeVisibilityRequest) (_result *ListEmpAttributeVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEmpAttributeVisibilityHeaders{}
	_result = &ListEmpAttributeVisibilityResponse{}
	_body, _err := client.ListEmpAttributeVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEmpAttributeVisibilityWithOptions(request *ListEmpAttributeVisibilityRequest, headers *ListEmpAttributeVisibilityHeaders, runtime *util.RuntimeOptions) (_result *ListEmpAttributeVisibilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &ListEmpAttributeVisibilityResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEmpAttributeVisibility"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/staffAttributes/visibilitySettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchUser(request *SearchUserRequest) (_result *SearchUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchUserHeaders{}
	_result = &SearchUserResponse{}
	_body, _err := client.SearchUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchUserWithOptions(request *SearchUserRequest, headers *SearchUserHeaders, runtime *util.RuntimeOptions) (_result *SearchUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryWord)) {
		body["queryWord"] = request.QueryWord
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
	_result = &SearchUserResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchUser"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/users/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranslateFileJobResult(request *GetTranslateFileJobResultRequest) (_result *GetTranslateFileJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTranslateFileJobResultHeaders{}
	_result = &GetTranslateFileJobResultResponse{}
	_body, _err := client.GetTranslateFileJobResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranslateFileJobResultWithOptions(request *GetTranslateFileJobResultRequest, headers *GetTranslateFileJobResultHeaders, runtime *util.RuntimeOptions) (_result *GetTranslateFileJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["jobId"] = request.JobId
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
	_result = &GetTranslateFileJobResultResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTranslateFileJobResult"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/files/translateResults"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserManagementResources(userId *string) (_result *QueryUserManagementResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserManagementResourcesHeaders{}
	_result = &QueryUserManagementResourcesResponse{}
	_body, _err := client.QueryUserManagementResourcesWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserManagementResourcesWithOptions(userId *string, headers *QueryUserManagementResourcesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserManagementResourcesResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserManagementResourcesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserManagementResources"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/users/"+tea.StringValue(userId)+"/managemementResources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactHideSetting(settingId *string) (_result *DeleteContactHideSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteContactHideSettingHeaders{}
	_result = &DeleteContactHideSettingResponse{}
	_body, _err := client.DeleteContactHideSettingWithOptions(settingId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactHideSettingWithOptions(settingId *string, headers *DeleteContactHideSettingHeaders, runtime *util.RuntimeOptions) (_result *DeleteContactHideSettingResponse, _err error) {
	settingId = openapiutil.GetEncodeParam(settingId)
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
	_result = &DeleteContactHideSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteContactHideSetting"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/contact/contactHideSettings/"+tea.StringValue(settingId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserOwnness(userId *string, request *UpdateUserOwnnessRequest) (_result *UpdateUserOwnnessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserOwnnessHeaders{}
	_result = &UpdateUserOwnnessResponse{}
	_body, _err := client.UpdateUserOwnnessWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserOwnnessWithOptions(userId *string, request *UpdateUserOwnnessRequest, headers *UpdateUserOwnnessHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserOwnnessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnenssType)) {
		body["ownenssType"] = request.OwnenssType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeletedFlag)) {
		body["deletedFlag"] = request.DeletedFlag
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
	_result = &UpdateUserOwnnessResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateUserOwnness"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/contact/user/"+tea.StringValue(userId)+"/ownness"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMigrationUnionIdByUnionId(request *GetMigrationUnionIdByUnionIdRequest) (_result *GetMigrationUnionIdByUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMigrationUnionIdByUnionIdHeaders{}
	_result = &GetMigrationUnionIdByUnionIdResponse{}
	_body, _err := client.GetMigrationUnionIdByUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMigrationUnionIdByUnionIdWithOptions(request *GetMigrationUnionIdByUnionIdRequest, headers *GetMigrationUnionIdByUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetMigrationUnionIdByUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
	_result = &GetMigrationUnionIdByUnionIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMigrationUnionIdByUnionId"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/orgAccount/getMigrationUnionIdByUnionIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDingIdByMigrationDingId(request *GetDingIdByMigrationDingIdRequest) (_result *GetDingIdByMigrationDingIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingIdByMigrationDingIdHeaders{}
	_result = &GetDingIdByMigrationDingIdResponse{}
	_body, _err := client.GetDingIdByMigrationDingIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingIdByMigrationDingIdWithOptions(request *GetDingIdByMigrationDingIdRequest, headers *GetDingIdByMigrationDingIdHeaders, runtime *util.RuntimeOptions) (_result *GetDingIdByMigrationDingIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationDingId)) {
		query["migrationDingId"] = request.MigrationDingId
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
	_result = &GetDingIdByMigrationDingIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDingIdByMigrationDingId"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/orgAccount/getDingIdByMigrationDingIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCooperateOrgInviteInfo(cooperateCorpId *string) (_result *GetCooperateOrgInviteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCooperateOrgInviteInfoHeaders{}
	_result = &GetCooperateOrgInviteInfoResponse{}
	_body, _err := client.GetCooperateOrgInviteInfoWithOptions(cooperateCorpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCooperateOrgInviteInfoWithOptions(cooperateCorpId *string, headers *GetCooperateOrgInviteInfoHeaders, runtime *util.RuntimeOptions) (_result *GetCooperateOrgInviteInfoResponse, _err error) {
	cooperateCorpId = openapiutil.GetEncodeParam(cooperateCorpId)
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
	_result = &GetCooperateOrgInviteInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCooperateOrgInviteInfo"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/cooperateCorps/"+tea.StringValue(cooperateCorpId)+"/inviteInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateManagementGroup(groupId *string, request *UpdateManagementGroupRequest) (_result *UpdateManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateManagementGroupHeaders{}
	_result = &UpdateManagementGroupResponse{}
	_body, _err := client.UpdateManagementGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateManagementGroupWithOptions(groupId *string, request *UpdateManagementGroupRequest, headers *UpdateManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateManagementGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	groupId = openapiutil.GetEncodeParam(groupId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Scope))) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
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
	_result = &UpdateManagementGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateManagementGroup"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/contact/managementGroups/"+tea.StringValue(groupId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBranchAuthData(request *GetBranchAuthDataRequest) (_result *GetBranchAuthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBranchAuthDataHeaders{}
	_result = &GetBranchAuthDataResponse{}
	_body, _err := client.GetBranchAuthDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBranchAuthDataWithOptions(request *GetBranchAuthDataRequest, headers *GetBranchAuthDataHeaders, runtime *util.RuntimeOptions) (_result *GetBranchAuthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchCorpId)) {
		query["branchCorpId"] = request.BranchCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	body := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body = request.Body
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetBranchAuthDataResponse{}
	_body, _err := client.DoROARequest(tea.String("GetBranchAuthData"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/branchAuthDatas/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(unionId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(unionId *string, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	unionId = openapiutil.GetEncodeParam(unionId)
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
	_result = &GetUserResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUser"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/users/"+tea.StringValue(unionId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResourceManagementMembers(resourceId *string) (_result *QueryResourceManagementMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryResourceManagementMembersHeaders{}
	_result = &QueryResourceManagementMembersResponse{}
	_body, _err := client.QueryResourceManagementMembersWithOptions(resourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResourceManagementMembersWithOptions(resourceId *string, headers *QueryResourceManagementMembersHeaders, runtime *util.RuntimeOptions) (_result *QueryResourceManagementMembersResponse, _err error) {
	resourceId = openapiutil.GetEncodeParam(resourceId)
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
	_result = &QueryResourceManagementMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryResourceManagementMembers"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/resources/"+tea.StringValue(resourceId)+"/managementMembers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateFile(request *TranslateFileRequest) (_result *TranslateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TranslateFileHeaders{}
	_result = &TranslateFileResponse{}
	_body, _err := client.TranslateFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateFileWithOptions(request *TranslateFileRequest, headers *TranslateFileHeaders, runtime *util.RuntimeOptions) (_result *TranslateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFileName)) {
		body["outputFileName"] = request.OutputFileName
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.EagleEyeTraceId)) {
		body["eagleEyeTraceId"] = request.EagleEyeTraceId
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
	_result = &TranslateFileResponse{}
	_body, _err := client.DoROARequest(tea.String("TranslateFile"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/files/translate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSeniorSetting(request *UpdateSeniorSettingRequest) (_result *UpdateSeniorSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSeniorSettingHeaders{}
	_result = &UpdateSeniorSettingResponse{}
	_body, _err := client.UpdateSeniorSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSeniorSettingWithOptions(request *UpdateSeniorSettingRequest, headers *UpdateSeniorSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateSeniorSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SeniorStaffId)) {
		body["seniorStaffId"] = request.SeniorStaffId
	}

	if !tea.BoolValue(util.IsUnset(request.Open)) {
		body["open"] = request.Open
	}

	if !tea.BoolValue(util.IsUnset(request.PermitStaffIds)) {
		body["permitStaffIds"] = request.PermitStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.PermitDeptIds)) {
		body["permitDeptIds"] = request.PermitDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.PermitTagIds)) {
		body["permitTagIds"] = request.PermitTagIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectScenes)) {
		body["protectScenes"] = request.ProtectScenes
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
	_result = &UpdateSeniorSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSeniorSetting"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/seniorSettings"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplyInviteInfo(request *GetApplyInviteInfoRequest) (_result *GetApplyInviteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplyInviteInfoHeaders{}
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.GetApplyInviteInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplyInviteInfoWithOptions(request *GetApplyInviteInfoRequest, headers *GetApplyInviteInfoHeaders, runtime *util.RuntimeOptions) (_result *GetApplyInviteInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviterUserId)) {
		query["inviterUserId"] = request.InviterUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApplyInviteInfo"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/invites/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCooperateOrg(request *CreateCooperateOrgRequest) (_result *CreateCooperateOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCooperateOrgHeaders{}
	_result = &CreateCooperateOrgResponse{}
	_body, _err := client.CreateCooperateOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCooperateOrgWithOptions(request *CreateCooperateOrgRequest, headers *CreateCooperateOrgHeaders, runtime *util.RuntimeOptions) (_result *CreateCooperateOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.LogoMediaId)) {
		body["logoMediaId"] = request.LogoMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.IndustryCode)) {
		body["industryCode"] = request.IndustryCode
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
	_result = &CreateCooperateOrgResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCooperateOrg"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/cooperateCorps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMigrationDingIdByDingId(request *GetMigrationDingIdByDingIdRequest) (_result *GetMigrationDingIdByDingIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMigrationDingIdByDingIdHeaders{}
	_result = &GetMigrationDingIdByDingIdResponse{}
	_body, _err := client.GetMigrationDingIdByDingIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMigrationDingIdByDingIdWithOptions(request *GetMigrationDingIdByDingIdRequest, headers *GetMigrationDingIdByDingIdHeaders, runtime *util.RuntimeOptions) (_result *GetMigrationDingIdByDingIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingId)) {
		query["dingId"] = request.DingId
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
	_result = &GetMigrationDingIdByDingIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMigrationDingIdByDingId"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/orgAccount/getMigrationDingIdByDingIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateManagementGroup(request *CreateManagementGroupRequest) (_result *CreateManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateManagementGroupHeaders{}
	_result = &CreateManagementGroupResponse{}
	_body, _err := client.CreateManagementGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateManagementGroupWithOptions(request *CreateManagementGroupRequest, headers *CreateManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateManagementGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Scope))) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
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
	_result = &CreateManagementGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateManagementGroup"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/managementGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteManagementGroup(groupId *string) (_result *DeleteManagementGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteManagementGroupHeaders{}
	_result = &DeleteManagementGroupResponse{}
	_body, _err := client.DeleteManagementGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteManagementGroupWithOptions(groupId *string, headers *DeleteManagementGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteManagementGroupResponse, _err error) {
	groupId = openapiutil.GetEncodeParam(groupId)
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
	_result = &DeleteManagementGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteManagementGroup"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/contact/managementGroups/"+tea.StringValue(groupId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransformToExclusiveAccount(request *TransformToExclusiveAccountRequest) (_result *TransformToExclusiveAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransformToExclusiveAccountHeaders{}
	_result = &TransformToExclusiveAccountResponse{}
	_body, _err := client.TransformToExclusiveAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransformToExclusiveAccountWithOptions(request *TransformToExclusiveAccountRequest, headers *TransformToExclusiveAccountHeaders, runtime *util.RuntimeOptions) (_result *TransformToExclusiveAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TransformType)) {
		body["transformType"] = request.TransformType
	}

	if !tea.BoolValue(util.IsUnset(request.IdpDingTalk)) {
		body["idpDingTalk"] = request.IdpDingTalk
	}

	if !tea.BoolValue(util.IsUnset(request.LoginId)) {
		body["loginId"] = request.LoginId
	}

	if !tea.BoolValue(util.IsUnset(request.InitPassword)) {
		body["initPassword"] = request.InitPassword
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &TransformToExclusiveAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("TransformToExclusiveAccount"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/contact/orgAccount/transformToExclusiveAccounts"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUnionIdByMigrationUnionId(request *GetUnionIdByMigrationUnionIdRequest) (_result *GetUnionIdByMigrationUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUnionIdByMigrationUnionIdHeaders{}
	_result = &GetUnionIdByMigrationUnionIdResponse{}
	_body, _err := client.GetUnionIdByMigrationUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUnionIdByMigrationUnionIdWithOptions(request *GetUnionIdByMigrationUnionIdRequest, headers *GetUnionIdByMigrationUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetUnionIdByMigrationUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationUnionId)) {
		query["migrationUnionId"] = request.MigrationUnionId
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
	_result = &GetUnionIdByMigrationUnionIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUnionIdByMigrationUnionId"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/orgAccount/getUnionIdByMigrationUnionIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLatestDingIndex() (_result *GetLatestDingIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLatestDingIndexHeaders{}
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.GetLatestDingIndexWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLatestDingIndexWithOptions(headers *GetLatestDingIndexHeaders, runtime *util.RuntimeOptions) (_result *GetLatestDingIndexResponse, _err error) {
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
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.DoROARequest(tea.String("GetLatestDingIndex"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/dingIndexs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
