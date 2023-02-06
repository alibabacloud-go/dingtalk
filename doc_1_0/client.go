// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package doc_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
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
	// 要追加的值(二维数组)
	// 最大size:
	// 	1000
	Values [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 操作人id
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
	// 待查询知识库id。
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
	// 是否有访问知识库权限。
	HasPermission *bool `json:"hasPermission,omitempty" xml:"hasPermission,omitempty"`
	// 知识库信息。
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
	// 知识库创建时间。
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否全员公开
	OrgPublished *bool `json:"orgPublished,omitempty" xml:"orgPublished,omitempty"`
	// 最近访问列表
	RecentList []*BatchGetWorkspacesResponseBodyWorkspacesWorkspaceRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// 知识库打开url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 知识库id。
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

type ClearHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ClearHeaders) String() string {
	return tea.Prettify(s)
}

func (s ClearHeaders) GoString() string {
	return s.String()
}

func (s *ClearHeaders) SetCommonHeaders(v map[string]*string) *ClearHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearHeaders) SetXAcsDingtalkAccessToken(v string) *ClearHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ClearRequest struct {
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ClearRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearRequest) GoString() string {
	return s.String()
}

func (s *ClearRequest) SetOperatorId(v string) *ClearRequest {
	s.OperatorId = &v
	return s
}

type ClearResponseBody struct {
	// 使用A1表示法的Range地址
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s ClearResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearResponseBody) GoString() string {
	return s.String()
}

func (s *ClearResponseBody) SetA1Notation(v string) *ClearResponseBody {
	s.A1Notation = &v
	return s
}

type ClearResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearResponse) GoString() string {
	return s.String()
}

func (s *ClearResponse) SetHeaders(v map[string]*string) *ClearResponse {
	s.Headers = v
	return s
}

func (s *ClearResponse) SetBody(v *ClearResponseBody) *ClearResponse {
	s.Body = v
	return s
}

type ClearDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ClearDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s ClearDataHeaders) GoString() string {
	return s.String()
}

func (s *ClearDataHeaders) SetCommonHeaders(v map[string]*string) *ClearDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearDataHeaders) SetXAcsDingtalkAccessToken(v string) *ClearDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ClearDataRequest struct {
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ClearDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearDataRequest) GoString() string {
	return s.String()
}

func (s *ClearDataRequest) SetOperatorId(v string) *ClearDataRequest {
	s.OperatorId = &v
	return s
}

type ClearDataResponseBody struct {
	// 使用A1表示法的Range地址
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s ClearDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearDataResponseBody) GoString() string {
	return s.String()
}

func (s *ClearDataResponseBody) SetA1Notation(v string) *ClearDataResponseBody {
	s.A1Notation = &v
	return s
}

type ClearDataResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearDataResponse) GoString() string {
	return s.String()
}

func (s *ClearDataResponse) SetHeaders(v map[string]*string) *ClearDataResponse {
	s.Headers = v
	return s
}

func (s *ClearDataResponse) SetBody(v *ClearDataResponseBody) *ClearDataResponse {
	s.Body = v
	return s
}

type CreateConditionalFormattingRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateConditionalFormattingRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleHeaders) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleHeaders) SetCommonHeaders(v map[string]*string) *CreateConditionalFormattingRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateConditionalFormattingRuleHeaders) SetXAcsDingtalkAccessToken(v string) *CreateConditionalFormattingRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateConditionalFormattingRuleRequest struct {
	// 设定当前配置的规则的单元格样式
	CellStyle *CreateConditionalFormattingRuleRequestCellStyle `json:"cellStyle,omitempty" xml:"cellStyle,omitempty" type:"Struct"`
	// 重复值规则
	DuplicateCondition *CreateConditionalFormattingRuleRequestDuplicateCondition `json:"duplicateCondition,omitempty" xml:"duplicateCondition,omitempty" type:"Struct"`
	// 条件格式生效的区域。
	Ranges []*string `json:"ranges,omitempty" xml:"ranges,omitempty" type:"Repeated"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateConditionalFormattingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleRequest) SetCellStyle(v *CreateConditionalFormattingRuleRequestCellStyle) *CreateConditionalFormattingRuleRequest {
	s.CellStyle = v
	return s
}

func (s *CreateConditionalFormattingRuleRequest) SetDuplicateCondition(v *CreateConditionalFormattingRuleRequestDuplicateCondition) *CreateConditionalFormattingRuleRequest {
	s.DuplicateCondition = v
	return s
}

func (s *CreateConditionalFormattingRuleRequest) SetRanges(v []*string) *CreateConditionalFormattingRuleRequest {
	s.Ranges = v
	return s
}

func (s *CreateConditionalFormattingRuleRequest) SetOperatorId(v string) *CreateConditionalFormattingRuleRequest {
	s.OperatorId = &v
	return s
}

type CreateConditionalFormattingRuleRequestCellStyle struct {
	// 背景色，使用十六进制颜色表示法，如#ff0000
	BackgroundColor *string `json:"backgroundColor,omitempty" xml:"backgroundColor,omitempty"`
}

func (s CreateConditionalFormattingRuleRequestCellStyle) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleRequestCellStyle) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleRequestCellStyle) SetBackgroundColor(v string) *CreateConditionalFormattingRuleRequestCellStyle {
	s.BackgroundColor = &v
	return s
}

type CreateConditionalFormattingRuleRequestDuplicateCondition struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s CreateConditionalFormattingRuleRequestDuplicateCondition) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleRequestDuplicateCondition) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleRequestDuplicateCondition) SetOperator(v string) *CreateConditionalFormattingRuleRequestDuplicateCondition {
	s.Operator = &v
	return s
}

type CreateConditionalFormattingRuleResponseBody struct {
	// 条件格式ID
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateConditionalFormattingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleResponseBody) SetId(v string) *CreateConditionalFormattingRuleResponseBody {
	s.Id = &v
	return s
}

type CreateConditionalFormattingRuleResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConditionalFormattingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConditionalFormattingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConditionalFormattingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateConditionalFormattingRuleResponse) SetHeaders(v map[string]*string) *CreateConditionalFormattingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateConditionalFormattingRuleResponse) SetBody(v *CreateConditionalFormattingRuleResponseBody) *CreateConditionalFormattingRuleResponse {
	s.Body = v
	return s
}

type CreateDeveloperMetadataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeveloperMetadataHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataHeaders) SetCommonHeaders(v map[string]*string) *CreateDeveloperMetadataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeveloperMetadataHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeveloperMetadataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeveloperMetadataRequest struct {
	// 元数据所关联到的列
	AssociatedColumn *CreateDeveloperMetadataRequestAssociatedColumn `json:"associatedColumn,omitempty" xml:"associatedColumn,omitempty" type:"Struct"`
	// 元数据所关联到的行
	AssociatedRow *CreateDeveloperMetadataRequestAssociatedRow `json:"associatedRow,omitempty" xml:"associatedRow,omitempty" type:"Struct"`
	// 元数据值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateDeveloperMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataRequest) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataRequest) SetAssociatedColumn(v *CreateDeveloperMetadataRequestAssociatedColumn) *CreateDeveloperMetadataRequest {
	s.AssociatedColumn = v
	return s
}

func (s *CreateDeveloperMetadataRequest) SetAssociatedRow(v *CreateDeveloperMetadataRequestAssociatedRow) *CreateDeveloperMetadataRequest {
	s.AssociatedRow = v
	return s
}

func (s *CreateDeveloperMetadataRequest) SetValue(v string) *CreateDeveloperMetadataRequest {
	s.Value = &v
	return s
}

func (s *CreateDeveloperMetadataRequest) SetOperatorId(v string) *CreateDeveloperMetadataRequest {
	s.OperatorId = &v
	return s
}

type CreateDeveloperMetadataRequestAssociatedColumn struct {
	// 列号，从0开始
	Column *int32  `json:"column,omitempty" xml:"column,omitempty"`
	Sheet  *string `json:"sheet,omitempty" xml:"sheet,omitempty"`
}

func (s CreateDeveloperMetadataRequestAssociatedColumn) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataRequestAssociatedColumn) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataRequestAssociatedColumn) SetColumn(v int32) *CreateDeveloperMetadataRequestAssociatedColumn {
	s.Column = &v
	return s
}

func (s *CreateDeveloperMetadataRequestAssociatedColumn) SetSheet(v string) *CreateDeveloperMetadataRequestAssociatedColumn {
	s.Sheet = &v
	return s
}

type CreateDeveloperMetadataRequestAssociatedRow struct {
	// 行号，从0开始
	Row *int32 `json:"row,omitempty" xml:"row,omitempty"`
	// 工作表ID或名称
	Sheet *string `json:"sheet,omitempty" xml:"sheet,omitempty"`
}

func (s CreateDeveloperMetadataRequestAssociatedRow) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataRequestAssociatedRow) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataRequestAssociatedRow) SetRow(v int32) *CreateDeveloperMetadataRequestAssociatedRow {
	s.Row = &v
	return s
}

func (s *CreateDeveloperMetadataRequestAssociatedRow) SetSheet(v string) *CreateDeveloperMetadataRequestAssociatedRow {
	s.Sheet = &v
	return s
}

type CreateDeveloperMetadataResponseBody struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateDeveloperMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataResponseBody) SetId(v string) *CreateDeveloperMetadataResponseBody {
	s.Id = &v
	return s
}

type CreateDeveloperMetadataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeveloperMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeveloperMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeveloperMetadataResponse) GoString() string {
	return s.String()
}

func (s *CreateDeveloperMetadataResponse) SetHeaders(v map[string]*string) *CreateDeveloperMetadataResponse {
	s.Headers = v
	return s
}

func (s *CreateDeveloperMetadataResponse) SetBody(v *CreateDeveloperMetadataResponseBody) *CreateDeveloperMetadataResponse {
	s.Body = v
	return s
}

type CreateRangeProtectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRangeProtectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRangeProtectionHeaders) GoString() string {
	return s.String()
}

func (s *CreateRangeProtectionHeaders) SetCommonHeaders(v map[string]*string) *CreateRangeProtectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRangeProtectionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRangeProtectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRangeProtectionRequest struct {
	// 对于拥有「可编辑」权限的用户的细化权限配置。
	EditableSetting *CreateRangeProtectionRequestEditableSetting `json:"editableSetting,omitempty" xml:"editableSetting,omitempty" type:"Struct"`
	// 其它用户的权限
	OtherUserPermission *string `json:"otherUserPermission,omitempty" xml:"otherUserPermission,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateRangeProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRangeProtectionRequest) GoString() string {
	return s.String()
}

