// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package diot_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AyunOnlienTestRequest struct {
	ReqId *string `json:"reqId,omitempty" xml:"reqId,omitempty"`
}

func (s AyunOnlienTestRequest) String() string {
	return tea.Prettify(s)
}

func (s AyunOnlienTestRequest) GoString() string {
	return s.String()
}

func (s *AyunOnlienTestRequest) SetReqId(v string) *AyunOnlienTestRequest {
	s.ReqId = &v
	return s
}

type AyunOnlienTestResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AyunOnlienTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AyunOnlienTestResponseBody) GoString() string {
	return s.String()
}

func (s *AyunOnlienTestResponseBody) SetRequestId(v string) *AyunOnlienTestResponseBody {
	s.RequestId = &v
	return s
}

type AyunOnlienTestResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AyunOnlienTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AyunOnlienTestResponse) String() string {
	return tea.Prettify(s)
}

func (s AyunOnlienTestResponse) GoString() string {
	return s.String()
}

func (s *AyunOnlienTestResponse) SetHeaders(v map[string]*string) *AyunOnlienTestResponse {
	s.Headers = v
	return s
}

func (s *AyunOnlienTestResponse) SetStatusCode(v int32) *AyunOnlienTestResponse {
	s.StatusCode = &v
	return s
}

