// This file is auto-generated, don't edit it. Thanks.
package conference_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type MetricMapValue struct {
	// example:
	//
	// 1682562120000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 1145172
	SendBitRate *string `json:"sendBitRate,omitempty" xml:"sendBitRate,omitempty"`
	// example:
	//
	// 111
	RecvBitRate *string `json:"recvBitRate,omitempty" xml:"recvBitRate,omitempty"`
	// example:
	//
	// 0
	LostRate *string `json:"lostRate,omitempty" xml:"lostRate,omitempty"`
	// example:
	//
	// 20
	RoundTripTime *string `json:"roundTripTime,omitempty" xml:"roundTripTime,omitempty"`
	// example:
	//
	// 59103
	AudioSendBitRate *string `json:"audioSendBitRate,omitempty" xml:"audioSendBitRate,omitempty"`
	// example:
	//
	// 52939
	AudioRecvBitRate *string `json:"audioRecvBitRate,omitempty" xml:"audioRecvBitRate,omitempty"`
	// example:
	//
	// 25
	AudioRecLevel *string `json:"audioRecLevel,omitempty" xml:"audioRecLevel,omitempty"`
	// example:
	//
	// 27
	AudioPlayLevel *string `json:"audioPlayLevel,omitempty" xml:"audioPlayLevel,omitempty"`
	// example:
	//
	// 1145172
	CameraSendBitRate *string `json:"cameraSendBitRate,omitempty" xml:"cameraSendBitRate,omitempty"`
	// example:
	//
	// 66160
	CameraRecvBitRate *string `json:"cameraRecvBitRate,omitempty" xml:"cameraRecvBitRate,omitempty"`
	// example:
	//
	// 1920*1080
	CameraSendResolutionActual *string `json:"cameraSendResolutionActual,omitempty" xml:"cameraSendResolutionActual,omitempty"`
	// example:
	//
	// 1920*1080
	CameraRecvResolutionActual *string `json:"cameraRecvResolutionActual,omitempty" xml:"cameraRecvResolutionActual,omitempty"`
	// example:
	//
	// 20
	CameraSendFrame *string `json:"cameraSendFrame,omitempty" xml:"cameraSendFrame,omitempty"`
	// example:
	//
	// 15701
	ScreenSendBitRate *string `json:"screenSendBitRate,omitempty" xml:"screenSendBitRate,omitempty"`
	// example:
	//
	// 20
	CameraRecvFrame *string `json:"cameraRecvFrame,omitempty" xml:"cameraRecvFrame,omitempty"`
	// example:
	//
	// 0
	ScreenRecvBitRate *string `json:"screenRecvBitRate,omitempty" xml:"screenRecvBitRate,omitempty"`
	// example:
	//
	// 1920*1080
	ScreenSendResolutionActual *string `json:"screenSendResolutionActual,omitempty" xml:"screenSendResolutionActual,omitempty"`
	// example:
	//
	// 1920*1080
	ScreenRecvResolutionActual *string `json:"screenRecvResolutionActual,omitempty" xml:"screenRecvResolutionActual,omitempty"`
	// example:
	//
	// 14
	ScreenSendFrame *string `json:"screenSendFrame,omitempty" xml:"screenSendFrame,omitempty"`
	// example:
	//
	// 0
	ScreenRecvFrame *string `json:"screenRecvFrame,omitempty" xml:"screenRecvFrame,omitempty"`
	// example:
	//
	// 0
	AudioJitterMax *string `json:"audioJitterMax,omitempty" xml:"audioJitterMax,omitempty"`
	// example:
	//
	// 0
	AudioJitterAvg *string `json:"audioJitterAvg,omitempty" xml:"audioJitterAvg,omitempty"`
}

func (s MetricMapValue) String() string {
	return tea.Prettify(s)
}

func (s MetricMapValue) GoString() string {
	return s.String()
}

func (s *MetricMapValue) SetTimestamp(v int64) *MetricMapValue {
	s.Timestamp = &v
	return s
}

func (s *MetricMapValue) SetSendBitRate(v string) *MetricMapValue {
	s.SendBitRate = &v
	return s
}

func (s *MetricMapValue) SetRecvBitRate(v string) *MetricMapValue {
	s.RecvBitRate = &v
	return s
}

func (s *MetricMapValue) SetLostRate(v string) *MetricMapValue {
	s.LostRate = &v
	return s
}

func (s *MetricMapValue) SetRoundTripTime(v string) *MetricMapValue {
	s.RoundTripTime = &v
	return s
}

func (s *MetricMapValue) SetAudioSendBitRate(v string) *MetricMapValue {
	s.AudioSendBitRate = &v
	return s
}

func (s *MetricMapValue) SetAudioRecvBitRate(v string) *MetricMapValue {
	s.AudioRecvBitRate = &v
	return s
}

func (s *MetricMapValue) SetAudioRecLevel(v string) *MetricMapValue {
	s.AudioRecLevel = &v
	return s
}

func (s *MetricMapValue) SetAudioPlayLevel(v string) *MetricMapValue {
	s.AudioPlayLevel = &v
	return s
}

func (s *MetricMapValue) SetCameraSendBitRate(v string) *MetricMapValue {
	s.CameraSendBitRate = &v
	return s
}

func (s *MetricMapValue) SetCameraRecvBitRate(v string) *MetricMapValue {
	s.CameraRecvBitRate = &v
	return s
}

func (s *MetricMapValue) SetCameraSendResolutionActual(v string) *MetricMapValue {
	s.CameraSendResolutionActual = &v
	return s
}

func (s *MetricMapValue) SetCameraRecvResolutionActual(v string) *MetricMapValue {
	s.CameraRecvResolutionActual = &v
	return s
}

func (s *MetricMapValue) SetCameraSendFrame(v string) *MetricMapValue {
	s.CameraSendFrame = &v
	return s
}

func (s *MetricMapValue) SetScreenSendBitRate(v string) *MetricMapValue {
	s.ScreenSendBitRate = &v
	return s
}

func (s *MetricMapValue) SetCameraRecvFrame(v string) *MetricMapValue {
	s.CameraRecvFrame = &v
	return s
}

func (s *MetricMapValue) SetScreenRecvBitRate(v string) *MetricMapValue {
	s.ScreenRecvBitRate = &v
	return s
}

func (s *MetricMapValue) SetScreenSendResolutionActual(v string) *MetricMapValue {
	s.ScreenSendResolutionActual = &v
	return s
}

func (s *MetricMapValue) SetScreenRecvResolutionActual(v string) *MetricMapValue {
	s.ScreenRecvResolutionActual = &v
	return s
}

func (s *MetricMapValue) SetScreenSendFrame(v string) *MetricMapValue {
	s.ScreenSendFrame = &v
	return s
}

func (s *MetricMapValue) SetScreenRecvFrame(v string) *MetricMapValue {
	s.ScreenRecvFrame = &v
	return s
}

func (s *MetricMapValue) SetAudioJitterMax(v string) *MetricMapValue {
	s.AudioJitterMax = &v
	return s
}

func (s *MetricMapValue) SetAudioJitterAvg(v string) *MetricMapValue {
	s.AudioJitterAvg = &v
	return s
}

type MemberModelMapValue struct {
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 654058f2411fe90147e68780
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 测试昵称
	UserNick *string `json:"userNick,omitempty" xml:"userNick,omitempty"`
	// example:
	//
	// 1699347295876
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// example:
	//
	// 1699347395876
	LeaveTime *int64 `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	// example:
	//
	// 100000
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1：初始化  2：呼叫中  3：活跃（在会）  4：入会失败（拒接等）  5：被踢  6：离会
	AttendStatus *int32 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	// example:
	//
	// true：是  false：否
	Host *bool `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// true：是  false：否
	CoHost *bool `json:"coHost,omitempty" xml:"coHost,omitempty"`
	// example:
	//
	// true：是  false：否
	OuterOrgMember *bool `json:"outerOrgMember,omitempty" xml:"outerOrgMember,omitempty"`
	// example:
	//
	// true：是  false：否
	PstnJoin *bool `json:"pstnJoin,omitempty" xml:"pstnJoin,omitempty"`
	// example:
	//
	// Win Mac iOS Android
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
}

func (s MemberModelMapValue) String() string {
	return tea.Prettify(s)
}

func (s MemberModelMapValue) GoString() string {
	return s.String()
}

func (s *MemberModelMapValue) SetUnionId(v string) *MemberModelMapValue {
	s.UnionId = &v
	return s
}

func (s *MemberModelMapValue) SetConferenceId(v string) *MemberModelMapValue {
	s.ConferenceId = &v
	return s
}

func (s *MemberModelMapValue) SetUserNick(v string) *MemberModelMapValue {
	s.UserNick = &v
	return s
}

func (s *MemberModelMapValue) SetJoinTime(v int64) *MemberModelMapValue {
	s.JoinTime = &v
	return s
}

func (s *MemberModelMapValue) SetLeaveTime(v int64) *MemberModelMapValue {
	s.LeaveTime = &v
	return s
}

func (s *MemberModelMapValue) SetDuration(v int64) *MemberModelMapValue {
	s.Duration = &v
	return s
}

func (s *MemberModelMapValue) SetAttendStatus(v int32) *MemberModelMapValue {
	s.AttendStatus = &v
	return s
}

func (s *MemberModelMapValue) SetHost(v bool) *MemberModelMapValue {
	s.Host = &v
	return s
}

func (s *MemberModelMapValue) SetCoHost(v bool) *MemberModelMapValue {
	s.CoHost = &v
	return s
}

func (s *MemberModelMapValue) SetOuterOrgMember(v bool) *MemberModelMapValue {
	s.OuterOrgMember = &v
	return s
}

func (s *MemberModelMapValue) SetPstnJoin(v bool) *MemberModelMapValue {
	s.PstnJoin = &v
	return s
}

func (s *MemberModelMapValue) SetDeviceType(v string) *MemberModelMapValue {
	s.DeviceType = &v
	return s
}

type AddRecordPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRecordPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPermissionHeaders) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionHeaders) SetCommonHeaders(v map[string]*string) *AddRecordPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRecordPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *AddRecordPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRecordPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_record
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFFXXXXXXX
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRGXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddRecordPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionRequest) SetBizType(v string) *AddRecordPermissionRequest {
	s.BizType = &v
	return s
}

func (s *AddRecordPermissionRequest) SetOwnerUnionId(v string) *AddRecordPermissionRequest {
	s.OwnerUnionId = &v
	return s
}

func (s *AddRecordPermissionRequest) SetUnionId(v string) *AddRecordPermissionRequest {
	s.UnionId = &v
	return s
}

type AddRecordPermissionResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s AddRecordPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionResponseBody) SetCode(v string) *AddRecordPermissionResponseBody {
	s.Code = &v
	return s
}

type AddRecordPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecordPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRecordPermissionResponse) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionResponse) SetHeaders(v map[string]*string) *AddRecordPermissionResponse {
	s.Headers = v
	return s
}

func (s *AddRecordPermissionResponse) SetStatusCode(v int32) *AddRecordPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordPermissionResponse) SetBody(v *AddRecordPermissionResponseBody) *AddRecordPermissionResponse {
	s.Body = v
	return s
}

type CancelScheduleConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelScheduleConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *CancelScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelScheduleConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CancelScheduleConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxx
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
}

func (s CancelScheduleConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceRequest) SetCreatorUnionId(v string) *CancelScheduleConferenceRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CancelScheduleConferenceRequest) SetScheduleConferenceId(v string) *CancelScheduleConferenceRequest {
	s.ScheduleConferenceId = &v
	return s
}

type CancelScheduleConferenceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelScheduleConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceResponseBody) SetSuccess(v bool) *CancelScheduleConferenceResponseBody {
	s.Success = &v
	return s
}

type CancelScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelScheduleConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceResponse) SetHeaders(v map[string]*string) *CancelScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *CancelScheduleConferenceResponse) SetStatusCode(v int32) *CancelScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelScheduleConferenceResponse) SetBody(v *CancelScheduleConferenceResponseBody) *CancelScheduleConferenceResponse {
	s.Body = v
	return s
}

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
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
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
	// This parameter is required.
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// add
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxxxxx
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CohostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateCustomShortLinkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomShortLinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomShortLinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomShortLinkHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomShortLinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomShortLinkHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomShortLinkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomShortLinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// COOLAPP-0-1026633886192127xxxB000W
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// example:
	//
	// bizData
	ExtensionAppBizData *string `json:"extensionAppBizData,omitempty" xml:"extensionAppBizData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f6fb627e-a7e8-403e-b1f8-26e85450f4a9
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// true：使用 false：不使用
	UseExtensionApp *bool `json:"useExtensionApp,omitempty" xml:"useExtensionApp,omitempty"`
}

func (s CreateCustomShortLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomShortLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomShortLinkRequest) SetCoolAppCode(v string) *CreateCustomShortLinkRequest {
	s.CoolAppCode = &v
	return s
}

func (s *CreateCustomShortLinkRequest) SetCreatorUnionId(v string) *CreateCustomShortLinkRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateCustomShortLinkRequest) SetExtensionAppBizData(v string) *CreateCustomShortLinkRequest {
	s.ExtensionAppBizData = &v
	return s
}

func (s *CreateCustomShortLinkRequest) SetScheduleConferenceId(v string) *CreateCustomShortLinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *CreateCustomShortLinkRequest) SetUseExtensionApp(v bool) *CreateCustomShortLinkRequest {
	s.UseExtensionApp = &v
	return s
}

type CreateCustomShortLinkResponseBody struct {
	Result *CreateCustomShortLinkResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateCustomShortLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomShortLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomShortLinkResponseBody) SetResult(v *CreateCustomShortLinkResponseBodyResult) *CreateCustomShortLinkResponseBody {
	s.Result = v
	return s
}

type CreateCustomShortLinkResponseBodyResult struct {
	CustomShortLink *string `json:"customShortLink,omitempty" xml:"customShortLink,omitempty"`
}

func (s CreateCustomShortLinkResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomShortLinkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCustomShortLinkResponseBodyResult) SetCustomShortLink(v string) *CreateCustomShortLinkResponseBodyResult {
	s.CustomShortLink = &v
	return s
}

type CreateCustomShortLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomShortLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomShortLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomShortLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomShortLinkResponse) SetHeaders(v map[string]*string) *CreateCustomShortLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomShortLinkResponse) SetStatusCode(v int32) *CreateCustomShortLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomShortLinkResponse) SetBody(v *CreateCustomShortLinkResponseBody) *CreateCustomShortLinkResponse {
	s.Body = v
	return s
}

type CreateScheduleConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateScheduleConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduleConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateScheduleConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxx
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000
	EndTime                  *int64                                                   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ScheduleConfSettingModel *CreateScheduleConferenceRequestScheduleConfSettingModel `json:"scheduleConfSettingModel,omitempty" xml:"scheduleConfSettingModel,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1687924800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 预约会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateScheduleConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequest) SetCreatorUnionId(v string) *CreateScheduleConferenceRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateScheduleConferenceRequest) SetEndTime(v int64) *CreateScheduleConferenceRequest {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleConferenceRequest) SetScheduleConfSettingModel(v *CreateScheduleConferenceRequestScheduleConfSettingModel) *CreateScheduleConferenceRequest {
	s.ScheduleConfSettingModel = v
	return s
}

func (s *CreateScheduleConferenceRequest) SetStartTime(v int64) *CreateScheduleConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *CreateScheduleConferenceRequest) SetTitle(v string) *CreateScheduleConferenceRequest {
	s.Title = &v
	return s
}

