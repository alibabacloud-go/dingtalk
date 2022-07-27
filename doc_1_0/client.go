// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package doc_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddWorkspaceDocMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceDocMembersHeaders) SetXAcsDingtalkAccessToken(v string) *AddWorkspaceDocMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddWorkspaceDocMembersRequest struct {
	// 被操作用户组
	Members []*AddWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s AddWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequest) SetMembers(v []*AddWorkspaceDocMembersRequestMembers) *AddWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetOperatorId(v string) *AddWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

type AddWorkspaceDocMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 用户权限
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s AddWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberId(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetRoleType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

type AddWorkspaceDocMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

type AddWorkspaceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *AddWorkspaceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddWorkspaceMembersRequest struct {
	// 被操作用户组
	Members []*AddWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s AddWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequest) SetMembers(v []*AddWorkspaceMembersRequestMembers) *AddWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceMembersRequest) SetOperatorId(v string) *AddWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

type AddWorkspaceMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 用户权限
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s AddWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberId(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberType(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetRoleType(v string) *AddWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

type AddWorkspaceMembersResponseBody struct {
	NotInOrgList []*string `json:"notInOrgList,omitempty" xml:"notInOrgList,omitempty" type:"Repeated"`
}

func (s AddWorkspaceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponseBody) SetNotInOrgList(v []*string) *AddWorkspaceMembersResponseBody {
	s.NotInOrgList = v
	return s
}

type AddWorkspaceMembersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceMembersResponse) SetBody(v *AddWorkspaceMembersResponseBody) *AddWorkspaceMembersResponse {
	s.Body = v
	return s
}

type AppendRowsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppendRowsHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppendRowsHeaders) GoString() string {
	return s.String()
}

func (s *AppendRowsHeaders) SetCommonHeaders(v map[string]*string) *AppendRowsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendRowsHeaders) SetXAcsDingtalkAccessToken(v string) *AppendRowsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppendRowsRequest struct {
	// 要追加的值
	Values [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s AppendRowsRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendRowsRequest) GoString() string {
	return s.String()
}

func (s *AppendRowsRequest) SetValues(v [][]*string) *AppendRowsRequest {
	s.Values = v
	return s
}

func (s *AppendRowsRequest) SetOperatorId(v string) *AppendRowsRequest {
	s.OperatorId = &v
	return s
}

type AppendRowsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AppendRowsResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendRowsResponse) GoString() string {
	return s.String()
}

func (s *AppendRowsResponse) SetHeaders(v map[string]*string) *AppendRowsResponse {
	s.Headers = v
	return s
}

type BatchGetWorkspaceDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetWorkspaceDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsHeaders) SetCommonHeaders(v map[string]*string) *BatchGetWorkspaceDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetWorkspaceDocsHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetWorkspaceDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetWorkspaceDocsRequest struct {
	// 查询节点Id
	NodeIds []*string `json:"nodeIds,omitempty" xml:"nodeIds,omitempty" type:"Repeated"`
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s BatchGetWorkspaceDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsRequest) SetNodeIds(v []*string) *BatchGetWorkspaceDocsRequest {
	s.NodeIds = v
	return s
}

func (s *BatchGetWorkspaceDocsRequest) SetOperatorId(v string) *BatchGetWorkspaceDocsRequest {
	s.OperatorId = &v
	return s
}

type BatchGetWorkspaceDocsResponseBody struct {
	Result []*BatchGetWorkspaceDocsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchGetWorkspaceDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsResponseBody) SetResult(v []*BatchGetWorkspaceDocsResponseBodyResult) *BatchGetWorkspaceDocsResponseBody {
	s.Result = v
	return s
}

type BatchGetWorkspaceDocsResponseBodyResult struct {
	HasPermission *bool                                               `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
	NodeBO        *BatchGetWorkspaceDocsResponseBodyResultNodeBO      `json:"nodeBO,omitempty" xml:"nodeBO,omitempty" type:"Struct"`
	WorkspaceBO   *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO `json:"workspaceBO,omitempty" xml:"workspaceBO,omitempty" type:"Struct"`
}

func (s BatchGetWorkspaceDocsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsResponseBodyResult) SetHasPermission(v bool) *BatchGetWorkspaceDocsResponseBodyResult {
	s.HasPermission = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResult) SetNodeBO(v *BatchGetWorkspaceDocsResponseBodyResultNodeBO) *BatchGetWorkspaceDocsResponseBodyResult {
	s.NodeBO = v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResult) SetWorkspaceBO(v *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO) *BatchGetWorkspaceDocsResponseBodyResult {
	s.WorkspaceBO = v
	return s
}

type BatchGetWorkspaceDocsResponseBodyResultNodeBO struct {
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 节点类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 最后编辑时间
	LastEditTime *int64  `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	NodeId       *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	Url          *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchGetWorkspaceDocsResponseBodyResultNodeBO) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsResponseBodyResultNodeBO) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetDeleted(v bool) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.Deleted = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetDocType(v string) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.DocType = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetLastEditTime(v int64) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.LastEditTime = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetName(v string) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.Name = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetNodeId(v string) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.NodeId = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultNodeBO) SetUrl(v string) *BatchGetWorkspaceDocsResponseBodyResultNodeBO {
	s.Url = &v
	return s
}

type BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO) SetName(v string) *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO {
	s.Name = &v
	return s
}

func (s *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO) SetWorkspaceId(v string) *BatchGetWorkspaceDocsResponseBodyResultWorkspaceBO {
	s.WorkspaceId = &v
	return s
}

type BatchGetWorkspaceDocsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetWorkspaceDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetWorkspaceDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspaceDocsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspaceDocsResponse) SetHeaders(v map[string]*string) *BatchGetWorkspaceDocsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetWorkspaceDocsResponse) SetBody(v *BatchGetWorkspaceDocsResponseBody) *BatchGetWorkspaceDocsResponse {
	s.Body = v
	return s
}

type BatchGetWorkspacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetWorkspacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *BatchGetWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetWorkspacesHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetWorkspacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetWorkspacesRequest struct {
	// 是否查询最近访问文档
	IncludeRecent *bool `json:"includeRecent,omitempty" xml:"includeRecent,omitempty"`
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 待查询空间Id
	WorkspaceIds []*string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty" type:"Repeated"`
}

func (s BatchGetWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesRequest) SetIncludeRecent(v bool) *BatchGetWorkspacesRequest {
	s.IncludeRecent = &v
	return s
}

func (s *BatchGetWorkspacesRequest) SetOperatorId(v string) *BatchGetWorkspacesRequest {
	s.OperatorId = &v
	return s
}

func (s *BatchGetWorkspacesRequest) SetWorkspaceIds(v []*string) *BatchGetWorkspacesRequest {
	s.WorkspaceIds = v
	return s
}

type BatchGetWorkspacesResponseBody struct {
	// workspace信息
	Workspaces []*BatchGetWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s BatchGetWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesResponseBody) SetWorkspaces(v []*BatchGetWorkspacesResponseBodyWorkspaces) *BatchGetWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type BatchGetWorkspacesResponseBodyWorkspaces struct {
	// 是否有访问团队空间权限
	HasPermission *bool `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
	// 团队空间信息
	Workspace *BatchGetWorkspacesResponseBodyWorkspacesWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s BatchGetWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesResponseBodyWorkspaces) SetHasPermission(v bool) *BatchGetWorkspacesResponseBodyWorkspaces {
	s.HasPermission = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspaces) SetWorkspace(v *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) *BatchGetWorkspacesResponseBodyWorkspaces {
	s.Workspace = v
	return s
}

type BatchGetWorkspacesResponseBodyWorkspacesWorkspace struct {
	// 团队空间创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 团队空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否全员公开
	OrgPublished *bool `json:"orgPublished,omitempty" xml:"orgPublished,omitempty"`
	// 最近访问列表
	RecentList []*BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// 团队空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchGetWorkspacesResponseBodyWorkspacesWorkspace) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesResponseBodyWorkspacesWorkspace) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetCreateTime(v int64) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.CreateTime = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetName(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.Name = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetOrgPublished(v bool) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.OrgPublished = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetRecentList(v []*BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.RecentList = v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetUrl(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.Url = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspace) SetWorkspaceId(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspace {
	s.WorkspaceId = &v
	return s
}

type BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList struct {
	// 最近编辑时间
	LastEditTime *string `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) SetLastEditTime(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList {
	s.LastEditTime = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) SetName(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList {
	s.Name = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) SetNodeId(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList {
	s.NodeId = &v
	return s
}

func (s *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList) SetUrl(v string) *BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList {
	s.Url = &v
	return s
}

type BatchGetWorkspacesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *BatchGetWorkspacesResponse) SetHeaders(v map[string]*string) *BatchGetWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *BatchGetWorkspacesResponse) SetBody(v *BatchGetWorkspacesResponseBody) *BatchGetWorkspacesResponse {
	s.Body = v
	return s
}

type CreateSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSheetRequest struct {
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetRequest) SetName(v string) *CreateSheetRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetRequest) SetOperatorId(v string) *CreateSheetRequest {
	s.OperatorId = &v
	return s
}

