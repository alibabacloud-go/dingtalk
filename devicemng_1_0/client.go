// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package devicemng_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchRegisterDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRegisterDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceHeaders) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceHeaders) SetCommonHeaders(v map[string]*string) *BatchRegisterDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRegisterDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRegisterDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRegisterDeviceRequest struct {
	// 设备列表
	DeviceList []*BatchRegisterDeviceRequestDeviceList `json:"deviceList,omitempty" xml:"deviceList,omitempty" type:"Repeated"`
	// 创建者userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchRegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequest) SetDeviceList(v []*BatchRegisterDeviceRequestDeviceList) *BatchRegisterDeviceRequest {
	s.DeviceList = v
	return s
}

func (s *BatchRegisterDeviceRequest) SetUserId(v string) *BatchRegisterDeviceRequest {
	s.UserId = &v
	return s
}

type BatchRegisterDeviceRequestDeviceList struct {
	// 协助者userId列表
	Collaborators *string `json:"collaborators,omitempty" xml:"collaborators,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 设备描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 设备标识
	DeviceKey *string `json:"deviceKey,omitempty" xml:"deviceKey,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 管理员userId列表
	Managers *string `json:"managers,omitempty" xml:"managers,omitempty"`
}

func (s BatchRegisterDeviceRequestDeviceList) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceRequestDeviceList) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceRequestDeviceList) SetCollaborators(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Collaborators = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDepartmentId(v int64) *BatchRegisterDeviceRequestDeviceList {
	s.DepartmentId = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDescription(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Description = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDeviceKey(v string) *BatchRegisterDeviceRequestDeviceList {
	s.DeviceKey = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetDeviceName(v string) *BatchRegisterDeviceRequestDeviceList {
	s.DeviceName = &v
	return s
}

func (s *BatchRegisterDeviceRequestDeviceList) SetManagers(v string) *BatchRegisterDeviceRequestDeviceList {
	s.Managers = &v
	return s
}

type BatchRegisterDeviceResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BatchRegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponseBody) SetResult(v string) *BatchRegisterDeviceResponseBody {
	s.Result = &v
	return s
}

type BatchRegisterDeviceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *BatchRegisterDeviceResponse) SetHeaders(v map[string]*string) *BatchRegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *BatchRegisterDeviceResponse) SetBody(v *BatchRegisterDeviceResponseBody) *BatchRegisterDeviceResponse {
	s.Body = v
	return s
}

type CreateChatRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateChatRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateChatRoomHeaders) GoString() string {
	return s.String()
}

func (s *CreateChatRoomHeaders) SetCommonHeaders(v map[string]*string) *CreateChatRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateChatRoomHeaders) SetXAcsDingtalkAccessToken(v string) *CreateChatRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateChatRoomRequest struct {
	ChatGroupName *string   `json:"chatGroupName,omitempty" xml:"chatGroupName,omitempty"`
	DeviceCodes   []*string `json:"deviceCodes,omitempty" xml:"deviceCodes,omitempty" type:"Repeated"`
	DeviceTypeId  *string   `json:"deviceTypeId,omitempty" xml:"deviceTypeId,omitempty"`
	RoleList      []*string `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
}

func (s CreateChatRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChatRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateChatRoomRequest) SetChatGroupName(v string) *CreateChatRoomRequest {
	s.ChatGroupName = &v
	return s
}

func (s *CreateChatRoomRequest) SetDeviceCodes(v []*string) *CreateChatRoomRequest {
	s.DeviceCodes = v
	return s
}

func (s *CreateChatRoomRequest) SetDeviceTypeId(v string) *CreateChatRoomRequest {
	s.DeviceTypeId = &v
	return s
}

func (s *CreateChatRoomRequest) SetRoleList(v []*string) *CreateChatRoomRequest {
	s.RoleList = v
	return s
}

type CreateChatRoomResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateChatRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChatRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatRoomResponseBody) SetResult(v string) *CreateChatRoomResponseBody {
	s.Result = &v
	return s
}

func (s *CreateChatRoomResponseBody) SetSuccess(v bool) *CreateChatRoomResponseBody {
	s.Success = &v
	return s
}

type CreateChatRoomResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateChatRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChatRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChatRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateChatRoomResponse) SetHeaders(v map[string]*string) *CreateChatRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateChatRoomResponse) SetBody(v *CreateChatRoomResponseBody) *CreateChatRoomResponse {
	s.Body = v
	return s
}

type CreateDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *CreateDepartmentHeaders) SetCommonHeaders(v map[string]*string) *CreateDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDepartmentRequest struct {
	// 认证信息
	AuthInfo *string `json:"authInfo,omitempty" xml:"authInfo,omitempty"`
	// 认证方式
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// 业务扩展
	BizExt *string `json:"bizExt,omitempty" xml:"bizExt,omitempty"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 部门类型
	DepartmentType *string `json:"departmentType,omitempty" xml:"departmentType,omitempty"`
	// 部门描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 业务系统地址
	SystemUrl *string `json:"systemUrl,omitempty" xml:"systemUrl,omitempty"`
	// 创建人工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) SetAuthInfo(v string) *CreateDepartmentRequest {
	s.AuthInfo = &v
	return s
}

