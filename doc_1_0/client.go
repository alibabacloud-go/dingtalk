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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*DeleteWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s DeleteWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequest) SetOperatorId(v string) *DeleteWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetMembers(v []*DeleteWorkspaceMembersRequestMembers) *DeleteWorkspaceMembersRequest {
	s.Members = v
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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*AddWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s AddWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequest) SetOperatorId(v string) *AddWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetMembers(v []*AddWorkspaceDocMembersRequestMembers) *AddWorkspaceDocMembersRequest {
	s.Members = v
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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*UpdateWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequest) SetOperatorId(v string) *UpdateWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetMembers(v []*UpdateWorkspaceMembersRequestMembers) *UpdateWorkspaceMembersRequest {
	s.Members = v
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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*UpdateWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequest) SetOperatorId(v string) *UpdateWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetMembers(v []*UpdateWorkspaceDocMembersRequestMembers) *UpdateWorkspaceDocMembersRequest {
	s.Members = v
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
	// 文档名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文档类型
	DocType *string `json:"docType,omitempty" xml:"docType,omitempty"`
	// 操作人unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateWorkspaceDocRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequest) SetName(v string) *CreateWorkspaceDocRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetDocType(v string) *CreateWorkspaceDocRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetOperatorId(v string) *CreateWorkspaceDocRequest {
	s.OperatorId = &v
	return s
}

type CreateWorkspaceDocResponseBody struct {
	// 团队空间Id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 文档Id
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 文档docKey
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	// 文档打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateWorkspaceDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponseBody) SetWorkspaceId(v string) *CreateWorkspaceDocResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetNodeId(v string) *CreateWorkspaceDocResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetDocKey(v string) *CreateWorkspaceDocResponseBody {
	s.DocKey = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetUrl(v string) *CreateWorkspaceDocResponseBody {
	s.Url = &v
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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*AddWorkspaceMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s AddWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequest) SetOperatorId(v string) *AddWorkspaceMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *AddWorkspaceMembersRequest) SetMembers(v []*AddWorkspaceMembersRequestMembers) *AddWorkspaceMembersRequest {
	s.Members = v
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
	// 团队空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 团队空间描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 用户id
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	DingOrgId           *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingUid             *int64  `json:"dingUid,omitempty" xml:"dingUid,omitempty"`
	DingAccessTokenType *string `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingIsvOrgId        *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOperatorId(v string) *CreateWorkspaceRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDingOrgId(v int64) *CreateWorkspaceRequest {
	s.DingOrgId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDingUid(v int64) *CreateWorkspaceRequest {
	s.DingUid = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDingAccessTokenType(v string) *CreateWorkspaceRequest {
	s.DingAccessTokenType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDingIsvOrgId(v int64) *CreateWorkspaceRequest {
	s.DingIsvOrgId = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	// 工作空间id
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// 工作空间名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 工作空间描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作空间打开url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetName(v string) *CreateWorkspaceResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetDescription(v string) *CreateWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetUrl(v string) *CreateWorkspaceResponseBody {
	s.Url = &v
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
	// 发起操作者unionId
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// 被操作用户组
	Members []*DeleteWorkspaceDocMembersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
}

func (s DeleteWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequest) SetOperatorId(v string) *DeleteWorkspaceDocMembersRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetMembers(v []*DeleteWorkspaceDocMembersRequestMembers) *DeleteWorkspaceDocMembersRequest {
	s.Members = v
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
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
	IsDeleted *bool   `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetUrl(v string) *GetWorkspaceResponseBody {
	s.Url = &v
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members/remove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["docType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateWorkspaceDoc"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddWorkspaceMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/members"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingUid)) {
		body["dingUid"] = request.DingUid
	}

	if !tea.BoolValue(util.IsUnset(request.DingAccessTokenType)) {
		body["dingAccessTokenType"] = request.DingAccessTokenType
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
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
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateWorkspace"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
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
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteWorkspaceDocMembers"), tea.String("doc_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/doc/workspaces/"+tea.StringValue(workspaceId)+"/docs/"+tea.StringValue(nodeId)+"/members/remove"), tea.String("none"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
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
