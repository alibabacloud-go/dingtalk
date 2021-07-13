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

type CheckClosingAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckClosingAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountHeaders) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountHeaders) SetCommonHeaders(v map[string]*string) *CheckClosingAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckClosingAccountHeaders) SetXAcsDingtalkAccessToken(v string) *CheckClosingAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckClosingAccountRequest struct {
	// 员工列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// 时间段
	UserTimeRange []*CheckClosingAccountRequestUserTimeRange `json:"userTimeRange,omitempty" xml:"userTimeRange,omitempty" type:"Repeated"`
	// 情景
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
}

func (s CheckClosingAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountRequest) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountRequest) SetUserIds(v []*string) *CheckClosingAccountRequest {
	s.UserIds = v
	return s
}

func (s *CheckClosingAccountRequest) SetUserTimeRange(v []*CheckClosingAccountRequestUserTimeRange) *CheckClosingAccountRequest {
	s.UserTimeRange = v
	return s
}

func (s *CheckClosingAccountRequest) SetBizCode(v string) *CheckClosingAccountRequest {
	s.BizCode = &v
	return s
}

type CheckClosingAccountRequestUserTimeRange struct {
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
}

func (s CheckClosingAccountRequestUserTimeRange) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountRequestUserTimeRange) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountRequestUserTimeRange) SetStartTime(v int64) *CheckClosingAccountRequestUserTimeRange {
	s.StartTime = &v
	return s
}

func (s *CheckClosingAccountRequestUserTimeRange) SetEndTime(v int64) *CheckClosingAccountRequestUserTimeRange {
	s.EndTime = &v
	return s
}