func (s *CreateDepartmentRequest) SetAuthType(v string) *CreateDepartmentRequest {
	s.AuthType = &v
	return s
}

func (s *CreateDepartmentRequest) SetBizExt(v string) *CreateDepartmentRequest {
	s.BizExt = &v
	return s
}

func (s *CreateDepartmentRequest) SetDepartmentName(v string) *CreateDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *CreateDepartmentRequest) SetDepartmentType(v string) *CreateDepartmentRequest {
	s.DepartmentType = &v
	return s
}

func (s *CreateDepartmentRequest) SetDescription(v string) *CreateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *CreateDepartmentRequest) SetSystemUrl(v string) *CreateDepartmentRequest {
	s.SystemUrl = &v
	return s
}

func (s *CreateDepartmentRequest) SetUserId(v string) *CreateDepartmentRequest {
	s.UserId = &v
	return s
}

type CreateDepartmentResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) SetResult(v string) *CreateDepartmentResponseBody {
	s.Result = &v
	return s
}

type CreateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
	s.Body = v
	return s
}

type DeviceDingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeviceDingHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeviceDingHeaders) GoString() string {
	return s.String()
}

func (s *DeviceDingHeaders) SetCommonHeaders(v map[string]*string) *DeviceDingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceDingHeaders) SetXAcsDingtalkAccessToken(v string) *DeviceDingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeviceDingRequest struct {
	// 设备标识
	DeviceKey *string `json:"deviceKey,omitempty" xml:"deviceKey,omitempty"`
	// 消息体动态参数
	ParamsJson *string `json:"paramsJson,omitempty" xml:"paramsJson,omitempty"`
	// staffId列表
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
}

func (s DeviceDingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeviceDingRequest) GoString() string {
	return s.String()
}

func (s *DeviceDingRequest) SetDeviceKey(v string) *DeviceDingRequest {
	s.DeviceKey = &v
	return s
}

func (s *DeviceDingRequest) SetParamsJson(v string) *DeviceDingRequest {
	s.ParamsJson = &v
	return s
}

func (s *DeviceDingRequest) SetReceiverUserIdList(v []*string) *DeviceDingRequest {
	s.ReceiverUserIdList = v
	return s
}

type DeviceDingResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeviceDingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeviceDingResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceDingResponseBody) SetResult(v string) *DeviceDingResponseBody {
	s.Result = &v
	return s
}

type DeviceDingResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeviceDingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeviceDingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeviceDingResponse) GoString() string {
	return s.String()
}

func (s *DeviceDingResponse) SetHeaders(v map[string]*string) *DeviceDingResponse {
	s.Headers = v
	return s
}

