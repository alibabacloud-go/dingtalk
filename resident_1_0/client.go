// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package resident_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateResideceGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResideceGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateResideceGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResideceGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResideceGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResideceGroupRequest struct {
	// 组长userid
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// 组名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s UpdateResideceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupRequest) SetManagerUserId(v string) *UpdateResideceGroupRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetDepartmentName(v string) *UpdateResideceGroupRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetDepartmentId(v int64) *UpdateResideceGroupRequest {
	s.DepartmentId = &v
	return s
}

type UpdateResideceGroupResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResideceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupResponseBody) SetResult(v bool) *UpdateResideceGroupResponseBody {
	s.Result = &v
	return s
}

type UpdateResideceGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResideceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResideceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupResponse) SetHeaders(v map[string]*string) *UpdateResideceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResideceGroupResponse) SetBody(v *UpdateResideceGroupResponseBody) *UpdateResideceGroupResponse {
	s.Body = v
	return s
}

type DeleteResidentDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteResidentDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentHeaders) SetCommonHeaders(v map[string]*string) *DeleteResidentDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteResidentDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteResidentDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteResidentDepartmentRequest struct {
	// 组/户id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s DeleteResidentDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentRequest) SetDepartmentId(v int64) *DeleteResidentDepartmentRequest {
	s.DepartmentId = &v
	return s
}

type DeleteResidentDepartmentResponseBody struct {
	// 是否删除成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteResidentDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentResponseBody) SetResult(v bool) *DeleteResidentDepartmentResponseBody {
	s.Result = &v
	return s
}

type DeleteResidentDepartmentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResidentDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteResidentDepartmentResponse) SetHeaders(v map[string]*string) *DeleteResidentDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteResidentDepartmentResponse) SetBody(v *DeleteResidentDepartmentResponseBody) *DeleteResidentDepartmentResponse {
	s.Body = v
	return s
}

type AddResidentUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddResidentUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersHeaders) GoString() string {
	return s.String()
}

func (s *AddResidentUsersHeaders) SetCommonHeaders(v map[string]*string) *AddResidentUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddResidentUsersHeaders) SetXAcsDingtalkAccessToken(v string) *AddResidentUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddResidentUsersRequest struct {
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否是租客
	IsLeaseholder *bool `json:"isLeaseholder,omitempty" xml:"isLeaseholder,omitempty"`
	// 居民名字
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*AddResidentUsersRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
}

func (s AddResidentUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersRequest) GoString() string {
	return s.String()
}

func (s *AddResidentUsersRequest) SetAddress(v string) *AddResidentUsersRequest {
	s.Address = &v
	return s
}

func (s *AddResidentUsersRequest) SetIsLeaseholder(v bool) *AddResidentUsersRequest {
	s.IsLeaseholder = &v
	return s
}

func (s *AddResidentUsersRequest) SetUserName(v string) *AddResidentUsersRequest {
	s.UserName = &v
	return s
}

func (s *AddResidentUsersRequest) SetMobile(v string) *AddResidentUsersRequest {
	s.Mobile = &v
	return s
}

func (s *AddResidentUsersRequest) SetDepartmentId(v int64) *AddResidentUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *AddResidentUsersRequest) SetExtField(v []*AddResidentUsersRequestExtField) *AddResidentUsersRequest {
	s.ExtField = v
	return s
}

func (s *AddResidentUsersRequest) SetRelateType(v string) *AddResidentUsersRequest {
	s.RelateType = &v
	return s
}

type AddResidentUsersRequestExtField struct {
	// 扩展字段值
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
}

func (s AddResidentUsersRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersRequestExtField) GoString() string {
	return s.String()
}

func (s *AddResidentUsersRequestExtField) SetItemValue(v string) *AddResidentUsersRequestExtField {
	s.ItemValue = &v
	return s
}

func (s *AddResidentUsersRequestExtField) SetItemName(v string) *AddResidentUsersRequestExtField {
	s.ItemName = &v
	return s
}

type AddResidentUsersResponseBody struct {
	// 创建成功的userId
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddResidentUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentUsersResponseBody) SetResult(v string) *AddResidentUsersResponseBody {
	s.Result = &v
	return s
}

type AddResidentUsersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResidentUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddResidentUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersResponse) GoString() string {
	return s.String()
}

func (s *AddResidentUsersResponse) SetHeaders(v map[string]*string) *AddResidentUsersResponse {
	s.Headers = v
	return s
}

func (s *AddResidentUsersResponse) SetBody(v *AddResidentUsersResponseBody) *AddResidentUsersResponse {
	s.Body = v
	return s
}

type AddResidentDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddResidentDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentHeaders) SetCommonHeaders(v map[string]*string) *AddResidentDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddResidentDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *AddResidentDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddResidentDepartmentRequest struct {
	// 是否为组
	IsResidenceGroup *bool `json:"isResidenceGroup,omitempty" xml:"isResidenceGroup,omitempty"`
	// 部门名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 父部门id
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s AddResidentDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentRequest) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentRequest) SetIsResidenceGroup(v bool) *AddResidentDepartmentRequest {
	s.IsResidenceGroup = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetDepartmentName(v string) *AddResidentDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetParentDepartmentId(v int64) *AddResidentDepartmentRequest {
	s.ParentDepartmentId = &v
	return s
}

type AddResidentDepartmentResponseBody struct {
	// 创建成功的deptId
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddResidentDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentResponseBody) SetResult(v int64) *AddResidentDepartmentResponseBody {
	s.Result = &v
	return s
}

type AddResidentDepartmentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddResidentDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentResponse) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentResponse) SetHeaders(v map[string]*string) *AddResidentDepartmentResponse {
	s.Headers = v
	return s
}

func (s *AddResidentDepartmentResponse) SetBody(v *AddResidentDepartmentResponseBody) *AddResidentDepartmentResponse {
	s.Body = v
	return s
}

type RemoveResidentUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveResidentUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserHeaders) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserHeaders) SetCommonHeaders(v map[string]*string) *RemoveResidentUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveResidentUserHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveResidentUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveResidentUserRequest struct {
	// 户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RemoveResidentUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserRequest) SetDepartmentId(v int64) *RemoveResidentUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *RemoveResidentUserRequest) SetUserId(v string) *RemoveResidentUserRequest {
	s.UserId = &v
	return s
}

type RemoveResidentUserResponseBody struct {
	// 是否移除成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveResidentUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserResponseBody) SetResult(v bool) *RemoveResidentUserResponseBody {
	s.Result = &v
	return s
}

type RemoveResidentUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveResidentUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveResidentUserResponse) SetHeaders(v map[string]*string) *RemoveResidentUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveResidentUserResponse) SetBody(v *RemoveResidentUserResponseBody) *RemoveResidentUserResponse {
	s.Body = v
	return s
}

type UpdateResidenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidenceHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidenceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidenceRequest struct {
	// 家庭管理员用户id
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// 户名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 所属网格
	Grid *string `json:"grid,omitempty" xml:"grid,omitempty"`
	// 家庭电话
	HomeTel *string `json:"homeTel,omitempty" xml:"homeTel,omitempty"`
	// 是否是贫困户
	Destitute *bool `json:"destitute,omitempty" xml:"destitute,omitempty"`
	// 组id
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s UpdateResidenceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidenceRequest) SetManagerUserId(v string) *UpdateResidenceRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateResidenceRequest) SetDepartmentName(v string) *UpdateResidenceRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResidenceRequest) SetDepartmentId(v int64) *UpdateResidenceRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidenceRequest) SetGrid(v string) *UpdateResidenceRequest {
	s.Grid = &v
	return s
}

func (s *UpdateResidenceRequest) SetHomeTel(v string) *UpdateResidenceRequest {
	s.HomeTel = &v
	return s
}

func (s *UpdateResidenceRequest) SetDestitute(v bool) *UpdateResidenceRequest {
	s.Destitute = &v
	return s
}

func (s *UpdateResidenceRequest) SetParentDepartmentId(v int64) *UpdateResidenceRequest {
	s.ParentDepartmentId = &v
	return s
}

type UpdateResidenceResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResidenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidenceResponseBody) SetResult(v bool) *UpdateResidenceResponseBody {
	s.Result = &v
	return s
}

type UpdateResidenceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResidenceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidenceResponse) SetHeaders(v map[string]*string) *UpdateResidenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidenceResponse) SetBody(v *UpdateResidenceResponseBody) *UpdateResidenceResponse {
	s.Body = v
	return s
}

type UpdateResidentUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidentUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidentUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidentUserHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidentUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidentUserRequest struct {
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 是否保留原部门
	IsRetainOldDept *bool `json:"isRetainOldDept,omitempty" xml:"isRetainOldDept,omitempty"`
	// 居民名字
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 所在新的户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*UpdateResidentUserRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// 人员userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 原所在部门id
	OldDepartmentId *int64 `json:"oldDepartmentId,omitempty" xml:"oldDepartmentId,omitempty"`
}

func (s UpdateResidentUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserRequest) SetAddress(v string) *UpdateResidentUserRequest {
	s.Address = &v
	return s
}

func (s *UpdateResidentUserRequest) SetIsRetainOldDept(v bool) *UpdateResidentUserRequest {
	s.IsRetainOldDept = &v
	return s
}

func (s *UpdateResidentUserRequest) SetUserName(v string) *UpdateResidentUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateResidentUserRequest) SetMobile(v string) *UpdateResidentUserRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateResidentUserRequest) SetDepartmentId(v int64) *UpdateResidentUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidentUserRequest) SetExtField(v []*UpdateResidentUserRequestExtField) *UpdateResidentUserRequest {
	s.ExtField = v
	return s
}

