// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package village_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeptHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptHeaders) SetCommonHeaders(v map[string]*string) *GetDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeptRequest struct {
	Language  *string `json:"language,omitempty" xml:"language,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeptRequest) GoString() string {
	return s.String()
}

func (s *GetDeptRequest) SetLanguage(v string) *GetDeptRequest {
	s.Language = &v
	return s
}

func (s *GetDeptRequest) SetSubCorpId(v string) *GetDeptRequest {
	s.SubCorpId = &v
	return s
}

type GetDeptResponseBody struct {
	DepartmentId       *int64  `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName     *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	FromUnionOrg       *bool   `json:"fromUnionOrg,omitempty" xml:"fromUnionOrg,omitempty"`
	Order              *int64  `json:"order,omitempty" xml:"order,omitempty"`
	ParentDepartmentId *int64  `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s GetDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeptResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeptResponseBody) SetDepartmentId(v int64) *GetDeptResponseBody {
	s.DepartmentId = &v
	return s
}

func (s *GetDeptResponseBody) SetDepartmentName(v string) *GetDeptResponseBody {
	s.DepartmentName = &v
	return s
}

func (s *GetDeptResponseBody) SetFromUnionOrg(v bool) *GetDeptResponseBody {
	s.FromUnionOrg = &v
	return s
}

func (s *GetDeptResponseBody) SetOrder(v int64) *GetDeptResponseBody {
	s.Order = &v
	return s
}

func (s *GetDeptResponseBody) SetParentDepartmentId(v int64) *GetDeptResponseBody {
	s.ParentDepartmentId = &v
	return s
}

type GetDeptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeptResponse) GoString() string {
	return s.String()
}

func (s *GetDeptResponse) SetHeaders(v map[string]*string) *GetDeptResponse {
	s.Headers = v
	return s
}

func (s *GetDeptResponse) SetStatusCode(v int32) *GetDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeptResponse) SetBody(v *GetDeptResponseBody) *GetDeptResponse {
	s.Body = v
	return s
}

type GetResidentDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentDeptHeaders) SetCommonHeaders(v map[string]*string) *GetResidentDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentDeptHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentDeptRequest struct {
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetResidentDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptRequest) GoString() string {
	return s.String()
}

func (s *GetResidentDeptRequest) SetSubCorpId(v string) *GetResidentDeptRequest {
	s.SubCorpId = &v
	return s
}

type GetResidentDeptResponseBody struct {
	ContactType    *string `json:"contactType,omitempty" xml:"contactType,omitempty"`
	DepartmentId   *int64  `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	DeptType       *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	Feature        *string `json:"feature,omitempty" xml:"feature,omitempty"`
}