func (s *CreateRangeProtectionRequest) SetEditableSetting(v *CreateRangeProtectionRequestEditableSetting) *CreateRangeProtectionRequest {
	s.EditableSetting = v
	return s
}

func (s *CreateRangeProtectionRequest) SetOtherUserPermission(v string) *CreateRangeProtectionRequest {
	s.OtherUserPermission = &v
	return s
}

func (s *CreateRangeProtectionRequest) SetOperatorId(v string) *CreateRangeProtectionRequest {
	s.OperatorId = &v
	return s
}

type CreateRangeProtectionRequestEditableSetting struct {
	// 是否可删除列
	DeleteColumns *bool `json:"deleteColumns,omitempty" xml:"deleteColumns,omitempty"`
	// 是否可删除行
	DeleteRows *bool `json:"deleteRows,omitempty" xml:"deleteRows,omitempty"`
	// 是否可修改单元格的值
	EditCells *bool `json:"editCells,omitempty" xml:"editCells,omitempty"`
	// 是否可修改单元格样式
	FormatCells *bool `json:"formatCells,omitempty" xml:"formatCells,omitempty"`
	// 是否可插入列
	InsertColumns *bool `json:"insertColumns,omitempty" xml:"insertColumns,omitempty"`
	// 是否可插入行
	InsertRows *bool `json:"insertRows,omitempty" xml:"insertRows,omitempty"`
}

func (s CreateRangeProtectionRequestEditableSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateRangeProtectionRequestEditableSetting) GoString() string {
	return s.String()
}

func (s *CreateRangeProtectionRequestEditableSetting) SetDeleteColumns(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.DeleteColumns = &v
	return s
}

func (s *CreateRangeProtectionRequestEditableSetting) SetDeleteRows(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.DeleteRows = &v
	return s
}

func (s *CreateRangeProtectionRequestEditableSetting) SetEditCells(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.EditCells = &v
	return s
}

func (s *CreateRangeProtectionRequestEditableSetting) SetFormatCells(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.FormatCells = &v
	return s
}

func (s *CreateRangeProtectionRequestEditableSetting) SetInsertColumns(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.InsertColumns = &v
	return s
}

func (s *CreateRangeProtectionRequestEditableSetting) SetInsertRows(v bool) *CreateRangeProtectionRequestEditableSetting {
	s.InsertRows = &v
	return s
}

type CreateRangeProtectionResponseBody struct {
	// 单元格锁定ID
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateRangeProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRangeProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRangeProtectionResponseBody) SetId(v string) *CreateRangeProtectionResponseBody {
	s.Id = &v
	return s
}

type CreateRangeProtectionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRangeProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRangeProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRangeProtectionResponse) GoString() string {
	return s.String()
}

func (s *CreateRangeProtectionResponse) SetHeaders(v map[string]*string) *CreateRangeProtectionResponse {
	s.Headers = v
	return s
}

func (s *CreateRangeProtectionResponse) SetBody(v *CreateRangeProtectionResponseBody) *CreateRangeProtectionResponse {
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
	// 操作人id
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
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作表可见性, 创建之后默认为visible
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
	// 知识库描述。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户id。
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
	// 知识库描述。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库打开url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 知识库id。
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
	// 知识库id。
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

type DeleteColumnsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteColumnsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteColumnsHeaders) SetCommonHeaders(v map[string]*string) *DeleteColumnsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteColumnsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteColumnsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteColumnsRequest struct {
	// 要删除的第一列的位置，从0开始。
	Column *int64 `json:"column,omitempty" xml:"column,omitempty"`
	// 要删除的列的数量。
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteColumnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnsRequest) GoString() string {
	return s.String()
}

func (s *DeleteColumnsRequest) SetColumn(v int64) *DeleteColumnsRequest {
	s.Column = &v
	return s
}

func (s *DeleteColumnsRequest) SetColumnCount(v int64) *DeleteColumnsRequest {
	s.ColumnCount = &v
	return s
}

