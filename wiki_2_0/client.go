// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package wiki_2_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddTeamHeaders) GoString() string {
	return s.String()
}

func (s *AddTeamHeaders) SetCommonHeaders(v map[string]*string) *AddTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddTeamHeaders) SetXAcsDingtalkAccessToken(v string) *AddTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddTeamRequest struct {
	Name       *string               `json:"name,omitempty" xml:"name,omitempty"`
	Option     *AddTeamRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	OperatorId *string               `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s AddTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTeamRequest) GoString() string {
	return s.String()
}

func (s *AddTeamRequest) SetName(v string) *AddTeamRequest {
	s.Name = &v
	return s
}

func (s *AddTeamRequest) SetOption(v *AddTeamRequestOption) *AddTeamRequest {
	s.Option = v
	return s
}

func (s *AddTeamRequest) SetOperatorId(v string) *AddTeamRequest {
	s.OperatorId = &v
	return s
}

type AddTeamRequestOption struct {
	Cover       *string                   `json:"cover,omitempty" xml:"cover,omitempty"`
	Description *string                   `json:"description,omitempty" xml:"description,omitempty"`
	Icon        *AddTeamRequestOptionIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
}

func (s AddTeamRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddTeamRequestOption) GoString() string {
	return s.String()
}

func (s *AddTeamRequestOption) SetCover(v string) *AddTeamRequestOption {
	s.Cover = &v
	return s
}

func (s *AddTeamRequestOption) SetDescription(v string) *AddTeamRequestOption {
	s.Description = &v
	return s
}

func (s *AddTeamRequestOption) SetIcon(v *AddTeamRequestOptionIcon) *AddTeamRequestOption {
	s.Icon = v
	return s
}

type AddTeamRequestOptionIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddTeamRequestOptionIcon) String() string {
	return tea.Prettify(s)
}

func (s AddTeamRequestOptionIcon) GoString() string {
	return s.String()
}

func (s *AddTeamRequestOptionIcon) SetType(v string) *AddTeamRequestOptionIcon {
	s.Type = &v
	return s
}

func (s *AddTeamRequestOptionIcon) SetValue(v string) *AddTeamRequestOptionIcon {
	s.Value = &v
	return s
}

type AddTeamResponseBody struct {
	Team *AddTeamResponseBodyTeam `json:"team,omitempty" xml:"team,omitempty" type:"Struct"`
}

func (s AddTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTeamResponseBody) GoString() string {
	return s.String()
}

func (s *AddTeamResponseBody) SetTeam(v *AddTeamResponseBodyTeam) *AddTeamResponseBody {
	s.Team = v
	return s
}

type AddTeamResponseBodyTeam struct {
	CorpId       *string                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover        *string                      `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime   *string                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId    *string                      `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description  *string                      `json:"description,omitempty" xml:"description,omitempty"`
	Icon         *AddTeamResponseBodyTeamIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime *string                      `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId   *string                      `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name         *string                      `json:"name,omitempty" xml:"name,omitempty"`
	TeamId       *string                      `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s AddTeamResponseBodyTeam) String() string {
	return tea.Prettify(s)
}

func (s AddTeamResponseBodyTeam) GoString() string {
	return s.String()
}

func (s *AddTeamResponseBodyTeam) SetCorpId(v string) *AddTeamResponseBodyTeam {
	s.CorpId = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetCover(v string) *AddTeamResponseBodyTeam {
	s.Cover = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetCreateTime(v string) *AddTeamResponseBodyTeam {
	s.CreateTime = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetCreatorId(v string) *AddTeamResponseBodyTeam {
	s.CreatorId = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetDescription(v string) *AddTeamResponseBodyTeam {
	s.Description = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetIcon(v *AddTeamResponseBodyTeamIcon) *AddTeamResponseBodyTeam {
	s.Icon = v
	return s
}

func (s *AddTeamResponseBodyTeam) SetModifiedTime(v string) *AddTeamResponseBodyTeam {
	s.ModifiedTime = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetModifierId(v string) *AddTeamResponseBodyTeam {
	s.ModifierId = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetName(v string) *AddTeamResponseBodyTeam {
	s.Name = &v
	return s
}

func (s *AddTeamResponseBodyTeam) SetTeamId(v string) *AddTeamResponseBodyTeam {
	s.TeamId = &v
	return s
}

type AddTeamResponseBodyTeamIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddTeamResponseBodyTeamIcon) String() string {
	return tea.Prettify(s)
}

func (s AddTeamResponseBodyTeamIcon) GoString() string {
	return s.String()
}

func (s *AddTeamResponseBodyTeamIcon) SetType(v string) *AddTeamResponseBodyTeamIcon {
	s.Type = &v
	return s
}

func (s *AddTeamResponseBodyTeamIcon) SetValue(v string) *AddTeamResponseBodyTeamIcon {
	s.Value = &v
	return s
}

type AddTeamResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTeamResponse) GoString() string {
	return s.String()
}

func (s *AddTeamResponse) SetHeaders(v map[string]*string) *AddTeamResponse {
	s.Headers = v
	return s
}

func (s *AddTeamResponse) SetStatusCode(v int32) *AddTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTeamResponse) SetBody(v *AddTeamResponseBody) *AddTeamResponse {
	s.Body = v
	return s
}

type AddWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *AddWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddWorkspaceRequest struct {
	Name       *string                    `json:"name,omitempty" xml:"name,omitempty"`
	Option     *AddWorkspaceRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	OperatorId *string                    `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s AddWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceRequest) SetName(v string) *AddWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *AddWorkspaceRequest) SetOption(v *AddWorkspaceRequestOption) *AddWorkspaceRequest {
	s.Option = v
	return s
}

func (s *AddWorkspaceRequest) SetOperatorId(v string) *AddWorkspaceRequest {
	s.OperatorId = &v
	return s
}

type AddWorkspaceRequestOption struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	TeamId      *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s AddWorkspaceRequestOption) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceRequestOption) GoString() string {
	return s.String()
}

func (s *AddWorkspaceRequestOption) SetDescription(v string) *AddWorkspaceRequestOption {
	s.Description = &v
	return s
}

func (s *AddWorkspaceRequestOption) SetTeamId(v string) *AddWorkspaceRequestOption {
	s.TeamId = &v
	return s
}

type AddWorkspaceResponseBody struct {
	Workspace *AddWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s AddWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBody) SetWorkspace(v *AddWorkspaceResponseBodyWorkspace) *AddWorkspaceResponseBody {
	s.Workspace = v
	return s
}

type AddWorkspaceResponseBodyWorkspace struct {
	CorpId         *string                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover          *string                                `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime     *string                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string                                `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string                                `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *AddWorkspaceResponseBodyWorkspaceIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime   *string                                `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string                                `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string                                `json:"name,omitempty" xml:"name,omitempty"`
	PermissionRole *string                                `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	RootNodeId     *string                                `json:"rootNodeId,omitempty" xml:"rootNodeId,omitempty"`
	TeamId         *string                                `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Type           *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Url            *string                                `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId    *string                                `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s AddWorkspaceResponseBodyWorkspace) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCorpId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCover(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetDescription(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetIcon(v *AddWorkspaceResponseBodyWorkspaceIcon) *AddWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *AddWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetModifierId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetName(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *AddWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetTeamId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetType(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetUrl(v string) *AddWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *AddWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceResponseBodyWorkspaceIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AddWorkspaceResponseBodyWorkspaceIcon) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *AddWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *AddWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *AddWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

type AddWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponse) SetHeaders(v map[string]*string) *AddWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceResponse) SetStatusCode(v int32) *AddWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceResponse) SetBody(v *AddWorkspaceResponseBody) *AddWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTeamHeaders) SetCommonHeaders(v map[string]*string) *DeleteTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTeamHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTeamRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s DeleteTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeamRequest) SetOperatorId(v string) *DeleteTeamRequest {
	s.OperatorId = &v
	return s
}

type DeleteTeamResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponseBody) SetSuccess(v bool) *DeleteTeamResponseBody {
	s.Success = &v
	return s
}

type DeleteTeamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeamResponse) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponse) SetHeaders(v map[string]*string) *DeleteTeamResponse {
	s.Headers = v
	return s
}

