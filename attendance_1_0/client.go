// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package attendance_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ResultDurationSettingsValue struct {
	CalcType                     *int32                                            `json:"calcType,omitempty" xml:"calcType,omitempty"`
	DurationType                 *int32                                            `json:"durationType,omitempty" xml:"durationType,omitempty"`
	OvertimeRedress              *bool                                             `json:"overtimeRedress,omitempty" xml:"overtimeRedress,omitempty"`
	Settings                     map[string]interface{}                            `json:"settings,omitempty" xml:"settings,omitempty"`
	OvertimeRedressBy            *string                                           `json:"overtimeRedressBy,omitempty" xml:"overtimeRedressBy,omitempty"`
	VacationRate                 *float32                                          `json:"vacationRate,omitempty" xml:"vacationRate,omitempty"`
	SkipTime                     *string                                           `json:"skipTime,omitempty" xml:"skipTime,omitempty"`
	SkipTimeByFrames             []*ResultDurationSettingsValueSkipTimeByFrames    `json:"skipTimeByFrames,omitempty" xml:"skipTimeByFrames,omitempty" type:"Repeated"`
	SkipTimeByDurations          []*ResultDurationSettingsValueSkipTimeByDurations `json:"skipTimeByDurations,omitempty" xml:"skipTimeByDurations,omitempty" type:"Repeated"`
	HolidayPlanOvertimeRedress   *bool                                             `json:"holidayPlanOvertimeRedress,omitempty" xml:"holidayPlanOvertimeRedress,omitempty"`
	HolidayPlanOvertimeRedressBy *string                                           `json:"holidayPlanOvertimeRedressBy,omitempty" xml:"holidayPlanOvertimeRedressBy,omitempty"`
	HolidayPlanVacationRate      *float32                                          `json:"holidayPlanVacationRate,omitempty" xml:"holidayPlanVacationRate,omitempty"`
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
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Valid     *bool   `json:"valid,omitempty" xml:"valid,omitempty"`
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
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Minus    *int64 `json:"minus,omitempty" xml:"minus,omitempty"`
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
	BizType          *string                               `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Extras           *string                               `json:"extras,omitempty" xml:"extras,omitempty"`
	HoursInPerDay    *int64                                `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate *AddLeaveTypeRequestLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveName        *string                               `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveViewUnit    *string                               `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	NaturalDayLeave  *bool                                 `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	SubmitTimeRule   *AddLeaveTypeRequestSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	VisibilityRules  []*AddLeaveTypeRequestVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
	OpUserId         *string                               `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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

func (s *AddLeaveTypeRequest) SetVisibilityRules(v []*AddLeaveTypeRequestVisibilityRules) *AddLeaveTypeRequest {
	s.VisibilityRules = v
	return s
}

func (s *AddLeaveTypeRequest) SetOpUserId(v string) *AddLeaveTypeRequest {
	s.OpUserId = &v
	return s
}

type AddLeaveTypeRequestLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
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

type AddLeaveTypeRequestVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s AddLeaveTypeRequestVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeRequestVisibilityRules) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeRequestVisibilityRules) SetType(v string) *AddLeaveTypeRequestVisibilityRules {
	s.Type = &v
	return s
}

func (s *AddLeaveTypeRequestVisibilityRules) SetVisible(v []*string) *AddLeaveTypeRequestVisibilityRules {
	s.Visible = v
	return s
}

type AddLeaveTypeResponseBody struct {
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
	BizType          *string                                          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	HoursInPerDay    *int64                                           `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate *AddLeaveTypeResponseBodyResultLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveCode        *string                                          `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveName        *string                                          `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveViewUnit    *string                                          `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	NaturalDayLeave  *bool                                            `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	SubmitTimeRule   *AddLeaveTypeResponseBodyResultSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	VisibilityRules  []*AddLeaveTypeResponseBodyResultVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
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

func (s *AddLeaveTypeResponseBodyResult) SetVisibilityRules(v []*AddLeaveTypeResponseBodyResultVisibilityRules) *AddLeaveTypeResponseBodyResult {
	s.VisibilityRules = v
	return s
}

type AddLeaveTypeResponseBodyResultLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
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

type AddLeaveTypeResponseBodyResultVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s AddLeaveTypeResponseBodyResultVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s AddLeaveTypeResponseBodyResultVisibilityRules) GoString() string {
	return s.String()
}

func (s *AddLeaveTypeResponseBodyResultVisibilityRules) SetType(v string) *AddLeaveTypeResponseBodyResultVisibilityRules {
	s.Type = &v
	return s
}

func (s *AddLeaveTypeResponseBodyResultVisibilityRules) SetVisible(v []*string) *AddLeaveTypeResponseBodyResultVisibilityRules {
	s.Visible = v
	return s
}

type AddLeaveTypeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddLeaveTypeResponse) SetStatusCode(v int32) *AddLeaveTypeResponse {
	s.StatusCode = &v
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
	DeviceIdList []*int64 `json:"deviceIdList,omitempty" xml:"deviceIdList,omitempty" type:"Repeated"`
	GroupKey     *string  `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
	OpUserId     *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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
	ErrorList   []*AttendanceBleDevicesAddResponseBodyErrorList   `json:"errorList,omitempty" xml:"errorList,omitempty" type:"Repeated"`
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
	Code        *string                                                    `json:"code,omitempty" xml:"code,omitempty"`
	FailureList []*AttendanceBleDevicesAddResponseBodyErrorListFailureList `json:"failureList,omitempty" xml:"failureList,omitempty" type:"Repeated"`
	Msg         *string                                                    `json:"msg,omitempty" xml:"msg,omitempty"`
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
	DeviceId   *int64  `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
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
	DeviceId   *int64  `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttendanceBleDevicesAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttendanceBleDevicesAddResponse) SetStatusCode(v int32) *AttendanceBleDevicesAddResponse {
	s.StatusCode = &v
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
	GroupKey *string `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
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
	DeviceId   *int64  `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttendanceBleDevicesQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttendanceBleDevicesQueryResponse) SetStatusCode(v int32) *AttendanceBleDevicesQueryResponse {
	s.StatusCode = &v
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
	DeviceIdList []*int64 `json:"deviceIdList,omitempty" xml:"deviceIdList,omitempty" type:"Repeated"`
	GroupKey     *string  `json:"groupKey,omitempty" xml:"groupKey,omitempty"`
	OpUserId     *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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
	ErrorList   []*AttendanceBleDevicesRemoveResponseBodyErrorList `json:"errorList,omitempty" xml:"errorList,omitempty" type:"Repeated"`
	SuccessList []*int64                                           `json:"successList,omitempty" xml:"successList,omitempty" type:"Repeated"`
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
	Code        *string  `json:"code,omitempty" xml:"code,omitempty"`
	FailureList []*int64 `json:"failureList,omitempty" xml:"failureList,omitempty" type:"Repeated"`
	Msg         *string  `json:"msg,omitempty" xml:"msg,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttendanceBleDevicesRemoveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttendanceBleDevicesRemoveResponse) SetStatusCode(v int32) *AttendanceBleDevicesRemoveResponse {
	s.StatusCode = &v
	return s
}

func (s *AttendanceBleDevicesRemoveResponse) SetBody(v *AttendanceBleDevicesRemoveResponseBody) *AttendanceBleDevicesRemoveResponse {
	s.Body = v
	return s
}

type BatchBossCheckHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchBossCheckHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchBossCheckHeaders) GoString() string {
	return s.String()
}

func (s *BatchBossCheckHeaders) SetCommonHeaders(v map[string]*string) *BatchBossCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchBossCheckHeaders) SetXAcsDingtalkAccessToken(v string) *BatchBossCheckHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchBossCheckRequest struct {
	Models   []*BatchBossCheckRequestModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	OpUserId *string                        `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s BatchBossCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchBossCheckRequest) GoString() string {
	return s.String()
}

func (s *BatchBossCheckRequest) SetModels(v []*BatchBossCheckRequestModels) *BatchBossCheckRequest {
	s.Models = v
	return s
}

func (s *BatchBossCheckRequest) SetOpUserId(v string) *BatchBossCheckRequest {
	s.OpUserId = &v
	return s
}

type BatchBossCheckRequestModels struct {
	AbsentMin  *int64  `json:"absentMin,omitempty" xml:"absentMin,omitempty"`
	PlanId     *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	TimeResult *string `json:"timeResult,omitempty" xml:"timeResult,omitempty"`
}

func (s BatchBossCheckRequestModels) String() string {
	return tea.Prettify(s)
}

func (s BatchBossCheckRequestModels) GoString() string {
	return s.String()
}

func (s *BatchBossCheckRequestModels) SetAbsentMin(v int64) *BatchBossCheckRequestModels {
	s.AbsentMin = &v
	return s
}

func (s *BatchBossCheckRequestModels) SetPlanId(v int64) *BatchBossCheckRequestModels {
	s.PlanId = &v
	return s
}

func (s *BatchBossCheckRequestModels) SetRemark(v string) *BatchBossCheckRequestModels {
	s.Remark = &v
	return s
}

func (s *BatchBossCheckRequestModels) SetTimeResult(v string) *BatchBossCheckRequestModels {
	s.TimeResult = &v
	return s
}

type BatchBossCheckResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchBossCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchBossCheckResponseBody) GoString() string {
	return s.String()
}

func (s *BatchBossCheckResponseBody) SetResult(v bool) *BatchBossCheckResponseBody {
	s.Result = &v
	return s
}

type BatchBossCheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchBossCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchBossCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchBossCheckResponse) GoString() string {
	return s.String()
}

func (s *BatchBossCheckResponse) SetHeaders(v map[string]*string) *BatchBossCheckResponse {
	s.Headers = v
	return s
}

func (s *BatchBossCheckResponse) SetStatusCode(v int32) *BatchBossCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBossCheckResponse) SetBody(v *BatchBossCheckResponseBody) *BatchBossCheckResponse {
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
	BizCode       *string                                    `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	UserIds       []*string                                  `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckClosingAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckClosingAccountResponse) SetStatusCode(v int32) *CheckClosingAccountResponse {
	s.StatusCode = &v
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
	Category    *string  `json:"category,omitempty" xml:"category,omitempty"`
	EntityIds   []*int64 `json:"entityIds,omitempty" xml:"entityIds,omitempty" type:"Repeated"`
	OpUserId    *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	ResourceKey *string  `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckWritePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckWritePermissionResponse) SetStatusCode(v int32) *CheckWritePermissionResponse {
	s.StatusCode = &v
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
	ApproveId  *string                         `json:"approveId,omitempty" xml:"approveId,omitempty"`
	OpUserid   *string                         `json:"opUserid,omitempty" xml:"opUserid,omitempty"`
	PunchParam *CreateApproveRequestPunchParam `json:"punchParam,omitempty" xml:"punchParam,omitempty" type:"Struct"`
	SubType    *string                         `json:"subType,omitempty" xml:"subType,omitempty"`
	TagName    *string                         `json:"tagName,omitempty" xml:"tagName,omitempty"`
	Userid     *string                         `json:"userid,omitempty" xml:"userid,omitempty"`
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
	PositionId   *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	PositionType *string `json:"positionType,omitempty" xml:"positionType,omitempty"`
	PunchTime    *int64  `json:"punchTime,omitempty" xml:"punchTime,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApproveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateApproveResponse) SetStatusCode(v int32) *CreateApproveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApproveResponse) SetBody(v *CreateApproveResponseBody) *CreateApproveResponse {
	s.Body = v
	return s
}

type DeleteWaterMarkTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteWaterMarkTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaterMarkTemplateHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateHeaders) SetCommonHeaders(v map[string]*string) *DeleteWaterMarkTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWaterMarkTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteWaterMarkTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteWaterMarkTemplateRequest struct {
	FormCode           *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	FormContent        *string `json:"formContent,omitempty" xml:"formContent,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SystemTemplate     *bool   `json:"systemTemplate,omitempty" xml:"systemTemplate,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteWaterMarkTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateRequest) SetFormCode(v string) *DeleteWaterMarkTemplateRequest {
	s.FormCode = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetFormContent(v string) *DeleteWaterMarkTemplateRequest {
	s.FormContent = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetOpenConversationId(v string) *DeleteWaterMarkTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetSystemTemplate(v bool) *DeleteWaterMarkTemplateRequest {
	s.SystemTemplate = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetUserId(v string) *DeleteWaterMarkTemplateRequest {
	s.UserId = &v
	return s
}

type DeleteWaterMarkTemplateResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteWaterMarkTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateResponseBody) SetResult(v string) *DeleteWaterMarkTemplateResponseBody {
	s.Result = &v
	return s
}

type DeleteWaterMarkTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWaterMarkTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *DeleteWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaterMarkTemplateResponse) SetStatusCode(v int32) *DeleteWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaterMarkTemplateResponse) SetBody(v *DeleteWaterMarkTemplateResponseBody) *DeleteWaterMarkTemplateResponse {
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
	ClientVer   *string `json:"clientVer,omitempty" xml:"clientVer,omitempty"`
	Platform    *string `json:"platform,omitempty" xml:"platform,omitempty"`
	PlatformVer *string `json:"platformVer,omitempty" xml:"platformVer,omitempty"`
	Sec         *string `json:"sec,omitempty" xml:"sec,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	HasRisk  *bool              `json:"hasRisk,omitempty" xml:"hasRisk,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DingTalkSecurityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DingTalkSecurityCheckResponse) SetStatusCode(v int32) *DingTalkSecurityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *DingTalkSecurityCheckResponse) SetBody(v *DingTalkSecurityCheckResponseBody) *DingTalkSecurityCheckResponse {
	s.Body = v
	return s
}

type GetATManageScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetATManageScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetATManageScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetATManageScopeHeaders) SetCommonHeaders(v map[string]*string) *GetATManageScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetATManageScopeHeaders) SetXAcsDingtalkAccessToken(v string) *GetATManageScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetATManageScopeRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetATManageScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetATManageScopeRequest) GoString() string {
	return s.String()
}

func (s *GetATManageScopeRequest) SetMaxResults(v int32) *GetATManageScopeRequest {
	s.MaxResults = &v
	return s
}

func (s *GetATManageScopeRequest) SetNextToken(v string) *GetATManageScopeRequest {
	s.NextToken = &v
	return s
}

func (s *GetATManageScopeRequest) SetUserId(v string) *GetATManageScopeRequest {
	s.UserId = &v
	return s
}

type GetATManageScopeResponseBody struct {
	Result []*GetATManageScopeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetATManageScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetATManageScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetATManageScopeResponseBody) SetResult(v []*GetATManageScopeResponseBodyResult) *GetATManageScopeResponseBody {
	s.Result = v
	return s
}

type GetATManageScopeResponseBodyResult struct {
	HasMore     *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	ManageScope *string   `json:"manageScope,omitempty" xml:"manageScope,omitempty"`
	UserIds     []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetATManageScopeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetATManageScopeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetATManageScopeResponseBodyResult) SetHasMore(v bool) *GetATManageScopeResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetATManageScopeResponseBodyResult) SetManageScope(v string) *GetATManageScopeResponseBodyResult {
	s.ManageScope = &v
	return s
}

func (s *GetATManageScopeResponseBodyResult) SetUserIds(v []*string) *GetATManageScopeResponseBodyResult {
	s.UserIds = v
	return s
}

type GetATManageScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetATManageScopeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetATManageScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetATManageScopeResponse) GoString() string {
	return s.String()
}

func (s *GetATManageScopeResponse) SetHeaders(v map[string]*string) *GetATManageScopeResponse {
	s.Headers = v
	return s
}

