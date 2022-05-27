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

type AddLeaveTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLeaveTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeHeaders) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeHeaders) SetCommonHeaders(v map[string]*string) *AddLeaveTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLeaveTypeHeaders) SetXAcsDingtalkAccessToken(v string) *AddLeaveTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLeaveTypeRequest struct {
	// 假期类型，普通假期或者加班转调休假期。(general_leave、lieu_leave其中一种)
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 调休假有效期规则(validity_type:有效类型 absolute_time(绝对时间)、relative_time(相对时间)其中一种 validity_value:延长日期(当validity_type为absolute_time该值该值不为空且满足yy-mm格式 validity_type为relative_time该值为大于1的整数))
	Extras *string `json:"extras,omitempty" xml:"extras,omitempty"`
	// 每天折算的工作时长(百分之一 例如1天=10小时=1000)
	HoursInPerDay *int64 `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	// 请假证明
	LeaveCertificate *AddLeaveTypeRequestLeaveCertificate `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	// 假期名称
	LeaveName *string `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	// 请假单位，可以按照天半天或者小时请假。(day、halfDay、hour其中一种)
	LeaveViewUnit *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	// 是否按照自然日统计请假时长，当为false的时候，用户发起请假时候会根据用户在请假时间段内的排班情况来计算请假时长。
	NaturalDayLeave *bool `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	// 限时提交规则
	SubmitTimeRule *AddLeaveTypeRequestSubmitTimeRule `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	// 操作者ID
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AddLeaveTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeRequest) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeRequest) SetBizType(v string) *AddLeaveTypeRequest {
	s.BizType = &v
	return s
}

func (s *AddLeaveTypeRequest) SetExtras(v string) *AddLeaveTypeRequest {
	s.Extras = &v
	return s
}

func (s *AddLeaveTypeRequest) SetHoursInPerDay(v int64) *AddLeaveTypeRequest {
	s.HoursInPerDay = &v
	return s
}

func (s *AddLeaveTypeRequest) SetLeaveCertificate(v *AddLeaveTypeRequestLeaveCertificate) *AddLeaveTypeRequest {
	s.LeaveCertificate = v
	return s
}

func (s *AddLeaveTypeRequest) SetLeaveName(v string) *AddLeaveTypeRequest {
	s.LeaveName = &v
	return s
}

func (s *AddLeaveTypeRequest) SetLeaveViewUnit(v string) *AddLeaveTypeRequest {
	s.LeaveViewUnit = &v
	return s
}

func (s *AddLeaveTypeRequest) SetNaturalDayLeave(v bool) *AddLeaveTypeRequest {
	s.NaturalDayLeave = &v
	return s
}

func (s *AddLeaveTypeRequest) SetSubmitTimeRule(v *AddLeaveTypeRequestSubmitTimeRule) *AddLeaveTypeRequest {
	s.SubmitTimeRule = v
	return s
}

func (s *AddLeaveTypeRequest) SetOpUserId(v string) *AddLeaveTypeRequest {
	s.OpUserId = &v
	return s
}

type AddLeaveTypeRequestLeaveCertificate struct {
	// 超过多长时间需提供请假证明
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否开启请假证明
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 请假提示文案
	PromptInformation *string `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	// 请假证明单位hour，day
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s AddLeaveTypeRequestLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeRequestLeaveCertificate) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeRequestLeaveCertificate) SetDuration(v float64) *AddLeaveTypeRequestLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *AddLeaveTypeRequestLeaveCertificate) SetEnable(v bool) *AddLeaveTypeRequestLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *AddLeaveTypeRequestLeaveCertificate) SetPromptInformation(v string) *AddLeaveTypeRequestLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *AddLeaveTypeRequestLeaveCertificate) SetUnit(v string) *AddLeaveTypeRequestLeaveCertificate {
	s.Unit = &v
	return s
}

type AddLeaveTypeRequestSubmitTimeRule struct {
	// 是否开启限时提交功能：仅且为true时开启
	EnableTimeLimit *bool `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	// 限制类型：before-提前；after-补交
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// 时间单位：day-天；hour-小时
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 限制值：timeUnit=day时，有效值范围[0~30] 天；timeUnit=hour时，有效值范围[0~24] 小时
	TimeValue *int64 `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s AddLeaveTypeRequestSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeRequestSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeRequestSubmitTimeRule) SetEnableTimeLimit(v bool) *AddLeaveTypeRequestSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *AddLeaveTypeRequestSubmitTimeRule) SetTimeType(v string) *AddLeaveTypeRequestSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *AddLeaveTypeRequestSubmitTimeRule) SetTimeUnit(v string) *AddLeaveTypeRequestSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *AddLeaveTypeRequestSubmitTimeRule) SetTimeValue(v int64) *AddLeaveTypeRequestSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type AddLeaveTypeResponseBody struct {
	// 返回参数
	Result *AddLeaveTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s AddLeaveTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponseBody) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponseBody) SetResult(v *AddLeaveTypeResponseBodyResult) *AddLeaveTypeResponseBody {
	s.Result = v
	return s
}

type AddLeaveTypeResponseBodyResult struct {
	// 假期类型，普通假期或者加班转调休假期。(general_leave、lieu_leave其中一种)
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 每天折算的工作时长(百分之一 例如1天=10小时=1000)
	HoursInPerDay *int64 `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	// 请假证明
	LeaveCertificate *AddLeaveTypeResponseBodyResultLeaveCertificate `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	// 假期类型唯一标识
	LeaveCode *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	// 假期名称
	LeaveName *string `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	// 请假单位，可以按照天半天或者小时请假。(day、halfDay、hour其中一种)
	LeaveViewUnit *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	// 是否按照自然日统计请假时长，当为false的时候，用户发起请假时候会根据用户在请假时间段内的排班情况来计算请假时长。
	NaturalDayLeave *bool `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	// 限时提交规则
	SubmitTimeRule *AddLeaveTypeResponseBodyResultSubmitTimeRule `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
}

func (s AddLeaveTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponseBodyResult) SetBizType(v string) *AddLeaveTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetHoursInPerDay(v int64) *AddLeaveTypeResponseBodyResult {
	s.HoursInPerDay = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetLeaveCertificate(v *AddLeaveTypeResponseBodyResultLeaveCertificate) *AddLeaveTypeResponseBodyResult {
	s.LeaveCertificate = v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetLeaveCode(v string) *AddLeaveTypeResponseBodyResult {
	s.LeaveCode = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetLeaveName(v string) *AddLeaveTypeResponseBodyResult {
	s.LeaveName = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetLeaveViewUnit(v string) *AddLeaveTypeResponseBodyResult {
	s.LeaveViewUnit = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetNaturalDayLeave(v bool) *AddLeaveTypeResponseBodyResult {
	s.NaturalDayLeave = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResult) SetSubmitTimeRule(v *AddLeaveTypeResponseBodyResultSubmitTimeRule) *AddLeaveTypeResponseBodyResult {
	s.SubmitTimeRule = v
	return s
}

type AddLeaveTypeResponseBodyResultLeaveCertificate struct {
	// 超过多长时间需提供请假证明
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否开启请假证明
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 请假提示文案
	PromptInformation *string `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	// 请假证明单位hour，day
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s AddLeaveTypeResponseBodyResultLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponseBodyResultLeaveCertificate) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponseBodyResultLeaveCertificate) SetDuration(v float64) *AddLeaveTypeResponseBodyResultLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultLeaveCertificate) SetEnable(v bool) *AddLeaveTypeResponseBodyResultLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultLeaveCertificate) SetPromptInformation(v string) *AddLeaveTypeResponseBodyResultLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultLeaveCertificate) SetUnit(v string) *AddLeaveTypeResponseBodyResultLeaveCertificate {
	s.Unit = &v
	return s
}

type AddLeaveTypeResponseBodyResultSubmitTimeRule struct {
	// 是否开启限时提交功能：仅且为true时开启
	EnableTimeLimit *bool `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	// 限制类型：before-提前；after-补交
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// 时间单位：day-天；hour-小时
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 限制值：timeUnit=day时，有效值范围[0~30] 天；timeUnit=hour时，有效值范围[0~24] 小时
	TimeValue *int64 `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s AddLeaveTypeResponseBodyResultSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponseBodyResultSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponseBodyResultSubmitTimeRule) SetEnableTimeLimit(v bool) *AddLeaveTypeResponseBodyResultSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeType(v string) *AddLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeUnit(v string) *AddLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeValue(v int64) *AddLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type AddLeaveTypeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLeaveTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponse) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponse) SetHeaders(v map[string]*string) *AddLeaveTypeResponse {
	s.Headers = v
	return s
}

func (s *AddLeaveTypeResponse) SetBody(v *AddLeaveTypeResponseBody) *AddLeaveTypeResponse {
	s.Body = v
	return s
}

type AttendanceBleDevicesAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AttendanceBleDevicesAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddHeaders) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddHeaders) SetCommonHeaders(v map[string]*string) *AttendanceBleDevicesAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AttendanceBleDevicesAddHeaders) SetXAcsDingtalkAccessToken(v string) *AttendanceBleDevicesAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AttendanceBleDevicesAddRequest struct {
	// 蓝牙设备Id列表
	DeviceIdList []*int64 `json:"deviceIdList,omitempty" xml:"deviceIdList,omitempty" type:"Repeated"`
	// 考勤组Id
	GroupKey *string `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
	// 操作人Id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AttendanceBleDevicesAddRequest) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddRequest) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddRequest) SetDeviceIdList(v []*int64) *AttendanceBleDevicesAddRequest {
	s.DeviceIdList = v
	return s
}

func (s *AttendanceBleDevicesAddRequest) SetGroupKey(v string) *AttendanceBleDevicesAddRequest {
	s.GroupKey = &v
	return s
}

func (s *AttendanceBleDevicesAddRequest) SetOpUserId(v string) *AttendanceBleDevicesAddRequest {
	s.OpUserId = &v
	return s
}

type AttendanceBleDevicesAddResponseBody struct {
	// 添加错误列表
	ErrorList []*AttendanceBleDevicesAddResponseBodyErrorList `json:"errorList,omitempty" xml:"errorList,omitempty" type:"Repeated"`
	// 添加成功蓝牙设备列表
	SuccessList []*AttendanceBleDevicesAddResponseBodySuccessList `json:"successList,omitempty" xml:"successList,omitempty" type:"Repeated"`
}

func (s AttendanceBleDevicesAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddResponseBody) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddResponseBody) SetErrorList(v []*AttendanceBleDevicesAddResponseBodyErrorList) *AttendanceBleDevicesAddResponseBody {
	s.ErrorList = v
	return s
}

func (s *AttendanceBleDevicesAddResponseBody) SetSuccessList(v []*AttendanceBleDevicesAddResponseBodySuccessList) *AttendanceBleDevicesAddResponseBody {
	s.SuccessList = v
	return s
}

type AttendanceBleDevicesAddResponseBodyErrorList struct {
	// 错误code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败蓝牙设备列表
	FailureList []*AttendanceBleDevicesAddResponseBodyErrorListFailureList `json:"failureList,omitempty" xml:"failureList,omitempty" type:"Repeated"`
	// errorMsg
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s AttendanceBleDevicesAddResponseBodyErrorList) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddResponseBodyErrorList) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddResponseBodyErrorList) SetCode(v string) *AttendanceBleDevicesAddResponseBodyErrorList {
	s.Code = &v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodyErrorList) SetFailureList(v []*AttendanceBleDevicesAddResponseBodyErrorListFailureList) *AttendanceBleDevicesAddResponseBodyErrorList {
	s.FailureList = v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodyErrorList) SetMsg(v string) *AttendanceBleDevicesAddResponseBodyErrorList {
	s.Msg = &v
	return s
}

type AttendanceBleDevicesAddResponseBodyErrorListFailureList struct {
	// 蓝牙设备Id
	DeviceId *int64 `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 蓝牙设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// sn
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s AttendanceBleDevicesAddResponseBodyErrorListFailureList) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddResponseBodyErrorListFailureList) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddResponseBodyErrorListFailureList) SetDeviceId(v int64) *AttendanceBleDevicesAddResponseBodyErrorListFailureList {
	s.DeviceId = &v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodyErrorListFailureList) SetDeviceName(v string) *AttendanceBleDevicesAddResponseBodyErrorListFailureList {
	s.DeviceName = &v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodyErrorListFailureList) SetSn(v string) *AttendanceBleDevicesAddResponseBodyErrorListFailureList {
	s.Sn = &v
	return s
}