type CheckClosingAccountResponseBody struct {
	Mesage *string `json:"mesage,omitempty" xml:"mesage,omitempty"`
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Pass   *bool   `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s CheckClosingAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountResponseBody) SetMesage(v string) *CheckClosingAccountResponseBody {
	s.Mesage = &v
	return s
}

func (s *CheckClosingAccountResponseBody) SetCode(v string) *CheckClosingAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CheckClosingAccountResponseBody) SetPass(v bool) *CheckClosingAccountResponseBody {
	s.Pass = &v
	return s
}

type CheckClosingAccountResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckClosingAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckClosingAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountResponse) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountResponse) SetHeaders(v map[string]*string) *CheckClosingAccountResponse {
	s.Headers = v
	return s
}

func (s *CheckClosingAccountResponse) SetBody(v *CheckClosingAccountResponseBody) *CheckClosingAccountResponse {
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

type CheckWritePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckWritePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckWritePermissionHeaders) GoString() string {
	return s.String()
}

func (s *CheckWritePermissionHeaders) SetCommonHeaders(v map[string]*string) *CheckWritePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckWritePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *CheckWritePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckWritePermissionRequest struct {
	// corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// category
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// resourceKey
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// entityIds
	EntityIds []*int64 `json:"entityIds,omitempty" xml:"entityIds,omitempty" type:"Repeated"`
}

func (s CheckWritePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckWritePermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckWritePermissionRequest) SetCorpId(v string) *CheckWritePermissionRequest {
	s.CorpId = &v
	return s
}

func (s *CheckWritePermissionRequest) SetOpUserId(v string) *CheckWritePermissionRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckWritePermissionRequest) SetCategory(v string) *CheckWritePermissionRequest {
	s.Category = &v
	return s
}

func (s *CheckWritePermissionRequest) SetResourceKey(v string) *CheckWritePermissionRequest {
	s.ResourceKey = &v
	return s
}

func (s *CheckWritePermissionRequest) SetEntityIds(v []*int64) *CheckWritePermissionRequest {
	s.EntityIds = v
	return s
}

type CheckWritePermissionResponseBody struct {
	// entityPermissionMap
	EntityPermissionMap map[string]*bool `json:"entityPermissionMap,omitempty" xml:"entityPermissionMap,omitempty"`
}

func (s CheckWritePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckWritePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckWritePermissionResponseBody) SetEntityPermissionMap(v map[string]*bool) *CheckWritePermissionResponseBody {
	s.EntityPermissionMap = v
	return s
}

type CheckWritePermissionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckWritePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckWritePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckWritePermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckWritePermissionResponse) SetHeaders(v map[string]*string) *CheckWritePermissionResponse {
	s.Headers = v
	return s
}

func (s *CheckWritePermissionResponse) SetBody(v *CheckWritePermissionResponseBody) *CheckWritePermissionResponse {
	s.Body = v
	return s
}

type GetClosingAccountsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetClosingAccountsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsHeaders) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsHeaders) SetCommonHeaders(v map[string]*string) *GetClosingAccountsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetClosingAccountsHeaders) SetXAcsDingtalkAccessToken(v string) *GetClosingAccountsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetClosingAccountsRequest struct {
	// 人员列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetClosingAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsRequest) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsRequest) SetUserIds(v []*string) *GetClosingAccountsRequest {
	s.UserIds = v
	return s
}

type GetClosingAccountsResponseBody struct {
	// 规则列表
	Result []*GetClosingAccountsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetClosingAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponseBody) SetResult(v []*GetClosingAccountsResponseBodyResult) *GetClosingAccountsResponseBody {
	s.Result = v
	return s
}

type GetClosingAccountsResponseBodyResult struct {
	// 人员ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 开关
	SwitchOn *bool `json:"switchOn,omitempty" xml:"switchOn,omitempty"`
	// 封账规则
	ClosingAccountModel *GetClosingAccountsResponseBodyResultClosingAccountModel `json:"closingAccountModel,omitempty" xml:"closingAccountModel,omitempty" type:"Struct"`
	// 解封规则
	UnsealClosingAccountModel *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel `json:"unsealClosingAccountModel,omitempty" xml:"unsealClosingAccountModel,omitempty" type:"Struct"`
}

func (s GetClosingAccountsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponseBodyResult) SetUserId(v string) *GetClosingAccountsResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetSwitchOn(v bool) *GetClosingAccountsResponseBodyResult {
	s.SwitchOn = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetClosingAccountModel(v *GetClosingAccountsResponseBodyResultClosingAccountModel) *GetClosingAccountsResponseBodyResult {
	s.ClosingAccountModel = v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetUnsealClosingAccountModel(v *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel) *GetClosingAccountsResponseBodyResult {
	s.UnsealClosingAccountModel = v
	return s
}

type GetClosingAccountsResponseBodyResultClosingAccountModel struct {
	// 封账时间-日
	ClosingDay *int32 `json:"closingDay,omitempty" xml:"closingDay,omitempty"`
	// 封账时间-时分
	ClosingHourMinutes *int64 `json:"closingHourMinutes,omitempty" xml:"closingHourMinutes,omitempty"`
	// 封账范围-开始月
	StartMonth *int32 `json:"startMonth,omitempty" xml:"startMonth,omitempty"`
	// 封账范围-开始日
	StartDay *int32 `json:"startDay,omitempty" xml:"startDay,omitempty"`
	// 封账范围-结束月
	EndMonth *int32 `json:"endMonth,omitempty" xml:"endMonth,omitempty"`
	// 封账范围-结束日
	EndDay *int32 `json:"endDay,omitempty" xml:"endDay,omitempty"`
}

func (s GetClosingAccountsResponseBodyResultClosingAccountModel) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponseBodyResultClosingAccountModel) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetClosingDay(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.ClosingDay = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetClosingHourMinutes(v int64) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.ClosingHourMinutes = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetStartMonth(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.StartMonth = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetStartDay(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.StartDay = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetEndMonth(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.EndMonth = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetEndDay(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.EndDay = &v
	return s
}

type GetClosingAccountsResponseBodyResultUnsealClosingAccountModel struct {
	// 解封时间点
	InvalidTimeStamp *int64 `json:"invalidTimeStamp,omitempty" xml:"invalidTimeStamp,omitempty"`
}

func (s GetClosingAccountsResponseBodyResultUnsealClosingAccountModel) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponseBodyResultUnsealClosingAccountModel) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel) SetInvalidTimeStamp(v int64) *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel {
	s.InvalidTimeStamp = &v
	return s
}

type GetClosingAccountsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClosingAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClosingAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponse) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponse) SetHeaders(v map[string]*string) *GetClosingAccountsResponse {
	s.Headers = v
	return s
}

func (s *GetClosingAccountsResponse) SetBody(v *GetClosingAccountsResponseBody) *GetClosingAccountsResponse {
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

func (client *Client) CheckClosingAccount(request *CheckClosingAccountRequest) (_result *CheckClosingAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckClosingAccountHeaders{}
	_result = &CheckClosingAccountResponse{}
	_body, _err := client.CheckClosingAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckClosingAccountWithOptions(request *CheckClosingAccountRequest, headers *CheckClosingAccountHeaders, runtime *util.RuntimeOptions) (_result *CheckClosingAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserTimeRange)) {
		body["userTimeRange"] = request.UserTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
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
	_result = &CheckClosingAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckClosingAccount"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/closingAccounts/status/query"), tea.String("json"), req, runtime)
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

func (client *Client) CheckWritePermission(request *CheckWritePermissionRequest) (_result *CheckWritePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckWritePermissionHeaders{}
	_result = &CheckWritePermissionResponse{}
	_body, _err := client.CheckWritePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckWritePermissionWithOptions(request *CheckWritePermissionRequest, headers *CheckWritePermissionHeaders, runtime *util.RuntimeOptions) (_result *CheckWritePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceKey)) {
		body["resourceKey"] = request.ResourceKey
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIds)) {
		body["entityIds"] = request.EntityIds
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CheckWritePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckWritePermission"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/writePermissions/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClosingAccounts(request *GetClosingAccountsRequest) (_result *GetClosingAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetClosingAccountsHeaders{}
	_result = &GetClosingAccountsResponse{}
	_body, _err := client.GetClosingAccountsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClosingAccountsWithOptions(request *GetClosingAccountsRequest, headers *GetClosingAccountsHeaders, runtime *util.RuntimeOptions) (_result *GetClosingAccountsResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetClosingAccountsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetClosingAccounts"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/closingAccounts/rules/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
