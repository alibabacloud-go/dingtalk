// This file is auto-generated, don't edit it. Thanks.
package smart_device_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
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
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// This parameter is required.
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
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeviceVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateExportDeviceStatisticTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateExportDeviceStatisticTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateExportDeviceStatisticTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateExportDeviceStatisticTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateExportDeviceStatisticTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateExportDeviceStatisticTaskRequest struct {
	AiSheetTemplateId *int64  `json:"aiSheetTemplateId,omitempty" xml:"aiSheetTemplateId,omitempty"`
	CreatorCorpId     *string `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	CreatorUnionId    *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	TaskName          *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateExportDeviceStatisticTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskRequest) SetAiSheetTemplateId(v int64) *CreateExportDeviceStatisticTaskRequest {
	s.AiSheetTemplateId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskRequest) SetCreatorCorpId(v string) *CreateExportDeviceStatisticTaskRequest {
	s.CreatorCorpId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskRequest) SetCreatorUnionId(v string) *CreateExportDeviceStatisticTaskRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskRequest) SetTaskName(v string) *CreateExportDeviceStatisticTaskRequest {
	s.TaskName = &v
	return s
}

type CreateExportDeviceStatisticTaskResponseBody struct {
	ExportStatisticTaskDTO *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO `json:"exportStatisticTaskDTO,omitempty" xml:"exportStatisticTaskDTO,omitempty" type:"Struct"`
}

func (s CreateExportDeviceStatisticTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskResponseBody) SetExportStatisticTaskDTO(v *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) *CreateExportDeviceStatisticTaskResponseBody {
	s.ExportStatisticTaskDTO = v
	return s
}

type CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO struct {
	AiSheetDocumentOpenDTO *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO `json:"aiSheetDocumentOpenDTO,omitempty" xml:"aiSheetDocumentOpenDTO,omitempty" type:"Struct"`
	CorpId                 *string                                                                                  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	TaskId                 *string                                                                                  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskName               *string                                                                                  `json:"taskName,omitempty" xml:"taskName,omitempty"`
	UnionId                *string                                                                                  `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) SetAiSheetDocumentOpenDTO(v *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO {
	s.AiSheetDocumentOpenDTO = v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) SetCorpId(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO {
	s.CorpId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) SetTaskId(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO {
	s.TaskId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) SetTaskName(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO {
	s.TaskName = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO) SetUnionId(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTO {
	s.UnionId = &v
	return s
}

type CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO struct {
	AiSheetTemplateId *int64  `json:"aiSheetTemplateId,omitempty" xml:"aiSheetTemplateId,omitempty"`
	CorpId            *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DocumentId        *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
	DocumentName      *string `json:"documentName,omitempty" xml:"documentName,omitempty"`
	DocumentUrl       *string `json:"documentUrl,omitempty" xml:"documentUrl,omitempty"`
}

func (s CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) SetAiSheetTemplateId(v int64) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO {
	s.AiSheetTemplateId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) SetCorpId(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO {
	s.CorpId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) SetDocumentId(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO {
	s.DocumentId = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) SetDocumentName(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO {
	s.DocumentName = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO) SetDocumentUrl(v string) *CreateExportDeviceStatisticTaskResponseBodyExportStatisticTaskDTOAiSheetDocumentOpenDTO {
	s.DocumentUrl = &v
	return s
}

type CreateExportDeviceStatisticTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExportDeviceStatisticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExportDeviceStatisticTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExportDeviceStatisticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateExportDeviceStatisticTaskResponse) SetHeaders(v map[string]*string) *CreateExportDeviceStatisticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponse) SetStatusCode(v int32) *CreateExportDeviceStatisticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExportDeviceStatisticTaskResponse) SetBody(v *CreateExportDeviceStatisticTaskResponseBody) *CreateExportDeviceStatisticTaskResponse {
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
	// This parameter is required.
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExtractFacialFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// example:
	//
	// 165441111
	DeviceId     *int64   `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	ScopeDeptIds []*int64 `json:"scopeDeptIds,omitempty" xml:"scopeDeptIds,omitempty" type:"Repeated"`
	// example:
	//
	// user01
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// example:
	//
	// true
	AttendancePersonManage *bool `json:"attendancePersonManage,omitempty" xml:"attendancePersonManage,omitempty"`
	// example:
	//
	// true
	BluetoothPunchManage *bool `json:"bluetoothPunchManage,omitempty" xml:"bluetoothPunchManage,omitempty"`
	// example:
	//
	// true
	DeviceReset *bool `json:"deviceReset,omitempty" xml:"deviceReset,omitempty"`
	// example:
	//
	// true
	DeviceSettings *bool `json:"deviceSettings,omitempty" xml:"deviceSettings,omitempty"`
	// example:
	//
	// true
	FacePunchManage *bool `json:"facePunchManage,omitempty" xml:"facePunchManage,omitempty"`
	// example:
	//
	// true
	FingerPunchManage *bool `json:"fingerPunchManage,omitempty" xml:"fingerPunchManage,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceVideoConferenceBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type TextToImageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TextToImageHeaders) String() string {
	return tea.Prettify(s)
}

func (s TextToImageHeaders) GoString() string {
	return s.String()
}

func (s *TextToImageHeaders) SetCommonHeaders(v map[string]*string) *TextToImageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TextToImageHeaders) SetXAcsDingtalkAccessToken(v string) *TextToImageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TextToImageRequest struct {
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 1
	PictureNum *int64 `json:"pictureNum,omitempty" xml:"pictureNum,omitempty"`
	// example:
	//
	// 1024*1024
	PictureSize *string `json:"pictureSize,omitempty" xml:"pictureSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 帮我生成一个小猫在草坪上奔跑的图片
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s TextToImageRequest) String() string {
	return tea.Prettify(s)
}

func (s TextToImageRequest) GoString() string {
	return s.String()
}

func (s *TextToImageRequest) SetModelId(v string) *TextToImageRequest {
	s.ModelId = &v
	return s
}

func (s *TextToImageRequest) SetPictureNum(v int64) *TextToImageRequest {
	s.PictureNum = &v
	return s
}

func (s *TextToImageRequest) SetPictureSize(v string) *TextToImageRequest {
	s.PictureSize = &v
	return s
}

func (s *TextToImageRequest) SetQuery(v string) *TextToImageRequest {
	s.Query = &v
	return s
}

type TextToImageResponseBody struct {
	Result  *TextToImageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TextToImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextToImageResponseBody) GoString() string {
	return s.String()
}

func (s *TextToImageResponseBody) SetResult(v *TextToImageResponseBodyResult) *TextToImageResponseBody {
	s.Result = v
	return s
}

func (s *TextToImageResponseBody) SetSuccess(v bool) *TextToImageResponseBody {
	s.Success = &v
	return s
}

type TextToImageResponseBodyResult struct {
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId     *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s TextToImageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TextToImageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TextToImageResponseBodyResult) SetRequestId(v string) *TextToImageResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *TextToImageResponseBodyResult) SetTaskId(v string) *TextToImageResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *TextToImageResponseBodyResult) SetTaskStatus(v string) *TextToImageResponseBodyResult {
	s.TaskStatus = &v
	return s
}

type TextToImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextToImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextToImageResponse) String() string {
	return tea.Prettify(s)
}

func (s TextToImageResponse) GoString() string {
	return s.String()
}

func (s *TextToImageResponse) SetHeaders(v map[string]*string) *TextToImageResponse {
	s.Headers = v
	return s
}

func (s *TextToImageResponse) SetStatusCode(v int32) *TextToImageResponse {
	s.StatusCode = &v
	return s
}

func (s *TextToImageResponse) SetBody(v *TextToImageResponseBody) *TextToImageResponse {
	s.Body = v
	return s
}

type UpdateExportDeviceStatisticHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateExportDeviceStatisticHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateExportDeviceStatisticHeaders) GoString() string {
	return s.String()
}

func (s *UpdateExportDeviceStatisticHeaders) SetCommonHeaders(v map[string]*string) *UpdateExportDeviceStatisticHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateExportDeviceStatisticHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateExportDeviceStatisticHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateExportDeviceStatisticRequest struct {
	CreatorCorpId  *string                                      `json:"creatorCorpId,omitempty" xml:"creatorCorpId,omitempty"`
	CreatorUnionId *string                                      `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Records        []*UpdateExportDeviceStatisticRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	TaskId         *string                                      `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateExportDeviceStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExportDeviceStatisticRequest) GoString() string {
	return s.String()
}

func (s *UpdateExportDeviceStatisticRequest) SetCreatorCorpId(v string) *UpdateExportDeviceStatisticRequest {
	s.CreatorCorpId = &v
	return s
}

func (s *UpdateExportDeviceStatisticRequest) SetCreatorUnionId(v string) *UpdateExportDeviceStatisticRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *UpdateExportDeviceStatisticRequest) SetRecords(v []*UpdateExportDeviceStatisticRequestRecords) *UpdateExportDeviceStatisticRequest {
	s.Records = v
	return s
}

func (s *UpdateExportDeviceStatisticRequest) SetTaskId(v string) *UpdateExportDeviceStatisticRequest {
	s.TaskId = &v
	return s
}

type UpdateExportDeviceStatisticRequestRecords struct {
	Fields    map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty"`
	SheetName *string                `json:"sheetName,omitempty" xml:"sheetName,omitempty"`
}

func (s UpdateExportDeviceStatisticRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s UpdateExportDeviceStatisticRequestRecords) GoString() string {
	return s.String()
}

func (s *UpdateExportDeviceStatisticRequestRecords) SetFields(v map[string]interface{}) *UpdateExportDeviceStatisticRequestRecords {
	s.Fields = v
	return s
}

func (s *UpdateExportDeviceStatisticRequestRecords) SetSheetName(v string) *UpdateExportDeviceStatisticRequestRecords {
	s.SheetName = &v
	return s
}

type UpdateExportDeviceStatisticResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateExportDeviceStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExportDeviceStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExportDeviceStatisticResponseBody) SetSuccess(v bool) *UpdateExportDeviceStatisticResponseBody {
	s.Success = &v
	return s
}

type UpdateExportDeviceStatisticResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExportDeviceStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExportDeviceStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExportDeviceStatisticResponse) GoString() string {
	return s.String()
}

func (s *UpdateExportDeviceStatisticResponse) SetHeaders(v map[string]*string) *UpdateExportDeviceStatisticResponse {
	s.Headers = v
	return s
}

func (s *UpdateExportDeviceStatisticResponse) SetStatusCode(v int32) *UpdateExportDeviceStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExportDeviceStatisticResponse) SetBody(v *UpdateExportDeviceStatisticResponseBody) *UpdateExportDeviceStatisticResponse {
	s.Body = v
	return s
}

type VoiceCloneHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s VoiceCloneHeaders) String() string {
	return tea.Prettify(s)
}