type CreateSheetResponseBody struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建的工作表的名称。当输入参数中的工作表名称在表格中已存在时，可能与输入参数指定的工作表名称不同。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作表可见性
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSheetResponseBody) SetId(v string) *CreateSheetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSheetResponseBody) SetName(v string) *CreateSheetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSheetResponseBody) SetVisibility(v string) *CreateSheetResponseBody {
	s.Visibility = &v
	return s
}

type CreateSheetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSheetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponse) GoString() string {
	return s.String()
}

func (s *CreateSheetResponse) SetHeaders(v map[string]*string) *CreateSheetResponse {
	s.Headers = v
	return s
}

func (s *CreateSheetResponse) SetBody(v *CreateSheetResponseBody) *CreateSheetResponse {
	s.Body = v
	return s
}

type CreateWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateWorkspaceRequest struct {
	// 团队空间描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 团队空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOperatorId(v string) *CreateWorkspaceRequest {
	s.OperatorId = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	// 工作空间描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 工作空间id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetDescription(v string) *CreateWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetName(v string) *CreateWorkspaceResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetUrl(v string) *CreateWorkspaceResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type CreateWorkspaceDocHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateWorkspaceDocHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceDocHeaders) SetXAcsDingtalkAccessToken(v string) *CreateWorkspaceDocHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateWorkspaceDocRequest struct {
	// 文档类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 文档名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 父节点nodeId
	ParentNodeId *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
	// 文档模板id
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s CreateWorkspaceDocRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequest) SetDocType(v string) *CreateWorkspaceDocRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetName(v string) *CreateWorkspaceDocRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetOperatorId(v string) *CreateWorkspaceDocRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetParentNodeId(v string) *CreateWorkspaceDocRequest {
	s.ParentNodeId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateId(v string) *CreateWorkspaceDocRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateType(v string) *CreateWorkspaceDocRequest {
	s.TemplateType = &v
	return s
}

type CreateWorkspaceDocResponseBody struct {
	// 文档docKey
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// 文档Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponseBody) SetDocKey(v string) *CreateWorkspaceDocResponseBody {
	s.DocKey = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetNodeId(v string) *CreateWorkspaceDocResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetUrl(v string) *CreateWorkspaceDocResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetWorkspaceId(v string) *CreateWorkspaceDocResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceDocResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkspaceDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceDocResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponse) SetHeaders(v map[string]*string) *CreateWorkspaceDocResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceDocResponse) SetBody(v *CreateWorkspaceDocResponseBody) *CreateWorkspaceDocResponse {
	s.Body = v
	return s
}

type DeleteSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSheetHeaders) SetCommonHeaders(v map[string]*string) *DeleteSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSheetHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSheetRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSheetRequest) SetOperatorId(v string) *DeleteSheetRequest {
	s.OperatorId = &v
	return s
}

type DeleteSheetResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponse) SetHeaders(v map[string]*string) *DeleteSheetResponse {
	s.Headers = v
	return s
}

type DeleteWorkspaceDocHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteWorkspaceDocHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteWorkspaceDocHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteWorkspaceDocRequest struct {
	// 发起删除请求的用户用户的unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteWorkspaceDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocRequest) SetOperatorId(v string) *DeleteWorkspaceDocRequest {
	s.OperatorId = &v
	return s
}

type DeleteWorkspaceDocResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteWorkspaceDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceDocResponse {
	s.Headers = v
	return s
}

type DeleteWorkspaceDocMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocMembersHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteWorkspaceDocMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteWorkspaceDocMembersRequest struct {
	// 被操作用户组
	Members []*DeleteWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequest) SetMembers(v []*DeleteWorkspaceDocMembersRequestMembers) *DeleteWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetOperatorId(v string) *DeleteWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

type DeleteWorkspaceDocMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

type DeleteWorkspaceDocMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

type DeleteWorkspaceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteWorkspaceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteWorkspaceMembersRequest struct {
	// 被操作用户组
	Members []*DeleteWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequest) SetMembers(v []*DeleteWorkspaceMembersRequestMembers) *DeleteWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetOperatorId(v string) *DeleteWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

type DeleteWorkspaceMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s DeleteWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

type DeleteWorkspaceMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceMembersResponse {
	s.Headers = v
	return s
}

type GetRangeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRangeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRangeHeaders) GoString() string {
	return s.String()
}

func (s *GetRangeHeaders) SetCommonHeaders(v map[string]*string) *GetRangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRangeHeaders) SetXAcsDingtalkAccessToken(v string) *GetRangeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRangeRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRangeRequest) GoString() string {
	return s.String()
}

func (s *GetRangeRequest) SetOperatorId(v string) *GetRangeRequest {
	s.OperatorId = &v
	return s
}

type GetRangeResponseBody struct {
	BackgroundColors [][]*GetRangeResponseBodyBackgroundColors `json:"backgroundColors,omitempty" xml:"backgroundColors,omitempty" type:"Repeated"`
	DisplayValues    [][]*string                               `json:"displayValues,omitempty" xml:"displayValues,omitempty" type:"Repeated"`
	Formulas         [][]*string                               `json:"formulas,omitempty" xml:"formulas,omitempty" type:"Repeated"`
	// 值
	Values [][]interface{} `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBody) SetBackgroundColors(v [][]*GetRangeResponseBodyBackgroundColors) *GetRangeResponseBody {
	s.BackgroundColors = v
	return s
}

func (s *GetRangeResponseBody) SetDisplayValues(v [][]*string) *GetRangeResponseBody {
	s.DisplayValues = v
	return s
}

func (s *GetRangeResponseBody) SetFormulas(v [][]*string) *GetRangeResponseBody {
	s.Formulas = v
	return s
}

func (s *GetRangeResponseBody) SetValues(v [][]interface{}) *GetRangeResponseBody {
	s.Values = v
	return s
}

type GetRangeResponseBodyBackgroundColors struct {
	// RGB值中的红色值
	Red *int64 `json:"red,omitempty" xml:"red,omitempty"`
	// RGB值中的绿色值
	Green *int64 `json:"green,omitempty" xml:"green,omitempty"`
	// RGB值中的蓝色值
	Blue *int64 `json:"blue,omitempty" xml:"blue,omitempty"`
	// 16进制表示的颜色
	HexString *string `json:"hexString,omitempty" xml:"hexString,omitempty"`
}

func (s GetRangeResponseBodyBackgroundColors) String() string {
	return tea.Prettify(s)
}

func (s GetRangeResponseBodyBackgroundColors) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBodyBackgroundColors) SetRed(v int64) *GetRangeResponseBodyBackgroundColors {
	s.Red = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetGreen(v int64) *GetRangeResponseBodyBackgroundColors {
	s.Green = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetBlue(v int64) *GetRangeResponseBodyBackgroundColors {
	s.Blue = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetHexString(v string) *GetRangeResponseBodyBackgroundColors {
	s.HexString = &v
	return s
}

type GetRangeResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRangeResponse) GoString() string {
	return s.String()
}

func (s *GetRangeResponse) SetHeaders(v map[string]*string) *GetRangeResponse {
	s.Headers = v
	return s
}

func (s *GetRangeResponse) SetBody(v *GetRangeResponseBody) *GetRangeResponse {
	s.Body = v
	return s
}

type GetRecentEditDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecentEditDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsHeaders) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsHeaders) SetCommonHeaders(v map[string]*string) *GetRecentEditDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecentEditDocsHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecentEditDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecentEditDocsRequest struct {
	// 查询size
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 发起操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetRecentEditDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsRequest) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsRequest) SetMaxResults(v int32) *GetRecentEditDocsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRecentEditDocsRequest) SetNextToken(v string) *GetRecentEditDocsRequest {
	s.NextToken = &v
	return s
}

func (s *GetRecentEditDocsRequest) SetOperatorId(v string) *GetRecentEditDocsRequest {
	s.OperatorId = &v
	return s
}

type GetRecentEditDocsResponseBody struct {
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 查询结果
	RecentList []*GetRecentEditDocsResponseBodyRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
}

func (s GetRecentEditDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsResponseBody) SetNextToken(v string) *GetRecentEditDocsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRecentEditDocsResponseBody) SetRecentList(v []*GetRecentEditDocsResponseBodyRecentList) *GetRecentEditDocsResponseBody {
	s.RecentList = v
	return s
}

type GetRecentEditDocsResponseBodyRecentList struct {
	// 文档信息
	NodeBO *GetRecentEditDocsResponseBodyRecentListNodeBO `json:"nodeBO,omitempty" xml:"nodeBO,omitempty" type:"Struct"`
	// 团队空间信息
	WorkspaceBO *GetRecentEditDocsResponseBodyRecentListWorkspaceBO `json:"workspaceBO,omitempty" xml:"workspaceBO,omitempty" type:"Struct"`
}

func (s GetRecentEditDocsResponseBodyRecentList) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsResponseBodyRecentList) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsResponseBodyRecentList) SetNodeBO(v *GetRecentEditDocsResponseBodyRecentListNodeBO) *GetRecentEditDocsResponseBodyRecentList {
	s.NodeBO = v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentList) SetWorkspaceBO(v *GetRecentEditDocsResponseBodyRecentListWorkspaceBO) *GetRecentEditDocsResponseBodyRecentList {
	s.WorkspaceBO = v
	return s
}

type GetRecentEditDocsResponseBodyRecentListNodeBO struct {
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 节点类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 是否被删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 内容的最后编辑时间
	LastEditTime *int64 `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// 文档Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档名称
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 文档更新时间，包括重命名、移动、内容编辑等操作都会刷新更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRecentEditDocsResponseBodyRecentListNodeBO) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsResponseBodyRecentListNodeBO) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetCreateTime(v int64) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.CreateTime = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetDocType(v string) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.DocType = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetIsDeleted(v bool) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.IsDeleted = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetLastEditTime(v int64) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.LastEditTime = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetNodeId(v string) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.NodeId = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetNodeName(v string) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.NodeName = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetUpdateTime(v int64) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.UpdateTime = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListNodeBO) SetUrl(v string) *GetRecentEditDocsResponseBodyRecentListNodeBO {
	s.Url = &v
	return s
}

