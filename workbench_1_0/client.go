// This file is auto-generated, don't edit it. Thanks.
package workbench_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddRecentUserAppListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRecentUserAppListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRecentUserAppListHeaders) GoString() string {
	return s.String()
}

func (s *AddRecentUserAppListHeaders) SetCommonHeaders(v map[string]*string) *AddRecentUserAppListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRecentUserAppListHeaders) SetXAcsDingtalkAccessToken(v string) *AddRecentUserAppListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRecentUserAppListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding48143d56cd15327624f2f5cc6abecb85
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	UsedAppDetailList []*AddRecentUserAppListRequestUsedAppDetailList `json:"usedAppDetailList,omitempty" xml:"usedAppDetailList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 642325391030949
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddRecentUserAppListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecentUserAppListRequest) GoString() string {
	return s.String()
}

func (s *AddRecentUserAppListRequest) SetCorpId(v string) *AddRecentUserAppListRequest {
	s.CorpId = &v
	return s
}

func (s *AddRecentUserAppListRequest) SetUsedAppDetailList(v []*AddRecentUserAppListRequestUsedAppDetailList) *AddRecentUserAppListRequest {
	s.UsedAppDetailList = v
	return s
}

func (s *AddRecentUserAppListRequest) SetUserId(v string) *AddRecentUserAppListRequest {
	s.UserId = &v
	return s
}

type AddRecentUserAppListRequestUsedAppDetailList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2636835622
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s AddRecentUserAppListRequestUsedAppDetailList) String() string {
	return tea.Prettify(s)
}

func (s AddRecentUserAppListRequestUsedAppDetailList) GoString() string {
	return s.String()
}

func (s *AddRecentUserAppListRequestUsedAppDetailList) SetAgentId(v string) *AddRecentUserAppListRequestUsedAppDetailList {
	s.AgentId = &v
	return s
}

type AddRecentUserAppListResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddRecentUserAppListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecentUserAppListResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecentUserAppListResponseBody) SetResult(v bool) *AddRecentUserAppListResponseBody {
	s.Result = &v
	return s
}

type AddRecentUserAppListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecentUserAppListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecentUserAppListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRecentUserAppListResponse) GoString() string {
	return s.String()
}

func (s *AddRecentUserAppListResponse) SetHeaders(v map[string]*string) *AddRecentUserAppListResponse {
	s.Headers = v
	return s
}

func (s *AddRecentUserAppListResponse) SetStatusCode(v int32) *AddRecentUserAppListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecentUserAppListResponse) SetBody(v *AddRecentUserAppListResponseBody) *AddRecentUserAppListResponse {
	s.Body = v
	return s
}

type GetDingPortalDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingPortalDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailHeaders) SetCommonHeaders(v map[string]*string) *GetDingPortalDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingPortalDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingPortalDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingPortalDetailResponseBody struct {
	// example:
	//
	// SWAPP-XXX
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 我的自定义工作台
	DingPortalName *string                                 `json:"dingPortalName,omitempty" xml:"dingPortalName,omitempty"`
	Pages          []*GetDingPortalDetailResponseBodyPages `json:"pages,omitempty" xml:"pages,omitempty" type:"Repeated"`
}

func (s GetDingPortalDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponseBody) SetAppUuid(v string) *GetDingPortalDetailResponseBody {
	s.AppUuid = &v
	return s
}

func (s *GetDingPortalDetailResponseBody) SetDingPortalName(v string) *GetDingPortalDetailResponseBody {
	s.DingPortalName = &v
	return s
}

func (s *GetDingPortalDetailResponseBody) SetPages(v []*GetDingPortalDetailResponseBodyPages) *GetDingPortalDetailResponseBody {
	s.Pages = v
	return s
}

type GetDingPortalDetailResponseBodyPages struct {
	// example:
	//
	// false
	AllVisible *bool    `json:"allVisible,omitempty" xml:"allVisible,omitempty"`
	DeptIds    []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// example:
	//
	// 我的工作台页面
	PageName *string `json:"pageName,omitempty" xml:"pageName,omitempty"`
	// example:
	//
	// XX-XX-XX
	PageUuid *string   `json:"pageUuid,omitempty" xml:"pageUuid,omitempty"`
	RoleIds  []*int64  `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	Userids  []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
}

