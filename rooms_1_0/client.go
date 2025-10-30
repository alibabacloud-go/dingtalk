// This file is auto-generated, don't edit it. Thanks.
package rooms_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateBookingBlacklistHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBookingBlacklistHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBookingBlacklistHeaders) GoString() string {
	return s.String()
}

func (s *CreateBookingBlacklistHeaders) SetCommonHeaders(v map[string]*string) *CreateBookingBlacklistHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBookingBlacklistHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBookingBlacklistHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBookingBlacklistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	BlacklistUnionId *string `json:"blacklistUnionId,omitempty" xml:"blacklistUnionId,omitempty"`
	// example:
	//
	// 1728539655110
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1728539655017
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateBookingBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBookingBlacklistRequest) GoString() string {
	return s.String()
}

func (s *CreateBookingBlacklistRequest) SetBlacklistUnionId(v string) *CreateBookingBlacklistRequest {
	s.BlacklistUnionId = &v
	return s
}

func (s *CreateBookingBlacklistRequest) SetEndTime(v int64) *CreateBookingBlacklistRequest {
	s.EndTime = &v
	return s
}

func (s *CreateBookingBlacklistRequest) SetMemo(v string) *CreateBookingBlacklistRequest {
	s.Memo = &v
	return s
}

func (s *CreateBookingBlacklistRequest) SetStartTime(v int64) *CreateBookingBlacklistRequest {
	s.StartTime = &v
	return s
}

func (s *CreateBookingBlacklistRequest) SetUnionId(v string) *CreateBookingBlacklistRequest {
	s.UnionId = &v
	return s
}

type CreateBookingBlacklistResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateBookingBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBookingBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBookingBlacklistResponseBody) SetResult(v bool) *CreateBookingBlacklistResponseBody {
	s.Result = &v
	return s
}

type CreateBookingBlacklistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBookingBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBookingBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBookingBlacklistResponse) GoString() string {
	return s.String()
}

func (s *CreateBookingBlacklistResponse) SetHeaders(v map[string]*string) *CreateBookingBlacklistResponse {
	s.Headers = v
	return s
}

func (s *CreateBookingBlacklistResponse) SetStatusCode(v int32) *CreateBookingBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBookingBlacklistResponse) SetBody(v *CreateBookingBlacklistResponseBody) *CreateBookingBlacklistResponse {
	s.Body = v
	return s
}

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
	BgImgList []*string `json:"bgImgList,omitempty" xml:"bgImgList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BgType *int32 `json:"bgType,omitempty" xml:"bgType,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01GWWCCR1y2D9D9EHej_!!6000000006520-2-tps-1920-470.png
	BgUrl *string `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	// example:
	//
	// 测试文本
	CustomDoc                    *string   `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	DesensitizeUserName          *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	DeviceUnionIds               []*string `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	GroupIds                     []*int64  `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	HideServerCodeWhenProjecting *bool     `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	Instruction                  *bool     `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// example:
	//
	// 1
	IsPicTop *int32 `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	// example:
	//
	// logo
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// example:
	//
	// 测试企业
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 10
	PicturePlayInterval *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	RoomIds             []*string `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
	ShowCalendarCard    *bool     `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	ShowCalendarTitle   *bool     `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	ShowFunctionCard    *bool     `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	EnableCycleReservation *bool `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	// example:
	//
	// 0
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                       `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	OpenReservation      *bool                                         `json:"openReservation,omitempty" xml:"openReservation,omitempty"`
	ReservationAuthority *CreateMeetingRoomRequestReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 10
	RoomCapacity *int32                                `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomLabelIds []*int64                              `json:"roomLabelIds,omitempty" xml:"roomLabelIds,omitempty" type:"Repeated"`
	RoomLocation *CreateMeetingRoomRequestRoomLocation `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 测试会议室
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPDgfLPFjNPu3NAWjNAWg_360_360.jpg
	RoomPicture *string `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.全员可用 1.仅管理员可用
	RoomStatus *int32 `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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

func (s *CreateMeetingRoomRequest) SetOpenReservation(v bool) *CreateMeetingRoomRequest {
	s.OpenReservation = &v
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
	// example:
	//
	// lPHhSZDLXXXXXXpBlC9lxLwiEiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 张三
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	// example:
	//
	// user
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
	// example:
	//
	// xx市xx区xx路xx号
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// xxx公司
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
	// example:
	//
	// 0ffb71843fbb7fc362cb1a0de97fd20b808b09d6ca6282ed
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateMeetingRoomControlPanelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMeetingRoomControlPanelHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomControlPanelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomControlPanelHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMeetingRoomControlPanelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMeetingRoomControlPanelRequest struct {
	Extra *CreateMeetingRoomControlPanelRequestExtra `json:"extra,omitempty" xml:"extra,omitempty" type:"Struct"`
	// This parameter is required.
	RoomConfig []*CreateMeetingRoomControlPanelRequestRoomConfig `json:"roomConfig,omitempty" xml:"roomConfig,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 25SDWxxxxxx
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateMeetingRoomControlPanelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelRequest) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelRequest) SetExtra(v *CreateMeetingRoomControlPanelRequestExtra) *CreateMeetingRoomControlPanelRequest {
	s.Extra = v
	return s
}

func (s *CreateMeetingRoomControlPanelRequest) SetRoomConfig(v []*CreateMeetingRoomControlPanelRequestRoomConfig) *CreateMeetingRoomControlPanelRequest {
	s.RoomConfig = v
	return s
}

func (s *CreateMeetingRoomControlPanelRequest) SetRoomId(v string) *CreateMeetingRoomControlPanelRequest {
	s.RoomId = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequest) SetStatus(v int32) *CreateMeetingRoomControlPanelRequest {
	s.Status = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequest) SetUnionId(v string) *CreateMeetingRoomControlPanelRequest {
	s.UnionId = &v
	return s
}

type CreateMeetingRoomControlPanelRequestExtra struct {
	Param map[string]*string `json:"param,omitempty" xml:"param,omitempty"`
}