func (s *AyunOnlienTestResponse) SetBody(v *AyunOnlienTestResponseBody) *AyunOnlienTestResponse {
	s.Body = v
	return s
}

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
	CorpId    *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchDeleteDeviceResponse) SetStatusCode(v int32) *BatchDeleteDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDeviceResponse) SetBody(v *BatchDeleteDeviceResponseBody) *BatchDeleteDeviceResponse {
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
	CorpId  *string                              `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	DeviceId       *string                                    `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName     *string                                    `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceStatus   *int32                                     `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	DeviceType     *string                                    `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	DeviceTypeName *string                                    `json:"deviceTypeName,omitempty" xml:"deviceTypeName,omitempty"`
	ExtraData      map[string]interface{}                     `json:"extraData,omitempty" xml:"extraData,omitempty"`
	LiveUrls       *BatchRegisterDeviceRequestDevicesLiveUrls `json:"liveUrls,omitempty" xml:"liveUrls,omitempty" type:"Struct"`
	Location       *string                                    `json:"location,omitempty" xml:"location,omitempty"`
	ParentId       *string                                    `json:"parentId,omitempty" xml:"parentId,omitempty"`
	ProductType    *string                                    `json:"productType,omitempty" xml:"productType,omitempty"`
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

func (s *BatchRegisterDeviceRequestDevices) SetExtraData(v map[string]interface{}) *BatchRegisterDeviceRequestDevices {
	s.ExtraData = v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetLiveUrls(v *BatchRegisterDeviceRequestDevicesLiveUrls) *BatchRegisterDeviceRequestDevices {
	s.LiveUrls = v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetLocation(v string) *BatchRegisterDeviceRequestDevices {
	s.Location = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetParentId(v string) *BatchRegisterDeviceRequestDevices {
	s.ParentId = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevices) SetProductType(v string) *BatchRegisterDeviceRequestDevices {
	s.ProductType = &v
	return s
}

type BatchRegisterDeviceRequestDevicesLiveUrls struct {
	Flv  *string `json:"flv,omitempty" xml:"flv,omitempty"`
	Hls  *string `json:"hls,omitempty" xml:"hls,omitempty"`
	Rtmp *string `json:"rtmp,omitempty" xml:"rtmp,omitempty"`
}

func (s BatchRegisterDeviceRequestDevicesLiveUrls) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequestDevicesLiveUrls) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequestDevicesLiveUrls) SetFlv(v string) *BatchRegisterDeviceRequestDevicesLiveUrls {
	s.Flv = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevicesLiveUrls) SetHls(v string) *BatchRegisterDeviceRequestDevicesLiveUrls {
	s.Hls = &v
	return s
}

func (s *BatchRegisterDeviceRequestDevicesLiveUrls) SetRtmp(v string) *BatchRegisterDeviceRequestDevicesLiveUrls {
	s.Rtmp = &v
	return s
}

type BatchRegisterDeviceResponseBody struct {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchRegisterDeviceResponse) SetStatusCode(v int32) *BatchRegisterDeviceResponse {
	s.StatusCode = &v
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
	CorpId     *string                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	EventType     *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRegisterEventTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchRegisterEventTypeResponse) SetStatusCode(v int32) *BatchRegisterEventTypeResponse {
	s.StatusCode = &v
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
	CorpId  *string                            `json:"corpId,omitempty" xml:"corpId,omitempty"`
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
	DeviceId     *string                                  `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName   *string                                  `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceStatus *int32                                   `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	ExtraData    map[string]interface{}                   `json:"extraData,omitempty" xml:"extraData,omitempty"`
	LiveUrls     *BatchUpdateDeviceRequestDevicesLiveUrls `json:"liveUrls,omitempty" xml:"liveUrls,omitempty" type:"Struct"`
	Location     *string                                  `json:"location,omitempty" xml:"location,omitempty"`
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

func (s *BatchUpdateDeviceRequestDevices) SetDeviceStatus(v int32) *BatchUpdateDeviceRequestDevices {
	s.DeviceStatus = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetExtraData(v map[string]interface{}) *BatchUpdateDeviceRequestDevices {
	s.ExtraData = v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetLiveUrls(v *BatchUpdateDeviceRequestDevicesLiveUrls) *BatchUpdateDeviceRequestDevices {
	s.LiveUrls = v
	return s
}

func (s *BatchUpdateDeviceRequestDevices) SetLocation(v string) *BatchUpdateDeviceRequestDevices {
	s.Location = &v
	return s
}

type BatchUpdateDeviceRequestDevicesLiveUrls struct {
	Flv  *string `json:"flv,omitempty" xml:"flv,omitempty"`
	Hls  *string `json:"hls,omitempty" xml:"hls,omitempty"`
	Rtmp *string `json:"rtmp,omitempty" xml:"rtmp,omitempty"`
}

func (s BatchUpdateDeviceRequestDevicesLiveUrls) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateDeviceRequestDevicesLiveUrls) GoString() string {
	return s.String()
}

func (s *BatchUpdateDeviceRequestDevicesLiveUrls) SetFlv(v string) *BatchUpdateDeviceRequestDevicesLiveUrls {
	s.Flv = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevicesLiveUrls) SetHls(v string) *BatchUpdateDeviceRequestDevicesLiveUrls {
	s.Hls = &v
	return s
}

func (s *BatchUpdateDeviceRequestDevicesLiveUrls) SetRtmp(v string) *BatchUpdateDeviceRequestDevicesLiveUrls {
	s.Rtmp = &v
	return s
}

type BatchUpdateDeviceResponseBody struct {
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchUpdateDeviceResponse) SetStatusCode(v int32) *BatchUpdateDeviceResponse {
	s.StatusCode = &v
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
	AuthCode   *string                `json:"authCode,omitempty" xml:"authCode,omitempty"`
	ClientId   *string                `json:"clientId,omitempty" xml:"clientId,omitempty"`
	ClientName *string                `json:"clientName,omitempty" xml:"clientName,omitempty"`
	CorpId     *string                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ExtraData  map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
}

func (s BindSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSystemRequest) GoString() string {
	return s.String()
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

func (s *BindSystemRequest) SetCorpId(v string) *BindSystemRequest {
	s.CorpId = &v
	return s
}

func (s *BindSystemRequest) SetExtraData(v map[string]interface{}) *BindSystemRequest {
	s.ExtraData = v
	return s
}

type BindSystemResponseBody struct {
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s BindSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSystemResponseBody) GoString() string {
	return s.String()
}

func (s *BindSystemResponseBody) SetClientId(v string) *BindSystemResponseBody {
	s.ClientId = &v
	return s
}

func (s *BindSystemResponseBody) SetCorpId(v string) *BindSystemResponseBody {
	s.CorpId = &v
	return s
}

type BindSystemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BindSystemResponse) SetStatusCode(v int32) *BindSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSystemResponse) SetBody(v *BindSystemResponseBody) *BindSystemResponse {
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
	ConfTitle          *string   `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	ConferenceId       *string   `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	ConferencePassword *string   `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	DeviceIds          []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeviceConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeviceConferenceResponse) SetStatusCode(v int32) *DeviceConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceConferenceResponse) SetBody(v *DeviceConferenceResponseBody) *DeviceConferenceResponse {
	s.Body = v
	return s
}

type DiotSystemMarkTestResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DiotSystemMarkTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DiotSystemMarkTestResponseBody) GoString() string {
	return s.String()
}

func (s *DiotSystemMarkTestResponseBody) SetRequestId(v string) *DiotSystemMarkTestResponseBody {
	s.RequestId = &v
	return s
}

type DiotSystemMarkTestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DiotSystemMarkTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DiotSystemMarkTestResponse) String() string {
	return tea.Prettify(s)
}

func (s DiotSystemMarkTestResponse) GoString() string {
	return s.String()
}

func (s *DiotSystemMarkTestResponse) SetHeaders(v map[string]*string) *DiotSystemMarkTestResponse {
	s.Headers = v
	return s
}

func (s *DiotSystemMarkTestResponse) SetStatusCode(v int32) *DiotSystemMarkTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DiotSystemMarkTestResponse) SetBody(v *DiotSystemMarkTestResponseBody) *DiotSystemMarkTestResponse {
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
	CorpId         *string                `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceId       *string                `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	EventId        *string                `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventName      *string                `json:"eventName,omitempty" xml:"eventName,omitempty"`
	EventType      *string                `json:"eventType,omitempty" xml:"eventType,omitempty"`
	ExtraData      map[string]interface{} `json:"extraData,omitempty" xml:"extraData,omitempty"`
	Location       *string                `json:"location,omitempty" xml:"location,omitempty"`
	Msg            *string                `json:"msg,omitempty" xml:"msg,omitempty"`
	OccurrenceTime *int64                 `json:"occurrenceTime,omitempty" xml:"occurrenceTime,omitempty"`
	PicUrls        []*string              `json:"picUrls,omitempty" xml:"picUrls,omitempty" type:"Repeated"`
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

func (s *PushEventRequest) SetDeviceId(v string) *PushEventRequest {
	s.DeviceId = &v
	return s
}

func (s *PushEventRequest) SetEventId(v string) *PushEventRequest {
	s.EventId = &v
	return s
}

func (s *PushEventRequest) SetEventName(v string) *PushEventRequest {
	s.EventName = &v
	return s
}

func (s *PushEventRequest) SetEventType(v string) *PushEventRequest {
	s.EventType = &v
	return s
}

func (s *PushEventRequest) SetExtraData(v map[string]interface{}) *PushEventRequest {
	s.ExtraData = v
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

func (s *PushEventRequest) SetOccurrenceTime(v int64) *PushEventRequest {
	s.OccurrenceTime = &v
	return s
}

func (s *PushEventRequest) SetPicUrls(v []*string) *PushEventRequest {
	s.PicUrls = v
	return s
}

type PushEventResponseBody struct {
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PushEventResponse) SetStatusCode(v int32) *PushEventResponse {
	s.StatusCode = &v
	return s
}

func (s *PushEventResponse) SetBody(v *PushEventResponseBody) *PushEventResponse {
	s.Body = v
	return s
}

type QueryDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceRequest struct {
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceRequest) SetCorpId(v string) *QueryDeviceRequest {
	s.CorpId = &v
	return s
}

func (s *QueryDeviceRequest) SetPageNumber(v int64) *QueryDeviceRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDeviceRequest) SetPageSize(v int64) *QueryDeviceRequest {
	s.PageSize = &v
	return s
}

type QueryDeviceResponseBody struct {
	Data       []*QueryDeviceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                         `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                         `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceResponseBody) SetData(v []*QueryDeviceResponseBodyData) *QueryDeviceResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceResponseBody) SetPageNumber(v int64) *QueryDeviceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDeviceResponseBody) SetPageSize(v int64) *QueryDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDeviceResponseBody) SetTotalCount(v int64) *QueryDeviceResponseBody {
	s.TotalCount = &v
	return s
}