type CreateScheduleConferenceRequestScheduleConfSettingModel struct {
	CohostUnionIds []*string `json:"cohostUnionIds,omitempty" xml:"cohostUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// dingc02f685fa06381c44ac5d6980864d335
	ConfAllowedCorpId *string `json:"confAllowedCorpId,omitempty" xml:"confAllowedCorpId,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	HostUnionId *string `json:"hostUnionId,omitempty" xml:"hostUnionId,omitempty"`
	// example:
	//
	// 0：取消锁定 1：锁定
	LockRoom                    *int32                                                                              `json:"lockRoom,omitempty" xml:"lockRoom,omitempty"`
	MoziConfOpenRecordSetting   *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting   `json:"moziConfOpenRecordSetting,omitempty" xml:"moziConfOpenRecordSetting,omitempty" type:"Struct"`
	MoziConfVirtualExtraSetting *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting `json:"moziConfVirtualExtraSetting,omitempty" xml:"moziConfVirtualExtraSetting,omitempty" type:"Struct"`
	// example:
	//
	// -1：开启 0：未开启 6：超过6人自动开启静音
	MuteOnJoin *int32 `json:"muteOnJoin,omitempty" xml:"muteOnJoin,omitempty"`
	// example:
	//
	// 0：允许共享 1：禁止共享
	ScreenShareForbidden *int32 `json:"screenShareForbidden,omitempty" xml:"screenShareForbidden,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModel) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModel) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetCohostUnionIds(v []*string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.CohostUnionIds = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetConfAllowedCorpId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.ConfAllowedCorpId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetHostUnionId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.HostUnionId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetLockRoom(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.LockRoom = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMoziConfOpenRecordSetting(v *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MoziConfOpenRecordSetting = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMoziConfVirtualExtraSetting(v *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MoziConfVirtualExtraSetting = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMuteOnJoin(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MuteOnJoin = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetScreenShareForbidden(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.ScreenShareForbidden = &v
	return s
}

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting struct {
	// example:
	//
	// true：跟随 false：不跟随
	IsFollowHost *bool `json:"isFollowHost,omitempty" xml:"isFollowHost,omitempty"`
	// example:
	//
	// grid：宫格模式,默认9宫格(3x3) speech：演讲者模式 full_screen：全屏模式 auto_grid：自动宫格模式，默认最大4x4宫格 screen_cast：屏幕共享模式，仅放置屏幕共享流 p2p：双人通话模式 full_screen_and_speaker：共享内容+发言人模式
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0：不自动开启 1：自动开启
	RecordAutoStart *int32 `json:"recordAutoStart,omitempty" xml:"recordAutoStart,omitempty"`
	// example:
	//
	// 0：我以主持人身份入会后自动开启 1：其他人以联席主持人身份入会后开启 2：任何人以任何身份入会后开启
	RecordAutoStartType *int32 `json:"recordAutoStartType,omitempty" xml:"recordAutoStartType,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetIsFollowHost(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.IsFollowHost = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetMode(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.Mode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStart(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStart = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStartType(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStartType = &v
	return s
}

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting struct {
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CloudRecordOwnerUnionId *string `json:"cloudRecordOwnerUnionId,omitempty" xml:"cloudRecordOwnerUnionId,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	EnableChat *int32 `json:"enableChat,omitempty" xml:"enableChat,omitempty"`
	// example:
	//
	// true：允许匿名登录入会 false：不允许匿名登录入会
	EnableWebAnonymousJoin *bool `json:"enableWebAnonymousJoin,omitempty" xml:"enableWebAnonymousJoin,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	JoinBeforeHost *int32 `json:"joinBeforeHost,omitempty" xml:"joinBeforeHost,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	LockMediaStatusMicMute *int32 `json:"lockMediaStatusMicMute,omitempty" xml:"lockMediaStatusMicMute,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	LockNick *int32 `json:"lockNick,omitempty" xml:"lockNick,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	MinutesOwnerUnionId          *string                                                                                                           `json:"minutesOwnerUnionId,omitempty" xml:"minutesOwnerUnionId,omitempty"`
	MoziConfExtensionAppSettings []*CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings `json:"moziConfExtensionAppSettings,omitempty" xml:"moziConfExtensionAppSettings,omitempty" type:"Repeated"`
	PushAllMeetingRecords        *bool                                                                                                             `json:"pushAllMeetingRecords,omitempty" xml:"pushAllMeetingRecords,omitempty"`
	PushCloudRecordCard          *bool                                                                                                             `json:"pushCloudRecordCard,omitempty" xml:"pushCloudRecordCard,omitempty"`
	PushMinutesCard              *bool                                                                                                             `json:"pushMinutesCard,omitempty" xml:"pushMinutesCard,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	WaitingRoom *int32 `json:"waitingRoom,omitempty" xml:"waitingRoom,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetCloudRecordOwnerUnionId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.CloudRecordOwnerUnionId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableChat(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableChat = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableWebAnonymousJoin(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableWebAnonymousJoin = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetJoinBeforeHost(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.JoinBeforeHost = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockMediaStatusMicMute(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockMediaStatusMicMute = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockNick(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockNick = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMinutesOwnerUnionId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MinutesOwnerUnionId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMoziConfExtensionAppSettings(v []*CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MoziConfExtensionAppSettings = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushAllMeetingRecords(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushAllMeetingRecords = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushCloudRecordCard(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushCloudRecordCard = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushMinutesCard(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushMinutesCard = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetWaitingRoom(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.WaitingRoom = &v
	return s
}

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings struct {
	// example:
	//
	// 0：不自动打开 1：仅主持人/联席主持人自动打开 2：全员自动打开
	AutoOpenMode *int32 `json:"autoOpenMode,omitempty" xml:"autoOpenMode,omitempty"`
	// example:
	//
	// COOLAPP-0-1026633886192127xxxB000W
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// example:
	//
	// bizData
	ExtensionAppBizData *string `json:"extensionAppBizData,omitempty" xml:"extensionAppBizData,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetAutoOpenMode(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetCoolAppCode(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.CoolAppCode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetExtensionAppBizData(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

type CreateScheduleConferenceResponseBody struct {
	Phones    []*string `json:"phones,omitempty" xml:"phones,omitempty" type:"Repeated"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 838 722 xxxxx
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// example:
	//
	// 2a489c68-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/Bsbp3ixxxxxUyJJ9
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateScheduleConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceResponseBody) SetPhones(v []*string) *CreateScheduleConferenceResponseBody {
	s.Phones = v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetRequestId(v string) *CreateScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetRoomCode(v string) *CreateScheduleConferenceResponseBody {
	s.RoomCode = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetScheduleConferenceId(v string) *CreateScheduleConferenceResponseBody {
	s.ScheduleConferenceId = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetUrl(v string) *CreateScheduleConferenceResponseBody {
	s.Url = &v
	return s
}

type CreateScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduleConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceResponse) SetHeaders(v map[string]*string) *CreateScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleConferenceResponse) SetStatusCode(v int32) *CreateScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleConferenceResponse) SetBody(v *CreateScheduleConferenceResponseBody) *CreateScheduleConferenceResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// XXX的视频会议
	ConfTitle *string `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	InviteCaller *bool `json:"inviteCaller,omitempty" xml:"inviteCaller,omitempty"`
	// if can be null:
	// true
	InviteUserIds []*string `json:"inviteUserIds,omitempty" xml:"inviteUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FocusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GenerateFlashMinutesDocumentUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GenerateFlashMinutesDocumentUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GenerateFlashMinutesDocumentUrlHeaders) GoString() string {
	return s.String()
}

func (s *GenerateFlashMinutesDocumentUrlHeaders) SetCommonHeaders(v map[string]*string) *GenerateFlashMinutesDocumentUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateFlashMinutesDocumentUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GenerateFlashMinutesDocumentUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GenerateFlashMinutesDocumentUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_record
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// if can be null:
	// false
	//
	// example:
	//
	// 1727185971000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRG9KXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GenerateFlashMinutesDocumentUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateFlashMinutesDocumentUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateFlashMinutesDocumentUrlRequest) SetBizType(v string) *GenerateFlashMinutesDocumentUrlRequest {
	s.BizType = &v
	return s
}

func (s *GenerateFlashMinutesDocumentUrlRequest) SetExpireTime(v int64) *GenerateFlashMinutesDocumentUrlRequest {
	s.ExpireTime = &v
	return s
}

func (s *GenerateFlashMinutesDocumentUrlRequest) SetUnionId(v string) *GenerateFlashMinutesDocumentUrlRequest {
	s.UnionId = &v
	return s
}

type GenerateFlashMinutesDocumentUrlResponseBody struct {
	MinutesDocUrl *string `json:"minutesDocUrl,omitempty" xml:"minutesDocUrl,omitempty"`
}

func (s GenerateFlashMinutesDocumentUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateFlashMinutesDocumentUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFlashMinutesDocumentUrlResponseBody) SetMinutesDocUrl(v string) *GenerateFlashMinutesDocumentUrlResponseBody {
	s.MinutesDocUrl = &v
	return s
}

type GenerateFlashMinutesDocumentUrlResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFlashMinutesDocumentUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFlashMinutesDocumentUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateFlashMinutesDocumentUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateFlashMinutesDocumentUrlResponse) SetHeaders(v map[string]*string) *GenerateFlashMinutesDocumentUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateFlashMinutesDocumentUrlResponse) SetStatusCode(v int32) *GenerateFlashMinutesDocumentUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFlashMinutesDocumentUrlResponse) SetBody(v *GenerateFlashMinutesDocumentUrlResponseBody) *GenerateFlashMinutesDocumentUrlResponse {
	s.Body = v
	return s
}

type GetConfDataByConferenceIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConfDataByConferenceIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConfDataByConferenceIdHeaders) GoString() string {
	return s.String()
}

func (s *GetConfDataByConferenceIdHeaders) SetCommonHeaders(v map[string]*string) *GetConfDataByConferenceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConfDataByConferenceIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetConfDataByConferenceIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConfDataByConferenceIdRequest struct {
	// This parameter is required.
	RealData *bool `json:"realData,omitempty" xml:"realData,omitempty"`
}

func (s GetConfDataByConferenceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfDataByConferenceIdRequest) GoString() string {
	return s.String()
}

func (s *GetConfDataByConferenceIdRequest) SetRealData(v bool) *GetConfDataByConferenceIdRequest {
	s.RealData = &v
	return s
}

type GetConfDataByConferenceIdResponseBody struct {
	// example:
	//
	// 6449d8a6414xxxxxxxx01464af9f0
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// njMTqKo9xxxxEiE
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// xxxx
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 1682561790000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 0
	FreeType *string `json:"freeType,omitempty" xml:"freeType,omitempty"`
	// example:
	//
	// ding_talk
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1682561190000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 600000
	TimeLength *int64 `json:"timeLength,omitempty" xml:"timeLength,omitempty"`
	// example:
	//
	// xxxxx测试会议
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
}

func (s GetConfDataByConferenceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfDataByConferenceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfDataByConferenceIdResponseBody) SetConferenceId(v string) *GetConfDataByConferenceIdResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetCreatorId(v string) *GetConfDataByConferenceIdResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetCreatorNick(v string) *GetConfDataByConferenceIdResponseBody {
	s.CreatorNick = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetDeptName(v string) *GetConfDataByConferenceIdResponseBody {
	s.DeptName = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetEndTime(v int64) *GetConfDataByConferenceIdResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetFreeType(v string) *GetConfDataByConferenceIdResponseBody {
	s.FreeType = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetScene(v string) *GetConfDataByConferenceIdResponseBody {
	s.Scene = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetStartTime(v int64) *GetConfDataByConferenceIdResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetTimeLength(v int64) *GetConfDataByConferenceIdResponseBody {
	s.TimeLength = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetTitle(v string) *GetConfDataByConferenceIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetConfDataByConferenceIdResponseBody) SetUserCount(v int32) *GetConfDataByConferenceIdResponseBody {
	s.UserCount = &v
	return s
}

type GetConfDataByConferenceIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfDataByConferenceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfDataByConferenceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfDataByConferenceIdResponse) GoString() string {
	return s.String()
}

func (s *GetConfDataByConferenceIdResponse) SetHeaders(v map[string]*string) *GetConfDataByConferenceIdResponse {
	s.Headers = v
	return s
}

func (s *GetConfDataByConferenceIdResponse) SetStatusCode(v int32) *GetConfDataByConferenceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfDataByConferenceIdResponse) SetBody(v *GetConfDataByConferenceIdResponseBody) *GetConfDataByConferenceIdResponse {
	s.Body = v
	return s
}

type GetConfDetailDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConfDetailDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConfDetailDataHeaders) GoString() string {
	return s.String()
}

func (s *GetConfDetailDataHeaders) SetCommonHeaders(v map[string]*string) *GetConfDetailDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConfDetailDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetConfDetailDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConfDetailDataRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxx9uZ0bVGoqjxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 张三
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
}

func (s GetConfDetailDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfDetailDataRequest) GoString() string {
	return s.String()
}

func (s *GetConfDetailDataRequest) SetMaxResults(v int32) *GetConfDetailDataRequest {
	s.MaxResults = &v
	return s
}

func (s *GetConfDetailDataRequest) SetNextToken(v string) *GetConfDetailDataRequest {
	s.NextToken = &v
	return s
}

func (s *GetConfDetailDataRequest) SetNick(v string) *GetConfDetailDataRequest {
	s.Nick = &v
	return s
}

type GetConfDetailDataResponseBody struct {
	List []*GetConfDetailDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxZ0bVGoqxxBGQbxdxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetConfDetailDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfDetailDataResponseBody) SetList(v []*GetConfDetailDataResponseBodyList) *GetConfDetailDataResponseBody {
	s.List = v
	return s
}

func (s *GetConfDetailDataResponseBody) SetNextToken(v string) *GetConfDetailDataResponseBody {
	s.NextToken = &v
	return s
}