func (s GetDingPortalDetailResponseBodyPages) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponseBodyPages) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponseBodyPages) SetAllVisible(v bool) *GetDingPortalDetailResponseBodyPages {
	s.AllVisible = &v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetDeptIds(v []*int64) *GetDingPortalDetailResponseBodyPages {
	s.DeptIds = v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetPageName(v string) *GetDingPortalDetailResponseBodyPages {
	s.PageName = &v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetPageUuid(v string) *GetDingPortalDetailResponseBodyPages {
	s.PageUuid = &v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetRoleIds(v []*int64) *GetDingPortalDetailResponseBodyPages {
	s.RoleIds = v
	return s
}

func (s *GetDingPortalDetailResponseBodyPages) SetUserids(v []*string) *GetDingPortalDetailResponseBodyPages {
	s.Userids = v
	return s
}

type GetDingPortalDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingPortalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDingPortalDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingPortalDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDingPortalDetailResponse) SetHeaders(v map[string]*string) *GetDingPortalDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDingPortalDetailResponse) SetStatusCode(v int32) *GetDingPortalDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingPortalDetailResponse) SetBody(v *GetDingPortalDetailResponseBody) *GetDingPortalDetailResponse {
	s.Body = v
	return s
}

type GetPluginPermissionPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPluginPermissionPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPluginPermissionPointHeaders) GoString() string {
	return s.String()
}

func (s *GetPluginPermissionPointHeaders) SetCommonHeaders(v map[string]*string) *GetPluginPermissionPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPluginPermissionPointHeaders) SetXAcsDingtalkAccessToken(v string) *GetPluginPermissionPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPluginPermissionPointRequest struct {
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s GetPluginPermissionPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPluginPermissionPointRequest) GoString() string {
	return s.String()
}

func (s *GetPluginPermissionPointRequest) SetMiniAppId(v string) *GetPluginPermissionPointRequest {
	s.MiniAppId = &v
	return s
}

type GetPluginPermissionPointResponseBody struct {
	PermissionPointList []*string `json:"permissionPointList,omitempty" xml:"permissionPointList,omitempty" type:"Repeated"`
}

func (s GetPluginPermissionPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPluginPermissionPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginPermissionPointResponseBody) SetPermissionPointList(v []*string) *GetPluginPermissionPointResponseBody {
	s.PermissionPointList = v
	return s
}

type GetPluginPermissionPointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginPermissionPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginPermissionPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPluginPermissionPointResponse) GoString() string {
	return s.String()
}

func (s *GetPluginPermissionPointResponse) SetHeaders(v map[string]*string) *GetPluginPermissionPointResponse {
	s.Headers = v
	return s
}

func (s *GetPluginPermissionPointResponse) SetStatusCode(v int32) *GetPluginPermissionPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginPermissionPointResponse) SetBody(v *GetPluginPermissionPointResponseBody) *GetPluginPermissionPointResponse {
	s.Body = v
	return s
}

type GetPluginRuleCheckInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPluginRuleCheckInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPluginRuleCheckInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPluginRuleCheckInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPluginRuleCheckInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPluginRuleCheckInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPluginRuleCheckInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPluginRuleCheckInfoRequest struct {
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
}

func (s GetPluginRuleCheckInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPluginRuleCheckInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPluginRuleCheckInfoRequest) SetMiniAppId(v string) *GetPluginRuleCheckInfoRequest {
	s.MiniAppId = &v
	return s
}

type GetPluginRuleCheckInfoResponseBody struct {
	PackCode              *string `json:"packCode,omitempty" xml:"packCode,omitempty"`
	PluginRuleCheckDetail *string `json:"pluginRuleCheckDetail,omitempty" xml:"pluginRuleCheckDetail,omitempty"`
}

func (s GetPluginRuleCheckInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPluginRuleCheckInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginRuleCheckInfoResponseBody) SetPackCode(v string) *GetPluginRuleCheckInfoResponseBody {
	s.PackCode = &v
	return s
}

func (s *GetPluginRuleCheckInfoResponseBody) SetPluginRuleCheckDetail(v string) *GetPluginRuleCheckInfoResponseBody {
	s.PluginRuleCheckDetail = &v
	return s
}

type GetPluginRuleCheckInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginRuleCheckInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginRuleCheckInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPluginRuleCheckInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPluginRuleCheckInfoResponse) SetHeaders(v map[string]*string) *GetPluginRuleCheckInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPluginRuleCheckInfoResponse) SetStatusCode(v int32) *GetPluginRuleCheckInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginRuleCheckInfoResponse) SetBody(v *GetPluginRuleCheckInfoResponseBody) *GetPluginRuleCheckInfoResponse {
	s.Body = v
	return s
}

type ListWorkBenchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListWorkBenchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListWorkBenchGroupHeaders) GoString() string {
	return s.String()
}

func (s *ListWorkBenchGroupHeaders) SetCommonHeaders(v map[string]*string) *ListWorkBenchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListWorkBenchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *ListWorkBenchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListWorkBenchGroupRequest struct {
	// example:
	//
	// corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 首页 = WORK_HOME 全部页 = WORK_ALL_APP
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
}

func (s ListWorkBenchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkBenchGroupRequest) GoString() string {
	return s.String()
}

func (s *ListWorkBenchGroupRequest) SetEcologicalCorpId(v string) *ListWorkBenchGroupRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *ListWorkBenchGroupRequest) SetGroupType(v string) *ListWorkBenchGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ListWorkBenchGroupRequest) SetOpUnionId(v string) *ListWorkBenchGroupRequest {
	s.OpUnionId = &v
	return s
}

type ListWorkBenchGroupResponseBody struct {
	GroupList []*ListWorkBenchGroupResponseBodyGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
}

func (s ListWorkBenchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkBenchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkBenchGroupResponseBody) SetGroupList(v []*ListWorkBenchGroupResponseBodyGroupList) *ListWorkBenchGroupResponseBody {
	s.GroupList = v
	return s
}

type ListWorkBenchGroupResponseBodyGroupList struct {
	// example:
	//
	// desc
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListWorkBenchGroupResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListWorkBenchGroupResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *ListWorkBenchGroupResponseBodyGroupList) SetComponentId(v string) *ListWorkBenchGroupResponseBodyGroupList {
	s.ComponentId = &v
	return s
}

func (s *ListWorkBenchGroupResponseBodyGroupList) SetName(v string) *ListWorkBenchGroupResponseBodyGroupList {
	s.Name = &v
	return s
}

type ListWorkBenchGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkBenchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkBenchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkBenchGroupResponse) GoString() string {
	return s.String()
}

func (s *ListWorkBenchGroupResponse) SetHeaders(v map[string]*string) *ListWorkBenchGroupResponse {
	s.Headers = v
	return s
}

func (s *ListWorkBenchGroupResponse) SetStatusCode(v int32) *ListWorkBenchGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkBenchGroupResponse) SetBody(v *ListWorkBenchGroupResponseBody) *ListWorkBenchGroupResponse {
	s.Body = v
	return s
}

type ModifyWorkbenchBadgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ModifyWorkbenchBadgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkbenchBadgeHeaders) GoString() string {
	return s.String()
}

