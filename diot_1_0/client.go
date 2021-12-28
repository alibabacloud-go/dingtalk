// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package diot_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchDeleteDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDeviceHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteDeviceHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchDeleteDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchDeleteDeviceRequest struct {
	// 钉钉物联组织ID, 第三方平台必填，企业内部系统忽略。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 设备ID列表，最多500条。
	DeviceIds []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s BatchDeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDeviceRequest) SetCorpId(v string) *BatchDeleteDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *BatchDeleteDeviceRequest) SetDeviceIds(v []*string) *BatchDeleteDeviceRequest {
	s.DeviceIds = v
	return s
}

type BatchDeleteDeviceResponseBody struct {
	// 成功删除设备ID列表。
	DeviceIds []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s BatchDeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDeviceResponseBody) SetDeviceIds(v []*string) *BatchDeleteDeviceResponseBody {
	s.DeviceIds = v
	return s
}

type BatchDeleteDeviceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDeviceResponse) SetHeaders(v map[string]*string) *BatchDeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDeviceResponse) SetBody(v *BatchDeleteDeviceResponseBody) *BatchDeleteDeviceResponse {
	s.Body = v
	return s
}

type PushEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushEventHeaders) GoString() string {
	return s.String()
}

func (s *PushEventHeaders) SetCommonHeaders(v map[string]*string) *PushEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushEventHeaders) SetXAcsDingtalkAccessToken(v string) *PushEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushEventRequest struct {
	// 钉钉物联组织ID, 第三方平台必填，企业内部系统忽略。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 事件ID。
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// 事件类型，最长20个字符。
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// 事件名称，长度4-20个字符，一个中文汉字算2个字符。
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// 事件发生事件，Unix时间戳，单位毫秒。
	OccurrenceTime *int64 `json:"occurrenceTime,omitempty" xml:"occurrenceTime,omitempty"`
	// 触发事件设备ID。
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 事件发生地点。
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 事件文字信息。
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 事件图片地址列表。
	PicUrls []*string `json:"picUrls,omitempty" xml:"picUrls,omitempty" type:"Repeated"`
	// 第三方平台定制参数，企业内部系统忽略。
	ExtraData map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
}

func (s PushEventRequest) String() string {
	return tea.Prettify(s)
}

func (s PushEventRequest) GoString() string {
	return s.String()
}

func (s *PushEventRequest) SetCorpId(v string) *PushEventRequest {
	s.CorpId = &v
	return s
}

func (s *PushEventRequest) SetEventId(v string) *PushEventRequest {
	s.EventId = &v
	return s
}

func (s *PushEventRequest) SetEventType(v string) *PushEventRequest {
	s.EventType = &v
	return s
}

func (s *PushEventRequest) SetEventName(v string) *PushEventRequest {
	s.EventName = &v
	return s
}

func (s *PushEventRequest) SetOccurrenceTime(v int64) *PushEventRequest {
	s.OccurrenceTime = &v
	return s
}

func (s *PushEventRequest) SetDeviceId(v string) *PushEventRequest {
	s.DeviceId = &v
	return s
}

func (s *PushEventRequest) SetLocation(v string) *PushEventRequest {
	s.Location = &v
	return s
}

func (s *PushEventRequest) SetMsg(v string) *PushEventRequest {
	s.Msg = &v
	return s
}

func (s *PushEventRequest) SetPicUrls(v []*string) *PushEventRequest {
	s.PicUrls = v
	return s
}

func (s *PushEventRequest) SetExtraData(v map[string]interface{}) *PushEventRequest {
	s.ExtraData = v
	return s
}

type PushEventResponseBody struct {
	// 事件ID。
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
}

func (s PushEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushEventResponseBody) GoString() string {
	return s.String()
}

func (s *PushEventResponseBody) SetEventId(v string) *PushEventResponseBody {
	s.EventId = &v
	return s
}

type PushEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushEventResponse) String() string {
	return tea.Prettify(s)
}

func (s PushEventResponse) GoString() string {
	return s.String()
}

func (s *PushEventResponse) SetHeaders(v map[string]*string) *PushEventResponse {
	s.Headers = v
	return s
}

func (s *PushEventResponse) SetBody(v *PushEventResponseBody) *PushEventResponse {
	s.Body = v
	return s
}

type DeviceConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeviceConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceConferenceHeaders) GoString() string {
	return s.String()
}