func (s *GetATManageScopeResponse) SetStatusCode(v int32) *GetATManageScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetATManageScopeResponse) SetBody(v *GetATManageScopeResponseBody) *GetATManageScopeResponse {
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
	Result *GetAdjustmentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdjustmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdjustmentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdjustmentsResponseBody) SetResult(v *GetAdjustmentsResponseBodyResult) *GetAdjustmentsResponseBody {
	s.Result = v
	return s
}

type GetAdjustmentsResponseBodyResult struct {
	Items      []*GetAdjustmentsResponseBodyResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int64                                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	TotalPage  *int64                                   `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
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
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	SettingId *int64  `json:"settingId,omitempty" xml:"settingId,omitempty"`
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

func (s *GetAdjustmentsResponseBodyResultItems) SetSettingId(v int64) *GetAdjustmentsResponseBodyResultItems {
	s.SettingId = &v
	return s
}

type GetAdjustmentsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAdjustmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAdjustmentsResponse) SetStatusCode(v int32) *GetAdjustmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdjustmentsResponse) SetBody(v *GetAdjustmentsResponseBody) *GetAdjustmentsResponse {
	s.Body = v
	return s
}

type GetCheckInSchemaTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCheckInSchemaTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateHeaders) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateHeaders) SetCommonHeaders(v map[string]*string) *GetCheckInSchemaTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCheckInSchemaTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *GetCheckInSchemaTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCheckInSchemaTemplateRequest struct {
	BizCode            *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SceneCode          *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetCheckInSchemaTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateRequest) SetBizCode(v string) *GetCheckInSchemaTemplateRequest {
	s.BizCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateRequest) SetOpenConversationId(v string) *GetCheckInSchemaTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetCheckInSchemaTemplateRequest) SetSceneCode(v string) *GetCheckInSchemaTemplateRequest {
	s.SceneCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateRequest) SetUserId(v string) *GetCheckInSchemaTemplateRequest {
	s.UserId = &v
	return s
}

type GetCheckInSchemaTemplateResponseBody struct {
	Result *GetCheckInSchemaTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetCheckInSchemaTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateResponseBody) SetResult(v *GetCheckInSchemaTemplateResponseBodyResult) *GetCheckInSchemaTemplateResponseBody {
	s.Result = v
	return s
}

type GetCheckInSchemaTemplateResponseBodyResult struct {
	BizCode                 *string                                                              `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	CanModifyAndAddTemplate *bool                                                                `json:"canModifyAndAddTemplate,omitempty" xml:"canModifyAndAddTemplate,omitempty"`
	ConversationAdmin       *bool                                                                `json:"conversationAdmin,omitempty" xml:"conversationAdmin,omitempty"`
	CustomTemplateMaxSize   *int32                                                               `json:"customTemplateMaxSize,omitempty" xml:"customTemplateMaxSize,omitempty"`
	OpenConversationId      *string                                                              `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ShowStat                *bool                                                                `json:"showStat,omitempty" xml:"showStat,omitempty"`
	TemplateDegrade         *bool                                                                `json:"templateDegrade,omitempty" xml:"templateDegrade,omitempty"`
	WaterMarkTemplateModels []*GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels `json:"waterMarkTemplateModels,omitempty" xml:"waterMarkTemplateModels,omitempty" type:"Repeated"`
}

func (s GetCheckInSchemaTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetBizCode(v string) *GetCheckInSchemaTemplateResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetCanModifyAndAddTemplate(v bool) *GetCheckInSchemaTemplateResponseBodyResult {
	s.CanModifyAndAddTemplate = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetConversationAdmin(v bool) *GetCheckInSchemaTemplateResponseBodyResult {
	s.ConversationAdmin = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetCustomTemplateMaxSize(v int32) *GetCheckInSchemaTemplateResponseBodyResult {
	s.CustomTemplateMaxSize = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetOpenConversationId(v string) *GetCheckInSchemaTemplateResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetShowStat(v bool) *GetCheckInSchemaTemplateResponseBodyResult {
	s.ShowStat = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetTemplateDegrade(v bool) *GetCheckInSchemaTemplateResponseBodyResult {
	s.TemplateDegrade = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResult) SetWaterMarkTemplateModels(v []*GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) *GetCheckInSchemaTemplateResponseBodyResult {
	s.WaterMarkTemplateModels = v
	return s
}

type GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels struct {
	CanModify      *bool   `json:"canModify,omitempty" xml:"canModify,omitempty"`
	FormCode       *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	LayoutDesign   *string `json:"layoutDesign,omitempty" xml:"layoutDesign,omitempty"`
	SceneCode      *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
	SchemaContent  *string `json:"schemaContent,omitempty" xml:"schemaContent,omitempty"`
	SuiteKey       *string `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	SystemTemplate *bool   `json:"systemTemplate,omitempty" xml:"systemTemplate,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	WaterMarkId    *string `json:"waterMarkId,omitempty" xml:"waterMarkId,omitempty"`
}

func (s GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetCanModify(v bool) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.CanModify = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetFormCode(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.FormCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetIcon(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.Icon = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetLayoutDesign(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.LayoutDesign = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetSceneCode(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.SceneCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetSchemaContent(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.SchemaContent = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetSuiteKey(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.SuiteKey = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetSystemTemplate(v bool) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.SystemTemplate = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetTitle(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.Title = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels) SetWaterMarkId(v string) *GetCheckInSchemaTemplateResponseBodyResultWaterMarkTemplateModels {
	s.WaterMarkId = &v
	return s
}

type GetCheckInSchemaTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCheckInSchemaTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCheckInSchemaTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckInSchemaTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetCheckInSchemaTemplateResponse) SetHeaders(v map[string]*string) *GetCheckInSchemaTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetCheckInSchemaTemplateResponse) SetStatusCode(v int32) *GetCheckInSchemaTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckInSchemaTemplateResponse) SetBody(v *GetCheckInSchemaTemplateResponseBody) *GetCheckInSchemaTemplateResponse {
	s.Body = v
	return s
}

type GetCheckinRecordByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCheckinRecordByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserHeaders) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserHeaders) SetCommonHeaders(v map[string]*string) *GetCheckinRecordByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCheckinRecordByUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetCheckinRecordByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCheckinRecordByUserRequest struct {
	EndTime        *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxResults     *int64    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OperatorUserId *string   `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	StartTime      *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserIdList     []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetCheckinRecordByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserRequest) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserRequest) SetEndTime(v int64) *GetCheckinRecordByUserRequest {
	s.EndTime = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetMaxResults(v int64) *GetCheckinRecordByUserRequest {
	s.MaxResults = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetNextToken(v int64) *GetCheckinRecordByUserRequest {
	s.NextToken = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetOperatorUserId(v string) *GetCheckinRecordByUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetStartTime(v int64) *GetCheckinRecordByUserRequest {
	s.StartTime = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetUserIdList(v []*string) *GetCheckinRecordByUserRequest {
	s.UserIdList = v
	return s
}

type GetCheckinRecordByUserResponseBody struct {
	Result *GetCheckinRecordByUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetCheckinRecordByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBody) SetResult(v *GetCheckinRecordByUserResponseBodyResult) *GetCheckinRecordByUserResponseBody {
	s.Result = v
	return s
}

type GetCheckinRecordByUserResponseBodyResult struct {
	NextToken *int64                                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PageList  []*GetCheckinRecordByUserResponseBodyResultPageList `json:"pageList,omitempty" xml:"pageList,omitempty" type:"Repeated"`
}

func (s GetCheckinRecordByUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResult) SetNextToken(v int64) *GetCheckinRecordByUserResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResult) SetPageList(v []*GetCheckinRecordByUserResponseBodyResultPageList) *GetCheckinRecordByUserResponseBodyResult {
	s.PageList = v
	return s
}

type GetCheckinRecordByUserResponseBodyResultPageList struct {
	CheckinTime    *int64                                                            `json:"checkinTime,omitempty" xml:"checkinTime,omitempty"`
	CustomDataList []*GetCheckinRecordByUserResponseBodyResultPageListCustomDataList `json:"customDataList,omitempty" xml:"customDataList,omitempty" type:"Repeated"`
	DetailPlace    *string                                                           `json:"detailPlace,omitempty" xml:"detailPlace,omitempty"`
	ImageList      []*string                                                         `json:"imageList,omitempty" xml:"imageList,omitempty" type:"Repeated"`
	Latitude       *string                                                           `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude      *string                                                           `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Place          *string                                                           `json:"place,omitempty" xml:"place,omitempty"`
	Remark         *string                                                           `json:"remark,omitempty" xml:"remark,omitempty"`
	UserId         *string                                                           `json:"userId,omitempty" xml:"userId,omitempty"`
	VisitUser      *string                                                           `json:"visitUser,omitempty" xml:"visitUser,omitempty"`
}

func (s GetCheckinRecordByUserResponseBodyResultPageList) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResultPageList) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetCheckinTime(v int64) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.CheckinTime = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetCustomDataList(v []*GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.CustomDataList = v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetDetailPlace(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.DetailPlace = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetImageList(v []*string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.ImageList = v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetLatitude(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Latitude = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetLongitude(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Longitude = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetPlace(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Place = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetRemark(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Remark = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetUserId(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.UserId = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetVisitUser(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.VisitUser = &v
	return s
}

type GetCheckinRecordByUserResponseBodyResultPageListCustomDataList struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetKey(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Key = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetLabel(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Label = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetValue(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Value = &v
	return s
}

type GetCheckinRecordByUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCheckinRecordByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCheckinRecordByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponse) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponse) SetHeaders(v map[string]*string) *GetCheckinRecordByUserResponse {
	s.Headers = v
	return s
}

func (s *GetCheckinRecordByUserResponse) SetStatusCode(v int32) *GetCheckinRecordByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckinRecordByUserResponse) SetBody(v *GetCheckinRecordByUserResponseBody) *GetCheckinRecordByUserResponse {
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
	ClosingAccountModel       *GetClosingAccountsResponseBodyResultClosingAccountModel       `json:"closingAccountModel,omitempty" xml:"closingAccountModel,omitempty" type:"Struct"`
	SwitchOn                  *bool                                                          `json:"switchOn,omitempty" xml:"switchOn,omitempty"`
	UnsealClosingAccountModel *GetClosingAccountsResponseBodyResultUnsealClosingAccountModel `json:"unsealClosingAccountModel,omitempty" xml:"unsealClosingAccountModel,omitempty" type:"Struct"`
	UserId                    *string                                                        `json:"userId,omitempty" xml:"userId,omitempty"`
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
	ClosingDay         *int32 `json:"closingDay,omitempty" xml:"closingDay,omitempty"`
	ClosingHourMinutes *int64 `json:"closingHourMinutes,omitempty" xml:"closingHourMinutes,omitempty"`
	EndDay             *int32 `json:"endDay,omitempty" xml:"endDay,omitempty"`
	EndMonth           *int32 `json:"endMonth,omitempty" xml:"endMonth,omitempty"`
	StartDay           *int32 `json:"startDay,omitempty" xml:"startDay,omitempty"`
	StartMonth         *int32 `json:"startMonth,omitempty" xml:"startMonth,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClosingAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetClosingAccountsResponse) SetStatusCode(v int32) *GetClosingAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClosingAccountsResponse) SetBody(v *GetClosingAccountsResponseBody) *GetClosingAccountsResponse {
	s.Body = v
	return s
}

type GetLeaveRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLeaveRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetLeaveRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLeaveRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetLeaveRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLeaveRecordsRequest struct {
	LeaveCode  *string   `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	OpUserId   *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	PageNumber *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetLeaveRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsRequest) SetLeaveCode(v string) *GetLeaveRecordsRequest {
	s.LeaveCode = &v
	return s
}

func (s *GetLeaveRecordsRequest) SetOpUserId(v string) *GetLeaveRecordsRequest {
	s.OpUserId = &v
	return s
}

func (s *GetLeaveRecordsRequest) SetPageNumber(v int64) *GetLeaveRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLeaveRecordsRequest) SetPageSize(v int32) *GetLeaveRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *GetLeaveRecordsRequest) SetUserIds(v []*string) *GetLeaveRecordsRequest {
	s.UserIds = v
	return s
}

type GetLeaveRecordsResponseBody struct {
	Result  *GetLeaveRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetLeaveRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsResponseBody) SetResult(v *GetLeaveRecordsResponseBodyResult) *GetLeaveRecordsResponseBody {
	s.Result = v
	return s
}

func (s *GetLeaveRecordsResponseBody) SetSuccess(v bool) *GetLeaveRecordsResponseBody {
	s.Success = &v
	return s
}

type GetLeaveRecordsResponseBodyResult struct {
	HasMore      *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	LeaveRecords []*GetLeaveRecordsResponseBodyResultLeaveRecords `json:"leaveRecords,omitempty" xml:"leaveRecords,omitempty" type:"Repeated"`
}

func (s GetLeaveRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsResponseBodyResult) SetHasMore(v bool) *GetLeaveRecordsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResult) SetLeaveRecords(v []*GetLeaveRecordsResponseBodyResultLeaveRecords) *GetLeaveRecordsResponseBodyResult {
	s.LeaveRecords = v
	return s
}

type GetLeaveRecordsResponseBodyResultLeaveRecords struct {
	CalType          *string `json:"calType,omitempty" xml:"calType,omitempty"`
	EndTime          *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtCreate        *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	LeaveCode        *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveReason      *string `json:"leaveReason,omitempty" xml:"leaveReason,omitempty"`
	LeaveRecordType  *string `json:"leaveRecordType,omitempty" xml:"leaveRecordType,omitempty"`
	LeaveStatus      *string `json:"leaveStatus,omitempty" xml:"leaveStatus,omitempty"`
	LeaveViewUnit    *string `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	OpUserId         *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	QuotaId          *string `json:"quotaId,omitempty" xml:"quotaId,omitempty"`
	RecordId         *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	RecordNumPerDay  *int64  `json:"recordNumPerDay,omitempty" xml:"recordNumPerDay,omitempty"`
	RecordNumPerHour *int64  `json:"recordNumPerHour,omitempty" xml:"recordNumPerHour,omitempty"`
	StartTime        *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetLeaveRecordsResponseBodyResultLeaveRecords) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsResponseBodyResultLeaveRecords) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetCalType(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.CalType = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetEndTime(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.EndTime = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetGmtCreate(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetGmtModified(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.GmtModified = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetLeaveCode(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.LeaveCode = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetLeaveReason(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.LeaveReason = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetLeaveRecordType(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.LeaveRecordType = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetLeaveStatus(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.LeaveStatus = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetLeaveViewUnit(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.LeaveViewUnit = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetOpUserId(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.OpUserId = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetQuotaId(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.QuotaId = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetRecordId(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.RecordId = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetRecordNumPerDay(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.RecordNumPerDay = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetRecordNumPerHour(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.RecordNumPerHour = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetStartTime(v int64) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.StartTime = &v
	return s
}

func (s *GetLeaveRecordsResponseBodyResultLeaveRecords) SetUserId(v string) *GetLeaveRecordsResponseBodyResultLeaveRecords {
	s.UserId = &v
	return s
}

type GetLeaveRecordsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLeaveRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLeaveRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetLeaveRecordsResponse) SetHeaders(v map[string]*string) *GetLeaveRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetLeaveRecordsResponse) SetStatusCode(v int32) *GetLeaveRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLeaveRecordsResponse) SetBody(v *GetLeaveRecordsResponseBody) *GetLeaveRecordsResponse {
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
	OpUserId       *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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
	BizType          *string                                          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	HoursInPerDay    *int64                                           `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate *GetLeaveTypeResponseBodyResultLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveCode        *string                                          `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveName        *string                                          `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveViewUnit    *string                                          `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	NaturalDayLeave  *bool                                            `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	Source           *string                                          `json:"source,omitempty" xml:"source,omitempty"`
	SubmitTimeRule   *GetLeaveTypeResponseBodyResultSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	ValidityType     *string                                          `json:"validityType,omitempty" xml:"validityType,omitempty"`
	ValidityValue    *string                                          `json:"validityValue,omitempty" xml:"validityValue,omitempty"`
	VisibilityRules  []*GetLeaveTypeResponseBodyResultVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
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

func (s *GetLeaveTypeResponseBodyResult) SetVisibilityRules(v []*GetLeaveTypeResponseBodyResultVisibilityRules) *GetLeaveTypeResponseBodyResult {
	s.VisibilityRules = v
	return s
}

type GetLeaveTypeResponseBodyResultLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
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

type GetLeaveTypeResponseBodyResultVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s GetLeaveTypeResponseBodyResultVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s GetLeaveTypeResponseBodyResultVisibilityRules) GoString() string {
	return s.String()
}

func (s *GetLeaveTypeResponseBodyResultVisibilityRules) SetType(v string) *GetLeaveTypeResponseBodyResultVisibilityRules {
	s.Type = &v
	return s
}

func (s *GetLeaveTypeResponseBodyResultVisibilityRules) SetVisible(v []*string) *GetLeaveTypeResponseBodyResultVisibilityRules {
	s.Visible = v
	return s
}

type GetLeaveTypeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetLeaveTypeResponse) SetStatusCode(v int32) *GetLeaveTypeResponse {
	s.StatusCode = &v
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
	AtmManagerList     []*string                                       `json:"atmManagerList,omitempty" xml:"atmManagerList,omitempty" type:"Repeated"`
	DevId              *int64                                          `json:"devId,omitempty" xml:"devId,omitempty"`
	DeviceId           *string                                         `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	DeviceName         *string                                         `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceSn           *string                                         `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	MachineBluetoothVO *GetMachineResponseBodyResultMachineBluetoothVO `json:"machineBluetoothVO,omitempty" xml:"machineBluetoothVO,omitempty" type:"Struct"`
	MaxFace            *int32                                          `json:"maxFace,omitempty" xml:"maxFace,omitempty"`
	NetStatus          *string                                         `json:"netStatus,omitempty" xml:"netStatus,omitempty"`
	ProductName        *string                                         `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion     *string                                         `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	VoiceMode          *int32                                          `json:"voiceMode,omitempty" xml:"voiceMode,omitempty"`
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
	Address                   *string  `json:"address,omitempty" xml:"address,omitempty"`
	BluetoothCheckWithFace    *bool    `json:"bluetoothCheckWithFace,omitempty" xml:"bluetoothCheckWithFace,omitempty"`
	BluetoothDistanceMode     *string  `json:"bluetoothDistanceMode,omitempty" xml:"bluetoothDistanceMode,omitempty"`
	BluetoothDistanceModeDesc *string  `json:"bluetoothDistanceModeDesc,omitempty" xml:"bluetoothDistanceModeDesc,omitempty"`
	BluetoothValue            *bool    `json:"bluetoothValue,omitempty" xml:"bluetoothValue,omitempty"`
	Latitude                  *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	LimitUserDeviceCount      *bool    `json:"limitUserDeviceCount,omitempty" xml:"limitUserDeviceCount,omitempty"`
	Longitude                 *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	MonitorLocationAbnormal   *bool    `json:"monitorLocationAbnormal,omitempty" xml:"monitorLocationAbnormal,omitempty"`
	UserDeviceCount           *int32   `json:"userDeviceCount,omitempty" xml:"userDeviceCount,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMachineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMachineResponse) SetStatusCode(v int32) *GetMachineResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMachineUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMachineUserResponse) SetStatusCode(v int32) *GetMachineUserResponse {
	s.StatusCode = &v
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
	Default           *bool                                                    `json:"default,omitempty" xml:"default,omitempty"`
	DurationSettings  map[string]*ResultDurationSettingsValue                  `json:"durationSettings,omitempty" xml:"durationSettings,omitempty"`
	Id                *int64                                                   `json:"id,omitempty" xml:"id,omitempty"`
	Name              *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
	OvertimeDivisions []*GetOvertimeSettingResponseBodyResultOvertimeDivisions `json:"overtimeDivisions,omitempty" xml:"overtimeDivisions,omitempty" type:"Repeated"`
	SettingId         *int64                                                   `json:"settingId,omitempty" xml:"settingId,omitempty"`
	StepType          *int32                                                   `json:"stepType,omitempty" xml:"stepType,omitempty"`
	StepValue         *float32                                                 `json:"stepValue,omitempty" xml:"stepValue,omitempty"`
	WarningSettings   []*GetOvertimeSettingResponseBodyResultWarningSettings   `json:"warningSettings,omitempty" xml:"warningSettings,omitempty" type:"Repeated"`
	WorkMinutesPerDay *int32                                                   `json:"workMinutesPerDay,omitempty" xml:"workMinutesPerDay,omitempty"`
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
	NextDayType     *string `json:"nextDayType,omitempty" xml:"nextDayType,omitempty"`
	PreviousDayType *string `json:"previousDayType,omitempty" xml:"previousDayType,omitempty"`
	TimeSplitPoint  *string `json:"timeSplitPoint,omitempty" xml:"timeSplitPoint,omitempty"`
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
	Action    *string `json:"action,omitempty" xml:"action,omitempty"`
	Threshold *int64  `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Time      *string `json:"time,omitempty" xml:"time,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOvertimeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetOvertimeSettingResponse) SetStatusCode(v int32) *GetOvertimeSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOvertimeSettingResponse) SetBody(v *GetOvertimeSettingResponseBody) *GetOvertimeSettingResponse {
	s.Body = v
	return s
}

type GetShiftHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetShiftHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetShiftHeaders) GoString() string {
	return s.String()
}

func (s *GetShiftHeaders) SetCommonHeaders(v map[string]*string) *GetShiftHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetShiftHeaders) SetXAcsDingtalkAccessToken(v string) *GetShiftHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetShiftRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	ShiftId  *int64  `json:"shiftId,omitempty" xml:"shiftId,omitempty"`
}

func (s GetShiftRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShiftRequest) GoString() string {
	return s.String()
}

func (s *GetShiftRequest) SetOpUserId(v string) *GetShiftRequest {
	s.OpUserId = &v
	return s
}

func (s *GetShiftRequest) SetShiftId(v int64) *GetShiftRequest {
	s.ShiftId = &v
	return s
}

type GetShiftResponseBody struct {
	Result *GetShiftResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetShiftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBody) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBody) SetResult(v *GetShiftResponseBodyResult) *GetShiftResponseBody {
	s.Result = v
	return s
}

type GetShiftResponseBodyResult struct {
	CorpId         *string                                 `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Id             *int64                                  `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Owner          *string                                 `json:"owner,omitempty" xml:"owner,omitempty"`
	Sections       []*GetShiftResponseBodyResultSections   `json:"sections,omitempty" xml:"sections,omitempty" type:"Repeated"`
	ShiftGroupId   *int64                                  `json:"shiftGroupId,omitempty" xml:"shiftGroupId,omitempty"`
	ShiftGroupName *string                                 `json:"shiftGroupName,omitempty" xml:"shiftGroupName,omitempty"`
	ShiftSetting   *GetShiftResponseBodyResultShiftSetting `json:"shiftSetting,omitempty" xml:"shiftSetting,omitempty" type:"Struct"`
}

func (s GetShiftResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResult) SetCorpId(v string) *GetShiftResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetId(v int64) *GetShiftResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetName(v string) *GetShiftResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetOwner(v string) *GetShiftResponseBodyResult {
	s.Owner = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetSections(v []*GetShiftResponseBodyResultSections) *GetShiftResponseBodyResult {
	s.Sections = v
	return s
}

func (s *GetShiftResponseBodyResult) SetShiftGroupId(v int64) *GetShiftResponseBodyResult {
	s.ShiftGroupId = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetShiftGroupName(v string) *GetShiftResponseBodyResult {
	s.ShiftGroupName = &v
	return s
}

func (s *GetShiftResponseBodyResult) SetShiftSetting(v *GetShiftResponseBodyResultShiftSetting) *GetShiftResponseBodyResult {
	s.ShiftSetting = v
	return s
}

type GetShiftResponseBodyResultSections struct {
	Punches   []*GetShiftResponseBodyResultSectionsPunches `json:"punches,omitempty" xml:"punches,omitempty" type:"Repeated"`
	Rests     []*GetShiftResponseBodyResultSectionsRests   `json:"rests,omitempty" xml:"rests,omitempty" type:"Repeated"`
	SectionId *int64                                       `json:"sectionId,omitempty" xml:"sectionId,omitempty"`
}

func (s GetShiftResponseBodyResultSections) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultSections) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultSections) SetPunches(v []*GetShiftResponseBodyResultSectionsPunches) *GetShiftResponseBodyResultSections {
	s.Punches = v
	return s
}

func (s *GetShiftResponseBodyResultSections) SetRests(v []*GetShiftResponseBodyResultSectionsRests) *GetShiftResponseBodyResultSections {
	s.Rests = v
	return s
}

func (s *GetShiftResponseBodyResultSections) SetSectionId(v int64) *GetShiftResponseBodyResultSections {
	s.SectionId = &v
	return s
}

type GetShiftResponseBodyResultSectionsPunches struct {
	AbsenteeismLateMinutes *int64                                                    `json:"absenteeismLateMinutes,omitempty" xml:"absenteeismLateMinutes,omitempty"`
	Across                 *int64                                                    `json:"across,omitempty" xml:"across,omitempty"`
	BeginMin               *int64                                                    `json:"beginMin,omitempty" xml:"beginMin,omitempty"`
	CheckTime              *string                                                   `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	CheckType              *string                                                   `json:"checkType,omitempty" xml:"checkType,omitempty"`
	EndMin                 *int64                                                    `json:"end_min,omitempty" xml:"end_min,omitempty"`
	FreeCheck              *bool                                                     `json:"freeCheck,omitempty" xml:"freeCheck,omitempty"`
	LateBackSetting        *GetShiftResponseBodyResultSectionsPunchesLateBackSetting `json:"lateBackSetting,omitempty" xml:"lateBackSetting,omitempty" type:"Struct"`
	PermitMinutes          *int64                                                    `json:"permitMinutes,omitempty" xml:"permitMinutes,omitempty"`
	PuncheId               *int64                                                    `json:"puncheId,omitempty" xml:"puncheId,omitempty"`
	SeriousLateMinutes     *int64                                                    `json:"seriousLateMinutes,omitempty" xml:"seriousLateMinutes,omitempty"`
}

func (s GetShiftResponseBodyResultSectionsPunches) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultSectionsPunches) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetAbsenteeismLateMinutes(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.AbsenteeismLateMinutes = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetAcross(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.Across = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetBeginMin(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.BeginMin = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetCheckTime(v string) *GetShiftResponseBodyResultSectionsPunches {
	s.CheckTime = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetCheckType(v string) *GetShiftResponseBodyResultSectionsPunches {
	s.CheckType = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetEndMin(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.EndMin = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetFreeCheck(v bool) *GetShiftResponseBodyResultSectionsPunches {
	s.FreeCheck = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetLateBackSetting(v *GetShiftResponseBodyResultSectionsPunchesLateBackSetting) *GetShiftResponseBodyResultSectionsPunches {
	s.LateBackSetting = v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetPermitMinutes(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.PermitMinutes = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetPuncheId(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.PuncheId = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunches) SetSeriousLateMinutes(v int64) *GetShiftResponseBodyResultSectionsPunches {
	s.SeriousLateMinutes = &v
	return s
}

type GetShiftResponseBodyResultSectionsPunchesLateBackSetting struct {
	LateBackPairs []*GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs `json:"lateBackPairs,omitempty" xml:"lateBackPairs,omitempty" type:"Repeated"`
}

func (s GetShiftResponseBodyResultSectionsPunchesLateBackSetting) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultSectionsPunchesLateBackSetting) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultSectionsPunchesLateBackSetting) SetLateBackPairs(v []*GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs) *GetShiftResponseBodyResultSectionsPunchesLateBackSetting {
	s.LateBackPairs = v
	return s
}

type GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs struct {
	Extra *int64 `json:"extra,omitempty" xml:"extra,omitempty"`
	Late  *int64 `json:"late,omitempty" xml:"late,omitempty"`
}

func (s GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs) SetExtra(v int64) *GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs {
	s.Extra = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs) SetLate(v int64) *GetShiftResponseBodyResultSectionsPunchesLateBackSettingLateBackPairs {
	s.Late = &v
	return s
}

type GetShiftResponseBodyResultSectionsRests struct {
	Across    *int64  `json:"across,omitempty" xml:"across,omitempty"`
	CheckTime *string `json:"checkTime,omitempty" xml:"checkTime,omitempty"`
	CheckType *string `json:"check_type,omitempty" xml:"check_type,omitempty"`
	RestId    *int64  `json:"restId,omitempty" xml:"restId,omitempty"`
}

func (s GetShiftResponseBodyResultSectionsRests) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultSectionsRests) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultSectionsRests) SetAcross(v int64) *GetShiftResponseBodyResultSectionsRests {
	s.Across = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsRests) SetCheckTime(v string) *GetShiftResponseBodyResultSectionsRests {
	s.CheckTime = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsRests) SetCheckType(v string) *GetShiftResponseBodyResultSectionsRests {
	s.CheckType = &v
	return s
}

func (s *GetShiftResponseBodyResultSectionsRests) SetRestId(v int64) *GetShiftResponseBodyResultSectionsRests {
	s.RestId = &v
	return s
}

type GetShiftResponseBodyResultShiftSetting struct {
	AttendDays      *string `json:"attendDays,omitempty" xml:"attendDays,omitempty"`
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GmtCreate       *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ShiftId         *int64  `json:"shiftId,omitempty" xml:"shiftId,omitempty"`
	ShiftSettingId  *int64  `json:"shiftSettingId,omitempty" xml:"shiftSettingId,omitempty"`
	WorkTimeMinutes *int64  `json:"workTimeMinutes,omitempty" xml:"workTimeMinutes,omitempty"`
}

func (s GetShiftResponseBodyResultShiftSetting) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponseBodyResultShiftSetting) GoString() string {
	return s.String()
}

func (s *GetShiftResponseBodyResultShiftSetting) SetAttendDays(v string) *GetShiftResponseBodyResultShiftSetting {
	s.AttendDays = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetCorpId(v string) *GetShiftResponseBodyResultShiftSetting {
	s.CorpId = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetGmtCreate(v string) *GetShiftResponseBodyResultShiftSetting {
	s.GmtCreate = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetGmtModified(v string) *GetShiftResponseBodyResultShiftSetting {
	s.GmtModified = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetShiftId(v int64) *GetShiftResponseBodyResultShiftSetting {
	s.ShiftId = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetShiftSettingId(v int64) *GetShiftResponseBodyResultShiftSetting {
	s.ShiftSettingId = &v
	return s
}

func (s *GetShiftResponseBodyResultShiftSetting) SetWorkTimeMinutes(v int64) *GetShiftResponseBodyResultShiftSetting {
	s.WorkTimeMinutes = &v
	return s
}

type GetShiftResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShiftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShiftResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShiftResponse) GoString() string {
	return s.String()
}

func (s *GetShiftResponse) SetHeaders(v map[string]*string) *GetShiftResponse {
	s.Headers = v
	return s
}

func (s *GetShiftResponse) SetStatusCode(v int32) *GetShiftResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShiftResponse) SetBody(v *GetShiftResponseBody) *GetShiftResponse {
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
	Result *GetSimpleOvertimeSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSimpleOvertimeSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimpleOvertimeSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimpleOvertimeSettingResponseBody) SetResult(v *GetSimpleOvertimeSettingResponseBodyResult) *GetSimpleOvertimeSettingResponseBody {
	s.Result = v
	return s
}

type GetSimpleOvertimeSettingResponseBodyResult struct {
	Items      []*GetSimpleOvertimeSettingResponseBodyResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int64                                             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	TotalPage  *int64                                             `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
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
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	SettingId *int64  `json:"settingId,omitempty" xml:"settingId,omitempty"`
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

func (s *GetSimpleOvertimeSettingResponseBodyResultItems) SetSettingId(v int64) *GetSimpleOvertimeSettingResponseBodyResultItems {
	s.SettingId = &v
	return s
}

type GetSimpleOvertimeSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSimpleOvertimeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSimpleOvertimeSettingResponse) SetStatusCode(v int32) *GetSimpleOvertimeSettingResponse {
	s.StatusCode = &v
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
	UserIds      []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	WorkDateFrom *int64    `json:"workDateFrom,omitempty" xml:"workDateFrom,omitempty"`
	WorkDateTo   *int64    `json:"workDateTo,omitempty" xml:"workDateTo,omitempty"`
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
	Holidays []*GetUserHolidaysResponseBodyResultHolidays `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	UserId   *string                                      `json:"userId,omitempty" xml:"userId,omitempty"`
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
	HolidayName  *string `json:"holidayName,omitempty" xml:"holidayName,omitempty"`
	HolidayType  *string `json:"holidayType,omitempty" xml:"holidayType,omitempty"`
	RealWorkDate *int64  `json:"realWorkDate,omitempty" xml:"realWorkDate,omitempty"`
	WorkDate     *int64  `json:"workDate,omitempty" xml:"workDate,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserHolidaysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserHolidaysResponse) SetStatusCode(v int32) *GetUserHolidaysResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserHolidaysResponse) SetBody(v *GetUserHolidaysResponseBody) *GetUserHolidaysResponse {
	s.Body = v
	return s
}

type GroupAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupAddHeaders) GoString() string {
	return s.String()
}