type QueryDeviceResponseBodyData struct {
	DeviceId       *string                              `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName     *string                              `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceStatus   *int64                               `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	DeviceType     *string                              `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	DeviceTypeName *string                              `json:"deviceTypeName,omitempty" xml:"deviceTypeName,omitempty"`
	LiveUrls       *QueryDeviceResponseBodyDataLiveUrls `json:"liveUrls,omitempty" xml:"liveUrls,omitempty" type:"Struct"`
	Location       *string                              `json:"location,omitempty" xml:"location,omitempty"`
	ParentId       *string                              `json:"parentId,omitempty" xml:"parentId,omitempty"`
	ProductType    *string                              `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s QueryDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceResponseBodyData) SetDeviceId(v string) *QueryDeviceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetDeviceName(v string) *QueryDeviceResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetDeviceStatus(v int64) *QueryDeviceResponseBodyData {
	s.DeviceStatus = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetDeviceType(v string) *QueryDeviceResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetDeviceTypeName(v string) *QueryDeviceResponseBodyData {
	s.DeviceTypeName = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetLiveUrls(v *QueryDeviceResponseBodyDataLiveUrls) *QueryDeviceResponseBodyData {
	s.LiveUrls = v
	return s
}

func (s *QueryDeviceResponseBodyData) SetLocation(v string) *QueryDeviceResponseBodyData {
	s.Location = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetParentId(v string) *QueryDeviceResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *QueryDeviceResponseBodyData) SetProductType(v string) *QueryDeviceResponseBodyData {
	s.ProductType = &v
	return s
}

type QueryDeviceResponseBodyDataLiveUrls struct {
	Flv  *string `json:"flv,omitempty" xml:"flv,omitempty"`
	Hls  *string `json:"hls,omitempty" xml:"hls,omitempty"`
	Rtmp *string `json:"rtmp,omitempty" xml:"rtmp,omitempty"`
}

func (s QueryDeviceResponseBodyDataLiveUrls) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceResponseBodyDataLiveUrls) GoString() string {
	return s.String()
}

func (s *QueryDeviceResponseBodyDataLiveUrls) SetFlv(v string) *QueryDeviceResponseBodyDataLiveUrls {
	s.Flv = &v
	return s
}

func (s *QueryDeviceResponseBodyDataLiveUrls) SetHls(v string) *QueryDeviceResponseBodyDataLiveUrls {
	s.Hls = &v
	return s
}

func (s *QueryDeviceResponseBodyDataLiveUrls) SetRtmp(v string) *QueryDeviceResponseBodyDataLiveUrls {
	s.Rtmp = &v
	return s
}

type QueryDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceResponse) SetHeaders(v map[string]*string) *QueryDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceResponse) SetStatusCode(v int32) *QueryDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceResponse) SetBody(v *QueryDeviceResponseBody) *QueryDeviceResponse {
	s.Body = v
	return s
}

type QueryEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEventHeaders) GoString() string {
	return s.String()
}

func (s *QueryEventHeaders) SetCommonHeaders(v map[string]*string) *QueryEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEventHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEventRequest struct {
	CorpId          *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceIdList    []*string `json:"deviceIdList,omitempty" xml:"deviceIdList,omitempty" type:"Repeated"`
	EndTime         *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EventId         *string   `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventStatusList []*int64  `json:"eventStatusList,omitempty" xml:"eventStatusList,omitempty" type:"Repeated"`
	EventTypeList   []*string `json:"eventTypeList,omitempty" xml:"eventTypeList,omitempty" type:"Repeated"`
	PageNumber      *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime       *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryEventRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEventRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRequest) SetCorpId(v string) *QueryEventRequest {
	s.CorpId = &v
	return s
}