type AttendanceBleDevicesAddResponseBodySuccessList struct {
	// 蓝牙设备Id
	DeviceId *int64 `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 蓝牙设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// sn
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s AttendanceBleDevicesAddResponseBodySuccessList) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddResponseBodySuccessList) SetDeviceId(v int64) *AttendanceBleDevicesAddResponseBodySuccessList {
	s.DeviceId = &v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodySuccessList) SetDeviceName(v string) *AttendanceBleDevicesAddResponseBodySuccessList {
	s.DeviceName = &v
	return s
}

func (s *AttendanceBleDevicesAddResponseBodySuccessList) SetSn(v string) *AttendanceBleDevicesAddResponseBodySuccessList {
	s.Sn = &v
	return s
}

type AttendanceBleDevicesAddResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttendanceBleDevicesAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttendanceBleDevicesAddResponse) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesAddResponse) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesAddResponse) SetHeaders(v map[string]*string) *AttendanceBleDevicesAddResponse {
	s.Headers = v
	return s
}

func (s *AttendanceBleDevicesAddResponse) SetBody(v *AttendanceBleDevicesAddResponseBody) *AttendanceBleDevicesAddResponse {
	s.Body = v
	return s
}

type AttendanceBleDevicesQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AttendanceBleDevicesQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesQueryHeaders) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesQueryHeaders) SetCommonHeaders(v map[string]*string) *AttendanceBleDevicesQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AttendanceBleDevicesQueryHeaders) SetXAcsDingtalkAccessToken(v string) *AttendanceBleDevicesQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AttendanceBleDevicesQueryRequest struct {
	// 考勤组Id
	GroupKey *string `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
	// 操作人Id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AttendanceBleDevicesQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesQueryRequest) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesQueryRequest) SetGroupKey(v string) *AttendanceBleDevicesQueryRequest {
	s.GroupKey = &v
	return s
}

func (s *AttendanceBleDevicesQueryRequest) SetOpUserId(v string) *AttendanceBleDevicesQueryRequest {
	s.OpUserId = &v
	return s
}

type AttendanceBleDevicesQueryResponseBody struct {
	// 蓝牙列表
	Result []*AttendanceBleDevicesQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s AttendanceBleDevicesQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesQueryResponseBody) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesQueryResponseBody) SetResult(v []*AttendanceBleDevicesQueryResponseBodyResult) *AttendanceBleDevicesQueryResponseBody {
	s.Result = v
	return s
}

type AttendanceBleDevicesQueryResponseBodyResult struct {
	// 蓝牙设备Id
	DeviceId *int64 `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 蓝牙设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// sn
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s AttendanceBleDevicesQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesQueryResponseBodyResult) SetDeviceId(v int64) *AttendanceBleDevicesQueryResponseBodyResult {
	s.DeviceId = &v
	return s
}

func (s *AttendanceBleDevicesQueryResponseBodyResult) SetDeviceName(v string) *AttendanceBleDevicesQueryResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *AttendanceBleDevicesQueryResponseBodyResult) SetSn(v string) *AttendanceBleDevicesQueryResponseBodyResult {
	s.Sn = &v
	return s
}

type AttendanceBleDevicesQueryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttendanceBleDevicesQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttendanceBleDevicesQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesQueryResponse) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesQueryResponse) SetHeaders(v map[string]*string) *AttendanceBleDevicesQueryResponse {
	s.Headers = v
	return s
}

func (s *AttendanceBleDevicesQueryResponse) SetBody(v *AttendanceBleDevicesQueryResponseBody) *AttendanceBleDevicesQueryResponse {
	s.Body = v
	return s
}

type AttendanceBleDevicesRemoveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AttendanceBleDevicesRemoveHeaders) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesRemoveHeaders) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesRemoveHeaders) SetCommonHeaders(v map[string]*string) *AttendanceBleDevicesRemoveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AttendanceBleDevicesRemoveHeaders) SetXAcsDingtalkAccessToken(v string) *AttendanceBleDevicesRemoveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AttendanceBleDevicesRemoveRequest struct {
	// 蓝牙设备Id列表
	DeviceIdList []*int64 `json:"deviceIdList,omitempty" xml:"deviceIdList,omitempty" type:"Repeated"`
	// 考勤组Id
	GroupKey *string `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
	// 操作人id
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s AttendanceBleDevicesRemoveRequest) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesRemoveRequest) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesRemoveRequest) SetDeviceIdList(v []*int64) *AttendanceBleDevicesRemoveRequest {
	s.DeviceIdList = v
	return s
}

func (s *AttendanceBleDevicesRemoveRequest) SetGroupKey(v string) *AttendanceBleDevicesRemoveRequest {
	s.GroupKey = &v
	return s
}

func (s *AttendanceBleDevicesRemoveRequest) SetOpUserId(v string) *AttendanceBleDevicesRemoveRequest {
	s.OpUserId = &v
	return s
}

type AttendanceBleDevicesRemoveResponseBody struct {
	// 移出错误列表
	ErrorList []*AttendanceBleDevicesRemoveResponseBodyErrorList `json:"errorList,omitempty" xml:"errorList,omitempty" type:"Repeated"`
	// 移除成功蓝牙设备Id列表
	SuccessList []*int64 `json:"successList,omitempty" xml:"successList,omitempty" type:"Repeated"`
}

func (s AttendanceBleDevicesRemoveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesRemoveResponseBody) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesRemoveResponseBody) SetErrorList(v []*AttendanceBleDevicesRemoveResponseBodyErrorList) *AttendanceBleDevicesRemoveResponseBody {
	s.ErrorList = v
	return s
}

func (s *AttendanceBleDevicesRemoveResponseBody) SetSuccessList(v []*int64) *AttendanceBleDevicesRemoveResponseBody {
	s.SuccessList = v
	return s
}

type AttendanceBleDevicesRemoveResponseBodyErrorList struct {
	// 错误code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 移除失败蓝牙设备Id列表
	FailureList []*int64 `json:"failureList,omitempty" xml:"failureList,omitempty" type:"Repeated"`
	// 错误信息
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s AttendanceBleDevicesRemoveResponseBodyErrorList) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesRemoveResponseBodyErrorList) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesRemoveResponseBodyErrorList) SetCode(v string) *AttendanceBleDevicesRemoveResponseBodyErrorList {
	s.Code = &v
	return s
}

func (s *AttendanceBleDevicesRemoveResponseBodyErrorList) SetFailureList(v []*int64) *AttendanceBleDevicesRemoveResponseBodyErrorList {
	s.FailureList = v
	return s
}

func (s *AttendanceBleDevicesRemoveResponseBodyErrorList) SetMsg(v string) *AttendanceBleDevicesRemoveResponseBodyErrorList {
	s.Msg = &v
	return s
}

type AttendanceBleDevicesRemoveResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttendanceBleDevicesRemoveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttendanceBleDevicesRemoveResponse) String() string {
	return tea.Prettify(s)
}

func (s AttendanceBleDevicesRemoveResponse) GoString() string {
	return s.String()
}

func (s *AttendanceBleDevicesRemoveResponse) SetHeaders(v map[string]*string) *AttendanceBleDevicesRemoveResponse {
	s.Headers = v
	return s
}

func (s *AttendanceBleDevicesRemoveResponse) SetBody(v *AttendanceBleDevicesRemoveResponseBody) *AttendanceBleDevicesRemoveResponse {
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
	// 情景
	BizCode *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	// 员工列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// 时间段
	UserTimeRange []*CheckClosingAccountRequestUserTimeRange `json:"userTimeRange,omitempty" xml:"userTimeRange,omitempty" type:"Repeated"`
}

func (s CheckClosingAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountRequest) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountRequest) SetBizCode(v string) *CheckClosingAccountRequest {
	s.BizCode = &v
	return s
}

func (s *CheckClosingAccountRequest) SetUserIds(v []*string) *CheckClosingAccountRequest {
	s.UserIds = v
	return s
}

func (s *CheckClosingAccountRequest) SetUserTimeRange(v []*CheckClosingAccountRequestUserTimeRange) *CheckClosingAccountRequest {
	s.UserTimeRange = v
	return s
}

type CheckClosingAccountRequestUserTimeRange struct {
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CheckClosingAccountRequestUserTimeRange) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountRequestUserTimeRange) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountRequestUserTimeRange) SetEndTime(v int64) *CheckClosingAccountRequestUserTimeRange {
	s.EndTime = &v
	return s
}

func (s *CheckClosingAccountRequestUserTimeRange) SetStartTime(v int64) *CheckClosingAccountRequestUserTimeRange {
	s.StartTime = &v
	return s
}

type CheckClosingAccountResponseBody struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Mesage *string `json:"mesage,omitempty" xml:"mesage,omitempty"`
	Pass   *bool   `json:"pass,omitempty" xml:"pass,omitempty"`
}

func (s CheckClosingAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckClosingAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CheckClosingAccountResponseBody) SetCode(v string) *CheckClosingAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CheckClosingAccountResponseBody) SetMesage(v string) *CheckClosingAccountResponseBody {
	s.Mesage = &v
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
	// category
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// entityIds
	EntityIds []*int64 `json:"entityIds,omitempty" xml:"entityIds,omitempty" type:"Repeated"`
	// opUserId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// resourceKey
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
}

func (s CheckWritePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckWritePermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckWritePermissionRequest) SetCategory(v string) *CheckWritePermissionRequest {
	s.Category = &v
	return s
}

func (s *CheckWritePermissionRequest) SetEntityIds(v []*int64) *CheckWritePermissionRequest {
	s.EntityIds = v
	return s
}

func (s *CheckWritePermissionRequest) SetOpUserId(v string) *CheckWritePermissionRequest {
	s.OpUserId = &v
	return s
}

