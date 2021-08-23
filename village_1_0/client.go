// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package village_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListSubCorpsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubCorpsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsHeaders) GoString() string {
	return s.String()
}

func (s *ListSubCorpsHeaders) SetCommonHeaders(v map[string]*string) *ListSubCorpsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubCorpsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubCorpsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubCorpsRequest struct {
	// 下级指定组织层级列表，组织层级为county,town,community,residential，依次为区/县、乡/镇/街道、社区/村、小区，如果查多个用 '|' 分隔
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
	// 真正查询的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 是否查询直接下级
	IsOnlyDirect *bool `json:"isOnlyDirect,omitempty" xml:"isOnlyDirect,omitempty"`
}

func (s ListSubCorpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsRequest) GoString() string {
	return s.String()
}

func (s *ListSubCorpsRequest) SetTypes(v string) *ListSubCorpsRequest {
	s.Types = &v
	return s
}

func (s *ListSubCorpsRequest) SetSubCorpId(v string) *ListSubCorpsRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListSubCorpsRequest) SetIsOnlyDirect(v bool) *ListSubCorpsRequest {
	s.IsOnlyDirect = &v
	return s
}

type ListSubCorpsResponseBody struct {
	// result
	Result []*ListSubCorpsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSubCorpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponseBody) SetResult(v []*ListSubCorpsResponseBodyResult) *ListSubCorpsResponseBody {
	s.Result = v
	return s
}

type ListSubCorpsResponseBodyResult struct {
	// 企业corpid
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 企业名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 区域类型，值为county,town,community,residential，依次为区/县、乡/镇/街道、社区/村、小区
	RegionType *string `json:"regionType,omitempty" xml:"regionType,omitempty"`
	// 区域码
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 区域详细信息
	RegionLocation *string `json:"regionLocation,omitempty" xml:"regionLocation,omitempty"`
	// 企业行业码
	IndustryCode *int32 `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	// 企业行业名称
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s ListSubCorpsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponseBodyResult) SetCorpId(v string) *ListSubCorpsResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetName(v string) *ListSubCorpsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetRegionType(v string) *ListSubCorpsResponseBodyResult {
	s.RegionType = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetRegionId(v string) *ListSubCorpsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetRegionLocation(v string) *ListSubCorpsResponseBodyResult {
	s.RegionLocation = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetIndustryCode(v int32) *ListSubCorpsResponseBodyResult {
	s.IndustryCode = &v
	return s
}

func (s *ListSubCorpsResponseBodyResult) SetIndustry(v string) *ListSubCorpsResponseBodyResult {
	s.Industry = &v
	return s
}

type ListSubCorpsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubCorpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubCorpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponse) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponse) SetHeaders(v map[string]*string) *ListSubCorpsResponse {
	s.Headers = v
	return s
}

func (s *ListSubCorpsResponse) SetBody(v *ListSubCorpsResponseBody) *ListSubCorpsResponse {
	s.Body = v
	return s
}

type GetVillageOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetVillageOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *GetVillageOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetVillageOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetVillageOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetVillageOrgInfoResponseBody struct {
	// 区域类型
	RegionType *string `json:"regionType,omitempty" xml:"regionType,omitempty"`
	// 行政区ID
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 具体的企业区域位置信息
	RegionLocation *string `json:"regionLocation,omitempty" xml:"regionLocation,omitempty"`
}

func (s GetVillageOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoResponseBody) SetRegionType(v string) *GetVillageOrgInfoResponseBody {
	s.RegionType = &v
	return s
}

func (s *GetVillageOrgInfoResponseBody) SetRegionId(v string) *GetVillageOrgInfoResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVillageOrgInfoResponseBody) SetRegionLocation(v string) *GetVillageOrgInfoResponseBody {
	s.RegionLocation = &v
	return s
}

type GetVillageOrgInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVillageOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVillageOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoResponse) SetHeaders(v map[string]*string) *GetVillageOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVillageOrgInfoResponse) SetBody(v *GetVillageOrgInfoResponseBody) *GetVillageOrgInfoResponse {
	s.Body = v
	return s
}

type ListResidentDeptUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentDeptUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersHeaders) SetCommonHeaders(v map[string]*string) *ListResidentDeptUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentDeptUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentDeptUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentDeptUsersRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 标签
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 游标，不传默认1
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 大小
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListResidentDeptUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersRequest) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersRequest) SetSubCorpId(v string) *ListResidentDeptUsersRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetRole(v string) *ListResidentDeptUsersRequest {
	s.Role = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetCursor(v int64) *ListResidentDeptUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetSize(v int32) *ListResidentDeptUsersRequest {
	s.Size = &v
	return s
}

type ListResidentDeptUsersResponseBody struct {
	// 下一个游标
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 是否还有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页数据
	List []*ListResidentDeptUsersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListResidentDeptUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBody) SetNextCursor(v int64) *ListResidentDeptUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListResidentDeptUsersResponseBody) SetHasMore(v bool) *ListResidentDeptUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListResidentDeptUsersResponseBody) SetList(v []*ListResidentDeptUsersResponseBodyList) *ListResidentDeptUsersResponseBody {
	s.List = v
	return s
}

type ListResidentDeptUsersResponseBodyList struct {
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 员工名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签列表
	Roles []*ListResidentDeptUsersResponseBodyListRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	// 员工特征
	Feature *string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 钉钉唯一标识
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListResidentDeptUsersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBodyList) SetUserId(v string) *ListResidentDeptUsersResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyList) SetName(v string) *ListResidentDeptUsersResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyList) SetRoles(v []*ListResidentDeptUsersResponseBodyListRoles) *ListResidentDeptUsersResponseBodyList {
	s.Roles = v
	return s
}

func (s *ListResidentDeptUsersResponseBodyList) SetFeature(v string) *ListResidentDeptUsersResponseBodyList {
	s.Feature = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyList) SetUnionId(v string) *ListResidentDeptUsersResponseBodyList {
	s.UnionId = &v
	return s
}

type ListResidentDeptUsersResponseBodyListRoles struct {
	// 标签id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 标签名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签名称 tagCode
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s ListResidentDeptUsersResponseBodyListRoles) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBodyListRoles) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBodyListRoles) SetId(v int64) *ListResidentDeptUsersResponseBodyListRoles {
	s.Id = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyListRoles) SetName(v string) *ListResidentDeptUsersResponseBodyListRoles {
	s.Name = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyListRoles) SetTagCode(v string) *ListResidentDeptUsersResponseBodyListRoles {
	s.TagCode = &v
	return s
}

type ListResidentDeptUsersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResidentDeptUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentDeptUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponse) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponse) SetHeaders(v map[string]*string) *ListResidentDeptUsersResponse {
	s.Headers = v
	return s
}

func (s *ListResidentDeptUsersResponse) SetBody(v *ListResidentDeptUsersResponseBody) *ListResidentDeptUsersResponse {
	s.Body = v
	return s
}

type ListDeptSimpleUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptSimpleUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersHeaders) SetCommonHeaders(v map[string]*string) *ListDeptSimpleUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptSimpleUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptSimpleUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptSimpleUsersRequest struct {
	// cursor
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// size
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// containAccessLimit
	ContainAccessLimit *bool `json:"containAccessLimit,omitempty" xml:"containAccessLimit,omitempty"`
	// subCorpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// language
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// orderField
	OrderField *string `json:"orderField,omitempty" xml:"orderField,omitempty"`
}

func (s ListDeptSimpleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersRequest) SetCursor(v int64) *ListDeptSimpleUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetSize(v int32) *ListDeptSimpleUsersRequest {
	s.Size = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetContainAccessLimit(v bool) *ListDeptSimpleUsersRequest {
	s.ContainAccessLimit = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetSubCorpId(v string) *ListDeptSimpleUsersRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetLanguage(v string) *ListDeptSimpleUsersRequest {
	s.Language = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetOrderField(v string) *ListDeptSimpleUsersRequest {
	s.OrderField = &v
	return s
}

type ListDeptSimpleUsersResponseBody struct {
	List       []*ListDeptSimpleUsersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TotalCount *int64                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	NextCursor *int64                                 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	HasMore    *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s ListDeptSimpleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponseBody) SetList(v []*ListDeptSimpleUsersResponseBodyList) *ListDeptSimpleUsersResponseBody {
	s.List = v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetTotalCount(v int64) *ListDeptSimpleUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetNextCursor(v int64) *ListDeptSimpleUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetHasMore(v bool) *ListDeptSimpleUsersResponseBody {
	s.HasMore = &v
	return s
}

type ListDeptSimpleUsersResponseBodyList struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListDeptSimpleUsersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponseBodyList) SetUserId(v string) *ListDeptSimpleUsersResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBodyList) SetName(v string) *ListDeptSimpleUsersResponseBodyList {
	s.Name = &v
	return s
}

type ListDeptSimpleUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeptSimpleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptSimpleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponse) SetHeaders(v map[string]*string) *ListDeptSimpleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDeptSimpleUsersResponse) SetBody(v *ListDeptSimpleUsersResponseBody) *ListDeptSimpleUsersResponse {
	s.Body = v
	return s
}

type GetUserByUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserByUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserByUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserByUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserByUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserByUnionIdRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 人员 id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 通讯录语言(默认zh_CN另外支持en_US)
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserByUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdRequest) SetSubCorpId(v string) *GetUserByUnionIdRequest {
	s.SubCorpId = &v
	return s
}

func (s *GetUserByUnionIdRequest) SetUnionId(v string) *GetUserByUnionIdRequest {
	s.UnionId = &v
	return s
}

func (s *GetUserByUnionIdRequest) SetLanguage(v string) *GetUserByUnionIdRequest {
	s.Language = &v
	return s
}

type GetUserByUnionIdResponseBody struct {
	// 联系类型，0表示企业内部员工，1表示企业外部联系人
	ContactType *int32 `json:"contactType,omitempty" xml:"contactType,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserByUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdResponseBody) SetContactType(v int32) *GetUserByUnionIdResponseBody {
	s.ContactType = &v
	return s
}

