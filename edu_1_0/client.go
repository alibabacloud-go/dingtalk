// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package edu_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ActivateDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ActivateDeviceHeaders) SetCommonHeaders(v map[string]*string) *ActivateDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ActivateDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *ActivateDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ActivateDeviceRequest struct {
	LicenseKey *string `json:"licenseKey,omitempty" xml:"licenseKey,omitempty"`
	Model      *string `json:"model,omitempty" xml:"model,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ActivateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActivateDeviceRequest) SetLicenseKey(v string) *ActivateDeviceRequest {
	s.LicenseKey = &v
	return s
}

func (s *ActivateDeviceRequest) SetModel(v string) *ActivateDeviceRequest {
	s.Model = &v
	return s
}

func (s *ActivateDeviceRequest) SetName(v string) *ActivateDeviceRequest {
	s.Name = &v
	return s
}

func (s *ActivateDeviceRequest) SetSn(v string) *ActivateDeviceRequest {
	s.Sn = &v
	return s
}

func (s *ActivateDeviceRequest) SetType(v string) *ActivateDeviceRequest {
	s.Type = &v
	return s
}

type ActivateDeviceResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ActivateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponseBody) SetResult(v bool) *ActivateDeviceResponseBody {
	s.Result = &v
	return s
}

type ActivateDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActivateDeviceResponse) SetHeaders(v map[string]*string) *ActivateDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActivateDeviceResponse) SetStatusCode(v int32) *ActivateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateDeviceResponse) SetBody(v *ActivateDeviceResponseBody) *ActivateDeviceResponse {
	s.Body = v
	return s
}

type AddCompetitionRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCompetitionRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCompetitionRecordHeaders) GoString() string {
	return s.String()
}

func (s *AddCompetitionRecordHeaders) SetCommonHeaders(v map[string]*string) *AddCompetitionRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCompetitionRecordHeaders) SetXAcsDingtalkAccessToken(v string) *AddCompetitionRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCompetitionRecordRequest struct {
	CompetitionCode   *string `json:"competitionCode,omitempty" xml:"competitionCode,omitempty"`
	GroupTemplateCode *string `json:"groupTemplateCode,omitempty" xml:"groupTemplateCode,omitempty"`
	JoinGroup         *bool   `json:"joinGroup,omitempty" xml:"joinGroup,omitempty"`
	ParticipantName   *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
	UnionId           *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddCompetitionRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCompetitionRecordRequest) GoString() string {
	return s.String()
}

func (s *AddCompetitionRecordRequest) SetCompetitionCode(v string) *AddCompetitionRecordRequest {
	s.CompetitionCode = &v
	return s
}

func (s *AddCompetitionRecordRequest) SetGroupTemplateCode(v string) *AddCompetitionRecordRequest {
	s.GroupTemplateCode = &v
	return s
}

func (s *AddCompetitionRecordRequest) SetJoinGroup(v bool) *AddCompetitionRecordRequest {
	s.JoinGroup = &v
	return s
}

func (s *AddCompetitionRecordRequest) SetParticipantName(v string) *AddCompetitionRecordRequest {
	s.ParticipantName = &v
	return s
}

func (s *AddCompetitionRecordRequest) SetUnionId(v string) *AddCompetitionRecordRequest {
	s.UnionId = &v
	return s
}

type AddCompetitionRecordResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddCompetitionRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCompetitionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddCompetitionRecordResponseBody) SetSuccess(v bool) *AddCompetitionRecordResponseBody {
	s.Success = &v
	return s
}

type AddCompetitionRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCompetitionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCompetitionRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCompetitionRecordResponse) GoString() string {
	return s.String()
}

func (s *AddCompetitionRecordResponse) SetHeaders(v map[string]*string) *AddCompetitionRecordResponse {
	s.Headers = v
	return s
}

func (s *AddCompetitionRecordResponse) SetStatusCode(v int32) *AddCompetitionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCompetitionRecordResponse) SetBody(v *AddCompetitionRecordResponseBody) *AddCompetitionRecordResponse {
	s.Body = v
	return s
}

type AddDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceHeaders) GoString() string {
	return s.String()
}

func (s *AddDeviceHeaders) SetCommonHeaders(v map[string]*string) *AddDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *AddDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddDeviceRequest struct {
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Model      *string `json:"model,omitempty" xml:"model,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Scene      *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
	Type       *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AddDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceRequest) SetMerchantId(v string) *AddDeviceRequest {
	s.MerchantId = &v
	return s
}

func (s *AddDeviceRequest) SetModel(v string) *AddDeviceRequest {
	s.Model = &v
	return s
}

func (s *AddDeviceRequest) SetName(v string) *AddDeviceRequest {
	s.Name = &v
	return s
}

func (s *AddDeviceRequest) SetScene(v int64) *AddDeviceRequest {
	s.Scene = &v
	return s
}

func (s *AddDeviceRequest) SetSn(v string) *AddDeviceRequest {
	s.Sn = &v
	return s
}

func (s *AddDeviceRequest) SetStatus(v int64) *AddDeviceRequest {
	s.Status = &v
	return s
}

func (s *AddDeviceRequest) SetType(v int64) *AddDeviceRequest {
	s.Type = &v
	return s
}

type AddDeviceResponseBody struct {
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AddDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDeviceResponseBody) SetCorpId(v string) *AddDeviceResponseBody {
	s.CorpId = &v
	return s
}

func (s *AddDeviceResponseBody) SetId(v int64) *AddDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *AddDeviceResponseBody) SetMerchantId(v string) *AddDeviceResponseBody {
	s.MerchantId = &v
	return s
}

func (s *AddDeviceResponseBody) SetSn(v string) *AddDeviceResponseBody {
	s.Sn = &v
	return s
}

func (s *AddDeviceResponseBody) SetStatus(v int64) *AddDeviceResponseBody {
	s.Status = &v
	return s
}

type AddDeviceResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddDeviceResponse) SetHeaders(v map[string]*string) *AddDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddDeviceResponse) SetStatusCode(v int32) *AddDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDeviceResponse) SetBody(v *AddDeviceResponseBody) *AddDeviceResponse {
	s.Body = v
	return s
}

type AddSchoolConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddSchoolConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddSchoolConfigHeaders) GoString() string {
	return s.String()
}

func (s *AddSchoolConfigHeaders) SetCommonHeaders(v map[string]*string) *AddSchoolConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddSchoolConfigHeaders) SetXAcsDingtalkAccessToken(v string) *AddSchoolConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddSchoolConfigRequest struct {
	OperatorId         *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OperatorName       *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	TemperatureUpLimit *int64  `json:"temperatureUpLimit,omitempty" xml:"temperatureUpLimit,omitempty"`
}

func (s AddSchoolConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSchoolConfigRequest) GoString() string {
	return s.String()
}

func (s *AddSchoolConfigRequest) SetOperatorId(v string) *AddSchoolConfigRequest {
	s.OperatorId = &v
	return s
}

func (s *AddSchoolConfigRequest) SetOperatorName(v string) *AddSchoolConfigRequest {
	s.OperatorName = &v
	return s
}

func (s *AddSchoolConfigRequest) SetTemperatureUpLimit(v int64) *AddSchoolConfigRequest {
	s.TemperatureUpLimit = &v
	return s
}

type AddSchoolConfigResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddSchoolConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSchoolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddSchoolConfigResponseBody) SetResult(v bool) *AddSchoolConfigResponseBody {
	s.Result = &v
	return s
}

type AddSchoolConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddSchoolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSchoolConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSchoolConfigResponse) GoString() string {
	return s.String()
}

func (s *AddSchoolConfigResponse) SetHeaders(v map[string]*string) *AddSchoolConfigResponse {
	s.Headers = v
	return s
}

func (s *AddSchoolConfigResponse) SetStatusCode(v int32) *AddSchoolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSchoolConfigResponse) SetBody(v *AddSchoolConfigResponseBody) *AddSchoolConfigResponse {
	s.Body = v
	return s
}

type AssignClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AssignClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s AssignClassHeaders) GoString() string {
	return s.String()
}

func (s *AssignClassHeaders) SetCommonHeaders(v map[string]*string) *AssignClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignClassHeaders) SetXAcsDingtalkAccessToken(v string) *AssignClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AssignClassRequest struct {
	ClassId   *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	IsFinish  *bool   `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	Operator  *string `json:"operator,omitempty" xml:"operator,omitempty"`
	StudentId *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
	TaskId    *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AssignClassRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignClassRequest) GoString() string {
	return s.String()
}

func (s *AssignClassRequest) SetClassId(v int64) *AssignClassRequest {
	s.ClassId = &v
	return s
}

func (s *AssignClassRequest) SetIsFinish(v bool) *AssignClassRequest {
	s.IsFinish = &v
	return s
}

func (s *AssignClassRequest) SetOperator(v string) *AssignClassRequest {
	s.Operator = &v
	return s
}

func (s *AssignClassRequest) SetStudentId(v int64) *AssignClassRequest {
	s.StudentId = &v
	return s
}

func (s *AssignClassRequest) SetTaskId(v int64) *AssignClassRequest {
	s.TaskId = &v
	return s
}

type AssignClassResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AssignClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignClassResponseBody) GoString() string {
	return s.String()
}

func (s *AssignClassResponseBody) SetSuccess(v bool) *AssignClassResponseBody {
	s.Success = &v
	return s
}

type AssignClassResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssignClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignClassResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignClassResponse) GoString() string {
	return s.String()
}

func (s *AssignClassResponse) SetHeaders(v map[string]*string) *AssignClassResponse {
	s.Headers = v
	return s
}

func (s *AssignClassResponse) SetStatusCode(v int32) *AssignClassResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignClassResponse) SetBody(v *AssignClassResponseBody) *AssignClassResponse {
	s.Body = v
	return s
}

type BatchCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateHeaders) GoString() string {
	return s.String()
}

func (s *BatchCreateHeaders) SetCommonHeaders(v map[string]*string) *BatchCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchCreateHeaders) SetXAcsDingtalkAccessToken(v string) *BatchCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchCreateRequest struct {
	CardBizCode *string                 `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	Data        *BatchCreateRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Identifier  *string                 `json:"identifier,omitempty" xml:"identifier,omitempty"`
	JsVersion   *int32                  `json:"jsVersion,omitempty" xml:"jsVersion,omitempty"`
	SourceType  *string                 `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Userid      *string                 `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s BatchCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateRequest) SetCardBizCode(v string) *BatchCreateRequest {
	s.CardBizCode = &v
	return s
}

func (s *BatchCreateRequest) SetData(v *BatchCreateRequestData) *BatchCreateRequest {
	s.Data = v
	return s
}

func (s *BatchCreateRequest) SetIdentifier(v string) *BatchCreateRequest {
	s.Identifier = &v
	return s
}

func (s *BatchCreateRequest) SetJsVersion(v int32) *BatchCreateRequest {
	s.JsVersion = &v
	return s
}

func (s *BatchCreateRequest) SetSourceType(v string) *BatchCreateRequest {
	s.SourceType = &v
	return s
}

func (s *BatchCreateRequest) SetUserid(v string) *BatchCreateRequest {
	s.Userid = &v
	return s
}

type BatchCreateRequestData struct {
	CanReissueCard           *bool                                             `json:"canReissueCard,omitempty" xml:"canReissueCard,omitempty"`
	CardCycle                *int32                                            `json:"cardCycle,omitempty" xml:"cardCycle,omitempty"`
	CardFrequency            []*int32                                          `json:"cardFrequency,omitempty" xml:"cardFrequency,omitempty" type:"Repeated"`
	CardRuleItemParamList    []*BatchCreateRequestDataCardRuleItemParamList    `json:"cardRuleItemParamList,omitempty" xml:"cardRuleItemParamList,omitempty" type:"Repeated"`
	ClassIds                 []*string                                         `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	ClassNames               []*string                                         `json:"classNames,omitempty" xml:"classNames,omitempty" type:"Repeated"`
	Content                  *string                                           `json:"content,omitempty" xml:"content,omitempty"`
	EffectDate               *int64                                            `json:"effectDate,omitempty" xml:"effectDate,omitempty"`
	Medias                   *string                                           `json:"medias,omitempty" xml:"medias,omitempty"`
	NeedMetering             *string                                           `json:"needMetering,omitempty" xml:"needMetering,omitempty"`
	OrgClassStudentGroupList []*BatchCreateRequestDataOrgClassStudentGroupList `json:"orgClassStudentGroupList,omitempty" xml:"orgClassStudentGroupList,omitempty" type:"Repeated"`
	RemindHour               *int32                                            `json:"remindHour,omitempty" xml:"remindHour,omitempty"`
	RemindMinute             *int32                                            `json:"remindMinute,omitempty" xml:"remindMinute,omitempty"`
	TargetRole               *string                                           `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	TemplateId               *int64                                            `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title                    *string                                           `json:"title,omitempty" xml:"title,omitempty"`
	UnitOfMeasurement        *string                                           `json:"unitOfMeasurement,omitempty" xml:"unitOfMeasurement,omitempty"`
}

func (s BatchCreateRequestData) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestData) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestData) SetCanReissueCard(v bool) *BatchCreateRequestData {
	s.CanReissueCard = &v
	return s
}

func (s *BatchCreateRequestData) SetCardCycle(v int32) *BatchCreateRequestData {
	s.CardCycle = &v
	return s
}

func (s *BatchCreateRequestData) SetCardFrequency(v []*int32) *BatchCreateRequestData {
	s.CardFrequency = v
	return s
}

func (s *BatchCreateRequestData) SetCardRuleItemParamList(v []*BatchCreateRequestDataCardRuleItemParamList) *BatchCreateRequestData {
	s.CardRuleItemParamList = v
	return s
}

func (s *BatchCreateRequestData) SetClassIds(v []*string) *BatchCreateRequestData {
	s.ClassIds = v
	return s
}

func (s *BatchCreateRequestData) SetClassNames(v []*string) *BatchCreateRequestData {
	s.ClassNames = v
	return s
}

func (s *BatchCreateRequestData) SetContent(v string) *BatchCreateRequestData {
	s.Content = &v
	return s
}

func (s *BatchCreateRequestData) SetEffectDate(v int64) *BatchCreateRequestData {
	s.EffectDate = &v
	return s
}

func (s *BatchCreateRequestData) SetMedias(v string) *BatchCreateRequestData {
	s.Medias = &v
	return s
}

func (s *BatchCreateRequestData) SetNeedMetering(v string) *BatchCreateRequestData {
	s.NeedMetering = &v
	return s
}

func (s *BatchCreateRequestData) SetOrgClassStudentGroupList(v []*BatchCreateRequestDataOrgClassStudentGroupList) *BatchCreateRequestData {
	s.OrgClassStudentGroupList = v
	return s
}

func (s *BatchCreateRequestData) SetRemindHour(v int32) *BatchCreateRequestData {
	s.RemindHour = &v
	return s
}

func (s *BatchCreateRequestData) SetRemindMinute(v int32) *BatchCreateRequestData {
	s.RemindMinute = &v
	return s
}

func (s *BatchCreateRequestData) SetTargetRole(v string) *BatchCreateRequestData {
	s.TargetRole = &v
	return s
}

func (s *BatchCreateRequestData) SetTemplateId(v int64) *BatchCreateRequestData {
	s.TemplateId = &v
	return s
}

func (s *BatchCreateRequestData) SetTitle(v string) *BatchCreateRequestData {
	s.Title = &v
	return s
}

func (s *BatchCreateRequestData) SetUnitOfMeasurement(v string) *BatchCreateRequestData {
	s.UnitOfMeasurement = &v
	return s
}

type BatchCreateRequestDataCardRuleItemParamList struct {
	CardRuleAttr  *string `json:"cardRuleAttr,omitempty" xml:"cardRuleAttr,omitempty"`
	CardTaskCode  *string `json:"cardTaskCode,omitempty" xml:"cardTaskCode,omitempty"`
	DailyDubbing  *int32  `json:"dailyDubbing,omitempty" xml:"dailyDubbing,omitempty"`
	RelationId    *string `json:"relationId,omitempty" xml:"relationId,omitempty"`
	RelationTitle *string `json:"relationTitle,omitempty" xml:"relationTitle,omitempty"`
	RelationUrl   *string `json:"relationUrl,omitempty" xml:"relationUrl,omitempty"`
}

func (s BatchCreateRequestDataCardRuleItemParamList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataCardRuleItemParamList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetCardRuleAttr(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.CardRuleAttr = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetCardTaskCode(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.CardTaskCode = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetDailyDubbing(v int32) *BatchCreateRequestDataCardRuleItemParamList {
	s.DailyDubbing = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationId(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationId = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationTitle(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationTitle = &v
	return s
}

func (s *BatchCreateRequestDataCardRuleItemParamList) SetRelationUrl(v string) *BatchCreateRequestDataCardRuleItemParamList {
	s.RelationUrl = &v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupList struct {
	ClassList []*BatchCreateRequestDataOrgClassStudentGroupListClassList `json:"classList,omitempty" xml:"classList,omitempty" type:"Repeated"`
	CorpId    *string                                                    `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupList) SetClassList(v []*BatchCreateRequestDataOrgClassStudentGroupListClassList) *BatchCreateRequestDataOrgClassStudentGroupList {
	s.ClassList = v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupList) SetCorpId(v string) *BatchCreateRequestDataOrgClassStudentGroupList {
	s.CorpId = &v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupListClassList struct {
	ClassId   *int64                                                             `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName *string                                                            `json:"className,omitempty" xml:"className,omitempty"`
	Students  []*BatchCreateRequestDataOrgClassStudentGroupListClassListStudents `json:"students,omitempty" xml:"students,omitempty" type:"Repeated"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassList) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetClassId(v int64) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.ClassId = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetClassName(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.ClassName = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassList) SetStudents(v []*BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) *BatchCreateRequestDataOrgClassStudentGroupListClassList {
	s.Students = v
	return s
}

type BatchCreateRequestDataOrgClassStudentGroupListClassListStudents struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) GoString() string {
	return s.String()
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetName(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.Name = &v
	return s
}

func (s *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents) SetStaffId(v string) *BatchCreateRequestDataOrgClassStudentGroupListClassListStudents {
	s.StaffId = &v
	return s
}

type BatchCreateResponseBody struct {
	Result *BatchCreateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s BatchCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateResponseBody) SetResult(v *BatchCreateResponseBodyResult) *BatchCreateResponseBody {
	s.Result = v
	return s
}

type BatchCreateResponseBodyResult struct {
	CorpIdCardIdMap map[string]*string `json:"corpIdCardIdMap,omitempty" xml:"corpIdCardIdMap,omitempty"`
}

func (s BatchCreateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchCreateResponseBodyResult) SetCorpIdCardIdMap(v map[string]*string) *BatchCreateResponseBodyResult {
	s.CorpIdCardIdMap = v
	return s
}

type BatchCreateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateResponse) SetHeaders(v map[string]*string) *BatchCreateResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateResponse) SetStatusCode(v int32) *BatchCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateResponse) SetBody(v *BatchCreateResponseBody) *BatchCreateResponse {
	s.Body = v
	return s
}

type BatchOrgCreateHWHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchOrgCreateHWHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWHeaders) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWHeaders) SetCommonHeaders(v map[string]*string) *BatchOrgCreateHWHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchOrgCreateHWHeaders) SetXAcsDingtalkAccessToken(v string) *BatchOrgCreateHWHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchOrgCreateHWRequest struct {
	Attributes         *string                                      `json:"attributes,omitempty" xml:"attributes,omitempty"`
	BizCode            *string                                      `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	CourseName         *string                                      `json:"courseName,omitempty" xml:"courseName,omitempty"`
	HwContent          *string                                      `json:"hwContent,omitempty" xml:"hwContent,omitempty"`
	HwDeadline         *int64                                       `json:"hwDeadline,omitempty" xml:"hwDeadline,omitempty"`
	HwDeadlineOpen     *string                                      `json:"hwDeadlineOpen,omitempty" xml:"hwDeadlineOpen,omitempty"`
	HwMedia            *string                                      `json:"hwMedia,omitempty" xml:"hwMedia,omitempty"`
	HwPhoto            *string                                      `json:"hwPhoto,omitempty" xml:"hwPhoto,omitempty"`
	HwTitle            *string                                      `json:"hwTitle,omitempty" xml:"hwTitle,omitempty"`
	HwType             *string                                      `json:"hwType,omitempty" xml:"hwType,omitempty"`
	HwVideo            *string                                      `json:"hwVideo,omitempty" xml:"hwVideo,omitempty"`
	Identifier         *string                                      `json:"identifier,omitempty" xml:"identifier,omitempty"`
	OpenSelectItemList []*BatchOrgCreateHWRequestOpenSelectItemList `json:"openSelectItemList,omitempty" xml:"openSelectItemList,omitempty" type:"Repeated"`
	ScheduledRelease   *string                                      `json:"scheduledRelease,omitempty" xml:"scheduledRelease,omitempty"`
	ScheduledTime      *string                                      `json:"scheduledTime,omitempty" xml:"scheduledTime,omitempty"`
	Status             *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	TargetRole         *string                                      `json:"targetRole,omitempty" xml:"targetRole,omitempty"`
	TeacherName        *string                                      `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	TeacherUserId      *string                                      `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s BatchOrgCreateHWRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequest) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequest) SetAttributes(v string) *BatchOrgCreateHWRequest {
	s.Attributes = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetBizCode(v string) *BatchOrgCreateHWRequest {
	s.BizCode = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetCourseName(v string) *BatchOrgCreateHWRequest {
	s.CourseName = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwContent(v string) *BatchOrgCreateHWRequest {
	s.HwContent = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwDeadline(v int64) *BatchOrgCreateHWRequest {
	s.HwDeadline = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwDeadlineOpen(v string) *BatchOrgCreateHWRequest {
	s.HwDeadlineOpen = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwMedia(v string) *BatchOrgCreateHWRequest {
	s.HwMedia = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwPhoto(v string) *BatchOrgCreateHWRequest {
	s.HwPhoto = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwTitle(v string) *BatchOrgCreateHWRequest {
	s.HwTitle = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwType(v string) *BatchOrgCreateHWRequest {
	s.HwType = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetHwVideo(v string) *BatchOrgCreateHWRequest {
	s.HwVideo = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetIdentifier(v string) *BatchOrgCreateHWRequest {
	s.Identifier = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetOpenSelectItemList(v []*BatchOrgCreateHWRequestOpenSelectItemList) *BatchOrgCreateHWRequest {
	s.OpenSelectItemList = v
	return s
}

func (s *BatchOrgCreateHWRequest) SetScheduledRelease(v string) *BatchOrgCreateHWRequest {
	s.ScheduledRelease = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetScheduledTime(v string) *BatchOrgCreateHWRequest {
	s.ScheduledTime = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetStatus(v string) *BatchOrgCreateHWRequest {
	s.Status = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTargetRole(v string) *BatchOrgCreateHWRequest {
	s.TargetRole = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTeacherName(v string) *BatchOrgCreateHWRequest {
	s.TeacherName = &v
	return s
}

func (s *BatchOrgCreateHWRequest) SetTeacherUserId(v string) *BatchOrgCreateHWRequest {
	s.TeacherUserId = &v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemList struct {
	ClassList           []*BatchOrgCreateHWRequestOpenSelectItemListClassList `json:"classList,omitempty" xml:"classList,omitempty" type:"Repeated"`
	CorpId              *string                                               `json:"corpId,omitempty" xml:"corpId,omitempty"`
	SelectedClassesDesc *string                                               `json:"selectedClassesDesc,omitempty" xml:"selectedClassesDesc,omitempty"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetClassList(v []*BatchOrgCreateHWRequestOpenSelectItemListClassList) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.ClassList = v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetCorpId(v string) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.CorpId = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemList) SetSelectedClassesDesc(v string) *BatchOrgCreateHWRequestOpenSelectItemList {
	s.SelectedClassesDesc = &v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemListClassList struct {
	All       *bool                                                         `json:"all,omitempty" xml:"all,omitempty"`
	ClassId   *string                                                       `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName *string                                                       `json:"className,omitempty" xml:"className,omitempty"`
	Students  []*BatchOrgCreateHWRequestOpenSelectItemListClassListStudents `json:"students,omitempty" xml:"students,omitempty" type:"Repeated"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetAll(v bool) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.All = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetClassId(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.ClassId = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetClassName(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.ClassName = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassList) SetStudents(v []*BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) *BatchOrgCreateHWRequestOpenSelectItemListClassList {
	s.Students = v
	return s
}

type BatchOrgCreateHWRequestOpenSelectItemListClassListStudents struct {
	Avatar  *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetAvatar(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.Avatar = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetName(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.Name = &v
	return s
}

func (s *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents) SetStaffId(v string) *BatchOrgCreateHWRequestOpenSelectItemListClassListStudents {
	s.StaffId = &v
	return s
}

type BatchOrgCreateHWResponseBody struct {
	Result *BatchOrgCreateHWResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s BatchOrgCreateHWResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBody) SetResult(v *BatchOrgCreateHWResponseBodyResult) *BatchOrgCreateHWResponseBody {
	s.Result = v
	return s
}

type BatchOrgCreateHWResponseBodyResult struct {
	PublishList []*BatchOrgCreateHWResponseBodyResultPublishList `json:"publishList,omitempty" xml:"publishList,omitempty" type:"Repeated"`
}

func (s BatchOrgCreateHWResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBodyResult) SetPublishList(v []*BatchOrgCreateHWResponseBodyResultPublishList) *BatchOrgCreateHWResponseBodyResult {
	s.PublishList = v
	return s
}

type BatchOrgCreateHWResponseBodyResultPublishList struct {
	Corpid *string `json:"corpid,omitempty" xml:"corpid,omitempty"`
	Hwid   *int64  `json:"hwid,omitempty" xml:"hwid,omitempty"`
}

func (s BatchOrgCreateHWResponseBodyResultPublishList) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponseBodyResultPublishList) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponseBodyResultPublishList) SetCorpid(v string) *BatchOrgCreateHWResponseBodyResultPublishList {
	s.Corpid = &v
	return s
}

func (s *BatchOrgCreateHWResponseBodyResultPublishList) SetHwid(v int64) *BatchOrgCreateHWResponseBodyResultPublishList {
	s.Hwid = &v
	return s
}

type BatchOrgCreateHWResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchOrgCreateHWResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchOrgCreateHWResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchOrgCreateHWResponse) GoString() string {
	return s.String()
}

func (s *BatchOrgCreateHWResponse) SetHeaders(v map[string]*string) *BatchOrgCreateHWResponse {
	s.Headers = v
	return s
}

func (s *BatchOrgCreateHWResponse) SetStatusCode(v int32) *BatchOrgCreateHWResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchOrgCreateHWResponse) SetBody(v *BatchOrgCreateHWResponseBody) *BatchOrgCreateHWResponse {
	s.Body = v
	return s
}

type CancelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderHeaders) GoString() string {
	return s.String()
}

func (s *CancelOrderHeaders) SetCommonHeaders(v map[string]*string) *CancelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CancelOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelOrderRequest struct {
	FaceId    *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	OrderNo   *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetFaceId(v string) *CancelOrderRequest {
	s.FaceId = &v
	return s
}

func (s *CancelOrderRequest) SetOrderNo(v string) *CancelOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *CancelOrderRequest) SetSignature(v string) *CancelOrderRequest {
	s.Signature = &v
	return s
}

func (s *CancelOrderRequest) SetSn(v string) *CancelOrderRequest {
	s.Sn = &v
	return s
}

func (s *CancelOrderRequest) SetTimestamp(v int64) *CancelOrderRequest {
	s.Timestamp = &v
	return s
}

func (s *CancelOrderRequest) SetUserId(v string) *CancelOrderRequest {
	s.UserId = &v
	return s
}

type CancelOrderResponseBody struct {
	NeedRetry   *bool   `json:"needRetry,omitempty" xml:"needRetry,omitempty"`
	TradeAction *string `json:"tradeAction,omitempty" xml:"tradeAction,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) SetNeedRetry(v bool) *CancelOrderResponseBody {
	s.NeedRetry = &v
	return s
}

func (s *CancelOrderResponseBody) SetTradeAction(v string) *CancelOrderResponseBody {
	s.TradeAction = &v
	return s
}

type CancelOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderResponse) SetHeaders(v map[string]*string) *CancelOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderResponse) SetStatusCode(v int32) *CancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderResponse) SetBody(v *CancelOrderResponseBody) *CancelOrderResponse {
	s.Body = v
	return s
}

type CancelSnsOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelSnsOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelSnsOrderHeaders) GoString() string {
	return s.String()
}

func (s *CancelSnsOrderHeaders) SetCommonHeaders(v map[string]*string) *CancelSnsOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelSnsOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CancelSnsOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelSnsOrderRequest struct {
	AlipayAppId *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId  *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo     *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Timestamp   *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s CancelSnsOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSnsOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelSnsOrderRequest) SetAlipayAppId(v string) *CancelSnsOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *CancelSnsOrderRequest) SetMerchantId(v string) *CancelSnsOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *CancelSnsOrderRequest) SetOrderNo(v string) *CancelSnsOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *CancelSnsOrderRequest) SetSignature(v string) *CancelSnsOrderRequest {
	s.Signature = &v
	return s
}

func (s *CancelSnsOrderRequest) SetTimestamp(v int64) *CancelSnsOrderRequest {
	s.Timestamp = &v
	return s
}

type CancelSnsOrderResponseBody struct {
	AlipayAppId     *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId      *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PayStatus       *int32  `json:"payStatus,omitempty" xml:"payStatus,omitempty"`
	RefundStatus    *int32  `json:"refundStatus,omitempty" xml:"refundStatus,omitempty"`
	TotalAmount     *int64  `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s CancelSnsOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSnsOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSnsOrderResponseBody) SetAlipayAppId(v string) *CancelSnsOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetMerchantId(v string) *CancelSnsOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetMerchantOrderNo(v string) *CancelSnsOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetOrderNo(v string) *CancelSnsOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetPayStatus(v int32) *CancelSnsOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetRefundStatus(v int32) *CancelSnsOrderResponseBody {
	s.RefundStatus = &v
	return s
}

func (s *CancelSnsOrderResponseBody) SetTotalAmount(v int64) *CancelSnsOrderResponseBody {
	s.TotalAmount = &v
	return s
}

type CancelSnsOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelSnsOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelSnsOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSnsOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelSnsOrderResponse) SetHeaders(v map[string]*string) *CancelSnsOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelSnsOrderResponse) SetStatusCode(v int32) *CancelSnsOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSnsOrderResponse) SetBody(v *CancelSnsOrderResponseBody) *CancelSnsOrderResponse {
	s.Body = v
	return s
}

type CancelUserOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelUserOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelUserOrderHeaders) GoString() string {
	return s.String()
}

func (s *CancelUserOrderHeaders) SetCommonHeaders(v map[string]*string) *CancelUserOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelUserOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CancelUserOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelUserOrderRequest struct {
	AlipayAppId *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId  *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo     *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Timestamp   *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s CancelUserOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUserOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelUserOrderRequest) SetAlipayAppId(v string) *CancelUserOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *CancelUserOrderRequest) SetMerchantId(v string) *CancelUserOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *CancelUserOrderRequest) SetOrderNo(v string) *CancelUserOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *CancelUserOrderRequest) SetSignature(v string) *CancelUserOrderRequest {
	s.Signature = &v
	return s
}

func (s *CancelUserOrderRequest) SetTimestamp(v int64) *CancelUserOrderRequest {
	s.Timestamp = &v
	return s
}

type CancelUserOrderResponseBody struct {
	AlipayAppId     *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId      *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PayStatus       *int32  `json:"payStatus,omitempty" xml:"payStatus,omitempty"`
	RefundStatus    *int32  `json:"refundStatus,omitempty" xml:"refundStatus,omitempty"`
	TotalAmount     *int64  `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s CancelUserOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUserOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUserOrderResponseBody) SetAlipayAppId(v string) *CancelUserOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetMerchantId(v string) *CancelUserOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetMerchantOrderNo(v string) *CancelUserOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetOrderNo(v string) *CancelUserOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetPayStatus(v int32) *CancelUserOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetRefundStatus(v int32) *CancelUserOrderResponseBody {
	s.RefundStatus = &v
	return s
}

func (s *CancelUserOrderResponseBody) SetTotalAmount(v int64) *CancelUserOrderResponseBody {
	s.TotalAmount = &v
	return s
}

type CancelUserOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelUserOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUserOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUserOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelUserOrderResponse) SetHeaders(v map[string]*string) *CancelUserOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelUserOrderResponse) SetStatusCode(v int32) *CancelUserOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUserOrderResponse) SetBody(v *CancelUserOrderResponseBody) *CancelUserOrderResponse {
	s.Body = v
	return s
}

type CardBatchQueryCardsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardBatchQueryCardsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsHeaders) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsHeaders) SetCommonHeaders(v map[string]*string) *CardBatchQueryCardsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardBatchQueryCardsHeaders) SetXAcsDingtalkAccessToken(v string) *CardBatchQueryCardsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardBatchQueryCardsRequest struct {
	CardBizCode *string  `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardIds     []*int64 `json:"cardIds,omitempty" xml:"cardIds,omitempty" type:"Repeated"`
	SourceType  *string  `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UserId      *string  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CardBatchQueryCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsRequest) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsRequest) SetCardBizCode(v string) *CardBatchQueryCardsRequest {
	s.CardBizCode = &v
	return s
}

func (s *CardBatchQueryCardsRequest) SetCardIds(v []*int64) *CardBatchQueryCardsRequest {
	s.CardIds = v
	return s
}

func (s *CardBatchQueryCardsRequest) SetSourceType(v string) *CardBatchQueryCardsRequest {
	s.SourceType = &v
	return s
}

func (s *CardBatchQueryCardsRequest) SetUserId(v string) *CardBatchQueryCardsRequest {
	s.UserId = &v
	return s
}

type CardBatchQueryCardsResponseBody struct {
	Result  *CardBatchQueryCardsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardBatchQueryCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsResponseBody) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsResponseBody) SetResult(v *CardBatchQueryCardsResponseBodyResult) *CardBatchQueryCardsResponseBody {
	s.Result = v
	return s
}

func (s *CardBatchQueryCardsResponseBody) SetSuccess(v bool) *CardBatchQueryCardsResponseBody {
	s.Success = &v
	return s
}

type CardBatchQueryCardsResponseBodyResult struct {
	Cards []*CardBatchQueryCardsResponseBodyResultCards `json:"cards,omitempty" xml:"cards,omitempty" type:"Repeated"`
}

func (s CardBatchQueryCardsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsResponseBodyResult) SetCards(v []*CardBatchQueryCardsResponseBodyResultCards) *CardBatchQueryCardsResponseBodyResult {
	s.Cards = v
	return s
}

type CardBatchQueryCardsResponseBodyResultCards struct {
	CardBizCode    *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardId         *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardStatus     *int32  `json:"cardStatus,omitempty" xml:"cardStatus,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	CorpId         *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EffectTime     *string `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	Finished       *bool   `json:"finished,omitempty" xml:"finished,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	OptEndTime     *string `json:"optEndTime,omitempty" xml:"optEndTime,omitempty"`
	OptEndUserId   *string `json:"optEndUserId,omitempty" xml:"optEndUserId,omitempty"`
	OptEndUserName *string `json:"optEndUserName,omitempty" xml:"optEndUserName,omitempty"`
	SendTime       *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	StartTime      *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TeacherId      *string `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
	TeacherName    *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	Type           *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CardBatchQueryCardsResponseBodyResultCards) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsResponseBodyResultCards) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetCardBizCode(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.CardBizCode = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetCardId(v int64) *CardBatchQueryCardsResponseBodyResultCards {
	s.CardId = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetCardStatus(v int32) *CardBatchQueryCardsResponseBodyResultCards {
	s.CardStatus = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetContent(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.Content = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetCorpId(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.CorpId = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetEffectTime(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.EffectTime = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetFinished(v bool) *CardBatchQueryCardsResponseBodyResultCards {
	s.Finished = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetGmtCreate(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.GmtCreate = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetOptEndTime(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.OptEndTime = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetOptEndUserId(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.OptEndUserId = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetOptEndUserName(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.OptEndUserName = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetSendTime(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.SendTime = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetStartTime(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.StartTime = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetStatus(v int32) *CardBatchQueryCardsResponseBodyResultCards {
	s.Status = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetTeacherId(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.TeacherId = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetTeacherName(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.TeacherName = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetTitle(v string) *CardBatchQueryCardsResponseBodyResultCards {
	s.Title = &v
	return s
}

func (s *CardBatchQueryCardsResponseBodyResultCards) SetType(v int32) *CardBatchQueryCardsResponseBodyResultCards {
	s.Type = &v
	return s
}

type CardBatchQueryCardsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardBatchQueryCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardBatchQueryCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s CardBatchQueryCardsResponse) GoString() string {
	return s.String()
}

func (s *CardBatchQueryCardsResponse) SetHeaders(v map[string]*string) *CardBatchQueryCardsResponse {
	s.Headers = v
	return s
}

func (s *CardBatchQueryCardsResponse) SetStatusCode(v int32) *CardBatchQueryCardsResponse {
	s.StatusCode = &v
	return s
}

func (s *CardBatchQueryCardsResponse) SetBody(v *CardBatchQueryCardsResponseBody) *CardBatchQueryCardsResponse {
	s.Body = v
	return s
}

type CardDeleteCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardDeleteCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteCardHeaders) GoString() string {
	return s.String()
}

func (s *CardDeleteCardHeaders) SetCommonHeaders(v map[string]*string) *CardDeleteCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardDeleteCardHeaders) SetXAcsDingtalkAccessToken(v string) *CardDeleteCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardDeleteCardRequest struct {
	CardBizCode *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardBizId   *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardId      *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CardDeleteCardRequest) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteCardRequest) GoString() string {
	return s.String()
}

func (s *CardDeleteCardRequest) SetCardBizCode(v string) *CardDeleteCardRequest {
	s.CardBizCode = &v
	return s
}

func (s *CardDeleteCardRequest) SetCardBizId(v string) *CardDeleteCardRequest {
	s.CardBizId = &v
	return s
}

func (s *CardDeleteCardRequest) SetCardId(v int64) *CardDeleteCardRequest {
	s.CardId = &v
	return s
}

func (s *CardDeleteCardRequest) SetSourceType(v string) *CardDeleteCardRequest {
	s.SourceType = &v
	return s
}

func (s *CardDeleteCardRequest) SetUserId(v string) *CardDeleteCardRequest {
	s.UserId = &v
	return s
}

type CardDeleteCardResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardDeleteCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteCardResponseBody) GoString() string {
	return s.String()
}

func (s *CardDeleteCardResponseBody) SetResult(v bool) *CardDeleteCardResponseBody {
	s.Result = &v
	return s
}

func (s *CardDeleteCardResponseBody) SetSuccess(v bool) *CardDeleteCardResponseBody {
	s.Success = &v
	return s
}

type CardDeleteCardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardDeleteCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardDeleteCardResponse) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteCardResponse) GoString() string {
	return s.String()
}

func (s *CardDeleteCardResponse) SetHeaders(v map[string]*string) *CardDeleteCardResponse {
	s.Headers = v
	return s
}

func (s *CardDeleteCardResponse) SetStatusCode(v int32) *CardDeleteCardResponse {
	s.StatusCode = &v
	return s
}

func (s *CardDeleteCardResponse) SetBody(v *CardDeleteCardResponseBody) *CardDeleteCardResponse {
	s.Body = v
	return s
}

type CardEndCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardEndCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardEndCardHeaders) GoString() string {
	return s.String()
}

func (s *CardEndCardHeaders) SetCommonHeaders(v map[string]*string) *CardEndCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardEndCardHeaders) SetXAcsDingtalkAccessToken(v string) *CardEndCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardEndCardRequest struct {
	CardBizCode *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardBizId   *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardId      *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CardEndCardRequest) String() string {
	return tea.Prettify(s)
}

func (s CardEndCardRequest) GoString() string {
	return s.String()
}

func (s *CardEndCardRequest) SetCardBizCode(v string) *CardEndCardRequest {
	s.CardBizCode = &v
	return s
}

func (s *CardEndCardRequest) SetCardBizId(v string) *CardEndCardRequest {
	s.CardBizId = &v
	return s
}

func (s *CardEndCardRequest) SetCardId(v int64) *CardEndCardRequest {
	s.CardId = &v
	return s
}

func (s *CardEndCardRequest) SetSourceType(v string) *CardEndCardRequest {
	s.SourceType = &v
	return s
}

func (s *CardEndCardRequest) SetUserId(v string) *CardEndCardRequest {
	s.UserId = &v
	return s
}

type CardEndCardResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardEndCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardEndCardResponseBody) GoString() string {
	return s.String()
}

func (s *CardEndCardResponseBody) SetResult(v bool) *CardEndCardResponseBody {
	s.Result = &v
	return s
}

func (s *CardEndCardResponseBody) SetSuccess(v bool) *CardEndCardResponseBody {
	s.Success = &v
	return s
}

type CardEndCardResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardEndCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardEndCardResponse) String() string {
	return tea.Prettify(s)
}

func (s CardEndCardResponse) GoString() string {
	return s.String()
}

func (s *CardEndCardResponse) SetHeaders(v map[string]*string) *CardEndCardResponse {
	s.Headers = v
	return s
}

func (s *CardEndCardResponse) SetStatusCode(v int32) *CardEndCardResponse {
	s.StatusCode = &v
	return s
}

func (s *CardEndCardResponse) SetBody(v *CardEndCardResponseBody) *CardEndCardResponse {
	s.Body = v
	return s
}

type CardGetCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardGetCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardHeaders) GoString() string {
	return s.String()
}

func (s *CardGetCardHeaders) SetCommonHeaders(v map[string]*string) *CardGetCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardGetCardHeaders) SetXAcsDingtalkAccessToken(v string) *CardGetCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardGetCardRequest struct {
	CardId     *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s CardGetCardRequest) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardRequest) GoString() string {
	return s.String()
}

func (s *CardGetCardRequest) SetCardId(v int64) *CardGetCardRequest {
	s.CardId = &v
	return s
}

func (s *CardGetCardRequest) SetSourceType(v string) *CardGetCardRequest {
	s.SourceType = &v
	return s
}

type CardGetCardResponseBody struct {
	Result  *CardGetCardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardGetCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardResponseBody) GoString() string {
	return s.String()
}

func (s *CardGetCardResponseBody) SetResult(v *CardGetCardResponseBodyResult) *CardGetCardResponseBody {
	s.Result = v
	return s
}

func (s *CardGetCardResponseBody) SetSuccess(v bool) *CardGetCardResponseBody {
	s.Success = &v
	return s
}

type CardGetCardResponseBodyResult struct {
	Attr                            *string   `json:"attr,omitempty" xml:"attr,omitempty"`
	CardBizCode                     *string   `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardBizId                       *string   `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardContentCount                *int32    `json:"cardContentCount,omitempty" xml:"cardContentCount,omitempty"`
	CardCycle                       *int32    `json:"cardCycle,omitempty" xml:"cardCycle,omitempty"`
	CardFrequency                   []*int32  `json:"cardFrequency,omitempty" xml:"cardFrequency,omitempty" type:"Repeated"`
	CardId                          *int64    `json:"cardId,omitempty" xml:"cardId,omitempty"`
	CardStatus                      *int32    `json:"cardStatus,omitempty" xml:"cardStatus,omitempty"`
	CardUrl                         *string   `json:"cardUrl,omitempty" xml:"cardUrl,omitempty"`
	CategoryContentTag              *string   `json:"categoryContentTag,omitempty" xml:"categoryContentTag,omitempty"`
	CategoryCoverImageUrl           *string   `json:"categoryCoverImageUrl,omitempty" xml:"categoryCoverImageUrl,omitempty"`
	CategoryCreateCardSmallImageUrl *string   `json:"categoryCreateCardSmallImageUrl,omitempty" xml:"categoryCreateCardSmallImageUrl,omitempty"`
	CategoryListSmallImageUrl       *string   `json:"categoryListSmallImageUrl,omitempty" xml:"categoryListSmallImageUrl,omitempty"`
	CategoryName                    *string   `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	ClassIds                        []*string `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	ClassNames                      []*string `json:"classNames,omitempty" xml:"classNames,omitempty" type:"Repeated"`
	Content                         *string   `json:"content,omitempty" xml:"content,omitempty"`
	CorpId                          *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	EffectTime                      *string   `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	Finished                        *bool     `json:"finished,omitempty" xml:"finished,omitempty"`
	Media                           *string   `json:"media,omitempty" xml:"media,omitempty"`
	OptEndTime                      *string   `json:"optEndTime,omitempty" xml:"optEndTime,omitempty"`
	OptEndUserId                    *string   `json:"optEndUserId,omitempty" xml:"optEndUserId,omitempty"`
	OptEndUserName                  *string   `json:"optEndUserName,omitempty" xml:"optEndUserName,omitempty"`
	RemindNotPunchCardHour          *int32    `json:"remindNotPunchCardHour,omitempty" xml:"remindNotPunchCardHour,omitempty"`
	RemindNotPunchCardMinute        *int32    `json:"remindNotPunchCardMinute,omitempty" xml:"remindNotPunchCardMinute,omitempty"`
	SendTime                        *string   `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	SourceType                      *string   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	StartTime                       *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status                          *int32    `json:"status,omitempty" xml:"status,omitempty"`
	SystemTime                      *int64    `json:"systemTime,omitempty" xml:"systemTime,omitempty"`
	TeacherId                       *string   `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
	TeacherName                     *string   `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	TemplateCoverImageUrl           *string   `json:"templateCoverImageUrl,omitempty" xml:"templateCoverImageUrl,omitempty"`
	Title                           *string   `json:"title,omitempty" xml:"title,omitempty"`
	Type                            *int32    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CardGetCardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardGetCardResponseBodyResult) SetAttr(v string) *CardGetCardResponseBodyResult {
	s.Attr = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardBizCode(v string) *CardGetCardResponseBodyResult {
	s.CardBizCode = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardBizId(v string) *CardGetCardResponseBodyResult {
	s.CardBizId = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardContentCount(v int32) *CardGetCardResponseBodyResult {
	s.CardContentCount = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardCycle(v int32) *CardGetCardResponseBodyResult {
	s.CardCycle = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardFrequency(v []*int32) *CardGetCardResponseBodyResult {
	s.CardFrequency = v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardId(v int64) *CardGetCardResponseBodyResult {
	s.CardId = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardStatus(v int32) *CardGetCardResponseBodyResult {
	s.CardStatus = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCardUrl(v string) *CardGetCardResponseBodyResult {
	s.CardUrl = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCategoryContentTag(v string) *CardGetCardResponseBodyResult {
	s.CategoryContentTag = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCategoryCoverImageUrl(v string) *CardGetCardResponseBodyResult {
	s.CategoryCoverImageUrl = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCategoryCreateCardSmallImageUrl(v string) *CardGetCardResponseBodyResult {
	s.CategoryCreateCardSmallImageUrl = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCategoryListSmallImageUrl(v string) *CardGetCardResponseBodyResult {
	s.CategoryListSmallImageUrl = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCategoryName(v string) *CardGetCardResponseBodyResult {
	s.CategoryName = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetClassIds(v []*string) *CardGetCardResponseBodyResult {
	s.ClassIds = v
	return s
}

func (s *CardGetCardResponseBodyResult) SetClassNames(v []*string) *CardGetCardResponseBodyResult {
	s.ClassNames = v
	return s
}

func (s *CardGetCardResponseBodyResult) SetContent(v string) *CardGetCardResponseBodyResult {
	s.Content = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetCorpId(v string) *CardGetCardResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetEffectTime(v string) *CardGetCardResponseBodyResult {
	s.EffectTime = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetFinished(v bool) *CardGetCardResponseBodyResult {
	s.Finished = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetMedia(v string) *CardGetCardResponseBodyResult {
	s.Media = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetOptEndTime(v string) *CardGetCardResponseBodyResult {
	s.OptEndTime = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetOptEndUserId(v string) *CardGetCardResponseBodyResult {
	s.OptEndUserId = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetOptEndUserName(v string) *CardGetCardResponseBodyResult {
	s.OptEndUserName = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetRemindNotPunchCardHour(v int32) *CardGetCardResponseBodyResult {
	s.RemindNotPunchCardHour = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetRemindNotPunchCardMinute(v int32) *CardGetCardResponseBodyResult {
	s.RemindNotPunchCardMinute = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetSendTime(v string) *CardGetCardResponseBodyResult {
	s.SendTime = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetSourceType(v string) *CardGetCardResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetStartTime(v string) *CardGetCardResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetStatus(v int32) *CardGetCardResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetSystemTime(v int64) *CardGetCardResponseBodyResult {
	s.SystemTime = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetTeacherId(v string) *CardGetCardResponseBodyResult {
	s.TeacherId = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetTeacherName(v string) *CardGetCardResponseBodyResult {
	s.TeacherName = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetTemplateCoverImageUrl(v string) *CardGetCardResponseBodyResult {
	s.TemplateCoverImageUrl = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetTitle(v string) *CardGetCardResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CardGetCardResponseBodyResult) SetType(v int32) *CardGetCardResponseBodyResult {
	s.Type = &v
	return s
}

type CardGetCardResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardGetCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardGetCardResponse) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardResponse) GoString() string {
	return s.String()
}

func (s *CardGetCardResponse) SetHeaders(v map[string]*string) *CardGetCardResponse {
	s.Headers = v
	return s
}

func (s *CardGetCardResponse) SetStatusCode(v int32) *CardGetCardResponse {
	s.StatusCode = &v
	return s
}

func (s *CardGetCardResponse) SetBody(v *CardGetCardResponseBody) *CardGetCardResponse {
	s.Body = v
	return s
}

type CardGetCardFinishProgressHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardGetCardFinishProgressHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressHeaders) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressHeaders) SetCommonHeaders(v map[string]*string) *CardGetCardFinishProgressHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardGetCardFinishProgressHeaders) SetXAcsDingtalkAccessToken(v string) *CardGetCardFinishProgressHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardGetCardFinishProgressRequest struct {
	CardBizCode *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardBizId   *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardId      *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	SourceType  *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	StudentId   *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CardGetCardFinishProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressRequest) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressRequest) SetCardBizCode(v string) *CardGetCardFinishProgressRequest {
	s.CardBizCode = &v
	return s
}

func (s *CardGetCardFinishProgressRequest) SetCardBizId(v string) *CardGetCardFinishProgressRequest {
	s.CardBizId = &v
	return s
}

func (s *CardGetCardFinishProgressRequest) SetCardId(v int64) *CardGetCardFinishProgressRequest {
	s.CardId = &v
	return s
}

func (s *CardGetCardFinishProgressRequest) SetSourceType(v string) *CardGetCardFinishProgressRequest {
	s.SourceType = &v
	return s
}

func (s *CardGetCardFinishProgressRequest) SetStudentId(v string) *CardGetCardFinishProgressRequest {
	s.StudentId = &v
	return s
}

func (s *CardGetCardFinishProgressRequest) SetUserId(v string) *CardGetCardFinishProgressRequest {
	s.UserId = &v
	return s
}

type CardGetCardFinishProgressResponseBody struct {
	Result  *CardGetCardFinishProgressResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardGetCardFinishProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponseBody) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponseBody) SetResult(v *CardGetCardFinishProgressResponseBodyResult) *CardGetCardFinishProgressResponseBody {
	s.Result = v
	return s
}

func (s *CardGetCardFinishProgressResponseBody) SetSuccess(v bool) *CardGetCardFinishProgressResponseBody {
	s.Success = &v
	return s
}

type CardGetCardFinishProgressResponseBodyResult struct {
	ClassStatistics     []*CardGetCardFinishProgressResponseBodyResultClassStatistics     `json:"classStatistics,omitempty" xml:"classStatistics,omitempty" type:"Repeated"`
	PatriarchStatistics []*CardGetCardFinishProgressResponseBodyResultPatriarchStatistics `json:"patriarchStatistics,omitempty" xml:"patriarchStatistics,omitempty" type:"Repeated"`
	StudentNameList     []*string                                                         `json:"studentNameList,omitempty" xml:"studentNameList,omitempty" type:"Repeated"`
}

func (s CardGetCardFinishProgressResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponseBodyResult) SetClassStatistics(v []*CardGetCardFinishProgressResponseBodyResultClassStatistics) *CardGetCardFinishProgressResponseBodyResult {
	s.ClassStatistics = v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResult) SetPatriarchStatistics(v []*CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) *CardGetCardFinishProgressResponseBodyResult {
	s.PatriarchStatistics = v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResult) SetStudentNameList(v []*string) *CardGetCardFinishProgressResponseBodyResult {
	s.StudentNameList = v
	return s
}

type CardGetCardFinishProgressResponseBodyResultClassStatistics struct {
	CardBizId   *string                                                              `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardBizName *string                                                              `json:"cardBizName,omitempty" xml:"cardBizName,omitempty"`
	ClassId     *string                                                              `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName   *string                                                              `json:"className,omitempty" xml:"className,omitempty"`
	Process     []*CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess `json:"process,omitempty" xml:"process,omitempty" type:"Repeated"`
}

func (s CardGetCardFinishProgressResponseBodyResultClassStatistics) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponseBodyResultClassStatistics) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatistics) SetCardBizId(v string) *CardGetCardFinishProgressResponseBodyResultClassStatistics {
	s.CardBizId = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatistics) SetCardBizName(v string) *CardGetCardFinishProgressResponseBodyResultClassStatistics {
	s.CardBizName = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatistics) SetClassId(v string) *CardGetCardFinishProgressResponseBodyResultClassStatistics {
	s.ClassId = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatistics) SetClassName(v string) *CardGetCardFinishProgressResponseBodyResultClassStatistics {
	s.ClassName = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatistics) SetProcess(v []*CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) *CardGetCardFinishProgressResponseBodyResultClassStatistics {
	s.Process = v
	return s
}

type CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess struct {
	Date                  *string `json:"date,omitempty" xml:"date,omitempty"`
	FinishedStudentsNum   *int64  `json:"finishedStudentsNum,omitempty" xml:"finishedStudentsNum,omitempty"`
	NeedFinishStudentsNum *int64  `json:"needFinishStudentsNum,omitempty" xml:"needFinishStudentsNum,omitempty"`
	TaskCode              *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
	Today                 *string `json:"today,omitempty" xml:"today,omitempty"`
}

func (s CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) SetDate(v string) *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess {
	s.Date = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) SetFinishedStudentsNum(v int64) *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess {
	s.FinishedStudentsNum = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) SetNeedFinishStudentsNum(v int64) *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess {
	s.NeedFinishStudentsNum = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) SetTaskCode(v string) *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess {
	s.TaskCode = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess) SetToday(v string) *CardGetCardFinishProgressResponseBodyResultClassStatisticsProcess {
	s.Today = &v
	return s
}

type CardGetCardFinishProgressResponseBodyResultPatriarchStatistics struct {
	CardTaskCode            *string `json:"cardTaskCode,omitempty" xml:"cardTaskCode,omitempty"`
	Date                    *string `json:"date,omitempty" xml:"date,omitempty"`
	IsFinished              *bool   `json:"isFinished,omitempty" xml:"isFinished,omitempty"`
	IsFinishedByReissueCard *bool   `json:"isFinishedByReissueCard,omitempty" xml:"isFinishedByReissueCard,omitempty"`
	IsLastDay               *bool   `json:"isLastDay,omitempty" xml:"isLastDay,omitempty"`
	ReissueCard             *bool   `json:"reissueCard,omitempty" xml:"reissueCard,omitempty"`
	StudentId               *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName             *string `json:"studentName,omitempty" xml:"studentName,omitempty"`
	Today                   *string `json:"today,omitempty" xml:"today,omitempty"`
	UserSubTaskId           *int64  `json:"userSubTaskId,omitempty" xml:"userSubTaskId,omitempty"`
}

func (s CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetCardTaskCode(v string) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.CardTaskCode = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetDate(v string) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.Date = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetIsFinished(v bool) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.IsFinished = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetIsFinishedByReissueCard(v bool) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.IsFinishedByReissueCard = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetIsLastDay(v bool) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.IsLastDay = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetReissueCard(v bool) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.ReissueCard = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetStudentId(v string) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.StudentId = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetStudentName(v string) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.StudentName = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetToday(v string) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.Today = &v
	return s
}

func (s *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics) SetUserSubTaskId(v int64) *CardGetCardFinishProgressResponseBodyResultPatriarchStatistics {
	s.UserSubTaskId = &v
	return s
}

type CardGetCardFinishProgressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardGetCardFinishProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardGetCardFinishProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s CardGetCardFinishProgressResponse) GoString() string {
	return s.String()
}

func (s *CardGetCardFinishProgressResponse) SetHeaders(v map[string]*string) *CardGetCardFinishProgressResponse {
	s.Headers = v
	return s
}

func (s *CardGetCardFinishProgressResponse) SetStatusCode(v int32) *CardGetCardFinishProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *CardGetCardFinishProgressResponse) SetBody(v *CardGetCardFinishProgressResponseBody) *CardGetCardFinishProgressResponse {
	s.Body = v
	return s
}

type CardQueryCardFeedsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardQueryCardFeedsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsHeaders) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsHeaders) SetCommonHeaders(v map[string]*string) *CardQueryCardFeedsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardQueryCardFeedsHeaders) SetXAcsDingtalkAccessToken(v string) *CardQueryCardFeedsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardQueryCardFeedsRequest struct {
	BizType           *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CardBizCode       *string `json:"cardBizCode,omitempty" xml:"cardBizCode,omitempty"`
	CardBizId         *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardId            *int64  `json:"cardId,omitempty" xml:"cardId,omitempty"`
	Count             *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Cursor            *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	FeedType          *int32  `json:"feedType,omitempty" xml:"feedType,omitempty"`
	NeedFinishProcess *bool   `json:"needFinishProcess,omitempty" xml:"needFinishProcess,omitempty"`
	SourceType        *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	StudentId         *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	SubBizId          *string `json:"subBizId,omitempty" xml:"subBizId,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CardQueryCardFeedsRequest) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsRequest) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsRequest) SetBizType(v int32) *CardQueryCardFeedsRequest {
	s.BizType = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetCardBizCode(v string) *CardQueryCardFeedsRequest {
	s.CardBizCode = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetCardBizId(v string) *CardQueryCardFeedsRequest {
	s.CardBizId = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetCardId(v int64) *CardQueryCardFeedsRequest {
	s.CardId = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetCount(v int32) *CardQueryCardFeedsRequest {
	s.Count = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetCursor(v int64) *CardQueryCardFeedsRequest {
	s.Cursor = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetFeedType(v int32) *CardQueryCardFeedsRequest {
	s.FeedType = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetNeedFinishProcess(v bool) *CardQueryCardFeedsRequest {
	s.NeedFinishProcess = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetSourceType(v string) *CardQueryCardFeedsRequest {
	s.SourceType = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetStudentId(v string) *CardQueryCardFeedsRequest {
	s.StudentId = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetSubBizId(v string) *CardQueryCardFeedsRequest {
	s.SubBizId = &v
	return s
}

func (s *CardQueryCardFeedsRequest) SetUserId(v string) *CardQueryCardFeedsRequest {
	s.UserId = &v
	return s
}

type CardQueryCardFeedsResponseBody struct {
	Result  *CardQueryCardFeedsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CardQueryCardFeedsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponseBody) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponseBody) SetResult(v *CardQueryCardFeedsResponseBodyResult) *CardQueryCardFeedsResponseBody {
	s.Result = v
	return s
}

func (s *CardQueryCardFeedsResponseBody) SetSuccess(v bool) *CardQueryCardFeedsResponseBody {
	s.Success = &v
	return s
}

type CardQueryCardFeedsResponseBodyResult struct {
	HasMore *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Posts   []*CardQueryCardFeedsResponseBodyResultPosts `json:"posts,omitempty" xml:"posts,omitempty" type:"Repeated"`
}

func (s CardQueryCardFeedsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponseBodyResult) SetHasMore(v bool) *CardQueryCardFeedsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResult) SetPosts(v []*CardQueryCardFeedsResponseBodyResultPosts) *CardQueryCardFeedsResponseBodyResult {
	s.Posts = v
	return s
}

type CardQueryCardFeedsResponseBodyResultPosts struct {
	Author   *CardQueryCardFeedsResponseBodyResultPostsAuthor  `json:"author,omitempty" xml:"author,omitempty" type:"Struct"`
	BizType  *int32                                            `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Content  *CardQueryCardFeedsResponseBodyResultPostsContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	CreateAt *int64                                            `json:"createAt,omitempty" xml:"createAt,omitempty"`
	FeedType *int32                                            `json:"feedType,omitempty" xml:"feedType,omitempty"`
	PostId   *int64                                            `json:"postId,omitempty" xml:"postId,omitempty"`
	Status   *int32                                            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CardQueryCardFeedsResponseBodyResultPosts) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponseBodyResultPosts) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetAuthor(v *CardQueryCardFeedsResponseBodyResultPostsAuthor) *CardQueryCardFeedsResponseBodyResultPosts {
	s.Author = v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetBizType(v int32) *CardQueryCardFeedsResponseBodyResultPosts {
	s.BizType = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetContent(v *CardQueryCardFeedsResponseBodyResultPostsContent) *CardQueryCardFeedsResponseBodyResultPosts {
	s.Content = v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetCreateAt(v int64) *CardQueryCardFeedsResponseBodyResultPosts {
	s.CreateAt = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetFeedType(v int32) *CardQueryCardFeedsResponseBodyResultPosts {
	s.FeedType = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetPostId(v int64) *CardQueryCardFeedsResponseBodyResultPosts {
	s.PostId = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPosts) SetStatus(v int32) *CardQueryCardFeedsResponseBodyResultPosts {
	s.Status = &v
	return s
}

type CardQueryCardFeedsResponseBodyResultPostsAuthor struct {
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserRole *string `json:"userRole,omitempty" xml:"userRole,omitempty"`
}

func (s CardQueryCardFeedsResponseBodyResultPostsAuthor) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponseBodyResultPostsAuthor) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponseBodyResultPostsAuthor) SetShowName(v string) *CardQueryCardFeedsResponseBodyResultPostsAuthor {
	s.ShowName = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPostsAuthor) SetUserId(v string) *CardQueryCardFeedsResponseBodyResultPostsAuthor {
	s.UserId = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPostsAuthor) SetUserRole(v string) *CardQueryCardFeedsResponseBodyResultPostsAuthor {
	s.UserRole = &v
	return s
}

type CardQueryCardFeedsResponseBodyResultPostsContent struct {
	ContentType *int32  `json:"contentType,omitempty" xml:"contentType,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CardQueryCardFeedsResponseBodyResultPostsContent) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponseBodyResultPostsContent) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponseBodyResultPostsContent) SetContentType(v int32) *CardQueryCardFeedsResponseBodyResultPostsContent {
	s.ContentType = &v
	return s
}

func (s *CardQueryCardFeedsResponseBodyResultPostsContent) SetText(v string) *CardQueryCardFeedsResponseBodyResultPostsContent {
	s.Text = &v
	return s
}

type CardQueryCardFeedsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardQueryCardFeedsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardQueryCardFeedsResponse) String() string {
	return tea.Prettify(s)
}

func (s CardQueryCardFeedsResponse) GoString() string {
	return s.String()
}

func (s *CardQueryCardFeedsResponse) SetHeaders(v map[string]*string) *CardQueryCardFeedsResponse {
	s.Headers = v
	return s
}

func (s *CardQueryCardFeedsResponse) SetStatusCode(v int32) *CardQueryCardFeedsResponse {
	s.StatusCode = &v
	return s
}

func (s *CardQueryCardFeedsResponse) SetBody(v *CardQueryCardFeedsResponseBody) *CardQueryCardFeedsResponse {
	s.Body = v
	return s
}

type CheckRestrictionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckRestrictionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckRestrictionHeaders) GoString() string {
	return s.String()
}

func (s *CheckRestrictionHeaders) SetCommonHeaders(v map[string]*string) *CheckRestrictionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckRestrictionHeaders) SetXAcsDingtalkAccessToken(v string) *CheckRestrictionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckRestrictionRequest struct {
	ActualAmount *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	FaceId       *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Scene        *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	Sn           *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CheckRestrictionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckRestrictionRequest) GoString() string {
	return s.String()
}

func (s *CheckRestrictionRequest) SetActualAmount(v int64) *CheckRestrictionRequest {
	s.ActualAmount = &v
	return s
}

func (s *CheckRestrictionRequest) SetFaceId(v string) *CheckRestrictionRequest {
	s.FaceId = &v
	return s
}

func (s *CheckRestrictionRequest) SetScene(v int64) *CheckRestrictionRequest {
	s.Scene = &v
	return s
}

func (s *CheckRestrictionRequest) SetSn(v string) *CheckRestrictionRequest {
	s.Sn = &v
	return s
}

func (s *CheckRestrictionRequest) SetUserId(v string) *CheckRestrictionRequest {
	s.UserId = &v
	return s
}

type CheckRestrictionResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckRestrictionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckRestrictionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRestrictionResponseBody) SetSuccess(v bool) *CheckRestrictionResponseBody {
	s.Success = &v
	return s
}

type CheckRestrictionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckRestrictionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckRestrictionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckRestrictionResponse) GoString() string {
	return s.String()
}

func (s *CheckRestrictionResponse) SetHeaders(v map[string]*string) *CheckRestrictionResponse {
	s.Headers = v
	return s
}

func (s *CheckRestrictionResponse) SetStatusCode(v int32) *CheckRestrictionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRestrictionResponse) SetBody(v *CheckRestrictionResponseBody) *CheckRestrictionResponse {
	s.Body = v
	return s
}

type ConsumePointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumePointHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointHeaders) GoString() string {
	return s.String()
}

func (s *ConsumePointHeaders) SetCommonHeaders(v map[string]*string) *ConsumePointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumePointHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumePointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumePointRequest struct {
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	BizId       *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
}

func (s ConsumePointRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointRequest) GoString() string {
	return s.String()
}

func (s *ConsumePointRequest) SetAmount(v int64) *ConsumePointRequest {
	s.Amount = &v
	return s
}

func (s *ConsumePointRequest) SetBizId(v string) *ConsumePointRequest {
	s.BizId = &v
	return s
}

func (s *ConsumePointRequest) SetDescription(v string) *ConsumePointRequest {
	s.Description = &v
	return s
}

func (s *ConsumePointRequest) SetProductCode(v string) *ConsumePointRequest {
	s.ProductCode = &v
	return s
}

type ConsumePointResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConsumePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumePointResponseBody) SetResult(v bool) *ConsumePointResponseBody {
	s.Result = &v
	return s
}

func (s *ConsumePointResponseBody) SetSuccess(v bool) *ConsumePointResponseBody {
	s.Success = &v
	return s
}

type ConsumePointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConsumePointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConsumePointResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumePointResponse) GoString() string {
	return s.String()
}

func (s *ConsumePointResponse) SetHeaders(v map[string]*string) *ConsumePointResponse {
	s.Headers = v
	return s
}

func (s *ConsumePointResponse) SetStatusCode(v int32) *ConsumePointResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumePointResponse) SetBody(v *ConsumePointResponseBody) *ConsumePointResponse {
	s.Body = v
	return s
}

type CourseSchedulingComplimentNoticeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CourseSchedulingComplimentNoticeHeaders) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeHeaders) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeHeaders) SetCommonHeaders(v map[string]*string) *CourseSchedulingComplimentNoticeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CourseSchedulingComplimentNoticeHeaders) SetXAcsDingtalkAccessToken(v string) *CourseSchedulingComplimentNoticeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CourseSchedulingComplimentNoticeRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CourseSchedulingComplimentNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeRequest) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeRequest) SetOpUserId(v string) *CourseSchedulingComplimentNoticeRequest {
	s.OpUserId = &v
	return s
}

type CourseSchedulingComplimentNoticeResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CourseSchedulingComplimentNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeResponseBody) SetResult(v bool) *CourseSchedulingComplimentNoticeResponseBody {
	s.Result = &v
	return s
}

type CourseSchedulingComplimentNoticeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CourseSchedulingComplimentNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CourseSchedulingComplimentNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s CourseSchedulingComplimentNoticeResponse) GoString() string {
	return s.String()
}

func (s *CourseSchedulingComplimentNoticeResponse) SetHeaders(v map[string]*string) *CourseSchedulingComplimentNoticeResponse {
	s.Headers = v
	return s
}

func (s *CourseSchedulingComplimentNoticeResponse) SetStatusCode(v int32) *CourseSchedulingComplimentNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *CourseSchedulingComplimentNoticeResponse) SetBody(v *CourseSchedulingComplimentNoticeResponseBody) *CourseSchedulingComplimentNoticeResponse {
	s.Body = v
	return s
}

type CreateAppOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAppOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateAppOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateAppOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAppOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAppOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAppOrderRequest struct {
	ActualAmount    *int64                             `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId     *string                            `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	BizCode         *int32                             `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	DetailList      []*CreateAppOrderRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
	LabelAmount     *int64                             `json:"labelAmount,omitempty" xml:"labelAmount,omitempty"`
	MerchantId      *string                            `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string                            `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OuterUserId     *string                            `json:"outerUserId,omitempty" xml:"outerUserId,omitempty"`
	Signature       *string                            `json:"signature,omitempty" xml:"signature,omitempty"`
	Subject         *string                            `json:"subject,omitempty" xml:"subject,omitempty"`
	Timestamp       *int64                             `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s CreateAppOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAppOrderRequest) SetActualAmount(v int64) *CreateAppOrderRequest {
	s.ActualAmount = &v
	return s
}

func (s *CreateAppOrderRequest) SetAlipayAppId(v string) *CreateAppOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *CreateAppOrderRequest) SetBizCode(v int32) *CreateAppOrderRequest {
	s.BizCode = &v
	return s
}

func (s *CreateAppOrderRequest) SetDetailList(v []*CreateAppOrderRequestDetailList) *CreateAppOrderRequest {
	s.DetailList = v
	return s
}

func (s *CreateAppOrderRequest) SetLabelAmount(v int64) *CreateAppOrderRequest {
	s.LabelAmount = &v
	return s
}

func (s *CreateAppOrderRequest) SetMerchantId(v string) *CreateAppOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *CreateAppOrderRequest) SetMerchantOrderNo(v string) *CreateAppOrderRequest {
	s.MerchantOrderNo = &v
	return s
}

func (s *CreateAppOrderRequest) SetOuterUserId(v string) *CreateAppOrderRequest {
	s.OuterUserId = &v
	return s
}

func (s *CreateAppOrderRequest) SetSignature(v string) *CreateAppOrderRequest {
	s.Signature = &v
	return s
}

func (s *CreateAppOrderRequest) SetSubject(v string) *CreateAppOrderRequest {
	s.Subject = &v
	return s
}

func (s *CreateAppOrderRequest) SetTimestamp(v int64) *CreateAppOrderRequest {
	s.Timestamp = &v
	return s
}

type CreateAppOrderRequestDetailList struct {
	GoodsId       *string `json:"goodsId,omitempty" xml:"goodsId,omitempty"`
	GoodsName     *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	GoodsPrice    *int64  `json:"goodsPrice,omitempty" xml:"goodsPrice,omitempty"`
	GoodsQuantity *int32  `json:"goodsQuantity,omitempty" xml:"goodsQuantity,omitempty"`
}

func (s CreateAppOrderRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOrderRequestDetailList) GoString() string {
	return s.String()
}

func (s *CreateAppOrderRequestDetailList) SetGoodsId(v string) *CreateAppOrderRequestDetailList {
	s.GoodsId = &v
	return s
}

func (s *CreateAppOrderRequestDetailList) SetGoodsName(v string) *CreateAppOrderRequestDetailList {
	s.GoodsName = &v
	return s
}

func (s *CreateAppOrderRequestDetailList) SetGoodsPrice(v int64) *CreateAppOrderRequestDetailList {
	s.GoodsPrice = &v
	return s
}

func (s *CreateAppOrderRequestDetailList) SetGoodsQuantity(v int32) *CreateAppOrderRequestDetailList {
	s.GoodsQuantity = &v
	return s
}

type CreateAppOrderResponseBody struct {
	ActualAmount    *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId     *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	Body            *string `json:"body,omitempty" xml:"body,omitempty"`
	MerchantId      *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s CreateAppOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppOrderResponseBody) SetActualAmount(v int64) *CreateAppOrderResponseBody {
	s.ActualAmount = &v
	return s
}

func (s *CreateAppOrderResponseBody) SetAlipayAppId(v string) *CreateAppOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *CreateAppOrderResponseBody) SetBody(v string) *CreateAppOrderResponseBody {
	s.Body = &v
	return s
}

func (s *CreateAppOrderResponseBody) SetMerchantId(v string) *CreateAppOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *CreateAppOrderResponseBody) SetMerchantOrderNo(v string) *CreateAppOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *CreateAppOrderResponseBody) SetOrderNo(v string) *CreateAppOrderResponseBody {
	s.OrderNo = &v
	return s
}

type CreateAppOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAppOrderResponse) SetHeaders(v map[string]*string) *CreateAppOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAppOrderResponse) SetStatusCode(v int32) *CreateAppOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppOrderResponse) SetBody(v *CreateAppOrderResponseBody) *CreateAppOrderResponse {
	s.Body = v
	return s
}

type CreateCustomClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomClassHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomClassHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomClassRequest struct {
	CustomClass *CreateCustomClassRequestCustomClass `json:"customClass,omitempty" xml:"customClass,omitempty" type:"Struct"`
	Operator    *string                              `json:"operator,omitempty" xml:"operator,omitempty"`
	SuperId     *int64                               `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s CreateCustomClassRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomClassRequest) SetCustomClass(v *CreateCustomClassRequestCustomClass) *CreateCustomClassRequest {
	s.CustomClass = v
	return s
}

func (s *CreateCustomClassRequest) SetOperator(v string) *CreateCustomClassRequest {
	s.Operator = &v
	return s
}

func (s *CreateCustomClassRequest) SetSuperId(v int64) *CreateCustomClassRequest {
	s.SuperId = &v
	return s
}

type CreateCustomClassRequestCustomClass struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCustomClassRequestCustomClass) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassRequestCustomClass) GoString() string {
	return s.String()
}

func (s *CreateCustomClassRequestCustomClass) SetName(v string) *CreateCustomClassRequestCustomClass {
	s.Name = &v
	return s
}

type CreateCustomClassResponseBody struct {
	Result  *CreateCustomClassResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCustomClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponseBody) SetResult(v *CreateCustomClassResponseBodyResult) *CreateCustomClassResponseBody {
	s.Result = v
	return s
}

func (s *CreateCustomClassResponseBody) SetSuccess(v bool) *CreateCustomClassResponseBody {
	s.Success = &v
	return s
}

type CreateCustomClassResponseBodyResult struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CreateCustomClassResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponseBodyResult) SetDeptId(v int64) *CreateCustomClassResponseBodyResult {
	s.DeptId = &v
	return s
}

type CreateCustomClassResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomClassResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomClassResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomClassResponse) SetHeaders(v map[string]*string) *CreateCustomClassResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomClassResponse) SetStatusCode(v int32) *CreateCustomClassResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomClassResponse) SetBody(v *CreateCustomClassResponseBody) *CreateCustomClassResponse {
	s.Body = v
	return s
}

type CreateCustomDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomDeptRequest struct {
	CustomDept *CreateCustomDeptRequestCustomDept `json:"customDept,omitempty" xml:"customDept,omitempty" type:"Struct"`
	Operator   *string                            `json:"operator,omitempty" xml:"operator,omitempty"`
	SuperId    *int64                             `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s CreateCustomDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptRequest) SetCustomDept(v *CreateCustomDeptRequestCustomDept) *CreateCustomDeptRequest {
	s.CustomDept = v
	return s
}

func (s *CreateCustomDeptRequest) SetOperator(v string) *CreateCustomDeptRequest {
	s.Operator = &v
	return s
}

func (s *CreateCustomDeptRequest) SetSuperId(v int64) *CreateCustomDeptRequest {
	s.SuperId = &v
	return s
}

type CreateCustomDeptRequestCustomDept struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateCustomDeptRequestCustomDept) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptRequestCustomDept) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptRequestCustomDept) SetName(v string) *CreateCustomDeptRequestCustomDept {
	s.Name = &v
	return s
}

func (s *CreateCustomDeptRequestCustomDept) SetType(v string) *CreateCustomDeptRequestCustomDept {
	s.Type = &v
	return s
}

type CreateCustomDeptResponseBody struct {
	Result  *CreateCustomDeptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCustomDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponseBody) SetResult(v *CreateCustomDeptResponseBodyResult) *CreateCustomDeptResponseBody {
	s.Result = v
	return s
}

func (s *CreateCustomDeptResponseBody) SetSuccess(v bool) *CreateCustomDeptResponseBody {
	s.Success = &v
	return s
}

type CreateCustomDeptResponseBodyResult struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CreateCustomDeptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponseBodyResult) SetDeptId(v int64) *CreateCustomDeptResponseBodyResult {
	s.DeptId = &v
	return s
}

type CreateCustomDeptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDeptResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomDeptResponse) SetHeaders(v map[string]*string) *CreateCustomDeptResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomDeptResponse) SetStatusCode(v int32) *CreateCustomDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomDeptResponse) SetBody(v *CreateCustomDeptResponseBody) *CreateCustomDeptResponse {
	s.Body = v
	return s
}

type CreateEduAssetSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateEduAssetSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEduAssetSpaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateEduAssetSpaceHeaders) SetCommonHeaders(v map[string]*string) *CreateEduAssetSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEduAssetSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateEduAssetSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateEduAssetSpaceRequest struct {
	BizCode   *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	SpaceDesc *string `json:"spaceDesc,omitempty" xml:"spaceDesc,omitempty"`
	SpaceIcon *string `json:"spaceIcon,omitempty" xml:"spaceIcon,omitempty"`
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateEduAssetSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEduAssetSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateEduAssetSpaceRequest) SetBizCode(v string) *CreateEduAssetSpaceRequest {
	s.BizCode = &v
	return s
}

func (s *CreateEduAssetSpaceRequest) SetSpaceDesc(v string) *CreateEduAssetSpaceRequest {
	s.SpaceDesc = &v
	return s
}

func (s *CreateEduAssetSpaceRequest) SetSpaceIcon(v string) *CreateEduAssetSpaceRequest {
	s.SpaceIcon = &v
	return s
}

func (s *CreateEduAssetSpaceRequest) SetSpaceName(v string) *CreateEduAssetSpaceRequest {
	s.SpaceName = &v
	return s
}

func (s *CreateEduAssetSpaceRequest) SetUserId(v string) *CreateEduAssetSpaceRequest {
	s.UserId = &v
	return s
}

type CreateEduAssetSpaceResponseBody struct {
	CreateTimeMillis *int64  `json:"createTimeMillis,omitempty" xml:"createTimeMillis,omitempty"`
	ModifyTimeMillis *int64  `json:"modifyTimeMillis,omitempty" xml:"modifyTimeMillis,omitempty"`
	PermissionMode   *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	Quota            *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
	SpaceId          *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceName        *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType        *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	UsedQuota        *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s CreateEduAssetSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEduAssetSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEduAssetSpaceResponseBody) SetCreateTimeMillis(v int64) *CreateEduAssetSpaceResponseBody {
	s.CreateTimeMillis = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetModifyTimeMillis(v int64) *CreateEduAssetSpaceResponseBody {
	s.ModifyTimeMillis = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetPermissionMode(v string) *CreateEduAssetSpaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetQuota(v int64) *CreateEduAssetSpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetSpaceId(v string) *CreateEduAssetSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetSpaceName(v string) *CreateEduAssetSpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetSpaceType(v string) *CreateEduAssetSpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *CreateEduAssetSpaceResponseBody) SetUsedQuota(v int64) *CreateEduAssetSpaceResponseBody {
	s.UsedQuota = &v
	return s
}

type CreateEduAssetSpaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEduAssetSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEduAssetSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEduAssetSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateEduAssetSpaceResponse) SetHeaders(v map[string]*string) *CreateEduAssetSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateEduAssetSpaceResponse) SetStatusCode(v int32) *CreateEduAssetSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEduAssetSpaceResponse) SetBody(v *CreateEduAssetSpaceResponseBody) *CreateEduAssetSpaceResponse {
	s.Body = v
	return s
}

type CreateFulfilRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateFulfilRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFulfilRecordHeaders) GoString() string {
	return s.String()
}

func (s *CreateFulfilRecordHeaders) SetCommonHeaders(v map[string]*string) *CreateFulfilRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFulfilRecordHeaders) SetXAcsDingtalkAccessToken(v string) *CreateFulfilRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateFulfilRecordRequest struct {
	BizTime *int64  `json:"bizTime,omitempty" xml:"bizTime,omitempty"`
	ExtInfo *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	FaceId  *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Scene   *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	Sn      *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateFulfilRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFulfilRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateFulfilRecordRequest) SetBizTime(v int64) *CreateFulfilRecordRequest {
	s.BizTime = &v
	return s
}

func (s *CreateFulfilRecordRequest) SetExtInfo(v string) *CreateFulfilRecordRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateFulfilRecordRequest) SetFaceId(v string) *CreateFulfilRecordRequest {
	s.FaceId = &v
	return s
}

func (s *CreateFulfilRecordRequest) SetScene(v int64) *CreateFulfilRecordRequest {
	s.Scene = &v
	return s
}

func (s *CreateFulfilRecordRequest) SetSn(v string) *CreateFulfilRecordRequest {
	s.Sn = &v
	return s
}

func (s *CreateFulfilRecordRequest) SetUserId(v string) *CreateFulfilRecordRequest {
	s.UserId = &v
	return s
}

type CreateFulfilRecordResponseBody struct {
	SuccessInfo *string `json:"successInfo,omitempty" xml:"successInfo,omitempty"`
}

func (s CreateFulfilRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFulfilRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFulfilRecordResponseBody) SetSuccessInfo(v string) *CreateFulfilRecordResponseBody {
	s.SuccessInfo = &v
	return s
}

type CreateFulfilRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFulfilRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFulfilRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFulfilRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateFulfilRecordResponse) SetHeaders(v map[string]*string) *CreateFulfilRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateFulfilRecordResponse) SetStatusCode(v int32) *CreateFulfilRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFulfilRecordResponse) SetBody(v *CreateFulfilRecordResponseBody) *CreateFulfilRecordResponse {
	s.Body = v
	return s
}

type CreateInviteUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInviteUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlHeaders) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlHeaders) SetCommonHeaders(v map[string]*string) *CreateInviteUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInviteUrlHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInviteUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInviteUrlRequest struct {
	AuthCode       *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	TargetCorpId   *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TargetOperator *string `json:"targetOperator,omitempty" xml:"targetOperator,omitempty"`
}

func (s CreateInviteUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlRequest) SetAuthCode(v string) *CreateInviteUrlRequest {
	s.AuthCode = &v
	return s
}

func (s *CreateInviteUrlRequest) SetTargetCorpId(v string) *CreateInviteUrlRequest {
	s.TargetCorpId = &v
	return s
}

func (s *CreateInviteUrlRequest) SetTargetOperator(v string) *CreateInviteUrlRequest {
	s.TargetOperator = &v
	return s
}

type CreateInviteUrlResponseBody struct {
	Result  *CreateInviteUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInviteUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponseBody) SetResult(v *CreateInviteUrlResponseBodyResult) *CreateInviteUrlResponseBody {
	s.Result = v
	return s
}

func (s *CreateInviteUrlResponseBody) SetSuccess(v bool) *CreateInviteUrlResponseBody {
	s.Success = &v
	return s
}

type CreateInviteUrlResponseBodyResult struct {
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
}

func (s CreateInviteUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponseBodyResult) SetInviteUrl(v string) *CreateInviteUrlResponseBodyResult {
	s.InviteUrl = &v
	return s
}

type CreateInviteUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInviteUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInviteUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInviteUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateInviteUrlResponse) SetHeaders(v map[string]*string) *CreateInviteUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateInviteUrlResponse) SetStatusCode(v int32) *CreateInviteUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInviteUrlResponse) SetBody(v *CreateInviteUrlResponseBody) *CreateInviteUrlResponse {
	s.Body = v
	return s
}

type CreateItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateItemHeaders) GoString() string {
	return s.String()
}

func (s *CreateItemHeaders) SetCommonHeaders(v map[string]*string) *CreateItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateItemHeaders) SetXAcsDingtalkAccessToken(v string) *CreateItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateItemRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EffectType  *int64  `json:"effectType,omitempty" xml:"effectType,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MerchantId  *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OptUser     *string `json:"optUser,omitempty" xml:"optUser,omitempty"`
	PeriodType  *int64  `json:"periodType,omitempty" xml:"periodType,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
	Scene       *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Type        *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateItemRequest) GoString() string {
	return s.String()
}

func (s *CreateItemRequest) SetDescription(v string) *CreateItemRequest {
	s.Description = &v
	return s
}

func (s *CreateItemRequest) SetEffectType(v int64) *CreateItemRequest {
	s.EffectType = &v
	return s
}

func (s *CreateItemRequest) SetEndTime(v int64) *CreateItemRequest {
	s.EndTime = &v
	return s
}

func (s *CreateItemRequest) SetMerchantId(v string) *CreateItemRequest {
	s.MerchantId = &v
	return s
}

func (s *CreateItemRequest) SetName(v string) *CreateItemRequest {
	s.Name = &v
	return s
}

func (s *CreateItemRequest) SetOptUser(v string) *CreateItemRequest {
	s.OptUser = &v
	return s
}

func (s *CreateItemRequest) SetPeriodType(v int64) *CreateItemRequest {
	s.PeriodType = &v
	return s
}

func (s *CreateItemRequest) SetPrice(v int64) *CreateItemRequest {
	s.Price = &v
	return s
}

func (s *CreateItemRequest) SetScene(v int64) *CreateItemRequest {
	s.Scene = &v
	return s
}

func (s *CreateItemRequest) SetStartTime(v int64) *CreateItemRequest {
	s.StartTime = &v
	return s
}

func (s *CreateItemRequest) SetStatus(v int32) *CreateItemRequest {
	s.Status = &v
	return s
}

func (s *CreateItemRequest) SetType(v int64) *CreateItemRequest {
	s.Type = &v
	return s
}

type CreateItemResponseBody struct {
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateItemResponseBody) SetCorpId(v string) *CreateItemResponseBody {
	s.CorpId = &v
	return s
}

func (s *CreateItemResponseBody) SetId(v int64) *CreateItemResponseBody {
	s.Id = &v
	return s
}

func (s *CreateItemResponseBody) SetMerchantId(v string) *CreateItemResponseBody {
	s.MerchantId = &v
	return s
}

func (s *CreateItemResponseBody) SetStatus(v int64) *CreateItemResponseBody {
	s.Status = &v
	return s
}

type CreateItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateItemResponse) GoString() string {
	return s.String()
}

func (s *CreateItemResponse) SetHeaders(v map[string]*string) *CreateItemResponse {
	s.Headers = v
	return s
}

func (s *CreateItemResponse) SetStatusCode(v int32) *CreateItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateItemResponse) SetBody(v *CreateItemResponseBody) *CreateItemResponse {
	s.Body = v
	return s
}

type CreateOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrderRequest struct {
	ActualAmount   *int64                          `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	CreateTime     *int64                          `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DetailList     []*CreateOrderRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
	FaceId         *string                         `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Ftoken         *string                         `json:"ftoken,omitempty" xml:"ftoken,omitempty"`
	Signature      *string                         `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn             *string                         `json:"sn,omitempty" xml:"sn,omitempty"`
	TerminalParams *string                         `json:"terminalParams,omitempty" xml:"terminalParams,omitempty"`
	Timestamp      *int64                          `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TotalAmount    *int64                          `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	UserId         *string                         `json:"userId,omitempty" xml:"userId,omitempty"`
	Version        *string                         `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetActualAmount(v int64) *CreateOrderRequest {
	s.ActualAmount = &v
	return s
}

func (s *CreateOrderRequest) SetCreateTime(v int64) *CreateOrderRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateOrderRequest) SetDetailList(v []*CreateOrderRequestDetailList) *CreateOrderRequest {
	s.DetailList = v
	return s
}

func (s *CreateOrderRequest) SetFaceId(v string) *CreateOrderRequest {
	s.FaceId = &v
	return s
}

func (s *CreateOrderRequest) SetFtoken(v string) *CreateOrderRequest {
	s.Ftoken = &v
	return s
}

func (s *CreateOrderRequest) SetSignature(v string) *CreateOrderRequest {
	s.Signature = &v
	return s
}

func (s *CreateOrderRequest) SetSn(v string) *CreateOrderRequest {
	s.Sn = &v
	return s
}

func (s *CreateOrderRequest) SetTerminalParams(v string) *CreateOrderRequest {
	s.TerminalParams = &v
	return s
}

func (s *CreateOrderRequest) SetTimestamp(v int64) *CreateOrderRequest {
	s.Timestamp = &v
	return s
}

func (s *CreateOrderRequest) SetTotalAmount(v int64) *CreateOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateOrderRequest) SetUserId(v string) *CreateOrderRequest {
	s.UserId = &v
	return s
}

func (s *CreateOrderRequest) SetVersion(v string) *CreateOrderRequest {
	s.Version = &v
	return s
}

type CreateOrderRequestDetailList struct {
	ActualAmount *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	ItemAmount   *int64  `json:"itemAmount,omitempty" xml:"itemAmount,omitempty"`
	ItemName     *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	Scene        *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s CreateOrderRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequestDetailList) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestDetailList) SetActualAmount(v int64) *CreateOrderRequestDetailList {
	s.ActualAmount = &v
	return s
}

func (s *CreateOrderRequestDetailList) SetItemAmount(v int64) *CreateOrderRequestDetailList {
	s.ItemAmount = &v
	return s
}

func (s *CreateOrderRequestDetailList) SetItemName(v string) *CreateOrderRequestDetailList {
	s.ItemName = &v
	return s
}

func (s *CreateOrderRequestDetailList) SetScene(v int64) *CreateOrderRequestDetailList {
	s.Scene = &v
	return s
}

type CreateOrderResponseBody struct {
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetOrderNo(v string) *CreateOrderResponseBody {
	s.OrderNo = &v
	return s
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreateOrderFlowHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrderFlowHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderFlowHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrderFlowHeaders) SetCommonHeaders(v map[string]*string) *CreateOrderFlowHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrderFlowHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrderFlowHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrderFlowRequest struct {
	ActualAmount   *int64                              `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayUid      *string                             `json:"alipayUid,omitempty" xml:"alipayUid,omitempty"`
	CreateTime     *int64                              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DetailList     []*CreateOrderFlowRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
	FaceId         *string                             `json:"faceId,omitempty" xml:"faceId,omitempty"`
	GuardianUserId *string                             `json:"guardianUserId,omitempty" xml:"guardianUserId,omitempty"`
	MerchantId     *string                             `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo        *string                             `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature      *string                             `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn             *string                             `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp      *int64                              `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TotalAmount    *int64                              `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	UserId         *string                             `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrderFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderFlowRequest) SetActualAmount(v int64) *CreateOrderFlowRequest {
	s.ActualAmount = &v
	return s
}

func (s *CreateOrderFlowRequest) SetAlipayUid(v string) *CreateOrderFlowRequest {
	s.AlipayUid = &v
	return s
}

func (s *CreateOrderFlowRequest) SetCreateTime(v int64) *CreateOrderFlowRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateOrderFlowRequest) SetDetailList(v []*CreateOrderFlowRequestDetailList) *CreateOrderFlowRequest {
	s.DetailList = v
	return s
}

func (s *CreateOrderFlowRequest) SetFaceId(v string) *CreateOrderFlowRequest {
	s.FaceId = &v
	return s
}

func (s *CreateOrderFlowRequest) SetGuardianUserId(v string) *CreateOrderFlowRequest {
	s.GuardianUserId = &v
	return s
}

func (s *CreateOrderFlowRequest) SetMerchantId(v string) *CreateOrderFlowRequest {
	s.MerchantId = &v
	return s
}

func (s *CreateOrderFlowRequest) SetOrderNo(v string) *CreateOrderFlowRequest {
	s.OrderNo = &v
	return s
}

func (s *CreateOrderFlowRequest) SetSignature(v string) *CreateOrderFlowRequest {
	s.Signature = &v
	return s
}

func (s *CreateOrderFlowRequest) SetSn(v string) *CreateOrderFlowRequest {
	s.Sn = &v
	return s
}

func (s *CreateOrderFlowRequest) SetTimestamp(v int64) *CreateOrderFlowRequest {
	s.Timestamp = &v
	return s
}

func (s *CreateOrderFlowRequest) SetTotalAmount(v int64) *CreateOrderFlowRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateOrderFlowRequest) SetUserId(v string) *CreateOrderFlowRequest {
	s.UserId = &v
	return s
}

type CreateOrderFlowRequestDetailList struct {
	ActualAmount *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	ItemAmount   *int64  `json:"itemAmount,omitempty" xml:"itemAmount,omitempty"`
	ItemId       *int64  `json:"itemId,omitempty" xml:"itemId,omitempty"`
	ItemName     *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	Scene        *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s CreateOrderFlowRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderFlowRequestDetailList) GoString() string {
	return s.String()
}

func (s *CreateOrderFlowRequestDetailList) SetActualAmount(v int64) *CreateOrderFlowRequestDetailList {
	s.ActualAmount = &v
	return s
}

func (s *CreateOrderFlowRequestDetailList) SetItemAmount(v int64) *CreateOrderFlowRequestDetailList {
	s.ItemAmount = &v
	return s
}

func (s *CreateOrderFlowRequestDetailList) SetItemId(v int64) *CreateOrderFlowRequestDetailList {
	s.ItemId = &v
	return s
}

func (s *CreateOrderFlowRequestDetailList) SetItemName(v string) *CreateOrderFlowRequestDetailList {
	s.ItemName = &v
	return s
}

func (s *CreateOrderFlowRequestDetailList) SetScene(v int64) *CreateOrderFlowRequestDetailList {
	s.Scene = &v
	return s
}

type CreateOrderFlowResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateOrderFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderFlowResponseBody) SetSuccess(v bool) *CreateOrderFlowResponseBody {
	s.Success = &v
	return s
}

type CreateOrderFlowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrderFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderFlowResponse) SetHeaders(v map[string]*string) *CreateOrderFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderFlowResponse) SetStatusCode(v int32) *CreateOrderFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderFlowResponse) SetBody(v *CreateOrderFlowResponseBody) *CreateOrderFlowResponse {
	s.Body = v
	return s
}

type CreatePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreatePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *CreatePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *CreatePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreatePhysicalClassroomRequest struct {
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	ClassroomCampus   *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	ClassroomFloor    *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	ClassroomName     *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	ClassroomNumber   *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
	DirectBroadcast   *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
	Ext               *string `json:"ext,omitempty" xml:"ext,omitempty"`
	OpUserId          *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreatePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomRequest) SetClassroomBuilding(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomBuilding = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomCampus(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomCampus = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomFloor(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomFloor = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomName(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomName = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetClassroomNumber(v string) *CreatePhysicalClassroomRequest {
	s.ClassroomNumber = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetDirectBroadcast(v string) *CreatePhysicalClassroomRequest {
	s.DirectBroadcast = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetExt(v string) *CreatePhysicalClassroomRequest {
	s.Ext = &v
	return s
}

func (s *CreatePhysicalClassroomRequest) SetOpUserId(v string) *CreatePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

type CreatePhysicalClassroomResponseBody struct {
	ClassroomId *int64 `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
}

func (s CreatePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomResponseBody) SetClassroomId(v int64) *CreatePhysicalClassroomResponseBody {
	s.ClassroomId = &v
	return s
}

type CreatePhysicalClassroomResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalClassroomResponse) SetHeaders(v map[string]*string) *CreatePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalClassroomResponse) SetStatusCode(v int32) *CreatePhysicalClassroomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhysicalClassroomResponse) SetBody(v *CreatePhysicalClassroomResponseBody) *CreatePhysicalClassroomResponse {
	s.Body = v
	return s
}

type CreateRefundFlowHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRefundFlowHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundFlowHeaders) GoString() string {
	return s.String()
}

func (s *CreateRefundFlowHeaders) SetCommonHeaders(v map[string]*string) *CreateRefundFlowHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRefundFlowHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRefundFlowHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRefundFlowRequest struct {
	FaceId       *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	OperatorId   *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	OrderNo      *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature    *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn           *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp    *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateRefundFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateRefundFlowRequest) SetFaceId(v string) *CreateRefundFlowRequest {
	s.FaceId = &v
	return s
}

func (s *CreateRefundFlowRequest) SetOperatorId(v string) *CreateRefundFlowRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateRefundFlowRequest) SetOperatorName(v string) *CreateRefundFlowRequest {
	s.OperatorName = &v
	return s
}

func (s *CreateRefundFlowRequest) SetOrderNo(v string) *CreateRefundFlowRequest {
	s.OrderNo = &v
	return s
}

func (s *CreateRefundFlowRequest) SetSignature(v string) *CreateRefundFlowRequest {
	s.Signature = &v
	return s
}

func (s *CreateRefundFlowRequest) SetSn(v string) *CreateRefundFlowRequest {
	s.Sn = &v
	return s
}

func (s *CreateRefundFlowRequest) SetTimestamp(v int64) *CreateRefundFlowRequest {
	s.Timestamp = &v
	return s
}

func (s *CreateRefundFlowRequest) SetUserId(v string) *CreateRefundFlowRequest {
	s.UserId = &v
	return s
}

type CreateRefundFlowResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRefundFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRefundFlowResponseBody) SetSuccess(v bool) *CreateRefundFlowResponseBody {
	s.Success = &v
	return s
}

type CreateRefundFlowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRefundFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRefundFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateRefundFlowResponse) SetHeaders(v map[string]*string) *CreateRefundFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateRefundFlowResponse) SetStatusCode(v int32) *CreateRefundFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRefundFlowResponse) SetBody(v *CreateRefundFlowResponseBody) *CreateRefundFlowResponse {
	s.Body = v
	return s
}

type CreateRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *CreateRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRemoteClassCourseRequest struct {
	AttendParticipants  []*CreateRemoteClassCourseRequestAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	AuthCode            *string                                             `json:"authCode,omitempty" xml:"authCode,omitempty"`
	CourseName          *string                                             `json:"courseName,omitempty" xml:"courseName,omitempty"`
	EndTime             *int64                                              `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime           *int64                                              `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TeachingParticipant *CreateRemoteClassCourseRequestTeachingParticipant  `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
}

func (s CreateRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequest) SetAttendParticipants(v []*CreateRemoteClassCourseRequestAttendParticipants) *CreateRemoteClassCourseRequest {
	s.AttendParticipants = v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetAuthCode(v string) *CreateRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetCourseName(v string) *CreateRemoteClassCourseRequest {
	s.CourseName = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetEndTime(v int64) *CreateRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetStartTime(v int64) *CreateRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

func (s *CreateRemoteClassCourseRequest) SetTeachingParticipant(v *CreateRemoteClassCourseRequestTeachingParticipant) *CreateRemoteClassCourseRequest {
	s.TeachingParticipant = v
	return s
}

type CreateRemoteClassCourseRequestAttendParticipants struct {
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
}

func (s CreateRemoteClassCourseRequestAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequestAttendParticipants) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequestAttendParticipants) SetCorpId(v string) *CreateRemoteClassCourseRequestAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *CreateRemoteClassCourseRequestAttendParticipants) SetParticipantId(v string) *CreateRemoteClassCourseRequestAttendParticipants {
	s.ParticipantId = &v
	return s
}

type CreateRemoteClassCourseRequestTeachingParticipant struct {
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
}

func (s CreateRemoteClassCourseRequestTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseRequestTeachingParticipant) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseRequestTeachingParticipant) SetCorpId(v string) *CreateRemoteClassCourseRequestTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *CreateRemoteClassCourseRequestTeachingParticipant) SetParticipantId(v string) *CreateRemoteClassCourseRequestTeachingParticipant {
	s.ParticipantId = &v
	return s
}

type CreateRemoteClassCourseResponseBody struct {
	Result  *CreateRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponseBody) SetResult(v *CreateRemoteClassCourseResponseBodyResult) *CreateRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

func (s *CreateRemoteClassCourseResponseBody) SetSuccess(v bool) *CreateRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type CreateRemoteClassCourseResponseBodyResult struct {
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
}

func (s CreateRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *CreateRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

type CreateRemoteClassCourseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *CreateRemoteClassCourseResponse) SetHeaders(v map[string]*string) *CreateRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *CreateRemoteClassCourseResponse) SetStatusCode(v int32) *CreateRemoteClassCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemoteClassCourseResponse) SetBody(v *CreateRemoteClassCourseResponseBody) *CreateRemoteClassCourseResponse {
	s.Body = v
	return s
}

type CreateSectionConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSectionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigHeaders) SetCommonHeaders(v map[string]*string) *CreateSectionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSectionConfigHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSectionConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSectionConfigRequest struct {
	Ext            *string                                     `json:"ext,omitempty" xml:"ext,omitempty"`
	SectionConfigs []*CreateSectionConfigRequestSectionConfigs `json:"sectionConfigs,omitempty" xml:"sectionConfigs,omitempty" type:"Repeated"`
	OpUserId       *string                                     `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreateSectionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequest) SetExt(v string) *CreateSectionConfigRequest {
	s.Ext = &v
	return s
}

func (s *CreateSectionConfigRequest) SetSectionConfigs(v []*CreateSectionConfigRequestSectionConfigs) *CreateSectionConfigRequest {
	s.SectionConfigs = v
	return s
}

func (s *CreateSectionConfigRequest) SetOpUserId(v string) *CreateSectionConfigRequest {
	s.OpUserId = &v
	return s
}

type CreateSectionConfigRequestSectionConfigs struct {
	ScheduleName      *string                                                    `json:"scheduleName,omitempty" xml:"scheduleName,omitempty"`
	SchoolYear        *string                                                    `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	SectionEndDate    *CreateSectionConfigRequestSectionConfigsSectionEndDate    `json:"sectionEndDate,omitempty" xml:"sectionEndDate,omitempty" type:"Struct"`
	SectionModels     []*CreateSectionConfigRequestSectionConfigsSectionModels   `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	SectionStartDate  *CreateSectionConfigRequestSectionConfigsSectionStartDate  `json:"sectionStartDate,omitempty" xml:"sectionStartDate,omitempty" type:"Struct"`
	Semester          *int32                                                     `json:"semester,omitempty" xml:"semester,omitempty"`
	SemesterEndDate   *CreateSectionConfigRequestSectionConfigsSemesterEndDate   `json:"semesterEndDate,omitempty" xml:"semesterEndDate,omitempty" type:"Struct"`
	SemesterStartDate *CreateSectionConfigRequestSectionConfigsSemesterStartDate `json:"semesterStartDate,omitempty" xml:"semesterStartDate,omitempty" type:"Struct"`
}

func (s CreateSectionConfigRequestSectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigs) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigs) SetScheduleName(v string) *CreateSectionConfigRequestSectionConfigs {
	s.ScheduleName = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSchoolYear(v string) *CreateSectionConfigRequestSectionConfigs {
	s.SchoolYear = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionEndDate(v *CreateSectionConfigRequestSectionConfigsSectionEndDate) *CreateSectionConfigRequestSectionConfigs {
	s.SectionEndDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionModels(v []*CreateSectionConfigRequestSectionConfigsSectionModels) *CreateSectionConfigRequestSectionConfigs {
	s.SectionModels = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSectionStartDate(v *CreateSectionConfigRequestSectionConfigsSectionStartDate) *CreateSectionConfigRequestSectionConfigs {
	s.SectionStartDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemester(v int32) *CreateSectionConfigRequestSectionConfigs {
	s.Semester = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemesterEndDate(v *CreateSectionConfigRequestSectionConfigsSemesterEndDate) *CreateSectionConfigRequestSectionConfigs {
	s.SemesterEndDate = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigs) SetSemesterStartDate(v *CreateSectionConfigRequestSectionConfigsSemesterStartDate) *CreateSectionConfigRequestSectionConfigs {
	s.SemesterStartDate = v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionEndDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionEndDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionEndDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSectionEndDate {
	s.Year = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModels struct {
	SectionEndTime   *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime   `json:"sectionEndTime,omitempty" xml:"sectionEndTime,omitempty" type:"Struct"`
	SectionIndex     *int32                                                                 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName      *string                                                                `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	SectionStartTime *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime `json:"sectionStartTime,omitempty" xml:"sectionStartTime,omitempty" type:"Struct"`
	SectionType      *string                                                                `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModels) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModels) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionEndTime(v *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionEndTime = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionIndex(v int32) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionName(v string) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionName = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionStartTime(v *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionStartTime = v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModels) SetSectionType(v string) *CreateSectionConfigRequestSectionConfigsSectionModels {
	s.SectionType = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) SetHour(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime {
	s.Hour = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime) SetMin(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionEndTime {
	s.Min = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) SetHour(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime {
	s.Hour = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime) SetMin(v int32) *CreateSectionConfigRequestSectionConfigsSectionModelsSectionStartTime {
	s.Min = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSectionStartDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSectionStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSectionStartDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSectionStartDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSectionStartDate {
	s.Year = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSemesterEndDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSemesterEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSemesterEndDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterEndDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSemesterEndDate {
	s.Year = &v
	return s
}

type CreateSectionConfigRequestSectionConfigsSemesterStartDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateSectionConfigRequestSectionConfigsSemesterStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigRequestSectionConfigsSemesterStartDate) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetDayOfMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetMonth(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.Month = &v
	return s
}

func (s *CreateSectionConfigRequestSectionConfigsSemesterStartDate) SetYear(v int32) *CreateSectionConfigRequestSectionConfigsSemesterStartDate {
	s.Year = &v
	return s
}

type CreateSectionConfigResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSectionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigResponseBody) SetResult(v bool) *CreateSectionConfigResponseBody {
	s.Result = &v
	return s
}

type CreateSectionConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSectionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSectionConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSectionConfigResponse) SetHeaders(v map[string]*string) *CreateSectionConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSectionConfigResponse) SetStatusCode(v int32) *CreateSectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSectionConfigResponse) SetBody(v *CreateSectionConfigResponseBody) *CreateSectionConfigResponse {
	s.Body = v
	return s
}

type CreateSnsAppOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSnsAppOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSnsAppOrderHeaders) GoString() string {
	return s.String()
}

func (s *CreateSnsAppOrderHeaders) SetCommonHeaders(v map[string]*string) *CreateSnsAppOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSnsAppOrderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSnsAppOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSnsAppOrderRequest struct {
	ActualAmount    *int64                                `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId     *string                               `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	BizCode         *int32                                `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	DetailList      []*CreateSnsAppOrderRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
	LabelAmount     *int64                                `json:"labelAmount,omitempty" xml:"labelAmount,omitempty"`
	MerchantId      *string                               `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string                               `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	Signature       *string                               `json:"signature,omitempty" xml:"signature,omitempty"`
	Subject         *string                               `json:"subject,omitempty" xml:"subject,omitempty"`
	Timestamp       *int64                                `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s CreateSnsAppOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnsAppOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateSnsAppOrderRequest) SetActualAmount(v int64) *CreateSnsAppOrderRequest {
	s.ActualAmount = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetAlipayAppId(v string) *CreateSnsAppOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetBizCode(v int32) *CreateSnsAppOrderRequest {
	s.BizCode = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetDetailList(v []*CreateSnsAppOrderRequestDetailList) *CreateSnsAppOrderRequest {
	s.DetailList = v
	return s
}

func (s *CreateSnsAppOrderRequest) SetLabelAmount(v int64) *CreateSnsAppOrderRequest {
	s.LabelAmount = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetMerchantId(v string) *CreateSnsAppOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetMerchantOrderNo(v string) *CreateSnsAppOrderRequest {
	s.MerchantOrderNo = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetSignature(v string) *CreateSnsAppOrderRequest {
	s.Signature = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetSubject(v string) *CreateSnsAppOrderRequest {
	s.Subject = &v
	return s
}

func (s *CreateSnsAppOrderRequest) SetTimestamp(v int64) *CreateSnsAppOrderRequest {
	s.Timestamp = &v
	return s
}

type CreateSnsAppOrderRequestDetailList struct {
	GoodsId       *string `json:"goodsId,omitempty" xml:"goodsId,omitempty"`
	GoodsName     *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	GoodsPrice    *int64  `json:"goodsPrice,omitempty" xml:"goodsPrice,omitempty"`
	GoodsQuantity *int32  `json:"goodsQuantity,omitempty" xml:"goodsQuantity,omitempty"`
}

func (s CreateSnsAppOrderRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s CreateSnsAppOrderRequestDetailList) GoString() string {
	return s.String()
}

func (s *CreateSnsAppOrderRequestDetailList) SetGoodsId(v string) *CreateSnsAppOrderRequestDetailList {
	s.GoodsId = &v
	return s
}

func (s *CreateSnsAppOrderRequestDetailList) SetGoodsName(v string) *CreateSnsAppOrderRequestDetailList {
	s.GoodsName = &v
	return s
}

func (s *CreateSnsAppOrderRequestDetailList) SetGoodsPrice(v int64) *CreateSnsAppOrderRequestDetailList {
	s.GoodsPrice = &v
	return s
}

func (s *CreateSnsAppOrderRequestDetailList) SetGoodsQuantity(v int32) *CreateSnsAppOrderRequestDetailList {
	s.GoodsQuantity = &v
	return s
}

type CreateSnsAppOrderResponseBody struct {
	ActualAmount    *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId     *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	Body            *string `json:"body,omitempty" xml:"body,omitempty"`
	MerchantId      *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantOrderNo *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
}

func (s CreateSnsAppOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnsAppOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnsAppOrderResponseBody) SetActualAmount(v int64) *CreateSnsAppOrderResponseBody {
	s.ActualAmount = &v
	return s
}

func (s *CreateSnsAppOrderResponseBody) SetAlipayAppId(v string) *CreateSnsAppOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *CreateSnsAppOrderResponseBody) SetBody(v string) *CreateSnsAppOrderResponseBody {
	s.Body = &v
	return s
}

func (s *CreateSnsAppOrderResponseBody) SetMerchantId(v string) *CreateSnsAppOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *CreateSnsAppOrderResponseBody) SetMerchantOrderNo(v string) *CreateSnsAppOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *CreateSnsAppOrderResponseBody) SetOrderNo(v string) *CreateSnsAppOrderResponseBody {
	s.OrderNo = &v
	return s
}

type CreateSnsAppOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnsAppOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnsAppOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnsAppOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateSnsAppOrderResponse) SetHeaders(v map[string]*string) *CreateSnsAppOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateSnsAppOrderResponse) SetStatusCode(v int32) *CreateSnsAppOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnsAppOrderResponse) SetBody(v *CreateSnsAppOrderResponseBody) *CreateSnsAppOrderResponse {
	s.Body = v
	return s
}

type CreateStsTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateStsTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateStsTokenHeaders) GoString() string {
	return s.String()
}

func (s *CreateStsTokenHeaders) SetCommonHeaders(v map[string]*string) *CreateStsTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateStsTokenHeaders) SetXAcsDingtalkAccessToken(v string) *CreateStsTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateStsTokenRequest struct {
	DeviceSn *string `json:"deviceSn,omitempty" xml:"deviceSn,omitempty"`
	StsType  *string `json:"stsType,omitempty" xml:"stsType,omitempty"`
}

func (s CreateStsTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStsTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateStsTokenRequest) SetDeviceSn(v string) *CreateStsTokenRequest {
	s.DeviceSn = &v
	return s
}

func (s *CreateStsTokenRequest) SetStsType(v string) *CreateStsTokenRequest {
	s.StsType = &v
	return s
}

type CreateStsTokenResponseBody struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Expiration      *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	ExtInfo         *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateStsTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStsTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStsTokenResponseBody) SetAccessKeyId(v string) *CreateStsTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *CreateStsTokenResponseBody) SetAccessKeySecret(v string) *CreateStsTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateStsTokenResponseBody) SetExpiration(v string) *CreateStsTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *CreateStsTokenResponseBody) SetExtInfo(v string) *CreateStsTokenResponseBody {
	s.ExtInfo = &v
	return s
}

func (s *CreateStsTokenResponseBody) SetSecurityToken(v string) *CreateStsTokenResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *CreateStsTokenResponseBody) SetStatus(v string) *CreateStsTokenResponseBody {
	s.Status = &v
	return s
}

type CreateStsTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStsTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStsTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStsTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateStsTokenResponse) SetHeaders(v map[string]*string) *CreateStsTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateStsTokenResponse) SetStatusCode(v int32) *CreateStsTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStsTokenResponse) SetBody(v *CreateStsTokenResponseBody) *CreateStsTokenResponse {
	s.Body = v
	return s
}

type CreateTokenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTokenHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenHeaders) GoString() string {
	return s.String()
}

func (s *CreateTokenHeaders) SetCommonHeaders(v map[string]*string) *CreateTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTokenHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTokenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTokenRequest struct {
	Sn   *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetSn(v string) *CreateTokenRequest {
	s.Sn = &v
	return s
}

func (s *CreateTokenRequest) SetType(v string) *CreateTokenRequest {
	s.Type = &v
	return s
}

type CreateTokenResponseBody struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Expiration      *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	ExtInfo         *string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetAccessKeyId(v string) *CreateTokenResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *CreateTokenResponseBody) SetAccessKeySecret(v string) *CreateTokenResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateTokenResponseBody) SetExpiration(v string) *CreateTokenResponseBody {
	s.Expiration = &v
	return s
}

func (s *CreateTokenResponseBody) SetExtInfo(v string) *CreateTokenResponseBody {
	s.ExtInfo = &v
	return s
}

func (s *CreateTokenResponseBody) SetSecurityToken(v string) *CreateTokenResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *CreateTokenResponseBody) SetStatus(v string) *CreateTokenResponseBody {
	s.Status = &v
	return s
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

type CreateUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUniversityCourseGroupRequest struct {
	CourseGroupIntroduce   *string                                                     `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	CourseGroupName        *string                                                     `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	CourserGroupItemModels []*CreateUniversityCourseGroupRequestCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
	Ext                    *string                                                     `json:"ext,omitempty" xml:"ext,omitempty"`
	IsvCourseGroupCode     *string                                                     `json:"isvCourseGroupCode,omitempty" xml:"isvCourseGroupCode,omitempty"`
	PeriodCode             *string                                                     `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	SchoolYear             *string                                                     `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	Semester               *int32                                                      `json:"semester,omitempty" xml:"semester,omitempty"`
	SubjectName            *string                                                     `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	TeacherInfos           []*CreateUniversityCourseGroupRequestTeacherInfos           `json:"teacherInfos,omitempty" xml:"teacherInfos,omitempty" type:"Repeated"`
	OpUserId               *string                                                     `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreateUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequest) SetCourseGroupIntroduce(v string) *CreateUniversityCourseGroupRequest {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetCourseGroupName(v string) *CreateUniversityCourseGroupRequest {
	s.CourseGroupName = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetCourserGroupItemModels(v []*CreateUniversityCourseGroupRequestCourserGroupItemModels) *CreateUniversityCourseGroupRequest {
	s.CourserGroupItemModels = v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetExt(v string) *CreateUniversityCourseGroupRequest {
	s.Ext = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetIsvCourseGroupCode(v string) *CreateUniversityCourseGroupRequest {
	s.IsvCourseGroupCode = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetPeriodCode(v string) *CreateUniversityCourseGroupRequest {
	s.PeriodCode = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSchoolYear(v string) *CreateUniversityCourseGroupRequest {
	s.SchoolYear = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSemester(v int32) *CreateUniversityCourseGroupRequest {
	s.Semester = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetSubjectName(v string) *CreateUniversityCourseGroupRequest {
	s.SubjectName = &v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetTeacherInfos(v []*CreateUniversityCourseGroupRequestTeacherInfos) *CreateUniversityCourseGroupRequest {
	s.TeacherInfos = v
	return s
}

func (s *CreateUniversityCourseGroupRequest) SetOpUserId(v string) *CreateUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModels struct {
	ClassPeriodType           *int32                                                                             `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	ClassroomId               *int64                                                                             `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	CourseType                *int32                                                                             `json:"courseType,omitempty" xml:"courseType,omitempty"`
	CourserGroupItemEndDate   *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate   `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	CourserGroupItemStartDate *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	DayOfWeek                 *int32                                                                             `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	SectionIndex              []*int32                                                                           `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetClassPeriodType(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetClassroomId(v int64) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourseType(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourseType = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemEndDate(v *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemStartDate(v *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetDayOfWeek(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModels) SetSectionIndex(v []*int32) *CreateUniversityCourseGroupRequestCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

type CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *CreateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

type CreateUniversityCourseGroupRequestTeacherInfos struct {
	ParticipantRole *string `json:"participantRole,omitempty" xml:"participantRole,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateUniversityCourseGroupRequestTeacherInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupRequestTeacherInfos) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupRequestTeacherInfos) SetParticipantRole(v string) *CreateUniversityCourseGroupRequestTeacherInfos {
	s.ParticipantRole = &v
	return s
}

func (s *CreateUniversityCourseGroupRequestTeacherInfos) SetUserId(v string) *CreateUniversityCourseGroupRequestTeacherInfos {
	s.UserId = &v
	return s
}

type CreateUniversityCourseGroupResponseBody struct {
	CourseGroupInfo *CreateUniversityCourseGroupResponseBodyCourseGroupInfo `json:"courseGroupInfo,omitempty" xml:"courseGroupInfo,omitempty" type:"Struct"`
}

func (s CreateUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponseBody) SetCourseGroupInfo(v *CreateUniversityCourseGroupResponseBodyCourseGroupInfo) *CreateUniversityCourseGroupResponseBody {
	s.CourseGroupInfo = v
	return s
}

type CreateUniversityCourseGroupResponseBodyCourseGroupInfo struct {
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
}

func (s CreateUniversityCourseGroupResponseBodyCourseGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponseBodyCourseGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponseBodyCourseGroupInfo) SetCourseGroupCode(v string) *CreateUniversityCourseGroupResponseBodyCourseGroupInfo {
	s.CourseGroupCode = &v
	return s
}

type CreateUniversityCourseGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *CreateUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUniversityCourseGroupResponse) SetStatusCode(v int32) *CreateUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUniversityCourseGroupResponse) SetBody(v *CreateUniversityCourseGroupResponseBody) *CreateUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type CreateUniversityStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUniversityStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityStudentHeaders) GoString() string {
	return s.String()
}

func (s *CreateUniversityStudentHeaders) SetCommonHeaders(v map[string]*string) *CreateUniversityStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUniversityStudentHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUniversityStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUniversityStudentRequest struct {
	ClassId        *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	Gender         *string `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentityNumber *string `json:"identityNumber,omitempty" xml:"identityNumber,omitempty"`
	Mobile         *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	StudentNumber  *string `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
	OpUserId       *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s CreateUniversityStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityStudentRequest) GoString() string {
	return s.String()
}

func (s *CreateUniversityStudentRequest) SetClassId(v int64) *CreateUniversityStudentRequest {
	s.ClassId = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetGender(v string) *CreateUniversityStudentRequest {
	s.Gender = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetIdentityNumber(v string) *CreateUniversityStudentRequest {
	s.IdentityNumber = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetMobile(v string) *CreateUniversityStudentRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetName(v string) *CreateUniversityStudentRequest {
	s.Name = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetStudentNumber(v string) *CreateUniversityStudentRequest {
	s.StudentNumber = &v
	return s
}

func (s *CreateUniversityStudentRequest) SetOpUserId(v string) *CreateUniversityStudentRequest {
	s.OpUserId = &v
	return s
}

type CreateUniversityStudentResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateUniversityStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityStudentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniversityStudentResponseBody) SetResult(v bool) *CreateUniversityStudentResponseBody {
	s.Result = &v
	return s
}

type CreateUniversityStudentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUniversityStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUniversityStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityStudentResponse) GoString() string {
	return s.String()
}

func (s *CreateUniversityStudentResponse) SetHeaders(v map[string]*string) *CreateUniversityStudentResponse {
	s.Headers = v
	return s
}

func (s *CreateUniversityStudentResponse) SetStatusCode(v int32) *CreateUniversityStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUniversityStudentResponse) SetBody(v *CreateUniversityStudentResponseBody) *CreateUniversityStudentResponse {
	s.Body = v
	return s
}

type CreateUniversityTeacherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUniversityTeacherHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityTeacherHeaders) GoString() string {
	return s.String()
}

func (s *CreateUniversityTeacherHeaders) SetCommonHeaders(v map[string]*string) *CreateUniversityTeacherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUniversityTeacherHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUniversityTeacherHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUniversityTeacherRequest struct {
	ClassId       *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	OpUserId      *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	Role          *string `json:"role,omitempty" xml:"role,omitempty"`
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s CreateUniversityTeacherRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityTeacherRequest) GoString() string {
	return s.String()
}

func (s *CreateUniversityTeacherRequest) SetClassId(v int64) *CreateUniversityTeacherRequest {
	s.ClassId = &v
	return s
}

func (s *CreateUniversityTeacherRequest) SetOpUserId(v string) *CreateUniversityTeacherRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateUniversityTeacherRequest) SetRole(v string) *CreateUniversityTeacherRequest {
	s.Role = &v
	return s
}

func (s *CreateUniversityTeacherRequest) SetTeacherUserId(v string) *CreateUniversityTeacherRequest {
	s.TeacherUserId = &v
	return s
}

type CreateUniversityTeacherResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateUniversityTeacherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityTeacherResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUniversityTeacherResponseBody) SetResult(v bool) *CreateUniversityTeacherResponseBody {
	s.Result = &v
	return s
}

type CreateUniversityTeacherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUniversityTeacherResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUniversityTeacherResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUniversityTeacherResponse) GoString() string {
	return s.String()
}

func (s *CreateUniversityTeacherResponse) SetHeaders(v map[string]*string) *CreateUniversityTeacherResponse {
	s.Headers = v
	return s
}

func (s *CreateUniversityTeacherResponse) SetStatusCode(v int32) *CreateUniversityTeacherResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUniversityTeacherResponse) SetBody(v *CreateUniversityTeacherResponseBody) *CreateUniversityTeacherResponse {
	s.Body = v
	return s
}

type DeactivateDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeactivateDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeactivateDeviceHeaders) GoString() string {
	return s.String()
}

func (s *DeactivateDeviceHeaders) SetCommonHeaders(v map[string]*string) *DeactivateDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeactivateDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *DeactivateDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeactivateDeviceRequest struct {
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	Sn    *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeactivateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactivateDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeactivateDeviceRequest) SetModel(v string) *DeactivateDeviceRequest {
	s.Model = &v
	return s
}

func (s *DeactivateDeviceRequest) SetSn(v string) *DeactivateDeviceRequest {
	s.Sn = &v
	return s
}

func (s *DeactivateDeviceRequest) SetType(v string) *DeactivateDeviceRequest {
	s.Type = &v
	return s
}

type DeactivateDeviceResponseBody struct {
	ActivateTimes *int32 `json:"activateTimes,omitempty" xml:"activateTimes,omitempty"`
}

func (s DeactivateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactivateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateDeviceResponseBody) SetActivateTimes(v int32) *DeactivateDeviceResponseBody {
	s.ActivateTimes = &v
	return s
}

type DeactivateDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeactivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactivateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactivateDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeactivateDeviceResponse) SetHeaders(v map[string]*string) *DeactivateDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeactivateDeviceResponse) SetStatusCode(v int32) *DeactivateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateDeviceResponse) SetBody(v *DeactivateDeviceResponseBody) *DeactivateDeviceResponse {
	s.Body = v
	return s
}

type DeductPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeductPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeductPointHeaders) GoString() string {
	return s.String()
}

func (s *DeductPointHeaders) SetCommonHeaders(v map[string]*string) *DeductPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeductPointHeaders) SetXAcsDingtalkAccessToken(v string) *DeductPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeductPointRequest struct {
	BizId           *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	DeductDesc      *string `json:"deductDesc,omitempty" xml:"deductDesc,omitempty"`
	DeductDetailUrl *string `json:"deductDetailUrl,omitempty" xml:"deductDetailUrl,omitempty"`
	DeductNum       *int32  `json:"deductNum,omitempty" xml:"deductNum,omitempty"`
	PointType       *string `json:"pointType,omitempty" xml:"pointType,omitempty"`
}

func (s DeductPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeductPointRequest) GoString() string {
	return s.String()
}

func (s *DeductPointRequest) SetBizId(v string) *DeductPointRequest {
	s.BizId = &v
	return s
}

func (s *DeductPointRequest) SetDeductDesc(v string) *DeductPointRequest {
	s.DeductDesc = &v
	return s
}

func (s *DeductPointRequest) SetDeductDetailUrl(v string) *DeductPointRequest {
	s.DeductDetailUrl = &v
	return s
}

func (s *DeductPointRequest) SetDeductNum(v int32) *DeductPointRequest {
	s.DeductNum = &v
	return s
}

func (s *DeductPointRequest) SetPointType(v string) *DeductPointRequest {
	s.PointType = &v
	return s
}

type DeductPointResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeductPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeductPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeductPointResponseBody) SetResult(v bool) *DeductPointResponseBody {
	s.Result = &v
	return s
}

func (s *DeductPointResponseBody) SetSuccess(v bool) *DeductPointResponseBody {
	s.Success = &v
	return s
}

type DeductPointResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeductPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeductPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeductPointResponse) GoString() string {
	return s.String()
}

func (s *DeductPointResponse) SetHeaders(v map[string]*string) *DeductPointResponse {
	s.Headers = v
	return s
}

func (s *DeductPointResponse) SetStatusCode(v int32) *DeductPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeductPointResponse) SetBody(v *DeductPointResponseBody) *DeductPointResponse {
	s.Body = v
	return s
}

type DeleteDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeptHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeptHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeptRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeptRequest) SetOperator(v string) *DeleteDeptRequest {
	s.Operator = &v
	return s
}

type DeleteDeptResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeptResponseBody) SetSuccess(v bool) *DeleteDeptResponseBody {
	s.Success = &v
	return s
}

type DeleteDeptResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeptResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeptResponse) SetHeaders(v map[string]*string) *DeleteDeptResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeptResponse) SetStatusCode(v int32) *DeleteDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeptResponse) SetBody(v *DeleteDeptResponseBody) *DeleteDeptResponse {
	s.Body = v
	return s
}

type DeleteDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeviceHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeviceRequest struct {
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s DeleteDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceRequest) SetSn(v string) *DeleteDeviceRequest {
	s.Sn = &v
	return s
}

type DeleteDeviceResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) SetResult(v bool) *DeleteDeviceResponseBody {
	s.Result = &v
	return s
}

type DeleteDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponse) SetHeaders(v map[string]*string) *DeleteDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceResponse) SetStatusCode(v int32) *DeleteDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceResponse) SetBody(v *DeleteDeviceResponseBody) *DeleteDeviceResponse {
	s.Body = v
	return s
}

type DeleteDeviceOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDeviceOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeviceOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeviceOrgHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDeviceOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDeviceOrgRequest struct {
	AuthCode   *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
}

func (s DeleteDeviceOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgRequest) SetAuthCode(v string) *DeleteDeviceOrgRequest {
	s.AuthCode = &v
	return s
}

func (s *DeleteDeviceOrgRequest) SetDeviceCode(v string) *DeleteDeviceOrgRequest {
	s.DeviceCode = &v
	return s
}

type DeleteDeviceOrgResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDeviceOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgResponseBody) SetSuccess(v bool) *DeleteDeviceOrgResponseBody {
	s.Success = &v
	return s
}

type DeleteDeviceOrgResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceOrgResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceOrgResponse) SetHeaders(v map[string]*string) *DeleteDeviceOrgResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceOrgResponse) SetStatusCode(v int32) *DeleteDeviceOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceOrgResponse) SetBody(v *DeleteDeviceOrgResponseBody) *DeleteDeviceOrgResponse {
	s.Body = v
	return s
}

type DeleteGuardianHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGuardianHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGuardianHeaders) SetCommonHeaders(v map[string]*string) *DeleteGuardianHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGuardianHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGuardianHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGuardianRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	StuId    *string `json:"stuId,omitempty" xml:"stuId,omitempty"`
}

func (s DeleteGuardianRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianRequest) GoString() string {
	return s.String()
}

func (s *DeleteGuardianRequest) SetOperator(v string) *DeleteGuardianRequest {
	s.Operator = &v
	return s
}

func (s *DeleteGuardianRequest) SetStuId(v string) *DeleteGuardianRequest {
	s.StuId = &v
	return s
}

type DeleteGuardianResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGuardianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGuardianResponseBody) SetSuccess(v bool) *DeleteGuardianResponseBody {
	s.Success = &v
	return s
}

type DeleteGuardianResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGuardianResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGuardianResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGuardianResponse) GoString() string {
	return s.String()
}

func (s *DeleteGuardianResponse) SetHeaders(v map[string]*string) *DeleteGuardianResponse {
	s.Headers = v
	return s
}

func (s *DeleteGuardianResponse) SetStatusCode(v int32) *DeleteGuardianResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGuardianResponse) SetBody(v *DeleteGuardianResponseBody) *DeleteGuardianResponse {
	s.Body = v
	return s
}

type DeleteOrgRelationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteOrgRelationHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrgRelationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrgRelationHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteOrgRelationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteOrgRelationRequest struct {
	AuthCode     *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s DeleteOrgRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationRequest) SetAuthCode(v string) *DeleteOrgRelationRequest {
	s.AuthCode = &v
	return s
}

func (s *DeleteOrgRelationRequest) SetTargetCorpId(v string) *DeleteOrgRelationRequest {
	s.TargetCorpId = &v
	return s
}

type DeleteOrgRelationResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteOrgRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationResponseBody) SetSuccess(v bool) *DeleteOrgRelationResponseBody {
	s.Success = &v
	return s
}

type DeleteOrgRelationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOrgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOrgRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrgRelationResponse) SetHeaders(v map[string]*string) *DeleteOrgRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrgRelationResponse) SetStatusCode(v int32) *DeleteOrgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOrgRelationResponse) SetBody(v *DeleteOrgRelationResponseBody) *DeleteOrgRelationResponse {
	s.Body = v
	return s
}

type DeletePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeletePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *DeletePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *DeletePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeletePhysicalClassroomRequest struct {
	ClassroomId *int64  `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	OpUserId    *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s DeletePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomRequest) SetClassroomId(v int64) *DeletePhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

func (s *DeletePhysicalClassroomRequest) SetOpUserId(v string) *DeletePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

type DeletePhysicalClassroomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeletePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomResponseBody) SetResult(v bool) *DeletePhysicalClassroomResponseBody {
	s.Result = &v
	return s
}

type DeletePhysicalClassroomResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *DeletePhysicalClassroomResponse) SetHeaders(v map[string]*string) *DeletePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *DeletePhysicalClassroomResponse) SetStatusCode(v int32) *DeletePhysicalClassroomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePhysicalClassroomResponse) SetBody(v *DeletePhysicalClassroomResponseBody) *DeletePhysicalClassroomResponse {
	s.Body = v
	return s
}

type DeleteRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *DeleteRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteRemoteClassCourseRequest struct {
	AuthCode *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
}

func (s DeleteRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseRequest) SetAuthCode(v string) *DeleteRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

type DeleteRemoteClassCourseResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseResponseBody) SetSuccess(v bool) *DeleteRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type DeleteRemoteClassCourseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemoteClassCourseResponse) SetHeaders(v map[string]*string) *DeleteRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemoteClassCourseResponse) SetStatusCode(v int32) *DeleteRemoteClassCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRemoteClassCourseResponse) SetBody(v *DeleteRemoteClassCourseResponseBody) *DeleteRemoteClassCourseResponse {
	s.Body = v
	return s
}

type DeleteStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteStudentHeaders) SetCommonHeaders(v map[string]*string) *DeleteStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteStudentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteStudentRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentRequest) GoString() string {
	return s.String()
}

func (s *DeleteStudentRequest) SetOperator(v string) *DeleteStudentRequest {
	s.Operator = &v
	return s
}

type DeleteStudentResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStudentResponseBody) SetSuccess(v bool) *DeleteStudentResponseBody {
	s.Success = &v
	return s
}

type DeleteStudentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStudentResponse) GoString() string {
	return s.String()
}

func (s *DeleteStudentResponse) SetHeaders(v map[string]*string) *DeleteStudentResponse {
	s.Headers = v
	return s
}

func (s *DeleteStudentResponse) SetStatusCode(v int32) *DeleteStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStudentResponse) SetBody(v *DeleteStudentResponseBody) *DeleteStudentResponse {
	s.Body = v
	return s
}

type DeleteTeacherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTeacherHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTeacherHeaders) SetCommonHeaders(v map[string]*string) *DeleteTeacherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTeacherHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTeacherHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTeacherRequest struct {
	Adviser  *int32  `json:"adviser,omitempty" xml:"adviser,omitempty"`
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s DeleteTeacherRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeacherRequest) SetAdviser(v int32) *DeleteTeacherRequest {
	s.Adviser = &v
	return s
}

func (s *DeleteTeacherRequest) SetOperator(v string) *DeleteTeacherRequest {
	s.Operator = &v
	return s
}

type DeleteTeacherResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTeacherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeacherResponseBody) SetSuccess(v bool) *DeleteTeacherResponseBody {
	s.Success = &v
	return s
}

type DeleteTeacherResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTeacherResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTeacherResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTeacherResponse) GoString() string {
	return s.String()
}

func (s *DeleteTeacherResponse) SetHeaders(v map[string]*string) *DeleteTeacherResponse {
	s.Headers = v
	return s
}

func (s *DeleteTeacherResponse) SetStatusCode(v int32) *DeleteTeacherResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTeacherResponse) SetBody(v *DeleteTeacherResponseBody) *DeleteTeacherResponse {
	s.Body = v
	return s
}

type DeleteUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteUniversityCourseGroupRequest struct {
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	OpUserId        *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s DeleteUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupRequest) SetCourseGroupCode(v string) *DeleteUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *DeleteUniversityCourseGroupRequest) SetOpUserId(v string) *DeleteUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type DeleteUniversityCourseGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupResponseBody) SetResult(v bool) *DeleteUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type DeleteUniversityCourseGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *DeleteUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUniversityCourseGroupResponse) SetStatusCode(v int32) *DeleteUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUniversityCourseGroupResponse) SetBody(v *DeleteUniversityCourseGroupResponseBody) *DeleteUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type DeleteUniversityStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteUniversityStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityStudentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUniversityStudentHeaders) SetCommonHeaders(v map[string]*string) *DeleteUniversityStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUniversityStudentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteUniversityStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteUniversityStudentRequest struct {
	ClassId       *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	OpUserId      *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	StudentUserId *string `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
}

func (s DeleteUniversityStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityStudentRequest) GoString() string {
	return s.String()
}

func (s *DeleteUniversityStudentRequest) SetClassId(v int64) *DeleteUniversityStudentRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteUniversityStudentRequest) SetOpUserId(v string) *DeleteUniversityStudentRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteUniversityStudentRequest) SetStudentUserId(v string) *DeleteUniversityStudentRequest {
	s.StudentUserId = &v
	return s
}

type DeleteUniversityStudentResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteUniversityStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityStudentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUniversityStudentResponseBody) SetResult(v bool) *DeleteUniversityStudentResponseBody {
	s.Result = &v
	return s
}

type DeleteUniversityStudentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUniversityStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUniversityStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityStudentResponse) GoString() string {
	return s.String()
}

func (s *DeleteUniversityStudentResponse) SetHeaders(v map[string]*string) *DeleteUniversityStudentResponse {
	s.Headers = v
	return s
}

func (s *DeleteUniversityStudentResponse) SetStatusCode(v int32) *DeleteUniversityStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUniversityStudentResponse) SetBody(v *DeleteUniversityStudentResponseBody) *DeleteUniversityStudentResponse {
	s.Body = v
	return s
}

type DeleteUniversityTeacherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteUniversityTeacherHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityTeacherHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUniversityTeacherHeaders) SetCommonHeaders(v map[string]*string) *DeleteUniversityTeacherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUniversityTeacherHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteUniversityTeacherHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteUniversityTeacherRequest struct {
	ClassId       *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	OpUserId      *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	Role          *string `json:"role,omitempty" xml:"role,omitempty"`
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s DeleteUniversityTeacherRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityTeacherRequest) GoString() string {
	return s.String()
}

func (s *DeleteUniversityTeacherRequest) SetClassId(v int64) *DeleteUniversityTeacherRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteUniversityTeacherRequest) SetOpUserId(v string) *DeleteUniversityTeacherRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteUniversityTeacherRequest) SetRole(v string) *DeleteUniversityTeacherRequest {
	s.Role = &v
	return s
}

func (s *DeleteUniversityTeacherRequest) SetTeacherUserId(v string) *DeleteUniversityTeacherRequest {
	s.TeacherUserId = &v
	return s
}

type DeleteUniversityTeacherResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteUniversityTeacherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityTeacherResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUniversityTeacherResponseBody) SetResult(v bool) *DeleteUniversityTeacherResponseBody {
	s.Result = &v
	return s
}

type DeleteUniversityTeacherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUniversityTeacherResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUniversityTeacherResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUniversityTeacherResponse) GoString() string {
	return s.String()
}

func (s *DeleteUniversityTeacherResponse) SetHeaders(v map[string]*string) *DeleteUniversityTeacherResponse {
	s.Headers = v
	return s
}

func (s *DeleteUniversityTeacherResponse) SetStatusCode(v int32) *DeleteUniversityTeacherResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUniversityTeacherResponse) SetBody(v *DeleteUniversityTeacherResponseBody) *DeleteUniversityTeacherResponse {
	s.Body = v
	return s
}

type DeviceHeartbeatHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeviceHeartbeatHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceHeartbeatHeaders) GoString() string {
	return s.String()
}

func (s *DeviceHeartbeatHeaders) SetCommonHeaders(v map[string]*string) *DeviceHeartbeatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceHeartbeatHeaders) SetXAcsDingtalkAccessToken(v string) *DeviceHeartbeatHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeviceHeartbeatRequest struct {
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s DeviceHeartbeatRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceHeartbeatRequest) GoString() string {
	return s.String()
}

func (s *DeviceHeartbeatRequest) SetSn(v string) *DeviceHeartbeatRequest {
	s.Sn = &v
	return s
}

type DeviceHeartbeatResponseBody struct {
	Command *int32 `json:"command,omitempty" xml:"command,omitempty"`
}

func (s DeviceHeartbeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceHeartbeatResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceHeartbeatResponseBody) SetCommand(v int32) *DeviceHeartbeatResponseBody {
	s.Command = &v
	return s
}

type DeviceHeartbeatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeviceHeartbeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeviceHeartbeatResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceHeartbeatResponse) GoString() string {
	return s.String()
}

func (s *DeviceHeartbeatResponse) SetHeaders(v map[string]*string) *DeviceHeartbeatResponse {
	s.Headers = v
	return s
}

func (s *DeviceHeartbeatResponse) SetStatusCode(v int32) *DeviceHeartbeatResponse {
	s.StatusCode = &v
	return s
}

func (s *DeviceHeartbeatResponse) SetBody(v *DeviceHeartbeatResponseBody) *DeviceHeartbeatResponse {
	s.Body = v
	return s
}

type EduFindUserRolesByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EduFindUserRolesByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s EduFindUserRolesByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *EduFindUserRolesByUserIdHeaders) SetCommonHeaders(v map[string]*string) *EduFindUserRolesByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EduFindUserRolesByUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *EduFindUserRolesByUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EduFindUserRolesByUserIdRequest struct {
	ClassId    *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	HasOrgRole *bool   `json:"hasOrgRole,omitempty" xml:"hasOrgRole,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EduFindUserRolesByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s EduFindUserRolesByUserIdRequest) GoString() string {
	return s.String()
}

func (s *EduFindUserRolesByUserIdRequest) SetClassId(v int64) *EduFindUserRolesByUserIdRequest {
	s.ClassId = &v
	return s
}

func (s *EduFindUserRolesByUserIdRequest) SetCorpId(v string) *EduFindUserRolesByUserIdRequest {
	s.CorpId = &v
	return s
}

func (s *EduFindUserRolesByUserIdRequest) SetHasOrgRole(v bool) *EduFindUserRolesByUserIdRequest {
	s.HasOrgRole = &v
	return s
}

func (s *EduFindUserRolesByUserIdRequest) SetUserId(v string) *EduFindUserRolesByUserIdRequest {
	s.UserId = &v
	return s
}

type EduFindUserRolesByUserIdResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s EduFindUserRolesByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EduFindUserRolesByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *EduFindUserRolesByUserIdResponseBody) SetResult(v []*string) *EduFindUserRolesByUserIdResponseBody {
	s.Result = v
	return s
}

type EduFindUserRolesByUserIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EduFindUserRolesByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EduFindUserRolesByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s EduFindUserRolesByUserIdResponse) GoString() string {
	return s.String()
}

func (s *EduFindUserRolesByUserIdResponse) SetHeaders(v map[string]*string) *EduFindUserRolesByUserIdResponse {
	s.Headers = v
	return s
}

func (s *EduFindUserRolesByUserIdResponse) SetStatusCode(v int32) *EduFindUserRolesByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *EduFindUserRolesByUserIdResponse) SetBody(v *EduFindUserRolesByUserIdResponseBody) *EduFindUserRolesByUserIdResponse {
	s.Body = v
	return s
}

type EduListUserByFromUserIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EduListUserByFromUserIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s EduListUserByFromUserIdsHeaders) GoString() string {
	return s.String()
}

func (s *EduListUserByFromUserIdsHeaders) SetCommonHeaders(v map[string]*string) *EduListUserByFromUserIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EduListUserByFromUserIdsHeaders) SetXAcsDingtalkAccessToken(v string) *EduListUserByFromUserIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EduListUserByFromUserIdsRequest struct {
	ClassId        *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	CorpId         *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GuardianUserId *string `json:"guardianUserId,omitempty" xml:"guardianUserId,omitempty"`
}

func (s EduListUserByFromUserIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s EduListUserByFromUserIdsRequest) GoString() string {
	return s.String()
}

func (s *EduListUserByFromUserIdsRequest) SetClassId(v int64) *EduListUserByFromUserIdsRequest {
	s.ClassId = &v
	return s
}

func (s *EduListUserByFromUserIdsRequest) SetCorpId(v string) *EduListUserByFromUserIdsRequest {
	s.CorpId = &v
	return s
}

func (s *EduListUserByFromUserIdsRequest) SetGuardianUserId(v string) *EduListUserByFromUserIdsRequest {
	s.GuardianUserId = &v
	return s
}

type EduListUserByFromUserIdsResponseBody struct {
	Result []*EduListUserByFromUserIdsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s EduListUserByFromUserIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EduListUserByFromUserIdsResponseBody) GoString() string {
	return s.String()
}

func (s *EduListUserByFromUserIdsResponseBody) SetResult(v []*EduListUserByFromUserIdsResponseBodyResult) *EduListUserByFromUserIdsResponseBody {
	s.Result = v
	return s
}

type EduListUserByFromUserIdsResponseBodyResult struct {
	CampusId *int64  `json:"campusId,omitempty" xml:"campusId,omitempty"`
	ClassId  *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	GradeId  *int64  `json:"gradeId,omitempty" xml:"gradeId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PeriodId *int64  `json:"periodId,omitempty" xml:"periodId,omitempty"`
	Role     *string `json:"role,omitempty" xml:"role,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EduListUserByFromUserIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EduListUserByFromUserIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetCampusId(v int64) *EduListUserByFromUserIdsResponseBodyResult {
	s.CampusId = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetClassId(v int64) *EduListUserByFromUserIdsResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetGradeId(v int64) *EduListUserByFromUserIdsResponseBodyResult {
	s.GradeId = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetName(v string) *EduListUserByFromUserIdsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetPeriodId(v int64) *EduListUserByFromUserIdsResponseBodyResult {
	s.PeriodId = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetRole(v string) *EduListUserByFromUserIdsResponseBodyResult {
	s.Role = &v
	return s
}

func (s *EduListUserByFromUserIdsResponseBodyResult) SetUserId(v string) *EduListUserByFromUserIdsResponseBodyResult {
	s.UserId = &v
	return s
}

type EduListUserByFromUserIdsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EduListUserByFromUserIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EduListUserByFromUserIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s EduListUserByFromUserIdsResponse) GoString() string {
	return s.String()
}

func (s *EduListUserByFromUserIdsResponse) SetHeaders(v map[string]*string) *EduListUserByFromUserIdsResponse {
	s.Headers = v
	return s
}

func (s *EduListUserByFromUserIdsResponse) SetStatusCode(v int32) *EduListUserByFromUserIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *EduListUserByFromUserIdsResponse) SetBody(v *EduListUserByFromUserIdsResponseBody) *EduListUserByFromUserIdsResponse {
	s.Body = v
	return s
}

type EduTeacherListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EduTeacherListHeaders) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListHeaders) GoString() string {
	return s.String()
}

func (s *EduTeacherListHeaders) SetCommonHeaders(v map[string]*string) *EduTeacherListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EduTeacherListHeaders) SetXAcsDingtalkAccessToken(v string) *EduTeacherListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EduTeacherListRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s EduTeacherListRequest) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListRequest) GoString() string {
	return s.String()
}

func (s *EduTeacherListRequest) SetPageNumber(v int32) *EduTeacherListRequest {
	s.PageNumber = &v
	return s
}

func (s *EduTeacherListRequest) SetPageSize(v int32) *EduTeacherListRequest {
	s.PageSize = &v
	return s
}

type EduTeacherListResponseBody struct {
	Result *EduTeacherListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s EduTeacherListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListResponseBody) GoString() string {
	return s.String()
}

func (s *EduTeacherListResponseBody) SetResult(v *EduTeacherListResponseBodyResult) *EduTeacherListResponseBody {
	s.Result = v
	return s
}

type EduTeacherListResponseBodyResult struct {
	HasMore        *bool                                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	TeacherDetails []*EduTeacherListResponseBodyResultTeacherDetails `json:"teacherDetails,omitempty" xml:"teacherDetails,omitempty" type:"Repeated"`
}

func (s EduTeacherListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EduTeacherListResponseBodyResult) SetHasMore(v bool) *EduTeacherListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *EduTeacherListResponseBodyResult) SetTeacherDetails(v []*EduTeacherListResponseBodyResultTeacherDetails) *EduTeacherListResponseBodyResult {
	s.TeacherDetails = v
	return s
}

type EduTeacherListResponseBodyResultTeacherDetails struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EduTeacherListResponseBodyResultTeacherDetails) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListResponseBodyResultTeacherDetails) GoString() string {
	return s.String()
}

func (s *EduTeacherListResponseBodyResultTeacherDetails) SetName(v string) *EduTeacherListResponseBodyResultTeacherDetails {
	s.Name = &v
	return s
}

func (s *EduTeacherListResponseBodyResultTeacherDetails) SetRole(v string) *EduTeacherListResponseBodyResultTeacherDetails {
	s.Role = &v
	return s
}

func (s *EduTeacherListResponseBodyResultTeacherDetails) SetUnionId(v string) *EduTeacherListResponseBodyResultTeacherDetails {
	s.UnionId = &v
	return s
}

func (s *EduTeacherListResponseBodyResultTeacherDetails) SetUserId(v string) *EduTeacherListResponseBodyResultTeacherDetails {
	s.UserId = &v
	return s
}

type EduTeacherListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EduTeacherListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EduTeacherListResponse) String() string {
	return tea.Prettify(s)
}

func (s EduTeacherListResponse) GoString() string {
	return s.String()
}

func (s *EduTeacherListResponse) SetHeaders(v map[string]*string) *EduTeacherListResponse {
	s.Headers = v
	return s
}

func (s *EduTeacherListResponse) SetStatusCode(v int32) *EduTeacherListResponse {
	s.StatusCode = &v
	return s
}

func (s *EduTeacherListResponse) SetBody(v *EduTeacherListResponseBody) *EduTeacherListResponse {
	s.Body = v
	return s
}

type EndCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EndCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s EndCourseHeaders) GoString() string {
	return s.String()
}

func (s *EndCourseHeaders) SetCommonHeaders(v map[string]*string) *EndCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EndCourseHeaders) SetXAcsDingtalkAccessToken(v string) *EndCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EndCourseRequest struct {
	CourseCode       *string                             `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Ext              *string                             `json:"ext,omitempty" xml:"ext,omitempty"`
	IsvCode          *string                             `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	LivePlayInfoList []*EndCourseRequestLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
	OpUserId         *string                             `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s EndCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s EndCourseRequest) GoString() string {
	return s.String()
}

func (s *EndCourseRequest) SetCourseCode(v string) *EndCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *EndCourseRequest) SetExt(v string) *EndCourseRequest {
	s.Ext = &v
	return s
}

func (s *EndCourseRequest) SetIsvCode(v string) *EndCourseRequest {
	s.IsvCode = &v
	return s
}

func (s *EndCourseRequest) SetLivePlayInfoList(v []*EndCourseRequestLivePlayInfoList) *EndCourseRequest {
	s.LivePlayInfoList = v
	return s
}

func (s *EndCourseRequest) SetOpUserId(v string) *EndCourseRequest {
	s.OpUserId = &v
	return s
}

type EndCourseRequestLivePlayInfoList struct {
	LiveInputUrl     *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	LiveOutputFlvUrl *string `json:"liveOutputFlvUrl,omitempty" xml:"liveOutputFlvUrl,omitempty"`
	LiveOutputHlsUrl *string `json:"liveOutputHlsUrl,omitempty" xml:"liveOutputHlsUrl,omitempty"`
	LiveType         *int32  `json:"liveType,omitempty" xml:"liveType,omitempty"`
	ReplayUrl        *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s EndCourseRequestLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s EndCourseRequestLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveInputUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveOutputFlvUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.LiveOutputFlvUrl = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveOutputHlsUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.LiveOutputHlsUrl = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetLiveType(v int32) *EndCourseRequestLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *EndCourseRequestLivePlayInfoList) SetReplayUrl(v string) *EndCourseRequestLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type EndCourseResponseBody struct {
	UniversityCourseCommonResponse *EndCourseResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s EndCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponseBody) GoString() string {
	return s.String()
}

func (s *EndCourseResponseBody) SetUniversityCourseCommonResponse(v *EndCourseResponseBodyUniversityCourseCommonResponse) *EndCourseResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type EndCourseResponseBodyUniversityCourseCommonResponse struct {
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EndCourseResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *EndCourseResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *EndCourseResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

func (s *EndCourseResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *EndCourseResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

type EndCourseResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EndCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EndCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s EndCourseResponse) GoString() string {
	return s.String()
}

func (s *EndCourseResponse) SetHeaders(v map[string]*string) *EndCourseResponse {
	s.Headers = v
	return s
}

func (s *EndCourseResponse) SetStatusCode(v int32) *EndCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *EndCourseResponse) SetBody(v *EndCourseResponseBody) *EndCourseResponse {
	s.Body = v
	return s
}

type GetBindChildInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBindChildInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBindChildInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetBindChildInfoHeaders) SetCommonHeaders(v map[string]*string) *GetBindChildInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBindChildInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetBindChildInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBindChildInfoRequest struct {
	SchoolCorpId  *string `json:"schoolCorpId,omitempty" xml:"schoolCorpId,omitempty"`
	StudentUserId *string `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetBindChildInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBindChildInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBindChildInfoRequest) SetSchoolCorpId(v string) *GetBindChildInfoRequest {
	s.SchoolCorpId = &v
	return s
}

func (s *GetBindChildInfoRequest) SetStudentUserId(v string) *GetBindChildInfoRequest {
	s.StudentUserId = &v
	return s
}

func (s *GetBindChildInfoRequest) SetUnionId(v string) *GetBindChildInfoRequest {
	s.UnionId = &v
	return s
}

type GetBindChildInfoResponseBody struct {
	ChildUserId   *string `json:"childUserId,omitempty" xml:"childUserId,omitempty"`
	CurrentUserId *string `json:"currentUserId,omitempty" xml:"currentUserId,omitempty"`
	FamilyCorpId  *string `json:"familyCorpId,omitempty" xml:"familyCorpId,omitempty"`
}

func (s GetBindChildInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBindChildInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindChildInfoResponseBody) SetChildUserId(v string) *GetBindChildInfoResponseBody {
	s.ChildUserId = &v
	return s
}

func (s *GetBindChildInfoResponseBody) SetCurrentUserId(v string) *GetBindChildInfoResponseBody {
	s.CurrentUserId = &v
	return s
}

func (s *GetBindChildInfoResponseBody) SetFamilyCorpId(v string) *GetBindChildInfoResponseBody {
	s.FamilyCorpId = &v
	return s
}

type GetBindChildInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBindChildInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBindChildInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBindChildInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBindChildInfoResponse) SetHeaders(v map[string]*string) *GetBindChildInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBindChildInfoResponse) SetStatusCode(v int32) *GetBindChildInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBindChildInfoResponse) SetBody(v *GetBindChildInfoResponseBody) *GetBindChildInfoResponse {
	s.Body = v
	return s
}

type GetDefaultChildHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDefaultChildHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildHeaders) GoString() string {
	return s.String()
}

func (s *GetDefaultChildHeaders) SetCommonHeaders(v map[string]*string) *GetDefaultChildHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDefaultChildHeaders) SetXAcsDingtalkAccessToken(v string) *GetDefaultChildHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDefaultChildResponseBody struct {
	Avatar       *string                                    `json:"avatar,omitempty" xml:"avatar,omitempty"`
	BindStudents []*GetDefaultChildResponseBodyBindStudents `json:"bindStudents,omitempty" xml:"bindStudents,omitempty" type:"Repeated"`
	Grade        *int32                                     `json:"grade,omitempty" xml:"grade,omitempty"`
	Name         *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	Nick         *string                                    `json:"nick,omitempty" xml:"nick,omitempty"`
	Period       *string                                    `json:"period,omitempty" xml:"period,omitempty"`
	UnionId      *string                                    `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetDefaultChildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponseBody) SetAvatar(v string) *GetDefaultChildResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetBindStudents(v []*GetDefaultChildResponseBodyBindStudents) *GetDefaultChildResponseBody {
	s.BindStudents = v
	return s
}

func (s *GetDefaultChildResponseBody) SetGrade(v int32) *GetDefaultChildResponseBody {
	s.Grade = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetName(v string) *GetDefaultChildResponseBody {
	s.Name = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetNick(v string) *GetDefaultChildResponseBody {
	s.Nick = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetPeriod(v string) *GetDefaultChildResponseBody {
	s.Period = &v
	return s
}

func (s *GetDefaultChildResponseBody) SetUnionId(v string) *GetDefaultChildResponseBody {
	s.UnionId = &v
	return s
}

type GetDefaultChildResponseBodyBindStudents struct {
	ClassId *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	CorpId  *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Period  *string `json:"period,omitempty" xml:"period,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s GetDefaultChildResponseBodyBindStudents) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponseBodyBindStudents) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponseBodyBindStudents) SetClassId(v int64) *GetDefaultChildResponseBodyBindStudents {
	s.ClassId = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetCorpId(v string) *GetDefaultChildResponseBodyBindStudents {
	s.CorpId = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetPeriod(v string) *GetDefaultChildResponseBodyBindStudents {
	s.Period = &v
	return s
}

func (s *GetDefaultChildResponseBodyBindStudents) SetStaffId(v string) *GetDefaultChildResponseBodyBindStudents {
	s.StaffId = &v
	return s
}

type GetDefaultChildResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDefaultChildResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultChildResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultChildResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultChildResponse) SetHeaders(v map[string]*string) *GetDefaultChildResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultChildResponse) SetStatusCode(v int32) *GetDefaultChildResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultChildResponse) SetBody(v *GetDefaultChildResponseBody) *GetDefaultChildResponse {
	s.Body = v
	return s
}

type GetEduUserIdentityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEduUserIdentityHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEduUserIdentityHeaders) GoString() string {
	return s.String()
}

func (s *GetEduUserIdentityHeaders) SetCommonHeaders(v map[string]*string) *GetEduUserIdentityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEduUserIdentityHeaders) SetXAcsDingtalkAccessToken(v string) *GetEduUserIdentityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEduUserIdentityRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetEduUserIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEduUserIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetEduUserIdentityRequest) SetUnionId(v string) *GetEduUserIdentityRequest {
	s.UnionId = &v
	return s
}

type GetEduUserIdentityResponseBody struct {
	Data    *GetEduUserIdentityResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEduUserIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEduUserIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetEduUserIdentityResponseBody) SetData(v *GetEduUserIdentityResponseBodyData) *GetEduUserIdentityResponseBody {
	s.Data = v
	return s
}

func (s *GetEduUserIdentityResponseBody) SetSuccess(v bool) *GetEduUserIdentityResponseBody {
	s.Success = &v
	return s
}

type GetEduUserIdentityResponseBodyData struct {
	MatchGuardianRule *bool   `json:"matchGuardianRule,omitempty" xml:"matchGuardianRule,omitempty"`
	MatchTeacherRule  *bool   `json:"matchTeacherRule,omitempty" xml:"matchTeacherRule,omitempty"`
	UnionId           *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetEduUserIdentityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEduUserIdentityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEduUserIdentityResponseBodyData) SetMatchGuardianRule(v bool) *GetEduUserIdentityResponseBodyData {
	s.MatchGuardianRule = &v
	return s
}

func (s *GetEduUserIdentityResponseBodyData) SetMatchTeacherRule(v bool) *GetEduUserIdentityResponseBodyData {
	s.MatchTeacherRule = &v
	return s
}

func (s *GetEduUserIdentityResponseBodyData) SetUnionId(v string) *GetEduUserIdentityResponseBodyData {
	s.UnionId = &v
	return s
}

type GetEduUserIdentityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEduUserIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEduUserIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEduUserIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetEduUserIdentityResponse) SetHeaders(v map[string]*string) *GetEduUserIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetEduUserIdentityResponse) SetStatusCode(v int32) *GetEduUserIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEduUserIdentityResponse) SetBody(v *GetEduUserIdentityResponseBody) *GetEduUserIdentityResponse {
	s.Body = v
	return s
}

type GetOpenCourseDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOpenCourseDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailHeaders) SetCommonHeaders(v map[string]*string) *GetOpenCourseDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenCourseDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetOpenCourseDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOpenCourseDetailResponseBody struct {
	CourseId     *string                                   `json:"courseId,omitempty" xml:"courseId,omitempty"`
	CourseType   *int64                                    `json:"courseType,omitempty" xml:"courseType,omitempty"`
	CoverUrl     *string                                   `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	Introduction *string                                   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	PushModel    *GetOpenCourseDetailResponseBodyPushModel `json:"pushModel,omitempty" xml:"pushModel,omitempty" type:"Struct"`
	StartTime    *int64                                    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TeacherId    *string                                   `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
	TeacherName  *string                                   `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	Title        *string                                   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetOpenCourseDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponseBody) SetCourseId(v string) *GetOpenCourseDetailResponseBody {
	s.CourseId = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetCourseType(v int64) *GetOpenCourseDetailResponseBody {
	s.CourseType = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetCoverUrl(v string) *GetOpenCourseDetailResponseBody {
	s.CoverUrl = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetIntroduction(v string) *GetOpenCourseDetailResponseBody {
	s.Introduction = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetPushModel(v *GetOpenCourseDetailResponseBodyPushModel) *GetOpenCourseDetailResponseBody {
	s.PushModel = v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetStartTime(v int64) *GetOpenCourseDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTeacherId(v string) *GetOpenCourseDetailResponseBody {
	s.TeacherId = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTeacherName(v string) *GetOpenCourseDetailResponseBody {
	s.TeacherName = &v
	return s
}

func (s *GetOpenCourseDetailResponseBody) SetTitle(v string) *GetOpenCourseDetailResponseBody {
	s.Title = &v
	return s
}

type GetOpenCourseDetailResponseBodyPushModel struct {
	PushOrgNameList  []*string `json:"pushOrgNameList,omitempty" xml:"pushOrgNameList,omitempty" type:"Repeated"`
	PushRoleNameList []*string `json:"pushRoleNameList,omitempty" xml:"pushRoleNameList,omitempty" type:"Repeated"`
}

func (s GetOpenCourseDetailResponseBodyPushModel) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponseBodyPushModel) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponseBodyPushModel) SetPushOrgNameList(v []*string) *GetOpenCourseDetailResponseBodyPushModel {
	s.PushOrgNameList = v
	return s
}

func (s *GetOpenCourseDetailResponseBodyPushModel) SetPushRoleNameList(v []*string) *GetOpenCourseDetailResponseBodyPushModel {
	s.PushRoleNameList = v
	return s
}

type GetOpenCourseDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOpenCourseDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenCourseDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCourseDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOpenCourseDetailResponse) SetHeaders(v map[string]*string) *GetOpenCourseDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOpenCourseDetailResponse) SetStatusCode(v int32) *GetOpenCourseDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenCourseDetailResponse) SetBody(v *GetOpenCourseDetailResponseBody) *GetOpenCourseDetailResponse {
	s.Body = v
	return s
}

type GetOpenCoursesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOpenCoursesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesHeaders) SetCommonHeaders(v map[string]*string) *GetOpenCoursesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenCoursesHeaders) SetXAcsDingtalkAccessToken(v string) *GetOpenCoursesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOpenCoursesRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetOpenCoursesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesRequest) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesRequest) SetPageNumber(v int64) *GetOpenCoursesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOpenCoursesRequest) SetPageSize(v int64) *GetOpenCoursesRequest {
	s.PageSize = &v
	return s
}

type GetOpenCoursesResponseBody struct {
	CourseList []*GetOpenCoursesResponseBodyCourseList `json:"courseList,omitempty" xml:"courseList,omitempty" type:"Repeated"`
	TotalCount *int64                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetOpenCoursesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponseBody) SetCourseList(v []*GetOpenCoursesResponseBodyCourseList) *GetOpenCoursesResponseBody {
	s.CourseList = v
	return s
}

func (s *GetOpenCoursesResponseBody) SetTotalCount(v int64) *GetOpenCoursesResponseBody {
	s.TotalCount = &v
	return s
}

type GetOpenCoursesResponseBodyCourseList struct {
	CourseId    *string `json:"courseId,omitempty" xml:"courseId,omitempty"`
	CoverUrl    *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	FeedType    *int64  `json:"feedType,omitempty" xml:"feedType,omitempty"`
	JumpUrl     *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	StartTime   *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TeacherId   *string `json:"teacherId,omitempty" xml:"teacherId,omitempty"`
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetOpenCoursesResponseBodyCourseList) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponseBodyCourseList) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponseBodyCourseList) SetCourseId(v string) *GetOpenCoursesResponseBodyCourseList {
	s.CourseId = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetCoverUrl(v string) *GetOpenCoursesResponseBodyCourseList {
	s.CoverUrl = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetFeedType(v int64) *GetOpenCoursesResponseBodyCourseList {
	s.FeedType = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetJumpUrl(v string) *GetOpenCoursesResponseBodyCourseList {
	s.JumpUrl = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetStartTime(v int64) *GetOpenCoursesResponseBodyCourseList {
	s.StartTime = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTeacherId(v string) *GetOpenCoursesResponseBodyCourseList {
	s.TeacherId = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTeacherName(v string) *GetOpenCoursesResponseBodyCourseList {
	s.TeacherName = &v
	return s
}

func (s *GetOpenCoursesResponseBodyCourseList) SetTitle(v string) *GetOpenCoursesResponseBodyCourseList {
	s.Title = &v
	return s
}

type GetOpenCoursesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOpenCoursesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenCoursesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenCoursesResponse) GoString() string {
	return s.String()
}

func (s *GetOpenCoursesResponse) SetHeaders(v map[string]*string) *GetOpenCoursesResponse {
	s.Headers = v
	return s
}

func (s *GetOpenCoursesResponse) SetStatusCode(v int32) *GetOpenCoursesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpenCoursesResponse) SetBody(v *GetOpenCoursesResponseBody) *GetOpenCoursesResponse {
	s.Body = v
	return s
}

type GetPointActionRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPointActionRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordHeaders) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordHeaders) SetCommonHeaders(v map[string]*string) *GetPointActionRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPointActionRecordHeaders) SetXAcsDingtalkAccessToken(v string) *GetPointActionRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPointActionRecordRequest struct {
	Body *GetPointActionRecordRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s GetPointActionRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordRequest) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordRequest) SetBody(v *GetPointActionRecordRequestBody) *GetPointActionRecordRequest {
	s.Body = v
	return s
}

type GetPointActionRecordRequestBody struct {
	BizId     *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	OwnerId   *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	PointType *string `json:"pointType,omitempty" xml:"pointType,omitempty"`
}

func (s GetPointActionRecordRequestBody) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordRequestBody) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordRequestBody) SetBizId(v string) *GetPointActionRecordRequestBody {
	s.BizId = &v
	return s
}

func (s *GetPointActionRecordRequestBody) SetOwnerId(v string) *GetPointActionRecordRequestBody {
	s.OwnerId = &v
	return s
}

func (s *GetPointActionRecordRequestBody) SetPointType(v string) *GetPointActionRecordRequestBody {
	s.PointType = &v
	return s
}

type GetPointActionRecordShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPointActionRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordShrinkRequest) SetBodyShrink(v string) *GetPointActionRecordShrinkRequest {
	s.BodyShrink = &v
	return s
}

type GetPointActionRecordResponseBody struct {
	Result  *GetPointActionRecordResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPointActionRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordResponseBody) SetResult(v *GetPointActionRecordResponseBodyResult) *GetPointActionRecordResponseBody {
	s.Result = v
	return s
}

func (s *GetPointActionRecordResponseBody) SetSuccess(v bool) *GetPointActionRecordResponseBody {
	s.Success = &v
	return s
}

type GetPointActionRecordResponseBodyResult struct {
	ActionTime *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	Quantity   *int64  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPointActionRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordResponseBodyResult) SetActionTime(v string) *GetPointActionRecordResponseBodyResult {
	s.ActionTime = &v
	return s
}

func (s *GetPointActionRecordResponseBodyResult) SetQuantity(v int64) *GetPointActionRecordResponseBodyResult {
	s.Quantity = &v
	return s
}

func (s *GetPointActionRecordResponseBodyResult) SetStatus(v string) *GetPointActionRecordResponseBodyResult {
	s.Status = &v
	return s
}

type GetPointActionRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPointActionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPointActionRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPointActionRecordResponse) GoString() string {
	return s.String()
}

func (s *GetPointActionRecordResponse) SetHeaders(v map[string]*string) *GetPointActionRecordResponse {
	s.Headers = v
	return s
}

func (s *GetPointActionRecordResponse) SetStatusCode(v int32) *GetPointActionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPointActionRecordResponse) SetBody(v *GetPointActionRecordResponseBody) *GetPointActionRecordResponse {
	s.Body = v
	return s
}

type GetPointInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPointInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPointInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPointInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPointInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPointInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPointInfoRequest struct {
	PointType *string `json:"pointType,omitempty" xml:"pointType,omitempty"`
}

func (s GetPointInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPointInfoRequest) SetPointType(v string) *GetPointInfoRequest {
	s.PointType = &v
	return s
}

type GetPointInfoResponseBody struct {
	Result  *GetPointInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPointInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponseBody) SetResult(v *GetPointInfoResponseBodyResult) *GetPointInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetPointInfoResponseBody) SetSuccess(v bool) *GetPointInfoResponseBody {
	s.Success = &v
	return s
}

type GetPointInfoResponseBodyResult struct {
	AvailableQuota *int64  `json:"availableQuota,omitempty" xml:"availableQuota,omitempty"`
	EndTime        *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime      *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetPointInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponseBodyResult) SetAvailableQuota(v int64) *GetPointInfoResponseBodyResult {
	s.AvailableQuota = &v
	return s
}

func (s *GetPointInfoResponseBodyResult) SetEndTime(v string) *GetPointInfoResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetPointInfoResponseBodyResult) SetStartTime(v string) *GetPointInfoResponseBodyResult {
	s.StartTime = &v
	return s
}

type GetPointInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPointInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPointInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPointInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPointInfoResponse) SetHeaders(v map[string]*string) *GetPointInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPointInfoResponse) SetStatusCode(v int32) *GetPointInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPointInfoResponse) SetBody(v *GetPointInfoResponseBody) *GetPointInfoResponse {
	s.Body = v
	return s
}

type GetRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *GetRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *GetRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRemoteClassCourseRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s GetRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseRequest) SetOperator(v string) *GetRemoteClassCourseRequest {
	s.Operator = &v
	return s
}

type GetRemoteClassCourseResponseBody struct {
	Result  *GetRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBody) SetResult(v *GetRemoteClassCourseResponseBodyResult) *GetRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

func (s *GetRemoteClassCourseResponseBody) SetSuccess(v bool) *GetRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type GetRemoteClassCourseResponseBodyResult struct {
	AttendParticipants  []*GetRemoteClassCourseResponseBodyResultAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	CanEdit             *bool                                                       `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	CourseCode          *string                                                     `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	CourseName          *string                                                     `json:"courseName,omitempty" xml:"courseName,omitempty"`
	EndTime             *int64                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	LiveUrl             *string                                                     `json:"liveUrl,omitempty" xml:"liveUrl,omitempty"`
	RecordInfos         []*GetRemoteClassCourseResponseBodyResultRecordInfos        `json:"recordInfos,omitempty" xml:"recordInfos,omitempty" type:"Repeated"`
	RoomStatus          *int32                                                      `json:"roomStatus,omitempty" xml:"roomStatus,omitempty"`
	StartTime           *int64                                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status              *int32                                                      `json:"status,omitempty" xml:"status,omitempty"`
	TeachingParticipant *GetRemoteClassCourseResponseBodyResultTeachingParticipant  `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
}

func (s GetRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResult) SetAttendParticipants(v []*GetRemoteClassCourseResponseBodyResultAttendParticipants) *GetRemoteClassCourseResponseBodyResult {
	s.AttendParticipants = v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCanEdit(v bool) *GetRemoteClassCourseResponseBodyResult {
	s.CanEdit = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *GetRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetCourseName(v string) *GetRemoteClassCourseResponseBodyResult {
	s.CourseName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetEndTime(v int64) *GetRemoteClassCourseResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetLiveUrl(v string) *GetRemoteClassCourseResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetRecordInfos(v []*GetRemoteClassCourseResponseBodyResultRecordInfos) *GetRemoteClassCourseResponseBodyResult {
	s.RecordInfos = v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetRoomStatus(v int32) *GetRemoteClassCourseResponseBodyResult {
	s.RoomStatus = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetStartTime(v int64) *GetRemoteClassCourseResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetStatus(v int32) *GetRemoteClassCourseResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResult) SetTeachingParticipant(v *GetRemoteClassCourseResponseBodyResultTeachingParticipant) *GetRemoteClassCourseResponseBodyResult {
	s.TeachingParticipant = v
	return s
}

type GetRemoteClassCourseResponseBodyResultAttendParticipants struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrgName         *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ParticipantId   *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
}

func (s GetRemoteClassCourseResponseBodyResultAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResultAttendParticipants) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetCorpId(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetOrgName(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.OrgName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantId(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantName(v string) *GetRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantName = &v
	return s
}

type GetRemoteClassCourseResponseBodyResultRecordInfos struct {
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StopTime  *string `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRemoteClassCourseResponseBodyResultRecordInfos) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResultRecordInfos) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResultRecordInfos) SetStartTime(v string) *GetRemoteClassCourseResponseBodyResultRecordInfos {
	s.StartTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultRecordInfos) SetStopTime(v string) *GetRemoteClassCourseResponseBodyResultRecordInfos {
	s.StopTime = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultRecordInfos) SetUrl(v string) *GetRemoteClassCourseResponseBodyResultRecordInfos {
	s.Url = &v
	return s
}

type GetRemoteClassCourseResponseBodyResultTeachingParticipant struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrgName         *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ParticipantId   *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
}

func (s GetRemoteClassCourseResponseBodyResultTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponseBodyResultTeachingParticipant) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetCorpId(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetOrgName(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.OrgName = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantId(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *GetRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantName(v string) *GetRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantName = &v
	return s
}

type GetRemoteClassCourseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *GetRemoteClassCourseResponse) SetHeaders(v map[string]*string) *GetRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *GetRemoteClassCourseResponse) SetStatusCode(v int32) *GetRemoteClassCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemoteClassCourseResponse) SetBody(v *GetRemoteClassCourseResponseBody) *GetRemoteClassCourseResponse {
	s.Body = v
	return s
}

type GetShareRoleMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetShareRoleMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersHeaders) SetCommonHeaders(v map[string]*string) *GetShareRoleMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetShareRoleMembersHeaders) SetXAcsDingtalkAccessToken(v string) *GetShareRoleMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetShareRoleMembersResponseBody struct {
	Result []*GetShareRoleMembersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetShareRoleMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponseBody) SetResult(v []*GetShareRoleMembersResponseBodyResult) *GetShareRoleMembersResponseBody {
	s.Result = v
	return s
}

type GetShareRoleMembersResponseBodyResult struct {
	CorpId                     *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	MemberUserIdListInTrunkOrg []*string `json:"memberUserIdListInTrunkOrg,omitempty" xml:"memberUserIdListInTrunkOrg,omitempty" type:"Repeated"`
}

func (s GetShareRoleMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponseBodyResult) SetCorpId(v string) *GetShareRoleMembersResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetShareRoleMembersResponseBodyResult) SetMemberUserIdListInTrunkOrg(v []*string) *GetShareRoleMembersResponseBodyResult {
	s.MemberUserIdListInTrunkOrg = v
	return s
}

type GetShareRoleMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShareRoleMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShareRoleMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRoleMembersResponse) GoString() string {
	return s.String()
}

func (s *GetShareRoleMembersResponse) SetHeaders(v map[string]*string) *GetShareRoleMembersResponse {
	s.Headers = v
	return s
}

func (s *GetShareRoleMembersResponse) SetStatusCode(v int32) *GetShareRoleMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareRoleMembersResponse) SetBody(v *GetShareRoleMembersResponseBody) *GetShareRoleMembersResponse {
	s.Body = v
	return s
}

type GetShareRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetShareRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesHeaders) GoString() string {
	return s.String()
}

func (s *GetShareRolesHeaders) SetCommonHeaders(v map[string]*string) *GetShareRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetShareRolesHeaders) SetXAcsDingtalkAccessToken(v string) *GetShareRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetShareRolesResponseBody struct {
	Result []*GetShareRolesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetShareRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponseBody) SetResult(v []*GetShareRolesResponseBodyResult) *GetShareRolesResponseBody {
	s.Result = v
	return s
}

type GetShareRolesResponseBodyResult struct {
	ShareRoleCode *string `json:"shareRoleCode,omitempty" xml:"shareRoleCode,omitempty"`
	ShareRoleName *string `json:"shareRoleName,omitempty" xml:"shareRoleName,omitempty"`
}

func (s GetShareRolesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponseBodyResult) SetShareRoleCode(v string) *GetShareRolesResponseBodyResult {
	s.ShareRoleCode = &v
	return s
}

func (s *GetShareRolesResponseBodyResult) SetShareRoleName(v string) *GetShareRolesResponseBodyResult {
	s.ShareRoleName = &v
	return s
}

type GetShareRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShareRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetShareRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareRolesResponse) GoString() string {
	return s.String()
}

func (s *GetShareRolesResponse) SetHeaders(v map[string]*string) *GetShareRolesResponse {
	s.Headers = v
	return s
}

func (s *GetShareRolesResponse) SetStatusCode(v int32) *GetShareRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareRolesResponse) SetBody(v *GetShareRolesResponseBody) *GetShareRolesResponse {
	s.Body = v
	return s
}

type GetTaskListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskListHeaders) SetCommonHeaders(v map[string]*string) *GetTaskListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskListHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskListRequest struct {
	Operator   *string `json:"operator,omitempty" xml:"operator,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TaskYear   *int64  `json:"taskYear,omitempty" xml:"taskYear,omitempty"`
}

func (s GetTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetTaskListRequest) SetOperator(v string) *GetTaskListRequest {
	s.Operator = &v
	return s
}

func (s *GetTaskListRequest) SetPageNumber(v int64) *GetTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskListRequest) SetPageSize(v int64) *GetTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskListRequest) SetTaskYear(v int64) *GetTaskListRequest {
	s.TaskYear = &v
	return s
}

type GetTaskListResponseBody struct {
	Count    *int64                             `json:"count,omitempty" xml:"count,omitempty"`
	TaskList []*GetTaskListResponseBodyTaskList `json:"taskList,omitempty" xml:"taskList,omitempty" type:"Repeated"`
}

func (s GetTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskListResponseBody) SetCount(v int64) *GetTaskListResponseBody {
	s.Count = &v
	return s
}

func (s *GetTaskListResponseBody) SetTaskList(v []*GetTaskListResponseBodyTaskList) *GetTaskListResponseBody {
	s.TaskList = v
	return s
}

type GetTaskListResponseBodyTaskList struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	TaskId   *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskYear *int64  `json:"taskYear,omitempty" xml:"taskYear,omitempty"`
}

func (s GetTaskListResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *GetTaskListResponseBodyTaskList) SetName(v string) *GetTaskListResponseBodyTaskList {
	s.Name = &v
	return s
}

func (s *GetTaskListResponseBodyTaskList) SetTaskId(v int64) *GetTaskListResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *GetTaskListResponseBodyTaskList) SetTaskYear(v int64) *GetTaskListResponseBodyTaskList {
	s.TaskYear = &v
	return s
}

type GetTaskListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetTaskListResponse) SetHeaders(v map[string]*string) *GetTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetTaskListResponse) SetStatusCode(v int32) *GetTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskListResponse) SetBody(v *GetTaskListResponseBody) *GetTaskListResponse {
	s.Body = v
	return s
}

type GetTaskStudentListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskStudentListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStudentListHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskStudentListHeaders) SetCommonHeaders(v map[string]*string) *GetTaskStudentListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskStudentListHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskStudentListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskStudentListRequest struct {
	Operator   *string `json:"operator,omitempty" xml:"operator,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TaskId     *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskStudentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStudentListRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStudentListRequest) SetOperator(v string) *GetTaskStudentListRequest {
	s.Operator = &v
	return s
}

func (s *GetTaskStudentListRequest) SetPageNumber(v int64) *GetTaskStudentListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskStudentListRequest) SetPageSize(v int64) *GetTaskStudentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskStudentListRequest) SetTaskId(v int64) *GetTaskStudentListRequest {
	s.TaskId = &v
	return s
}

type GetTaskStudentListResponseBody struct {
	Count       *int64                                       `json:"count,omitempty" xml:"count,omitempty"`
	StudentList []*GetTaskStudentListResponseBodyStudentList `json:"studentList,omitempty" xml:"studentList,omitempty" type:"Repeated"`
	TaskId      *int64                                       `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskStudentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStudentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStudentListResponseBody) SetCount(v int64) *GetTaskStudentListResponseBody {
	s.Count = &v
	return s
}

func (s *GetTaskStudentListResponseBody) SetStudentList(v []*GetTaskStudentListResponseBodyStudentList) *GetTaskStudentListResponseBody {
	s.StudentList = v
	return s
}

func (s *GetTaskStudentListResponseBody) SetTaskId(v int64) *GetTaskStudentListResponseBody {
	s.TaskId = &v
	return s
}

type GetTaskStudentListResponseBodyStudentList struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Sexuality *string `json:"sexuality,omitempty" xml:"sexuality,omitempty"`
	StudentId *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s GetTaskStudentListResponseBodyStudentList) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStudentListResponseBodyStudentList) GoString() string {
	return s.String()
}

func (s *GetTaskStudentListResponseBodyStudentList) SetName(v string) *GetTaskStudentListResponseBodyStudentList {
	s.Name = &v
	return s
}

func (s *GetTaskStudentListResponseBodyStudentList) SetSexuality(v string) *GetTaskStudentListResponseBodyStudentList {
	s.Sexuality = &v
	return s
}

func (s *GetTaskStudentListResponseBodyStudentList) SetStudentId(v int64) *GetTaskStudentListResponseBodyStudentList {
	s.StudentId = &v
	return s
}

type GetTaskStudentListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskStudentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskStudentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStudentListResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStudentListResponse) SetHeaders(v map[string]*string) *GetTaskStudentListResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStudentListResponse) SetStatusCode(v int32) *GetTaskStudentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStudentListResponse) SetBody(v *GetTaskStudentListResponseBody) *GetTaskStudentListResponse {
	s.Body = v
	return s
}

type InitCoursesOfClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitCoursesOfClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassHeaders) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassHeaders) SetCommonHeaders(v map[string]*string) *InitCoursesOfClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitCoursesOfClassHeaders) SetXAcsDingtalkAccessToken(v string) *InitCoursesOfClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitCoursesOfClassRequest struct {
	Courses       []*InitCoursesOfClassRequestCourses     `json:"courses,omitempty" xml:"courses,omitempty" type:"Repeated"`
	SectionConfig *InitCoursesOfClassRequestSectionConfig `json:"sectionConfig,omitempty" xml:"sectionConfig,omitempty" type:"Struct"`
	OpUserId      *string                                 `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s InitCoursesOfClassRequest) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequest) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequest) SetCourses(v []*InitCoursesOfClassRequestCourses) *InitCoursesOfClassRequest {
	s.Courses = v
	return s
}

func (s *InitCoursesOfClassRequest) SetSectionConfig(v *InitCoursesOfClassRequestSectionConfig) *InitCoursesOfClassRequest {
	s.SectionConfig = v
	return s
}

func (s *InitCoursesOfClassRequest) SetOpUserId(v string) *InitCoursesOfClassRequest {
	s.OpUserId = &v
	return s
}

type InitCoursesOfClassRequestCourses struct {
	CourseName      *string                                       `json:"courseName,omitempty" xml:"courseName,omitempty"`
	CreatorName     *string                                       `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	DateModel       *InitCoursesOfClassRequestCoursesDateModel    `json:"dateModel,omitempty" xml:"dateModel,omitempty" type:"Struct"`
	Location        *string                                       `json:"location,omitempty" xml:"location,omitempty"`
	SectionModel    *InitCoursesOfClassRequestCoursesSectionModel `json:"sectionModel,omitempty" xml:"sectionModel,omitempty" type:"Struct"`
	TeacherStaffIds []*string                                     `json:"teacherStaffIds,omitempty" xml:"teacherStaffIds,omitempty" type:"Repeated"`
}

func (s InitCoursesOfClassRequestCourses) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCourses) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCourses) SetCourseName(v string) *InitCoursesOfClassRequestCourses {
	s.CourseName = &v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetCreatorName(v string) *InitCoursesOfClassRequestCourses {
	s.CreatorName = &v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetDateModel(v *InitCoursesOfClassRequestCoursesDateModel) *InitCoursesOfClassRequestCourses {
	s.DateModel = v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetLocation(v string) *InitCoursesOfClassRequestCourses {
	s.Location = &v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetSectionModel(v *InitCoursesOfClassRequestCoursesSectionModel) *InitCoursesOfClassRequestCourses {
	s.SectionModel = v
	return s
}

func (s *InitCoursesOfClassRequestCourses) SetTeacherStaffIds(v []*string) *InitCoursesOfClassRequestCourses {
	s.TeacherStaffIds = v
	return s
}

type InitCoursesOfClassRequestCoursesDateModel struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s InitCoursesOfClassRequestCoursesDateModel) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCoursesDateModel) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetDayOfMonth(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.DayOfMonth = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetMonth(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesDateModel) SetYear(v int32) *InitCoursesOfClassRequestCoursesDateModel {
	s.Year = &v
	return s
}

type InitCoursesOfClassRequestCoursesSectionModel struct {
	SectionIndex *int32  `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName  *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
}

func (s InitCoursesOfClassRequestCoursesSectionModel) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestCoursesSectionModel) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestCoursesSectionModel) SetSectionIndex(v int32) *InitCoursesOfClassRequestCoursesSectionModel {
	s.SectionIndex = &v
	return s
}

func (s *InitCoursesOfClassRequestCoursesSectionModel) SetSectionName(v string) *InitCoursesOfClassRequestCoursesSectionModel {
	s.SectionName = &v
	return s
}

type InitCoursesOfClassRequestSectionConfig struct {
	End           *InitCoursesOfClassRequestSectionConfigEnd             `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionModels []*InitCoursesOfClassRequestSectionConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	Start         *InitCoursesOfClassRequestSectionConfigStart           `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s InitCoursesOfClassRequestSectionConfig) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfig) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfig) SetEnd(v *InitCoursesOfClassRequestSectionConfigEnd) *InitCoursesOfClassRequestSectionConfig {
	s.End = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfig) SetSectionModels(v []*InitCoursesOfClassRequestSectionConfigSectionModels) *InitCoursesOfClassRequestSectionConfig {
	s.SectionModels = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfig) SetStart(v *InitCoursesOfClassRequestSectionConfigStart) *InitCoursesOfClassRequestSectionConfig {
	s.Start = v
	return s
}

type InitCoursesOfClassRequestSectionConfigEnd struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigEnd) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigEnd) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetDayOfMonth(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.DayOfMonth = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetMonth(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigEnd) SetYear(v int32) *InitCoursesOfClassRequestSectionConfigEnd {
	s.Year = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModels struct {
	End          *InitCoursesOfClassRequestSectionConfigSectionModelsEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionIndex *int32                                                    `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionType  *string                                                   `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	Start        *InitCoursesOfClassRequestSectionConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModels) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetEnd(v *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.End = v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetSectionIndex(v int32) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetSectionType(v string) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModels) SetStart(v *InitCoursesOfClassRequestSectionConfigSectionModelsStart) *InitCoursesOfClassRequestSectionConfigSectionModels {
	s.Start = v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModelsEnd struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) SetHour(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsEnd) SetMin(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Min = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigSectionModelsStart struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsStart) SetHour(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigSectionModelsStart) SetMin(v int32) *InitCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Min = &v
	return s
}

type InitCoursesOfClassRequestSectionConfigStart struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s InitCoursesOfClassRequestSectionConfigStart) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassRequestSectionConfigStart) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetDayOfMonth(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.DayOfMonth = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetMonth(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.Month = &v
	return s
}

func (s *InitCoursesOfClassRequestSectionConfigStart) SetYear(v int32) *InitCoursesOfClassRequestSectionConfigStart {
	s.Year = &v
	return s
}

type InitCoursesOfClassResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InitCoursesOfClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassResponseBody) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassResponseBody) SetResult(v bool) *InitCoursesOfClassResponseBody {
	s.Result = &v
	return s
}

type InitCoursesOfClassResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitCoursesOfClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitCoursesOfClassResponse) String() string {
	return tea.Prettify(s)
}

func (s InitCoursesOfClassResponse) GoString() string {
	return s.String()
}

func (s *InitCoursesOfClassResponse) SetHeaders(v map[string]*string) *InitCoursesOfClassResponse {
	s.Headers = v
	return s
}

func (s *InitCoursesOfClassResponse) SetStatusCode(v int32) *InitCoursesOfClassResponse {
	s.StatusCode = &v
	return s
}

func (s *InitCoursesOfClassResponse) SetBody(v *InitCoursesOfClassResponseBody) *InitCoursesOfClassResponse {
	s.Body = v
	return s
}

type InitDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceHeaders) GoString() string {
	return s.String()
}

func (s *InitDeviceHeaders) SetCommonHeaders(v map[string]*string) *InitDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *InitDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitDeviceRequest struct {
	EncryptPubKey *string `json:"encryptPubKey,omitempty" xml:"encryptPubKey,omitempty"`
	Signature     *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn            *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp     *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InitDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceRequest) GoString() string {
	return s.String()
}

func (s *InitDeviceRequest) SetEncryptPubKey(v string) *InitDeviceRequest {
	s.EncryptPubKey = &v
	return s
}

func (s *InitDeviceRequest) SetSignature(v string) *InitDeviceRequest {
	s.Signature = &v
	return s
}

func (s *InitDeviceRequest) SetSn(v string) *InitDeviceRequest {
	s.Sn = &v
	return s
}

func (s *InitDeviceRequest) SetTimestamp(v int64) *InitDeviceRequest {
	s.Timestamp = &v
	return s
}

func (s *InitDeviceRequest) SetVersion(v string) *InitDeviceRequest {
	s.Version = &v
	return s
}

type InitDeviceResponseBody struct {
	SuccessInfo *string `json:"successInfo,omitempty" xml:"successInfo,omitempty"`
}

func (s InitDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *InitDeviceResponseBody) SetSuccessInfo(v string) *InitDeviceResponseBody {
	s.SuccessInfo = &v
	return s
}

type InitDeviceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitDeviceResponse) GoString() string {
	return s.String()
}

func (s *InitDeviceResponse) SetHeaders(v map[string]*string) *InitDeviceResponse {
	s.Headers = v
	return s
}

func (s *InitDeviceResponse) SetStatusCode(v int32) *InitDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *InitDeviceResponse) SetBody(v *InitDeviceResponseBody) *InitDeviceResponse {
	s.Body = v
	return s
}

type InitVPaasDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InitVPaasDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitVPaasDeviceHeaders) GoString() string {
	return s.String()
}

func (s *InitVPaasDeviceHeaders) SetCommonHeaders(v map[string]*string) *InitVPaasDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitVPaasDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *InitVPaasDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InitVPaasDeviceRequest struct {
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InitVPaasDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitVPaasDeviceRequest) GoString() string {
	return s.String()
}

func (s *InitVPaasDeviceRequest) SetSn(v string) *InitVPaasDeviceRequest {
	s.Sn = &v
	return s
}

func (s *InitVPaasDeviceRequest) SetTimestamp(v int64) *InitVPaasDeviceRequest {
	s.Timestamp = &v
	return s
}

func (s *InitVPaasDeviceRequest) SetType(v string) *InitVPaasDeviceRequest {
	s.Type = &v
	return s
}

type InitVPaasDeviceResponseBody struct {
	Pspk *string `json:"pspk,omitempty" xml:"pspk,omitempty"`
}

func (s InitVPaasDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitVPaasDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *InitVPaasDeviceResponseBody) SetPspk(v string) *InitVPaasDeviceResponseBody {
	s.Pspk = &v
	return s
}

type InitVPaasDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitVPaasDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitVPaasDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitVPaasDeviceResponse) GoString() string {
	return s.String()
}

func (s *InitVPaasDeviceResponse) SetHeaders(v map[string]*string) *InitVPaasDeviceResponse {
	s.Headers = v
	return s
}

func (s *InitVPaasDeviceResponse) SetStatusCode(v int32) *InitVPaasDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *InitVPaasDeviceResponse) SetBody(v *InitVPaasDeviceResponseBody) *InitVPaasDeviceResponse {
	s.Body = v
	return s
}

type InsertSectionConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertSectionConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigHeaders) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigHeaders) SetCommonHeaders(v map[string]*string) *InsertSectionConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertSectionConfigHeaders) SetXAcsDingtalkAccessToken(v string) *InsertSectionConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertSectionConfigRequest struct {
	End           *InsertSectionConfigRequestEnd             `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ScheduleName  *string                                    `json:"scheduleName,omitempty" xml:"scheduleName,omitempty"`
	SectionModels []*InsertSectionConfigRequestSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	Start         *InsertSectionConfigRequestStart           `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	OpUserId      *string                                    `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s InsertSectionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequest) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequest) SetEnd(v *InsertSectionConfigRequestEnd) *InsertSectionConfigRequest {
	s.End = v
	return s
}

func (s *InsertSectionConfigRequest) SetScheduleName(v string) *InsertSectionConfigRequest {
	s.ScheduleName = &v
	return s
}

func (s *InsertSectionConfigRequest) SetSectionModels(v []*InsertSectionConfigRequestSectionModels) *InsertSectionConfigRequest {
	s.SectionModels = v
	return s
}

func (s *InsertSectionConfigRequest) SetStart(v *InsertSectionConfigRequestStart) *InsertSectionConfigRequest {
	s.Start = v
	return s
}

func (s *InsertSectionConfigRequest) SetOpUserId(v string) *InsertSectionConfigRequest {
	s.OpUserId = &v
	return s
}

type InsertSectionConfigRequestEnd struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s InsertSectionConfigRequestEnd) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestEnd) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestEnd) SetDayOfMonth(v int32) *InsertSectionConfigRequestEnd {
	s.DayOfMonth = &v
	return s
}

func (s *InsertSectionConfigRequestEnd) SetMonth(v int32) *InsertSectionConfigRequestEnd {
	s.Month = &v
	return s
}

func (s *InsertSectionConfigRequestEnd) SetYear(v int32) *InsertSectionConfigRequestEnd {
	s.Year = &v
	return s
}

type InsertSectionConfigRequestSectionModels struct {
	End          *InsertSectionConfigRequestSectionModelsEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionIndex *int32                                        `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName  *string                                       `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	SectionType  *string                                       `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	Start        *InsertSectionConfigRequestSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s InsertSectionConfigRequestSectionModels) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModels) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModels) SetEnd(v *InsertSectionConfigRequestSectionModelsEnd) *InsertSectionConfigRequestSectionModels {
	s.End = v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionIndex(v int32) *InsertSectionConfigRequestSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionName(v string) *InsertSectionConfigRequestSectionModels {
	s.SectionName = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetSectionType(v string) *InsertSectionConfigRequestSectionModels {
	s.SectionType = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModels) SetStart(v *InsertSectionConfigRequestSectionModelsStart) *InsertSectionConfigRequestSectionModels {
	s.Start = v
	return s
}

type InsertSectionConfigRequestSectionModelsEnd struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s InsertSectionConfigRequestSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModelsEnd) SetHour(v int32) *InsertSectionConfigRequestSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModelsEnd) SetMin(v int32) *InsertSectionConfigRequestSectionModelsEnd {
	s.Min = &v
	return s
}

type InsertSectionConfigRequestSectionModelsStart struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s InsertSectionConfigRequestSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestSectionModelsStart) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestSectionModelsStart) SetHour(v int32) *InsertSectionConfigRequestSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *InsertSectionConfigRequestSectionModelsStart) SetMin(v int32) *InsertSectionConfigRequestSectionModelsStart {
	s.Min = &v
	return s
}

type InsertSectionConfigRequestStart struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s InsertSectionConfigRequestStart) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigRequestStart) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigRequestStart) SetDayOfMonth(v int32) *InsertSectionConfigRequestStart {
	s.DayOfMonth = &v
	return s
}

func (s *InsertSectionConfigRequestStart) SetMonth(v int32) *InsertSectionConfigRequestStart {
	s.Month = &v
	return s
}

func (s *InsertSectionConfigRequestStart) SetYear(v int32) *InsertSectionConfigRequestStart {
	s.Year = &v
	return s
}

type InsertSectionConfigResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s InsertSectionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigResponseBody) SetResult(v bool) *InsertSectionConfigResponseBody {
	s.Result = &v
	return s
}

type InsertSectionConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertSectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertSectionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertSectionConfigResponse) GoString() string {
	return s.String()
}

func (s *InsertSectionConfigResponse) SetHeaders(v map[string]*string) *InsertSectionConfigResponse {
	s.Headers = v
	return s
}

func (s *InsertSectionConfigResponse) SetStatusCode(v int32) *InsertSectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertSectionConfigResponse) SetBody(v *InsertSectionConfigResponseBody) *InsertSectionConfigResponse {
	s.Body = v
	return s
}

type IsvDataWriteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IsvDataWriteHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteHeaders) GoString() string {
	return s.String()
}

func (s *IsvDataWriteHeaders) SetCommonHeaders(v map[string]*string) *IsvDataWriteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvDataWriteHeaders) SetXAcsDingtalkAccessToken(v string) *IsvDataWriteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IsvDataWriteRequest struct {
	ObjectCode   *string                              `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
	RowValueList [][]*IsvDataWriteRequestRowValueList `json:"rowValueList,omitempty" xml:"rowValueList,omitempty" type:"Repeated"`
}

func (s IsvDataWriteRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteRequest) GoString() string {
	return s.String()
}

func (s *IsvDataWriteRequest) SetObjectCode(v string) *IsvDataWriteRequest {
	s.ObjectCode = &v
	return s
}

func (s *IsvDataWriteRequest) SetRowValueList(v [][]*IsvDataWriteRequestRowValueList) *IsvDataWriteRequest {
	s.RowValueList = v
	return s
}

type IsvDataWriteRequestRowValueList struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s IsvDataWriteRequestRowValueList) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteRequestRowValueList) GoString() string {
	return s.String()
}

func (s *IsvDataWriteRequestRowValueList) SetName(v string) *IsvDataWriteRequestRowValueList {
	s.Name = &v
	return s
}

func (s *IsvDataWriteRequestRowValueList) SetValue(v string) *IsvDataWriteRequestRowValueList {
	s.Value = &v
	return s
}

type IsvDataWriteResponseBody struct {
	Result  *IsvDataWriteResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IsvDataWriteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteResponseBody) GoString() string {
	return s.String()
}

func (s *IsvDataWriteResponseBody) SetResult(v *IsvDataWriteResponseBodyResult) *IsvDataWriteResponseBody {
	s.Result = v
	return s
}

func (s *IsvDataWriteResponseBody) SetSuccess(v bool) *IsvDataWriteResponseBody {
	s.Success = &v
	return s
}

type IsvDataWriteResponseBodyResult struct {
	NeedRetry *bool `json:"needRetry,omitempty" xml:"needRetry,omitempty"`
	Success   *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IsvDataWriteResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IsvDataWriteResponseBodyResult) SetNeedRetry(v bool) *IsvDataWriteResponseBodyResult {
	s.NeedRetry = &v
	return s
}

func (s *IsvDataWriteResponseBodyResult) SetSuccess(v bool) *IsvDataWriteResponseBodyResult {
	s.Success = &v
	return s
}

type IsvDataWriteResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsvDataWriteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsvDataWriteResponse) String() string {
	return tea.Prettify(s)
}

func (s IsvDataWriteResponse) GoString() string {
	return s.String()
}

func (s *IsvDataWriteResponse) SetHeaders(v map[string]*string) *IsvDataWriteResponse {
	s.Headers = v
	return s
}

func (s *IsvDataWriteResponse) SetStatusCode(v int32) *IsvDataWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvDataWriteResponse) SetBody(v *IsvDataWriteResponseBody) *IsvDataWriteResponse {
	s.Body = v
	return s
}

type IsvMetadataQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IsvMetadataQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryHeaders) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryHeaders) SetCommonHeaders(v map[string]*string) *IsvMetadataQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsvMetadataQueryHeaders) SetXAcsDingtalkAccessToken(v string) *IsvMetadataQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IsvMetadataQueryRequest struct {
	ObjectCode *string `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
}

func (s IsvMetadataQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryRequest) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryRequest) SetObjectCode(v string) *IsvMetadataQueryRequest {
	s.ObjectCode = &v
	return s
}

type IsvMetadataQueryResponseBody struct {
	Result  *IsvMetadataQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s IsvMetadataQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryResponseBody) SetResult(v *IsvMetadataQueryResponseBodyResult) *IsvMetadataQueryResponseBody {
	s.Result = v
	return s
}

func (s *IsvMetadataQueryResponseBody) SetSuccess(v bool) *IsvMetadataQueryResponseBody {
	s.Success = &v
	return s
}

type IsvMetadataQueryResponseBodyResult struct {
	Fields     []*IsvMetadataQueryResponseBodyResultFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	TableCode  *string                                     `json:"tableCode,omitempty" xml:"tableCode,omitempty"`
	TableExist *bool                                       `json:"tableExist,omitempty" xml:"tableExist,omitempty"`
}

func (s IsvMetadataQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryResponseBodyResult) SetFields(v []*IsvMetadataQueryResponseBodyResultFields) *IsvMetadataQueryResponseBodyResult {
	s.Fields = v
	return s
}

func (s *IsvMetadataQueryResponseBodyResult) SetTableCode(v string) *IsvMetadataQueryResponseBodyResult {
	s.TableCode = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResult) SetTableExist(v bool) *IsvMetadataQueryResponseBodyResult {
	s.TableExist = &v
	return s
}

type IsvMetadataQueryResponseBodyResultFields struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	FieldKey    *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldName   *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	FieldType   *string `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	PrimaryKey  *bool   `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	Required    *bool   `json:"required,omitempty" xml:"required,omitempty"`
}

func (s IsvMetadataQueryResponseBodyResultFields) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryResponseBodyResultFields) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetDescription(v string) *IsvMetadataQueryResponseBodyResultFields {
	s.Description = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetFieldKey(v string) *IsvMetadataQueryResponseBodyResultFields {
	s.FieldKey = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetFieldName(v string) *IsvMetadataQueryResponseBodyResultFields {
	s.FieldName = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetFieldType(v string) *IsvMetadataQueryResponseBodyResultFields {
	s.FieldType = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetPrimaryKey(v bool) *IsvMetadataQueryResponseBodyResultFields {
	s.PrimaryKey = &v
	return s
}

func (s *IsvMetadataQueryResponseBodyResultFields) SetRequired(v bool) *IsvMetadataQueryResponseBodyResultFields {
	s.Required = &v
	return s
}

type IsvMetadataQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsvMetadataQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsvMetadataQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s IsvMetadataQueryResponse) GoString() string {
	return s.String()
}

func (s *IsvMetadataQueryResponse) SetHeaders(v map[string]*string) *IsvMetadataQueryResponse {
	s.Headers = v
	return s
}

func (s *IsvMetadataQueryResponse) SetStatusCode(v int32) *IsvMetadataQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvMetadataQueryResponse) SetBody(v *IsvMetadataQueryResponseBody) *IsvMetadataQueryResponse {
	s.Body = v
	return s
}

type ListOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrderHeaders) GoString() string {
	return s.String()
}

func (s *ListOrderHeaders) SetCommonHeaders(v map[string]*string) *ListOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ListOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListOrderRequest struct {
	CreateTimeEnd   *int64  `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	CreateTimeStart *int64  `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	MerchantId      *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo         *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PageNumber      *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Scene           *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	Status          *int64  `json:"status,omitempty" xml:"status,omitempty"`
	TradeNo         *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrderRequest) GoString() string {
	return s.String()
}

func (s *ListOrderRequest) SetCreateTimeEnd(v int64) *ListOrderRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListOrderRequest) SetCreateTimeStart(v int64) *ListOrderRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListOrderRequest) SetMerchantId(v string) *ListOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *ListOrderRequest) SetOrderNo(v string) *ListOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *ListOrderRequest) SetPageNumber(v int64) *ListOrderRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrderRequest) SetPageSize(v int64) *ListOrderRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrderRequest) SetScene(v int64) *ListOrderRequest {
	s.Scene = &v
	return s
}

func (s *ListOrderRequest) SetStatus(v int64) *ListOrderRequest {
	s.Status = &v
	return s
}

func (s *ListOrderRequest) SetTradeNo(v string) *ListOrderRequest {
	s.TradeNo = &v
	return s
}

func (s *ListOrderRequest) SetUserId(v string) *ListOrderRequest {
	s.UserId = &v
	return s
}

type ListOrderResponseBody struct {
	List  []*ListOrderResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Total *int64                       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBody) SetList(v []*ListOrderResponseBodyList) *ListOrderResponseBody {
	s.List = v
	return s
}

func (s *ListOrderResponseBody) SetTotal(v int64) *ListOrderResponseBody {
	s.Total = &v
	return s
}

type ListOrderResponseBodyList struct {
	ActualAmount *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	BuyerId      *string `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreateTime   *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OrderNo      *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	PayTime      *int64  `json:"payTime,omitempty" xml:"payTime,omitempty"`
	RefundNo     *string `json:"refundNo,omitempty" xml:"refundNo,omitempty"`
	Scene        *int64  `json:"scene,omitempty" xml:"scene,omitempty"`
	StartTime    *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status       *int64  `json:"status,omitempty" xml:"status,omitempty"`
	TradeNo      *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListOrderResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBodyList) SetActualAmount(v int64) *ListOrderResponseBodyList {
	s.ActualAmount = &v
	return s
}

func (s *ListOrderResponseBodyList) SetBuyerId(v string) *ListOrderResponseBodyList {
	s.BuyerId = &v
	return s
}

func (s *ListOrderResponseBodyList) SetCorpId(v string) *ListOrderResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *ListOrderResponseBodyList) SetCreateTime(v int64) *ListOrderResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListOrderResponseBodyList) SetEndTime(v int64) *ListOrderResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *ListOrderResponseBodyList) SetOrderNo(v string) *ListOrderResponseBodyList {
	s.OrderNo = &v
	return s
}

func (s *ListOrderResponseBodyList) SetPayTime(v int64) *ListOrderResponseBodyList {
	s.PayTime = &v
	return s
}

func (s *ListOrderResponseBodyList) SetRefundNo(v string) *ListOrderResponseBodyList {
	s.RefundNo = &v
	return s
}

func (s *ListOrderResponseBodyList) SetScene(v int64) *ListOrderResponseBodyList {
	s.Scene = &v
	return s
}

func (s *ListOrderResponseBodyList) SetStartTime(v int64) *ListOrderResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListOrderResponseBodyList) SetStatus(v int64) *ListOrderResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListOrderResponseBodyList) SetTradeNo(v string) *ListOrderResponseBodyList {
	s.TradeNo = &v
	return s
}

func (s *ListOrderResponseBodyList) SetUserId(v string) *ListOrderResponseBodyList {
	s.UserId = &v
	return s
}

type ListOrderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponse) GoString() string {
	return s.String()
}

func (s *ListOrderResponse) SetHeaders(v map[string]*string) *ListOrderResponse {
	s.Headers = v
	return s
}

func (s *ListOrderResponse) SetStatusCode(v int32) *ListOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrderResponse) SetBody(v *ListOrderResponseBody) *ListOrderResponse {
	s.Body = v
	return s
}

type MoveStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentHeaders) GoString() string {
	return s.String()
}

func (s *MoveStudentHeaders) SetCommonHeaders(v map[string]*string) *MoveStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveStudentHeaders) SetXAcsDingtalkAccessToken(v string) *MoveStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveStudentRequest struct {
	Operator      *string `json:"operator,omitempty" xml:"operator,omitempty"`
	OriginClassId *int64  `json:"originClassId,omitempty" xml:"originClassId,omitempty"`
	TargetClassId *int64  `json:"targetClassId,omitempty" xml:"targetClassId,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s MoveStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentRequest) GoString() string {
	return s.String()
}

func (s *MoveStudentRequest) SetOperator(v string) *MoveStudentRequest {
	s.Operator = &v
	return s
}

func (s *MoveStudentRequest) SetOriginClassId(v int64) *MoveStudentRequest {
	s.OriginClassId = &v
	return s
}

func (s *MoveStudentRequest) SetTargetClassId(v int64) *MoveStudentRequest {
	s.TargetClassId = &v
	return s
}

func (s *MoveStudentRequest) SetUserId(v string) *MoveStudentRequest {
	s.UserId = &v
	return s
}

type MoveStudentResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MoveStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentResponseBody) GoString() string {
	return s.String()
}

func (s *MoveStudentResponseBody) SetSuccess(v bool) *MoveStudentResponseBody {
	s.Success = &v
	return s
}

type MoveStudentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveStudentResponse) GoString() string {
	return s.String()
}

func (s *MoveStudentResponse) SetHeaders(v map[string]*string) *MoveStudentResponse {
	s.Headers = v
	return s
}

func (s *MoveStudentResponse) SetStatusCode(v int32) *MoveStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveStudentResponse) SetBody(v *MoveStudentResponseBody) *MoveStudentResponse {
	s.Body = v
	return s
}

type PageQueryDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageQueryDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageQueryDevicesHeaders) GoString() string {
	return s.String()
}

func (s *PageQueryDevicesHeaders) SetCommonHeaders(v map[string]*string) *PageQueryDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageQueryDevicesHeaders) SetXAcsDingtalkAccessToken(v string) *PageQueryDevicesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageQueryDevicesRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PageQueryDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s PageQueryDevicesRequest) GoString() string {
	return s.String()
}

func (s *PageQueryDevicesRequest) SetMaxResults(v int64) *PageQueryDevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *PageQueryDevicesRequest) SetNextToken(v string) *PageQueryDevicesRequest {
	s.NextToken = &v
	return s
}

func (s *PageQueryDevicesRequest) SetType(v string) *PageQueryDevicesRequest {
	s.Type = &v
	return s
}

type PageQueryDevicesResponseBody struct {
	List       []*PageQueryDevicesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken  *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	TotalCount *int64                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PageQueryDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageQueryDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryDevicesResponseBody) SetList(v []*PageQueryDevicesResponseBodyList) *PageQueryDevicesResponseBody {
	s.List = v
	return s
}

func (s *PageQueryDevicesResponseBody) SetNextToken(v string) *PageQueryDevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *PageQueryDevicesResponseBody) SetTotalCount(v int64) *PageQueryDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type PageQueryDevicesResponseBodyList struct {
	GmtExpiry *int64  `json:"gmtExpiry,omitempty" xml:"gmtExpiry,omitempty"`
	Model     *string `json:"model,omitempty" xml:"model,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PageQueryDevicesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PageQueryDevicesResponseBodyList) GoString() string {
	return s.String()
}

func (s *PageQueryDevicesResponseBodyList) SetGmtExpiry(v int64) *PageQueryDevicesResponseBodyList {
	s.GmtExpiry = &v
	return s
}

func (s *PageQueryDevicesResponseBodyList) SetModel(v string) *PageQueryDevicesResponseBodyList {
	s.Model = &v
	return s
}

func (s *PageQueryDevicesResponseBodyList) SetName(v string) *PageQueryDevicesResponseBodyList {
	s.Name = &v
	return s
}

func (s *PageQueryDevicesResponseBodyList) SetSn(v string) *PageQueryDevicesResponseBodyList {
	s.Sn = &v
	return s
}

func (s *PageQueryDevicesResponseBodyList) SetType(v string) *PageQueryDevicesResponseBodyList {
	s.Type = &v
	return s
}

type PageQueryDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageQueryDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageQueryDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s PageQueryDevicesResponse) GoString() string {
	return s.String()
}

func (s *PageQueryDevicesResponse) SetHeaders(v map[string]*string) *PageQueryDevicesResponse {
	s.Headers = v
	return s
}

func (s *PageQueryDevicesResponse) SetStatusCode(v int32) *PageQueryDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryDevicesResponse) SetBody(v *PageQueryDevicesResponseBody) *PageQueryDevicesResponse {
	s.Body = v
	return s
}

type PayOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PayOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s PayOrderHeaders) GoString() string {
	return s.String()
}

func (s *PayOrderHeaders) SetCommonHeaders(v map[string]*string) *PayOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PayOrderHeaders) SetXAcsDingtalkAccessToken(v string) *PayOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PayOrderRequest struct {
	FaceId    *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	OrderNo   *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s PayOrderRequest) GoString() string {
	return s.String()
}

func (s *PayOrderRequest) SetFaceId(v string) *PayOrderRequest {
	s.FaceId = &v
	return s
}

func (s *PayOrderRequest) SetOrderNo(v string) *PayOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *PayOrderRequest) SetSignature(v string) *PayOrderRequest {
	s.Signature = &v
	return s
}

func (s *PayOrderRequest) SetSn(v string) *PayOrderRequest {
	s.Sn = &v
	return s
}

func (s *PayOrderRequest) SetTimestamp(v int64) *PayOrderRequest {
	s.Timestamp = &v
	return s
}

func (s *PayOrderRequest) SetUserId(v string) *PayOrderRequest {
	s.UserId = &v
	return s
}

func (s *PayOrderRequest) SetVersion(v string) *PayOrderRequest {
	s.Version = &v
	return s
}

type PayOrderResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *PayOrderResponseBody) SetSuccess(v bool) *PayOrderResponseBody {
	s.Success = &v
	return s
}

type PayOrderResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s PayOrderResponse) GoString() string {
	return s.String()
}

func (s *PayOrderResponse) SetHeaders(v map[string]*string) *PayOrderResponse {
	s.Headers = v
	return s
}

func (s *PayOrderResponse) SetStatusCode(v int32) *PayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *PayOrderResponse) SetBody(v *PayOrderResponseBody) *PayOrderResponse {
	s.Body = v
	return s
}

type PollingConfirmStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PollingConfirmStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusHeaders) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusHeaders) SetCommonHeaders(v map[string]*string) *PollingConfirmStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PollingConfirmStatusHeaders) SetXAcsDingtalkAccessToken(v string) *PollingConfirmStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PollingConfirmStatusRequest struct {
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Ext        *string `json:"ext,omitempty" xml:"ext,omitempty"`
	IsvCode    *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	OpUserId   *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s PollingConfirmStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusRequest) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusRequest) SetCourseCode(v string) *PollingConfirmStatusRequest {
	s.CourseCode = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetExt(v string) *PollingConfirmStatusRequest {
	s.Ext = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetIsvCode(v string) *PollingConfirmStatusRequest {
	s.IsvCode = &v
	return s
}

func (s *PollingConfirmStatusRequest) SetOpUserId(v string) *PollingConfirmStatusRequest {
	s.OpUserId = &v
	return s
}

type PollingConfirmStatusResponseBody struct {
	UniversityPollingCourseStatusResponse *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse `json:"universityPollingCourseStatusResponse,omitempty" xml:"universityPollingCourseStatusResponse,omitempty" type:"Struct"`
}

func (s PollingConfirmStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBody) SetUniversityPollingCourseStatusResponse(v *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) *PollingConfirmStatusResponseBody {
	s.UniversityPollingCourseStatusResponse = v
	return s
}

type PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse struct {
	ConfirmStatus    *bool                                                                                    `json:"confirmStatus,omitempty" xml:"confirmStatus,omitempty"`
	CourseCode       *string                                                                                  `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	LivePlayInfoList []*PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetConfirmStatus(v bool) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.ConfirmStatus = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetCourseCode(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.CourseCode = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse) SetLivePlayInfoList(v []*PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponse {
	s.LivePlayInfoList = v
	return s
}

type PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList struct {
	LiveInputUrl  *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	LiveOutputUrl *string `json:"liveOutputUrl,omitempty" xml:"liveOutputUrl,omitempty"`
	LiveType      *int64  `json:"liveType,omitempty" xml:"liveType,omitempty"`
	ReplayUrl     *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveInputUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveOutputUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveOutputUrl = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetLiveType(v int64) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList) SetReplayUrl(v string) *PollingConfirmStatusResponseBodyUniversityPollingCourseStatusResponseLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type PollingConfirmStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PollingConfirmStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PollingConfirmStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PollingConfirmStatusResponse) GoString() string {
	return s.String()
}

func (s *PollingConfirmStatusResponse) SetHeaders(v map[string]*string) *PollingConfirmStatusResponse {
	s.Headers = v
	return s
}

func (s *PollingConfirmStatusResponse) SetStatusCode(v int32) *PollingConfirmStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PollingConfirmStatusResponse) SetBody(v *PollingConfirmStatusResponseBody) *PollingConfirmStatusResponse {
	s.Body = v
	return s
}

type PreDialHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PreDialHeaders) String() string {
	return tea.Prettify(s)
}

func (s PreDialHeaders) GoString() string {
	return s.String()
}

func (s *PreDialHeaders) SetCommonHeaders(v map[string]*string) *PreDialHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreDialHeaders) SetXAcsDingtalkAccessToken(v string) *PreDialHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PreDialRequest struct {
	CallerUserId   *string `json:"callerUserId,omitempty" xml:"callerUserId,omitempty"`
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
	Sn             *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PreDialRequest) String() string {
	return tea.Prettify(s)
}

func (s PreDialRequest) GoString() string {
	return s.String()
}

func (s *PreDialRequest) SetCallerUserId(v string) *PreDialRequest {
	s.CallerUserId = &v
	return s
}

func (s *PreDialRequest) SetReceiverUserId(v string) *PreDialRequest {
	s.ReceiverUserId = &v
	return s
}

func (s *PreDialRequest) SetSn(v string) *PreDialRequest {
	s.Sn = &v
	return s
}

func (s *PreDialRequest) SetType(v string) *PreDialRequest {
	s.Type = &v
	return s
}

type PreDialResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PreDialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreDialResponseBody) GoString() string {
	return s.String()
}

func (s *PreDialResponseBody) SetResult(v bool) *PreDialResponseBody {
	s.Result = &v
	return s
}

type PreDialResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreDialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreDialResponse) String() string {
	return tea.Prettify(s)
}

func (s PreDialResponse) GoString() string {
	return s.String()
}

func (s *PreDialResponse) SetHeaders(v map[string]*string) *PreDialResponse {
	s.Headers = v
	return s
}

func (s *PreDialResponse) SetStatusCode(v int32) *PreDialResponse {
	s.StatusCode = &v
	return s
}

func (s *PreDialResponse) SetBody(v *PreDialResponseBody) *PreDialResponse {
	s.Body = v
	return s
}

type ProvidePointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ProvidePointHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProvidePointHeaders) GoString() string {
	return s.String()
}

func (s *ProvidePointHeaders) SetCommonHeaders(v map[string]*string) *ProvidePointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProvidePointHeaders) SetXAcsDingtalkAccessToken(v string) *ProvidePointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ProvidePointRequest struct {
	ActionCode *string `json:"actionCode,omitempty" xml:"actionCode,omitempty"`
	BizId      *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	PointType  *string `json:"pointType,omitempty" xml:"pointType,omitempty"`
}

func (s ProvidePointRequest) String() string {
	return tea.Prettify(s)
}

func (s ProvidePointRequest) GoString() string {
	return s.String()
}

func (s *ProvidePointRequest) SetActionCode(v string) *ProvidePointRequest {
	s.ActionCode = &v
	return s
}

func (s *ProvidePointRequest) SetBizId(v string) *ProvidePointRequest {
	s.BizId = &v
	return s
}

func (s *ProvidePointRequest) SetPointType(v string) *ProvidePointRequest {
	s.PointType = &v
	return s
}

type ProvidePointResponseBody struct {
	Result  *ProvidePointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ProvidePointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProvidePointResponseBody) GoString() string {
	return s.String()
}

func (s *ProvidePointResponseBody) SetResult(v *ProvidePointResponseBodyResult) *ProvidePointResponseBody {
	s.Result = v
	return s
}

func (s *ProvidePointResponseBody) SetSuccess(v bool) *ProvidePointResponseBody {
	s.Success = &v
	return s
}

type ProvidePointResponseBodyResult struct {
	AvailableQuota *int64 `json:"availableQuota,omitempty" xml:"availableQuota,omitempty"`
	ProvideNum     *int64 `json:"provideNum,omitempty" xml:"provideNum,omitempty"`
	ProvideSuccess *bool  `json:"provideSuccess,omitempty" xml:"provideSuccess,omitempty"`
}

func (s ProvidePointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ProvidePointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ProvidePointResponseBodyResult) SetAvailableQuota(v int64) *ProvidePointResponseBodyResult {
	s.AvailableQuota = &v
	return s
}

func (s *ProvidePointResponseBodyResult) SetProvideNum(v int64) *ProvidePointResponseBodyResult {
	s.ProvideNum = &v
	return s
}

func (s *ProvidePointResponseBodyResult) SetProvideSuccess(v bool) *ProvidePointResponseBodyResult {
	s.ProvideSuccess = &v
	return s
}

type ProvidePointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ProvidePointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProvidePointResponse) String() string {
	return tea.Prettify(s)
}

func (s ProvidePointResponse) GoString() string {
	return s.String()
}

func (s *ProvidePointResponse) SetHeaders(v map[string]*string) *ProvidePointResponse {
	s.Headers = v
	return s
}

func (s *ProvidePointResponse) SetStatusCode(v int32) *ProvidePointResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvidePointResponse) SetBody(v *ProvidePointResponseBody) *ProvidePointResponse {
	s.Body = v
	return s
}

type QueryAllSubjectsFromClassScheduleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleHeaders) SetCommonHeaders(v map[string]*string) *QueryAllSubjectsFromClassScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllSubjectsFromClassScheduleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllSubjectsFromClassScheduleRequest struct {
	ClassIds   []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	OpUserId   *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	PeriodCode *string  `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetClassIds(v []*int64) *QueryAllSubjectsFromClassScheduleRequest {
	s.ClassIds = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetOpUserId(v string) *QueryAllSubjectsFromClassScheduleRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleRequest) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleRequest {
	s.PeriodCode = &v
	return s
}

type QueryAllSubjectsFromClassScheduleShrinkRequest struct {
	ClassIdsShrink *string `json:"classIds,omitempty" xml:"classIds,omitempty"`
	OpUserId       *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	PeriodCode     *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetClassIdsShrink(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.ClassIdsShrink = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetOpUserId(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleShrinkRequest) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleShrinkRequest {
	s.PeriodCode = &v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBody struct {
	Result []*QueryAllSubjectsFromClassScheduleResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBody) SetResult(v []*QueryAllSubjectsFromClassScheduleResponseBodyResult) *QueryAllSubjectsFromClassScheduleResponseBody {
	s.Result = v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResult struct {
	CreatorOrgId *int64                                                  `json:"creatorOrgId,omitempty" xml:"creatorOrgId,omitempty"`
	Ext          *QueryAllSubjectsFromClassScheduleResponseBodyResultExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	PeriodCode   *string                                                 `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	SubjectCode  *string                                                 `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	SubjectName  *string                                                 `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetCreatorOrgId(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.CreatorOrgId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetExt(v *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.Ext = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetPeriodCode(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetSubjectCode(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResult) SetSubjectName(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResult {
	s.SubjectName = &v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResultExt struct {
	BackgroundColor *string                                                              `json:"backgroundColor,omitempty" xml:"backgroundColor,omitempty"`
	ClassId         *int64                                                               `json:"classId,omitempty" xml:"classId,omitempty"`
	FontColor       *string                                                              `json:"fontColor,omitempty" xml:"fontColor,omitempty"`
	TeacherList     []*QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList `json:"teacherList,omitempty" xml:"teacherList,omitempty" type:"Repeated"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExt) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExt) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetBackgroundColor(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.BackgroundColor = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetClassId(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.ClassId = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetFontColor(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.FontColor = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExt) SetTeacherList(v []*QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) *QueryAllSubjectsFromClassScheduleResponseBodyResultExt {
	s.TeacherList = v
	return s
}

type QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList struct {
	Avator *string `json:"avator,omitempty" xml:"avator,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Uid    *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetAvator(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Avator = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetName(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Name = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetUid(v int64) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.Uid = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList) SetUserId(v string) *QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList {
	s.UserId = &v
	return s
}

type QueryAllSubjectsFromClassScheduleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllSubjectsFromClassScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllSubjectsFromClassScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllSubjectsFromClassScheduleResponse) GoString() string {
	return s.String()
}

func (s *QueryAllSubjectsFromClassScheduleResponse) SetHeaders(v map[string]*string) *QueryAllSubjectsFromClassScheduleResponse {
	s.Headers = v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponse) SetStatusCode(v int32) *QueryAllSubjectsFromClassScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllSubjectsFromClassScheduleResponse) SetBody(v *QueryAllSubjectsFromClassScheduleResponseBody) *QueryAllSubjectsFromClassScheduleResponse {
	s.Body = v
	return s
}

type QueryClassScheduleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleRequest struct {
	SectionIndexList []*int64  `json:"sectionIndexList,omitempty" xml:"sectionIndexList,omitempty" type:"Repeated"`
	SubscriberIds    []*string `json:"subscriberIds,omitempty" xml:"subscriberIds,omitempty" type:"Repeated"`
	EndTime          *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpUserId         *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	StartTime        *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscriberType   *string   `json:"subscriberType,omitempty" xml:"subscriberType,omitempty"`
}

func (s QueryClassScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleRequest) SetSectionIndexList(v []*int64) *QueryClassScheduleRequest {
	s.SectionIndexList = v
	return s
}

func (s *QueryClassScheduleRequest) SetSubscriberIds(v []*string) *QueryClassScheduleRequest {
	s.SubscriberIds = v
	return s
}

func (s *QueryClassScheduleRequest) SetEndTime(v int64) *QueryClassScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleRequest) SetOpUserId(v string) *QueryClassScheduleRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryClassScheduleRequest) SetStartTime(v int64) *QueryClassScheduleRequest {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleRequest) SetSubscriberType(v string) *QueryClassScheduleRequest {
	s.SubscriberType = &v
	return s
}

type QueryClassScheduleResponseBody struct {
	Config     *QueryClassScheduleResponseBodyConfig       `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	CourseDTOS []*QueryClassScheduleResponseBodyCourseDTOS `json:"courseDTOS,omitempty" xml:"courseDTOS,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBody) SetConfig(v *QueryClassScheduleResponseBodyConfig) *QueryClassScheduleResponseBody {
	s.Config = v
	return s
}

func (s *QueryClassScheduleResponseBody) SetCourseDTOS(v []*QueryClassScheduleResponseBodyCourseDTOS) *QueryClassScheduleResponseBody {
	s.CourseDTOS = v
	return s
}

type QueryClassScheduleResponseBodyConfig struct {
	End           *QueryClassScheduleResponseBodyConfigEnd             `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionModels []*QueryClassScheduleResponseBodyConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	Start         *QueryClassScheduleResponseBodyConfigStart           `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s QueryClassScheduleResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfig) SetEnd(v *QueryClassScheduleResponseBodyConfigEnd) *QueryClassScheduleResponseBodyConfig {
	s.End = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfig) SetSectionModels(v []*QueryClassScheduleResponseBodyConfigSectionModels) *QueryClassScheduleResponseBodyConfig {
	s.SectionModels = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfig) SetStart(v *QueryClassScheduleResponseBodyConfigStart) *QueryClassScheduleResponseBodyConfig {
	s.Start = v
	return s
}

type QueryClassScheduleResponseBodyConfigEnd struct {
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int64 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int64 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetDayOfMonth(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.DayOfMonth = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetMonth(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigEnd) SetYear(v int64) *QueryClassScheduleResponseBodyConfigEnd {
	s.Year = &v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModels struct {
	End          *QueryClassScheduleResponseBodyConfigSectionModelsEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionIndex *int64                                                  `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName  *string                                                 `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	SectionType  *string                                                 `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	Start        *QueryClassScheduleResponseBodyConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetEnd(v *QueryClassScheduleResponseBodyConfigSectionModelsEnd) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.End = v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionIndex(v int64) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionName(v string) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetSectionType(v string) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModels) SetStart(v *QueryClassScheduleResponseBodyConfigSectionModelsStart) *QueryClassScheduleResponseBodyConfigSectionModels {
	s.Start = v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModelsEnd struct {
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int64 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsEnd) SetHour(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsEnd) SetMin(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsEnd {
	s.Min = &v
	return s
}

type QueryClassScheduleResponseBodyConfigSectionModelsStart struct {
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int64 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsStart) SetHour(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigSectionModelsStart) SetMin(v int64) *QueryClassScheduleResponseBodyConfigSectionModelsStart {
	s.Min = &v
	return s
}

type QueryClassScheduleResponseBodyConfigStart struct {
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int64 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int64 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryClassScheduleResponseBodyConfigStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyConfigStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetDayOfMonth(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.DayOfMonth = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetMonth(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleResponseBodyConfigStart) SetYear(v int64) *QueryClassScheduleResponseBodyConfigStart {
	s.Year = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOS struct {
	ClassId         *int64                                                   `json:"classId,omitempty" xml:"classId,omitempty"`
	Classrooms      []*QueryClassScheduleResponseBodyCourseDTOSClassrooms    `json:"classrooms,omitempty" xml:"classrooms,omitempty" type:"Repeated"`
	Code            *string                                                  `json:"code,omitempty" xml:"code,omitempty"`
	CourseGroupCode *string                                                  `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	CoverUrl        *string                                                  `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CreatorCorpId   *string                                                  `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	CreatorUserId   *string                                                  `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	CreatorUserName *string                                                  `json:"creatorUserName,omitempty" xml:"creatorUserName,omitempty"`
	EduUserModels   []*QueryClassScheduleResponseBodyCourseDTOSEduUserModels `json:"eduUserModels,omitempty" xml:"eduUserModels,omitempty" type:"Repeated"`
	EndTime         *int64                                                   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExtInfo         *string                                                  `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Introduce       *string                                                  `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name            *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
	SectionIndex    *int64                                                   `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName     *string                                                  `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	StartTime       *int64                                                   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status          *int64                                                   `json:"status,omitempty" xml:"status,omitempty"`
	SubjectCode     *string                                                  `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	TeacherCorpId   *string                                                  `json:"teacherCorpId,omitempty" xml:"teacherCorpId,omitempty"`
	TeacherUserId   *string                                                  `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	TeacherUserName *string                                                  `json:"teacherUserName,omitempty" xml:"teacherUserName,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOS) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOS) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetClassId(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetClassrooms(v []*QueryClassScheduleResponseBodyCourseDTOSClassrooms) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Classrooms = v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Code = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCourseGroupCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCoverUrl(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CoverUrl = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorCorpId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorCorpId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorUserId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorUserId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetCreatorUserName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.CreatorUserName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetEduUserModels(v []*QueryClassScheduleResponseBodyCourseDTOSEduUserModels) *QueryClassScheduleResponseBodyCourseDTOS {
	s.EduUserModels = v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetEndTime(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetExtInfo(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.ExtInfo = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetIntroduce(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Introduce = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSectionIndex(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSectionName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetStartTime(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetStatus(v int64) *QueryClassScheduleResponseBodyCourseDTOS {
	s.Status = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetSubjectCode(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.SubjectCode = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherCorpId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherCorpId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherUserId(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherUserId = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOS) SetTeacherUserName(v string) *QueryClassScheduleResponseBodyCourseDTOS {
	s.TeacherUserName = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOSClassrooms struct {
	InteractInfo *string `json:"interactInfo,omitempty" xml:"interactInfo,omitempty"`
	TargetId     *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOSClassrooms) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOSClassrooms) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOSClassrooms) SetInteractInfo(v string) *QueryClassScheduleResponseBodyCourseDTOSClassrooms {
	s.InteractInfo = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSClassrooms) SetTargetId(v string) *QueryClassScheduleResponseBodyCourseDTOSClassrooms {
	s.TargetId = &v
	return s
}

type QueryClassScheduleResponseBodyCourseDTOSEduUserModels struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Uid    *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryClassScheduleResponseBodyCourseDTOSEduUserModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponseBodyCourseDTOSEduUserModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetName(v string) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetUid(v int64) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.Uid = &v
	return s
}

func (s *QueryClassScheduleResponseBodyCourseDTOSEduUserModels) SetUserId(v string) *QueryClassScheduleResponseBodyCourseDTOSEduUserModels {
	s.UserId = &v
	return s
}

type QueryClassScheduleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryClassScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleResponse) SetHeaders(v map[string]*string) *QueryClassScheduleResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleResponse) SetStatusCode(v int32) *QueryClassScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClassScheduleResponse) SetBody(v *QueryClassScheduleResponseBody) *QueryClassScheduleResponse {
	s.Body = v
	return s
}

type QueryClassScheduleByTimeSchoolHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleByTimeSchoolHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleByTimeSchoolHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleByTimeSchoolRequest struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpUserId  *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetEndTime(v int64) *QueryClassScheduleByTimeSchoolRequest {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetOpUserId(v string) *QueryClassScheduleByTimeSchoolRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolRequest) SetStartTime(v int64) *QueryClassScheduleByTimeSchoolRequest {
	s.StartTime = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBody struct {
	Result []*QueryClassScheduleByTimeSchoolResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleByTimeSchoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBody) SetResult(v []*QueryClassScheduleByTimeSchoolResponseBodyResult) *QueryClassScheduleByTimeSchoolResponseBody {
	s.Result = v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResult struct {
	BizKey          *string                                                          `json:"bizKey,omitempty" xml:"bizKey,omitempty"`
	ClassId         *int64                                                           `json:"classId,omitempty" xml:"classId,omitempty"`
	Classrooms      []*QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms    `json:"classrooms,omitempty" xml:"classrooms,omitempty" type:"Repeated"`
	Code            *string                                                          `json:"code,omitempty" xml:"code,omitempty"`
	CourseGroupCode *string                                                          `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	CoverUrl        *string                                                          `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CreatorCorpId   *string                                                          `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	CreatorUserId   *string                                                          `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	CreatorUserName *string                                                          `json:"creatorUserName,omitempty" xml:"creatorUserName,omitempty"`
	EduUserModels   []*QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels `json:"eduUserModels,omitempty" xml:"eduUserModels,omitempty" type:"Repeated"`
	EndTime         *int64                                                           `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ExtInfo         *string                                                          `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Introduce       *string                                                          `json:"introduce,omitempty" xml:"introduce,omitempty"`
	Name            *string                                                          `json:"name,omitempty" xml:"name,omitempty"`
	SectionIndex    *int64                                                           `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName     *string                                                          `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	StartTime       *int64                                                           `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status          *int64                                                           `json:"status,omitempty" xml:"status,omitempty"`
	SubjectCode     *string                                                          `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	TeacherCorpId   *string                                                          `json:"teacherCorpId,omitempty" xml:"teacherCorpId,omitempty"`
	TeacherUserId   *string                                                          `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
	TeacherUserName *string                                                          `json:"teacherUserName,omitempty" xml:"teacherUserName,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetBizKey(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.BizKey = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetClassId(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetClassrooms(v []*QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Classrooms = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Code = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCourseGroupCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCoverUrl(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorCorpId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorCorpId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetCreatorUserName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.CreatorUserName = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetEduUserModels(v []*QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.EduUserModels = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetEndTime(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetExtInfo(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.ExtInfo = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetIntroduce(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Introduce = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSectionIndex(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSectionName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetStartTime(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetStatus(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetSubjectCode(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherCorpId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherCorpId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResult) SetTeacherUserName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResult {
	s.TeacherUserName = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms struct {
	InteractInfo *string `json:"interactInfo,omitempty" xml:"interactInfo,omitempty"`
	TargetId     *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) SetInteractInfo(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms {
	s.InteractInfo = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms) SetTargetId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultClassrooms {
	s.TargetId = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Uid    *int64  `json:"uid,omitempty" xml:"uid,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetName(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.Name = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetUid(v int64) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.Uid = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels) SetUserId(v string) *QueryClassScheduleByTimeSchoolResponseBodyResultEduUserModels {
	s.UserId = &v
	return s
}

type QueryClassScheduleByTimeSchoolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryClassScheduleByTimeSchoolResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleByTimeSchoolResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleByTimeSchoolResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleByTimeSchoolResponse) SetHeaders(v map[string]*string) *QueryClassScheduleByTimeSchoolResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponse) SetStatusCode(v int32) *QueryClassScheduleByTimeSchoolResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClassScheduleByTimeSchoolResponse) SetBody(v *QueryClassScheduleByTimeSchoolResponseBody) *QueryClassScheduleByTimeSchoolResponse {
	s.Body = v
	return s
}

type QueryClassScheduleConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryClassScheduleConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigHeaders) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigHeaders) SetCommonHeaders(v map[string]*string) *QueryClassScheduleConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryClassScheduleConfigHeaders) SetXAcsDingtalkAccessToken(v string) *QueryClassScheduleConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryClassScheduleConfigRequest struct {
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	OpUserId *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryClassScheduleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigRequest) SetClassIds(v []*int64) *QueryClassScheduleConfigRequest {
	s.ClassIds = v
	return s
}

func (s *QueryClassScheduleConfigRequest) SetOpUserId(v string) *QueryClassScheduleConfigRequest {
	s.OpUserId = &v
	return s
}

type QueryClassScheduleConfigShrinkRequest struct {
	ClassIdsShrink *string `json:"classIds,omitempty" xml:"classIds,omitempty"`
	OpUserId       *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryClassScheduleConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigShrinkRequest) SetClassIdsShrink(v string) *QueryClassScheduleConfigShrinkRequest {
	s.ClassIdsShrink = &v
	return s
}

func (s *QueryClassScheduleConfigShrinkRequest) SetOpUserId(v string) *QueryClassScheduleConfigShrinkRequest {
	s.OpUserId = &v
	return s
}

type QueryClassScheduleConfigResponseBody struct {
	Result []*QueryClassScheduleConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryClassScheduleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBody) SetResult(v []*QueryClassScheduleConfigResponseBodyResult) *QueryClassScheduleConfigResponseBody {
	s.Result = v
	return s
}

type QueryClassScheduleConfigResponseBodyResult struct {
	ClassId       *int64                                                     `json:"classId,omitempty" xml:"classId,omitempty"`
	End           *QueryClassScheduleConfigResponseBodyResultEnd             `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionModels []*QueryClassScheduleConfigResponseBodyResultSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
	Start         *QueryClassScheduleConfigResponseBodyResultStart           `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s QueryClassScheduleConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetClassId(v int64) *QueryClassScheduleConfigResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetEnd(v *QueryClassScheduleConfigResponseBodyResultEnd) *QueryClassScheduleConfigResponseBodyResult {
	s.End = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetSectionModels(v []*QueryClassScheduleConfigResponseBodyResultSectionModels) *QueryClassScheduleConfigResponseBodyResult {
	s.SectionModels = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResult) SetStart(v *QueryClassScheduleConfigResponseBodyResultStart) *QueryClassScheduleConfigResponseBodyResult {
	s.Start = v
	return s
}

type QueryClassScheduleConfigResponseBodyResultEnd struct {
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int64 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int64 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetDayOfMonth(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.DayOfMonth = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetMonth(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultEnd) SetYear(v int64) *QueryClassScheduleConfigResponseBodyResultEnd {
	s.Year = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModels struct {
	End          *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionIndex *int32                                                        `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName  *string                                                       `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	SectionType  *string                                                       `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	Start        *QueryClassScheduleConfigResponseBodyResultSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModels) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetEnd(v *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.End = v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionIndex(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionName(v string) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionName = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetSectionType(v string) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.SectionType = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModels) SetStart(v *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) *QueryClassScheduleConfigResponseBodyResultSectionModels {
	s.Start = v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModelsEnd struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) SetHour(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd) SetMin(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsEnd {
	s.Min = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultSectionModelsStart struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultSectionModelsStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) SetHour(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultSectionModelsStart) SetMin(v int32) *QueryClassScheduleConfigResponseBodyResultSectionModelsStart {
	s.Min = &v
	return s
}

type QueryClassScheduleConfigResponseBodyResultStart struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryClassScheduleConfigResponseBodyResultStart) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponseBodyResultStart) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetDayOfMonth(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.DayOfMonth = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetMonth(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.Month = &v
	return s
}

func (s *QueryClassScheduleConfigResponseBodyResultStart) SetYear(v int32) *QueryClassScheduleConfigResponseBodyResultStart {
	s.Year = &v
	return s
}

type QueryClassScheduleConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryClassScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClassScheduleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClassScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryClassScheduleConfigResponse) SetHeaders(v map[string]*string) *QueryClassScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryClassScheduleConfigResponse) SetStatusCode(v int32) *QueryClassScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClassScheduleConfigResponse) SetBody(v *QueryClassScheduleConfigResponseBody) *QueryClassScheduleConfigResponse {
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
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s QueryDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceRequest) SetSn(v string) *QueryDeviceRequest {
	s.Sn = &v
	return s
}

type QueryDeviceResponseBody struct {
	GmtExpiry *int64  `json:"gmtExpiry,omitempty" xml:"gmtExpiry,omitempty"`
	Model     *string `json:"model,omitempty" xml:"model,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceResponseBody) SetGmtExpiry(v int64) *QueryDeviceResponseBody {
	s.GmtExpiry = &v
	return s
}

func (s *QueryDeviceResponseBody) SetModel(v string) *QueryDeviceResponseBody {
	s.Model = &v
	return s
}

func (s *QueryDeviceResponseBody) SetName(v string) *QueryDeviceResponseBody {
	s.Name = &v
	return s
}

func (s *QueryDeviceResponseBody) SetSn(v string) *QueryDeviceResponseBody {
	s.Sn = &v
	return s
}

func (s *QueryDeviceResponseBody) SetType(v string) *QueryDeviceResponseBody {
	s.Type = &v
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

type QueryDeviceListByCorpIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDeviceListByCorpIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceListByCorpIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceListByCorpIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDeviceListByCorpIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDeviceListByCorpIdRequest struct {
	Operator   *string `json:"operator,omitempty" xml:"operator,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryDeviceListByCorpIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdRequest) SetOperator(v string) *QueryDeviceListByCorpIdRequest {
	s.Operator = &v
	return s
}

func (s *QueryDeviceListByCorpIdRequest) SetPageNumber(v int32) *QueryDeviceListByCorpIdRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDeviceListByCorpIdRequest) SetPageSize(v int32) *QueryDeviceListByCorpIdRequest {
	s.PageSize = &v
	return s
}

type QueryDeviceListByCorpIdResponseBody struct {
	Result  *QueryDeviceListByCorpIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryDeviceListByCorpIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBody) SetResult(v *QueryDeviceListByCorpIdResponseBodyResult) *QueryDeviceListByCorpIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBody) SetSuccess(v bool) *QueryDeviceListByCorpIdResponseBody {
	s.Success = &v
	return s
}

type QueryDeviceListByCorpIdResponseBodyResult struct {
	HasMore *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryDeviceListByCorpIdResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryDeviceListByCorpIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBodyResult) SetHasMore(v bool) *QueryDeviceListByCorpIdResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResult) SetList(v []*QueryDeviceListByCorpIdResponseBodyResultList) *QueryDeviceListByCorpIdResponseBodyResult {
	s.List = v
	return s
}

type QueryDeviceListByCorpIdResponseBodyResultList struct {
	AppStatus  *int32  `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
}

func (s QueryDeviceListByCorpIdResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetAppStatus(v int32) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.AppStatus = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetDeviceCode(v string) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.DeviceCode = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponseBodyResultList) SetDeviceName(v string) *QueryDeviceListByCorpIdResponseBodyResultList {
	s.DeviceName = &v
	return s
}

type QueryDeviceListByCorpIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDeviceListByCorpIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDeviceListByCorpIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeviceListByCorpIdResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceListByCorpIdResponse) SetHeaders(v map[string]*string) *QueryDeviceListByCorpIdResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceListByCorpIdResponse) SetStatusCode(v int32) *QueryDeviceListByCorpIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceListByCorpIdResponse) SetBody(v *QueryDeviceListByCorpIdResponseBody) *QueryDeviceListByCorpIdResponse {
	s.Body = v
	return s
}

type QueryEduAssetSpacesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryEduAssetSpacesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryEduAssetSpacesHeaders) GoString() string {
	return s.String()
}

func (s *QueryEduAssetSpacesHeaders) SetCommonHeaders(v map[string]*string) *QueryEduAssetSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryEduAssetSpacesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryEduAssetSpacesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryEduAssetSpacesRequest struct {
	BizCode    *string `json:"bizCode,omitempty" xml:"bizCode,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryEduAssetSpacesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEduAssetSpacesRequest) GoString() string {
	return s.String()
}

func (s *QueryEduAssetSpacesRequest) SetBizCode(v string) *QueryEduAssetSpacesRequest {
	s.BizCode = &v
	return s
}

func (s *QueryEduAssetSpacesRequest) SetMaxResults(v int32) *QueryEduAssetSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryEduAssetSpacesRequest) SetNextToken(v int64) *QueryEduAssetSpacesRequest {
	s.NextToken = &v
	return s
}

type QueryEduAssetSpacesResponseBody struct {
	HasMore   *bool                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *string                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Spaces    []*QueryEduAssetSpacesResponseBodySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
}

func (s QueryEduAssetSpacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEduAssetSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEduAssetSpacesResponseBody) SetHasMore(v bool) *QueryEduAssetSpacesResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBody) SetNextToken(v string) *QueryEduAssetSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBody) SetSpaces(v []*QueryEduAssetSpacesResponseBodySpaces) *QueryEduAssetSpacesResponseBody {
	s.Spaces = v
	return s
}

type QueryEduAssetSpacesResponseBodySpaces struct {
	CreateTimeMillis *int64  `json:"createTimeMillis,omitempty" xml:"createTimeMillis,omitempty"`
	ModifyTimeMillis *int64  `json:"modifyTimeMillis,omitempty" xml:"modifyTimeMillis,omitempty"`
	PermissionMode   *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	Quota            *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
	SpaceId          *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceName        *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType        *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	UsedQuota        *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s QueryEduAssetSpacesResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s QueryEduAssetSpacesResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetCreateTimeMillis(v int64) *QueryEduAssetSpacesResponseBodySpaces {
	s.CreateTimeMillis = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetModifyTimeMillis(v int64) *QueryEduAssetSpacesResponseBodySpaces {
	s.ModifyTimeMillis = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetPermissionMode(v string) *QueryEduAssetSpacesResponseBodySpaces {
	s.PermissionMode = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetQuota(v int64) *QueryEduAssetSpacesResponseBodySpaces {
	s.Quota = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetSpaceId(v string) *QueryEduAssetSpacesResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetSpaceName(v string) *QueryEduAssetSpacesResponseBodySpaces {
	s.SpaceName = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetSpaceType(v string) *QueryEduAssetSpacesResponseBodySpaces {
	s.SpaceType = &v
	return s
}

func (s *QueryEduAssetSpacesResponseBodySpaces) SetUsedQuota(v int64) *QueryEduAssetSpacesResponseBodySpaces {
	s.UsedQuota = &v
	return s
}

type QueryEduAssetSpacesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEduAssetSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEduAssetSpacesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEduAssetSpacesResponse) GoString() string {
	return s.String()
}

func (s *QueryEduAssetSpacesResponse) SetHeaders(v map[string]*string) *QueryEduAssetSpacesResponse {
	s.Headers = v
	return s
}

func (s *QueryEduAssetSpacesResponse) SetStatusCode(v int32) *QueryEduAssetSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEduAssetSpacesResponse) SetBody(v *QueryEduAssetSpacesResponseBody) *QueryEduAssetSpacesResponse {
	s.Body = v
	return s
}

type QueryGroupIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupIdHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupIdRequest struct {
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s QueryGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupIdRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupIdRequest) SetSn(v string) *QueryGroupIdRequest {
	s.Sn = &v
	return s
}

type QueryGroupIdResponseBody struct {
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GroupId      *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MerchantId   *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantName *string `json:"merchantName,omitempty" xml:"merchantName,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Pid          *string `json:"pid,omitempty" xml:"pid,omitempty"`
}

func (s QueryGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupIdResponseBody) SetCorpId(v string) *QueryGroupIdResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryGroupIdResponseBody) SetGroupId(v string) *QueryGroupIdResponseBody {
	s.GroupId = &v
	return s
}

func (s *QueryGroupIdResponseBody) SetMerchantId(v string) *QueryGroupIdResponseBody {
	s.MerchantId = &v
	return s
}

func (s *QueryGroupIdResponseBody) SetMerchantName(v string) *QueryGroupIdResponseBody {
	s.MerchantName = &v
	return s
}

func (s *QueryGroupIdResponseBody) SetName(v string) *QueryGroupIdResponseBody {
	s.Name = &v
	return s
}

func (s *QueryGroupIdResponseBody) SetPid(v string) *QueryGroupIdResponseBody {
	s.Pid = &v
	return s
}

type QueryGroupIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupIdResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupIdResponse) SetHeaders(v map[string]*string) *QueryGroupIdResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupIdResponse) SetStatusCode(v int32) *QueryGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupIdResponse) SetBody(v *QueryGroupIdResponseBody) *QueryGroupIdResponse {
	s.Body = v
	return s
}

type QueryOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrderHeaders) SetCommonHeaders(v map[string]*string) *QueryOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrderRequest struct {
	AlipayAppId *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId  *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo     *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s QueryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) SetAlipayAppId(v string) *QueryOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *QueryOrderRequest) SetMerchantId(v string) *QueryOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *QueryOrderRequest) SetOrderNo(v string) *QueryOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *QueryOrderRequest) SetSignature(v string) *QueryOrderRequest {
	s.Signature = &v
	return s
}

type QueryOrderResponseBody struct {
	ActualAmount         *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId          *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	CloseTime            *string `json:"closeTime,omitempty" xml:"closeTime,omitempty"`
	CloseTimestamp       *int64  `json:"closeTimestamp,omitempty" xml:"closeTimestamp,omitempty"`
	CreateTime           *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreateTimestamp      *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	LabelAmount          *int64  `json:"labelAmount,omitempty" xml:"labelAmount,omitempty"`
	MerchantId           *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantMergeOrderNo *string `json:"merchantMergeOrderNo,omitempty" xml:"merchantMergeOrderNo,omitempty"`
	MerchantOrderNo      *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo              *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrderType            *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	OuterUserId          *string `json:"outerUserId,omitempty" xml:"outerUserId,omitempty"`
	PayLogonId           *string `json:"payLogonId,omitempty" xml:"payLogonId,omitempty"`
	PayStatus            *int32  `json:"payStatus,omitempty" xml:"payStatus,omitempty"`
	PayTime              *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	PayTimestamp         *int64  `json:"payTimestamp,omitempty" xml:"payTimestamp,omitempty"`
	PayType              *string `json:"payType,omitempty" xml:"payType,omitempty"`
	RefundAmount         *int64  `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	RefundStatus         *int32  `json:"refundStatus,omitempty" xml:"refundStatus,omitempty"`
	RefundTime           *string `json:"refundTime,omitempty" xml:"refundTime,omitempty"`
	RefundTimestamp      *int64  `json:"refundTimestamp,omitempty" xml:"refundTimestamp,omitempty"`
	Subject              *string `json:"subject,omitempty" xml:"subject,omitempty"`
	TradeNo              *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
}

func (s QueryOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBody) SetActualAmount(v int64) *QueryOrderResponseBody {
	s.ActualAmount = &v
	return s
}

func (s *QueryOrderResponseBody) SetAlipayAppId(v string) *QueryOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *QueryOrderResponseBody) SetCloseTime(v string) *QueryOrderResponseBody {
	s.CloseTime = &v
	return s
}

func (s *QueryOrderResponseBody) SetCloseTimestamp(v int64) *QueryOrderResponseBody {
	s.CloseTimestamp = &v
	return s
}

func (s *QueryOrderResponseBody) SetCreateTime(v string) *QueryOrderResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryOrderResponseBody) SetCreateTimestamp(v int64) *QueryOrderResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryOrderResponseBody) SetLabelAmount(v int64) *QueryOrderResponseBody {
	s.LabelAmount = &v
	return s
}

func (s *QueryOrderResponseBody) SetMerchantId(v string) *QueryOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *QueryOrderResponseBody) SetMerchantMergeOrderNo(v string) *QueryOrderResponseBody {
	s.MerchantMergeOrderNo = &v
	return s
}

func (s *QueryOrderResponseBody) SetMerchantOrderNo(v string) *QueryOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *QueryOrderResponseBody) SetOrderNo(v string) *QueryOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *QueryOrderResponseBody) SetOrderType(v string) *QueryOrderResponseBody {
	s.OrderType = &v
	return s
}

func (s *QueryOrderResponseBody) SetOuterUserId(v string) *QueryOrderResponseBody {
	s.OuterUserId = &v
	return s
}

func (s *QueryOrderResponseBody) SetPayLogonId(v string) *QueryOrderResponseBody {
	s.PayLogonId = &v
	return s
}

func (s *QueryOrderResponseBody) SetPayStatus(v int32) *QueryOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *QueryOrderResponseBody) SetPayTime(v string) *QueryOrderResponseBody {
	s.PayTime = &v
	return s
}

func (s *QueryOrderResponseBody) SetPayTimestamp(v int64) *QueryOrderResponseBody {
	s.PayTimestamp = &v
	return s
}

func (s *QueryOrderResponseBody) SetPayType(v string) *QueryOrderResponseBody {
	s.PayType = &v
	return s
}

func (s *QueryOrderResponseBody) SetRefundAmount(v int64) *QueryOrderResponseBody {
	s.RefundAmount = &v
	return s
}

func (s *QueryOrderResponseBody) SetRefundStatus(v int32) *QueryOrderResponseBody {
	s.RefundStatus = &v
	return s
}

func (s *QueryOrderResponseBody) SetRefundTime(v string) *QueryOrderResponseBody {
	s.RefundTime = &v
	return s
}

func (s *QueryOrderResponseBody) SetRefundTimestamp(v int64) *QueryOrderResponseBody {
	s.RefundTimestamp = &v
	return s
}

func (s *QueryOrderResponseBody) SetSubject(v string) *QueryOrderResponseBody {
	s.Subject = &v
	return s
}

func (s *QueryOrderResponseBody) SetTradeNo(v string) *QueryOrderResponseBody {
	s.TradeNo = &v
	return s
}

type QueryOrderResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) SetHeaders(v map[string]*string) *QueryOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderResponse) SetStatusCode(v int32) *QueryOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderResponse) SetBody(v *QueryOrderResponseBody) *QueryOrderResponse {
	s.Body = v
	return s
}

type QueryOrgRelationListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgRelationListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgRelationListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgRelationListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgRelationListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgRelationListRequest struct {
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s QueryOrgRelationListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListRequest) SetOperator(v string) *QueryOrgRelationListRequest {
	s.Operator = &v
	return s
}

type QueryOrgRelationListResponseBody struct {
	Result  []*QueryOrgRelationListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgRelationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponseBody) SetResult(v []*QueryOrgRelationListResponseBodyResult) *QueryOrgRelationListResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrgRelationListResponseBody) SetSuccess(v bool) *QueryOrgRelationListResponseBody {
	s.Success = &v
	return s
}

type QueryOrgRelationListResponseBodyResult struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceCount *int32  `json:"deviceCount,omitempty" xml:"deviceCount,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryOrgRelationListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponseBodyResult) SetCorpId(v string) *QueryOrgRelationListResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryOrgRelationListResponseBodyResult) SetDeviceCount(v int32) *QueryOrgRelationListResponseBodyResult {
	s.DeviceCount = &v
	return s
}

func (s *QueryOrgRelationListResponseBodyResult) SetName(v string) *QueryOrgRelationListResponseBodyResult {
	s.Name = &v
	return s
}

type QueryOrgRelationListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrgRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgRelationListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgRelationListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgRelationListResponse) SetHeaders(v map[string]*string) *QueryOrgRelationListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgRelationListResponse) SetStatusCode(v int32) *QueryOrgRelationListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgRelationListResponse) SetBody(v *QueryOrgRelationListResponseBody) *QueryOrgRelationListResponse {
	s.Body = v
	return s
}

type QueryOrgSecretKeyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgSecretKeyHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgSecretKeyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgSecretKeyHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgSecretKeyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgSecretKeyRequest struct {
	IsvCode  *string `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryOrgSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyRequest) SetIsvCode(v string) *QueryOrgSecretKeyRequest {
	s.IsvCode = &v
	return s
}

func (s *QueryOrgSecretKeyRequest) SetOpUserId(v string) *QueryOrgSecretKeyRequest {
	s.OpUserId = &v
	return s
}

type QueryOrgSecretKeyResponseBody struct {
	UniversitySecretKeyInfo *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo `json:"universitySecretKeyInfo,omitempty" xml:"universitySecretKeyInfo,omitempty" type:"Struct"`
}

func (s QueryOrgSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponseBody) SetUniversitySecretKeyInfo(v *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) *QueryOrgSecretKeyResponseBody {
	s.UniversitySecretKeyInfo = v
	return s
}

type QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo struct {
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
}

func (s QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo) SetSecretKey(v string) *QueryOrgSecretKeyResponseBodyUniversitySecretKeyInfo {
	s.SecretKey = &v
	return s
}

type QueryOrgSecretKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrgSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgSecretKeyResponse) SetHeaders(v map[string]*string) *QueryOrgSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgSecretKeyResponse) SetStatusCode(v int32) *QueryOrgSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgSecretKeyResponse) SetBody(v *QueryOrgSecretKeyResponseBody) *QueryOrgSecretKeyResponse {
	s.Body = v
	return s
}

type QueryOrgTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgTypeResponseBody struct {
	OrgType *int64 `json:"orgType,omitempty" xml:"orgType,omitempty"`
}

func (s QueryOrgTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeResponseBody) SetOrgType(v int64) *QueryOrgTypeResponseBody {
	s.OrgType = &v
	return s
}

type QueryOrgTypeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrgTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTypeResponse) SetHeaders(v map[string]*string) *QueryOrgTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTypeResponse) SetStatusCode(v int32) *QueryOrgTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgTypeResponse) SetBody(v *QueryOrgTypeResponseBody) *QueryOrgTypeResponse {
	s.Body = v
	return s
}

type QueryPayResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPayResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPayResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryPayResultHeaders) SetCommonHeaders(v map[string]*string) *QueryPayResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPayResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPayResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPayResultRequest struct {
	FaceId    *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	OrderNo   *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Sn        *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryPayResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPayResultRequest) GoString() string {
	return s.String()
}

func (s *QueryPayResultRequest) SetFaceId(v string) *QueryPayResultRequest {
	s.FaceId = &v
	return s
}

func (s *QueryPayResultRequest) SetOrderNo(v string) *QueryPayResultRequest {
	s.OrderNo = &v
	return s
}

func (s *QueryPayResultRequest) SetSignature(v string) *QueryPayResultRequest {
	s.Signature = &v
	return s
}

func (s *QueryPayResultRequest) SetSn(v string) *QueryPayResultRequest {
	s.Sn = &v
	return s
}

func (s *QueryPayResultRequest) SetTimestamp(v int64) *QueryPayResultRequest {
	s.Timestamp = &v
	return s
}

func (s *QueryPayResultRequest) SetUserId(v string) *QueryPayResultRequest {
	s.UserId = &v
	return s
}

func (s *QueryPayResultRequest) SetVersion(v string) *QueryPayResultRequest {
	s.Version = &v
	return s
}

type QueryPayResultResponseBody struct {
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryPayResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPayResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPayResultResponseBody) SetStatus(v int32) *QueryPayResultResponseBody {
	s.Status = &v
	return s
}

type QueryPayResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPayResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPayResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPayResultResponse) GoString() string {
	return s.String()
}

func (s *QueryPayResultResponse) SetHeaders(v map[string]*string) *QueryPayResultResponse {
	s.Headers = v
	return s
}

func (s *QueryPayResultResponse) SetStatusCode(v int32) *QueryPayResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPayResultResponse) SetBody(v *QueryPayResultResponseBody) *QueryPayResultResponse {
	s.Body = v
	return s
}

type QueryPhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *QueryPhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPhysicalClassroomRequest struct {
	ClassroomId *int64  `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	OpUserId    *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryPhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomRequest) SetClassroomId(v int64) *QueryPhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

func (s *QueryPhysicalClassroomRequest) SetOpUserId(v string) *QueryPhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

type QueryPhysicalClassroomResponseBody struct {
	Result  *QueryPhysicalClassroomResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponseBody) SetResult(v *QueryPhysicalClassroomResponseBodyResult) *QueryPhysicalClassroomResponseBody {
	s.Result = v
	return s
}

func (s *QueryPhysicalClassroomResponseBody) SetSuccess(v bool) *QueryPhysicalClassroomResponseBody {
	s.Success = &v
	return s
}

type QueryPhysicalClassroomResponseBodyResult struct {
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	ClassroomCampus   *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	ClassroomFloor    *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	ClassroomId       *int64  `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	ClassroomName     *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	ClassroomNumber   *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
	DirectBroadcast   *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
}

func (s QueryPhysicalClassroomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomBuilding(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomBuilding = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomCampus(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomCampus = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomFloor(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomFloor = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomId(v int64) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomId = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomName(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomName = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetClassroomNumber(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.ClassroomNumber = &v
	return s
}

func (s *QueryPhysicalClassroomResponseBodyResult) SetDirectBroadcast(v string) *QueryPhysicalClassroomResponseBodyResult {
	s.DirectBroadcast = &v
	return s
}

type QueryPhysicalClassroomResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *QueryPhysicalClassroomResponse) SetHeaders(v map[string]*string) *QueryPhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *QueryPhysicalClassroomResponse) SetStatusCode(v int32) *QueryPhysicalClassroomResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPhysicalClassroomResponse) SetBody(v *QueryPhysicalClassroomResponseBody) *QueryPhysicalClassroomResponse {
	s.Body = v
	return s
}

type QueryPurchaseInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPurchaseInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchaseInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryPurchaseInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryPurchaseInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPurchaseInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPurchaseInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPurchaseInfoRequest struct {
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Scene      *int32  `json:"scene,omitempty" xml:"scene,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPurchaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchaseInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPurchaseInfoRequest) SetMerchantId(v string) *QueryPurchaseInfoRequest {
	s.MerchantId = &v
	return s
}

func (s *QueryPurchaseInfoRequest) SetScene(v int32) *QueryPurchaseInfoRequest {
	s.Scene = &v
	return s
}

func (s *QueryPurchaseInfoRequest) SetSn(v string) *QueryPurchaseInfoRequest {
	s.Sn = &v
	return s
}

func (s *QueryPurchaseInfoRequest) SetUserId(v string) *QueryPurchaseInfoRequest {
	s.UserId = &v
	return s
}

type QueryPurchaseInfoResponseBody struct {
	CorpId     *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Scene      *int32  `json:"scene,omitempty" xml:"scene,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPurchaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchaseInfoResponseBody) SetCorpId(v string) *QueryPurchaseInfoResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryPurchaseInfoResponseBody) SetMerchantId(v string) *QueryPurchaseInfoResponseBody {
	s.MerchantId = &v
	return s
}

func (s *QueryPurchaseInfoResponseBody) SetName(v string) *QueryPurchaseInfoResponseBody {
	s.Name = &v
	return s
}

func (s *QueryPurchaseInfoResponseBody) SetScene(v int32) *QueryPurchaseInfoResponseBody {
	s.Scene = &v
	return s
}

func (s *QueryPurchaseInfoResponseBody) SetStatus(v int32) *QueryPurchaseInfoResponseBody {
	s.Status = &v
	return s
}

func (s *QueryPurchaseInfoResponseBody) SetUserId(v string) *QueryPurchaseInfoResponseBody {
	s.UserId = &v
	return s
}

type QueryPurchaseInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPurchaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPurchaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchaseInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchaseInfoResponse) SetHeaders(v map[string]*string) *QueryPurchaseInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchaseInfoResponse) SetStatusCode(v int32) *QueryPurchaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchaseInfoResponse) SetBody(v *QueryPurchaseInfoResponseBody) *QueryPurchaseInfoResponse {
	s.Body = v
	return s
}

type QueryRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *QueryRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRemoteClassCourseRequest struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Operator  *string `json:"operator,omitempty" xml:"operator,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseRequest) SetEndTime(v int64) *QueryRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRemoteClassCourseRequest) SetOperator(v string) *QueryRemoteClassCourseRequest {
	s.Operator = &v
	return s
}

func (s *QueryRemoteClassCourseRequest) SetStartTime(v int64) *QueryRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

type QueryRemoteClassCourseResponseBody struct {
	Result  []*QueryRemoteClassCourseResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBody) SetResult(v []*QueryRemoteClassCourseResponseBodyResult) *QueryRemoteClassCourseResponseBody {
	s.Result = v
	return s
}

func (s *QueryRemoteClassCourseResponseBody) SetSuccess(v bool) *QueryRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type QueryRemoteClassCourseResponseBodyResult struct {
	AttendParticipants  []*QueryRemoteClassCourseResponseBodyResultAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	CanEdit             *bool                                                         `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	CourseCode          *string                                                       `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	CourseName          *string                                                       `json:"courseName,omitempty" xml:"courseName,omitempty"`
	CourseWays          []*string                                                     `json:"courseWays,omitempty" xml:"courseWays,omitempty" type:"Repeated"`
	EndTime             *int64                                                        `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime           *int64                                                        `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status              *int32                                                        `json:"status,omitempty" xml:"status,omitempty"`
	TeachingParticipant *QueryRemoteClassCourseResponseBodyResultTeachingParticipant  `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
}

func (s QueryRemoteClassCourseResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetAttendParticipants(v []*QueryRemoteClassCourseResponseBodyResultAttendParticipants) *QueryRemoteClassCourseResponseBodyResult {
	s.AttendParticipants = v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCanEdit(v bool) *QueryRemoteClassCourseResponseBodyResult {
	s.CanEdit = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCourseCode(v string) *QueryRemoteClassCourseResponseBodyResult {
	s.CourseCode = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCourseName(v string) *QueryRemoteClassCourseResponseBodyResult {
	s.CourseName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetCourseWays(v []*string) *QueryRemoteClassCourseResponseBodyResult {
	s.CourseWays = v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetEndTime(v int64) *QueryRemoteClassCourseResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetStartTime(v int64) *QueryRemoteClassCourseResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetStatus(v int32) *QueryRemoteClassCourseResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResult) SetTeachingParticipant(v *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) *QueryRemoteClassCourseResponseBodyResult {
	s.TeachingParticipant = v
	return s
}

type QueryRemoteClassCourseResponseBodyResultAttendParticipants struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrgName         *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ParticipantId   *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
}

func (s QueryRemoteClassCourseResponseBodyResultAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResultAttendParticipants) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetCorpId(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetOrgName(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.OrgName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantId(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultAttendParticipants) SetParticipantName(v string) *QueryRemoteClassCourseResponseBodyResultAttendParticipants {
	s.ParticipantName = &v
	return s
}

type QueryRemoteClassCourseResponseBodyResultTeachingParticipant struct {
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OrgName         *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ParticipantId   *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
	ParticipantName *string `json:"participantName,omitempty" xml:"participantName,omitempty"`
}

func (s QueryRemoteClassCourseResponseBodyResultTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponseBodyResultTeachingParticipant) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetCorpId(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetOrgName(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.OrgName = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantId(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantId = &v
	return s
}

func (s *QueryRemoteClassCourseResponseBodyResultTeachingParticipant) SetParticipantName(v string) *QueryRemoteClassCourseResponseBodyResultTeachingParticipant {
	s.ParticipantName = &v
	return s
}

type QueryRemoteClassCourseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *QueryRemoteClassCourseResponse) SetHeaders(v map[string]*string) *QueryRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *QueryRemoteClassCourseResponse) SetStatusCode(v int32) *QueryRemoteClassCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRemoteClassCourseResponse) SetBody(v *QueryRemoteClassCourseResponseBody) *QueryRemoteClassCourseResponse {
	s.Body = v
	return s
}

type QuerySchoolUserFaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySchoolUserFaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySchoolUserFaceHeaders) GoString() string {
	return s.String()
}

func (s *QuerySchoolUserFaceHeaders) SetCommonHeaders(v map[string]*string) *QuerySchoolUserFaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySchoolUserFaceHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySchoolUserFaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySchoolUserFaceRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QuerySchoolUserFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchoolUserFaceRequest) GoString() string {
	return s.String()
}

func (s *QuerySchoolUserFaceRequest) SetPageNumber(v int32) *QuerySchoolUserFaceRequest {
	s.PageNumber = &v
	return s
}

func (s *QuerySchoolUserFaceRequest) SetPageSize(v int32) *QuerySchoolUserFaceRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySchoolUserFaceRequest) SetSn(v string) *QuerySchoolUserFaceRequest {
	s.Sn = &v
	return s
}

func (s *QuerySchoolUserFaceRequest) SetType(v int32) *QuerySchoolUserFaceRequest {
	s.Type = &v
	return s
}

type QuerySchoolUserFaceResponseBody struct {
	CorpId       *string                                        `json:"corpId,omitempty" xml:"corpId,omitempty"`
	HasMore      *bool                                          `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	UserFaceList []*QuerySchoolUserFaceResponseBodyUserFaceList `json:"userFaceList,omitempty" xml:"userFaceList,omitempty" type:"Repeated"`
}

func (s QuerySchoolUserFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchoolUserFaceResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchoolUserFaceResponseBody) SetCorpId(v string) *QuerySchoolUserFaceResponseBody {
	s.CorpId = &v
	return s
}

func (s *QuerySchoolUserFaceResponseBody) SetHasMore(v bool) *QuerySchoolUserFaceResponseBody {
	s.HasMore = &v
	return s
}

func (s *QuerySchoolUserFaceResponseBody) SetUserFaceList(v []*QuerySchoolUserFaceResponseBodyUserFaceList) *QuerySchoolUserFaceResponseBody {
	s.UserFaceList = v
	return s
}

type QuerySchoolUserFaceResponseBodyUserFaceList struct {
	FaceId *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySchoolUserFaceResponseBodyUserFaceList) String() string {
	return tea.Prettify(s)
}

func (s QuerySchoolUserFaceResponseBodyUserFaceList) GoString() string {
	return s.String()
}

func (s *QuerySchoolUserFaceResponseBodyUserFaceList) SetFaceId(v string) *QuerySchoolUserFaceResponseBodyUserFaceList {
	s.FaceId = &v
	return s
}

func (s *QuerySchoolUserFaceResponseBodyUserFaceList) SetName(v string) *QuerySchoolUserFaceResponseBodyUserFaceList {
	s.Name = &v
	return s
}

func (s *QuerySchoolUserFaceResponseBodyUserFaceList) SetStatus(v int32) *QuerySchoolUserFaceResponseBodyUserFaceList {
	s.Status = &v
	return s
}

func (s *QuerySchoolUserFaceResponseBodyUserFaceList) SetUserId(v string) *QuerySchoolUserFaceResponseBodyUserFaceList {
	s.UserId = &v
	return s
}

type QuerySchoolUserFaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySchoolUserFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySchoolUserFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchoolUserFaceResponse) GoString() string {
	return s.String()
}

func (s *QuerySchoolUserFaceResponse) SetHeaders(v map[string]*string) *QuerySchoolUserFaceResponse {
	s.Headers = v
	return s
}

func (s *QuerySchoolUserFaceResponse) SetStatusCode(v int32) *QuerySchoolUserFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchoolUserFaceResponse) SetBody(v *QuerySchoolUserFaceResponseBody) *QuerySchoolUserFaceResponse {
	s.Body = v
	return s
}

type QuerySnsOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySnsOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySnsOrderHeaders) GoString() string {
	return s.String()
}

func (s *QuerySnsOrderHeaders) SetCommonHeaders(v map[string]*string) *QuerySnsOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySnsOrderHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySnsOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySnsOrderRequest struct {
	AlipayAppId *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	MerchantId  *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	OrderNo     *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s QuerySnsOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySnsOrderRequest) GoString() string {
	return s.String()
}

func (s *QuerySnsOrderRequest) SetAlipayAppId(v string) *QuerySnsOrderRequest {
	s.AlipayAppId = &v
	return s
}

func (s *QuerySnsOrderRequest) SetMerchantId(v string) *QuerySnsOrderRequest {
	s.MerchantId = &v
	return s
}

func (s *QuerySnsOrderRequest) SetOrderNo(v string) *QuerySnsOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *QuerySnsOrderRequest) SetSignature(v string) *QuerySnsOrderRequest {
	s.Signature = &v
	return s
}

type QuerySnsOrderResponseBody struct {
	ActualAmount         *int64  `json:"actualAmount,omitempty" xml:"actualAmount,omitempty"`
	AlipayAppId          *string `json:"alipayAppId,omitempty" xml:"alipayAppId,omitempty"`
	CloseTime            *string `json:"closeTime,omitempty" xml:"closeTime,omitempty"`
	CloseTimestamp       *int64  `json:"closeTimestamp,omitempty" xml:"closeTimestamp,omitempty"`
	CreateTime           *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreateTimestamp      *int64  `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	LabelAmount          *int64  `json:"labelAmount,omitempty" xml:"labelAmount,omitempty"`
	MerchantId           *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	MerchantMergeOrderNo *string `json:"merchantMergeOrderNo,omitempty" xml:"merchantMergeOrderNo,omitempty"`
	MerchantOrderNo      *string `json:"merchantOrderNo,omitempty" xml:"merchantOrderNo,omitempty"`
	OrderNo              *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrderType            *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	OuterUserId          *string `json:"outerUserId,omitempty" xml:"outerUserId,omitempty"`
	PayLogonId           *string `json:"payLogonId,omitempty" xml:"payLogonId,omitempty"`
	PayStatus            *int32  `json:"payStatus,omitempty" xml:"payStatus,omitempty"`
	PayTime              *string `json:"payTime,omitempty" xml:"payTime,omitempty"`
	PayTimestamp         *int64  `json:"payTimestamp,omitempty" xml:"payTimestamp,omitempty"`
	PayType              *string `json:"payType,omitempty" xml:"payType,omitempty"`
	RefundAmount         *int64  `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"`
	RefundStatus         *int32  `json:"refundStatus,omitempty" xml:"refundStatus,omitempty"`
	RefundTime           *string `json:"refundTime,omitempty" xml:"refundTime,omitempty"`
	RefundTimestamp      *int64  `json:"refundTimestamp,omitempty" xml:"refundTimestamp,omitempty"`
	Subject              *string `json:"subject,omitempty" xml:"subject,omitempty"`
	TradeNo              *string `json:"tradeNo,omitempty" xml:"tradeNo,omitempty"`
}

func (s QuerySnsOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySnsOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySnsOrderResponseBody) SetActualAmount(v int64) *QuerySnsOrderResponseBody {
	s.ActualAmount = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetAlipayAppId(v string) *QuerySnsOrderResponseBody {
	s.AlipayAppId = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetCloseTime(v string) *QuerySnsOrderResponseBody {
	s.CloseTime = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetCloseTimestamp(v int64) *QuerySnsOrderResponseBody {
	s.CloseTimestamp = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetCreateTime(v string) *QuerySnsOrderResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetCreateTimestamp(v int64) *QuerySnsOrderResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetLabelAmount(v int64) *QuerySnsOrderResponseBody {
	s.LabelAmount = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetMerchantId(v string) *QuerySnsOrderResponseBody {
	s.MerchantId = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetMerchantMergeOrderNo(v string) *QuerySnsOrderResponseBody {
	s.MerchantMergeOrderNo = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetMerchantOrderNo(v string) *QuerySnsOrderResponseBody {
	s.MerchantOrderNo = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetOrderNo(v string) *QuerySnsOrderResponseBody {
	s.OrderNo = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetOrderType(v string) *QuerySnsOrderResponseBody {
	s.OrderType = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetOuterUserId(v string) *QuerySnsOrderResponseBody {
	s.OuterUserId = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetPayLogonId(v string) *QuerySnsOrderResponseBody {
	s.PayLogonId = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetPayStatus(v int32) *QuerySnsOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetPayTime(v string) *QuerySnsOrderResponseBody {
	s.PayTime = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetPayTimestamp(v int64) *QuerySnsOrderResponseBody {
	s.PayTimestamp = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetPayType(v string) *QuerySnsOrderResponseBody {
	s.PayType = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetRefundAmount(v int64) *QuerySnsOrderResponseBody {
	s.RefundAmount = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetRefundStatus(v int32) *QuerySnsOrderResponseBody {
	s.RefundStatus = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetRefundTime(v string) *QuerySnsOrderResponseBody {
	s.RefundTime = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetRefundTimestamp(v int64) *QuerySnsOrderResponseBody {
	s.RefundTimestamp = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetSubject(v string) *QuerySnsOrderResponseBody {
	s.Subject = &v
	return s
}

func (s *QuerySnsOrderResponseBody) SetTradeNo(v string) *QuerySnsOrderResponseBody {
	s.TradeNo = &v
	return s
}

type QuerySnsOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySnsOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySnsOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySnsOrderResponse) GoString() string {
	return s.String()
}

func (s *QuerySnsOrderResponse) SetHeaders(v map[string]*string) *QuerySnsOrderResponse {
	s.Headers = v
	return s
}

func (s *QuerySnsOrderResponse) SetStatusCode(v int32) *QuerySnsOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySnsOrderResponse) SetBody(v *QuerySnsOrderResponseBody) *QuerySnsOrderResponse {
	s.Body = v
	return s
}

type QueryStatisticsDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryStatisticsDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataHeaders) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataHeaders) SetCommonHeaders(v map[string]*string) *QueryStatisticsDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryStatisticsDataHeaders) SetXAcsDingtalkAccessToken(v string) *QueryStatisticsDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryStatisticsDataRequest struct {
	SectionIndexList []*int64  `json:"sectionIndexList,omitempty" xml:"sectionIndexList,omitempty" type:"Repeated"`
	TeacherUserIds   []*string `json:"teacherUserIds,omitempty" xml:"teacherUserIds,omitempty" type:"Repeated"`
	EndTime          *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpUserId         *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	StartTime        *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s QueryStatisticsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataRequest) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataRequest) SetSectionIndexList(v []*int64) *QueryStatisticsDataRequest {
	s.SectionIndexList = v
	return s
}

func (s *QueryStatisticsDataRequest) SetTeacherUserIds(v []*string) *QueryStatisticsDataRequest {
	s.TeacherUserIds = v
	return s
}

func (s *QueryStatisticsDataRequest) SetEndTime(v int64) *QueryStatisticsDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStatisticsDataRequest) SetOpUserId(v string) *QueryStatisticsDataRequest {
	s.OpUserId = &v
	return s
}

func (s *QueryStatisticsDataRequest) SetStartTime(v int64) *QueryStatisticsDataRequest {
	s.StartTime = &v
	return s
}

type QueryStatisticsDataResponseBody struct {
	Result []*QueryStatisticsDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryStatisticsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponseBody) SetResult(v []*QueryStatisticsDataResponseBodyResult) *QueryStatisticsDataResponseBody {
	s.Result = v
	return s
}

type QueryStatisticsDataResponseBodyResult struct {
	ClassId     *int64   `json:"classId,omitempty" xml:"classId,omitempty"`
	CourseCount *int64   `json:"courseCount,omitempty" xml:"courseCount,omitempty"`
	CourseHours *float32 `json:"courseHours,omitempty" xml:"courseHours,omitempty"`
	SubjectCode *string  `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	SubjectName *int64   `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	UserId      *string  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName    *string  `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryStatisticsDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponseBodyResult) SetClassId(v int64) *QueryStatisticsDataResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetCourseCount(v int64) *QueryStatisticsDataResponseBodyResult {
	s.CourseCount = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetCourseHours(v float32) *QueryStatisticsDataResponseBodyResult {
	s.CourseHours = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetSubjectCode(v string) *QueryStatisticsDataResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetSubjectName(v int64) *QueryStatisticsDataResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetUserId(v string) *QueryStatisticsDataResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryStatisticsDataResponseBodyResult) SetUserName(v string) *QueryStatisticsDataResponseBodyResult {
	s.UserName = &v
	return s
}

type QueryStatisticsDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStatisticsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatisticsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsDataResponse) GoString() string {
	return s.String()
}

func (s *QueryStatisticsDataResponse) SetHeaders(v map[string]*string) *QueryStatisticsDataResponse {
	s.Headers = v
	return s
}

func (s *QueryStatisticsDataResponse) SetStatusCode(v int32) *QueryStatisticsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStatisticsDataResponse) SetBody(v *QueryStatisticsDataResponseBody) *QueryStatisticsDataResponse {
	s.Body = v
	return s
}

type QuerySubjectTeachersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySubjectTeachersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersHeaders) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersHeaders) SetCommonHeaders(v map[string]*string) *QuerySubjectTeachersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySubjectTeachersHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySubjectTeachersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySubjectTeachersRequest struct {
	ClassIds    []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	OpUserId    *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	SubjectCode *string  `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
}

func (s QuerySubjectTeachersRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersRequest) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersRequest) SetClassIds(v []*int64) *QuerySubjectTeachersRequest {
	s.ClassIds = v
	return s
}

func (s *QuerySubjectTeachersRequest) SetOpUserId(v string) *QuerySubjectTeachersRequest {
	s.OpUserId = &v
	return s
}

func (s *QuerySubjectTeachersRequest) SetSubjectCode(v string) *QuerySubjectTeachersRequest {
	s.SubjectCode = &v
	return s
}

type QuerySubjectTeachersResponseBody struct {
	Result []*QuerySubjectTeachersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QuerySubjectTeachersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponseBody) SetResult(v []*QuerySubjectTeachersResponseBodyResult) *QuerySubjectTeachersResponseBody {
	s.Result = v
	return s
}

type QuerySubjectTeachersResponseBodyResult struct {
	ClassId       *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	PeriodCode    *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	SubjectCode   *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	SubjectName   *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	TeacherName   *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s QuerySubjectTeachersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponseBodyResult) SetClassId(v int64) *QuerySubjectTeachersResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetCorpId(v string) *QuerySubjectTeachersResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetPeriodCode(v string) *QuerySubjectTeachersResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetSubjectCode(v string) *QuerySubjectTeachersResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetSubjectName(v string) *QuerySubjectTeachersResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetTeacherName(v string) *QuerySubjectTeachersResponseBodyResult {
	s.TeacherName = &v
	return s
}

func (s *QuerySubjectTeachersResponseBodyResult) SetTeacherUserId(v string) *QuerySubjectTeachersResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

type QuerySubjectTeachersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySubjectTeachersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySubjectTeachersResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubjectTeachersResponse) GoString() string {
	return s.String()
}

func (s *QuerySubjectTeachersResponse) SetHeaders(v map[string]*string) *QuerySubjectTeachersResponse {
	s.Headers = v
	return s
}

func (s *QuerySubjectTeachersResponse) SetStatusCode(v int32) *QuerySubjectTeachersResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubjectTeachersResponse) SetBody(v *QuerySubjectTeachersResponseBody) *QuerySubjectTeachersResponse {
	s.Body = v
	return s
}

type QueryTeachSubjectsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTeachSubjectsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsHeaders) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsHeaders) SetCommonHeaders(v map[string]*string) *QueryTeachSubjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTeachSubjectsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTeachSubjectsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTeachSubjectsRequest struct {
	ClassIds []*int64 `json:"classIds,omitempty" xml:"classIds,omitempty" type:"Repeated"`
	OpUserId *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryTeachSubjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsRequest) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsRequest) SetClassIds(v []*int64) *QueryTeachSubjectsRequest {
	s.ClassIds = v
	return s
}

func (s *QueryTeachSubjectsRequest) SetOpUserId(v string) *QueryTeachSubjectsRequest {
	s.OpUserId = &v
	return s
}

type QueryTeachSubjectsResponseBody struct {
	Result []*QueryTeachSubjectsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryTeachSubjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponseBody) SetResult(v []*QueryTeachSubjectsResponseBodyResult) *QueryTeachSubjectsResponseBody {
	s.Result = v
	return s
}

type QueryTeachSubjectsResponseBodyResult struct {
	ClassId       *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	PeriodCode    *string `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	SubjectCode   *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	SubjectName   *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	TeacherName   *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
	TeacherUserId *string `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s QueryTeachSubjectsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponseBodyResult) SetClassId(v int64) *QueryTeachSubjectsResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetCorpId(v string) *QueryTeachSubjectsResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetPeriodCode(v string) *QueryTeachSubjectsResponseBodyResult {
	s.PeriodCode = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetSubjectCode(v string) *QueryTeachSubjectsResponseBodyResult {
	s.SubjectCode = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetSubjectName(v string) *QueryTeachSubjectsResponseBodyResult {
	s.SubjectName = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetTeacherName(v string) *QueryTeachSubjectsResponseBodyResult {
	s.TeacherName = &v
	return s
}

func (s *QueryTeachSubjectsResponseBodyResult) SetTeacherUserId(v string) *QueryTeachSubjectsResponseBodyResult {
	s.TeacherUserId = &v
	return s
}

type QueryTeachSubjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTeachSubjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTeachSubjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTeachSubjectsResponse) GoString() string {
	return s.String()
}

func (s *QueryTeachSubjectsResponse) SetHeaders(v map[string]*string) *QueryTeachSubjectsResponse {
	s.Headers = v
	return s
}

func (s *QueryTeachSubjectsResponse) SetStatusCode(v int32) *QueryTeachSubjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTeachSubjectsResponse) SetBody(v *QueryTeachSubjectsResponseBody) *QueryTeachSubjectsResponse {
	s.Body = v
	return s
}

type QueryUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUniversityCourseGroupRequest struct {
	CourseGroupCode *string `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	OpUserId        *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s QueryUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupRequest) SetCourseGroupCode(v string) *QueryUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryUniversityCourseGroupRequest) SetOpUserId(v string) *QueryUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type QueryUniversityCourseGroupResponseBody struct {
	UniversityCourseGroupInfo *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo `json:"universityCourseGroupInfo,omitempty" xml:"universityCourseGroupInfo,omitempty" type:"Struct"`
}

func (s QueryUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBody) SetUniversityCourseGroupInfo(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) *QueryUniversityCourseGroupResponseBody {
	s.UniversityCourseGroupInfo = v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo struct {
	CourseGroupCode        *string                                                                                  `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	CourseGroupIntroduce   *string                                                                                  `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	CourseGroupName        *string                                                                                  `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	CourserGroupItemModels []*QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
	IsvCourseGroupCode     *string                                                                                  `json:"isvCourseGroupCode,omitempty" xml:"isvCourseGroupCode,omitempty"`
	PeriodCode             *string                                                                                  `json:"periodCode,omitempty" xml:"periodCode,omitempty"`
	SchoolYear             *string                                                                                  `json:"schoolYear,omitempty" xml:"schoolYear,omitempty"`
	Semester               *int32                                                                                   `json:"semester,omitempty" xml:"semester,omitempty"`
	SubjectName            *string                                                                                  `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupIntroduce(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourseGroupName(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourseGroupName = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetCourserGroupItemModels(v []*QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.CourserGroupItemModels = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetIsvCourseGroupCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.IsvCourseGroupCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetPeriodCode(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.PeriodCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSchoolYear(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.SchoolYear = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSemester(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.Semester = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo) SetSubjectName(v string) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfo {
	s.SubjectName = &v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels struct {
	ClassPeriodType           *int32                                                                                                          `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	ClassroomId               *int64                                                                                                          `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	CourseType                *int32                                                                                                          `json:"courseType,omitempty" xml:"courseType,omitempty"`
	CourserGroupItemEndDate   *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate   `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	CourserGroupItemStartDate *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	DayOfWeek                 *int32                                                                                                          `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	SectionIndex              []*int32                                                                                                        `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetClassPeriodType(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetClassroomId(v int64) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourseType(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourseType = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourserGroupItemEndDate(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetCourserGroupItemStartDate(v *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetDayOfWeek(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels) SetSectionIndex(v []*int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

type QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *QueryUniversityCourseGroupResponseBodyUniversityCourseGroupInfoCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

type QueryUniversityCourseGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *QueryUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryUniversityCourseGroupResponse) SetStatusCode(v int32) *QueryUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUniversityCourseGroupResponse) SetBody(v *QueryUniversityCourseGroupResponseBody) *QueryUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type QueryUserFaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserFaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFaceHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserFaceHeaders) SetCommonHeaders(v map[string]*string) *QueryUserFaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserFaceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserFaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserFaceRequest struct {
	FaceId *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Sn     *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s QueryUserFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFaceRequest) GoString() string {
	return s.String()
}

func (s *QueryUserFaceRequest) SetFaceId(v string) *QueryUserFaceRequest {
	s.FaceId = &v
	return s
}

func (s *QueryUserFaceRequest) SetSn(v string) *QueryUserFaceRequest {
	s.Sn = &v
	return s
}

type QueryUserFaceResponseBody struct {
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	FaceId *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserFaceResponseBody) SetCorpId(v string) *QueryUserFaceResponseBody {
	s.CorpId = &v
	return s
}

func (s *QueryUserFaceResponseBody) SetFaceId(v string) *QueryUserFaceResponseBody {
	s.FaceId = &v
	return s
}

func (s *QueryUserFaceResponseBody) SetName(v string) *QueryUserFaceResponseBody {
	s.Name = &v
	return s
}

func (s *QueryUserFaceResponseBody) SetUserId(v string) *QueryUserFaceResponseBody {
	s.UserId = &v
	return s
}

type QueryUserFaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFaceResponse) GoString() string {
	return s.String()
}

func (s *QueryUserFaceResponse) SetHeaders(v map[string]*string) *QueryUserFaceResponse {
	s.Headers = v
	return s
}

func (s *QueryUserFaceResponse) SetStatusCode(v int32) *QueryUserFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserFaceResponse) SetBody(v *QueryUserFaceResponseBody) *QueryUserFaceResponse {
	s.Body = v
	return s
}

type QueryUserPayInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserPayInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserPayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserPayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserPayInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserPayInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserPayInfoRequest struct {
	FaceId *string `json:"faceId,omitempty" xml:"faceId,omitempty"`
	Sn     *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserPayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUserPayInfoRequest) SetFaceId(v string) *QueryUserPayInfoRequest {
	s.FaceId = &v
	return s
}

func (s *QueryUserPayInfoRequest) SetSn(v string) *QueryUserPayInfoRequest {
	s.Sn = &v
	return s
}

func (s *QueryUserPayInfoRequest) SetUserId(v string) *QueryUserPayInfoRequest {
	s.UserId = &v
	return s
}

type QueryUserPayInfoResponseBody struct {
	SignNo *string `json:"signNo,omitempty" xml:"signNo,omitempty"`
}

func (s QueryUserPayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserPayInfoResponseBody) SetSignNo(v string) *QueryUserPayInfoResponseBody {
	s.SignNo = &v
	return s
}

type QueryUserPayInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserPayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserPayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserPayInfoResponse) SetHeaders(v map[string]*string) *QueryUserPayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserPayInfoResponse) SetStatusCode(v int32) *QueryUserPayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserPayInfoResponse) SetBody(v *QueryUserPayInfoResponseBody) *QueryUserPayInfoResponse {
	s.Body = v
	return s
}

type RemoveDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceHeaders) GoString() string {
	return s.String()
}

func (s *RemoveDeviceHeaders) SetCommonHeaders(v map[string]*string) *RemoveDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveDeviceRequest struct {
	MerchantId *string `json:"merchantId,omitempty" xml:"merchantId,omitempty"`
	Sn         *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s RemoveDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceRequest) GoString() string {
	return s.String()
}

func (s *RemoveDeviceRequest) SetMerchantId(v string) *RemoveDeviceRequest {
	s.MerchantId = &v
	return s
}

func (s *RemoveDeviceRequest) SetSn(v string) *RemoveDeviceRequest {
	s.Sn = &v
	return s
}

type RemoveDeviceResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDeviceResponseBody) SetSuccess(v string) *RemoveDeviceResponseBody {
	s.Success = &v
	return s
}

type RemoveDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceResponse) GoString() string {
	return s.String()
}

func (s *RemoveDeviceResponse) SetHeaders(v map[string]*string) *RemoveDeviceResponse {
	s.Headers = v
	return s
}

func (s *RemoveDeviceResponse) SetStatusCode(v int32) *RemoveDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDeviceResponse) SetBody(v *RemoveDeviceResponseBody) *RemoveDeviceResponse {
	s.Body = v
	return s
}

type ReportDeviceLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReportDeviceLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceLogHeaders) GoString() string {
	return s.String()
}

func (s *ReportDeviceLogHeaders) SetCommonHeaders(v map[string]*string) *ReportDeviceLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReportDeviceLogHeaders) SetXAcsDingtalkAccessToken(v string) *ReportDeviceLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReportDeviceLogRequest struct {
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Sn      *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ReportDeviceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceLogRequest) GoString() string {
	return s.String()
}

func (s *ReportDeviceLogRequest) SetMediaId(v string) *ReportDeviceLogRequest {
	s.MediaId = &v
	return s
}

func (s *ReportDeviceLogRequest) SetSn(v string) *ReportDeviceLogRequest {
	s.Sn = &v
	return s
}

func (s *ReportDeviceLogRequest) SetType(v string) *ReportDeviceLogRequest {
	s.Type = &v
	return s
}

type ReportDeviceLogResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReportDeviceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceLogResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDeviceLogResponseBody) SetSuccess(v bool) *ReportDeviceLogResponseBody {
	s.Success = &v
	return s
}

type ReportDeviceLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportDeviceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportDeviceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceLogResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceLogResponse) SetHeaders(v map[string]*string) *ReportDeviceLogResponse {
	s.Headers = v
	return s
}

func (s *ReportDeviceLogResponse) SetStatusCode(v int32) *ReportDeviceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportDeviceLogResponse) SetBody(v *ReportDeviceLogResponseBody) *ReportDeviceLogResponse {
	s.Body = v
	return s
}

type ReportDeviceUseLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReportDeviceUseLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceUseLogHeaders) GoString() string {
	return s.String()
}

func (s *ReportDeviceUseLogHeaders) SetCommonHeaders(v map[string]*string) *ReportDeviceUseLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReportDeviceUseLogHeaders) SetXAcsDingtalkAccessToken(v string) *ReportDeviceUseLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReportDeviceUseLogRequest struct {
	Action  *string `json:"action,omitempty" xml:"action,omitempty"`
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Sn      *string `json:"sn,omitempty" xml:"sn,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ReportDeviceUseLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceUseLogRequest) GoString() string {
	return s.String()
}

func (s *ReportDeviceUseLogRequest) SetAction(v string) *ReportDeviceUseLogRequest {
	s.Action = &v
	return s
}

func (s *ReportDeviceUseLogRequest) SetOrderNo(v string) *ReportDeviceUseLogRequest {
	s.OrderNo = &v
	return s
}

func (s *ReportDeviceUseLogRequest) SetSn(v string) *ReportDeviceUseLogRequest {
	s.Sn = &v
	return s
}

func (s *ReportDeviceUseLogRequest) SetUserId(v string) *ReportDeviceUseLogRequest {
	s.UserId = &v
	return s
}

type ReportDeviceUseLogResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReportDeviceUseLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceUseLogResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDeviceUseLogResponseBody) SetSuccess(v bool) *ReportDeviceUseLogResponseBody {
	s.Success = &v
	return s
}

type ReportDeviceUseLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportDeviceUseLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportDeviceUseLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDeviceUseLogResponse) GoString() string {
	return s.String()
}

func (s *ReportDeviceUseLogResponse) SetHeaders(v map[string]*string) *ReportDeviceUseLogResponse {
	s.Headers = v
	return s
}

func (s *ReportDeviceUseLogResponse) SetStatusCode(v int32) *ReportDeviceUseLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportDeviceUseLogResponse) SetBody(v *ReportDeviceUseLogResponseBody) *ReportDeviceUseLogResponse {
	s.Body = v
	return s
}

type RollbackDeductPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RollbackDeductPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s RollbackDeductPointHeaders) GoString() string {
	return s.String()
}

func (s *RollbackDeductPointHeaders) SetCommonHeaders(v map[string]*string) *RollbackDeductPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RollbackDeductPointHeaders) SetXAcsDingtalkAccessToken(v string) *RollbackDeductPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RollbackDeductPointRequest struct {
	BizId     *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	PointType *string `json:"pointType,omitempty" xml:"pointType,omitempty"`
}

func (s RollbackDeductPointRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackDeductPointRequest) GoString() string {
	return s.String()
}

func (s *RollbackDeductPointRequest) SetBizId(v string) *RollbackDeductPointRequest {
	s.BizId = &v
	return s
}

func (s *RollbackDeductPointRequest) SetPointType(v string) *RollbackDeductPointRequest {
	s.PointType = &v
	return s
}

type RollbackDeductPointResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RollbackDeductPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackDeductPointResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackDeductPointResponseBody) SetResult(v bool) *RollbackDeductPointResponseBody {
	s.Result = &v
	return s
}

func (s *RollbackDeductPointResponseBody) SetSuccess(v bool) *RollbackDeductPointResponseBody {
	s.Success = &v
	return s
}

type RollbackDeductPointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackDeductPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackDeductPointResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackDeductPointResponse) GoString() string {
	return s.String()
}

func (s *RollbackDeductPointResponse) SetHeaders(v map[string]*string) *RollbackDeductPointResponse {
	s.Headers = v
	return s
}

func (s *RollbackDeductPointResponse) SetStatusCode(v int32) *RollbackDeductPointResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackDeductPointResponse) SetBody(v *RollbackDeductPointResponseBody) *RollbackDeductPointResponse {
	s.Body = v
	return s
}

type SaveClassLearningDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveClassLearningDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveClassLearningDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveClassLearningDataHeaders) SetCommonHeaders(v map[string]*string) *SaveClassLearningDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveClassLearningDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveClassLearningDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveClassLearningDataRequest struct {
	AssignNum                *int32    `json:"assignNum,omitempty" xml:"assignNum,omitempty"`
	AssignStudentUserIds     []*string `json:"assignStudentUserIds,omitempty" xml:"assignStudentUserIds,omitempty" type:"Repeated"`
	BizId                    *string   `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType                  *string   `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CorpId                   *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId                   *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	FileSuffix               *string   `json:"fileSuffix,omitempty" xml:"fileSuffix,omitempty"`
	GeneratedTime            *int64    `json:"generatedTime,omitempty" xml:"generatedTime,omitempty"`
	QuestionNum              *int32    `json:"questionNum,omitempty" xml:"questionNum,omitempty"`
	QuestionPictureNum       *int32    `json:"questionPictureNum,omitempty" xml:"questionPictureNum,omitempty"`
	StandardAnswerPictureNum *int32    `json:"standardAnswerPictureNum,omitempty" xml:"standardAnswerPictureNum,omitempty"`
	SubjectCode              *string   `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	TeacherUserId            *string   `json:"teacherUserId,omitempty" xml:"teacherUserId,omitempty"`
}

func (s SaveClassLearningDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveClassLearningDataRequest) GoString() string {
	return s.String()
}

func (s *SaveClassLearningDataRequest) SetAssignNum(v int32) *SaveClassLearningDataRequest {
	s.AssignNum = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetAssignStudentUserIds(v []*string) *SaveClassLearningDataRequest {
	s.AssignStudentUserIds = v
	return s
}

func (s *SaveClassLearningDataRequest) SetBizId(v string) *SaveClassLearningDataRequest {
	s.BizId = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetBizType(v string) *SaveClassLearningDataRequest {
	s.BizType = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetCorpId(v string) *SaveClassLearningDataRequest {
	s.CorpId = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetDeptId(v int64) *SaveClassLearningDataRequest {
	s.DeptId = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetFileSuffix(v string) *SaveClassLearningDataRequest {
	s.FileSuffix = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetGeneratedTime(v int64) *SaveClassLearningDataRequest {
	s.GeneratedTime = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetQuestionNum(v int32) *SaveClassLearningDataRequest {
	s.QuestionNum = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetQuestionPictureNum(v int32) *SaveClassLearningDataRequest {
	s.QuestionPictureNum = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetStandardAnswerPictureNum(v int32) *SaveClassLearningDataRequest {
	s.StandardAnswerPictureNum = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetSubjectCode(v string) *SaveClassLearningDataRequest {
	s.SubjectCode = &v
	return s
}

func (s *SaveClassLearningDataRequest) SetTeacherUserId(v string) *SaveClassLearningDataRequest {
	s.TeacherUserId = &v
	return s
}

type SaveClassLearningDataResponseBody struct {
	Result  *SaveClassLearningDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveClassLearningDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveClassLearningDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveClassLearningDataResponseBody) SetResult(v *SaveClassLearningDataResponseBodyResult) *SaveClassLearningDataResponseBody {
	s.Result = v
	return s
}

func (s *SaveClassLearningDataResponseBody) SetSuccess(v bool) *SaveClassLearningDataResponseBody {
	s.Success = &v
	return s
}

type SaveClassLearningDataResponseBodyResult struct {
	QuestionUploadUrlList       []*string `json:"questionUploadUrlList,omitempty" xml:"questionUploadUrlList,omitempty" type:"Repeated"`
	StandardAnswerUploadUrlList []*string `json:"standardAnswerUploadUrlList,omitempty" xml:"standardAnswerUploadUrlList,omitempty" type:"Repeated"`
}

func (s SaveClassLearningDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveClassLearningDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveClassLearningDataResponseBodyResult) SetQuestionUploadUrlList(v []*string) *SaveClassLearningDataResponseBodyResult {
	s.QuestionUploadUrlList = v
	return s
}

func (s *SaveClassLearningDataResponseBodyResult) SetStandardAnswerUploadUrlList(v []*string) *SaveClassLearningDataResponseBodyResult {
	s.StandardAnswerUploadUrlList = v
	return s
}

type SaveClassLearningDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveClassLearningDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveClassLearningDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveClassLearningDataResponse) GoString() string {
	return s.String()
}

func (s *SaveClassLearningDataResponse) SetHeaders(v map[string]*string) *SaveClassLearningDataResponse {
	s.Headers = v
	return s
}

func (s *SaveClassLearningDataResponse) SetStatusCode(v int32) *SaveClassLearningDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveClassLearningDataResponse) SetBody(v *SaveClassLearningDataResponseBody) *SaveClassLearningDataResponse {
	s.Body = v
	return s
}

type SaveStudentLearningDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveStudentLearningDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataHeaders) SetCommonHeaders(v map[string]*string) *SaveStudentLearningDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveStudentLearningDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveStudentLearningDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveStudentLearningDataRequest struct {
	AssignNum      *int32                                          `json:"assignNum,omitempty" xml:"assignNum,omitempty"`
	BizId          *string                                         `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType        *string                                         `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CorpId         *string                                         `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorrectNum     *int32                                          `json:"correctNum,omitempty" xml:"correctNum,omitempty"`
	DeptId         *int64                                          `json:"deptId,omitempty" xml:"deptId,omitempty"`
	FileSuffix     *string                                         `json:"fileSuffix,omitempty" xml:"fileSuffix,omitempty"`
	GeneratedTime  *int64                                          `json:"generatedTime,omitempty" xml:"generatedTime,omitempty"`
	QuestionNum    *int32                                          `json:"questionNum,omitempty" xml:"questionNum,omitempty"`
	StudentUserId  *string                                         `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
	SubjectCode    *string                                         `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
	SubmitNum      *int32                                          `json:"submitNum,omitempty" xml:"submitNum,omitempty"`
	WrongQuestions []*SaveStudentLearningDataRequestWrongQuestions `json:"wrongQuestions,omitempty" xml:"wrongQuestions,omitempty" type:"Repeated"`
}

func (s SaveStudentLearningDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataRequest) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataRequest) SetAssignNum(v int32) *SaveStudentLearningDataRequest {
	s.AssignNum = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetBizId(v string) *SaveStudentLearningDataRequest {
	s.BizId = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetBizType(v string) *SaveStudentLearningDataRequest {
	s.BizType = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetCorpId(v string) *SaveStudentLearningDataRequest {
	s.CorpId = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetCorrectNum(v int32) *SaveStudentLearningDataRequest {
	s.CorrectNum = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetDeptId(v int64) *SaveStudentLearningDataRequest {
	s.DeptId = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetFileSuffix(v string) *SaveStudentLearningDataRequest {
	s.FileSuffix = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetGeneratedTime(v int64) *SaveStudentLearningDataRequest {
	s.GeneratedTime = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetQuestionNum(v int32) *SaveStudentLearningDataRequest {
	s.QuestionNum = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetStudentUserId(v string) *SaveStudentLearningDataRequest {
	s.StudentUserId = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetSubjectCode(v string) *SaveStudentLearningDataRequest {
	s.SubjectCode = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetSubmitNum(v int32) *SaveStudentLearningDataRequest {
	s.SubmitNum = &v
	return s
}

func (s *SaveStudentLearningDataRequest) SetWrongQuestions(v []*SaveStudentLearningDataRequestWrongQuestions) *SaveStudentLearningDataRequest {
	s.WrongQuestions = v
	return s
}

type SaveStudentLearningDataRequestWrongQuestions struct {
	KnowledgePoints          []*string `json:"knowledgePoints,omitempty" xml:"knowledgePoints,omitempty" type:"Repeated"`
	QuestionNo               *string   `json:"questionNo,omitempty" xml:"questionNo,omitempty"`
	QuestionPictureNum       *int32    `json:"questionPictureNum,omitempty" xml:"questionPictureNum,omitempty"`
	StandardAnswerPictureNum *int32    `json:"standardAnswerPictureNum,omitempty" xml:"standardAnswerPictureNum,omitempty"`
	UserAnswerPictureNum     *int32    `json:"userAnswerPictureNum,omitempty" xml:"userAnswerPictureNum,omitempty"`
}

func (s SaveStudentLearningDataRequestWrongQuestions) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataRequestWrongQuestions) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataRequestWrongQuestions) SetKnowledgePoints(v []*string) *SaveStudentLearningDataRequestWrongQuestions {
	s.KnowledgePoints = v
	return s
}

func (s *SaveStudentLearningDataRequestWrongQuestions) SetQuestionNo(v string) *SaveStudentLearningDataRequestWrongQuestions {
	s.QuestionNo = &v
	return s
}

func (s *SaveStudentLearningDataRequestWrongQuestions) SetQuestionPictureNum(v int32) *SaveStudentLearningDataRequestWrongQuestions {
	s.QuestionPictureNum = &v
	return s
}

func (s *SaveStudentLearningDataRequestWrongQuestions) SetStandardAnswerPictureNum(v int32) *SaveStudentLearningDataRequestWrongQuestions {
	s.StandardAnswerPictureNum = &v
	return s
}

func (s *SaveStudentLearningDataRequestWrongQuestions) SetUserAnswerPictureNum(v int32) *SaveStudentLearningDataRequestWrongQuestions {
	s.UserAnswerPictureNum = &v
	return s
}

type SaveStudentLearningDataResponseBody struct {
	Result  *SaveStudentLearningDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveStudentLearningDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataResponseBody) SetResult(v *SaveStudentLearningDataResponseBodyResult) *SaveStudentLearningDataResponseBody {
	s.Result = v
	return s
}

func (s *SaveStudentLearningDataResponseBody) SetSuccess(v bool) *SaveStudentLearningDataResponseBody {
	s.Success = &v
	return s
}

type SaveStudentLearningDataResponseBodyResult struct {
	SaveSuccess    *bool                                                      `json:"saveSuccess,omitempty" xml:"saveSuccess,omitempty"`
	WrongQuestions []*SaveStudentLearningDataResponseBodyResultWrongQuestions `json:"wrongQuestions,omitempty" xml:"wrongQuestions,omitempty" type:"Repeated"`
}

func (s SaveStudentLearningDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataResponseBodyResult) SetSaveSuccess(v bool) *SaveStudentLearningDataResponseBodyResult {
	s.SaveSuccess = &v
	return s
}

func (s *SaveStudentLearningDataResponseBodyResult) SetWrongQuestions(v []*SaveStudentLearningDataResponseBodyResultWrongQuestions) *SaveStudentLearningDataResponseBodyResult {
	s.WrongQuestions = v
	return s
}

type SaveStudentLearningDataResponseBodyResultWrongQuestions struct {
	QuestionNo                  *string   `json:"questionNo,omitempty" xml:"questionNo,omitempty"`
	QuestionUploadUrlList       []*string `json:"questionUploadUrlList,omitempty" xml:"questionUploadUrlList,omitempty" type:"Repeated"`
	StandardAnswerUploadUrlList []*string `json:"standardAnswerUploadUrlList,omitempty" xml:"standardAnswerUploadUrlList,omitempty" type:"Repeated"`
	UserAnswerUploadUrlList     []*string `json:"userAnswerUploadUrlList,omitempty" xml:"userAnswerUploadUrlList,omitempty" type:"Repeated"`
}

func (s SaveStudentLearningDataResponseBodyResultWrongQuestions) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataResponseBodyResultWrongQuestions) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataResponseBodyResultWrongQuestions) SetQuestionNo(v string) *SaveStudentLearningDataResponseBodyResultWrongQuestions {
	s.QuestionNo = &v
	return s
}

func (s *SaveStudentLearningDataResponseBodyResultWrongQuestions) SetQuestionUploadUrlList(v []*string) *SaveStudentLearningDataResponseBodyResultWrongQuestions {
	s.QuestionUploadUrlList = v
	return s
}

func (s *SaveStudentLearningDataResponseBodyResultWrongQuestions) SetStandardAnswerUploadUrlList(v []*string) *SaveStudentLearningDataResponseBodyResultWrongQuestions {
	s.StandardAnswerUploadUrlList = v
	return s
}

func (s *SaveStudentLearningDataResponseBodyResultWrongQuestions) SetUserAnswerUploadUrlList(v []*string) *SaveStudentLearningDataResponseBodyResultWrongQuestions {
	s.UserAnswerUploadUrlList = v
	return s
}

type SaveStudentLearningDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveStudentLearningDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveStudentLearningDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveStudentLearningDataResponse) GoString() string {
	return s.String()
}

func (s *SaveStudentLearningDataResponse) SetHeaders(v map[string]*string) *SaveStudentLearningDataResponse {
	s.Headers = v
	return s
}

func (s *SaveStudentLearningDataResponse) SetStatusCode(v int32) *SaveStudentLearningDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveStudentLearningDataResponse) SetBody(v *SaveStudentLearningDataResponseBody) *SaveStudentLearningDataResponse {
	s.Body = v
	return s
}

type SearchTeachersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchTeachersHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersHeaders) GoString() string {
	return s.String()
}

func (s *SearchTeachersHeaders) SetCommonHeaders(v map[string]*string) *SearchTeachersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchTeachersHeaders) SetXAcsDingtalkAccessToken(v string) *SearchTeachersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchTeachersRequest struct {
	NameKeyword *string `json:"nameKeyword,omitempty" xml:"nameKeyword,omitempty"`
}

func (s SearchTeachersRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersRequest) GoString() string {
	return s.String()
}

func (s *SearchTeachersRequest) SetNameKeyword(v string) *SearchTeachersRequest {
	s.NameKeyword = &v
	return s
}

type SearchTeachersResponseBody struct {
	Users []*SearchTeachersResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s SearchTeachersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponseBody) SetUsers(v []*SearchTeachersResponseBodyUsers) *SearchTeachersResponseBody {
	s.Users = v
	return s
}

type SearchTeachersResponseBodyUsers struct {
	ClassId  *int64  `json:"classId,omitempty" xml:"classId,omitempty"`
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchTeachersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponseBodyUsers) SetClassId(v int64) *SearchTeachersResponseBodyUsers {
	s.ClassId = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetDeptName(v string) *SearchTeachersResponseBodyUsers {
	s.DeptName = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetName(v string) *SearchTeachersResponseBodyUsers {
	s.Name = &v
	return s
}

func (s *SearchTeachersResponseBodyUsers) SetUserId(v string) *SearchTeachersResponseBodyUsers {
	s.UserId = &v
	return s
}

type SearchTeachersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTeachersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTeachersResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTeachersResponse) GoString() string {
	return s.String()
}

func (s *SearchTeachersResponse) SetHeaders(v map[string]*string) *SearchTeachersResponse {
	s.Headers = v
	return s
}

func (s *SearchTeachersResponse) SetStatusCode(v int32) *SearchTeachersResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTeachersResponse) SetBody(v *SearchTeachersResponseBody) *SearchTeachersResponse {
	s.Body = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	BizId        *string   `json:"bizId,omitempty" xml:"bizId,omitempty"`
	FromUserId   *string   `json:"fromUserId,omitempty" xml:"fromUserId,omitempty"`
	Sn           *string   `json:"sn,omitempty" xml:"sn,omitempty"`
	ToUserIdList []*string `json:"toUserIdList,omitempty" xml:"toUserIdList,omitempty" type:"Repeated"`
	Type         *int64    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetBizId(v string) *SendMessageRequest {
	s.BizId = &v
	return s
}

func (s *SendMessageRequest) SetFromUserId(v string) *SendMessageRequest {
	s.FromUserId = &v
	return s
}

func (s *SendMessageRequest) SetSn(v string) *SendMessageRequest {
	s.Sn = &v
	return s
}

func (s *SendMessageRequest) SetToUserIdList(v []*string) *SendMessageRequest {
	s.ToUserIdList = v
	return s
}

func (s *SendMessageRequest) SetType(v int64) *SendMessageRequest {
	s.Type = &v
	return s
}

type SendMessageResponseBody struct {
	SuccessInfo *string `json:"successInfo,omitempty" xml:"successInfo,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetSuccessInfo(v string) *SendMessageResponseBody {
	s.SuccessInfo = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type StartCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCourseHeaders) GoString() string {
	return s.String()
}

func (s *StartCourseHeaders) SetCommonHeaders(v map[string]*string) *StartCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCourseHeaders) SetXAcsDingtalkAccessToken(v string) *StartCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCourseRequest struct {
	CourseCode       *string                               `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Ext              *string                               `json:"ext,omitempty" xml:"ext,omitempty"`
	IsvCode          *string                               `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	LivePlayInfoList []*StartCourseRequestLivePlayInfoList `json:"livePlayInfoList,omitempty" xml:"livePlayInfoList,omitempty" type:"Repeated"`
	OpUserId         *string                               `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s StartCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCourseRequest) GoString() string {
	return s.String()
}

func (s *StartCourseRequest) SetCourseCode(v string) *StartCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *StartCourseRequest) SetExt(v string) *StartCourseRequest {
	s.Ext = &v
	return s
}

func (s *StartCourseRequest) SetIsvCode(v string) *StartCourseRequest {
	s.IsvCode = &v
	return s
}

func (s *StartCourseRequest) SetLivePlayInfoList(v []*StartCourseRequestLivePlayInfoList) *StartCourseRequest {
	s.LivePlayInfoList = v
	return s
}

func (s *StartCourseRequest) SetOpUserId(v string) *StartCourseRequest {
	s.OpUserId = &v
	return s
}

type StartCourseRequestLivePlayInfoList struct {
	LiveInputUrl     *string `json:"liveInputUrl,omitempty" xml:"liveInputUrl,omitempty"`
	LiveOutputFlvUrl *string `json:"liveOutputFlvUrl,omitempty" xml:"liveOutputFlvUrl,omitempty"`
	LiveOutputHlsUrl *string `json:"liveOutputHlsUrl,omitempty" xml:"liveOutputHlsUrl,omitempty"`
	LiveType         *int32  `json:"liveType,omitempty" xml:"liveType,omitempty"`
	ReplayUrl        *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s StartCourseRequestLivePlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s StartCourseRequestLivePlayInfoList) GoString() string {
	return s.String()
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveInputUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.LiveInputUrl = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveOutputFlvUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.LiveOutputFlvUrl = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveOutputHlsUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.LiveOutputHlsUrl = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetLiveType(v int32) *StartCourseRequestLivePlayInfoList {
	s.LiveType = &v
	return s
}

func (s *StartCourseRequestLivePlayInfoList) SetReplayUrl(v string) *StartCourseRequestLivePlayInfoList {
	s.ReplayUrl = &v
	return s
}

type StartCourseResponseBody struct {
	UniversityCourseCommonResponse *StartCourseResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s StartCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponseBody) GoString() string {
	return s.String()
}

func (s *StartCourseResponseBody) SetUniversityCourseCommonResponse(v *StartCourseResponseBodyUniversityCourseCommonResponse) *StartCourseResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type StartCourseResponseBodyUniversityCourseCommonResponse struct {
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartCourseResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *StartCourseResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *StartCourseResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

func (s *StartCourseResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *StartCourseResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

type StartCourseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCourseResponse) GoString() string {
	return s.String()
}

func (s *StartCourseResponse) SetHeaders(v map[string]*string) *StartCourseResponse {
	s.Headers = v
	return s
}

func (s *StartCourseResponse) SetStatusCode(v int32) *StartCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCourseResponse) SetBody(v *StartCourseResponseBody) *StartCourseResponse {
	s.Body = v
	return s
}

type StartCoursePrepareHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCoursePrepareHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareHeaders) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareHeaders) SetCommonHeaders(v map[string]*string) *StartCoursePrepareHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCoursePrepareHeaders) SetXAcsDingtalkAccessToken(v string) *StartCoursePrepareHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCoursePrepareRequest struct {
	CourseDate      *string  `json:"courseDate,omitempty" xml:"courseDate,omitempty"`
	CourseGroupCode *string  `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	DeviceId        *string  `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	Ext             *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	IsvCode         *string  `json:"isvCode,omitempty" xml:"isvCode,omitempty"`
	LiveCoverImage  *string  `json:"liveCoverImage,omitempty" xml:"liveCoverImage,omitempty"`
	SectionIndex    []*int32 `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
	OpUserId        *string  `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s StartCoursePrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareRequest) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareRequest) SetCourseDate(v string) *StartCoursePrepareRequest {
	s.CourseDate = &v
	return s
}

func (s *StartCoursePrepareRequest) SetCourseGroupCode(v string) *StartCoursePrepareRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *StartCoursePrepareRequest) SetDeviceId(v string) *StartCoursePrepareRequest {
	s.DeviceId = &v
	return s
}

func (s *StartCoursePrepareRequest) SetExt(v string) *StartCoursePrepareRequest {
	s.Ext = &v
	return s
}

func (s *StartCoursePrepareRequest) SetIsvCode(v string) *StartCoursePrepareRequest {
	s.IsvCode = &v
	return s
}

func (s *StartCoursePrepareRequest) SetLiveCoverImage(v string) *StartCoursePrepareRequest {
	s.LiveCoverImage = &v
	return s
}

func (s *StartCoursePrepareRequest) SetSectionIndex(v []*int32) *StartCoursePrepareRequest {
	s.SectionIndex = v
	return s
}

func (s *StartCoursePrepareRequest) SetOpUserId(v string) *StartCoursePrepareRequest {
	s.OpUserId = &v
	return s
}

type StartCoursePrepareResponseBody struct {
	UniversityCourseCommonResponse *StartCoursePrepareResponseBodyUniversityCourseCommonResponse `json:"universityCourseCommonResponse,omitempty" xml:"universityCourseCommonResponse,omitempty" type:"Struct"`
}

func (s StartCoursePrepareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponseBody) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponseBody) SetUniversityCourseCommonResponse(v *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) *StartCoursePrepareResponseBody {
	s.UniversityCourseCommonResponse = v
	return s
}

type StartCoursePrepareResponseBodyUniversityCourseCommonResponse struct {
	CourseCode *string `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartCoursePrepareResponseBodyUniversityCourseCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponseBodyUniversityCourseCommonResponse) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) SetCourseCode(v string) *StartCoursePrepareResponseBodyUniversityCourseCommonResponse {
	s.CourseCode = &v
	return s
}

func (s *StartCoursePrepareResponseBodyUniversityCourseCommonResponse) SetSuccess(v bool) *StartCoursePrepareResponseBodyUniversityCourseCommonResponse {
	s.Success = &v
	return s
}

type StartCoursePrepareResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartCoursePrepareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCoursePrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCoursePrepareResponse) GoString() string {
	return s.String()
}

func (s *StartCoursePrepareResponse) SetHeaders(v map[string]*string) *StartCoursePrepareResponse {
	s.Headers = v
	return s
}

func (s *StartCoursePrepareResponse) SetStatusCode(v int32) *StartCoursePrepareResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCoursePrepareResponse) SetBody(v *StartCoursePrepareResponseBody) *StartCoursePrepareResponse {
	s.Body = v
	return s
}

type SubscribeUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *SubscribeUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeUniversityCourseGroupRequest struct {
	CourseGroupCode *string   `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	StudentUserIds  []*string `json:"studentUserIds,omitempty" xml:"studentUserIds,omitempty" type:"Repeated"`
	OpUserId        *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s SubscribeUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupRequest) SetCourseGroupCode(v string) *SubscribeUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *SubscribeUniversityCourseGroupRequest) SetStudentUserIds(v []*string) *SubscribeUniversityCourseGroupRequest {
	s.StudentUserIds = v
	return s
}

func (s *SubscribeUniversityCourseGroupRequest) SetOpUserId(v string) *SubscribeUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type SubscribeUniversityCourseGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SubscribeUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupResponseBody) SetResult(v bool) *SubscribeUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type SubscribeUniversityCourseGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubscribeUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *SubscribeUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *SubscribeUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *SubscribeUniversityCourseGroupResponse) SetStatusCode(v int32) *SubscribeUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeUniversityCourseGroupResponse) SetBody(v *SubscribeUniversityCourseGroupResponseBody) *SubscribeUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type UnsubscribeUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UnsubscribeUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnsubscribeUniversityCourseGroupRequest struct {
	CourseGroupCode *string   `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	StudentUserIds  []*string `json:"studentUserIds,omitempty" xml:"studentUserIds,omitempty" type:"Repeated"`
	OpUserId        *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetCourseGroupCode(v string) *UnsubscribeUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetStudentUserIds(v []*string) *UnsubscribeUniversityCourseGroupRequest {
	s.StudentUserIds = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupRequest) SetOpUserId(v string) *UnsubscribeUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type UnsubscribeUniversityCourseGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnsubscribeUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupResponseBody) SetResult(v bool) *UnsubscribeUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type UnsubscribeUniversityCourseGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnsubscribeUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *UnsubscribeUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeUniversityCourseGroupResponse) SetStatusCode(v int32) *UnsubscribeUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeUniversityCourseGroupResponse) SetBody(v *UnsubscribeUniversityCourseGroupResponseBody) *UnsubscribeUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type UpdateCoursesOfClassHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCoursesOfClassHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassHeaders) SetCommonHeaders(v map[string]*string) *UpdateCoursesOfClassHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCoursesOfClassHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCoursesOfClassHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCoursesOfClassRequest struct {
	Courses       []*UpdateCoursesOfClassRequestCourses     `json:"courses,omitempty" xml:"courses,omitempty" type:"Repeated"`
	SectionConfig *UpdateCoursesOfClassRequestSectionConfig `json:"sectionConfig,omitempty" xml:"sectionConfig,omitempty" type:"Struct"`
	OpUserId      *string                                   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateCoursesOfClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequest) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequest) SetCourses(v []*UpdateCoursesOfClassRequestCourses) *UpdateCoursesOfClassRequest {
	s.Courses = v
	return s
}

func (s *UpdateCoursesOfClassRequest) SetSectionConfig(v *UpdateCoursesOfClassRequestSectionConfig) *UpdateCoursesOfClassRequest {
	s.SectionConfig = v
	return s
}

func (s *UpdateCoursesOfClassRequest) SetOpUserId(v string) *UpdateCoursesOfClassRequest {
	s.OpUserId = &v
	return s
}

type UpdateCoursesOfClassRequestCourses struct {
	CourseCode      *string                                         `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	CourseGroupCode *string                                         `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	CourseName      *string                                         `json:"courseName,omitempty" xml:"courseName,omitempty"`
	CreatorName     *string                                         `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	DateModel       *UpdateCoursesOfClassRequestCoursesDateModel    `json:"dateModel,omitempty" xml:"dateModel,omitempty" type:"Struct"`
	DeleteTag       *bool                                           `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	Location        *string                                         `json:"location,omitempty" xml:"location,omitempty"`
	SectionModel    *UpdateCoursesOfClassRequestCoursesSectionModel `json:"sectionModel,omitempty" xml:"sectionModel,omitempty" type:"Struct"`
	TeacherStaffIds []*string                                       `json:"teacherStaffIds,omitempty" xml:"teacherStaffIds,omitempty" type:"Repeated"`
}

func (s UpdateCoursesOfClassRequestCourses) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCourses) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseCode(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseCode = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseGroupCode(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseGroupCode = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCourseName(v string) *UpdateCoursesOfClassRequestCourses {
	s.CourseName = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetCreatorName(v string) *UpdateCoursesOfClassRequestCourses {
	s.CreatorName = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetDateModel(v *UpdateCoursesOfClassRequestCoursesDateModel) *UpdateCoursesOfClassRequestCourses {
	s.DateModel = v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetDeleteTag(v bool) *UpdateCoursesOfClassRequestCourses {
	s.DeleteTag = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetLocation(v string) *UpdateCoursesOfClassRequestCourses {
	s.Location = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetSectionModel(v *UpdateCoursesOfClassRequestCoursesSectionModel) *UpdateCoursesOfClassRequestCourses {
	s.SectionModel = v
	return s
}

func (s *UpdateCoursesOfClassRequestCourses) SetTeacherStaffIds(v []*string) *UpdateCoursesOfClassRequestCourses {
	s.TeacherStaffIds = v
	return s
}

type UpdateCoursesOfClassRequestCoursesDateModel struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s UpdateCoursesOfClassRequestCoursesDateModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCoursesDateModel) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetDayOfMonth(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.DayOfMonth = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetMonth(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.Month = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesDateModel) SetYear(v int32) *UpdateCoursesOfClassRequestCoursesDateModel {
	s.Year = &v
	return s
}

type UpdateCoursesOfClassRequestCoursesSectionModel struct {
	SectionIndex *int32  `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionName  *string `json:"sectionName,omitempty" xml:"sectionName,omitempty"`
	SectionType  *string `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
}

func (s UpdateCoursesOfClassRequestCoursesSectionModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestCoursesSectionModel) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionIndex(v int32) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionIndex = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionName(v string) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionName = &v
	return s
}

func (s *UpdateCoursesOfClassRequestCoursesSectionModel) SetSectionType(v string) *UpdateCoursesOfClassRequestCoursesSectionModel {
	s.SectionType = &v
	return s
}

type UpdateCoursesOfClassRequestSectionConfig struct {
	SectionModels []*UpdateCoursesOfClassRequestSectionConfigSectionModels `json:"sectionModels,omitempty" xml:"sectionModels,omitempty" type:"Repeated"`
}

func (s UpdateCoursesOfClassRequestSectionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfig) SetSectionModels(v []*UpdateCoursesOfClassRequestSectionConfigSectionModels) *UpdateCoursesOfClassRequestSectionConfig {
	s.SectionModels = v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModels struct {
	End          *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd   `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	SectionIndex *int32                                                      `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty"`
	SectionType  *string                                                     `json:"sectionType,omitempty" xml:"sectionType,omitempty"`
	Start        *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModels) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModels) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetEnd(v *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.End = v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetSectionIndex(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionIndex = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetSectionType(v string) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.SectionType = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModels) SetStart(v *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) *UpdateCoursesOfClassRequestSectionConfigSectionModels {
	s.Start = v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) SetHour(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Hour = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd) SetMin(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsEnd {
	s.Min = &v
	return s
}

type UpdateCoursesOfClassRequestSectionConfigSectionModelsStart struct {
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Min  *int32 `json:"min,omitempty" xml:"min,omitempty"`
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) SetHour(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Hour = &v
	return s
}

func (s *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart) SetMin(v int32) *UpdateCoursesOfClassRequestSectionConfigSectionModelsStart {
	s.Min = &v
	return s
}

type UpdateCoursesOfClassResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateCoursesOfClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassResponseBody) SetResult(v bool) *UpdateCoursesOfClassResponseBody {
	s.Result = &v
	return s
}

type UpdateCoursesOfClassResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCoursesOfClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCoursesOfClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCoursesOfClassResponse) GoString() string {
	return s.String()
}

func (s *UpdateCoursesOfClassResponse) SetHeaders(v map[string]*string) *UpdateCoursesOfClassResponse {
	s.Headers = v
	return s
}

func (s *UpdateCoursesOfClassResponse) SetStatusCode(v int32) *UpdateCoursesOfClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCoursesOfClassResponse) SetBody(v *UpdateCoursesOfClassResponseBody) *UpdateCoursesOfClassResponse {
	s.Body = v
	return s
}

type UpdatePhysicalClassroomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePhysicalClassroomHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomHeaders) SetCommonHeaders(v map[string]*string) *UpdatePhysicalClassroomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePhysicalClassroomHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePhysicalClassroomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePhysicalClassroomRequest struct {
	ClassroomBuilding *string `json:"classroomBuilding,omitempty" xml:"classroomBuilding,omitempty"`
	ClassroomCampus   *string `json:"classroomCampus,omitempty" xml:"classroomCampus,omitempty"`
	ClassroomFloor    *string `json:"classroomFloor,omitempty" xml:"classroomFloor,omitempty"`
	ClassroomId       *int64  `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	ClassroomName     *string `json:"classroomName,omitempty" xml:"classroomName,omitempty"`
	ClassroomNumber   *string `json:"classroomNumber,omitempty" xml:"classroomNumber,omitempty"`
	DirectBroadcast   *string `json:"directBroadcast,omitempty" xml:"directBroadcast,omitempty"`
	Ext               *string `json:"ext,omitempty" xml:"ext,omitempty"`
	OpUserId          *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdatePhysicalClassroomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomBuilding(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomBuilding = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomCampus(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomCampus = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomFloor(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomFloor = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomId(v int64) *UpdatePhysicalClassroomRequest {
	s.ClassroomId = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomName(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomName = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetClassroomNumber(v string) *UpdatePhysicalClassroomRequest {
	s.ClassroomNumber = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetDirectBroadcast(v string) *UpdatePhysicalClassroomRequest {
	s.DirectBroadcast = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetExt(v string) *UpdatePhysicalClassroomRequest {
	s.Ext = &v
	return s
}

func (s *UpdatePhysicalClassroomRequest) SetOpUserId(v string) *UpdatePhysicalClassroomRequest {
	s.OpUserId = &v
	return s
}

type UpdatePhysicalClassroomResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdatePhysicalClassroomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomResponseBody) SetResult(v bool) *UpdatePhysicalClassroomResponseBody {
	s.Result = &v
	return s
}

type UpdatePhysicalClassroomResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePhysicalClassroomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePhysicalClassroomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePhysicalClassroomResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhysicalClassroomResponse) SetHeaders(v map[string]*string) *UpdatePhysicalClassroomResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhysicalClassroomResponse) SetStatusCode(v int32) *UpdatePhysicalClassroomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhysicalClassroomResponse) SetBody(v *UpdatePhysicalClassroomResponseBody) *UpdatePhysicalClassroomResponse {
	s.Body = v
	return s
}

type UpdateRemoteClassCourseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRemoteClassCourseHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseHeaders) SetCommonHeaders(v map[string]*string) *UpdateRemoteClassCourseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRemoteClassCourseHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRemoteClassCourseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRemoteClassCourseRequest struct {
	AttendParticipants  []*UpdateRemoteClassCourseRequestAttendParticipants `json:"attendParticipants,omitempty" xml:"attendParticipants,omitempty" type:"Repeated"`
	AuthCode            *string                                             `json:"authCode,omitempty" xml:"authCode,omitempty"`
	CourseCode          *string                                             `json:"courseCode,omitempty" xml:"courseCode,omitempty"`
	CourseName          *string                                             `json:"courseName,omitempty" xml:"courseName,omitempty"`
	EndTime             *int64                                              `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime           *int64                                              `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TeachingParticipant *UpdateRemoteClassCourseRequestTeachingParticipant  `json:"teachingParticipant,omitempty" xml:"teachingParticipant,omitempty" type:"Struct"`
}

func (s UpdateRemoteClassCourseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequest) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequest) SetAttendParticipants(v []*UpdateRemoteClassCourseRequestAttendParticipants) *UpdateRemoteClassCourseRequest {
	s.AttendParticipants = v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetAuthCode(v string) *UpdateRemoteClassCourseRequest {
	s.AuthCode = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetCourseCode(v string) *UpdateRemoteClassCourseRequest {
	s.CourseCode = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetCourseName(v string) *UpdateRemoteClassCourseRequest {
	s.CourseName = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetEndTime(v int64) *UpdateRemoteClassCourseRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetStartTime(v int64) *UpdateRemoteClassCourseRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateRemoteClassCourseRequest) SetTeachingParticipant(v *UpdateRemoteClassCourseRequestTeachingParticipant) *UpdateRemoteClassCourseRequest {
	s.TeachingParticipant = v
	return s
}

type UpdateRemoteClassCourseRequestAttendParticipants struct {
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
}

func (s UpdateRemoteClassCourseRequestAttendParticipants) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequestAttendParticipants) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequestAttendParticipants) SetCorpId(v string) *UpdateRemoteClassCourseRequestAttendParticipants {
	s.CorpId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequestAttendParticipants) SetParticipantId(v string) *UpdateRemoteClassCourseRequestAttendParticipants {
	s.ParticipantId = &v
	return s
}

type UpdateRemoteClassCourseRequestTeachingParticipant struct {
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	ParticipantId *string `json:"participantId,omitempty" xml:"participantId,omitempty"`
}

func (s UpdateRemoteClassCourseRequestTeachingParticipant) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseRequestTeachingParticipant) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseRequestTeachingParticipant) SetCorpId(v string) *UpdateRemoteClassCourseRequestTeachingParticipant {
	s.CorpId = &v
	return s
}

func (s *UpdateRemoteClassCourseRequestTeachingParticipant) SetParticipantId(v string) *UpdateRemoteClassCourseRequestTeachingParticipant {
	s.ParticipantId = &v
	return s
}

type UpdateRemoteClassCourseResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRemoteClassCourseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseResponseBody) SetResult(v string) *UpdateRemoteClassCourseResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateRemoteClassCourseResponseBody) SetSuccess(v bool) *UpdateRemoteClassCourseResponseBody {
	s.Success = &v
	return s
}

type UpdateRemoteClassCourseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRemoteClassCourseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRemoteClassCourseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassCourseResponse) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassCourseResponse) SetHeaders(v map[string]*string) *UpdateRemoteClassCourseResponse {
	s.Headers = v
	return s
}

func (s *UpdateRemoteClassCourseResponse) SetStatusCode(v int32) *UpdateRemoteClassCourseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRemoteClassCourseResponse) SetBody(v *UpdateRemoteClassCourseResponseBody) *UpdateRemoteClassCourseResponse {
	s.Body = v
	return s
}

type UpdateRemoteClassDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRemoteClassDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassDeviceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassDeviceHeaders) SetCommonHeaders(v map[string]*string) *UpdateRemoteClassDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRemoteClassDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRemoteClassDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRemoteClassDeviceRequest struct {
	AuthCode   *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
}

func (s UpdateRemoteClassDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassDeviceRequest) SetAuthCode(v string) *UpdateRemoteClassDeviceRequest {
	s.AuthCode = &v
	return s
}

func (s *UpdateRemoteClassDeviceRequest) SetDeviceCode(v string) *UpdateRemoteClassDeviceRequest {
	s.DeviceCode = &v
	return s
}

func (s *UpdateRemoteClassDeviceRequest) SetDeviceName(v string) *UpdateRemoteClassDeviceRequest {
	s.DeviceName = &v
	return s
}

type UpdateRemoteClassDeviceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRemoteClassDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassDeviceResponseBody) SetSuccess(v bool) *UpdateRemoteClassDeviceResponseBody {
	s.Success = &v
	return s
}

type UpdateRemoteClassDeviceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRemoteClassDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRemoteClassDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRemoteClassDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateRemoteClassDeviceResponse) SetHeaders(v map[string]*string) *UpdateRemoteClassDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateRemoteClassDeviceResponse) SetStatusCode(v int32) *UpdateRemoteClassDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRemoteClassDeviceResponse) SetBody(v *UpdateRemoteClassDeviceResponseBody) *UpdateRemoteClassDeviceResponse {
	s.Body = v
	return s
}

type UpdateUniversityCourseGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUniversityCourseGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateUniversityCourseGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUniversityCourseGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUniversityCourseGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUniversityCourseGroupRequest struct {
	CourseGroupCode        *string                                                     `json:"courseGroupCode,omitempty" xml:"courseGroupCode,omitempty"`
	CourseGroupIntroduce   *string                                                     `json:"courseGroupIntroduce,omitempty" xml:"courseGroupIntroduce,omitempty"`
	CourseGroupName        *string                                                     `json:"courseGroupName,omitempty" xml:"courseGroupName,omitempty"`
	CourserGroupItemModels []*UpdateUniversityCourseGroupRequestCourserGroupItemModels `json:"courserGroupItemModels,omitempty" xml:"courserGroupItemModels,omitempty" type:"Repeated"`
	Ext                    *string                                                     `json:"ext,omitempty" xml:"ext,omitempty"`
	OpUserId               *string                                                     `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateUniversityCourseGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupCode(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupCode = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupIntroduce(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupIntroduce = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourseGroupName(v string) *UpdateUniversityCourseGroupRequest {
	s.CourseGroupName = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetCourserGroupItemModels(v []*UpdateUniversityCourseGroupRequestCourserGroupItemModels) *UpdateUniversityCourseGroupRequest {
	s.CourserGroupItemModels = v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetExt(v string) *UpdateUniversityCourseGroupRequest {
	s.Ext = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequest) SetOpUserId(v string) *UpdateUniversityCourseGroupRequest {
	s.OpUserId = &v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModels struct {
	ClassPeriodType           *int32                                                                             `json:"classPeriodType,omitempty" xml:"classPeriodType,omitempty"`
	ClassroomId               *int64                                                                             `json:"classroomId,omitempty" xml:"classroomId,omitempty"`
	CourseType                *int32                                                                             `json:"courseType,omitempty" xml:"courseType,omitempty"`
	CourserGroupItemEndDate   *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate   `json:"courserGroupItemEndDate,omitempty" xml:"courserGroupItemEndDate,omitempty" type:"Struct"`
	CourserGroupItemStartDate *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate `json:"courserGroupItemStartDate,omitempty" xml:"courserGroupItemStartDate,omitempty" type:"Struct"`
	DayOfWeek                 *int32                                                                             `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty"`
	SectionIndex              []*int32                                                                           `json:"sectionIndex,omitempty" xml:"sectionIndex,omitempty" type:"Repeated"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModels) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModels) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetClassPeriodType(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassPeriodType = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetClassroomId(v int64) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.ClassroomId = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourseType(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourseType = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemEndDate(v *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemEndDate = v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetCourserGroupItemStartDate(v *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.CourserGroupItemStartDate = v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetDayOfWeek(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.DayOfWeek = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModels) SetSectionIndex(v []*int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModels {
	s.SectionIndex = v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetDayOfMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.DayOfMonth = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Month = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate) SetYear(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemEndDate {
	s.Year = &v
	return s
}

type UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate struct {
	DayOfMonth *int32 `json:"dayOfMonth,omitempty" xml:"dayOfMonth,omitempty"`
	Month      *int32 `json:"month,omitempty" xml:"month,omitempty"`
	Year       *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetDayOfMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.DayOfMonth = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetMonth(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Month = &v
	return s
}

func (s *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate) SetYear(v int32) *UpdateUniversityCourseGroupRequestCourserGroupItemModelsCourserGroupItemStartDate {
	s.Year = &v
	return s
}

type UpdateUniversityCourseGroupResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateUniversityCourseGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupResponseBody) SetResult(v bool) *UpdateUniversityCourseGroupResponseBody {
	s.Result = &v
	return s
}

type UpdateUniversityCourseGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUniversityCourseGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUniversityCourseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUniversityCourseGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUniversityCourseGroupResponse) SetHeaders(v map[string]*string) *UpdateUniversityCourseGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUniversityCourseGroupResponse) SetStatusCode(v int32) *UpdateUniversityCourseGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUniversityCourseGroupResponse) SetBody(v *UpdateUniversityCourseGroupResponseBody) *UpdateUniversityCourseGroupResponse {
	s.Body = v
	return s
}

type UploadLearningDataCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadLearningDataCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadLearningDataCallbackHeaders) GoString() string {
	return s.String()
}

func (s *UploadLearningDataCallbackHeaders) SetCommonHeaders(v map[string]*string) *UploadLearningDataCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadLearningDataCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *UploadLearningDataCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadLearningDataCallbackRequest struct {
	BizId         *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	BizType       *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CorpId        *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	GeneratedTime *int64  `json:"generatedTime,omitempty" xml:"generatedTime,omitempty"`
	StudentUserId *string `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
	SubjectCode   *string `json:"subjectCode,omitempty" xml:"subjectCode,omitempty"`
}

func (s UploadLearningDataCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadLearningDataCallbackRequest) GoString() string {
	return s.String()
}

func (s *UploadLearningDataCallbackRequest) SetBizId(v string) *UploadLearningDataCallbackRequest {
	s.BizId = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetBizType(v string) *UploadLearningDataCallbackRequest {
	s.BizType = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetCorpId(v string) *UploadLearningDataCallbackRequest {
	s.CorpId = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetDeptId(v int64) *UploadLearningDataCallbackRequest {
	s.DeptId = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetGeneratedTime(v int64) *UploadLearningDataCallbackRequest {
	s.GeneratedTime = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetStudentUserId(v string) *UploadLearningDataCallbackRequest {
	s.StudentUserId = &v
	return s
}

func (s *UploadLearningDataCallbackRequest) SetSubjectCode(v string) *UploadLearningDataCallbackRequest {
	s.SubjectCode = &v
	return s
}

type UploadLearningDataCallbackResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadLearningDataCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadLearningDataCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *UploadLearningDataCallbackResponseBody) SetResult(v bool) *UploadLearningDataCallbackResponseBody {
	s.Result = &v
	return s
}

func (s *UploadLearningDataCallbackResponseBody) SetSuccess(v bool) *UploadLearningDataCallbackResponseBody {
	s.Success = &v
	return s
}

type UploadLearningDataCallbackResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadLearningDataCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadLearningDataCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadLearningDataCallbackResponse) GoString() string {
	return s.String()
}

func (s *UploadLearningDataCallbackResponse) SetHeaders(v map[string]*string) *UploadLearningDataCallbackResponse {
	s.Headers = v
	return s
}

func (s *UploadLearningDataCallbackResponse) SetStatusCode(v int32) *UploadLearningDataCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadLearningDataCallbackResponse) SetBody(v *UploadLearningDataCallbackResponseBody) *UploadLearningDataCallbackResponse {
	s.Body = v
	return s
}

type VPaasProxyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s VPaasProxyHeaders) String() string {
	return tea.Prettify(s)
}

func (s VPaasProxyHeaders) GoString() string {
	return s.String()
}

func (s *VPaasProxyHeaders) SetCommonHeaders(v map[string]*string) *VPaasProxyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VPaasProxyHeaders) SetXAcsDingtalkAccessToken(v string) *VPaasProxyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type VPaasProxyRequest struct {
	ActionCode *string `json:"actionCode,omitempty" xml:"actionCode,omitempty"`
	Params     *string `json:"params,omitempty" xml:"params,omitempty"`
	PublicKey  *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s VPaasProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s VPaasProxyRequest) GoString() string {
	return s.String()
}

func (s *VPaasProxyRequest) SetActionCode(v string) *VPaasProxyRequest {
	s.ActionCode = &v
	return s
}

func (s *VPaasProxyRequest) SetParams(v string) *VPaasProxyRequest {
	s.Params = &v
	return s
}

func (s *VPaasProxyRequest) SetPublicKey(v string) *VPaasProxyRequest {
	s.PublicKey = &v
	return s
}

type VPaasProxyResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s VPaasProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VPaasProxyResponseBody) GoString() string {
	return s.String()
}

func (s *VPaasProxyResponseBody) SetResult(v string) *VPaasProxyResponseBody {
	s.Result = &v
	return s
}

func (s *VPaasProxyResponseBody) SetTicket(v string) *VPaasProxyResponseBody {
	s.Ticket = &v
	return s
}

type VPaasProxyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VPaasProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VPaasProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s VPaasProxyResponse) GoString() string {
	return s.String()
}

func (s *VPaasProxyResponse) SetHeaders(v map[string]*string) *VPaasProxyResponse {
	s.Headers = v
	return s
}

func (s *VPaasProxyResponse) SetStatusCode(v int32) *VPaasProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *VPaasProxyResponse) SetBody(v *VPaasProxyResponseBody) *VPaasProxyResponse {
	s.Body = v
	return s
}

type ValidateNewGradeManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateNewGradeManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateNewGradeManagerHeaders) GoString() string {
	return s.String()
}

func (s *ValidateNewGradeManagerHeaders) SetCommonHeaders(v map[string]*string) *ValidateNewGradeManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateNewGradeManagerHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateNewGradeManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateNewGradeManagerRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ValidateNewGradeManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateNewGradeManagerRequest) GoString() string {
	return s.String()
}

func (s *ValidateNewGradeManagerRequest) SetUnionId(v string) *ValidateNewGradeManagerRequest {
	s.UnionId = &v
	return s
}

type ValidateNewGradeManagerResponseBody struct {
	MatchRule *bool `json:"matchRule,omitempty" xml:"matchRule,omitempty"`
}

func (s ValidateNewGradeManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateNewGradeManagerResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateNewGradeManagerResponseBody) SetMatchRule(v bool) *ValidateNewGradeManagerResponseBody {
	s.MatchRule = &v
	return s
}

type ValidateNewGradeManagerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateNewGradeManagerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateNewGradeManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateNewGradeManagerResponse) GoString() string {
	return s.String()
}

func (s *ValidateNewGradeManagerResponse) SetHeaders(v map[string]*string) *ValidateNewGradeManagerResponse {
	s.Headers = v
	return s
}

func (s *ValidateNewGradeManagerResponse) SetStatusCode(v int32) *ValidateNewGradeManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateNewGradeManagerResponse) SetBody(v *ValidateNewGradeManagerResponseBody) *ValidateNewGradeManagerResponse {
	s.Body = v
	return s
}

type ValidateUserRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateUserRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateUserRoleHeaders) GoString() string {
	return s.String()
}

func (s *ValidateUserRoleHeaders) SetCommonHeaders(v map[string]*string) *ValidateUserRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateUserRoleHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateUserRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateUserRoleRequest struct {
	TimeThreshold *int64  `json:"timeThreshold,omitempty" xml:"timeThreshold,omitempty"`
	UnionId       *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ValidateUserRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateUserRoleRequest) GoString() string {
	return s.String()
}

func (s *ValidateUserRoleRequest) SetTimeThreshold(v int64) *ValidateUserRoleRequest {
	s.TimeThreshold = &v
	return s
}

func (s *ValidateUserRoleRequest) SetUnionId(v string) *ValidateUserRoleRequest {
	s.UnionId = &v
	return s
}

type ValidateUserRoleResponseBody struct {
	MatchParentIdentity  *bool `json:"matchParentIdentity,omitempty" xml:"matchParentIdentity,omitempty"`
	MatchTeacherIdentity *bool `json:"matchTeacherIdentity,omitempty" xml:"matchTeacherIdentity,omitempty"`
}

func (s ValidateUserRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateUserRoleResponseBody) SetMatchParentIdentity(v bool) *ValidateUserRoleResponseBody {
	s.MatchParentIdentity = &v
	return s
}

func (s *ValidateUserRoleResponseBody) SetMatchTeacherIdentity(v bool) *ValidateUserRoleResponseBody {
	s.MatchTeacherIdentity = &v
	return s
}

type ValidateUserRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateUserRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateUserRoleResponse) GoString() string {
	return s.String()
}

func (s *ValidateUserRoleResponse) SetHeaders(v map[string]*string) *ValidateUserRoleResponse {
	s.Headers = v
	return s
}

func (s *ValidateUserRoleResponse) SetStatusCode(v int32) *ValidateUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateUserRoleResponse) SetBody(v *ValidateUserRoleResponseBody) *ValidateUserRoleResponse {
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

func (client *Client) ActivateDeviceWithOptions(request *ActivateDeviceRequest, headers *ActivateDeviceHeaders, runtime *util.RuntimeOptions) (_result *ActivateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LicenseKey)) {
		body["licenseKey"] = request.LicenseKey
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("ActivateDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices/activate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateDevice(request *ActivateDeviceRequest) (_result *ActivateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ActivateDeviceHeaders{}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.ActivateDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCompetitionRecordWithOptions(request *AddCompetitionRecordRequest, headers *AddCompetitionRecordHeaders, runtime *util.RuntimeOptions) (_result *AddCompetitionRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompetitionCode)) {
		body["competitionCode"] = request.CompetitionCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateCode)) {
		body["groupTemplateCode"] = request.GroupTemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.JoinGroup)) {
		body["joinGroup"] = request.JoinGroup
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantName)) {
		body["participantName"] = request.ParticipantName
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("AddCompetitionRecord"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/competitions/records"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCompetitionRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCompetitionRecord(request *AddCompetitionRecordRequest) (_result *AddCompetitionRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCompetitionRecordHeaders{}
	_result = &AddCompetitionRecordResponse{}
	_body, _err := client.AddCompetitionRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDeviceWithOptions(request *AddDeviceRequest, headers *AddDeviceHeaders, runtime *util.RuntimeOptions) (_result *AddDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("AddDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/devices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDevice(request *AddDeviceRequest) (_result *AddDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddDeviceHeaders{}
	_result = &AddDeviceResponse{}
	_body, _err := client.AddDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSchoolConfigWithOptions(request *AddSchoolConfigRequest, headers *AddSchoolConfigHeaders, runtime *util.RuntimeOptions) (_result *AddSchoolConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorName)) {
		body["operatorName"] = request.OperatorName
	}

	if !tea.BoolValue(util.IsUnset(request.TemperatureUpLimit)) {
		body["temperatureUpLimit"] = request.TemperatureUpLimit
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
		Action:      tea.String("AddSchoolConfig"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schools/configurations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSchoolConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSchoolConfig(request *AddSchoolConfigRequest) (_result *AddSchoolConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddSchoolConfigHeaders{}
	_result = &AddSchoolConfigResponse{}
	_body, _err := client.AddSchoolConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignClassWithOptions(request *AssignClassRequest, headers *AssignClassHeaders, runtime *util.RuntimeOptions) (_result *AssignClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.IsFinish)) {
		body["isFinish"] = request.IsFinish
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		body["studentId"] = request.StudentId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("AssignClass"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/newGrades/tasks/students/classes/assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignClassResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignClass(request *AssignClassRequest) (_result *AssignClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AssignClassHeaders{}
	_result = &AssignClassResponse{}
	_body, _err := client.AssignClassWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateWithOptions(request *BatchCreateRequest, headers *BatchCreateHeaders, runtime *util.RuntimeOptions) (_result *BatchCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		body["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.JsVersion)) {
		body["jsVersion"] = request.JsVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
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
		Action:      tea.String("BatchCreate"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreate(request *BatchCreateRequest) (_result *BatchCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchCreateHeaders{}
	_result = &BatchCreateResponse{}
	_body, _err := client.BatchCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchOrgCreateHWWithOptions(request *BatchOrgCreateHWRequest, headers *BatchOrgCreateHWHeaders, runtime *util.RuntimeOptions) (_result *BatchOrgCreateHWResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		body["attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.HwContent)) {
		body["hwContent"] = request.HwContent
	}

	if !tea.BoolValue(util.IsUnset(request.HwDeadline)) {
		body["hwDeadline"] = request.HwDeadline
	}

	if !tea.BoolValue(util.IsUnset(request.HwDeadlineOpen)) {
		body["hwDeadlineOpen"] = request.HwDeadlineOpen
	}

	if !tea.BoolValue(util.IsUnset(request.HwMedia)) {
		body["hwMedia"] = request.HwMedia
	}

	if !tea.BoolValue(util.IsUnset(request.HwPhoto)) {
		body["hwPhoto"] = request.HwPhoto
	}

	if !tea.BoolValue(util.IsUnset(request.HwTitle)) {
		body["hwTitle"] = request.HwTitle
	}

	if !tea.BoolValue(util.IsUnset(request.HwType)) {
		body["hwType"] = request.HwType
	}

	if !tea.BoolValue(util.IsUnset(request.HwVideo)) {
		body["hwVideo"] = request.HwVideo
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSelectItemList)) {
		body["openSelectItemList"] = request.OpenSelectItemList
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledRelease)) {
		body["scheduledRelease"] = request.ScheduledRelease
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTime)) {
		body["scheduledTime"] = request.ScheduledTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRole)) {
		body["targetRole"] = request.TargetRole
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherName)) {
		body["teacherName"] = request.TeacherName
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserId)) {
		body["teacherUserId"] = request.TeacherUserId
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
		Action:      tea.String("BatchOrgCreateHW"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/homeworks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchOrgCreateHWResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchOrgCreateHW(request *BatchOrgCreateHWRequest) (_result *BatchOrgCreateHWResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchOrgCreateHWHeaders{}
	_result = &BatchOrgCreateHWResponse{}
	_body, _err := client.BatchOrgCreateHWWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, headers *CancelOrderHeaders, runtime *util.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CancelOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelOrderHeaders{}
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelSnsOrderWithOptions(request *CancelSnsOrderRequest, headers *CancelSnsOrderHeaders, runtime *util.RuntimeOptions) (_result *CancelSnsOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		body["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CancelSnsOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/snsUserOrders/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelSnsOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelSnsOrder(request *CancelSnsOrderRequest) (_result *CancelSnsOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelSnsOrderHeaders{}
	_result = &CancelSnsOrderResponse{}
	_body, _err := client.CancelSnsOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelUserOrderWithOptions(request *CancelUserOrderRequest, headers *CancelUserOrderHeaders, runtime *util.RuntimeOptions) (_result *CancelUserOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		body["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CancelUserOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/userOrders/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelUserOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUserOrder(request *CancelUserOrderRequest) (_result *CancelUserOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelUserOrderHeaders{}
	_result = &CancelUserOrderResponse{}
	_body, _err := client.CancelUserOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardBatchQueryCardsWithOptions(request *CardBatchQueryCardsRequest, headers *CardBatchQueryCardsHeaders, runtime *util.RuntimeOptions) (_result *CardBatchQueryCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		body["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardIds)) {
		body["cardIds"] = request.CardIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
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
		Action:      tea.String("CardBatchQueryCards"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards/tasks/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardBatchQueryCardsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardBatchQueryCards(request *CardBatchQueryCardsRequest) (_result *CardBatchQueryCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardBatchQueryCardsHeaders{}
	_result = &CardBatchQueryCardsResponse{}
	_body, _err := client.CardBatchQueryCardsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardDeleteCardWithOptions(request *CardDeleteCardRequest, headers *CardDeleteCardHeaders, runtime *util.RuntimeOptions) (_result *CardDeleteCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		query["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		query["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardId)) {
		query["cardId"] = request.CardId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
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
		Action:      tea.String("CardDeleteCard"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardDeleteCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardDeleteCard(request *CardDeleteCardRequest) (_result *CardDeleteCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardDeleteCardHeaders{}
	_result = &CardDeleteCardResponse{}
	_body, _err := client.CardDeleteCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardEndCardWithOptions(request *CardEndCardRequest, headers *CardEndCardHeaders, runtime *util.RuntimeOptions) (_result *CardEndCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		body["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		body["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardId)) {
		body["cardId"] = request.CardId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
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
		Action:      tea.String("CardEndCard"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards/finish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardEndCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardEndCard(request *CardEndCardRequest) (_result *CardEndCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardEndCardHeaders{}
	_result = &CardEndCardResponse{}
	_body, _err := client.CardEndCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardGetCardWithOptions(request *CardGetCardRequest, headers *CardGetCardHeaders, runtime *util.RuntimeOptions) (_result *CardGetCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardId)) {
		query["cardId"] = request.CardId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
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
		Action:      tea.String("CardGetCard"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardGetCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardGetCard(request *CardGetCardRequest) (_result *CardGetCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardGetCardHeaders{}
	_result = &CardGetCardResponse{}
	_body, _err := client.CardGetCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardGetCardFinishProgressWithOptions(request *CardGetCardFinishProgressRequest, headers *CardGetCardFinishProgressHeaders, runtime *util.RuntimeOptions) (_result *CardGetCardFinishProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		query["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		query["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardId)) {
		query["cardId"] = request.CardId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CardGetCardFinishProgress"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards/completionProgress"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardGetCardFinishProgressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardGetCardFinishProgress(request *CardGetCardFinishProgressRequest) (_result *CardGetCardFinishProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardGetCardFinishProgressHeaders{}
	_result = &CardGetCardFinishProgressResponse{}
	_body, _err := client.CardGetCardFinishProgressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardQueryCardFeedsWithOptions(request *CardQueryCardFeedsRequest, headers *CardQueryCardFeedsHeaders, runtime *util.RuntimeOptions) (_result *CardQueryCardFeedsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizCode)) {
		query["cardBizCode"] = request.CardBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		query["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardId)) {
		query["cardId"] = request.CardId
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.FeedType)) {
		query["feedType"] = request.FeedType
	}

	if !tea.BoolValue(util.IsUnset(request.NeedFinishProcess)) {
		query["needFinishProcess"] = request.NeedFinishProcess
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
	}

	if !tea.BoolValue(util.IsUnset(request.SubBizId)) {
		query["subBizId"] = request.SubBizId
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
		Action:      tea.String("CardQueryCardFeeds"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/cards/feeds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardQueryCardFeedsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardQueryCardFeeds(request *CardQueryCardFeedsRequest) (_result *CardQueryCardFeedsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardQueryCardFeedsHeaders{}
	_result = &CardQueryCardFeedsResponse{}
	_body, _err := client.CardQueryCardFeedsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckRestrictionWithOptions(request *CheckRestrictionRequest, headers *CheckRestrictionHeaders, runtime *util.RuntimeOptions) (_result *CheckRestrictionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualAmount)) {
		body["actualAmount"] = request.ActualAmount
	}

	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
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
		Action:      tea.String("CheckRestriction"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/restrictions/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckRestrictionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckRestriction(request *CheckRestrictionRequest) (_result *CheckRestrictionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckRestrictionHeaders{}
	_result = &CheckRestrictionResponse{}
	_body, _err := client.CheckRestrictionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumePointWithOptions(request *ConsumePointRequest, headers *ConsumePointHeaders, runtime *util.RuntimeOptions) (_result *ConsumePointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
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
		Action:      tea.String("ConsumePoint"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/poins/consume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsumePointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConsumePoint(request *ConsumePointRequest) (_result *ConsumePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumePointHeaders{}
	_result = &ConsumePointResponse{}
	_body, _err := client.ConsumePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CourseSchedulingComplimentNoticeWithOptions(request *CourseSchedulingComplimentNoticeRequest, headers *CourseSchedulingComplimentNoticeHeaders, runtime *util.RuntimeOptions) (_result *CourseSchedulingComplimentNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("CourseSchedulingComplimentNotice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schedules/finishNotify"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CourseSchedulingComplimentNoticeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CourseSchedulingComplimentNotice(request *CourseSchedulingComplimentNoticeRequest) (_result *CourseSchedulingComplimentNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CourseSchedulingComplimentNoticeHeaders{}
	_result = &CourseSchedulingComplimentNoticeResponse{}
	_body, _err := client.CourseSchedulingComplimentNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppOrderWithOptions(request *CreateAppOrderRequest, headers *CreateAppOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateAppOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualAmount)) {
		body["actualAmount"] = request.ActualAmount
	}

	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		body["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
	}

	if !tea.BoolValue(util.IsUnset(request.LabelAmount)) {
		body["labelAmount"] = request.LabelAmount
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantOrderNo)) {
		body["merchantOrderNo"] = request.MerchantOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OuterUserId)) {
		body["outerUserId"] = request.OuterUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CreateAppOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/appOrders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppOrder(request *CreateAppOrderRequest) (_result *CreateAppOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAppOrderHeaders{}
	_result = &CreateAppOrderResponse{}
	_body, _err := client.CreateAppOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomClassWithOptions(request *CreateCustomClassRequest, headers *CreateCustomClassHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomClass)) {
		body["customClass"] = request.CustomClass
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		body["superId"] = request.SuperId
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
		Action:      tea.String("CreateCustomClass"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/customClasses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomClassResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomClass(request *CreateCustomClassRequest) (_result *CreateCustomClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomClassHeaders{}
	_result = &CreateCustomClassResponse{}
	_body, _err := client.CreateCustomClassWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomDeptWithOptions(request *CreateCustomDeptRequest, headers *CreateCustomDeptHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomDept)) {
		body["customDept"] = request.CustomDept
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		body["superId"] = request.SuperId
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
		Action:      tea.String("CreateCustomDept"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/customDepts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomDept(request *CreateCustomDeptRequest) (_result *CreateCustomDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomDeptHeaders{}
	_result = &CreateCustomDeptResponse{}
	_body, _err := client.CreateCustomDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEduAssetSpaceWithOptions(request *CreateEduAssetSpaceRequest, headers *CreateEduAssetSpaceHeaders, runtime *util.RuntimeOptions) (_result *CreateEduAssetSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceDesc)) {
		body["spaceDesc"] = request.SpaceDesc
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceIcon)) {
		body["spaceIcon"] = request.SpaceIcon
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceName)) {
		body["spaceName"] = request.SpaceName
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
		Action:      tea.String("CreateEduAssetSpace"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/assets/spaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEduAssetSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEduAssetSpace(request *CreateEduAssetSpaceRequest) (_result *CreateEduAssetSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEduAssetSpaceHeaders{}
	_result = &CreateEduAssetSpaceResponse{}
	_body, _err := client.CreateEduAssetSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFulfilRecordWithOptions(request *CreateFulfilRecordRequest, headers *CreateFulfilRecordHeaders, runtime *util.RuntimeOptions) (_result *CreateFulfilRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTime)) {
		body["bizTime"] = request.BizTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
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
		Action:      tea.String("CreateFulfilRecord"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/fulfilRecords"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFulfilRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFulfilRecord(request *CreateFulfilRecordRequest) (_result *CreateFulfilRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFulfilRecordHeaders{}
	_result = &CreateFulfilRecordResponse{}
	_body, _err := client.CreateFulfilRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInviteUrlWithOptions(request *CreateInviteUrlRequest, headers *CreateInviteUrlHeaders, runtime *util.RuntimeOptions) (_result *CreateInviteUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOperator)) {
		body["targetOperator"] = request.TargetOperator
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
		Action:      tea.String("CreateInviteUrl"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/orgRelations/inviteUrls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInviteUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInviteUrl(request *CreateInviteUrlRequest) (_result *CreateInviteUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInviteUrlHeaders{}
	_result = &CreateInviteUrlResponse{}
	_body, _err := client.CreateInviteUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateItemWithOptions(request *CreateItemRequest, headers *CreateItemHeaders, runtime *util.RuntimeOptions) (_result *CreateItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EffectType)) {
		body["effectType"] = request.EffectType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptUser)) {
		body["optUser"] = request.OptUser
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodType)) {
		body["periodType"] = request.PeriodType
	}

	if !tea.BoolValue(util.IsUnset(request.Price)) {
		body["price"] = request.Price
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("CreateItem"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/items"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateItemResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateItem(request *CreateItemRequest) (_result *CreateItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateItemHeaders{}
	_result = &CreateItemResponse{}
	_body, _err := client.CreateItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, headers *CreateOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualAmount)) {
		body["actualAmount"] = request.ActualAmount
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
	}

	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ftoken)) {
		body["ftoken"] = request.Ftoken
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalParams)) {
		body["terminalParams"] = request.TerminalParams
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrderHeaders{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderFlowWithOptions(request *CreateOrderFlowRequest, headers *CreateOrderFlowHeaders, runtime *util.RuntimeOptions) (_result *CreateOrderFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualAmount)) {
		body["actualAmount"] = request.ActualAmount
	}

	if !tea.BoolValue(util.IsUnset(request.AlipayUid)) {
		body["alipayUid"] = request.AlipayUid
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
	}

	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GuardianUserId)) {
		body["guardianUserId"] = request.GuardianUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.TotalAmount)) {
		body["totalAmount"] = request.TotalAmount
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
		Action:      tea.String("CreateOrderFlow"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders/flows"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderFlowResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrderFlow(request *CreateOrderFlowRequest) (_result *CreateOrderFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrderFlowHeaders{}
	_result = &CreateOrderFlowResponse{}
	_body, _err := client.CreateOrderFlowWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhysicalClassroomWithOptions(request *CreatePhysicalClassroomRequest, headers *CreatePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *CreatePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomBuilding)) {
		body["classroomBuilding"] = request.ClassroomBuilding
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomCampus)) {
		body["classroomCampus"] = request.ClassroomCampus
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomFloor)) {
		body["classroomFloor"] = request.ClassroomFloor
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomName)) {
		body["classroomName"] = request.ClassroomName
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomNumber)) {
		body["classroomNumber"] = request.ClassroomNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DirectBroadcast)) {
		body["directBroadcast"] = request.DirectBroadcast
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
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
		Action:      tea.String("CreatePhysicalClassroom"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/physicalClassrooms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePhysicalClassroomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhysicalClassroom(request *CreatePhysicalClassroomRequest) (_result *CreatePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreatePhysicalClassroomHeaders{}
	_result = &CreatePhysicalClassroomResponse{}
	_body, _err := client.CreatePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRefundFlowWithOptions(request *CreateRefundFlowRequest, headers *CreateRefundFlowHeaders, runtime *util.RuntimeOptions) (_result *CreateRefundFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorName)) {
		body["operatorName"] = request.OperatorName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CreateRefundFlow"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/refunds/flows"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRefundFlowResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRefundFlow(request *CreateRefundFlowRequest) (_result *CreateRefundFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRefundFlowHeaders{}
	_result = &CreateRefundFlowResponse{}
	_body, _err := client.CreateRefundFlowWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRemoteClassCourseWithOptions(request *CreateRemoteClassCourseRequest, headers *CreateRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *CreateRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttendParticipants)) {
		body["attendParticipants"] = request.AttendParticipants
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TeachingParticipant)) {
		body["teachingParticipant"] = request.TeachingParticipant
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
		Action:      tea.String("CreateRemoteClassCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/courses"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRemoteClassCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRemoteClassCourse(request *CreateRemoteClassCourseRequest) (_result *CreateRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRemoteClassCourseHeaders{}
	_result = &CreateRemoteClassCourseResponse{}
	_body, _err := client.CreateRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSectionConfigWithOptions(request *CreateSectionConfigRequest, headers *CreateSectionConfigHeaders, runtime *util.RuntimeOptions) (_result *CreateSectionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.SectionConfigs)) {
		body["sectionConfigs"] = request.SectionConfigs
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
		Action:      tea.String("CreateSectionConfig"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/sectionConfigs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSectionConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSectionConfig(request *CreateSectionConfigRequest) (_result *CreateSectionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSectionConfigHeaders{}
	_result = &CreateSectionConfigResponse{}
	_body, _err := client.CreateSectionConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnsAppOrderWithOptions(request *CreateSnsAppOrderRequest, headers *CreateSnsAppOrderHeaders, runtime *util.RuntimeOptions) (_result *CreateSnsAppOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualAmount)) {
		body["actualAmount"] = request.ActualAmount
	}

	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		body["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		body["bizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
	}

	if !tea.BoolValue(util.IsUnset(request.LabelAmount)) {
		body["labelAmount"] = request.LabelAmount
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantOrderNo)) {
		body["merchantOrderNo"] = request.MerchantOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("CreateSnsAppOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/snsAppOrders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnsAppOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnsAppOrder(request *CreateSnsAppOrderRequest) (_result *CreateSnsAppOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSnsAppOrderHeaders{}
	_result = &CreateSnsAppOrderResponse{}
	_body, _err := client.CreateSnsAppOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStsTokenWithOptions(request *CreateStsTokenRequest, headers *CreateStsTokenHeaders, runtime *util.RuntimeOptions) (_result *CreateStsTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceSn)) {
		body["deviceSn"] = request.DeviceSn
	}

	if !tea.BoolValue(util.IsUnset(request.StsType)) {
		body["stsType"] = request.StsType
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
		Action:      tea.String("CreateStsToken"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/ststoken"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStsTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStsToken(request *CreateStsTokenRequest) (_result *CreateStsTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateStsTokenHeaders{}
	_result = &CreateStsTokenResponse{}
	_body, _err := client.CreateStsTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, headers *CreateTokenHeaders, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("CreateToken"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/tokens"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTokenHeaders{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUniversityCourseGroupWithOptions(request *CreateUniversityCourseGroupRequest, headers *CreateUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *CreateUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupIntroduce)) {
		body["courseGroupIntroduce"] = request.CourseGroupIntroduce
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupName)) {
		body["courseGroupName"] = request.CourseGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.CourserGroupItemModels)) {
		body["courserGroupItemModels"] = request.CourserGroupItemModels
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCourseGroupCode)) {
		body["isvCourseGroupCode"] = request.IsvCourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCode)) {
		body["periodCode"] = request.PeriodCode
	}

	if !tea.BoolValue(util.IsUnset(request.SchoolYear)) {
		body["schoolYear"] = request.SchoolYear
	}

	if !tea.BoolValue(util.IsUnset(request.Semester)) {
		body["semester"] = request.Semester
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectName)) {
		body["subjectName"] = request.SubjectName
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherInfos)) {
		body["teacherInfos"] = request.TeacherInfos
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
		Action:      tea.String("CreateUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUniversityCourseGroup(request *CreateUniversityCourseGroupRequest) (_result *CreateUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUniversityCourseGroupHeaders{}
	_result = &CreateUniversityCourseGroupResponse{}
	_body, _err := client.CreateUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUniversityStudentWithOptions(request *CreateUniversityStudentRequest, headers *CreateUniversityStudentHeaders, runtime *util.RuntimeOptions) (_result *CreateUniversityStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityNumber)) {
		body["identityNumber"] = request.IdentityNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StudentNumber)) {
		body["studentNumber"] = request.StudentNumber
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
		Action:      tea.String("CreateUniversityStudent"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/students"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUniversityStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUniversityStudent(request *CreateUniversityStudentRequest) (_result *CreateUniversityStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUniversityStudentHeaders{}
	_result = &CreateUniversityStudentResponse{}
	_body, _err := client.CreateUniversityStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUniversityTeacherWithOptions(request *CreateUniversityTeacherRequest, headers *CreateUniversityTeacherHeaders, runtime *util.RuntimeOptions) (_result *CreateUniversityTeacherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserId)) {
		body["teacherUserId"] = request.TeacherUserId
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
		Action:      tea.String("CreateUniversityTeacher"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/teachers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUniversityTeacherResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUniversityTeacher(request *CreateUniversityTeacherRequest) (_result *CreateUniversityTeacherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUniversityTeacherHeaders{}
	_result = &CreateUniversityTeacherResponse{}
	_body, _err := client.CreateUniversityTeacherWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactivateDeviceWithOptions(request *DeactivateDeviceRequest, headers *DeactivateDeviceHeaders, runtime *util.RuntimeOptions) (_result *DeactivateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("DeactivateDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices/deactivate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactivateDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactivateDevice(request *DeactivateDeviceRequest) (_result *DeactivateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeactivateDeviceHeaders{}
	_result = &DeactivateDeviceResponse{}
	_body, _err := client.DeactivateDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeductPointWithOptions(request *DeductPointRequest, headers *DeductPointHeaders, runtime *util.RuntimeOptions) (_result *DeductPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.DeductDesc)) {
		body["deductDesc"] = request.DeductDesc
	}

	if !tea.BoolValue(util.IsUnset(request.DeductDetailUrl)) {
		body["deductDetailUrl"] = request.DeductDetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DeductNum)) {
		body["deductNum"] = request.DeductNum
	}

	if !tea.BoolValue(util.IsUnset(request.PointType)) {
		body["pointType"] = request.PointType
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
		Action:      tea.String("DeductPoint"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/points/deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeductPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeductPoint(request *DeductPointRequest) (_result *DeductPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeductPointHeaders{}
	_result = &DeductPointResponse{}
	_body, _err := client.DeductPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeptWithOptions(deptId *string, request *DeleteDeptRequest, headers *DeleteDeptHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("DeleteDept"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/depts/" + tea.StringValue(deptId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDept(deptId *string, request *DeleteDeptRequest) (_result *DeleteDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeptHeaders{}
	_result = &DeleteDeptResponse{}
	_body, _err := client.DeleteDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceWithOptions(request *DeleteDeviceRequest, headers *DeleteDeviceHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("DeleteDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (_result *DeleteDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeviceHeaders{}
	_result = &DeleteDeviceResponse{}
	_body, _err := client.DeleteDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceOrgWithOptions(request *DeleteDeviceOrgRequest, headers *DeleteDeviceOrgHeaders, runtime *util.RuntimeOptions) (_result *DeleteDeviceOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["deviceCode"] = request.DeviceCode
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
		Action:      tea.String("DeleteDeviceOrg"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/deviceOrgs"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceOrg(request *DeleteDeviceOrgRequest) (_result *DeleteDeviceOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDeviceOrgHeaders{}
	_result = &DeleteDeviceOrgResponse{}
	_body, _err := client.DeleteDeviceOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGuardianWithOptions(classId *string, userId *string, request *DeleteGuardianRequest, headers *DeleteGuardianHeaders, runtime *util.RuntimeOptions) (_result *DeleteGuardianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StuId)) {
		query["stuId"] = request.StuId
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
		Action:      tea.String("DeleteGuardian"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/" + tea.StringValue(classId) + "/guardians/" + tea.StringValue(userId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGuardianResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGuardian(classId *string, userId *string, request *DeleteGuardianRequest) (_result *DeleteGuardianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGuardianHeaders{}
	_result = &DeleteGuardianResponse{}
	_body, _err := client.DeleteGuardianWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOrgRelationWithOptions(request *DeleteOrgRelationRequest, headers *DeleteOrgRelationHeaders, runtime *util.RuntimeOptions) (_result *DeleteOrgRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("DeleteOrgRelation"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/orgRelations"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOrgRelationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOrgRelation(request *DeleteOrgRelationRequest) (_result *DeleteOrgRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteOrgRelationHeaders{}
	_result = &DeleteOrgRelationResponse{}
	_body, _err := client.DeleteOrgRelationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePhysicalClassroomWithOptions(request *DeletePhysicalClassroomRequest, headers *DeletePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *DeletePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		query["classroomId"] = request.ClassroomId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("DeletePhysicalClassroom"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/physicalClassrooms"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePhysicalClassroomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePhysicalClassroom(request *DeletePhysicalClassroomRequest) (_result *DeletePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePhysicalClassroomHeaders{}
	_result = &DeletePhysicalClassroomResponse{}
	_body, _err := client.DeletePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRemoteClassCourseWithOptions(courseCode *string, request *DeleteRemoteClassCourseRequest, headers *DeleteRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *DeleteRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
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
		Action:      tea.String("DeleteRemoteClassCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/courses/" + tea.StringValue(courseCode)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRemoteClassCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRemoteClassCourse(courseCode *string, request *DeleteRemoteClassCourseRequest) (_result *DeleteRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteRemoteClassCourseHeaders{}
	_result = &DeleteRemoteClassCourseResponse{}
	_body, _err := client.DeleteRemoteClassCourseWithOptions(courseCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStudentWithOptions(classId *string, userId *string, request *DeleteStudentRequest, headers *DeleteStudentHeaders, runtime *util.RuntimeOptions) (_result *DeleteStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("DeleteStudent"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/" + tea.StringValue(classId) + "/students/" + tea.StringValue(userId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStudent(classId *string, userId *string, request *DeleteStudentRequest) (_result *DeleteStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteStudentHeaders{}
	_result = &DeleteStudentResponse{}
	_body, _err := client.DeleteStudentWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTeacherWithOptions(classId *string, userId *string, request *DeleteTeacherRequest, headers *DeleteTeacherHeaders, runtime *util.RuntimeOptions) (_result *DeleteTeacherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adviser)) {
		query["adviser"] = request.Adviser
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("DeleteTeacher"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/" + tea.StringValue(classId) + "/teachers/" + tea.StringValue(userId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTeacherResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTeacher(classId *string, userId *string, request *DeleteTeacherRequest) (_result *DeleteTeacherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTeacherHeaders{}
	_result = &DeleteTeacherResponse{}
	_body, _err := client.DeleteTeacherWithOptions(classId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUniversityCourseGroupWithOptions(request *DeleteUniversityCourseGroupRequest, headers *DeleteUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *DeleteUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		query["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("DeleteUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUniversityCourseGroup(request *DeleteUniversityCourseGroupRequest) (_result *DeleteUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUniversityCourseGroupHeaders{}
	_result = &DeleteUniversityCourseGroupResponse{}
	_body, _err := client.DeleteUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUniversityStudentWithOptions(request *DeleteUniversityStudentRequest, headers *DeleteUniversityStudentHeaders, runtime *util.RuntimeOptions) (_result *DeleteUniversityStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		query["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserId)) {
		query["studentUserId"] = request.StudentUserId
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
		Action:      tea.String("DeleteUniversityStudent"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/students"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUniversityStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUniversityStudent(request *DeleteUniversityStudentRequest) (_result *DeleteUniversityStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUniversityStudentHeaders{}
	_result = &DeleteUniversityStudentResponse{}
	_body, _err := client.DeleteUniversityStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUniversityTeacherWithOptions(request *DeleteUniversityTeacherRequest, headers *DeleteUniversityTeacherHeaders, runtime *util.RuntimeOptions) (_result *DeleteUniversityTeacherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		query["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserId)) {
		query["teacherUserId"] = request.TeacherUserId
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
		Action:      tea.String("DeleteUniversityTeacher"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/teachers"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUniversityTeacherResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUniversityTeacher(request *DeleteUniversityTeacherRequest) (_result *DeleteUniversityTeacherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUniversityTeacherHeaders{}
	_result = &DeleteUniversityTeacherResponse{}
	_body, _err := client.DeleteUniversityTeacherWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeviceHeartbeatWithOptions(request *DeviceHeartbeatRequest, headers *DeviceHeartbeatHeaders, runtime *util.RuntimeOptions) (_result *DeviceHeartbeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("DeviceHeartbeat"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/heartbeats/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeviceHeartbeatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeviceHeartbeat(request *DeviceHeartbeatRequest) (_result *DeviceHeartbeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceHeartbeatHeaders{}
	_result = &DeviceHeartbeatResponse{}
	_body, _err := client.DeviceHeartbeatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EduFindUserRolesByUserIdWithOptions(request *EduFindUserRolesByUserIdRequest, headers *EduFindUserRolesByUserIdHeaders, runtime *util.RuntimeOptions) (_result *EduFindUserRolesByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		query["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.HasOrgRole)) {
		query["hasOrgRole"] = request.HasOrgRole
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
		Action:      tea.String("EduFindUserRolesByUserId"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/users/allRoles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EduFindUserRolesByUserIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EduFindUserRolesByUserId(request *EduFindUserRolesByUserIdRequest) (_result *EduFindUserRolesByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EduFindUserRolesByUserIdHeaders{}
	_result = &EduFindUserRolesByUserIdResponse{}
	_body, _err := client.EduFindUserRolesByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EduListUserByFromUserIdsWithOptions(request *EduListUserByFromUserIdsRequest, headers *EduListUserByFromUserIdsHeaders, runtime *util.RuntimeOptions) (_result *EduListUserByFromUserIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		query["classId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GuardianUserId)) {
		query["guardianUserId"] = request.GuardianUserId
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
		Action:      tea.String("EduListUserByFromUserIds"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/users/allRelations/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EduListUserByFromUserIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EduListUserByFromUserIds(request *EduListUserByFromUserIdsRequest) (_result *EduListUserByFromUserIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EduListUserByFromUserIdsHeaders{}
	_result = &EduListUserByFromUserIdsResponse{}
	_body, _err := client.EduListUserByFromUserIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EduTeacherListWithOptions(request *EduTeacherListRequest, headers *EduTeacherListHeaders, runtime *util.RuntimeOptions) (_result *EduTeacherListResponse, _err error) {
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
		Action:      tea.String("EduTeacherList"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/teachers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EduTeacherListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EduTeacherList(request *EduTeacherListRequest) (_result *EduTeacherListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EduTeacherListHeaders{}
	_result = &EduTeacherListResponse{}
	_body, _err := client.EduTeacherListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EndCourseWithOptions(request *EndCourseRequest, headers *EndCourseHeaders, runtime *util.RuntimeOptions) (_result *EndCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.LivePlayInfoList)) {
		body["livePlayInfoList"] = request.LivePlayInfoList
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
		Action:      tea.String("EndCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courses/end"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EndCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EndCourse(request *EndCourseRequest) (_result *EndCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EndCourseHeaders{}
	_result = &EndCourseResponse{}
	_body, _err := client.EndCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBindChildInfoWithOptions(request *GetBindChildInfoRequest, headers *GetBindChildInfoHeaders, runtime *util.RuntimeOptions) (_result *GetBindChildInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchoolCorpId)) {
		query["schoolCorpId"] = request.SchoolCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserId)) {
		query["studentUserId"] = request.StudentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("GetBindChildInfo"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/families/childs/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBindChildInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBindChildInfo(request *GetBindChildInfoRequest) (_result *GetBindChildInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBindChildInfoHeaders{}
	_result = &GetBindChildInfoResponse{}
	_body, _err := client.GetBindChildInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultChildWithOptions(headers *GetDefaultChildHeaders, runtime *util.RuntimeOptions) (_result *GetDefaultChildResponse, _err error) {
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
		Action:      tea.String("GetDefaultChild"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/defaultChildren"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultChildResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultChild() (_result *GetDefaultChildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDefaultChildHeaders{}
	_result = &GetDefaultChildResponse{}
	_body, _err := client.GetDefaultChildWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEduUserIdentityWithOptions(request *GetEduUserIdentityRequest, headers *GetEduUserIdentityHeaders, runtime *util.RuntimeOptions) (_result *GetEduUserIdentityResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEduUserIdentity"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/apollos/activities/userIdentities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEduUserIdentityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEduUserIdentity(request *GetEduUserIdentityRequest) (_result *GetEduUserIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEduUserIdentityHeaders{}
	_result = &GetEduUserIdentityResponse{}
	_body, _err := client.GetEduUserIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenCourseDetailWithOptions(courseId *string, headers *GetOpenCourseDetailHeaders, runtime *util.RuntimeOptions) (_result *GetOpenCourseDetailResponse, _err error) {
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
		Action:      tea.String("GetOpenCourseDetail"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/openCourse/" + tea.StringValue(courseId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenCourseDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenCourseDetail(courseId *string) (_result *GetOpenCourseDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOpenCourseDetailHeaders{}
	_result = &GetOpenCourseDetailResponse{}
	_body, _err := client.GetOpenCourseDetailWithOptions(courseId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenCoursesWithOptions(request *GetOpenCoursesRequest, headers *GetOpenCoursesHeaders, runtime *util.RuntimeOptions) (_result *GetOpenCoursesResponse, _err error) {
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
		Action:      tea.String("GetOpenCourses"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/openCourses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpenCoursesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenCourses(request *GetOpenCoursesRequest) (_result *GetOpenCoursesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOpenCoursesHeaders{}
	_result = &GetOpenCoursesResponse{}
	_body, _err := client.GetOpenCoursesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPointActionRecordWithOptions(tmpReq *GetPointActionRecordRequest, headers *GetPointActionRecordHeaders, runtime *util.RuntimeOptions) (_result *GetPointActionRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPointActionRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("GetPointActionRecord"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/points/actionRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPointActionRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPointActionRecord(request *GetPointActionRecordRequest) (_result *GetPointActionRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPointActionRecordHeaders{}
	_result = &GetPointActionRecordResponse{}
	_body, _err := client.GetPointActionRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPointInfoWithOptions(request *GetPointInfoRequest, headers *GetPointInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPointInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PointType)) {
		query["pointType"] = request.PointType
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
		Action:      tea.String("GetPointInfo"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/points/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPointInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPointInfo(request *GetPointInfoRequest) (_result *GetPointInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPointInfoHeaders{}
	_result = &GetPointInfoResponse{}
	_body, _err := client.GetPointInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRemoteClassCourseWithOptions(courseCode *string, request *GetRemoteClassCourseRequest, headers *GetRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *GetRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("GetRemoteClassCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/courses/" + tea.StringValue(courseCode)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRemoteClassCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRemoteClassCourse(courseCode *string, request *GetRemoteClassCourseRequest) (_result *GetRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRemoteClassCourseHeaders{}
	_result = &GetRemoteClassCourseResponse{}
	_body, _err := client.GetRemoteClassCourseWithOptions(courseCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShareRoleMembersWithOptions(shareRoleCode *string, headers *GetShareRoleMembersHeaders, runtime *util.RuntimeOptions) (_result *GetShareRoleMembersResponse, _err error) {
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
		Action:      tea.String("GetShareRoleMembers"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/shareRoles/" + tea.StringValue(shareRoleCode) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShareRoleMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShareRoleMembers(shareRoleCode *string) (_result *GetShareRoleMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetShareRoleMembersHeaders{}
	_result = &GetShareRoleMembersResponse{}
	_body, _err := client.GetShareRoleMembersWithOptions(shareRoleCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetShareRolesWithOptions(headers *GetShareRolesHeaders, runtime *util.RuntimeOptions) (_result *GetShareRolesResponse, _err error) {
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
		Action:      tea.String("GetShareRoles"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/shareRoles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShareRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetShareRoles() (_result *GetShareRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetShareRolesHeaders{}
	_result = &GetShareRolesResponse{}
	_body, _err := client.GetShareRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskListWithOptions(request *GetTaskListRequest, headers *GetTaskListHeaders, runtime *util.RuntimeOptions) (_result *GetTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskYear)) {
		query["taskYear"] = request.TaskYear
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
		Action:      tea.String("GetTaskList"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/newGrades/tasks/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskList(request *GetTaskListRequest) (_result *GetTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskListHeaders{}
	_result = &GetTaskListResponse{}
	_body, _err := client.GetTaskListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskStudentListWithOptions(request *GetTaskStudentListRequest, headers *GetTaskStudentListHeaders, runtime *util.RuntimeOptions) (_result *GetTaskStudentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
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
		Action:      tea.String("GetTaskStudentList"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/newGrades/tasks/students/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStudentListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskStudentList(request *GetTaskStudentListRequest) (_result *GetTaskStudentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskStudentListHeaders{}
	_result = &GetTaskStudentListResponse{}
	_body, _err := client.GetTaskStudentListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitCoursesOfClassWithOptions(classId *string, request *InitCoursesOfClassRequest, headers *InitCoursesOfClassHeaders, runtime *util.RuntimeOptions) (_result *InitCoursesOfClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Courses)) {
		body["courses"] = request.Courses
	}

	if !tea.BoolValue(util.IsUnset(request.SectionConfig)) {
		body["sectionConfig"] = request.SectionConfig
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
		Action:      tea.String("InitCoursesOfClass"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/" + tea.StringValue(classId) + "/courses/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InitCoursesOfClassResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitCoursesOfClass(classId *string, request *InitCoursesOfClassRequest) (_result *InitCoursesOfClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitCoursesOfClassHeaders{}
	_result = &InitCoursesOfClassResponse{}
	_body, _err := client.InitCoursesOfClassWithOptions(classId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitDeviceWithOptions(request *InitDeviceRequest, headers *InitDeviceHeaders, runtime *util.RuntimeOptions) (_result *InitDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptPubKey)) {
		body["encryptPubKey"] = request.EncryptPubKey
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("InitDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/devices/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InitDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitDevice(request *InitDeviceRequest) (_result *InitDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitDeviceHeaders{}
	_result = &InitDeviceResponse{}
	_body, _err := client.InitDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitVPaasDeviceWithOptions(request *InitVPaasDeviceRequest, headers *InitVPaasDeviceHeaders, runtime *util.RuntimeOptions) (_result *InitVPaasDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("InitVPaasDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices/init"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InitVPaasDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitVPaasDevice(request *InitVPaasDeviceRequest) (_result *InitVPaasDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitVPaasDeviceHeaders{}
	_result = &InitVPaasDeviceResponse{}
	_body, _err := client.InitVPaasDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertSectionConfigWithOptions(request *InsertSectionConfigRequest, headers *InsertSectionConfigHeaders, runtime *util.RuntimeOptions) (_result *InsertSectionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.End)) {
		body["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleName)) {
		body["scheduleName"] = request.ScheduleName
	}

	if !tea.BoolValue(util.IsUnset(request.SectionModels)) {
		body["sectionModels"] = request.SectionModels
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["start"] = request.Start
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
		Action:      tea.String("InsertSectionConfig"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schedules/configs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertSectionConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertSectionConfig(request *InsertSectionConfigRequest) (_result *InsertSectionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertSectionConfigHeaders{}
	_result = &InsertSectionConfigResponse{}
	_body, _err := client.InsertSectionConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsvDataWriteWithOptions(request *IsvDataWriteRequest, headers *IsvDataWriteHeaders, runtime *util.RuntimeOptions) (_result *IsvDataWriteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectCode)) {
		body["objectCode"] = request.ObjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.RowValueList)) {
		body["rowValueList"] = request.RowValueList
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
		Action:      tea.String("IsvDataWrite"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/datas/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IsvDataWriteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsvDataWrite(request *IsvDataWriteRequest) (_result *IsvDataWriteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsvDataWriteHeaders{}
	_result = &IsvDataWriteResponse{}
	_body, _err := client.IsvDataWriteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsvMetadataQueryWithOptions(request *IsvMetadataQueryRequest, headers *IsvMetadataQueryHeaders, runtime *util.RuntimeOptions) (_result *IsvMetadataQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectCode)) {
		query["objectCode"] = request.ObjectCode
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
		Action:      tea.String("IsvMetadataQuery"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/datas/metadatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IsvMetadataQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsvMetadataQuery(request *IsvMetadataQueryRequest) (_result *IsvMetadataQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsvMetadataQueryHeaders{}
	_result = &IsvMetadataQueryResponse{}
	_body, _err := client.IsvMetadataQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrderWithOptions(request *ListOrderRequest, headers *ListOrderHeaders, runtime *util.RuntimeOptions) (_result *ListOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["createTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["createTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		body["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TradeNo)) {
		body["tradeNo"] = request.TradeNo
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
		Action:      tea.String("ListOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrder(request *ListOrderRequest) (_result *ListOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrderHeaders{}
	_result = &ListOrderResponse{}
	_body, _err := client.ListOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveStudentWithOptions(request *MoveStudentRequest, headers *MoveStudentHeaders, runtime *util.RuntimeOptions) (_result *MoveStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OriginClassId)) {
		body["originClassId"] = request.OriginClassId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetClassId)) {
		body["targetClassId"] = request.TargetClassId
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
		Action:      tea.String("MoveStudent"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/students/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveStudent(request *MoveStudentRequest) (_result *MoveStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveStudentHeaders{}
	_result = &MoveStudentResponse{}
	_body, _err := client.MoveStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageQueryDevicesWithOptions(request *PageQueryDevicesRequest, headers *PageQueryDevicesHeaders, runtime *util.RuntimeOptions) (_result *PageQueryDevicesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("PageQueryDevices"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PageQueryDevicesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageQueryDevices(request *PageQueryDevicesRequest) (_result *PageQueryDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageQueryDevicesHeaders{}
	_result = &PageQueryDevicesResponse{}
	_body, _err := client.PageQueryDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PayOrderWithOptions(request *PayOrderRequest, headers *PayOrderHeaders, runtime *util.RuntimeOptions) (_result *PayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("PayOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders/pay"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PayOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PayOrder(request *PayOrderRequest) (_result *PayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PayOrderHeaders{}
	_result = &PayOrderResponse{}
	_body, _err := client.PayOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollingConfirmStatusWithOptions(request *PollingConfirmStatusRequest, headers *PollingConfirmStatusHeaders, runtime *util.RuntimeOptions) (_result *PollingConfirmStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		query["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		query["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("PollingConfirmStatus"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courses/pollingConfirmStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PollingConfirmStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollingConfirmStatus(request *PollingConfirmStatusRequest) (_result *PollingConfirmStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PollingConfirmStatusHeaders{}
	_result = &PollingConfirmStatusResponse{}
	_body, _err := client.PollingConfirmStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreDialWithOptions(request *PreDialRequest, headers *PreDialHeaders, runtime *util.RuntimeOptions) (_result *PreDialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallerUserId)) {
		body["callerUserId"] = request.CallerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserId)) {
		body["receiverUserId"] = request.ReceiverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("PreDial"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/devices/preDial"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PreDialResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreDial(request *PreDialRequest) (_result *PreDialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PreDialHeaders{}
	_result = &PreDialResponse{}
	_body, _err := client.PreDialWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProvidePointWithOptions(request *ProvidePointRequest, headers *ProvidePointHeaders, runtime *util.RuntimeOptions) (_result *ProvidePointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCode)) {
		body["actionCode"] = request.ActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.PointType)) {
		body["pointType"] = request.PointType
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
		Action:      tea.String("ProvidePoint"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/points/provide"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ProvidePointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProvidePoint(request *ProvidePointRequest) (_result *ProvidePointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProvidePointHeaders{}
	_result = &ProvidePointResponse{}
	_body, _err := client.ProvidePointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllSubjectsFromClassScheduleWithOptions(tmpReq *QueryAllSubjectsFromClassScheduleRequest, headers *QueryAllSubjectsFromClassScheduleHeaders, runtime *util.RuntimeOptions) (_result *QueryAllSubjectsFromClassScheduleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryAllSubjectsFromClassScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClassIds)) {
		request.ClassIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClassIds, tea.String("classIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIdsShrink)) {
		query["classIds"] = request.ClassIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodCode)) {
		query["periodCode"] = request.PeriodCode
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
		Action:      tea.String("QueryAllSubjectsFromClassSchedule"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/subjects/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllSubjectsFromClassScheduleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllSubjectsFromClassSchedule(request *QueryAllSubjectsFromClassScheduleRequest) (_result *QueryAllSubjectsFromClassScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllSubjectsFromClassScheduleHeaders{}
	_result = &QueryAllSubjectsFromClassScheduleResponse{}
	_body, _err := client.QueryAllSubjectsFromClassScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleWithOptions(request *QueryClassScheduleRequest, headers *QueryClassScheduleHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriberType)) {
		query["subscriberType"] = request.SubscriberType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SectionIndexList)) {
		body["sectionIndexList"] = request.SectionIndexList
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriberIds)) {
		body["subscriberIds"] = request.SubscriberIds
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
		Action:      tea.String("QueryClassSchedule"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/schedules/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClassScheduleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassSchedule(request *QueryClassScheduleRequest) (_result *QueryClassScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleHeaders{}
	_result = &QueryClassScheduleResponse{}
	_body, _err := client.QueryClassScheduleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleByTimeSchoolWithOptions(request *QueryClassScheduleByTimeSchoolRequest, headers *QueryClassScheduleByTimeSchoolHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleByTimeSchoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("QueryClassScheduleByTimeSchool"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schools/classes/courses "),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClassScheduleByTimeSchoolResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassScheduleByTimeSchool(request *QueryClassScheduleByTimeSchoolRequest) (_result *QueryClassScheduleByTimeSchoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleByTimeSchoolHeaders{}
	_result = &QueryClassScheduleByTimeSchoolResponse{}
	_body, _err := client.QueryClassScheduleByTimeSchoolWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClassScheduleConfigWithOptions(tmpReq *QueryClassScheduleConfigRequest, headers *QueryClassScheduleConfigHeaders, runtime *util.RuntimeOptions) (_result *QueryClassScheduleConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryClassScheduleConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ClassIds)) {
		request.ClassIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClassIds, tea.String("classIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIdsShrink)) {
		query["classIds"] = request.ClassIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("QueryClassScheduleConfig"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schedules/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClassScheduleConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClassScheduleConfig(request *QueryClassScheduleConfigRequest) (_result *QueryClassScheduleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryClassScheduleConfigHeaders{}
	_result = &QueryClassScheduleConfigResponse{}
	_body, _err := client.QueryClassScheduleConfigWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpass/devices/query"),
		Method:      tea.String("POST"),
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

func (client *Client) QueryDeviceListByCorpIdWithOptions(request *QueryDeviceListByCorpIdRequest, headers *QueryDeviceListByCorpIdHeaders, runtime *util.RuntimeOptions) (_result *QueryDeviceListByCorpIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("QueryDeviceListByCorpId"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/devices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDeviceListByCorpIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDeviceListByCorpId(request *QueryDeviceListByCorpIdRequest) (_result *QueryDeviceListByCorpIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDeviceListByCorpIdHeaders{}
	_result = &QueryDeviceListByCorpIdResponse{}
	_body, _err := client.QueryDeviceListByCorpIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEduAssetSpacesWithOptions(request *QueryEduAssetSpacesRequest, headers *QueryEduAssetSpacesHeaders, runtime *util.RuntimeOptions) (_result *QueryEduAssetSpacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["bizCode"] = request.BizCode
	}

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
		Action:      tea.String("QueryEduAssetSpaces"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/assets/spaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEduAssetSpacesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEduAssetSpaces(request *QueryEduAssetSpacesRequest) (_result *QueryEduAssetSpacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryEduAssetSpacesHeaders{}
	_result = &QueryEduAssetSpacesResponse{}
	_body, _err := client.QueryEduAssetSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupIdWithOptions(request *QueryGroupIdRequest, headers *QueryGroupIdHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("QueryGroupId"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/faces/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupId(request *QueryGroupIdRequest) (_result *QueryGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupIdHeaders{}
	_result = &QueryGroupIdResponse{}
	_body, _err := client.QueryGroupIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderWithOptions(request *QueryOrderRequest, headers *QueryOrderHeaders, runtime *util.RuntimeOptions) (_result *QueryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		query["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		query["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
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
		Action:      tea.String("QueryOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrder(request *QueryOrderRequest) (_result *QueryOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrderHeaders{}
	_result = &QueryOrderResponse{}
	_body, _err := client.QueryOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgRelationListWithOptions(request *QueryOrgRelationListRequest, headers *QueryOrgRelationListHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgRelationListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
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
		Action:      tea.String("QueryOrgRelationList"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/orgRelations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgRelationListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgRelationList(request *QueryOrgRelationListRequest) (_result *QueryOrgRelationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgRelationListHeaders{}
	_result = &QueryOrgRelationListResponse{}
	_body, _err := client.QueryOrgRelationListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgSecretKeyWithOptions(request *QueryOrgSecretKeyRequest, headers *QueryOrgSecretKeyHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		query["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("QueryOrgSecretKey"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/secretKeys"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgSecretKeyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgSecretKey(request *QueryOrgSecretKeyRequest) (_result *QueryOrgSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgSecretKeyHeaders{}
	_result = &QueryOrgSecretKeyResponse{}
	_body, _err := client.QueryOrgSecretKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgTypeWithOptions(headers *QueryOrgTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgTypeResponse, _err error) {
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
		Action:      tea.String("QueryOrgType"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orgTypes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgType() (_result *QueryOrgTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgTypeHeaders{}
	_result = &QueryOrgTypeResponse{}
	_body, _err := client.QueryOrgTypeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPayResultWithOptions(request *QueryPayResultRequest, headers *QueryPayResultHeaders, runtime *util.RuntimeOptions) (_result *QueryPayResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		body["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
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
		Action:      tea.String("QueryPayResult"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/payResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPayResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPayResult(request *QueryPayResultRequest) (_result *QueryPayResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPayResultHeaders{}
	_result = &QueryPayResultResponse{}
	_body, _err := client.QueryPayResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPhysicalClassroomWithOptions(request *QueryPhysicalClassroomRequest, headers *QueryPhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *QueryPhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		query["classroomId"] = request.ClassroomId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("QueryPhysicalClassroom"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/physicalClassrooms"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPhysicalClassroomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPhysicalClassroom(request *QueryPhysicalClassroomRequest) (_result *QueryPhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPhysicalClassroomHeaders{}
	_result = &QueryPhysicalClassroomResponse{}
	_body, _err := client.QueryPhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPurchaseInfoWithOptions(request *QueryPurchaseInfoRequest, headers *QueryPurchaseInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryPurchaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		query["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("QueryPurchaseInfo"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/users/purchases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPurchaseInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPurchaseInfo(request *QueryPurchaseInfoRequest) (_result *QueryPurchaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPurchaseInfoHeaders{}
	_result = &QueryPurchaseInfoResponse{}
	_body, _err := client.QueryPurchaseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRemoteClassCourseWithOptions(request *QueryRemoteClassCourseRequest, headers *QueryRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *QueryRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		query["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("QueryRemoteClassCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/courses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRemoteClassCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRemoteClassCourse(request *QueryRemoteClassCourseRequest) (_result *QueryRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRemoteClassCourseHeaders{}
	_result = &QueryRemoteClassCourseResponse{}
	_body, _err := client.QueryRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySchoolUserFaceWithOptions(request *QuerySchoolUserFaceRequest, headers *QuerySchoolUserFaceHeaders, runtime *util.RuntimeOptions) (_result *QuerySchoolUserFaceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("QuerySchoolUserFace"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/schools/faces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchoolUserFaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySchoolUserFace(request *QuerySchoolUserFaceRequest) (_result *QuerySchoolUserFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySchoolUserFaceHeaders{}
	_result = &QuerySchoolUserFaceResponse{}
	_body, _err := client.QuerySchoolUserFaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySnsOrderWithOptions(request *QuerySnsOrderRequest, headers *QuerySnsOrderHeaders, runtime *util.RuntimeOptions) (_result *QuerySnsOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlipayAppId)) {
		query["alipayAppId"] = request.AlipayAppId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		query["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		query["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		query["signature"] = request.Signature
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
		Action:      tea.String("QuerySnsOrder"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/snsOrders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySnsOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySnsOrder(request *QuerySnsOrderRequest) (_result *QuerySnsOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySnsOrderHeaders{}
	_result = &QuerySnsOrderResponse{}
	_body, _err := client.QuerySnsOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatisticsDataWithOptions(request *QueryStatisticsDataRequest, headers *QueryStatisticsDataHeaders, runtime *util.RuntimeOptions) (_result *QueryStatisticsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SectionIndexList)) {
		body["sectionIndexList"] = request.SectionIndexList
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserIds)) {
		body["teacherUserIds"] = request.TeacherUserIds
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
		Action:      tea.String("QueryStatisticsData"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/schedules/statisticData/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStatisticsDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStatisticsData(request *QueryStatisticsDataRequest) (_result *QueryStatisticsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryStatisticsDataHeaders{}
	_result = &QueryStatisticsDataResponse{}
	_body, _err := client.QueryStatisticsDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySubjectTeachersWithOptions(request *QuerySubjectTeachersRequest, headers *QuerySubjectTeachersHeaders, runtime *util.RuntimeOptions) (_result *QuerySubjectTeachersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIds)) {
		query["classIds"] = request.ClassIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectCode)) {
		query["subjectCode"] = request.SubjectCode
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
		Action:      tea.String("QuerySubjectTeachers"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/subjects/teachers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySubjectTeachersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySubjectTeachers(request *QuerySubjectTeachersRequest) (_result *QuerySubjectTeachersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySubjectTeachersHeaders{}
	_result = &QuerySubjectTeachersResponse{}
	_body, _err := client.QuerySubjectTeachersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTeachSubjectsWithOptions(request *QueryTeachSubjectsRequest, headers *QueryTeachSubjectsHeaders, runtime *util.RuntimeOptions) (_result *QueryTeachSubjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassIds)) {
		query["classIds"] = request.ClassIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("QueryTeachSubjects"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/teachers/subjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTeachSubjectsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTeachSubjects(request *QueryTeachSubjectsRequest) (_result *QueryTeachSubjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTeachSubjectsHeaders{}
	_result = &QueryTeachSubjectsResponse{}
	_body, _err := client.QueryTeachSubjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUniversityCourseGroupWithOptions(request *QueryUniversityCourseGroupRequest, headers *QueryUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		query["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
		Action:      tea.String("QueryUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUniversityCourseGroup(request *QueryUniversityCourseGroupRequest) (_result *QueryUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUniversityCourseGroupHeaders{}
	_result = &QueryUniversityCourseGroupResponse{}
	_body, _err := client.QueryUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserFaceWithOptions(request *QueryUserFaceRequest, headers *QueryUserFaceHeaders, runtime *util.RuntimeOptions) (_result *QueryUserFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		query["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("QueryUserFace"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/users/faces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserFaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserFace(request *QueryUserFaceRequest) (_result *QueryUserFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserFaceHeaders{}
	_result = &QueryUserFaceResponse{}
	_body, _err := client.QueryUserFaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserPayInfoWithOptions(request *QueryUserPayInfoRequest, headers *QueryUserPayInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserPayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FaceId)) {
		query["faceId"] = request.FaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("QueryUserPayInfo"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/orders/payInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserPayInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserPayInfo(request *QueryUserPayInfoRequest) (_result *QueryUserPayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserPayInfoHeaders{}
	_result = &QueryUserPayInfoResponse{}
	_body, _err := client.QueryUserPayInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDeviceWithOptions(request *RemoveDeviceRequest, headers *RemoveDeviceHeaders, runtime *util.RuntimeOptions) (_result *RemoveDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantId)) {
		query["merchantId"] = request.MerchantId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
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
		Action:      tea.String("RemoveDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/devices"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDevice(request *RemoveDeviceRequest) (_result *RemoveDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveDeviceHeaders{}
	_result = &RemoveDeviceResponse{}
	_body, _err := client.RemoveDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportDeviceLogWithOptions(request *ReportDeviceLogRequest, headers *ReportDeviceLogHeaders, runtime *util.RuntimeOptions) (_result *ReportDeviceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("ReportDeviceLog"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/deviceLogs/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportDeviceLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportDeviceLog(request *ReportDeviceLogRequest) (_result *ReportDeviceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReportDeviceLogHeaders{}
	_result = &ReportDeviceLogResponse{}
	_body, _err := client.ReportDeviceLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportDeviceUseLogWithOptions(request *ReportDeviceUseLogRequest, headers *ReportDeviceUseLogHeaders, runtime *util.RuntimeOptions) (_result *ReportDeviceUseLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
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
		Action:      tea.String("ReportDeviceUseLog"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/deviceUseLogs/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportDeviceUseLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportDeviceUseLog(request *ReportDeviceUseLogRequest) (_result *ReportDeviceUseLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReportDeviceUseLogHeaders{}
	_result = &ReportDeviceUseLogResponse{}
	_body, _err := client.ReportDeviceUseLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackDeductPointWithOptions(request *RollbackDeductPointRequest, headers *RollbackDeductPointHeaders, runtime *util.RuntimeOptions) (_result *RollbackDeductPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.PointType)) {
		body["pointType"] = request.PointType
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
		Action:      tea.String("RollbackDeductPoint"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/deductPoints/rollback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackDeductPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackDeductPoint(request *RollbackDeductPointRequest) (_result *RollbackDeductPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RollbackDeductPointHeaders{}
	_result = &RollbackDeductPointResponse{}
	_body, _err := client.RollbackDeductPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveClassLearningDataWithOptions(request *SaveClassLearningDataRequest, headers *SaveClassLearningDataHeaders, runtime *util.RuntimeOptions) (_result *SaveClassLearningDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignNum)) {
		body["assignNum"] = request.AssignNum
	}

	if !tea.BoolValue(util.IsUnset(request.AssignStudentUserIds)) {
		body["assignStudentUserIds"] = request.AssignStudentUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSuffix)) {
		body["fileSuffix"] = request.FileSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratedTime)) {
		body["generatedTime"] = request.GeneratedTime
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionNum)) {
		body["questionNum"] = request.QuestionNum
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionPictureNum)) {
		body["questionPictureNum"] = request.QuestionPictureNum
	}

	if !tea.BoolValue(util.IsUnset(request.StandardAnswerPictureNum)) {
		body["standardAnswerPictureNum"] = request.StandardAnswerPictureNum
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectCode)) {
		body["subjectCode"] = request.SubjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.TeacherUserId)) {
		body["teacherUserId"] = request.TeacherUserId
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
		Action:      tea.String("SaveClassLearningData"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/learnings/datas/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveClassLearningDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveClassLearningData(request *SaveClassLearningDataRequest) (_result *SaveClassLearningDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveClassLearningDataHeaders{}
	_result = &SaveClassLearningDataResponse{}
	_body, _err := client.SaveClassLearningDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveStudentLearningDataWithOptions(request *SaveStudentLearningDataRequest, headers *SaveStudentLearningDataHeaders, runtime *util.RuntimeOptions) (_result *SaveStudentLearningDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssignNum)) {
		body["assignNum"] = request.AssignNum
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CorrectNum)) {
		body["correctNum"] = request.CorrectNum
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSuffix)) {
		body["fileSuffix"] = request.FileSuffix
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratedTime)) {
		body["generatedTime"] = request.GeneratedTime
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionNum)) {
		body["questionNum"] = request.QuestionNum
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserId)) {
		body["studentUserId"] = request.StudentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectCode)) {
		body["subjectCode"] = request.SubjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitNum)) {
		body["submitNum"] = request.SubmitNum
	}

	if !tea.BoolValue(util.IsUnset(request.WrongQuestions)) {
		body["wrongQuestions"] = request.WrongQuestions
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
		Action:      tea.String("SaveStudentLearningData"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/students/learnings/datas/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveStudentLearningDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveStudentLearningData(request *SaveStudentLearningDataRequest) (_result *SaveStudentLearningDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveStudentLearningDataHeaders{}
	_result = &SaveStudentLearningDataResponse{}
	_body, _err := client.SaveStudentLearningDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTeachersWithOptions(request *SearchTeachersRequest, headers *SearchTeachersHeaders, runtime *util.RuntimeOptions) (_result *SearchTeachersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NameKeyword)) {
		query["nameKeyword"] = request.NameKeyword
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
		Action:      tea.String("SearchTeachers"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/teachers/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTeachersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTeachers(request *SearchTeachersRequest) (_result *SearchTeachersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchTeachersHeaders{}
	_result = &SearchTeachersResponse{}
	_body, _err := client.SearchTeachersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["fromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		body["sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserIdList)) {
		body["toUserIdList"] = request.ToUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("SendMessage"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCourseWithOptions(request *StartCourseRequest, headers *StartCourseHeaders, runtime *util.RuntimeOptions) (_result *StartCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.LivePlayInfoList)) {
		body["livePlayInfoList"] = request.LivePlayInfoList
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
		Action:      tea.String("StartCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courses/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCourse(request *StartCourseRequest) (_result *StartCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCourseHeaders{}
	_result = &StartCourseResponse{}
	_body, _err := client.StartCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCoursePrepareWithOptions(request *StartCoursePrepareRequest, headers *StartCoursePrepareHeaders, runtime *util.RuntimeOptions) (_result *StartCoursePrepareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseDate)) {
		body["courseDate"] = request.CourseDate
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
	}

	if !tea.BoolValue(util.IsUnset(request.IsvCode)) {
		body["isvCode"] = request.IsvCode
	}

	if !tea.BoolValue(util.IsUnset(request.LiveCoverImage)) {
		body["liveCoverImage"] = request.LiveCoverImage
	}

	if !tea.BoolValue(util.IsUnset(request.SectionIndex)) {
		body["sectionIndex"] = request.SectionIndex
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
		Action:      tea.String("StartCoursePrepare"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courses/prepare"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCoursePrepareResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCoursePrepare(request *StartCoursePrepareRequest) (_result *StartCoursePrepareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCoursePrepareHeaders{}
	_result = &StartCoursePrepareResponse{}
	_body, _err := client.StartCoursePrepareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeUniversityCourseGroupWithOptions(request *SubscribeUniversityCourseGroupRequest, headers *SubscribeUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *SubscribeUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserIds)) {
		body["studentUserIds"] = request.StudentUserIds
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
		Action:      tea.String("SubscribeUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups/subscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeUniversityCourseGroup(request *SubscribeUniversityCourseGroupRequest) (_result *SubscribeUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeUniversityCourseGroupHeaders{}
	_result = &SubscribeUniversityCourseGroupResponse{}
	_body, _err := client.SubscribeUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeUniversityCourseGroupWithOptions(request *UnsubscribeUniversityCourseGroupRequest, headers *UnsubscribeUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *UnsubscribeUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserIds)) {
		body["studentUserIds"] = request.StudentUserIds
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
		Action:      tea.String("UnsubscribeUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups/unsubscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnsubscribeUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeUniversityCourseGroup(request *UnsubscribeUniversityCourseGroupRequest) (_result *UnsubscribeUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnsubscribeUniversityCourseGroupHeaders{}
	_result = &UnsubscribeUniversityCourseGroupResponse{}
	_body, _err := client.UnsubscribeUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCoursesOfClassWithOptions(classId *string, request *UpdateCoursesOfClassRequest, headers *UpdateCoursesOfClassHeaders, runtime *util.RuntimeOptions) (_result *UpdateCoursesOfClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Courses)) {
		body["courses"] = request.Courses
	}

	if !tea.BoolValue(util.IsUnset(request.SectionConfig)) {
		body["sectionConfig"] = request.SectionConfig
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
		Action:      tea.String("UpdateCoursesOfClass"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/classes/" + tea.StringValue(classId) + "/courses/schedules"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCoursesOfClassResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCoursesOfClass(classId *string, request *UpdateCoursesOfClassRequest) (_result *UpdateCoursesOfClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCoursesOfClassHeaders{}
	_result = &UpdateCoursesOfClassResponse{}
	_body, _err := client.UpdateCoursesOfClassWithOptions(classId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePhysicalClassroomWithOptions(request *UpdatePhysicalClassroomRequest, headers *UpdatePhysicalClassroomHeaders, runtime *util.RuntimeOptions) (_result *UpdatePhysicalClassroomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassroomBuilding)) {
		body["classroomBuilding"] = request.ClassroomBuilding
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomCampus)) {
		body["classroomCampus"] = request.ClassroomCampus
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomFloor)) {
		body["classroomFloor"] = request.ClassroomFloor
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomId)) {
		body["classroomId"] = request.ClassroomId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomName)) {
		body["classroomName"] = request.ClassroomName
	}

	if !tea.BoolValue(util.IsUnset(request.ClassroomNumber)) {
		body["classroomNumber"] = request.ClassroomNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DirectBroadcast)) {
		body["directBroadcast"] = request.DirectBroadcast
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
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
		Action:      tea.String("UpdatePhysicalClassroom"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/physicalClassrooms"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePhysicalClassroomResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePhysicalClassroom(request *UpdatePhysicalClassroomRequest) (_result *UpdatePhysicalClassroomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePhysicalClassroomHeaders{}
	_result = &UpdatePhysicalClassroomResponse{}
	_body, _err := client.UpdatePhysicalClassroomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRemoteClassCourseWithOptions(request *UpdateRemoteClassCourseRequest, headers *UpdateRemoteClassCourseHeaders, runtime *util.RuntimeOptions) (_result *UpdateRemoteClassCourseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttendParticipants)) {
		body["attendParticipants"] = request.AttendParticipants
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		body["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseCode)) {
		body["courseCode"] = request.CourseCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseName)) {
		body["courseName"] = request.CourseName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TeachingParticipant)) {
		body["teachingParticipant"] = request.TeachingParticipant
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
		Action:      tea.String("UpdateRemoteClassCourse"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/courses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRemoteClassCourseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRemoteClassCourse(request *UpdateRemoteClassCourseRequest) (_result *UpdateRemoteClassCourseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRemoteClassCourseHeaders{}
	_result = &UpdateRemoteClassCourseResponse{}
	_body, _err := client.UpdateRemoteClassCourseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRemoteClassDeviceWithOptions(request *UpdateRemoteClassDeviceRequest, headers *UpdateRemoteClassDeviceHeaders, runtime *util.RuntimeOptions) (_result *UpdateRemoteClassDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["authCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["deviceName"] = request.DeviceName
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
		Action:      tea.String("UpdateRemoteClassDevice"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/remoteClasses/deviceNames"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRemoteClassDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRemoteClassDevice(request *UpdateRemoteClassDeviceRequest) (_result *UpdateRemoteClassDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRemoteClassDeviceHeaders{}
	_result = &UpdateRemoteClassDeviceResponse{}
	_body, _err := client.UpdateRemoteClassDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUniversityCourseGroupWithOptions(request *UpdateUniversityCourseGroupRequest, headers *UpdateUniversityCourseGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateUniversityCourseGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseGroupCode)) {
		body["courseGroupCode"] = request.CourseGroupCode
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupIntroduce)) {
		body["courseGroupIntroduce"] = request.CourseGroupIntroduce
	}

	if !tea.BoolValue(util.IsUnset(request.CourseGroupName)) {
		body["courseGroupName"] = request.CourseGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.CourserGroupItemModels)) {
		body["courserGroupItemModels"] = request.CourserGroupItemModels
	}

	if !tea.BoolValue(util.IsUnset(request.Ext)) {
		body["ext"] = request.Ext
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
		Action:      tea.String("UpdateUniversityCourseGroup"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/universities/courseGroups"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUniversityCourseGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUniversityCourseGroup(request *UpdateUniversityCourseGroupRequest) (_result *UpdateUniversityCourseGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUniversityCourseGroupHeaders{}
	_result = &UpdateUniversityCourseGroupResponse{}
	_body, _err := client.UpdateUniversityCourseGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadLearningDataCallbackWithOptions(request *UploadLearningDataCallbackRequest, headers *UploadLearningDataCallbackHeaders, runtime *util.RuntimeOptions) (_result *UploadLearningDataCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratedTime)) {
		body["generatedTime"] = request.GeneratedTime
	}

	if !tea.BoolValue(util.IsUnset(request.StudentUserId)) {
		body["studentUserId"] = request.StudentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectCode)) {
		body["subjectCode"] = request.SubjectCode
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
		Action:      tea.String("UploadLearningDataCallback"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/uploadLearnings/datas/callback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadLearningDataCallbackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadLearningDataCallback(request *UploadLearningDataCallbackRequest) (_result *UploadLearningDataCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadLearningDataCallbackHeaders{}
	_result = &UploadLearningDataCallbackResponse{}
	_body, _err := client.UploadLearningDataCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VPaasProxyWithOptions(request *VPaasProxyRequest, headers *VPaasProxyHeaders, runtime *util.RuntimeOptions) (_result *VPaasProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCode)) {
		body["actionCode"] = request.ActionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKey)) {
		body["publicKey"] = request.PublicKey
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
		Action:      tea.String("VPaasProxy"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/vpaas/proxy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &VPaasProxyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VPaasProxy(request *VPaasProxyRequest) (_result *VPaasProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &VPaasProxyHeaders{}
	_result = &VPaasProxyResponse{}
	_body, _err := client.VPaasProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateNewGradeManagerWithOptions(request *ValidateNewGradeManagerRequest, headers *ValidateNewGradeManagerHeaders, runtime *util.RuntimeOptions) (_result *ValidateNewGradeManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("ValidateNewGradeManager"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/newGrades/tasks/validate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateNewGradeManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateNewGradeManager(request *ValidateNewGradeManagerRequest) (_result *ValidateNewGradeManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateNewGradeManagerHeaders{}
	_result = &ValidateNewGradeManagerResponse{}
	_body, _err := client.ValidateNewGradeManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateUserRoleWithOptions(request *ValidateUserRoleRequest, headers *ValidateUserRoleHeaders, runtime *util.RuntimeOptions) (_result *ValidateUserRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TimeThreshold)) {
		body["timeThreshold"] = request.TimeThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("ValidateUserRole"),
		Version:     tea.String("edu_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/edu/users/roles/validate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateUserRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateUserRole(request *ValidateUserRoleRequest) (_result *ValidateUserRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateUserRoleHeaders{}
	_result = &ValidateUserRoleResponse{}
	_body, _err := client.ValidateUserRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
