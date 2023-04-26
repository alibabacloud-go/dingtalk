// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package conference_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseVideoConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseVideoConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CloseVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseVideoConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CloseVideoConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseVideoConferenceRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CloseVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceRequest) SetUnionId(v string) *CloseVideoConferenceRequest {
	s.UnionId = &v
	return s
}

type CloseVideoConferenceResponseBody struct {
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	Code  *int64  `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CloseVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponseBody) SetCause(v string) *CloseVideoConferenceResponseBody {
	s.Cause = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetCode(v int64) *CloseVideoConferenceResponseBody {
	s.Code = &v
	return s
}

type CloseVideoConferenceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseVideoConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponse) SetHeaders(v map[string]*string) *CloseVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CloseVideoConferenceResponse) SetStatusCode(v int32) *CloseVideoConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseVideoConferenceResponse) SetBody(v *CloseVideoConferenceResponseBody) *CloseVideoConferenceResponse {
	s.Body = v
	return s
}

type CohostsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CohostsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CohostsHeaders) GoString() string {
	return s.String()
}

func (s *CohostsHeaders) SetCommonHeaders(v map[string]*string) *CohostsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CohostsHeaders) SetXAcsDingtalkAccessToken(v string) *CohostsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CohostsRequest struct {
	Action   *string                   `json:"action,omitempty" xml:"action,omitempty"`
	UserList []*CohostsRequestUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s CohostsRequest) String() string {
	return tea.Prettify(s)
}

func (s CohostsRequest) GoString() string {
	return s.String()
}

func (s *CohostsRequest) SetAction(v string) *CohostsRequest {
	s.Action = &v
	return s
}

func (s *CohostsRequest) SetUserList(v []*CohostsRequestUserList) *CohostsRequest {
	s.UserList = v
	return s
}

type CohostsRequestUserList struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CohostsRequestUserList) String() string {
	return tea.Prettify(s)
}

func (s CohostsRequestUserList) GoString() string {
	return s.String()
}

func (s *CohostsRequestUserList) SetUnionId(v string) *CohostsRequestUserList {
	s.UnionId = &v
	return s
}

type CohostsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CohostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CohostsResponseBody) GoString() string {
	return s.String()
}

func (s *CohostsResponseBody) SetSuccess(v bool) *CohostsResponseBody {
	s.Success = &v
	return s
}

type CohostsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CohostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CohostsResponse) String() string {
	return tea.Prettify(s)
}

func (s CohostsResponse) GoString() string {
	return s.String()
}

func (s *CohostsResponse) SetHeaders(v map[string]*string) *CohostsResponse {
	s.Headers = v
	return s
}

func (s *CohostsResponse) SetStatusCode(v int32) *CohostsResponse {
	s.StatusCode = &v
	return s
}

func (s *CohostsResponse) SetBody(v *CohostsResponseBody) *CohostsResponse {
	s.Body = v
	return s
}

type CreateVideoConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateVideoConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVideoConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateVideoConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateVideoConferenceRequest struct {
	ConfTitle     *string   `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	InviteCaller  *bool     `json:"inviteCaller,omitempty" xml:"inviteCaller,omitempty"`
	InviteUserIds []*string `json:"inviteUserIds,omitempty" xml:"inviteUserIds,omitempty" type:"Repeated"`
	UserId        *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceRequest) SetConfTitle(v string) *CreateVideoConferenceRequest {
	s.ConfTitle = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteCaller(v bool) *CreateVideoConferenceRequest {
	s.InviteCaller = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteUserIds(v []*string) *CreateVideoConferenceRequest {
	s.InviteUserIds = v
	return s
}

func (s *CreateVideoConferenceRequest) SetUserId(v string) *CreateVideoConferenceRequest {
	s.UserId = &v
	return s
}

type CreateVideoConferenceResponseBody struct {
	ConferenceId       *string   `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ConferencePassword *string   `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	ExternalLinkUrl    *string   `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	HostPassword       *string   `json:"hostPassword,omitempty" xml:"hostPassword,omitempty"`
	PhoneNumbers       []*string `json:"phoneNumbers,omitempty" xml:"phoneNumbers,omitempty" type:"Repeated"`
	RoomCode           *string   `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s CreateVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponseBody) SetConferenceId(v string) *CreateVideoConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetConferencePassword(v string) *CreateVideoConferenceResponseBody {
	s.ConferencePassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetExternalLinkUrl(v string) *CreateVideoConferenceResponseBody {
	s.ExternalLinkUrl = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetHostPassword(v string) *CreateVideoConferenceResponseBody {
	s.HostPassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetPhoneNumbers(v []*string) *CreateVideoConferenceResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetRoomCode(v string) *CreateVideoConferenceResponseBody {
	s.RoomCode = &v
	return s
}

type CreateVideoConferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponse) SetHeaders(v map[string]*string) *CreateVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoConferenceResponse) SetStatusCode(v int32) *CreateVideoConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoConferenceResponse) SetBody(v *CreateVideoConferenceResponseBody) *CreateVideoConferenceResponse {
	s.Body = v
	return s
}

type FocusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FocusHeaders) String() string {
	return tea.Prettify(s)
}

func (s FocusHeaders) GoString() string {
	return s.String()
}

func (s *FocusHeaders) SetCommonHeaders(v map[string]*string) *FocusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FocusHeaders) SetXAcsDingtalkAccessToken(v string) *FocusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FocusRequest struct {
	Action  *string `json:"action,omitempty" xml:"action,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s FocusRequest) String() string {
	return tea.Prettify(s)
}

func (s FocusRequest) GoString() string {
	return s.String()
}

func (s *FocusRequest) SetAction(v string) *FocusRequest {
	s.Action = &v
	return s
}

func (s *FocusRequest) SetUnionId(v string) *FocusRequest {
	s.UnionId = &v
	return s
}

type FocusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FocusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FocusResponseBody) GoString() string {
	return s.String()
}

func (s *FocusResponseBody) SetSuccess(v bool) *FocusResponseBody {
	s.Success = &v
	return s
}

type FocusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FocusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FocusResponse) String() string {
	return tea.Prettify(s)
}

func (s FocusResponse) GoString() string {
	return s.String()
}

func (s *FocusResponse) SetHeaders(v map[string]*string) *FocusResponse {
	s.Headers = v
	return s
}

func (s *FocusResponse) SetStatusCode(v int32) *FocusResponse {
	s.StatusCode = &v
	return s
}

func (s *FocusResponse) SetBody(v *FocusResponseBody) *FocusResponse {
	s.Body = v
	return s
}

type InviteUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InviteUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersHeaders) GoString() string {
	return s.String()
}

func (s *InviteUsersHeaders) SetCommonHeaders(v map[string]*string) *InviteUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InviteUsersHeaders) SetXAcsDingtalkAccessToken(v string) *InviteUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InviteUsersRequest struct {
	InviteeList []*InviteUsersRequestInviteeList `json:"inviteeList,omitempty" xml:"inviteeList,omitempty" type:"Repeated"`
	UnionId     *string                          `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s InviteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersRequest) GoString() string {
	return s.String()
}

