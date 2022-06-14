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
	OwnerUserId   *string   `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
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

func (s *CreateChatRoomRequest) SetOwnerUserId(v string) *CreateChatRoomRequest {
	s.OwnerUserId = &v
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

type CreateDeviceChatRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDeviceChatRoomHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceChatRoomHeaders) GoString() string {
	return s.String()
}

func (s *CreateDeviceChatRoomHeaders) SetCommonHeaders(v map[string]*string) *CreateDeviceChatRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeviceChatRoomHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDeviceChatRoomHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDeviceChatRoomRequest struct {
	// 场景群类型，不填，为默认普通设备群，示例值--维修群：REPAIR_GROUP，点检群:INSPECT_GROUP,保养群：MAINTAIN_GROUP。
	ChatType *string `json:"chatType,omitempty" xml:"chatType,omitempty"`
	// 设备编码，客户侧生成的设备标识，能够唯一标识一个设备，该字段与deviceUuids字段需要二选一，并且不能都填充。
	DeviceCodes []*string `json:"deviceCodes,omitempty" xml:"deviceCodes,omitempty" type:"Repeated"`
	// 设备唯一标识，钉钉侧生成的设备标识，能够唯一标识一个设备，该字段与deviceCodes字段需要二选一，并且不能都填充。
	DeviceUuids []*string `json:"deviceUuids,omitempty" xml:"deviceUuids,omitempty" type:"Repeated"`
	// 群主钉钉账户。
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// 设备场景群名称。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 需要被拉入群聊的钉钉账号，可以传多个，通过英文逗号分隔。
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateDeviceChatRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceChatRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceChatRoomRequest) SetChatType(v string) *CreateDeviceChatRoomRequest {
	s.ChatType = &v
	return s
}

func (s *CreateDeviceChatRoomRequest) SetDeviceCodes(v []*string) *CreateDeviceChatRoomRequest {
	s.DeviceCodes = v
	return s
}

func (s *CreateDeviceChatRoomRequest) SetDeviceUuids(v []*string) *CreateDeviceChatRoomRequest {
	s.DeviceUuids = v
	return s
}

func (s *CreateDeviceChatRoomRequest) SetOwnerUserId(v string) *CreateDeviceChatRoomRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateDeviceChatRoomRequest) SetTitle(v string) *CreateDeviceChatRoomRequest {
	s.Title = &v
	return s
}

func (s *CreateDeviceChatRoomRequest) SetUserIds(v []*string) *CreateDeviceChatRoomRequest {
	s.UserIds = v
	return s
}

type CreateDeviceChatRoomResponseBody struct {
	Result  *CreateDeviceChatRoomResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeviceChatRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceChatRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceChatRoomResponseBody) SetResult(v *CreateDeviceChatRoomResponseBodyResult) *CreateDeviceChatRoomResponseBody {
	s.Result = v
	return s
}

func (s *CreateDeviceChatRoomResponseBody) SetSuccess(v bool) *CreateDeviceChatRoomResponseBody {
	s.Success = &v
	return s
}

type CreateDeviceChatRoomResponseBodyResult struct {
	ChatId             *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	EncodedCid         *string `json:"encodedCid,omitempty" xml:"encodedCid,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateDeviceChatRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceChatRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDeviceChatRoomResponseBodyResult) SetChatId(v string) *CreateDeviceChatRoomResponseBodyResult {
	s.ChatId = &v
	return s
}

func (s *CreateDeviceChatRoomResponseBodyResult) SetEncodedCid(v string) *CreateDeviceChatRoomResponseBodyResult {
	s.EncodedCid = &v
	return s
}

func (s *CreateDeviceChatRoomResponseBodyResult) SetOpenConversationId(v string) *CreateDeviceChatRoomResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

type CreateDeviceChatRoomResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeviceChatRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeviceChatRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeviceChatRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceChatRoomResponse) SetHeaders(v map[string]*string) *CreateDeviceChatRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceChatRoomResponse) SetBody(v *CreateDeviceChatRoomResponseBody) *CreateDeviceChatRoomResponse {
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

type DissolveGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DissolveGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s DissolveGroupHeaders) GoString() string {
	return s.String()
}