func (s *ModifyWorkbenchBadgeHeaders) SetCommonHeaders(v map[string]*string) *ModifyWorkbenchBadgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ModifyWorkbenchBadgeHeaders) SetXAcsDingtalkAccessToken(v string) *ModifyWorkbenchBadgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ModifyWorkbenchBadgeRequest struct {
	// if can be null:
	// false
	BizIdList []*string `json:"bizIdList,omitempty" xml:"bizIdList,omitempty" type:"Repeated"`
	IsAdded   *bool     `json:"isAdded,omitempty" xml:"isAdded,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// full
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5000000004278081/test
	RedDotRelationId *string `json:"redDotRelationId,omitempty" xml:"redDotRelationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workbench_component
	RedDotType *string `json:"redDotType,omitempty" xml:"redDotType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0123465
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ModifyWorkbenchBadgeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkbenchBadgeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWorkbenchBadgeRequest) SetBizIdList(v []*string) *ModifyWorkbenchBadgeRequest {
	s.BizIdList = v
	return s
}

func (s *ModifyWorkbenchBadgeRequest) SetIsAdded(v bool) *ModifyWorkbenchBadgeRequest {
	s.IsAdded = &v
	return s
}

func (s *ModifyWorkbenchBadgeRequest) SetModifyMode(v string) *ModifyWorkbenchBadgeRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyWorkbenchBadgeRequest) SetRedDotRelationId(v string) *ModifyWorkbenchBadgeRequest {
	s.RedDotRelationId = &v
	return s
}

func (s *ModifyWorkbenchBadgeRequest) SetRedDotType(v string) *ModifyWorkbenchBadgeRequest {
	s.RedDotType = &v
	return s
}

func (s *ModifyWorkbenchBadgeRequest) SetUserId(v string) *ModifyWorkbenchBadgeRequest {
	s.UserId = &v
	return s
}

type ModifyWorkbenchBadgeResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyWorkbenchBadgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkbenchBadgeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWorkbenchBadgeResponseBody) SetResult(v bool) *ModifyWorkbenchBadgeResponseBody {
	s.Result = &v
	return s
}

type ModifyWorkbenchBadgeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWorkbenchBadgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWorkbenchBadgeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWorkbenchBadgeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWorkbenchBadgeResponse) SetHeaders(v map[string]*string) *ModifyWorkbenchBadgeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWorkbenchBadgeResponse) SetStatusCode(v int32) *ModifyWorkbenchBadgeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWorkbenchBadgeResponse) SetBody(v *ModifyWorkbenchBadgeResponseBody) *ModifyWorkbenchBadgeResponse {
	s.Body = v
	return s
}

type QueryComponentScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryComponentScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesHeaders) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesHeaders) SetCommonHeaders(v map[string]*string) *QueryComponentScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryComponentScopesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryComponentScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryComponentScopesResponseBody struct {
	DeptVisibleScopes []*int64  `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
	UserVisibleScopes []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
}

func (s QueryComponentScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesResponseBody) SetDeptVisibleScopes(v []*int64) *QueryComponentScopesResponseBody {
	s.DeptVisibleScopes = v
	return s
}

func (s *QueryComponentScopesResponseBody) SetUserVisibleScopes(v []*string) *QueryComponentScopesResponseBody {
	s.UserVisibleScopes = v
	return s
}

type QueryComponentScopesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryComponentScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryComponentScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentScopesResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentScopesResponse) SetHeaders(v map[string]*string) *QueryComponentScopesResponse {
	s.Headers = v
	return s
}

func (s *QueryComponentScopesResponse) SetStatusCode(v int32) *QueryComponentScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryComponentScopesResponse) SetBody(v *QueryComponentScopesResponseBody) *QueryComponentScopesResponse {
	s.Body = v
	return s
}

type QueryShortcutScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryShortcutScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesHeaders) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesHeaders) SetCommonHeaders(v map[string]*string) *QueryShortcutScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryShortcutScopesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryShortcutScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryShortcutScopesResponseBody struct {
	DeptVisibleScopes []*int64  `json:"deptVisibleScopes,omitempty" xml:"deptVisibleScopes,omitempty" type:"Repeated"`
	UserVisibleScopes []*string `json:"userVisibleScopes,omitempty" xml:"userVisibleScopes,omitempty" type:"Repeated"`
}

func (s QueryShortcutScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesResponseBody) SetDeptVisibleScopes(v []*int64) *QueryShortcutScopesResponseBody {
	s.DeptVisibleScopes = v
	return s
}

func (s *QueryShortcutScopesResponseBody) SetUserVisibleScopes(v []*string) *QueryShortcutScopesResponseBody {
	s.UserVisibleScopes = v
	return s
}

type QueryShortcutScopesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryShortcutScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryShortcutScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShortcutScopesResponse) GoString() string {
	return s.String()
}

func (s *QueryShortcutScopesResponse) SetHeaders(v map[string]*string) *QueryShortcutScopesResponse {
	s.Headers = v
	return s
}

func (s *QueryShortcutScopesResponse) SetStatusCode(v int32) *QueryShortcutScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryShortcutScopesResponse) SetBody(v *QueryShortcutScopesResponseBody) *QueryShortcutScopesResponse {
	s.Body = v
	return s
}

type UndoDeletionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UndoDeletionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UndoDeletionHeaders) GoString() string {
	return s.String()
}