func (s *GroupAddHeaders) SetCommonHeaders(v map[string]*string) *GroupAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupAddHeaders) SetXAcsDingtalkAccessToken(v string) *GroupAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupAddRequest struct {
	AdjustmentSettingId            *int64                           `json:"adjustmentSettingId,omitempty" xml:"adjustmentSettingId,omitempty"`
	BleDeviceList                  []*GroupAddRequestBleDeviceList  `json:"bleDeviceList,omitempty" xml:"bleDeviceList,omitempty" type:"Repeated"`
	CheckNeedHealthyCode           *bool                            `json:"checkNeedHealthyCode,omitempty" xml:"checkNeedHealthyCode,omitempty"`
	DefaultClassId                 *int64                           `json:"defaultClassId,omitempty" xml:"defaultClassId,omitempty"`
	DisableCheckWhenRest           *bool                            `json:"disableCheckWhenRest,omitempty" xml:"disableCheckWhenRest,omitempty"`
	DisableCheckWithoutSchedule    *bool                            `json:"disableCheckWithoutSchedule,omitempty" xml:"disableCheckWithoutSchedule,omitempty"`
	EnableCameraCheck              *bool                            `json:"enableCameraCheck,omitempty" xml:"enableCameraCheck,omitempty"`
	EnableEmpSelectClass           *bool                            `json:"enableEmpSelectClass,omitempty" xml:"enableEmpSelectClass,omitempty"`
	EnableFaceCheck                *bool                            `json:"enableFaceCheck,omitempty" xml:"enableFaceCheck,omitempty"`
	EnableFaceStrictMode           *bool                            `json:"enableFaceStrictMode,omitempty" xml:"enableFaceStrictMode,omitempty"`
	EnableNextDay                  *bool                            `json:"enableNextDay,omitempty" xml:"enableNextDay,omitempty"`
	EnableOutSideUpdateNormalCheck *bool                            `json:"enableOutSideUpdateNormalCheck,omitempty" xml:"enableOutSideUpdateNormalCheck,omitempty"`
	EnableOutsideApply             *bool                            `json:"enableOutsideApply,omitempty" xml:"enableOutsideApply,omitempty"`
	EnableOutsideCameraCheck       *bool                            `json:"enableOutsideCameraCheck,omitempty" xml:"enableOutsideCameraCheck,omitempty"`
	EnableOutsideCheck             *bool                            `json:"enableOutsideCheck,omitempty" xml:"enableOutsideCheck,omitempty"`
	EnableOutsideRemark            *bool                            `json:"enableOutsideRemark,omitempty" xml:"enableOutsideRemark,omitempty"`
	EnablePositionBle              *bool                            `json:"enablePositionBle,omitempty" xml:"enablePositionBle,omitempty"`
	EnableTrimDistance             *bool                            `json:"enableTrimDistance,omitempty" xml:"enableTrimDistance,omitempty"`
	ForbidHideOutSideAddress       *bool                            `json:"forbidHideOutSideAddress,omitempty" xml:"forbidHideOutSideAddress,omitempty"`
	FreeCheckSetting               *GroupAddRequestFreeCheckSetting `json:"freeCheckSetting,omitempty" xml:"freeCheckSetting,omitempty" type:"Struct"`
	FreeCheckTypeId                *int32                           `json:"freeCheckTypeId,omitempty" xml:"freeCheckTypeId,omitempty"`
	FreecheckDayStartMinOffset     *int32                           `json:"freecheckDayStartMinOffset,omitempty" xml:"freecheckDayStartMinOffset,omitempty"`
	FreecheckWorkDays              []*int64                         `json:"freecheckWorkDays,omitempty" xml:"freecheckWorkDays,omitempty" type:"Repeated"`
	GroupId                        *int64                           `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName                      *string                          `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ManagerList                    []*string                        `json:"managerList,omitempty" xml:"managerList,omitempty" type:"Repeated"`
	Members                        []*GroupAddRequestMembers        `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	ModifyMember                   *bool                            `json:"modifyMember,omitempty" xml:"modifyMember,omitempty"`
	Offset                         *int32                           `json:"offset,omitempty" xml:"offset,omitempty"`
	OpenFaceCheck                  *bool                            `json:"openFaceCheck,omitempty" xml:"openFaceCheck,omitempty"`
	OutsideCheckApproveModeId      *int32                           `json:"outsideCheckApproveModeId,omitempty" xml:"outsideCheckApproveModeId,omitempty"`
	OvertimeSettingId              *int64                           `json:"overtimeSettingId,omitempty" xml:"overtimeSettingId,omitempty"`
	Owner                          *string                          `json:"owner,omitempty" xml:"owner,omitempty"`
	Positions                      []*GroupAddRequestPositions      `json:"positions,omitempty" xml:"positions,omitempty" type:"Repeated"`
	ResourcePermissionMap          map[string]interface{}           `json:"resourcePermissionMap,omitempty" xml:"resourcePermissionMap,omitempty"`
	ShiftVOList                    []*GroupAddRequestShiftVOList    `json:"shiftVOList,omitempty" xml:"shiftVOList,omitempty" type:"Repeated"`
	SkipHolidays                   *bool                            `json:"skipHolidays,omitempty" xml:"skipHolidays,omitempty"`
	SpecialDays                    *string                          `json:"specialDays,omitempty" xml:"specialDays,omitempty"`
	TrimDistance                   *int32                           `json:"trimDistance,omitempty" xml:"trimDistance,omitempty"`
	Type                           *string                          `json:"type,omitempty" xml:"type,omitempty"`
	Wifis                          []*GroupAddRequestWifis          `json:"wifis,omitempty" xml:"wifis,omitempty" type:"Repeated"`
	WorkdayClassList               []*int64                         `json:"workdayClassList,omitempty" xml:"workdayClassList,omitempty" type:"Repeated"`
	OpUserId                       *string                          `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GroupAddRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequest) GoString() string {
	return s.String()
}