func (s *DissolveGroupHeaders) SetCommonHeaders(v map[string]*string) *DissolveGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DissolveGroupHeaders) SetXAcsDingtalkAccessToken(v string) *DissolveGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DissolveGroupRequest struct {
	// 场景群唯一标识，调用创建场景群接口时，会返回该标识。
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s DissolveGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DissolveGroupRequest) GoString() string {
	return s.String()
}

func (s *DissolveGroupRequest) SetOpenConversationId(v string) *DissolveGroupRequest {
	s.OpenConversationId = &v
	return s
}

type DissolveGroupResponseBody struct {
	// 接口处理返回结果。
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 接口处理是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DissolveGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissolveGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DissolveGroupResponseBody) SetResult(v string) *DissolveGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DissolveGroupResponseBody) SetSuccess(v bool) *DissolveGroupResponseBody {
	s.Success = &v
	return s
}

type DissolveGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissolveGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissolveGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DissolveGroupResponse) GoString() string {
	return s.String()
}

func (s *DissolveGroupResponse) SetHeaders(v map[string]*string) *DissolveGroupResponse {
	s.Headers = v
	return s
}

func (s *DissolveGroupResponse) SetBody(v *DissolveGroupResponseBody) *DissolveGroupResponse {
	s.Body = v
	return s
}

type EditDeviceAdminHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditDeviceAdminHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditDeviceAdminHeaders) GoString() string {
	return s.String()
}

func (s *EditDeviceAdminHeaders) SetCommonHeaders(v map[string]*string) *EditDeviceAdminHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditDeviceAdminHeaders) SetXAcsDingtalkAccessToken(v string) *EditDeviceAdminHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditDeviceAdminRequest struct {
	// 需要处理的设备编号。客户侧生成的设备标识，能够唯一标识一个设备，该字段与uuid字段需要二选一，并且不能都填充。
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	// 角色唯一标识
	RoleUuid *string `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	// 需要编辑的角色唯一标识，非必填，不传默认为管理员。
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// 设备唯一标识，钉钉侧生成的设备标识，能够唯一标识一个设备，该字段与deviceCode字段需要二选一，并且不能都填充。
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s EditDeviceAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s EditDeviceAdminRequest) GoString() string {
	return s.String()
}

func (s *EditDeviceAdminRequest) SetDeviceCode(v string) *EditDeviceAdminRequest {
	s.DeviceCode = &v
	return s
}

func (s *EditDeviceAdminRequest) SetRoleUuid(v string) *EditDeviceAdminRequest {
	s.RoleUuid = &v
	return s
}

func (s *EditDeviceAdminRequest) SetUserIds(v []*string) *EditDeviceAdminRequest {
	s.UserIds = v
	return s
}

func (s *EditDeviceAdminRequest) SetUuid(v string) *EditDeviceAdminRequest {
	s.Uuid = &v
	return s
}

type EditDeviceAdminResponseBody struct {
	// 接口处理返回结果。
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 接口处理是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EditDeviceAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditDeviceAdminResponseBody) GoString() string {
	return s.String()
}

func (s *EditDeviceAdminResponseBody) SetResult(v string) *EditDeviceAdminResponseBody {
	s.Result = &v
	return s
}

func (s *EditDeviceAdminResponseBody) SetSuccess(v bool) *EditDeviceAdminResponseBody {
	s.Success = &v
	return s
}

type EditDeviceAdminResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditDeviceAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditDeviceAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s EditDeviceAdminResponse) GoString() string {
	return s.String()
}

func (s *EditDeviceAdminResponse) SetHeaders(v map[string]*string) *EditDeviceAdminResponse {
	s.Headers = v
	return s
}

func (s *EditDeviceAdminResponse) SetBody(v *EditDeviceAdminResponseBody) *EditDeviceAdminResponse {
	s.Body = v
	return s
}

type GetDeviceGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeviceGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeviceGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeviceGroupInfoRequest struct {
	// 开放群唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetDeviceGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoRequest) SetOpenConversationId(v string) *GetDeviceGroupInfoRequest {
	s.OpenConversationId = &v
	return s
}

type GetDeviceGroupInfoResponseBody struct {
	Result  *GetDeviceGroupInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDeviceGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoResponseBody) SetResult(v *GetDeviceGroupInfoResponseBodyResult) *GetDeviceGroupInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceGroupInfoResponseBody) SetSuccess(v bool) *GetDeviceGroupInfoResponseBody {
	s.Success = &v
	return s
}

type GetDeviceGroupInfoResponseBodyResult struct {
	Devices       []*GetDeviceGroupInfoResponseBodyResultDevices `json:"devices,omitempty" xml:"devices,omitempty" type:"Repeated"`
	OwnerUser     *string                                        `json:"ownerUser,omitempty" xml:"ownerUser,omitempty"`
	SubAdminUsers []*string                                      `json:"subAdminUsers,omitempty" xml:"subAdminUsers,omitempty" type:"Repeated"`
	Title         *string                                        `json:"title,omitempty" xml:"title,omitempty"`
	Users         map[string]*string                             `json:"users,omitempty" xml:"users,omitempty"`
}

func (s GetDeviceGroupInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoResponseBodyResult) SetDevices(v []*GetDeviceGroupInfoResponseBodyResultDevices) *GetDeviceGroupInfoResponseBodyResult {
	s.Devices = v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResult) SetOwnerUser(v string) *GetDeviceGroupInfoResponseBodyResult {
	s.OwnerUser = &v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResult) SetSubAdminUsers(v []*string) *GetDeviceGroupInfoResponseBodyResult {
	s.SubAdminUsers = v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResult) SetTitle(v string) *GetDeviceGroupInfoResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResult) SetUsers(v map[string]*string) *GetDeviceGroupInfoResponseBodyResult {
	s.Users = v
	return s
}

type GetDeviceGroupInfoResponseBodyResultDevices struct {
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
	DeviceUuid *string `json:"deviceUuid,omitempty" xml:"deviceUuid,omitempty"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s GetDeviceGroupInfoResponseBodyResultDevices) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoResponseBodyResultDevices) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoResponseBodyResultDevices) SetDeviceCode(v string) *GetDeviceGroupInfoResponseBodyResultDevices {
	s.DeviceCode = &v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResultDevices) SetDeviceName(v string) *GetDeviceGroupInfoResponseBodyResultDevices {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResultDevices) SetDeviceUuid(v string) *GetDeviceGroupInfoResponseBodyResultDevices {
	s.DeviceUuid = &v
	return s
}