type GetRecentEditDocsResponseBodyRecentListWorkspaceBO struct {
	// 团队空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 团队空间名称
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetRecentEditDocsResponseBodyRecentListWorkspaceBO) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsResponseBodyRecentListWorkspaceBO) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsResponseBodyRecentListWorkspaceBO) SetUrl(v string) *GetRecentEditDocsResponseBodyRecentListWorkspaceBO {
	s.Url = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListWorkspaceBO) SetWorkspaceId(v string) *GetRecentEditDocsResponseBodyRecentListWorkspaceBO {
	s.WorkspaceId = &v
	return s
}

func (s *GetRecentEditDocsResponseBodyRecentListWorkspaceBO) SetWorkspaceName(v string) *GetRecentEditDocsResponseBodyRecentListWorkspaceBO {
	s.WorkspaceName = &v
	return s
}

type GetRecentEditDocsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecentEditDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecentEditDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecentEditDocsResponse) GoString() string {
	return s.String()
}

func (s *GetRecentEditDocsResponse) SetHeaders(v map[string]*string) *GetRecentEditDocsResponse {
	s.Headers = v
	return s
}

func (s *GetRecentEditDocsResponse) SetBody(v *GetRecentEditDocsResponseBody) *GetRecentEditDocsResponse {
	s.Body = v
	return s
}

type GetRecentOpenDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecentOpenDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsHeaders) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsHeaders) SetCommonHeaders(v map[string]*string) *GetRecentOpenDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecentOpenDocsHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecentOpenDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecentOpenDocsRequest struct {
	// 查询size
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 发起操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetRecentOpenDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsRequest) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsRequest) SetMaxResults(v int32) *GetRecentOpenDocsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRecentOpenDocsRequest) SetNextToken(v string) *GetRecentOpenDocsRequest {
	s.NextToken = &v
	return s
}

func (s *GetRecentOpenDocsRequest) SetOperatorId(v string) *GetRecentOpenDocsRequest {
	s.OperatorId = &v
	return s
}

type GetRecentOpenDocsResponseBody struct {
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 查询结果
	RecentList []*GetRecentOpenDocsResponseBodyRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
}

func (s GetRecentOpenDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsResponseBody) SetNextToken(v string) *GetRecentOpenDocsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRecentOpenDocsResponseBody) SetRecentList(v []*GetRecentOpenDocsResponseBodyRecentList) *GetRecentOpenDocsResponseBody {
	s.RecentList = v
	return s
}

type GetRecentOpenDocsResponseBodyRecentList struct {
	// 文档信息
	NodeBO *GetRecentOpenDocsResponseBodyRecentListNodeBO `json:"nodeBO,omitempty" xml:"nodeBO,omitempty" type:"Struct"`
	// 团队空间信息
	WorkspaceBO *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO `json:"workspaceBO,omitempty" xml:"workspaceBO,omitempty" type:"Struct"`
}

func (s GetRecentOpenDocsResponseBodyRecentList) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsResponseBodyRecentList) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsResponseBodyRecentList) SetNodeBO(v *GetRecentOpenDocsResponseBodyRecentListNodeBO) *GetRecentOpenDocsResponseBodyRecentList {
	s.NodeBO = v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentList) SetWorkspaceBO(v *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) *GetRecentOpenDocsResponseBodyRecentList {
	s.WorkspaceBO = v
	return s
}

type GetRecentOpenDocsResponseBodyRecentListNodeBO struct {
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 节点类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 是否被删除
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// 最后编辑时间
	LastOpenTime *int64 `json:"lastOpenTime,omitempty" xml:"lastOpenTime,omitempty"`
	// 文档Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档名称
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 文档更新时间，包括重命名、移动、内容编辑等操作都会刷新更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRecentOpenDocsResponseBodyRecentListNodeBO) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsResponseBodyRecentListNodeBO) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetCreateTime(v int64) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.CreateTime = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetDocType(v string) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.DocType = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetIsDeleted(v bool) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.IsDeleted = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetLastOpenTime(v int64) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.LastOpenTime = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetNodeId(v string) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.NodeId = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetNodeName(v string) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.NodeName = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetUpdateTime(v int64) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.UpdateTime = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListNodeBO) SetUrl(v string) *GetRecentOpenDocsResponseBodyRecentListNodeBO {
	s.Url = &v
	return s
}

type GetRecentOpenDocsResponseBodyRecentListWorkspaceBO struct {
	// 团队空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 团队空间名称
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) SetUrl(v string) *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO {
	s.Url = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) SetWorkspaceId(v string) *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO {
	s.WorkspaceId = &v
	return s
}

func (s *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO) SetWorkspaceName(v string) *GetRecentOpenDocsResponseBodyRecentListWorkspaceBO {
	s.WorkspaceName = &v
	return s
}

type GetRecentOpenDocsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRecentOpenDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecentOpenDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecentOpenDocsResponse) GoString() string {
	return s.String()
}

func (s *GetRecentOpenDocsResponse) SetHeaders(v map[string]*string) *GetRecentOpenDocsResponse {
	s.Headers = v
	return s
}

func (s *GetRecentOpenDocsResponse) SetBody(v *GetRecentOpenDocsResponseBody) *GetRecentOpenDocsResponse {
	s.Body = v
	return s
}

type GetRelatedWorkspacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRelatedWorkspacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *GetRelatedWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelatedWorkspacesHeaders) SetXAcsDingtalkAccessToken(v string) *GetRelatedWorkspacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRelatedWorkspacesRequest struct {
	// 是否查询最近访问文档列表
	IncludeRecent *bool `json:"includeRecent,omitempty" xml:"includeRecent,omitempty"`
	// 发起操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetRelatedWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesRequest) SetIncludeRecent(v bool) *GetRelatedWorkspacesRequest {
	s.IncludeRecent = &v
	return s
}

func (s *GetRelatedWorkspacesRequest) SetOperatorId(v string) *GetRelatedWorkspacesRequest {
	s.OperatorId = &v
	return s
}