func (s *GroupAddRequest) SetAdjustmentSettingId(v int64) *GroupAddRequest {
	s.AdjustmentSettingId = &v
	return s
}

func (s *GroupAddRequest) SetBleDeviceList(v []*GroupAddRequestBleDeviceList) *GroupAddRequest {
	s.BleDeviceList = v
	return s
}

func (s *GroupAddRequest) SetCheckNeedHealthyCode(v bool) *GroupAddRequest {
	s.CheckNeedHealthyCode = &v
	return s
}

func (s *GroupAddRequest) SetDefaultClassId(v int64) *GroupAddRequest {
	s.DefaultClassId = &v
	return s
}

func (s *GroupAddRequest) SetDisableCheckWhenRest(v bool) *GroupAddRequest {
	s.DisableCheckWhenRest = &v
	return s
}

func (s *GroupAddRequest) SetDisableCheckWithoutSchedule(v bool) *GroupAddRequest {
	s.DisableCheckWithoutSchedule = &v
	return s
}

func (s *GroupAddRequest) SetEnableCameraCheck(v bool) *GroupAddRequest {
	s.EnableCameraCheck = &v
	return s
}

func (s *GroupAddRequest) SetEnableEmpSelectClass(v bool) *GroupAddRequest {
	s.EnableEmpSelectClass = &v
	return s
}

func (s *GroupAddRequest) SetEnableFaceCheck(v bool) *GroupAddRequest {
	s.EnableFaceCheck = &v
	return s
}

func (s *GroupAddRequest) SetEnableFaceStrictMode(v bool) *GroupAddRequest {
	s.EnableFaceStrictMode = &v
	return s
}

func (s *GroupAddRequest) SetEnableNextDay(v bool) *GroupAddRequest {
	s.EnableNextDay = &v
	return s
}

func (s *GroupAddRequest) SetEnableOutSideUpdateNormalCheck(v bool) *GroupAddRequest {
	s.EnableOutSideUpdateNormalCheck = &v
	return s
}

func (s *GroupAddRequest) SetEnableOutsideApply(v bool) *GroupAddRequest {
	s.EnableOutsideApply = &v
	return s
}

func (s *GroupAddRequest) SetEnableOutsideCameraCheck(v bool) *GroupAddRequest {
	s.EnableOutsideCameraCheck = &v
	return s
}

func (s *GroupAddRequest) SetEnableOutsideCheck(v bool) *GroupAddRequest {
	s.EnableOutsideCheck = &v
	return s
}

func (s *GroupAddRequest) SetEnableOutsideRemark(v bool) *GroupAddRequest {
	s.EnableOutsideRemark = &v
	return s
}

func (s *GroupAddRequest) SetEnablePositionBle(v bool) *GroupAddRequest {
	s.EnablePositionBle = &v
	return s
}

func (s *GroupAddRequest) SetEnableTrimDistance(v bool) *GroupAddRequest {
	s.EnableTrimDistance = &v
	return s
}

func (s *GroupAddRequest) SetForbidHideOutSideAddress(v bool) *GroupAddRequest {
	s.ForbidHideOutSideAddress = &v
	return s
}

func (s *GroupAddRequest) SetFreeCheckSetting(v *GroupAddRequestFreeCheckSetting) *GroupAddRequest {
	s.FreeCheckSetting = v
	return s
}

func (s *GroupAddRequest) SetFreeCheckTypeId(v int32) *GroupAddRequest {
	s.FreeCheckTypeId = &v
	return s
}

func (s *GroupAddRequest) SetFreecheckDayStartMinOffset(v int32) *GroupAddRequest {
	s.FreecheckDayStartMinOffset = &v
	return s
}

func (s *GroupAddRequest) SetFreecheckWorkDays(v []*int64) *GroupAddRequest {
	s.FreecheckWorkDays = v
	return s
}

func (s *GroupAddRequest) SetGroupId(v int64) *GroupAddRequest {
	s.GroupId = &v
	return s
}

func (s *GroupAddRequest) SetGroupName(v string) *GroupAddRequest {
	s.GroupName = &v
	return s
}

func (s *GroupAddRequest) SetManagerList(v []*string) *GroupAddRequest {
	s.ManagerList = v
	return s
}

func (s *GroupAddRequest) SetMembers(v []*GroupAddRequestMembers) *GroupAddRequest {
	s.Members = v
	return s
}

func (s *GroupAddRequest) SetModifyMember(v bool) *GroupAddRequest {
	s.ModifyMember = &v
	return s
}

func (s *GroupAddRequest) SetOffset(v int32) *GroupAddRequest {
	s.Offset = &v
	return s
}

func (s *GroupAddRequest) SetOpenFaceCheck(v bool) *GroupAddRequest {
	s.OpenFaceCheck = &v
	return s
}

func (s *GroupAddRequest) SetOutsideCheckApproveModeId(v int32) *GroupAddRequest {
	s.OutsideCheckApproveModeId = &v
	return s
}

func (s *GroupAddRequest) SetOvertimeSettingId(v int64) *GroupAddRequest {
	s.OvertimeSettingId = &v
	return s
}

func (s *GroupAddRequest) SetOwner(v string) *GroupAddRequest {
	s.Owner = &v
	return s
}

func (s *GroupAddRequest) SetPositions(v []*GroupAddRequestPositions) *GroupAddRequest {
	s.Positions = v
	return s
}

func (s *GroupAddRequest) SetResourcePermissionMap(v map[string]interface{}) *GroupAddRequest {
	s.ResourcePermissionMap = v
	return s
}

func (s *GroupAddRequest) SetShiftVOList(v []*GroupAddRequestShiftVOList) *GroupAddRequest {
	s.ShiftVOList = v
	return s
}

func (s *GroupAddRequest) SetSkipHolidays(v bool) *GroupAddRequest {
	s.SkipHolidays = &v
	return s
}

func (s *GroupAddRequest) SetSpecialDays(v string) *GroupAddRequest {
	s.SpecialDays = &v
	return s
}

func (s *GroupAddRequest) SetTrimDistance(v int32) *GroupAddRequest {
	s.TrimDistance = &v
	return s
}

func (s *GroupAddRequest) SetType(v string) *GroupAddRequest {
	s.Type = &v
	return s
}

func (s *GroupAddRequest) SetWifis(v []*GroupAddRequestWifis) *GroupAddRequest {
	s.Wifis = v
	return s
}

func (s *GroupAddRequest) SetWorkdayClassList(v []*int64) *GroupAddRequest {
	s.WorkdayClassList = v
	return s
}

func (s *GroupAddRequest) SetOpUserId(v string) *GroupAddRequest {
	s.OpUserId = &v
	return s
}

type GroupAddRequestBleDeviceList struct {
	DeviceId *int64 `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
}

func (s GroupAddRequestBleDeviceList) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestBleDeviceList) GoString() string {
	return s.String()
}

func (s *GroupAddRequestBleDeviceList) SetDeviceId(v int64) *GroupAddRequestBleDeviceList {
	s.DeviceId = &v
	return s
}

type GroupAddRequestFreeCheckSetting struct {
	DelimitOffsetMinutesBetweenDays *int32                                       `json:"delimitOffsetMinutesBetweenDays,omitempty" xml:"delimitOffsetMinutesBetweenDays,omitempty"`
	FreeCheckGap                    *GroupAddRequestFreeCheckSettingFreeCheckGap `json:"freeCheckGap,omitempty" xml:"freeCheckGap,omitempty" type:"Struct"`
}

func (s GroupAddRequestFreeCheckSetting) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestFreeCheckSetting) GoString() string {
	return s.String()
}

func (s *GroupAddRequestFreeCheckSetting) SetDelimitOffsetMinutesBetweenDays(v int32) *GroupAddRequestFreeCheckSetting {
	s.DelimitOffsetMinutesBetweenDays = &v
	return s
}

func (s *GroupAddRequestFreeCheckSetting) SetFreeCheckGap(v *GroupAddRequestFreeCheckSettingFreeCheckGap) *GroupAddRequestFreeCheckSetting {
	s.FreeCheckGap = v
	return s
}

type GroupAddRequestFreeCheckSettingFreeCheckGap struct {
	OffOnCheckGapMinutes *int32 `json:"offOnCheckGapMinutes,omitempty" xml:"offOnCheckGapMinutes,omitempty"`
	OnOffCheckGapMinutes *int32 `json:"onOffCheckGapMinutes,omitempty" xml:"onOffCheckGapMinutes,omitempty"`
}

func (s GroupAddRequestFreeCheckSettingFreeCheckGap) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestFreeCheckSettingFreeCheckGap) GoString() string {
	return s.String()
}

func (s *GroupAddRequestFreeCheckSettingFreeCheckGap) SetOffOnCheckGapMinutes(v int32) *GroupAddRequestFreeCheckSettingFreeCheckGap {
	s.OffOnCheckGapMinutes = &v
	return s
}

func (s *GroupAddRequestFreeCheckSettingFreeCheckGap) SetOnOffCheckGapMinutes(v int32) *GroupAddRequestFreeCheckSettingFreeCheckGap {
	s.OnOffCheckGapMinutes = &v
	return s
}

type GroupAddRequestMembers struct {
	Role   *string `json:"role,omitempty" xml:"role,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GroupAddRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestMembers) GoString() string {
	return s.String()
}