func (s *GetDeviceGroupInfoResponseBodyResultDevices) SetUuid(v string) *GetDeviceGroupInfoResponseBodyResultDevices {
	s.Uuid = &v
	return s
}

type GetDeviceGroupInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupInfoResponse) SetHeaders(v map[string]*string) *GetDeviceGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceGroupInfoResponse) SetBody(v *GetDeviceGroupInfoResponseBody) *GetDeviceGroupInfoResponse {
	s.Body = v
	return s
}

type GetWholeDeviceGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetWholeDeviceGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetWholeDeviceGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetWholeDeviceGroupHeaders) SetCommonHeaders(v map[string]*string) *GetWholeDeviceGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWholeDeviceGroupHeaders) SetXAcsDingtalkAccessToken(v string) *GetWholeDeviceGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetWholeDeviceGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWholeDeviceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWholeDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetWholeDeviceGroupResponseBody) SetResult(v string) *GetWholeDeviceGroupResponseBody {
	s.Result = &v
	return s
}

func (s *GetWholeDeviceGroupResponseBody) SetSuccess(v bool) *GetWholeDeviceGroupResponseBody {
	s.Success = &v
	return s
}

type GetWholeDeviceGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWholeDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWholeDeviceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWholeDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetWholeDeviceGroupResponse) SetHeaders(v map[string]*string) *GetWholeDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetWholeDeviceGroupResponse) SetBody(v *GetWholeDeviceGroupResponseBody) *GetWholeDeviceGroupResponse {
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

type PullDeviceToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullDeviceToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullDeviceToGroupHeaders) GoString() string {
	return s.String()
}