func (s *DeleteTeamResponse) SetStatusCode(v int32) *DeleteTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTeamResponse) SetBody(v *DeleteTeamResponseBody) *DeleteTeamResponse {
	s.Body = v
	return s
}

type GetDefaultHandOverUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDefaultHandOverUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultHandOverUserHeaders) GoString() string {
	return s.String()
}

func (s *GetDefaultHandOverUserHeaders) SetCommonHeaders(v map[string]*string) *GetDefaultHandOverUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDefaultHandOverUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetDefaultHandOverUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDefaultHandOverUserRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetDefaultHandOverUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultHandOverUserRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultHandOverUserRequest) SetOperatorId(v string) *GetDefaultHandOverUserRequest {
	s.OperatorId = &v
	return s
}

type GetDefaultHandOverUserResponseBody struct {
	DefaultHandoverUserId *string `json:"defaultHandoverUserId,omitempty" xml:"defaultHandoverUserId,omitempty"`
}

func (s GetDefaultHandOverUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultHandOverUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultHandOverUserResponseBody) SetDefaultHandoverUserId(v string) *GetDefaultHandOverUserResponseBody {
	s.DefaultHandoverUserId = &v
	return s
}

type GetDefaultHandOverUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDefaultHandOverUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultHandOverUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultHandOverUserResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultHandOverUserResponse) SetHeaders(v map[string]*string) *GetDefaultHandOverUserResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultHandOverUserResponse) SetStatusCode(v int32) *GetDefaultHandOverUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultHandOverUserResponse) SetBody(v *GetDefaultHandOverUserResponseBody) *GetDefaultHandOverUserResponse {
	s.Body = v
	return s
}

type GetMineWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMineWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *GetMineWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMineWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetMineWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMineWorkspaceRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetMineWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceRequest) SetOperatorId(v string) *GetMineWorkspaceRequest {
	s.OperatorId = &v
	return s
}

type GetMineWorkspaceResponseBody struct {
	Workspace *GetMineWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetMineWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBody) SetWorkspace(v *GetMineWorkspaceResponseBodyWorkspace) *GetMineWorkspaceResponseBody {
	s.Workspace = v
	return s
}

type GetMineWorkspaceResponseBodyWorkspace struct {
	CorpId         *string                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover          *string                                    `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime     *string                                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string                                    `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *GetMineWorkspaceResponseBodyWorkspaceIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime   *string                                    `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string                                    `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	PermissionRole *string                                    `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	RootNodeId     *string                                    `json:"rootNodeId,omitempty" xml:"rootNodeId,omitempty"`
	TeamId         *string                                    `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Type           *string                                    `json:"type,omitempty" xml:"type,omitempty"`
	Url            *string                                    `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId    *string                                    `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetMineWorkspaceResponseBodyWorkspace) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCorpId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCover(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetDescription(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetIcon(v *GetMineWorkspaceResponseBodyWorkspaceIcon) *GetMineWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetModifierId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetName(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetTeamId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetType(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetUrl(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

type GetMineWorkspaceResponseBodyWorkspaceIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMineWorkspaceResponseBodyWorkspaceIcon) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *GetMineWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *GetMineWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

type GetMineWorkspaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMineWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMineWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMineWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponse) SetHeaders(v map[string]*string) *GetMineWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetMineWorkspaceResponse) SetStatusCode(v int32) *GetMineWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMineWorkspaceResponse) SetBody(v *GetMineWorkspaceResponseBody) *GetMineWorkspaceResponse {
	s.Body = v
	return s
}

type GetNodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNodeHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeHeaders) SetCommonHeaders(v map[string]*string) *GetNodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeHeaders) SetXAcsDingtalkAccessToken(v string) *GetNodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNodeRequest struct {
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	WithPermissionRole  *bool   `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
	WithStatisticalInfo *bool   `json:"withStatisticalInfo,omitempty" xml:"withStatisticalInfo,omitempty"`
}

func (s GetNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) SetOperatorId(v string) *GetNodeRequest {
	s.OperatorId = &v
	return s
}

func (s *GetNodeRequest) SetWithPermissionRole(v bool) *GetNodeRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodeRequest) SetWithStatisticalInfo(v bool) *GetNodeRequest {
	s.WithStatisticalInfo = &v
	return s
}

type GetNodeResponseBody struct {
	Node *GetNodeResponseBodyNode `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
}