type GetConfDetailDataResponseBodyList struct {
	// example:
	//
	// true
	BelongOrg *string `json:"belongOrg,omitempty" xml:"belongOrg,omitempty"`
	// example:
	//
	// 6449d8a6414xxxxxxxx01464af9f0
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// Mac
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// example:
	//
	// 974000
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1682561199000
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// example:
	//
	// 1682562173000
	LeaveTime *int64 `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	// example:
	//
	// -1
	NetworkQuality *string `json:"networkQuality,omitempty" xml:"networkQuality,omitempty"`
	// example:
	//
	// 张三
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// 参会人
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 已离开
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// njMTqKo9xxxxEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 6.1.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetConfDetailDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetConfDetailDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetConfDetailDataResponseBodyList) SetBelongOrg(v string) *GetConfDetailDataResponseBodyList {
	s.BelongOrg = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetConferenceId(v string) *GetConfDetailDataResponseBodyList {
	s.ConferenceId = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetDeviceType(v string) *GetConfDetailDataResponseBodyList {
	s.DeviceType = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetDuration(v int64) *GetConfDetailDataResponseBodyList {
	s.Duration = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetJoinTime(v int64) *GetConfDetailDataResponseBodyList {
	s.JoinTime = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetLeaveTime(v int64) *GetConfDetailDataResponseBodyList {
	s.LeaveTime = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetNetworkQuality(v string) *GetConfDetailDataResponseBodyList {
	s.NetworkQuality = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetNick(v string) *GetConfDetailDataResponseBodyList {
	s.Nick = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetRole(v string) *GetConfDetailDataResponseBodyList {
	s.Role = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetSessionId(v string) *GetConfDetailDataResponseBodyList {
	s.SessionId = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetStatus(v string) *GetConfDetailDataResponseBodyList {
	s.Status = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetUnionId(v string) *GetConfDetailDataResponseBodyList {
	s.UnionId = &v
	return s
}

func (s *GetConfDetailDataResponseBodyList) SetVersion(v string) *GetConfDetailDataResponseBodyList {
	s.Version = &v
	return s
}

type GetConfDetailDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfDetailDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfDetailDataResponse) GoString() string {
	return s.String()
}

func (s *GetConfDetailDataResponse) SetHeaders(v map[string]*string) *GetConfDetailDataResponse {
	s.Headers = v
	return s
}

func (s *GetConfDetailDataResponse) SetStatusCode(v int32) *GetConfDetailDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfDetailDataResponse) SetBody(v *GetConfDetailDataResponseBody) *GetConfDetailDataResponse {
	s.Body = v
	return s
}

type GetHistoryConfDataListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetHistoryConfDataListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryConfDataListHeaders) GoString() string {
	return s.String()
}

func (s *GetHistoryConfDataListHeaders) SetCommonHeaders(v map[string]*string) *GetHistoryConfDataListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHistoryConfDataListHeaders) SetXAcsDingtalkAccessToken(v string) *GetHistoryConfDataListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetHistoryConfDataListRequest struct {
	// example:
	//
	// 张三
	CreatorNike *string `json:"creatorNike,omitempty" xml:"creatorNike,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1682611199000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 0
	FreeType *string `json:"freeType,omitempty" xml:"freeType,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxx9uZ0bVGoqjxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	RealData *bool `json:"realData,omitempty" xml:"realData,omitempty"`
	// example:
	//
	// ding_talk
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1682524800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// xxxxx视频会议
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetHistoryConfDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryConfDataListRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryConfDataListRequest) SetCreatorNike(v string) *GetHistoryConfDataListRequest {
	s.CreatorNike = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetEndTime(v int64) *GetHistoryConfDataListRequest {
	s.EndTime = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetFreeType(v string) *GetHistoryConfDataListRequest {
	s.FreeType = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetMaxResults(v int32) *GetHistoryConfDataListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetNextToken(v string) *GetHistoryConfDataListRequest {
	s.NextToken = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetRealData(v bool) *GetHistoryConfDataListRequest {
	s.RealData = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetScene(v string) *GetHistoryConfDataListRequest {
	s.Scene = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetStartTime(v int64) *GetHistoryConfDataListRequest {
	s.StartTime = &v
	return s
}

func (s *GetHistoryConfDataListRequest) SetTitle(v string) *GetHistoryConfDataListRequest {
	s.Title = &v
	return s
}

type GetHistoryConfDataListResponseBody struct {
	List []*GetHistoryConfDataListResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// xxx9uZ0bVGoqjxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetHistoryConfDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryConfDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryConfDataListResponseBody) SetList(v []*GetHistoryConfDataListResponseBodyList) *GetHistoryConfDataListResponseBody {
	s.List = v
	return s
}

func (s *GetHistoryConfDataListResponseBody) SetNextToken(v string) *GetHistoryConfDataListResponseBody {
	s.NextToken = &v
	return s
}

type GetHistoryConfDataListResponseBodyList struct {
	// example:
	//
	// 6449d8a6414xxxxxxxx01464af9f0
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// njMTqKo9xxxxEiE
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 张三
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// example:
	//
	// xxxxx
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 1682561790000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 0
	FreeType *string `json:"freeType,omitempty" xml:"freeType,omitempty"`
	// example:
	//
	// ding_talk
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1682561190000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 600000
	TimeLength *int64 `json:"timeLength,omitempty" xml:"timeLength,omitempty"`
	// example:
	//
	// xxxxx视频会议
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
}

func (s GetHistoryConfDataListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryConfDataListResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetHistoryConfDataListResponseBodyList) SetConferenceId(v string) *GetHistoryConfDataListResponseBodyList {
	s.ConferenceId = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetCreatorId(v string) *GetHistoryConfDataListResponseBodyList {
	s.CreatorId = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetCreatorNick(v string) *GetHistoryConfDataListResponseBodyList {
	s.CreatorNick = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetDeptName(v string) *GetHistoryConfDataListResponseBodyList {
	s.DeptName = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetEndTime(v int64) *GetHistoryConfDataListResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetFreeType(v string) *GetHistoryConfDataListResponseBodyList {
	s.FreeType = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetScene(v string) *GetHistoryConfDataListResponseBodyList {
	s.Scene = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetStartTime(v int64) *GetHistoryConfDataListResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetTimeLength(v int64) *GetHistoryConfDataListResponseBodyList {
	s.TimeLength = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetTitle(v string) *GetHistoryConfDataListResponseBodyList {
	s.Title = &v
	return s
}

func (s *GetHistoryConfDataListResponseBodyList) SetUserCount(v int32) *GetHistoryConfDataListResponseBodyList {
	s.UserCount = &v
	return s
}

type GetHistoryConfDataListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryConfDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryConfDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryConfDataListResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryConfDataListResponse) SetHeaders(v map[string]*string) *GetHistoryConfDataListResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryConfDataListResponse) SetStatusCode(v int32) *GetHistoryConfDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryConfDataListResponse) SetBody(v *GetHistoryConfDataListResponseBody) *GetHistoryConfDataListResponse {
	s.Body = v
	return s
}

type GetUserLastMetricHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserLastMetricHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserLastMetricHeaders) GoString() string {
	return s.String()
}

func (s *GetUserLastMetricHeaders) SetCommonHeaders(v map[string]*string) *GetUserLastMetricHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserLastMetricHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserLastMetricHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserLastMetricRequest struct {
	// This parameter is required.
	UnionIdList []*string `json:"unionIdList,omitempty" xml:"unionIdList,omitempty" type:"Repeated"`
}

func (s GetUserLastMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserLastMetricRequest) GoString() string {
	return s.String()
}

func (s *GetUserLastMetricRequest) SetUnionIdList(v []*string) *GetUserLastMetricRequest {
	s.UnionIdList = v
	return s
}

type GetUserLastMetricResponseBody struct {
	MetricMap map[string]*MetricMapValue `json:"metricMap,omitempty" xml:"metricMap,omitempty"`
}

func (s GetUserLastMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserLastMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLastMetricResponseBody) SetMetricMap(v map[string]*MetricMapValue) *GetUserLastMetricResponseBody {
	s.MetricMap = v
	return s
}

type GetUserLastMetricResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserLastMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserLastMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserLastMetricResponse) GoString() string {
	return s.String()
}

func (s *GetUserLastMetricResponse) SetHeaders(v map[string]*string) *GetUserLastMetricResponse {
	s.Headers = v
	return s
}

func (s *GetUserLastMetricResponse) SetStatusCode(v int32) *GetUserLastMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLastMetricResponse) SetBody(v *GetUserLastMetricResponseBody) *GetUserLastMetricResponse {
	s.Body = v
	return s
}

type GetUserMetricDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserMetricDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserMetricDataHeaders) GoString() string {
	return s.String()
}

func (s *GetUserMetricDataHeaders) SetCommonHeaders(v map[string]*string) *GetUserMetricDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserMetricDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserMetricDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserMetricDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1682524800000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1682611199000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// njMTqKo9iiyxxxxxxxxiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMetricDataRequest) GoString() string {
	return s.String()
}

func (s *GetUserMetricDataRequest) SetBeginTime(v int64) *GetUserMetricDataRequest {
	s.BeginTime = &v
	return s
}

func (s *GetUserMetricDataRequest) SetEndTime(v int64) *GetUserMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetUserMetricDataRequest) SetUnionId(v string) *GetUserMetricDataRequest {
	s.UnionId = &v
	return s
}

type GetUserMetricDataResponseBody struct {
	MetricDataList []*GetUserMetricDataResponseBodyMetricDataList `json:"metricDataList,omitempty" xml:"metricDataList,omitempty" type:"Repeated"`
}

func (s GetUserMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMetricDataResponseBody) SetMetricDataList(v []*GetUserMetricDataResponseBodyMetricDataList) *GetUserMetricDataResponseBody {
	s.MetricDataList = v
	return s
}

type GetUserMetricDataResponseBodyMetricDataList struct {
	// example:
	//
	// 27
	AudioPlayLevel *string `json:"audioPlayLevel,omitempty" xml:"audioPlayLevel,omitempty"`
	// example:
	//
	// 25
	AudioRecLevel *string `json:"audioRecLevel,omitempty" xml:"audioRecLevel,omitempty"`
	// example:
	//
	// 52939
	AudioRecvBitRate *string `json:"audioRecvBitRate,omitempty" xml:"audioRecvBitRate,omitempty"`
	// example:
	//
	// 59103
	AudioSendBitRate *string `json:"audioSendBitRate,omitempty" xml:"audioSendBitRate,omitempty"`
	// example:
	//
	// 66160
	CameraRecvBitRate *string `json:"cameraRecvBitRate,omitempty" xml:"cameraRecvBitRate,omitempty"`
	// example:
	//
	// 20
	CameraRecvFrame *string `json:"cameraRecvFrame,omitempty" xml:"cameraRecvFrame,omitempty"`
	// example:
	//
	// 1920*1080
	CameraRecvResolutionActual *string `json:"cameraRecvResolutionActual,omitempty" xml:"cameraRecvResolutionActual,omitempty"`
	// example:
	//
	// 1145172
	CameraSendBitRate *string `json:"cameraSendBitRate,omitempty" xml:"cameraSendBitRate,omitempty"`
	// example:
	//
	// 20
	CameraSendFrame *string `json:"cameraSendFrame,omitempty" xml:"cameraSendFrame,omitempty"`
	// example:
	//
	// 1920*1080
	CameraSendResolutionActual *string `json:"cameraSendResolutionActual,omitempty" xml:"cameraSendResolutionActual,omitempty"`
	// example:
	//
	// 0
	LostRate *string `json:"lostRate,omitempty" xml:"lostRate,omitempty"`
	// example:
	//
	// 66160
	RecvBitRate *string `json:"recvBitRate,omitempty" xml:"recvBitRate,omitempty"`
	// example:
	//
	// 20
	RoundTripTime *string `json:"roundTripTime,omitempty" xml:"roundTripTime,omitempty"`
	// example:
	//
	// 0
	ScreenRecvBitRate *string `json:"screenRecvBitRate,omitempty" xml:"screenRecvBitRate,omitempty"`
	// example:
	//
	// 0
	ScreenRecvFrame *string `json:"screenRecvFrame,omitempty" xml:"screenRecvFrame,omitempty"`
	// example:
	//
	// 1920*1080
	ScreenRecvResolutionActual *string `json:"screenRecvResolutionActual,omitempty" xml:"screenRecvResolutionActual,omitempty"`
	// example:
	//
	// 15701
	ScreenSendBitRate *string `json:"screenSendBitRate,omitempty" xml:"screenSendBitRate,omitempty"`
	// example:
	//
	// 14
	ScreenSendFrame *string `json:"screenSendFrame,omitempty" xml:"screenSendFrame,omitempty"`
	// example:
	//
	// 1920*1080
	ScreenSendResolutionActual *string `json:"screenSendResolutionActual,omitempty" xml:"screenSendResolutionActual,omitempty"`
	// example:
	//
	// 1145172
	SendBitRate *string `json:"sendBitRate,omitempty" xml:"sendBitRate,omitempty"`
	// example:
	//
	// 1682562120000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetUserMetricDataResponseBodyMetricDataList) String() string {
	return tea.Prettify(s)
}

func (s GetUserMetricDataResponseBodyMetricDataList) GoString() string {
	return s.String()
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetAudioPlayLevel(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.AudioPlayLevel = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetAudioRecLevel(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.AudioRecLevel = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetAudioRecvBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.AudioRecvBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetAudioSendBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.AudioSendBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraRecvBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraRecvBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraRecvFrame(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraRecvFrame = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraRecvResolutionActual(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraRecvResolutionActual = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraSendBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraSendBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraSendFrame(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraSendFrame = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetCameraSendResolutionActual(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.CameraSendResolutionActual = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetLostRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.LostRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetRecvBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.RecvBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetRoundTripTime(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.RoundTripTime = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenRecvBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenRecvBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenRecvFrame(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenRecvFrame = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenRecvResolutionActual(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenRecvResolutionActual = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenSendBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenSendBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenSendFrame(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenSendFrame = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetScreenSendResolutionActual(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.ScreenSendResolutionActual = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetSendBitRate(v string) *GetUserMetricDataResponseBodyMetricDataList {
	s.SendBitRate = &v
	return s
}

func (s *GetUserMetricDataResponseBodyMetricDataList) SetTimestamp(v int64) *GetUserMetricDataResponseBodyMetricDataList {
	s.Timestamp = &v
	return s
}

type GetUserMetricDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMetricDataResponse) GoString() string {
	return s.String()
}

func (s *GetUserMetricDataResponse) SetHeaders(v map[string]*string) *GetUserMetricDataResponse {
	s.Headers = v
	return s
}

func (s *GetUserMetricDataResponse) SetStatusCode(v int32) *GetUserMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMetricDataResponse) SetBody(v *GetUserMetricDataResponseBody) *GetUserMetricDataResponse {
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
	InviteeList      []*InviteUsersRequestInviteeList      `json:"inviteeList,omitempty" xml:"inviteeList,omitempty" type:"Repeated"`
	PhoneInviteeList []*InviteUsersRequestPhoneInviteeList `json:"phoneInviteeList,omitempty" xml:"phoneInviteeList,omitempty" type:"Repeated"`
	// example:
	//
	// qzR1iSMDvzR9iPXXXXXXXXXXXXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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

func (s *InviteUsersRequest) SetPhoneInviteeList(v []*InviteUsersRequestPhoneInviteeList) *InviteUsersRequest {
	s.PhoneInviteeList = v
	return s
}

func (s *InviteUsersRequest) SetUnionId(v string) *InviteUsersRequest {
	s.UnionId = &v
	return s
}

type InviteUsersRequestInviteeList struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试用户
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// qzR1iSMDvzR9kXXXXXXXx
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

type InviteUsersRequestPhoneInviteeList struct {
	// example:
	//
	// false
	InviteClient *bool `json:"inviteClient,omitempty" xml:"inviteClient,omitempty"`
	// example:
	//
	// 测试电话用户
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// 1xxxxxxxxxx9
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// example:
	//
	// 86
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InviteUsersRequestPhoneInviteeList) String() string {
	return tea.Prettify(s)
}

func (s InviteUsersRequestPhoneInviteeList) GoString() string {
	return s.String()
}

func (s *InviteUsersRequestPhoneInviteeList) SetInviteClient(v bool) *InviteUsersRequestPhoneInviteeList {
	s.InviteClient = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetNick(v string) *InviteUsersRequestPhoneInviteeList {
	s.Nick = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetPhoneNumber(v string) *InviteUsersRequestPhoneInviteeList {
	s.PhoneNumber = &v
	return s
}

func (s *InviteUsersRequestPhoneInviteeList) SetStatusCode(v string) *InviteUsersRequestPhoneInviteeList {
	s.StatusCode = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	ForbiddenRejoin *bool `json:"forbiddenRejoin,omitempty" xml:"forbiddenRejoin,omitempty"`
	// This parameter is required.
	UserList []*KickMembersRequestUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
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
	// example:
	//
	// 644272f14ba3311xxxxxxxxx
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxxxxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KickMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type LockConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LockConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s LockConferenceHeaders) GoString() string {
	return s.String()
}

func (s *LockConferenceHeaders) SetCommonHeaders(v map[string]*string) *LockConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LockConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *LockConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LockConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// lock
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s LockConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockConferenceRequest) GoString() string {
	return s.String()
}

func (s *LockConferenceRequest) SetAction(v string) *LockConferenceRequest {
	s.Action = &v
	return s
}

type LockConferenceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LockConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *LockConferenceResponseBody) SetSuccess(v bool) *LockConferenceResponseBody {
	s.Success = &v
	return s
}

type LockConferenceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockConferenceResponse) GoString() string {
	return s.String()
}

func (s *LockConferenceResponse) SetHeaders(v map[string]*string) *LockConferenceResponse {
	s.Headers = v
	return s
}

func (s *LockConferenceResponse) SetStatusCode(v int32) *LockConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockConferenceResponse) SetBody(v *LockConferenceResponseBody) *LockConferenceResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// mute
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// mute
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
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
	// example:
	//
	// 644272f14ba3311xxxxxxxxx
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxxxxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1
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
	// example:
	//
	// true
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
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1631172045153000
	NextTtoken *int64 `json:"nextTtoken,omitempty" xml:"nextTtoken,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 嘿！你好，这里是小钉
	Paragraph *string `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	// example:
	//
	// 44444
	RecordId     *int64                                                       `json:"recordId,omitempty" xml:"recordId,omitempty"`
	SentenceList []*QueryCloudRecordTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// WFBkgJvt0xxxxSaA1jK4sgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 这里是小钉
	Sentence *string `json:"sentence,omitempty" xml:"sentence,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// WFBkgJvt0xxxxSaA1jK4sgiEiE
	UnionId  *string                                                              `json:"unionId,omitempty" xml:"unionId,omitempty"`
	WordList []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
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
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 这里
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
	// example:
	//
	// 1631172050535000#0
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1631172094000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// faa1566c5bc24f21821ae2394f82db2e
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 290882268xxx1172033231
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// example:
	//
	// 0
	RecordType *int64 `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 1631172094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// WFBkgJvtxxxxtSaA1jK4sgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// https://vod.mcs.dingtalk.com/faa1566c5bc24f21821ae2394f82db2e/8bbd1612e686462ab4717919f67bb721-b8531e0d534b2f9747a9fdfxxxxxxxxc-sd.mp4
	Mp4FileUrl *string `json:"mp4FileUrl,omitempty" xml:"mp4FileUrl,omitempty"`
	// example:
	//
	// https://vod.mcs.dingtalk.com/faa1566c5bc24f21821ae2394f82db2e/8bbd1612e686462ab4717919f67bb721-ab85cc044a163568c9485xxxxxxxx76d-sd.m3u8
	PlayUrl *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	Status  *int64  `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCloudRecordVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 10
	ActiveNum *int32 `json:"activeNum,omitempty" xml:"activeNum,omitempty"`
	// example:
	//
	// 15
	AttendNum               *int32  `json:"attendNum,omitempty" xml:"attendNum,omitempty"`
	BizType                 *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CloudRecordOwnerUnionId *string `json:"cloudRecordOwnerUnionId,omitempty" xml:"cloudRecordOwnerUnionId,omitempty"`
	CloudRecordStatus       *int32  `json:"cloudRecordStatus,omitempty" xml:"cloudRecordStatus,omitempty"`
	// example:
	//
	// 1000000
	ConfDuration *int64 `json:"confDuration,omitempty" xml:"confDuration,omitempty"`
	// example:
	//
	// 6323d7568777190142ab9d10
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 昵称
	CreatorNick          *string                                                        `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	EndTime              *int64                                                         `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExtensionAppSettings []*QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings `json:"extensionAppSettings,omitempty" xml:"extensionAppSettings,omitempty" type:"Repeated"`
	// example:
	//
	// https://meeting.dingtalk.com/app?roomCode=42726033559&token=1_7ac974ac-6e4f-47c3-b82b-bfb32fd94d2c
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	// example:
	//
	// 20
	InvitedNum          *int32  `json:"invitedNum,omitempty" xml:"invitedNum,omitempty"`
	MinutesOwnerUnionId *string `json:"minutesOwnerUnionId,omitempty" xml:"minutesOwnerUnionId,omitempty"`
	MinutesStatus       *int32  `json:"minutesStatus,omitempty" xml:"minutesStatus,omitempty"`
	// example:
	//
	// 42726033559
	RoomCode             *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// 1663293270000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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

func (s *QueryConferenceInfoResponseBodyConfInfo) SetBizType(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.BizType = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCloudRecordOwnerUnionId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CloudRecordOwnerUnionId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCloudRecordStatus(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.CloudRecordStatus = &v
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

func (s *QueryConferenceInfoResponseBodyConfInfo) SetExtensionAppSettings(v []*QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) *QueryConferenceInfoResponseBodyConfInfo {
	s.ExtensionAppSettings = v
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

func (s *QueryConferenceInfoResponseBodyConfInfo) SetMinutesOwnerUnionId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.MinutesOwnerUnionId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetMinutesStatus(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.MinutesStatus = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetRoomCode(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetScheduleConferenceId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ScheduleConferenceId = &v
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

type QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings struct {
	AppCode             *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	AppId               *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AutoOpenMode        *int32  `json:"autoOpenMode,omitempty" xml:"autoOpenMode,omitempty"`
	ExtensionAppBizData *string `json:"extensionAppBizData,omitempty" xml:"extensionAppBizData,omitempty"`
}

func (s QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) SetAppCode(v string) *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings {
	s.AppCode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) SetAppId(v string) *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings {
	s.AppId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) SetAutoOpenMode(v int32) *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings) SetExtensionAppBizData(v string) *QueryConferenceInfoResponseBodyConfInfoExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

type QueryConferenceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 0-正常，1-麦克风静音，2-摄像头关闭，4-强制全员静音
	MediaStatus *int64 `json:"mediaStatus,omitempty" xml:"mediaStatus,omitempty"`
	StartTime   *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 0-初始化，1-会议结束，2-会议开始
	Status   *int64                                               `json:"status,omitempty" xml:"status,omitempty"`
	Title    *string                                              `json:"title,omitempty" xml:"title,omitempty"`
	UserList []*QueryConferenceInfoBatchResponseBodyInfosUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
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
	// example:
	//
	// 0-未定义,1-初始化,2-加入中,3-在会,4-加入失败,5,被踢出,6-离开
	AttendStatus *int64 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	// example:
	//
	// 0-初始化，1-关闭，2-打开
	CameraStatus *int64 `json:"cameraStatus,omitempty" xml:"cameraStatus,omitempty"`
	// example:
	//
	// 0-初始化，1-关闭，2-打开
	MicStatus *int64  `json:"micStatus,omitempty" xml:"micStatus,omitempty"`
	Nick      *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// 抱歉，正在开会
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceInfoBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryConferenceInfoByRoomCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceInfoByRoomCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceInfoByRoomCodeRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetMaxResults(v int32) *QueryConferenceInfoByRoomCodeRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeRequest) SetNextToken(v string) *QueryConferenceInfoByRoomCodeRequest {
	s.NextToken = &v
	return s
}

type QueryConferenceInfoByRoomCodeResponseBody struct {
	ConferenceList []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList `json:"conferenceList,omitempty" xml:"conferenceList,omitempty" type:"Repeated"`
	HasMore        *bool                                                      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken      *string                                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount     *int32                                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetConferenceList(v []*QueryConferenceInfoByRoomCodeResponseBodyConferenceList) *QueryConferenceInfoByRoomCodeResponseBody {
	s.ConferenceList = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetHasMore(v bool) *QueryConferenceInfoByRoomCodeResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetNextToken(v string) *QueryConferenceInfoByRoomCodeResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBody) SetTotalCount(v int32) *QueryConferenceInfoByRoomCodeResponseBody {
	s.TotalCount = &v
	return s
}

type QueryConferenceInfoByRoomCodeResponseBodyConferenceList struct {
	BizType              *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	ConfDuration         *int64  `json:"confDuration,omitempty" xml:"confDuration,omitempty"`
	ConferenceId         *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	CreatorId            *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CreatorNick          *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExternalLinkUrl      *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	RoomCode             *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	StartTime            *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status               *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponseBodyConferenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponseBodyConferenceList) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetBizType(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.BizType = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetConfDuration(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ConfDuration = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetConferenceId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetCreatorId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.CreatorId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetCreatorNick(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.CreatorNick = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetEndTime(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.EndTime = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetExternalLinkUrl(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ExternalLinkUrl = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetRoomCode(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetScheduleConferenceId(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetStartTime(v int64) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetStatus(v int32) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponseBodyConferenceList) SetTitle(v string) *QueryConferenceInfoByRoomCodeResponseBodyConferenceList {
	s.Title = &v
	return s
}

type QueryConferenceInfoByRoomCodeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceInfoByRoomCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetStatusCode(v int32) *QueryConferenceInfoByRoomCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetBody(v *QueryConferenceInfoByRoomCodeResponseBody) *QueryConferenceInfoByRoomCodeResponse {
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// 123000000
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// example:
	//
	// 6
	AttendStatus *int32 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	CoHost       *bool  `json:"coHost,omitempty" xml:"coHost,omitempty"`
	// example:
	//
	// 6323d7562b18000142ab9d10
	ConferenceId   *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	Duration       *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Host           *bool   `json:"host,omitempty" xml:"host,omitempty"`
	JoinTime       *int64  `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	LeaveTime      *int64  `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	OuterOrgMember *bool   `json:"outerOrgMember,omitempty" xml:"outerOrgMember,omitempty"`
	PstnJoin       *bool   `json:"pstnJoin,omitempty" xml:"pstnJoin,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 昵称
	UserNick *string `json:"userNick,omitempty" xml:"userNick,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryFlashMinutesSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFlashMinutesSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryHeaders) SetCommonHeaders(v map[string]*string) *QueryFlashMinutesSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFlashMinutesSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFlashMinutesSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFlashMinutesSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_record
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRG9KXXXX
	RecorderUnionId *string `json:"recorderUnionId,omitempty" xml:"recorderUnionId,omitempty"`
}

func (s QueryFlashMinutesSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryRequest) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryRequest) SetBizType(v string) *QueryFlashMinutesSummaryRequest {
	s.BizType = &v
	return s
}

func (s *QueryFlashMinutesSummaryRequest) SetRecorderUnionId(v string) *QueryFlashMinutesSummaryRequest {
	s.RecorderUnionId = &v
	return s
}

type QueryFlashMinutesSummaryResponseBody struct {
	FlashMinutesSummary *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary `json:"flashMinutesSummary,omitempty" xml:"flashMinutesSummary,omitempty" type:"Struct"`
}

func (s QueryFlashMinutesSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryResponseBody) SetFlashMinutesSummary(v *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary) *QueryFlashMinutesSummaryResponseBody {
	s.FlashMinutesSummary = v
	return s
}

type QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary struct {
	Status  *int32                                                            `json:"status,omitempty" xml:"status,omitempty"`
	Summary []*QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Repeated"`
}

func (s QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary) SetStatus(v int32) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary {
	s.Status = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary) SetSummary(v []*QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummary {
	s.Summary = v
	return s
}

type QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary struct {
	End      *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Headline *string `json:"headline,omitempty" xml:"headline,omitempty"`
	Id       *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Start    *int64  `json:"start,omitempty" xml:"start,omitempty"`
	Summary  *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) SetEnd(v int64) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary {
	s.End = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) SetHeadline(v string) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary {
	s.Headline = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) SetId(v int32) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary {
	s.Id = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) SetStart(v int64) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary {
	s.Start = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary) SetSummary(v string) *QueryFlashMinutesSummaryResponseBodyFlashMinutesSummarySummary {
	s.Summary = &v
	return s
}

type QueryFlashMinutesSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFlashMinutesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFlashMinutesSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFlashMinutesSummaryResponse) GoString() string {
	return s.String()
}

func (s *QueryFlashMinutesSummaryResponse) SetHeaders(v map[string]*string) *QueryFlashMinutesSummaryResponse {
	s.Headers = v
	return s
}

func (s *QueryFlashMinutesSummaryResponse) SetStatusCode(v int32) *QueryFlashMinutesSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFlashMinutesSummaryResponse) SetBody(v *QueryFlashMinutesSummaryResponseBody) *QueryFlashMinutesSummaryResponse {
	s.Body = v
	return s
}

type QueryMinutesAudioHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesAudioHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesAudioHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesAudioHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesAudioHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesAudioHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesAudioHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesAudioRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesAudioRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesAudioRequest) SetUnionId(v string) *QueryMinutesAudioRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesAudioResponseBody struct {
	AudioList []*QueryMinutesAudioResponseBodyAudioList `json:"audioList,omitempty" xml:"audioList,omitempty" type:"Repeated"`
}

func (s QueryMinutesAudioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesAudioResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesAudioResponseBody) SetAudioList(v []*QueryMinutesAudioResponseBodyAudioList) *QueryMinutesAudioResponseBody {
	s.AudioList = v
	return s
}

type QueryMinutesAudioResponseBodyAudioList struct {
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1631172094000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// https://xxx-hangzhou.oss-cn-hangzhou.aliyuncs.com/record_xxxx.mp3?Expires=1718045081&OSSAccessKeyId=TMP.3KdwHtvZxopmwacMZEdyb4WHLVmbArrNRB9CTKnR1MaJgmRjdmZczs6Rip66cgKgk2HhQon1yygvBnbY3uqEaZNeHBLcBa&Signature=OFWyAIY%2FdlzfwM9wIfEaKoAudkxxxxx
	PlayUrl *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	// example:
	//
	// 290882268xxx1172033231
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// example:
	//
	// 1631172094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// WFBkgJvtxxxxtSaA1jK4sgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesAudioResponseBodyAudioList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesAudioResponseBodyAudioList) GoString() string {
	return s.String()
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetDuration(v int64) *QueryMinutesAudioResponseBodyAudioList {
	s.Duration = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetEndTime(v int64) *QueryMinutesAudioResponseBodyAudioList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetFileSize(v int64) *QueryMinutesAudioResponseBodyAudioList {
	s.FileSize = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetPlayUrl(v string) *QueryMinutesAudioResponseBodyAudioList {
	s.PlayUrl = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetRecordId(v string) *QueryMinutesAudioResponseBodyAudioList {
	s.RecordId = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetStartTime(v int64) *QueryMinutesAudioResponseBodyAudioList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesAudioResponseBodyAudioList) SetUnionId(v string) *QueryMinutesAudioResponseBodyAudioList {
	s.UnionId = &v
	return s
}

type QueryMinutesAudioResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesAudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesAudioResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesAudioResponse) SetHeaders(v map[string]*string) *QueryMinutesAudioResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesAudioResponse) SetStatusCode(v int32) *QueryMinutesAudioResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesAudioResponse) SetBody(v *QueryMinutesAudioResponseBody) *QueryMinutesAudioResponse {
	s.Body = v
	return s
}

type QueryMinutesSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesSummaryRequest struct {
	SummaryTypeList []*string `json:"summaryTypeList,omitempty" xml:"summaryTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryRequest) SetSummaryTypeList(v []*string) *QueryMinutesSummaryRequest {
	s.SummaryTypeList = v
	return s
}

func (s *QueryMinutesSummaryRequest) SetUnionId(v string) *QueryMinutesSummaryRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesSummaryResponseBody struct {
	Summary *QueryMinutesSummaryResponseBodySummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
}

func (s QueryMinutesSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBody) SetSummary(v *QueryMinutesSummaryResponseBodySummary) *QueryMinutesSummaryResponseBody {
	s.Summary = v
	return s
}

type QueryMinutesSummaryResponseBodySummary struct {
	Actions                   []*QueryMinutesSummaryResponseBodySummaryActions                   `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	AutoChapters              []*QueryMinutesSummaryResponseBodySummaryAutoChapters              `json:"autoChapters,omitempty" xml:"autoChapters,omitempty" type:"Repeated"`
	ConversationalSummary     []*QueryMinutesSummaryResponseBodySummaryConversationalSummary     `json:"conversationalSummary,omitempty" xml:"conversationalSummary,omitempty" type:"Repeated"`
	KeySentences              []*QueryMinutesSummaryResponseBodySummaryKeySentences              `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	Keywords                  []*string                                                          `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	ParagraphSummary          *string                                                            `json:"paragraphSummary,omitempty" xml:"paragraphSummary,omitempty"`
	QuestionsAnsweringSummary []*QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary `json:"questionsAnsweringSummary,omitempty" xml:"questionsAnsweringSummary,omitempty" type:"Repeated"`
}

func (s QueryMinutesSummaryResponseBodySummary) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummary) SetActions(v []*QueryMinutesSummaryResponseBodySummaryActions) *QueryMinutesSummaryResponseBodySummary {
	s.Actions = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetAutoChapters(v []*QueryMinutesSummaryResponseBodySummaryAutoChapters) *QueryMinutesSummaryResponseBodySummary {
	s.AutoChapters = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetConversationalSummary(v []*QueryMinutesSummaryResponseBodySummaryConversationalSummary) *QueryMinutesSummaryResponseBodySummary {
	s.ConversationalSummary = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetKeySentences(v []*QueryMinutesSummaryResponseBodySummaryKeySentences) *QueryMinutesSummaryResponseBodySummary {
	s.KeySentences = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetKeywords(v []*string) *QueryMinutesSummaryResponseBodySummary {
	s.Keywords = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetParagraphSummary(v string) *QueryMinutesSummaryResponseBodySummary {
	s.ParagraphSummary = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetQuestionsAnsweringSummary(v []*QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) *QueryMinutesSummaryResponseBodySummary {
	s.QuestionsAnsweringSummary = v
	return s
}

type QueryMinutesSummaryResponseBodySummaryActions struct {
	End        *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	SentenceId *int64  `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	Start      *int64  `json:"start,omitempty" xml:"start,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryActions) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryActions) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetSentenceId(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.SentenceId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetText(v string) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Text = &v
	return s
}

type QueryMinutesSummaryResponseBodySummaryAutoChapters struct {
	End      *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Headline *string `json:"headline,omitempty" xml:"headline,omitempty"`
	Id       *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Start    *int64  `json:"start,omitempty" xml:"start,omitempty"`
	Summary  *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryAutoChapters) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryAutoChapters) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetHeadline(v string) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Headline = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetSummary(v string) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Summary = &v
	return s
}

type QueryMinutesSummaryResponseBodySummaryConversationalSummary struct {
	SpeakerId   *string `json:"speakerId,omitempty" xml:"speakerId,omitempty"`
	SpeakerName *string `json:"speakerName,omitempty" xml:"speakerName,omitempty"`
	Summary     *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryConversationalSummary) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryConversationalSummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSpeakerId(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.SpeakerId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSpeakerName(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.SpeakerName = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSummary(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.Summary = &v
	return s
}

type QueryMinutesSummaryResponseBodySummaryKeySentences struct {
	End        *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	SentenceId *int64  `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	Start      *int64  `json:"start,omitempty" xml:"start,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryKeySentences) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryKeySentences) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetSentenceId(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.SentenceId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetText(v string) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Text = &v
	return s
}

type QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary struct {
	Answer                *string  `json:"answer,omitempty" xml:"answer,omitempty"`
	Question              *string  `json:"question,omitempty" xml:"question,omitempty"`
	SentenceIdsOfAnswer   []*int64 `json:"sentenceIdsOfAnswer,omitempty" xml:"sentenceIdsOfAnswer,omitempty" type:"Repeated"`
	SentenceIdsOfQuestion []*int64 `json:"sentenceIdsOfQuestion,omitempty" xml:"sentenceIdsOfQuestion,omitempty" type:"Repeated"`
}

func (s QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetAnswer(v string) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.Answer = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetQuestion(v string) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.Question = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetSentenceIdsOfAnswer(v []*int64) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.SentenceIdsOfAnswer = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetSentenceIdsOfQuestion(v []*int64) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.SentenceIdsOfQuestion = v
	return s
}

type QueryMinutesSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesSummaryResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponse) SetHeaders(v map[string]*string) *QueryMinutesSummaryResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesSummaryResponse) SetStatusCode(v int32) *QueryMinutesSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesSummaryResponse) SetBody(v *QueryMinutesSummaryResponseBody) *QueryMinutesSummaryResponse {
	s.Body = v
	return s
}

type QueryMinutesTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMinutesTextHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesTextHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMinutesTextHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMinutesTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Direction  *string `json:"direction,omitempty" xml:"direction,omitempty"`
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTextRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextRequest) SetDirection(v string) *QueryMinutesTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryMinutesTextRequest) SetMaxResults(v int64) *QueryMinutesTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryMinutesTextRequest) SetNextToken(v string) *QueryMinutesTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextRequest) SetUnionId(v string) *QueryMinutesTextRequest {
	s.UnionId = &v
	return s
}

type QueryMinutesTextResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1631172045153000_7940
	NextToken     *string                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ParagraphList []*QueryMinutesTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBody) SetHasMore(v bool) *QueryMinutesTextResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetNextToken(v string) *QueryMinutesTextResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryMinutesTextResponseBody) SetParagraphList(v []*QueryMinutesTextResponseBodyParagraphList) *QueryMinutesTextResponseBody {
	s.ParagraphList = v
	return s
}

type QueryMinutesTextResponseBodyParagraphList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 小钉
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// 嘿！你好，这里是小钉
	Paragraph   *string `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	ParagraphId *int64  `json:"paragraphId,omitempty" xml:"paragraphId,omitempty"`
	// example:
	//
	// 44444
	RecordId     *int64                                                   `json:"recordId,omitempty" xml:"recordId,omitempty"`
	SentenceList []*QueryMinutesTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// WFBkgJvt0xxxxSaA1jK4sgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetNickName(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraph(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetParagraphId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.ParagraphId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetRecordId(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetSentenceList(v []*QueryMinutesTextResponseBodyParagraphListSentenceList) *QueryMinutesTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphList) SetUnionId(v string) *QueryMinutesTextResponseBodyParagraphList {
	s.UnionId = &v
	return s
}

type QueryMinutesTextResponseBodyParagraphListSentenceList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 这里是小钉
	Sentence *string `json:"sentence,omitempty" xml:"sentence,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// WFBkgJvt0xxxxSaA1jK4sgiEiE
	UnionId  *string                                                          `json:"unionId,omitempty" xml:"unionId,omitempty"`
	WordList []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetUnionId(v string) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.UnionId = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryMinutesTextResponseBodyParagraphListSentenceListWordList) *QueryMinutesTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

type QueryMinutesTextResponseBodyParagraphListSentenceListWordList struct {
	// example:
	//
	// 7940
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 7940
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 这里
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
	// example:
	//
	// 1631172050535000#0
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryMinutesTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryMinutesTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

type QueryMinutesTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMinutesTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMinutesTextResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMinutesTextResponse) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextResponse) SetHeaders(v map[string]*string) *QueryMinutesTextResponse {
	s.Headers = v
	return s
}

func (s *QueryMinutesTextResponse) SetStatusCode(v int32) *QueryMinutesTextResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMinutesTextResponse) SetBody(v *QueryMinutesTextResponseBody) *QueryMinutesTextResponse {
	s.Body = v
	return s
}

type QueryOrgConferenceListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgConferenceListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConferenceListHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgConferenceListHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgConferenceListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgConferenceListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgConferenceListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgConferenceListRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgConferenceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConferenceListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgConferenceListRequest) SetMaxResults(v int32) *QueryOrgConferenceListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgConferenceListRequest) SetNextToken(v string) *QueryOrgConferenceListRequest {
	s.NextToken = &v
	return s
}