func (s *PullDeviceToGroupHeaders) SetCommonHeaders(v map[string]*string) *PullDeviceToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullDeviceToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *PullDeviceToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullDeviceToGroupRequest struct {
	// 设备业务标识
	DeviceCodes []*string `json:"deviceCodes,omitempty" xml:"deviceCodes,omitempty" type:"Repeated"`
	// 设备uuid，系统唯一标识
	DeviceUuids []*string `json:"deviceUuids,omitempty" xml:"deviceUuids,omitempty" type:"Repeated"`
	// 群id，群的唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 操作人userId
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s PullDeviceToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDeviceToGroupRequest) GoString() string {
	return s.String()
}

func (s *PullDeviceToGroupRequest) SetDeviceCodes(v []*string) *PullDeviceToGroupRequest {
	s.DeviceCodes = v
	return s
}

func (s *PullDeviceToGroupRequest) SetDeviceUuids(v []*string) *PullDeviceToGroupRequest {
	s.DeviceUuids = v
	return s
}

func (s *PullDeviceToGroupRequest) SetOpenConversationId(v string) *PullDeviceToGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PullDeviceToGroupRequest) SetOperator(v string) *PullDeviceToGroupRequest {
	s.Operator = &v
	return s
}

type PullDeviceToGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PullDeviceToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullDeviceToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *PullDeviceToGroupResponseBody) SetResult(v string) *PullDeviceToGroupResponseBody {
	s.Result = &v
	return s
}

func (s *PullDeviceToGroupResponseBody) SetSuccess(v bool) *PullDeviceToGroupResponseBody {
	s.Success = &v
	return s
}

type PullDeviceToGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullDeviceToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullDeviceToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDeviceToGroupResponse) GoString() string {
	return s.String()
}

func (s *PullDeviceToGroupResponse) SetHeaders(v map[string]*string) *PullDeviceToGroupResponse {
	s.Headers = v
	return s
}

func (s *PullDeviceToGroupResponse) SetBody(v *PullDeviceToGroupResponseBody) *PullDeviceToGroupResponse {
	s.Body = v
	return s
}

type PullUserToGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullUserToGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullUserToGroupHeaders) GoString() string {
	return s.String()
}

func (s *PullUserToGroupHeaders) SetCommonHeaders(v map[string]*string) *PullUserToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullUserToGroupHeaders) SetXAcsDingtalkAccessToken(v string) *PullUserToGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullUserToGroupRequest struct {
	// 开放群唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 入群用户唯一标识列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s PullUserToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s PullUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *PullUserToGroupRequest) SetOpenConversationId(v string) *PullUserToGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PullUserToGroupRequest) SetUserIds(v []*string) *PullUserToGroupRequest {
	s.UserIds = v
	return s
}

type PullUserToGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PullUserToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullUserToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *PullUserToGroupResponseBody) SetResult(v string) *PullUserToGroupResponseBody {
	s.Result = &v
	return s
}

func (s *PullUserToGroupResponseBody) SetSuccess(v bool) *PullUserToGroupResponseBody {
	s.Success = &v
	return s
}

type PullUserToGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullUserToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s PullUserToGroupResponse) GoString() string {
	return s.String()
}

func (s *PullUserToGroupResponse) SetHeaders(v map[string]*string) *PullUserToGroupResponse {
	s.Headers = v
	return s
}

func (s *PullUserToGroupResponse) SetBody(v *PullUserToGroupResponseBody) *PullUserToGroupResponse {
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

type RemoveDeviceFromGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveDeviceFromGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *RemoveDeviceFromGroupHeaders) SetCommonHeaders(v map[string]*string) *RemoveDeviceFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveDeviceFromGroupHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveDeviceFromGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveDeviceFromGroupRequest struct {
	// 设备编号列表（与设备唯一标识列表二选一）
	DeviceCodes []*string `json:"deviceCodes,omitempty" xml:"deviceCodes,omitempty" type:"Repeated"`
	// 设备唯一标识列表（与设备编码列表二选一）
	DeviceUuids []*string `json:"deviceUuids,omitempty" xml:"deviceUuids,omitempty" type:"Repeated"`
	// 开放群唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 操作人唯一标识
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s RemoveDeviceFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveDeviceFromGroupRequest) SetDeviceCodes(v []*string) *RemoveDeviceFromGroupRequest {
	s.DeviceCodes = v
	return s
}