func (s *InviteUsersRequest) SetInviteeList(v []*InviteUsersRequestInviteeList) *InviteUsersRequest {
	s.InviteeList = v
	return s
}

func (s *InviteUsersRequest) SetUnionId(v string) *InviteUsersRequest {
	s.UnionId = &v
	return s
}

type InviteUsersRequestInviteeList struct {
	Nick    *string `json:"nick,omitempty" xml:"nick,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s InviteUsersRequestInviteeList) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersRequestInviteeList) GoString() string {
	return s.String()
}

func (s *InviteUsersRequestInviteeList) SetNick(v string) *InviteUsersRequestInviteeList {
	s.Nick = &v
	return s
}

func (s *InviteUsersRequestInviteeList) SetUnionId(v string) *InviteUsersRequestInviteeList {
	s.UnionId = &v
	return s
}

type InviteUsersResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InviteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *InviteUsersResponseBody) SetSuccess(v bool) *InviteUsersResponseBody {
	s.Success = &v
	return s
}

type InviteUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InviteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InviteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersResponse) GoString() string {
	return s.String()
}

func (s *InviteUsersResponse) SetHeaders(v map[string]*string) *InviteUsersResponse {
	s.Headers = v
	return s
}

func (s *InviteUsersResponse) SetStatusCode(v int32) *InviteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteUsersResponse) SetBody(v *InviteUsersResponseBody) *InviteUsersResponse {
	s.Body = v
	return s
}

type KickMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s KickMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s KickMembersHeaders) GoString() string {
	return s.String()
}

func (s *KickMembersHeaders) SetCommonHeaders(v map[string]*string) *KickMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *KickMembersHeaders) SetXAcsDingtalkAccessToken(v string) *KickMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type KickMembersRequest struct {
	ForbiddenRejoin *bool                         `json:"forbiddenRejoin,omitempty" xml:"forbiddenRejoin,omitempty"`
	UserList        []*KickMembersRequestUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s KickMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s KickMembersRequest) GoString() string {
	return s.String()
}

func (s *KickMembersRequest) SetForbiddenRejoin(v bool) *KickMembersRequest {
	s.ForbiddenRejoin = &v
	return s
}

func (s *KickMembersRequest) SetUserList(v []*KickMembersRequestUserList) *KickMembersRequest {
	s.UserList = v
	return s
}

type KickMembersRequestUserList struct {
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s KickMembersRequestUserList) String() string {
	return tea.Prettify(s)
}

func (s KickMembersRequestUserList) GoString() string {
	return s.String()
}

func (s *KickMembersRequestUserList) SetParticipantId(v string) *KickMembersRequestUserList {
	s.ParticipantId = &v
	return s
}

func (s *KickMembersRequestUserList) SetUnionId(v string) *KickMembersRequestUserList {
	s.UnionId = &v
	return s
}

type KickMembersResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s KickMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KickMembersResponseBody) GoString() string {
	return s.String()
}

func (s *KickMembersResponseBody) SetSuccess(v bool) *KickMembersResponseBody {
	s.Success = &v
	return s
}

type KickMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *KickMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KickMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s KickMembersResponse) GoString() string {
	return s.String()
}

func (s *KickMembersResponse) SetHeaders(v map[string]*string) *KickMembersResponse {
	s.Headers = v
	return s
}

func (s *KickMembersResponse) SetStatusCode(v int32) *KickMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *KickMembersResponse) SetBody(v *KickMembersResponseBody) *KickMembersResponse {
	s.Body = v
	return s
}

type MuteAllHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MuteAllHeaders) String() string {
	return tea.Prettify(s)
}

func (s MuteAllHeaders) GoString() string {
	return s.String()
}

func (s *MuteAllHeaders) SetCommonHeaders(v map[string]*string) *MuteAllHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteAllHeaders) SetXAcsDingtalkAccessToken(v string) *MuteAllHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MuteAllRequest struct {
	Action    *string `json:"action,omitempty" xml:"action,omitempty"`
	ForceMute *bool   `json:"forceMute,omitempty" xml:"forceMute,omitempty"`
}

func (s MuteAllRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteAllRequest) GoString() string {
	return s.String()
}

func (s *MuteAllRequest) SetAction(v string) *MuteAllRequest {
	s.Action = &v
	return s
}

func (s *MuteAllRequest) SetForceMute(v bool) *MuteAllRequest {
	s.ForceMute = &v
	return s
}

type MuteAllResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MuteAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteAllResponseBody) GoString() string {
	return s.String()
}

func (s *MuteAllResponseBody) SetSuccess(v bool) *MuteAllResponseBody {
	s.Success = &v
	return s
}

type MuteAllResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MuteAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteAllResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteAllResponse) GoString() string {
	return s.String()
}

func (s *MuteAllResponse) SetHeaders(v map[string]*string) *MuteAllResponse {
	s.Headers = v
	return s
}

func (s *MuteAllResponse) SetStatusCode(v int32) *MuteAllResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteAllResponse) SetBody(v *MuteAllResponseBody) *MuteAllResponse {
	s.Body = v
	return s
}

type MuteMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MuteMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s MuteMembersHeaders) GoString() string {
	return s.String()
}

func (s *MuteMembersHeaders) SetCommonHeaders(v map[string]*string) *MuteMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MuteMembersHeaders) SetXAcsDingtalkAccessToken(v string) *MuteMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MuteMembersRequest struct {
	Action   *string                       `json:"action,omitempty" xml:"action,omitempty"`
	UserList []*MuteMembersRequestUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s MuteMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteMembersRequest) GoString() string {
	return s.String()
}

func (s *MuteMembersRequest) SetAction(v string) *MuteMembersRequest {
	s.Action = &v
	return s
}

func (s *MuteMembersRequest) SetUserList(v []*MuteMembersRequestUserList) *MuteMembersRequest {
	s.UserList = v
	return s
}

type MuteMembersRequestUserList struct {
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MuteMembersRequestUserList) String() string {
	return tea.Prettify(s)
}

func (s MuteMembersRequestUserList) GoString() string {
	return s.String()
}

func (s *MuteMembersRequestUserList) SetParticipantId(v string) *MuteMembersRequestUserList {
	s.ParticipantId = &v
	return s
}

func (s *MuteMembersRequestUserList) SetUnionId(v string) *MuteMembersRequestUserList {
	s.UnionId = &v
	return s
}

type MuteMembersResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MuteMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *MuteMembersResponseBody) SetSuccess(v bool) *MuteMembersResponseBody {
	s.Success = &v
	return s
}

type MuteMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MuteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteMembersResponse) GoString() string {
	return s.String()
}

func (s *MuteMembersResponse) SetHeaders(v map[string]*string) *MuteMembersResponse {
	s.Headers = v
	return s
}

func (s *MuteMembersResponse) SetStatusCode(v int32) *MuteMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteMembersResponse) SetBody(v *MuteMembersResponseBody) *MuteMembersResponse {
	s.Body = v
	return s
}

type QueryCloudRecordTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordTextHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordTextHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordTextHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordTextRequest struct {
	Direction  *string `json:"direction,omitempty" xml:"direction,omitempty"`
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordTextRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextRequest) SetDirection(v string) *QueryCloudRecordTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetMaxResults(v int64) *QueryCloudRecordTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetNextToken(v int64) *QueryCloudRecordTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetStartTime(v int64) *QueryCloudRecordTextRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetUnionId(v string) *QueryCloudRecordTextRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordTextResponseBody struct {
	HasMore       *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	ParagraphList []*QueryCloudRecordTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBody) SetHasMore(v bool) *QueryCloudRecordTextResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCloudRecordTextResponseBody) SetParagraphList(v []*QueryCloudRecordTextResponseBodyParagraphList) *QueryCloudRecordTextResponseBody {
	s.ParagraphList = v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphList struct {
	EndTime      *int64                                                       `json:"endTime,omitempty" xml:"endTime,omitempty"`
	NextTtoken   *int64                                                       `json:"nextTtoken,omitempty" xml:"nextTtoken,omitempty"`
	NickName     *string                                                      `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Paragraph    *string                                                      `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	RecordId     *int64                                                       `json:"recordId,omitempty" xml:"recordId,omitempty"`
	SentenceList []*QueryCloudRecordTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	StartTime    *int64                                                       `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *int64                                                       `json:"status,omitempty" xml:"status,omitempty"`
	UnionId      *string                                                      `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNextTtoken(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NextTtoken = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNickName(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetParagraph(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetRecordId(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetSentenceList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceList) *QueryCloudRecordTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStatus(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Status = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetUnionId(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.UnionId = &v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphListSentenceList struct {
	EndTime   *int64                                                               `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Sentence  *string                                                              `json:"sentence,omitempty" xml:"sentence,omitempty"`
	StartTime *int64                                                               `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId   *string                                                              `json:"unionId,omitempty" xml:"unionId,omitempty"`
	WordList  []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetUnionId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.UnionId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Word      *string `json:"word,omitempty" xml:"word,omitempty"`
	WordId    *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

type QueryCloudRecordTextResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCloudRecordTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordTextResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponse) SetHeaders(v map[string]*string) *QueryCloudRecordTextResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordTextResponse) SetStatusCode(v int32) *QueryCloudRecordTextResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordTextResponse) SetBody(v *QueryCloudRecordTextResponseBody) *QueryCloudRecordTextResponse {
	s.Body = v
	return s
}

type QueryCloudRecordVideoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordVideoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordVideoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordVideoRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoRequest) SetUnionId(v string) *QueryCloudRecordVideoRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoResponseBody struct {
	VideoList []*QueryCloudRecordVideoResponseBodyVideoList `json:"videoList,omitempty" xml:"videoList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBody) SetVideoList(v []*QueryCloudRecordVideoResponseBodyVideoList) *QueryCloudRecordVideoResponseBody {
	s.VideoList = v
	return s
}