func (s *CheckWritePermissionRequest) SetResourceKey(v string) *CheckWritePermissionRequest {
	s.ResourceKey = &v
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
	// 三方审批单id，全局唯一
	ApproveId *string `json:"approveId,omitempty" xml:"approveId,omitempty"`
	// 审批人员工id
	OpUserid *string `json:"opUserid,omitempty" xml:"opUserid,omitempty"`
	// 审批单关联的打卡信息
	PunchParam *CreateApproveRequestPunchParam `json:"punchParam,omitempty" xml:"punchParam,omitempty" type:"Struct"`
	// 子类型名称，最大长度20个字符
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// 审批单类型名称，最大长度20个字符
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// 员工id
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s CreateApproveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveRequest) GoString() string {
	return s.String()
}

func (s *CreateApproveRequest) SetApproveId(v string) *CreateApproveRequest {
	s.ApproveId = &v
	return s
}

func (s *CreateApproveRequest) SetOpUserid(v string) *CreateApproveRequest {
	s.OpUserid = &v
	return s
}

func (s *CreateApproveRequest) SetPunchParam(v *CreateApproveRequestPunchParam) *CreateApproveRequest {
	s.PunchParam = v
	return s
}

func (s *CreateApproveRequest) SetSubType(v string) *CreateApproveRequest {
	s.SubType = &v
	return s
}

func (s *CreateApproveRequest) SetTagName(v string) *CreateApproveRequest {
	s.TagName = &v
	return s
}

func (s *CreateApproveRequest) SetUserid(v string) *CreateApproveRequest {
	s.Userid = &v
	return s
}

type CreateApproveRequestPunchParam struct {
	// 地理位置标识：wifi:ssid_macAddress ble: deviceId gps:longitude_latitude
	PositionId *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	// 地理位置名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 地理位置类型：wifi/ble/gps
	PositionType *string `json:"positionType,omitempty" xml:"positionType,omitempty"`
	// 打卡时间，单位毫秒
	PunchTime *int64 `json:"punchTime,omitempty" xml:"punchTime,omitempty"`
}

func (s CreateApproveRequestPunchParam) String() string {
	return tea.Prettify(s)
}

func (s CreateApproveRequestPunchParam) GoString() string {
	return s.String()
}

