// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package conference_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 会议发起人UID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 会议主题： 文字，不超过20中文
	ConfTitle *string `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	// 邀请参会人员UID列表（必须好友或同事）
	InviteUserIds []*string `json:"inviteUserIds,omitempty" xml:"inviteUserIds,omitempty" type:"Repeated"`
}

func (s CreateVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceRequest) SetUserId(v string) *CreateVideoConferenceRequest {
	s.UserId = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetConfTitle(v string) *CreateVideoConferenceRequest {
	s.ConfTitle = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteUserIds(v []*string) *CreateVideoConferenceRequest {
	s.InviteUserIds = v
	return s
}

type CreateVideoConferenceResponseBody struct {
	// conferenceId
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 会议密码
	ConferencePassword *string `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	// 主持人密码
	HostPassword *string `json:"hostPassword,omitempty" xml:"hostPassword,omitempty"`
	// 入会链接
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	// 电话入会号码
	PhoneNumbers []*string `json:"phoneNumbers,omitempty" xml:"phoneNumbers,omitempty" type:"Repeated"`
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

func (s *CreateVideoConferenceResponseBody) SetHostPassword(v string) *CreateVideoConferenceResponseBody {
	s.HostPassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetExternalLinkUrl(v string) *CreateVideoConferenceResponseBody {
	s.ExternalLinkUrl = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetPhoneNumbers(v []*string) *CreateVideoConferenceResponseBody {
	s.PhoneNumbers = v
	return s
}

type CreateVideoConferenceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVideoConferenceResponse) SetBody(v *CreateVideoConferenceResponseBody) *CreateVideoConferenceResponse {
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
	// 失败原因
	Case *string `json:"case,omitempty" xml:"case,omitempty"`
	// 返回编码
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVideoConferenceExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateVideoConferenceExtInfoResponse) SetBody(v *UpdateVideoConferenceExtInfoResponseBody) *UpdateVideoConferenceExtInfoResponse {
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
	// 员工在当前开发者企业账号范围内的唯一标识
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
	Code  *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
}

func (s CloseVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponseBody) SetCode(v int64) *CloseVideoConferenceResponseBody {
	s.Code = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetCause(v string) *CloseVideoConferenceResponseBody {
	s.Cause = &v
	return s
}

type CloseVideoConferenceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CloseVideoConferenceResponse) SetBody(v *CloseVideoConferenceResponseBody) *CloseVideoConferenceResponse {
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
	// 会议详情列表
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
	// 会议iD
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 会议名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 会议开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 会议状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 媒体状态
	MediaStatus *int64 `json:"mediaStatus,omitempty" xml:"mediaStatus,omitempty"`
	// 参会用户列表
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

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetTitle(v string) *QueryConferenceInfoBatchResponseBodyInfos {
	s.Title = &v
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

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetMediaStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.MediaStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetUserList(v []*QueryConferenceInfoBatchResponseBodyInfosUserList) *QueryConferenceInfoBatchResponseBodyInfos {
	s.UserList = v
	return s
}

type QueryConferenceInfoBatchResponseBodyInfosUserList struct {
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 名称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 在会状态
	AttendStatus *int64 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	// 摄像头状态
	CameraStatus *int64 `json:"cameraStatus,omitempty" xml:"cameraStatus,omitempty"`
	// 麦克风状态
	MicStatus *int64 `json:"micStatus,omitempty" xml:"micStatus,omitempty"`
	// 拒绝原因
	RejectDescription *string `json:"rejectDescription,omitempty" xml:"rejectDescription,omitempty"`
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetUserId(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.UserId = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetNick(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.Nick = &v
	return s
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

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetRejectDescription(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.RejectDescription = &v
	return s
}

type QueryConferenceInfoBatchResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConferenceInfoBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryConferenceInfoBatchResponse) SetBody(v *QueryConferenceInfoBatchResponseBody) *QueryConferenceInfoBatchResponse {
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

func (client *Client) CreateVideoConferenceWithOptions(request *CreateVideoConferenceRequest, headers *CreateVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfTitle)) {
		body["confTitle"] = request.ConfTitle
	}

	if !tea.BoolValue(util.IsUnset(request.InviteUserIds)) {
		body["inviteUserIds"] = request.InviteUserIds
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
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateVideoConference"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateVideoConferenceExtInfoWithOptions(conferenceId *string, headers *UpdateVideoConferenceExtInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateVideoConferenceExtInfoResponse, _err error) {
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
	_result = &UpdateVideoConferenceExtInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateVideoConferenceExtInfo"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/extInfo"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("CloseVideoConference"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryConferenceInfoBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryConferenceInfoBatch"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