func (s GetResidentDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentDeptResponseBody) SetContactType(v string) *GetResidentDeptResponseBody {
	s.ContactType = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetDepartmentId(v int64) *GetResidentDeptResponseBody {
	s.DepartmentId = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetDepartmentName(v string) *GetResidentDeptResponseBody {
	s.DepartmentName = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetDeptType(v string) *GetResidentDeptResponseBody {
	s.DeptType = &v
	return s
}

func (s *GetResidentDeptResponseBody) SetFeature(v string) *GetResidentDeptResponseBody {
	s.Feature = &v
	return s
}

type GetResidentDeptResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResidentDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResidentDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentDeptResponse) GoString() string {
	return s.String()
}

func (s *GetResidentDeptResponse) SetHeaders(v map[string]*string) *GetResidentDeptResponse {
	s.Headers = v
	return s
}

func (s *GetResidentDeptResponse) SetStatusCode(v int32) *GetResidentDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResidentDeptResponse) SetBody(v *GetResidentDeptResponseBody) *GetResidentDeptResponse {
	s.Body = v
	return s
}

type GetResidentUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetResidentUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentUserInfoRequest struct {
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetResidentUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoRequest) SetSubCorpId(v string) *GetResidentUserInfoRequest {
	s.SubCorpId = &v
	return s
}

type GetResidentUserInfoResponseBody struct {
	Feature *string                                 `json:"feature,omitempty" xml:"feature,omitempty"`
	Name    *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Roles   []*GetResidentUserInfoResponseBodyRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	UnionId *string                                 `json:"unionId,omitempty" xml:"unionId,omitempty"`
	Userid  *string                                 `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetResidentUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponseBody) SetFeature(v string) *GetResidentUserInfoResponseBody {
	s.Feature = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetName(v string) *GetResidentUserInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetRoles(v []*GetResidentUserInfoResponseBodyRoles) *GetResidentUserInfoResponseBody {
	s.Roles = v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetUnionId(v string) *GetResidentUserInfoResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetResidentUserInfoResponseBody) SetUserid(v string) *GetResidentUserInfoResponseBody {
	s.Userid = &v
	return s
}

type GetResidentUserInfoResponseBodyRoles struct {
	RoleId   *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	TagCode  *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s GetResidentUserInfoResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponseBodyRoles) SetRoleId(v int64) *GetResidentUserInfoResponseBodyRoles {
	s.RoleId = &v
	return s
}

func (s *GetResidentUserInfoResponseBodyRoles) SetRoleName(v string) *GetResidentUserInfoResponseBodyRoles {
	s.RoleName = &v
	return s
}

func (s *GetResidentUserInfoResponseBodyRoles) SetTagCode(v string) *GetResidentUserInfoResponseBodyRoles {
	s.TagCode = &v
	return s
}

type GetResidentUserInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResidentUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResidentUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResidentUserInfoResponse) SetHeaders(v map[string]*string) *GetResidentUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResidentUserInfoResponse) SetStatusCode(v int32) *GetResidentUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResidentUserInfoResponse) SetBody(v *GetResidentUserInfoResponseBody) *GetResidentUserInfoResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserRequest struct {
	Language  *string `json:"language,omitempty" xml:"language,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetLanguage(v string) *GetUserRequest {
	s.Language = &v
	return s
}

func (s *GetUserRequest) SetSubCorpId(v string) *GetUserRequest {
	s.SubCorpId = &v
	return s
}

type GetUserResponseBody struct {
	Active               *bool                                    `json:"active,omitempty" xml:"active,omitempty"`
	Admin                *bool                                    `json:"admin,omitempty" xml:"admin,omitempty"`
	Boss                 *bool                                    `json:"boss,omitempty" xml:"boss,omitempty"`
	DepartmentIdList     []*int64                                 `json:"departmentIdList,omitempty" xml:"departmentIdList,omitempty" type:"Repeated"`
	DepartmentOrderSet   []*GetUserResponseBodyDepartmentOrderSet `json:"departmentOrderSet,omitempty" xml:"departmentOrderSet,omitempty" type:"Repeated"`
	ExclusiveAccount     *bool                                    `json:"exclusiveAccount,omitempty" xml:"exclusiveAccount,omitempty"`
	ExclusiveAccountType *string                                  `json:"exclusiveAccountType,omitempty" xml:"exclusiveAccountType,omitempty"`
	Extension            *string                                  `json:"extension,omitempty" xml:"extension,omitempty"`
	HiredDate            *int64                                   `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	JobNumber            *string                                  `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	LeaderInDepartment   []*GetUserResponseBodyLeaderInDepartment `json:"leaderInDepartment,omitempty" xml:"leaderInDepartment,omitempty" type:"Repeated"`
	ManagerUserId        *string                                  `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	Name                 *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	RealAuthed           *bool                                    `json:"realAuthed,omitempty" xml:"realAuthed,omitempty"`
	Remark               *string                                  `json:"remark,omitempty" xml:"remark,omitempty"`
	RoleList             []*GetUserResponseBodyRoleList           `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	Senior               *bool                                    `json:"senior,omitempty" xml:"senior,omitempty"`
	Title                *string                                  `json:"title,omitempty" xml:"title,omitempty"`
	UnionEmpExt          *GetUserResponseBodyUnionEmpExt          `json:"unionEmpExt,omitempty" xml:"unionEmpExt,omitempty" type:"Struct"`
	UnionId              *string                                  `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId               *string                                  `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkPlace            *string                                  `json:"workPlace,omitempty" xml:"workPlace,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetActive(v bool) *GetUserResponseBody {
	s.Active = &v
	return s
}

func (s *GetUserResponseBody) SetAdmin(v bool) *GetUserResponseBody {
	s.Admin = &v
	return s
}

func (s *GetUserResponseBody) SetBoss(v bool) *GetUserResponseBody {
	s.Boss = &v
	return s
}

func (s *GetUserResponseBody) SetDepartmentIdList(v []*int64) *GetUserResponseBody {
	s.DepartmentIdList = v
	return s
}

func (s *GetUserResponseBody) SetDepartmentOrderSet(v []*GetUserResponseBodyDepartmentOrderSet) *GetUserResponseBody {
	s.DepartmentOrderSet = v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccount(v bool) *GetUserResponseBody {
	s.ExclusiveAccount = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountType(v string) *GetUserResponseBody {
	s.ExclusiveAccountType = &v
	return s
}

func (s *GetUserResponseBody) SetExtension(v string) *GetUserResponseBody {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBody) SetHiredDate(v int64) *GetUserResponseBody {
	s.HiredDate = &v
	return s
}

func (s *GetUserResponseBody) SetJobNumber(v string) *GetUserResponseBody {
	s.JobNumber = &v
	return s
}

func (s *GetUserResponseBody) SetLeaderInDepartment(v []*GetUserResponseBodyLeaderInDepartment) *GetUserResponseBody {
	s.LeaderInDepartment = v
	return s
}

func (s *GetUserResponseBody) SetManagerUserId(v string) *GetUserResponseBody {
	s.ManagerUserId = &v
	return s
}

func (s *GetUserResponseBody) SetName(v string) *GetUserResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserResponseBody) SetRealAuthed(v bool) *GetUserResponseBody {
	s.RealAuthed = &v
	return s
}

func (s *GetUserResponseBody) SetRemark(v string) *GetUserResponseBody {
	s.Remark = &v
	return s
}

func (s *GetUserResponseBody) SetRoleList(v []*GetUserResponseBodyRoleList) *GetUserResponseBody {
	s.RoleList = v
	return s
}

func (s *GetUserResponseBody) SetSenior(v bool) *GetUserResponseBody {
	s.Senior = &v
	return s
}

func (s *GetUserResponseBody) SetTitle(v string) *GetUserResponseBody {
	s.Title = &v
	return s
}

func (s *GetUserResponseBody) SetUnionEmpExt(v *GetUserResponseBodyUnionEmpExt) *GetUserResponseBody {
	s.UnionEmpExt = v
	return s
}

func (s *GetUserResponseBody) SetUnionId(v string) *GetUserResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserResponseBody) SetUserId(v string) *GetUserResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBody) SetWorkPlace(v string) *GetUserResponseBody {
	s.WorkPlace = &v
	return s
}

type GetUserResponseBodyDepartmentOrderSet struct {
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	Order        *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetUserResponseBodyDepartmentOrderSet) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDepartmentOrderSet) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDepartmentOrderSet) SetDepartmentId(v int64) *GetUserResponseBodyDepartmentOrderSet {
	s.DepartmentId = &v
	return s
}

func (s *GetUserResponseBodyDepartmentOrderSet) SetOrder(v int64) *GetUserResponseBodyDepartmentOrderSet {
	s.Order = &v
	return s
}

type GetUserResponseBodyLeaderInDepartment struct {
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	Leader       *bool  `json:"leader,omitempty" xml:"leader,omitempty"`
}

func (s GetUserResponseBodyLeaderInDepartment) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyLeaderInDepartment) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyLeaderInDepartment) SetDepartmentId(v int64) *GetUserResponseBodyLeaderInDepartment {
	s.DepartmentId = &v
	return s
}

func (s *GetUserResponseBodyLeaderInDepartment) SetLeader(v bool) *GetUserResponseBodyLeaderInDepartment {
	s.Leader = &v
	return s
}

type GetUserResponseBodyRoleList struct {
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	RoleId    *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName  *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s GetUserResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyRoleList) SetGroupName(v string) *GetUserResponseBodyRoleList {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetRoleId(v int64) *GetUserResponseBodyRoleList {
	s.RoleId = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetRoleName(v string) *GetUserResponseBodyRoleList {
	s.RoleName = &v
	return s
}

type GetUserResponseBodyUnionEmpExt struct {
	CorpId          *string                                          `json:"corpId,omitempty" xml:"corpId,omitempty"`
	StaffId         *string                                          `json:"staffId,omitempty" xml:"staffId,omitempty"`
	UnionEmpMapList []*GetUserResponseBodyUnionEmpExtUnionEmpMapList `json:"unionEmpMapList,omitempty" xml:"unionEmpMapList,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUnionEmpExt) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExt) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExt) SetCorpId(v string) *GetUserResponseBodyUnionEmpExt {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetStaffId(v string) *GetUserResponseBodyUnionEmpExt {
	s.StaffId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUnionEmpMapList(v []*GetUserResponseBodyUnionEmpExtUnionEmpMapList) *GetUserResponseBodyUnionEmpExt {
	s.UnionEmpMapList = v
	return s
}

type GetUserResponseBodyUnionEmpExtUnionEmpMapList struct {
	CorpId  *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetCorpId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetStaffId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.StaffId = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserByUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserByUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserByUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserByUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserByUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserByUnionIdRequest struct {
	Language  *string `json:"language,omitempty" xml:"language,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserByUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdRequest) SetLanguage(v string) *GetUserByUnionIdRequest {
	s.Language = &v
	return s
}

func (s *GetUserByUnionIdRequest) SetSubCorpId(v string) *GetUserByUnionIdRequest {
	s.SubCorpId = &v
	return s
}

func (s *GetUserByUnionIdRequest) SetUnionId(v string) *GetUserByUnionIdRequest {
	s.UnionId = &v
	return s
}

type GetUserByUnionIdResponseBody struct {
	ContactType *int32  `json:"contactType,omitempty" xml:"contactType,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserByUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdResponseBody) SetContactType(v int32) *GetUserByUnionIdResponseBody {
	s.ContactType = &v
	return s
}

func (s *GetUserByUnionIdResponseBody) SetUserId(v string) *GetUserByUnionIdResponseBody {
	s.UserId = &v
	return s
}

type GetUserByUnionIdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserByUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserByUnionIdResponse) SetHeaders(v map[string]*string) *GetUserByUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserByUnionIdResponse) SetStatusCode(v int32) *GetUserByUnionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserByUnionIdResponse) SetBody(v *GetUserByUnionIdResponseBody) *GetUserByUnionIdResponse {
	s.Body = v
	return s
}

type GetVillageOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetVillageOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *GetVillageOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetVillageOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetVillageOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetVillageOrgInfoResponseBody struct {
	RegionId       *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RegionLocation *string `json:"regionLocation,omitempty" xml:"regionLocation,omitempty"`
	RegionType     *string `json:"regionType,omitempty" xml:"regionType,omitempty"`
}

func (s GetVillageOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoResponseBody) SetRegionId(v string) *GetVillageOrgInfoResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVillageOrgInfoResponseBody) SetRegionLocation(v string) *GetVillageOrgInfoResponseBody {
	s.RegionLocation = &v
	return s
}

func (s *GetVillageOrgInfoResponseBody) SetRegionType(v string) *GetVillageOrgInfoResponseBody {
	s.RegionType = &v
	return s
}

type GetVillageOrgInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVillageOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVillageOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVillageOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVillageOrgInfoResponse) SetHeaders(v map[string]*string) *GetVillageOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVillageOrgInfoResponse) SetStatusCode(v int32) *GetVillageOrgInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVillageOrgInfoResponse) SetBody(v *GetVillageOrgInfoResponseBody) *GetVillageOrgInfoResponse {
	s.Body = v
	return s
}

type ListDeptSimpleUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptSimpleUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersHeaders) SetCommonHeaders(v map[string]*string) *ListDeptSimpleUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptSimpleUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptSimpleUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptSimpleUsersRequest struct {
	ContainAccessLimit *bool   `json:"containAccessLimit,omitempty" xml:"containAccessLimit,omitempty"`
	Cursor             *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Language           *string `json:"language,omitempty" xml:"language,omitempty"`
	OrderField         *string `json:"orderField,omitempty" xml:"orderField,omitempty"`
	Size               *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SubCorpId          *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListDeptSimpleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersRequest) SetContainAccessLimit(v bool) *ListDeptSimpleUsersRequest {
	s.ContainAccessLimit = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetCursor(v int64) *ListDeptSimpleUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetLanguage(v string) *ListDeptSimpleUsersRequest {
	s.Language = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetOrderField(v string) *ListDeptSimpleUsersRequest {
	s.OrderField = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetSize(v int32) *ListDeptSimpleUsersRequest {
	s.Size = &v
	return s
}

func (s *ListDeptSimpleUsersRequest) SetSubCorpId(v string) *ListDeptSimpleUsersRequest {
	s.SubCorpId = &v
	return s
}

type ListDeptSimpleUsersResponseBody struct {
	HasMore    *bool                                      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextCursor *int64                                     `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	UserList   []*ListDeptSimpleUsersResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListDeptSimpleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponseBody) SetHasMore(v bool) *ListDeptSimpleUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetNextCursor(v int64) *ListDeptSimpleUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetTotalCount(v int64) *ListDeptSimpleUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBody) SetUserList(v []*ListDeptSimpleUsersResponseBodyUserList) *ListDeptSimpleUsersResponseBody {
	s.UserList = v
	return s
}

type ListDeptSimpleUsersResponseBodyUserList struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListDeptSimpleUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponseBodyUserList) SetName(v string) *ListDeptSimpleUsersResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *ListDeptSimpleUsersResponseBodyUserList) SetUserId(v string) *ListDeptSimpleUsersResponseBodyUserList {
	s.UserId = &v
	return s
}

type ListDeptSimpleUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeptSimpleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptSimpleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptSimpleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDeptSimpleUsersResponse) SetHeaders(v map[string]*string) *ListDeptSimpleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDeptSimpleUsersResponse) SetStatusCode(v int32) *ListDeptSimpleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeptSimpleUsersResponse) SetBody(v *ListDeptSimpleUsersResponseBody) *ListDeptSimpleUsersResponse {
	s.Body = v
	return s
}

type ListDeptUserIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptUserIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsHeaders) SetCommonHeaders(v map[string]*string) *ListDeptUserIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptUserIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptUserIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptUserIdsRequest struct {
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListDeptUserIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsRequest) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsRequest) SetSubCorpId(v string) *ListDeptUserIdsRequest {
	s.SubCorpId = &v
	return s
}

type ListDeptUserIdsResponseBody struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ListDeptUserIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsResponseBody) SetUserIdList(v []*string) *ListDeptUserIdsResponseBody {
	s.UserIdList = v
	return s
}

type ListDeptUserIdsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeptUserIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptUserIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUserIdsResponse) GoString() string {
	return s.String()
}

func (s *ListDeptUserIdsResponse) SetHeaders(v map[string]*string) *ListDeptUserIdsResponse {
	s.Headers = v
	return s
}

func (s *ListDeptUserIdsResponse) SetStatusCode(v int32) *ListDeptUserIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeptUserIdsResponse) SetBody(v *ListDeptUserIdsResponseBody) *ListDeptUserIdsResponse {
	s.Body = v
	return s
}

type ListDeptUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListDeptUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListDeptUsersHeaders) SetCommonHeaders(v map[string]*string) *ListDeptUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeptUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListDeptUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListDeptUsersRequest struct {
	ContainAccessLimit *bool   `json:"containAccessLimit,omitempty" xml:"containAccessLimit,omitempty"`
	Cursor             *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Language           *string `json:"language,omitempty" xml:"language,omitempty"`
	OrderField         *string `json:"orderField,omitempty" xml:"orderField,omitempty"`
	Size               *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SubCorpId          *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListDeptUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDeptUsersRequest) SetContainAccessLimit(v bool) *ListDeptUsersRequest {
	s.ContainAccessLimit = &v
	return s
}

func (s *ListDeptUsersRequest) SetCursor(v int64) *ListDeptUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeptUsersRequest) SetLanguage(v string) *ListDeptUsersRequest {
	s.Language = &v
	return s
}

func (s *ListDeptUsersRequest) SetOrderField(v string) *ListDeptUsersRequest {
	s.OrderField = &v
	return s
}

func (s *ListDeptUsersRequest) SetSize(v int32) *ListDeptUsersRequest {
	s.Size = &v
	return s
}

func (s *ListDeptUsersRequest) SetSubCorpId(v string) *ListDeptUsersRequest {
	s.SubCorpId = &v
	return s
}

type ListDeptUsersResponseBody struct {
	HasMore    *bool                                `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextCursor *int64                               `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	UserList   []*ListDeptUsersResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListDeptUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponseBody) SetHasMore(v bool) *ListDeptUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListDeptUsersResponseBody) SetNextCursor(v int64) *ListDeptUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListDeptUsersResponseBody) SetUserList(v []*ListDeptUsersResponseBodyUserList) *ListDeptUsersResponseBody {
	s.UserList = v
	return s
}

type ListDeptUsersResponseBodyUserList struct {
	Active         *bool    `json:"active,omitempty" xml:"active,omitempty"`
	DepartmentList []*int64 `json:"departmentList,omitempty" xml:"departmentList,omitempty" type:"Repeated"`
	JobNumber      *string  `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	UnionId        *string  `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId         *string  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListDeptUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponseBodyUserList) SetActive(v bool) *ListDeptUsersResponseBodyUserList {
	s.Active = &v
	return s
}

func (s *ListDeptUsersResponseBodyUserList) SetDepartmentList(v []*int64) *ListDeptUsersResponseBodyUserList {
	s.DepartmentList = v
	return s
}

func (s *ListDeptUsersResponseBodyUserList) SetJobNumber(v string) *ListDeptUsersResponseBodyUserList {
	s.JobNumber = &v
	return s
}

func (s *ListDeptUsersResponseBodyUserList) SetName(v string) *ListDeptUsersResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *ListDeptUsersResponseBodyUserList) SetUnionId(v string) *ListDeptUsersResponseBodyUserList {
	s.UnionId = &v
	return s
}

func (s *ListDeptUsersResponseBodyUserList) SetUserId(v string) *ListDeptUsersResponseBodyUserList {
	s.UserId = &v
	return s
}

type ListDeptUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeptUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeptUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeptUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDeptUsersResponse) SetHeaders(v map[string]*string) *ListDeptUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDeptUsersResponse) SetStatusCode(v int32) *ListDeptUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeptUsersResponse) SetBody(v *ListDeptUsersResponseBody) *ListDeptUsersResponse {
	s.Body = v
	return s
}

type ListParentByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListParentByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptHeaders) GoString() string {
	return s.String()
}

func (s *ListParentByDeptHeaders) SetCommonHeaders(v map[string]*string) *ListParentByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListParentByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *ListParentByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListParentByDeptRequest struct {
	DepartmentId *int64  `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	SubCorpId    *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListParentByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptRequest) GoString() string {
	return s.String()
}

func (s *ListParentByDeptRequest) SetDepartmentId(v int64) *ListParentByDeptRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListParentByDeptRequest) SetSubCorpId(v string) *ListParentByDeptRequest {
	s.SubCorpId = &v
	return s
}

type ListParentByDeptResponseBody struct {
	DepartmentIdList []*int64 `json:"departmentIdList,omitempty" xml:"departmentIdList,omitempty" type:"Repeated"`
}

func (s ListParentByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *ListParentByDeptResponseBody) SetDepartmentIdList(v []*int64) *ListParentByDeptResponseBody {
	s.DepartmentIdList = v
	return s
}

type ListParentByDeptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListParentByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParentByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParentByDeptResponse) GoString() string {
	return s.String()
}

func (s *ListParentByDeptResponse) SetHeaders(v map[string]*string) *ListParentByDeptResponse {
	s.Headers = v
	return s
}

func (s *ListParentByDeptResponse) SetStatusCode(v int32) *ListParentByDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParentByDeptResponse) SetBody(v *ListParentByDeptResponseBody) *ListParentByDeptResponse {
	s.Body = v
	return s
}

type ListParentByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListParentByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserHeaders) GoString() string {
	return s.String()
}

func (s *ListParentByUserHeaders) SetCommonHeaders(v map[string]*string) *ListParentByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListParentByUserHeaders) SetXAcsDingtalkAccessToken(v string) *ListParentByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListParentByUserRequest struct {
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListParentByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserRequest) GoString() string {
	return s.String()
}

func (s *ListParentByUserRequest) SetSubCorpId(v string) *ListParentByUserRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListParentByUserRequest) SetUserId(v string) *ListParentByUserRequest {
	s.UserId = &v
	return s
}

type ListParentByUserResponseBody struct {
	DepartmentIdList []*int64 `json:"departmentIdList,omitempty" xml:"departmentIdList,omitempty" type:"Repeated"`
}

func (s ListParentByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListParentByUserResponseBody) SetDepartmentIdList(v []*int64) *ListParentByUserResponseBody {
	s.DepartmentIdList = v
	return s
}

type ListParentByUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListParentByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListParentByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParentByUserResponse) GoString() string {
	return s.String()
}

func (s *ListParentByUserResponse) SetHeaders(v map[string]*string) *ListParentByUserResponse {
	s.Headers = v
	return s
}

func (s *ListParentByUserResponse) SetStatusCode(v int32) *ListParentByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParentByUserResponse) SetBody(v *ListParentByUserResponseBody) *ListParentByUserResponse {
	s.Body = v
	return s
}

type ListResidentDeptUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentDeptUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersHeaders) SetCommonHeaders(v map[string]*string) *ListResidentDeptUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentDeptUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentDeptUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentDeptUsersRequest struct {
	Cursor    *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Role      *string `json:"role,omitempty" xml:"role,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListResidentDeptUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersRequest) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersRequest) SetCursor(v int64) *ListResidentDeptUsersRequest {
	s.Cursor = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetRole(v string) *ListResidentDeptUsersRequest {
	s.Role = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetSize(v int32) *ListResidentDeptUsersRequest {
	s.Size = &v
	return s
}

func (s *ListResidentDeptUsersRequest) SetSubCorpId(v string) *ListResidentDeptUsersRequest {
	s.SubCorpId = &v
	return s
}

type ListResidentDeptUsersResponseBody struct {
	HasMore    *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextCursor *int64                                       `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	UserList   []*ListResidentDeptUsersResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListResidentDeptUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBody) SetHasMore(v bool) *ListResidentDeptUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListResidentDeptUsersResponseBody) SetNextCursor(v int64) *ListResidentDeptUsersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListResidentDeptUsersResponseBody) SetUserList(v []*ListResidentDeptUsersResponseBodyUserList) *ListResidentDeptUsersResponseBody {
	s.UserList = v
	return s
}