func (s *GetUserByUnionIdResponseBody) SetUserId(v string) *GetUserByUnionIdResponseBody {
	s.UserId = &v
	return s
}

type GetUserByUnionIdResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserByUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdResponse) SetHeaders(v map[string]*string) *GetUserByUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserByUnionIdResponse) SetBody(v *GetUserByUnionIdResponseBody) *GetUserByUnionIdResponse {
	s.Body = v
	return s
}

type GetResidentDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentDeptHeaders) SetCommonHeaders(v map[string]*string) *GetResidentDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentDeptHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentDeptRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetResidentDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptRequest) GoString() string {
	return s.String()
}

func (s *GetResidentDeptRequest) SetSubCorpId(v string) *GetResidentDeptRequest {
	s.SubCorpId = &v
	return s
}

type GetResidentDeptResponseBody struct {
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 部门类型
	DeptType *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	// 通讯录架构类型
	ContactType *string `json:"contactType,omitempty" xml:"contactType,omitempty"`
	// 部门属性
	Feature *string `json:"feature,omitempty" xml:"feature,omitempty"`
}

func (s GetResidentDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentDeptResponseBody) SetDeptId(v int64) *GetResidentDeptResponseBody {
	s.DeptId = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetName(v string) *GetResidentDeptResponseBody {
	s.Name = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetDeptType(v string) *GetResidentDeptResponseBody {
	s.DeptType = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetContactType(v string) *GetResidentDeptResponseBody {
	s.ContactType = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetFeature(v string) *GetResidentDeptResponseBody {
	s.Feature = &v
	return s
}

type GetResidentDeptResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResidentDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResidentDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptResponse) GoString() string {
	return s.String()
}

func (s *GetResidentDeptResponse) SetHeaders(v map[string]*string) *GetResidentDeptResponse {
	s.Headers = v
	return s
}

func (s *GetResidentDeptResponse) SetBody(v *GetResidentDeptResponseBody) *GetResidentDeptResponse {
	s.Body = v
	return s
}

type GetResidentUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetResidentUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentUserInfoRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetResidentUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoRequest) SetSubCorpId(v string) *GetResidentUserInfoRequest {
	s.SubCorpId = &v
	return s
}

type GetResidentUserInfoResponseBody struct {
	// 员工id
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
	// 员工名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签列表
	Roles []*GetResidentUserInfoResponseBodyRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	// 员工特征
	Feature *string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 钉钉唯一标识
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetResidentUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponseBody) SetUserid(v string) *GetResidentUserInfoResponseBody {
	s.Userid = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetName(v string) *GetResidentUserInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetRoles(v []*GetResidentUserInfoResponseBodyRoles) *GetResidentUserInfoResponseBody {
	s.Roles = v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetFeature(v string) *GetResidentUserInfoResponseBody {
	s.Feature = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetUnionId(v string) *GetResidentUserInfoResponseBody {
	s.UnionId = &v
	return s
}

type GetResidentUserInfoResponseBodyRoles struct {
	// 标签id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 标签名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签名称 tagCode
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s GetResidentUserInfoResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponseBodyRoles) SetId(v int64) *GetResidentUserInfoResponseBodyRoles {
	s.Id = &v
	return s
}

func (s *GetResidentUserInfoResponseBodyRoles) SetName(v string) *GetResidentUserInfoResponseBodyRoles {
	s.Name = &v
	return s
}

func (s *GetResidentUserInfoResponseBodyRoles) SetTagCode(v string) *GetResidentUserInfoResponseBodyRoles {
	s.TagCode = &v
	return s
}

type GetResidentUserInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResidentUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResidentUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponse) SetHeaders(v map[string]*string) *GetResidentUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResidentUserInfoResponse) SetBody(v *GetResidentUserInfoResponseBody) *GetResidentUserInfoResponse {
	s.Body = v
	return s
}

type GetDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeptHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptHeaders) SetCommonHeaders(v map[string]*string) *GetDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeptRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 通讯录语言(默认zh_CN另外支持en_US)
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeptRequest) GoString() string {
	return s.String()
}

func (s *GetDeptRequest) SetSubCorpId(v string) *GetDeptRequest {
	s.SubCorpId = &v
	return s
}

func (s *GetDeptRequest) SetLanguage(v string) *GetDeptRequest {
	s.Language = &v
	return s
}

type GetDeptResponseBody struct {
	// 在父部门中的次序值
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父部门id
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 部门是否来自关联组织
	FromUnionOrg *bool `json:"fromUnionOrg,omitempty" xml:"fromUnionOrg,omitempty"`
}

func (s GetDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeptResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeptResponseBody) SetOrder(v int64) *GetDeptResponseBody {
	s.Order = &v
	return s
}

func (s *GetDeptResponseBody) SetDeptId(v int64) *GetDeptResponseBody {
	s.DeptId = &v
	return s
}

func (s *GetDeptResponseBody) SetName(v string) *GetDeptResponseBody {
	s.Name = &v
	return s
}

func (s *GetDeptResponseBody) SetParentId(v int64) *GetDeptResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetDeptResponseBody) SetFromUnionOrg(v bool) *GetDeptResponseBody {
	s.FromUnionOrg = &v
	return s
}

type GetDeptResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeptResponse) GoString() string {
	return s.String()
}

func (s *GetDeptResponse) SetHeaders(v map[string]*string) *GetDeptResponse {
	s.Headers = v
	return s
}

func (s *GetDeptResponse) SetBody(v *GetDeptResponseBody) *GetDeptResponse {
	s.Body = v
	return s
}

type ListParentByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListParentByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptHeaders) GoString() string {
	return s.String()
}

func (s *ListParentByDeptHeaders) SetCommonHeaders(v map[string]*string) *ListParentByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListParentByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *ListParentByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListParentByDeptRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s ListParentByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptRequest) GoString() string {
	return s.String()
}

func (s *ListParentByDeptRequest) SetSubCorpId(v string) *ListParentByDeptRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListParentByDeptRequest) SetDeptId(v int64) *ListParentByDeptRequest {
	s.DeptId = &v
	return s
}

type ListParentByDeptResponseBody struct {
	// 父部门列表
	ParentIdList []*int64 `json:"parentIdList,omitempty" xml:"parentIdList,omitempty" type:"Repeated"`
}

func (s ListParentByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *ListParentByDeptResponseBody) SetParentIdList(v []*int64) *ListParentByDeptResponseBody {
	s.ParentIdList = v
	return s
}

type ListParentByDeptResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListParentByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParentByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptResponse) GoString() string {
	return s.String()
}

func (s *ListParentByDeptResponse) SetHeaders(v map[string]*string) *ListParentByDeptResponse {
	s.Headers = v
	return s
}

func (s *ListParentByDeptResponse) SetBody(v *ListParentByDeptResponseBody) *ListParentByDeptResponse {
	s.Body = v
	return s
}

type ListDeptUserIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptUserIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsHeaders) SetCommonHeaders(v map[string]*string) *ListDeptUserIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptUserIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptUserIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptUserIdsRequest struct {
	// subCorpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListDeptUserIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsRequest) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsRequest) SetSubCorpId(v string) *ListDeptUserIdsRequest {
	s.SubCorpId = &v
	return s
}

type ListDeptUserIdsResponseBody struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ListDeptUserIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsResponseBody) SetUserIdList(v []*string) *ListDeptUserIdsResponseBody {
	s.UserIdList = v
	return s
}

type ListDeptUserIdsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeptUserIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptUserIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsResponse) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsResponse) SetHeaders(v map[string]*string) *ListDeptUserIdsResponse {
	s.Headers = v
	return s
}

func (s *ListDeptUserIdsResponse) SetBody(v *ListDeptUserIdsResponseBody) *ListDeptUserIdsResponse {
	s.Body = v
	return s
}

type ListSimpleUsersByRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSimpleUsersByRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleHeaders) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleHeaders) SetCommonHeaders(v map[string]*string) *ListSimpleUsersByRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSimpleUsersByRoleHeaders) SetXAcsDingtalkAccessToken(v string) *ListSimpleUsersByRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSimpleUsersByRoleRequest struct {
	// cursor
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// size
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// roleId
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// subCorpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListSimpleUsersByRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleRequest) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleRequest) SetOffset(v int64) *ListSimpleUsersByRoleRequest {
	s.Offset = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetSize(v int32) *ListSimpleUsersByRoleRequest {
	s.Size = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetRoleId(v int64) *ListSimpleUsersByRoleRequest {
	s.RoleId = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetSubCorpId(v string) *ListSimpleUsersByRoleRequest {
	s.SubCorpId = &v
	return s
}

type ListSimpleUsersByRoleResponseBody struct {
	List             []*ListSimpleUsersByRoleResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursorString *string                                  `json:"nextCursorString,omitempty" xml:"nextCursorString,omitempty"`
	NextCursor       *int64                                   `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	HasMore          *bool                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s ListSimpleUsersByRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponseBody) SetList(v []*ListSimpleUsersByRoleResponseBodyList) *ListSimpleUsersByRoleResponseBody {
	s.List = v
	return s
}

func (s *ListSimpleUsersByRoleResponseBody) SetNextCursorString(v string) *ListSimpleUsersByRoleResponseBody {
	s.NextCursorString = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBody) SetNextCursor(v int64) *ListSimpleUsersByRoleResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBody) SetHasMore(v bool) *ListSimpleUsersByRoleResponseBody {
	s.HasMore = &v
	return s
}

type ListSimpleUsersByRoleResponseBodyList struct {
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSimpleUsersByRoleResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponseBodyList) SetUserId(v string) *ListSimpleUsersByRoleResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyList) SetUnionId(v string) *ListSimpleUsersByRoleResponseBodyList {
	s.UnionId = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyList) SetJobNumber(v string) *ListSimpleUsersByRoleResponseBodyList {
	s.JobNumber = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyList) SetName(v string) *ListSimpleUsersByRoleResponseBodyList {
	s.Name = &v
	return s
}

type ListSimpleUsersByRoleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSimpleUsersByRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSimpleUsersByRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponse) SetHeaders(v map[string]*string) *ListSimpleUsersByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListSimpleUsersByRoleResponse) SetBody(v *ListSimpleUsersByRoleResponseBody) *ListSimpleUsersByRoleResponse {
	s.Body = v
	return s
}

type ListResidentSubDeptsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentSubDeptsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsHeaders) SetCommonHeaders(v map[string]*string) *ListResidentSubDeptsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentSubDeptsHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentSubDeptsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentSubDeptsRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 游标，不传默认1
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 大小
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListResidentSubDeptsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsRequest) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsRequest) SetSubCorpId(v string) *ListResidentSubDeptsRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentSubDeptsRequest) SetCursor(v int64) *ListResidentSubDeptsRequest {
	s.Cursor = &v
	return s
}

func (s *ListResidentSubDeptsRequest) SetSize(v int32) *ListResidentSubDeptsRequest {
	s.Size = &v
	return s
}

type ListResidentSubDeptsResponseBody struct {
	List       []*ListResidentSubDeptsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                  `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	HasMore    *bool                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Total      *int64                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListResidentSubDeptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponseBody) SetList(v []*ListResidentSubDeptsResponseBodyList) *ListResidentSubDeptsResponseBody {
	s.List = v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetNextCursor(v int64) *ListResidentSubDeptsResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetHasMore(v bool) *ListResidentSubDeptsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetTotal(v int64) *ListResidentSubDeptsResponseBody {
	s.Total = &v
	return s
}

type ListResidentSubDeptsResponseBodyList struct {
	DeptId  *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	SuperId *int64  `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s ListResidentSubDeptsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponseBodyList) SetDeptId(v int64) *ListResidentSubDeptsResponseBodyList {
	s.DeptId = &v
	return s
}

func (s *ListResidentSubDeptsResponseBodyList) SetName(v string) *ListResidentSubDeptsResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListResidentSubDeptsResponseBodyList) SetSuperId(v int64) *ListResidentSubDeptsResponseBodyList {
	s.SuperId = &v
	return s
}

type ListResidentSubDeptsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResidentSubDeptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentSubDeptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponse) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponse) SetHeaders(v map[string]*string) *ListResidentSubDeptsResponse {
	s.Headers = v
	return s
}

func (s *ListResidentSubDeptsResponse) SetBody(v *ListResidentSubDeptsResponseBody) *ListResidentSubDeptsResponse {
	s.Body = v
	return s
}

type ListParentByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListParentByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserHeaders) GoString() string {
	return s.String()
}

func (s *ListParentByUserHeaders) SetCommonHeaders(v map[string]*string) *ListParentByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListParentByUserHeaders) SetXAcsDingtalkAccessToken(v string) *ListParentByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListParentByUserRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListParentByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserRequest) GoString() string {
	return s.String()
}

func (s *ListParentByUserRequest) SetSubCorpId(v string) *ListParentByUserRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListParentByUserRequest) SetUserId(v string) *ListParentByUserRequest {
	s.UserId = &v
	return s
}

type ListParentByUserResponseBody struct {
	// 上级部门id链
	ParentDeptIdList []*int64 `json:"parentDeptIdList,omitempty" xml:"parentDeptIdList,omitempty" type:"Repeated"`
}

func (s ListParentByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListParentByUserResponseBody) SetParentDeptIdList(v []*int64) *ListParentByUserResponseBody {
	s.ParentDeptIdList = v
	return s
}

type ListParentByUserResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListParentByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParentByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserResponse) GoString() string {
	return s.String()
}

func (s *ListParentByUserResponse) SetHeaders(v map[string]*string) *ListParentByUserResponse {
	s.Headers = v
	return s
}

func (s *ListParentByUserResponse) SetBody(v *ListParentByUserResponseBody) *ListParentByUserResponse {
	s.Body = v
	return s
}

type ListSubDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptHeaders) GoString() string {
	return s.String()
}

func (s *ListSubDeptHeaders) SetCommonHeaders(v map[string]*string) *ListSubDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubDeptHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubDeptRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 通讯录语言(默认zh_CN另外支持en_US)
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s ListSubDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptRequest) GoString() string {
	return s.String()
}

func (s *ListSubDeptRequest) SetSubCorpId(v string) *ListSubDeptRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListSubDeptRequest) SetLanguage(v string) *ListSubDeptRequest {
	s.Language = &v
	return s
}

type ListSubDeptResponseBody struct {
	// 返回列表
	Result []*ListSubDeptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSubDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponseBody) SetResult(v []*ListSubDeptResponseBodyResult) *ListSubDeptResponseBody {
	s.Result = v
	return s
}

type ListSubDeptResponseBodyResult struct {
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 部门名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSubDeptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponseBodyResult) SetDeptId(v int64) *ListSubDeptResponseBodyResult {
	s.DeptId = &v
	return s
}

func (s *ListSubDeptResponseBodyResult) SetName(v string) *ListSubDeptResponseBodyResult {
	s.Name = &v
	return s
}

type ListSubDeptResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponse) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponse) SetHeaders(v map[string]*string) *ListSubDeptResponse {
	s.Headers = v
	return s
}

func (s *ListSubDeptResponse) SetBody(v *ListSubDeptResponseBody) *ListSubDeptResponse {
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

type GetUserRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 通讯录语言(默认zh_CN另外支持en_US)
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetSubCorpId(v string) *GetUserRequest {
	s.SubCorpId = &v
	return s
}

func (s *GetUserRequest) SetLanguage(v string) *GetUserRequest {
	s.Language = &v
	return s
}

type GetUserResponseBody struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 员工在当前开发者企业账号范围内的唯一标识
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 员工工号
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	// 职位
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 办公地点
	WorkPlace *string `json:"workPlace,omitempty" xml:"workPlace,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 所属部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工在对应的部门中的排序。
	DeptOrderSet []*GetUserResponseBodyDeptOrderSet `json:"deptOrderSet,omitempty" xml:"deptOrderSet,omitempty" type:"Repeated"`
	// 扩展属性，长度最大2000个字符。可以设置多种属性（手机上最多显示10个扩展属性，具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置）。 该字段的值支持链接类型填写，同时链接支持变量通配符自动替换，目前支持通配符有：userid，corpid。示例： [工位地址](http://www.dingtalk.com?userid=#userid#&corpid=#corpid#)
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 入职时间，Unix时间戳，单位ms。
	HiredDate *int64 `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	// 是否激活
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 是否实名认证
	RealAuthed *bool `json:"realAuthed,omitempty" xml:"realAuthed,omitempty"`
	// 是否高管
	Senior *bool `json:"senior,omitempty" xml:"senior,omitempty"`
	// 是否管理员
	Admin *bool `json:"admin,omitempty" xml:"admin,omitempty"`
	// 是否老板
	Boss *bool `json:"boss,omitempty" xml:"boss,omitempty"`
	// 是否专属帐号
	ExclusiveAccount *bool `json:"exclusiveAccount,omitempty" xml:"exclusiveAccount,omitempty"`
	// 专属帐号类型：{sso: 企业自定义idp;dingtalk: 钉钉idp}
	ExclusiveAccountType *string `json:"exclusiveAccountType,omitempty" xml:"exclusiveAccountType,omitempty"`
	// 员工在对应的部门中是否领导。
	LeaderInDept []*GetUserResponseBodyLeaderInDept `json:"leaderInDept,omitempty" xml:"leaderInDept,omitempty" type:"Repeated"`
	// 角色列表
	RoleList []*GetUserResponseBodyRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	// 关联信息
	UnionEmpExt *GetUserResponseBodyUnionEmpExt `json:"unionEmpExt,omitempty" xml:"unionEmpExt,omitempty" type:"Struct"`
	// 主管的ID，仅限企业内部开发调用
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetUserId(v string) *GetUserResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBody) SetUnionId(v string) *GetUserResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserResponseBody) SetName(v string) *GetUserResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserResponseBody) SetAvatar(v string) *GetUserResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetUserResponseBody) SetJobNumber(v string) *GetUserResponseBody {
	s.JobNumber = &v
	return s
}

func (s *GetUserResponseBody) SetTitle(v string) *GetUserResponseBody {
	s.Title = &v
	return s
}

func (s *GetUserResponseBody) SetWorkPlace(v string) *GetUserResponseBody {
	s.WorkPlace = &v
	return s
}

func (s *GetUserResponseBody) SetRemark(v string) *GetUserResponseBody {
	s.Remark = &v
	return s
}

func (s *GetUserResponseBody) SetDeptIdList(v []*int64) *GetUserResponseBody {
	s.DeptIdList = v
	return s
}

func (s *GetUserResponseBody) SetDeptOrderSet(v []*GetUserResponseBodyDeptOrderSet) *GetUserResponseBody {
	s.DeptOrderSet = v
	return s
}

func (s *GetUserResponseBody) SetExtension(v string) *GetUserResponseBody {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBody) SetHiredDate(v int64) *GetUserResponseBody {
	s.HiredDate = &v
	return s
}

func (s *GetUserResponseBody) SetActive(v bool) *GetUserResponseBody {
	s.Active = &v
	return s
}

func (s *GetUserResponseBody) SetRealAuthed(v bool) *GetUserResponseBody {
	s.RealAuthed = &v
	return s
}

func (s *GetUserResponseBody) SetSenior(v bool) *GetUserResponseBody {
	s.Senior = &v
	return s
}

func (s *GetUserResponseBody) SetAdmin(v bool) *GetUserResponseBody {
	s.Admin = &v
	return s
}

func (s *GetUserResponseBody) SetBoss(v bool) *GetUserResponseBody {
	s.Boss = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccount(v bool) *GetUserResponseBody {
	s.ExclusiveAccount = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountType(v string) *GetUserResponseBody {
	s.ExclusiveAccountType = &v
	return s
}

func (s *GetUserResponseBody) SetLeaderInDept(v []*GetUserResponseBodyLeaderInDept) *GetUserResponseBody {
	s.LeaderInDept = v
	return s
}

func (s *GetUserResponseBody) SetRoleList(v []*GetUserResponseBodyRoleList) *GetUserResponseBody {
	s.RoleList = v
	return s
}

func (s *GetUserResponseBody) SetUnionEmpExt(v *GetUserResponseBodyUnionEmpExt) *GetUserResponseBody {
	s.UnionEmpExt = v
	return s
}

func (s *GetUserResponseBody) SetManagerUserId(v string) *GetUserResponseBody {
	s.ManagerUserId = &v
	return s
}

type GetUserResponseBodyDeptOrderSet struct {
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 员工在部门中的排序。
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetUserResponseBodyDeptOrderSet) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDeptOrderSet) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDeptOrderSet) SetDeptId(v int64) *GetUserResponseBodyDeptOrderSet {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyDeptOrderSet) SetOrder(v int64) *GetUserResponseBodyDeptOrderSet {
	s.Order = &v
	return s
}

type GetUserResponseBodyLeaderInDept struct {
	// 部门 id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 员工在对应的部门中是否领导
	Leader *bool `json:"leader,omitempty" xml:"leader,omitempty"`
}

func (s GetUserResponseBodyLeaderInDept) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyLeaderInDept) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyLeaderInDept) SetDeptId(v int64) *GetUserResponseBodyLeaderInDept {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyLeaderInDept) SetLeader(v bool) *GetUserResponseBodyLeaderInDept {
	s.Leader = &v
	return s
}

type GetUserResponseBodyRoleList struct {
	// 角色id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 角色名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 角色组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s GetUserResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyRoleList) SetId(v int64) *GetUserResponseBodyRoleList {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetName(v string) *GetUserResponseBodyRoleList {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetGroupName(v string) *GetUserResponseBodyRoleList {
	s.GroupName = &v
	return s
}

type GetUserResponseBodyUnionEmpExt struct {
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 员工id
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// 关联映射关系
	UnionEmpMapList []*GetUserResponseBodyUnionEmpExtUnionEmpMapList `json:"unionEmpMapList,omitempty" xml:"unionEmpMapList,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUnionEmpExt) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExt) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExt) SetCorpId(v string) *GetUserResponseBodyUnionEmpExt {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetStaffId(v string) *GetUserResponseBodyUnionEmpExt {
	s.StaffId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUnionEmpMapList(v []*GetUserResponseBodyUnionEmpExtUnionEmpMapList) *GetUserResponseBodyUnionEmpExt {
	s.UnionEmpMapList = v
	return s
}

type GetUserResponseBodyUnionEmpExtUnionEmpMapList struct {
	// 企业 id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 员工 id
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetCorpId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetStaffId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.StaffId = &v
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

type ListDeptUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptUsersHeaders) SetCommonHeaders(v map[string]*string) *ListDeptUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptUsersRequest struct {
	// cursor
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// size
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// containAccessLimit
	ContainAccessLimit *bool `json:"containAccessLimit,omitempty" xml:"containAccessLimit,omitempty"`
	// subCorpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// language
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// orderField
	OrderField *string `json:"orderField,omitempty" xml:"orderField,omitempty"`
}

func (s ListDeptUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDeptUsersRequest) SetCursor(v int64) *ListDeptUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeptUsersRequest) SetSize(v int32) *ListDeptUsersRequest {
	s.Size = &v
	return s
}

func (s *ListDeptUsersRequest) SetContainAccessLimit(v bool) *ListDeptUsersRequest {
	s.ContainAccessLimit = &v
	return s
}

func (s *ListDeptUsersRequest) SetSubCorpId(v string) *ListDeptUsersRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListDeptUsersRequest) SetLanguage(v string) *ListDeptUsersRequest {
	s.Language = &v
	return s
}

func (s *ListDeptUsersRequest) SetOrderField(v string) *ListDeptUsersRequest {
	s.OrderField = &v
	return s
}

type ListDeptUsersResponseBody struct {
	List       []*ListDeptUsersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	HasMore    *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s ListDeptUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponseBody) SetList(v []*ListDeptUsersResponseBodyList) *ListDeptUsersResponseBody {
	s.List = v
	return s
}

func (s *ListDeptUsersResponseBody) SetNextCursor(v int64) *ListDeptUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListDeptUsersResponseBody) SetHasMore(v bool) *ListDeptUsersResponseBody {
	s.HasMore = &v
	return s
}

type ListDeptUsersResponseBodyList struct {
	UserId     *string  `json:"userId,omitempty" xml:"userId,omitempty"`
	UnionId    *string  `json:"unionId,omitempty" xml:"unionId,omitempty"`
	JobNumber  *string  `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name       *string  `json:"name,omitempty" xml:"name,omitempty"`
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	Active     *bool    `json:"active,omitempty" xml:"active,omitempty"`
}

func (s ListDeptUsersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponseBodyList) SetUserId(v string) *ListDeptUsersResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListDeptUsersResponseBodyList) SetUnionId(v string) *ListDeptUsersResponseBodyList {
	s.UnionId = &v
	return s
}

func (s *ListDeptUsersResponseBodyList) SetJobNumber(v string) *ListDeptUsersResponseBodyList {
	s.JobNumber = &v
	return s
}

func (s *ListDeptUsersResponseBodyList) SetName(v string) *ListDeptUsersResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListDeptUsersResponseBodyList) SetDeptIdList(v []*int64) *ListDeptUsersResponseBodyList {
	s.DeptIdList = v
	return s
}

func (s *ListDeptUsersResponseBodyList) SetActive(v bool) *ListDeptUsersResponseBodyList {
	s.Active = &v
	return s
}

type ListDeptUsersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeptUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponse) SetHeaders(v map[string]*string) *ListDeptUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDeptUsersResponse) SetBody(v *ListDeptUsersResponseBody) *ListDeptUsersResponse {
	s.Body = v
	return s
}

type ListResidentUserInfosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentUserInfosHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosHeaders) SetCommonHeaders(v map[string]*string) *ListResidentUserInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentUserInfosHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentUserInfosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentUserInfosRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 用户id列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ListResidentUserInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosRequest) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosRequest) SetSubCorpId(v string) *ListResidentUserInfosRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentUserInfosRequest) SetUserIds(v []*string) *ListResidentUserInfosRequest {
	s.UserIds = v
	return s
}

type ListResidentUserInfosShrinkRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// 用户id列表
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s ListResidentUserInfosShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosShrinkRequest) SetSubCorpId(v string) *ListResidentUserInfosShrinkRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentUserInfosShrinkRequest) SetUserIdsShrink(v string) *ListResidentUserInfosShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

type ListResidentUserInfosResponseBody struct {
	// result
	Result []*ListResidentUserInfosResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListResidentUserInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBody) SetResult(v []*ListResidentUserInfosResponseBodyResult) *ListResidentUserInfosResponseBody {
	s.Result = v
	return s
}

type ListResidentUserInfosResponseBodyResult struct {
	// 员工id
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
	// 员工名字
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签列表
	Roles []*ListResidentUserInfosResponseBodyResultRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	// 员工特征
	Feature *string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 钉钉唯一标识
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListResidentUserInfosResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBodyResult) SetUserid(v string) *ListResidentUserInfosResponseBodyResult {
	s.Userid = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyResult) SetName(v string) *ListResidentUserInfosResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyResult) SetRoles(v []*ListResidentUserInfosResponseBodyResultRoles) *ListResidentUserInfosResponseBodyResult {
	s.Roles = v
	return s
}

func (s *ListResidentUserInfosResponseBodyResult) SetFeature(v string) *ListResidentUserInfosResponseBodyResult {
	s.Feature = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyResult) SetUnionId(v string) *ListResidentUserInfosResponseBodyResult {
	s.UnionId = &v
	return s
}

type ListResidentUserInfosResponseBodyResultRoles struct {
	// 标签id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 标签名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签名称 tagCode
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s ListResidentUserInfosResponseBodyResultRoles) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBodyResultRoles) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBodyResultRoles) SetId(v int64) *ListResidentUserInfosResponseBodyResultRoles {
	s.Id = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyResultRoles) SetName(v string) *ListResidentUserInfosResponseBodyResultRoles {
	s.Name = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyResultRoles) SetTagCode(v string) *ListResidentUserInfosResponseBodyResultRoles {
	s.TagCode = &v
	return s
}

type ListResidentUserInfosResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResidentUserInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentUserInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponse) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponse) SetHeaders(v map[string]*string) *ListResidentUserInfosResponse {
	s.Headers = v
	return s
}

func (s *ListResidentUserInfosResponse) SetBody(v *ListResidentUserInfosResponseBody) *ListResidentUserInfosResponse {
	s.Body = v
	return s
}

type ListSubDeptIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubDeptIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsHeaders) SetCommonHeaders(v map[string]*string) *ListSubDeptIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubDeptIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubDeptIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubDeptIdsRequest struct {
	// 真实请求的组织corpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListSubDeptIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsRequest) SetSubCorpId(v string) *ListSubDeptIdsRequest {
	s.SubCorpId = &v
	return s
}

type ListSubDeptIdsResponseBody struct {
	// 部门ID列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
}

func (s ListSubDeptIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsResponseBody) SetDeptIdList(v []*int64) *ListSubDeptIdsResponseBody {
	s.DeptIdList = v
	return s
}

type ListSubDeptIdsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubDeptIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubDeptIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsResponse) SetHeaders(v map[string]*string) *ListSubDeptIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSubDeptIdsResponse) SetBody(v *ListSubDeptIdsResponseBody) *ListSubDeptIdsResponse {
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

func (client *Client) ListSubCorps(request *ListSubCorpsRequest) (_result *ListSubCorpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubCorpsHeaders{}
	_result = &ListSubCorpsResponse{}
	_body, _err := client.ListSubCorpsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubCorpsWithOptions(request *ListSubCorpsRequest, headers *ListSubCorpsHeaders, runtime *util.RuntimeOptions) (_result *ListSubCorpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.IsOnlyDirect)) {
		query["isOnlyDirect"] = request.IsOnlyDirect
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
	_result = &ListSubCorpsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSubCorps"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/village/corps/subCorps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVillageOrgInfo(subCorpId *string) (_result *GetVillageOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetVillageOrgInfoHeaders{}
	_result = &GetVillageOrgInfoResponse{}
	_body, _err := client.GetVillageOrgInfoWithOptions(subCorpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVillageOrgInfoWithOptions(subCorpId *string, headers *GetVillageOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *GetVillageOrgInfoResponse, _err error) {
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
	_result = &GetVillageOrgInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetVillageOrgInfo"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/corps/"+tea.StringValue(subCorpId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentDeptUsers(deptId *string, request *ListResidentDeptUsersRequest) (_result *ListResidentDeptUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentDeptUsersHeaders{}
	_result = &ListResidentDeptUsersResponse{}
	_body, _err := client.ListResidentDeptUsersWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentDeptUsersWithOptions(deptId *string, request *ListResidentDeptUsersRequest, headers *ListResidentDeptUsersHeaders, runtime *util.RuntimeOptions) (_result *ListResidentDeptUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
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
	_result = &ListResidentDeptUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListResidentDeptUsers"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/residentDepartments/"+tea.StringValue(deptId)+"/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptSimpleUsers(deptId *string, request *ListDeptSimpleUsersRequest) (_result *ListDeptSimpleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptSimpleUsersHeaders{}
	_result = &ListDeptSimpleUsersResponse{}
	_body, _err := client.ListDeptSimpleUsersWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptSimpleUsersWithOptions(deptId *string, request *ListDeptSimpleUsersRequest, headers *ListDeptSimpleUsersHeaders, runtime *util.RuntimeOptions) (_result *ListDeptSimpleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.ContainAccessLimit)) {
		query["containAccessLimit"] = request.ContainAccessLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["orderField"] = request.OrderField
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
	_result = &ListDeptSimpleUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDeptSimpleUsers"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/"+tea.StringValue(deptId)+"/simpleUsers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByUnionId(request *GetUserByUnionIdRequest) (_result *GetUserByUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserByUnionIdHeaders{}
	_result = &GetUserByUnionIdResponse{}
	_body, _err := client.GetUserByUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserByUnionIdWithOptions(request *GetUserByUnionIdRequest, headers *GetUserByUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserByUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
	_result = &GetUserByUnionIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserByUnionId"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/users/getByUnionId"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResidentDept(departmentId *string, request *GetResidentDeptRequest) (_result *GetResidentDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentDeptHeaders{}
	_result = &GetResidentDeptResponse{}
	_body, _err := client.GetResidentDeptWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResidentDeptWithOptions(departmentId *string, request *GetResidentDeptRequest, headers *GetResidentDeptHeaders, runtime *util.RuntimeOptions) (_result *GetResidentDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &GetResidentDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("GetResidentDept"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/residentDepartments/departments/"+tea.StringValue(departmentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResidentUserInfo(deptId *string, userId *string, request *GetResidentUserInfoRequest) (_result *GetResidentUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentUserInfoHeaders{}
	_result = &GetResidentUserInfoResponse{}
	_body, _err := client.GetResidentUserInfoWithOptions(deptId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResidentUserInfoWithOptions(deptId *string, userId *string, request *GetResidentUserInfoRequest, headers *GetResidentUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetResidentUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &GetResidentUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetResidentUserInfo"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/residentDepartments/"+tea.StringValue(deptId)+"/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDept(departmentId *string, request *GetDeptRequest) (_result *GetDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeptHeaders{}
	_result = &GetDeptResponse{}
	_body, _err := client.GetDeptWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeptWithOptions(departmentId *string, request *GetDeptRequest, headers *GetDeptHeaders, runtime *util.RuntimeOptions) (_result *GetDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &GetDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDept"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/deptartments/"+tea.StringValue(departmentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParentByDept(request *ListParentByDeptRequest) (_result *ListParentByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListParentByDeptHeaders{}
	_result = &ListParentByDeptResponse{}
	_body, _err := client.ListParentByDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParentByDeptWithOptions(request *ListParentByDeptRequest, headers *ListParentByDeptHeaders, runtime *util.RuntimeOptions) (_result *ListParentByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &ListParentByDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("ListParentByDept"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/listParentByDepartment"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptUserIds(deptId *string, request *ListDeptUserIdsRequest) (_result *ListDeptUserIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptUserIdsHeaders{}
	_result = &ListDeptUserIdsResponse{}
	_body, _err := client.ListDeptUserIdsWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptUserIdsWithOptions(deptId *string, request *ListDeptUserIdsRequest, headers *ListDeptUserIdsHeaders, runtime *util.RuntimeOptions) (_result *ListDeptUserIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &ListDeptUserIdsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDeptUserIds"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/"+tea.StringValue(deptId)+"/userIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSimpleUsersByRole(request *ListSimpleUsersByRoleRequest) (_result *ListSimpleUsersByRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSimpleUsersByRoleHeaders{}
	_result = &ListSimpleUsersByRoleResponse{}
	_body, _err := client.ListSimpleUsersByRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSimpleUsersByRoleWithOptions(request *ListSimpleUsersByRoleRequest, headers *ListSimpleUsersByRoleHeaders, runtime *util.RuntimeOptions) (_result *ListSimpleUsersByRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["roleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &ListSimpleUsersByRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSimpleUsersByRole"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/users/listByRole"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentSubDepts(deptId *string, request *ListResidentSubDeptsRequest) (_result *ListResidentSubDeptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentSubDeptsHeaders{}
	_result = &ListResidentSubDeptsResponse{}
	_body, _err := client.ListResidentSubDeptsWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentSubDeptsWithOptions(deptId *string, request *ListResidentSubDeptsRequest, headers *ListResidentSubDeptsHeaders, runtime *util.RuntimeOptions) (_result *ListResidentSubDeptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
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
	_result = &ListResidentSubDeptsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListResidentSubDepts"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/residentDepartments/"+tea.StringValue(deptId)+"/subDepartments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParentByUser(request *ListParentByUserRequest) (_result *ListParentByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListParentByUserHeaders{}
	_result = &ListParentByUserResponse{}
	_body, _err := client.ListParentByUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParentByUserWithOptions(request *ListParentByUserRequest, headers *ListParentByUserHeaders, runtime *util.RuntimeOptions) (_result *ListParentByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

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
	_result = &ListParentByUserResponse{}
	_body, _err := client.DoROARequest(tea.String("ListParentByUser"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/listParentByUser"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubDept(deptId *string, request *ListSubDeptRequest) (_result *ListSubDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubDeptHeaders{}
	_result = &ListSubDeptResponse{}
	_body, _err := client.ListSubDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubDeptWithOptions(deptId *string, request *ListSubDeptRequest, headers *ListSubDeptHeaders, runtime *util.RuntimeOptions) (_result *ListSubDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &ListSubDeptResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSubDept"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/"+tea.StringValue(deptId)+"/subDepartments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(userId *string, request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(userId *string, request *GetUserRequest, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &GetUserResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUser"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/users/getByUserId"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptUsers(deptId *string, request *ListDeptUsersRequest) (_result *ListDeptUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptUsersHeaders{}
	_result = &ListDeptUsersResponse{}
	_body, _err := client.ListDeptUsersWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptUsersWithOptions(deptId *string, request *ListDeptUsersRequest, headers *ListDeptUsersHeaders, runtime *util.RuntimeOptions) (_result *ListDeptUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.ContainAccessLimit)) {
		query["containAccessLimit"] = request.ContainAccessLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["orderField"] = request.OrderField
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
	_result = &ListDeptUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDeptUsers"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/"+tea.StringValue(deptId)+"/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentUserInfos(request *ListResidentUserInfosRequest) (_result *ListResidentUserInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentUserInfosHeaders{}
	_result = &ListResidentUserInfosResponse{}
	_body, _err := client.ListResidentUserInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentUserInfosWithOptions(tmpReq *ListResidentUserInfosRequest, headers *ListResidentUserInfosHeaders, runtime *util.RuntimeOptions) (_result *ListResidentUserInfosResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResidentUserInfosShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIds)) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, tea.String("userIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdsShrink)) {
		query["userIds"] = request.UserIdsShrink
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
	_result = &ListResidentUserInfosResponse{}
	_body, _err := client.DoROARequest(tea.String("ListResidentUserInfos"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/residentUsers/getByUserIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubDeptIds(deptId *string, request *ListSubDeptIdsRequest) (_result *ListSubDeptIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubDeptIdsHeaders{}
	_result = &ListSubDeptIdsResponse{}
	_body, _err := client.ListSubDeptIdsWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubDeptIdsWithOptions(deptId *string, request *ListSubDeptIdsRequest, headers *ListSubDeptIdsHeaders, runtime *util.RuntimeOptions) (_result *ListSubDeptIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
	_result = &ListSubDeptIdsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSubDeptIds"), tea.String("village_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/village/departments/"+tea.StringValue(deptId)+"/subDepartmentIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