func (s *RemoveDeviceFromGroupRequest) SetDeviceUuids(v []*string) *RemoveDeviceFromGroupRequest {
	s.DeviceUuids = v
	return s
}

func (s *RemoveDeviceFromGroupRequest) SetOpenConversationId(v string) *RemoveDeviceFromGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RemoveDeviceFromGroupRequest) SetOperator(v string) *RemoveDeviceFromGroupRequest {
	s.Operator = &v
	return s
}

type RemoveDeviceFromGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveDeviceFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDeviceFromGroupResponseBody) SetResult(v string) *RemoveDeviceFromGroupResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveDeviceFromGroupResponseBody) SetSuccess(v bool) *RemoveDeviceFromGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveDeviceFromGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDeviceFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDeviceFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDeviceFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveDeviceFromGroupResponse) SetHeaders(v map[string]*string) *RemoveDeviceFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveDeviceFromGroupResponse) SetBody(v *RemoveDeviceFromGroupResponseBody) *RemoveDeviceFromGroupResponse {
	s.Body = v
	return s
}

type RemoveUserFromGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveUserFromGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupHeaders) SetCommonHeaders(v map[string]*string) *RemoveUserFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveUserFromGroupHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveUserFromGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveUserFromGroupRequest struct {
	// 开放群唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 用户唯一标识列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) SetOpenConversationId(v string) *RemoveUserFromGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserIds(v []*string) *RemoveUserFromGroupRequest {
	s.UserIds = v
	return s
}

type RemoveUserFromGroupResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveUserFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponseBody) SetResult(v string) *RemoveUserFromGroupResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveUserFromGroupResponseBody) SetSuccess(v bool) *RemoveUserFromGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveUserFromGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromGroupResponse) SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse {
	s.Body = v
	return s
}

type SendCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendCardHeaders) GoString() string {
	return s.String()
}

func (s *SendCardHeaders) SetCommonHeaders(v map[string]*string) *SendCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendCardRequest struct {
	// 卡片实例唯一标识
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 卡片变量赋值，json结构
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// 设备业务标识
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	// 设备uuid，唯一标识
	DeviceUuid *string `json:"deviceUuid,omitempty" xml:"deviceUuid,omitempty"`
	// 群id，群的唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 卡片模板唯一标识，开放平台获取
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 是否为吊顶卡片
	Topbox *bool `json:"topbox,omitempty" xml:"topbox,omitempty"`
	// 用户通讯录唯一标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCardRequest) GoString() string {
	return s.String()
}

func (s *SendCardRequest) SetBizId(v string) *SendCardRequest {
	s.BizId = &v
	return s
}

func (s *SendCardRequest) SetCardData(v string) *SendCardRequest {
	s.CardData = &v
	return s
}

func (s *SendCardRequest) SetDeviceCode(v string) *SendCardRequest {
	s.DeviceCode = &v
	return s
}

func (s *SendCardRequest) SetDeviceUuid(v string) *SendCardRequest {
	s.DeviceUuid = &v
	return s
}

func (s *SendCardRequest) SetOpenConversationId(v string) *SendCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendCardRequest) SetTemplateId(v string) *SendCardRequest {
	s.TemplateId = &v
	return s
}

func (s *SendCardRequest) SetTopbox(v bool) *SendCardRequest {
	s.Topbox = &v
	return s
}

func (s *SendCardRequest) SetUserId(v string) *SendCardRequest {
	s.UserId = &v
	return s
}

type SendCardResponseBody struct {
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendCardResponseBody) SetResult(v string) *SendCardResponseBody {
	s.Result = &v
	return s
}

func (s *SendCardResponseBody) SetSuccess(v bool) *SendCardResponseBody {
	s.Success = &v
	return s
}

type SendCardResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCardResponse) GoString() string {
	return s.String()
}