func (s *DeleteColumnsRequest) SetOperatorId(v string) *DeleteColumnsRequest {
	s.OperatorId = &v
	return s
}

type DeleteColumnsResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteColumnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteColumnsResponseBody) SetId(v string) *DeleteColumnsResponseBody {
	s.Id = &v
	return s
}

type DeleteColumnsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteColumnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnsResponse) GoString() string {
	return s.String()
}

func (s *DeleteColumnsResponse) SetHeaders(v map[string]*string) *DeleteColumnsResponse {
	s.Headers = v
	return s
}

func (s *DeleteColumnsResponse) SetBody(v *DeleteColumnsResponseBody) *DeleteColumnsResponse {
	s.Body = v
	return s
}

type DeleteDropdownListsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDropdownListsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDropdownListsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDropdownListsHeaders) SetCommonHeaders(v map[string]*string) *DeleteDropdownListsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDropdownListsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDropdownListsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDropdownListsRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteDropdownListsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDropdownListsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDropdownListsRequest) SetOperatorId(v string) *DeleteDropdownListsRequest {
	s.OperatorId = &v
	return s
}

type DeleteDropdownListsResponseBody struct {
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s DeleteDropdownListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDropdownListsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDropdownListsResponseBody) SetA1Notation(v string) *DeleteDropdownListsResponseBody {
	s.A1Notation = &v
	return s
}

type DeleteDropdownListsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDropdownListsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDropdownListsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDropdownListsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDropdownListsResponse) SetHeaders(v map[string]*string) *DeleteDropdownListsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDropdownListsResponse) SetBody(v *DeleteDropdownListsResponseBody) *DeleteDropdownListsResponse {
	s.Body = v
	return s
}

type DeleteRangeProtectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRangeProtectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRangeProtectionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRangeProtectionHeaders) SetCommonHeaders(v map[string]*string) *DeleteRangeProtectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRangeProtectionHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRangeProtectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRangeProtectionRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteRangeProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRangeProtectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRangeProtectionRequest) SetOperatorId(v string) *DeleteRangeProtectionRequest {
	s.OperatorId = &v
	return s
}

type DeleteRangeProtectionResponseBody struct {
	// Range地址
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s DeleteRangeProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRangeProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRangeProtectionResponseBody) SetA1Notation(v string) *DeleteRangeProtectionResponseBody {
	s.A1Notation = &v
	return s
}

type DeleteRangeProtectionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRangeProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRangeProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRangeProtectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRangeProtectionResponse) SetHeaders(v map[string]*string) *DeleteRangeProtectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRangeProtectionResponse) SetBody(v *DeleteRangeProtectionResponseBody) *DeleteRangeProtectionResponse {
	s.Body = v
	return s
}

type DeleteRowsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRowsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRowsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRowsHeaders) SetCommonHeaders(v map[string]*string) *DeleteRowsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRowsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRowsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRowsRequest struct {
	// 要删除的第一行的位置，从0开始。
	Row *int64 `json:"row,omitempty" xml:"row,omitempty"`
	// 要删除的行的数量。
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteRowsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRowsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRowsRequest) SetRow(v int64) *DeleteRowsRequest {
	s.Row = &v
	return s
}

func (s *DeleteRowsRequest) SetRowCount(v int64) *DeleteRowsRequest {
	s.RowCount = &v
	return s
}

func (s *DeleteRowsRequest) SetOperatorId(v string) *DeleteRowsRequest {
	s.OperatorId = &v
	return s
}

type DeleteRowsResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteRowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRowsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRowsResponseBody) SetId(v string) *DeleteRowsResponseBody {
	s.Id = &v
	return s
}

type DeleteRowsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRowsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRowsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRowsResponse) SetHeaders(v map[string]*string) *DeleteRowsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRowsResponse) SetBody(v *DeleteRowsResponseBody) *DeleteRowsResponse {
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
	// 操作人id
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

type DeleteSheetResponseBody struct {
	// 本次操作是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponseBody) SetSuccess(v bool) *DeleteSheetResponseBody {
	s.Success = &v
	return s
}

type DeleteSheetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSheetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSheetResponse) SetBody(v *DeleteSheetResponseBody) *DeleteSheetResponse {
	s.Body = v
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

type GetAllSheetsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllSheetsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllSheetsHeaders) SetCommonHeaders(v map[string]*string) *GetAllSheetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllSheetsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllSheetsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllSheetsRequest struct {
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetAllSheetsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsRequest) GoString() string {
	return s.String()
}

func (s *GetAllSheetsRequest) SetOperatorId(v string) *GetAllSheetsRequest {
	s.OperatorId = &v
	return s
}

type GetAllSheetsResponseBody struct {
	// 工作表列表
	Value []*GetAllSheetsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetAllSheetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBody) SetValue(v []*GetAllSheetsResponseBodyValue) *GetAllSheetsResponseBody {
	s.Value = v
	return s
}

type GetAllSheetsResponseBodyValue struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAllSheetsResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBodyValue) SetId(v string) *GetAllSheetsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *GetAllSheetsResponseBodyValue) SetName(v string) *GetAllSheetsResponseBodyValue {
	s.Name = &v
	return s
}

type GetAllSheetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllSheetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllSheetsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponse) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponse) SetHeaders(v map[string]*string) *GetAllSheetsResponse {
	s.Headers = v
	return s
}

func (s *GetAllSheetsResponse) SetBody(v *GetAllSheetsResponseBody) *GetAllSheetsResponse {
	s.Body = v
	return s
}

type GetDeveloperMetadataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeveloperMetadataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataHeaders) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataHeaders) SetCommonHeaders(v map[string]*string) *GetDeveloperMetadataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeveloperMetadataHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeveloperMetadataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeveloperMetadataRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetDeveloperMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataRequest) SetOperatorId(v string) *GetDeveloperMetadataRequest {
	s.OperatorId = &v
	return s
}