func (s *CreateApproveRequestPunchParam) SetPositionId(v string) *CreateApproveRequestPunchParam {
	s.PositionId = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPositionName(v string) *CreateApproveRequestPunchParam {
	s.PositionName = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPositionType(v string) *CreateApproveRequestPunchParam {
	s.PositionType = &v
	return s
}

func (s *CreateApproveRequestPunchParam) SetPunchTime(v int64) *CreateApproveRequestPunchParam {
	s.PunchTime = &v
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

type DingTalkSecurityCheckHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DingTalkSecurityCheckHeaders) String() string {
	return tea.Prettify(s)
}

func (s DingTalkSecurityCheckHeaders) GoString() string {
	return s.String()
}

func (s *DingTalkSecurityCheckHeaders) SetCommonHeaders(v map[string]*string) *DingTalkSecurityCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DingTalkSecurityCheckHeaders) SetXAcsDingtalkAccessToken(v string) *DingTalkSecurityCheckHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DingTalkSecurityCheckRequest struct {
	// 客户端版本号
	ClientVer *string `json:"clientVer,omitempty" xml:"clientVer,omitempty"`
	// 客户端平台类型(iOS,Android)
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 客户端平台平台版本
	PlatformVer *string `json:"platformVer,omitempty" xml:"platformVer,omitempty"`
	// 加密字符
	Sec *string `json:"sec,omitempty" xml:"sec,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DingTalkSecurityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s DingTalkSecurityCheckRequest) GoString() string {
	return s.String()
}

func (s *DingTalkSecurityCheckRequest) SetClientVer(v string) *DingTalkSecurityCheckRequest {
	s.ClientVer = &v
	return s
}

func (s *DingTalkSecurityCheckRequest) SetPlatform(v string) *DingTalkSecurityCheckRequest {
	s.Platform = &v
	return s
}

func (s *DingTalkSecurityCheckRequest) SetPlatformVer(v string) *DingTalkSecurityCheckRequest {
	s.PlatformVer = &v
	return s
}

func (s *DingTalkSecurityCheckRequest) SetSec(v string) *DingTalkSecurityCheckRequest {
	s.Sec = &v
	return s
}

func (s *DingTalkSecurityCheckRequest) SetUserId(v string) *DingTalkSecurityCheckRequest {
	s.UserId = &v
	return s
}

type DingTalkSecurityCheckResponseBody struct {
	// 返回参数
	Result *DingTalkSecurityCheckResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DingTalkSecurityCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DingTalkSecurityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DingTalkSecurityCheckResponseBody) SetResult(v *DingTalkSecurityCheckResponseBodyResult) *DingTalkSecurityCheckResponseBody {
	s.Result = v
	return s
}

type DingTalkSecurityCheckResponseBodyResult struct {
	// 是否有风险
	HasRisk *bool `json:"hasRisk,omitempty" xml:"hasRisk,omitempty"`
	// 风险信息
	RiskInfo map[string]*string `json:"riskInfo,omitempty" xml:"riskInfo,omitempty"`
}

func (s DingTalkSecurityCheckResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DingTalkSecurityCheckResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DingTalkSecurityCheckResponseBodyResult) SetHasRisk(v bool) *DingTalkSecurityCheckResponseBodyResult {
	s.HasRisk = &v
	return s
}

func (s *DingTalkSecurityCheckResponseBodyResult) SetRiskInfo(v map[string]*string) *DingTalkSecurityCheckResponseBodyResult {
	s.RiskInfo = v
	return s
}

type DingTalkSecurityCheckResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DingTalkSecurityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DingTalkSecurityCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s DingTalkSecurityCheckResponse) GoString() string {
	return s.String()
}

func (s *DingTalkSecurityCheckResponse) SetHeaders(v map[string]*string) *DingTalkSecurityCheckResponse {
	s.Headers = v
	return s
}

func (s *DingTalkSecurityCheckResponse) SetBody(v *DingTalkSecurityCheckResponseBody) *DingTalkSecurityCheckResponse {
	s.Body = v
	return s
}

type GetAdjustmentsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAdjustmentsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsHeaders) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsHeaders) SetCommonHeaders(v map[string]*string) *GetAdjustmentsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAdjustmentsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAdjustmentsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAdjustmentsRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetAdjustmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsRequest) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsRequest) SetPageNumber(v int64) *GetAdjustmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAdjustmentsRequest) SetPageSize(v int64) *GetAdjustmentsRequest {
	s.PageSize = &v
	return s
}

type GetAdjustmentsResponseBody struct {
	// Id of the request
	Result []*GetAdjustmentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetAdjustmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsResponseBody) SetResult(v []*GetAdjustmentsResponseBodyResult) *GetAdjustmentsResponseBody {
	s.Result = v
	return s
}

type GetAdjustmentsResponseBodyResult struct {
	// 补卡规则集合
	Items []*GetAdjustmentsResponseBodyResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 当前页码
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总页数
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetAdjustmentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsResponseBodyResult) SetItems(v []*GetAdjustmentsResponseBodyResultItems) *GetAdjustmentsResponseBodyResult {
	s.Items = v
	return s
}

func (s *GetAdjustmentsResponseBodyResult) SetPageNumber(v int64) *GetAdjustmentsResponseBodyResult {
	s.PageNumber = &v
	return s
}

func (s *GetAdjustmentsResponseBodyResult) SetTotalPage(v int64) *GetAdjustmentsResponseBodyResult {
	s.TotalPage = &v
	return s
}

type GetAdjustmentsResponseBodyResultItems struct {
	// 补卡规则id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 补卡规则名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAdjustmentsResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsResponseBodyResultItems) SetId(v int64) *GetAdjustmentsResponseBodyResultItems {
	s.Id = &v
	return s
}

func (s *GetAdjustmentsResponseBodyResultItems) SetName(v string) *GetAdjustmentsResponseBodyResultItems {
	s.Name = &v
	return s
}

type GetAdjustmentsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAdjustmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAdjustmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsResponse) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsResponse) SetHeaders(v map[string]*string) *GetAdjustmentsResponse {
	s.Headers = v
	return s
}

func (s *GetAdjustmentsResponse) SetBody(v *GetAdjustmentsResponseBody) *GetAdjustmentsResponse {
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
	// 封账规则
	ClosingAccountModel *GetClosingAccountsResponseBodyResultClosingAccountModel `json:"closingAccountModel,omitempty" xml:"closingAccountModel,omitempty" type:"Struct"`
	// 开关
	SwitchOn *bool `json:"switchOn,omitempty" xml:"switchOn,omitempty"`
	// 解封规则
	UnsealClosingAccountModel *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel `json:"unsealClosingAccountModel,omitempty" xml:"unsealClosingAccountModel,omitempty" type:"Struct"`
	// 人员ID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetClosingAccountsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClosingAccountsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClosingAccountsResponseBodyResult) SetClosingAccountModel(v *GetClosingAccountsResponseBodyResultClosingAccountModel) *GetClosingAccountsResponseBodyResult {
	s.ClosingAccountModel = v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetSwitchOn(v bool) *GetClosingAccountsResponseBodyResult {
	s.SwitchOn = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetUnsealClosingAccountModel(v *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel) *GetClosingAccountsResponseBodyResult {
	s.UnsealClosingAccountModel = v
	return s
}

func (s *GetClosingAccountsResponseBodyResult) SetUserId(v string) *GetClosingAccountsResponseBodyResult {
	s.UserId = &v
	return s
}

type GetClosingAccountsResponseBodyResultClosingAccountModel struct {
	// 封账时间-日
	ClosingDay *int32 `json:"closingDay,omitempty" xml:"closingDay,omitempty"`
	// 封账时间-时分
	ClosingHourMinutes *int64 `json:"closingHourMinutes,omitempty" xml:"closingHourMinutes,omitempty"`
	// 封账范围-结束日
	EndDay *int32 `json:"endDay,omitempty" xml:"endDay,omitempty"`
	// 封账范围-结束月
	EndMonth *int32 `json:"endMonth,omitempty" xml:"endMonth,omitempty"`
	// 封账范围-开始日
	StartDay *int32 `json:"startDay,omitempty" xml:"startDay,omitempty"`
	// 封账范围-开始月
	StartMonth *int32 `json:"startMonth,omitempty" xml:"startMonth,omitempty"`
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

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetEndDay(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.EndDay = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetEndMonth(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.EndMonth = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetStartDay(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.StartDay = &v
	return s
}

func (s *GetClosingAccountsResponseBodyResultClosingAccountModel) SetStartMonth(v int32) *GetClosingAccountsResponseBodyResultClosingAccountModel {
	s.StartMonth = &v
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

type GetLeaveTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLeaveTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeHeaders) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeHeaders) SetCommonHeaders(v map[string]*string) *GetLeaveTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLeaveTypeHeaders) SetXAcsDingtalkAccessToken(v string) *GetLeaveTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLeaveTypeRequest struct {
	// 操作者ID
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 空:开放接口定义假期类型;all:所有假期类型
	VacationSource *string `json:"vacationSource,omitempty" xml:"vacationSource,omitempty"`
}

func (s GetLeaveTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeRequest) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeRequest) SetOpUserId(v string) *GetLeaveTypeRequest {
	s.OpUserId = &v
	return s
}

func (s *GetLeaveTypeRequest) SetVacationSource(v string) *GetLeaveTypeRequest {
	s.VacationSource = &v
	return s
}

type GetLeaveTypeResponseBody struct {
	// 返回参数
	Result []*GetLeaveTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetLeaveTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponseBody) SetResult(v []*GetLeaveTypeResponseBodyResult) *GetLeaveTypeResponseBody {
	s.Result = v
	return s
}

type GetLeaveTypeResponseBodyResult struct {
	// 假期类型，普通假期或者加班转调休假期。(general_leave、lieu_leave其中一种)
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 每天折算的工作时长(百分之一 例如1天=10小时=1000)
	HoursInPerDay *int64 `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	// 请假证明
	LeaveCertificate *GetLeaveTypeResponseBodyResultLeaveCertificate `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	// 假期类型唯一标识
	LeaveCode *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	// 假期名称
	LeaveName *string `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	// 请假单位，可以按照天半天或者小时请假。(day、halfDay、hour其中一种)
	LeaveViewUnit *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	// 是否按照自然日统计请假时长，当为false的时候，用户发起请假时候会根据用户在请假时间段内的排班情况来计算请假时长。
	NaturalDayLeave *bool `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	// 开放接口自定义的:external oa后台新建的：inner
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 限时提交规则
	SubmitTimeRule *GetLeaveTypeResponseBodyResultSubmitTimeRule `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	// 有效类型 absolute_time(绝对时间)、relative_time(相对时间)其中一种
	ValidityType *string `json:"validityType,omitempty" xml:"validityType,omitempty"`
	// 延长日期(当validity_type为absolute_time该值该值不为空且满足yy-mm格式 validity_type为relative_time该值为大于1的整数)
	ValidityValue *string `json:"validityValue,omitempty" xml:"validityValue,omitempty"`
}

func (s GetLeaveTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponseBodyResult) SetBizType(v string) *GetLeaveTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetHoursInPerDay(v int64) *GetLeaveTypeResponseBodyResult {
	s.HoursInPerDay = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetLeaveCertificate(v *GetLeaveTypeResponseBodyResultLeaveCertificate) *GetLeaveTypeResponseBodyResult {
	s.LeaveCertificate = v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetLeaveCode(v string) *GetLeaveTypeResponseBodyResult {
	s.LeaveCode = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetLeaveName(v string) *GetLeaveTypeResponseBodyResult {
	s.LeaveName = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetLeaveViewUnit(v string) *GetLeaveTypeResponseBodyResult {
	s.LeaveViewUnit = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetNaturalDayLeave(v bool) *GetLeaveTypeResponseBodyResult {
	s.NaturalDayLeave = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetSource(v string) *GetLeaveTypeResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetSubmitTimeRule(v *GetLeaveTypeResponseBodyResultSubmitTimeRule) *GetLeaveTypeResponseBodyResult {
	s.SubmitTimeRule = v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetValidityType(v string) *GetLeaveTypeResponseBodyResult {
	s.ValidityType = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResult) SetValidityValue(v string) *GetLeaveTypeResponseBodyResult {
	s.ValidityValue = &v
	return s
}

type GetLeaveTypeResponseBodyResultLeaveCertificate struct {
	// 超过多长时间需提供请假证明
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否开启请假证明
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 请假提示文案
	PromptInformation *string `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	// 请假证明单位hour，day
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetLeaveTypeResponseBodyResultLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponseBodyResultLeaveCertificate) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponseBodyResultLeaveCertificate) SetDuration(v float64) *GetLeaveTypeResponseBodyResultLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultLeaveCertificate) SetEnable(v bool) *GetLeaveTypeResponseBodyResultLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultLeaveCertificate) SetPromptInformation(v string) *GetLeaveTypeResponseBodyResultLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultLeaveCertificate) SetUnit(v string) *GetLeaveTypeResponseBodyResultLeaveCertificate {
	s.Unit = &v
	return s
}

type GetLeaveTypeResponseBodyResultSubmitTimeRule struct {
	// 是否开启限时提交功能：仅且为true时开启
	EnableTimeLimit *bool `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	// 限制类型：before-提前；after-补交
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// 时间单位：day-天；hour-小时
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 限制值：timeUnit=day时，有效值范围[0~30] 天；timeUnit=hour时，有效值范围[0~24] 小时
	TimeValue *int64 `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s GetLeaveTypeResponseBodyResultSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponseBodyResultSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponseBodyResultSubmitTimeRule) SetEnableTimeLimit(v bool) *GetLeaveTypeResponseBodyResultSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeType(v string) *GetLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeUnit(v string) *GetLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeValue(v int64) *GetLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type GetLeaveTypeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLeaveTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponse) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponse) SetHeaders(v map[string]*string) *GetLeaveTypeResponse {
	s.Headers = v
	return s
}

func (s *GetLeaveTypeResponse) SetBody(v *GetLeaveTypeResponseBody) *GetLeaveTypeResponse {
	s.Body = v
	return s
}

type GetMachineHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMachineHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMachineHeaders) GoString() string {
	return s.String()
}

func (s *GetMachineHeaders) SetCommonHeaders(v map[string]*string) *GetMachineHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMachineHeaders) SetXAcsDingtalkAccessToken(v string) *GetMachineHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMachineResponseBody struct {
	// 查询结果
	Result *GetMachineResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetMachineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMachineResponseBody) GoString() string {
	return s.String()
}

func (s *GetMachineResponseBody) SetResult(v *GetMachineResponseBodyResult) *GetMachineResponseBody {
	s.Result = v
	return s
}

type GetMachineResponseBodyResult struct {
	// 设备管理员列表
	AtmManagerList []*string `json:"atmManagerList,omitempty" xml:"atmManagerList,omitempty" type:"Repeated"`
	// 设备id (deviceId)
	DevId *int64 `json:"devId,omitempty" xml:"devId,omitempty"`
	// 设备id (deviceUid加密之后)
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 设备sn号
	DeviceSn *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	// 考勤机蓝牙相关设置信息
	MachineBluetoothVO *GetMachineResponseBodyResultMachineBluetoothVO `json:"machineBluetoothVO,omitempty" xml:"machineBluetoothVO,omitempty" type:"Struct"`
	// 人脸容量
	MaxFace *int32 `json:"maxFace,omitempty" xml:"maxFace,omitempty"`
	// 网络状态
	NetStatus *string `json:"netStatus,omitempty" xml:"netStatus,omitempty"`
	// 设备类型名称
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 固件版本
	ProductVersion *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	// 音量模式
	VoiceMode *int32 `json:"voiceMode,omitempty" xml:"voiceMode,omitempty"`
}

func (s GetMachineResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMachineResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMachineResponseBodyResult) SetAtmManagerList(v []*string) *GetMachineResponseBodyResult {
	s.AtmManagerList = v
	return s
}

func (s *GetMachineResponseBodyResult) SetDevId(v int64) *GetMachineResponseBodyResult {
	s.DevId = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetDeviceId(v string) *GetMachineResponseBodyResult {
	s.DeviceId = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetDeviceName(v string) *GetMachineResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetDeviceSn(v string) *GetMachineResponseBodyResult {
	s.DeviceSn = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetMachineBluetoothVO(v *GetMachineResponseBodyResultMachineBluetoothVO) *GetMachineResponseBodyResult {
	s.MachineBluetoothVO = v
	return s
}

func (s *GetMachineResponseBodyResult) SetMaxFace(v int32) *GetMachineResponseBodyResult {
	s.MaxFace = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetNetStatus(v string) *GetMachineResponseBodyResult {
	s.NetStatus = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetProductName(v string) *GetMachineResponseBodyResult {
	s.ProductName = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetProductVersion(v string) *GetMachineResponseBodyResult {
	s.ProductVersion = &v
	return s
}

func (s *GetMachineResponseBodyResult) SetVoiceMode(v int32) *GetMachineResponseBodyResult {
	s.VoiceMode = &v
	return s
}

type GetMachineResponseBodyResultMachineBluetoothVO struct {
	// 地址位置描述
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 蓝牙打卡人脸识别开关值
	BluetoothCheckWithFace *bool `json:"bluetoothCheckWithFace,omitempty" xml:"bluetoothCheckWithFace,omitempty"`
	// 蓝牙打卡范围
	BluetoothDistanceMode *string `json:"bluetoothDistanceMode,omitempty" xml:"bluetoothDistanceMode,omitempty"`
	// 蓝牙打卡范围描述
	BluetoothDistanceModeDesc *string `json:"bluetoothDistanceModeDesc,omitempty" xml:"bluetoothDistanceModeDesc,omitempty"`
	// 蓝牙打卡开关
	BluetoothValue *bool `json:"bluetoothValue,omitempty" xml:"bluetoothValue,omitempty"`
	// 纬度
	Latitude *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 是否限制员工常用手机
	LimitUserDeviceCount *bool `json:"limitUserDeviceCount,omitempty" xml:"limitUserDeviceCount,omitempty"`
	// 经度
	Longitude *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 是否打开位置异常监控
	MonitorLocationAbnormal *bool `json:"monitorLocationAbnormal,omitempty" xml:"monitorLocationAbnormal,omitempty"`
	// 员工常用手机数量
	UserDeviceCount *int32 `json:"userDeviceCount,omitempty" xml:"userDeviceCount,omitempty"`
}

func (s GetMachineResponseBodyResultMachineBluetoothVO) String() string {
	return tea.Prettify(s)
}

func (s GetMachineResponseBodyResultMachineBluetoothVO) GoString() string {
	return s.String()
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetAddress(v string) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.Address = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetBluetoothCheckWithFace(v bool) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.BluetoothCheckWithFace = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetBluetoothDistanceMode(v string) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.BluetoothDistanceMode = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetBluetoothDistanceModeDesc(v string) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.BluetoothDistanceModeDesc = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetBluetoothValue(v bool) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.BluetoothValue = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetLatitude(v float64) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.Latitude = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetLimitUserDeviceCount(v bool) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.LimitUserDeviceCount = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetLongitude(v float64) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.Longitude = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetMonitorLocationAbnormal(v bool) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.MonitorLocationAbnormal = &v
	return s
}

func (s *GetMachineResponseBodyResultMachineBluetoothVO) SetUserDeviceCount(v int32) *GetMachineResponseBodyResultMachineBluetoothVO {
	s.UserDeviceCount = &v
	return s
}

type GetMachineResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMachineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMachineResponse) GoString() string {
	return s.String()
}

func (s *GetMachineResponse) SetHeaders(v map[string]*string) *GetMachineResponse {
	s.Headers = v
	return s
}

func (s *GetMachineResponse) SetBody(v *GetMachineResponseBody) *GetMachineResponse {
	s.Body = v
	return s
}

type GetMachineUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMachineUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserHeaders) GoString() string {
	return s.String()
}

func (s *GetMachineUserHeaders) SetCommonHeaders(v map[string]*string) *GetMachineUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMachineUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetMachineUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMachineUserRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetMachineUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserRequest) GoString() string {
	return s.String()
}

func (s *GetMachineUserRequest) SetMaxResults(v int32) *GetMachineUserRequest {
	s.MaxResults = &v
	return s
}

func (s *GetMachineUserRequest) SetNextToken(v string) *GetMachineUserRequest {
	s.NextToken = &v
	return s
}

type GetMachineUserResponseBody struct {
	Result *GetMachineUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetMachineUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetMachineUserResponseBody) SetResult(v *GetMachineUserResponseBodyResult) *GetMachineUserResponseBody {
	s.Result = v
	return s
}

type GetMachineUserResponseBodyResult struct {
	HasMore   *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserList  []*GetMachineUserResponseBodyResultUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s GetMachineUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMachineUserResponseBodyResult) SetHasMore(v bool) *GetMachineUserResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetMachineUserResponseBodyResult) SetNextToken(v string) *GetMachineUserResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetMachineUserResponseBodyResult) SetUserList(v []*GetMachineUserResponseBodyResultUserList) *GetMachineUserResponseBodyResult {
	s.UserList = v
	return s
}

type GetMachineUserResponseBodyResultUserList struct {
	HasFace *bool   `json:"hasFace,omitempty" xml:"hasFace,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMachineUserResponseBodyResultUserList) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserResponseBodyResultUserList) GoString() string {
	return s.String()
}

func (s *GetMachineUserResponseBodyResultUserList) SetHasFace(v bool) *GetMachineUserResponseBodyResultUserList {
	s.HasFace = &v
	return s
}

func (s *GetMachineUserResponseBodyResultUserList) SetName(v string) *GetMachineUserResponseBodyResultUserList {
	s.Name = &v
	return s
}

func (s *GetMachineUserResponseBodyResultUserList) SetUserId(v string) *GetMachineUserResponseBodyResultUserList {
	s.UserId = &v
	return s
}

type GetMachineUserResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMachineUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMachineUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMachineUserResponse) GoString() string {
	return s.String()
}

func (s *GetMachineUserResponse) SetHeaders(v map[string]*string) *GetMachineUserResponse {
	s.Headers = v
	return s
}

func (s *GetMachineUserResponse) SetBody(v *GetMachineUserResponseBody) *GetMachineUserResponse {
	s.Body = v
	return s
}

type GetOvertimeSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOvertimeSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingHeaders) SetCommonHeaders(v map[string]*string) *GetOvertimeSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOvertimeSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetOvertimeSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOvertimeSettingRequest struct {
	OvertimeSettingIds []*int64 `json:"overtimeSettingIds,omitempty" xml:"overtimeSettingIds,omitempty" type:"Repeated"`
}

func (s GetOvertimeSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingRequest) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingRequest) SetOvertimeSettingIds(v []*int64) *GetOvertimeSettingRequest {
	s.OvertimeSettingIds = v
	return s
}

type GetOvertimeSettingResponseBody struct {
	Result []*GetOvertimeSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetOvertimeSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingResponseBody) SetResult(v []*GetOvertimeSettingResponseBodyResult) *GetOvertimeSettingResponseBody {
	s.Result = v
	return s
}

type GetOvertimeSettingResponseBodyResult struct {
	// 是否默认
	Default          *bool                                   `json:"default,omitempty" xml:"default,omitempty"`
	DurationSettings map[string]*ResultDurationSettingsValue `json:"durationSettings,omitempty" xml:"durationSettings,omitempty"`
	// 历史加班规则设置id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 规则名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 时间分割规则
	OvertimeDivisions []*GetOvertimeSettingResponseBodyResultOvertimeDivisions `json:"overtimeDivisions,omitempty" xml:"overtimeDivisions,omitempty" type:"Repeated"`
	// 设置id
	SettingId *int64 `json:"settingId,omitempty" xml:"settingId,omitempty"`
	// 加班时长单位
	StepType *int32 `json:"stepType,omitempty" xml:"stepType,omitempty"`
	// 加班时长是否取整 单位 小时
	StepValue       *float32                                               `json:"stepValue,omitempty" xml:"stepValue,omitempty"`
	WarningSettings []*GetOvertimeSettingResponseBodyResultWarningSettings `json:"warningSettings,omitempty" xml:"warningSettings,omitempty" type:"Repeated"`
	// 日折算时长 单位：分钟
	WorkMinutesPerDay *int32 `json:"workMinutesPerDay,omitempty" xml:"workMinutesPerDay,omitempty"`
}

func (s GetOvertimeSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingResponseBodyResult) SetDefault(v bool) *GetOvertimeSettingResponseBodyResult {
	s.Default = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetDurationSettings(v map[string]*ResultDurationSettingsValue) *GetOvertimeSettingResponseBodyResult {
	s.DurationSettings = v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetId(v int64) *GetOvertimeSettingResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetName(v string) *GetOvertimeSettingResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetOvertimeDivisions(v []*GetOvertimeSettingResponseBodyResultOvertimeDivisions) *GetOvertimeSettingResponseBodyResult {
	s.OvertimeDivisions = v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetSettingId(v int64) *GetOvertimeSettingResponseBodyResult {
	s.SettingId = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetStepType(v int32) *GetOvertimeSettingResponseBodyResult {
	s.StepType = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetStepValue(v float32) *GetOvertimeSettingResponseBodyResult {
	s.StepValue = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetWarningSettings(v []*GetOvertimeSettingResponseBodyResultWarningSettings) *GetOvertimeSettingResponseBodyResult {
	s.WarningSettings = v
	return s
}

func (s *GetOvertimeSettingResponseBodyResult) SetWorkMinutesPerDay(v int32) *GetOvertimeSettingResponseBodyResult {
	s.WorkMinutesPerDay = &v
	return s
}

type GetOvertimeSettingResponseBodyResultOvertimeDivisions struct {
	// 后一日类型
	NextDayType *string `json:"nextDayType,omitempty" xml:"nextDayType,omitempty"`
	// 前一日类型
	PreviousDayType *string `json:"previousDayType,omitempty" xml:"previousDayType,omitempty"`
	// 分割时间点
	TimeSplitPoint *string `json:"timeSplitPoint,omitempty" xml:"timeSplitPoint,omitempty"`
}

func (s GetOvertimeSettingResponseBodyResultOvertimeDivisions) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingResponseBodyResultOvertimeDivisions) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingResponseBodyResultOvertimeDivisions) SetNextDayType(v string) *GetOvertimeSettingResponseBodyResultOvertimeDivisions {
	s.NextDayType = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResultOvertimeDivisions) SetPreviousDayType(v string) *GetOvertimeSettingResponseBodyResultOvertimeDivisions {
	s.PreviousDayType = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResultOvertimeDivisions) SetTimeSplitPoint(v string) *GetOvertimeSettingResponseBodyResultOvertimeDivisions {
	s.TimeSplitPoint = &v
	return s
}

type GetOvertimeSettingResponseBodyResultWarningSettings struct {
	// 风险预警 或 最大加班时间
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 提醒阈值
	Threshold *int64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// 预警类型
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetOvertimeSettingResponseBodyResultWarningSettings) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingResponseBodyResultWarningSettings) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingResponseBodyResultWarningSettings) SetAction(v string) *GetOvertimeSettingResponseBodyResultWarningSettings {
	s.Action = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResultWarningSettings) SetThreshold(v int64) *GetOvertimeSettingResponseBodyResultWarningSettings {
	s.Threshold = &v
	return s
}

func (s *GetOvertimeSettingResponseBodyResultWarningSettings) SetTime(v string) *GetOvertimeSettingResponseBodyResultWarningSettings {
	s.Time = &v
	return s
}

type GetOvertimeSettingResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOvertimeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOvertimeSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOvertimeSettingResponse) GoString() string {
	return s.String()
}

func (s *GetOvertimeSettingResponse) SetHeaders(v map[string]*string) *GetOvertimeSettingResponse {
	s.Headers = v
	return s
}

func (s *GetOvertimeSettingResponse) SetBody(v *GetOvertimeSettingResponseBody) *GetOvertimeSettingResponse {
	s.Body = v
	return s
}

type GetSimpleOvertimeSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSimpleOvertimeSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingHeaders) SetCommonHeaders(v map[string]*string) *GetSimpleOvertimeSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSimpleOvertimeSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetSimpleOvertimeSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSimpleOvertimeSettingRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetSimpleOvertimeSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingRequest) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingRequest) SetPageNumber(v int64) *GetSimpleOvertimeSettingRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSimpleOvertimeSettingRequest) SetPageSize(v int64) *GetSimpleOvertimeSettingRequest {
	s.PageSize = &v
	return s
}

type GetSimpleOvertimeSettingResponseBody struct {
	// Id of the request
	Result []*GetSimpleOvertimeSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetSimpleOvertimeSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingResponseBody) SetResult(v []*GetSimpleOvertimeSettingResponseBodyResult) *GetSimpleOvertimeSettingResponseBody {
	s.Result = v
	return s
}

type GetSimpleOvertimeSettingResponseBodyResult struct {
	// 加班规则集合
	Items []*GetSimpleOvertimeSettingResponseBodyResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 当前页码
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总页数
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetSimpleOvertimeSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingResponseBodyResult) SetItems(v []*GetSimpleOvertimeSettingResponseBodyResultItems) *GetSimpleOvertimeSettingResponseBodyResult {
	s.Items = v
	return s
}

func (s *GetSimpleOvertimeSettingResponseBodyResult) SetPageNumber(v int64) *GetSimpleOvertimeSettingResponseBodyResult {
	s.PageNumber = &v
	return s
}

func (s *GetSimpleOvertimeSettingResponseBodyResult) SetTotalPage(v int64) *GetSimpleOvertimeSettingResponseBodyResult {
	s.TotalPage = &v
	return s
}

type GetSimpleOvertimeSettingResponseBodyResultItems struct {
	// 加班规则id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 加班规则名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetSimpleOvertimeSettingResponseBodyResultItems) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingResponseBodyResultItems) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingResponseBodyResultItems) SetId(v int64) *GetSimpleOvertimeSettingResponseBodyResultItems {
	s.Id = &v
	return s
}

func (s *GetSimpleOvertimeSettingResponseBodyResultItems) SetName(v string) *GetSimpleOvertimeSettingResponseBodyResultItems {
	s.Name = &v
	return s
}

type GetSimpleOvertimeSettingResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSimpleOvertimeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSimpleOvertimeSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingResponse) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingResponse) SetHeaders(v map[string]*string) *GetSimpleOvertimeSettingResponse {
	s.Headers = v
	return s
}