func (s *DeviceDingResponse) SetBody(v *DeviceDingResponseBody) *DeviceDingResponse {
	s.Body = v
	return s
}

type ListActivateDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListActivateDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListActivateDevicesHeaders) GoString() string {
	return s.String()
}

func (s *ListActivateDevicesHeaders) SetCommonHeaders(v map[string]*string) *ListActivateDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListActivateDevicesHeaders) SetXAcsDingtalkAccessToken(v string) *ListActivateDevicesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListActivateDevicesRequest struct {
	// deviceCode
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	// deviceTypeId
	DeviceTypeId *string `json:"deviceTypeId,omitempty" xml:"deviceTypeId,omitempty"`
	// groupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// pageNo
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListActivateDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActivateDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListActivateDevicesRequest) SetDeviceCode(v string) *ListActivateDevicesRequest {
	s.DeviceCode = &v
	return s
}

func (s *ListActivateDevicesRequest) SetDeviceTypeId(v string) *ListActivateDevicesRequest {
	s.DeviceTypeId = &v
	return s
}

func (s *ListActivateDevicesRequest) SetGroupId(v string) *ListActivateDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *ListActivateDevicesRequest) SetPageNumber(v int32) *ListActivateDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListActivateDevicesRequest) SetPageSize(v int32) *ListActivateDevicesRequest {
	s.PageSize = &v
	return s
}

type ListActivateDevicesResponseBody struct {
	Result     []*ListActivateDevicesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success    *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListActivateDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActivateDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivateDevicesResponseBody) SetResult(v []*ListActivateDevicesResponseBodyResult) *ListActivateDevicesResponseBody {
	s.Result = v
	return s
}