func (s GetNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

type GetNodeResponseBodyNode struct {
	Category        *string                                 `json:"category,omitempty" xml:"category,omitempty"`
	CreateTime      *string                                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId       *string                                 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension       *string                                 `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren     *bool                                   `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	ModifiedTime    *string                                 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId      *string                                 `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name            *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	NodeId          *string                                 `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	PermissionRole  *string                                 `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	Size            *int64                                  `json:"size,omitempty" xml:"size,omitempty"`
	StatisticalInfo *GetNodeResponseBodyNodeStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	Type            *string                                 `json:"type,omitempty" xml:"type,omitempty"`
	Url             *string                                 `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId     *string                                 `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetNodeResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) SetCategory(v string) *GetNodeResponseBodyNode {
	s.Category = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v string) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetCreatorId(v string) *GetNodeResponseBodyNode {
	s.CreatorId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetExtension(v string) *GetNodeResponseBodyNode {
	s.Extension = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetHasChildren(v bool) *GetNodeResponseBodyNode {
	s.HasChildren = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifiedTime(v string) *GetNodeResponseBodyNode {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifierId(v string) *GetNodeResponseBodyNode {
	s.ModifierId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetNodeId(v string) *GetNodeResponseBodyNode {
	s.NodeId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetPermissionRole(v string) *GetNodeResponseBodyNode {
	s.PermissionRole = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSize(v int64) *GetNodeResponseBodyNode {
	s.Size = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetStatisticalInfo(v *GetNodeResponseBodyNodeStatisticalInfo) *GetNodeResponseBodyNode {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodeResponseBodyNode) SetType(v string) *GetNodeResponseBodyNode {
	s.Type = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetUrl(v string) *GetNodeResponseBodyNode {
	s.Url = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetWorkspaceId(v string) *GetNodeResponseBodyNode {
	s.WorkspaceId = &v
	return s
}

type GetNodeResponseBodyNodeStatisticalInfo struct {
	WordCount *int64 `json:"wordCount,omitempty" xml:"wordCount,omitempty"`
}

func (s GetNodeResponseBodyNodeStatisticalInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyNodeStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNodeStatisticalInfo) SetWordCount(v int64) *GetNodeResponseBodyNodeStatisticalInfo {
	s.WordCount = &v
	return s
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

type GetNodeByUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNodeByUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlHeaders) SetCommonHeaders(v map[string]*string) *GetNodeByUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeByUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetNodeByUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNodeByUrlRequest struct {
	Option     *GetNodeByUrlRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	Url        *string                    `json:"url,omitempty" xml:"url,omitempty"`
	OperatorId *string                    `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetNodeByUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlRequest) SetOption(v *GetNodeByUrlRequestOption) *GetNodeByUrlRequest {
	s.Option = v
	return s
}

func (s *GetNodeByUrlRequest) SetUrl(v string) *GetNodeByUrlRequest {
	s.Url = &v
	return s
}

func (s *GetNodeByUrlRequest) SetOperatorId(v string) *GetNodeByUrlRequest {
	s.OperatorId = &v
	return s
}

type GetNodeByUrlRequestOption struct {
	WithPermissionRole  *bool `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
	WithStatisticalInfo *bool `json:"withStatisticalInfo,omitempty" xml:"withStatisticalInfo,omitempty"`
}

func (s GetNodeByUrlRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlRequestOption) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlRequestOption) SetWithPermissionRole(v bool) *GetNodeByUrlRequestOption {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodeByUrlRequestOption) SetWithStatisticalInfo(v bool) *GetNodeByUrlRequestOption {
	s.WithStatisticalInfo = &v
	return s
}

type GetNodeByUrlResponseBody struct {
	Node *GetNodeByUrlResponseBodyNode `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
}

func (s GetNodeByUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBody) SetNode(v *GetNodeByUrlResponseBodyNode) *GetNodeByUrlResponseBody {
	s.Node = v
	return s
}

type GetNodeByUrlResponseBodyNode struct {
	Category        *string                                      `json:"category,omitempty" xml:"category,omitempty"`
	CreateTime      *string                                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId       *string                                      `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension       *string                                      `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren     *bool                                        `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	ModifiedTime    *string                                      `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId      *string                                      `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name            *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	NodeId          *string                                      `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	PermissionRole  *string                                      `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	Size            *int64                                       `json:"size,omitempty" xml:"size,omitempty"`
	StatisticalInfo *GetNodeByUrlResponseBodyNodeStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	Type            *string                                      `json:"type,omitempty" xml:"type,omitempty"`
	Url             *string                                      `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId     *string                                      `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetNodeByUrlResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBodyNode) SetCategory(v string) *GetNodeByUrlResponseBodyNode {
	s.Category = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetCreateTime(v string) *GetNodeByUrlResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetCreatorId(v string) *GetNodeByUrlResponseBodyNode {
	s.CreatorId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetExtension(v string) *GetNodeByUrlResponseBodyNode {
	s.Extension = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetHasChildren(v bool) *GetNodeByUrlResponseBodyNode {
	s.HasChildren = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetModifiedTime(v string) *GetNodeByUrlResponseBodyNode {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetModifierId(v string) *GetNodeByUrlResponseBodyNode {
	s.ModifierId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetName(v string) *GetNodeByUrlResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetNodeId(v string) *GetNodeByUrlResponseBodyNode {
	s.NodeId = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetPermissionRole(v string) *GetNodeByUrlResponseBodyNode {
	s.PermissionRole = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetSize(v int64) *GetNodeByUrlResponseBodyNode {
	s.Size = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetStatisticalInfo(v *GetNodeByUrlResponseBodyNodeStatisticalInfo) *GetNodeByUrlResponseBodyNode {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetType(v string) *GetNodeByUrlResponseBodyNode {
	s.Type = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetUrl(v string) *GetNodeByUrlResponseBodyNode {
	s.Url = &v
	return s
}

func (s *GetNodeByUrlResponseBodyNode) SetWorkspaceId(v string) *GetNodeByUrlResponseBodyNode {
	s.WorkspaceId = &v
	return s
}

type GetNodeByUrlResponseBodyNodeStatisticalInfo struct {
	WordCount *int64 `json:"wordCount,omitempty" xml:"wordCount,omitempty"`
}

func (s GetNodeByUrlResponseBodyNodeStatisticalInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlResponseBodyNodeStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponseBodyNodeStatisticalInfo) SetWordCount(v int64) *GetNodeByUrlResponseBodyNodeStatisticalInfo {
	s.WordCount = &v
	return s
}

type GetNodeByUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodeByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodeByUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponse) SetHeaders(v map[string]*string) *GetNodeByUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByUrlResponse) SetStatusCode(v int32) *GetNodeByUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByUrlResponse) SetBody(v *GetNodeByUrlResponseBody) *GetNodeByUrlResponse {
	s.Body = v
	return s
}

type GetNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNodesHeaders) GoString() string {
	return s.String()
}

func (s *GetNodesHeaders) SetCommonHeaders(v map[string]*string) *GetNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodesHeaders) SetXAcsDingtalkAccessToken(v string) *GetNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNodesRequest struct {
	NodeIds    []*string              `json:"nodeIds,omitempty" xml:"nodeIds,omitempty" type:"Repeated"`
	Option     *GetNodesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	OperatorId *string                `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodesRequest) GoString() string {
	return s.String()
}

func (s *GetNodesRequest) SetNodeIds(v []*string) *GetNodesRequest {
	s.NodeIds = v
	return s
}

func (s *GetNodesRequest) SetOption(v *GetNodesRequestOption) *GetNodesRequest {
	s.Option = v
	return s
}

func (s *GetNodesRequest) SetOperatorId(v string) *GetNodesRequest {
	s.OperatorId = &v
	return s
}

type GetNodesRequestOption struct {
	WithPermissionRole  *bool `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
	WithStatisticalInfo *bool `json:"withStatisticalInfo,omitempty" xml:"withStatisticalInfo,omitempty"`
}

func (s GetNodesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetNodesRequestOption) GoString() string {
	return s.String()
}

func (s *GetNodesRequestOption) SetWithPermissionRole(v bool) *GetNodesRequestOption {
	s.WithPermissionRole = &v
	return s
}

func (s *GetNodesRequestOption) SetWithStatisticalInfo(v bool) *GetNodesRequestOption {
	s.WithStatisticalInfo = &v
	return s
}

type GetNodesResponseBody struct {
	Nodes []*GetNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s GetNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBody) SetNodes(v []*GetNodesResponseBodyNodes) *GetNodesResponseBody {
	s.Nodes = v
	return s
}

type GetNodesResponseBodyNodes struct {
	Category        *string                                   `json:"category,omitempty" xml:"category,omitempty"`
	CreateTime      *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId       *string                                   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension       *string                                   `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren     *bool                                     `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	ModifiedTime    *string                                   `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId      *string                                   `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name            *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	NodeId          *string                                   `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	PermissionRole  *string                                   `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	Size            *int64                                    `json:"size,omitempty" xml:"size,omitempty"`
	StatisticalInfo *GetNodesResponseBodyNodesStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	Type            *string                                   `json:"type,omitempty" xml:"type,omitempty"`
	Url             *string                                   `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId     *string                                   `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s GetNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBodyNodes) SetCategory(v string) *GetNodesResponseBodyNodes {
	s.Category = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetCreateTime(v string) *GetNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetCreatorId(v string) *GetNodesResponseBodyNodes {
	s.CreatorId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetExtension(v string) *GetNodesResponseBodyNodes {
	s.Extension = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetHasChildren(v bool) *GetNodesResponseBodyNodes {
	s.HasChildren = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetModifiedTime(v string) *GetNodesResponseBodyNodes {
	s.ModifiedTime = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetModifierId(v string) *GetNodesResponseBodyNodes {
	s.ModifierId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetName(v string) *GetNodesResponseBodyNodes {
	s.Name = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetNodeId(v string) *GetNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetPermissionRole(v string) *GetNodesResponseBodyNodes {
	s.PermissionRole = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetSize(v int64) *GetNodesResponseBodyNodes {
	s.Size = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetStatisticalInfo(v *GetNodesResponseBodyNodesStatisticalInfo) *GetNodesResponseBodyNodes {
	s.StatisticalInfo = v
	return s
}

func (s *GetNodesResponseBodyNodes) SetType(v string) *GetNodesResponseBodyNodes {
	s.Type = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetUrl(v string) *GetNodesResponseBodyNodes {
	s.Url = &v
	return s
}

func (s *GetNodesResponseBodyNodes) SetWorkspaceId(v string) *GetNodesResponseBodyNodes {
	s.WorkspaceId = &v
	return s
}

type GetNodesResponseBodyNodesStatisticalInfo struct {
	WordCount *int64 `json:"wordCount,omitempty" xml:"wordCount,omitempty"`
}

func (s GetNodesResponseBodyNodesStatisticalInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNodesResponseBodyNodesStatisticalInfo) GoString() string {
	return s.String()
}

func (s *GetNodesResponseBodyNodesStatisticalInfo) SetWordCount(v int64) *GetNodesResponseBodyNodesStatisticalInfo {
	s.WordCount = &v
	return s
}

type GetNodesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodesResponse) GoString() string {
	return s.String()
}

func (s *GetNodesResponse) SetHeaders(v map[string]*string) *GetNodesResponse {
	s.Headers = v
	return s
}

func (s *GetNodesResponse) SetStatusCode(v int32) *GetNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodesResponse) SetBody(v *GetNodesResponseBody) *GetNodesResponse {
	s.Body = v
	return s
}

type GetTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTeamHeaders) GoString() string {
	return s.String()
}

func (s *GetTeamHeaders) SetCommonHeaders(v map[string]*string) *GetTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTeamHeaders) SetXAcsDingtalkAccessToken(v string) *GetTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTeamRequest struct {
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTeamRequest) GoString() string {
	return s.String()
}

func (s *GetTeamRequest) SetOperatorId(v string) *GetTeamRequest {
	s.OperatorId = &v
	return s
}

type GetTeamResponseBody struct {
	Team *GetTeamResponseBodyTeam `json:"team,omitempty" xml:"team,omitempty" type:"Struct"`
}

func (s GetTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTeamResponseBody) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBody) SetTeam(v *GetTeamResponseBodyTeam) *GetTeamResponseBody {
	s.Team = v
	return s
}

type GetTeamResponseBodyTeam struct {
	CorpId       *string                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover        *string                      `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime   *string                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId    *string                      `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description  *string                      `json:"description,omitempty" xml:"description,omitempty"`
	Icon         *GetTeamResponseBodyTeamIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime *string                      `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId   *string                      `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name         *string                      `json:"name,omitempty" xml:"name,omitempty"`
	TeamId       *string                      `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s GetTeamResponseBodyTeam) String() string {
	return tea.Prettify(s)
}

func (s GetTeamResponseBodyTeam) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyTeam) SetCorpId(v string) *GetTeamResponseBodyTeam {
	s.CorpId = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetCover(v string) *GetTeamResponseBodyTeam {
	s.Cover = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetCreateTime(v string) *GetTeamResponseBodyTeam {
	s.CreateTime = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetCreatorId(v string) *GetTeamResponseBodyTeam {
	s.CreatorId = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetDescription(v string) *GetTeamResponseBodyTeam {
	s.Description = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetIcon(v *GetTeamResponseBodyTeamIcon) *GetTeamResponseBodyTeam {
	s.Icon = v
	return s
}

func (s *GetTeamResponseBodyTeam) SetModifiedTime(v string) *GetTeamResponseBodyTeam {
	s.ModifiedTime = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetModifierId(v string) *GetTeamResponseBodyTeam {
	s.ModifierId = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetName(v string) *GetTeamResponseBodyTeam {
	s.Name = &v
	return s
}

func (s *GetTeamResponseBodyTeam) SetTeamId(v string) *GetTeamResponseBodyTeam {
	s.TeamId = &v
	return s
}

type GetTeamResponseBodyTeamIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetTeamResponseBodyTeamIcon) String() string {
	return tea.Prettify(s)
}

func (s GetTeamResponseBodyTeamIcon) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBodyTeamIcon) SetType(v string) *GetTeamResponseBodyTeamIcon {
	s.Type = &v
	return s
}

func (s *GetTeamResponseBodyTeamIcon) SetValue(v string) *GetTeamResponseBodyTeamIcon {
	s.Value = &v
	return s
}

type GetTeamResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTeamResponse) GoString() string {
	return s.String()
}

func (s *GetTeamResponse) SetHeaders(v map[string]*string) *GetTeamResponse {
	s.Headers = v
	return s
}

func (s *GetTeamResponse) SetStatusCode(v int32) *GetTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTeamResponse) SetBody(v *GetTeamResponseBody) *GetTeamResponse {
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

type GetWorkspaceRequest struct {
	OperatorId         *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	WithPermissionRole *bool   `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) SetOperatorId(v string) *GetWorkspaceRequest {
	s.OperatorId = &v
	return s
}

func (s *GetWorkspaceRequest) SetWithPermissionRole(v bool) *GetWorkspaceRequest {
	s.WithPermissionRole = &v
	return s
}

type GetWorkspaceResponseBody struct {
	Workspace *GetWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody {
	s.Workspace = v
	return s
}

type GetWorkspaceResponseBodyWorkspace struct {
	CorpId         *string                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover          *string                                `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime     *string                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string                                `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string                                `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *GetWorkspaceResponseBodyWorkspaceIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime   *string                                `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string                                `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string                                `json:"name,omitempty" xml:"name,omitempty"`
	PermissionRole *string                                `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	RootNodeId     *string                                `json:"rootNodeId,omitempty" xml:"rootNodeId,omitempty"`
	TeamId         *string                                `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Type           *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Url            *string                                `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId    *string                                `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspace) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCorpId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCover(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetDescription(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetIcon(v *GetWorkspaceResponseBodyWorkspaceIcon) *GetWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetModifierId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetName(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *GetWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetTeamId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetType(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetUrl(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

type GetWorkspaceResponseBodyWorkspaceIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspaceIcon) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *GetWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *GetWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type GetWorkspacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWorkspacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspacesHeaders) SetXAcsDingtalkAccessToken(v string) *GetWorkspacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWorkspacesRequest struct {
	Option       *GetWorkspacesRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	WorkspaceIds []*string                   `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty" type:"Repeated"`
	OperatorId   *string                     `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GetWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspacesRequest) SetOption(v *GetWorkspacesRequestOption) *GetWorkspacesRequest {
	s.Option = v
	return s
}

func (s *GetWorkspacesRequest) SetWorkspaceIds(v []*string) *GetWorkspacesRequest {
	s.WorkspaceIds = v
	return s
}

func (s *GetWorkspacesRequest) SetOperatorId(v string) *GetWorkspacesRequest {
	s.OperatorId = &v
	return s
}

type GetWorkspacesRequestOption struct {
	WithPermissionRole *bool `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
}

func (s GetWorkspacesRequestOption) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesRequestOption) GoString() string {
	return s.String()
}

func (s *GetWorkspacesRequestOption) SetWithPermissionRole(v bool) *GetWorkspacesRequestOption {
	s.WithPermissionRole = &v
	return s
}

type GetWorkspacesResponseBody struct {
	Workspaces []*GetWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s GetWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBody) SetWorkspaces(v []*GetWorkspacesResponseBodyWorkspaces) *GetWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type GetWorkspacesResponseBodyWorkspaces struct {
	CorpId         *string                                  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover          *string                                  `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime     *string                                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string                                  `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *GetWorkspacesResponseBodyWorkspacesIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime   *string                                  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string                                  `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	PermissionRole *string                                  `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	RootNodeId     *string                                  `json:"rootNodeId,omitempty" xml:"rootNodeId,omitempty"`
	TeamId         *string                                  `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Type           *string                                  `json:"type,omitempty" xml:"type,omitempty"`
	Url            *string                                  `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId    *string                                  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetCorpId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.CorpId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetCover(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.Cover = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetCreatorId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.CreatorId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetDescription(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetIcon(v *GetWorkspacesResponseBodyWorkspacesIcon) *GetWorkspacesResponseBodyWorkspaces {
	s.Icon = v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetModifiedTime(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetModifierId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.ModifierId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetName(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetPermissionRole(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.PermissionRole = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetRootNodeId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.RootNodeId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetTeamId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.TeamId = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetType(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.Type = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetUrl(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.Url = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *GetWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

type GetWorkspacesResponseBodyWorkspacesIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetWorkspacesResponseBodyWorkspacesIcon) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesResponseBodyWorkspacesIcon) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponseBodyWorkspacesIcon) SetType(v string) *GetWorkspacesResponseBodyWorkspacesIcon {
	s.Type = &v
	return s
}

func (s *GetWorkspacesResponseBodyWorkspacesIcon) SetValue(v string) *GetWorkspacesResponseBodyWorkspacesIcon {
	s.Value = &v
	return s
}

type GetWorkspacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspacesResponse) SetHeaders(v map[string]*string) *GetWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspacesResponse) SetStatusCode(v int32) *GetWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspacesResponse) SetBody(v *GetWorkspacesResponseBody) *GetWorkspacesResponse {
	s.Body = v
	return s
}

type HandOverWorkspaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HandOverWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s HandOverWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *HandOverWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *HandOverWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HandOverWorkspaceHeaders) SetXAcsDingtalkAccessToken(v string) *HandOverWorkspaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HandOverWorkspaceRequest struct {
	SourceOwnerId *string `json:"sourceOwnerId,omitempty" xml:"sourceOwnerId,omitempty"`
	TargetOwnerId *string `json:"targetOwnerId,omitempty" xml:"targetOwnerId,omitempty"`
	WorkspaceId   *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	OperatorId    *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s HandOverWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s HandOverWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *HandOverWorkspaceRequest) SetSourceOwnerId(v string) *HandOverWorkspaceRequest {
	s.SourceOwnerId = &v
	return s
}

func (s *HandOverWorkspaceRequest) SetTargetOwnerId(v string) *HandOverWorkspaceRequest {
	s.TargetOwnerId = &v
	return s
}

func (s *HandOverWorkspaceRequest) SetWorkspaceId(v string) *HandOverWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *HandOverWorkspaceRequest) SetOperatorId(v string) *HandOverWorkspaceRequest {
	s.OperatorId = &v
	return s
}

type HandOverWorkspaceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HandOverWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandOverWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *HandOverWorkspaceResponseBody) SetSuccess(v bool) *HandOverWorkspaceResponseBody {
	s.Success = &v
	return s
}

type HandOverWorkspaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HandOverWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandOverWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s HandOverWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *HandOverWorkspaceResponse) SetHeaders(v map[string]*string) *HandOverWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *HandOverWorkspaceResponse) SetStatusCode(v int32) *HandOverWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *HandOverWorkspaceResponse) SetBody(v *HandOverWorkspaceResponseBody) *HandOverWorkspaceResponse {
	s.Body = v
	return s
}

type ListNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListNodesHeaders) GoString() string {
	return s.String()
}

func (s *ListNodesHeaders) SetCommonHeaders(v map[string]*string) *ListNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNodesHeaders) SetXAcsDingtalkAccessToken(v string) *ListNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListNodesRequest struct {
	MaxResults         *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId         *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParentNodeId       *string `json:"parentNodeId,omitempty" xml:"parentNodeId,omitempty"`
	WithPermissionRole *bool   `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetMaxResults(v int32) *ListNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodesRequest) SetNextToken(v string) *ListNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodesRequest) SetOperatorId(v string) *ListNodesRequest {
	s.OperatorId = &v
	return s
}

func (s *ListNodesRequest) SetParentNodeId(v string) *ListNodesRequest {
	s.ParentNodeId = &v
	return s
}

func (s *ListNodesRequest) SetWithPermissionRole(v bool) *ListNodesRequest {
	s.WithPermissionRole = &v
	return s
}

type ListNodesResponseBody struct {
	NextToken *string                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Nodes     []*ListNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNextToken(v string) *ListNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodesResponseBody) SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

type ListNodesResponseBodyNodes struct {
	Category        *string                                    `json:"category,omitempty" xml:"category,omitempty"`
	CreateTime      *string                                    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId       *string                                    `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Extension       *string                                    `json:"extension,omitempty" xml:"extension,omitempty"`
	HasChildren     *bool                                      `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	ModifiedTime    *string                                    `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId      *string                                    `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name            *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	NodeId          *string                                    `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	PermissionRole  *string                                    `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	Size            *int64                                     `json:"size,omitempty" xml:"size,omitempty"`
	StatisticalInfo *ListNodesResponseBodyNodesStatisticalInfo `json:"statisticalInfo,omitempty" xml:"statisticalInfo,omitempty" type:"Struct"`
	Type            *string                                    `json:"type,omitempty" xml:"type,omitempty"`
	Url             *string                                    `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId     *string                                    `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetCategory(v string) *ListNodesResponseBodyNodes {
	s.Category = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreateTime(v string) *ListNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreatorId(v string) *ListNodesResponseBodyNodes {
	s.CreatorId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExtension(v string) *ListNodesResponseBodyNodes {
	s.Extension = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHasChildren(v bool) *ListNodesResponseBodyNodes {
	s.HasChildren = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetModifiedTime(v string) *ListNodesResponseBodyNodes {
	s.ModifiedTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetModifierId(v string) *ListNodesResponseBodyNodes {
	s.ModifierId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetName(v string) *ListNodesResponseBodyNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetNodeId(v string) *ListNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetPermissionRole(v string) *ListNodesResponseBodyNodes {
	s.PermissionRole = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetSize(v int64) *ListNodesResponseBodyNodes {
	s.Size = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetStatisticalInfo(v *ListNodesResponseBodyNodesStatisticalInfo) *ListNodesResponseBodyNodes {
	s.StatisticalInfo = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetType(v string) *ListNodesResponseBodyNodes {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetUrl(v string) *ListNodesResponseBodyNodes {
	s.Url = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetWorkspaceId(v string) *ListNodesResponseBodyNodes {
	s.WorkspaceId = &v
	return s
}

type ListNodesResponseBodyNodesStatisticalInfo struct {
	WordCount *int64 `json:"wordCount,omitempty" xml:"wordCount,omitempty"`
}

func (s ListNodesResponseBodyNodesStatisticalInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesStatisticalInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesStatisticalInfo) SetWordCount(v int64) *ListNodesResponseBodyNodesStatisticalInfo {
	s.WordCount = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListTeamsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTeamsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTeamsHeaders) SetXAcsDingtalkAccessToken(v string) *ListTeamsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTeamsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s ListTeamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListTeamsRequest) SetMaxResults(v int32) *ListTeamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsRequest) SetNextToken(v string) *ListTeamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamsRequest) SetOperatorId(v string) *ListTeamsRequest {
	s.OperatorId = &v
	return s
}

type ListTeamsResponseBody struct {
	NextToken *string                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Teams     []*ListTeamsResponseBodyTeams `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
}

func (s ListTeamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBody) SetNextToken(v string) *ListTeamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamsResponseBody) SetTeams(v []*ListTeamsResponseBodyTeams) *ListTeamsResponseBody {
	s.Teams = v
	return s
}