func (s *GetSimpleOvertimeSettingResponse) SetBody(v *GetSimpleOvertimeSettingResponseBody) *GetSimpleOvertimeSettingResponse {
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
	// 假期列表
	Holidays []*GetUserHolidaysResponseBodyResultHolidays `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	// 员工id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserHolidaysResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserHolidaysResponseBodyResult) SetHolidays(v []*GetUserHolidaysResponseBodyResultHolidays) *GetUserHolidaysResponseBodyResult {
	s.Holidays = v
	return s
}

func (s *GetUserHolidaysResponseBodyResult) SetUserId(v string) *GetUserHolidaysResponseBodyResult {
	s.UserId = &v
	return s
}

type GetUserHolidaysResponseBodyResultHolidays struct {
	// 假期名称
	HolidayName *string `json:"holidayName,omitempty" xml:"holidayName,omitempty"`
	// 假期类型，festival：法定节假日；rest：调休日；overtime：加班日；
	HolidayType *string `json:"holidayType,omitempty" xml:"holidayType,omitempty"`
	// 补休日，只有假期类型为加班日时才有值
	RealWorkDate *int64 `json:"realWorkDate,omitempty" xml:"realWorkDate,omitempty"`
	// 日期
	WorkDate *int64 `json:"workDate,omitempty" xml:"workDate,omitempty"`
}

func (s GetUserHolidaysResponseBodyResultHolidays) String() string {
	return tea.Prettify(s)
}

func (s GetUserHolidaysResponseBodyResultHolidays) GoString() string {
	return s.String()
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

func (s *GetUserHolidaysResponseBodyResultHolidays) SetWorkDate(v int64) *GetUserHolidaysResponseBodyResultHolidays {
	s.WorkDate = &v
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

type ProcessApproveCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ProcessApproveCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateHeaders) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateHeaders) SetCommonHeaders(v map[string]*string) *ProcessApproveCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProcessApproveCreateHeaders) SetXAcsDingtalkAccessToken(v string) *ProcessApproveCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ProcessApproveCreateRequest struct {
	// 三方审批单id，全局唯一
	ApproveId *string `json:"approveId,omitempty" xml:"approveId,omitempty"`
	// 审批人员工userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 审批单关联的打卡信息
	PunchParam *ProcessApproveCreateRequestPunchParam `json:"punchParam,omitempty" xml:"punchParam,omitempty" type:"Struct"`
	// 审批单子类型名称：调店:shiftGroup
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// 审批单类型名称
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// 员工的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ProcessApproveCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateRequest) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateRequest) SetApproveId(v string) *ProcessApproveCreateRequest {
	s.ApproveId = &v
	return s
}

func (s *ProcessApproveCreateRequest) SetOpUserId(v string) *ProcessApproveCreateRequest {
	s.OpUserId = &v
	return s
}

func (s *ProcessApproveCreateRequest) SetPunchParam(v *ProcessApproveCreateRequestPunchParam) *ProcessApproveCreateRequest {
	s.PunchParam = v
	return s
}

func (s *ProcessApproveCreateRequest) SetSubType(v string) *ProcessApproveCreateRequest {
	s.SubType = &v
	return s
}

func (s *ProcessApproveCreateRequest) SetTagName(v string) *ProcessApproveCreateRequest {
	s.TagName = &v
	return s
}

func (s *ProcessApproveCreateRequest) SetUserId(v string) *ProcessApproveCreateRequest {
	s.UserId = &v
	return s
}

type ProcessApproveCreateRequestPunchParam struct {
	// 地理位置标识：wifi:ssid_macAddress ble: deviceId gps:longitude_latitude
	PositionId *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	// 地理位置名称
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// 地理位置类型：wifi/ble/gps
	PositionType *string `json:"positionType,omitempty" xml:"positionType,omitempty"`
	// 审批单关联的打卡时间，单位毫秒
	PunchTime *int64 `json:"punchTime,omitempty" xml:"punchTime,omitempty"`
}

func (s ProcessApproveCreateRequestPunchParam) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateRequestPunchParam) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateRequestPunchParam) SetPositionId(v string) *ProcessApproveCreateRequestPunchParam {
	s.PositionId = &v
	return s
}

func (s *ProcessApproveCreateRequestPunchParam) SetPositionName(v string) *ProcessApproveCreateRequestPunchParam {
	s.PositionName = &v
	return s
}

func (s *ProcessApproveCreateRequestPunchParam) SetPositionType(v string) *ProcessApproveCreateRequestPunchParam {
	s.PositionType = &v
	return s
}

func (s *ProcessApproveCreateRequestPunchParam) SetPunchTime(v int64) *ProcessApproveCreateRequestPunchParam {
	s.PunchTime = &v
	return s
}

type ProcessApproveCreateResponseBody struct {
	// 审批单返回对象
	Result *ProcessApproveCreateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ProcessApproveCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateResponseBody) SetResult(v *ProcessApproveCreateResponseBodyResult) *ProcessApproveCreateResponseBody {
	s.Result = v
	return s
}

type ProcessApproveCreateResponseBodyResult struct {
	// 钉钉侧生成的审批单id
	DingtalkApproveId *string `json:"dingtalkApproveId,omitempty" xml:"dingtalkApproveId,omitempty"`
}

func (s ProcessApproveCreateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateResponseBodyResult) SetDingtalkApproveId(v string) *ProcessApproveCreateResponseBodyResult {
	s.DingtalkApproveId = &v
	return s
}

type ProcessApproveCreateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessApproveCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessApproveCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessApproveCreateResponse) GoString() string {
	return s.String()
}

func (s *ProcessApproveCreateResponse) SetHeaders(v map[string]*string) *ProcessApproveCreateResponse {
	s.Headers = v
	return s
}

func (s *ProcessApproveCreateResponse) SetBody(v *ProcessApproveCreateResponseBody) *ProcessApproveCreateResponse {
	s.Body = v
	return s
}

type SyncScheduleInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncScheduleInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncScheduleInfoHeaders) GoString() string {
	return s.String()
}

func (s *SyncScheduleInfoHeaders) SetCommonHeaders(v map[string]*string) *SyncScheduleInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncScheduleInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SyncScheduleInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncScheduleInfoRequest struct {
	OpUserId      *string                                 `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	ScheduleInfos []*SyncScheduleInfoRequestScheduleInfos `json:"scheduleInfos,omitempty" xml:"scheduleInfos,omitempty" type:"Repeated"`
}

