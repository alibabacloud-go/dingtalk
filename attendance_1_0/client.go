// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package attendance_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateApproveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateApproveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveHeaders) GoString() string {
	return s.String()
}

func (s *CreateApproveHeaders) SetCommonHeaders(v map[string]*string) *CreateApproveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateApproveHeaders) SetXAcsDingtalkAccessToken(v string) *CreateApproveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateApproveRequest struct {
	// 员工id
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
	// 审批单类型名称，最大长度20个字符
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// 子类型名称，最大长度20个字符
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// 审批单关联的打卡信息
	PunchParam *CreateApproveRequestPunchParam `json:"punchParam,omitempty" xml:"punchParam,omitempty" type:"Struct"`
	// 三方审批单id，全局唯一
	ApproveId *string `json:"approveId,omitempty" xml:"approveId,omitempty"`
	// 审批人员工id
	OpUserid *string `json:"opUserid,omitempty" xml:"opUserid,omitempty"`
}

func (s CreateApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveRequest) GoString() string {
	return s.String()
}

func (s *CreateApproveRequest) SetUserid(v string) *CreateApproveRequest {
	s.Userid = &v
	return s
}

func (s *CreateApproveRequest) SetTagName(v string) *CreateApproveRequest {
	s.TagName = &v
	return s
}

func (s *CreateApproveRequest) SetSubType(v string) *CreateApproveRequest {
	s.SubType = &v
	return s
}

func (s *CreateApproveRequest) SetPunchParam(v *CreateApproveRequestPunchParam) *CreateApproveRequest {
	s.PunchParam = v
	return s
}

func (s *CreateApproveRequest) SetApproveId(v string) *CreateApproveRequest {
	s.ApproveId = &v
	return s
}

func (s *CreateApproveRequest) SetOpUserid(v string) *CreateApproveRequest {
	s.OpUserid = &v
	return s
}

type CreateApproveRequestPunchParam struct {
	// 打卡时间，单位毫秒
	PunchTime *int64 `json:"punchTime,omitempty" xml:"punchTime,omitempty"`
	// 地理位置标识：wifi:ssid_macAddress ble: deviceId gps:longitude_latitude
	PositionId *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	// 地理位置类型：wifi/ble/gps
	PositionType *string `json:"positionType,omitempty" xml:"positionType,omitempty"`
	// 地理位置名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
}

func (s CreateApproveRequestPunchParam) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveRequestPunchParam) GoString() string {
	return s.String()
}

func (s *CreateApproveRequestPunchParam) SetPunchTime(v int64) *CreateApproveRequestPunchParam {
	s.PunchTime = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPositionId(v string) *CreateApproveRequestPunchParam {
	s.PositionId = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPositionType(v string) *CreateApproveRequestPunchParam {
	s.PositionType = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPositionName(v string) *CreateApproveRequestPunchParam {
	s.PositionName = &v
	return s
}

type CreateApproveResponseBody struct {
	// 返回结果
	DingtalkApproveId *string `json:"dingtalkApproveId,omitempty" xml:"dingtalkApproveId,omitempty"`
}

func (s CreateApproveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApproveResponseBody) SetDingtalkApproveId(v string) *CreateApproveResponseBody {
	s.DingtalkApproveId = &v
	return s
}

type CreateApproveResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApproveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveResponse) GoString() string {
	return s.String()
}

func (s *CreateApproveResponse) SetHeaders(v map[string]*string) *CreateApproveResponse {
	s.Headers = v
	return s
}

func (s *CreateApproveResponse) SetBody(v *CreateApproveResponseBody) *CreateApproveResponse {
	s.Body = v
	return s
}

type GetUserHolidaysHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHolidaysHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysHeaders) SetCommonHeaders(v map[string]*string) *GetUserHolidaysHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHolidaysHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHolidaysHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserHolidaysRequest struct {
	// 员工列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// 开始日期
	WorkDateFrom *int64 `json:"workDateFrom,omitempty" xml:"workDateFrom,omitempty"`
	// 结束日期
	WorkDateTo *int64 `json:"workDateTo,omitempty" xml:"workDateTo,omitempty"`
}

func (s GetUserHolidaysRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysRequest) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysRequest) SetUserIds(v []*string) *GetUserHolidaysRequest {
	s.UserIds = v
	return s
}

