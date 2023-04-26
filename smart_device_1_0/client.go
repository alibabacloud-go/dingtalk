// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package smart_device_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDeviceVideoConferenceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddDeviceVideoConferenceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceVideoConferenceMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddDeviceVideoConferenceMembersHeaders) SetCommonHeaders(v map[string]*string) *AddDeviceVideoConferenceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDeviceVideoConferenceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *AddDeviceVideoConferenceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddDeviceVideoConferenceMembersRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddDeviceVideoConferenceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceVideoConferenceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceVideoConferenceMembersRequest) SetUserIds(v []*string) *AddDeviceVideoConferenceMembersRequest {
	s.UserIds = v
	return s
}

type AddDeviceVideoConferenceMembersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s AddDeviceVideoConferenceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceVideoConferenceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceVideoConferenceMembersResponse) SetHeaders(v map[string]*string) *AddDeviceVideoConferenceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceVideoConferenceMembersResponse) SetStatusCode(v int32) *AddDeviceVideoConferenceMembersResponse {
	s.StatusCode = &v
	return s
}

type CreateDeviceVideoConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeviceVideoConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeviceVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateDeviceVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeviceVideoConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeviceVideoConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeviceVideoConferenceRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateDeviceVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceVideoConferenceRequest) SetUserIds(v []*string) *CreateDeviceVideoConferenceRequest {
	s.UserIds = v
	return s
}

type CreateDeviceVideoConferenceResponseBody struct {
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s CreateDeviceVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceVideoConferenceResponseBody) SetCode(v string) *CreateDeviceVideoConferenceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDeviceVideoConferenceResponseBody) SetConferenceId(v string) *CreateDeviceVideoConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

type CreateDeviceVideoConferenceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeviceVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceVideoConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceVideoConferenceResponse) SetHeaders(v map[string]*string) *CreateDeviceVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceVideoConferenceResponse) SetStatusCode(v int32) *CreateDeviceVideoConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceVideoConferenceResponse) SetBody(v *CreateDeviceVideoConferenceResponseBody) *CreateDeviceVideoConferenceResponse {
	s.Body = v
	return s
}

type ExtractFacialFeatureHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExtractFacialFeatureHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExtractFacialFeatureHeaders) GoString() string {
	return s.String()
}

func (s *ExtractFacialFeatureHeaders) SetCommonHeaders(v map[string]*string) *ExtractFacialFeatureHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExtractFacialFeatureHeaders) SetXAcsDingtalkAccessToken(v string) *ExtractFacialFeatureHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExtractFacialFeatureRequest struct {
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Userid  *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s ExtractFacialFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractFacialFeatureRequest) GoString() string {
	return s.String()
}

func (s *ExtractFacialFeatureRequest) SetMediaId(v string) *ExtractFacialFeatureRequest {
	s.MediaId = &v
	return s
}

func (s *ExtractFacialFeatureRequest) SetUserid(v string) *ExtractFacialFeatureRequest {
	s.Userid = &v
	return s
}

type ExtractFacialFeatureResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExtractFacialFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtractFacialFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *ExtractFacialFeatureResponseBody) SetResult(v bool) *ExtractFacialFeatureResponseBody {
	s.Result = &v
	return s
}

type ExtractFacialFeatureResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExtractFacialFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExtractFacialFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtractFacialFeatureResponse) GoString() string {
	return s.String()
}

func (s *ExtractFacialFeatureResponse) SetHeaders(v map[string]*string) *ExtractFacialFeatureResponse {
	s.Headers = v
	return s
}

func (s *ExtractFacialFeatureResponse) SetStatusCode(v int32) *ExtractFacialFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtractFacialFeatureResponse) SetBody(v *ExtractFacialFeatureResponseBody) *ExtractFacialFeatureResponse {
	s.Body = v
	return s
}

type KickDeviceVideoConferenceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s KickDeviceVideoConferenceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s KickDeviceVideoConferenceMembersHeaders) GoString() string {
	return s.String()
}

func (s *KickDeviceVideoConferenceMembersHeaders) SetCommonHeaders(v map[string]*string) *KickDeviceVideoConferenceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *KickDeviceVideoConferenceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *KickDeviceVideoConferenceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type KickDeviceVideoConferenceMembersRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s KickDeviceVideoConferenceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s KickDeviceVideoConferenceMembersRequest) GoString() string {
	return s.String()
}

func (s *KickDeviceVideoConferenceMembersRequest) SetUserIds(v []*string) *KickDeviceVideoConferenceMembersRequest {
	s.UserIds = v
	return s
}

type KickDeviceVideoConferenceMembersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s KickDeviceVideoConferenceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s KickDeviceVideoConferenceMembersResponse) GoString() string {
	return s.String()
}

func (s *KickDeviceVideoConferenceMembersResponse) SetHeaders(v map[string]*string) *KickDeviceVideoConferenceMembersResponse {
	s.Headers = v
	return s
}

func (s *KickDeviceVideoConferenceMembersResponse) SetStatusCode(v int32) *KickDeviceVideoConferenceMembersResponse {
	s.StatusCode = &v
	return s
}

type MachineManagerUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MachineManagerUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s MachineManagerUpdateHeaders) GoString() string {
	return s.String()
}

func (s *MachineManagerUpdateHeaders) SetCommonHeaders(v map[string]*string) *MachineManagerUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MachineManagerUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *MachineManagerUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MachineManagerUpdateRequest struct {
	AtmManagerRightMap *MachineManagerUpdateRequestAtmManagerRightMap `json:"atmManagerRightMap,omitempty" xml:"atmManagerRightMap,omitempty" type:"Struct"`
	DeviceId           *int64                                         `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	ScopeDeptIds       []*int64                                       `json:"scopeDeptIds,omitempty" xml:"scopeDeptIds,omitempty" type:"Repeated"`
	UserId             *string                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MachineManagerUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s MachineManagerUpdateRequest) GoString() string {
	return s.String()
}

func (s *MachineManagerUpdateRequest) SetAtmManagerRightMap(v *MachineManagerUpdateRequestAtmManagerRightMap) *MachineManagerUpdateRequest {
	s.AtmManagerRightMap = v
	return s
}

func (s *MachineManagerUpdateRequest) SetDeviceId(v int64) *MachineManagerUpdateRequest {
	s.DeviceId = &v
	return s
}

func (s *MachineManagerUpdateRequest) SetScopeDeptIds(v []*int64) *MachineManagerUpdateRequest {
	s.ScopeDeptIds = v
	return s
}

func (s *MachineManagerUpdateRequest) SetUserId(v string) *MachineManagerUpdateRequest {
	s.UserId = &v
	return s
}

type MachineManagerUpdateRequestAtmManagerRightMap struct {
	AttendancePersonManage *bool `json:"attendancePersonManage,omitempty" xml:"attendancePersonManage,omitempty"`
	BluetoothPunchManage   *bool `json:"bluetoothPunchManage,omitempty" xml:"bluetoothPunchManage,omitempty"`
	DeviceReset            *bool `json:"deviceReset,omitempty" xml:"deviceReset,omitempty"`
	DeviceSettings         *bool `json:"deviceSettings,omitempty" xml:"deviceSettings,omitempty"`
	FacePunchManage        *bool `json:"facePunchManage,omitempty" xml:"facePunchManage,omitempty"`
	FingerPunchManage      *bool `json:"fingerPunchManage,omitempty" xml:"fingerPunchManage,omitempty"`
}

func (s MachineManagerUpdateRequestAtmManagerRightMap) String() string {
	return tea.Prettify(s)
}

func (s MachineManagerUpdateRequestAtmManagerRightMap) GoString() string {
	return s.String()
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetAttendancePersonManage(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.AttendancePersonManage = &v
	return s
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetBluetoothPunchManage(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.BluetoothPunchManage = &v
	return s
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetDeviceReset(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.DeviceReset = &v
	return s
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetDeviceSettings(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.DeviceSettings = &v
	return s
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetFacePunchManage(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.FacePunchManage = &v
	return s
}

func (s *MachineManagerUpdateRequestAtmManagerRightMap) SetFingerPunchManage(v bool) *MachineManagerUpdateRequestAtmManagerRightMap {
	s.FingerPunchManage = &v
	return s
}

type MachineManagerUpdateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s MachineManagerUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s MachineManagerUpdateResponse) GoString() string {
	return s.String()
}

func (s *MachineManagerUpdateResponse) SetHeaders(v map[string]*string) *MachineManagerUpdateResponse {
	s.Headers = v
	return s
}

func (s *MachineManagerUpdateResponse) SetStatusCode(v int32) *MachineManagerUpdateResponse {
	s.StatusCode = &v
	return s
}

type MachineUsersUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MachineUsersUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s MachineUsersUpdateHeaders) GoString() string {
	return s.String()
}

func (s *MachineUsersUpdateHeaders) SetCommonHeaders(v map[string]*string) *MachineUsersUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MachineUsersUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *MachineUsersUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MachineUsersUpdateRequest struct {
	AddDeptIds []*int64  `json:"addDeptIds,omitempty" xml:"addDeptIds,omitempty" type:"Repeated"`
	AddUserIds []*string `json:"addUserIds,omitempty" xml:"addUserIds,omitempty" type:"Repeated"`
	DelDeptIds []*int64  `json:"delDeptIds,omitempty" xml:"delDeptIds,omitempty" type:"Repeated"`
	DelUserIds []*string `json:"delUserIds,omitempty" xml:"delUserIds,omitempty" type:"Repeated"`
	DevIds     []*int64  `json:"devIds,omitempty" xml:"devIds,omitempty" type:"Repeated"`
	DeviceIds  []*string `json:"deviceIds,omitempty" xml:"deviceIds,omitempty" type:"Repeated"`
}

func (s MachineUsersUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s MachineUsersUpdateRequest) GoString() string {
	return s.String()
}

func (s *MachineUsersUpdateRequest) SetAddDeptIds(v []*int64) *MachineUsersUpdateRequest {
	s.AddDeptIds = v
	return s
}

func (s *MachineUsersUpdateRequest) SetAddUserIds(v []*string) *MachineUsersUpdateRequest {
	s.AddUserIds = v
	return s
}

func (s *MachineUsersUpdateRequest) SetDelDeptIds(v []*int64) *MachineUsersUpdateRequest {
	s.DelDeptIds = v
	return s
}

func (s *MachineUsersUpdateRequest) SetDelUserIds(v []*string) *MachineUsersUpdateRequest {
	s.DelUserIds = v
	return s
}

func (s *MachineUsersUpdateRequest) SetDevIds(v []*int64) *MachineUsersUpdateRequest {
	s.DevIds = v
	return s
}

func (s *MachineUsersUpdateRequest) SetDeviceIds(v []*string) *MachineUsersUpdateRequest {
	s.DeviceIds = v
	return s
}

type MachineUsersUpdateResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s MachineUsersUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s MachineUsersUpdateResponse) GoString() string {
	return s.String()
}

func (s *MachineUsersUpdateResponse) SetHeaders(v map[string]*string) *MachineUsersUpdateResponse {
	s.Headers = v
	return s
}

func (s *MachineUsersUpdateResponse) SetStatusCode(v int32) *MachineUsersUpdateResponse {
	s.StatusCode = &v
	return s
}

type QueryDeviceVideoConferenceBookHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceVideoConferenceBookHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVideoConferenceBookHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceVideoConferenceBookHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceVideoConferenceBookHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceVideoConferenceBookHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceVideoConferenceBookHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceVideoConferenceBookResponseBody struct {
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryDeviceVideoConferenceBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVideoConferenceBookResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVideoConferenceBookResponseBody) SetCode(v string) *QueryDeviceVideoConferenceBookResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVideoConferenceBookResponseBody) SetConferenceId(v string) *QueryDeviceVideoConferenceBookResponseBody {
	s.ConferenceId = &v
	return s
}

type QueryDeviceVideoConferenceBookResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceVideoConferenceBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceVideoConferenceBookResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceVideoConferenceBookResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceVideoConferenceBookResponse) SetHeaders(v map[string]*string) *QueryDeviceVideoConferenceBookResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceVideoConferenceBookResponse) SetStatusCode(v int32) *QueryDeviceVideoConferenceBookResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceVideoConferenceBookResponse) SetBody(v *QueryDeviceVideoConferenceBookResponseBody) *QueryDeviceVideoConferenceBookResponse {
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

func (client *Client) AddDeviceVideoConferenceMembersWithOptions(deviceId *string, conferenceId *string, request *AddDeviceVideoConferenceMembersRequest, headers *AddDeviceVideoConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *AddDeviceVideoConferenceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("AddDeviceVideoConferenceMembers"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/devices/" + tea.StringValue(deviceId) + "/videoConferences/" + tea.StringValue(conferenceId) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &AddDeviceVideoConferenceMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDeviceVideoConferenceMembers(deviceId *string, conferenceId *string, request *AddDeviceVideoConferenceMembersRequest) (_result *AddDeviceVideoConferenceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddDeviceVideoConferenceMembersHeaders{}
	_result = &AddDeviceVideoConferenceMembersResponse{}
	_body, _err := client.AddDeviceVideoConferenceMembersWithOptions(deviceId, conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceVideoConferenceWithOptions(deviceId *string, request *CreateDeviceVideoConferenceRequest, headers *CreateDeviceVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateDeviceVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("CreateDeviceVideoConference"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/devices/" + tea.StringValue(deviceId) + "/videoConferences"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeviceVideoConferenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeviceVideoConference(deviceId *string, request *CreateDeviceVideoConferenceRequest) (_result *CreateDeviceVideoConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeviceVideoConferenceHeaders{}
	_result = &CreateDeviceVideoConferenceResponse{}
	_body, _err := client.CreateDeviceVideoConferenceWithOptions(deviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtractFacialFeatureWithOptions(request *ExtractFacialFeatureRequest, headers *ExtractFacialFeatureHeaders, runtime *util.RuntimeOptions) (_result *ExtractFacialFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Userid)) {
		body["userid"] = request.Userid
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
		Action:      tea.String("ExtractFacialFeature"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/faceRecognitions/features/extract"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtractFacialFeatureResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtractFacialFeature(request *ExtractFacialFeatureRequest) (_result *ExtractFacialFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExtractFacialFeatureHeaders{}
	_result = &ExtractFacialFeatureResponse{}
	_body, _err := client.ExtractFacialFeatureWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KickDeviceVideoConferenceMembersWithOptions(deviceId *string, conferenceId *string, request *KickDeviceVideoConferenceMembersRequest, headers *KickDeviceVideoConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *KickDeviceVideoConferenceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("KickDeviceVideoConferenceMembers"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/devices/" + tea.StringValue(deviceId) + "/videoConferences/" + tea.StringValue(conferenceId) + "/members/batchDelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &KickDeviceVideoConferenceMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KickDeviceVideoConferenceMembers(deviceId *string, conferenceId *string, request *KickDeviceVideoConferenceMembersRequest) (_result *KickDeviceVideoConferenceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &KickDeviceVideoConferenceMembersHeaders{}
	_result = &KickDeviceVideoConferenceMembersResponse{}
	_body, _err := client.KickDeviceVideoConferenceMembersWithOptions(deviceId, conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MachineManagerUpdateWithOptions(request *MachineManagerUpdateRequest, headers *MachineManagerUpdateHeaders, runtime *util.RuntimeOptions) (_result *MachineManagerUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtmManagerRightMap)) {
		body["atmManagerRightMap"] = request.AtmManagerRightMap
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeDeptIds)) {
		body["scopeDeptIds"] = request.ScopeDeptIds
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
		Action:      tea.String("MachineManagerUpdate"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/atmachines/managers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &MachineManagerUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MachineManagerUpdate(request *MachineManagerUpdateRequest) (_result *MachineManagerUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MachineManagerUpdateHeaders{}
	_result = &MachineManagerUpdateResponse{}
	_body, _err := client.MachineManagerUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MachineUsersUpdateWithOptions(request *MachineUsersUpdateRequest, headers *MachineUsersUpdateHeaders, runtime *util.RuntimeOptions) (_result *MachineUsersUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddDeptIds)) {
		body["addDeptIds"] = request.AddDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.AddUserIds)) {
		body["addUserIds"] = request.AddUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelDeptIds)) {
		body["delDeptIds"] = request.DelDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelUserIds)) {
		body["delUserIds"] = request.DelUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.DevIds)) {
		body["devIds"] = request.DevIds
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
		Action:      tea.String("MachineUsersUpdate"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/atmachines/users"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &MachineUsersUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MachineUsersUpdate(request *MachineUsersUpdateRequest) (_result *MachineUsersUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MachineUsersUpdateHeaders{}
	_result = &MachineUsersUpdateResponse{}
	_body, _err := client.MachineUsersUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDeviceVideoConferenceBookWithOptions(deviceId *string, bookId *string, headers *QueryDeviceVideoConferenceBookHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceVideoConferenceBookResponse, _err error) {
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
		Action:      tea.String("QueryDeviceVideoConferenceBook"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/devices/" + tea.StringValue(deviceId) + "/books/" + tea.StringValue(bookId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceVideoConferenceBookResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceVideoConferenceBook(deviceId *string, bookId *string) (_result *QueryDeviceVideoConferenceBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceVideoConferenceBookHeaders{}
	_result = &QueryDeviceVideoConferenceBookResponse{}
	_body, _err := client.QueryDeviceVideoConferenceBookWithOptions(deviceId, bookId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