func (s SyncScheduleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncScheduleInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncScheduleInfoRequest) SetOpUserId(v string) *SyncScheduleInfoRequest {
	s.OpUserId = &v
	return s
}

func (s *SyncScheduleInfoRequest) SetScheduleInfos(v []*SyncScheduleInfoRequestScheduleInfos) *SyncScheduleInfoRequest {
	s.ScheduleInfos = v
	return s
}

type SyncScheduleInfoRequestScheduleInfos struct {
	PlanId                *int64    `json:"planId,omitempty" xml:"planId,omitempty"`
	PositionKeys          []*string `json:"positionKeys,omitempty" xml:"positionKeys,omitempty" type:"Repeated"`
	RetainAttendanceCheck *bool     `json:"retainAttendanceCheck,omitempty" xml:"retainAttendanceCheck,omitempty"`
	WifiKeys              []*string `json:"wifiKeys,omitempty" xml:"wifiKeys,omitempty" type:"Repeated"`
}

func (s SyncScheduleInfoRequestScheduleInfos) String() string {
	return tea.Prettify(s)
}

func (s SyncScheduleInfoRequestScheduleInfos) GoString() string {
	return s.String()
}

func (s *SyncScheduleInfoRequestScheduleInfos) SetPlanId(v int64) *SyncScheduleInfoRequestScheduleInfos {
	s.PlanId = &v
	return s
}

func (s *SyncScheduleInfoRequestScheduleInfos) SetPositionKeys(v []*string) *SyncScheduleInfoRequestScheduleInfos {
	s.PositionKeys = v
	return s
}

func (s *SyncScheduleInfoRequestScheduleInfos) SetRetainAttendanceCheck(v bool) *SyncScheduleInfoRequestScheduleInfos {
	s.RetainAttendanceCheck = &v
	return s
}

func (s *SyncScheduleInfoRequestScheduleInfos) SetWifiKeys(v []*string) *SyncScheduleInfoRequestScheduleInfos {
	s.WifiKeys = v
	return s
}

type SyncScheduleInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SyncScheduleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncScheduleInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncScheduleInfoResponse) SetHeaders(v map[string]*string) *SyncScheduleInfoResponse {
	s.Headers = v
	return s
}

type UpdateLeaveTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateLeaveTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeHeaders) SetCommonHeaders(v map[string]*string) *UpdateLeaveTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLeaveTypeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateLeaveTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateLeaveTypeRequest struct {
	// 假期类型，普通假期或者加班转调休假期。(general_leave、lieu_leave其中一种)
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 调休假有效期规则(validity_type:有效类型 absolute_time(绝对时间)、relative_time(相对时间)其中一种 validity_value:延长日期(当validity_type为absolute_time该值该值不为空且满足yy-mm格式 validity_type为relative_time该值为大于1的整数))
	Extras *string `json:"extras,omitempty" xml:"extras,omitempty"`
	// 每天折算的工作时长(百分之一 例如1天=10小时=1000)
	HoursInPerDay *int64 `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	// 请假证明
	LeaveCertificate *UpdateLeaveTypeRequestLeaveCertificate `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	// 假期类型唯一标识
	LeaveCode *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	// 假期名称
	LeaveName *string `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	// 请假单位，可以按照天半天或者小时请假。(day、halfDay、hour其中一种)
	LeaveViewUnit *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	// 是否按照自然日统计请假时长，当为false的时候，用户发起请假时候会根据用户在请假时间段内的排班情况来计算请假时长。
	NaturalDayLeave *bool `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	// 限时提交规则
	SubmitTimeRule *UpdateLeaveTypeRequestSubmitTimeRule `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	// 操作者ID
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateLeaveTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeRequest) SetBizType(v string) *UpdateLeaveTypeRequest {
	s.BizType = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetExtras(v string) *UpdateLeaveTypeRequest {
	s.Extras = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetHoursInPerDay(v int64) *UpdateLeaveTypeRequest {
	s.HoursInPerDay = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetLeaveCertificate(v *UpdateLeaveTypeRequestLeaveCertificate) *UpdateLeaveTypeRequest {
	s.LeaveCertificate = v
	return s
}

func (s *UpdateLeaveTypeRequest) SetLeaveCode(v string) *UpdateLeaveTypeRequest {
	s.LeaveCode = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetLeaveName(v string) *UpdateLeaveTypeRequest {
	s.LeaveName = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetLeaveViewUnit(v string) *UpdateLeaveTypeRequest {
	s.LeaveViewUnit = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetNaturalDayLeave(v bool) *UpdateLeaveTypeRequest {
	s.NaturalDayLeave = &v
	return s
}

func (s *UpdateLeaveTypeRequest) SetSubmitTimeRule(v *UpdateLeaveTypeRequestSubmitTimeRule) *UpdateLeaveTypeRequest {
	s.SubmitTimeRule = v
	return s
}

func (s *UpdateLeaveTypeRequest) SetOpUserId(v string) *UpdateLeaveTypeRequest {
	s.OpUserId = &v
	return s
}

type UpdateLeaveTypeRequestLeaveCertificate struct {
	// 超过多长时间需提供请假证明
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否开启请假证明
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 请假提示文案
	PromptInformation *string `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	// 请假证明单位hour，day
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s UpdateLeaveTypeRequestLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeRequestLeaveCertificate) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeRequestLeaveCertificate) SetDuration(v float64) *UpdateLeaveTypeRequestLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *UpdateLeaveTypeRequestLeaveCertificate) SetEnable(v bool) *UpdateLeaveTypeRequestLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *UpdateLeaveTypeRequestLeaveCertificate) SetPromptInformation(v string) *UpdateLeaveTypeRequestLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *UpdateLeaveTypeRequestLeaveCertificate) SetUnit(v string) *UpdateLeaveTypeRequestLeaveCertificate {
	s.Unit = &v
	return s
}