func (s *GetUserHolidaysRequest) SetWorkDateFrom(v int64) *GetUserHolidaysRequest {
	s.WorkDateFrom = &v
	return s
}

func (s *GetUserHolidaysRequest) SetWorkDateTo(v int64) *GetUserHolidaysRequest {
	s.WorkDateTo = &v
	return s
}

type GetUserHolidaysResponseBody struct {
	// 员工假期列表
	Result []*GetUserHolidaysResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetUserHolidaysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysResponseBody) SetResult(v []*GetUserHolidaysResponseBodyResult) *GetUserHolidaysResponseBody {
	s.Result = v
	return s
}

type GetUserHolidaysResponseBodyResult struct {
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 假期列表
	Holidays []*GetUserHolidaysResponseBodyResultHolidays `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
}

func (s GetUserHolidaysResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysResponseBodyResult) SetUserId(v string) *GetUserHolidaysResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetUserHolidaysResponseBodyResult) SetHolidays(v []*GetUserHolidaysResponseBodyResultHolidays) *GetUserHolidaysResponseBodyResult {
	s.Holidays = v
	return s
}

type GetUserHolidaysResponseBodyResultHolidays struct {
	// 日期
	WorkDate *int64 `json:"workDate,omitempty" xml:"workDate,omitempty"`
	// 假期名称
	HolidayName *string `json:"holidayName,omitempty" xml:"holidayName,omitempty"`
	// 假期类型，festival：法定节假日；rest：调休日；overtime：加班日；
	HolidayType *string `json:"holidayType,omitempty" xml:"holidayType,omitempty"`
	// 补休日，只有假期类型为加班日时才有值
	RealWorkDate *int64 `json:"realWorkDate,omitempty" xml:"realWorkDate,omitempty"`
}

func (s GetUserHolidaysResponseBodyResultHolidays) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponseBodyResultHolidays) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysResponseBodyResultHolidays) SetWorkDate(v int64) *GetUserHolidaysResponseBodyResultHolidays {
	s.WorkDate = &v
	return s
}

func (s *GetUserHolidaysResponseBodyResultHolidays) SetHolidayName(v string) *GetUserHolidaysResponseBodyResultHolidays {
	s.HolidayName = &v
	return s
}

func (s *GetUserHolidaysResponseBodyResultHolidays) SetHolidayType(v string) *GetUserHolidaysResponseBodyResultHolidays {
	s.HolidayType = &v
	return s
}

func (s *GetUserHolidaysResponseBodyResultHolidays) SetRealWorkDate(v int64) *GetUserHolidaysResponseBodyResultHolidays {
	s.RealWorkDate = &v
	return s
}

type GetUserHolidaysResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserHolidaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserHolidaysResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponse) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysResponse) SetHeaders(v map[string]*string) *GetUserHolidaysResponse {
	s.Headers = v
	return s
}

func (s *GetUserHolidaysResponse) SetBody(v *GetUserHolidaysResponseBody) *GetUserHolidaysResponse {
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

func (client *Client) CreateApprove(request *CreateApproveRequest) (_result *CreateApproveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateApproveHeaders{}
	_result = &CreateApproveResponse{}
	_body, _err := client.CreateApproveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApproveWithOptions(request *CreateApproveRequest, headers *CreateApproveHeaders, runtime *util.RuntimeOptions) (_result *CreateApproveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Userid)) {
		body["userid"] = request.Userid
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		body["subType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PunchParam))) {
		body["punchParam"] = request.PunchParam
	}

	if !tea.BoolValue(util.IsUnset(request.ApproveId)) {
		body["approveId"] = request.ApproveId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserid)) {
		body["opUserid"] = request.OpUserid
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
	_result = &CreateApproveResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateApprove"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/approves"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserHolidays(request *GetUserHolidaysRequest) (_result *GetUserHolidaysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHolidaysHeaders{}
	_result = &GetUserHolidaysResponse{}
	_body, _err := client.GetUserHolidaysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserHolidaysWithOptions(request *GetUserHolidaysRequest, headers *GetUserHolidaysHeaders, runtime *util.RuntimeOptions) (_result *GetUserHolidaysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkDateFrom)) {
		body["workDateFrom"] = request.WorkDateFrom
	}

	if !tea.BoolValue(util.IsUnset(request.WorkDateTo)) {
		body["workDateTo"] = request.WorkDateTo
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
	_result = &GetUserHolidaysResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserHolidays"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/holidays"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
