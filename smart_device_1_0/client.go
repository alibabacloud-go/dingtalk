// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package smart_device_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 入会口令
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 会议id
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDeviceVideoConferenceBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryDeviceVideoConferenceBookResponse) SetBody(v *QueryDeviceVideoConferenceBookResponseBody) *QueryDeviceVideoConferenceBookResponse {
	s.Body = v
	return s
}

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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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
	// 会议id
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 入会口令
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CreateDeviceVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceVideoConferenceResponseBody) SetConferenceId(v string) *CreateDeviceVideoConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *CreateDeviceVideoConferenceResponseBody) SetCode(v string) *CreateDeviceVideoConferenceResponseBody {
	s.Code = &v
	return s
}

type CreateDeviceVideoConferenceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDeviceVideoConferenceResponse) SetBody(v *CreateDeviceVideoConferenceResponseBody) *CreateDeviceVideoConferenceResponse {
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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
	Userid  *string `json:"userid,omitempty" xml:"userid,omitempty"`
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s ExtractFacialFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtractFacialFeatureRequest) GoString() string {
	return s.String()
}

func (s *ExtractFacialFeatureRequest) SetUserid(v string) *ExtractFacialFeatureRequest {
	s.Userid = &v
	return s
}

func (s *ExtractFacialFeatureRequest) SetMediaId(v string) *ExtractFacialFeatureRequest {
	s.MediaId = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExtractFacialFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ExtractFacialFeatureResponse) SetBody(v *ExtractFacialFeatureResponseBody) *ExtractFacialFeatureResponse {
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

func (client *Client) QueryDeviceVideoConferenceBookWithOptions(deviceId *string, bookId *string, headers *QueryDeviceVideoConferenceBookHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceVideoConferenceBookResponse, _err error) {
	deviceId = openapiutil.GetEncodeParam(deviceId)
	bookId = openapiutil.GetEncodeParam(bookId)
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
	_result = &QueryDeviceVideoConferenceBookResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryDeviceVideoConferenceBook"), tea.String("smartDevice_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/smartDevice/devices/"+tea.StringValue(deviceId)+"/books/"+tea.StringValue(bookId)), tea.String("json"), req, runtime)
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

func (client *Client) AddDeviceVideoConferenceMembersWithOptions(deviceId *string, conferenceId *string, request *AddDeviceVideoConferenceMembersRequest, headers *AddDeviceVideoConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *AddDeviceVideoConferenceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deviceId = openapiutil.GetEncodeParam(deviceId)
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
	_result = &AddDeviceVideoConferenceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddDeviceVideoConferenceMembers"), tea.String("smartDevice_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/smartDevice/devices/"+tea.StringValue(deviceId)+"/videoConferences/"+tea.StringValue(conferenceId)+"/members"), tea.String("none"), req, runtime)
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

func (client *Client) CreateDeviceVideoConferenceWithOptions(deviceId *string, request *CreateDeviceVideoConferenceRequest, headers *CreateDeviceVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateDeviceVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deviceId = openapiutil.GetEncodeParam(deviceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
	_result = &CreateDeviceVideoConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDeviceVideoConference"), tea.String("smartDevice_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/smartDevice/devices/"+tea.StringValue(deviceId)+"/videoConferences"), tea.String("json"), req, runtime)
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

func (client *Client) KickDeviceVideoConferenceMembersWithOptions(deviceId *string, conferenceId *string, request *KickDeviceVideoConferenceMembersRequest, headers *KickDeviceVideoConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *KickDeviceVideoConferenceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	deviceId = openapiutil.GetEncodeParam(deviceId)
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
	_result = &KickDeviceVideoConferenceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("KickDeviceVideoConferenceMembers"), tea.String("smartDevice_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/smartDevice/devices/"+tea.StringValue(deviceId)+"/videoConferences/"+tea.StringValue(conferenceId)+"/members/batchDelete"), tea.String("none"), req, runtime)
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

func (client *Client) ExtractFacialFeatureWithOptions(request *ExtractFacialFeatureRequest, headers *ExtractFacialFeatureHeaders, runtime *util.RuntimeOptions) (_result *ExtractFacialFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Userid)) {
		body["userid"] = request.Userid
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
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
	_result = &ExtractFacialFeatureResponse{}
	_body, _err := client.DoROARequest(tea.String("ExtractFacialFeature"), tea.String("smartDevice_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/smartDevice/faceRecognitions/features/extract"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