func (s *DeviceConferenceHeaders) SetCommonHeaders(v map[string]*string) *DeviceConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *DeviceConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeviceConferenceRequest struct {
	// 会议主题，最多不能超20个中文。
	ConfTitle *string `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	// 钉钉会议ID，加入已存在的会议必填。
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 钉钉会议密码，加入已存在的会议必填。
	ConferencePassword *string `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	// 需要邀请的设备ID，最多5个。
	DeviceIds []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s DeviceConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceConferenceRequest) GoString() string {
	return s.String()
}

func (s *DeviceConferenceRequest) SetConfTitle(v string) *DeviceConferenceRequest {
	s.ConfTitle = &v
	return s
}

func (s *DeviceConferenceRequest) SetConferenceId(v string) *DeviceConferenceRequest {
	s.ConferenceId = &v
	return s
}

func (s *DeviceConferenceRequest) SetConferencePassword(v string) *DeviceConferenceRequest {
	s.ConferencePassword = &v
	return s
}

func (s *DeviceConferenceRequest) SetDeviceIds(v []*string) *DeviceConferenceRequest {
	s.DeviceIds = v
	return s
}

type DeviceConferenceResponseBody struct {
	// 会议ID
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s DeviceConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceConferenceResponseBody) SetConferenceId(v string) *DeviceConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

type DeviceConferenceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeviceConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeviceConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceConferenceResponse) GoString() string {
	return s.String()
}

func (s *DeviceConferenceResponse) SetHeaders(v map[string]*string) *DeviceConferenceResponse {
	s.Headers = v
	return s
}

func (s *DeviceConferenceResponse) SetBody(v *DeviceConferenceResponseBody) *DeviceConferenceResponse {
	s.Body = v
	return s
}

type RegisterDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceHeaders) GoString() string {
	return s.String()
}

func (s *RegisterDeviceHeaders) SetCommonHeaders(v map[string]*string) *RegisterDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterDeviceRequest struct {
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 设备id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 设备昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 设备地址
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 设备状态 0:在线 1:离线
	DeviceStatus *int32 `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	// 设备类型
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// 设备类型名称
	DeviceTypeName *string `json:"deviceTypeName,omitempty" xml:"deviceTypeName,omitempty"`
	// 设备父节点id
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 设备类型 摄像头:CAMERA 其它:OTHERS
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// 视频流地址
	LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetCorpId(v string) *RegisterDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *RegisterDeviceRequest) SetId(v string) *RegisterDeviceRequest {
	s.Id = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceName(v string) *RegisterDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *RegisterDeviceRequest) SetNickName(v string) *RegisterDeviceRequest {
	s.NickName = &v
	return s
}

func (s *RegisterDeviceRequest) SetLocation(v string) *RegisterDeviceRequest {
	s.Location = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceStatus(v int32) *RegisterDeviceRequest {
	s.DeviceStatus = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceType(v string) *RegisterDeviceRequest {
	s.DeviceType = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceTypeName(v string) *RegisterDeviceRequest {
	s.DeviceTypeName = &v
	return s
}

func (s *RegisterDeviceRequest) SetParentId(v string) *RegisterDeviceRequest {
	s.ParentId = &v
	return s
}

func (s *RegisterDeviceRequest) SetProductType(v string) *RegisterDeviceRequest {
	s.ProductType = &v
	return s
}

func (s *RegisterDeviceRequest) SetLiveUrl(v string) *RegisterDeviceRequest {
	s.LiveUrl = &v
	return s
}

type RegisterDeviceResponseBody struct {
	// 设备id
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetDeviceId(v string) *RegisterDeviceResponseBody {
	s.DeviceId = &v
	return s
}

type RegisterDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type BatchRegisterDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRegisterDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceHeaders) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceHeaders) SetCommonHeaders(v map[string]*string) *BatchRegisterDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRegisterDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRegisterDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRegisterDeviceRequest struct {
	// 钉钉物联组织ID, 第三方平台必填，企业内部系统忽略。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 设备列表。
	Devices []*BatchRegisterDeviceRequestDevices `json:"devices,omitempty" xml:"devices,omitempty" type:"Repeated"`
}

func (s BatchRegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequest) SetCorpId(v string) *BatchRegisterDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *BatchRegisterDeviceRequest) SetDevices(v []*BatchRegisterDeviceRequestDevices) *BatchRegisterDeviceRequest {
	s.Devices = v
	return s
}

type BatchRegisterDeviceRequestDevices struct {
	// 设备ID。
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 设备名称。
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 设备状态  0:在线  1:离线
	DeviceStatus *int32 `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	// 设备类型，自定义传入，最多128个字节。
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// 设备类型名称，自定义传入，最多128个字节，与deviceType一一对应。
	DeviceTypeName *string `json:"deviceTypeName,omitempty" xml:"deviceTypeName,omitempty"`
	// 产品类型 CAMERA：摄像头，可看直播 OTHERS：非摄像头
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// 视频流地址直播流地址，支持rtmp、flv、hls等格式，需要https协议。
	LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty"`
	// 父设备ID。
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 设备地址。
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 第三方平台定制参数，企业内部系统忽略。
	ExtraData map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
}