type QueryCloudRecordVideoResponseBodyVideoList struct {
	Duration   *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	MediaId    *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	RecordId   *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	RecordType *int64  `json:"recordType,omitempty" xml:"recordType,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoResponseBodyVideoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetDuration(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetEndTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetFileSize(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetMediaId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordType(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordType = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRegionId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetStartTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetUnionId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCloudRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoResponse) SetStatusCode(v int32) *QueryCloudRecordVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordVideoResponse) SetBody(v *QueryCloudRecordVideoResponseBody) *QueryCloudRecordVideoResponse {
	s.Body = v
	return s
}

type QueryCloudRecordVideoPlayInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordVideoPlayInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordVideoPlayInfoRequest struct {
	MediaId  *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetMediaId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetRegionId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetUnionId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoPlayInfoResponseBody struct {
	Duration   *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Mp4FileUrl *string `json:"mp4FileUrl,omitempty" xml:"mp4FileUrl,omitempty"`
	PlayUrl    *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetDuration(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetFileSize(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetMp4FileUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Mp4FileUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetPlayUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.PlayUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetStatus(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Status = &v
	return s
}

type QueryCloudRecordVideoPlayInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCloudRecordVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordVideoPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetStatusCode(v int32) *QueryCloudRecordVideoPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetBody(v *QueryCloudRecordVideoPlayInfoResponseBody) *QueryCloudRecordVideoPlayInfoResponse {
	s.Body = v
	return s
}

type QueryConferenceInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceInfoResponseBody struct {
	ConfInfo *QueryConferenceInfoResponseBodyConfInfo `json:"confInfo,omitempty" xml:"confInfo,omitempty" type:"Struct"`
}

func (s QueryConferenceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBody) SetConfInfo(v *QueryConferenceInfoResponseBodyConfInfo) *QueryConferenceInfoResponseBody {
	s.ConfInfo = v
	return s
}

type QueryConferenceInfoResponseBodyConfInfo struct {
	ActiveNum       *int32  `json:"activeNum,omitempty" xml:"activeNum,omitempty"`
	AttendNum       *int32  `json:"attendNum,omitempty" xml:"attendNum,omitempty"`
	ConfDuration    *int64  `json:"confDuration,omitempty" xml:"confDuration,omitempty"`
	ConferenceId    *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	CreatorId       *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CreatorNick     *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	EndTime         *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	InvitedNum      *int32  `json:"invitedNum,omitempty" xml:"invitedNum,omitempty"`
	RoomCode        *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	StartTime       *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryConferenceInfoResponseBodyConfInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponseBodyConfInfo) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetActiveNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.ActiveNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetAttendNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.AttendNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConfDuration(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConfDuration = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConferenceId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorNick(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorNick = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetEndTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.EndTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetExternalLinkUrl(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ExternalLinkUrl = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetInvitedNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.InvitedNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetRoomCode(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStartTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStatus(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetTitle(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.Title = &v
	return s
}

type QueryConferenceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoResponse) SetStatusCode(v int32) *QueryConferenceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceInfoResponse) SetBody(v *QueryConferenceInfoResponseBody) *QueryConferenceInfoResponse {
	s.Body = v
	return s
}

type QueryConferenceInfoBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceInfoBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoBatchHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceInfoBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceInfoBatchRequest struct {
	ConferenceIdList []*string `json:"conferenceIdList,omitempty" xml:"conferenceIdList,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchRequest) SetConferenceIdList(v []*string) *QueryConferenceInfoBatchRequest {
	s.ConferenceIdList = v
	return s
}