type UpdateLeaveTypeRequestSubmitTimeRule struct {
	// 是否开启限时提交功能：仅且为true时开启
	EnableTimeLimit *bool `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	// 限制类型：before-提前；after-补交
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// 时间单位：day-天；hour-小时
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 限制值：timeUnit=day时，有效值范围[0~30] 天；timeUnit=hour时，有效值范围[0~24] 小时
	TimeValue *int64 `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s UpdateLeaveTypeRequestSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeRequestSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeRequestSubmitTimeRule) SetEnableTimeLimit(v bool) *UpdateLeaveTypeRequestSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *UpdateLeaveTypeRequestSubmitTimeRule) SetTimeType(v string) *UpdateLeaveTypeRequestSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *UpdateLeaveTypeRequestSubmitTimeRule) SetTimeUnit(v string) *UpdateLeaveTypeRequestSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *UpdateLeaveTypeRequestSubmitTimeRule) SetTimeValue(v int64) *UpdateLeaveTypeRequestSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type UpdateLeaveTypeResponseBody struct {
	// 返回参数
	Result *UpdateLeaveTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateLeaveTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponseBody) SetResult(v *UpdateLeaveTypeResponseBodyResult) *UpdateLeaveTypeResponseBody {
	s.Result = v
	return s
}

type UpdateLeaveTypeResponseBodyResult struct {
	// 假期类型，普通假期或者加班转调休假期。(general_leave、lieu_leave其中一种)
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 每天折算的工作时长(百分之一 例如1天=10小时=1000)
	HoursInPerDay *int64 `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	// 请假证明
	LeaveCertificate *UpdateLeaveTypeResponseBodyResultLeaveCertificate `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	// 假期类型唯一标识
	LeaveCode *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	// 假期名称
	LeaveName *string `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	// 请假单位，可以按照天半天或者小时请假。(day、halfDay、hour其中一种)
	LeaveViewUnit *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	// 是否按照自然日统计请假时长，当为false的时候，用户发起请假时候会根据用户在请假时间段内的排班情况来计算请假时长。
	NaturalDayLeave *bool `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	// 限时提交规则
	SubmitTimeRule *UpdateLeaveTypeResponseBodyResultSubmitTimeRule `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
}

func (s UpdateLeaveTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponseBodyResult) SetBizType(v string) *UpdateLeaveTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetHoursInPerDay(v int64) *UpdateLeaveTypeResponseBodyResult {
	s.HoursInPerDay = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetLeaveCertificate(v *UpdateLeaveTypeResponseBodyResultLeaveCertificate) *UpdateLeaveTypeResponseBodyResult {
	s.LeaveCertificate = v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetLeaveCode(v string) *UpdateLeaveTypeResponseBodyResult {
	s.LeaveCode = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetLeaveName(v string) *UpdateLeaveTypeResponseBodyResult {
	s.LeaveName = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetLeaveViewUnit(v string) *UpdateLeaveTypeResponseBodyResult {
	s.LeaveViewUnit = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetNaturalDayLeave(v bool) *UpdateLeaveTypeResponseBodyResult {
	s.NaturalDayLeave = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResult) SetSubmitTimeRule(v *UpdateLeaveTypeResponseBodyResultSubmitTimeRule) *UpdateLeaveTypeResponseBodyResult {
	s.SubmitTimeRule = v
	return s
}

type UpdateLeaveTypeResponseBodyResultLeaveCertificate struct {
	// 超过多长时间需提供请假证明
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否开启请假证明
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 请假提示文案
	PromptInformation *string `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	// 请假证明单位hour，day
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s UpdateLeaveTypeResponseBodyResultLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponseBodyResultLeaveCertificate) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponseBodyResultLeaveCertificate) SetDuration(v float64) *UpdateLeaveTypeResponseBodyResultLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultLeaveCertificate) SetEnable(v bool) *UpdateLeaveTypeResponseBodyResultLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultLeaveCertificate) SetPromptInformation(v string) *UpdateLeaveTypeResponseBodyResultLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultLeaveCertificate) SetUnit(v string) *UpdateLeaveTypeResponseBodyResultLeaveCertificate {
	s.Unit = &v
	return s
}

type UpdateLeaveTypeResponseBodyResultSubmitTimeRule struct {
	// 是否开启限时提交功能：仅且为true时开启
	EnableTimeLimit *bool `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	// 限制类型：before-提前；after-补交
	TimeType *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	// 时间单位：day-天；hour-小时
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	// 限制值：timeUnit=day时，有效值范围[0~30] 天；timeUnit=hour时，有效值范围[0~24] 小时
	TimeValue *int64 `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s UpdateLeaveTypeResponseBodyResultSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponseBodyResultSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponseBodyResultSubmitTimeRule) SetEnableTimeLimit(v bool) *UpdateLeaveTypeResponseBodyResultSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeType(v string) *UpdateLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeUnit(v string) *UpdateLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultSubmitTimeRule) SetTimeValue(v int64) *UpdateLeaveTypeResponseBodyResultSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type UpdateLeaveTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLeaveTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponse) SetHeaders(v map[string]*string) *UpdateLeaveTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLeaveTypeResponse) SetBody(v *UpdateLeaveTypeResponseBody) *UpdateLeaveTypeResponse {
	s.Body = v
	return s
}

type ResultDurationSettingsValue struct {
	CalcType     *int32 `json:"calcType,omitempty" xml:"calcType,omitempty"`
	DurationType *int32 `json:"durationType,omitempty" xml:"durationType,omitempty"`
	// 加班时长计为调休或加班费开关
	OvertimeRedress *bool `json:"overtimeRedress,omitempty" xml:"overtimeRedress,omitempty"`
	// 加班开始时间 或 最小加班时间
	Settings map[string]interface{} `json:"settings,omitempty" xml:"settings,omitempty"`
	// 加班时长计为方式
	OvertimeRedressBy *string `json:"overtimeRedressBy,omitempty" xml:"overtimeRedressBy,omitempty"`
	// 调休时长计算
	VacationRate *float32 `json:"vacationRate,omitempty" xml:"vacationRate,omitempty"`
	// 扣除休息时间
	SkipTime *string `json:"skipTime,omitempty" xml:"skipTime,omitempty"`
	// 休息时段
	SkipTimeByFrames []*ResultDurationSettingsValueSkipTimeByFrames `json:"skipTimeByFrames,omitempty" xml:"skipTimeByFrames,omitempty" type:"Repeated"`
	// 加班时长
	SkipTimeByDurations []*ResultDurationSettingsValueSkipTimeByDurations `json:"skipTimeByDurations,omitempty" xml:"skipTimeByDurations,omitempty" type:"Repeated"`
	// 休息日或节假日排班加班时长计为调休或加班费开关
	HolidayPlanOvertimeRedress *bool `json:"holidayPlanOvertimeRedress,omitempty" xml:"holidayPlanOvertimeRedress,omitempty"`
	// 休息日或节假日排班加班时长计为方式
	HolidayPlanOvertimeRedressBy *string `json:"holidayPlanOvertimeRedressBy,omitempty" xml:"holidayPlanOvertimeRedressBy,omitempty"`
	// 休息日或节假日排班调休时长计算
	HolidayPlanVacationRate *float32 `json:"holidayPlanVacationRate,omitempty" xml:"holidayPlanVacationRate,omitempty"`
}

func (s ResultDurationSettingsValue) String() string {
	return tea.Prettify(s)
}

func (s ResultDurationSettingsValue) GoString() string {
	return s.String()
}

func (s *ResultDurationSettingsValue) SetCalcType(v int32) *ResultDurationSettingsValue {
	s.CalcType = &v
	return s
}

func (s *ResultDurationSettingsValue) SetDurationType(v int32) *ResultDurationSettingsValue {
	s.DurationType = &v
	return s
}

func (s *ResultDurationSettingsValue) SetOvertimeRedress(v bool) *ResultDurationSettingsValue {
	s.OvertimeRedress = &v
	return s
}

func (s *ResultDurationSettingsValue) SetSettings(v map[string]interface{}) *ResultDurationSettingsValue {
	s.Settings = v
	return s
}

func (s *ResultDurationSettingsValue) SetOvertimeRedressBy(v string) *ResultDurationSettingsValue {
	s.OvertimeRedressBy = &v
	return s
}

func (s *ResultDurationSettingsValue) SetVacationRate(v float32) *ResultDurationSettingsValue {
	s.VacationRate = &v
	return s
}

func (s *ResultDurationSettingsValue) SetSkipTime(v string) *ResultDurationSettingsValue {
	s.SkipTime = &v
	return s
}

func (s *ResultDurationSettingsValue) SetSkipTimeByFrames(v []*ResultDurationSettingsValueSkipTimeByFrames) *ResultDurationSettingsValue {
	s.SkipTimeByFrames = v
	return s
}

func (s *ResultDurationSettingsValue) SetSkipTimeByDurations(v []*ResultDurationSettingsValueSkipTimeByDurations) *ResultDurationSettingsValue {
	s.SkipTimeByDurations = v
	return s
}

func (s *ResultDurationSettingsValue) SetHolidayPlanOvertimeRedress(v bool) *ResultDurationSettingsValue {
	s.HolidayPlanOvertimeRedress = &v
	return s
}

func (s *ResultDurationSettingsValue) SetHolidayPlanOvertimeRedressBy(v string) *ResultDurationSettingsValue {
	s.HolidayPlanOvertimeRedressBy = &v
	return s
}

func (s *ResultDurationSettingsValue) SetHolidayPlanVacationRate(v float32) *ResultDurationSettingsValue {
	s.HolidayPlanVacationRate = &v
	return s
}

type ResultDurationSettingsValueSkipTimeByFrames struct {
	// 开始时间，格式为"HH:mm"
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间，格式为"HH:mm"
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 是否生效
	Valid *bool `json:"valid,omitempty" xml:"valid,omitempty"`
}

func (s ResultDurationSettingsValueSkipTimeByFrames) String() string {
	return tea.Prettify(s)
}

func (s ResultDurationSettingsValueSkipTimeByFrames) GoString() string {
	return s.String()
}

func (s *ResultDurationSettingsValueSkipTimeByFrames) SetStartTime(v string) *ResultDurationSettingsValueSkipTimeByFrames {
	s.StartTime = &v
	return s
}

func (s *ResultDurationSettingsValueSkipTimeByFrames) SetEndTime(v string) *ResultDurationSettingsValueSkipTimeByFrames {
	s.EndTime = &v
	return s
}

func (s *ResultDurationSettingsValueSkipTimeByFrames) SetValid(v bool) *ResultDurationSettingsValueSkipTimeByFrames {
	s.Valid = &v
	return s
}

type ResultDurationSettingsValueSkipTimeByDurations struct {
	// 每天加班满 x小时，单位 秒
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 扣除 x小时，单位 秒
	Minus *int64 `json:"minus,omitempty" xml:"minus,omitempty"`
}

func (s ResultDurationSettingsValueSkipTimeByDurations) String() string {
	return tea.Prettify(s)
}

func (s ResultDurationSettingsValueSkipTimeByDurations) GoString() string {
	return s.String()
}

func (s *ResultDurationSettingsValueSkipTimeByDurations) SetDuration(v int64) *ResultDurationSettingsValueSkipTimeByDurations {
	s.Duration = &v
	return s
}