func (s CreateMeetingRoomControlPanelRequestExtra) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelRequestExtra) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelRequestExtra) SetParam(v map[string]*string) *CreateMeetingRoomControlPanelRequestExtra {
	s.Param = v
	return s
}

type CreateMeetingRoomControlPanelRequestRoomConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// name
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.xxx.com
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 栗子xx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	ShowTime *int32 `json:"showTime,omitempty" xml:"showTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Sort *int32 `json:"sort,omitempty" xml:"sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.taoxxx.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateMeetingRoomControlPanelRequestRoomConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelRequestRoomConfig) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetEnName(v string) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.EnName = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetIcon(v string) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.Icon = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetName(v string) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.Name = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetShowTime(v int32) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.ShowTime = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetSort(v int32) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.Sort = &v
	return s
}

func (s *CreateMeetingRoomControlPanelRequestRoomConfig) SetUrl(v string) *CreateMeetingRoomControlPanelRequestRoomConfig {
	s.Url = &v
	return s
}

type CreateMeetingRoomControlPanelResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateMeetingRoomControlPanelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelResponseBody) SetResult(v string) *CreateMeetingRoomControlPanelResponseBody {
	s.Result = &v
	return s
}

type CreateMeetingRoomControlPanelResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMeetingRoomControlPanelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMeetingRoomControlPanelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMeetingRoomControlPanelResponse) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomControlPanelResponse) SetHeaders(v map[string]*string) *CreateMeetingRoomControlPanelResponse {
	s.Headers = v
	return s
}

func (s *CreateMeetingRoomControlPanelResponse) SetStatusCode(v int32) *CreateMeetingRoomControlPanelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMeetingRoomControlPanelResponse) SetBody(v *CreateMeetingRoomControlPanelResponseBody) *CreateMeetingRoomControlPanelResponse {
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
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 172
	ParentGroupId *int64 `json:"parentGroupId,omitempty" xml:"parentGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// 172
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteBookingBlacklistHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteBookingBlacklistHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteBookingBlacklistHeaders) GoString() string {
	return s.String()
}

func (s *DeleteBookingBlacklistHeaders) SetCommonHeaders(v map[string]*string) *DeleteBookingBlacklistHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteBookingBlacklistHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteBookingBlacklistHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteBookingBlacklistRequest struct {
	BlacklistUnionIds []*string `json:"blacklistUnionIds,omitempty" xml:"blacklistUnionIds,omitempty" type:"Repeated"`
	UnionId           *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteBookingBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBookingBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteBookingBlacklistRequest) SetBlacklistUnionIds(v []*string) *DeleteBookingBlacklistRequest {
	s.BlacklistUnionIds = v
	return s
}

func (s *DeleteBookingBlacklistRequest) SetUnionId(v string) *DeleteBookingBlacklistRequest {
	s.UnionId = &v
	return s
}

type DeleteBookingBlacklistResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteBookingBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBookingBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBookingBlacklistResponseBody) SetResult(v bool) *DeleteBookingBlacklistResponseBody {
	s.Result = &v
	return s
}

type DeleteBookingBlacklistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBookingBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBookingBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBookingBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteBookingBlacklistResponse) SetHeaders(v map[string]*string) *DeleteBookingBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteBookingBlacklistResponse) SetStatusCode(v int32) *DeleteBookingBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBookingBlacklistResponse) SetBody(v *DeleteBookingBlacklistResponseBody) *DeleteBookingBlacklistResponse {
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
	// This parameter is required.
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
	// example:
	//
	// true
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	// example:
	//
	// true
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteMeetingRoomControlPanelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMeetingRoomControlPanelHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomControlPanelHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomControlPanelHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomControlPanelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomControlPanelHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMeetingRoomControlPanelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMeetingRoomControlPanelRequest struct {
	// This parameter is required.
	RoomIds []*string `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// A1FAxxxxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteMeetingRoomControlPanelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomControlPanelRequest) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomControlPanelRequest) SetRoomIds(v []*string) *DeleteMeetingRoomControlPanelRequest {
	s.RoomIds = v
	return s
}

func (s *DeleteMeetingRoomControlPanelRequest) SetUnionId(v string) *DeleteMeetingRoomControlPanelRequest {
	s.UnionId = &v
	return s
}

type DeleteMeetingRoomControlPanelResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteMeetingRoomControlPanelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomControlPanelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomControlPanelResponseBody) SetResult(v string) *DeleteMeetingRoomControlPanelResponseBody {
	s.Result = &v
	return s
}

type DeleteMeetingRoomControlPanelResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMeetingRoomControlPanelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMeetingRoomControlPanelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMeetingRoomControlPanelResponse) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomControlPanelResponse) SetHeaders(v map[string]*string) *DeleteMeetingRoomControlPanelResponse {
	s.Headers = v
	return s
}