func (s VoiceCloneHeaders) GoString() string {
	return s.String()
}

func (s *VoiceCloneHeaders) SetCommonHeaders(v map[string]*string) *VoiceCloneHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VoiceCloneHeaders) SetXAcsDingtalkAccessToken(v string) *VoiceCloneHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type VoiceCloneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 你好，我叫小智，是来自阿里云的超大规模语言模型。我是一个能够回答问题、创作文字，还能表达观点、撰写代码的全能型AI助手。如果您有任何问题或需要帮助，请随时告诉我，我会尽我所能为您提供帮助！
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// manager4224
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qhtestvoice-01
	VoiceId *string `json:"voiceId,omitempty" xml:"voiceId,omitempty"`
}

func (s VoiceCloneRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceCloneRequest) GoString() string {
	return s.String()
}

func (s *VoiceCloneRequest) SetText(v string) *VoiceCloneRequest {
	s.Text = &v
	return s
}

func (s *VoiceCloneRequest) SetUserId(v string) *VoiceCloneRequest {
	s.UserId = &v
	return s
}

func (s *VoiceCloneRequest) SetVoiceId(v string) *VoiceCloneRequest {
	s.VoiceId = &v
	return s
}

type VoiceCloneResponseBody struct {
	Result  *VoiceCloneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VoiceCloneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceCloneResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceCloneResponseBody) SetResult(v *VoiceCloneResponseBodyResult) *VoiceCloneResponseBody {
	s.Result = v
	return s
}