type ListTeamsResponseBodyTeams struct {
	CorpId       *string                         `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover        *string                         `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime   *string                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId    *string                         `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description  *string                         `json:"description,omitempty" xml:"description,omitempty"`
	Icon         *ListTeamsResponseBodyTeamsIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime *string                         `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId   *string                         `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name         *string                         `json:"name,omitempty" xml:"name,omitempty"`
	TeamId       *string                         `json:"teamId,omitempty" xml:"teamId,omitempty"`
}

func (s ListTeamsResponseBodyTeams) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsResponseBodyTeams) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyTeams) SetCorpId(v string) *ListTeamsResponseBodyTeams {
	s.CorpId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCover(v string) *ListTeamsResponseBodyTeams {
	s.Cover = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCreateTime(v string) *ListTeamsResponseBodyTeams {
	s.CreateTime = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCreatorId(v string) *ListTeamsResponseBodyTeams {
	s.CreatorId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetDescription(v string) *ListTeamsResponseBodyTeams {
	s.Description = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetIcon(v *ListTeamsResponseBodyTeamsIcon) *ListTeamsResponseBodyTeams {
	s.Icon = v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetModifiedTime(v string) *ListTeamsResponseBodyTeams {
	s.ModifiedTime = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetModifierId(v string) *ListTeamsResponseBodyTeams {
	s.ModifierId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetName(v string) *ListTeamsResponseBodyTeams {
	s.Name = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetTeamId(v string) *ListTeamsResponseBodyTeams {
	s.TeamId = &v
	return s
}

type ListTeamsResponseBodyTeamsIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTeamsResponseBodyTeamsIcon) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsResponseBodyTeamsIcon) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyTeamsIcon) SetType(v string) *ListTeamsResponseBodyTeamsIcon {
	s.Type = &v
	return s
}

func (s *ListTeamsResponseBodyTeamsIcon) SetValue(v string) *ListTeamsResponseBodyTeamsIcon {
	s.Value = &v
	return s
}

type ListTeamsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTeamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListTeamsResponse) SetHeaders(v map[string]*string) *ListTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListTeamsResponse) SetStatusCode(v int32) *ListTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTeamsResponse) SetBody(v *ListTeamsResponseBody) *ListTeamsResponse {
	s.Body = v
	return s
}

type ListWorkspacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListWorkspacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesHeaders) GoString() string {
	return s.String()
}

func (s *ListWorkspacesHeaders) SetCommonHeaders(v map[string]*string) *ListWorkspacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListWorkspacesHeaders) SetXAcsDingtalkAccessToken(v string) *ListWorkspacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListWorkspacesRequest struct {
	MaxResults         *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorId         *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OrderBy            *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	TeamId             *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	WithPermissionRole *bool   `json:"withPermissionRole,omitempty" xml:"withPermissionRole,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetOperatorId(v string) *ListWorkspacesRequest {
	s.OperatorId = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrderBy(v string) *ListWorkspacesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetTeamId(v string) *ListWorkspacesRequest {
	s.TeamId = &v
	return s
}

func (s *ListWorkspacesRequest) SetWithPermissionRole(v bool) *ListWorkspacesRequest {
	s.WithPermissionRole = &v
	return s
}

type ListWorkspacesResponseBody struct {
	NextToken  *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	CorpId         *string                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cover          *string                                   `json:"cover,omitempty" xml:"cover,omitempty"`
	CreateTime     *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId      *string                                   `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description    *string                                   `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *ListWorkspacesResponseBodyWorkspacesIcon `json:"icon,omitempty" xml:"icon,omitempty" type:"Struct"`
	ModifiedTime   *string                                   `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId     *string                                   `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	Name           *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	PermissionRole *string                                   `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	RootNodeId     *string                                   `json:"rootNodeId,omitempty" xml:"rootNodeId,omitempty"`
	TeamId         *string                                   `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Type           *string                                   `json:"type,omitempty" xml:"type,omitempty"`
	Url            *string                                   `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId    *string                                   `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCorpId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CorpId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCover(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Cover = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreatorId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreatorId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetIcon(v *ListWorkspacesResponseBodyWorkspacesIcon) *ListWorkspacesResponseBodyWorkspaces {
	s.Icon = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetModifierId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ModifierId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPermissionRole(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PermissionRole = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRootNodeId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RootNodeId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetTeamId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.TeamId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Type = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetUrl(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Url = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

type ListWorkspacesResponseBodyWorkspacesIcon struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesIcon) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesIcon) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) SetType(v string) *ListWorkspacesResponseBodyWorkspacesIcon {
	s.Type = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) SetValue(v string) *ListWorkspacesResponseBodyWorkspacesIcon {
	s.Value = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type SetDefaultHandOverUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDefaultHandOverUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultHandOverUserHeaders) GoString() string {
	return s.String()
}

func (s *SetDefaultHandOverUserHeaders) SetCommonHeaders(v map[string]*string) *SetDefaultHandOverUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDefaultHandOverUserHeaders) SetXAcsDingtalkAccessToken(v string) *SetDefaultHandOverUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDefaultHandOverUserRequest struct {
	DefaultHandoverUserId *string `json:"defaultHandoverUserId,omitempty" xml:"defaultHandoverUserId,omitempty"`
	OperatorId            *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s SetDefaultHandOverUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultHandOverUserRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultHandOverUserRequest) SetDefaultHandoverUserId(v string) *SetDefaultHandOverUserRequest {
	s.DefaultHandoverUserId = &v
	return s
}

func (s *SetDefaultHandOverUserRequest) SetOperatorId(v string) *SetDefaultHandOverUserRequest {
	s.OperatorId = &v
	return s
}

type SetDefaultHandOverUserResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetDefaultHandOverUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultHandOverUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultHandOverUserResponseBody) SetSuccess(v bool) *SetDefaultHandOverUserResponseBody {
	s.Success = &v
	return s
}

type SetDefaultHandOverUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDefaultHandOverUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultHandOverUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultHandOverUserResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultHandOverUserResponse) SetHeaders(v map[string]*string) *SetDefaultHandOverUserResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultHandOverUserResponse) SetStatusCode(v int32) *SetDefaultHandOverUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultHandOverUserResponse) SetBody(v *SetDefaultHandOverUserResponseBody) *SetDefaultHandOverUserResponse {
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) AddTeamWithOptions(request *AddTeamRequest, headers *AddTeamHeaders, runtime *util.RuntimeOptions) (_result *AddTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
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
		Action:      tea.String("AddTeam"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/teams"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTeam(request *AddTeamRequest) (_result *AddTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddTeamHeaders{}
	_result = &AddTeamResponse{}
	_body, _err := client.AddTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceWithOptions(request *AddWorkspaceRequest, headers *AddWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *AddWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
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
		Action:      tea.String("AddWorkspace"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/workspaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWorkspace(request *AddWorkspaceRequest) (_result *AddWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddWorkspaceHeaders{}
	_result = &AddWorkspaceResponse{}
	_body, _err := client.AddWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTeamWithOptions(teamId *string, request *DeleteTeamRequest, headers *DeleteTeamHeaders, runtime *util.RuntimeOptions) (_result *DeleteTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("DeleteTeam"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/teams/" + tea.StringValue(teamId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTeam(teamId *string, request *DeleteTeamRequest) (_result *DeleteTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTeamHeaders{}
	_result = &DeleteTeamResponse{}
	_body, _err := client.DeleteTeamWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultHandOverUserWithOptions(request *GetDefaultHandOverUserRequest, headers *GetDefaultHandOverUserHeaders, runtime *util.RuntimeOptions) (_result *GetDefaultHandOverUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetDefaultHandOverUser"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/managementSettings/defaultHandOverUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultHandOverUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultHandOverUser(request *GetDefaultHandOverUserRequest) (_result *GetDefaultHandOverUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDefaultHandOverUserHeaders{}
	_result = &GetDefaultHandOverUserResponse{}
	_body, _err := client.GetDefaultHandOverUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMineWorkspaceWithOptions(request *GetMineWorkspaceRequest, headers *GetMineWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *GetMineWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetMineWorkspace"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/mineWorkspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMineWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMineWorkspace(request *GetMineWorkspaceRequest) (_result *GetMineWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMineWorkspaceHeaders{}
	_result = &GetMineWorkspaceResponse{}
	_body, _err := client.GetMineWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeWithOptions(nodeId *string, request *GetNodeRequest, headers *GetNodeHeaders, runtime *util.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.WithPermissionRole)) {
		query["withPermissionRole"] = request.WithPermissionRole
	}

	if !tea.BoolValue(util.IsUnset(request.WithStatisticalInfo)) {
		query["withStatisticalInfo"] = request.WithStatisticalInfo
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
		Action:      tea.String("GetNode"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/nodes/" + tea.StringValue(nodeId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNode(nodeId *string, request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNodeHeaders{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(nodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodeByUrlWithOptions(request *GetNodeByUrlRequest, headers *GetNodeByUrlHeaders, runtime *util.RuntimeOptions) (_result *GetNodeByUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
		Action:      tea.String("GetNodeByUrl"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/nodes/queryByUrl"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeByUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodeByUrl(request *GetNodeByUrlRequest) (_result *GetNodeByUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNodeByUrlHeaders{}
	_result = &GetNodeByUrlResponse{}
	_body, _err := client.GetNodeByUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNodesWithOptions(request *GetNodesRequest, headers *GetNodesHeaders, runtime *util.RuntimeOptions) (_result *GetNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIds)) {
		body["nodeIds"] = request.NodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
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
		Action:      tea.String("GetNodes"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/nodes/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNodes(request *GetNodesRequest) (_result *GetNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNodesHeaders{}
	_result = &GetNodesResponse{}
	_body, _err := client.GetNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTeamWithOptions(teamId *string, request *GetTeamRequest, headers *GetTeamHeaders, runtime *util.RuntimeOptions) (_result *GetTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetTeam"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/teams/" + tea.StringValue(teamId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTeam(teamId *string, request *GetTeamRequest) (_result *GetTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTeamHeaders{}
	_result = &GetTeamResponse{}
	_body, _err := client.GetTeamWithOptions(teamId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(workspaceId *string, request *GetWorkspaceRequest, headers *GetWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.WithPermissionRole)) {
		query["withPermissionRole"] = request.WithPermissionRole
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
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/workspaces/" + tea.StringValue(workspaceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(workspaceId *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkspaceHeaders{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspacesWithOptions(request *GetWorkspacesRequest, headers *GetWorkspacesHeaders, runtime *util.RuntimeOptions) (_result *GetWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["option"] = request.Option
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspaces"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/workspaces/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspaces(request *GetWorkspacesRequest) (_result *GetWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWorkspacesHeaders{}
	_result = &GetWorkspacesResponse{}
	_body, _err := client.GetWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandOverWorkspaceWithOptions(request *HandOverWorkspaceRequest, headers *HandOverWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *HandOverWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceOwnerId)) {
		body["sourceOwnerId"] = request.SourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOwnerId)) {
		body["targetOwnerId"] = request.TargetOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
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
		Action:      tea.String("HandOverWorkspace"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/managementOperations/workspaces/handOver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HandOverWorkspaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandOverWorkspace(request *HandOverWorkspaceRequest) (_result *HandOverWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HandOverWorkspaceHeaders{}
	_result = &HandOverWorkspaceResponse{}
	_body, _err := client.HandOverWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, headers *ListNodesHeaders, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ParentNodeId)) {
		query["parentNodeId"] = request.ParentNodeId
	}

	if !tea.BoolValue(util.IsUnset(request.WithPermissionRole)) {
		query["withPermissionRole"] = request.WithPermissionRole
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
		Action:      tea.String("ListNodes"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/nodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListNodesHeaders{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTeamsWithOptions(request *ListTeamsRequest, headers *ListTeamsHeaders, runtime *util.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListTeams"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/teams"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTeamsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTeams(request *ListTeamsRequest) (_result *ListTeamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTeamsHeaders{}
	_result = &ListTeamsResponse{}
	_body, _err := client.ListTeamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers *ListWorkspacesHeaders, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		query["teamId"] = request.TeamId
	}

	if !tea.BoolValue(util.IsUnset(request.WithPermissionRole)) {
		query["withPermissionRole"] = request.WithPermissionRole
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
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/workspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListWorkspacesHeaders{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultHandOverUserWithOptions(request *SetDefaultHandOverUserRequest, headers *SetDefaultHandOverUserHeaders, runtime *util.RuntimeOptions) (_result *SetDefaultHandOverUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultHandoverUserId)) {
		body["defaultHandoverUserId"] = request.DefaultHandoverUserId
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
		Action:      tea.String("SetDefaultHandOverUser"),
		Version:     tea.String("wiki_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/wiki/managementSettings/defaultHandOverUsers/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDefaultHandOverUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultHandOverUser(request *SetDefaultHandOverUserRequest) (_result *SetDefaultHandOverUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDefaultHandOverUserHeaders{}
	_result = &SetDefaultHandOverUserResponse{}
	_body, _err := client.SetDefaultHandOverUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