type ListResidentDeptUsersResponseBodyUserList struct {
	Feature *string                                           `json:"feature,omitempty" xml:"feature,omitempty"`
	Name    *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Roles   []*ListResidentDeptUsersResponseBodyUserListRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	UnionId *string                                           `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string                                           `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListResidentDeptUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBodyUserList) SetFeature(v string) *ListResidentDeptUsersResponseBodyUserList {
	s.Feature = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserList) SetName(v string) *ListResidentDeptUsersResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserList) SetRoles(v []*ListResidentDeptUsersResponseBodyUserListRoles) *ListResidentDeptUsersResponseBodyUserList {
	s.Roles = v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserList) SetUnionId(v string) *ListResidentDeptUsersResponseBodyUserList {
	s.UnionId = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserList) SetUserId(v string) *ListResidentDeptUsersResponseBodyUserList {
	s.UserId = &v
	return s
}

type ListResidentDeptUsersResponseBodyUserListRoles struct {
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagId   *int64  `json:"tagId,omitempty" xml:"tagId,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s ListResidentDeptUsersResponseBodyUserListRoles) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponseBodyUserListRoles) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponseBodyUserListRoles) SetTagCode(v string) *ListResidentDeptUsersResponseBodyUserListRoles {
	s.TagCode = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserListRoles) SetTagId(v int64) *ListResidentDeptUsersResponseBodyUserListRoles {
	s.TagId = &v
	return s
}

func (s *ListResidentDeptUsersResponseBodyUserListRoles) SetTagName(v string) *ListResidentDeptUsersResponseBodyUserListRoles {
	s.TagName = &v
	return s
}

type ListResidentDeptUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResidentDeptUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentDeptUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentDeptUsersResponse) GoString() string {
	return s.String()
}

func (s *ListResidentDeptUsersResponse) SetHeaders(v map[string]*string) *ListResidentDeptUsersResponse {
	s.Headers = v
	return s
}

func (s *ListResidentDeptUsersResponse) SetStatusCode(v int32) *ListResidentDeptUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResidentDeptUsersResponse) SetBody(v *ListResidentDeptUsersResponseBody) *ListResidentDeptUsersResponse {
	s.Body = v
	return s
}

type ListResidentSubDeptsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentSubDeptsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsHeaders) SetCommonHeaders(v map[string]*string) *ListResidentSubDeptsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentSubDeptsHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentSubDeptsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentSubDeptsRequest struct {
	Cursor    *int64  `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListResidentSubDeptsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsRequest) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsRequest) SetCursor(v int64) *ListResidentSubDeptsRequest {
	s.Cursor = &v
	return s
}

func (s *ListResidentSubDeptsRequest) SetSize(v int32) *ListResidentSubDeptsRequest {
	s.Size = &v
	return s
}

func (s *ListResidentSubDeptsRequest) SetSubCorpId(v string) *ListResidentSubDeptsRequest {
	s.SubCorpId = &v
	return s
}

type ListResidentSubDeptsResponseBody struct {
	DepartmentList []*ListResidentSubDeptsResponseBodyDepartmentList `json:"departmentList,omitempty" xml:"departmentList,omitempty" type:"Repeated"`
	HasMore        *bool                                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextCursor     *int64                                            `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Total          *int64                                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListResidentSubDeptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponseBody) SetDepartmentList(v []*ListResidentSubDeptsResponseBodyDepartmentList) *ListResidentSubDeptsResponseBody {
	s.DepartmentList = v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetHasMore(v bool) *ListResidentSubDeptsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetNextCursor(v int64) *ListResidentSubDeptsResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListResidentSubDeptsResponseBody) SetTotal(v int64) *ListResidentSubDeptsResponseBody {
	s.Total = &v
	return s
}

type ListResidentSubDeptsResponseBodyDepartmentList struct {
	DepartmentId      *int64  `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName    *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	SuperDepartmentId *int64  `json:"superDepartmentId,omitempty" xml:"superDepartmentId,omitempty"`
}

func (s ListResidentSubDeptsResponseBodyDepartmentList) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponseBodyDepartmentList) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponseBodyDepartmentList) SetDepartmentId(v int64) *ListResidentSubDeptsResponseBodyDepartmentList {
	s.DepartmentId = &v
	return s
}

func (s *ListResidentSubDeptsResponseBodyDepartmentList) SetDepartmentName(v string) *ListResidentSubDeptsResponseBodyDepartmentList {
	s.DepartmentName = &v
	return s
}

func (s *ListResidentSubDeptsResponseBodyDepartmentList) SetSuperDepartmentId(v int64) *ListResidentSubDeptsResponseBodyDepartmentList {
	s.SuperDepartmentId = &v
	return s
}

type ListResidentSubDeptsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResidentSubDeptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentSubDeptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentSubDeptsResponse) GoString() string {
	return s.String()
}

func (s *ListResidentSubDeptsResponse) SetHeaders(v map[string]*string) *ListResidentSubDeptsResponse {
	s.Headers = v
	return s
}

func (s *ListResidentSubDeptsResponse) SetStatusCode(v int32) *ListResidentSubDeptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResidentSubDeptsResponse) SetBody(v *ListResidentSubDeptsResponseBody) *ListResidentSubDeptsResponse {
	s.Body = v
	return s
}

type ListResidentUserInfosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListResidentUserInfosHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosHeaders) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosHeaders) SetCommonHeaders(v map[string]*string) *ListResidentUserInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResidentUserInfosHeaders) SetXAcsDingtalkAccessToken(v string) *ListResidentUserInfosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListResidentUserInfosRequest struct {
	SubCorpId *string   `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	UserIds   []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ListResidentUserInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosRequest) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosRequest) SetSubCorpId(v string) *ListResidentUserInfosRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentUserInfosRequest) SetUserIds(v []*string) *ListResidentUserInfosRequest {
	s.UserIds = v
	return s
}

type ListResidentUserInfosShrinkRequest struct {
	SubCorpId     *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s ListResidentUserInfosShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosShrinkRequest) SetSubCorpId(v string) *ListResidentUserInfosShrinkRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListResidentUserInfosShrinkRequest) SetUserIdsShrink(v string) *ListResidentUserInfosShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

type ListResidentUserInfosResponseBody struct {
	UserList []*ListResidentUserInfosResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListResidentUserInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBody) SetUserList(v []*ListResidentUserInfosResponseBodyUserList) *ListResidentUserInfosResponseBody {
	s.UserList = v
	return s
}

type ListResidentUserInfosResponseBodyUserList struct {
	Feature  *string                                           `json:"feature,omitempty" xml:"feature,omitempty"`
	Roles    []*ListResidentUserInfosResponseBodyUserListRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
	UnionId  *string                                           `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId   *string                                           `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *string                                           `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListResidentUserInfosResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBodyUserList) SetFeature(v string) *ListResidentUserInfosResponseBodyUserList {
	s.Feature = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserList) SetRoles(v []*ListResidentUserInfosResponseBodyUserListRoles) *ListResidentUserInfosResponseBodyUserList {
	s.Roles = v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserList) SetUnionId(v string) *ListResidentUserInfosResponseBodyUserList {
	s.UnionId = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserList) SetUserId(v string) *ListResidentUserInfosResponseBodyUserList {
	s.UserId = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserList) SetUserName(v string) *ListResidentUserInfosResponseBodyUserList {
	s.UserName = &v
	return s
}