type GetRelatedWorkspacesResponseBody struct {
	// 团队空间结果集
	Workspaces []*GetRelatedWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s GetRelatedWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBody) SetWorkspaces(v []*GetRelatedWorkspacesResponseBodyWorkspaces) *GetRelatedWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type GetRelatedWorkspacesResponseBodyWorkspaces struct {
	// 团队空间创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 团队空间是否被删除
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 团队空间名称
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// 团队空间最近访问文档列表
	RecentList []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// 用户的角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 团队空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetRelatedWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetCreateTime(v int64) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetDeleted(v bool) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Deleted = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetName(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetOwner(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Owner = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetRecentList(v []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.RecentList = v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetRole(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Role = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetUrl(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.Url = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *GetRelatedWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

type GetRelatedWorkspacesResponseBodyWorkspacesRecentList struct {
	// 文档最后编辑时间
	LastEditTime *int64 `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRelatedWorkspacesResponseBodyWorkspacesRecentList) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesResponseBodyWorkspacesRecentList) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetLastEditTime(v int64) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.LastEditTime = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetName(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.Name = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetNodeId(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.NodeId = &v
	return s
}

func (s *GetRelatedWorkspacesResponseBodyWorkspacesRecentList) SetUrl(v string) *GetRelatedWorkspacesResponseBodyWorkspacesRecentList {
	s.Url = &v
	return s
}

type GetRelatedWorkspacesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRelatedWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRelatedWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelatedWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesResponse) SetHeaders(v map[string]*string) *GetRelatedWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *GetRelatedWorkspacesResponse) SetBody(v *GetRelatedWorkspacesResponseBody) *GetRelatedWorkspacesResponse {
	s.Body = v
	return s
}

type GetSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSheetHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetHeaders) SetCommonHeaders(v map[string]*string) *GetSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetHeaders) SetXAcsDingtalkAccessToken(v string) *GetSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSheetRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSheetRequest) GoString() string {
	return s.String()
}

func (s *GetSheetRequest) SetOperatorId(v string) *GetSheetRequest {
	s.OperatorId = &v
	return s
}

type GetSheetResponseBody struct {
	// 工作表列数
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// 最后一列非空列的位置，从0开始。表为空时返回-1。
	LastNonEmptyColumn *int64 `json:"lastNonEmptyColumn,omitempty" xml:"lastNonEmptyColumn,omitempty"`
	// 最后一行非空行的位置，从0开始。表为空时返回-1。
	LastNonEmptyRow *int64 `json:"lastNonEmptyRow,omitempty" xml:"lastNonEmptyRow,omitempty"`
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作表行数
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 工作表可见性
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s GetSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSheetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSheetResponseBody) SetColumnCount(v int64) *GetSheetResponseBody {
	s.ColumnCount = &v
	return s
}

func (s *GetSheetResponseBody) SetLastNonEmptyColumn(v int64) *GetSheetResponseBody {
	s.LastNonEmptyColumn = &v
	return s
}

func (s *GetSheetResponseBody) SetLastNonEmptyRow(v int64) *GetSheetResponseBody {
	s.LastNonEmptyRow = &v
	return s
}

func (s *GetSheetResponseBody) SetName(v string) *GetSheetResponseBody {
	s.Name = &v
	return s
}

func (s *GetSheetResponseBody) SetRowCount(v int64) *GetSheetResponseBody {
	s.RowCount = &v
	return s
}

func (s *GetSheetResponseBody) SetVisibility(v string) *GetSheetResponseBody {
	s.Visibility = &v
	return s
}

type GetSheetResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSheetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSheetResponse) GoString() string {
	return s.String()
}

func (s *GetSheetResponse) SetHeaders(v map[string]*string) *GetSheetResponse {
	s.Headers = v
	return s
}

func (s *GetSheetResponse) SetBody(v *GetSheetResponseBody) *GetSheetResponse {
	s.Body = v
	return s
}

type GetTemplateByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTemplateByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetTemplateByIdHeaders) SetCommonHeaders(v map[string]*string) *GetTemplateByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTemplateByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetTemplateByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTemplateByIdRequest struct {
	// 模版归属
	// public_template //公共模板
	// team_template //团队模板
	// user_template //个人模板
	Belong *string `json:"belong,omitempty" xml:"belong,omitempty"`
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetTemplateByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateByIdRequest) SetBelong(v string) *GetTemplateByIdRequest {
	s.Belong = &v
	return s
}

func (s *GetTemplateByIdRequest) SetOperatorId(v string) *GetTemplateByIdRequest {
	s.OperatorId = &v
	return s
}

type GetTemplateByIdResponseBody struct {
	// 模版预览url
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 模版创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 模版对应文档类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 模版id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 模版类型
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// 模版标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 模版修改时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 模版归属空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetTemplateByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateByIdResponseBody) SetCoverUrl(v string) *GetTemplateByIdResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetCreateTime(v int64) *GetTemplateByIdResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetDocType(v string) *GetTemplateByIdResponseBody {
	s.DocType = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetId(v string) *GetTemplateByIdResponseBody {
	s.Id = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetTemplateType(v string) *GetTemplateByIdResponseBody {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetTitle(v string) *GetTemplateByIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetUpdateTime(v int64) *GetTemplateByIdResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateByIdResponseBody) SetWorkspaceId(v string) *GetTemplateByIdResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetTemplateByIdResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateByIdResponse) SetHeaders(v map[string]*string) *GetTemplateByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateByIdResponse) SetBody(v *GetTemplateByIdResponseBody) *GetTemplateByIdResponse {
	s.Body = v
	return s
}

type GetWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWorkspaceResponseBody struct {
	// 团队空间所属企业id
	CorpId    *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	IsDeleted *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetCorpId(v string) *GetWorkspaceResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetIsDeleted(v bool) *GetWorkspaceResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetOwner(v string) *GetWorkspaceResponseBody {
	s.Owner = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetUrl(v string) *GetWorkspaceResponseBody {
	s.Url = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type GetWorkspaceNodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWorkspaceNodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspaceNodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspaceNodeHeaders) SetXAcsDingtalkAccessToken(v string) *GetWorkspaceNodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWorkspaceNodeRequest struct {
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetWorkspaceNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeRequest) SetOperatorId(v string) *GetWorkspaceNodeRequest {
	s.OperatorId = &v
	return s
}

type GetWorkspaceNodeResponseBody struct {
	// 是否有权限
	HasPermission *bool `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
	// 节点信息
	NodeBO *GetWorkspaceNodeResponseBodyNodeBO `json:"nodeBO,omitempty" xml:"nodeBO,omitempty" type:"Struct"`
	// 节点所属团队空间信息
	WorkspaceBO *GetWorkspaceNodeResponseBodyWorkspaceBO `json:"workspaceBO,omitempty" xml:"workspaceBO,omitempty" type:"Struct"`
}

func (s GetWorkspaceNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeResponseBody) SetHasPermission(v bool) *GetWorkspaceNodeResponseBody {
	s.HasPermission = &v
	return s
}

func (s *GetWorkspaceNodeResponseBody) SetNodeBO(v *GetWorkspaceNodeResponseBodyNodeBO) *GetWorkspaceNodeResponseBody {
	s.NodeBO = v
	return s
}

func (s *GetWorkspaceNodeResponseBody) SetWorkspaceBO(v *GetWorkspaceNodeResponseBodyWorkspaceBO) *GetWorkspaceNodeResponseBody {
	s.WorkspaceBO = v
	return s
}

type GetWorkspaceNodeResponseBodyNodeBO struct {
	// 节点类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 最后编辑时间
	LastEditTime *int64 `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// 节点名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetWorkspaceNodeResponseBodyNodeBO) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeResponseBodyNodeBO) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeResponseBodyNodeBO) SetDocType(v string) *GetWorkspaceNodeResponseBodyNodeBO {
	s.DocType = &v
	return s
}

func (s *GetWorkspaceNodeResponseBodyNodeBO) SetLastEditTime(v int64) *GetWorkspaceNodeResponseBodyNodeBO {
	s.LastEditTime = &v
	return s
}

func (s *GetWorkspaceNodeResponseBodyNodeBO) SetName(v string) *GetWorkspaceNodeResponseBodyNodeBO {
	s.Name = &v
	return s
}

func (s *GetWorkspaceNodeResponseBodyNodeBO) SetNodeId(v string) *GetWorkspaceNodeResponseBodyNodeBO {
	s.NodeId = &v
	return s
}

func (s *GetWorkspaceNodeResponseBodyNodeBO) SetUrl(v string) *GetWorkspaceNodeResponseBodyNodeBO {
	s.Url = &v
	return s
}

type GetWorkspaceNodeResponseBodyWorkspaceBO struct {
	// 团队空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspaceNodeResponseBodyWorkspaceBO) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeResponseBodyWorkspaceBO) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeResponseBodyWorkspaceBO) SetName(v string) *GetWorkspaceNodeResponseBodyWorkspaceBO {
	s.Name = &v
	return s
}

func (s *GetWorkspaceNodeResponseBodyWorkspaceBO) SetWorkspaceId(v string) *GetWorkspaceNodeResponseBodyWorkspaceBO {
	s.WorkspaceId = &v
	return s
}

type GetWorkspaceNodeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkspaceNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspaceNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceNodeResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceNodeResponse) SetHeaders(v map[string]*string) *GetWorkspaceNodeResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceNodeResponse) SetBody(v *GetWorkspaceNodeResponseBody) *GetWorkspaceNodeResponse {
	s.Body = v
	return s
}

type InsertBlocksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertBlocksHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksHeaders) GoString() string {
	return s.String()
}

func (s *InsertBlocksHeaders) SetCommonHeaders(v map[string]*string) *InsertBlocksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertBlocksHeaders) SetXAcsDingtalkAccessToken(v string) *InsertBlocksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertBlocksRequest struct {
	// 元素数组
	Blocks []*InsertBlocksRequestBlocks `json:"blocks,omitempty" xml:"blocks,omitempty" type:"Repeated"`
	// 位置信息
	Location *InsertBlocksRequestLocation `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s InsertBlocksRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequest) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequest) SetBlocks(v []*InsertBlocksRequestBlocks) *InsertBlocksRequest {
	s.Blocks = v
	return s
}