type QueryOrgConferenceListResponseBody struct {
	HasMore         *bool                                                `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken       *string                                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OnGoingConfList []*QueryOrgConferenceListResponseBodyOnGoingConfList `json:"onGoingConfList,omitempty" xml:"onGoingConfList,omitempty" type:"Repeated"`
	TotalCount      *int32                                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryOrgConferenceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConferenceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgConferenceListResponseBody) SetHasMore(v bool) *QueryOrgConferenceListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryOrgConferenceListResponseBody) SetNextToken(v string) *QueryOrgConferenceListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgConferenceListResponseBody) SetOnGoingConfList(v []*QueryOrgConferenceListResponseBodyOnGoingConfList) *QueryOrgConferenceListResponseBody {
	s.OnGoingConfList = v
	return s
}

func (s *QueryOrgConferenceListResponseBody) SetTotalCount(v int32) *QueryOrgConferenceListResponseBody {
	s.TotalCount = &v
	return s
}

type QueryOrgConferenceListResponseBodyOnGoingConfList struct {
	BizType      *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	CreatorId    *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	CreatorNick  *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	RoomCode     *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryOrgConferenceListResponseBodyOnGoingConfList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConferenceListResponseBodyOnGoingConfList) GoString() string {
	return s.String()
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetBizType(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.BizType = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetConferenceId(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.ConferenceId = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetCreatorId(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.CreatorId = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetCreatorNick(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.CreatorNick = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetEndTime(v int64) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.EndTime = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetRoomCode(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.RoomCode = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetStartTime(v int64) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.StartTime = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetStatus(v int32) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.Status = &v
	return s
}

func (s *QueryOrgConferenceListResponseBodyOnGoingConfList) SetTitle(v string) *QueryOrgConferenceListResponseBodyOnGoingConfList {
	s.Title = &v
	return s
}

type QueryOrgConferenceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrgConferenceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrgConferenceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgConferenceListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgConferenceListResponse) SetHeaders(v map[string]*string) *QueryOrgConferenceListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgConferenceListResponse) SetStatusCode(v int32) *QueryOrgConferenceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgConferenceListResponse) SetBody(v *QueryOrgConferenceListResponseBody) *QueryOrgConferenceListResponse {
	s.Body = v
	return s
}

type QueryRecordMinutesUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRecordMinutesUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordMinutesUrlHeaders) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlHeaders) SetCommonHeaders(v map[string]*string) *QueryRecordMinutesUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRecordMinutesUrlHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRecordMinutesUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRecordMinutesUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_record
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lJcRnm39OsU4jlFVmRG9KXXXX
	RecorderUnionId *string `json:"recorderUnionId,omitempty" xml:"recorderUnionId,omitempty"`
}

func (s QueryRecordMinutesUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordMinutesUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlRequest) SetBizType(v string) *QueryRecordMinutesUrlRequest {
	s.BizType = &v
	return s
}

func (s *QueryRecordMinutesUrlRequest) SetRecorderUnionId(v string) *QueryRecordMinutesUrlRequest {
	s.RecorderUnionId = &v
	return s
}

type QueryRecordMinutesUrlResponseBody struct {
	RecordMinutesUrls []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls `json:"recordMinutesUrls,omitempty" xml:"recordMinutesUrls,omitempty" type:"Repeated"`
}

func (s QueryRecordMinutesUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordMinutesUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponseBody) SetRecordMinutesUrls(v []*QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) *QueryRecordMinutesUrlResponseBody {
	s.RecordMinutesUrls = v
	return s
}

type QueryRecordMinutesUrlResponseBodyRecordMinutesUrls struct {
	RecordMinutesUrl *string `json:"recordMinutesUrl,omitempty" xml:"recordMinutesUrl,omitempty"`
}

func (s QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls) SetRecordMinutesUrl(v string) *QueryRecordMinutesUrlResponseBodyRecordMinutesUrls {
	s.RecordMinutesUrl = &v
	return s
}

type QueryRecordMinutesUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecordMinutesUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecordMinutesUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordMinutesUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlResponse) SetHeaders(v map[string]*string) *QueryRecordMinutesUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordMinutesUrlResponse) SetStatusCode(v int32) *QueryRecordMinutesUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordMinutesUrlResponse) SetBody(v *QueryRecordMinutesUrlResponseBody) *QueryRecordMinutesUrlResponse {
	s.Body = v
	return s
}

type QueryScheduleConfSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScheduleConfSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConfSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConfSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScheduleConfSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScheduleConfSettingsRequest struct {
	// This parameter is required.
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
}

func (s QueryScheduleConfSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsRequest) SetScheduleConferenceId(v string) *QueryScheduleConfSettingsRequest {
	s.ScheduleConferenceId = &v
	return s
}

type QueryScheduleConfSettingsResponseBody struct {
	ScheduleConfSettingModel *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel `json:"scheduleConfSettingModel,omitempty" xml:"scheduleConfSettingModel,omitempty" type:"Struct"`
}

func (s QueryScheduleConfSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsResponseBody) SetScheduleConfSettingModel(v *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) *QueryScheduleConfSettingsResponseBody {
	s.ScheduleConfSettingModel = v
	return s
}

type QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel struct {
	CohostUnionIds              []*string                                                                                 `json:"cohostUnionIds,omitempty" xml:"cohostUnionIds,omitempty" type:"Repeated"`
	ConfAllowedCorpId           *string                                                                                   `json:"confAllowedCorpId,omitempty" xml:"confAllowedCorpId,omitempty"`
	HostUnionId                 *string                                                                                   `json:"hostUnionId,omitempty" xml:"hostUnionId,omitempty"`
	LockRoom                    *int32                                                                                    `json:"lockRoom,omitempty" xml:"lockRoom,omitempty"`
	MoziConfVirtualExtraSetting *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting `json:"moziConfVirtualExtraSetting,omitempty" xml:"moziConfVirtualExtraSetting,omitempty" type:"Struct"`
	MuteOnJoin                  *int32                                                                                    `json:"muteOnJoin,omitempty" xml:"muteOnJoin,omitempty"`
	ScreenShareForbidden        *int32                                                                                    `json:"screenShareForbidden,omitempty" xml:"screenShareForbidden,omitempty"`
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetCohostUnionIds(v []*string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.CohostUnionIds = v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetConfAllowedCorpId(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.ConfAllowedCorpId = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetHostUnionId(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.HostUnionId = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetLockRoom(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.LockRoom = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetMoziConfVirtualExtraSetting(v *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.MoziConfVirtualExtraSetting = v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetMuteOnJoin(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.MuteOnJoin = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel) SetScreenShareForbidden(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModel {
	s.ScreenShareForbidden = &v
	return s
}

type QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting struct {
	EnableChat                   *int32                                                                                                                  `json:"enableChat,omitempty" xml:"enableChat,omitempty"`
	EnableWebAnonymousJoin       *bool                                                                                                                   `json:"enableWebAnonymousJoin,omitempty" xml:"enableWebAnonymousJoin,omitempty"`
	JoinBeforeHost               *int32                                                                                                                  `json:"joinBeforeHost,omitempty" xml:"joinBeforeHost,omitempty"`
	LockMediaStatusMicMute       *int32                                                                                                                  `json:"lockMediaStatusMicMute,omitempty" xml:"lockMediaStatusMicMute,omitempty"`
	LockNick                     *int32                                                                                                                  `json:"lockNick,omitempty" xml:"lockNick,omitempty"`
	MoziConfExtensionAppSettings []*QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings `json:"moziConfExtensionAppSettings,omitempty" xml:"moziConfExtensionAppSettings,omitempty" type:"Repeated"`
	WaitingRoom                  *int32                                                                                                                  `json:"waitingRoom,omitempty" xml:"waitingRoom,omitempty"`
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableChat(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableChat = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableWebAnonymousJoin(v bool) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableWebAnonymousJoin = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetJoinBeforeHost(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.JoinBeforeHost = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockMediaStatusMicMute(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockMediaStatusMicMute = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockNick(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockNick = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMoziConfExtensionAppSettings(v []*QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MoziConfExtensionAppSettings = v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting) SetWaitingRoom(v int32) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.WaitingRoom = &v
	return s
}

type QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings struct {
	AutoOpenMode        *string `json:"autoOpenMode,omitempty" xml:"autoOpenMode,omitempty"`
	ClientId            *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	CoolAppCode         *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	ExtensionAppBizData *string `json:"extensionAppBizData,omitempty" xml:"extensionAppBizData,omitempty"`
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetAutoOpenMode(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetClientId(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ClientId = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetCoolAppCode(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.CoolAppCode = &v
	return s
}

func (s *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetExtensionAppBizData(v string) *QueryScheduleConfSettingsResponseBodyScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

type QueryScheduleConfSettingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConfSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConfSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConfSettingsResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConfSettingsResponse) SetHeaders(v map[string]*string) *QueryScheduleConfSettingsResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConfSettingsResponse) SetStatusCode(v int32) *QueryScheduleConfSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConfSettingsResponse) SetBody(v *QueryScheduleConfSettingsResponseBody) *QueryScheduleConfSettingsResponse {
	s.Body = v
	return s
}

type QueryScheduleConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScheduleConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScheduleConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxx
	RequestUnionId *string `json:"requestUnionId,omitempty" xml:"requestUnionId,omitempty"`
}

func (s QueryScheduleConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceRequest) SetRequestUnionId(v string) *QueryScheduleConferenceRequest {
	s.RequestUnionId = &v
	return s
}

type QueryScheduleConferenceResponseBody struct {
	// example:
	//
	// 1687928400000
	EndTime *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Phones  []*string `json:"phones,omitempty" xml:"phones,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 838 722 xxxxx
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// example:
	//
	// 2a489c68-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// 1687924800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 预约会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/Bsbp3ixxxxxUyJJ9
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryScheduleConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceResponseBody) SetEndTime(v int64) *QueryScheduleConferenceResponseBody {
	s.EndTime = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetPhones(v []*string) *QueryScheduleConferenceResponseBody {
	s.Phones = v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetRequestId(v string) *QueryScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetRoomCode(v string) *QueryScheduleConferenceResponseBody {
	s.RoomCode = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetScheduleConferenceId(v string) *QueryScheduleConferenceResponseBody {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetStartTime(v int64) *QueryScheduleConferenceResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetTitle(v string) *QueryScheduleConferenceResponseBody {
	s.Title = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetUrl(v string) *QueryScheduleConferenceResponseBody {
	s.Url = &v
	return s
}

type QueryScheduleConferenceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceResponse) SetHeaders(v map[string]*string) *QueryScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConferenceResponse) SetStatusCode(v int32) *QueryScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConferenceResponse) SetBody(v *QueryScheduleConferenceResponseBody) *QueryScheduleConferenceResponse {
	s.Body = v
	return s
}

type QueryScheduleConferenceInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryScheduleConferenceInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryScheduleConferenceInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryScheduleConferenceInfoRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryScheduleConferenceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoRequest) SetMaxResults(v int32) *QueryScheduleConferenceInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryScheduleConferenceInfoRequest) SetNextToken(v string) *QueryScheduleConferenceInfoRequest {
	s.NextToken = &v
	return s
}