func (s *SendCardResponse) SetHeaders(v map[string]*string) *SendCardResponse {
	s.Headers = v
	return s
}

func (s *SendCardResponse) SetBody(v *SendCardResponseBody) *SendCardResponse {
	s.Body = v
	return s
}

type SendMsgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMsgHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMsgHeaders) GoString() string {
	return s.String()
}

func (s *SendMsgHeaders) SetCommonHeaders(v map[string]*string) *SendMsgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMsgHeaders) SetXAcsDingtalkAccessToken(v string) *SendMsgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMsgRequest struct {
	// 消息内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 设备业务标识
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	// 设备唯一系统标识
	DeviceUuid *string `json:"deviceUuid,omitempty" xml:"deviceUuid,omitempty"`
	// 开放群唯一标识
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// 用户列表，群聊时为被@的人，单聊时为目标对象
	UserList []*string `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s SendMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequest) GoString() string {
	return s.String()
}

func (s *SendMsgRequest) SetContent(v string) *SendMsgRequest {
	s.Content = &v
	return s
}

func (s *SendMsgRequest) SetDeviceCode(v string) *SendMsgRequest {
	s.DeviceCode = &v
	return s
}

func (s *SendMsgRequest) SetDeviceUuid(v string) *SendMsgRequest {
	s.DeviceUuid = &v
	return s
}

func (s *SendMsgRequest) SetOpenConversationId(v string) *SendMsgRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendMsgRequest) SetUserList(v []*string) *SendMsgRequest {
	s.UserList = v
	return s
}

type SendMsgResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMsgResponseBody) GoString() string {
	return s.String()
}

func (s *SendMsgResponseBody) SetResult(v string) *SendMsgResponseBody {
	s.Result = &v
	return s
}

func (s *SendMsgResponseBody) SetSuccess(v bool) *SendMsgResponseBody {
	s.Success = &v
	return s
}

type SendMsgResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMsgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMsgResponse) GoString() string {
	return s.String()
}

func (s *SendMsgResponse) SetHeaders(v map[string]*string) *SendMsgResponse {
	s.Headers = v
	return s
}

func (s *SendMsgResponse) SetBody(v *SendMsgResponseBody) *SendMsgResponse {
	s.Body = v
	return s
}

type UninstallDeviceRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UninstallDeviceRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s UninstallDeviceRobotHeaders) GoString() string {
	return s.String()
}

func (s *UninstallDeviceRobotHeaders) SetCommonHeaders(v map[string]*string) *UninstallDeviceRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UninstallDeviceRobotHeaders) SetXAcsDingtalkAccessToken(v string) *UninstallDeviceRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UninstallDeviceRobotRequest struct {
	// 设备编码，客户侧生成的设备标识，能够唯一标识一个设备，该字段与deviceUuid字段需要二选一，并且不能都填充。
	DeviceCode *string `json:"deviceCode,omitempty" xml:"deviceCode,omitempty"`
	// 设备唯一标识，钉钉侧生成的设备标识，能够唯一标识一个设备，该字段与deviceCode字段需要二选一，并且不能都填充。
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s UninstallDeviceRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallDeviceRobotRequest) GoString() string {
	return s.String()
}

func (s *UninstallDeviceRobotRequest) SetDeviceCode(v string) *UninstallDeviceRobotRequest {
	s.DeviceCode = &v
	return s
}

func (s *UninstallDeviceRobotRequest) SetUuid(v string) *UninstallDeviceRobotRequest {
	s.Uuid = &v
	return s
}

type UninstallDeviceRobotResponseBody struct {
	// 接口处理返回结果，内容为群的唯一标识。
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 接口处理是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UninstallDeviceRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallDeviceRobotResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallDeviceRobotResponseBody) SetResult(v string) *UninstallDeviceRobotResponseBody {
	s.Result = &v
	return s
}

func (s *UninstallDeviceRobotResponseBody) SetSuccess(v bool) *UninstallDeviceRobotResponseBody {
	s.Success = &v
	return s
}

type UninstallDeviceRobotResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallDeviceRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallDeviceRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallDeviceRobotResponse) GoString() string {
	return s.String()
}

func (s *UninstallDeviceRobotResponse) SetHeaders(v map[string]*string) *UninstallDeviceRobotResponse {
	s.Headers = v
	return s
}

func (s *UninstallDeviceRobotResponse) SetBody(v *UninstallDeviceRobotResponseBody) *UninstallDeviceRobotResponse {
	s.Body = v
	return s
}

type UpdateCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCardRequest struct {
	// 卡片实例唯一标识
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 卡片变量赋值，json结构
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
}

func (s UpdateCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateCardRequest) SetBizId(v string) *UpdateCardRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCardRequest) SetCardData(v string) *UpdateCardRequest {
	s.CardData = &v
	return s
}

type UpdateCardResponseBody struct {
	// 响应数据
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCardResponseBody) SetResult(v string) *UpdateCardResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCardResponseBody) SetSuccess(v bool) *UpdateCardResponseBody {
	s.Success = &v
	return s
}

type UpdateCardResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateCardResponse) SetHeaders(v map[string]*string) *UpdateCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateCardResponse) SetBody(v *UpdateCardResponseBody) *UpdateCardResponse {
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
	CoverUrl   *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
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

func (s *UploadEventRequest) SetCoverUrl(v string) *UploadEventRequest {
	s.CoverUrl = &v
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

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
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

func (client *Client) CreateDeviceChatRoom(request *CreateDeviceChatRoomRequest) (_result *CreateDeviceChatRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDeviceChatRoomHeaders{}
	_result = &CreateDeviceChatRoomResponse{}
	_body, _err := client.CreateDeviceChatRoomWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDeviceChatRoomWithOptions(request *CreateDeviceChatRoomRequest, headers *CreateDeviceChatRoomHeaders, runtime *util.RuntimeOptions) (_result *CreateDeviceChatRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatType)) {
		body["chatType"] = request.ChatType
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCodes)) {
		body["deviceCodes"] = request.DeviceCodes
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUuids)) {
		body["deviceUuids"] = request.DeviceUuids
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
	_result = &CreateDeviceChatRoomResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDeviceChatRoom"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/groups"), tea.String("json"), req, runtime)
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

func (client *Client) DissolveGroup(request *DissolveGroupRequest) (_result *DissolveGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DissolveGroupHeaders{}
	_result = &DissolveGroupResponse{}
	_body, _err := client.DissolveGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissolveGroupWithOptions(request *DissolveGroupRequest, headers *DissolveGroupHeaders, runtime *util.RuntimeOptions) (_result *DissolveGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &DissolveGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DissolveGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/groups/dissolve"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditDeviceAdmin(request *EditDeviceAdminRequest) (_result *EditDeviceAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditDeviceAdminHeaders{}
	_result = &EditDeviceAdminResponse{}
	_body, _err := client.EditDeviceAdminWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditDeviceAdminWithOptions(request *EditDeviceAdminRequest, headers *EditDeviceAdminHeaders, runtime *util.RuntimeOptions) (_result *EditDeviceAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		body["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.RoleUuid)) {
		body["roleUuid"] = request.RoleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
	_result = &EditDeviceAdminResponse{}
	_body, _err := client.DoROARequest(tea.String("EditDeviceAdmin"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/devices/administrators/edit"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceGroupInfo(request *GetDeviceGroupInfoRequest) (_result *GetDeviceGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeviceGroupInfoHeaders{}
	_result = &GetDeviceGroupInfoResponse{}
	_body, _err := client.GetDeviceGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceGroupInfoWithOptions(request *GetDeviceGroupInfoRequest, headers *GetDeviceGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *GetDeviceGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &GetDeviceGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDeviceGroupInfo"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/groupInfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWholeDeviceGroup() (_result *GetWholeDeviceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetWholeDeviceGroupHeaders{}
	_result = &GetWholeDeviceGroupResponse{}
	_body, _err := client.GetWholeDeviceGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWholeDeviceGroupWithOptions(headers *GetWholeDeviceGroupHeaders, runtime *util.RuntimeOptions) (_result *GetWholeDeviceGroupResponse, _err error) {
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
	_result = &GetWholeDeviceGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("GetWholeDeviceGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/wholeGroupId"), tea.String("json"), req, runtime)
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

func (client *Client) PullDeviceToGroup(request *PullDeviceToGroupRequest) (_result *PullDeviceToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullDeviceToGroupHeaders{}
	_result = &PullDeviceToGroupResponse{}
	_body, _err := client.PullDeviceToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDeviceToGroupWithOptions(request *PullDeviceToGroupRequest, headers *PullDeviceToGroupHeaders, runtime *util.RuntimeOptions) (_result *PullDeviceToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCodes)) {
		body["deviceCodes"] = request.DeviceCodes
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUuids)) {
		body["deviceUuids"] = request.DeviceUuids
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
	_result = &PullDeviceToGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("PullDeviceToGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/devices"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullUserToGroup(request *PullUserToGroupRequest) (_result *PullUserToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullUserToGroupHeaders{}
	_result = &PullUserToGroupResponse{}
	_body, _err := client.PullUserToGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullUserToGroupWithOptions(request *PullUserToGroupRequest, headers *PullUserToGroupHeaders, runtime *util.RuntimeOptions) (_result *PullUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &PullUserToGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("PullUserToGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/users"), tea.String("json"), req, runtime)
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

func (client *Client) RemoveDeviceFromGroup(request *RemoveDeviceFromGroupRequest) (_result *RemoveDeviceFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveDeviceFromGroupHeaders{}
	_result = &RemoveDeviceFromGroupResponse{}
	_body, _err := client.RemoveDeviceFromGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDeviceFromGroupWithOptions(request *RemoveDeviceFromGroupRequest, headers *RemoveDeviceFromGroupHeaders, runtime *util.RuntimeOptions) (_result *RemoveDeviceFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCodes)) {
		body["deviceCodes"] = request.DeviceCodes
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUuids)) {
		body["deviceUuids"] = request.DeviceUuids
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
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
	_result = &RemoveDeviceFromGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveDeviceFromGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/devices/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveUserFromGroupHeaders{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, headers *RemoveUserFromGroupHeaders, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveUserFromGroup"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/chatRooms/users/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCard(request *SendCardRequest) (_result *SendCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendCardHeaders{}
	_result = &SendCardResponse{}
	_body, _err := client.SendCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCardWithOptions(request *SendCardRequest, headers *SendCardHeaders, runtime *util.RuntimeOptions) (_result *SendCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		body["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceUuid)) {
		body["deviceUuid"] = request.DeviceUuid
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Topbox)) {
		body["topbox"] = request.Topbox
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
	_result = &SendCardResponse{}
	_body, _err := client.DoROARequest(tea.String("SendCard"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/cards/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMsg(request *SendMsgRequest) (_result *SendMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMsgHeaders{}
	_result = &SendMsgResponse{}
	_body, _err := client.SendMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMsgWithOptions(request *SendMsgRequest, headers *SendMsgHeaders, runtime *util.RuntimeOptions) (_result *SendMsgResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserList)) {
		body["userList"] = request.UserList
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
	_result = &SendMsgResponse{}
	_body, _err := client.DoROARequest(tea.String("SendMsg"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/messages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallDeviceRobot(request *UninstallDeviceRobotRequest) (_result *UninstallDeviceRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UninstallDeviceRobotHeaders{}
	_result = &UninstallDeviceRobotResponse{}
	_body, _err := client.UninstallDeviceRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallDeviceRobotWithOptions(request *UninstallDeviceRobotRequest, headers *UninstallDeviceRobotHeaders, runtime *util.RuntimeOptions) (_result *UninstallDeviceRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceCode)) {
		body["deviceCode"] = request.DeviceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
	_result = &UninstallDeviceRobotResponse{}
	_body, _err := client.DoROARequest(tea.String("UninstallDeviceRobot"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/devices/uninstall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCard(request *UpdateCardRequest) (_result *UpdateCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCardHeaders{}
	_result = &UpdateCardResponse{}
	_body, _err := client.UpdateCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCardWithOptions(request *UpdateCardRequest, headers *UpdateCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
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
	_result = &UpdateCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCard"), tea.String("devicemng_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/devicemng/customers/cards"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
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