func (s *InsertBlocksRequest) SetLocation(v *InsertBlocksRequestLocation) *InsertBlocksRequest {
	s.Location = v
	return s
}

func (s *InsertBlocksRequest) SetOperatorId(v string) *InsertBlocksRequest {
	s.OperatorId = &v
	return s
}

type InsertBlocksRequestBlocks struct {
	// 元素类型标识
	BlockType *string `json:"blockType,omitempty" xml:"blockType,omitempty"`
	// 段落元素
	Paragraph *InsertBlocksRequestBlocksParagraph `json:"paragraph,omitempty" xml:"paragraph,omitempty" type:"Struct"`
}

func (s InsertBlocksRequestBlocks) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocks) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocks) SetBlockType(v string) *InsertBlocksRequestBlocks {
	s.BlockType = &v
	return s
}

func (s *InsertBlocksRequestBlocks) SetParagraph(v *InsertBlocksRequestBlocksParagraph) *InsertBlocksRequestBlocks {
	s.Paragraph = v
	return s
}

type InsertBlocksRequestBlocksParagraph struct {
	// 子节点
	Children []*InsertBlocksRequestBlocksParagraphChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// 段落样式
	Style *InsertBlocksRequestBlocksParagraphStyle `json:"style,omitempty" xml:"style,omitempty" type:"Struct"`
}

func (s InsertBlocksRequestBlocksParagraph) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocksParagraph) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocksParagraph) SetChildren(v []*InsertBlocksRequestBlocksParagraphChildren) *InsertBlocksRequestBlocksParagraph {
	s.Children = v
	return s
}

func (s *InsertBlocksRequestBlocksParagraph) SetStyle(v *InsertBlocksRequestBlocksParagraphStyle) *InsertBlocksRequestBlocksParagraph {
	s.Style = v
	return s
}