func (s *QueryEventRequest) SetDeviceIdList(v []*string) *QueryEventRequest {
	s.DeviceIdList = v
	return s
}

func (s *QueryEventRequest) SetEndTime(v int64) *QueryEventRequest {
	s.EndTime = &v
	return s
}

func (s *QueryEventRequest) SetEventId(v string) *QueryEventRequest {
	s.EventId = &v
	return s
}

func (s *QueryEventRequest) SetEventStatusList(v []*int64) *QueryEventRequest {
	s.EventStatusList = v
	return s
}

func (s *QueryEventRequest) SetEventTypeList(v []*string) *QueryEventRequest {
	s.EventTypeList = v
	return s
}

func (s *QueryEventRequest) SetPageNumber(v int64) *QueryEventRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryEventRequest) SetPageSize(v int64) *QueryEventRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEventRequest) SetStartTime(v int64) *QueryEventRequest {
	s.StartTime = &v
	return s
}

type QueryEventResponseBody struct {
	Data       []*QueryEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageNumber *int64                        `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEventResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEventResponseBody) SetData(v []*QueryEventResponseBodyData) *QueryEventResponseBody {
	s.Data = v
	return s
}

func (s *QueryEventResponseBody) SetPageNumber(v int64) *QueryEventResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryEventResponseBody) SetPageSize(v int64) *QueryEventResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryEventResponseBody) SetTotalCount(v int64) *QueryEventResponseBody {
	s.TotalCount = &v
	return s
}

type QueryEventResponseBodyData struct {
	EventId        *string   `json:"eventId,omitempty" xml:"eventId,omitempty"`
	EventName      *string   `json:"eventName,omitempty" xml:"eventName,omitempty"`
	EventStatus    *int64    `json:"eventStatus,omitempty" xml:"eventStatus,omitempty"`
	EventType      *string   `json:"eventType,omitempty" xml:"eventType,omitempty"`
	Msg            *string   `json:"msg,omitempty" xml:"msg,omitempty"`
	OccurrenceTime *int64    `json:"occurrenceTime,omitempty" xml:"occurrenceTime,omitempty"`
	PicUrls        []*string `json:"picUrls,omitempty" xml:"picUrls,omitempty" type:"Repeated"`
}

func (s QueryEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEventResponseBodyData) SetEventId(v string) *QueryEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *QueryEventResponseBodyData) SetEventName(v string) *QueryEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *QueryEventResponseBodyData) SetEventStatus(v int64) *QueryEventResponseBodyData {
	s.EventStatus = &v
	return s
}

func (s *QueryEventResponseBodyData) SetEventType(v string) *QueryEventResponseBodyData {
	s.EventType = &v
	return s
}

func (s *QueryEventResponseBodyData) SetMsg(v string) *QueryEventResponseBodyData {
	s.Msg = &v
	return s
}

func (s *QueryEventResponseBodyData) SetOccurrenceTime(v int64) *QueryEventResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *QueryEventResponseBodyData) SetPicUrls(v []*string) *QueryEventResponseBodyData {
	s.PicUrls = v
	return s
}

type QueryEventResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEventResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEventResponse) GoString() string {
	return s.String()
}

func (s *QueryEventResponse) SetHeaders(v map[string]*string) *QueryEventResponse {
	s.Headers = v
	return s
}

func (s *QueryEventResponse) SetStatusCode(v int32) *QueryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventResponse) SetBody(v *QueryEventResponseBody) *QueryEventResponse {
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
	CorpId         *string                        `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceName     *string                        `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceStatus   *int32                         `json:"deviceStatus,omitempty" xml:"deviceStatus,omitempty"`
	DeviceType     *string                        `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	DeviceTypeName *string                        `json:"deviceTypeName,omitempty" xml:"deviceTypeName,omitempty"`
	Id             *string                        `json:"id,omitempty" xml:"id,omitempty"`
	LiveUrls       *RegisterDeviceRequestLiveUrls `json:"liveUrls,omitempty" xml:"liveUrls,omitempty" type:"Struct"`
	Location       *string                        `json:"location,omitempty" xml:"location,omitempty"`
	NickName       *string                        `json:"nickName,omitempty" xml:"nickName,omitempty"`
	ParentId       *string                        `json:"parentId,omitempty" xml:"parentId,omitempty"`
	ProductType    *string                        `json:"productType,omitempty" xml:"productType,omitempty"`
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

func (s *RegisterDeviceRequest) SetDeviceName(v string) *RegisterDeviceRequest {
	s.DeviceName = &v
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

func (s *RegisterDeviceRequest) SetId(v string) *RegisterDeviceRequest {
	s.Id = &v
	return s
}

func (s *RegisterDeviceRequest) SetLiveUrls(v *RegisterDeviceRequestLiveUrls) *RegisterDeviceRequest {
	s.LiveUrls = v
	return s
}

func (s *RegisterDeviceRequest) SetLocation(v string) *RegisterDeviceRequest {
	s.Location = &v
	return s
}

func (s *RegisterDeviceRequest) SetNickName(v string) *RegisterDeviceRequest {
	s.NickName = &v
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

type RegisterDeviceRequestLiveUrls struct {
	Flv  *string `json:"flv,omitempty" xml:"flv,omitempty"`
	Hls  *string `json:"hls,omitempty" xml:"hls,omitempty"`
	Rtmp *string `json:"rtmp,omitempty" xml:"rtmp,omitempty"`
}

func (s RegisterDeviceRequestLiveUrls) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequestLiveUrls) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequestLiveUrls) SetFlv(v string) *RegisterDeviceRequestLiveUrls {
	s.Flv = &v
	return s
}

func (s *RegisterDeviceRequestLiveUrls) SetHls(v string) *RegisterDeviceRequestLiveUrls {
	s.Hls = &v
	return s
}

func (s *RegisterDeviceRequestLiveUrls) SetRtmp(v string) *RegisterDeviceRequestLiveUrls {
	s.Rtmp = &v
	return s
}

type RegisterDeviceResponseBody struct {
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RegisterDeviceResponse) SetStatusCode(v int32) *RegisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type WorkbenchTransformInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s WorkbenchTransformInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WorkbenchTransformInfoResponseBody) GoString() string {
	return s.String()
}