func (s *UndoDeletionHeaders) SetCommonHeaders(v map[string]*string) *UndoDeletionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UndoDeletionHeaders) SetXAcsDingtalkAccessToken(v string) *UndoDeletionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UndoDeletionRequest struct {
	// This parameter is required.
	BizIdList []*string `json:"bizIdList,omitempty" xml:"bizIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	RedDotRelationId *string `json:"redDotRelationId,omitempty" xml:"redDotRelationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workbench_component
	RedDotType *string `json:"redDotType,omitempty" xml:"redDotType,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UndoDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s UndoDeletionRequest) GoString() string {
	return s.String()
}

func (s *UndoDeletionRequest) SetBizIdList(v []*string) *UndoDeletionRequest {
	s.BizIdList = v
	return s
}

func (s *UndoDeletionRequest) SetRedDotRelationId(v string) *UndoDeletionRequest {
	s.RedDotRelationId = &v
	return s
}

func (s *UndoDeletionRequest) SetRedDotType(v string) *UndoDeletionRequest {
	s.RedDotType = &v
	return s
}

func (s *UndoDeletionRequest) SetUserId(v string) *UndoDeletionRequest {
	s.UserId = &v
	return s
}

type UndoDeletionResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UndoDeletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UndoDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *UndoDeletionResponseBody) SetResult(v bool) *UndoDeletionResponseBody {
	s.Result = &v
	return s
}

type UndoDeletionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UndoDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UndoDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s UndoDeletionResponse) GoString() string {
	return s.String()
}

func (s *UndoDeletionResponse) SetHeaders(v map[string]*string) *UndoDeletionResponse {
	s.Headers = v
	return s
}

func (s *UndoDeletionResponse) SetStatusCode(v int32) *UndoDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *UndoDeletionResponse) SetBody(v *UndoDeletionResponseBody) *UndoDeletionResponse {
	s.Body = v
	return s
}

type UpdateDingPortalPageScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateDingPortalPageScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeHeaders) SetCommonHeaders(v map[string]*string) *UpdateDingPortalPageScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDingPortalPageScopeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateDingPortalPageScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateDingPortalPageScopeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	AllVisible *bool     `json:"allVisible,omitempty" xml:"allVisible,omitempty"`
	DeptIds    []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	RoleIds    []*int64  `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	Userids    []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
}