type QueryScheduleConferenceInfoResponseBody struct {
	ConferenceList []*QueryScheduleConferenceInfoResponseBodyConferenceList `json:"conferenceList,omitempty" xml:"conferenceList,omitempty" type:"Repeated"`
	// This parameter is required.
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryScheduleConferenceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponseBody) SetConferenceList(v []*QueryScheduleConferenceInfoResponseBodyConferenceList) *QueryScheduleConferenceInfoResponseBody {
	s.ConferenceList = v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetNextToken(v string) *QueryScheduleConferenceInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBody) SetTotalCount(v int32) *QueryScheduleConferenceInfoResponseBody {
	s.TotalCount = &v
	return s
}

type QueryScheduleConferenceInfoResponseBodyConferenceList struct {
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	RoomCode     *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryScheduleConferenceInfoResponseBodyConferenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponseBodyConferenceList) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetConferenceId(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.ConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetEndTime(v int64) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.EndTime = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetRoomCode(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.RoomCode = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetStartTime(v int64) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.StartTime = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetStatus(v int32) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.Status = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponseBodyConferenceList) SetTitle(v string) *QueryScheduleConferenceInfoResponseBodyConferenceList {
	s.Title = &v
	return s
}

type QueryScheduleConferenceInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConferenceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponse) SetHeaders(v map[string]*string) *QueryScheduleConferenceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConferenceInfoResponse) SetStatusCode(v int32) *QueryScheduleConferenceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponse) SetBody(v *QueryScheduleConferenceInfoResponseBody) *QueryScheduleConferenceInfoResponse {
	s.Body = v
	return s
}

type QueryUserOnGoingConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserOnGoingConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOnGoingConferenceHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserOnGoingConferenceHeaders) SetCommonHeaders(v map[string]*string) *QueryUserOnGoingConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserOnGoingConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserOnGoingConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserOnGoingConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryUserOnGoingConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOnGoingConferenceRequest) GoString() string {
	return s.String()
}

func (s *QueryUserOnGoingConferenceRequest) SetUnionId(v string) *QueryUserOnGoingConferenceRequest {
	s.UnionId = &v
	return s
}

type QueryUserOnGoingConferenceResponseBody struct {
	MemberModelMap    map[string]*MemberModelMapValue `json:"memberModelMap,omitempty" xml:"memberModelMap,omitempty"`
	OnGoingConfIdList []*string                       `json:"onGoingConfIdList,omitempty" xml:"onGoingConfIdList,omitempty" type:"Repeated"`
}

func (s QueryUserOnGoingConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOnGoingConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserOnGoingConferenceResponseBody) SetMemberModelMap(v map[string]*MemberModelMapValue) *QueryUserOnGoingConferenceResponseBody {
	s.MemberModelMap = v
	return s
}

func (s *QueryUserOnGoingConferenceResponseBody) SetOnGoingConfIdList(v []*string) *QueryUserOnGoingConferenceResponseBody {
	s.OnGoingConfIdList = v
	return s
}

type QueryUserOnGoingConferenceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserOnGoingConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserOnGoingConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOnGoingConferenceResponse) GoString() string {
	return s.String()
}

func (s *QueryUserOnGoingConferenceResponse) SetHeaders(v map[string]*string) *QueryUserOnGoingConferenceResponse {
	s.Headers = v
	return s
}