func (s *GroupAddRequestMembers) SetRole(v string) *GroupAddRequestMembers {
	s.Role = &v
	return s
}

func (s *GroupAddRequestMembers) SetType(v string) *GroupAddRequestMembers {
	s.Type = &v
	return s
}

func (s *GroupAddRequestMembers) SetUserId(v string) *GroupAddRequestMembers {
	s.UserId = &v
	return s
}

type GroupAddRequestPositions struct {
	Address   *string `json:"address,omitempty" xml:"address,omitempty"`
	Latitude  *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Offset    *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GroupAddRequestPositions) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestPositions) GoString() string {
	return s.String()
}

func (s *GroupAddRequestPositions) SetAddress(v string) *GroupAddRequestPositions {
	s.Address = &v
	return s
}

func (s *GroupAddRequestPositions) SetLatitude(v string) *GroupAddRequestPositions {
	s.Latitude = &v
	return s
}

func (s *GroupAddRequestPositions) SetLongitude(v string) *GroupAddRequestPositions {
	s.Longitude = &v
	return s
}

func (s *GroupAddRequestPositions) SetOffset(v int32) *GroupAddRequestPositions {
	s.Offset = &v
	return s
}

func (s *GroupAddRequestPositions) SetTitle(v string) *GroupAddRequestPositions {
	s.Title = &v
	return s
}

type GroupAddRequestShiftVOList struct {
	ShiftId *int64 `json:"shiftId,omitempty" xml:"shiftId,omitempty"`
}

func (s GroupAddRequestShiftVOList) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestShiftVOList) GoString() string {
	return s.String()
}

func (s *GroupAddRequestShiftVOList) SetShiftId(v int64) *GroupAddRequestShiftVOList {
	s.ShiftId = &v
	return s
}

type GroupAddRequestWifis struct {
	MacAddr *string `json:"macAddr,omitempty" xml:"macAddr,omitempty"`
	Ssid    *string `json:"ssid,omitempty" xml:"ssid,omitempty"`
}

func (s GroupAddRequestWifis) String() string {
	return tea.Prettify(s)
}

func (s GroupAddRequestWifis) GoString() string {
	return s.String()
}

func (s *GroupAddRequestWifis) SetMacAddr(v string) *GroupAddRequestWifis {
	s.MacAddr = &v
	return s
}

func (s *GroupAddRequestWifis) SetSsid(v string) *GroupAddRequestWifis {
	s.Ssid = &v
	return s
}

type GroupAddResponseBody struct {
	Result  *GroupAddResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GroupAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupAddResponseBody) GoString() string {
	return s.String()
}

func (s *GroupAddResponseBody) SetResult(v *GroupAddResponseBodyResult) *GroupAddResponseBody {
	s.Result = v
	return s
}

func (s *GroupAddResponseBody) SetSuccess(v bool) *GroupAddResponseBody {
	s.Success = &v
	return s
}

type GroupAddResponseBodyResult struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GroupAddResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GroupAddResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GroupAddResponseBodyResult) SetId(v int64) *GroupAddResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GroupAddResponseBodyResult) SetName(v string) *GroupAddResponseBodyResult {
	s.Name = &v
	return s
}

type GroupAddResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupAddResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupAddResponse) GoString() string {
	return s.String()
}

func (s *GroupAddResponse) SetHeaders(v map[string]*string) *GroupAddResponse {
	s.Headers = v
	return s
}

func (s *GroupAddResponse) SetStatusCode(v int32) *GroupAddResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupAddResponse) SetBody(v *GroupAddResponseBody) *GroupAddResponse {
	s.Body = v
	return s
}

type GroupUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateHeaders) GoString() string {
	return s.String()
}

func (s *GroupUpdateHeaders) SetCommonHeaders(v map[string]*string) *GroupUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *GroupUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupUpdateRequest struct {
	AdjustmentSettingId            *int64                              `json:"adjustmentSettingId,omitempty" xml:"adjustmentSettingId,omitempty"`
	DisableCheckWhenRest           *bool                               `json:"disableCheckWhenRest,omitempty" xml:"disableCheckWhenRest,omitempty"`
	DisableCheckWithoutSchedule    *bool                               `json:"disableCheckWithoutSchedule,omitempty" xml:"disableCheckWithoutSchedule,omitempty"`
	EnableCameraCheck              *bool                               `json:"enableCameraCheck,omitempty" xml:"enableCameraCheck,omitempty"`
	EnableEmpSelectClass           *bool                               `json:"enableEmpSelectClass,omitempty" xml:"enableEmpSelectClass,omitempty"`
	EnableFaceCheck                *bool                               `json:"enableFaceCheck,omitempty" xml:"enableFaceCheck,omitempty"`
	EnableFaceStrictMode           *bool                               `json:"enableFaceStrictMode,omitempty" xml:"enableFaceStrictMode,omitempty"`
	EnableOutSideUpdateNormalCheck *bool                               `json:"enableOutSideUpdateNormalCheck,omitempty" xml:"enableOutSideUpdateNormalCheck,omitempty"`
	EnableOutsideApply             *bool                               `json:"enableOutsideApply,omitempty" xml:"enableOutsideApply,omitempty"`
	EnableOutsideCheck             *bool                               `json:"enableOutsideCheck,omitempty" xml:"enableOutsideCheck,omitempty"`
	EnableOutsideRemark            *bool                               `json:"enableOutsideRemark,omitempty" xml:"enableOutsideRemark,omitempty"`
	EnableTrimDistance             *bool                               `json:"enableTrimDistance,omitempty" xml:"enableTrimDistance,omitempty"`
	ForbidHideOutSideAddress       *bool                               `json:"forbidHideOutSideAddress,omitempty" xml:"forbidHideOutSideAddress,omitempty"`
	FreeCheckSetting               *GroupUpdateRequestFreeCheckSetting `json:"freeCheckSetting,omitempty" xml:"freeCheckSetting,omitempty" type:"Struct"`
	FreeCheckTypeId                *int32                              `json:"freeCheckTypeId,omitempty" xml:"freeCheckTypeId,omitempty"`
	GroupId                        *int64                              `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName                      *string                             `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ManagerList                    []*string                           `json:"managerList,omitempty" xml:"managerList,omitempty" type:"Repeated"`
	Offset                         *int32                              `json:"offset,omitempty" xml:"offset,omitempty"`
	OpenFaceCheck                  *bool                               `json:"openFaceCheck,omitempty" xml:"openFaceCheck,omitempty"`
	OutsideCheckApproveModeId      *int32                              `json:"outsideCheckApproveModeId,omitempty" xml:"outsideCheckApproveModeId,omitempty"`
	OvertimeSettingId              *int64                              `json:"overtimeSettingId,omitempty" xml:"overtimeSettingId,omitempty"`
	Owner                          *string                             `json:"owner,omitempty" xml:"owner,omitempty"`
	Positions                      []*GroupUpdateRequestPositions      `json:"positions,omitempty" xml:"positions,omitempty" type:"Repeated"`
	ResourcePermissionMap          map[string]interface{}              `json:"resourcePermissionMap,omitempty" xml:"resourcePermissionMap,omitempty"`
	ShiftVOList                    []*GroupUpdateRequestShiftVOList    `json:"shiftVOList,omitempty" xml:"shiftVOList,omitempty" type:"Repeated"`
	SkipHolidays                   *bool                               `json:"skipHolidays,omitempty" xml:"skipHolidays,omitempty"`
	TrimDistance                   *int32                              `json:"trimDistance,omitempty" xml:"trimDistance,omitempty"`
	WorkdayClassList               []*int64                            `json:"workdayClassList,omitempty" xml:"workdayClassList,omitempty" type:"Repeated"`
	OpUserId                       *string                             `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GroupUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateRequest) GoString() string {
	return s.String()
}

func (s *GroupUpdateRequest) SetAdjustmentSettingId(v int64) *GroupUpdateRequest {
	s.AdjustmentSettingId = &v
	return s
}

func (s *GroupUpdateRequest) SetDisableCheckWhenRest(v bool) *GroupUpdateRequest {
	s.DisableCheckWhenRest = &v
	return s
}

func (s *GroupUpdateRequest) SetDisableCheckWithoutSchedule(v bool) *GroupUpdateRequest {
	s.DisableCheckWithoutSchedule = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableCameraCheck(v bool) *GroupUpdateRequest {
	s.EnableCameraCheck = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableEmpSelectClass(v bool) *GroupUpdateRequest {
	s.EnableEmpSelectClass = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableFaceCheck(v bool) *GroupUpdateRequest {
	s.EnableFaceCheck = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableFaceStrictMode(v bool) *GroupUpdateRequest {
	s.EnableFaceStrictMode = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableOutSideUpdateNormalCheck(v bool) *GroupUpdateRequest {
	s.EnableOutSideUpdateNormalCheck = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableOutsideApply(v bool) *GroupUpdateRequest {
	s.EnableOutsideApply = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableOutsideCheck(v bool) *GroupUpdateRequest {
	s.EnableOutsideCheck = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableOutsideRemark(v bool) *GroupUpdateRequest {
	s.EnableOutsideRemark = &v
	return s
}

func (s *GroupUpdateRequest) SetEnableTrimDistance(v bool) *GroupUpdateRequest {
	s.EnableTrimDistance = &v
	return s
}

func (s *GroupUpdateRequest) SetForbidHideOutSideAddress(v bool) *GroupUpdateRequest {
	s.ForbidHideOutSideAddress = &v
	return s
}

func (s *GroupUpdateRequest) SetFreeCheckSetting(v *GroupUpdateRequestFreeCheckSetting) *GroupUpdateRequest {
	s.FreeCheckSetting = v
	return s
}

func (s *GroupUpdateRequest) SetFreeCheckTypeId(v int32) *GroupUpdateRequest {
	s.FreeCheckTypeId = &v
	return s
}

func (s *GroupUpdateRequest) SetGroupId(v int64) *GroupUpdateRequest {
	s.GroupId = &v
	return s
}

func (s *GroupUpdateRequest) SetGroupName(v string) *GroupUpdateRequest {
	s.GroupName = &v
	return s
}

func (s *GroupUpdateRequest) SetManagerList(v []*string) *GroupUpdateRequest {
	s.ManagerList = v
	return s
}

func (s *GroupUpdateRequest) SetOffset(v int32) *GroupUpdateRequest {
	s.Offset = &v
	return s
}

func (s *GroupUpdateRequest) SetOpenFaceCheck(v bool) *GroupUpdateRequest {
	s.OpenFaceCheck = &v
	return s
}

func (s *GroupUpdateRequest) SetOutsideCheckApproveModeId(v int32) *GroupUpdateRequest {
	s.OutsideCheckApproveModeId = &v
	return s
}

func (s *GroupUpdateRequest) SetOvertimeSettingId(v int64) *GroupUpdateRequest {
	s.OvertimeSettingId = &v
	return s
}

func (s *GroupUpdateRequest) SetOwner(v string) *GroupUpdateRequest {
	s.Owner = &v
	return s
}

func (s *GroupUpdateRequest) SetPositions(v []*GroupUpdateRequestPositions) *GroupUpdateRequest {
	s.Positions = v
	return s
}

func (s *GroupUpdateRequest) SetResourcePermissionMap(v map[string]interface{}) *GroupUpdateRequest {
	s.ResourcePermissionMap = v
	return s
}

func (s *GroupUpdateRequest) SetShiftVOList(v []*GroupUpdateRequestShiftVOList) *GroupUpdateRequest {
	s.ShiftVOList = v
	return s
}

func (s *GroupUpdateRequest) SetSkipHolidays(v bool) *GroupUpdateRequest {
	s.SkipHolidays = &v
	return s
}

func (s *GroupUpdateRequest) SetTrimDistance(v int32) *GroupUpdateRequest {
	s.TrimDistance = &v
	return s
}

func (s *GroupUpdateRequest) SetWorkdayClassList(v []*int64) *GroupUpdateRequest {
	s.WorkdayClassList = v
	return s
}

func (s *GroupUpdateRequest) SetOpUserId(v string) *GroupUpdateRequest {
	s.OpUserId = &v
	return s
}

type GroupUpdateRequestFreeCheckSetting struct {
	DelimitOffsetMinutesBetweenDays *int32                                          `json:"delimitOffsetMinutesBetweenDays,omitempty" xml:"delimitOffsetMinutesBetweenDays,omitempty"`
	FreeCheckGap                    *GroupUpdateRequestFreeCheckSettingFreeCheckGap `json:"freeCheckGap,omitempty" xml:"freeCheckGap,omitempty" type:"Struct"`
}

func (s GroupUpdateRequestFreeCheckSetting) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateRequestFreeCheckSetting) GoString() string {
	return s.String()
}

func (s *GroupUpdateRequestFreeCheckSetting) SetDelimitOffsetMinutesBetweenDays(v int32) *GroupUpdateRequestFreeCheckSetting {
	s.DelimitOffsetMinutesBetweenDays = &v
	return s
}

func (s *GroupUpdateRequestFreeCheckSetting) SetFreeCheckGap(v *GroupUpdateRequestFreeCheckSettingFreeCheckGap) *GroupUpdateRequestFreeCheckSetting {
	s.FreeCheckGap = v
	return s
}

type GroupUpdateRequestFreeCheckSettingFreeCheckGap struct {
	OffOnCheckGapMinutes *int32 `json:"offOnCheckGapMinutes,omitempty" xml:"offOnCheckGapMinutes,omitempty"`
	OnOffCheckGapMinutes *int32 `json:"onOffCheckGapMinutes,omitempty" xml:"onOffCheckGapMinutes,omitempty"`
}

func (s GroupUpdateRequestFreeCheckSettingFreeCheckGap) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateRequestFreeCheckSettingFreeCheckGap) GoString() string {
	return s.String()
}

func (s *GroupUpdateRequestFreeCheckSettingFreeCheckGap) SetOffOnCheckGapMinutes(v int32) *GroupUpdateRequestFreeCheckSettingFreeCheckGap {
	s.OffOnCheckGapMinutes = &v
	return s
}

func (s *GroupUpdateRequestFreeCheckSettingFreeCheckGap) SetOnOffCheckGapMinutes(v int32) *GroupUpdateRequestFreeCheckSettingFreeCheckGap {
	s.OnOffCheckGapMinutes = &v
	return s
}

type GroupUpdateRequestPositions struct {
	Address   *string `json:"address,omitempty" xml:"address,omitempty"`
	Latitude  *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Offset    *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GroupUpdateRequestPositions) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateRequestPositions) GoString() string {
	return s.String()
}

func (s *GroupUpdateRequestPositions) SetAddress(v string) *GroupUpdateRequestPositions {
	s.Address = &v
	return s
}

func (s *GroupUpdateRequestPositions) SetLatitude(v string) *GroupUpdateRequestPositions {
	s.Latitude = &v
	return s
}

func (s *GroupUpdateRequestPositions) SetLongitude(v string) *GroupUpdateRequestPositions {
	s.Longitude = &v
	return s
}

func (s *GroupUpdateRequestPositions) SetOffset(v int32) *GroupUpdateRequestPositions {
	s.Offset = &v
	return s
}

func (s *GroupUpdateRequestPositions) SetTitle(v string) *GroupUpdateRequestPositions {
	s.Title = &v
	return s
}

type GroupUpdateRequestShiftVOList struct {
	ShiftId *int64 `json:"shiftId,omitempty" xml:"shiftId,omitempty"`
}

func (s GroupUpdateRequestShiftVOList) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateRequestShiftVOList) GoString() string {
	return s.String()
}

func (s *GroupUpdateRequestShiftVOList) SetShiftId(v int64) *GroupUpdateRequestShiftVOList {
	s.ShiftId = &v
	return s
}

type GroupUpdateResponseBody struct {
	Result  *GroupUpdateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GroupUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *GroupUpdateResponseBody) SetResult(v *GroupUpdateResponseBodyResult) *GroupUpdateResponseBody {
	s.Result = v
	return s
}

func (s *GroupUpdateResponseBody) SetSuccess(v bool) *GroupUpdateResponseBody {
	s.Success = &v
	return s
}

type GroupUpdateResponseBodyResult struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GroupUpdateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GroupUpdateResponseBodyResult) SetId(v int64) *GroupUpdateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GroupUpdateResponseBodyResult) SetName(v string) *GroupUpdateResponseBodyResult {
	s.Name = &v
	return s
}

type GroupUpdateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateResponse) GoString() string {
	return s.String()
}

func (s *GroupUpdateResponse) SetHeaders(v map[string]*string) *GroupUpdateResponse {
	s.Headers = v
	return s
}

func (s *GroupUpdateResponse) SetStatusCode(v int32) *GroupUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupUpdateResponse) SetBody(v *GroupUpdateResponseBody) *GroupUpdateResponse {
	s.Body = v
	return s
}

type InitAndGetLeaveALlocationQuotasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitAndGetLeaveALlocationQuotasHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitAndGetLeaveALlocationQuotasHeaders) GoString() string {
	return s.String()
}

func (s *InitAndGetLeaveALlocationQuotasHeaders) SetCommonHeaders(v map[string]*string) *InitAndGetLeaveALlocationQuotasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasHeaders) SetXAcsDingtalkAccessToken(v string) *InitAndGetLeaveALlocationQuotasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitAndGetLeaveALlocationQuotasRequest struct {
	LeaveCode *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	OpUserId  *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s InitAndGetLeaveALlocationQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s InitAndGetLeaveALlocationQuotasRequest) GoString() string {
	return s.String()
}

func (s *InitAndGetLeaveALlocationQuotasRequest) SetLeaveCode(v string) *InitAndGetLeaveALlocationQuotasRequest {
	s.LeaveCode = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasRequest) SetOpUserId(v string) *InitAndGetLeaveALlocationQuotasRequest {
	s.OpUserId = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasRequest) SetUserId(v string) *InitAndGetLeaveALlocationQuotasRequest {
	s.UserId = &v
	return s
}

type InitAndGetLeaveALlocationQuotasResponseBody struct {
	Result []*InitAndGetLeaveALlocationQuotasResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s InitAndGetLeaveALlocationQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitAndGetLeaveALlocationQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *InitAndGetLeaveALlocationQuotasResponseBody) SetResult(v []*InitAndGetLeaveALlocationQuotasResponseBodyResult) *InitAndGetLeaveALlocationQuotasResponseBody {
	s.Result = v
	return s
}

type InitAndGetLeaveALlocationQuotasResponseBodyResult struct {
	EndTime         *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	LeaveCode       *string `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	QuotaCycle      *string `json:"quotaCycle,omitempty" xml:"quotaCycle,omitempty"`
	QuotaId         *string `json:"quotaId,omitempty" xml:"quotaId,omitempty"`
	QuotaNumPerDay  *int64  `json:"quotaNumPerDay,omitempty" xml:"quotaNumPerDay,omitempty"`
	QuotaNumPerHour *int64  `json:"quotaNumPerHour,omitempty" xml:"quotaNumPerHour,omitempty"`
	StartTime       *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UsedNumPerDay   *int64  `json:"usedNumPerDay,omitempty" xml:"usedNumPerDay,omitempty"`
	UsedNumPerHour  *int64  `json:"usedNumPerHour,omitempty" xml:"usedNumPerHour,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s InitAndGetLeaveALlocationQuotasResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InitAndGetLeaveALlocationQuotasResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetEndTime(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetLeaveCode(v string) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.LeaveCode = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetQuotaCycle(v string) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.QuotaCycle = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetQuotaId(v string) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.QuotaId = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetQuotaNumPerDay(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.QuotaNumPerDay = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetQuotaNumPerHour(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.QuotaNumPerHour = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetStartTime(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetUsedNumPerDay(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.UsedNumPerDay = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetUsedNumPerHour(v int64) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.UsedNumPerHour = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponseBodyResult) SetUserId(v string) *InitAndGetLeaveALlocationQuotasResponseBodyResult {
	s.UserId = &v
	return s
}

type InitAndGetLeaveALlocationQuotasResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitAndGetLeaveALlocationQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitAndGetLeaveALlocationQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s InitAndGetLeaveALlocationQuotasResponse) GoString() string {
	return s.String()
}

func (s *InitAndGetLeaveALlocationQuotasResponse) SetHeaders(v map[string]*string) *InitAndGetLeaveALlocationQuotasResponse {
	s.Headers = v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponse) SetStatusCode(v int32) *InitAndGetLeaveALlocationQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *InitAndGetLeaveALlocationQuotasResponse) SetBody(v *InitAndGetLeaveALlocationQuotasResponseBody) *InitAndGetLeaveALlocationQuotasResponse {
	s.Body = v
	return s
}

type ListApproveByUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListApproveByUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApproveByUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListApproveByUsersHeaders) SetCommonHeaders(v map[string]*string) *ListApproveByUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApproveByUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListApproveByUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListApproveByUsersRequest struct {
	BizTypes     []*int32 `json:"bizTypes,omitempty" xml:"bizTypes,omitempty" type:"Repeated"`
	FromDateTime *int64   `json:"fromDateTime,omitempty" xml:"fromDateTime,omitempty"`
	ToDateTime   *int64   `json:"toDateTime,omitempty" xml:"toDateTime,omitempty"`
	UserIds      *string  `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s ListApproveByUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApproveByUsersRequest) GoString() string {
	return s.String()
}