func (s *ResultDurationSettingsValueSkipTimeByDurations) SetMinus(v int64) *ResultDurationSettingsValueSkipTimeByDurations {
	s.Minus = &v
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

func (client *Client) AddLeaveType(request *AddLeaveTypeRequest) (_result *AddLeaveTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLeaveTypeHeaders{}
	_result = &AddLeaveTypeResponse{}
	_body, _err := client.AddLeaveTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLeaveTypeWithOptions(request *AddLeaveTypeRequest, headers *AddLeaveTypeHeaders, runtime *util.RuntimeOptions) (_result *AddLeaveTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Extras)) {
		body["extras"] = request.Extras
	}

	if !tea.BoolValue(util.IsUnset(request.HoursInPerDay)) {
		body["hoursInPerDay"] = request.HoursInPerDay
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LeaveCertificate))) {
		body["leaveCertificate"] = request.LeaveCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveName)) {
		body["leaveName"] = request.LeaveName
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveViewUnit)) {
		body["leaveViewUnit"] = request.LeaveViewUnit
	}

	if !tea.BoolValue(util.IsUnset(request.NaturalDayLeave)) {
		body["naturalDayLeave"] = request.NaturalDayLeave
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubmitTimeRule))) {
		body["submitTimeRule"] = request.SubmitTimeRule
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
	_result = &AddLeaveTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddLeaveType"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/leaves/types"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttendanceBleDevicesAdd(request *AttendanceBleDevicesAddRequest) (_result *AttendanceBleDevicesAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AttendanceBleDevicesAddHeaders{}
	_result = &AttendanceBleDevicesAddResponse{}
	_body, _err := client.AttendanceBleDevicesAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttendanceBleDevicesAddWithOptions(request *AttendanceBleDevicesAddRequest, headers *AttendanceBleDevicesAddHeaders, runtime *util.RuntimeOptions) (_result *AttendanceBleDevicesAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdList)) {
		body["deviceIdList"] = request.DeviceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.GroupKey)) {
		body["groupKey"] = request.GroupKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
	_result = &AttendanceBleDevicesAddResponse{}
	_body, _err := client.DoROARequest(tea.String("AttendanceBleDevicesAdd"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/group/bledevices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttendanceBleDevicesQuery(request *AttendanceBleDevicesQueryRequest) (_result *AttendanceBleDevicesQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AttendanceBleDevicesQueryHeaders{}
	_result = &AttendanceBleDevicesQueryResponse{}
	_body, _err := client.AttendanceBleDevicesQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttendanceBleDevicesQueryWithOptions(request *AttendanceBleDevicesQueryRequest, headers *AttendanceBleDevicesQueryHeaders, runtime *util.RuntimeOptions) (_result *AttendanceBleDevicesQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupKey)) {
		body["groupKey"] = request.GroupKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
	_result = &AttendanceBleDevicesQueryResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("AttendanceBleDevicesQuery"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/group/bledevices/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttendanceBleDevicesRemove(request *AttendanceBleDevicesRemoveRequest) (_result *AttendanceBleDevicesRemoveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AttendanceBleDevicesRemoveHeaders{}
	_result = &AttendanceBleDevicesRemoveResponse{}
	_body, _err := client.AttendanceBleDevicesRemoveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttendanceBleDevicesRemoveWithOptions(request *AttendanceBleDevicesRemoveRequest, headers *AttendanceBleDevicesRemoveHeaders, runtime *util.RuntimeOptions) (_result *AttendanceBleDevicesRemoveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceIdList)) {
		body["deviceIdList"] = request.DeviceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.GroupKey)) {
		body["groupKey"] = request.GroupKey
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
	_result = &AttendanceBleDevicesRemoveResponse{}
	_body, _err := client.DoROARequest(tea.String("AttendanceBleDevicesRemove"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/group/bledevices/remove"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserTimeRange)) {
		body["userTimeRange"] = request.UserTimeRange
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
	_result = &CheckClosingAccountResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckClosingAccount"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/closingAccounts/status/query"), tea.String("json"), req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIds)) {
		body["entityIds"] = request.EntityIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceKey)) {
		body["resourceKey"] = request.ResourceKey
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
	_result = &CheckWritePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckWritePermission"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/writePermissions/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	if !tea.BoolValue(util.IsUnset(request.ApproveId)) {
		body["approveId"] = request.ApproveId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserid)) {
		body["opUserid"] = request.OpUserid
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PunchParam))) {
		body["punchParam"] = request.PunchParam
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		body["subType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
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
	_result = &CreateApproveResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateApprove"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/approves"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DingTalkSecurityCheck(request *DingTalkSecurityCheckRequest) (_result *DingTalkSecurityCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DingTalkSecurityCheckHeaders{}
	_result = &DingTalkSecurityCheckResponse{}
	_body, _err := client.DingTalkSecurityCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DingTalkSecurityCheckWithOptions(request *DingTalkSecurityCheckRequest, headers *DingTalkSecurityCheckHeaders, runtime *util.RuntimeOptions) (_result *DingTalkSecurityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientVer)) {
		body["clientVer"] = request.ClientVer
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformVer)) {
		body["platformVer"] = request.PlatformVer
	}

	if !tea.BoolValue(util.IsUnset(request.Sec)) {
		body["sec"] = request.Sec
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
	_result = &DingTalkSecurityCheckResponse{}
	_body, _err := client.DoROARequest(tea.String("DingTalkSecurityCheck"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/securities/check"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAdjustments(request *GetAdjustmentsRequest) (_result *GetAdjustmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAdjustmentsHeaders{}
	_result = &GetAdjustmentsResponse{}
	_body, _err := client.GetAdjustmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAdjustmentsWithOptions(request *GetAdjustmentsRequest, headers *GetAdjustmentsHeaders, runtime *util.RuntimeOptions) (_result *GetAdjustmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_result = &GetAdjustmentsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAdjustments"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/attendance/adjustments"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) GetLeaveType(request *GetLeaveTypeRequest) (_result *GetLeaveTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLeaveTypeHeaders{}
	_result = &GetLeaveTypeResponse{}
	_body, _err := client.GetLeaveTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLeaveTypeWithOptions(request *GetLeaveTypeRequest, headers *GetLeaveTypeHeaders, runtime *util.RuntimeOptions) (_result *GetLeaveTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.VacationSource)) {
		query["vacationSource"] = request.VacationSource
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
	_result = &GetLeaveTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetLeaveType"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/attendance/leaves/types"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMachine(devId *string) (_result *GetMachineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMachineHeaders{}
	_result = &GetMachineResponse{}
	_body, _err := client.GetMachineWithOptions(devId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMachineWithOptions(devId *string, headers *GetMachineHeaders, runtime *util.RuntimeOptions) (_result *GetMachineResponse, _err error) {
	devId = openapiutil.GetEncodeParam(devId)
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
	_result = &GetMachineResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMachine"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/attendance/machines/"+tea.StringValue(devId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMachineUser(devId *string, request *GetMachineUserRequest) (_result *GetMachineUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMachineUserHeaders{}
	_result = &GetMachineUserResponse{}
	_body, _err := client.GetMachineUserWithOptions(devId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMachineUserWithOptions(devId *string, request *GetMachineUserRequest, headers *GetMachineUserHeaders, runtime *util.RuntimeOptions) (_result *GetMachineUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	devId = openapiutil.GetEncodeParam(devId)
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
	_result = &GetMachineUserResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMachineUser"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/attendance/machines/getUser/"+tea.StringValue(devId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOvertimeSetting(request *GetOvertimeSettingRequest) (_result *GetOvertimeSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOvertimeSettingHeaders{}
	_result = &GetOvertimeSettingResponse{}
	_body, _err := client.GetOvertimeSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOvertimeSettingWithOptions(request *GetOvertimeSettingRequest, headers *GetOvertimeSettingHeaders, runtime *util.RuntimeOptions) (_result *GetOvertimeSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OvertimeSettingIds)) {
		body["overtimeSettingIds"] = request.OvertimeSettingIds
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
	_result = &GetOvertimeSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOvertimeSetting"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/overtimeSettings/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimpleOvertimeSetting(request *GetSimpleOvertimeSettingRequest) (_result *GetSimpleOvertimeSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSimpleOvertimeSettingHeaders{}
	_result = &GetSimpleOvertimeSettingResponse{}
	_body, _err := client.GetSimpleOvertimeSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimpleOvertimeSettingWithOptions(request *GetSimpleOvertimeSettingRequest, headers *GetSimpleOvertimeSettingHeaders, runtime *util.RuntimeOptions) (_result *GetSimpleOvertimeSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_result = &GetSimpleOvertimeSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSimpleOvertimeSetting"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/attendance/overtimeSettings"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) ProcessApproveCreate(request *ProcessApproveCreateRequest) (_result *ProcessApproveCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProcessApproveCreateHeaders{}
	_result = &ProcessApproveCreateResponse{}
	_body, _err := client.ProcessApproveCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessApproveCreateWithOptions(request *ProcessApproveCreateRequest, headers *ProcessApproveCreateHeaders, runtime *util.RuntimeOptions) (_result *ProcessApproveCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApproveId)) {
		body["approveId"] = request.ApproveId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.PunchParam))) {
		body["punchParam"] = request.PunchParam
	}

	if !tea.BoolValue(util.IsUnset(request.SubType)) {
		body["subType"] = request.SubType
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
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
	_result = &ProcessApproveCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("ProcessApproveCreate"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/attendance/workflows/checkInForms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncScheduleInfo(request *SyncScheduleInfoRequest) (_result *SyncScheduleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncScheduleInfoHeaders{}
	_result = &SyncScheduleInfoResponse{}
	_body, _err := client.SyncScheduleInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncScheduleInfoWithOptions(request *SyncScheduleInfoRequest, headers *SyncScheduleInfoHeaders, runtime *util.RuntimeOptions) (_result *SyncScheduleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleInfos)) {
		body["scheduleInfos"] = request.ScheduleInfos
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
	_result = &SyncScheduleInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncScheduleInfo"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/attendance/schedules/additionalInfo"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLeaveType(request *UpdateLeaveTypeRequest) (_result *UpdateLeaveTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateLeaveTypeHeaders{}
	_result = &UpdateLeaveTypeResponse{}
	_body, _err := client.UpdateLeaveTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLeaveTypeWithOptions(request *UpdateLeaveTypeRequest, headers *UpdateLeaveTypeHeaders, runtime *util.RuntimeOptions) (_result *UpdateLeaveTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Extras)) {
		body["extras"] = request.Extras
	}

	if !tea.BoolValue(util.IsUnset(request.HoursInPerDay)) {
		body["hoursInPerDay"] = request.HoursInPerDay
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.LeaveCertificate))) {
		body["leaveCertificate"] = request.LeaveCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveCode)) {
		body["leaveCode"] = request.LeaveCode
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveName)) {
		body["leaveName"] = request.LeaveName
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveViewUnit)) {
		body["leaveViewUnit"] = request.LeaveViewUnit
	}

	if !tea.BoolValue(util.IsUnset(request.NaturalDayLeave)) {
		body["naturalDayLeave"] = request.NaturalDayLeave
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SubmitTimeRule))) {
		body["submitTimeRule"] = request.SubmitTimeRule
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
	_result = &UpdateLeaveTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateLeaveType"), tea.String("attendance_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/attendance/leaves/types"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