func (s *UpdateResidentUserRequest) SetRelateType(v string) *UpdateResidentUserRequest {
	s.RelateType = &v
	return s
}

func (s *UpdateResidentUserRequest) SetUserId(v string) *UpdateResidentUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateResidentUserRequest) SetOldDepartmentId(v int64) *UpdateResidentUserRequest {
	s.OldDepartmentId = &v
	return s
}

type UpdateResidentUserRequestExtField struct {
	// 扩展字段值
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
}

func (s UpdateResidentUserRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserRequestExtField) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserRequestExtField) SetItemValue(v string) *UpdateResidentUserRequestExtField {
	s.ItemValue = &v
	return s
}

func (s *UpdateResidentUserRequestExtField) SetItemName(v string) *UpdateResidentUserRequestExtField {
	s.ItemName = &v
	return s
}

type UpdateResidentUserResponseBody struct {
	// 是否更新成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateResidentUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserResponseBody) SetResult(v bool) *UpdateResidentUserResponseBody {
	s.Result = &v
	return s
}

type UpdateResidentUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResidentUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserResponse) SetHeaders(v map[string]*string) *UpdateResidentUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidentUserResponse) SetBody(v *UpdateResidentUserResponseBody) *UpdateResidentUserResponse {
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

func (client *Client) UpdateResideceGroup(request *UpdateResideceGroupRequest) (_result *UpdateResideceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResideceGroupHeaders{}
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.UpdateResideceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResideceGroupWithOptions(request *UpdateResideceGroupRequest, headers *UpdateResideceGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateResideceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResideceGroup"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResideceGroup"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResidentDepartment(request *DeleteResidentDepartmentRequest) (_result *DeleteResidentDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteResidentDepartmentHeaders{}
	_result = &DeleteResidentDepartmentResponse{}
	_body, _err := client.DeleteResidentDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResidentDepartmentWithOptions(request *DeleteResidentDepartmentRequest, headers *DeleteResidentDepartmentHeaders, runtime *util.RuntimeOptions) (_result *DeleteResidentDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
	_result = &DeleteResidentDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteResidentDepartment"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/resident/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResidentUsers(request *AddResidentUsersRequest) (_result *AddResidentUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddResidentUsersHeaders{}
	_result = &AddResidentUsersResponse{}
	_body, _err := client.AddResidentUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddResidentUsersWithOptions(request *AddResidentUsersRequest, headers *AddResidentUsersHeaders, runtime *util.RuntimeOptions) (_result *AddResidentUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.IsLeaseholder)) {
		query["isLeaseholder"] = request.IsLeaseholder
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
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
	_result = &AddResidentUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentUsers"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResidentDepartment(request *AddResidentDepartmentRequest) (_result *AddResidentDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddResidentDepartmentHeaders{}
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.AddResidentDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddResidentDepartmentWithOptions(request *AddResidentDepartmentRequest, headers *AddResidentDepartmentHeaders, runtime *util.RuntimeOptions) (_result *AddResidentDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsResidenceGroup)) {
		query["isResidenceGroup"] = request.IsResidenceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentDepartment"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveResidentUser(request *RemoveResidentUserRequest) (_result *RemoveResidentUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveResidentUserHeaders{}
	_result = &RemoveResidentUserResponse{}
	_body, _err := client.RemoveResidentUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveResidentUserWithOptions(request *RemoveResidentUserRequest, headers *RemoveResidentUserHeaders, runtime *util.RuntimeOptions) (_result *RemoveResidentUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
	_result = &RemoveResidentUserResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveResidentUser"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/users/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResidence(request *UpdateResidenceRequest) (_result *UpdateResidenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidenceHeaders{}
	_result = &UpdateResidenceResponse{}
	_body, _err := client.UpdateResidenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResidenceWithOptions(request *UpdateResidenceRequest, headers *UpdateResidenceHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Grid)) {
		query["grid"] = request.Grid
	}

	if !tea.BoolValue(util.IsUnset(request.HomeTel)) {
		query["homeTel"] = request.HomeTel
	}

	if !tea.BoolValue(util.IsUnset(request.Destitute)) {
		query["destitute"] = request.Destitute
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
	_result = &UpdateResidenceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidence"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResidece"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResidentUser(request *UpdateResidentUserRequest) (_result *UpdateResidentUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidentUserHeaders{}
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.UpdateResidentUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResidentUserWithOptions(request *UpdateResidentUserRequest, headers *UpdateResidentUserHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.IsRetainOldDept)) {
		query["isRetainOldDept"] = request.IsRetainOldDept
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.OldDepartmentId)) {
		query["oldDepartmentId"] = request.OldDepartmentId
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
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentUser"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