func (s *ListApproveByUsersRequest) SetBizTypes(v []*int32) *ListApproveByUsersRequest {
	s.BizTypes = v
	return s
}

func (s *ListApproveByUsersRequest) SetFromDateTime(v int64) *ListApproveByUsersRequest {
	s.FromDateTime = &v
	return s
}

func (s *ListApproveByUsersRequest) SetToDateTime(v int64) *ListApproveByUsersRequest {
	s.ToDateTime = &v
	return s
}

func (s *ListApproveByUsersRequest) SetUserIds(v string) *ListApproveByUsersRequest {
	s.UserIds = &v
	return s
}

type ListApproveByUsersResponseBody struct {
	Result []*ListApproveByUsersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListApproveByUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApproveByUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListApproveByUsersResponseBody) SetResult(v []*ListApproveByUsersResponseBodyResult) *ListApproveByUsersResponseBody {
	s.Result = v
	return s
}

type ListApproveByUsersResponseBodyResult struct {
	ApproveId      *string `json:"approveId,omitempty" xml:"approveId,omitempty"`
	BeginTime      *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	BizType        *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CalculateModel *int32  `json:"calculateModel,omitempty" xml:"calculateModel,omitempty"`
	DurationUnit   *string `json:"durationUnit,omitempty" xml:"durationUnit,omitempty"`
	EndTime        *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	SubType        *string `json:"subType,omitempty" xml:"subType,omitempty"`
	TagName        *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListApproveByUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListApproveByUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApproveByUsersResponseBodyResult) SetApproveId(v string) *ListApproveByUsersResponseBodyResult {
	s.ApproveId = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetBeginTime(v string) *ListApproveByUsersResponseBodyResult {
	s.BeginTime = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetBizType(v int32) *ListApproveByUsersResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetCalculateModel(v int32) *ListApproveByUsersResponseBodyResult {
	s.CalculateModel = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetDurationUnit(v string) *ListApproveByUsersResponseBodyResult {
	s.DurationUnit = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetEndTime(v string) *ListApproveByUsersResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetSubType(v string) *ListApproveByUsersResponseBodyResult {
	s.SubType = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetTagName(v string) *ListApproveByUsersResponseBodyResult {
	s.TagName = &v
	return s
}

func (s *ListApproveByUsersResponseBodyResult) SetUserId(v string) *ListApproveByUsersResponseBodyResult {
	s.UserId = &v
	return s
}

type ListApproveByUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApproveByUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApproveByUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApproveByUsersResponse) GoString() string {
	return s.String()
}

func (s *ListApproveByUsersResponse) SetHeaders(v map[string]*string) *ListApproveByUsersResponse {
	s.Headers = v
	return s
}

func (s *ListApproveByUsersResponse) SetStatusCode(v int32) *ListApproveByUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApproveByUsersResponse) SetBody(v *ListApproveByUsersResponseBody) *ListApproveByUsersResponse {
	s.Body = v
	return s
}

type ModifyWaterMarkTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ModifyWaterMarkTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyWaterMarkTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ModifyWaterMarkTemplateHeaders) SetCommonHeaders(v map[string]*string) *ModifyWaterMarkTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ModifyWaterMarkTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *ModifyWaterMarkTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ModifyWaterMarkTemplateRequest struct {
	FormCode           *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	LayoutDesignId     *string `json:"layoutDesignId,omitempty" xml:"layoutDesignId,omitempty"`
	SchemaContent      *string `json:"schemaContent,omitempty" xml:"schemaContent,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	WaterMarkId        *string `json:"waterMarkId,omitempty" xml:"waterMarkId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ModifyWaterMarkTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyWaterMarkTemplateRequest) SetFormCode(v string) *ModifyWaterMarkTemplateRequest {
	s.FormCode = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetIcon(v string) *ModifyWaterMarkTemplateRequest {
	s.Icon = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetLayoutDesignId(v string) *ModifyWaterMarkTemplateRequest {
	s.LayoutDesignId = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetSchemaContent(v string) *ModifyWaterMarkTemplateRequest {
	s.SchemaContent = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetTitle(v string) *ModifyWaterMarkTemplateRequest {
	s.Title = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetWaterMarkId(v string) *ModifyWaterMarkTemplateRequest {
	s.WaterMarkId = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetOpenConversationId(v string) *ModifyWaterMarkTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ModifyWaterMarkTemplateRequest) SetUserId(v string) *ModifyWaterMarkTemplateRequest {
	s.UserId = &v
	return s
}

type ModifyWaterMarkTemplateResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyWaterMarkTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWaterMarkTemplateResponseBody) SetResult(v string) *ModifyWaterMarkTemplateResponseBody {
	s.Result = &v
	return s
}

type ModifyWaterMarkTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWaterMarkTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *ModifyWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyWaterMarkTemplateResponse) SetStatusCode(v int32) *ModifyWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWaterMarkTemplateResponse) SetBody(v *ModifyWaterMarkTemplateResponseBody) *ModifyWaterMarkTemplateResponse {
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
	ApproveId  *string                                `json:"approveId,omitempty" xml:"approveId,omitempty"`
	OpUserId   *string                                `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	PunchParam *ProcessApproveCreateRequestPunchParam `json:"punchParam,omitempty" xml:"punchParam,omitempty" type:"Struct"`
	SubType    *string                                `json:"subType,omitempty" xml:"subType,omitempty"`
	TagName    *string                                `json:"tagName,omitempty" xml:"tagName,omitempty"`
	UserId     *string                                `json:"userId,omitempty" xml:"userId,omitempty"`
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
	PositionId   *string `json:"positionId,omitempty" xml:"positionId,omitempty"`
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	PositionType *string `json:"positionType,omitempty" xml:"positionType,omitempty"`
	PunchTime    *int64  `json:"punchTime,omitempty" xml:"punchTime,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProcessApproveCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ProcessApproveCreateResponse) SetStatusCode(v int32) *ProcessApproveCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessApproveCreateResponse) SetBody(v *ProcessApproveCreateResponseBody) *ProcessApproveCreateResponse {
	s.Body = v
	return s
}

type RetainLeaveTypesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RetainLeaveTypesHeaders) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesHeaders) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesHeaders) SetCommonHeaders(v map[string]*string) *RetainLeaveTypesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RetainLeaveTypesHeaders) SetXAcsDingtalkAccessToken(v string) *RetainLeaveTypesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RetainLeaveTypesRequest struct {
	LeaveCodes []*string `json:"leaveCodes,omitempty" xml:"leaveCodes,omitempty" type:"Repeated"`
	OpUserId   *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	Source     *int64    `json:"source,omitempty" xml:"source,omitempty"`
}

func (s RetainLeaveTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesRequest) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesRequest) SetLeaveCodes(v []*string) *RetainLeaveTypesRequest {
	s.LeaveCodes = v
	return s
}

func (s *RetainLeaveTypesRequest) SetOpUserId(v string) *RetainLeaveTypesRequest {
	s.OpUserId = &v
	return s
}

func (s *RetainLeaveTypesRequest) SetSource(v int64) *RetainLeaveTypesRequest {
	s.Source = &v
	return s
}

type RetainLeaveTypesResponseBody struct {
	Result []*RetainLeaveTypesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RetainLeaveTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponseBody) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponseBody) SetResult(v []*RetainLeaveTypesResponseBodyResult) *RetainLeaveTypesResponseBody {
	s.Result = v
	return s
}