type QueryConferenceInfoBatchResponseBody struct {
	Infos []*QueryConferenceInfoBatchResponseBodyInfos `json:"infos,omitempty" xml:"infos,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBody) SetInfos(v []*QueryConferenceInfoBatchResponseBodyInfos) *QueryConferenceInfoBatchResponseBody {
	s.Infos = v
	return s
}

type QueryConferenceInfoBatchResponseBodyInfos struct {
	ConferenceId *string                                              `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	MediaStatus  *int64                                               `json:"mediaStatus,omitempty" xml:"mediaStatus,omitempty"`
	StartTime    *int64                                               `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *int64                                               `json:"status,omitempty" xml:"status,omitempty"`
	Title        *string                                              `json:"title,omitempty" xml:"title,omitempty"`
	UserList     []*QueryConferenceInfoBatchResponseBodyInfosUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchResponseBodyInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBodyInfos) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetConferenceId(v string) *QueryConferenceInfoBatchResponseBodyInfos {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetMediaStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.MediaStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetStartTime(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetTitle(v string) *QueryConferenceInfoBatchResponseBodyInfos {
	s.Title = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetUserList(v []*QueryConferenceInfoBatchResponseBodyInfosUserList) *QueryConferenceInfoBatchResponseBodyInfos {
	s.UserList = v
	return s
}

type QueryConferenceInfoBatchResponseBodyInfosUserList struct {
	AttendStatus      *int64  `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	CameraStatus      *int64  `json:"cameraStatus,omitempty" xml:"cameraStatus,omitempty"`
	MicStatus         *int64  `json:"micStatus,omitempty" xml:"micStatus,omitempty"`
	Nick              *string `json:"nick,omitempty" xml:"nick,omitempty"`
	RejectDescription *string `json:"rejectDescription,omitempty" xml:"rejectDescription,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetAttendStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.AttendStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetCameraStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.CameraStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetMicStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.MicStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetNick(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.Nick = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetRejectDescription(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.RejectDescription = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetUserId(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.UserId = &v
	return s
}

type QueryConferenceInfoBatchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryConferenceInfoBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceInfoBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoBatchResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoBatchResponse) SetStatusCode(v int32) *QueryConferenceInfoBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceInfoBatchResponse) SetBody(v *QueryConferenceInfoBatchResponseBody) *QueryConferenceInfoBatchResponse {
	s.Body = v
	return s
}

type QueryConferenceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceMembersRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryConferenceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersRequest) SetMaxResults(v int32) *QueryConferenceMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceMembersRequest) SetNextToken(v string) *QueryConferenceMembersRequest {
	s.NextToken = &v
	return s
}

type QueryConferenceMembersResponseBody struct {
	MemberModels []*QueryConferenceMembersResponseBodyMemberModels `json:"memberModels,omitempty" xml:"memberModels,omitempty" type:"Repeated"`
	NextToken    *string                                           `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount   *int32                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryConferenceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBody) SetMemberModels(v []*QueryConferenceMembersResponseBodyMemberModels) *QueryConferenceMembersResponseBody {
	s.MemberModels = v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetNextToken(v string) *QueryConferenceMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetTotalCount(v int32) *QueryConferenceMembersResponseBody {
	s.TotalCount = &v
	return s
}

type QueryConferenceMembersResponseBodyMemberModels struct {
	AttendStatus   *int32  `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	CoHost         *bool   `json:"coHost,omitempty" xml:"coHost,omitempty"`
	ConferenceId   *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	Duration       *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Host           *bool   `json:"host,omitempty" xml:"host,omitempty"`
	JoinTime       *int64  `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	LeaveTime      *int64  `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	OuterOrgMember *bool   `json:"outerOrgMember,omitempty" xml:"outerOrgMember,omitempty"`
	PstnJoin       *bool   `json:"pstnJoin,omitempty" xml:"pstnJoin,omitempty"`
	UnionId        *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserNick       *string `json:"userNick,omitempty" xml:"userNick,omitempty"`
}

func (s QueryConferenceMembersResponseBodyMemberModels) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponseBodyMemberModels) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetAttendStatus(v int32) *QueryConferenceMembersResponseBodyMemberModels {
	s.AttendStatus = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetCoHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.CoHost = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetConferenceId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetDuration(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.Duration = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.Host = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetJoinTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.JoinTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetLeaveTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.LeaveTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetOuterOrgMember(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.OuterOrgMember = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetPstnJoin(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.PstnJoin = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUnionId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UnionId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUserNick(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UserNick = &v
	return s
}

type QueryConferenceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryConferenceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponse) SetHeaders(v map[string]*string) *QueryConferenceMembersResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceMembersResponse) SetStatusCode(v int32) *QueryConferenceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceMembersResponse) SetBody(v *QueryConferenceMembersResponseBody) *QueryConferenceMembersResponse {
	s.Body = v
	return s
}

type StartCloudRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCloudRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StartCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudRecordHeaders) SetXAcsDingtalkAccessToken(v string) *StartCloudRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCloudRecordRequest struct {
	Mode                *string `json:"mode,omitempty" xml:"mode,omitempty"`
	SmallWindowPosition *string `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	UnionId             *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StartCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequest) SetMode(v string) *StartCloudRecordRequest {
	s.Mode = &v
	return s
}

func (s *StartCloudRecordRequest) SetSmallWindowPosition(v string) *StartCloudRecordRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartCloudRecordRequest) SetUnionId(v string) *StartCloudRecordRequest {
	s.UnionId = &v
	return s
}

type StartCloudRecordResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StartCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponseBody) SetCode(v string) *StartCloudRecordResponseBody {
	s.Code = &v
	return s
}

type StartCloudRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponse) SetHeaders(v map[string]*string) *StartCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StartCloudRecordResponse) SetStatusCode(v int32) *StartCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCloudRecordResponse) SetBody(v *StartCloudRecordResponseBody) *StartCloudRecordResponse {
	s.Body = v
	return s
}

type StartStreamOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartStreamOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutHeaders) GoString() string {
	return s.String()
}

func (s *StartStreamOutHeaders) SetCommonHeaders(v map[string]*string) *StartStreamOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartStreamOutHeaders) SetXAcsDingtalkAccessToken(v string) *StartStreamOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartStreamOutRequest struct {
	Mode                *string   `json:"mode,omitempty" xml:"mode,omitempty"`
	NeedHostJoin        *bool     `json:"needHostJoin,omitempty" xml:"needHostJoin,omitempty"`
	SmallWindowPosition *string   `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	StreamName          *string   `json:"streamName,omitempty" xml:"streamName,omitempty"`
	StreamUrlList       []*string `json:"streamUrlList,omitempty" xml:"streamUrlList,omitempty" type:"Repeated"`
	UnionId             *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StartStreamOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutRequest) GoString() string {
	return s.String()
}

func (s *StartStreamOutRequest) SetMode(v string) *StartStreamOutRequest {
	s.Mode = &v
	return s
}

func (s *StartStreamOutRequest) SetNeedHostJoin(v bool) *StartStreamOutRequest {
	s.NeedHostJoin = &v
	return s
}

func (s *StartStreamOutRequest) SetSmallWindowPosition(v string) *StartStreamOutRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartStreamOutRequest) SetStreamName(v string) *StartStreamOutRequest {
	s.StreamName = &v
	return s
}

func (s *StartStreamOutRequest) SetStreamUrlList(v []*string) *StartStreamOutRequest {
	s.StreamUrlList = v
	return s
}

func (s *StartStreamOutRequest) SetUnionId(v string) *StartStreamOutRequest {
	s.UnionId = &v
	return s
}

type StartStreamOutResponseBody struct {
	FailStreamMap    map[string]interface{} `json:"failStreamMap,omitempty" xml:"failStreamMap,omitempty"`
	SuccessStreamMap map[string]interface{} `json:"successStreamMap,omitempty" xml:"successStreamMap,omitempty"`
}

func (s StartStreamOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamOutResponseBody) SetFailStreamMap(v map[string]interface{}) *StartStreamOutResponseBody {
	s.FailStreamMap = v
	return s
}

func (s *StartStreamOutResponseBody) SetSuccessStreamMap(v map[string]interface{}) *StartStreamOutResponseBody {
	s.SuccessStreamMap = v
	return s
}

type StartStreamOutResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartStreamOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutResponse) GoString() string {
	return s.String()
}

func (s *StartStreamOutResponse) SetHeaders(v map[string]*string) *StartStreamOutResponse {
	s.Headers = v
	return s
}

func (s *StartStreamOutResponse) SetStatusCode(v int32) *StartStreamOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StartStreamOutResponse) SetBody(v *StartStreamOutResponseBody) *StartStreamOutResponse {
	s.Body = v
	return s
}

type StopCloudRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopCloudRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StopCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudRecordHeaders) SetXAcsDingtalkAccessToken(v string) *StopCloudRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopCloudRecordRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StopCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequest) SetUnionId(v string) *StopCloudRecordRequest {
	s.UnionId = &v
	return s
}

type StopCloudRecordResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StopCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponseBody) SetCode(v string) *StopCloudRecordResponseBody {
	s.Code = &v
	return s
}

type StopCloudRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponse) SetHeaders(v map[string]*string) *StopCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StopCloudRecordResponse) SetStatusCode(v int32) *StopCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCloudRecordResponse) SetBody(v *StopCloudRecordResponseBody) *StopCloudRecordResponse {
	s.Body = v
	return s
}

type StopStreamOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopStreamOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutHeaders) GoString() string {
	return s.String()
}

func (s *StopStreamOutHeaders) SetCommonHeaders(v map[string]*string) *StopStreamOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopStreamOutHeaders) SetXAcsDingtalkAccessToken(v string) *StopStreamOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopStreamOutRequest struct {
	StopAllStream *bool   `json:"stopAllStream,omitempty" xml:"stopAllStream,omitempty"`
	StreamId      *string `json:"streamId,omitempty" xml:"streamId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StopStreamOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutRequest) GoString() string {
	return s.String()
}