func (s UpdateDingPortalPageScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeRequest) SetAllVisible(v bool) *UpdateDingPortalPageScopeRequest {
	s.AllVisible = &v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetDeptIds(v []*int64) *UpdateDingPortalPageScopeRequest {
	s.DeptIds = v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetRoleIds(v []*int64) *UpdateDingPortalPageScopeRequest {
	s.RoleIds = v
	return s
}

func (s *UpdateDingPortalPageScopeRequest) SetUserids(v []*string) *UpdateDingPortalPageScopeRequest {
	s.Userids = v
	return s
}

type UpdateDingPortalPageScopeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateDingPortalPageScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingPortalPageScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDingPortalPageScopeResponse) SetHeaders(v map[string]*string) *UpdateDingPortalPageScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDingPortalPageScopeResponse) SetStatusCode(v int32) *UpdateDingPortalPageScopeResponse {
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
// 批量添加最近使用记录
//
// @param request - AddRecentUserAppListRequest
//
// @param headers - AddRecentUserAppListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecentUserAppListResponse
func (client *Client) AddRecentUserAppListWithOptions(request *AddRecentUserAppListRequest, headers *AddRecentUserAppListHeaders, runtime *util.RuntimeOptions) (_result *AddRecentUserAppListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UsedAppDetailList)) {
		body["usedAppDetailList"] = request.UsedAppDetailList
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
		Action:      tea.String("AddRecentUserAppList"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/components/recentUsed/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRecentUserAppListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加最近使用记录
//
// @param request - AddRecentUserAppListRequest
//
// @return AddRecentUserAppListResponse
func (client *Client) AddRecentUserAppList(request *AddRecentUserAppListRequest) (_result *AddRecentUserAppListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRecentUserAppListHeaders{}
	_result = &AddRecentUserAppListResponse{}
	_body, _err := client.AddRecentUserAppListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义工作台
//
// @param headers - GetDingPortalDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingPortalDetailResponse
func (client *Client) GetDingPortalDetailWithOptions(appUuid *string, headers *GetDingPortalDetailHeaders, runtime *util.RuntimeOptions) (_result *GetDingPortalDetailResponse, _err error) {
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
		Action:      tea.String("GetDingPortalDetail"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/dingPortals/" + tea.StringValue(appUuid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDingPortalDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义工作台
//
// @return GetDingPortalDetailResponse
func (client *Client) GetDingPortalDetail(appUuid *string) (_result *GetDingPortalDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingPortalDetailHeaders{}
	_result = &GetDingPortalDetailResponse{}
	_body, _err := client.GetDingPortalDetailWithOptions(appUuid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作台插件的权限点
//
// @param request - GetPluginPermissionPointRequest
//
// @param headers - GetPluginPermissionPointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginPermissionPointResponse
func (client *Client) GetPluginPermissionPointWithOptions(request *GetPluginPermissionPointRequest, headers *GetPluginPermissionPointHeaders, runtime *util.RuntimeOptions) (_result *GetPluginPermissionPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

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
		Action:      tea.String("GetPluginPermissionPoint"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/plugins/permissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPluginPermissionPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作台插件的权限点
//
// @param request - GetPluginPermissionPointRequest
//
// @return GetPluginPermissionPointResponse
func (client *Client) GetPluginPermissionPoint(request *GetPluginPermissionPointRequest) (_result *GetPluginPermissionPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPluginPermissionPointHeaders{}
	_result = &GetPluginPermissionPointResponse{}
	_body, _err := client.GetPluginPermissionPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取插件的校验规则
//
// @param request - GetPluginRuleCheckInfoRequest
//
// @param headers - GetPluginRuleCheckInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginRuleCheckInfoResponse
func (client *Client) GetPluginRuleCheckInfoWithOptions(request *GetPluginRuleCheckInfoRequest, headers *GetPluginRuleCheckInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPluginRuleCheckInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

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
		Action:      tea.String("GetPluginRuleCheckInfo"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/plugins/validationRules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPluginRuleCheckInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取插件的校验规则
//
// @param request - GetPluginRuleCheckInfoRequest
//
// @return GetPluginRuleCheckInfoResponse
func (client *Client) GetPluginRuleCheckInfo(request *GetPluginRuleCheckInfoRequest) (_result *GetPluginRuleCheckInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPluginRuleCheckInfoHeaders{}
	_result = &GetPluginRuleCheckInfoResponse{}
	_body, _err := client.GetPluginRuleCheckInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作台分组列表
//
// @param request - ListWorkBenchGroupRequest
//
// @param headers - ListWorkBenchGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkBenchGroupResponse
func (client *Client) ListWorkBenchGroupWithOptions(request *ListWorkBenchGroupRequest, headers *ListWorkBenchGroupHeaders, runtime *util.RuntimeOptions) (_result *ListWorkBenchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
	}

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
		Action:      tea.String("ListWorkBenchGroup"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkBenchGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作台分组列表
//
// @param request - ListWorkBenchGroupRequest
//
// @return ListWorkBenchGroupResponse
func (client *Client) ListWorkBenchGroup(request *ListWorkBenchGroupRequest) (_result *ListWorkBenchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListWorkBenchGroupHeaders{}
	_result = &ListWorkBenchGroupResponse{}
	_body, _err := client.ListWorkBenchGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工作台支持数字红点
//
// @param request - ModifyWorkbenchBadgeRequest
//
// @param headers - ModifyWorkbenchBadgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWorkbenchBadgeResponse
func (client *Client) ModifyWorkbenchBadgeWithOptions(request *ModifyWorkbenchBadgeRequest, headers *ModifyWorkbenchBadgeHeaders, runtime *util.RuntimeOptions) (_result *ModifyWorkbenchBadgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizIdList)) {
		body["bizIdList"] = request.BizIdList
	}

	if !tea.BoolValue(util.IsUnset(request.IsAdded)) {
		body["isAdded"] = request.IsAdded
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMode)) {
		body["modifyMode"] = request.ModifyMode
	}

	if !tea.BoolValue(util.IsUnset(request.RedDotRelationId)) {
		body["redDotRelationId"] = request.RedDotRelationId
	}

	if !tea.BoolValue(util.IsUnset(request.RedDotType)) {
		body["redDotType"] = request.RedDotType
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
		Action:      tea.String("ModifyWorkbenchBadge"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/badges/modify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWorkbenchBadgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作台支持数字红点
//
// @param request - ModifyWorkbenchBadgeRequest
//
// @return ModifyWorkbenchBadgeResponse
func (client *Client) ModifyWorkbenchBadge(request *ModifyWorkbenchBadgeRequest) (_result *ModifyWorkbenchBadgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ModifyWorkbenchBadgeHeaders{}
	_result = &ModifyWorkbenchBadgeResponse{}
	_body, _err := client.ModifyWorkbenchBadgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工作台组件授权范围查询
//
// @param headers - QueryComponentScopesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryComponentScopesResponse
func (client *Client) QueryComponentScopesWithOptions(componentId *string, headers *QueryComponentScopesHeaders, runtime *util.RuntimeOptions) (_result *QueryComponentScopesResponse, _err error) {
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
		Action:      tea.String("QueryComponentScopes"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/components/" + tea.StringValue(componentId) + "/scopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryComponentScopesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作台组件授权范围查询
//
// @return QueryComponentScopesResponse
func (client *Client) QueryComponentScopes(componentId *string) (_result *QueryComponentScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryComponentScopesHeaders{}
	_result = &QueryComponentScopesResponse{}
	_body, _err := client.QueryComponentScopesWithOptions(componentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询快捷方式可见范围
//
// @param headers - QueryShortcutScopesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShortcutScopesResponse
func (client *Client) QueryShortcutScopesWithOptions(shortcutKey *string, headers *QueryShortcutScopesHeaders, runtime *util.RuntimeOptions) (_result *QueryShortcutScopesResponse, _err error) {
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
		Action:      tea.String("QueryShortcutScopes"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/shortcuts/" + tea.StringValue(shortcutKey) + "/scopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryShortcutScopesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询快捷方式可见范围
//
// @return QueryShortcutScopesResponse
func (client *Client) QueryShortcutScopes(shortcutKey *string) (_result *QueryShortcutScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryShortcutScopesHeaders{}
	_result = &QueryShortcutScopesResponse{}
	_body, _err := client.QueryShortcutScopesWithOptions(shortcutKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工作台数字红点支持撤销已被删除的资源
//
// @param request - UndoDeletionRequest
//
// @param headers - UndoDeletionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UndoDeletionResponse
func (client *Client) UndoDeletionWithOptions(request *UndoDeletionRequest, headers *UndoDeletionHeaders, runtime *util.RuntimeOptions) (_result *UndoDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizIdList)) {
		body["bizIdList"] = request.BizIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RedDotRelationId)) {
		body["redDotRelationId"] = request.RedDotRelationId
	}

	if !tea.BoolValue(util.IsUnset(request.RedDotType)) {
		body["redDotType"] = request.RedDotType
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
		Action:      tea.String("UndoDeletion"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/badges/undoDeleted"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UndoDeletionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作台数字红点支持撤销已被删除的资源
//
// @param request - UndoDeletionRequest
//
// @return UndoDeletionResponse
func (client *Client) UndoDeletion(request *UndoDeletionRequest) (_result *UndoDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UndoDeletionHeaders{}
	_result = &UndoDeletionResponse{}
	_body, _err := client.UndoDeletionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新自定义工作台页面可见性
//
// @param request - UpdateDingPortalPageScopeRequest
//
// @param headers - UpdateDingPortalPageScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDingPortalPageScopeResponse
func (client *Client) UpdateDingPortalPageScopeWithOptions(pageUuid *string, appUuid *string, request *UpdateDingPortalPageScopeRequest, headers *UpdateDingPortalPageScopeHeaders, runtime *util.RuntimeOptions) (_result *UpdateDingPortalPageScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllVisible)) {
		body["allVisible"] = request.AllVisible
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		body["roleIds"] = request.RoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.Userids)) {
		body["userids"] = request.Userids
	}

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
		Action:      tea.String("UpdateDingPortalPageScope"),
		Version:     tea.String("workbench_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workbench/dingPortals/" + tea.StringValue(appUuid) + "/pageScopes/" + tea.StringValue(pageUuid)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDingPortalPageScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义工作台页面可见性
//
// @param request - UpdateDingPortalPageScopeRequest
//
// @return UpdateDingPortalPageScopeResponse
func (client *Client) UpdateDingPortalPageScope(pageUuid *string, appUuid *string, request *UpdateDingPortalPageScopeRequest) (_result *UpdateDingPortalPageScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDingPortalPageScopeHeaders{}
	_result = &UpdateDingPortalPageScopeResponse{}
	_body, _err := client.UpdateDingPortalPageScopeWithOptions(pageUuid, appUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