func (s *WorkbenchTransformInfoResponseBody) SetRequestId(v string) *WorkbenchTransformInfoResponseBody {
	s.RequestId = &v
	return s
}

type WorkbenchTransformInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WorkbenchTransformInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WorkbenchTransformInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s WorkbenchTransformInfoResponse) GoString() string {
	return s.String()
}

func (s *WorkbenchTransformInfoResponse) SetHeaders(v map[string]*string) *WorkbenchTransformInfoResponse {
	s.Headers = v
	return s
}

func (s *WorkbenchTransformInfoResponse) SetStatusCode(v int32) *WorkbenchTransformInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *WorkbenchTransformInfoResponse) SetBody(v *WorkbenchTransformInfoResponseBody) *WorkbenchTransformInfoResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) AyunOnlienTestWithOptions(request *AyunOnlienTestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AyunOnlienTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["reqId"] = request.ReqId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AyunOnlienTest"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/ayunTest"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AyunOnlienTestResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AyunOnlienTest(request *AyunOnlienTestRequest) (_result *AyunOnlienTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AyunOnlienTestResponse{}
	_body, _err := client.AyunOnlienTestWithOptions(request, headers, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteDevice"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/devices/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRegisterDevice"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/devices/registrations/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRegisterEventType"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/eventTypes/registrations/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRegisterEventTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateDevice"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/devices/batch"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) BindSystemWithOptions(request *BindSystemRequest, headers *BindSystemHeaders, runtime *util.RuntimeOptions) (_result *BindSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientName)) {
		body["clientName"] = request.ClientName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraData)) {
		body["extraData"] = request.ExtraData
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
		Action:      tea.String("BindSystem"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/systems/bind"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BindSystemResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeviceConference"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/deviceConferences/initiate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) DiotSystemMarkTestWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DiotSystemMarkTestResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DiotSystemMarkTest"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/sys/mark/test"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DiotSystemMarkTestResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DiotSystemMarkTest() (_result *DiotSystemMarkTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DiotSystemMarkTestResponse{}
	_body, _err := client.DiotSystemMarkTestWithOptions(headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		body["eventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["eventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraData)) {
		body["extraData"] = request.ExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		body["msg"] = request.Msg
	}

	if !tea.BoolValue(util.IsUnset(request.OccurrenceTime)) {
		body["occurrenceTime"] = request.OccurrenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.PicUrls)) {
		body["picUrls"] = request.PicUrls
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
		Action:      tea.String("PushEvent"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/events/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PushEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) QueryDeviceWithOptions(request *QueryDeviceRequest, headers *QueryDeviceHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryDevice"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/devices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevice(request *QueryDeviceRequest) (_result *QueryDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceHeaders{}
	_result = &QueryDeviceResponse{}
	_body, _err := client.QueryDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEventWithOptions(request *QueryEventRequest, headers *QueryEventHeaders, runtime *util.RuntimeOptions) (_result *QueryEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceIdList)) {
		body["deviceIdList"] = request.DeviceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventId)) {
		body["eventId"] = request.EventId
	}

	if !tea.BoolValue(util.IsUnset(request.EventStatusList)) {
		body["eventStatusList"] = request.EventStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.EventTypeList)) {
		body["eventTypeList"] = request.EventTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("QueryEvent"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/events/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEvent(request *QueryEventRequest) (_result *QueryEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEventHeaders{}
	_result = &QueryEventResponse{}
	_body, _err := client.QueryEventWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
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

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LiveUrls)) {
		body["liveUrls"] = request.LiveUrls
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
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
		Action:      tea.String("RegisterDevice"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/devices/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) WorkbenchTransformInfoWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *WorkbenchTransformInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("WorkbenchTransformInfo"),
		Version:     tea.String("diot_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/diot/workbench/transform"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WorkbenchTransformInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WorkbenchTransformInfo() (_result *WorkbenchTransformInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WorkbenchTransformInfoResponse{}
	_body, _err := client.WorkbenchTransformInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