func (s *DeleteMeetingRoomControlPanelResponse) SetStatusCode(v int32) *DeleteMeetingRoomControlPanelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMeetingRoomControlPanelResponse) SetBody(v *DeleteMeetingRoomControlPanelResponseBody) *DeleteMeetingRoomControlPanelResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	// example:
	//
	// true
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	BgImageList []*string `json:"bgImageList,omitempty" xml:"bgImageList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BgType *int32 `json:"bgType,omitempty" xml:"bgType,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01GWWCCR1y2D9D9EHej_!!6000000006520-2-tps-1920-470.png
	BgUrl *string `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	// example:
	//
	// 1
	ConfSubType *int32 `json:"confSubType,omitempty" xml:"confSubType,omitempty"`
	// example:
	//
	// 1
	ConfType *int32 `json:"confType,omitempty" xml:"confType,omitempty"`
	// example:
	//
	// dingc02f685faxxxxc44ac5d6980864d335
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 测试文本
	CustomDoc *string `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	// example:
	//
	// true：脱敏 false：不脱敏
	DesensitizeUserName *bool `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	// example:
	//
	// true：隐藏 false：不隐藏
	HideServerCodeWhenProjecting *bool `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	// example:
	//
	// true：显示 false：不显示
	Instruction *bool `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// example:
	//
	// 1
	IsPicTop *int32 `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	// example:
	//
	// logo
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// example:
	//
	// 测试企业
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 10
	PicturePlayInterval *int32 `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarCard *bool `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarTitle *bool `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowFunctionCard *bool `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	// example:
	//
	// 89
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 测试模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	BgImageList []*string `json:"bgImageList,omitempty" xml:"bgImageList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BgType *int32 `json:"bgType,omitempty" xml:"bgType,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01GWWCCR1y2D9D9EHej_!!6000000006520-2-tps-1920-470.png
	BgUrl *string `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	// example:
	//
	// 1
	ConfSubType *int32 `json:"confSubType,omitempty" xml:"confSubType,omitempty"`
	// example:
	//
	// 1
	ConfType *int32 `json:"confType,omitempty" xml:"confType,omitempty"`
	// example:
	//
	// dingc02f685fxxxx81c44ac5d6980864d335
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 测试文本
	CustomDoc *string `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	// example:
	//
	// true：脱敏 false：不脱敏
	DesensitizeUserName *bool `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	// example:
	//
	// true：隐藏 false：不隐藏
	HideServerCodeWhenProjecting *bool `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	// example:
	//
	// true：显示 false：不显示
	Instruction *bool `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// example:
	//
	// 1
	IsPicTop *int32 `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	// example:
	//
	// logo
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// example:
	//
	// 测试企业
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 10
	PicturePlayInterval *int32 `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarCard *bool `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarTitle *bool `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowFunctionCard *bool `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	// example:
	//
	// 89
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 测试模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceCustomTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1005F1K203604N000676
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
	// example:
	//
	// 30.12.1.100
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceIpByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	PropertyNames []*string `json:"propertyNames,omitempty" xml:"propertyNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1234
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// "lmvUrRkpboRrSMtgsiS9V3AiEiE"
	DeviceUnionId *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "lmvUrEjpboFrSMtgsiS9V3AiEiE"
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
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
	// example:
	//
	// "dev_app_status"
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty"`
	// example:
	//
	// "idle"
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	// example:
	//
	// ding994a046bca84545935c2f4657eb6378f
	CorpId                 *string                                            `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceUnionIds         []*string                                          `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	EnableCycleReservation *bool                                              `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	ExtensionConfig        *QueryMeetingRoomResponseBodyResultExtensionConfig `json:"extensionConfig,omitempty" xml:"extensionConfig,omitempty" type:"Struct"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                                 `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	ReservationAuthority *QueryMeetingRoomResponseBodyResultReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 10
	RoomCapacity *int32                                       `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomGroup    *QueryMeetingRoomResponseBodyResultRoomGroup `json:"roomGroup,omitempty" xml:"roomGroup,omitempty" type:"Struct"`
	// example:
	//
	// 0ffb71843fbb7fc362cb1a0de97fd20b808b09d6ca6282ed
	RoomId       *string                                         `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabels   []*QueryMeetingRoomResponseBodyResultRoomLabels `json:"roomLabels,omitempty" xml:"roomLabels,omitempty" type:"Repeated"`
	RoomLocation *QueryMeetingRoomResponseBodyResultRoomLocation `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPDgfLPFjNPu3NAWjNAWg_360_360.jpg
	RoomPicture *string `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	// example:
	//
	// 01224148194623278976
	RoomStaffId *string `json:"roomStaffId,omitempty" xml:"roomStaffId,omitempty"`
	// example:
	//
	// 0.全员可用 1.仅管理员可用
	RoomStatus *int32 `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	// example:
	//
	// DtB8VDzXXXXXX41rgiE
	RoomUnionId *string `json:"roomUnionId,omitempty" xml:"roomUnionId,omitempty"`
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