func (s *StopStreamOutRequest) SetStopAllStream(v bool) *StopStreamOutRequest {
	s.StopAllStream = &v
	return s
}

func (s *StopStreamOutRequest) SetStreamId(v string) *StopStreamOutRequest {
	s.StreamId = &v
	return s
}

func (s *StopStreamOutRequest) SetUnionId(v string) *StopStreamOutRequest {
	s.UnionId = &v
	return s
}

type StopStreamOutResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StopStreamOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamOutResponseBody) SetCode(v string) *StopStreamOutResponseBody {
	s.Code = &v
	return s
}

type StopStreamOutResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStreamOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutResponse) GoString() string {
	return s.String()
}

func (s *StopStreamOutResponse) SetHeaders(v map[string]*string) *StopStreamOutResponse {
	s.Headers = v
	return s
}

func (s *StopStreamOutResponse) SetStatusCode(v int32) *StopStreamOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStreamOutResponse) SetBody(v *StopStreamOutResponseBody) *StopStreamOutResponse {
	s.Body = v
	return s
}

type UpdateVideoConferenceExtInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVideoConferenceExtInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceExtInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceExtInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVideoConferenceExtInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVideoConferenceExtInfoResponseBody struct {
	Case *string `json:"case,omitempty" xml:"case,omitempty"`
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s UpdateVideoConferenceExtInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoResponseBody) SetCase(v string) *UpdateVideoConferenceExtInfoResponseBody {
	s.Case = &v
	return s
}

func (s *UpdateVideoConferenceExtInfoResponseBody) SetCode(v string) *UpdateVideoConferenceExtInfoResponseBody {
	s.Code = &v
	return s
}

type UpdateVideoConferenceExtInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVideoConferenceExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoConferenceExtInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoResponse) SetHeaders(v map[string]*string) *UpdateVideoConferenceExtInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoConferenceExtInfoResponse) SetStatusCode(v int32) *UpdateVideoConferenceExtInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoConferenceExtInfoResponse) SetBody(v *UpdateVideoConferenceExtInfoResponseBody) *UpdateVideoConferenceExtInfoResponse {
	s.Body = v
	return s
}

type UpdateVideoConferenceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVideoConferenceSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVideoConferenceSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVideoConferenceSettingRequest struct {
	AllowUnmuteSelf           *bool `json:"allowUnmuteSelf,omitempty" xml:"allowUnmuteSelf,omitempty"`
	AutoTransferHost          *bool `json:"autoTransferHost,omitempty" xml:"autoTransferHost,omitempty"`
	ForbiddenShareScreen      *bool `json:"forbiddenShareScreen,omitempty" xml:"forbiddenShareScreen,omitempty"`
	LockConference            *bool `json:"lockConference,omitempty" xml:"lockConference,omitempty"`
	MuteAll                   *bool `json:"muteAll,omitempty" xml:"muteAll,omitempty"`
	OnlyInternalEmployeesJoin *bool `json:"onlyInternalEmployeesJoin,omitempty" xml:"onlyInternalEmployeesJoin,omitempty"`
}