type InsertBlocksRequestBlocksParagraphChildren struct {
	// 元素类型
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// 文本元素
	Text *InsertBlocksRequestBlocksParagraphChildrenText `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s InsertBlocksRequestBlocksParagraphChildren) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocksParagraphChildren) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocksParagraphChildren) SetElementType(v string) *InsertBlocksRequestBlocksParagraphChildren {
	s.ElementType = &v
	return s
}

func (s *InsertBlocksRequestBlocksParagraphChildren) SetText(v *InsertBlocksRequestBlocksParagraphChildrenText) *InsertBlocksRequestBlocksParagraphChildren {
	s.Text = v
	return s
}

type InsertBlocksRequestBlocksParagraphChildrenText struct {
	// 文本内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 文字样式
	TextStyle *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle `json:"textStyle,omitempty" xml:"textStyle,omitempty" type:"Struct"`
}

func (s InsertBlocksRequestBlocksParagraphChildrenText) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocksParagraphChildrenText) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocksParagraphChildrenText) SetContent(v string) *InsertBlocksRequestBlocksParagraphChildrenText {
	s.Content = &v
	return s
}

func (s *InsertBlocksRequestBlocksParagraphChildrenText) SetTextStyle(v *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) *InsertBlocksRequestBlocksParagraphChildrenText {
	s.TextStyle = v
	return s
}

type InsertBlocksRequestBlocksParagraphChildrenTextTextStyle struct {
	// 是否加粗
	Bold *bool `json:"bold,omitempty" xml:"bold,omitempty"`
	// 数据类型
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// 字体大小
	FontSize *int32 `json:"fontSize,omitempty" xml:"fontSize,omitempty"`
	// 字体大小单位
	SizeUnit *string `json:"sizeUnit,omitempty" xml:"sizeUnit,omitempty"`
}

func (s InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) SetBold(v bool) *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle {
	s.Bold = &v
	return s
}

func (s *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) SetDataType(v string) *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle {
	s.DataType = &v
	return s
}

func (s *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) SetFontSize(v int32) *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle {
	s.FontSize = &v
	return s
}

func (s *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle) SetSizeUnit(v string) *InsertBlocksRequestBlocksParagraphChildrenTextTextStyle {
	s.SizeUnit = &v
	return s
}

type InsertBlocksRequestBlocksParagraphStyle struct {
	// 标题样式
	HeadingLevel *string `json:"headingLevel,omitempty" xml:"headingLevel,omitempty"`
}

func (s InsertBlocksRequestBlocksParagraphStyle) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestBlocksParagraphStyle) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestBlocksParagraphStyle) SetHeadingLevel(v string) *InsertBlocksRequestBlocksParagraphStyle {
	s.HeadingLevel = &v
	return s
}

type InsertBlocksRequestLocation struct {
	// 头部插入
	Head *bool `json:"head,omitempty" xml:"head,omitempty"`
}

func (s InsertBlocksRequestLocation) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksRequestLocation) GoString() string {
	return s.String()
}

func (s *InsertBlocksRequestLocation) SetHead(v bool) *InsertBlocksRequestLocation {
	s.Head = &v
	return s
}

type InsertBlocksResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s InsertBlocksResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertBlocksResponse) GoString() string {
	return s.String()
}

func (s *InsertBlocksResponse) SetHeaders(v map[string]*string) *InsertBlocksResponse {
	s.Headers = v
	return s
}

type ListTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *ListTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTemplateRequest struct {
	// 查询模版数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 翻页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 模版类型
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRequest) SetMaxResults(v int32) *ListTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateRequest) SetNextToken(v string) *ListTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateRequest) SetOperatorId(v string) *ListTemplateRequest {
	s.OperatorId = &v
	return s
}

func (s *ListTemplateRequest) SetTemplateType(v string) *ListTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplateRequest) SetWorkspaceId(v string) *ListTemplateRequest {
	s.WorkspaceId = &v
	return s
}

type ListTemplateResponseBody struct {
	// 是否还有更多模版
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 后续结果的偏移
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 模版信息列表
	TemplateList []*ListTemplateResponseBodyTemplateList `json:"templateList,omitempty" xml:"templateList,omitempty" type:"Repeated"`
}

func (s ListTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBody) SetHasMore(v bool) *ListTemplateResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListTemplateResponseBody) SetNextToken(v string) *ListTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateResponseBody) SetTemplateList(v []*ListTemplateResponseBodyTemplateList) *ListTemplateResponseBody {
	s.TemplateList = v
	return s
}

type ListTemplateResponseBodyTemplateList struct {
	// 模版预览url
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 模版创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 模版对应文档类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 模版Id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 模版类型
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// 模版标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 模版修改时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 模版归属空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTemplateResponseBodyTemplateList) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBodyTemplateList) SetCoverUrl(v string) *ListTemplateResponseBodyTemplateList {
	s.CoverUrl = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetCreateTime(v int64) *ListTemplateResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetDocType(v string) *ListTemplateResponseBodyTemplateList {
	s.DocType = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetId(v string) *ListTemplateResponseBodyTemplateList {
	s.Id = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetTemplateType(v string) *ListTemplateResponseBodyTemplateList {
	s.TemplateType = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetTitle(v string) *ListTemplateResponseBodyTemplateList {
	s.Title = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetUpdateTime(v int64) *ListTemplateResponseBodyTemplateList {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplateResponseBodyTemplateList) SetWorkspaceId(v string) *ListTemplateResponseBodyTemplateList {
	s.WorkspaceId = &v
	return s
}

type ListTemplateResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateResponse) SetHeaders(v map[string]*string) *ListTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateResponse) SetBody(v *ListTemplateResponseBody) *ListTemplateResponse {
	s.Body = v
	return s
}

type RangeFindNextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RangeFindNextHeaders) String() string {
	return tea.Prettify(s)
}

func (s RangeFindNextHeaders) GoString() string {
	return s.String()
}

func (s *RangeFindNextHeaders) SetCommonHeaders(v map[string]*string) *RangeFindNextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RangeFindNextHeaders) SetXAcsDingtalkAccessToken(v string) *RangeFindNextHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RangeFindNextRequest struct {
	// 查找选项
	FindOptions *RangeFindNextRequestFindOptions `json:"findOptions,omitempty" xml:"findOptions,omitempty" type:"Struct"`
	// 要查找的文本
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s RangeFindNextRequest) String() string {
	return tea.Prettify(s)
}

func (s RangeFindNextRequest) GoString() string {
	return s.String()
}

func (s *RangeFindNextRequest) SetFindOptions(v *RangeFindNextRequestFindOptions) *RangeFindNextRequest {
	s.FindOptions = v
	return s
}

func (s *RangeFindNextRequest) SetText(v string) *RangeFindNextRequest {
	s.Text = &v
	return s
}

func (s *RangeFindNextRequest) SetOperatorId(v string) *RangeFindNextRequest {
	s.OperatorId = &v
	return s
}

type RangeFindNextRequestFindOptions struct {
	// 匹配大小写
	MatchCase *bool `json:"matchCase,omitempty" xml:"matchCase,omitempty"`
	// 匹配整个单元格
	MatchEntireCell *bool `json:"matchEntireCell,omitempty" xml:"matchEntireCell,omitempty"`
	// 在公式内搜索
	MatchFormulaText *bool   `json:"matchFormulaText,omitempty" xml:"matchFormulaText,omitempty"`
	Scope            *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// text是正则表达式
	UseRegExp *bool `json:"useRegExp,omitempty" xml:"useRegExp,omitempty"`
}

func (s RangeFindNextRequestFindOptions) String() string {
	return tea.Prettify(s)
}

func (s RangeFindNextRequestFindOptions) GoString() string {
	return s.String()
}

func (s *RangeFindNextRequestFindOptions) SetMatchCase(v bool) *RangeFindNextRequestFindOptions {
	s.MatchCase = &v
	return s
}

func (s *RangeFindNextRequestFindOptions) SetMatchEntireCell(v bool) *RangeFindNextRequestFindOptions {
	s.MatchEntireCell = &v
	return s
}

func (s *RangeFindNextRequestFindOptions) SetMatchFormulaText(v bool) *RangeFindNextRequestFindOptions {
	s.MatchFormulaText = &v
	return s
}

func (s *RangeFindNextRequestFindOptions) SetScope(v string) *RangeFindNextRequestFindOptions {
	s.Scope = &v
	return s
}

func (s *RangeFindNextRequestFindOptions) SetUseRegExp(v bool) *RangeFindNextRequestFindOptions {
	s.UseRegExp = &v
	return s
}

type RangeFindNextResponseBody struct {
	// 找到的单元格的地址，使用A1表示法
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s RangeFindNextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RangeFindNextResponseBody) GoString() string {
	return s.String()
}

func (s *RangeFindNextResponseBody) SetA1Notation(v string) *RangeFindNextResponseBody {
	s.A1Notation = &v
	return s
}

type RangeFindNextResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RangeFindNextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RangeFindNextResponse) String() string {
	return tea.Prettify(s)
}

func (s RangeFindNextResponse) GoString() string {
	return s.String()
}

func (s *RangeFindNextResponse) SetHeaders(v map[string]*string) *RangeFindNextResponse {
	s.Headers = v
	return s
}

func (s *RangeFindNextResponse) SetBody(v *RangeFindNextResponseBody) *RangeFindNextResponse {
	s.Body = v
	return s
}

type SearchWorkspaceDocsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchWorkspaceDocsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsHeaders) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsHeaders) SetCommonHeaders(v map[string]*string) *SearchWorkspaceDocsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchWorkspaceDocsHeaders) SetXAcsDingtalkAccessToken(v string) *SearchWorkspaceDocsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchWorkspaceDocsRequest struct {
	// 搜索关键字
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 搜索数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 翻页Id
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 发起操作用户unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SearchWorkspaceDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsRequest) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsRequest) SetKeyword(v string) *SearchWorkspaceDocsRequest {
	s.Keyword = &v
	return s
}

func (s *SearchWorkspaceDocsRequest) SetMaxResults(v int32) *SearchWorkspaceDocsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchWorkspaceDocsRequest) SetNextToken(v string) *SearchWorkspaceDocsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchWorkspaceDocsRequest) SetOperatorId(v string) *SearchWorkspaceDocsRequest {
	s.OperatorId = &v
	return s
}

func (s *SearchWorkspaceDocsRequest) SetWorkspaceId(v string) *SearchWorkspaceDocsRequest {
	s.WorkspaceId = &v
	return s
}

type SearchWorkspaceDocsResponseBody struct {
	Docs []*SearchWorkspaceDocsResponseBodyDocs `json:"docs,omitempty" xml:"docs,omitempty" type:"Repeated"`
	// 是否还有可搜索内容
	HasMore   *bool   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s SearchWorkspaceDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsResponseBody) SetDocs(v []*SearchWorkspaceDocsResponseBodyDocs) *SearchWorkspaceDocsResponseBody {
	s.Docs = v
	return s
}

func (s *SearchWorkspaceDocsResponseBody) SetHasMore(v bool) *SearchWorkspaceDocsResponseBody {
	s.HasMore = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBody) SetNextToken(v string) *SearchWorkspaceDocsResponseBody {
	s.NextToken = &v
	return s
}

type SearchWorkspaceDocsResponseBodyDocs struct {
	NodeBO      *SearchWorkspaceDocsResponseBodyDocsNodeBO      `json:"nodeBO,omitempty" xml:"nodeBO,omitempty" type:"Struct"`
	WorkspaceBO *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO `json:"workspaceBO,omitempty" xml:"workspaceBO,omitempty" type:"Struct"`
}

func (s SearchWorkspaceDocsResponseBodyDocs) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsResponseBodyDocs) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsResponseBodyDocs) SetNodeBO(v *SearchWorkspaceDocsResponseBodyDocsNodeBO) *SearchWorkspaceDocsResponseBodyDocs {
	s.NodeBO = v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocs) SetWorkspaceBO(v *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO) *SearchWorkspaceDocsResponseBodyDocs {
	s.WorkspaceBO = v
	return s
}

type SearchWorkspaceDocsResponseBodyDocsNodeBO struct {
	// 节点类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 最近编辑时间
	LastEditTime *int64 `json:"lastEditTime,omitempty" xml:"lastEditTime,omitempty"`
	// 节点名称，如果命中了搜索关键词会包含高亮标签
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 节点Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点原始名称
	OriginName *string `json:"originName,omitempty" xml:"originName,omitempty"`
	// 节点打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchWorkspaceDocsResponseBodyDocsNodeBO) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsResponseBodyDocsNodeBO) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetDocType(v string) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.DocType = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetLastEditTime(v int64) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.LastEditTime = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetName(v string) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.Name = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetNodeId(v string) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.NodeId = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetOriginName(v string) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.OriginName = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsNodeBO) SetUrl(v string) *SearchWorkspaceDocsResponseBodyDocsNodeBO {
	s.Url = &v
	return s
}

type SearchWorkspaceDocsResponseBodyDocsWorkspaceBO struct {
	// 团队空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SearchWorkspaceDocsResponseBodyDocsWorkspaceBO) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsResponseBodyDocsWorkspaceBO) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO) SetName(v string) *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO {
	s.Name = &v
	return s
}

func (s *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO) SetWorkspaceId(v string) *SearchWorkspaceDocsResponseBodyDocsWorkspaceBO {
	s.WorkspaceId = &v
	return s
}

type SearchWorkspaceDocsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchWorkspaceDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchWorkspaceDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkspaceDocsResponse) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceDocsResponse) SetHeaders(v map[string]*string) *SearchWorkspaceDocsResponse {
	s.Headers = v
	return s
}

func (s *SearchWorkspaceDocsResponse) SetBody(v *SearchWorkspaceDocsResponseBody) *SearchWorkspaceDocsResponse {
	s.Body = v
	return s
}

type UpdateRangeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRangeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRangeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRangeHeaders) SetCommonHeaders(v map[string]*string) *UpdateRangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRangeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRangeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRangeRequest struct {
	// 背景色
	BackgroundColors [][]*string                       `json:"backgroundColors,omitempty" xml:"backgroundColors,omitempty" type:"Repeated"`
	Hyperlinks       [][]*UpdateRangeRequestHyperlinks `json:"hyperlinks,omitempty" xml:"hyperlinks,omitempty" type:"Repeated"`
	// 值
	Values [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRangeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRangeRequest) SetBackgroundColors(v [][]*string) *UpdateRangeRequest {
	s.BackgroundColors = v
	return s
}

func (s *UpdateRangeRequest) SetHyperlinks(v [][]*UpdateRangeRequestHyperlinks) *UpdateRangeRequest {
	s.Hyperlinks = v
	return s
}

func (s *UpdateRangeRequest) SetValues(v [][]*string) *UpdateRangeRequest {
	s.Values = v
	return s
}

func (s *UpdateRangeRequest) SetOperatorId(v string) *UpdateRangeRequest {
	s.OperatorId = &v
	return s
}

type UpdateRangeRequestHyperlinks struct {
	// 超链接类型，可选 "path", "sheet", "range"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 链接文本
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s UpdateRangeRequestHyperlinks) String() string {
	return tea.Prettify(s)
}

func (s UpdateRangeRequestHyperlinks) GoString() string {
	return s.String()
}

func (s *UpdateRangeRequestHyperlinks) SetType(v string) *UpdateRangeRequestHyperlinks {
	s.Type = &v
	return s
}

func (s *UpdateRangeRequestHyperlinks) SetLink(v string) *UpdateRangeRequestHyperlinks {
	s.Link = &v
	return s
}

func (s *UpdateRangeRequestHyperlinks) SetText(v string) *UpdateRangeRequestHyperlinks {
	s.Text = &v
	return s
}

type UpdateRangeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRangeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRangeResponse) SetHeaders(v map[string]*string) *UpdateRangeResponse {
	s.Headers = v
	return s
}

type UpdateSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSheetHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSheetHeaders) SetCommonHeaders(v map[string]*string) *UpdateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSheetHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSheetRequest struct {
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作表可见性
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSheetRequest) GoString() string {
	return s.String()
}

func (s *UpdateSheetRequest) SetName(v string) *UpdateSheetRequest {
	s.Name = &v
	return s
}

func (s *UpdateSheetRequest) SetVisibility(v string) *UpdateSheetRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateSheetRequest) SetOperatorId(v string) *UpdateSheetRequest {
	s.OperatorId = &v
	return s
}

type UpdateSheetResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSheetResponse) GoString() string {
	return s.String()
}

func (s *UpdateSheetResponse) SetHeaders(v map[string]*string) *UpdateSheetResponse {
	s.Headers = v
	return s
}

type UpdateWorkspaceDocMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceDocMembersHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateWorkspaceDocMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateWorkspaceDocMembersRequest struct {
	// 被操作用户组
	Members []*UpdateWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequest) SetMembers(v []*UpdateWorkspaceDocMembersRequestMembers) *UpdateWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetOperatorId(v string) *UpdateWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

type UpdateWorkspaceDocMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 用户权限
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

type UpdateWorkspaceDocMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

type UpdateWorkspaceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateWorkspaceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateWorkspaceMembersRequest struct {
	// 被操作用户组
	Members []*UpdateWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s UpdateWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequest) SetMembers(v []*UpdateWorkspaceMembersRequestMembers) *UpdateWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetOperatorId(v string) *UpdateWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

type UpdateWorkspaceMembersRequestMembers struct {
	// 被操作用户unionId
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 用户类型
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 用户权限
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s UpdateWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

type UpdateWorkspaceMembersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceMembersResponse {
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

func (client *Client) AddWorkspaceDocMembers(workspaceId *string, nodeId *string, request *AddWorkspaceDocMembersRequest) (_result *AddWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddWorkspaceDocMembersHeaders{}
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.AddWorkspaceDocMembersWithOptions(workspaceId, nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceDocMembersWithOptions(workspaceId *string, nodeId *string, request *AddWorkspaceDocMembersRequest, headers *AddWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *AddWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	nodeId = openapiutil.GetEncodeParam(nodeId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWorkspaceMembers(workspaceId *string, request *AddWorkspaceMembersRequest) (_result *AddWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddWorkspaceMembersHeaders{}
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.AddWorkspaceMembersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceMembersWithOptions(workspaceId *string, request *AddWorkspaceMembersRequest, headers *AddWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *AddWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendRows(workbookId *string, sheetId *string, request *AppendRowsRequest) (_result *AppendRowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppendRowsHeaders{}
	_result = &AppendRowsResponse{}
	_body, _err := client.AppendRowsWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendRowsWithOptions(workbookId *string, sheetId *string, request *AppendRowsRequest, headers *AppendRowsHeaders, runtime *util.RuntimeOptions) (_result *AppendRowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Values)) {
		body["values"] = request.Values
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
	_result = &AppendRowsResponse{}
	_body, _err := client.DoROARequest(tea.String("AppendRows"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/appendRows"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetWorkspaceDocs(request *BatchGetWorkspaceDocsRequest) (_result *BatchGetWorkspaceDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetWorkspaceDocsHeaders{}
	_result = &BatchGetWorkspaceDocsResponse{}
	_body, _err := client.BatchGetWorkspaceDocsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetWorkspaceDocsWithOptions(request *BatchGetWorkspaceDocsRequest, headers *BatchGetWorkspaceDocsHeaders, runtime *util.RuntimeOptions) (_result *BatchGetWorkspaceDocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		body["nodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &BatchGetWorkspaceDocsResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetWorkspaceDocs"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/docs/infos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetWorkspaces(request *BatchGetWorkspacesRequest) (_result *BatchGetWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetWorkspacesHeaders{}
	_result = &BatchGetWorkspacesResponse{}
	_body, _err := client.BatchGetWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetWorkspacesWithOptions(request *BatchGetWorkspacesRequest, headers *BatchGetWorkspacesHeaders, runtime *util.RuntimeOptions) (_result *BatchGetWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeRecent)) {
		body["includeRecent"] = request.IncludeRecent
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		body["workspaceIds"] = request.WorkspaceIds
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
	_result = &BatchGetWorkspacesResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetWorkspaces"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/infos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSheet(workbookId *string, request *CreateSheetRequest) (_result *CreateSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSheetHeaders{}
	_result = &CreateSheetResponse{}
	_body, _err := client.CreateSheetWithOptions(workbookId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSheetWithOptions(workbookId *string, request *CreateSheetRequest, headers *CreateSheetHeaders, runtime *util.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
	_result = &CreateSheetResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSheet"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkspaceHeaders{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers *CreateWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateWorkspace"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspaceDoc(workspaceId *string, request *CreateWorkspaceDocRequest) (_result *CreateWorkspaceDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkspaceDocHeaders{}
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.CreateWorkspaceDocWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceDocWithOptions(workspaceId *string, request *CreateWorkspaceDocRequest, headers *CreateWorkspaceDocHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkspaceDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["docType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentNodeId)) {
		body["parentNodeId"] = request.ParentNodeId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["templateType"] = request.TemplateType
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
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateWorkspaceDoc"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSheet(workbookId *string, sheetId *string, request *DeleteSheetRequest) (_result *DeleteSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSheetHeaders{}
	_result = &DeleteSheetResponse{}
	_body, _err := client.DeleteSheetWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSheetWithOptions(workbookId *string, sheetId *string, request *DeleteSheetRequest, headers *DeleteSheetHeaders, runtime *util.RuntimeOptions) (_result *DeleteSheetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
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
	_result = &DeleteSheetResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSheet"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceDoc(workspaceId *string, nodeId *string, request *DeleteWorkspaceDocRequest) (_result *DeleteWorkspaceDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWorkspaceDocHeaders{}
	_result = &DeleteWorkspaceDocResponse{}
	_body, _err := client.DeleteWorkspaceDocWithOptions(workspaceId, nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceDocWithOptions(workspaceId *string, nodeId *string, request *DeleteWorkspaceDocRequest, headers *DeleteWorkspaceDocHeaders, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	nodeId = openapiutil.GetEncodeParam(nodeId)
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
	_result = &DeleteWorkspaceDocResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteWorkspaceDoc"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceDocMembers(workspaceId *string, nodeId *string, request *DeleteWorkspaceDocMembersRequest) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWorkspaceDocMembersHeaders{}
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.DeleteWorkspaceDocMembersWithOptions(workspaceId, nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceDocMembersWithOptions(workspaceId *string, nodeId *string, request *DeleteWorkspaceDocMembersRequest, headers *DeleteWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	nodeId = openapiutil.GetEncodeParam(nodeId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members/remove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceMembers(workspaceId *string, request *DeleteWorkspaceMembersRequest) (_result *DeleteWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWorkspaceMembersHeaders{}
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.DeleteWorkspaceMembersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceMembersWithOptions(workspaceId *string, request *DeleteWorkspaceMembersRequest, headers *DeleteWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members/remove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRange(workbookId *string, sheetId *string, rangeAddress *string, request *GetRangeRequest) (_result *GetRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRangeHeaders{}
	_result = &GetRangeResponse{}
	_body, _err := client.GetRangeWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRangeWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *GetRangeRequest, headers *GetRangeHeaders, runtime *util.RuntimeOptions) (_result *GetRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	rangeAddress = openapiutil.GetEncodeParam(rangeAddress)
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
	_result = &GetRangeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRange"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecentEditDocs(request *GetRecentEditDocsRequest) (_result *GetRecentEditDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecentEditDocsHeaders{}
	_result = &GetRecentEditDocsResponse{}
	_body, _err := client.GetRecentEditDocsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecentEditDocsWithOptions(request *GetRecentEditDocsRequest, headers *GetRecentEditDocsHeaders, runtime *util.RuntimeOptions) (_result *GetRecentEditDocsResponse, _err error) {
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
	_result = &GetRecentEditDocsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRecentEditDocs"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/docs/recentEditDocs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecentOpenDocs(request *GetRecentOpenDocsRequest) (_result *GetRecentOpenDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecentOpenDocsHeaders{}
	_result = &GetRecentOpenDocsResponse{}
	_body, _err := client.GetRecentOpenDocsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecentOpenDocsWithOptions(request *GetRecentOpenDocsRequest, headers *GetRecentOpenDocsHeaders, runtime *util.RuntimeOptions) (_result *GetRecentOpenDocsResponse, _err error) {
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
	_result = &GetRecentOpenDocsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRecentOpenDocs"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/docs/recentOpenDocs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRelatedWorkspaces(request *GetRelatedWorkspacesRequest) (_result *GetRelatedWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelatedWorkspacesHeaders{}
	_result = &GetRelatedWorkspacesResponse{}
	_body, _err := client.GetRelatedWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRelatedWorkspacesWithOptions(request *GetRelatedWorkspacesRequest, headers *GetRelatedWorkspacesHeaders, runtime *util.RuntimeOptions) (_result *GetRelatedWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeRecent)) {
		query["includeRecent"] = request.IncludeRecent
	}

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
	_result = &GetRelatedWorkspacesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRelatedWorkspaces"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workspaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSheet(workbookId *string, sheetId *string, request *GetSheetRequest) (_result *GetSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSheetHeaders{}
	_result = &GetSheetResponse{}
	_body, _err := client.GetSheetWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSheetWithOptions(workbookId *string, sheetId *string, request *GetSheetRequest, headers *GetSheetHeaders, runtime *util.RuntimeOptions) (_result *GetSheetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
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
	_result = &GetSheetResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSheet"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplateById(templateId *string, request *GetTemplateByIdRequest) (_result *GetTemplateByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTemplateByIdHeaders{}
	_result = &GetTemplateByIdResponse{}
	_body, _err := client.GetTemplateByIdWithOptions(templateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateByIdWithOptions(templateId *string, request *GetTemplateByIdRequest, headers *GetTemplateByIdHeaders, runtime *util.RuntimeOptions) (_result *GetTemplateByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	templateId = openapiutil.GetEncodeParam(templateId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Belong)) {
		query["belong"] = request.Belong
	}

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
	_result = &GetTemplateByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTemplateById"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/templates/"+tea.StringValue(templateId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(workspaceId *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkspaceHeaders{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(workspaceId *string, headers *GetWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
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
	_result = &GetWorkspaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetWorkspace"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspaceNode(workspaceId *string, nodeId *string, request *GetWorkspaceNodeRequest) (_result *GetWorkspaceNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkspaceNodeHeaders{}
	_result = &GetWorkspaceNodeResponse{}
	_body, _err := client.GetWorkspaceNodeWithOptions(workspaceId, nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceNodeWithOptions(workspaceId *string, nodeId *string, request *GetWorkspaceNodeRequest, headers *GetWorkspaceNodeHeaders, runtime *util.RuntimeOptions) (_result *GetWorkspaceNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	nodeId = openapiutil.GetEncodeParam(nodeId)
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
	_result = &GetWorkspaceNodeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetWorkspaceNode"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertBlocks(documentId *string, request *InsertBlocksRequest) (_result *InsertBlocksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertBlocksHeaders{}
	_result = &InsertBlocksResponse{}
	_body, _err := client.InsertBlocksWithOptions(documentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertBlocksWithOptions(documentId *string, request *InsertBlocksRequest, headers *InsertBlocksHeaders, runtime *util.RuntimeOptions) (_result *InsertBlocksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	documentId = openapiutil.GetEncodeParam(documentId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Blocks)) {
		body["blocks"] = request.Blocks
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Location))) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &InsertBlocksResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertBlocks"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/documents/"+tea.StringValue(documentId)+"/blocks"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplate(request *ListTemplateRequest) (_result *ListTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTemplateHeaders{}
	_result = &ListTemplateResponse{}
	_body, _err := client.ListTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplateWithOptions(request *ListTemplateRequest, headers *ListTemplateHeaders, runtime *util.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["templateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
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
	_result = &ListTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTemplate"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/templates"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RangeFindNext(workbookId *string, sheetId *string, rangeAddress *string, request *RangeFindNextRequest) (_result *RangeFindNextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RangeFindNextHeaders{}
	_result = &RangeFindNextResponse{}
	_body, _err := client.RangeFindNextWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RangeFindNextWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *RangeFindNextRequest, headers *RangeFindNextHeaders, runtime *util.RuntimeOptions) (_result *RangeFindNextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	rangeAddress = openapiutil.GetEncodeParam(rangeAddress)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FindOptions))) {
		body["findOptions"] = request.FindOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
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
	_result = &RangeFindNextResponse{}
	_body, _err := client.DoROARequest(tea.String("RangeFindNext"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/findNext"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchWorkspaceDocs(request *SearchWorkspaceDocsRequest) (_result *SearchWorkspaceDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchWorkspaceDocsHeaders{}
	_result = &SearchWorkspaceDocsResponse{}
	_body, _err := client.SearchWorkspaceDocsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchWorkspaceDocsWithOptions(request *SearchWorkspaceDocsRequest, headers *SearchWorkspaceDocsHeaders, runtime *util.RuntimeOptions) (_result *SearchWorkspaceDocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
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
	_result = &SearchWorkspaceDocsResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchWorkspaceDocs"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/docs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRange(workbookId *string, sheetId *string, rangeAddress *string, request *UpdateRangeRequest) (_result *UpdateRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRangeHeaders{}
	_result = &UpdateRangeResponse{}
	_body, _err := client.UpdateRangeWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRangeWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *UpdateRangeRequest, headers *UpdateRangeHeaders, runtime *util.RuntimeOptions) (_result *UpdateRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	rangeAddress = openapiutil.GetEncodeParam(rangeAddress)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackgroundColors)) {
		body["backgroundColors"] = request.BackgroundColors
	}

	if !tea.BoolValue(util.IsUnset(request.Hyperlinks)) {
		body["hyperlinks"] = request.Hyperlinks
	}

	if !tea.BoolValue(util.IsUnset(request.Values)) {
		body["values"] = request.Values
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
	_result = &UpdateRangeResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateRange"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSheet(workbookId *string, sheetId *string, request *UpdateSheetRequest) (_result *UpdateSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSheetHeaders{}
	_result = &UpdateSheetResponse{}
	_body, _err := client.UpdateSheetWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSheetWithOptions(workbookId *string, sheetId *string, request *UpdateSheetRequest, headers *UpdateSheetHeaders, runtime *util.RuntimeOptions) (_result *UpdateSheetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		body["visibility"] = request.Visibility
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
	_result = &UpdateSheetResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSheet"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceDocMembers(workspaceId *string, nodeId *string, request *UpdateWorkspaceDocMembersRequest) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateWorkspaceDocMembersHeaders{}
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.UpdateWorkspaceDocMembersWithOptions(workspaceId, nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceDocMembersWithOptions(workspaceId *string, nodeId *string, request *UpdateWorkspaceDocMembersRequest, headers *UpdateWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	nodeId = openapiutil.GetEncodeParam(nodeId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceMembers(workspaceId *string, request *UpdateWorkspaceMembersRequest) (_result *UpdateWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateWorkspaceMembersHeaders{}
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.UpdateWorkspaceMembersWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceMembersWithOptions(workspaceId *string, request *UpdateWorkspaceMembersRequest, headers *UpdateWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workspaceId = openapiutil.GetEncodeParam(workspaceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