type ListResidentUserInfosResponseBodyUserListRoles struct {
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagId   *int64  `json:"tagId,omitempty" xml:"tagId,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s ListResidentUserInfosResponseBodyUserListRoles) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponseBodyUserListRoles) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponseBodyUserListRoles) SetTagCode(v string) *ListResidentUserInfosResponseBodyUserListRoles {
	s.TagCode = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserListRoles) SetTagId(v int64) *ListResidentUserInfosResponseBodyUserListRoles {
	s.TagId = &v
	return s
}

func (s *ListResidentUserInfosResponseBodyUserListRoles) SetTagName(v string) *ListResidentUserInfosResponseBodyUserListRoles {
	s.TagName = &v
	return s
}

type ListResidentUserInfosResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResidentUserInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResidentUserInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResidentUserInfosResponse) GoString() string {
	return s.String()
}

func (s *ListResidentUserInfosResponse) SetHeaders(v map[string]*string) *ListResidentUserInfosResponse {
	s.Headers = v
	return s
}

func (s *ListResidentUserInfosResponse) SetStatusCode(v int32) *ListResidentUserInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResidentUserInfosResponse) SetBody(v *ListResidentUserInfosResponseBody) *ListResidentUserInfosResponse {
	s.Body = v
	return s
}

type ListSimpleUsersByRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSimpleUsersByRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleHeaders) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleHeaders) SetCommonHeaders(v map[string]*string) *ListSimpleUsersByRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSimpleUsersByRoleHeaders) SetXAcsDingtalkAccessToken(v string) *ListSimpleUsersByRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSimpleUsersByRoleRequest struct {
	Offset    *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	RoleId    *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
	Size      *int32  `json:"size,omitempty" xml:"size,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListSimpleUsersByRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleRequest) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleRequest) SetOffset(v int64) *ListSimpleUsersByRoleRequest {
	s.Offset = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetRoleId(v int64) *ListSimpleUsersByRoleRequest {
	s.RoleId = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetSize(v int32) *ListSimpleUsersByRoleRequest {
	s.Size = &v
	return s
}

func (s *ListSimpleUsersByRoleRequest) SetSubCorpId(v string) *ListSimpleUsersByRoleRequest {
	s.SubCorpId = &v
	return s
}

type ListSimpleUsersByRoleResponseBody struct {
	HasMore    *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextCursor *int64                                       `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	UserList   []*ListSimpleUsersByRoleResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListSimpleUsersByRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponseBody) SetHasMore(v bool) *ListSimpleUsersByRoleResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBody) SetNextCursor(v int64) *ListSimpleUsersByRoleResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBody) SetUserList(v []*ListSimpleUsersByRoleResponseBodyUserList) *ListSimpleUsersByRoleResponseBody {
	s.UserList = v
	return s
}

type ListSimpleUsersByRoleResponseBodyUserList struct {
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListSimpleUsersByRoleResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponseBodyUserList) SetJobNumber(v string) *ListSimpleUsersByRoleResponseBodyUserList {
	s.JobNumber = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyUserList) SetName(v string) *ListSimpleUsersByRoleResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyUserList) SetUnionId(v string) *ListSimpleUsersByRoleResponseBodyUserList {
	s.UnionId = &v
	return s
}

func (s *ListSimpleUsersByRoleResponseBodyUserList) SetUserId(v string) *ListSimpleUsersByRoleResponseBodyUserList {
	s.UserId = &v
	return s
}

type ListSimpleUsersByRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSimpleUsersByRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSimpleUsersByRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSimpleUsersByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListSimpleUsersByRoleResponse) SetHeaders(v map[string]*string) *ListSimpleUsersByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListSimpleUsersByRoleResponse) SetStatusCode(v int32) *ListSimpleUsersByRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSimpleUsersByRoleResponse) SetBody(v *ListSimpleUsersByRoleResponseBody) *ListSimpleUsersByRoleResponse {
	s.Body = v
	return s
}

type ListSubCorpsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubCorpsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsHeaders) GoString() string {
	return s.String()
}

func (s *ListSubCorpsHeaders) SetCommonHeaders(v map[string]*string) *ListSubCorpsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubCorpsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubCorpsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubCorpsRequest struct {
	IsOnlyDirect *bool   `json:"isOnlyDirect,omitempty" xml:"isOnlyDirect,omitempty"`
	SubCorpId    *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	Types        *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListSubCorpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsRequest) GoString() string {
	return s.String()
}

func (s *ListSubCorpsRequest) SetIsOnlyDirect(v bool) *ListSubCorpsRequest {
	s.IsOnlyDirect = &v
	return s
}

func (s *ListSubCorpsRequest) SetSubCorpId(v string) *ListSubCorpsRequest {
	s.SubCorpId = &v
	return s
}

func (s *ListSubCorpsRequest) SetTypes(v string) *ListSubCorpsRequest {
	s.Types = &v
	return s
}

type ListSubCorpsResponseBody struct {
	CorpList []*ListSubCorpsResponseBodyCorpList `json:"corpList,omitempty" xml:"corpList,omitempty" type:"Repeated"`
}

func (s ListSubCorpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponseBody) SetCorpList(v []*ListSubCorpsResponseBodyCorpList) *ListSubCorpsResponseBody {
	s.CorpList = v
	return s
}

type ListSubCorpsResponseBodyCorpList struct {
	CorpId         *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorpName       *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	Industry       *string `json:"industry,omitempty" xml:"industry,omitempty"`
	IndustryCode   *int32  `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	RegionId       *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RegionLocation *string `json:"regionLocation,omitempty" xml:"regionLocation,omitempty"`
	RegionType     *string `json:"regionType,omitempty" xml:"regionType,omitempty"`
}

func (s ListSubCorpsResponseBodyCorpList) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponseBodyCorpList) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponseBodyCorpList) SetCorpId(v string) *ListSubCorpsResponseBodyCorpList {
	s.CorpId = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetCorpName(v string) *ListSubCorpsResponseBodyCorpList {
	s.CorpName = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetIndustry(v string) *ListSubCorpsResponseBodyCorpList {
	s.Industry = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetIndustryCode(v int32) *ListSubCorpsResponseBodyCorpList {
	s.IndustryCode = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetRegionId(v string) *ListSubCorpsResponseBodyCorpList {
	s.RegionId = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetRegionLocation(v string) *ListSubCorpsResponseBodyCorpList {
	s.RegionLocation = &v
	return s
}

func (s *ListSubCorpsResponseBodyCorpList) SetRegionType(v string) *ListSubCorpsResponseBodyCorpList {
	s.RegionType = &v
	return s
}

type ListSubCorpsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubCorpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubCorpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubCorpsResponse) GoString() string {
	return s.String()
}

func (s *ListSubCorpsResponse) SetHeaders(v map[string]*string) *ListSubCorpsResponse {
	s.Headers = v
	return s
}

func (s *ListSubCorpsResponse) SetStatusCode(v int32) *ListSubCorpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubCorpsResponse) SetBody(v *ListSubCorpsResponseBody) *ListSubCorpsResponse {
	s.Body = v
	return s
}

type ListSubDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptHeaders) GoString() string {
	return s.String()
}

func (s *ListSubDeptHeaders) SetCommonHeaders(v map[string]*string) *ListSubDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubDeptHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubDeptRequest struct {
	Language  *string `json:"language,omitempty" xml:"language,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListSubDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptRequest) GoString() string {
	return s.String()
}

func (s *ListSubDeptRequest) SetLanguage(v string) *ListSubDeptRequest {
	s.Language = &v
	return s
}

func (s *ListSubDeptRequest) SetSubCorpId(v string) *ListSubDeptRequest {
	s.SubCorpId = &v
	return s
}

type ListSubDeptResponseBody struct {
	Result []*ListSubDeptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSubDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponseBody) SetResult(v []*ListSubDeptResponseBodyResult) *ListSubDeptResponseBody {
	s.Result = v
	return s
}

type ListSubDeptResponseBodyResult struct {
	DepartmentId *int64  `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListSubDeptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponseBodyResult) SetDepartmentId(v int64) *ListSubDeptResponseBodyResult {
	s.DepartmentId = &v
	return s
}

func (s *ListSubDeptResponseBodyResult) SetName(v string) *ListSubDeptResponseBodyResult {
	s.Name = &v
	return s
}

type ListSubDeptResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptResponse) GoString() string {
	return s.String()
}

func (s *ListSubDeptResponse) SetHeaders(v map[string]*string) *ListSubDeptResponse {
	s.Headers = v
	return s
}

func (s *ListSubDeptResponse) SetStatusCode(v int32) *ListSubDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubDeptResponse) SetBody(v *ListSubDeptResponseBody) *ListSubDeptResponse {
	s.Body = v
	return s
}

type ListSubDeptIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubDeptIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsHeaders) SetCommonHeaders(v map[string]*string) *ListSubDeptIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubDeptIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubDeptIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubDeptIdsRequest struct {
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListSubDeptIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsRequest) SetSubCorpId(v string) *ListSubDeptIdsRequest {
	s.SubCorpId = &v
	return s
}

type ListSubDeptIdsResponseBody struct {
	DepartmentIdList []*int64 `json:"departmentIdList,omitempty" xml:"departmentIdList,omitempty" type:"Repeated"`
}

func (s ListSubDeptIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsResponseBody) SetDepartmentIdList(v []*int64) *ListSubDeptIdsResponseBody {
	s.DepartmentIdList = v
	return s
}

type ListSubDeptIdsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSubDeptIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubDeptIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubDeptIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSubDeptIdsResponse) SetHeaders(v map[string]*string) *ListSubDeptIdsResponse {
	s.Headers = v
	return s
}

func (s *ListSubDeptIdsResponse) SetStatusCode(v int32) *ListSubDeptIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubDeptIdsResponse) SetBody(v *ListSubDeptIdsResponseBody) *ListSubDeptIdsResponse {
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

func (client *Client) GetDeptWithOptions(departmentId *string, request *GetDeptRequest, headers *GetDeptHeaders, runtime *util.RuntimeOptions) (_result *GetDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("GetDept"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/deptartments/" + tea.StringValue(departmentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDept(departmentId *string, request *GetDeptRequest) (_result *GetDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeptHeaders{}
	_result = &GetDeptResponse{}
	_body, _err := client.GetDeptWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResidentDeptWithOptions(departmentId *string, request *GetResidentDeptRequest, headers *GetResidentDeptHeaders, runtime *util.RuntimeOptions) (_result *GetResidentDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("GetResidentDept"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/residentDepartments/departments/" + tea.StringValue(departmentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResidentDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResidentDept(departmentId *string, request *GetResidentDeptRequest) (_result *GetResidentDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentDeptHeaders{}
	_result = &GetResidentDeptResponse{}
	_body, _err := client.GetResidentDeptWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResidentUserInfoWithOptions(departmentId *string, userId *string, request *GetResidentUserInfoRequest, headers *GetResidentUserInfoHeaders, runtime *util.RuntimeOptions) (_result *GetResidentUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("GetResidentUserInfo"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/residentDepartments/" + tea.StringValue(departmentId) + "/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResidentUserInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResidentUserInfo(departmentId *string, userId *string, request *GetResidentUserInfoRequest) (_result *GetResidentUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentUserInfoHeaders{}
	_result = &GetResidentUserInfoResponse{}
	_body, _err := client.GetResidentUserInfoWithOptions(departmentId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(userId *string, request *GetUserRequest, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("GetUser"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/users/getByUserId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(userId *string, request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserByUnionIdWithOptions(request *GetUserByUnionIdRequest, headers *GetUserByUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetUserByUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("GetUserByUnionId"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/users/getByUnionId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserByUnionIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByUnionId(request *GetUserByUnionIdRequest) (_result *GetUserByUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserByUnionIdHeaders{}
	_result = &GetUserByUnionIdResponse{}
	_body, _err := client.GetUserByUnionIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVillageOrgInfoWithOptions(subCorpId *string, headers *GetVillageOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *GetVillageOrgInfoResponse, _err error) {
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
		Action:      tea.String("GetVillageOrgInfo"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/corps/" + tea.StringValue(subCorpId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVillageOrgInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVillageOrgInfo(subCorpId *string) (_result *GetVillageOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetVillageOrgInfoHeaders{}
	_result = &GetVillageOrgInfoResponse{}
	_body, _err := client.GetVillageOrgInfoWithOptions(subCorpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptSimpleUsersWithOptions(departmentId *string, request *ListDeptSimpleUsersRequest, headers *ListDeptSimpleUsersHeaders, runtime *util.RuntimeOptions) (_result *ListDeptSimpleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainAccessLimit)) {
		query["containAccessLimit"] = request.ContainAccessLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["orderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListDeptSimpleUsers"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/" + tea.StringValue(departmentId) + "/simpleUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeptSimpleUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptSimpleUsers(departmentId *string, request *ListDeptSimpleUsersRequest) (_result *ListDeptSimpleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptSimpleUsersHeaders{}
	_result = &ListDeptSimpleUsersResponse{}
	_body, _err := client.ListDeptSimpleUsersWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptUserIdsWithOptions(departmentId *string, request *ListDeptUserIdsRequest, headers *ListDeptUserIdsHeaders, runtime *util.RuntimeOptions) (_result *ListDeptUserIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListDeptUserIds"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/" + tea.StringValue(departmentId) + "/userIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeptUserIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptUserIds(departmentId *string, request *ListDeptUserIdsRequest) (_result *ListDeptUserIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptUserIdsHeaders{}
	_result = &ListDeptUserIdsResponse{}
	_body, _err := client.ListDeptUserIdsWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeptUsersWithOptions(departmentId *string, request *ListDeptUsersRequest, headers *ListDeptUsersHeaders, runtime *util.RuntimeOptions) (_result *ListDeptUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainAccessLimit)) {
		query["containAccessLimit"] = request.ContainAccessLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["orderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListDeptUsers"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/" + tea.StringValue(departmentId) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeptUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeptUsers(departmentId *string, request *ListDeptUsersRequest) (_result *ListDeptUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListDeptUsersHeaders{}
	_result = &ListDeptUsersResponse{}
	_body, _err := client.ListDeptUsersWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParentByDeptWithOptions(request *ListParentByDeptRequest, headers *ListParentByDeptHeaders, runtime *util.RuntimeOptions) (_result *ListParentByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListParentByDept"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/listParentByDepartment"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParentByDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParentByDept(request *ListParentByDeptRequest) (_result *ListParentByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListParentByDeptHeaders{}
	_result = &ListParentByDeptResponse{}
	_body, _err := client.ListParentByDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListParentByUserWithOptions(request *ListParentByUserRequest, headers *ListParentByUserHeaders, runtime *util.RuntimeOptions) (_result *ListParentByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListParentByUser"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/listParentByUser"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParentByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParentByUser(request *ListParentByUserRequest) (_result *ListParentByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListParentByUserHeaders{}
	_result = &ListParentByUserResponse{}
	_body, _err := client.ListParentByUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentDeptUsersWithOptions(departmentId *string, request *ListResidentDeptUsersRequest, headers *ListResidentDeptUsersHeaders, runtime *util.RuntimeOptions) (_result *ListResidentDeptUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListResidentDeptUsers"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/residentDepartments/" + tea.StringValue(departmentId) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResidentDeptUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentDeptUsers(departmentId *string, request *ListResidentDeptUsersRequest) (_result *ListResidentDeptUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentDeptUsersHeaders{}
	_result = &ListResidentDeptUsersResponse{}
	_body, _err := client.ListResidentDeptUsersWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentSubDeptsWithOptions(departmentId *string, request *ListResidentSubDeptsRequest, headers *ListResidentSubDeptsHeaders, runtime *util.RuntimeOptions) (_result *ListResidentSubDeptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListResidentSubDepts"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/residentDepartments/" + tea.StringValue(departmentId) + "/subDepartments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResidentSubDeptsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentSubDepts(departmentId *string, request *ListResidentSubDeptsRequest) (_result *ListResidentSubDeptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentSubDeptsHeaders{}
	_result = &ListResidentSubDeptsResponse{}
	_body, _err := client.ListResidentSubDeptsWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResidentUserInfosWithOptions(tmpReq *ListResidentUserInfosRequest, headers *ListResidentUserInfosHeaders, runtime *util.RuntimeOptions) (_result *ListResidentUserInfosResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResidentUserInfosShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIds)) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, tea.String("userIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdsShrink)) {
		query["userIds"] = request.UserIdsShrink
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
		Action:      tea.String("ListResidentUserInfos"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/residentUsers/getByUserIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResidentUserInfosResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResidentUserInfos(request *ListResidentUserInfosRequest) (_result *ListResidentUserInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResidentUserInfosHeaders{}
	_result = &ListResidentUserInfosResponse{}
	_body, _err := client.ListResidentUserInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSimpleUsersByRoleWithOptions(request *ListSimpleUsersByRoleRequest, headers *ListSimpleUsersByRoleHeaders, runtime *util.RuntimeOptions) (_result *ListSimpleUsersByRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["roleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListSimpleUsersByRole"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/users/listByRole"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSimpleUsersByRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSimpleUsersByRole(request *ListSimpleUsersByRoleRequest) (_result *ListSimpleUsersByRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSimpleUsersByRoleHeaders{}
	_result = &ListSimpleUsersByRoleResponse{}
	_body, _err := client.ListSimpleUsersByRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubCorpsWithOptions(request *ListSubCorpsRequest, headers *ListSubCorpsHeaders, runtime *util.RuntimeOptions) (_result *ListSubCorpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsOnlyDirect)) {
		query["isOnlyDirect"] = request.IsOnlyDirect
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
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
		Action:      tea.String("ListSubCorps"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/corps/subCorps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubCorpsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubCorps(request *ListSubCorpsRequest) (_result *ListSubCorpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubCorpsHeaders{}
	_result = &ListSubCorpsResponse{}
	_body, _err := client.ListSubCorpsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubDeptWithOptions(departmentId *string, request *ListSubDeptRequest, headers *ListSubDeptHeaders, runtime *util.RuntimeOptions) (_result *ListSubDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListSubDept"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/" + tea.StringValue(departmentId) + "/subDepartments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubDept(departmentId *string, request *ListSubDeptRequest) (_result *ListSubDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubDeptHeaders{}
	_result = &ListSubDeptResponse{}
	_body, _err := client.ListSubDeptWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubDeptIdsWithOptions(departmentId *string, request *ListSubDeptIdsRequest, headers *ListSubDeptIdsHeaders, runtime *util.RuntimeOptions) (_result *ListSubDeptIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		query["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("ListSubDeptIds"),
		Version:     tea.String("village_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/village/departments/" + tea.StringValue(departmentId) + "/subDepartmentIds"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubDeptIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubDeptIds(departmentId *string, request *ListSubDeptIdsRequest) (_result *ListSubDeptIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubDeptIdsHeaders{}
	_result = &ListSubDeptIdsResponse{}
	_body, _err := client.ListSubDeptIdsWithOptions(departmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