func (s *ListActivateDevicesResponseBody) SetSuccess(v bool) *ListActivateDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListActivateDevicesResponseBody) SetTotalCount(v int64) *ListActivateDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListActivateDevicesResponseBodyResult struct {
	BizExt            *string `json:"bizExt,omitempty" xml:"bizExt,omitempty"`
	CorpId            *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceCallbackUrl *string `json:"deviceCallbackUrl,omitempty" xml:"deviceCallbackUrl,omitempty"`
	DeviceCode        *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl   *string `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName        *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	GroupUuid         *string `json:"groupUuid,omitempty" xml:"groupUuid,omitempty"`
	Icon              *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Introduction      *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	TypeUuid          *string `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	Uuid              *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ListActivateDevicesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListActivateDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListActivateDevicesResponseBodyResult) SetBizExt(v string) *ListActivateDevicesResponseBodyResult {
	s.BizExt = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetCorpId(v string) *ListActivateDevicesResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetDeviceCallbackUrl(v string) *ListActivateDevicesResponseBodyResult {
	s.DeviceCallbackUrl = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetDeviceCode(v string) *ListActivateDevicesResponseBodyResult {
	s.DeviceCode = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetDeviceDetailUrl(v string) *ListActivateDevicesResponseBodyResult {
	s.DeviceDetailUrl = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetDeviceName(v string) *ListActivateDevicesResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetGroupUuid(v string) *ListActivateDevicesResponseBodyResult {
	s.GroupUuid = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetIcon(v string) *ListActivateDevicesResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetIntroduction(v string) *ListActivateDevicesResponseBodyResult {
	s.Introduction = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetTypeUuid(v string) *ListActivateDevicesResponseBodyResult {
	s.TypeUuid = &v
	return s
}

func (s *ListActivateDevicesResponseBodyResult) SetUuid(v string) *ListActivateDevicesResponseBodyResult {
	s.Uuid = &v
	return s
}

type ListActivateDevicesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListActivateDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActivateDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActivateDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListActivateDevicesResponse) SetHeaders(v map[string]*string) *ListActivateDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListActivateDevicesResponse) SetBody(v *ListActivateDevicesResponseBody) *ListActivateDevicesResponse {
	s.Body = v
	return s
}

type RegisterAndActivateDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterAndActivateDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceHeaders) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceHeaders) SetCommonHeaders(v map[string]*string) *RegisterAndActivateDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterAndActivateDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterAndActivateDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterAndActivateDeviceRequest struct {
	DeviceCallbackUrl *string   `json:"deviceCallbackUrl,omitempty" xml:"deviceCallbackUrl,omitempty"`
	DeviceCode        *string   `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl   *string   `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName        *string   `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	Introduction      *string   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	RoleUuid          *string   `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	TypeUuid          *string   `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RegisterAndActivateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceRequest) SetDeviceCallbackUrl(v string) *RegisterAndActivateDeviceRequest {
	s.DeviceCallbackUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetDeviceCode(v string) *RegisterAndActivateDeviceRequest {
	s.DeviceCode = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetDeviceDetailUrl(v string) *RegisterAndActivateDeviceRequest {
	s.DeviceDetailUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetDeviceName(v string) *RegisterAndActivateDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetIntroduction(v string) *RegisterAndActivateDeviceRequest {
	s.Introduction = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetRoleUuid(v string) *RegisterAndActivateDeviceRequest {
	s.RoleUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetTypeUuid(v string) *RegisterAndActivateDeviceRequest {
	s.TypeUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceRequest) SetUserIds(v []*string) *RegisterAndActivateDeviceRequest {
	s.UserIds = v
	return s
}

type RegisterAndActivateDeviceResponseBody struct {
	Result *RegisterAndActivateDeviceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// Id of the request
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterAndActivateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceResponseBody) SetResult(v *RegisterAndActivateDeviceResponseBodyResult) *RegisterAndActivateDeviceResponseBody {
	s.Result = v
	return s
}

func (s *RegisterAndActivateDeviceResponseBody) SetSuccess(v bool) *RegisterAndActivateDeviceResponseBody {
	s.Success = &v
	return s
}

type RegisterAndActivateDeviceResponseBodyResult struct {
	CorpId          *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceCode      *string   `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl *string   `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName      *string   `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceUuid      *string   `json:"deviceUuid,omitempty" xml:"deviceUuid,omitempty"`
	Introduction    *string   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	RoleUuid        *string   `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	TypeUuid        *string   `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	UserIds         []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RegisterAndActivateDeviceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetCorpId(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetDeviceCode(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.DeviceCode = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetDeviceDetailUrl(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.DeviceDetailUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetDeviceName(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetDeviceUuid(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.DeviceUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetIntroduction(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.Introduction = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetRoleUuid(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.RoleUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetTypeUuid(v string) *RegisterAndActivateDeviceResponseBodyResult {
	s.TypeUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceResponseBodyResult) SetUserIds(v []*string) *RegisterAndActivateDeviceResponseBodyResult {
	s.UserIds = v
	return s
}

type RegisterAndActivateDeviceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterAndActivateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterAndActivateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceResponse) SetHeaders(v map[string]*string) *RegisterAndActivateDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterAndActivateDeviceResponse) SetBody(v *RegisterAndActivateDeviceResponseBody) *RegisterAndActivateDeviceResponse {
	s.Body = v
	return s
}

type RegisterAndActivateDeviceBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterAndActivateDeviceBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchHeaders) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchHeaders) SetCommonHeaders(v map[string]*string) *RegisterAndActivateDeviceBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterAndActivateDeviceBatchHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterAndActivateDeviceBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterAndActivateDeviceBatchRequest struct {
	RegisterAndActivateVOS []*RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS `json:"registerAndActivateVOS,omitempty" xml:"registerAndActivateVOS,omitempty" type:"Repeated"`
}

func (s RegisterAndActivateDeviceBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchRequest) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchRequest) SetRegisterAndActivateVOS(v []*RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) *RegisterAndActivateDeviceBatchRequest {
	s.RegisterAndActivateVOS = v
	return s
}

type RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS struct {
	DeviceCallbackUrl *string   `json:"deviceCallbackUrl,omitempty" xml:"deviceCallbackUrl,omitempty"`
	DeviceCode        *string   `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl   *string   `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName        *string   `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	GroupUuid         *string   `json:"groupUuid,omitempty" xml:"groupUuid,omitempty"`
	Introduction      *string   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	RoleUuid          *string   `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	TypeUuid          *string   `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetDeviceCallbackUrl(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.DeviceCallbackUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetDeviceCode(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.DeviceCode = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetDeviceDetailUrl(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.DeviceDetailUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetDeviceName(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.DeviceName = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetGroupUuid(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.GroupUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetIntroduction(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.Introduction = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetRoleUuid(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.RoleUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetTypeUuid(v string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.TypeUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS) SetUserIds(v []*string) *RegisterAndActivateDeviceBatchRequestRegisterAndActivateVOS {
	s.UserIds = v
	return s
}

type RegisterAndActivateDeviceBatchResponseBody struct {
	FailItems    []*RegisterAndActivateDeviceBatchResponseBodyFailItems    `json:"failItems,omitempty" xml:"failItems,omitempty" type:"Repeated"`
	Success      *bool                                                     `json:"success,omitempty" xml:"success,omitempty"`
	SuccessItems []*RegisterAndActivateDeviceBatchResponseBodySuccessItems `json:"successItems,omitempty" xml:"successItems,omitempty" type:"Repeated"`
}

func (s RegisterAndActivateDeviceBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponseBody) SetFailItems(v []*RegisterAndActivateDeviceBatchResponseBodyFailItems) *RegisterAndActivateDeviceBatchResponseBody {
	s.FailItems = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBody) SetSuccess(v bool) *RegisterAndActivateDeviceBatchResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBody) SetSuccessItems(v []*RegisterAndActivateDeviceBatchResponseBodySuccessItems) *RegisterAndActivateDeviceBatchResponseBody {
	s.SuccessItems = v
	return s
}

type RegisterAndActivateDeviceBatchResponseBodyFailItems struct {
	ErrorCode *string                                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterAndActivateDeviceBatchResponseBodyFailItems) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponseBodyFailItems) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItems) SetErrorCode(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItems {
	s.ErrorCode = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItems) SetErrorMsg(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItems {
	s.ErrorMsg = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItems) SetResult(v *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) *RegisterAndActivateDeviceBatchResponseBodyFailItems {
	s.Result = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItems) SetSuccess(v bool) *RegisterAndActivateDeviceBatchResponseBodyFailItems {
	s.Success = &v
	return s
}

type RegisterAndActivateDeviceBatchResponseBodyFailItemsResult struct {
	CorpId            *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceCallbackUrl *string   `json:"deviceCallbackUrl,omitempty" xml:"deviceCallbackUrl,omitempty"`
	DeviceCode        *string   `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl   *string   `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName        *string   `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	GroupUuid         *string   `json:"groupUuid,omitempty" xml:"groupUuid,omitempty"`
	Icon              *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Introduction      *string   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	RoleUuid          *string   `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	Status            *int64    `json:"status,omitempty" xml:"status,omitempty"`
	TypeUuid          *string   `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	Uuid              *string   `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetCorpId(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.CorpId = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetDeviceCallbackUrl(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.DeviceCallbackUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetDeviceCode(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.DeviceCode = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetDeviceDetailUrl(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.DeviceDetailUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetDeviceName(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.DeviceName = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetGroupUuid(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.GroupUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetIcon(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.Icon = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetIntroduction(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.Introduction = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetRoleUuid(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.RoleUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetStatus(v int64) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.Status = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetTypeUuid(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.TypeUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetUserIds(v []*string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.UserIds = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult) SetUuid(v string) *RegisterAndActivateDeviceBatchResponseBodyFailItemsResult {
	s.Uuid = &v
	return s
}

type RegisterAndActivateDeviceBatchResponseBodySuccessItems struct {
	ErrorCode *string                                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                                       `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success   *bool                                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterAndActivateDeviceBatchResponseBodySuccessItems) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponseBodySuccessItems) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItems) SetErrorCode(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItems {
	s.ErrorCode = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItems) SetErrorMsg(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItems {
	s.ErrorMsg = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItems) SetResult(v *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) *RegisterAndActivateDeviceBatchResponseBodySuccessItems {
	s.Result = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItems) SetSuccess(v bool) *RegisterAndActivateDeviceBatchResponseBodySuccessItems {
	s.Success = &v
	return s
}

type RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult struct {
	CorpId            *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeviceCallbackUrl *string   `json:"deviceCallbackUrl,omitempty" xml:"deviceCallbackUrl,omitempty"`
	DeviceCode        *string   `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceDetailUrl   *string   `json:"deviceDetailUrl,omitempty" xml:"deviceDetailUrl,omitempty"`
	DeviceName        *string   `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	GroupUuid         *string   `json:"groupUuid,omitempty" xml:"groupUuid,omitempty"`
	Icon              *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	Introduction      *string   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	RoleUuid          *string   `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	Status            *int64    `json:"status,omitempty" xml:"status,omitempty"`
	TypeUuid          *string   `json:"typeUuid,omitempty" xml:"typeUuid,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	Uuid              *string   `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetCorpId(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.CorpId = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetDeviceCallbackUrl(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.DeviceCallbackUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetDeviceCode(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.DeviceCode = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetDeviceDetailUrl(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.DeviceDetailUrl = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetDeviceName(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.DeviceName = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetGroupUuid(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.GroupUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetIcon(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.Icon = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetIntroduction(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.Introduction = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetRoleUuid(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.RoleUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetStatus(v int64) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.Status = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetTypeUuid(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.TypeUuid = &v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetUserIds(v []*string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.UserIds = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult) SetUuid(v string) *RegisterAndActivateDeviceBatchResponseBodySuccessItemsResult {
	s.Uuid = &v
	return s
}

type RegisterAndActivateDeviceBatchResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterAndActivateDeviceBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterAndActivateDeviceBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterAndActivateDeviceBatchResponse) GoString() string {
	return s.String()
}

func (s *RegisterAndActivateDeviceBatchResponse) SetHeaders(v map[string]*string) *RegisterAndActivateDeviceBatchResponse {
	s.Headers = v
	return s
}

func (s *RegisterAndActivateDeviceBatchResponse) SetBody(v *RegisterAndActivateDeviceBatchResponseBody) *RegisterAndActivateDeviceBatchResponse {
	s.Body = v
	return s
}

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
	// 协助者userId列表
	Collaborators *string `json:"collaborators,omitempty" xml:"collaborators,omitempty"`
	// 部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 设备描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 设备标识
	DeviceKey *string `json:"deviceKey,omitempty" xml:"deviceKey,omitempty"`
	// 设备名称
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	// 管理员userId列表
	Managers *string `json:"managers,omitempty" xml:"managers,omitempty"`
	// 创建者userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetCollaborators(v string) *RegisterDeviceRequest {
	s.Collaborators = &v
	return s
}

func (s *RegisterDeviceRequest) SetDepartmentId(v int64) *RegisterDeviceRequest {
	s.DepartmentId = &v
	return s
}

func (s *RegisterDeviceRequest) SetDescription(v string) *RegisterDeviceRequest {
	s.Description = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceKey(v string) *RegisterDeviceRequest {
	s.DeviceKey = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceName(v string) *RegisterDeviceRequest {
	s.DeviceName = &v
	return s
}

func (s *RegisterDeviceRequest) SetManagers(v string) *RegisterDeviceRequest {
	s.Managers = &v
	return s
}

func (s *RegisterDeviceRequest) SetUserId(v string) *RegisterDeviceRequest {
	s.UserId = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetResult(v string) *RegisterDeviceResponseBody {
	s.Result = &v
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

type UploadEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadEventHeaders) GoString() string {
	return s.String()
}

func (s *UploadEventHeaders) SetCommonHeaders(v map[string]*string) *UploadEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadEventHeaders) SetXAcsDingtalkAccessToken(v string) *UploadEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadEventRequest struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceUuid *string `json:"deviceUuid,omitempty" xml:"deviceUuid,omitempty"`
	EventTime  *string `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	EventType  *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	Level      *string `json:"level,omitempty" xml:"level,omitempty"`
}

func (s UploadEventRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadEventRequest) GoString() string {
	return s.String()
}

func (s *UploadEventRequest) SetContent(v string) *UploadEventRequest {
	s.Content = &v
	return s
}

func (s *UploadEventRequest) SetDeviceCode(v string) *UploadEventRequest {
	s.DeviceCode = &v
	return s
}

func (s *UploadEventRequest) SetDeviceUuid(v string) *UploadEventRequest {
	s.DeviceUuid = &v
	return s
}

func (s *UploadEventRequest) SetEventTime(v string) *UploadEventRequest {
	s.EventTime = &v
	return s
}

func (s *UploadEventRequest) SetEventType(v string) *UploadEventRequest {
	s.EventType = &v
	return s
}

func (s *UploadEventRequest) SetLevel(v string) *UploadEventRequest {
	s.Level = &v
	return s
}

type UploadEventResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadEventResponseBody) GoString() string {
	return s.String()
}

func (s *UploadEventResponseBody) SetResult(v string) *UploadEventResponseBody {
	s.Result = &v
	return s
}

func (s *UploadEventResponseBody) SetSuccess(v bool) *UploadEventResponseBody {
	s.Success = &v
	return s
}

type UploadEventResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadEventResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadEventResponse) GoString() string {
	return s.String()
}

func (s *UploadEventResponse) SetHeaders(v map[string]*string) *UploadEventResponse {
	s.Headers = v
	return s
}

func (s *UploadEventResponse) SetBody(v *UploadEventResponseBody) *UploadEventResponse {
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

func (client *Client) BatchRegisterDevice(request *BatchRegisterDeviceRequest) (_result *BatchRegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRegisterDeviceHeaders{}
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.BatchRegisterDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRegisterDeviceWithOptions(request *BatchRegisterDeviceRequest, headers *BatchRegisterDeviceHeaders, runtime *util.RuntimeOptions) (_result *BatchRegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceList)) {
		body["deviceList"] = request.DeviceList
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
	_result = &BatchRegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRegisterDevice"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/devices/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateChatRoom(request *CreateChatRoomRequest) (_result *CreateChatRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateChatRoomHeaders{}
	_result = &CreateChatRoomResponse{}
	_body, _err := client.CreateChatRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateChatRoomWithOptions(request *CreateChatRoomRequest, headers *CreateChatRoomHeaders, runtime *util.RuntimeOptions) (_result *CreateChatRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatGroupName)) {
		body["chatGroupName"] = request.ChatGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCodes)) {
		body["deviceCodes"] = request.DeviceCodes
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTypeId)) {
		body["deviceTypeId"] = request.DeviceTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleList)) {
		body["roleList"] = request.RoleList
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
	_result = &CreateChatRoomResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateChatRoom"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRoom"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDepartment(request *CreateDepartmentRequest) (_result *CreateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDepartmentHeaders{}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CreateDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDepartmentWithOptions(request *CreateDepartmentRequest, headers *CreateDepartmentHeaders, runtime *util.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthInfo)) {
		body["authInfo"] = request.AuthInfo
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		body["authType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.BizExt)) {
		body["bizExt"] = request.BizExt
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		body["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentType)) {
		body["departmentType"] = request.DepartmentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SystemUrl)) {
		body["systemUrl"] = request.SystemUrl
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
	_result = &CreateDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDepartment"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeviceDing(request *DeviceDingRequest) (_result *DeviceDingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeviceDingHeaders{}
	_result = &DeviceDingResponse{}
	_body, _err := client.DeviceDingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeviceDingWithOptions(request *DeviceDingRequest, headers *DeviceDingHeaders, runtime *util.RuntimeOptions) (_result *DeviceDingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceKey)) {
		body["deviceKey"] = request.DeviceKey
	}

	if !tea.BoolValue(util.IsUnset(request.ParamsJson)) {
		body["paramsJson"] = request.ParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
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
	_result = &DeviceDingResponse{}
	_body, _err := client.DoROARequest(tea.String("DeviceDing"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/ding"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActivateDevices(request *ListActivateDevicesRequest) (_result *ListActivateDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListActivateDevicesHeaders{}
	_result = &ListActivateDevicesResponse{}
	_body, _err := client.ListActivateDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActivateDevicesWithOptions(request *ListActivateDevicesRequest, headers *ListActivateDevicesHeaders, runtime *util.RuntimeOptions) (_result *ListActivateDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		query["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTypeId)) {
		query["deviceTypeId"] = request.DeviceTypeId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
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
	_result = &ListActivateDevicesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListActivateDevices"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/devices/activations/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterAndActivateDevice(request *RegisterAndActivateDeviceRequest) (_result *RegisterAndActivateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterAndActivateDeviceHeaders{}
	_result = &RegisterAndActivateDeviceResponse{}
	_body, _err := client.RegisterAndActivateDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterAndActivateDeviceWithOptions(request *RegisterAndActivateDeviceRequest, headers *RegisterAndActivateDeviceHeaders, runtime *util.RuntimeOptions) (_result *RegisterAndActivateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCallbackUrl)) {
		body["deviceCallbackUrl"] = request.DeviceCallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		body["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceDetailUrl)) {
		body["deviceDetailUrl"] = request.DeviceDetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.RoleUuid)) {
		body["roleUuid"] = request.RoleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TypeUuid)) {
		body["typeUuid"] = request.TypeUuid
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
	_result = &RegisterAndActivateDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterAndActivateDevice"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/devices/registerAndActivate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterAndActivateDeviceBatch(request *RegisterAndActivateDeviceBatchRequest) (_result *RegisterAndActivateDeviceBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterAndActivateDeviceBatchHeaders{}
	_result = &RegisterAndActivateDeviceBatchResponse{}
	_body, _err := client.RegisterAndActivateDeviceBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterAndActivateDeviceBatchWithOptions(request *RegisterAndActivateDeviceBatchRequest, headers *RegisterAndActivateDeviceBatchHeaders, runtime *util.RuntimeOptions) (_result *RegisterAndActivateDeviceBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegisterAndActivateVOS)) {
		body["registerAndActivateVOS"] = request.RegisterAndActivateVOS
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
	_result = &RegisterAndActivateDeviceBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterAndActivateDeviceBatch"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/devices/registrationActivations/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	if !tea.BoolValue(util.IsUnset(request.Collaborators)) {
		body["collaborators"] = request.Collaborators
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceKey)) {
		body["deviceKey"] = request.DeviceKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		body["deviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.Managers)) {
		body["managers"] = request.Managers
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
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterDevice"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/devices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadEvent(request *UploadEventRequest) (_result *UploadEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadEventHeaders{}
	_result = &UploadEventResponse{}
	_body, _err := client.UploadEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadEventWithOptions(request *UploadEventRequest, headers *UploadEventHeaders, runtime *util.RuntimeOptions) (_result *UploadEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		body["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUuid)) {
		body["deviceUuid"] = request.DeviceUuid
	}

	if !tea.BoolValue(util.IsUnset(request.EventTime)) {
		body["eventTime"] = request.EventTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["level"] = request.Level
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
	_result = &UploadEventResponse{}
	_body, _err := client.DoROARequest(tea.String("UploadEvent"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/suppliers/events/upload"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