type GetDeveloperMetadataResponseBody struct {
	// 元数据所关联到的列
	AssociatedColumn *GetDeveloperMetadataResponseBodyAssociatedColumn `json:"associatedColumn,omitempty" xml:"associatedColumn,omitempty" type:"Struct"`
	// 元数据所关联到的行
	AssociatedRow *GetDeveloperMetadataResponseBodyAssociatedRow `json:"associatedRow,omitempty" xml:"associatedRow,omitempty" type:"Struct"`
	Value         interface{}                                    `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDeveloperMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataResponseBody) SetAssociatedColumn(v *GetDeveloperMetadataResponseBodyAssociatedColumn) *GetDeveloperMetadataResponseBody {
	s.AssociatedColumn = v
	return s
}

func (s *GetDeveloperMetadataResponseBody) SetAssociatedRow(v *GetDeveloperMetadataResponseBodyAssociatedRow) *GetDeveloperMetadataResponseBody {
	s.AssociatedRow = v
	return s
}

func (s *GetDeveloperMetadataResponseBody) SetValue(v interface{}) *GetDeveloperMetadataResponseBody {
	s.Value = v
	return s
}

type GetDeveloperMetadataResponseBodyAssociatedColumn struct {
	// 列号，从0开始
	Column  *int32  `json:"column,omitempty" xml:"column,omitempty"`
	SheetId *string `json:"sheetId,omitempty" xml:"sheetId,omitempty"`
}

func (s GetDeveloperMetadataResponseBodyAssociatedColumn) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataResponseBodyAssociatedColumn) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataResponseBodyAssociatedColumn) SetColumn(v int32) *GetDeveloperMetadataResponseBodyAssociatedColumn {
	s.Column = &v
	return s
}

func (s *GetDeveloperMetadataResponseBodyAssociatedColumn) SetSheetId(v string) *GetDeveloperMetadataResponseBodyAssociatedColumn {
	s.SheetId = &v
	return s
}

type GetDeveloperMetadataResponseBodyAssociatedRow struct {
	// 行号，从0开始
	Row *int32 `json:"row,omitempty" xml:"row,omitempty"`
	// 工作表ID或名称
	SheetId *string `json:"sheetId,omitempty" xml:"sheetId,omitempty"`
}

func (s GetDeveloperMetadataResponseBodyAssociatedRow) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataResponseBodyAssociatedRow) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataResponseBodyAssociatedRow) SetRow(v int32) *GetDeveloperMetadataResponseBodyAssociatedRow {
	s.Row = &v
	return s
}

func (s *GetDeveloperMetadataResponseBodyAssociatedRow) SetSheetId(v string) *GetDeveloperMetadataResponseBodyAssociatedRow {
	s.SheetId = &v
	return s
}

type GetDeveloperMetadataResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeveloperMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeveloperMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeveloperMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetDeveloperMetadataResponse) SetHeaders(v map[string]*string) *GetDeveloperMetadataResponse {
	s.Headers = v
	return s
}

func (s *GetDeveloperMetadataResponse) SetBody(v *GetDeveloperMetadataResponseBody) *GetDeveloperMetadataResponse {
	s.Body = v
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
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 限定要返回的字段
	Select *string `json:"select,omitempty" xml:"select,omitempty"`
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

func (s *GetRangeRequest) SetSelect(v string) *GetRangeRequest {
	s.Select = &v
	return s
}

type GetRangeResponseBody struct {
	// 背景颜色
	BackgroundColors [][]*GetRangeResponseBodyBackgroundColors `json:"backgroundColors,omitempty" xml:"backgroundColors,omitempty" type:"Repeated"`
	// 展示值
	DisplayValues [][]*string `json:"displayValues,omitempty" xml:"displayValues,omitempty" type:"Repeated"`
	// 公式
	Formulas [][]*string `json:"formulas,omitempty" xml:"formulas,omitempty" type:"Repeated"`
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
	Red *int32 `json:"red,omitempty" xml:"red,omitempty"`
	// RGB值中的绿色值
	Green *int32 `json:"green,omitempty" xml:"green,omitempty"`
	// RGB值中的蓝色值
	Blue *int32 `json:"blue,omitempty" xml:"blue,omitempty"`
	// 16进制表示的颜色
	HexString *string `json:"hexString,omitempty" xml:"hexString,omitempty"`
}

func (s GetRangeResponseBodyBackgroundColors) String() string {
	return tea.Prettify(s)
}

func (s GetRangeResponseBodyBackgroundColors) GoString() string {
	return s.String()
}

func (s *GetRangeResponseBodyBackgroundColors) SetRed(v int32) *GetRangeResponseBodyBackgroundColors {
	s.Red = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetGreen(v int32) *GetRangeResponseBodyBackgroundColors {
	s.Green = &v
	return s
}

func (s *GetRangeResponseBodyBackgroundColors) SetBlue(v int32) *GetRangeResponseBodyBackgroundColors {
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
	// 知识库信息。
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
	// 知识库打开url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 知识库id。
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 知识库名称。
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
	// 知识库信息。
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
	// 知识库打开url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 知识库id。
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 知识库名称。
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
	// 知识库结果集。
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
	// 知识库创建时间。
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 知识库是否被删除。
	Deleted *bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 知识库名称。
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// 知识库最近访问文档列表。
	RecentList []*GetRelatedWorkspacesResponseBodyWorkspacesRecentList `json:"recentList,omitempty" xml:"recentList,omitempty" type:"Repeated"`
	// 用户的角色
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// 知识库打开url。
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 知识库id。
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
	// 操作人id
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
	// 工作表列数。
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 最后一列非空列的位置，从0开始，表为空时返回-1。
	LastNonEmptyColumn *int64 `json:"lastNonEmptyColumn,omitempty" xml:"lastNonEmptyColumn,omitempty"`
	// 最后一行非空行的位置，从0开始，表为空时返回-1。
	LastNonEmptyRow *int64 `json:"lastNonEmptyRow,omitempty" xml:"lastNonEmptyRow,omitempty"`
	// 工作表名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作表行数。
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 工作表可见性
	// 枚举值:
	//    visible: 可见
	//    hidden: 隐藏
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

func (s *GetSheetResponseBody) SetId(v string) *GetSheetResponseBody {
	s.Id = &v
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
	// 模版归属知识库id。
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
	// 节点所属知识库信息。
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
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库id。
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

type InsertColumnsBeforeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertColumnsBeforeHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertColumnsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertColumnsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertColumnsBeforeHeaders) SetXAcsDingtalkAccessToken(v string) *InsertColumnsBeforeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertColumnsBeforeRequest struct {
	// 要插入列的位置，从0开始。
	Column *int64 `json:"column,omitempty" xml:"column,omitempty"`
	// 要插入列的数量。
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s InsertColumnsBeforeRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertColumnsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeRequest) SetColumn(v int64) *InsertColumnsBeforeRequest {
	s.Column = &v
	return s
}

func (s *InsertColumnsBeforeRequest) SetColumnCount(v int64) *InsertColumnsBeforeRequest {
	s.ColumnCount = &v
	return s
}

func (s *InsertColumnsBeforeRequest) SetOperatorId(v string) *InsertColumnsBeforeRequest {
	s.OperatorId = &v
	return s
}

type InsertColumnsBeforeResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s InsertColumnsBeforeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertColumnsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeResponseBody) SetId(v string) *InsertColumnsBeforeResponseBody {
	s.Id = &v
	return s
}

type InsertColumnsBeforeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertColumnsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertColumnsBeforeResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertColumnsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeResponse) SetHeaders(v map[string]*string) *InsertColumnsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertColumnsBeforeResponse) SetBody(v *InsertColumnsBeforeResponseBody) *InsertColumnsBeforeResponse {
	s.Body = v
	return s
}

type InsertDropdownListsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertDropdownListsHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertDropdownListsHeaders) GoString() string {
	return s.String()
}

func (s *InsertDropdownListsHeaders) SetCommonHeaders(v map[string]*string) *InsertDropdownListsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertDropdownListsHeaders) SetXAcsDingtalkAccessToken(v string) *InsertDropdownListsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertDropdownListsRequest struct {
	Options []*InsertDropdownListsRequestOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s InsertDropdownListsRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertDropdownListsRequest) GoString() string {
	return s.String()
}

func (s *InsertDropdownListsRequest) SetOptions(v []*InsertDropdownListsRequestOptions) *InsertDropdownListsRequest {
	s.Options = v
	return s
}

func (s *InsertDropdownListsRequest) SetOperatorId(v string) *InsertDropdownListsRequest {
	s.OperatorId = &v
	return s
}

type InsertDropdownListsRequestOptions struct {
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InsertDropdownListsRequestOptions) String() string {
	return tea.Prettify(s)
}

func (s InsertDropdownListsRequestOptions) GoString() string {
	return s.String()
}

func (s *InsertDropdownListsRequestOptions) SetColor(v string) *InsertDropdownListsRequestOptions {
	s.Color = &v
	return s
}

func (s *InsertDropdownListsRequestOptions) SetValue(v string) *InsertDropdownListsRequestOptions {
	s.Value = &v
	return s
}

type InsertDropdownListsResponseBody struct {
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s InsertDropdownListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertDropdownListsResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDropdownListsResponseBody) SetA1Notation(v string) *InsertDropdownListsResponseBody {
	s.A1Notation = &v
	return s
}

type InsertDropdownListsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertDropdownListsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertDropdownListsResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertDropdownListsResponse) GoString() string {
	return s.String()
}

func (s *InsertDropdownListsResponse) SetHeaders(v map[string]*string) *InsertDropdownListsResponse {
	s.Headers = v
	return s
}

func (s *InsertDropdownListsResponse) SetBody(v *InsertDropdownListsResponseBody) *InsertDropdownListsResponse {
	s.Body = v
	return s
}

type InsertRowsBeforeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertRowsBeforeHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeHeaders) SetXAcsDingtalkAccessToken(v string) *InsertRowsBeforeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertRowsBeforeRequest struct {
	// 要插入行的位置，从0开始。
	Row *int64 `json:"row,omitempty" xml:"row,omitempty"`
	// 要插入行的数量。
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s InsertRowsBeforeRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequest) SetRow(v int64) *InsertRowsBeforeRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetRowCount(v int64) *InsertRowsBeforeRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetOperatorId(v string) *InsertRowsBeforeRequest {
	s.OperatorId = &v
	return s
}

type InsertRowsBeforeResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s InsertRowsBeforeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponseBody) SetId(v string) *InsertRowsBeforeResponseBody {
	s.Id = &v
	return s
}

type InsertRowsBeforeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertRowsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertRowsBeforeResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponse) SetHeaders(v map[string]*string) *InsertRowsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertRowsBeforeResponse) SetBody(v *InsertRowsBeforeResponseBody) *InsertRowsBeforeResponse {
	s.Body = v
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
	// 知识库id。
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
	// 模版归属知识库id。
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

type MergeRangeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MergeRangeHeaders) String() string {
	return tea.Prettify(s)
}

func (s MergeRangeHeaders) GoString() string {
	return s.String()
}

func (s *MergeRangeHeaders) SetCommonHeaders(v map[string]*string) *MergeRangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MergeRangeHeaders) SetXAcsDingtalkAccessToken(v string) *MergeRangeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MergeRangeRequest struct {
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s MergeRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeRangeRequest) GoString() string {
	return s.String()
}

func (s *MergeRangeRequest) SetOperatorId(v string) *MergeRangeRequest {
	s.OperatorId = &v
	return s
}

type MergeRangeResponseBody struct {
	// 合并的单元格地址
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s MergeRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *MergeRangeResponseBody) SetA1Notation(v string) *MergeRangeResponseBody {
	s.A1Notation = &v
	return s
}

type MergeRangeResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeRangeResponse) GoString() string {
	return s.String()
}

func (s *MergeRangeResponse) SetHeaders(v map[string]*string) *MergeRangeResponse {
	s.Headers = v
	return s
}

func (s *MergeRangeResponse) SetBody(v *MergeRangeResponseBody) *MergeRangeResponse {
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
	IncludeHidden *bool `json:"includeHidden,omitempty" xml:"includeHidden,omitempty"`
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

func (s *RangeFindNextRequestFindOptions) SetIncludeHidden(v bool) *RangeFindNextRequestFindOptions {
	s.IncludeHidden = &v
	return s
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
	// 知识库id。
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
	// 知识库名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 知识库id。
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

type SetColumnsVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetColumnsVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetColumnsVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityHeaders) SetCommonHeaders(v map[string]*string) *SetColumnsVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetColumnsVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *SetColumnsVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetColumnsVisibilityRequest struct {
	// 要显示、隐藏的第一列的位置，从0开始。
	Column *int64 `json:"column,omitempty" xml:"column,omitempty"`
	// 要显示、隐藏的列的数量。
	ColumnCount *int64 `json:"columnCount,omitempty" xml:"columnCount,omitempty"`
	// 列可见性
	// 枚举值:
	// 	visible: 可见
	// 	hidden: 隐藏
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SetColumnsVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s SetColumnsVisibilityRequest) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityRequest) SetColumn(v int64) *SetColumnsVisibilityRequest {
	s.Column = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetColumnCount(v int64) *SetColumnsVisibilityRequest {
	s.ColumnCount = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetVisibility(v string) *SetColumnsVisibilityRequest {
	s.Visibility = &v
	return s
}

func (s *SetColumnsVisibilityRequest) SetOperatorId(v string) *SetColumnsVisibilityRequest {
	s.OperatorId = &v
	return s
}

type SetColumnsVisibilityResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SetColumnsVisibilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetColumnsVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityResponseBody) SetId(v string) *SetColumnsVisibilityResponseBody {
	s.Id = &v
	return s
}

type SetColumnsVisibilityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetColumnsVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetColumnsVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s SetColumnsVisibilityResponse) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityResponse) SetHeaders(v map[string]*string) *SetColumnsVisibilityResponse {
	s.Headers = v
	return s
}

func (s *SetColumnsVisibilityResponse) SetBody(v *SetColumnsVisibilityResponseBody) *SetColumnsVisibilityResponse {
	s.Body = v
	return s
}

type SetRowsVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetRowsVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetRowsVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityHeaders) SetCommonHeaders(v map[string]*string) *SetRowsVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRowsVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *SetRowsVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetRowsVisibilityRequest struct {
	// 要显示、隐藏的第一行的位置，从0开始。
	Row *int64 `json:"row,omitempty" xml:"row,omitempty"`
	// 要显示、隐藏的行的数量。
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 行可见性
	// 枚举值:
	// 	visible: 可见
	// 	hidden: 隐藏
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 操作人id
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SetRowsVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRowsVisibilityRequest) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityRequest) SetRow(v int64) *SetRowsVisibilityRequest {
	s.Row = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetRowCount(v int64) *SetRowsVisibilityRequest {
	s.RowCount = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetVisibility(v string) *SetRowsVisibilityRequest {
	s.Visibility = &v
	return s
}

func (s *SetRowsVisibilityRequest) SetOperatorId(v string) *SetRowsVisibilityRequest {
	s.OperatorId = &v
	return s
}

type SetRowsVisibilityResponseBody struct {
	// 工作表id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SetRowsVisibilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRowsVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityResponseBody) SetId(v string) *SetRowsVisibilityResponseBody {
	s.Id = &v
	return s
}

type SetRowsVisibilityResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRowsVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRowsVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRowsVisibilityResponse) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityResponse) SetHeaders(v map[string]*string) *SetRowsVisibilityResponse {
	s.Headers = v
	return s
}

func (s *SetRowsVisibilityResponse) SetBody(v *SetRowsVisibilityResponseBody) *SetRowsVisibilityResponse {
	s.Body = v
	return s
}

type SheetAutofitRowsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SheetAutofitRowsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SheetAutofitRowsHeaders) GoString() string {
	return s.String()
}

func (s *SheetAutofitRowsHeaders) SetCommonHeaders(v map[string]*string) *SheetAutofitRowsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SheetAutofitRowsHeaders) SetXAcsDingtalkAccessToken(v string) *SheetAutofitRowsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SheetAutofitRowsRequest struct {
	FontWidth *int64 `json:"fontWidth,omitempty" xml:"fontWidth,omitempty"`
	// 行号，从0开始
	Row *int64 `json:"row,omitempty" xml:"row,omitempty"`
	// 行数
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SheetAutofitRowsRequest) String() string {
	return tea.Prettify(s)
}

func (s SheetAutofitRowsRequest) GoString() string {
	return s.String()
}

func (s *SheetAutofitRowsRequest) SetFontWidth(v int64) *SheetAutofitRowsRequest {
	s.FontWidth = &v
	return s
}

func (s *SheetAutofitRowsRequest) SetRow(v int64) *SheetAutofitRowsRequest {
	s.Row = &v
	return s
}

func (s *SheetAutofitRowsRequest) SetRowCount(v int64) *SheetAutofitRowsRequest {
	s.RowCount = &v
	return s
}

func (s *SheetAutofitRowsRequest) SetOperatorId(v string) *SheetAutofitRowsRequest {
	s.OperatorId = &v
	return s
}

type SheetAutofitRowsResponseBody struct {
	// 当前工作表ID
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SheetAutofitRowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SheetAutofitRowsResponseBody) GoString() string {
	return s.String()
}

func (s *SheetAutofitRowsResponseBody) SetId(v string) *SheetAutofitRowsResponseBody {
	s.Id = &v
	return s
}

type SheetAutofitRowsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SheetAutofitRowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SheetAutofitRowsResponse) String() string {
	return tea.Prettify(s)
}

func (s SheetAutofitRowsResponse) GoString() string {
	return s.String()
}

func (s *SheetAutofitRowsResponse) SetHeaders(v map[string]*string) *SheetAutofitRowsResponse {
	s.Headers = v
	return s
}

func (s *SheetAutofitRowsResponse) SetBody(v *SheetAutofitRowsResponseBody) *SheetAutofitRowsResponse {
	s.Body = v
	return s
}

type SheetFindAllHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SheetFindAllHeaders) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllHeaders) GoString() string {
	return s.String()
}

func (s *SheetFindAllHeaders) SetCommonHeaders(v map[string]*string) *SheetFindAllHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SheetFindAllHeaders) SetXAcsDingtalkAccessToken(v string) *SheetFindAllHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SheetFindAllRequest struct {
	// 查找选项
	FindOptions *SheetFindAllRequestFindOptions `json:"findOptions,omitempty" xml:"findOptions,omitempty" type:"Struct"`
	// 要查找的文本
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	Select     *string `json:"select,omitempty" xml:"select,omitempty"`
}

func (s SheetFindAllRequest) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllRequest) GoString() string {
	return s.String()
}

func (s *SheetFindAllRequest) SetFindOptions(v *SheetFindAllRequestFindOptions) *SheetFindAllRequest {
	s.FindOptions = v
	return s
}

func (s *SheetFindAllRequest) SetText(v string) *SheetFindAllRequest {
	s.Text = &v
	return s
}

func (s *SheetFindAllRequest) SetOperatorId(v string) *SheetFindAllRequest {
	s.OperatorId = &v
	return s
}

func (s *SheetFindAllRequest) SetSelect(v string) *SheetFindAllRequest {
	s.Select = &v
	return s
}

type SheetFindAllRequestFindOptions struct {
	IncludeHidden *bool `json:"includeHidden,omitempty" xml:"includeHidden,omitempty"`
	// 匹配大小写
	MatchCase *bool `json:"matchCase,omitempty" xml:"matchCase,omitempty"`
	// 匹配整个单元格
	MatchEntireCell *bool `json:"matchEntireCell,omitempty" xml:"matchEntireCell,omitempty"`
	// 在公式内搜索
	MatchFormulaText *bool   `json:"matchFormulaText,omitempty" xml:"matchFormulaText,omitempty"`
	Scope            *string `json:"scope,omitempty" xml:"scope,omitempty"`
	UnionCells       *bool   `json:"unionCells,omitempty" xml:"unionCells,omitempty"`
	// text是正则表达式
	UseRegExp *bool `json:"useRegExp,omitempty" xml:"useRegExp,omitempty"`
}

func (s SheetFindAllRequestFindOptions) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllRequestFindOptions) GoString() string {
	return s.String()
}

func (s *SheetFindAllRequestFindOptions) SetIncludeHidden(v bool) *SheetFindAllRequestFindOptions {
	s.IncludeHidden = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetMatchCase(v bool) *SheetFindAllRequestFindOptions {
	s.MatchCase = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetMatchEntireCell(v bool) *SheetFindAllRequestFindOptions {
	s.MatchEntireCell = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetMatchFormulaText(v bool) *SheetFindAllRequestFindOptions {
	s.MatchFormulaText = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetScope(v string) *SheetFindAllRequestFindOptions {
	s.Scope = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetUnionCells(v bool) *SheetFindAllRequestFindOptions {
	s.UnionCells = &v
	return s
}

func (s *SheetFindAllRequestFindOptions) SetUseRegExp(v bool) *SheetFindAllRequestFindOptions {
	s.UseRegExp = &v
	return s
}

type SheetFindAllResponseBody struct {
	Value []*SheetFindAllResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s SheetFindAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllResponseBody) GoString() string {
	return s.String()
}

func (s *SheetFindAllResponseBody) SetValue(v []*SheetFindAllResponseBodyValue) *SheetFindAllResponseBody {
	s.Value = v
	return s
}

type SheetFindAllResponseBodyValue struct {
	A1Notation *string         `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
	Values     [][]interface{} `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s SheetFindAllResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllResponseBodyValue) GoString() string {
	return s.String()
}

func (s *SheetFindAllResponseBodyValue) SetA1Notation(v string) *SheetFindAllResponseBodyValue {
	s.A1Notation = &v
	return s
}

func (s *SheetFindAllResponseBodyValue) SetValues(v [][]interface{}) *SheetFindAllResponseBodyValue {
	s.Values = v
	return s
}

type SheetFindAllResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SheetFindAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SheetFindAllResponse) String() string {
	return tea.Prettify(s)
}

func (s SheetFindAllResponse) GoString() string {
	return s.String()
}

func (s *SheetFindAllResponse) SetHeaders(v map[string]*string) *SheetFindAllResponse {
	s.Headers = v
	return s
}

func (s *SheetFindAllResponse) SetBody(v *SheetFindAllResponseBody) *SheetFindAllResponse {
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
	// 最大size:
	// 	1000
	BackgroundColors [][]*string `json:"backgroundColors,omitempty" xml:"backgroundColors,omitempty" type:"Repeated"`
	// 超链接
	// 最大size:
	// 	1000
	Hyperlinks [][]*UpdateRangeRequestHyperlinks `json:"hyperlinks,omitempty" xml:"hyperlinks,omitempty" type:"Repeated"`
	// 数字格式
	NumberFormat *string `json:"numberFormat,omitempty" xml:"numberFormat,omitempty"`
	// 值
	// 最大size:
	// 	1000
	Values [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 操作人id
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

func (s *UpdateRangeRequest) SetNumberFormat(v string) *UpdateRangeRequest {
	s.NumberFormat = &v
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
	// 超链接类型，可选path、sheet、range
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 超链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 超链接文本
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

type UpdateRangeResponseBody struct {
	// 使用A1表示法的Range地址
	A1Notation *string `json:"a1Notation,omitempty" xml:"a1Notation,omitempty"`
}

func (s UpdateRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRangeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRangeResponseBody) SetA1Notation(v string) *UpdateRangeResponseBody {
	s.A1Notation = &v
	return s
}

type UpdateRangeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateRangeResponse) SetBody(v *UpdateRangeResponseBody) *UpdateRangeResponse {
	s.Body = v
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
	// 枚举值:
	// 	visible: 可见
	// 	hidden: 隐藏
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 操作人id
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

func (client *Client) Clear(workbookId *string, sheetId *string, rangeAddress *string, request *ClearRequest) (_result *ClearResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearHeaders{}
	_result = &ClearResponse{}
	_body, _err := client.ClearWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *ClearRequest, headers *ClearHeaders, runtime *util.RuntimeOptions) (_result *ClearResponse, _err error) {
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
	_result = &ClearResponse{}
	_body, _err := client.DoROARequest(tea.String("Clear"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/clear"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearData(workbookId *string, sheetId *string, rangeAddress *string, request *ClearDataRequest) (_result *ClearDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ClearDataHeaders{}
	_result = &ClearDataResponse{}
	_body, _err := client.ClearDataWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearDataWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *ClearDataRequest, headers *ClearDataHeaders, runtime *util.RuntimeOptions) (_result *ClearDataResponse, _err error) {
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
	_result = &ClearDataResponse{}
	_body, _err := client.DoROARequest(tea.String("ClearData"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/clearData"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConditionalFormattingRule(workbookId *string, sheetId *string, request *CreateConditionalFormattingRuleRequest) (_result *CreateConditionalFormattingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateConditionalFormattingRuleHeaders{}
	_result = &CreateConditionalFormattingRuleResponse{}
	_body, _err := client.CreateConditionalFormattingRuleWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConditionalFormattingRuleWithOptions(workbookId *string, sheetId *string, request *CreateConditionalFormattingRuleRequest, headers *CreateConditionalFormattingRuleHeaders, runtime *util.RuntimeOptions) (_result *CreateConditionalFormattingRuleResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.CellStyle)) {
		body["cellStyle"] = request.CellStyle
	}

	if !tea.BoolValue(util.IsUnset(request.DuplicateCondition)) {
		body["duplicateCondition"] = request.DuplicateCondition
	}

	if !tea.BoolValue(util.IsUnset(request.Ranges)) {
		body["ranges"] = request.Ranges
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
	_result = &CreateConditionalFormattingRuleResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateConditionalFormattingRule"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/conditionalFormattingRules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeveloperMetadata(workbookId *string, request *CreateDeveloperMetadataRequest) (_result *CreateDeveloperMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeveloperMetadataHeaders{}
	_result = &CreateDeveloperMetadataResponse{}
	_body, _err := client.CreateDeveloperMetadataWithOptions(workbookId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeveloperMetadataWithOptions(workbookId *string, request *CreateDeveloperMetadataRequest, headers *CreateDeveloperMetadataHeaders, runtime *util.RuntimeOptions) (_result *CreateDeveloperMetadataResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.AssociatedColumn)) {
		body["associatedColumn"] = request.AssociatedColumn
	}

	if !tea.BoolValue(util.IsUnset(request.AssociatedRow)) {
		body["associatedRow"] = request.AssociatedRow
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
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
	_result = &CreateDeveloperMetadataResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDeveloperMetadata"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/developerMetadatas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRangeProtection(workbookId *string, sheetId *string, rangeAddress *string, request *CreateRangeProtectionRequest) (_result *CreateRangeProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRangeProtectionHeaders{}
	_result = &CreateRangeProtectionResponse{}
	_body, _err := client.CreateRangeProtectionWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRangeProtectionWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *CreateRangeProtectionRequest, headers *CreateRangeProtectionHeaders, runtime *util.RuntimeOptions) (_result *CreateRangeProtectionResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.EditableSetting)) {
		body["editableSetting"] = request.EditableSetting
	}

	if !tea.BoolValue(util.IsUnset(request.OtherUserPermission)) {
		body["otherUserPermission"] = request.OtherUserPermission
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
	_result = &CreateRangeProtectionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRangeProtection"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/protections"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteColumns(workbookId *string, sheetId *string, request *DeleteColumnsRequest) (_result *DeleteColumnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteColumnsHeaders{}
	_result = &DeleteColumnsResponse{}
	_body, _err := client.DeleteColumnsWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteColumnsWithOptions(workbookId *string, sheetId *string, request *DeleteColumnsRequest, headers *DeleteColumnsHeaders, runtime *util.RuntimeOptions) (_result *DeleteColumnsResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnCount)) {
		body["columnCount"] = request.ColumnCount
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
	_result = &DeleteColumnsResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteColumns"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/deleteColumns"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDropdownLists(workbookId *string, sheetId *string, rangeAddress *string, request *DeleteDropdownListsRequest) (_result *DeleteDropdownListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDropdownListsHeaders{}
	_result = &DeleteDropdownListsResponse{}
	_body, _err := client.DeleteDropdownListsWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDropdownListsWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *DeleteDropdownListsRequest, headers *DeleteDropdownListsHeaders, runtime *util.RuntimeOptions) (_result *DeleteDropdownListsResponse, _err error) {
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
	_result = &DeleteDropdownListsResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDropdownLists"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/deleteDropdownLists"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRangeProtection(workbookId *string, sheetId *string, rangeAddress *string, protectionId *string, request *DeleteRangeProtectionRequest) (_result *DeleteRangeProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRangeProtectionHeaders{}
	_result = &DeleteRangeProtectionResponse{}
	_body, _err := client.DeleteRangeProtectionWithOptions(workbookId, sheetId, rangeAddress, protectionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRangeProtectionWithOptions(workbookId *string, sheetId *string, rangeAddress *string, protectionId *string, request *DeleteRangeProtectionRequest, headers *DeleteRangeProtectionHeaders, runtime *util.RuntimeOptions) (_result *DeleteRangeProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	sheetId = openapiutil.GetEncodeParam(sheetId)
	rangeAddress = openapiutil.GetEncodeParam(rangeAddress)
	protectionId = openapiutil.GetEncodeParam(protectionId)
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
	_result = &DeleteRangeProtectionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRangeProtection"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/protections/"+tea.StringValue(protectionId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRows(workbookId *string, sheetId *string, request *DeleteRowsRequest) (_result *DeleteRowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRowsHeaders{}
	_result = &DeleteRowsResponse{}
	_body, _err := client.DeleteRowsWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRowsWithOptions(workbookId *string, sheetId *string, request *DeleteRowsRequest, headers *DeleteRowsHeaders, runtime *util.RuntimeOptions) (_result *DeleteRowsResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["rowCount"] = request.RowCount
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
	_result = &DeleteRowsResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRows"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/deleteRows"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoROARequest(tea.String("DeleteSheet"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)), tea.String("json"), req, runtime)
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

func (client *Client) GetAllSheets(workbookId *string, request *GetAllSheetsRequest) (_result *GetAllSheetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllSheetsHeaders{}
	_result = &GetAllSheetsResponse{}
	_body, _err := client.GetAllSheetsWithOptions(workbookId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllSheetsWithOptions(workbookId *string, request *GetAllSheetsRequest, headers *GetAllSheetsHeaders, runtime *util.RuntimeOptions) (_result *GetAllSheetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
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
	_result = &GetAllSheetsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAllSheets"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeveloperMetadata(workbookId *string, developerMetadataId *string, request *GetDeveloperMetadataRequest) (_result *GetDeveloperMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeveloperMetadataHeaders{}
	_result = &GetDeveloperMetadataResponse{}
	_body, _err := client.GetDeveloperMetadataWithOptions(workbookId, developerMetadataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeveloperMetadataWithOptions(workbookId *string, developerMetadataId *string, request *GetDeveloperMetadataRequest, headers *GetDeveloperMetadataHeaders, runtime *util.RuntimeOptions) (_result *GetDeveloperMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	workbookId = openapiutil.GetEncodeParam(workbookId)
	developerMetadataId = openapiutil.GetEncodeParam(developerMetadataId)
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
	_result = &GetDeveloperMetadataResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDeveloperMetadata"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/developerMetadatas/"+tea.StringValue(developerMetadataId)), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Select)) {
		query["select"] = request.Select
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

	if !tea.BoolValue(util.IsUnset(request.Location)) {
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

func (client *Client) InsertColumnsBefore(workbookId *string, sheetId *string, request *InsertColumnsBeforeRequest) (_result *InsertColumnsBeforeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertColumnsBeforeHeaders{}
	_result = &InsertColumnsBeforeResponse{}
	_body, _err := client.InsertColumnsBeforeWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertColumnsBeforeWithOptions(workbookId *string, sheetId *string, request *InsertColumnsBeforeRequest, headers *InsertColumnsBeforeHeaders, runtime *util.RuntimeOptions) (_result *InsertColumnsBeforeResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnCount)) {
		body["columnCount"] = request.ColumnCount
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
	_result = &InsertColumnsBeforeResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertColumnsBefore"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/insertColumnsBefore"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertDropdownLists(workbookId *string, sheetId *string, rangeAddress *string, request *InsertDropdownListsRequest) (_result *InsertDropdownListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertDropdownListsHeaders{}
	_result = &InsertDropdownListsResponse{}
	_body, _err := client.InsertDropdownListsWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertDropdownListsWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *InsertDropdownListsRequest, headers *InsertDropdownListsHeaders, runtime *util.RuntimeOptions) (_result *InsertDropdownListsResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
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
	_result = &InsertDropdownListsResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertDropdownLists"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/insertDropdownLists"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRowsBefore(workbookId *string, sheetId *string, request *InsertRowsBeforeRequest) (_result *InsertRowsBeforeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertRowsBeforeHeaders{}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.InsertRowsBeforeWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRowsBeforeWithOptions(workbookId *string, sheetId *string, request *InsertRowsBeforeRequest, headers *InsertRowsBeforeHeaders, runtime *util.RuntimeOptions) (_result *InsertRowsBeforeResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["rowCount"] = request.RowCount
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
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertRowsBefore"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/insertRowsBefore"), tea.String("json"), req, runtime)
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

func (client *Client) MergeRange(workbookId *string, sheetId *string, rangeAddress *string, request *MergeRangeRequest) (_result *MergeRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MergeRangeHeaders{}
	_result = &MergeRangeResponse{}
	_body, _err := client.MergeRangeWithOptions(workbookId, sheetId, rangeAddress, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeRangeWithOptions(workbookId *string, sheetId *string, rangeAddress *string, request *MergeRangeRequest, headers *MergeRangeHeaders, runtime *util.RuntimeOptions) (_result *MergeRangeResponse, _err error) {
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
	_result = &MergeRangeResponse{}
	_body, _err := client.DoROARequest(tea.String("MergeRange"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)+"/merge"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.FindOptions)) {
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

func (client *Client) SetColumnsVisibility(workbookId *string, sheetId *string, request *SetColumnsVisibilityRequest) (_result *SetColumnsVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetColumnsVisibilityHeaders{}
	_result = &SetColumnsVisibilityResponse{}
	_body, _err := client.SetColumnsVisibilityWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetColumnsVisibilityWithOptions(workbookId *string, sheetId *string, request *SetColumnsVisibilityRequest, headers *SetColumnsVisibilityHeaders, runtime *util.RuntimeOptions) (_result *SetColumnsVisibilityResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Column)) {
		body["column"] = request.Column
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnCount)) {
		body["columnCount"] = request.ColumnCount
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
	_result = &SetColumnsVisibilityResponse{}
	_body, _err := client.DoROARequest(tea.String("SetColumnsVisibility"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/setColumnsVisibility"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRowsVisibility(workbookId *string, sheetId *string, request *SetRowsVisibilityRequest) (_result *SetRowsVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetRowsVisibilityHeaders{}
	_result = &SetRowsVisibilityResponse{}
	_body, _err := client.SetRowsVisibilityWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRowsVisibilityWithOptions(workbookId *string, sheetId *string, request *SetRowsVisibilityRequest, headers *SetRowsVisibilityHeaders, runtime *util.RuntimeOptions) (_result *SetRowsVisibilityResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["rowCount"] = request.RowCount
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
	_result = &SetRowsVisibilityResponse{}
	_body, _err := client.DoROARequest(tea.String("SetRowsVisibility"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/setRowsVisibility"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SheetAutofitRows(workbookId *string, sheetId *string, request *SheetAutofitRowsRequest) (_result *SheetAutofitRowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SheetAutofitRowsHeaders{}
	_result = &SheetAutofitRowsResponse{}
	_body, _err := client.SheetAutofitRowsWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SheetAutofitRowsWithOptions(workbookId *string, sheetId *string, request *SheetAutofitRowsRequest, headers *SheetAutofitRowsHeaders, runtime *util.RuntimeOptions) (_result *SheetAutofitRowsResponse, _err error) {
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
	if !tea.BoolValue(util.IsUnset(request.FontWidth)) {
		body["fontWidth"] = request.FontWidth
	}

	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["rowCount"] = request.RowCount
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
	_result = &SheetAutofitRowsResponse{}
	_body, _err := client.DoROARequest(tea.String("SheetAutofitRows"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/autofitRows"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SheetFindAll(workbookId *string, sheetId *string, request *SheetFindAllRequest) (_result *SheetFindAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SheetFindAllHeaders{}
	_result = &SheetFindAllResponse{}
	_body, _err := client.SheetFindAllWithOptions(workbookId, sheetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SheetFindAllWithOptions(workbookId *string, sheetId *string, request *SheetFindAllRequest, headers *SheetFindAllHeaders, runtime *util.RuntimeOptions) (_result *SheetFindAllResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Select)) {
		query["select"] = request.Select
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FindOptions)) {
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
	_result = &SheetFindAllResponse{}
	_body, _err := client.DoROARequest(tea.String("SheetFindAll"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/findAll"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.NumberFormat)) {
		body["numberFormat"] = request.NumberFormat
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
	_body, _err := client.DoROARequest(tea.String("UpdateRange"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workbooks/"+tea.StringValue(workbookId)+"/sheets/"+tea.StringValue(sheetId)+"/ranges/"+tea.StringValue(rangeAddress)), tea.String("json"), req, runtime)
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