func (s UpdateVideoConferenceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingRequest) SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingRequest {
	s.AllowUnmuteSelf = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingRequest {
	s.AutoTransferHost = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingRequest {
	s.ForbiddenShareScreen = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetLockConference(v bool) *UpdateVideoConferenceSettingRequest {
	s.LockConference = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetMuteAll(v bool) *UpdateVideoConferenceSettingRequest {
	s.MuteAll = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingRequest {
	s.OnlyInternalEmployeesJoin = &v
	return s
}

type UpdateVideoConferenceSettingResponseBody struct {
	Case *string `json:"case,omitempty" xml:"case,omitempty"`
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s UpdateVideoConferenceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCase(v string) *UpdateVideoConferenceSettingResponseBody {
	s.Case = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCode(v string) *UpdateVideoConferenceSettingResponseBody {
	s.Code = &v
	return s
}

type UpdateVideoConferenceSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVideoConferenceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoConferenceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponse) SetHeaders(v map[string]*string) *UpdateVideoConferenceSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) SetStatusCode(v int32) *UpdateVideoConferenceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) SetBody(v *UpdateVideoConferenceSettingResponseBody) *UpdateVideoConferenceSettingResponse {
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

func (client *Client) CloseVideoConferenceWithOptions(conferenceId *string, request *CloseVideoConferenceRequest, headers *CloseVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CloseVideoConferenceResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseVideoConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseVideoConference(conferenceId *string, request *CloseVideoConferenceRequest) (_result *CloseVideoConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseVideoConferenceHeaders{}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.CloseVideoConferenceWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CohostsWithOptions(conferenceId *string, request *CohostsRequest, headers *CohostsHeaders, runtime *util.RuntimeOptions) (_result *CohostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.UserList)) {
		body["userList"] = request.UserList
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
		Action:      tea.String("Cohosts"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/coHosts/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CohostsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Cohosts(conferenceId *string, request *CohostsRequest) (_result *CohostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CohostsHeaders{}
	_result = &CohostsResponse{}
	_body, _err := client.CohostsWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoConferenceWithOptions(request *CreateVideoConferenceRequest, headers *CreateVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfTitle)) {
		body["confTitle"] = request.ConfTitle
	}

	if !tea.BoolValue(util.IsUnset(request.InviteCaller)) {
		body["inviteCaller"] = request.InviteCaller
	}

	if !tea.BoolValue(util.IsUnset(request.InviteUserIds)) {
		body["inviteUserIds"] = request.InviteUserIds
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
		Action:      tea.String("CreateVideoConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoConference(request *CreateVideoConferenceRequest) (_result *CreateVideoConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVideoConferenceHeaders{}
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.CreateVideoConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FocusWithOptions(conferenceId *string, request *FocusRequest, headers *FocusHeaders, runtime *util.RuntimeOptions) (_result *FocusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("Focus"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/focus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FocusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Focus(conferenceId *string, request *FocusRequest) (_result *FocusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FocusHeaders{}
	_result = &FocusResponse{}
	_body, _err := client.FocusWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InviteUsersWithOptions(conferenceId *string, request *InviteUsersRequest, headers *InviteUsersHeaders, runtime *util.RuntimeOptions) (_result *InviteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteeList)) {
		body["inviteeList"] = request.InviteeList
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("InviteUsers"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/users/invite"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InviteUsers(conferenceId *string, request *InviteUsersRequest) (_result *InviteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InviteUsersHeaders{}
	_result = &InviteUsersResponse{}
	_body, _err := client.InviteUsersWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KickMembersWithOptions(conferenceId *string, request *KickMembersRequest, headers *KickMembersHeaders, runtime *util.RuntimeOptions) (_result *KickMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForbiddenRejoin)) {
		body["forbiddenRejoin"] = request.ForbiddenRejoin
	}

	if !tea.BoolValue(util.IsUnset(request.UserList)) {
		body["userList"] = request.UserList
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
		Action:      tea.String("KickMembers"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/members/kick"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &KickMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KickMembers(conferenceId *string, request *KickMembersRequest) (_result *KickMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &KickMembersHeaders{}
	_result = &KickMembersResponse{}
	_body, _err := client.KickMembersWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MuteAllWithOptions(conferenceId *string, request *MuteAllRequest, headers *MuteAllHeaders, runtime *util.RuntimeOptions) (_result *MuteAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.ForceMute)) {
		body["forceMute"] = request.ForceMute
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
		Action:      tea.String("MuteAll"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/allMembers/mute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MuteAllResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteAll(conferenceId *string, request *MuteAllRequest) (_result *MuteAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MuteAllHeaders{}
	_result = &MuteAllResponse{}
	_body, _err := client.MuteAllWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MuteMembersWithOptions(conferenceId *string, request *MuteMembersRequest, headers *MuteMembersHeaders, runtime *util.RuntimeOptions) (_result *MuteMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.UserList)) {
		body["userList"] = request.UserList
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
		Action:      tea.String("MuteMembers"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/members/mute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MuteMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteMembers(conferenceId *string, request *MuteMembersRequest) (_result *MuteMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MuteMembersHeaders{}
	_result = &MuteMembersResponse{}
	_body, _err := client.MuteMembersWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordTextWithOptions(conferenceId *string, request *QueryCloudRecordTextRequest, headers *QueryCloudRecordTextHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryCloudRecordText"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/cloudRecords/getTexts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordText(conferenceId *string, request *QueryCloudRecordTextRequest) (_result *QueryCloudRecordTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordTextHeaders{}
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.QueryCloudRecordTextWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoWithOptions(conferenceId *string, request *QueryCloudRecordVideoRequest, headers *QueryCloudRecordVideoHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordVideoResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCloudRecordVideo"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/cloudRecords/getVideos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordVideo(conferenceId *string, request *QueryCloudRecordVideoRequest) (_result *QueryCloudRecordVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordVideoHeaders{}
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.QueryCloudRecordVideoWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoPlayInfoWithOptions(conferenceId *string, request *QueryCloudRecordVideoPlayInfoRequest, headers *QueryCloudRecordVideoPlayInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("QueryCloudRecordVideoPlayInfo"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/cloudRecords/videos/getPlayInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoPlayInfo(conferenceId *string, request *QueryCloudRecordVideoPlayInfoRequest) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordVideoPlayInfoHeaders{}
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.QueryCloudRecordVideoPlayInfoWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceInfoWithOptions(conferenceId *string, headers *QueryConferenceInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceInfoResponse, _err error) {
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
		Action:      tea.String("QueryConferenceInfo"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceInfo(conferenceId *string) (_result *QueryConferenceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceInfoHeaders{}
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.QueryConferenceInfoWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceInfoBatchWithOptions(request *QueryConferenceInfoBatchRequest, headers *QueryConferenceInfoBatchHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceInfoBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceIdList)) {
		body["conferenceIdList"] = request.ConferenceIdList
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
		Action:      tea.String("QueryConferenceInfoBatch"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConferenceInfoBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceInfoBatch(request *QueryConferenceInfoBatchRequest) (_result *QueryConferenceInfoBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceInfoBatchHeaders{}
	_result = &QueryConferenceInfoBatchResponse{}
	_body, _err := client.QueryConferenceInfoBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceMembersWithOptions(conferenceId *string, request *QueryConferenceMembersRequest, headers *QueryConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceMembersResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("QueryConferenceMembers"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceMembers(conferenceId *string, request *QueryConferenceMembersRequest) (_result *QueryConferenceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceMembersHeaders{}
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.QueryConferenceMembersWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCloudRecordWithOptions(conferenceId *string, request *StartCloudRecordRequest, headers *StartCloudRecordHeaders, runtime *util.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.SmallWindowPosition)) {
		body["smallWindowPosition"] = request.SmallWindowPosition
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("StartCloudRecord"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/cloudRecords/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCloudRecord(conferenceId *string, request *StartCloudRecordRequest) (_result *StartCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCloudRecordHeaders{}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.StartCloudRecordWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartStreamOutWithOptions(conferenceId *string, request *StartStreamOutRequest, headers *StartStreamOutHeaders, runtime *util.RuntimeOptions) (_result *StartStreamOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.NeedHostJoin)) {
		body["needHostJoin"] = request.NeedHostJoin
	}

	if !tea.BoolValue(util.IsUnset(request.SmallWindowPosition)) {
		body["smallWindowPosition"] = request.SmallWindowPosition
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		body["streamName"] = request.StreamName
	}

	if !tea.BoolValue(util.IsUnset(request.StreamUrlList)) {
		body["streamUrlList"] = request.StreamUrlList
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("StartStreamOut"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/streamOuts/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartStreamOutResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartStreamOut(conferenceId *string, request *StartStreamOutRequest) (_result *StartStreamOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartStreamOutHeaders{}
	_result = &StartStreamOutResponse{}
	_body, _err := client.StartStreamOutWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCloudRecordWithOptions(conferenceId *string, request *StopCloudRecordRequest, headers *StopCloudRecordHeaders, runtime *util.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("StopCloudRecord"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/cloudRecords/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCloudRecord(conferenceId *string, request *StopCloudRecordRequest) (_result *StopCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopCloudRecordHeaders{}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.StopCloudRecordWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStreamOutWithOptions(conferenceId *string, request *StopStreamOutRequest, headers *StopStreamOutHeaders, runtime *util.RuntimeOptions) (_result *StopStreamOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StopAllStream)) {
		body["stopAllStream"] = request.StopAllStream
	}

	if !tea.BoolValue(util.IsUnset(request.StreamId)) {
		body["streamId"] = request.StreamId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("StopStreamOut"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/streamOuts/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StopStreamOutResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStreamOut(conferenceId *string, request *StopStreamOutRequest) (_result *StopStreamOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopStreamOutHeaders{}
	_result = &StopStreamOutResponse{}
	_body, _err := client.StopStreamOutWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoConferenceExtInfoWithOptions(conferenceId *string, headers *UpdateVideoConferenceExtInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateVideoConferenceExtInfoResponse, _err error) {
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
		Action:      tea.String("UpdateVideoConferenceExtInfo"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/extInfo"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVideoConferenceExtInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoConferenceExtInfo(conferenceId *string) (_result *UpdateVideoConferenceExtInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVideoConferenceExtInfoHeaders{}
	_result = &UpdateVideoConferenceExtInfoResponse{}
	_body, _err := client.UpdateVideoConferenceExtInfoWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoConferenceSettingWithOptions(conferenceId *string, request *UpdateVideoConferenceSettingRequest, headers *UpdateVideoConferenceSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowUnmuteSelf)) {
		body["allowUnmuteSelf"] = request.AllowUnmuteSelf
	}

	if !tea.BoolValue(util.IsUnset(request.AutoTransferHost)) {
		body["autoTransferHost"] = request.AutoTransferHost
	}

	if !tea.BoolValue(util.IsUnset(request.ForbiddenShareScreen)) {
		body["forbiddenShareScreen"] = request.ForbiddenShareScreen
	}

	if !tea.BoolValue(util.IsUnset(request.LockConference)) {
		body["lockConference"] = request.LockConference
	}

	if !tea.BoolValue(util.IsUnset(request.MuteAll)) {
		body["muteAll"] = request.MuteAll
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyInternalEmployeesJoin)) {
		body["onlyInternalEmployeesJoin"] = request.OnlyInternalEmployeesJoin
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
		Action:      tea.String("UpdateVideoConferenceSetting"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoConferenceSetting(conferenceId *string, request *UpdateVideoConferenceSettingRequest) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVideoConferenceSettingHeaders{}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.UpdateVideoConferenceSettingWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
