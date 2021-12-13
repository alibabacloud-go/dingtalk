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