func (s BatchRegisterDeviceRequestDevices) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequestDevices) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequestDevices) SetDeviceId(v string) *BatchRegisterDeviceRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetDeviceName(v string) *BatchRegisterDeviceRequestDevices {
	s.DeviceName = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetDeviceStatus(v int32) *BatchRegisterDeviceRequestDevices {
	s.DeviceStatus = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetDeviceType(v string) *BatchRegisterDeviceRequestDevices {
	s.DeviceType = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetDeviceTypeName(v string) *BatchRegisterDeviceRequestDevices {
	s.DeviceTypeName = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetProductType(v string) *BatchRegisterDeviceRequestDevices {
	s.ProductType = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetLiveUrl(v string) *BatchRegisterDeviceRequestDevices {
	s.LiveUrl = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetParentId(v string) *BatchRegisterDeviceRequestDevices {
	s.ParentId = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetLocation(v string) *BatchRegisterDeviceRequestDevices {
	s.Location = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetExtraData(v map[string]interface{}) *BatchRegisterDeviceRequestDevices {
	s.ExtraData = v
	return s
}

type BatchRegisterDeviceResponseBody struct {
	// 注册成功的设备ID列表。
	DeviceIds []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s BatchRegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponseBody) SetDeviceIds(v []*string) *BatchRegisterDeviceResponseBody {
	s.DeviceIds = v
	return s
}

type BatchRegisterDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponse) SetHeaders(v map[string]*string) *BatchRegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchRegisterDeviceResponse) SetBody(v *BatchRegisterDeviceResponseBody) *BatchRegisterDeviceResponse {
	s.Body = v
	return s
}

type BatchRegisterEventTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRegisterEventTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterEventTypeHeaders) GoString() string {
	return s.String()
}

func (s *BatchRegisterEventTypeHeaders) SetCommonHeaders(v map[string]*string) *BatchRegisterEventTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRegisterEventTypeHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRegisterEventTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRegisterEventTypeRequest struct {
	// 钉钉物联组织ID, 第三方平台必填，企业内部系统忽略。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 事件类型列表，最多支持添加500个。
	EventTypes []*BatchRegisterEventTypeRequestEventTypes `json:"eventTypes,omitempty" xml:"eventTypes,omitempty" type:"Repeated"`
}

func (s BatchRegisterEventTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterEventTypeRequest) GoString() string {
	return s.String()
}

func (s *BatchRegisterEventTypeRequest) SetCorpId(v string) *BatchRegisterEventTypeRequest {
	s.CorpId = &v
	return s
}

func (s *BatchRegisterEventTypeRequest) SetEventTypes(v []*BatchRegisterEventTypeRequestEventTypes) *BatchRegisterEventTypeRequest {
	s.EventTypes = v
	return s
}

type BatchRegisterEventTypeRequestEventTypes struct {
	// 事件类型(唯一)，最长20个字符。
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// 事件类型名称，长度4-20个字符，一个中文汉字算2个字符。
	EventTypeName *string `json:"eventTypeName,omitempty" xml:"eventTypeName,omitempty"`
}

func (s BatchRegisterEventTypeRequestEventTypes) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterEventTypeRequestEventTypes) GoString() string {
	return s.String()
}

func (s *BatchRegisterEventTypeRequestEventTypes) SetEventType(v string) *BatchRegisterEventTypeRequestEventTypes {
	s.EventType = &v
	return s
}

func (s *BatchRegisterEventTypeRequestEventTypes) SetEventTypeName(v string) *BatchRegisterEventTypeRequestEventTypes {
	s.EventTypeName = &v
	return s
}

type BatchRegisterEventTypeResponseBody struct {
	// 注册成功的事件类型列表。
	EventTypes []*string `json:"eventTypes,omitempty" xml:"eventTypes,omitempty" type:"Repeated"`
}

func (s BatchRegisterEventTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterEventTypeResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRegisterEventTypeResponseBody) SetEventTypes(v []*string) *BatchRegisterEventTypeResponseBody {
	s.EventTypes = v
	return s
}

type BatchRegisterEventTypeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRegisterEventTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRegisterEventTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterEventTypeResponse) GoString() string {
	return s.String()
}

func (s *BatchRegisterEventTypeResponse) SetHeaders(v map[string]*string) *BatchRegisterEventTypeResponse {
	s.Headers = v
	return s
}

func (s *BatchRegisterEventTypeResponse) SetBody(v *BatchRegisterEventTypeResponseBody) *BatchRegisterEventTypeResponse {
	s.Body = v
	return s
}

type BatchUpdateDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateDeviceRequest struct {
	// 钉钉物联组织ID, 第三方平台必填，企业内部系统忽略。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 设备列表。
	Devices []*BatchUpdateDeviceRequestDevices `json:"devices,omitempty" xml:"devices,omitempty" type:"Repeated"`
}

func (s BatchUpdateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceRequest) SetCorpId(v string) *BatchUpdateDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *BatchUpdateDeviceRequest) SetDevices(v []*BatchUpdateDeviceRequestDevices) *BatchUpdateDeviceRequest {
	s.Devices = v
	return s
}

type BatchUpdateDeviceRequestDevices struct {
	// 设备ID。
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 设备名称。
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 设备地址。
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 设备状态 0:在线 1:离线
	DeviceStatus *int32 `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	// 视频流地址直播流地址，支持rtmp、flv、hls等格式，需要https协议。
	LiveUrl *string `json:"liveUrl,omitempty" xml:"liveUrl,omitempty"`
	// 第三方平台定制参数，企业内部系统忽略。
	ExtraData map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
}

func (s BatchUpdateDeviceRequestDevices) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceRequestDevices) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceRequestDevices) SetDeviceId(v string) *BatchUpdateDeviceRequestDevices {
	s.DeviceId = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetDeviceName(v string) *BatchUpdateDeviceRequestDevices {
	s.DeviceName = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetLocation(v string) *BatchUpdateDeviceRequestDevices {
	s.Location = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetDeviceStatus(v int32) *BatchUpdateDeviceRequestDevices {
	s.DeviceStatus = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetLiveUrl(v string) *BatchUpdateDeviceRequestDevices {
	s.LiveUrl = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetExtraData(v map[string]interface{}) *BatchUpdateDeviceRequestDevices {
	s.ExtraData = v
	return s
}

type BatchUpdateDeviceResponseBody struct {
	// 修改成功的设备ID列表。
	DeviceIds []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s BatchUpdateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceResponseBody) SetDeviceIds(v []*string) *BatchUpdateDeviceResponseBody {
	s.DeviceIds = v
	return s
}

type BatchUpdateDeviceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceResponse) SetHeaders(v map[string]*string) *BatchUpdateDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateDeviceResponse) SetBody(v *BatchUpdateDeviceResponseBody) *BatchUpdateDeviceResponse {
	s.Body = v
	return s
}

type BindSystemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BindSystemHeaders) String() string {
	return tea.Prettify(s)
}

func (s BindSystemHeaders) GoString() string {
	return s.String()
}

func (s *BindSystemHeaders) SetCommonHeaders(v map[string]*string) *BindSystemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BindSystemHeaders) SetXAcsDingtalkAccessToken(v string) *BindSystemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BindSystemRequest struct {
	// 三方平台的用户的钉钉物联组织ID。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 与三方平台绑定验证的临时授权码。
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	// 三方平台的用户ID。
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// 三方平台的用户名。
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// 三方平台协定的其它参数。
	ExtraData map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
}

func (s BindSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSystemRequest) GoString() string {
	return s.String()
}

func (s *BindSystemRequest) SetCorpId(v string) *BindSystemRequest {
	s.CorpId = &v
	return s
}

func (s *BindSystemRequest) SetAuthCode(v string) *BindSystemRequest {
	s.AuthCode = &v
	return s
}

func (s *BindSystemRequest) SetClientId(v string) *BindSystemRequest {
	s.ClientId = &v
	return s
}

func (s *BindSystemRequest) SetClientName(v string) *BindSystemRequest {
	s.ClientName = &v
	return s
}

func (s *BindSystemRequest) SetExtraData(v map[string]interface{}) *BindSystemRequest {
	s.ExtraData = v
	return s
}

type BindSystemResponseBody struct {
	// 钉钉物联组织ID。
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 三方平台的用户ID。
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s BindSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSystemResponseBody) GoString() string {
	return s.String()
}

func (s *BindSystemResponseBody) SetCorpId(v string) *BindSystemResponseBody {
	s.CorpId = &v
	return s
}

func (s *BindSystemResponseBody) SetClientId(v string) *BindSystemResponseBody {
	s.ClientId = &v
	return s
}

type BindSystemResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSystemResponse) GoString() string {
	return s.String()
}

func (s *BindSystemResponse) SetHeaders(v map[string]*string) *BindSystemResponse {
	s.Headers = v
	return s
}

func (s *BindSystemResponse) SetBody(v *BindSystemResponseBody) *BindSystemResponse {
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

func (client *Client) BatchDeleteDevice(request *BatchDeleteDeviceRequest) (_result *BatchDeleteDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchDeleteDeviceHeaders{}
	_result = &BatchDeleteDeviceResponse{}
	_body, _err := client.BatchDeleteDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteDeviceWithOptions(request *BatchDeleteDeviceRequest, headers *BatchDeleteDeviceHeaders, runtime *util.RuntimeOptions) (_result *BatchDeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		body["deviceIds"] = request.DeviceIds
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
	_result = &BatchDeleteDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchDeleteDevice"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/devices/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushEvent(request *PushEventRequest) (_result *PushEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushEventHeaders{}
	_result = &PushEventResponse{}
	_body, _err := client.PushEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushEventWithOptions(request *PushEventRequest, headers *PushEventHeaders, runtime *util.RuntimeOptions) (_result *PushEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		body["eventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["eventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.OccurrenceTime)) {
		body["occurrenceTime"] = request.OccurrenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		body["msg"] = request.Msg
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrls)) {
		body["picUrls"] = request.PicUrls
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraData)) {
		body["extraData"] = request.ExtraData
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
	_result = &PushEventResponse{}
	_body, _err := client.DoROARequest(tea.String("PushEvent"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/events/push"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeviceConference(request *DeviceConferenceRequest) (_result *DeviceConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceConferenceHeaders{}
	_result = &DeviceConferenceResponse{}
	_body, _err := client.DeviceConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeviceConferenceWithOptions(request *DeviceConferenceRequest, headers *DeviceConferenceHeaders, runtime *util.RuntimeOptions) (_result *DeviceConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfTitle)) {
		body["confTitle"] = request.ConfTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["conferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.ConferencePassword)) {
		body["conferencePassword"] = request.ConferencePassword
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIds)) {
		body["deviceIds"] = request.DeviceIds
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
	_result = &DeviceConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeviceConference"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/deviceConferences/initiate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterDeviceHeaders{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, headers *RegisterDeviceHeaders, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceStatus)) {
		body["deviceStatus"] = request.DeviceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTypeName)) {
		body["deviceTypeName"] = request.DeviceTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.LiveUrl)) {
		body["liveUrl"] = request.LiveUrl
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
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterDevice"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/devices/register"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRegisterDevice(request *BatchRegisterDeviceRequest) (_result *BatchRegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRegisterDeviceHeaders{}
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.BatchRegisterDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRegisterDeviceWithOptions(request *BatchRegisterDeviceRequest, headers *BatchRegisterDeviceHeaders, runtime *util.RuntimeOptions) (_result *BatchRegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Devices)) {
		body["devices"] = request.Devices
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
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRegisterDevice"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/devices/registrations/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRegisterEventType(request *BatchRegisterEventTypeRequest) (_result *BatchRegisterEventTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRegisterEventTypeHeaders{}
	_result = &BatchRegisterEventTypeResponse{}
	_body, _err := client.BatchRegisterEventTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRegisterEventTypeWithOptions(request *BatchRegisterEventTypeRequest, headers *BatchRegisterEventTypeHeaders, runtime *util.RuntimeOptions) (_result *BatchRegisterEventTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EventTypes)) {
		body["eventTypes"] = request.EventTypes
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
	_result = &BatchRegisterEventTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRegisterEventType"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/eventTypes/registrations/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateDevice(request *BatchUpdateDeviceRequest) (_result *BatchUpdateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateDeviceHeaders{}
	_result = &BatchUpdateDeviceResponse{}
	_body, _err := client.BatchUpdateDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateDeviceWithOptions(request *BatchUpdateDeviceRequest, headers *BatchUpdateDeviceHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Devices)) {
		body["devices"] = request.Devices
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
	_result = &BatchUpdateDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchUpdateDevice"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/diot/devices/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindSystem(request *BindSystemRequest) (_result *BindSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BindSystemHeaders{}
	_result = &BindSystemResponse{}
	_body, _err := client.BindSystemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindSystemWithOptions(request *BindSystemRequest, headers *BindSystemHeaders, runtime *util.RuntimeOptions) (_result *BindSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientName)) {
		body["clientName"] = request.ClientName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraData)) {
		body["extraData"] = request.ExtraData
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
	_result = &BindSystemResponse{}
	_body, _err := client.DoROARequest(tea.String("BindSystem"), tea.String("diot_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/diot/systems/bind"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
