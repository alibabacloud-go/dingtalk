// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package rooms_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDeviceCustomTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeviceCustomTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceCustomTemplateHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeviceCustomTemplateHeaders) SetCommonHeaders(v map[string]*string) *CreateDeviceCustomTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeviceCustomTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeviceCustomTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeviceCustomTemplateRequest struct {
	BgImgList                    []*string `json:"bgImgList,omitempty" xml:"bgImgList,omitempty" type:"Repeated"`
	BgType                       *int32    `json:"bgType,omitempty" xml:"bgType,omitempty"`
	BgUrl                        *string   `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	CustomDoc                    *string   `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	DesensitizeUserName          *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	DeviceUnionIds               []*string `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	GroupIds                     []*int64  `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	HideServerCodeWhenProjecting *bool     `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	Instruction                  *bool     `json:"instruction,omitempty" xml:"instruction,omitempty"`
	IsPicTop                     *int32    `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	Logo                         *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	OrgName                      *string   `json:"orgName,omitempty" xml:"orgName,omitempty"`
	PicturePlayInterval          *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	RoomIds                      []*string `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
	ShowCalendarCard             *bool     `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	ShowCalendarTitle            *bool     `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	ShowFunctionCard             *bool     `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	TemplateName                 *string   `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s CreateDeviceCustomTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceCustomTemplateRequest) SetBgImgList(v []*string) *CreateDeviceCustomTemplateRequest {
	s.BgImgList = v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetBgType(v int32) *CreateDeviceCustomTemplateRequest {
	s.BgType = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetBgUrl(v string) *CreateDeviceCustomTemplateRequest {
	s.BgUrl = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetCustomDoc(v string) *CreateDeviceCustomTemplateRequest {
	s.CustomDoc = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetDesensitizeUserName(v bool) *CreateDeviceCustomTemplateRequest {
	s.DesensitizeUserName = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetDeviceUnionIds(v []*string) *CreateDeviceCustomTemplateRequest {
	s.DeviceUnionIds = v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetGroupIds(v []*int64) *CreateDeviceCustomTemplateRequest {
	s.GroupIds = v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetHideServerCodeWhenProjecting(v bool) *CreateDeviceCustomTemplateRequest {
	s.HideServerCodeWhenProjecting = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetInstruction(v bool) *CreateDeviceCustomTemplateRequest {
	s.Instruction = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetIsPicTop(v int32) *CreateDeviceCustomTemplateRequest {
	s.IsPicTop = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetLogo(v string) *CreateDeviceCustomTemplateRequest {
	s.Logo = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetOrgName(v string) *CreateDeviceCustomTemplateRequest {
	s.OrgName = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetPicturePlayInterval(v int32) *CreateDeviceCustomTemplateRequest {
	s.PicturePlayInterval = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetRoomIds(v []*string) *CreateDeviceCustomTemplateRequest {
	s.RoomIds = v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetShowCalendarCard(v bool) *CreateDeviceCustomTemplateRequest {
	s.ShowCalendarCard = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetShowCalendarTitle(v bool) *CreateDeviceCustomTemplateRequest {
	s.ShowCalendarTitle = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetShowFunctionCard(v bool) *CreateDeviceCustomTemplateRequest {
	s.ShowFunctionCard = &v
	return s
}

func (s *CreateDeviceCustomTemplateRequest) SetTemplateName(v string) *CreateDeviceCustomTemplateRequest {
	s.TemplateName = &v
	return s
}

type CreateDeviceCustomTemplateResponseBody struct {
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateDeviceCustomTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceCustomTemplateResponseBody) SetTemplateId(v int64) *CreateDeviceCustomTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateDeviceCustomTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceCustomTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceCustomTemplateResponse) SetHeaders(v map[string]*string) *CreateDeviceCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceCustomTemplateResponse) SetStatusCode(v int32) *CreateDeviceCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceCustomTemplateResponse) SetBody(v *CreateDeviceCustomTemplateResponseBody) *CreateDeviceCustomTemplateResponse {
	s.Body = v
	return s
}

type CreateMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMeetingRoomRequest struct {
	EnableCycleReservation *bool                                         `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	GroupId                *int64                                        `json:"groupId,omitempty" xml:"groupId,omitempty"`
	IsvRoomId              *string                                       `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	ReservationAuthority   *CreateMeetingRoomRequestReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	RoomCapacity           *int32                                        `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomLabelIds           []*int64                                      `json:"roomLabelIds,omitempty" xml:"roomLabelIds,omitempty" type:"Repeated"`
	RoomLocation           *CreateMeetingRoomRequestRoomLocation         `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	RoomName               *string                                       `json:"roomName,omitempty" xml:"roomName,omitempty"`
	RoomPicture            *string                                       `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	RoomStatus             *int32                                        `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	UnionId                *string                                       `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequest) SetEnableCycleReservation(v bool) *CreateMeetingRoomRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetGroupId(v int64) *CreateMeetingRoomRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetIsvRoomId(v string) *CreateMeetingRoomRequest {
	s.IsvRoomId = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetReservationAuthority(v *CreateMeetingRoomRequestReservationAuthority) *CreateMeetingRoomRequest {
	s.ReservationAuthority = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomCapacity(v int32) *CreateMeetingRoomRequest {
	s.RoomCapacity = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomLabelIds(v []*int64) *CreateMeetingRoomRequest {
	s.RoomLabelIds = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomLocation(v *CreateMeetingRoomRequestRoomLocation) *CreateMeetingRoomRequest {
	s.RoomLocation = v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomName(v string) *CreateMeetingRoomRequest {
	s.RoomName = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomPicture(v string) *CreateMeetingRoomRequest {
	s.RoomPicture = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetRoomStatus(v int32) *CreateMeetingRoomRequest {
	s.RoomStatus = &v
	return s
}

func (s *CreateMeetingRoomRequest) SetUnionId(v string) *CreateMeetingRoomRequest {
	s.UnionId = &v
	return s
}

type CreateMeetingRoomRequestReservationAuthority struct {
	AuthorizedMembers []*CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers `json:"authorizedMembers,omitempty" xml:"authorizedMembers,omitempty" type:"Repeated"`
}

func (s CreateMeetingRoomRequestReservationAuthority) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomRequestReservationAuthority) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestReservationAuthority) SetAuthorizedMembers(v []*CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) *CreateMeetingRoomRequestReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

type CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberId(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberName(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberType(v string) *CreateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

type CreateMeetingRoomRequestRoomLocation struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateMeetingRoomRequestRoomLocation) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomRequestRoomLocation) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomRequestRoomLocation) SetDesc(v string) *CreateMeetingRoomRequestRoomLocation {
	s.Desc = &v
	return s
}

func (s *CreateMeetingRoomRequestRoomLocation) SetTitle(v string) *CreateMeetingRoomRequestRoomLocation {
	s.Title = &v
	return s
}

type CreateMeetingRoomResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomResponseBody) SetResult(v string) *CreateMeetingRoomResponseBody {
	s.Result = &v
	return s
}

type CreateMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomResponse) SetHeaders(v map[string]*string) *CreateMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingRoomResponse) SetStatusCode(v int32) *CreateMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingRoomResponse) SetBody(v *CreateMeetingRoomResponseBody) *CreateMeetingRoomResponse {
	s.Body = v
	return s
}

type CreateMeetingRoomGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMeetingRoomGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMeetingRoomGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMeetingRoomGroupRequest struct {
	GroupName     *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ParentGroupId *int64  `json:"parentGroupId,omitempty" xml:"parentGroupId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateMeetingRoomGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupRequest) SetGroupName(v string) *CreateMeetingRoomGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMeetingRoomGroupRequest) SetParentGroupId(v int64) *CreateMeetingRoomGroupRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateMeetingRoomGroupRequest) SetUnionId(v string) *CreateMeetingRoomGroupRequest {
	s.UnionId = &v
	return s
}

type CreateMeetingRoomGroupResponseBody struct {
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateMeetingRoomGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupResponseBody) SetResult(v int64) *CreateMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

type CreateMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMeetingRoomGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *CreateMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingRoomGroupResponse) SetStatusCode(v int32) *CreateMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingRoomGroupResponse) SetBody(v *CreateMeetingRoomGroupResponseBody) *CreateMeetingRoomGroupResponse {
	s.Body = v
	return s
}

type DeleteDeviceCustomTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeviceCustomTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCustomTemplateHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCustomTemplateHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeviceCustomTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeviceCustomTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeviceCustomTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeviceCustomTemplateRequest struct {
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s DeleteDeviceCustomTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCustomTemplateRequest) SetTemplateId(v int64) *DeleteDeviceCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteDeviceCustomTemplateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDeviceCustomTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCustomTemplateResponseBody) SetResult(v bool) *DeleteDeviceCustomTemplateResponseBody {
	s.Result = &v
	return s
}

type DeleteDeviceCustomTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceCustomTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCustomTemplateResponse) SetHeaders(v map[string]*string) *DeleteDeviceCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceCustomTemplateResponse) SetStatusCode(v int32) *DeleteDeviceCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceCustomTemplateResponse) SetBody(v *DeleteDeviceCustomTemplateResponseBody) *DeleteDeviceCustomTemplateResponse {
	s.Body = v
	return s
}

type DeleteMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMeetingRoomRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomRequest) SetUnionId(v string) *DeleteMeetingRoomRequest {
	s.UnionId = &v
	return s
}

type DeleteMeetingRoomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomResponseBody) SetResult(v bool) *DeleteMeetingRoomResponseBody {
	s.Result = &v
	return s
}

type DeleteMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomResponse) SetHeaders(v map[string]*string) *DeleteMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingRoomResponse) SetStatusCode(v int32) *DeleteMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingRoomResponse) SetBody(v *DeleteMeetingRoomResponseBody) *DeleteMeetingRoomResponse {
	s.Body = v
	return s
}

type DeleteMeetingRoomGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMeetingRoomGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMeetingRoomGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMeetingRoomGroupRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteMeetingRoomGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupRequest) SetUnionId(v string) *DeleteMeetingRoomGroupRequest {
	s.UnionId = &v
	return s
}

type DeleteMeetingRoomGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteMeetingRoomGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupResponseBody) SetResult(v bool) *DeleteMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

type DeleteMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMeetingRoomGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *DeleteMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingRoomGroupResponse) SetStatusCode(v int32) *DeleteMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingRoomGroupResponse) SetBody(v *DeleteMeetingRoomGroupResponseBody) *DeleteMeetingRoomGroupResponse {
	s.Body = v
	return s
}

type QueryDeviceCustomTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceCustomTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceCustomTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceCustomTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceCustomTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceCustomTemplateResponseBody struct {
	Result *QueryDeviceCustomTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryDeviceCustomTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateResponseBody) SetResult(v *QueryDeviceCustomTemplateResponseBodyResult) *QueryDeviceCustomTemplateResponseBody {
	s.Result = v
	return s
}

type QueryDeviceCustomTemplateResponseBodyResult struct {
	DeviceCustomTemplate *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate `json:"deviceCustomTemplate,omitempty" xml:"deviceCustomTemplate,omitempty" type:"Struct"`
	DeviceUnionIds       []*string                                                        `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	GroupIds             []*int64                                                         `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	RoomIds              []*string                                                        `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
}

func (s QueryDeviceCustomTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateResponseBodyResult) SetDeviceCustomTemplate(v *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) *QueryDeviceCustomTemplateResponseBodyResult {
	s.DeviceCustomTemplate = v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResult) SetDeviceUnionIds(v []*string) *QueryDeviceCustomTemplateResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResult) SetGroupIds(v []*int64) *QueryDeviceCustomTemplateResponseBodyResult {
	s.GroupIds = v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResult) SetRoomIds(v []*string) *QueryDeviceCustomTemplateResponseBodyResult {
	s.RoomIds = v
	return s
}

type QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate struct {
	BgImageList                  []*string `json:"bgImageList,omitempty" xml:"bgImageList,omitempty" type:"Repeated"`
	BgType                       *int32    `json:"bgType,omitempty" xml:"bgType,omitempty"`
	BgUrl                        *string   `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	ConfSubType                  *int32    `json:"confSubType,omitempty" xml:"confSubType,omitempty"`
	ConfType                     *int32    `json:"confType,omitempty" xml:"confType,omitempty"`
	CorpId                       *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CustomDoc                    *string   `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	DesensitizeUserName          *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	HideServerCodeWhenProjecting *bool     `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	Instruction                  *bool     `json:"instruction,omitempty" xml:"instruction,omitempty"`
	IsPicTop                     *int32    `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	Logo                         *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	OrgName                      *string   `json:"orgName,omitempty" xml:"orgName,omitempty"`
	PicturePlayInterval          *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	ShowCalendarCard             *bool     `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	ShowCalendarTitle            *bool     `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	ShowFunctionCard             *bool     `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	TemplateId                   *int64    `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName                 *string   `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetBgImageList(v []*string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.BgImageList = v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetBgType(v int32) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.BgType = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetBgUrl(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.BgUrl = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetConfSubType(v int32) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.ConfSubType = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetConfType(v int32) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.ConfType = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetCorpId(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.CorpId = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetCustomDoc(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.CustomDoc = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetDesensitizeUserName(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.DesensitizeUserName = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetHideServerCodeWhenProjecting(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.HideServerCodeWhenProjecting = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetInstruction(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.Instruction = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetIsPicTop(v int32) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.IsPicTop = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetLogo(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.Logo = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetOrgName(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.OrgName = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetPicturePlayInterval(v int32) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.PicturePlayInterval = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetShowCalendarCard(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.ShowCalendarCard = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetShowCalendarTitle(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.ShowCalendarTitle = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetShowFunctionCard(v bool) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.ShowFunctionCard = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetTemplateId(v int64) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.TemplateId = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate) SetTemplateName(v string) *QueryDeviceCustomTemplateResponseBodyResultDeviceCustomTemplate {
	s.TemplateName = &v
	return s
}

type QueryDeviceCustomTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceCustomTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateResponse) SetHeaders(v map[string]*string) *QueryDeviceCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceCustomTemplateResponse) SetStatusCode(v int32) *QueryDeviceCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceCustomTemplateResponse) SetBody(v *QueryDeviceCustomTemplateResponseBody) *QueryDeviceCustomTemplateResponse {
	s.Body = v
	return s
}

type QueryDeviceCustomTemplateListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceCustomTemplateListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateListHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateListHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceCustomTemplateListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceCustomTemplateListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceCustomTemplateListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceCustomTemplateListResponseBody struct {
	Result *QueryDeviceCustomTemplateListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryDeviceCustomTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateListResponseBody) SetResult(v *QueryDeviceCustomTemplateListResponseBodyResult) *QueryDeviceCustomTemplateListResponseBody {
	s.Result = v
	return s
}

type QueryDeviceCustomTemplateListResponseBodyResult struct {
	DeviceCustomTemplates []*QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates `json:"deviceCustomTemplates,omitempty" xml:"deviceCustomTemplates,omitempty" type:"Repeated"`
	DeviceTemplateMap     map[string][]*string                                                    `json:"deviceTemplateMap,omitempty" xml:"deviceTemplateMap,omitempty"`
	GroupIdTemplateMap    map[string][]*int64                                                     `json:"groupIdTemplateMap,omitempty" xml:"groupIdTemplateMap,omitempty"`
	RoomIdTemplateMap     map[string][]*string                                                    `json:"roomIdTemplateMap,omitempty" xml:"roomIdTemplateMap,omitempty"`
}

func (s QueryDeviceCustomTemplateListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateListResponseBodyResult) SetDeviceCustomTemplates(v []*QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) *QueryDeviceCustomTemplateListResponseBodyResult {
	s.DeviceCustomTemplates = v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResult) SetDeviceTemplateMap(v map[string][]*string) *QueryDeviceCustomTemplateListResponseBodyResult {
	s.DeviceTemplateMap = v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResult) SetGroupIdTemplateMap(v map[string][]*int64) *QueryDeviceCustomTemplateListResponseBodyResult {
	s.GroupIdTemplateMap = v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResult) SetRoomIdTemplateMap(v map[string][]*string) *QueryDeviceCustomTemplateListResponseBodyResult {
	s.RoomIdTemplateMap = v
	return s
}

type QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates struct {
	BgImageList                  []*string `json:"bgImageList,omitempty" xml:"bgImageList,omitempty" type:"Repeated"`
	BgType                       *int32    `json:"bgType,omitempty" xml:"bgType,omitempty"`
	BgUrl                        *string   `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	ConfSubType                  *int32    `json:"confSubType,omitempty" xml:"confSubType,omitempty"`
	ConfType                     *int32    `json:"confType,omitempty" xml:"confType,omitempty"`
	CorpId                       *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CustomDoc                    *string   `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	DesensitizeUserName          *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	HideServerCodeWhenProjecting *bool     `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	Instruction                  *bool     `json:"instruction,omitempty" xml:"instruction,omitempty"`
	IsPicTop                     *int32    `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	Logo                         *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	OrgName                      *string   `json:"orgName,omitempty" xml:"orgName,omitempty"`
	PicturePlayInterval          *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	ShowCalendarCard             *bool     `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	ShowCalendarTitle            *bool     `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	ShowFunctionCard             *bool     `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	TemplateId                   *int64    `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName                 *string   `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetBgImageList(v []*string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.BgImageList = v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetBgType(v int32) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.BgType = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetBgUrl(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.BgUrl = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetConfSubType(v int32) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.ConfSubType = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetConfType(v int32) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.ConfType = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetCorpId(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.CorpId = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetCustomDoc(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.CustomDoc = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetDesensitizeUserName(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.DesensitizeUserName = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetHideServerCodeWhenProjecting(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.HideServerCodeWhenProjecting = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetInstruction(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.Instruction = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetIsPicTop(v int32) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.IsPicTop = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetLogo(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.Logo = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetOrgName(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.OrgName = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetPicturePlayInterval(v int32) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.PicturePlayInterval = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetShowCalendarCard(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.ShowCalendarCard = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetShowCalendarTitle(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.ShowCalendarTitle = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetShowFunctionCard(v bool) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.ShowFunctionCard = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetTemplateId(v int64) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.TemplateId = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates) SetTemplateName(v string) *QueryDeviceCustomTemplateListResponseBodyResultDeviceCustomTemplates {
	s.TemplateName = &v
	return s
}

type QueryDeviceCustomTemplateListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceCustomTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceCustomTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceCustomTemplateListResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceCustomTemplateListResponse) SetHeaders(v map[string]*string) *QueryDeviceCustomTemplateListResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceCustomTemplateListResponse) SetStatusCode(v int32) *QueryDeviceCustomTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceCustomTemplateListResponse) SetBody(v *QueryDeviceCustomTemplateListResponseBody) *QueryDeviceCustomTemplateListResponse {
	s.Body = v
	return s
}

type QueryDeviceIpByCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceIpByCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceIpByCodeHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceIpByCodeHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceIpByCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceIpByCodeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceIpByCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceIpByCodeRequest struct {
	DeviceSn *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
}

func (s QueryDeviceIpByCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceIpByCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceIpByCodeRequest) SetDeviceSn(v string) *QueryDeviceIpByCodeRequest {
	s.DeviceSn = &v
	return s
}

type QueryDeviceIpByCodeResponseBody struct {
	Result  *QueryDeviceIpByCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryDeviceIpByCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceIpByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceIpByCodeResponseBody) SetResult(v *QueryDeviceIpByCodeResponseBodyResult) *QueryDeviceIpByCodeResponseBody {
	s.Result = v
	return s
}

func (s *QueryDeviceIpByCodeResponseBody) SetSuccess(v bool) *QueryDeviceIpByCodeResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceIpByCodeResponseBodyResult struct {
	DeviceIp *string `json:"deviceIp,omitempty" xml:"deviceIp,omitempty"`
}

func (s QueryDeviceIpByCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceIpByCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceIpByCodeResponseBodyResult) SetDeviceIp(v string) *QueryDeviceIpByCodeResponseBodyResult {
	s.DeviceIp = &v
	return s
}

type QueryDeviceIpByCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceIpByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceIpByCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceIpByCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceIpByCodeResponse) SetHeaders(v map[string]*string) *QueryDeviceIpByCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceIpByCodeResponse) SetStatusCode(v int32) *QueryDeviceIpByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceIpByCodeResponse) SetBody(v *QueryDeviceIpByCodeResponseBody) *QueryDeviceIpByCodeResponse {
	s.Body = v
	return s
}

type QueryDevicePropertiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDevicePropertiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePropertiesHeaders) GoString() string {
	return s.String()
}

func (s *QueryDevicePropertiesHeaders) SetCommonHeaders(v map[string]*string) *QueryDevicePropertiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDevicePropertiesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDevicePropertiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDevicePropertiesRequest struct {
	PropertyNames   []*string `json:"propertyNames,omitempty" xml:"propertyNames,omitempty" type:"Repeated"`
	DeviceId        *string   `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceUnionId   *string   `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	OperatorUnionId *string   `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s QueryDevicePropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePropertiesRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicePropertiesRequest) SetPropertyNames(v []*string) *QueryDevicePropertiesRequest {
	s.PropertyNames = v
	return s
}

func (s *QueryDevicePropertiesRequest) SetDeviceId(v string) *QueryDevicePropertiesRequest {
	s.DeviceId = &v
	return s
}

func (s *QueryDevicePropertiesRequest) SetDeviceUnionId(v string) *QueryDevicePropertiesRequest {
	s.DeviceUnionId = &v
	return s
}

func (s *QueryDevicePropertiesRequest) SetOperatorUnionId(v string) *QueryDevicePropertiesRequest {
	s.OperatorUnionId = &v
	return s
}

type QueryDevicePropertiesResponseBody struct {
	Result []*QueryDevicePropertiesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryDevicePropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicePropertiesResponseBody) SetResult(v []*QueryDevicePropertiesResponseBodyResult) *QueryDevicePropertiesResponseBody {
	s.Result = v
	return s
}

type QueryDevicePropertiesResponseBodyResult struct {
	PropertyName  *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s QueryDevicePropertiesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePropertiesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDevicePropertiesResponseBodyResult) SetPropertyName(v string) *QueryDevicePropertiesResponseBodyResult {
	s.PropertyName = &v
	return s
}

func (s *QueryDevicePropertiesResponseBodyResult) SetPropertyValue(v string) *QueryDevicePropertiesResponseBodyResult {
	s.PropertyValue = &v
	return s
}

type QueryDevicePropertiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDevicePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicePropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicePropertiesResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicePropertiesResponse) SetHeaders(v map[string]*string) *QueryDevicePropertiesResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicePropertiesResponse) SetStatusCode(v int32) *QueryDevicePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicePropertiesResponse) SetBody(v *QueryDevicePropertiesResponseBody) *QueryDevicePropertiesResponse {
	s.Body = v
	return s
}

type QueryMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomRequest) SetUnionId(v string) *QueryMeetingRoomRequest {
	s.UnionId = &v
	return s
}

type QueryMeetingRoomResponseBody struct {
	Result *QueryMeetingRoomResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBody) SetResult(v *QueryMeetingRoomResponseBodyResult) *QueryMeetingRoomResponseBody {
	s.Result = v
	return s
}

type QueryMeetingRoomResponseBodyResult struct {
	CorpId                 *string                                                 `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceUnionIds         []*string                                               `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	EnableCycleReservation *bool                                                   `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	IsvRoomId              *string                                                 `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	ReservationAuthority   *QueryMeetingRoomResponseBodyResultReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	RoomCapacity           *int32                                                  `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomGroup              *QueryMeetingRoomResponseBodyResultRoomGroup            `json:"roomGroup,omitempty" xml:"roomGroup,omitempty" type:"Struct"`
	RoomId                 *string                                                 `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabels             []*QueryMeetingRoomResponseBodyResultRoomLabels         `json:"roomLabels,omitempty" xml:"roomLabels,omitempty" type:"Repeated"`
	RoomLocation           *QueryMeetingRoomResponseBodyResultRoomLocation         `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	RoomName               *string                                                 `json:"roomName,omitempty" xml:"roomName,omitempty"`
	RoomPicture            *string                                                 `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	RoomStaffId            *string                                                 `json:"roomStaffId,omitempty" xml:"roomStaffId,omitempty"`
	RoomStatus             *int32                                                  `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResult) SetCorpId(v string) *QueryMeetingRoomResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetDeviceUnionIds(v []*string) *QueryMeetingRoomResponseBodyResult {
	s.DeviceUnionIds = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetEnableCycleReservation(v bool) *QueryMeetingRoomResponseBodyResult {
	s.EnableCycleReservation = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetIsvRoomId(v string) *QueryMeetingRoomResponseBodyResult {
	s.IsvRoomId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetReservationAuthority(v *QueryMeetingRoomResponseBodyResultReservationAuthority) *QueryMeetingRoomResponseBodyResult {
	s.ReservationAuthority = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomCapacity(v int32) *QueryMeetingRoomResponseBodyResult {
	s.RoomCapacity = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomGroup(v *QueryMeetingRoomResponseBodyResultRoomGroup) *QueryMeetingRoomResponseBodyResult {
	s.RoomGroup = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomLabels(v []*QueryMeetingRoomResponseBodyResultRoomLabels) *QueryMeetingRoomResponseBodyResult {
	s.RoomLabels = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomLocation(v *QueryMeetingRoomResponseBodyResultRoomLocation) *QueryMeetingRoomResponseBodyResult {
	s.RoomLocation = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomName(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomPicture(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomPicture = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomStaffId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomStaffId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResult) SetRoomStatus(v int32) *QueryMeetingRoomResponseBodyResult {
	s.RoomStatus = &v
	return s
}

type QueryMeetingRoomResponseBodyResultReservationAuthority struct {
	AuthorizedMembers []*QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers `json:"authorizedMembers,omitempty" xml:"authorizedMembers,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthority) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthority) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthority) SetAuthorizedMembers(v []*QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) *QueryMeetingRoomResponseBodyResultReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

type QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberId(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberName(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers) SetMemberType(v string) *QueryMeetingRoomResponseBodyResultReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

type QueryMeetingRoomResponseBodyResultRoomGroup struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ParentId  *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomGroup) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetGroupId(v int64) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetGroupName(v string) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomGroup) SetParentId(v int64) *QueryMeetingRoomResponseBodyResultRoomGroup {
	s.ParentId = &v
	return s
}

type QueryMeetingRoomResponseBodyResultRoomLabels struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomLabels) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomLabels) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) SetLabelId(v int64) *QueryMeetingRoomResponseBodyResultRoomLabels {
	s.LabelId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLabels) SetLabelName(v string) *QueryMeetingRoomResponseBodyResultRoomLabels {
	s.LabelName = &v
	return s
}

type QueryMeetingRoomResponseBodyResultRoomLocation struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultRoomLocation) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultRoomLocation) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) SetDesc(v string) *QueryMeetingRoomResponseBodyResultRoomLocation {
	s.Desc = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultRoomLocation) SetTitle(v string) *QueryMeetingRoomResponseBodyResultRoomLocation {
	s.Title = &v
	return s
}

type QueryMeetingRoomResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomResponse) SetStatusCode(v int32) *QueryMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomResponse) SetBody(v *QueryMeetingRoomResponseBody) *QueryMeetingRoomResponse {
	s.Body = v
	return s
}

type QueryMeetingRoomDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomDeviceRequest struct {
	DeviceId        *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceUnionId   *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
}

func (s QueryMeetingRoomDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceRequest) SetDeviceId(v string) *QueryMeetingRoomDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *QueryMeetingRoomDeviceRequest) SetDeviceUnionId(v string) *QueryMeetingRoomDeviceRequest {
	s.DeviceUnionId = &v
	return s
}

func (s *QueryMeetingRoomDeviceRequest) SetOperatorUnionId(v string) *QueryMeetingRoomDeviceRequest {
	s.OperatorUnionId = &v
	return s
}

type QueryMeetingRoomDeviceResponseBody struct {
	Result *QueryMeetingRoomDeviceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceResponseBody) SetResult(v *QueryMeetingRoomDeviceResponseBodyResult) *QueryMeetingRoomDeviceResponseBody {
	s.Result = v
	return s
}

type QueryMeetingRoomDeviceResponseBodyResult struct {
	ActiveTime      *int64                                                 `json:"activeTime,omitempty" xml:"activeTime,omitempty"`
	Controllers     []*QueryMeetingRoomDeviceResponseBodyResultControllers `json:"controllers,omitempty" xml:"controllers,omitempty" type:"Repeated"`
	CorpId          *string                                                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreatorUnionId  *string                                                `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	DevCamera       *string                                                `json:"devCamera,omitempty" xml:"devCamera,omitempty"`
	DevHdmi         *string                                                `json:"devHdmi,omitempty" xml:"devHdmi,omitempty"`
	DevMic          *string                                                `json:"devMic,omitempty" xml:"devMic,omitempty"`
	DevMirror       *string                                                `json:"devMirror,omitempty" xml:"devMirror,omitempty"`
	DevNetIp        *string                                                `json:"devNetIp,omitempty" xml:"devNetIp,omitempty"`
	DevNetType      *string                                                `json:"devNetType,omitempty" xml:"devNetType,omitempty"`
	DevVoice        *string                                                `json:"devVoice,omitempty" xml:"devVoice,omitempty"`
	DevWifiMac      *string                                                `json:"devWifiMac,omitempty" xml:"devWifiMac,omitempty"`
	DevWireMac      *string                                                `json:"devWireMac,omitempty" xml:"devWireMac,omitempty"`
	DeviceId        *string                                                `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceMac       *string                                                `json:"deviceMac,omitempty" xml:"deviceMac,omitempty"`
	DeviceModel     *string                                                `json:"deviceModel,omitempty" xml:"deviceModel,omitempty"`
	DeviceName      *string                                                `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceServiceId *int32                                                 `json:"deviceServiceId,omitempty" xml:"deviceServiceId,omitempty"`
	DeviceSn        *string                                                `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	DeviceStatus    *string                                                `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	DeviceType      *string                                                `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	DeviceUnionId   *string                                                `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	FirmwareVersion *string                                                `json:"firmwareVersion,omitempty" xml:"firmwareVersion,omitempty"`
	OpenRoomId      *string                                                `json:"openRoomId,omitempty" xml:"openRoomId,omitempty"`
	RoomName        *string                                                `json:"roomName,omitempty" xml:"roomName,omitempty"`
	ShareCode       *string                                                `json:"shareCode,omitempty" xml:"shareCode,omitempty"`
	SipAccountName  *string                                                `json:"sipAccountName,omitempty" xml:"sipAccountName,omitempty"`
	SoftwareVersion *string                                                `json:"softwareVersion,omitempty" xml:"softwareVersion,omitempty"`
}

func (s QueryMeetingRoomDeviceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetActiveTime(v int64) *QueryMeetingRoomDeviceResponseBodyResult {
	s.ActiveTime = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetControllers(v []*QueryMeetingRoomDeviceResponseBodyResultControllers) *QueryMeetingRoomDeviceResponseBodyResult {
	s.Controllers = v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetCorpId(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetCreatorUnionId(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.CreatorUnionId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevCamera(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevCamera = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevHdmi(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevHdmi = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevMic(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevMic = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevMirror(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevMirror = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevNetIp(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevNetIp = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevNetType(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevNetType = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevVoice(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevVoice = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevWifiMac(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevWifiMac = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDevWireMac(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DevWireMac = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceId(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceMac(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceMac = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceModel(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceModel = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceName(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceServiceId(v int32) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceServiceId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceSn(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceSn = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceStatus(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceStatus = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceType(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceType = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetDeviceUnionId(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.DeviceUnionId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetFirmwareVersion(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetOpenRoomId(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.OpenRoomId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetRoomName(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.RoomName = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetShareCode(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.ShareCode = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetSipAccountName(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.SipAccountName = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResult) SetSoftwareVersion(v string) *QueryMeetingRoomDeviceResponseBodyResult {
	s.SoftwareVersion = &v
	return s
}

type QueryMeetingRoomDeviceResponseBodyResultControllers struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceId        *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceMac       *string `json:"deviceMac,omitempty" xml:"deviceMac,omitempty"`
	DeviceModel     *string `json:"deviceModel,omitempty" xml:"deviceModel,omitempty"`
	DeviceName      *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceServiceId *int32  `json:"deviceServiceId,omitempty" xml:"deviceServiceId,omitempty"`
	DeviceSn        *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	DeviceStatus    *string `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	DeviceType      *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	DeviceUnionId   *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	OpenRoomId      *string `json:"openRoomId,omitempty" xml:"openRoomId,omitempty"`
	ShareCode       *string `json:"shareCode,omitempty" xml:"shareCode,omitempty"`
}

func (s QueryMeetingRoomDeviceResponseBodyResultControllers) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceResponseBodyResultControllers) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetCorpId(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceId(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceMac(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceMac = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceModel(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceModel = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceName(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceName = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceServiceId(v int32) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceServiceId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceSn(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceSn = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceStatus(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceStatus = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceType(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceType = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetDeviceUnionId(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.DeviceUnionId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetOpenRoomId(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.OpenRoomId = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponseBodyResultControllers) SetShareCode(v string) *QueryMeetingRoomDeviceResponseBodyResultControllers {
	s.ShareCode = &v
	return s
}

type QueryMeetingRoomDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingRoomDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomDeviceResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomDeviceResponse) SetStatusCode(v int32) *QueryMeetingRoomDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomDeviceResponse) SetBody(v *QueryMeetingRoomDeviceResponseBody) *QueryMeetingRoomDeviceResponse {
	s.Body = v
	return s
}

type QueryMeetingRoomGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomGroupRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMeetingRoomGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupRequest) SetUnionId(v string) *QueryMeetingRoomGroupRequest {
	s.UnionId = &v
	return s
}

type QueryMeetingRoomGroupResponseBody struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ParentId  *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s QueryMeetingRoomGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupResponseBody) SetGroupId(v int64) *QueryMeetingRoomGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) SetGroupName(v string) *QueryMeetingRoomGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomGroupResponseBody) SetParentId(v int64) *QueryMeetingRoomGroupResponseBody {
	s.ParentId = &v
	return s
}

type QueryMeetingRoomGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingRoomGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomGroupResponse) SetStatusCode(v int32) *QueryMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomGroupResponse) SetBody(v *QueryMeetingRoomGroupResponseBody) *QueryMeetingRoomGroupResponse {
	s.Body = v
	return s
}

type QueryMeetingRoomGroupListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomGroupListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomGroupListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomGroupListRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMeetingRoomGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupListRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListRequest) SetUnionId(v string) *QueryMeetingRoomGroupListRequest {
	s.UnionId = &v
	return s
}

type QueryMeetingRoomGroupListResponseBody struct {
	Result []*QueryMeetingRoomGroupListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponseBody) SetResult(v []*QueryMeetingRoomGroupListResponseBodyResult) *QueryMeetingRoomGroupListResponseBody {
	s.Result = v
	return s
}

type QueryMeetingRoomGroupListResponseBodyResult struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ParentId  *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s QueryMeetingRoomGroupListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetGroupId(v int64) *QueryMeetingRoomGroupListResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetGroupName(v string) *QueryMeetingRoomGroupListResponseBodyResult {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponseBodyResult) SetParentId(v int64) *QueryMeetingRoomGroupListResponseBodyResult {
	s.ParentId = &v
	return s
}

type QueryMeetingRoomGroupListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingRoomGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingRoomGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomGroupListResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomGroupListResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomGroupListResponse) SetStatusCode(v int32) *QueryMeetingRoomGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomGroupListResponse) SetBody(v *QueryMeetingRoomGroupListResponseBody) *QueryMeetingRoomGroupListResponse {
	s.Body = v
	return s
}

type QueryMeetingRoomListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomListRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMeetingRoomListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListRequest) SetMaxResults(v int32) *QueryMeetingRoomListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMeetingRoomListRequest) SetNextToken(v int64) *QueryMeetingRoomListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomListRequest) SetUnionId(v string) *QueryMeetingRoomListRequest {
	s.UnionId = &v
	return s
}

type QueryMeetingRoomListResponseBody struct {
	HasMore   *bool                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*QueryMeetingRoomListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBody) SetHasMore(v bool) *QueryMeetingRoomListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryMeetingRoomListResponseBody) SetNextToken(v int64) *QueryMeetingRoomListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomListResponseBody) SetResult(v []*QueryMeetingRoomListResponseBodyResult) *QueryMeetingRoomListResponseBody {
	s.Result = v
	return s
}

type QueryMeetingRoomListResponseBodyResult struct {
	CorpId       *string                                             `json:"corpId,omitempty" xml:"corpId,omitempty"`
	IsvRoomId    *string                                             `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	RoomCapacity *int32                                              `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomGroup    *QueryMeetingRoomListResponseBodyResultRoomGroup    `json:"roomGroup,omitempty" xml:"roomGroup,omitempty" type:"Struct"`
	RoomId       *string                                             `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabels   []*QueryMeetingRoomListResponseBodyResultRoomLabels `json:"roomLabels,omitempty" xml:"roomLabels,omitempty" type:"Repeated"`
	RoomLocation *QueryMeetingRoomListResponseBodyResultRoomLocation `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	RoomName     *string                                             `json:"roomName,omitempty" xml:"roomName,omitempty"`
	RoomPicture  *string                                             `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	RoomStaffId  *string                                             `json:"roomStaffId,omitempty" xml:"roomStaffId,omitempty"`
	RoomStatus   *int32                                              `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResult) SetCorpId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetIsvRoomId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.IsvRoomId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomCapacity(v int32) *QueryMeetingRoomListResponseBodyResult {
	s.RoomCapacity = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomGroup(v *QueryMeetingRoomListResponseBodyResultRoomGroup) *QueryMeetingRoomListResponseBodyResult {
	s.RoomGroup = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomLabels(v []*QueryMeetingRoomListResponseBodyResultRoomLabels) *QueryMeetingRoomListResponseBodyResult {
	s.RoomLabels = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomLocation(v *QueryMeetingRoomListResponseBodyResultRoomLocation) *QueryMeetingRoomListResponseBodyResult {
	s.RoomLocation = v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomName(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomName = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomPicture(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomPicture = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomStaffId(v string) *QueryMeetingRoomListResponseBodyResult {
	s.RoomStaffId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResult) SetRoomStatus(v int32) *QueryMeetingRoomListResponseBodyResult {
	s.RoomStatus = &v
	return s
}

type QueryMeetingRoomListResponseBodyResultRoomGroup struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ParentId  *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResultRoomGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomGroup) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetGroupId(v int64) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.GroupId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetGroupName(v string) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.GroupName = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomGroup) SetParentId(v int64) *QueryMeetingRoomListResponseBodyResultRoomGroup {
	s.ParentId = &v
	return s
}

type QueryMeetingRoomListResponseBodyResultRoomLabels struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResultRoomLabels) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomLabels) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) SetLabelId(v int64) *QueryMeetingRoomListResponseBodyResultRoomLabels {
	s.LabelId = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLabels) SetLabelName(v string) *QueryMeetingRoomListResponseBodyResultRoomLabels {
	s.LabelName = &v
	return s
}

type QueryMeetingRoomListResponseBodyResultRoomLocation struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryMeetingRoomListResponseBodyResultRoomLocation) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponseBodyResultRoomLocation) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) SetDesc(v string) *QueryMeetingRoomListResponseBodyResultRoomLocation {
	s.Desc = &v
	return s
}

func (s *QueryMeetingRoomListResponseBodyResultRoomLocation) SetTitle(v string) *QueryMeetingRoomListResponseBodyResultRoomLocation {
	s.Title = &v
	return s
}

type QueryMeetingRoomListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMeetingRoomListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMeetingRoomListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomListResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomListResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomListResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomListResponse) SetStatusCode(v int32) *QueryMeetingRoomListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomListResponse) SetBody(v *QueryMeetingRoomListResponseBody) *QueryMeetingRoomListResponse {
	s.Body = v
	return s
}

type RemoveSuperUserMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveSuperUserMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveSuperUserMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *RemoveSuperUserMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *RemoveSuperUserMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveSuperUserMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveSuperUserMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveSuperUserMeetingRoomRequest struct {
	RoomId  *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s RemoveSuperUserMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSuperUserMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *RemoveSuperUserMeetingRoomRequest) SetRoomId(v string) *RemoveSuperUserMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *RemoveSuperUserMeetingRoomRequest) SetUnionId(v string) *RemoveSuperUserMeetingRoomRequest {
	s.UnionId = &v
	return s
}

type RemoveSuperUserMeetingRoomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveSuperUserMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSuperUserMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSuperUserMeetingRoomResponseBody) SetResult(v bool) *RemoveSuperUserMeetingRoomResponseBody {
	s.Result = &v
	return s
}

type RemoveSuperUserMeetingRoomResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSuperUserMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSuperUserMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSuperUserMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *RemoveSuperUserMeetingRoomResponse) SetHeaders(v map[string]*string) *RemoveSuperUserMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *RemoveSuperUserMeetingRoomResponse) SetStatusCode(v int32) *RemoveSuperUserMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSuperUserMeetingRoomResponse) SetBody(v *RemoveSuperUserMeetingRoomResponseBody) *RemoveSuperUserMeetingRoomResponse {
	s.Body = v
	return s
}

type SetSuperUserMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetSuperUserMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetSuperUserMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *SetSuperUserMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *SetSuperUserMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetSuperUserMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *SetSuperUserMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetSuperUserMeetingRoomRequest struct {
	DeptIdWhiteList []*int64  `json:"deptIdWhiteList,omitempty" xml:"deptIdWhiteList,omitempty" type:"Repeated"`
	RoomId          *string   `json:"roomId,omitempty" xml:"roomId,omitempty"`
	UnionId         *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserIdWhiteList []*string `json:"userIdWhiteList,omitempty" xml:"userIdWhiteList,omitempty" type:"Repeated"`
}

func (s SetSuperUserMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSuperUserMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *SetSuperUserMeetingRoomRequest) SetDeptIdWhiteList(v []*int64) *SetSuperUserMeetingRoomRequest {
	s.DeptIdWhiteList = v
	return s
}

func (s *SetSuperUserMeetingRoomRequest) SetRoomId(v string) *SetSuperUserMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *SetSuperUserMeetingRoomRequest) SetUnionId(v string) *SetSuperUserMeetingRoomRequest {
	s.UnionId = &v
	return s
}

func (s *SetSuperUserMeetingRoomRequest) SetUserIdWhiteList(v []*string) *SetSuperUserMeetingRoomRequest {
	s.UserIdWhiteList = v
	return s
}

type SetSuperUserMeetingRoomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetSuperUserMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSuperUserMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *SetSuperUserMeetingRoomResponseBody) SetResult(v bool) *SetSuperUserMeetingRoomResponseBody {
	s.Result = &v
	return s
}

type SetSuperUserMeetingRoomResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetSuperUserMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSuperUserMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSuperUserMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *SetSuperUserMeetingRoomResponse) SetHeaders(v map[string]*string) *SetSuperUserMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *SetSuperUserMeetingRoomResponse) SetStatusCode(v int32) *SetSuperUserMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSuperUserMeetingRoomResponse) SetBody(v *SetSuperUserMeetingRoomResponseBody) *SetSuperUserMeetingRoomResponse {
	s.Body = v
	return s
}

type UpdateDeviceCustomTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateDeviceCustomTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCustomTemplateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCustomTemplateHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeviceCustomTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeviceCustomTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateDeviceCustomTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateDeviceCustomTemplateRequest struct {
	BgImgList                    []*string `json:"bgImgList,omitempty" xml:"bgImgList,omitempty" type:"Repeated"`
	BgType                       *int32    `json:"bgType,omitempty" xml:"bgType,omitempty"`
	BgUrl                        *string   `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	CustomDoc                    *string   `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	DesensitizeUserName          *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	DeviceUnionIds               []*string `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	GroupIds                     []*int64  `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	HideServerCodeWhenProjecting *bool     `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	Instruction                  *bool     `json:"instruction,omitempty" xml:"instruction,omitempty"`
	IsPicTop                     *int32    `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	Logo                         *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	OrgName                      *string   `json:"orgName,omitempty" xml:"orgName,omitempty"`
	PicturePlayInterval          *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	RoomIds                      []*string `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
	ShowCalendarCard             *bool     `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	ShowCalendarTitle            *bool     `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	ShowFunctionCard             *bool     `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	TemplateId                   *int64    `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName                 *string   `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s UpdateDeviceCustomTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCustomTemplateRequest) SetBgImgList(v []*string) *UpdateDeviceCustomTemplateRequest {
	s.BgImgList = v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetBgType(v int32) *UpdateDeviceCustomTemplateRequest {
	s.BgType = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetBgUrl(v string) *UpdateDeviceCustomTemplateRequest {
	s.BgUrl = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetCustomDoc(v string) *UpdateDeviceCustomTemplateRequest {
	s.CustomDoc = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetDesensitizeUserName(v bool) *UpdateDeviceCustomTemplateRequest {
	s.DesensitizeUserName = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetDeviceUnionIds(v []*string) *UpdateDeviceCustomTemplateRequest {
	s.DeviceUnionIds = v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetGroupIds(v []*int64) *UpdateDeviceCustomTemplateRequest {
	s.GroupIds = v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetHideServerCodeWhenProjecting(v bool) *UpdateDeviceCustomTemplateRequest {
	s.HideServerCodeWhenProjecting = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetInstruction(v bool) *UpdateDeviceCustomTemplateRequest {
	s.Instruction = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetIsPicTop(v int32) *UpdateDeviceCustomTemplateRequest {
	s.IsPicTop = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetLogo(v string) *UpdateDeviceCustomTemplateRequest {
	s.Logo = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetOrgName(v string) *UpdateDeviceCustomTemplateRequest {
	s.OrgName = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetPicturePlayInterval(v int32) *UpdateDeviceCustomTemplateRequest {
	s.PicturePlayInterval = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetRoomIds(v []*string) *UpdateDeviceCustomTemplateRequest {
	s.RoomIds = v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetShowCalendarCard(v bool) *UpdateDeviceCustomTemplateRequest {
	s.ShowCalendarCard = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetShowCalendarTitle(v bool) *UpdateDeviceCustomTemplateRequest {
	s.ShowCalendarTitle = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetShowFunctionCard(v bool) *UpdateDeviceCustomTemplateRequest {
	s.ShowFunctionCard = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetTemplateId(v int64) *UpdateDeviceCustomTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDeviceCustomTemplateRequest) SetTemplateName(v string) *UpdateDeviceCustomTemplateRequest {
	s.TemplateName = &v
	return s
}

type UpdateDeviceCustomTemplateResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateDeviceCustomTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCustomTemplateResponseBody) SetResult(v bool) *UpdateDeviceCustomTemplateResponseBody {
	s.Result = &v
	return s
}

type UpdateDeviceCustomTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceCustomTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceCustomTemplateResponse) SetHeaders(v map[string]*string) *UpdateDeviceCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceCustomTemplateResponse) SetStatusCode(v int32) *UpdateDeviceCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeviceCustomTemplateResponse) SetBody(v *UpdateDeviceCustomTemplateResponseBody) *UpdateDeviceCustomTemplateResponse {
	s.Body = v
	return s
}

type UpdateMeetingRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMeetingRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMeetingRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMeetingRoomRequest struct {
	EnableCycleReservation *bool                                         `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	GroupId                *int64                                        `json:"groupId,omitempty" xml:"groupId,omitempty"`
	IsvRoomId              *string                                       `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	ReservationAuthority   *UpdateMeetingRoomRequestReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	RoomCapacity           *int32                                        `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomId                 *string                                       `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabelIds           []*int64                                      `json:"roomLabelIds,omitempty" xml:"roomLabelIds,omitempty" type:"Repeated"`
	RoomLocation           *UpdateMeetingRoomRequestRoomLocation         `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	RoomName               *string                                       `json:"roomName,omitempty" xml:"roomName,omitempty"`
	RoomPicture            *string                                       `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	RoomStatus             *int32                                        `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	UnionId                *string                                       `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateMeetingRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequest) SetEnableCycleReservation(v bool) *UpdateMeetingRoomRequest {
	s.EnableCycleReservation = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetGroupId(v int64) *UpdateMeetingRoomRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetIsvRoomId(v string) *UpdateMeetingRoomRequest {
	s.IsvRoomId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetReservationAuthority(v *UpdateMeetingRoomRequestReservationAuthority) *UpdateMeetingRoomRequest {
	s.ReservationAuthority = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomCapacity(v int32) *UpdateMeetingRoomRequest {
	s.RoomCapacity = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomId(v string) *UpdateMeetingRoomRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomLabelIds(v []*int64) *UpdateMeetingRoomRequest {
	s.RoomLabelIds = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomLocation(v *UpdateMeetingRoomRequestRoomLocation) *UpdateMeetingRoomRequest {
	s.RoomLocation = v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomName(v string) *UpdateMeetingRoomRequest {
	s.RoomName = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomPicture(v string) *UpdateMeetingRoomRequest {
	s.RoomPicture = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetRoomStatus(v int32) *UpdateMeetingRoomRequest {
	s.RoomStatus = &v
	return s
}

func (s *UpdateMeetingRoomRequest) SetUnionId(v string) *UpdateMeetingRoomRequest {
	s.UnionId = &v
	return s
}

type UpdateMeetingRoomRequestReservationAuthority struct {
	AuthorizedMembers []*UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers `json:"authorizedMembers,omitempty" xml:"authorizedMembers,omitempty" type:"Repeated"`
}

func (s UpdateMeetingRoomRequestReservationAuthority) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomRequestReservationAuthority) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestReservationAuthority) SetAuthorizedMembers(v []*UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) *UpdateMeetingRoomRequestReservationAuthority {
	s.AuthorizedMembers = v
	return s
}

type UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers struct {
	MemberId   *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
}

func (s UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberId(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberName(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberName = &v
	return s
}

func (s *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers) SetMemberType(v string) *UpdateMeetingRoomRequestReservationAuthorityAuthorizedMembers {
	s.MemberType = &v
	return s
}

type UpdateMeetingRoomRequestRoomLocation struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateMeetingRoomRequestRoomLocation) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomRequestRoomLocation) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomRequestRoomLocation) SetDesc(v string) *UpdateMeetingRoomRequestRoomLocation {
	s.Desc = &v
	return s
}

func (s *UpdateMeetingRoomRequestRoomLocation) SetTitle(v string) *UpdateMeetingRoomRequestRoomLocation {
	s.Title = &v
	return s
}

type UpdateMeetingRoomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateMeetingRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomResponseBody) SetResult(v bool) *UpdateMeetingRoomResponseBody {
	s.Result = &v
	return s
}

type UpdateMeetingRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMeetingRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomResponse) SetHeaders(v map[string]*string) *UpdateMeetingRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeetingRoomResponse) SetStatusCode(v int32) *UpdateMeetingRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeetingRoomResponse) SetBody(v *UpdateMeetingRoomResponseBody) *UpdateMeetingRoomResponse {
	s.Body = v
	return s
}

type UpdateMeetingRoomGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMeetingRoomGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMeetingRoomGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMeetingRoomGroupRequest struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateMeetingRoomGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupRequest) SetGroupId(v int64) *UpdateMeetingRoomGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMeetingRoomGroupRequest) SetGroupName(v string) *UpdateMeetingRoomGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateMeetingRoomGroupRequest) SetUnionId(v string) *UpdateMeetingRoomGroupRequest {
	s.UnionId = &v
	return s
}

type UpdateMeetingRoomGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateMeetingRoomGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupResponseBody) SetResult(v bool) *UpdateMeetingRoomGroupResponseBody {
	s.Result = &v
	return s
}

type UpdateMeetingRoomGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMeetingRoomGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeetingRoomGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupResponse) SetHeaders(v map[string]*string) *UpdateMeetingRoomGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeetingRoomGroupResponse) SetStatusCode(v int32) *UpdateMeetingRoomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeetingRoomGroupResponse) SetBody(v *UpdateMeetingRoomGroupResponseBody) *UpdateMeetingRoomGroupResponse {
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

func (client *Client) CreateDeviceCustomTemplateWithOptions(request *CreateDeviceCustomTemplateRequest, headers *CreateDeviceCustomTemplateHeaders, runtime *util.RuntimeOptions) (_result *CreateDeviceCustomTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BgImgList)) {
		body["bgImgList"] = request.BgImgList
	}

	if !tea.BoolValue(util.IsUnset(request.BgType)) {
		body["bgType"] = request.BgType
	}

	if !tea.BoolValue(util.IsUnset(request.BgUrl)) {
		body["bgUrl"] = request.BgUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CustomDoc)) {
		body["customDoc"] = request.CustomDoc
	}

	if !tea.BoolValue(util.IsUnset(request.DesensitizeUserName)) {
		body["desensitizeUserName"] = request.DesensitizeUserName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUnionIds)) {
		body["deviceUnionIds"] = request.DeviceUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideServerCodeWhenProjecting)) {
		body["hideServerCodeWhenProjecting"] = request.HideServerCodeWhenProjecting
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.IsPicTop)) {
		body["isPicTop"] = request.IsPicTop
	}

	if !tea.BoolValue(util.IsUnset(request.Logo)) {
		body["logo"] = request.Logo
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.PicturePlayInterval)) {
		body["picturePlayInterval"] = request.PicturePlayInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RoomIds)) {
		body["roomIds"] = request.RoomIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCalendarCard)) {
		body["showCalendarCard"] = request.ShowCalendarCard
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCalendarTitle)) {
		body["showCalendarTitle"] = request.ShowCalendarTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ShowFunctionCard)) {
		body["showFunctionCard"] = request.ShowFunctionCard
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["templateName"] = request.TemplateName
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
		Action:      tea.String("CreateDeviceCustomTemplate"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/screens/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceCustomTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceCustomTemplate(request *CreateDeviceCustomTemplateRequest) (_result *CreateDeviceCustomTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeviceCustomTemplateHeaders{}
	_result = &CreateDeviceCustomTemplateResponse{}
	_body, _err := client.CreateDeviceCustomTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMeetingRoomWithOptions(request *CreateMeetingRoomRequest, headers *CreateMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *CreateMeetingRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableCycleReservation)) {
		body["enableCycleReservation"] = request.EnableCycleReservation
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvRoomId)) {
		body["isvRoomId"] = request.IsvRoomId
	}

	if !tea.BoolValue(util.IsUnset(request.ReservationAuthority)) {
		body["reservationAuthority"] = request.ReservationAuthority
	}

	if !tea.BoolValue(util.IsUnset(request.RoomCapacity)) {
		body["roomCapacity"] = request.RoomCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.RoomLabelIds)) {
		body["roomLabelIds"] = request.RoomLabelIds
	}

	if !tea.BoolValue(util.IsUnset(request.RoomLocation)) {
		body["roomLocation"] = request.RoomLocation
	}

	if !tea.BoolValue(util.IsUnset(request.RoomName)) {
		body["roomName"] = request.RoomName
	}

	if !tea.BoolValue(util.IsUnset(request.RoomPicture)) {
		body["roomPicture"] = request.RoomPicture
	}

	if !tea.BoolValue(util.IsUnset(request.RoomStatus)) {
		body["roomStatus"] = request.RoomStatus
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
		Action:      tea.String("CreateMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingrooms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMeetingRoom(request *CreateMeetingRoomRequest) (_result *CreateMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMeetingRoomHeaders{}
	_result = &CreateMeetingRoomResponse{}
	_body, _err := client.CreateMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMeetingRoomGroupWithOptions(request *CreateMeetingRoomGroupRequest, headers *CreateMeetingRoomGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateMeetingRoomGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentGroupId)) {
		body["parentGroupId"] = request.ParentGroupId
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
		Action:      tea.String("CreateMeetingRoomGroup"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingRoomGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMeetingRoomGroup(request *CreateMeetingRoomGroupRequest) (_result *CreateMeetingRoomGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMeetingRoomGroupHeaders{}
	_result = &CreateMeetingRoomGroupResponse{}
	_body, _err := client.CreateMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceCustomTemplateWithOptions(request *DeleteDeviceCustomTemplateRequest, headers *DeleteDeviceCustomTemplateHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeviceCustomTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("DeleteDeviceCustomTemplate"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/screens/templates/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceCustomTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceCustomTemplate(request *DeleteDeviceCustomTemplateRequest) (_result *DeleteDeviceCustomTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeviceCustomTemplateHeaders{}
	_result = &DeleteDeviceCustomTemplateResponse{}
	_body, _err := client.DeleteDeviceCustomTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMeetingRoomWithOptions(roomId *string, request *DeleteMeetingRoomRequest, headers *DeleteMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *DeleteMeetingRoomResponse, _err error) {
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
		Action:      tea.String("DeleteMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRooms/" + tea.StringValue(roomId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMeetingRoom(roomId *string, request *DeleteMeetingRoomRequest) (_result *DeleteMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMeetingRoomHeaders{}
	_result = &DeleteMeetingRoomResponse{}
	_body, _err := client.DeleteMeetingRoomWithOptions(roomId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMeetingRoomGroupWithOptions(groupId *string, request *DeleteMeetingRoomGroupRequest, headers *DeleteMeetingRoomGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteMeetingRoomGroupResponse, _err error) {
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
		Action:      tea.String("DeleteMeetingRoomGroup"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMeetingRoomGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMeetingRoomGroup(groupId *string, request *DeleteMeetingRoomGroupRequest) (_result *DeleteMeetingRoomGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMeetingRoomGroupHeaders{}
	_result = &DeleteMeetingRoomGroupResponse{}
	_body, _err := client.DeleteMeetingRoomGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceCustomTemplateWithOptions(templateId *string, headers *QueryDeviceCustomTemplateHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceCustomTemplateResponse, _err error) {
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
		Action:      tea.String("QueryDeviceCustomTemplate"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/screens/templates/" + tea.StringValue(templateId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceCustomTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceCustomTemplate(templateId *string) (_result *QueryDeviceCustomTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceCustomTemplateHeaders{}
	_result = &QueryDeviceCustomTemplateResponse{}
	_body, _err := client.QueryDeviceCustomTemplateWithOptions(templateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceCustomTemplateListWithOptions(headers *QueryDeviceCustomTemplateListHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceCustomTemplateListResponse, _err error) {
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
		Action:      tea.String("QueryDeviceCustomTemplateList"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/screens/templateLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceCustomTemplateListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceCustomTemplateList() (_result *QueryDeviceCustomTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceCustomTemplateListHeaders{}
	_result = &QueryDeviceCustomTemplateListResponse{}
	_body, _err := client.QueryDeviceCustomTemplateListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceIpByCodeWithOptions(shareCode *string, request *QueryDeviceIpByCodeRequest, headers *QueryDeviceIpByCodeHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceIpByCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceSn)) {
		query["deviceSn"] = request.DeviceSn
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
		Action:      tea.String("QueryDeviceIpByCode"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/shareCodes/" + tea.StringValue(shareCode)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceIpByCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceIpByCode(shareCode *string, request *QueryDeviceIpByCodeRequest) (_result *QueryDeviceIpByCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceIpByCodeHeaders{}
	_result = &QueryDeviceIpByCodeResponse{}
	_body, _err := client.QueryDeviceIpByCodeWithOptions(shareCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicePropertiesWithOptions(request *QueryDevicePropertiesRequest, headers *QueryDevicePropertiesHeaders, runtime *util.RuntimeOptions) (_result *QueryDevicePropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUnionId)) {
		query["deviceUnionId"] = request.DeviceUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		query["operatorUnionId"] = request.OperatorUnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyNames)) {
		body["propertyNames"] = request.PropertyNames
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
		Action:      tea.String("QueryDeviceProperties"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/properties/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicePropertiesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceProperties(request *QueryDevicePropertiesRequest) (_result *QueryDevicePropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDevicePropertiesHeaders{}
	_result = &QueryDevicePropertiesResponse{}
	_body, _err := client.QueryDevicePropertiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingRoomWithOptions(roomId *string, request *QueryMeetingRoomRequest, headers *QueryMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomResponse, _err error) {
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
		Action:      tea.String("QueryMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRooms/" + tea.StringValue(roomId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingRoom(roomId *string, request *QueryMeetingRoomRequest) (_result *QueryMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomHeaders{}
	_result = &QueryMeetingRoomResponse{}
	_body, _err := client.QueryMeetingRoomWithOptions(roomId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingRoomDeviceWithOptions(request *QueryMeetingRoomDeviceRequest, headers *QueryMeetingRoomDeviceHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUnionId)) {
		query["deviceUnionId"] = request.DeviceUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		query["operatorUnionId"] = request.OperatorUnionId
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
		Action:      tea.String("QueryMeetingRoomDevice"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingRoomDevice(request *QueryMeetingRoomDeviceRequest) (_result *QueryMeetingRoomDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomDeviceHeaders{}
	_result = &QueryMeetingRoomDeviceResponse{}
	_body, _err := client.QueryMeetingRoomDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingRoomGroupWithOptions(groupId *string, request *QueryMeetingRoomGroupRequest, headers *QueryMeetingRoomGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomGroupResponse, _err error) {
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
		Action:      tea.String("QueryMeetingRoomGroup"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingRoomGroup(groupId *string, request *QueryMeetingRoomGroupRequest) (_result *QueryMeetingRoomGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomGroupHeaders{}
	_result = &QueryMeetingRoomGroupResponse{}
	_body, _err := client.QueryMeetingRoomGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingRoomGroupListWithOptions(request *QueryMeetingRoomGroupListRequest, headers *QueryMeetingRoomGroupListHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomGroupListResponse, _err error) {
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
		Action:      tea.String("QueryMeetingRoomGroupList"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/groupLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomGroupListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingRoomGroupList(request *QueryMeetingRoomGroupListRequest) (_result *QueryMeetingRoomGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomGroupListHeaders{}
	_result = &QueryMeetingRoomGroupListResponse{}
	_body, _err := client.QueryMeetingRoomGroupListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMeetingRoomListWithOptions(request *QueryMeetingRoomListRequest, headers *QueryMeetingRoomListHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomListResponse, _err error) {
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
		Action:      tea.String("QueryMeetingRoomList"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRoomLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMeetingRoomList(request *QueryMeetingRoomListRequest) (_result *QueryMeetingRoomListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomListHeaders{}
	_result = &QueryMeetingRoomListResponse{}
	_body, _err := client.QueryMeetingRoomListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSuperUserMeetingRoomWithOptions(request *RemoveSuperUserMeetingRoomRequest, headers *RemoveSuperUserMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *RemoveSuperUserMeetingRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["roomId"] = request.RoomId
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
		Action:      tea.String("RemoveSuperUserMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRooms/superUsers/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSuperUserMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSuperUserMeetingRoom(request *RemoveSuperUserMeetingRoomRequest) (_result *RemoveSuperUserMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveSuperUserMeetingRoomHeaders{}
	_result = &RemoveSuperUserMeetingRoomResponse{}
	_body, _err := client.RemoveSuperUserMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSuperUserMeetingRoomWithOptions(request *SetSuperUserMeetingRoomRequest, headers *SetSuperUserMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *SetSuperUserMeetingRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdWhiteList)) {
		body["deptIdWhiteList"] = request.DeptIdWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["roomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdWhiteList)) {
		body["userIdWhiteList"] = request.UserIdWhiteList
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
		Action:      tea.String("SetSuperUserMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRooms/superUsers/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSuperUserMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSuperUserMeetingRoom(request *SetSuperUserMeetingRoomRequest) (_result *SetSuperUserMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetSuperUserMeetingRoomHeaders{}
	_result = &SetSuperUserMeetingRoomResponse{}
	_body, _err := client.SetSuperUserMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceCustomTemplateWithOptions(request *UpdateDeviceCustomTemplateRequest, headers *UpdateDeviceCustomTemplateHeaders, runtime *util.RuntimeOptions) (_result *UpdateDeviceCustomTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BgImgList)) {
		body["bgImgList"] = request.BgImgList
	}

	if !tea.BoolValue(util.IsUnset(request.BgType)) {
		body["bgType"] = request.BgType
	}

	if !tea.BoolValue(util.IsUnset(request.BgUrl)) {
		body["bgUrl"] = request.BgUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CustomDoc)) {
		body["customDoc"] = request.CustomDoc
	}

	if !tea.BoolValue(util.IsUnset(request.DesensitizeUserName)) {
		body["desensitizeUserName"] = request.DesensitizeUserName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUnionIds)) {
		body["deviceUnionIds"] = request.DeviceUnionIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.HideServerCodeWhenProjecting)) {
		body["hideServerCodeWhenProjecting"] = request.HideServerCodeWhenProjecting
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.IsPicTop)) {
		body["isPicTop"] = request.IsPicTop
	}

	if !tea.BoolValue(util.IsUnset(request.Logo)) {
		body["logo"] = request.Logo
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.PicturePlayInterval)) {
		body["picturePlayInterval"] = request.PicturePlayInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RoomIds)) {
		body["roomIds"] = request.RoomIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCalendarCard)) {
		body["showCalendarCard"] = request.ShowCalendarCard
	}

	if !tea.BoolValue(util.IsUnset(request.ShowCalendarTitle)) {
		body["showCalendarTitle"] = request.ShowCalendarTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ShowFunctionCard)) {
		body["showFunctionCard"] = request.ShowFunctionCard
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		body["templateName"] = request.TemplateName
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
		Action:      tea.String("UpdateDeviceCustomTemplate"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/devices/screens/templates"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDeviceCustomTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceCustomTemplate(request *UpdateDeviceCustomTemplateRequest) (_result *UpdateDeviceCustomTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateDeviceCustomTemplateHeaders{}
	_result = &UpdateDeviceCustomTemplateResponse{}
	_body, _err := client.UpdateDeviceCustomTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMeetingRoomWithOptions(request *UpdateMeetingRoomRequest, headers *UpdateMeetingRoomHeaders, runtime *util.RuntimeOptions) (_result *UpdateMeetingRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableCycleReservation)) {
		body["enableCycleReservation"] = request.EnableCycleReservation
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvRoomId)) {
		body["isvRoomId"] = request.IsvRoomId
	}

	if !tea.BoolValue(util.IsUnset(request.ReservationAuthority)) {
		body["reservationAuthority"] = request.ReservationAuthority
	}

	if !tea.BoolValue(util.IsUnset(request.RoomCapacity)) {
		body["roomCapacity"] = request.RoomCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["roomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomLabelIds)) {
		body["roomLabelIds"] = request.RoomLabelIds
	}

	if !tea.BoolValue(util.IsUnset(request.RoomLocation)) {
		body["roomLocation"] = request.RoomLocation
	}

	if !tea.BoolValue(util.IsUnset(request.RoomName)) {
		body["roomName"] = request.RoomName
	}

	if !tea.BoolValue(util.IsUnset(request.RoomPicture)) {
		body["roomPicture"] = request.RoomPicture
	}

	if !tea.BoolValue(util.IsUnset(request.RoomStatus)) {
		body["roomStatus"] = request.RoomStatus
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
		Action:      tea.String("UpdateMeetingRoom"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/meetingRooms"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMeetingRoomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMeetingRoom(request *UpdateMeetingRoomRequest) (_result *UpdateMeetingRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMeetingRoomHeaders{}
	_result = &UpdateMeetingRoomResponse{}
	_body, _err := client.UpdateMeetingRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMeetingRoomGroupWithOptions(request *UpdateMeetingRoomGroupRequest, headers *UpdateMeetingRoomGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateMeetingRoomGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
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
		Action:      tea.String("UpdateMeetingRoomGroup"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/groups"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMeetingRoomGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMeetingRoomGroup(request *UpdateMeetingRoomGroupRequest) (_result *UpdateMeetingRoomGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMeetingRoomGroupHeaders{}
	_result = &UpdateMeetingRoomGroupResponse{}
	_body, _err := client.UpdateMeetingRoomGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