func (s *QueryMeetingRoomResponseBodyResult) SetExtensionConfig(v *QueryMeetingRoomResponseBodyResultExtensionConfig) *QueryMeetingRoomResponseBodyResult {
	s.ExtensionConfig = v
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

func (s *QueryMeetingRoomResponseBodyResult) SetRoomUnionId(v string) *QueryMeetingRoomResponseBodyResult {
	s.RoomUnionId = &v
	return s
}

type QueryMeetingRoomResponseBodyResultExtensionConfig struct {
	AdvanceReservation *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation `json:"advanceReservation,omitempty" xml:"advanceReservation,omitempty" type:"Struct"`
	ApprovalSwitch     *bool                                                                `json:"approvalSwitch,omitempty" xml:"approvalSwitch,omitempty"`
	ApprovalType       *int32                                                               `json:"approvalType,omitempty" xml:"approvalType,omitempty"`
	// example:
	//
	// 60
	MaxReservationTimeInterval *int32 `json:"maxReservationTimeInterval,omitempty" xml:"maxReservationTimeInterval,omitempty"`
	// example:
	//
	// 15
	MinReservationTimeInterval *int32                                                                   `json:"minReservationTimeInterval,omitempty" xml:"minReservationTimeInterval,omitempty"`
	OpenReservation            *bool                                                                    `json:"openReservation,omitempty" xml:"openReservation,omitempty"`
	ReservationCloseDetail     *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail `json:"reservationCloseDetail,omitempty" xml:"reservationCloseDetail,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfig) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetAdvanceReservation(v *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.AdvanceReservation = v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetApprovalSwitch(v bool) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.ApprovalSwitch = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetApprovalType(v int32) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.ApprovalType = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetMaxReservationTimeInterval(v int32) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.MaxReservationTimeInterval = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetMinReservationTimeInterval(v int32) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.MinReservationTimeInterval = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetOpenReservation(v bool) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.OpenReservation = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfig) SetReservationCloseDetail(v *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) *QueryMeetingRoomResponseBodyResultExtensionConfig {
	s.ReservationCloseDetail = v
	return s
}

type QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation struct {
	// example:
	//
	// 09:00
	AdvanceBookTimeFormat *string `json:"advanceBookTimeFormat,omitempty" xml:"advanceBookTimeFormat,omitempty"`
	// example:
	//
	// 3
	AdvanceReservationTime *int32 `json:"advanceReservationTime,omitempty" xml:"advanceReservationTime,omitempty"`
	// example:
	//
	// days
	AdvanceReservationTimeUnit *string `json:"advanceReservationTimeUnit,omitempty" xml:"advanceReservationTimeUnit,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) SetAdvanceBookTimeFormat(v string) *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation {
	s.AdvanceBookTimeFormat = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) SetAdvanceReservationTime(v int32) *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation {
	s.AdvanceReservationTime = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation) SetAdvanceReservationTimeUnit(v string) *QueryMeetingRoomResponseBodyResultExtensionConfigAdvanceReservation {
	s.AdvanceReservationTimeUnit = &v
	return s
}

type QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail struct {
	// example:
	//
	// 因为装修临时关闭预定
	CloseReason *string `json:"closeReason,omitempty" xml:"closeReason,omitempty"`
	// example:
	//
	// nick
	ContactNick *string `json:"contactNick,omitempty" xml:"contactNick,omitempty"`
	// example:
	//
	// 2iPOLbpxxxxuwggiiqiPwiEiF
	ContactUnionId *string `json:"contactUnionId,omitempty" xml:"contactUnionId,omitempty"`
	SendNotify     *bool   `json:"sendNotify,omitempty" xml:"sendNotify,omitempty"`
	// example:
	//
	// 1740045030000
	TaskEndTime *int64 `json:"taskEndTime,omitempty" xml:"taskEndTime,omitempty"`
	// example:
	//
	// 1740463800000
	TaskStartTime *int64 `json:"taskStartTime,omitempty" xml:"taskStartTime,omitempty"`
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetCloseReason(v string) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.CloseReason = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetContactNick(v string) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.ContactNick = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetContactUnionId(v string) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.ContactUnionId = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetSendNotify(v bool) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.SendNotify = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetTaskEndTime(v int64) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.TaskEndTime = &v
	return s
}

func (s *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail) SetTaskStartTime(v int64) *QueryMeetingRoomResponseBodyResultExtensionConfigReservationCloseDetail {
	s.TaskStartTime = &v
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
	// example:
	//
	// lPHhSZDLXXXXXXpBlC9lxLwiEiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 张三
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	// example:
	//
	// user
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
	// example:
	//
	// 1
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
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
	// example:
	//
	// xx市xx区xx街道xx号
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// xxx公司
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryMeetingRoomControlPanelListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMeetingRoomControlPanelListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomControlPanelListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomControlPanelListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMeetingRoomControlPanelListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMeetingRoomControlPanelListRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 2iPOLxxxxx
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMeetingRoomControlPanelListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListRequest) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListRequest) SetMaxResults(v int32) *QueryMeetingRoomControlPanelListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListRequest) SetNextToken(v int64) *QueryMeetingRoomControlPanelListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListRequest) SetRoomId(v string) *QueryMeetingRoomControlPanelListRequest {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListRequest) SetUnionId(v string) *QueryMeetingRoomControlPanelListRequest {
	s.UnionId = &v
	return s
}

type QueryMeetingRoomControlPanelListResponseBody struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 0
	NextToken *int64                                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Result    []*QueryMeetingRoomControlPanelListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryMeetingRoomControlPanelListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListResponseBody) SetHasMore(v bool) *QueryMeetingRoomControlPanelListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBody) SetNextToken(v int64) *QueryMeetingRoomControlPanelListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBody) SetResult(v []*QueryMeetingRoomControlPanelListResponseBodyResult) *QueryMeetingRoomControlPanelListResponseBody {
	s.Result = v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBody) SetTotalCount(v int32) *QueryMeetingRoomControlPanelListResponseBody {
	s.TotalCount = &v
	return s
}

type QueryMeetingRoomControlPanelListResponseBodyResult struct {
	// example:
	//
	// 1WADFxxxxxx
	RoomId        *string                                                            `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomIotConfig []*QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig `json:"roomIotConfig,omitempty" xml:"roomIotConfig,omitempty" type:"Repeated"`
}

func (s QueryMeetingRoomControlPanelListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResult) SetRoomId(v string) *QueryMeetingRoomControlPanelListResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResult) SetRoomIotConfig(v []*QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) *QueryMeetingRoomControlPanelListResponseBodyResult {
	s.RoomIotConfig = v
	return s
}

type QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig struct {
	// example:
	//
	// name
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
	// example:
	//
	// https://www.taoxxxxx.com
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 栗子xx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 30
	ShowTime *int32 `json:"showTime,omitempty" xml:"showTime,omitempty"`
	// example:
	//
	// 0
	Sort *int32 `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// https://www.taoxxxxx.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetEnName(v string) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.EnName = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetIcon(v string) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.Icon = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetName(v string) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.Name = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetShowTime(v int32) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.ShowTime = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetSort(v int32) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.Sort = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig) SetUrl(v string) *QueryMeetingRoomControlPanelListResponseBodyResultRoomIotConfig {
	s.Url = &v
	return s
}

type QueryMeetingRoomControlPanelListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomControlPanelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMeetingRoomControlPanelListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMeetingRoomControlPanelListResponse) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomControlPanelListResponse) SetHeaders(v map[string]*string) *QueryMeetingRoomControlPanelListResponse {
	s.Headers = v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponse) SetStatusCode(v int32) *QueryMeetingRoomControlPanelListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMeetingRoomControlPanelListResponse) SetBody(v *QueryMeetingRoomControlPanelListResponseBody) *QueryMeetingRoomControlPanelListResponse {
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
	// example:
	//
	// 1234
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// "lmvUrRkpboRrSMtgsiS9V3AiEiE"
	DeviceUnionId *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "lmvUrEjpboFrSMtgsiS9V3AiEiE"
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
	// example:
	//
	// 1697198045000
	ActiveTime  *int64                                                 `json:"activeTime,omitempty" xml:"activeTime,omitempty"`
	Controllers []*QueryMeetingRoomDeviceResponseBodyResultControllers `json:"controllers,omitempty" xml:"controllers,omitempty" type:"Repeated"`
	// example:
	//
	// "ding994a046bca84545935c2f4657eb6378f"
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// lPHhSZDLXXXXXXpBlC9lxLwiEiE
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// example:
	//
	// Smart Camera
	DevCamera *string `json:"devCamera,omitempty" xml:"devCamera,omitempty"`
	// example:
	//
	// false
	DevHdmi *string `json:"devHdmi,omitempty" xml:"devHdmi,omitempty"`
	// example:
	//
	// Microphone (2- Built-in Audio)
	DevMic *string `json:"devMic,omitempty" xml:"devMic,omitempty"`
	// example:
	//
	// false
	DevMirror *string `json:"devMirror,omitempty" xml:"devMirror,omitempty"`
	// example:
	//
	// 127.0.0.10
	DevNetIp *string `json:"devNetIp,omitempty" xml:"devNetIp,omitempty"`
	// example:
	//
	// net_wired
	DevNetType *string `json:"devNetType,omitempty" xml:"devNetType,omitempty"`
	// example:
	//
	// Speaker (2- Built-in Audio)
	DevVoice *string `json:"devVoice,omitempty" xml:"devVoice,omitempty"`
	// example:
	//
	// d4:aa:ee:e8:4d:55
	DevWifiMac *string `json:"devWifiMac,omitempty" xml:"devWifiMac,omitempty"`
	// example:
	//
	// d4:3a:ee:aa:45:85
	DevWireMac *string `json:"devWireMac,omitempty" xml:"devWireMac,omitempty"`
	// example:
	//
	// 1234
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// "14:85:7f:e5:f3:f3"
	DeviceMac *string `json:"deviceMac,omitempty" xml:"deviceMac,omitempty"`
	// example:
	//
	// winbox
	DeviceModel *string `json:"deviceModel,omitempty" xml:"deviceModel,omitempty"`
	// example:
	//
	// 钉钉会议设备_xxxx
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// 1204
	DeviceServiceId *int32 `json:"deviceServiceId,omitempty" xml:"deviceServiceId,omitempty"`
	// example:
	//
	// "02caa8169c80f74a2d375093a6107016"
	DeviceSn *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	// example:
	//
	// 空闲：idle  投屏中：projection   会议响铃中：conf_incoming   会议中：conf_running   使用白板中：white_board   离线: offline
	DeviceStatus *string `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	// example:
	//
	// 视频会议设备:"touyingyi"   设备控制器:"meetingaccessory"
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// "lmvUrRkpboRrSMtgsiS9V3AiEiE"
	DeviceUnionId *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	// example:
	//
	// LMVXXX.20XX0818
	FirmwareVersion *string `json:"firmwareVersion,omitempty" xml:"firmwareVersion,omitempty"`
	// example:
	//
	// "7263defed6b361fedf0fe6a3b578b96e808b09d6ca6282ed"
	OpenRoomId *string `json:"openRoomId,omitempty" xml:"openRoomId,omitempty"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// 123456
	ShareCode *string `json:"shareCode,omitempty" xml:"shareCode,omitempty"`
	// example:
	//
	// sip13492
	SipAccountName *string `json:"sipAccountName,omitempty" xml:"sipAccountName,omitempty"`
	// example:
	//
	// 7.14.1
	SoftwareVersion *string `json:"softwareVersion,omitempty" xml:"softwareVersion,omitempty"`
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
	// example:
	//
	// "ding994a046bca84545935c2f4657eb6378f"
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 2345
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// "d8:2f:e6:d9:ab:5b"
	DeviceMac *string `json:"deviceMac,omitempty" xml:"deviceMac,omitempty"`
	// example:
	//
	// "AILABS_S3_T1"
	DeviceModel *string `json:"deviceModel,omitempty" xml:"deviceModel,omitempty"`
	// example:
	//
	// 会控平板_xxxx
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// example:
	//
	// 1190
	DeviceServiceId *int32 `json:"deviceServiceId,omitempty" xml:"deviceServiceId,omitempty"`
	// example:
	//
	// "02caa8169c80f74a2d375093a6107017"
	DeviceSn *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	// example:
	//
	// 空闲：idle  投屏中：projection   会议响铃中：conf_incoming   会议中：conf_running   使用白板中：white_board   离线: offline
	DeviceStatus *string `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	// example:
	//
	// 视频会议设备:"touyingyi"   设备控制器:"meetingaccessory"
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// "lmvUrRkpboRrSMtgsiS9V4AiEiE"
	DeviceUnionId *string `json:"deviceUnionId,omitempty" xml:"deviceUnionId,omitempty"`
	// example:
	//
	// "7263defed6b361fedf0fe6a3b578b96e808b09d6ca6282ed"
	OpenRoomId *string `json:"openRoomId,omitempty" xml:"openRoomId,omitempty"`
	// example:
	//
	// 234567
	ShareCode *string `json:"shareCode,omitempty" xml:"shareCode,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	// example:
	//
	// 172
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	// example:
	//
	// 172
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 124
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 123
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
	// example:
	//
	// ding994a046bca84545935c2f4657eb6378f
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId *string `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	// example:
	//
	// 10
	RoomCapacity *int32                                           `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	RoomGroup    *QueryMeetingRoomListResponseBodyResultRoomGroup `json:"roomGroup,omitempty" xml:"roomGroup,omitempty" type:"Struct"`
	// example:
	//
	// 0ffb71843fbb7fc362cb1a0de97fd20b808b09d6ca6282ed
	RoomId       *string                                             `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabels   []*QueryMeetingRoomListResponseBodyResultRoomLabels `json:"roomLabels,omitempty" xml:"roomLabels,omitempty" type:"Repeated"`
	RoomLocation *QueryMeetingRoomListResponseBodyResultRoomLocation `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPDgfLPFjNPu3NAWjNAWg_360_360.jpg
	RoomPicture *string `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	// example:
	//
	// 01224148194623278976
	RoomStaffId *string `json:"roomStaffId,omitempty" xml:"roomStaffId,omitempty"`
	// example:
	//
	// 0.全员可用 1.仅管理员可用
	RoomStatus *int32 `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
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
	// example:
	//
	// 1
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 0
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
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
	// example:
	//
	// xx市xx区xx街道xx号
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// xxx公司
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMeetingRoomListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 0ffb71843fbb7fc362cb1a0de97fd20b808b09d6ca6282ed
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSuperUserMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	DeptIdWhiteList []*int64 `json:"deptIdWhiteList,omitempty" xml:"deptIdWhiteList,omitempty" type:"Repeated"`
	// This parameter is required.
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OcMXXXXXM2eRogiEiE
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSuperUserMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	BgImgList []*string `json:"bgImgList,omitempty" xml:"bgImgList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BgType *int32 `json:"bgType,omitempty" xml:"bgType,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01GWWCCR1y2D9D9EHej_!!6000000006520-2-tps-1920-470.png
	BgUrl *string `json:"bgUrl,omitempty" xml:"bgUrl,omitempty"`
	// example:
	//
	// 测试文本
	CustomDoc *string `json:"customDoc,omitempty" xml:"customDoc,omitempty"`
	// example:
	//
	// true：脱敏 false：不脱敏
	DesensitizeUserName *bool     `json:"desensitizeUserName,omitempty" xml:"desensitizeUserName,omitempty"`
	DeviceUnionIds      []*string `json:"deviceUnionIds,omitempty" xml:"deviceUnionIds,omitempty" type:"Repeated"`
	GroupIds            []*int64  `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// example:
	//
	// true：隐藏 false：不隐藏
	HideServerCodeWhenProjecting *bool `json:"hideServerCodeWhenProjecting,omitempty" xml:"hideServerCodeWhenProjecting,omitempty"`
	// example:
	//
	// true：显示 false：不显示
	Instruction *bool `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// example:
	//
	// 1
	IsPicTop *int32 `json:"isPicTop,omitempty" xml:"isPicTop,omitempty"`
	// example:
	//
	// logo
	Logo *string `json:"logo,omitempty" xml:"logo,omitempty"`
	// example:
	//
	// 测试企业
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 10
	PicturePlayInterval *int32    `json:"picturePlayInterval,omitempty" xml:"picturePlayInterval,omitempty"`
	RoomIds             []*string `json:"roomIds,omitempty" xml:"roomIds,omitempty" type:"Repeated"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarCard *bool `json:"showCalendarCard,omitempty" xml:"showCalendarCard,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowCalendarTitle *bool `json:"showCalendarTitle,omitempty" xml:"showCalendarTitle,omitempty"`
	// example:
	//
	// true：展示 false：不展示
	ShowFunctionCard *bool `json:"showFunctionCard,omitempty" xml:"showFunctionCard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 89
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeviceCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	EnableCycleReservation *bool `json:"enableCycleReservation,omitempty" xml:"enableCycleReservation,omitempty"`
	// example:
	//
	// 0
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// xxxIsvRoomId
	IsvRoomId            *string                                       `json:"isvRoomId,omitempty" xml:"isvRoomId,omitempty"`
	OpenReservation      *bool                                         `json:"openReservation,omitempty" xml:"openReservation,omitempty"`
	ReservationAuthority *UpdateMeetingRoomRequestReservationAuthority `json:"reservationAuthority,omitempty" xml:"reservationAuthority,omitempty" type:"Struct"`
	// example:
	//
	// 10
	RoomCapacity *int32 `json:"roomCapacity,omitempty" xml:"roomCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0ffb71843fbb7fc362cb1a0de97fd20b808b09d6ca6282ed
	RoomId       *string                               `json:"roomId,omitempty" xml:"roomId,omitempty"`
	RoomLabelIds []*int64                              `json:"roomLabelIds,omitempty" xml:"roomLabelIds,omitempty" type:"Repeated"`
	RoomLocation *UpdateMeetingRoomRequestRoomLocation `json:"roomLocation,omitempty" xml:"roomLocation,omitempty" type:"Struct"`
	// example:
	//
	// 测试会议室
	RoomName *string `json:"roomName,omitempty" xml:"roomName,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/lADPDgfLPFjNPu3NAWjNAWg_360_360.jpg
	RoomPicture *string `json:"roomPicture,omitempty" xml:"roomPicture,omitempty"`
	// example:
	//
	// 0.全员可用 1.仅管理员可用
	RoomStatus *int32 `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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

func (s *UpdateMeetingRoomRequest) SetOpenReservation(v bool) *UpdateMeetingRoomRequest {
	s.OpenReservation = &v
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
	// example:
	//
	// lPHhSZDLXXXXXXpBlC9lxLwiEiE
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// example:
	//
	// 张三
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	// example:
	//
	// user
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
	// example:
	//
	// xx市xx区xx路xx号
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// xxx公司
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
	// example:
	//
	// true
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeetingRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 172
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 测试分组
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeetingRoomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
// 创建会议室预定黑名单
//
// @param request - CreateBookingBlacklistRequest
//
// @param headers - CreateBookingBlacklistHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBookingBlacklistResponse
func (client *Client) CreateBookingBlacklistWithOptions(request *CreateBookingBlacklistRequest, headers *CreateBookingBlacklistHeaders, runtime *util.RuntimeOptions) (_result *CreateBookingBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlacklistUnionId)) {
		body["blacklistUnionId"] = request.BlacklistUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		body["memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("CreateBookingBlacklist"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/bookings/blacklist"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBookingBlacklistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会议室预定黑名单
//
// @param request - CreateBookingBlacklistRequest
//
// @return CreateBookingBlacklistResponse
func (client *Client) CreateBookingBlacklist(request *CreateBookingBlacklistRequest) (_result *CreateBookingBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBookingBlacklistHeaders{}
	_result = &CreateBookingBlacklistResponse{}
	_body, _err := client.CreateBookingBlacklistWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义屏幕模版
//
// @param request - CreateDeviceCustomTemplateRequest
//
// @param headers - CreateDeviceCustomTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceCustomTemplateResponse
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

// Summary:
//
// 创建自定义屏幕模版
//
// @param request - CreateDeviceCustomTemplateRequest
//
// @return CreateDeviceCustomTemplateResponse
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

// Summary:
//
// 创建智能会议室
//
// @param request - CreateMeetingRoomRequest
//
// @param headers - CreateMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMeetingRoomResponse
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

	if !tea.BoolValue(util.IsUnset(request.OpenReservation)) {
		body["openReservation"] = request.OpenReservation
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

// Summary:
//
// 创建智能会议室
//
// @param request - CreateMeetingRoomRequest
//
// @return CreateMeetingRoomResponse
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

// Summary:
//
// 创建智能会议室IOT配置
//
// @param request - CreateMeetingRoomControlPanelRequest
//
// @param headers - CreateMeetingRoomControlPanelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMeetingRoomControlPanelResponse
func (client *Client) CreateMeetingRoomControlPanelWithOptions(request *CreateMeetingRoomControlPanelRequest, headers *CreateMeetingRoomControlPanelHeaders, runtime *util.RuntimeOptions) (_result *CreateMeetingRoomControlPanelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.RoomConfig)) {
		body["roomConfig"] = request.RoomConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["roomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("CreateMeetingRoomControlPanel"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/panels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMeetingRoomControlPanelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建智能会议室IOT配置
//
// @param request - CreateMeetingRoomControlPanelRequest
//
// @return CreateMeetingRoomControlPanelResponse
func (client *Client) CreateMeetingRoomControlPanel(request *CreateMeetingRoomControlPanelRequest) (_result *CreateMeetingRoomControlPanelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMeetingRoomControlPanelHeaders{}
	_result = &CreateMeetingRoomControlPanelResponse{}
	_body, _err := client.CreateMeetingRoomControlPanelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建会议室分组
//
// @param request - CreateMeetingRoomGroupRequest
//
// @param headers - CreateMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMeetingRoomGroupResponse
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

// Summary:
//
// 创建会议室分组
//
// @param request - CreateMeetingRoomGroupRequest
//
// @return CreateMeetingRoomGroupResponse
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

// Summary:
//
// 删除会议室预定黑名单
//
// @param request - DeleteBookingBlacklistRequest
//
// @param headers - DeleteBookingBlacklistHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBookingBlacklistResponse
func (client *Client) DeleteBookingBlacklistWithOptions(request *DeleteBookingBlacklistRequest, headers *DeleteBookingBlacklistHeaders, runtime *util.RuntimeOptions) (_result *DeleteBookingBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlacklistUnionIds)) {
		body["blacklistUnionIds"] = request.BlacklistUnionIds
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
		Action:      tea.String("DeleteBookingBlacklist"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/bookings/blacklist/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBookingBlacklistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会议室预定黑名单
//
// @param request - DeleteBookingBlacklistRequest
//
// @return DeleteBookingBlacklistResponse
func (client *Client) DeleteBookingBlacklist(request *DeleteBookingBlacklistRequest) (_result *DeleteBookingBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteBookingBlacklistHeaders{}
	_result = &DeleteBookingBlacklistResponse{}
	_body, _err := client.DeleteBookingBlacklistWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义屏幕模板
//
// @param request - DeleteDeviceCustomTemplateRequest
//
// @param headers - DeleteDeviceCustomTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeviceCustomTemplateResponse
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

// Summary:
//
// 删除自定义屏幕模板
//
// @param request - DeleteDeviceCustomTemplateRequest
//
// @return DeleteDeviceCustomTemplateResponse
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

// Summary:
//
// 删除会议室
//
// @param request - DeleteMeetingRoomRequest
//
// @param headers - DeleteMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMeetingRoomResponse
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

// Summary:
//
// 删除会议室
//
// @param request - DeleteMeetingRoomRequest
//
// @return DeleteMeetingRoomResponse
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

// Summary:
//
// 删除会议室配置
//
// @param request - DeleteMeetingRoomControlPanelRequest
//
// @param headers - DeleteMeetingRoomControlPanelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMeetingRoomControlPanelResponse
func (client *Client) DeleteMeetingRoomControlPanelWithOptions(request *DeleteMeetingRoomControlPanelRequest, headers *DeleteMeetingRoomControlPanelHeaders, runtime *util.RuntimeOptions) (_result *DeleteMeetingRoomControlPanelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoomIds)) {
		body["roomIds"] = request.RoomIds
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
		Action:      tea.String("DeleteMeetingRoomControlPanel"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/panels/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMeetingRoomControlPanelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除会议室配置
//
// @param request - DeleteMeetingRoomControlPanelRequest
//
// @return DeleteMeetingRoomControlPanelResponse
func (client *Client) DeleteMeetingRoomControlPanel(request *DeleteMeetingRoomControlPanelRequest) (_result *DeleteMeetingRoomControlPanelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMeetingRoomControlPanelHeaders{}
	_result = &DeleteMeetingRoomControlPanelResponse{}
	_body, _err := client.DeleteMeetingRoomControlPanelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除会议室分组
//
// @param request - DeleteMeetingRoomGroupRequest
//
// @param headers - DeleteMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMeetingRoomGroupResponse
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

// Summary:
//
// 删除会议室分组
//
// @param request - DeleteMeetingRoomGroupRequest
//
// @return DeleteMeetingRoomGroupResponse
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

// Summary:
//
// 查询自定义屏幕模板
//
// @param headers - QueryDeviceCustomTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceCustomTemplateResponse
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

// Summary:
//
// 查询自定义屏幕模板
//
// @return QueryDeviceCustomTemplateResponse
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

// Summary:
//
// 查询自定义屏幕模板列表
//
// @param headers - QueryDeviceCustomTemplateListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceCustomTemplateListResponse
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

// Summary:
//
// 查询自定义屏幕模板列表
//
// @return QueryDeviceCustomTemplateListResponse
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

// Summary:
//
// 根据设备投屏码查询设备ip
//
// @param request - QueryDeviceIpByCodeRequest
//
// @param headers - QueryDeviceIpByCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceIpByCodeResponse
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

// Summary:
//
// 根据设备投屏码查询设备ip
//
// @param request - QueryDeviceIpByCodeRequest
//
// @return QueryDeviceIpByCodeResponse
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

// Summary:
//
// 查询设备属性
//
// @param request - QueryDevicePropertiesRequest
//
// @param headers - QueryDevicePropertiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicePropertiesResponse
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

// Summary:
//
// 查询设备属性
//
// @param request - QueryDevicePropertiesRequest
//
// @return QueryDevicePropertiesResponse
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

// Summary:
//
// 查询会议室详情
//
// @param request - QueryMeetingRoomRequest
//
// @param headers - QueryMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomResponse
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

// Summary:
//
// 查询会议室详情
//
// @param request - QueryMeetingRoomRequest
//
// @return QueryMeetingRoomResponse
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

// Summary:
//
// 获取会议室IOT配置列表
//
// @param request - QueryMeetingRoomControlPanelListRequest
//
// @param headers - QueryMeetingRoomControlPanelListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomControlPanelListResponse
func (client *Client) QueryMeetingRoomControlPanelListWithOptions(request *QueryMeetingRoomControlPanelListRequest, headers *QueryMeetingRoomControlPanelListHeaders, runtime *util.RuntimeOptions) (_result *QueryMeetingRoomControlPanelListResponse, _err error) {
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
		Action:      tea.String("QueryMeetingRoomControlPanelList"),
		Version:     tea.String("rooms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/rooms/panels/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMeetingRoomControlPanelListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会议室IOT配置列表
//
// @param request - QueryMeetingRoomControlPanelListRequest
//
// @return QueryMeetingRoomControlPanelListResponse
func (client *Client) QueryMeetingRoomControlPanelList(request *QueryMeetingRoomControlPanelListRequest) (_result *QueryMeetingRoomControlPanelListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMeetingRoomControlPanelListHeaders{}
	_result = &QueryMeetingRoomControlPanelListResponse{}
	_body, _err := client.QueryMeetingRoomControlPanelListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备信息
//
// @param request - QueryMeetingRoomDeviceRequest
//
// @param headers - QueryMeetingRoomDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomDeviceResponse
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

// Summary:
//
// 查询设备信息
//
// @param request - QueryMeetingRoomDeviceRequest
//
// @return QueryMeetingRoomDeviceResponse
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

// Summary:
//
// 查询会议室分组信息
//
// @param request - QueryMeetingRoomGroupRequest
//
// @param headers - QueryMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomGroupResponse
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

// Summary:
//
// 查询会议室分组信息
//
// @param request - QueryMeetingRoomGroupRequest
//
// @return QueryMeetingRoomGroupResponse
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

// Summary:
//
// 查询会议室分组列表
//
// @param request - QueryMeetingRoomGroupListRequest
//
// @param headers - QueryMeetingRoomGroupListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomGroupListResponse
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

// Summary:
//
// 查询会议室分组列表
//
// @param request - QueryMeetingRoomGroupListRequest
//
// @return QueryMeetingRoomGroupListResponse
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

// Summary:
//
// 查询会议室列表
//
// @param request - QueryMeetingRoomListRequest
//
// @param headers - QueryMeetingRoomListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMeetingRoomListResponse
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

// Summary:
//
// 查询会议室列表
//
// @param request - QueryMeetingRoomListRequest
//
// @return QueryMeetingRoomListResponse
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

// Summary:
//
// 取消会议室高级用户模式。
//
// @param request - RemoveSuperUserMeetingRoomRequest
//
// @param headers - RemoveSuperUserMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSuperUserMeetingRoomResponse
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

// Summary:
//
// 取消会议室高级用户模式。
//
// @param request - RemoveSuperUserMeetingRoomRequest
//
// @return RemoveSuperUserMeetingRoomResponse
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

// Summary:
//
// 设置会议室成为高级用户模式。只有设置在白名单里的人员或部门，才能呼叫此会议室。
//
// @param request - SetSuperUserMeetingRoomRequest
//
// @param headers - SetSuperUserMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuperUserMeetingRoomResponse
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

// Summary:
//
// 设置会议室成为高级用户模式。只有设置在白名单里的人员或部门，才能呼叫此会议室。
//
// @param request - SetSuperUserMeetingRoomRequest
//
// @return SetSuperUserMeetingRoomResponse
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

// Summary:
//
// 更新自定义屏幕模板
//
// @param request - UpdateDeviceCustomTemplateRequest
//
// @param headers - UpdateDeviceCustomTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeviceCustomTemplateResponse
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

// Summary:
//
// 更新自定义屏幕模板
//
// @param request - UpdateDeviceCustomTemplateRequest
//
// @return UpdateDeviceCustomTemplateResponse
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

// Summary:
//
// 更新会议室信息
//
// @param request - UpdateMeetingRoomRequest
//
// @param headers - UpdateMeetingRoomHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeetingRoomResponse
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

	if !tea.BoolValue(util.IsUnset(request.OpenReservation)) {
		body["openReservation"] = request.OpenReservation
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

// Summary:
//
// 更新会议室信息
//
// @param request - UpdateMeetingRoomRequest
//
// @return UpdateMeetingRoomResponse
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

// Summary:
//
// 更新会议室分组
//
// @param request - UpdateMeetingRoomGroupRequest
//
// @param headers - UpdateMeetingRoomGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMeetingRoomGroupResponse
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

// Summary:
//
// 更新会议室分组
//
// @param request - UpdateMeetingRoomGroupRequest
//
// @return UpdateMeetingRoomGroupResponse
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