func (s *QueryUserOnGoingConferenceResponse) SetStatusCode(v int32) *QueryUserOnGoingConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserOnGoingConferenceResponse) SetBody(v *QueryUserOnGoingConferenceResponseBody) *QueryUserOnGoingConferenceResponse {
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
	// example:
	//
	// 演讲
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 左上
	SmallWindowPosition *string `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type StartMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartMinutesHeaders) GoString() string {
	return s.String()
}

func (s *StartMinutesHeaders) SetCommonHeaders(v map[string]*string) *StartMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *StartMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartMinutesRequest struct {
	// example:
	//
	// 左上
	OwnerUnionId *string `json:"ownerUnionId,omitempty" xml:"ownerUnionId,omitempty"`
	// example:
	//
	// true
	RecordAudio *bool `json:"recordAudio,omitempty" xml:"recordAudio,omitempty"`
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StartMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMinutesRequest) GoString() string {
	return s.String()
}

func (s *StartMinutesRequest) SetOwnerUnionId(v string) *StartMinutesRequest {
	s.OwnerUnionId = &v
	return s
}

func (s *StartMinutesRequest) SetRecordAudio(v bool) *StartMinutesRequest {
	s.RecordAudio = &v
	return s
}

func (s *StartMinutesRequest) SetUnionId(v string) *StartMinutesRequest {
	s.UnionId = &v
	return s
}

type StartMinutesResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StartMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *StartMinutesResponseBody) SetCode(v string) *StartMinutesResponseBody {
	s.Code = &v
	return s
}

type StartMinutesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMinutesResponse) GoString() string {
	return s.String()
}

func (s *StartMinutesResponse) SetHeaders(v map[string]*string) *StartMinutesResponse {
	s.Headers = v
	return s
}

func (s *StartMinutesResponse) SetStatusCode(v int32) *StartMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMinutesResponse) SetBody(v *StartMinutesResponseBody) *StartMinutesResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// grip
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedHostJoin *bool `json:"needHostJoin,omitempty" xml:"needHostJoin,omitempty"`
	// This parameter is required.
	SmallWindowPosition *string `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 推流名称
	StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
	// This parameter is required.
	StreamUrlList []*string `json:"streamUrlList,omitempty" xml:"streamUrlList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type StopMinutesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopMinutesHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopMinutesHeaders) GoString() string {
	return s.String()
}

func (s *StopMinutesHeaders) SetCommonHeaders(v map[string]*string) *StopMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopMinutesHeaders) SetXAcsDingtalkAccessToken(v string) *StopMinutesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopMinutesRequest struct {
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StopMinutesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMinutesRequest) GoString() string {
	return s.String()
}

func (s *StopMinutesRequest) SetUnionId(v string) *StopMinutesRequest {
	s.UnionId = &v
	return s
}

type StopMinutesResponseBody struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StopMinutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *StopMinutesResponseBody) SetCode(v string) *StopMinutesResponseBody {
	s.Code = &v
	return s
}

type StopMinutesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMinutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMinutesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMinutesResponse) GoString() string {
	return s.String()
}

func (s *StopMinutesResponse) SetHeaders(v map[string]*string) *StopMinutesResponse {
	s.Headers = v
	return s
}

func (s *StopMinutesResponse) SetStatusCode(v int32) *StopMinutesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMinutesResponse) SetBody(v *StopMinutesResponseBody) *StopMinutesResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// true
	StopAllStream *bool `json:"stopAllStream,omitempty" xml:"stopAllStream,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 27SaQ3iiHLN0uwqcPisedfreNwiEiE
	StreamId *string `json:"streamId,omitempty" xml:"streamId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateScheduleConfSettingsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateScheduleConfSettingsHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConfSettingsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConfSettingsHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateScheduleConfSettingsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateScheduleConfSettingsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CreatorUnionId           *string                                                    `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	ScheduleConfSettingModel *UpdateScheduleConfSettingsRequestScheduleConfSettingModel `json:"scheduleConfSettingModel,omitempty" xml:"scheduleConfSettingModel,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// f6fb627e-a7e8-403e-b1f8-26e85450f4a9
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
}

func (s UpdateScheduleConfSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequest) SetCreatorUnionId(v string) *UpdateScheduleConfSettingsRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequest) SetScheduleConfSettingModel(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) *UpdateScheduleConfSettingsRequest {
	s.ScheduleConfSettingModel = v
	return s
}

func (s *UpdateScheduleConfSettingsRequest) SetScheduleConferenceId(v string) *UpdateScheduleConfSettingsRequest {
	s.ScheduleConferenceId = &v
	return s
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModel struct {
	CohostUnionIds []*string `json:"cohostUnionIds,omitempty" xml:"cohostUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// dingc02f685fa06381c44ac5d6980864d335
	ConfAllowedCorpId *string `json:"confAllowedCorpId,omitempty" xml:"confAllowedCorpId,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	HostUnionId *string `json:"hostUnionId,omitempty" xml:"hostUnionId,omitempty"`
	// example:
	//
	// 0：取消锁定 1：锁定
	LockRoom                    *int32                                                                                `json:"lockRoom,omitempty" xml:"lockRoom,omitempty"`
	MoziConfOpenRecordSetting   *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting   `json:"moziConfOpenRecordSetting,omitempty" xml:"moziConfOpenRecordSetting,omitempty" type:"Struct"`
	MoziConfVirtualExtraSetting *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting `json:"moziConfVirtualExtraSetting,omitempty" xml:"moziConfVirtualExtraSetting,omitempty" type:"Struct"`
	// example:
	//
	// -1：开启 0：未开启 6：超过6人自动开启静音
	MuteOnJoin *int32 `json:"muteOnJoin,omitempty" xml:"muteOnJoin,omitempty"`
	// example:
	//
	// 0：允许共享 1：禁止共享
	ScreenShareForbidden *int32 `json:"screenShareForbidden,omitempty" xml:"screenShareForbidden,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetCohostUnionIds(v []*string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.CohostUnionIds = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetConfAllowedCorpId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.ConfAllowedCorpId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetHostUnionId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.HostUnionId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetLockRoom(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.LockRoom = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMoziConfOpenRecordSetting(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MoziConfOpenRecordSetting = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMoziConfVirtualExtraSetting(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MoziConfVirtualExtraSetting = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMuteOnJoin(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MuteOnJoin = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetScreenShareForbidden(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.ScreenShareForbidden = &v
	return s
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting struct {
	// example:
	//
	// true：跟随 false：不跟随
	IsFollowHost *bool `json:"isFollowHost,omitempty" xml:"isFollowHost,omitempty"`
	// example:
	//
	// grid：宫格模式,默认9宫格(3x3) speech：演讲者模式 full_screen：全屏模式 auto_grid：自动宫格模式，默认最大4x4宫格 screen_cast：屏幕共享模式，仅放置屏幕共享流 p2p：双人通话模式 full_screen_and_speaker：共享内容+发言人模式
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// 0：不自动开启 1：自动开启
	RecordAutoStart *int32 `json:"recordAutoStart,omitempty" xml:"recordAutoStart,omitempty"`
	// example:
	//
	// 0：我以主持人身份入会后自动开启 1：其他人以联席主持人身份入会后开启 2：任何人以任何身份入会后开启
	RecordAutoStartType *int32 `json:"recordAutoStartType,omitempty" xml:"recordAutoStartType,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetIsFollowHost(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.IsFollowHost = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetMode(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.Mode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStart(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStart = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStartType(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStartType = &v
	return s
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting struct {
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CloudRecordOwnerUnionId *string `json:"cloudRecordOwnerUnionId,omitempty" xml:"cloudRecordOwnerUnionId,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	EnableChat *int32 `json:"enableChat,omitempty" xml:"enableChat,omitempty"`
	// example:
	//
	// true：允许匿名登录入会 false：不允许匿名登录入会
	EnableWebAnonymousJoin *bool `json:"enableWebAnonymousJoin,omitempty" xml:"enableWebAnonymousJoin,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	JoinBeforeHost *int32 `json:"joinBeforeHost,omitempty" xml:"joinBeforeHost,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	LockMediaStatusMicMute *int32 `json:"lockMediaStatusMicMute,omitempty" xml:"lockMediaStatusMicMute,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	LockNick *int32 `json:"lockNick,omitempty" xml:"lockNick,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	MinutesOwnerUnionId          *string                                                                                                             `json:"minutesOwnerUnionId,omitempty" xml:"minutesOwnerUnionId,omitempty"`
	MoziConfExtensionAppSettings []*UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings `json:"moziConfExtensionAppSettings,omitempty" xml:"moziConfExtensionAppSettings,omitempty" type:"Repeated"`
	PushAllMeetingRecords        *bool                                                                                                               `json:"pushAllMeetingRecords,omitempty" xml:"pushAllMeetingRecords,omitempty"`
	PushCloudRecordCard          *bool                                                                                                               `json:"pushCloudRecordCard,omitempty" xml:"pushCloudRecordCard,omitempty"`
	PushMinutesCard              *bool                                                                                                               `json:"pushMinutesCard,omitempty" xml:"pushMinutesCard,omitempty"`
	// example:
	//
	// 0：未开启 1：开启
	WaitingRoom *int32 `json:"waitingRoom,omitempty" xml:"waitingRoom,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetCloudRecordOwnerUnionId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.CloudRecordOwnerUnionId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableChat(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableChat = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableWebAnonymousJoin(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableWebAnonymousJoin = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetJoinBeforeHost(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.JoinBeforeHost = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockMediaStatusMicMute(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockMediaStatusMicMute = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockNick(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockNick = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMinutesOwnerUnionId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MinutesOwnerUnionId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMoziConfExtensionAppSettings(v []*UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MoziConfExtensionAppSettings = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushAllMeetingRecords(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushAllMeetingRecords = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushCloudRecordCard(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushCloudRecordCard = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushMinutesCard(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushMinutesCard = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetWaitingRoom(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.WaitingRoom = &v
	return s
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings struct {
	// example:
	//
	// 0：不自动打开 1：仅主持人/联席主持人自动打开 2：全员自动打开
	AutoOpenMode *int32 `json:"autoOpenMode,omitempty" xml:"autoOpenMode,omitempty"`
	// example:
	//
	// COOLAPP-0-1026633886192127xxxB000W
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// example:
	//
	// bizData
	ExtensionAppBizData *string `json:"extensionAppBizData,omitempty" xml:"extensionAppBizData,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetAutoOpenMode(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetCoolAppCode(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.CoolAppCode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetExtensionAppBizData(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

type UpdateScheduleConfSettingsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScheduleConfSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsResponseBody) SetSuccess(v bool) *UpdateScheduleConfSettingsResponseBody {
	s.Success = &v
	return s
}

type UpdateScheduleConfSettingsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleConfSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleConfSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConfSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsResponse) SetHeaders(v map[string]*string) *UpdateScheduleConfSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleConfSettingsResponse) SetStatusCode(v int32) *UpdateScheduleConfSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleConfSettingsResponse) SetBody(v *UpdateScheduleConfSettingsResponseBody) *UpdateScheduleConfSettingsResponse {
	s.Body = v
	return s
}

type UpdateScheduleConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateScheduleConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateScheduleConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qzR1iSMDvzR9iP7Pxxxxxxxxxxxx
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687924800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 预约会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateScheduleConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceRequest) SetCreatorUnionId(v string) *UpdateScheduleConferenceRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetEndTime(v int64) *UpdateScheduleConferenceRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetScheduleConferenceId(v string) *UpdateScheduleConferenceRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetStartTime(v int64) *UpdateScheduleConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetTitle(v string) *UpdateScheduleConferenceRequest {
	s.Title = &v
	return s
}

type UpdateScheduleConferenceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScheduleConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceResponseBody) SetSuccess(v bool) *UpdateScheduleConferenceResponseBody {
	s.Success = &v
	return s
}

type UpdateScheduleConferenceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceResponse) SetHeaders(v map[string]*string) *UpdateScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleConferenceResponse) SetStatusCode(v int32) *UpdateScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleConferenceResponse) SetBody(v *UpdateScheduleConferenceResponseBody) *UpdateScheduleConferenceResponse {
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoConferenceExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoConferenceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
// 增加闪记权限
//
// @param request - AddRecordPermissionRequest
//
// @param headers - AddRecordPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordPermissionResponse
func (client *Client) AddRecordPermissionWithOptions(conferenceId *string, request *AddRecordPermissionRequest, headers *AddRecordPermissionHeaders, runtime *util.RuntimeOptions) (_result *AddRecordPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUnionId)) {
		body["ownerUnionId"] = request.OwnerUnionId
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
		Action:      tea.String("AddRecordPermission"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/flashMinutes/recordPermissions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRecordPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加闪记权限
//
// @param request - AddRecordPermissionRequest
//
// @return AddRecordPermissionResponse
func (client *Client) AddRecordPermission(conferenceId *string, request *AddRecordPermissionRequest) (_result *AddRecordPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRecordPermissionHeaders{}
	_result = &AddRecordPermissionResponse{}
	_body, _err := client.AddRecordPermissionWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消预约会议
//
// @param request - CancelScheduleConferenceRequest
//
// @param headers - CancelScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelScheduleConferenceResponse
func (client *Client) CancelScheduleConferenceWithOptions(request *CancelScheduleConferenceRequest, headers *CancelScheduleConferenceHeaders, runtime *util.RuntimeOptions) (_result *CancelScheduleConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConferenceId)) {
		body["scheduleConferenceId"] = request.ScheduleConferenceId
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
		Action:      tea.String("CancelScheduleConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelScheduleConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消预约会议
//
// @param request - CancelScheduleConferenceRequest
//
// @return CancelScheduleConferenceResponse
func (client *Client) CancelScheduleConference(request *CancelScheduleConferenceRequest) (_result *CancelScheduleConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelScheduleConferenceHeaders{}
	_result = &CancelScheduleConferenceResponse{}
	_body, _err := client.CancelScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭视频会议
//
// @param request - CloseVideoConferenceRequest
//
// @param headers - CloseVideoConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseVideoConferenceResponse
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

// Summary:
//
// 关闭视频会议
//
// @param request - CloseVideoConferenceRequest
//
// @return CloseVideoConferenceResponse
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

// Summary:
//
// 设置联席主持人
//
// @param request - CohostsRequest
//
// @param headers - CohostsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CohostsResponse
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

// Summary:
//
// 设置联席主持人
//
// @param request - CohostsRequest
//
// @return CohostsResponse
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

// Summary:
//
// 创建专属短链
//
// @param request - CreateCustomShortLinkRequest
//
// @param headers - CreateCustomShortLinkHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomShortLinkResponse
func (client *Client) CreateCustomShortLinkWithOptions(request *CreateCustomShortLinkRequest, headers *CreateCustomShortLinkHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomShortLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionAppBizData)) {
		body["extensionAppBizData"] = request.ExtensionAppBizData
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConferenceId)) {
		body["scheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.UseExtensionApp)) {
		body["useExtensionApp"] = request.UseExtensionApp
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
		Action:      tea.String("CreateCustomShortLink"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/customShortLinks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomShortLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建专属短链
//
// @param request - CreateCustomShortLinkRequest
//
// @return CreateCustomShortLinkResponse
func (client *Client) CreateCustomShortLink(request *CreateCustomShortLinkRequest) (_result *CreateCustomShortLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomShortLinkHeaders{}
	_result = &CreateCustomShortLinkResponse{}
	_body, _err := client.CreateCustomShortLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预约会议
//
// @param request - CreateScheduleConferenceRequest
//
// @param headers - CreateScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduleConferenceResponse
func (client *Client) CreateScheduleConferenceWithOptions(request *CreateScheduleConferenceRequest, headers *CreateScheduleConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateScheduleConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConfSettingModel)) {
		body["scheduleConfSettingModel"] = request.ScheduleConfSettingModel
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("CreateScheduleConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduleConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预约会议
//
// @param request - CreateScheduleConferenceRequest
//
// @return CreateScheduleConferenceResponse
func (client *Client) CreateScheduleConference(request *CreateScheduleConferenceRequest) (_result *CreateScheduleConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateScheduleConferenceHeaders{}
	_result = &CreateScheduleConferenceResponse{}
	_body, _err := client.CreateScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建视频会议
//
// @param request - CreateVideoConferenceRequest
//
// @param headers - CreateVideoConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoConferenceResponse
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

// Summary:
//
// 创建视频会议
//
// @param request - CreateVideoConferenceRequest
//
// @return CreateVideoConferenceResponse
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

// Summary:
//
// 设置全员看他
//
// @param request - FocusRequest
//
// @param headers - FocusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FocusResponse
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

// Summary:
//
// 设置全员看他
//
// @param request - FocusRequest
//
// @return FocusResponse
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

// Summary:
//
// 生成会议闪记文档的下载链接
//
// @param request - GenerateFlashMinutesDocumentUrlRequest
//
// @param headers - GenerateFlashMinutesDocumentUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateFlashMinutesDocumentUrlResponse
func (client *Client) GenerateFlashMinutesDocumentUrlWithOptions(conferenceId *string, request *GenerateFlashMinutesDocumentUrlRequest, headers *GenerateFlashMinutesDocumentUrlHeaders, runtime *util.RuntimeOptions) (_result *GenerateFlashMinutesDocumentUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["expireTime"] = request.ExpireTime
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
		Action:      tea.String("GenerateFlashMinutesDocumentUrl"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/flashMinutes/document/generate"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateFlashMinutesDocumentUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成会议闪记文档的下载链接
//
// @param request - GenerateFlashMinutesDocumentUrlRequest
//
// @return GenerateFlashMinutesDocumentUrlResponse
func (client *Client) GenerateFlashMinutesDocumentUrl(conferenceId *string, request *GenerateFlashMinutesDocumentUrlRequest) (_result *GenerateFlashMinutesDocumentUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GenerateFlashMinutesDocumentUrlHeaders{}
	_result = &GenerateFlashMinutesDocumentUrlResponse{}
	_body, _err := client.GenerateFlashMinutesDocumentUrlWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过conferenceId获取指定音视频会议信息
//
// @param request - GetConfDataByConferenceIdRequest
//
// @param headers - GetConfDataByConferenceIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfDataByConferenceIdResponse
func (client *Client) GetConfDataByConferenceIdWithOptions(conferenceId *string, request *GetConfDataByConferenceIdRequest, headers *GetConfDataByConferenceIdHeaders, runtime *util.RuntimeOptions) (_result *GetConfDataByConferenceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RealData)) {
		query["realData"] = request.RealData
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
		Action:      tea.String("GetConfDataByConferenceId"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfDataByConferenceIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过conferenceId获取指定音视频会议信息
//
// @param request - GetConfDataByConferenceIdRequest
//
// @return GetConfDataByConferenceIdResponse
func (client *Client) GetConfDataByConferenceId(conferenceId *string, request *GetConfDataByConferenceIdRequest) (_result *GetConfDataByConferenceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConfDataByConferenceIdHeaders{}
	_result = &GetConfDataByConferenceIdResponse{}
	_body, _err := client.GetConfDataByConferenceIdWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过conferenceId获取指定音视频会议成员信息
//
// @param request - GetConfDetailDataRequest
//
// @param headers - GetConfDetailDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfDetailDataResponse
func (client *Client) GetConfDetailDataWithOptions(conferenceId *string, request *GetConfDetailDataRequest, headers *GetConfDetailDataHeaders, runtime *util.RuntimeOptions) (_result *GetConfDetailDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		query["nick"] = request.Nick
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
		Action:      tea.String("GetConfDetailData"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfDetailDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过conferenceId获取指定音视频会议成员信息
//
// @param request - GetConfDetailDataRequest
//
// @return GetConfDetailDataResponse
func (client *Client) GetConfDetailData(conferenceId *string, request *GetConfDetailDataRequest) (_result *GetConfDetailDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConfDetailDataHeaders{}
	_result = &GetConfDetailDataResponse{}
	_body, _err := client.GetConfDetailDataWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取音视频会议列表数据
//
// @param request - GetHistoryConfDataListRequest
//
// @param headers - GetHistoryConfDataListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryConfDataListResponse
func (client *Client) GetHistoryConfDataListWithOptions(request *GetHistoryConfDataListRequest, headers *GetHistoryConfDataListHeaders, runtime *util.RuntimeOptions) (_result *GetHistoryConfDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorNike)) {
		query["creatorNike"] = request.CreatorNike
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FreeType)) {
		query["freeType"] = request.FreeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RealData)) {
		query["realData"] = request.RealData
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
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
		Action:      tea.String("GetHistoryConfDataList"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/histories/dataLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryConfDataListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取音视频会议列表数据
//
// @param request - GetHistoryConfDataListRequest
//
// @return GetHistoryConfDataListResponse
func (client *Client) GetHistoryConfDataList(request *GetHistoryConfDataListRequest) (_result *GetHistoryConfDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHistoryConfDataListHeaders{}
	_result = &GetHistoryConfDataListResponse{}
	_body, _err := client.GetHistoryConfDataListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过conferenceId和unionId获取最新会议质量数据
//
// @param request - GetUserLastMetricRequest
//
// @param headers - GetUserLastMetricHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserLastMetricResponse
func (client *Client) GetUserLastMetricWithOptions(conferenceId *string, request *GetUserLastMetricRequest, headers *GetUserLastMetricHeaders, runtime *util.RuntimeOptions) (_result *GetUserLastMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionIdList)) {
		body["unionIdList"] = request.UnionIdList
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
		Action:      tea.String("GetUserLastMetric"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/lastMetricDatas/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserLastMetricResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过conferenceId和unionId获取最新会议质量数据
//
// @param request - GetUserLastMetricRequest
//
// @return GetUserLastMetricResponse
func (client *Client) GetUserLastMetric(conferenceId *string, request *GetUserLastMetricRequest) (_result *GetUserLastMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserLastMetricHeaders{}
	_result = &GetUserLastMetricResponse{}
	_body, _err := client.GetUserLastMetricWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过conferenceId和unionId获取指定音视频会议人员的会议质量数据
//
// @param request - GetUserMetricDataRequest
//
// @param headers - GetUserMetricDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMetricDataResponse
func (client *Client) GetUserMetricDataWithOptions(conferenceId *string, request *GetUserMetricDataRequest, headers *GetUserMetricDataHeaders, runtime *util.RuntimeOptions) (_result *GetUserMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["beginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
		Action:      tea.String("GetUserMetricData"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/metricDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserMetricDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过conferenceId和unionId获取指定音视频会议人员的会议质量数据
//
// @param request - GetUserMetricDataRequest
//
// @return GetUserMetricDataResponse
func (client *Client) GetUserMetricData(conferenceId *string, request *GetUserMetricDataRequest) (_result *GetUserMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserMetricDataHeaders{}
	_result = &GetUserMetricDataResponse{}
	_body, _err := client.GetUserMetricDataWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 邀请其他人员
//
// @param request - InviteUsersRequest
//
// @param headers - InviteUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteUsersResponse
func (client *Client) InviteUsersWithOptions(conferenceId *string, request *InviteUsersRequest, headers *InviteUsersHeaders, runtime *util.RuntimeOptions) (_result *InviteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteeList)) {
		body["inviteeList"] = request.InviteeList
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneInviteeList)) {
		body["phoneInviteeList"] = request.PhoneInviteeList
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

// Summary:
//
// 邀请其他人员
//
// @param request - InviteUsersRequest
//
// @return InviteUsersResponse
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

// Summary:
//
// 会议踢出成员
//
// @param request - KickMembersRequest
//
// @param headers - KickMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickMembersResponse
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

// Summary:
//
// 会议踢出成员
//
// @param request - KickMembersRequest
//
// @return KickMembersResponse
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

// Summary:
//
// 锁定会议
//
// @param request - LockConferenceRequest
//
// @param headers - LockConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockConferenceResponse
func (client *Client) LockConferenceWithOptions(conferenceId *string, request *LockConferenceRequest, headers *LockConferenceHeaders, runtime *util.RuntimeOptions) (_result *LockConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
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
		Action:      tea.String("LockConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/lock"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LockConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 锁定会议
//
// @param request - LockConferenceRequest
//
// @return LockConferenceResponse
func (client *Client) LockConference(conferenceId *string, request *LockConferenceRequest) (_result *LockConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LockConferenceHeaders{}
	_result = &LockConferenceResponse{}
	_body, _err := client.LockConferenceWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会议全员静音或解除静音
//
// @param request - MuteAllRequest
//
// @param headers - MuteAllHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MuteAllResponse
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

// Summary:
//
// 会议全员静音或解除静音
//
// @param request - MuteAllRequest
//
// @return MuteAllResponse
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

// Summary:
//
// 指定人员静音或取消静音
//
// @param request - MuteMembersRequest
//
// @param headers - MuteMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MuteMembersResponse
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

// Summary:
//
// 指定人员静音或取消静音
//
// @param request - MuteMembersRequest
//
// @return MuteMembersResponse
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

// Summary:
//
// 查询云录制文本信息
//
// @param request - QueryCloudRecordTextRequest
//
// @param headers - QueryCloudRecordTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordTextResponse
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
		ReqBodyType: tea.String("none"),
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

// Summary:
//
// 查询云录制文本信息
//
// @param request - QueryCloudRecordTextRequest
//
// @return QueryCloudRecordTextResponse
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

// Summary:
//
// 查询云录制视频
//
// @param request - QueryCloudRecordVideoRequest
//
// @param headers - QueryCloudRecordVideoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordVideoResponse
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
		ReqBodyType: tea.String("none"),
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

// Summary:
//
// 查询云录制视频
//
// @param request - QueryCloudRecordVideoRequest
//
// @return QueryCloudRecordVideoResponse
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

// Summary:
//
// 查询云录制视频播放信息
//
// @param request - QueryCloudRecordVideoPlayInfoRequest
//
// @param headers - QueryCloudRecordVideoPlayInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCloudRecordVideoPlayInfoResponse
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
		ReqBodyType: tea.String("none"),
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

// Summary:
//
// 查询云录制视频播放信息
//
// @param request - QueryCloudRecordVideoPlayInfoRequest
//
// @return QueryCloudRecordVideoPlayInfoResponse
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

// Summary:
//
// 查询视频会议信息
//
// @param headers - QueryConferenceInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceInfoResponse
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

// Summary:
//
// 查询视频会议信息
//
// @return QueryConferenceInfoResponse
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

// Summary:
//
// 批量查询视频会议信息
//
// @param request - QueryConferenceInfoBatchRequest
//
// @param headers - QueryConferenceInfoBatchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceInfoBatchResponse
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

// Summary:
//
// 批量查询视频会议信息
//
// @param request - QueryConferenceInfoBatchRequest
//
// @return QueryConferenceInfoBatchResponse
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

// Summary:
//
// 根据会议号查询会议信息
//
// @param request - QueryConferenceInfoByRoomCodeRequest
//
// @param headers - QueryConferenceInfoByRoomCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceInfoByRoomCodeResponse
func (client *Client) QueryConferenceInfoByRoomCodeWithOptions(roomCode *string, request *QueryConferenceInfoByRoomCodeRequest, headers *QueryConferenceInfoByRoomCodeHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceInfoByRoomCodeResponse, _err error) {
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
		Action:      tea.String("QueryConferenceInfoByRoomCode"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/roomCodes/" + tea.StringValue(roomCode) + "/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConferenceInfoByRoomCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据会议号查询会议信息
//
// @param request - QueryConferenceInfoByRoomCodeRequest
//
// @return QueryConferenceInfoByRoomCodeResponse
func (client *Client) QueryConferenceInfoByRoomCode(roomCode *string, request *QueryConferenceInfoByRoomCodeRequest) (_result *QueryConferenceInfoByRoomCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceInfoByRoomCodeHeaders{}
	_result = &QueryConferenceInfoByRoomCodeResponse{}
	_body, _err := client.QueryConferenceInfoByRoomCodeWithOptions(roomCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询视频会议成员
//
// @param request - QueryConferenceMembersRequest
//
// @param headers - QueryConferenceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConferenceMembersResponse
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

// Summary:
//
// 查询视频会议成员
//
// @param request - QueryConferenceMembersRequest
//
// @return QueryConferenceMembersResponse
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

// Summary:
//
// 查询云录制摘要请求
//
// @param request - QueryFlashMinutesSummaryRequest
//
// @param headers - QueryFlashMinutesSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFlashMinutesSummaryResponse
func (client *Client) QueryFlashMinutesSummaryWithOptions(conferenceId *string, request *QueryFlashMinutesSummaryRequest, headers *QueryFlashMinutesSummaryHeaders, runtime *util.RuntimeOptions) (_result *QueryFlashMinutesSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.RecorderUnionId)) {
		query["recorderUnionId"] = request.RecorderUnionId
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
		Action:      tea.String("QueryFlashMinutesSummary"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/flashMinutes/summaries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFlashMinutesSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云录制摘要请求
//
// @param request - QueryFlashMinutesSummaryRequest
//
// @return QueryFlashMinutesSummaryResponse
func (client *Client) QueryFlashMinutesSummary(conferenceId *string, request *QueryFlashMinutesSummaryRequest) (_result *QueryFlashMinutesSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFlashMinutesSummaryHeaders{}
	_result = &QueryFlashMinutesSummaryResponse{}
	_body, _err := client.QueryFlashMinutesSummaryWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会议闪记的音频信息
//
// @param request - QueryMinutesAudioRequest
//
// @param headers - QueryMinutesAudioHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesAudioResponse
func (client *Client) QueryMinutesAudioWithOptions(conferenceId *string, request *QueryMinutesAudioRequest, headers *QueryMinutesAudioHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesAudioResponse, _err error) {
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
		Action:      tea.String("QueryMinutesAudio"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/minutes/audioInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesAudioResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议闪记的音频信息
//
// @param request - QueryMinutesAudioRequest
//
// @return QueryMinutesAudioResponse
func (client *Client) QueryMinutesAudio(conferenceId *string, request *QueryMinutesAudioRequest) (_result *QueryMinutesAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesAudioHeaders{}
	_result = &QueryMinutesAudioResponse{}
	_body, _err := client.QueryMinutesAudioWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会议闪记智能纪要
//
// @param request - QueryMinutesSummaryRequest
//
// @param headers - QueryMinutesSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesSummaryResponse
func (client *Client) QueryMinutesSummaryWithOptions(conferenceId *string, request *QueryMinutesSummaryRequest, headers *QueryMinutesSummaryHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SummaryTypeList)) {
		body["summaryTypeList"] = request.SummaryTypeList
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
		Action:      tea.String("QueryMinutesSummary"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/minutes/summaries/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议闪记智能纪要
//
// @param request - QueryMinutesSummaryRequest
//
// @return QueryMinutesSummaryResponse
func (client *Client) QueryMinutesSummary(conferenceId *string, request *QueryMinutesSummaryRequest) (_result *QueryMinutesSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesSummaryHeaders{}
	_result = &QueryMinutesSummaryResponse{}
	_body, _err := client.QueryMinutesSummaryWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会议闪记文本信息
//
// @param request - QueryMinutesTextRequest
//
// @param headers - QueryMinutesTextHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesTextWithOptions(conferenceId *string, request *QueryMinutesTextRequest, headers *QueryMinutesTextHeaders, runtime *util.RuntimeOptions) (_result *QueryMinutesTextResponse, _err error) {
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
		Action:      tea.String("QueryMinutesText"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/minutes/textInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会议闪记文本信息
//
// @param request - QueryMinutesTextRequest
//
// @return QueryMinutesTextResponse
func (client *Client) QueryMinutesText(conferenceId *string, request *QueryMinutesTextRequest) (_result *QueryMinutesTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMinutesTextHeaders{}
	_result = &QueryMinutesTextResponse{}
	_body, _err := client.QueryMinutesTextWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业进行中会议列表
//
// @param request - QueryOrgConferenceListRequest
//
// @param headers - QueryOrgConferenceListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrgConferenceListResponse
func (client *Client) QueryOrgConferenceListWithOptions(request *QueryOrgConferenceListRequest, headers *QueryOrgConferenceListHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgConferenceListResponse, _err error) {
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
		Action:      tea.String("QueryOrgConferenceList"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/orgConferences"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgConferenceListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业进行中会议列表
//
// @param request - QueryOrgConferenceListRequest
//
// @return QueryOrgConferenceListResponse
func (client *Client) QueryOrgConferenceList(request *QueryOrgConferenceListRequest) (_result *QueryOrgConferenceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgConferenceListHeaders{}
	_result = &QueryOrgConferenceListResponse{}
	_body, _err := client.QueryOrgConferenceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询闪记链接
//
// @param request - QueryRecordMinutesUrlRequest
//
// @param headers - QueryRecordMinutesUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecordMinutesUrlResponse
func (client *Client) QueryRecordMinutesUrlWithOptions(conferenceId *string, request *QueryRecordMinutesUrlRequest, headers *QueryRecordMinutesUrlHeaders, runtime *util.RuntimeOptions) (_result *QueryRecordMinutesUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.RecorderUnionId)) {
		query["recorderUnionId"] = request.RecorderUnionId
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
		Action:      tea.String("QueryRecordMinutesUrl"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/flashMinutes/recordUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordMinutesUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询闪记链接
//
// @param request - QueryRecordMinutesUrlRequest
//
// @return QueryRecordMinutesUrlResponse
func (client *Client) QueryRecordMinutesUrl(conferenceId *string, request *QueryRecordMinutesUrlRequest) (_result *QueryRecordMinutesUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRecordMinutesUrlHeaders{}
	_result = &QueryRecordMinutesUrlResponse{}
	_body, _err := client.QueryRecordMinutesUrlWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询预约会议设置
//
// @param request - QueryScheduleConfSettingsRequest
//
// @param headers - QueryScheduleConfSettingsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConfSettingsResponse
func (client *Client) QueryScheduleConfSettingsWithOptions(request *QueryScheduleConfSettingsRequest, headers *QueryScheduleConfSettingsHeaders, runtime *util.RuntimeOptions) (_result *QueryScheduleConfSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScheduleConferenceId)) {
		query["scheduleConferenceId"] = request.ScheduleConferenceId
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
		Action:      tea.String("QueryScheduleConfSettings"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences/settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScheduleConfSettingsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预约会议设置
//
// @param request - QueryScheduleConfSettingsRequest
//
// @return QueryScheduleConfSettingsResponse
func (client *Client) QueryScheduleConfSettings(request *QueryScheduleConfSettingsRequest) (_result *QueryScheduleConfSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScheduleConfSettingsHeaders{}
	_result = &QueryScheduleConfSettingsResponse{}
	_body, _err := client.QueryScheduleConfSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询预约会议信息
//
// @param request - QueryScheduleConferenceRequest
//
// @param headers - QueryScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConferenceResponse
func (client *Client) QueryScheduleConferenceWithOptions(scheduleConferenceId *string, request *QueryScheduleConferenceRequest, headers *QueryScheduleConferenceHeaders, runtime *util.RuntimeOptions) (_result *QueryScheduleConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestUnionId)) {
		query["requestUnionId"] = request.RequestUnionId
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
		Action:      tea.String("QueryScheduleConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences/" + tea.StringValue(scheduleConferenceId) + "/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScheduleConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预约会议信息
//
// @param request - QueryScheduleConferenceRequest
//
// @return QueryScheduleConferenceResponse
func (client *Client) QueryScheduleConference(scheduleConferenceId *string, request *QueryScheduleConferenceRequest) (_result *QueryScheduleConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScheduleConferenceHeaders{}
	_result = &QueryScheduleConferenceResponse{}
	_body, _err := client.QueryScheduleConferenceWithOptions(scheduleConferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取预约会议历史会议信息，当前仅返回最后一次的会议信息
//
// @param request - QueryScheduleConferenceInfoRequest
//
// @param headers - QueryScheduleConferenceInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScheduleConferenceInfoResponse
func (client *Client) QueryScheduleConferenceInfoWithOptions(scheduleConferenceId *string, request *QueryScheduleConferenceInfoRequest, headers *QueryScheduleConferenceInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryScheduleConferenceInfoResponse, _err error) {
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
		Action:      tea.String("QueryScheduleConferenceInfo"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/scheduleConferences/" + tea.StringValue(scheduleConferenceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScheduleConferenceInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取预约会议历史会议信息，当前仅返回最后一次的会议信息
//
// @param request - QueryScheduleConferenceInfoRequest
//
// @return QueryScheduleConferenceInfoResponse
func (client *Client) QueryScheduleConferenceInfo(scheduleConferenceId *string, request *QueryScheduleConferenceInfoRequest) (_result *QueryScheduleConferenceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryScheduleConferenceInfoHeaders{}
	_result = &QueryScheduleConferenceInfoResponse{}
	_body, _err := client.QueryScheduleConferenceInfoWithOptions(scheduleConferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户进行中会议
//
// @param request - QueryUserOnGoingConferenceRequest
//
// @param headers - QueryUserOnGoingConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserOnGoingConferenceResponse
func (client *Client) QueryUserOnGoingConferenceWithOptions(request *QueryUserOnGoingConferenceRequest, headers *QueryUserOnGoingConferenceHeaders, runtime *util.RuntimeOptions) (_result *QueryUserOnGoingConferenceResponse, _err error) {
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
		Action:      tea.String("QueryUserOnGoingConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/users/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserOnGoingConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户进行中会议
//
// @param request - QueryUserOnGoingConferenceRequest
//
// @return QueryUserOnGoingConferenceResponse
func (client *Client) QueryUserOnGoingConference(request *QueryUserOnGoingConferenceRequest) (_result *QueryUserOnGoingConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserOnGoingConferenceHeaders{}
	_result = &QueryUserOnGoingConferenceResponse{}
	_body, _err := client.QueryUserOnGoingConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启云录制
//
// @param request - StartCloudRecordRequest
//
// @param headers - StartCloudRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
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

// Summary:
//
// 开启云录制
//
// @param request - StartCloudRecordRequest
//
// @return StartCloudRecordResponse
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

// Summary:
//
// 开启会议闪记
//
// @param request - StartMinutesRequest
//
// @param headers - StartMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMinutesResponse
func (client *Client) StartMinutesWithOptions(conferenceId *string, request *StartMinutesRequest, headers *StartMinutesHeaders, runtime *util.RuntimeOptions) (_result *StartMinutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerUnionId)) {
		body["ownerUnionId"] = request.OwnerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.RecordAudio)) {
		body["recordAudio"] = request.RecordAudio
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
		Action:      tea.String("StartMinutes"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/minutes/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启会议闪记
//
// @param request - StartMinutesRequest
//
// @return StartMinutesResponse
func (client *Client) StartMinutes(conferenceId *string, request *StartMinutesRequest) (_result *StartMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartMinutesHeaders{}
	_result = &StartMinutesResponse{}
	_body, _err := client.StartMinutesWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会议开始直播推流
//
// @param request - StartStreamOutRequest
//
// @param headers - StartStreamOutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamOutResponse
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

// Summary:
//
// 会议开始直播推流
//
// @param request - StartStreamOutRequest
//
// @return StartStreamOutResponse
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

// Summary:
//
// 关闭云录制
//
// @param request - StopCloudRecordRequest
//
// @param headers - StopCloudRecordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
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

// Summary:
//
// 关闭云录制
//
// @param request - StopCloudRecordRequest
//
// @return StopCloudRecordResponse
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

// Summary:
//
// 暂停会议闪记
//
// @param request - StopMinutesRequest
//
// @param headers - StopMinutesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMinutesResponse
func (client *Client) StopMinutesWithOptions(conferenceId *string, request *StopMinutesRequest, headers *StopMinutesHeaders, runtime *util.RuntimeOptions) (_result *StopMinutesResponse, _err error) {
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
		Action:      tea.String("StopMinutes"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/videoConferences/" + tea.StringValue(conferenceId) + "/minutes/pause"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMinutesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停会议闪记
//
// @param request - StopMinutesRequest
//
// @return StopMinutesResponse
func (client *Client) StopMinutes(conferenceId *string, request *StopMinutesRequest) (_result *StopMinutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopMinutesHeaders{}
	_result = &StopMinutesResponse{}
	_body, _err := client.StopMinutesWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会议停止直播推流
//
// @param request - StopStreamOutRequest
//
// @param headers - StopStreamOutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamOutResponse
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

// Summary:
//
// 会议停止直播推流
//
// @param request - StopStreamOutRequest
//
// @return StopStreamOutResponse
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

// Summary:
//
// 更新预约会议设置
//
// @param request - UpdateScheduleConfSettingsRequest
//
// @param headers - UpdateScheduleConfSettingsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleConfSettingsResponse
func (client *Client) UpdateScheduleConfSettingsWithOptions(request *UpdateScheduleConfSettingsRequest, headers *UpdateScheduleConfSettingsHeaders, runtime *util.RuntimeOptions) (_result *UpdateScheduleConfSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConfSettingModel)) {
		body["scheduleConfSettingModel"] = request.ScheduleConfSettingModel
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConferenceId)) {
		body["scheduleConferenceId"] = request.ScheduleConferenceId
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
		Action:      tea.String("UpdateScheduleConfSettings"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences/settings"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScheduleConfSettingsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新预约会议设置
//
// @param request - UpdateScheduleConfSettingsRequest
//
// @return UpdateScheduleConfSettingsResponse
func (client *Client) UpdateScheduleConfSettings(request *UpdateScheduleConfSettingsRequest) (_result *UpdateScheduleConfSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateScheduleConfSettingsHeaders{}
	_result = &UpdateScheduleConfSettingsResponse{}
	_body, _err := client.UpdateScheduleConfSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新预约会议
//
// @param request - UpdateScheduleConferenceRequest
//
// @param headers - UpdateScheduleConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleConferenceResponse
func (client *Client) UpdateScheduleConferenceWithOptions(request *UpdateScheduleConferenceRequest, headers *UpdateScheduleConferenceHeaders, runtime *util.RuntimeOptions) (_result *UpdateScheduleConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleConferenceId)) {
		body["scheduleConferenceId"] = request.ScheduleConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("UpdateScheduleConference"),
		Version:     tea.String("conference_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/conference/scheduleConferences"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScheduleConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新预约会议
//
// @param request - UpdateScheduleConferenceRequest
//
// @return UpdateScheduleConferenceResponse
func (client *Client) UpdateScheduleConference(request *UpdateScheduleConferenceRequest) (_result *UpdateScheduleConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateScheduleConferenceHeaders{}
	_result = &UpdateScheduleConferenceResponse{}
	_body, _err := client.UpdateScheduleConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新会议额外信息
//
// @param headers - UpdateVideoConferenceExtInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoConferenceExtInfoResponse
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

// Summary:
//
// 更新会议额外信息
//
// @return UpdateVideoConferenceExtInfoResponse
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

// Summary:
//
// 设置会议中的会议属性
//
// @param request - UpdateVideoConferenceSettingRequest
//
// @param headers - UpdateVideoConferenceSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVideoConferenceSettingResponse
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

// Summary:
//
// 设置会议中的会议属性
//
// @param request - UpdateVideoConferenceSettingRequest
//
// @return UpdateVideoConferenceSettingResponse
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