type RetainLeaveTypesResponseBodyResult struct {
	BizType              *string                                              `json:"bizType,omitempty" xml:"bizType,omitempty"`
	HoursInPerDay        *int64                                               `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate     *RetainLeaveTypesResponseBodyResultLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveCode            *string                                              `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveHourCeil        *string                                              `json:"leaveHourCeil,omitempty" xml:"leaveHourCeil,omitempty"`
	LeaveName            *string                                              `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveTimeCeil        *bool                                                `json:"leaveTimeCeil,omitempty" xml:"leaveTimeCeil,omitempty"`
	LeaveTimeCeilMinUnit *string                                              `json:"leaveTimeCeilMinUnit,omitempty" xml:"leaveTimeCeilMinUnit,omitempty"`
	LeaveViewUnit        *string                                              `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	LieuDelayNum         *int64                                               `json:"lieuDelayNum,omitempty" xml:"lieuDelayNum,omitempty"`
	LieuDelayUnit        *string                                              `json:"lieuDelayUnit,omitempty" xml:"lieuDelayUnit,omitempty"`
	MaxLeaveTime         *int64                                               `json:"maxLeaveTime,omitempty" xml:"maxLeaveTime,omitempty"`
	MinLeaveHour         *float64                                             `json:"minLeaveHour,omitempty" xml:"minLeaveHour,omitempty"`
	NaturalDayLeave      *bool                                                `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	PaidLeave            *bool                                                `json:"paidLeave,omitempty" xml:"paidLeave,omitempty"`
	SubmitTimeRule       *RetainLeaveTypesResponseBodyResultSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	VisibilityRules      []*RetainLeaveTypesResponseBodyResultVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
	WhenCanLeave         *string                                              `json:"whenCanLeave,omitempty" xml:"whenCanLeave,omitempty"`
}

func (s RetainLeaveTypesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponseBodyResult) SetBizType(v string) *RetainLeaveTypesResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetHoursInPerDay(v int64) *RetainLeaveTypesResponseBodyResult {
	s.HoursInPerDay = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveCertificate(v *RetainLeaveTypesResponseBodyResultLeaveCertificate) *RetainLeaveTypesResponseBodyResult {
	s.LeaveCertificate = v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveCode(v string) *RetainLeaveTypesResponseBodyResult {
	s.LeaveCode = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveHourCeil(v string) *RetainLeaveTypesResponseBodyResult {
	s.LeaveHourCeil = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveName(v string) *RetainLeaveTypesResponseBodyResult {
	s.LeaveName = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveTimeCeil(v bool) *RetainLeaveTypesResponseBodyResult {
	s.LeaveTimeCeil = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveTimeCeilMinUnit(v string) *RetainLeaveTypesResponseBodyResult {
	s.LeaveTimeCeilMinUnit = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLeaveViewUnit(v string) *RetainLeaveTypesResponseBodyResult {
	s.LeaveViewUnit = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLieuDelayNum(v int64) *RetainLeaveTypesResponseBodyResult {
	s.LieuDelayNum = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetLieuDelayUnit(v string) *RetainLeaveTypesResponseBodyResult {
	s.LieuDelayUnit = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetMaxLeaveTime(v int64) *RetainLeaveTypesResponseBodyResult {
	s.MaxLeaveTime = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetMinLeaveHour(v float64) *RetainLeaveTypesResponseBodyResult {
	s.MinLeaveHour = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetNaturalDayLeave(v bool) *RetainLeaveTypesResponseBodyResult {
	s.NaturalDayLeave = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetPaidLeave(v bool) *RetainLeaveTypesResponseBodyResult {
	s.PaidLeave = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetSubmitTimeRule(v *RetainLeaveTypesResponseBodyResultSubmitTimeRule) *RetainLeaveTypesResponseBodyResult {
	s.SubmitTimeRule = v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetVisibilityRules(v []*RetainLeaveTypesResponseBodyResultVisibilityRules) *RetainLeaveTypesResponseBodyResult {
	s.VisibilityRules = v
	return s
}

func (s *RetainLeaveTypesResponseBodyResult) SetWhenCanLeave(v string) *RetainLeaveTypesResponseBodyResult {
	s.WhenCanLeave = &v
	return s
}

type RetainLeaveTypesResponseBodyResultLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RetainLeaveTypesResponseBodyResultLeaveCertificate) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponseBodyResultLeaveCertificate) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponseBodyResultLeaveCertificate) SetDuration(v float64) *RetainLeaveTypesResponseBodyResultLeaveCertificate {
	s.Duration = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultLeaveCertificate) SetEnable(v bool) *RetainLeaveTypesResponseBodyResultLeaveCertificate {
	s.Enable = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultLeaveCertificate) SetPromptInformation(v string) *RetainLeaveTypesResponseBodyResultLeaveCertificate {
	s.PromptInformation = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultLeaveCertificate) SetUnit(v string) *RetainLeaveTypesResponseBodyResultLeaveCertificate {
	s.Unit = &v
	return s
}

type RetainLeaveTypesResponseBodyResultSubmitTimeRule struct {
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
}

func (s RetainLeaveTypesResponseBodyResultSubmitTimeRule) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponseBodyResultSubmitTimeRule) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponseBodyResultSubmitTimeRule) SetEnableTimeLimit(v bool) *RetainLeaveTypesResponseBodyResultSubmitTimeRule {
	s.EnableTimeLimit = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultSubmitTimeRule) SetTimeType(v string) *RetainLeaveTypesResponseBodyResultSubmitTimeRule {
	s.TimeType = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultSubmitTimeRule) SetTimeUnit(v string) *RetainLeaveTypesResponseBodyResultSubmitTimeRule {
	s.TimeUnit = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultSubmitTimeRule) SetTimeValue(v int64) *RetainLeaveTypesResponseBodyResultSubmitTimeRule {
	s.TimeValue = &v
	return s
}

type RetainLeaveTypesResponseBodyResultVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s RetainLeaveTypesResponseBodyResultVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponseBodyResultVisibilityRules) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponseBodyResultVisibilityRules) SetType(v string) *RetainLeaveTypesResponseBodyResultVisibilityRules {
	s.Type = &v
	return s
}

func (s *RetainLeaveTypesResponseBodyResultVisibilityRules) SetVisible(v []*string) *RetainLeaveTypesResponseBodyResultVisibilityRules {
	s.Visible = v
	return s
}

type RetainLeaveTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RetainLeaveTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetainLeaveTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s RetainLeaveTypesResponse) GoString() string {
	return s.String()
}

func (s *RetainLeaveTypesResponse) SetHeaders(v map[string]*string) *RetainLeaveTypesResponse {
	s.Headers = v
	return s
}

func (s *RetainLeaveTypesResponse) SetStatusCode(v int32) *RetainLeaveTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *RetainLeaveTypesResponse) SetBody(v *RetainLeaveTypesResponseBody) *RetainLeaveTypesResponse {
	s.Body = v
	return s
}

type SaveCustomWaterMarkTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveCustomWaterMarkTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomWaterMarkTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SaveCustomWaterMarkTemplateHeaders) SetCommonHeaders(v map[string]*string) *SaveCustomWaterMarkTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveCustomWaterMarkTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SaveCustomWaterMarkTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveCustomWaterMarkTemplateRequest struct {
	BizCode            *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	LayoutDesignId     *string `json:"layoutDesignId,omitempty" xml:"layoutDesignId,omitempty"`
	SceneCode          *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
	SchemaContent      *string `json:"schemaContent,omitempty" xml:"schemaContent,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SaveCustomWaterMarkTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *SaveCustomWaterMarkTemplateRequest) SetBizCode(v string) *SaveCustomWaterMarkTemplateRequest {
	s.BizCode = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetIcon(v string) *SaveCustomWaterMarkTemplateRequest {
	s.Icon = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetLayoutDesignId(v string) *SaveCustomWaterMarkTemplateRequest {
	s.LayoutDesignId = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetSceneCode(v string) *SaveCustomWaterMarkTemplateRequest {
	s.SceneCode = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetSchemaContent(v string) *SaveCustomWaterMarkTemplateRequest {
	s.SchemaContent = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetTitle(v string) *SaveCustomWaterMarkTemplateRequest {
	s.Title = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetOpenConversationId(v string) *SaveCustomWaterMarkTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateRequest) SetUserId(v string) *SaveCustomWaterMarkTemplateRequest {
	s.UserId = &v
	return s
}

type SaveCustomWaterMarkTemplateResponseBody struct {
	Result *SaveCustomWaterMarkTemplateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SaveCustomWaterMarkTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCustomWaterMarkTemplateResponseBody) SetResult(v *SaveCustomWaterMarkTemplateResponseBodyResult) *SaveCustomWaterMarkTemplateResponseBody {
	s.Result = v
	return s
}

type SaveCustomWaterMarkTemplateResponseBodyResult struct {
	FormCode    *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	WaterMarkId *string `json:"waterMarkId,omitempty" xml:"waterMarkId,omitempty"`
}

func (s SaveCustomWaterMarkTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomWaterMarkTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveCustomWaterMarkTemplateResponseBodyResult) SetFormCode(v string) *SaveCustomWaterMarkTemplateResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateResponseBodyResult) SetWaterMarkId(v string) *SaveCustomWaterMarkTemplateResponseBodyResult {
	s.WaterMarkId = &v
	return s
}

type SaveCustomWaterMarkTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveCustomWaterMarkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveCustomWaterMarkTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomWaterMarkTemplateResponse) GoString() string {
	return s.String()
}

func (s *SaveCustomWaterMarkTemplateResponse) SetHeaders(v map[string]*string) *SaveCustomWaterMarkTemplateResponse {
	s.Headers = v
	return s
}

func (s *SaveCustomWaterMarkTemplateResponse) SetStatusCode(v int32) *SaveCustomWaterMarkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCustomWaterMarkTemplateResponse) SetBody(v *SaveCustomWaterMarkTemplateResponseBody) *SaveCustomWaterMarkTemplateResponse {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

func (s *SyncScheduleInfoResponse) SetStatusCode(v int32) *SyncScheduleInfoResponse {
	s.StatusCode = &v
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
	BizType          *string                                  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Extras           *string                                  `json:"extras,omitempty" xml:"extras,omitempty"`
	HoursInPerDay    *int64                                   `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate *UpdateLeaveTypeRequestLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveCode        *string                                  `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveName        *string                                  `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveViewUnit    *string                                  `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	NaturalDayLeave  *bool                                    `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	SubmitTimeRule   *UpdateLeaveTypeRequestSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	VisibilityRules  []*UpdateLeaveTypeRequestVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
	OpUserId         *string                                  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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

func (s *UpdateLeaveTypeRequest) SetVisibilityRules(v []*UpdateLeaveTypeRequestVisibilityRules) *UpdateLeaveTypeRequest {
	s.VisibilityRules = v
	return s
}

func (s *UpdateLeaveTypeRequest) SetOpUserId(v string) *UpdateLeaveTypeRequest {
	s.OpUserId = &v
	return s
}

type UpdateLeaveTypeRequestLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
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

type UpdateLeaveTypeRequestVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s UpdateLeaveTypeRequestVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeRequestVisibilityRules) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeRequestVisibilityRules) SetType(v string) *UpdateLeaveTypeRequestVisibilityRules {
	s.Type = &v
	return s
}

func (s *UpdateLeaveTypeRequestVisibilityRules) SetVisible(v []*string) *UpdateLeaveTypeRequestVisibilityRules {
	s.Visible = v
	return s
}

type UpdateLeaveTypeResponseBody struct {
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
	BizType          *string                                             `json:"bizType,omitempty" xml:"bizType,omitempty"`
	HoursInPerDay    *int64                                              `json:"hoursInPerDay,omitempty" xml:"hoursInPerDay,omitempty"`
	LeaveCertificate *UpdateLeaveTypeResponseBodyResultLeaveCertificate  `json:"leaveCertificate,omitempty" xml:"leaveCertificate,omitempty" type:"Struct"`
	LeaveCode        *string                                             `json:"leaveCode,omitempty" xml:"leaveCode,omitempty"`
	LeaveName        *string                                             `json:"leaveName,omitempty" xml:"leaveName,omitempty"`
	LeaveViewUnit    *string                                             `json:"leaveViewUnit,omitempty" xml:"leaveViewUnit,omitempty"`
	NaturalDayLeave  *bool                                               `json:"naturalDayLeave,omitempty" xml:"naturalDayLeave,omitempty"`
	SubmitTimeRule   *UpdateLeaveTypeResponseBodyResultSubmitTimeRule    `json:"submitTimeRule,omitempty" xml:"submitTimeRule,omitempty" type:"Struct"`
	VisibilityRules  []*UpdateLeaveTypeResponseBodyResultVisibilityRules `json:"visibilityRules,omitempty" xml:"visibilityRules,omitempty" type:"Repeated"`
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

func (s *UpdateLeaveTypeResponseBodyResult) SetVisibilityRules(v []*UpdateLeaveTypeResponseBodyResultVisibilityRules) *UpdateLeaveTypeResponseBodyResult {
	s.VisibilityRules = v
	return s
}

type UpdateLeaveTypeResponseBodyResultLeaveCertificate struct {
	Duration          *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Enable            *bool    `json:"enable,omitempty" xml:"enable,omitempty"`
	PromptInformation *string  `json:"promptInformation,omitempty" xml:"promptInformation,omitempty"`
	Unit              *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	EnableTimeLimit *bool   `json:"enableTimeLimit,omitempty" xml:"enableTimeLimit,omitempty"`
	TimeType        *string `json:"timeType,omitempty" xml:"timeType,omitempty"`
	TimeUnit        *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
	TimeValue       *int64  `json:"timeValue,omitempty" xml:"timeValue,omitempty"`
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

type UpdateLeaveTypeResponseBodyResultVisibilityRules struct {
	Type    *string   `json:"type,omitempty" xml:"type,omitempty"`
	Visible []*string `json:"visible,omitempty" xml:"visible,omitempty" type:"Repeated"`
}

func (s UpdateLeaveTypeResponseBodyResultVisibilityRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateLeaveTypeResponseBodyResultVisibilityRules) GoString() string {
	return s.String()
}

func (s *UpdateLeaveTypeResponseBodyResultVisibilityRules) SetType(v string) *UpdateLeaveTypeResponseBodyResultVisibilityRules {
	s.Type = &v
	return s
}

func (s *UpdateLeaveTypeResponseBodyResultVisibilityRules) SetVisible(v []*string) *UpdateLeaveTypeResponseBodyResultVisibilityRules {
	s.Visible = v
	return s
}

type UpdateLeaveTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLeaveTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateLeaveTypeResponse) SetStatusCode(v int32) *UpdateLeaveTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLeaveTypeResponse) SetBody(v *UpdateLeaveTypeResponseBody) *UpdateLeaveTypeResponse {
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

	if !tea.BoolValue(util.IsUnset(request.LeaveCertificate)) {
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

	if !tea.BoolValue(util.IsUnset(request.SubmitTimeRule)) {
		body["submitTimeRule"] = request.SubmitTimeRule
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityRules)) {
		body["visibilityRules"] = request.VisibilityRules
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
	params := &openapi.Params{
		Action:      tea.String("AddLeaveType"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/leaves/types"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLeaveTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	params := &openapi.Params{
		Action:      tea.String("AttendanceBleDevicesAdd"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/group/bledevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AttendanceBleDevicesAddResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("AttendanceBleDevicesQuery"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/group/bledevices/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AttendanceBleDevicesQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("AttendanceBleDevicesRemove"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/group/bledevices/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AttendanceBleDevicesRemoveResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) BatchBossCheckWithOptions(request *BatchBossCheckRequest, headers *BatchBossCheckHeaders, runtime *util.RuntimeOptions) (_result *BatchBossCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Models)) {
		body["models"] = request.Models
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
	params := &openapi.Params{
		Action:      tea.String("BatchBossCheck"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/results/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchBossCheckResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchBossCheck(request *BatchBossCheckRequest) (_result *BatchBossCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchBossCheckHeaders{}
	_result = &BatchBossCheckResponse{}
	_body, _err := client.BatchBossCheckWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CheckClosingAccount"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/closingAccounts/status/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckClosingAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CheckWritePermission"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/writePermissions/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckWritePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.PunchParam)) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateApprove"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/approves"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApproveResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) DeleteWaterMarkTemplateWithOptions(request *DeleteWaterMarkTemplateRequest, headers *DeleteWaterMarkTemplateHeaders, runtime *util.RuntimeOptions) (_result *DeleteWaterMarkTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.FormContent)) {
		query["formContent"] = request.FormContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemTemplate)) {
		query["systemTemplate"] = request.SystemTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("DeleteWaterMarkTemplate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/watermarks/templates"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWaterMarkTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWaterMarkTemplate(request *DeleteWaterMarkTemplateRequest) (_result *DeleteWaterMarkTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWaterMarkTemplateHeaders{}
	_result = &DeleteWaterMarkTemplateResponse{}
	_body, _err := client.DeleteWaterMarkTemplateWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DingTalkSecurityCheck"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/securities/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DingTalkSecurityCheckResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetATManageScopeWithOptions(request *GetATManageScopeRequest, headers *GetATManageScopeHeaders, runtime *util.RuntimeOptions) (_result *GetATManageScopeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetATManageScope"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/manageScopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetATManageScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetATManageScope(request *GetATManageScopeRequest) (_result *GetATManageScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetATManageScopeHeaders{}
	_result = &GetATManageScopeResponse{}
	_body, _err := client.GetATManageScopeWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetAdjustments"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/adjustments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdjustmentsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetCheckInSchemaTemplateWithOptions(request *GetCheckInSchemaTemplateRequest, headers *GetCheckInSchemaTemplateHeaders, runtime *util.RuntimeOptions) (_result *GetCheckInSchemaTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["sceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetCheckInSchemaTemplate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/watermarks/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCheckInSchemaTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCheckInSchemaTemplate(request *GetCheckInSchemaTemplateRequest) (_result *GetCheckInSchemaTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCheckInSchemaTemplateHeaders{}
	_result = &GetCheckInSchemaTemplateResponse{}
	_body, _err := client.GetCheckInSchemaTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCheckinRecordByUserWithOptions(request *GetCheckinRecordByUserRequest, headers *GetCheckinRecordByUserHeaders, runtime *util.RuntimeOptions) (_result *GetCheckinRecordByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("GetCheckinRecordByUser"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/checkin/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCheckinRecordByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCheckinRecordByUser(request *GetCheckinRecordByUserRequest) (_result *GetCheckinRecordByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCheckinRecordByUserHeaders{}
	_result = &GetCheckinRecordByUserResponse{}
	_body, _err := client.GetCheckinRecordByUserWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetClosingAccounts"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/closingAccounts/rules/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClosingAccountsResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetLeaveRecordsWithOptions(request *GetLeaveRecordsRequest, headers *GetLeaveRecordsHeaders, runtime *util.RuntimeOptions) (_result *GetLeaveRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LeaveCode)) {
		body["leaveCode"] = request.LeaveCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

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
		Action:      tea.String("GetLeaveRecords"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/vacations/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLeaveRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLeaveRecords(request *GetLeaveRecordsRequest) (_result *GetLeaveRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLeaveRecordsHeaders{}
	_result = &GetLeaveRecordsResponse{}
	_body, _err := client.GetLeaveRecordsWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetLeaveType"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/leaves/types"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLeaveTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetMachineWithOptions(devId *string, headers *GetMachineHeaders, runtime *util.RuntimeOptions) (_result *GetMachineResponse, _err error) {
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
		Action:      tea.String("GetMachine"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/machines/" + tea.StringValue(devId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMachineResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetMachineUserWithOptions(devId *string, request *GetMachineUserRequest, headers *GetMachineUserHeaders, runtime *util.RuntimeOptions) (_result *GetMachineUserResponse, _err error) {
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
		Action:      tea.String("GetMachineUser"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/machines/getUser/" + tea.StringValue(devId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMachineUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetOvertimeSetting"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/overtimeSettings/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOvertimeSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GetShiftWithOptions(request *GetShiftRequest, headers *GetShiftHeaders, runtime *util.RuntimeOptions) (_result *GetShiftResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ShiftId)) {
		query["shiftId"] = request.ShiftId
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
		Action:      tea.String("GetShift"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/shifts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShiftResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShift(request *GetShiftRequest) (_result *GetShiftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetShiftHeaders{}
	_result = &GetShiftResponse{}
	_body, _err := client.GetShiftWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetSimpleOvertimeSetting"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/overtimeSettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimpleOvertimeSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetUserHolidays"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/holidays"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserHolidaysResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) GroupAddWithOptions(request *GroupAddRequest, headers *GroupAddHeaders, runtime *util.RuntimeOptions) (_result *GroupAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentSettingId)) {
		body["adjustmentSettingId"] = request.AdjustmentSettingId
	}

	if !tea.BoolValue(util.IsUnset(request.BleDeviceList)) {
		body["bleDeviceList"] = request.BleDeviceList
	}

	if !tea.BoolValue(util.IsUnset(request.CheckNeedHealthyCode)) {
		body["checkNeedHealthyCode"] = request.CheckNeedHealthyCode
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultClassId)) {
		body["defaultClassId"] = request.DefaultClassId
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCheckWhenRest)) {
		body["disableCheckWhenRest"] = request.DisableCheckWhenRest
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCheckWithoutSchedule)) {
		body["disableCheckWithoutSchedule"] = request.DisableCheckWithoutSchedule
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCameraCheck)) {
		body["enableCameraCheck"] = request.EnableCameraCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableEmpSelectClass)) {
		body["enableEmpSelectClass"] = request.EnableEmpSelectClass
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFaceCheck)) {
		body["enableFaceCheck"] = request.EnableFaceCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFaceStrictMode)) {
		body["enableFaceStrictMode"] = request.EnableFaceStrictMode
	}

	if !tea.BoolValue(util.IsUnset(request.EnableNextDay)) {
		body["enableNextDay"] = request.EnableNextDay
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutSideUpdateNormalCheck)) {
		body["enableOutSideUpdateNormalCheck"] = request.EnableOutSideUpdateNormalCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideApply)) {
		body["enableOutsideApply"] = request.EnableOutsideApply
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideCameraCheck)) {
		body["enableOutsideCameraCheck"] = request.EnableOutsideCameraCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideCheck)) {
		body["enableOutsideCheck"] = request.EnableOutsideCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideRemark)) {
		body["enableOutsideRemark"] = request.EnableOutsideRemark
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePositionBle)) {
		body["enablePositionBle"] = request.EnablePositionBle
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTrimDistance)) {
		body["enableTrimDistance"] = request.EnableTrimDistance
	}

	if !tea.BoolValue(util.IsUnset(request.ForbidHideOutSideAddress)) {
		body["forbidHideOutSideAddress"] = request.ForbidHideOutSideAddress
	}

	if !tea.BoolValue(util.IsUnset(request.FreeCheckSetting)) {
		body["freeCheckSetting"] = request.FreeCheckSetting
	}

	if !tea.BoolValue(util.IsUnset(request.FreeCheckTypeId)) {
		body["freeCheckTypeId"] = request.FreeCheckTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.FreecheckDayStartMinOffset)) {
		body["freecheckDayStartMinOffset"] = request.FreecheckDayStartMinOffset
	}

	if !tea.BoolValue(util.IsUnset(request.FreecheckWorkDays)) {
		body["freecheckWorkDays"] = request.FreecheckWorkDays
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerList)) {
		body["managerList"] = request.ManagerList
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyMember)) {
		body["modifyMember"] = request.ModifyMember
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OpenFaceCheck)) {
		body["openFaceCheck"] = request.OpenFaceCheck
	}

	if !tea.BoolValue(util.IsUnset(request.OutsideCheckApproveModeId)) {
		body["outsideCheckApproveModeId"] = request.OutsideCheckApproveModeId
	}

	if !tea.BoolValue(util.IsUnset(request.OvertimeSettingId)) {
		body["overtimeSettingId"] = request.OvertimeSettingId
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.Positions)) {
		body["positions"] = request.Positions
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePermissionMap)) {
		body["resourcePermissionMap"] = request.ResourcePermissionMap
	}

	if !tea.BoolValue(util.IsUnset(request.ShiftVOList)) {
		body["shiftVOList"] = request.ShiftVOList
	}

	if !tea.BoolValue(util.IsUnset(request.SkipHolidays)) {
		body["skipHolidays"] = request.SkipHolidays
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialDays)) {
		body["specialDays"] = request.SpecialDays
	}

	if !tea.BoolValue(util.IsUnset(request.TrimDistance)) {
		body["trimDistance"] = request.TrimDistance
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Wifis)) {
		body["wifis"] = request.Wifis
	}

	if !tea.BoolValue(util.IsUnset(request.WorkdayClassList)) {
		body["workdayClassList"] = request.WorkdayClassList
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
	params := &openapi.Params{
		Action:      tea.String("GroupAdd"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupAddResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupAdd(request *GroupAddRequest) (_result *GroupAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupAddHeaders{}
	_result = &GroupAddResponse{}
	_body, _err := client.GroupAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupUpdateWithOptions(request *GroupUpdateRequest, headers *GroupUpdateHeaders, runtime *util.RuntimeOptions) (_result *GroupUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentSettingId)) {
		body["adjustmentSettingId"] = request.AdjustmentSettingId
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCheckWhenRest)) {
		body["disableCheckWhenRest"] = request.DisableCheckWhenRest
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCheckWithoutSchedule)) {
		body["disableCheckWithoutSchedule"] = request.DisableCheckWithoutSchedule
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCameraCheck)) {
		body["enableCameraCheck"] = request.EnableCameraCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableEmpSelectClass)) {
		body["enableEmpSelectClass"] = request.EnableEmpSelectClass
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFaceCheck)) {
		body["enableFaceCheck"] = request.EnableFaceCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFaceStrictMode)) {
		body["enableFaceStrictMode"] = request.EnableFaceStrictMode
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutSideUpdateNormalCheck)) {
		body["enableOutSideUpdateNormalCheck"] = request.EnableOutSideUpdateNormalCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideApply)) {
		body["enableOutsideApply"] = request.EnableOutsideApply
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideCheck)) {
		body["enableOutsideCheck"] = request.EnableOutsideCheck
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutsideRemark)) {
		body["enableOutsideRemark"] = request.EnableOutsideRemark
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTrimDistance)) {
		body["enableTrimDistance"] = request.EnableTrimDistance
	}

	if !tea.BoolValue(util.IsUnset(request.ForbidHideOutSideAddress)) {
		body["forbidHideOutSideAddress"] = request.ForbidHideOutSideAddress
	}

	if !tea.BoolValue(util.IsUnset(request.FreeCheckSetting)) {
		body["freeCheckSetting"] = request.FreeCheckSetting
	}

	if !tea.BoolValue(util.IsUnset(request.FreeCheckTypeId)) {
		body["freeCheckTypeId"] = request.FreeCheckTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerList)) {
		body["managerList"] = request.ManagerList
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OpenFaceCheck)) {
		body["openFaceCheck"] = request.OpenFaceCheck
	}

	if !tea.BoolValue(util.IsUnset(request.OutsideCheckApproveModeId)) {
		body["outsideCheckApproveModeId"] = request.OutsideCheckApproveModeId
	}

	if !tea.BoolValue(util.IsUnset(request.OvertimeSettingId)) {
		body["overtimeSettingId"] = request.OvertimeSettingId
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.Positions)) {
		body["positions"] = request.Positions
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePermissionMap)) {
		body["resourcePermissionMap"] = request.ResourcePermissionMap
	}

	if !tea.BoolValue(util.IsUnset(request.ShiftVOList)) {
		body["shiftVOList"] = request.ShiftVOList
	}

	if !tea.BoolValue(util.IsUnset(request.SkipHolidays)) {
		body["skipHolidays"] = request.SkipHolidays
	}

	if !tea.BoolValue(util.IsUnset(request.TrimDistance)) {
		body["trimDistance"] = request.TrimDistance
	}

	if !tea.BoolValue(util.IsUnset(request.WorkdayClassList)) {
		body["workdayClassList"] = request.WorkdayClassList
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
	params := &openapi.Params{
		Action:      tea.String("GroupUpdate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/groups"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupUpdate(request *GroupUpdateRequest) (_result *GroupUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupUpdateHeaders{}
	_result = &GroupUpdateResponse{}
	_body, _err := client.GroupUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitAndGetLeaveALlocationQuotasWithOptions(request *InitAndGetLeaveALlocationQuotasRequest, headers *InitAndGetLeaveALlocationQuotasHeaders, runtime *util.RuntimeOptions) (_result *InitAndGetLeaveALlocationQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LeaveCode)) {
		query["leaveCode"] = request.LeaveCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("InitAndGetLeaveALlocationQuotas"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/leaves/initializations/balances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InitAndGetLeaveALlocationQuotasResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitAndGetLeaveALlocationQuotas(request *InitAndGetLeaveALlocationQuotasRequest) (_result *InitAndGetLeaveALlocationQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitAndGetLeaveALlocationQuotasHeaders{}
	_result = &InitAndGetLeaveALlocationQuotasResponse{}
	_body, _err := client.InitAndGetLeaveALlocationQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApproveByUsersWithOptions(request *ListApproveByUsersRequest, headers *ListApproveByUsersHeaders, runtime *util.RuntimeOptions) (_result *ListApproveByUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		body["bizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.FromDateTime)) {
		body["fromDateTime"] = request.FromDateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ToDateTime)) {
		body["toDateTime"] = request.ToDateTime
	}

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
		Action:      tea.String("ListApproveByUsers"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/approvals/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApproveByUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApproveByUsers(request *ListApproveByUsersRequest) (_result *ListApproveByUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListApproveByUsersHeaders{}
	_result = &ListApproveByUsersResponse{}
	_body, _err := client.ListApproveByUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWaterMarkTemplateWithOptions(request *ModifyWaterMarkTemplateRequest, headers *ModifyWaterMarkTemplateHeaders, runtime *util.RuntimeOptions) (_result *ModifyWaterMarkTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		body["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutDesignId)) {
		body["layoutDesignId"] = request.LayoutDesignId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaContent)) {
		body["schemaContent"] = request.SchemaContent
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.WaterMarkId)) {
		body["waterMarkId"] = request.WaterMarkId
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
	params := &openapi.Params{
		Action:      tea.String("ModifyWaterMarkTemplate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/watermarks/templates"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWaterMarkTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWaterMarkTemplate(request *ModifyWaterMarkTemplateRequest) (_result *ModifyWaterMarkTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ModifyWaterMarkTemplateHeaders{}
	_result = &ModifyWaterMarkTemplateResponse{}
	_body, _err := client.ModifyWaterMarkTemplateWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.PunchParam)) {
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
	params := &openapi.Params{
		Action:      tea.String("ProcessApproveCreate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/workflows/checkInForms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ProcessApproveCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) RetainLeaveTypesWithOptions(request *RetainLeaveTypesRequest, headers *RetainLeaveTypesHeaders, runtime *util.RuntimeOptions) (_result *RetainLeaveTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LeaveCodes)) {
		body["leaveCodes"] = request.LeaveCodes
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
		Action:      tea.String("RetainLeaveTypes"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/vacations/types/change"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RetainLeaveTypesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetainLeaveTypes(request *RetainLeaveTypesRequest) (_result *RetainLeaveTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RetainLeaveTypesHeaders{}
	_result = &RetainLeaveTypesResponse{}
	_body, _err := client.RetainLeaveTypesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveCustomWaterMarkTemplateWithOptions(request *SaveCustomWaterMarkTemplateRequest, headers *SaveCustomWaterMarkTemplateHeaders, runtime *util.RuntimeOptions) (_result *SaveCustomWaterMarkTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutDesignId)) {
		body["layoutDesignId"] = request.LayoutDesignId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		body["sceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaContent)) {
		body["schemaContent"] = request.SchemaContent
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveCustomWaterMarkTemplate"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/watermarks/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveCustomWaterMarkTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveCustomWaterMarkTemplate(request *SaveCustomWaterMarkTemplateRequest) (_result *SaveCustomWaterMarkTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveCustomWaterMarkTemplateHeaders{}
	_result = &SaveCustomWaterMarkTemplateResponse{}
	_body, _err := client.SaveCustomWaterMarkTemplateWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("SyncScheduleInfo"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/schedules/additionalInfo"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &SyncScheduleInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.LeaveCertificate)) {
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

	if !tea.BoolValue(util.IsUnset(request.SubmitTimeRule)) {
		body["submitTimeRule"] = request.SubmitTimeRule
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityRules)) {
		body["visibilityRules"] = request.VisibilityRules
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
	params := &openapi.Params{
		Action:      tea.String("UpdateLeaveType"),
		Version:     tea.String("attendance_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/attendance/leaves/types"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLeaveTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