func (s *VoiceCloneResponseBody) SetSuccess(v bool) *VoiceCloneResponseBody {
	s.Success = &v
	return s
}

type VoiceCloneResponseBodyResult struct {
	// example:
	//
	// https://xxxx
	MediaUrl  *string `json:"mediaUrl,omitempty" xml:"mediaUrl,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s VoiceCloneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s VoiceCloneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *VoiceCloneResponseBodyResult) SetMediaUrl(v string) *VoiceCloneResponseBodyResult {
	s.MediaUrl = &v
	return s
}

func (s *VoiceCloneResponseBodyResult) SetRequestId(v string) *VoiceCloneResponseBodyResult {
	s.RequestId = &v
	return s
}

type VoiceCloneResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceCloneResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceCloneResponse) GoString() string {
	return s.String()
}

func (s *VoiceCloneResponse) SetHeaders(v map[string]*string) *VoiceCloneResponse {
	s.Headers = v
	return s
}

func (s *VoiceCloneResponse) SetStatusCode(v int32) *VoiceCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceCloneResponse) SetBody(v *VoiceCloneResponseBody) *VoiceCloneResponse {
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
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 添加硬件视频会议参会人
//
// @param request - AddDeviceVideoConferenceMembersRequest
//
// @param headers - AddDeviceVideoConferenceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceVideoConferenceMembersResponse
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

// Summary:
//
// 添加硬件视频会议参会人
//
// @param request - AddDeviceVideoConferenceMembersRequest
//
// @return AddDeviceVideoConferenceMembersResponse
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

// Summary:
//
// 创建硬件视频会议
//
// @param request - CreateDeviceVideoConferenceRequest
//
// @param headers - CreateDeviceVideoConferenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceVideoConferenceResponse
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

// Summary:
//
// 创建硬件视频会议
//
// @param request - CreateDeviceVideoConferenceRequest
//
// @return CreateDeviceVideoConferenceResponse
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

// Summary:
//
// 创建导出设备数据任务
//
// @param request - CreateExportDeviceStatisticTaskRequest
//
// @param headers - CreateExportDeviceStatisticTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportDeviceStatisticTaskResponse
func (client *Client) CreateExportDeviceStatisticTaskWithOptions(request *CreateExportDeviceStatisticTaskRequest, headers *CreateExportDeviceStatisticTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateExportDeviceStatisticTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiSheetTemplateId)) {
		body["aiSheetTemplateId"] = request.AiSheetTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorCorpId)) {
		body["creatorCorpId"] = request.CreatorCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["taskName"] = request.TaskName
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
		Action:      tea.String("CreateExportDeviceStatisticTask"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/statistic/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExportDeviceStatisticTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建导出设备数据任务
//
// @param request - CreateExportDeviceStatisticTaskRequest
//
// @return CreateExportDeviceStatisticTaskResponse
func (client *Client) CreateExportDeviceStatisticTask(request *CreateExportDeviceStatisticTaskRequest) (_result *CreateExportDeviceStatisticTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateExportDeviceStatisticTaskHeaders{}
	_result = &CreateExportDeviceStatisticTaskResponse{}
	_body, _err := client.CreateExportDeviceStatisticTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 基于企业员工照片为终端提取人脸识别特征
//
// @param request - ExtractFacialFeatureRequest
//
// @param headers - ExtractFacialFeatureHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtractFacialFeatureResponse
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

// Summary:
//
// 基于企业员工照片为终端提取人脸识别特征
//
// @param request - ExtractFacialFeatureRequest
//
// @return ExtractFacialFeatureResponse
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

// Summary:
//
// 踢出硬件视频会议参会人
//
// @param request - KickDeviceVideoConferenceMembersRequest
//
// @param headers - KickDeviceVideoConferenceMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickDeviceVideoConferenceMembersResponse
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

// Summary:
//
// 踢出硬件视频会议参会人
//
// @param request - KickDeviceVideoConferenceMembersRequest
//
// @return KickDeviceVideoConferenceMembersResponse
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

// Summary:
//
// 变更智能考勤机设备管理员
//
// @param request - MachineManagerUpdateRequest
//
// @param headers - MachineManagerUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MachineManagerUpdateResponse
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

// Summary:
//
// 变更智能考勤机设备管理员
//
// @param request - MachineManagerUpdateRequest
//
// @return MachineManagerUpdateResponse
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

// Summary:
//
// 变更智能考勤机员工
//
// @param request - MachineUsersUpdateRequest
//
// @param headers - MachineUsersUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MachineUsersUpdateResponse
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

// Summary:
//
// 变更智能考勤机员工
//
// @param request - MachineUsersUpdateRequest
//
// @return MachineUsersUpdateResponse
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

// Summary:
//
// 查询硬件视频会议预约信息
//
// @param headers - QueryDeviceVideoConferenceBookHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceVideoConferenceBookResponse
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

// Summary:
//
// 查询硬件视频会议预约信息
//
// @return QueryDeviceVideoConferenceBookResponse
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

// Summary:
//
// 文生图开放接口
//
// @param request - TextToImageRequest
//
// @param headers - TextToImageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextToImageResponse
func (client *Client) TextToImageWithOptions(request *TextToImageRequest, headers *TextToImageHeaders, runtime *util.RuntimeOptions) (_result *TextToImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PictureNum)) {
		body["pictureNum"] = request.PictureNum
	}

	if !tea.BoolValue(util.IsUnset(request.PictureSize)) {
		body["pictureSize"] = request.PictureSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
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
		Action:      tea.String("TextToImage"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/textToImages/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TextToImageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文生图开放接口
//
// @param request - TextToImageRequest
//
// @return TextToImageResponse
func (client *Client) TextToImage(request *TextToImageRequest) (_result *TextToImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TextToImageHeaders{}
	_result = &TextToImageResponse{}
	_body, _err := client.TextToImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新导出设备数据
//
// @param request - UpdateExportDeviceStatisticRequest
//
// @param headers - UpdateExportDeviceStatisticHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExportDeviceStatisticResponse
func (client *Client) UpdateExportDeviceStatisticWithOptions(request *UpdateExportDeviceStatisticRequest, headers *UpdateExportDeviceStatisticHeaders, runtime *util.RuntimeOptions) (_result *UpdateExportDeviceStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorCorpId)) {
		body["creatorCorpId"] = request.CreatorCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
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
		Action:      tea.String("UpdateExportDeviceStatistic"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/statistic"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExportDeviceStatisticResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新导出设备数据
//
// @param request - UpdateExportDeviceStatisticRequest
//
// @return UpdateExportDeviceStatisticResponse
func (client *Client) UpdateExportDeviceStatistic(request *UpdateExportDeviceStatisticRequest) (_result *UpdateExportDeviceStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateExportDeviceStatisticHeaders{}
	_result = &UpdateExportDeviceStatisticResponse{}
	_body, _err := client.UpdateExportDeviceStatisticWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 音频复刻
//
// @param request - VoiceCloneRequest
//
// @param headers - VoiceCloneHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceCloneResponse
func (client *Client) VoiceCloneWithOptions(request *VoiceCloneRequest, headers *VoiceCloneHeaders, runtime *util.RuntimeOptions) (_result *VoiceCloneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceId)) {
		body["voiceId"] = request.VoiceId
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
		Action:      tea.String("VoiceClone"),
		Version:     tea.String("smartDevice_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/smartDevice/voices/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceCloneResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 音频复刻
//
// @param request - VoiceCloneRequest
//
// @return VoiceCloneResponse
func (client *Client) VoiceClone(request *VoiceCloneRequest) (_result *VoiceCloneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &VoiceCloneHeaders{}
	_result = &VoiceCloneResponse{}
	_body, _err := client.VoiceCloneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
