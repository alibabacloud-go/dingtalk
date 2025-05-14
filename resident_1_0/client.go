// This file is auto-generated, don't edit it. Thanks.
package resident_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddPointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddPointHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddPointHeaders) GoString() string {
	return s.String()
}

func (s *AddPointHeaders) SetCommonHeaders(v map[string]*string) *AddPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddPointHeaders) SetXAcsDingtalkAccessToken(v string) *AddPointHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddPointRequest struct {
	// example:
	//
	// 1634630147
	ActionTime *int64 `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// example:
	//
	// rule_1
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 发动态
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7645
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s AddPointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPointRequest) GoString() string {
	return s.String()
}

func (s *AddPointRequest) SetActionTime(v int64) *AddPointRequest {
	s.ActionTime = &v
	return s
}

func (s *AddPointRequest) SetIsCircle(v bool) *AddPointRequest {
	s.IsCircle = &v
	return s
}

func (s *AddPointRequest) SetRuleCode(v string) *AddPointRequest {
	s.RuleCode = &v
	return s
}

func (s *AddPointRequest) SetRuleName(v string) *AddPointRequest {
	s.RuleName = &v
	return s
}

func (s *AddPointRequest) SetScore(v int32) *AddPointRequest {
	s.Score = &v
	return s
}

func (s *AddPointRequest) SetUserId(v string) *AddPointRequest {
	s.UserId = &v
	return s
}

func (s *AddPointRequest) SetUuid(v string) *AddPointRequest {
	s.Uuid = &v
	return s
}

type AddPointResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddPointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPointResponse) GoString() string {
	return s.String()
}

func (s *AddPointResponse) SetHeaders(v map[string]*string) *AddPointResponse {
	s.Headers = v
	return s
}

func (s *AddPointResponse) SetStatusCode(v int32) *AddPointResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// 第一网格组
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// example:
	//
	// true
	IsResidenceGroup *bool `json:"isResidenceGroup,omitempty" xml:"isResidenceGroup,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s AddResidentDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentDepartmentRequest) GoString() string {
	return s.String()
}

func (s *AddResidentDepartmentRequest) SetDepartmentName(v string) *AddResidentDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetIsResidenceGroup(v bool) *AddResidentDepartmentRequest {
	s.IsResidenceGroup = &v
	return s
}

func (s *AddResidentDepartmentRequest) SetParentDepartmentId(v int64) *AddResidentDepartmentRequest {
	s.ParentDepartmentId = &v
	return s
}

type AddResidentDepartmentResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddResidentDepartmentResponse) SetStatusCode(v int32) *AddResidentDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResidentDepartmentResponse) SetBody(v *AddResidentDepartmentResponseBody) *AddResidentDepartmentResponse {
	s.Body = v
	return s
}

type AddResidentMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddResidentMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddResidentMemberHeaders) SetCommonHeaders(v map[string]*string) *AddResidentMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddResidentMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddResidentMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddResidentMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A栋
	ResidentAddInfo *AddResidentMemberRequestResidentAddInfo `json:"residentAddInfo,omitempty" xml:"residentAddInfo,omitempty" type:"Struct"`
}

func (s AddResidentMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberRequest) GoString() string {
	return s.String()
}

func (s *AddResidentMemberRequest) SetResidentAddInfo(v *AddResidentMemberRequestResidentAddInfo) *AddResidentMemberRequest {
	s.ResidentAddInfo = v
	return s
}

type AddResidentMemberRequestResidentAddInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 11112
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// true
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// example:
	//
	// {"startTime":1652358627106,"endTime":1652445027106}
	MemberDeptExtension map[string]interface{} `json:"memberDeptExtension,omitempty" xml:"memberDeptExtension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 148********
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
}

func (s AddResidentMemberRequestResidentAddInfo) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberRequestResidentAddInfo) GoString() string {
	return s.String()
}

func (s *AddResidentMemberRequestResidentAddInfo) SetDeptId(v int64) *AddResidentMemberRequestResidentAddInfo {
	s.DeptId = &v
	return s
}

func (s *AddResidentMemberRequestResidentAddInfo) SetIsPropertyOwner(v bool) *AddResidentMemberRequestResidentAddInfo {
	s.IsPropertyOwner = &v
	return s
}

func (s *AddResidentMemberRequestResidentAddInfo) SetMemberDeptExtension(v map[string]interface{}) *AddResidentMemberRequestResidentAddInfo {
	s.MemberDeptExtension = v
	return s
}

func (s *AddResidentMemberRequestResidentAddInfo) SetMobile(v string) *AddResidentMemberRequestResidentAddInfo {
	s.Mobile = &v
	return s
}

func (s *AddResidentMemberRequestResidentAddInfo) SetName(v string) *AddResidentMemberRequestResidentAddInfo {
	s.Name = &v
	return s
}

func (s *AddResidentMemberRequestResidentAddInfo) SetRelateType(v string) *AddResidentMemberRequestResidentAddInfo {
	s.RelateType = &v
	return s
}

type AddResidentMemberResponseBody struct {
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 10005
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddResidentMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentMemberResponseBody) SetStatus(v int32) *AddResidentMemberResponseBody {
	s.Status = &v
	return s
}

func (s *AddResidentMemberResponseBody) SetUnionId(v string) *AddResidentMemberResponseBody {
	s.UnionId = &v
	return s
}

func (s *AddResidentMemberResponseBody) SetUserId(v string) *AddResidentMemberResponseBody {
	s.UserId = &v
	return s
}

type AddResidentMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddResidentMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberResponse) GoString() string {
	return s.String()
}

func (s *AddResidentMemberResponse) SetHeaders(v map[string]*string) *AddResidentMemberResponse {
	s.Headers = v
	return s
}

func (s *AddResidentMemberResponse) SetStatusCode(v int32) *AddResidentMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResidentMemberResponse) SetBody(v *AddResidentMemberResponseBody) *AddResidentMemberResponse {
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
	// example:
	//
	// 美好社区创景街道万通公寓
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DepartmentId *int64                             `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	ExtField     []*AddResidentUsersRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsLeaseholder *bool `json:"isLeaseholder,omitempty" xml:"isLeaseholder,omitempty"`
	// example:
	//
	// 15612345678
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// SELF
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 王建国
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
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

func (s *AddResidentUsersRequest) SetDepartmentId(v int64) *AddResidentUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *AddResidentUsersRequest) SetExtField(v []*AddResidentUsersRequestExtField) *AddResidentUsersRequest {
	s.ExtField = v
	return s
}

func (s *AddResidentUsersRequest) SetIsLeaseholder(v bool) *AddResidentUsersRequest {
	s.IsLeaseholder = &v
	return s
}

func (s *AddResidentUsersRequest) SetMobile(v string) *AddResidentUsersRequest {
	s.Mobile = &v
	return s
}

func (s *AddResidentUsersRequest) SetRelateType(v string) *AddResidentUsersRequest {
	s.RelateType = &v
	return s
}

func (s *AddResidentUsersRequest) SetUserName(v string) *AddResidentUsersRequest {
	s.UserName = &v
	return s
}

type AddResidentUsersRequestExtField struct {
	// example:
	//
	// 性别
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 女
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
}

func (s AddResidentUsersRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s AddResidentUsersRequestExtField) GoString() string {
	return s.String()
}

func (s *AddResidentUsersRequestExtField) SetItemName(v string) *AddResidentUsersRequestExtField {
	s.ItemName = &v
	return s
}

func (s *AddResidentUsersRequestExtField) SetItemValue(v string) *AddResidentUsersRequestExtField {
	s.ItemValue = &v
	return s
}

type AddResidentUsersResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResidentUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddResidentUsersResponse) SetStatusCode(v int32) *AddResidentUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResidentUsersResponse) SetBody(v *AddResidentUsersResponseBody) *AddResidentUsersResponse {
	s.Body = v
	return s
}

type CreateResidentBlackBoardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateResidentBlackBoardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateResidentBlackBoardHeaders) GoString() string {
	return s.String()
}

func (s *CreateResidentBlackBoardHeaders) SetCommonHeaders(v map[string]*string) *CreateResidentBlackBoardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateResidentBlackBoardHeaders) SetXAcsDingtalkAccessToken(v string) *CreateResidentBlackBoardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateResidentBlackBoardRequest struct {
	// This parameter is required.
	Context  *string `json:"context,omitempty" xml:"context,omitempty"`
	MediaId  *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	SendTime *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateResidentBlackBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResidentBlackBoardRequest) GoString() string {
	return s.String()
}

func (s *CreateResidentBlackBoardRequest) SetContext(v string) *CreateResidentBlackBoardRequest {
	s.Context = &v
	return s
}

func (s *CreateResidentBlackBoardRequest) SetMediaId(v string) *CreateResidentBlackBoardRequest {
	s.MediaId = &v
	return s
}

func (s *CreateResidentBlackBoardRequest) SetSendTime(v string) *CreateResidentBlackBoardRequest {
	s.SendTime = &v
	return s
}

func (s *CreateResidentBlackBoardRequest) SetTitle(v string) *CreateResidentBlackBoardRequest {
	s.Title = &v
	return s
}

type CreateResidentBlackBoardResponseBody struct {
	BlackBoardId *string `json:"blackBoardId,omitempty" xml:"blackBoardId,omitempty"`
}

func (s CreateResidentBlackBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResidentBlackBoardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResidentBlackBoardResponseBody) SetBlackBoardId(v string) *CreateResidentBlackBoardResponseBody {
	s.BlackBoardId = &v
	return s
}

type CreateResidentBlackBoardResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResidentBlackBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResidentBlackBoardResponse) GoString() string {
	return s.String()
}

func (s *CreateResidentBlackBoardResponse) SetHeaders(v map[string]*string) *CreateResidentBlackBoardResponse {
	s.Headers = v
	return s
}

func (s *CreateResidentBlackBoardResponse) SetStatusCode(v int32) *CreateResidentBlackBoardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResidentBlackBoardResponse) SetBody(v *CreateResidentBlackBoardResponseBody) *CreateResidentBlackBoardResponse {
	s.Body = v
	return s
}

type CreateSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateSpaceHeaders) SetCommonHeaders(v map[string]*string) *CreateSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSpaceRequest struct {
	BillingArea  *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	Floor        *string  `json:"floor,omitempty" xml:"floor,omitempty"`
	HouseState   *int64   `json:"houseState,omitempty" xml:"houseState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A栋
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -7
	ParentDeptId *string `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// House
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// example:
	//
	// 2
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSpaceRequest) SetBillingArea(v float32) *CreateSpaceRequest {
	s.BillingArea = &v
	return s
}

func (s *CreateSpaceRequest) SetBuildingArea(v float32) *CreateSpaceRequest {
	s.BuildingArea = &v
	return s
}

func (s *CreateSpaceRequest) SetFloor(v string) *CreateSpaceRequest {
	s.Floor = &v
	return s
}

func (s *CreateSpaceRequest) SetHouseState(v int64) *CreateSpaceRequest {
	s.HouseState = &v
	return s
}

func (s *CreateSpaceRequest) SetName(v string) *CreateSpaceRequest {
	s.Name = &v
	return s
}

func (s *CreateSpaceRequest) SetParentDeptId(v string) *CreateSpaceRequest {
	s.ParentDeptId = &v
	return s
}

func (s *CreateSpaceRequest) SetTagCode(v string) *CreateSpaceRequest {
	s.TagCode = &v
	return s
}

func (s *CreateSpaceRequest) SetType(v string) *CreateSpaceRequest {
	s.Type = &v
	return s
}

type CreateSpaceResponseBody struct {
	// example:
	//
	// 10005
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CreateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponseBody) SetDeptId(v string) *CreateSpaceResponseBody {
	s.DeptId = &v
	return s
}

type CreateSpaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponse) SetHeaders(v map[string]*string) *CreateSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSpaceResponse) SetStatusCode(v int32) *CreateSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSpaceResponse) SetBody(v *CreateSpaceResponseBody) *CreateSpaceResponse {
	s.Body = v
	return s
}

type DeleteResidentBlackBoardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteResidentBlackBoardHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentBlackBoardHeaders) GoString() string {
	return s.String()
}

func (s *DeleteResidentBlackBoardHeaders) SetCommonHeaders(v map[string]*string) *DeleteResidentBlackBoardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteResidentBlackBoardHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteResidentBlackBoardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteResidentBlackBoardRequest struct {
	// This parameter is required.
	BlackboardId *string `json:"blackboardId,omitempty" xml:"blackboardId,omitempty"`
}

func (s DeleteResidentBlackBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentBlackBoardRequest) GoString() string {
	return s.String()
}

func (s *DeleteResidentBlackBoardRequest) SetBlackboardId(v string) *DeleteResidentBlackBoardRequest {
	s.BlackboardId = &v
	return s
}

type DeleteResidentBlackBoardResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteResidentBlackBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentBlackBoardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResidentBlackBoardResponseBody) SetSuccess(v bool) *DeleteResidentBlackBoardResponseBody {
	s.Success = &v
	return s
}

type DeleteResidentBlackBoardResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResidentBlackBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResidentBlackBoardResponse) GoString() string {
	return s.String()
}

func (s *DeleteResidentBlackBoardResponse) SetHeaders(v map[string]*string) *DeleteResidentBlackBoardResponse {
	s.Headers = v
	return s
}

func (s *DeleteResidentBlackBoardResponse) SetStatusCode(v int32) *DeleteResidentBlackBoardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResidentBlackBoardResponse) SetBody(v *DeleteResidentBlackBoardResponseBody) *DeleteResidentBlackBoardResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 12345
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResidentDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteResidentDepartmentResponse) SetStatusCode(v int32) *DeleteResidentDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResidentDepartmentResponse) SetBody(v *DeleteResidentDepartmentResponseBody) *DeleteResidentDepartmentResponse {
	s.Body = v
	return s
}

type DeleteSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSpaceHeaders) SetCommonHeaders(v map[string]*string) *DeleteSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSpaceRequest struct {
	// example:
	//
	// 忘川路1号
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
}

func (s DeleteSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpaceRequest) SetDeptIds(v []*int64) *DeleteSpaceRequest {
	s.DeptIds = v
	return s
}

type DeleteSpaceResponseBody struct {
	// This parameter is required.
	DelFailedDept []*DeleteSpaceResponseBodyDelFailedDept `json:"delFailedDept,omitempty" xml:"delFailedDept,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DelSuccessCount *bool `json:"delSuccessCount,omitempty" xml:"delSuccessCount,omitempty"`
}

func (s DeleteSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponseBody) SetDelFailedDept(v []*DeleteSpaceResponseBodyDelFailedDept) *DeleteSpaceResponseBody {
	s.DelFailedDept = v
	return s
}

func (s *DeleteSpaceResponseBody) SetDelSuccessCount(v bool) *DeleteSpaceResponseBody {
	s.DelSuccessCount = &v
	return s
}

type DeleteSpaceResponseBodyDelFailedDept struct {
	// This parameter is required.
	//
	// example:
	//
	// 122222
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s DeleteSpaceResponseBodyDelFailedDept) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponseBodyDelFailedDept) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponseBodyDelFailedDept) SetDeptId(v int64) *DeleteSpaceResponseBodyDelFailedDept {
	s.DeptId = &v
	return s
}

type DeleteSpaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponse) SetHeaders(v map[string]*string) *DeleteSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpaceResponse) SetStatusCode(v int32) *DeleteSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSpaceResponse) SetBody(v *DeleteSpaceResponseBody) *DeleteSpaceResponse {
	s.Body = v
	return s
}

type GetConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationIdHeaders) SetCommonHeaders(v map[string]*string) *GetConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chatd575783672bb40c005ba4e8b2*****ab
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
}

func (s GetConversationIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdRequest) GoString() string {
	return s.String()
}

func (s *GetConversationIdRequest) SetChatId(v string) *GetConversationIdRequest {
	s.ChatId = &v
	return s
}

type GetConversationIdResponseBody struct {
	// example:
	//
	// cidAX+2NwjqR3Y81Sxic5jtag==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponseBody) SetOpenConversationId(v string) *GetConversationIdResponseBody {
	s.OpenConversationId = &v
	return s
}

type GetConversationIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationIdResponse) GoString() string {
	return s.String()
}

func (s *GetConversationIdResponse) SetHeaders(v map[string]*string) *GetConversationIdResponse {
	s.Headers = v
	return s
}

func (s *GetConversationIdResponse) SetStatusCode(v int32) *GetConversationIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationIdResponse) SetBody(v *GetConversationIdResponseBody) *GetConversationIdResponse {
	s.Body = v
	return s
}

type GetIndustryTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetIndustryTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetIndustryTypeHeaders) GoString() string {
	return s.String()
}

func (s *GetIndustryTypeHeaders) SetCommonHeaders(v map[string]*string) *GetIndustryTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetIndustryTypeHeaders) SetXAcsDingtalkAccessToken(v string) *GetIndustryTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetIndustryTypeResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// RESIDENCE
	IndustryType *string `json:"industryType,omitempty" xml:"industryType,omitempty"`
}

func (s GetIndustryTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndustryTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndustryTypeResponseBody) SetIndustryType(v string) *GetIndustryTypeResponseBody {
	s.IndustryType = &v
	return s
}

type GetIndustryTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndustryTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndustryTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndustryTypeResponse) GoString() string {
	return s.String()
}

func (s *GetIndustryTypeResponse) SetHeaders(v map[string]*string) *GetIndustryTypeResponse {
	s.Headers = v
	return s
}

func (s *GetIndustryTypeResponse) SetStatusCode(v int32) *GetIndustryTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndustryTypeResponse) SetBody(v *GetIndustryTypeResponseBody) *GetIndustryTypeResponse {
	s.Body = v
	return s
}

type GetPropertyInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPropertyInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPropertyInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPropertyInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPropertyInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPropertyInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPropertyInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPropertyInfoRequest struct {
	PropertyCorpId *string `json:"propertyCorpId,omitempty" xml:"propertyCorpId,omitempty"`
}

func (s GetPropertyInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPropertyInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPropertyInfoRequest) SetPropertyCorpId(v string) *GetPropertyInfoRequest {
	s.PropertyCorpId = &v
	return s
}

type GetPropertyInfoResponseBody struct {
	AdminName           *string `json:"adminName,omitempty" xml:"adminName,omitempty"`
	AdminUserId         *string `json:"adminUserId,omitempty" xml:"adminUserId,omitempty"`
	Name                *string `json:"name,omitempty" xml:"name,omitempty"`
	OrgId               *int64  `json:"orgId,omitempty" xml:"orgId,omitempty"`
	UnifiedSocialCredit *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
}

func (s GetPropertyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPropertyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPropertyInfoResponseBody) SetAdminName(v string) *GetPropertyInfoResponseBody {
	s.AdminName = &v
	return s
}

func (s *GetPropertyInfoResponseBody) SetAdminUserId(v string) *GetPropertyInfoResponseBody {
	s.AdminUserId = &v
	return s
}

func (s *GetPropertyInfoResponseBody) SetName(v string) *GetPropertyInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetPropertyInfoResponseBody) SetOrgId(v int64) *GetPropertyInfoResponseBody {
	s.OrgId = &v
	return s
}

func (s *GetPropertyInfoResponseBody) SetUnifiedSocialCredit(v string) *GetPropertyInfoResponseBody {
	s.UnifiedSocialCredit = &v
	return s
}

type GetPropertyInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPropertyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPropertyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPropertyInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPropertyInfoResponse) SetHeaders(v map[string]*string) *GetPropertyInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPropertyInfoResponse) SetStatusCode(v int32) *GetPropertyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPropertyInfoResponse) SetBody(v *GetPropertyInfoResponseBody) *GetPropertyInfoResponse {
	s.Body = v
	return s
}

type GetResidentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentInfoHeaders) SetCommonHeaders(v map[string]*string) *GetResidentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentInfoRequest struct {
	ResidentCorpId *string `json:"residentCorpId,omitempty" xml:"residentCorpId,omitempty"`
}

func (s GetResidentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResidentInfoRequest) SetResidentCorpId(v string) *GetResidentInfoRequest {
	s.ResidentCorpId = &v
	return s
}

type GetResidentInfoResponseBody struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	AllUserGroupOpenConversationId *string `json:"allUserGroupOpenConversationId,omitempty" xml:"allUserGroupOpenConversationId,omitempty"`
	// This parameter is required.
	AllUserGroupOwnerUserId *string                                    `json:"allUserGroupOwnerUserId,omitempty" xml:"allUserGroupOwnerUserId,omitempty"`
	BuildingArea            *float32                                   `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	CityId                  *int32                                     `json:"cityId,omitempty" xml:"cityId,omitempty"`
	ContactMode             *int32                                     `json:"contactMode,omitempty" xml:"contactMode,omitempty"`
	CountyId                *int32                                     `json:"countyId,omitempty" xml:"countyId,omitempty"`
	DeliveryTime            *int64                                     `json:"deliveryTime,omitempty" xml:"deliveryTime,omitempty"`
	Location                *string                                    `json:"location,omitempty" xml:"location,omitempty"`
	Name                    *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	ProjectManager          *GetResidentInfoResponseBodyProjectManager `json:"projectManager,omitempty" xml:"projectManager,omitempty" type:"Struct"`
	// This parameter is required.
	PropertyDeptGroupOpenConversationId *string `json:"propertyDeptGroupOpenConversationId,omitempty" xml:"propertyDeptGroupOpenConversationId,omitempty"`
	// This parameter is required.
	PropertyDeptGroupOwnerUserId *string `json:"propertyDeptGroupOwnerUserId,omitempty" xml:"propertyDeptGroupOwnerUserId,omitempty"`
	ProvId                       *int64  `json:"provId,omitempty" xml:"provId,omitempty"`
	ScopeEast                    *string `json:"scopeEast,omitempty" xml:"scopeEast,omitempty"`
	ScopeNorth                   *string `json:"scopeNorth,omitempty" xml:"scopeNorth,omitempty"`
	ScopeSouth                   *string `json:"scopeSouth,omitempty" xml:"scopeSouth,omitempty"`
	ScopeWest                    *string `json:"scopeWest,omitempty" xml:"scopeWest,omitempty"`
	Telephone                    *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	TownId                       *int32  `json:"townId,omitempty" xml:"townId,omitempty"`
	Type                         *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetResidentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentInfoResponseBody) SetAddress(v string) *GetResidentInfoResponseBody {
	s.Address = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetAllUserGroupOpenConversationId(v string) *GetResidentInfoResponseBody {
	s.AllUserGroupOpenConversationId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetAllUserGroupOwnerUserId(v string) *GetResidentInfoResponseBody {
	s.AllUserGroupOwnerUserId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetBuildingArea(v float32) *GetResidentInfoResponseBody {
	s.BuildingArea = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetCityId(v int32) *GetResidentInfoResponseBody {
	s.CityId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetContactMode(v int32) *GetResidentInfoResponseBody {
	s.ContactMode = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetCountyId(v int32) *GetResidentInfoResponseBody {
	s.CountyId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetDeliveryTime(v int64) *GetResidentInfoResponseBody {
	s.DeliveryTime = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetLocation(v string) *GetResidentInfoResponseBody {
	s.Location = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetName(v string) *GetResidentInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetProjectManager(v *GetResidentInfoResponseBodyProjectManager) *GetResidentInfoResponseBody {
	s.ProjectManager = v
	return s
}

func (s *GetResidentInfoResponseBody) SetPropertyDeptGroupOpenConversationId(v string) *GetResidentInfoResponseBody {
	s.PropertyDeptGroupOpenConversationId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetPropertyDeptGroupOwnerUserId(v string) *GetResidentInfoResponseBody {
	s.PropertyDeptGroupOwnerUserId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetProvId(v int64) *GetResidentInfoResponseBody {
	s.ProvId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetScopeEast(v string) *GetResidentInfoResponseBody {
	s.ScopeEast = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetScopeNorth(v string) *GetResidentInfoResponseBody {
	s.ScopeNorth = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetScopeSouth(v string) *GetResidentInfoResponseBody {
	s.ScopeSouth = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetScopeWest(v string) *GetResidentInfoResponseBody {
	s.ScopeWest = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetTelephone(v string) *GetResidentInfoResponseBody {
	s.Telephone = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetTownId(v int32) *GetResidentInfoResponseBody {
	s.TownId = &v
	return s
}

func (s *GetResidentInfoResponseBody) SetType(v int32) *GetResidentInfoResponseBody {
	s.Type = &v
	return s
}

type GetResidentInfoResponseBodyProjectManager struct {
	Avatar   *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetResidentInfoResponseBodyProjectManager) String() string {
	return tea.Prettify(s)
}

func (s GetResidentInfoResponseBodyProjectManager) GoString() string {
	return s.String()
}

func (s *GetResidentInfoResponseBodyProjectManager) SetAvatar(v string) *GetResidentInfoResponseBodyProjectManager {
	s.Avatar = &v
	return s
}

func (s *GetResidentInfoResponseBodyProjectManager) SetUserId(v string) *GetResidentInfoResponseBodyProjectManager {
	s.UserId = &v
	return s
}

func (s *GetResidentInfoResponseBodyProjectManager) SetUserName(v string) *GetResidentInfoResponseBodyProjectManager {
	s.UserName = &v
	return s
}

type GetResidentInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResidentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResidentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResidentInfoResponse) SetHeaders(v map[string]*string) *GetResidentInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResidentInfoResponse) SetStatusCode(v int32) *GetResidentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResidentInfoResponse) SetBody(v *GetResidentInfoResponseBody) *GetResidentInfoResponse {
	s.Body = v
	return s
}

type GetResidentMembersInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetResidentMembersInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResidentMembersInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetResidentMembersInfoHeaders) SetCommonHeaders(v map[string]*string) *GetResidentMembersInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResidentMembersInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetResidentMembersInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetResidentMembersInfoRequest struct {
	// This parameter is required.
	ResidentCropId *string `json:"residentCropId,omitempty" xml:"residentCropId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetResidentMembersInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResidentMembersInfoRequest) GoString() string {
	return s.String()
}

func (s *GetResidentMembersInfoRequest) SetResidentCropId(v string) *GetResidentMembersInfoRequest {
	s.ResidentCropId = &v
	return s
}

func (s *GetResidentMembersInfoRequest) SetUserIdList(v []*string) *GetResidentMembersInfoRequest {
	s.UserIdList = v
	return s
}

type GetResidentMembersInfoResponseBody struct {
	ResidenceList []*GetResidentMembersInfoResponseBodyResidenceList `json:"residenceList,omitempty" xml:"residenceList,omitempty" type:"Repeated"`
}

func (s GetResidentMembersInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResidentMembersInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetResidentMembersInfoResponseBody) SetResidenceList(v []*GetResidentMembersInfoResponseBodyResidenceList) *GetResidentMembersInfoResponseBody {
	s.ResidenceList = v
	return s
}

type GetResidentMembersInfoResponseBodyResidenceList struct {
	Active          *bool   `json:"active,omitempty" xml:"active,omitempty"`
	ExtField        *string `json:"extField,omitempty" xml:"extField,omitempty"`
	IsPropertyOwner *bool   `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	RelateType      *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
}

func (s GetResidentMembersInfoResponseBodyResidenceList) String() string {
	return tea.Prettify(s)
}

func (s GetResidentMembersInfoResponseBodyResidenceList) GoString() string {
	return s.String()
}

func (s *GetResidentMembersInfoResponseBodyResidenceList) SetActive(v bool) *GetResidentMembersInfoResponseBodyResidenceList {
	s.Active = &v
	return s
}

func (s *GetResidentMembersInfoResponseBodyResidenceList) SetExtField(v string) *GetResidentMembersInfoResponseBodyResidenceList {
	s.ExtField = &v
	return s
}

func (s *GetResidentMembersInfoResponseBodyResidenceList) SetIsPropertyOwner(v bool) *GetResidentMembersInfoResponseBodyResidenceList {
	s.IsPropertyOwner = &v
	return s
}

func (s *GetResidentMembersInfoResponseBodyResidenceList) SetName(v string) *GetResidentMembersInfoResponseBodyResidenceList {
	s.Name = &v
	return s
}

func (s *GetResidentMembersInfoResponseBodyResidenceList) SetRelateType(v string) *GetResidentMembersInfoResponseBodyResidenceList {
	s.RelateType = &v
	return s
}

type GetResidentMembersInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResidentMembersInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResidentMembersInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResidentMembersInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResidentMembersInfoResponse) SetHeaders(v map[string]*string) *GetResidentMembersInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResidentMembersInfoResponse) SetStatusCode(v int32) *GetResidentMembersInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResidentMembersInfoResponse) SetBody(v *GetResidentMembersInfoResponseBody) *GetResidentMembersInfoResponse {
	s.Body = v
	return s
}

type GetSpaceIdByTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceIdByTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceIdByTypeHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceIdByTypeHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceIdByTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceIdByTypeHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceIdByTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceIdByTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PROPERTY_STAFF_DEPT
	DepartmentType *string `json:"departmentType,omitempty" xml:"departmentType,omitempty"`
}

func (s GetSpaceIdByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceIdByTypeRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceIdByTypeRequest) SetDepartmentType(v string) *GetSpaceIdByTypeRequest {
	s.DepartmentType = &v
	return s
}

type GetSpaceIdByTypeResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 12343
	ReferId *int64 `json:"referId,omitempty" xml:"referId,omitempty"`
}

func (s GetSpaceIdByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceIdByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceIdByTypeResponseBody) SetReferId(v int64) *GetSpaceIdByTypeResponseBody {
	s.ReferId = &v
	return s
}

type GetSpaceIdByTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceIdByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpaceIdByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceIdByTypeResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceIdByTypeResponse) SetHeaders(v map[string]*string) *GetSpaceIdByTypeResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceIdByTypeResponse) SetStatusCode(v int32) *GetSpaceIdByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceIdByTypeResponse) SetBody(v *GetSpaceIdByTypeResponseBody) *GetSpaceIdByTypeResponse {
	s.Body = v
	return s
}

type GetSpacesInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpacesInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpacesInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetSpacesInfoHeaders) SetCommonHeaders(v map[string]*string) *GetSpacesInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpacesInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpacesInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpacesInfoRequest struct {
	// This parameter is required.
	ReferIds []*int64 `json:"referIds,omitempty" xml:"referIds,omitempty" type:"Repeated"`
	// This parameter is required.
	ResidentCorpId *string `json:"residentCorpId,omitempty" xml:"residentCorpId,omitempty"`
}

func (s GetSpacesInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpacesInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSpacesInfoRequest) SetReferIds(v []*int64) *GetSpacesInfoRequest {
	s.ReferIds = v
	return s
}

func (s *GetSpacesInfoRequest) SetResidentCorpId(v string) *GetSpacesInfoRequest {
	s.ResidentCorpId = &v
	return s
}

type GetSpacesInfoResponseBody struct {
	SpaceList []*GetSpacesInfoResponseBodySpaceList `json:"spaceList,omitempty" xml:"spaceList,omitempty" type:"Repeated"`
}

func (s GetSpacesInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpacesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpacesInfoResponseBody) SetSpaceList(v []*GetSpacesInfoResponseBodySpaceList) *GetSpacesInfoResponseBody {
	s.SpaceList = v
	return s
}

type GetSpacesInfoResponseBodySpaceList struct {
	BillingArea   *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	BuildingArea  *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	Floor         *string  `json:"floor,omitempty" xml:"floor,omitempty"`
	HouseState    *int32   `json:"houseState,omitempty" xml:"houseState,omitempty"`
	IsVirtual     *int32   `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	ParentReferId *int64   `json:"parentReferId,omitempty" xml:"parentReferId,omitempty"`
	ReferId       *int64   `json:"referId,omitempty" xml:"referId,omitempty"`
	SpaceName     *string  `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	TagCode       *string  `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	Type          *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSpacesInfoResponseBodySpaceList) String() string {
	return tea.Prettify(s)
}

func (s GetSpacesInfoResponseBodySpaceList) GoString() string {
	return s.String()
}

func (s *GetSpacesInfoResponseBodySpaceList) SetBillingArea(v float32) *GetSpacesInfoResponseBodySpaceList {
	s.BillingArea = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetBuildingArea(v float32) *GetSpacesInfoResponseBodySpaceList {
	s.BuildingArea = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetFloor(v string) *GetSpacesInfoResponseBodySpaceList {
	s.Floor = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetHouseState(v int32) *GetSpacesInfoResponseBodySpaceList {
	s.HouseState = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetIsVirtual(v int32) *GetSpacesInfoResponseBodySpaceList {
	s.IsVirtual = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetParentReferId(v int64) *GetSpacesInfoResponseBodySpaceList {
	s.ParentReferId = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetReferId(v int64) *GetSpacesInfoResponseBodySpaceList {
	s.ReferId = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetSpaceName(v string) *GetSpacesInfoResponseBodySpaceList {
	s.SpaceName = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetTagCode(v string) *GetSpacesInfoResponseBodySpaceList {
	s.TagCode = &v
	return s
}

func (s *GetSpacesInfoResponseBodySpaceList) SetType(v string) *GetSpacesInfoResponseBodySpaceList {
	s.Type = &v
	return s
}

type GetSpacesInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpacesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpacesInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpacesInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSpacesInfoResponse) SetHeaders(v map[string]*string) *GetSpacesInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSpacesInfoResponse) SetStatusCode(v int32) *GetSpacesInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpacesInfoResponse) SetBody(v *GetSpacesInfoResponseBody) *GetSpacesInfoResponse {
	s.Body = v
	return s
}

type ListIndustryRoleUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListIndustryRoleUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListIndustryRoleUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListIndustryRoleUsersHeaders) SetCommonHeaders(v map[string]*string) *ListIndustryRoleUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListIndustryRoleUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListIndustryRoleUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListIndustryRoleUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SecurityManager
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s ListIndustryRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndustryRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListIndustryRoleUsersRequest) SetTagCode(v string) *ListIndustryRoleUsersRequest {
	s.TagCode = &v
	return s
}

type ListIndustryRoleUsersResponseBody struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ListIndustryRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndustryRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndustryRoleUsersResponseBody) SetUserIdList(v []*string) *ListIndustryRoleUsersResponseBody {
	s.UserIdList = v
	return s
}

type ListIndustryRoleUsersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndustryRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndustryRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndustryRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListIndustryRoleUsersResponse) SetHeaders(v map[string]*string) *ListIndustryRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListIndustryRoleUsersResponse) SetStatusCode(v int32) *ListIndustryRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndustryRoleUsersResponse) SetBody(v *ListIndustryRoleUsersResponseBody) *ListIndustryRoleUsersResponse {
	s.Body = v
	return s
}

type ListPointRulesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPointRulesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesHeaders) GoString() string {
	return s.String()
}

func (s *ListPointRulesHeaders) SetCommonHeaders(v map[string]*string) *ListPointRulesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPointRulesHeaders) SetXAcsDingtalkAccessToken(v string) *ListPointRulesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPointRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
}

func (s ListPointRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPointRulesRequest) SetIsCircle(v bool) *ListPointRulesRequest {
	s.IsCircle = &v
	return s
}

type ListPointRulesResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	PointRuleList []*ListPointRulesResponseBodyPointRuleList `json:"pointRuleList,omitempty" xml:"pointRuleList,omitempty" type:"Repeated"`
}

func (s ListPointRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponseBody) SetPointRuleList(v []*ListPointRulesResponseBodyPointRuleList) *ListPointRulesResponseBody {
	s.PointRuleList = v
	return s
}

type ListPointRulesResponseBodyPointRuleList struct {
	// This parameter is required.
	//
	// example:
	//
	// 50
	DayLimitTimes *int32 `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 排序Id
	OrderId *int32 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// rule_1
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 发动态
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPointRulesResponseBodyPointRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponseBodyPointRuleList) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponseBodyPointRuleList) SetDayLimitTimes(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.DayLimitTimes = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetExtension(v string) *ListPointRulesResponseBodyPointRuleList {
	s.Extension = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetGroupId(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.GroupId = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetOrderId(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.OrderId = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetRuleCode(v string) *ListPointRulesResponseBodyPointRuleList {
	s.RuleCode = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetRuleName(v string) *ListPointRulesResponseBodyPointRuleList {
	s.RuleName = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetScore(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.Score = &v
	return s
}

func (s *ListPointRulesResponseBodyPointRuleList) SetStatus(v int32) *ListPointRulesResponseBodyPointRuleList {
	s.Status = &v
	return s
}

type ListPointRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPointRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPointRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponse) SetHeaders(v map[string]*string) *ListPointRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPointRulesResponse) SetStatusCode(v int32) *ListPointRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPointRulesResponse) SetBody(v *ListPointRulesResponseBody) *ListPointRulesResponse {
	s.Body = v
	return s
}

type ListSubSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSubSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSubSpaceHeaders) GoString() string {
	return s.String()
}

func (s *ListSubSpaceHeaders) SetCommonHeaders(v map[string]*string) *ListSubSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *ListSubSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSubSpaceRequest struct {
	ReferId        *int64  `json:"referId,omitempty" xml:"referId,omitempty"`
	ResidentCorpId *string `json:"residentCorpId,omitempty" xml:"residentCorpId,omitempty"`
}

func (s ListSubSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubSpaceRequest) GoString() string {
	return s.String()
}

func (s *ListSubSpaceRequest) SetReferId(v int64) *ListSubSpaceRequest {
	s.ReferId = &v
	return s
}

func (s *ListSubSpaceRequest) SetResidentCorpId(v string) *ListSubSpaceRequest {
	s.ResidentCorpId = &v
	return s
}

type ListSubSpaceResponseBody struct {
	SpaceList []*ListSubSpaceResponseBodySpaceList `json:"spaceList,omitempty" xml:"spaceList,omitempty" type:"Repeated"`
}

func (s ListSubSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubSpaceResponseBody) SetSpaceList(v []*ListSubSpaceResponseBodySpaceList) *ListSubSpaceResponseBody {
	s.SpaceList = v
	return s
}

type ListSubSpaceResponseBodySpaceList struct {
	BillingArea   *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	BuildingArea  *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	Floor         *string  `json:"floor,omitempty" xml:"floor,omitempty"`
	HouseState    *int32   `json:"houseState,omitempty" xml:"houseState,omitempty"`
	IsVirtual     *int32   `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	ParentReferId *int64   `json:"parentReferId,omitempty" xml:"parentReferId,omitempty"`
	ReferId       *int64   `json:"referId,omitempty" xml:"referId,omitempty"`
	SpaceName     *string  `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	TagCode       *string  `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	Type          *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSubSpaceResponseBodySpaceList) String() string {
	return tea.Prettify(s)
}

func (s ListSubSpaceResponseBodySpaceList) GoString() string {
	return s.String()
}

func (s *ListSubSpaceResponseBodySpaceList) SetBillingArea(v float32) *ListSubSpaceResponseBodySpaceList {
	s.BillingArea = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetBuildingArea(v float32) *ListSubSpaceResponseBodySpaceList {
	s.BuildingArea = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetFloor(v string) *ListSubSpaceResponseBodySpaceList {
	s.Floor = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetHouseState(v int32) *ListSubSpaceResponseBodySpaceList {
	s.HouseState = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetIsVirtual(v int32) *ListSubSpaceResponseBodySpaceList {
	s.IsVirtual = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetParentReferId(v int64) *ListSubSpaceResponseBodySpaceList {
	s.ParentReferId = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetReferId(v int64) *ListSubSpaceResponseBodySpaceList {
	s.ReferId = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetSpaceName(v string) *ListSubSpaceResponseBodySpaceList {
	s.SpaceName = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetTagCode(v string) *ListSubSpaceResponseBodySpaceList {
	s.TagCode = &v
	return s
}

func (s *ListSubSpaceResponseBodySpaceList) SetType(v string) *ListSubSpaceResponseBodySpaceList {
	s.Type = &v
	return s
}

type ListSubSpaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubSpaceResponse) GoString() string {
	return s.String()
}

func (s *ListSubSpaceResponse) SetHeaders(v map[string]*string) *ListSubSpaceResponse {
	s.Headers = v
	return s
}

func (s *ListSubSpaceResponse) SetStatusCode(v int32) *ListSubSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubSpaceResponse) SetBody(v *ListSubSpaceResponseBody) *ListSubSpaceResponse {
	s.Body = v
	return s
}

type ListUncheckUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUncheckUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUncheckUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListUncheckUsersHeaders) SetCommonHeaders(v map[string]*string) *ListUncheckUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUncheckUsersHeaders) SetXAcsDingtalkAccessToken(v string) *ListUncheckUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUncheckUsersRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1652698991669
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListUncheckUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUncheckUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUncheckUsersRequest) SetMaxResults(v int32) *ListUncheckUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUncheckUsersRequest) SetNextToken(v int32) *ListUncheckUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUncheckUsersRequest) SetStartTime(v int64) *ListUncheckUsersRequest {
	s.StartTime = &v
	return s
}

func (s *ListUncheckUsersRequest) SetStatus(v int32) *ListUncheckUsersRequest {
	s.Status = &v
	return s
}

type ListUncheckUsersResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *int64                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Values    []*ListUncheckUsersResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListUncheckUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUncheckUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUncheckUsersResponseBody) SetHasMore(v bool) *ListUncheckUsersResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListUncheckUsersResponseBody) SetNextToken(v int64) *ListUncheckUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUncheckUsersResponseBody) SetValues(v []*ListUncheckUsersResponseBodyValues) *ListUncheckUsersResponseBody {
	s.Values = v
	return s
}

type ListUncheckUsersResponseBodyValues struct {
	// example:
	//
	// 5345345
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// "{\"startTime\":\"1654746593623\",\"endTime\":\"1656042593623\"}"
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// 1652683318162
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1652683318162
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// true
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// example:
	//
	// 张工
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 312423423
	UnionId *int64 `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s ListUncheckUsersResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s ListUncheckUsersResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListUncheckUsersResponseBodyValues) SetDeptId(v int64) *ListUncheckUsersResponseBodyValues {
	s.DeptId = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetExtension(v string) *ListUncheckUsersResponseBodyValues {
	s.Extension = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetGmtCreate(v int64) *ListUncheckUsersResponseBodyValues {
	s.GmtCreate = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetGmtModified(v int64) *ListUncheckUsersResponseBodyValues {
	s.GmtModified = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetIsPropertyOwner(v bool) *ListUncheckUsersResponseBodyValues {
	s.IsPropertyOwner = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetName(v string) *ListUncheckUsersResponseBodyValues {
	s.Name = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetStatus(v int32) *ListUncheckUsersResponseBodyValues {
	s.Status = &v
	return s
}

func (s *ListUncheckUsersResponseBodyValues) SetUnionId(v int64) *ListUncheckUsersResponseBodyValues {
	s.UnionId = &v
	return s
}

type ListUncheckUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUncheckUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUncheckUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUncheckUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUncheckUsersResponse) SetHeaders(v map[string]*string) *ListUncheckUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUncheckUsersResponse) SetStatusCode(v int32) *ListUncheckUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUncheckUsersResponse) SetBody(v *ListUncheckUsersResponseBody) *ListUncheckUsersResponse {
	s.Body = v
	return s
}

type ListUserIndustryRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUserIndustryRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserIndustryRolesHeaders) GoString() string {
	return s.String()
}

func (s *ListUserIndustryRolesHeaders) SetCommonHeaders(v map[string]*string) *ListUserIndustryRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserIndustryRolesHeaders) SetXAcsDingtalkAccessToken(v string) *ListUserIndustryRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUserIndustryRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUserIndustryRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserIndustryRolesRequest) GoString() string {
	return s.String()
}

func (s *ListUserIndustryRolesRequest) SetUserId(v string) *ListUserIndustryRolesRequest {
	s.UserId = &v
	return s
}

type ListUserIndustryRolesResponseBody struct {
	RoleList []*ListUserIndustryRolesResponseBodyRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
}

func (s ListUserIndustryRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserIndustryRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserIndustryRolesResponseBody) SetRoleList(v []*ListUserIndustryRolesResponseBodyRoleList) *ListUserIndustryRolesResponseBody {
	s.RoleList = v
	return s
}

type ListUserIndustryRolesResponseBodyRoleList struct {
	// example:
	//
	// 312423423
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// 安保部经理
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// SecurityManager
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s ListUserIndustryRolesResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s ListUserIndustryRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListUserIndustryRolesResponseBodyRoleList) SetRoleId(v int64) *ListUserIndustryRolesResponseBodyRoleList {
	s.RoleId = &v
	return s
}

func (s *ListUserIndustryRolesResponseBodyRoleList) SetRoleName(v string) *ListUserIndustryRolesResponseBodyRoleList {
	s.RoleName = &v
	return s
}

func (s *ListUserIndustryRolesResponseBodyRoleList) SetTagCode(v string) *ListUserIndustryRolesResponseBodyRoleList {
	s.TagCode = &v
	return s
}

type ListUserIndustryRolesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserIndustryRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserIndustryRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserIndustryRolesResponse) GoString() string {
	return s.String()
}

func (s *ListUserIndustryRolesResponse) SetHeaders(v map[string]*string) *ListUserIndustryRolesResponse {
	s.Headers = v
	return s
}

func (s *ListUserIndustryRolesResponse) SetStatusCode(v int32) *ListUserIndustryRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserIndustryRolesResponse) SetBody(v *ListUserIndustryRolesResponseBody) *ListUserIndustryRolesResponse {
	s.Body = v
	return s
}

type PagePointHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PagePointHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryHeaders) GoString() string {
	return s.String()
}

func (s *PagePointHistoryHeaders) SetCommonHeaders(v map[string]*string) *PagePointHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PagePointHistoryHeaders) SetXAcsDingtalkAccessToken(v string) *PagePointHistoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PagePointHistoryRequest struct {
	// example:
	//
	// 1631260866105
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1630345050858
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PagePointHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryRequest) GoString() string {
	return s.String()
}

func (s *PagePointHistoryRequest) SetEndTime(v int64) *PagePointHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *PagePointHistoryRequest) SetIsCircle(v bool) *PagePointHistoryRequest {
	s.IsCircle = &v
	return s
}

func (s *PagePointHistoryRequest) SetMaxResults(v int32) *PagePointHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *PagePointHistoryRequest) SetNextToken(v int64) *PagePointHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *PagePointHistoryRequest) SetStartTime(v int64) *PagePointHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *PagePointHistoryRequest) SetUserId(v string) *PagePointHistoryRequest {
	s.UserId = &v
	return s
}

type PagePointHistoryResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3276
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	PointRecordList []*PagePointHistoryResponseBodyPointRecordList `json:"pointRecordList,omitempty" xml:"pointRecordList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PagePointHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponseBody) SetHasMore(v bool) *PagePointHistoryResponseBody {
	s.HasMore = &v
	return s
}

func (s *PagePointHistoryResponseBody) SetNextToken(v int64) *PagePointHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *PagePointHistoryResponseBody) SetPointRecordList(v []*PagePointHistoryResponseBodyPointRecordList) *PagePointHistoryResponseBody {
	s.PointRecordList = v
	return s
}

func (s *PagePointHistoryResponseBody) SetTotalCount(v int64) *PagePointHistoryResponseBody {
	s.TotalCount = &v
	return s
}

type PagePointHistoryResponseBodyPointRecordList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1634630147
	CreateAt *int64 `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// example:
	//
	// rule_1
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 发动态
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7653
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PagePointHistoryResponseBodyPointRecordList) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponseBodyPointRecordList) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetCreateAt(v int64) *PagePointHistoryResponseBodyPointRecordList {
	s.CreateAt = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetRuleCode(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.RuleCode = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetRuleName(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.RuleName = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetScore(v int32) *PagePointHistoryResponseBodyPointRecordList {
	s.Score = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetUserId(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.UserId = &v
	return s
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetUuid(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.Uuid = &v
	return s
}

type PagePointHistoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PagePointHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PagePointHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponse) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponse) SetHeaders(v map[string]*string) *PagePointHistoryResponse {
	s.Headers = v
	return s
}

func (s *PagePointHistoryResponse) SetStatusCode(v int32) *PagePointHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *PagePointHistoryResponse) SetBody(v *PagePointHistoryResponseBody) *PagePointHistoryResponse {
	s.Body = v
	return s
}

type RemoveResidentMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveResidentMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentMemberHeaders) GoString() string {
	return s.String()
}

func (s *RemoveResidentMemberHeaders) SetCommonHeaders(v map[string]*string) *RemoveResidentMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveResidentMemberHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveResidentMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveResidentMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 111112***
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RemoveResidentMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveResidentMemberRequest) SetDeptId(v int64) *RemoveResidentMemberRequest {
	s.DeptId = &v
	return s
}

func (s *RemoveResidentMemberRequest) SetUnionId(v string) *RemoveResidentMemberRequest {
	s.UnionId = &v
	return s
}

func (s *RemoveResidentMemberRequest) SetUserId(v string) *RemoveResidentMemberRequest {
	s.UserId = &v
	return s
}

type RemoveResidentMemberResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemoveResidentMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveResidentMemberResponseBody) SetSuccess(v bool) *RemoveResidentMemberResponseBody {
	s.Success = &v
	return s
}

type RemoveResidentMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveResidentMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveResidentMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveResidentMemberResponse) SetHeaders(v map[string]*string) *RemoveResidentMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveResidentMemberResponse) SetStatusCode(v int32) *RemoveResidentMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveResidentMemberResponse) SetBody(v *RemoveResidentMemberResponseBody) *RemoveResidentMemberResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RemoveResidentUserResponse) SetStatusCode(v int32) *RemoveResidentUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveResidentUserResponse) SetBody(v *RemoveResidentUserResponseBody) *RemoveResidentUserResponse {
	s.Body = v
	return s
}

type SearchResidentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchResidentHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchResidentHeaders) GoString() string {
	return s.String()
}

func (s *SearchResidentHeaders) SetCommonHeaders(v map[string]*string) *SearchResidentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchResidentHeaders) SetXAcsDingtalkAccessToken(v string) *SearchResidentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchResidentRequest struct {
	// This parameter is required.
	ResidentCropId *string `json:"residentCropId,omitempty" xml:"residentCropId,omitempty"`
	// This parameter is required.
	SearchWord *string `json:"searchWord,omitempty" xml:"searchWord,omitempty"`
}

func (s SearchResidentRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchResidentRequest) GoString() string {
	return s.String()
}

func (s *SearchResidentRequest) SetResidentCropId(v string) *SearchResidentRequest {
	s.ResidentCropId = &v
	return s
}

func (s *SearchResidentRequest) SetSearchWord(v string) *SearchResidentRequest {
	s.SearchWord = &v
	return s
}

type SearchResidentResponseBody struct {
	ResidenceList []*SearchResidentResponseBodyResidenceList `json:"residenceList,omitempty" xml:"residenceList,omitempty" type:"Repeated"`
}

func (s SearchResidentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResidentResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResidentResponseBody) SetResidenceList(v []*SearchResidentResponseBodyResidenceList) *SearchResidentResponseBody {
	s.ResidenceList = v
	return s
}

type SearchResidentResponseBodyResidenceList struct {
	Active          *bool   `json:"active,omitempty" xml:"active,omitempty"`
	ExtField        *string `json:"extField,omitempty" xml:"extField,omitempty"`
	IsPropertyOwner *bool   `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	RelateType      *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
}

func (s SearchResidentResponseBodyResidenceList) String() string {
	return tea.Prettify(s)
}

func (s SearchResidentResponseBodyResidenceList) GoString() string {
	return s.String()
}

func (s *SearchResidentResponseBodyResidenceList) SetActive(v bool) *SearchResidentResponseBodyResidenceList {
	s.Active = &v
	return s
}

func (s *SearchResidentResponseBodyResidenceList) SetExtField(v string) *SearchResidentResponseBodyResidenceList {
	s.ExtField = &v
	return s
}

func (s *SearchResidentResponseBodyResidenceList) SetIsPropertyOwner(v bool) *SearchResidentResponseBodyResidenceList {
	s.IsPropertyOwner = &v
	return s
}

func (s *SearchResidentResponseBodyResidenceList) SetName(v string) *SearchResidentResponseBodyResidenceList {
	s.Name = &v
	return s
}

func (s *SearchResidentResponseBodyResidenceList) SetRelateType(v string) *SearchResidentResponseBodyResidenceList {
	s.RelateType = &v
	return s
}

type SearchResidentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResidentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResidentResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResidentResponse) GoString() string {
	return s.String()
}

func (s *SearchResidentResponse) SetHeaders(v map[string]*string) *SearchResidentResponse {
	s.Headers = v
	return s
}

func (s *SearchResidentResponse) SetStatusCode(v int32) *SearchResidentResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResidentResponse) SetBody(v *SearchResidentResponseBody) *SearchResidentResponse {
	s.Body = v
	return s
}

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
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 第一网格组
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// example:
	//
	// 1234
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
}

func (s UpdateResideceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResideceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResideceGroupRequest) SetDepartmentId(v int64) *UpdateResideceGroupRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetDepartmentName(v string) *UpdateResideceGroupRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResideceGroupRequest) SetManagerUserId(v string) *UpdateResideceGroupRequest {
	s.ManagerUserId = &v
	return s
}

type UpdateResideceGroupResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResideceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateResideceGroupResponse) SetStatusCode(v int32) *UpdateResideceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResideceGroupResponse) SetBody(v *UpdateResideceGroupResponseBody) *UpdateResideceGroupResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 101户
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// example:
	//
	// false
	Destitute *bool `json:"destitute,omitempty" xml:"destitute,omitempty"`
	// example:
	//
	// 第1网格
	Grid *string `json:"grid,omitempty" xml:"grid,omitempty"`
	// example:
	//
	// 16612345678
	HomeTel *string `json:"homeTel,omitempty" xml:"homeTel,omitempty"`
	// example:
	//
	// 1234
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ParentDepartmentId *int64 `json:"parentDepartmentId,omitempty" xml:"parentDepartmentId,omitempty"`
}

func (s UpdateResidenceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidenceRequest) SetDepartmentId(v int64) *UpdateResidenceRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidenceRequest) SetDepartmentName(v string) *UpdateResidenceRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateResidenceRequest) SetDestitute(v bool) *UpdateResidenceRequest {
	s.Destitute = &v
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

func (s *UpdateResidenceRequest) SetManagerUserId(v string) *UpdateResidenceRequest {
	s.ManagerUserId = &v
	return s
}

func (s *UpdateResidenceRequest) SetParentDepartmentId(v int64) *UpdateResidenceRequest {
	s.ParentDepartmentId = &v
	return s
}

type UpdateResidenceResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResidenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateResidenceResponse) SetStatusCode(v int32) *UpdateResidenceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResidenceResponse) SetBody(v *UpdateResidenceResponseBody) *UpdateResidenceResponse {
	s.Body = v
	return s
}

type UpdateResidentBlackBoardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidentBlackBoardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentBlackBoardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidentBlackBoardHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidentBlackBoardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidentBlackBoardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidentBlackBoardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidentBlackBoardRequest struct {
	BlackboardId *string `json:"blackboardId,omitempty" xml:"blackboardId,omitempty"`
	// This parameter is required.
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateResidentBlackBoardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentBlackBoardRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidentBlackBoardRequest) SetBlackboardId(v string) *UpdateResidentBlackBoardRequest {
	s.BlackboardId = &v
	return s
}

func (s *UpdateResidentBlackBoardRequest) SetContext(v string) *UpdateResidentBlackBoardRequest {
	s.Context = &v
	return s
}

func (s *UpdateResidentBlackBoardRequest) SetMediaId(v string) *UpdateResidentBlackBoardRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateResidentBlackBoardRequest) SetTitle(v string) *UpdateResidentBlackBoardRequest {
	s.Title = &v
	return s
}

type UpdateResidentBlackBoardResponseBody struct {
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateResidentBlackBoardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentBlackBoardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidentBlackBoardResponseBody) SetSuccess(v bool) *UpdateResidentBlackBoardResponseBody {
	s.Success = &v
	return s
}

type UpdateResidentBlackBoardResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResidentBlackBoardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentBlackBoardResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidentBlackBoardResponse) SetHeaders(v map[string]*string) *UpdateResidentBlackBoardResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidentBlackBoardResponse) SetStatusCode(v int32) *UpdateResidentBlackBoardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResidentBlackBoardResponse) SetBody(v *UpdateResidentBlackBoardResponseBody) *UpdateResidentBlackBoardResponse {
	s.Body = v
	return s
}

type UpdateResidentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidentInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidentInfoRequest struct {
	Address       *string  `json:"address,omitempty" xml:"address,omitempty"`
	BuildingArea  *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	CityName      *string  `json:"cityName,omitempty" xml:"cityName,omitempty"`
	CommunityType *int64   `json:"communityType,omitempty" xml:"communityType,omitempty"`
	CountyName    *string  `json:"countyName,omitempty" xml:"countyName,omitempty"`
	Location      *string  `json:"location,omitempty" xml:"location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试小区1
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	ProvName *string `json:"provName,omitempty" xml:"provName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	State     *int64  `json:"state,omitempty" xml:"state,omitempty"`
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s UpdateResidentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidentInfoRequest) SetAddress(v string) *UpdateResidentInfoRequest {
	s.Address = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetBuildingArea(v float32) *UpdateResidentInfoRequest {
	s.BuildingArea = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetCityName(v string) *UpdateResidentInfoRequest {
	s.CityName = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetCommunityType(v int64) *UpdateResidentInfoRequest {
	s.CommunityType = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetCountyName(v string) *UpdateResidentInfoRequest {
	s.CountyName = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetLocation(v string) *UpdateResidentInfoRequest {
	s.Location = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetName(v string) *UpdateResidentInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetProvName(v string) *UpdateResidentInfoRequest {
	s.ProvName = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetState(v int64) *UpdateResidentInfoRequest {
	s.State = &v
	return s
}

func (s *UpdateResidentInfoRequest) SetTelephone(v string) *UpdateResidentInfoRequest {
	s.Telephone = &v
	return s
}

type UpdateResidentInfoResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateResidentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidentInfoResponseBody) SetSuccess(v bool) *UpdateResidentInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateResidentInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResidentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResidentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidentInfoResponse) SetHeaders(v map[string]*string) *UpdateResidentInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidentInfoResponse) SetStatusCode(v int32) *UpdateResidentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResidentInfoResponse) SetBody(v *UpdateResidentInfoResponseBody) *UpdateResidentInfoResponse {
	s.Body = v
	return s
}

type UpdateResidentMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateResidentMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentMemberHeaders) GoString() string {
	return s.String()
}

func (s *UpdateResidentMemberHeaders) SetCommonHeaders(v map[string]*string) *UpdateResidentMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateResidentMemberHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateResidentMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateResidentMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试小区1
	ResidentUpdateInfo *UpdateResidentMemberRequestResidentUpdateInfo `json:"residentUpdateInfo,omitempty" xml:"residentUpdateInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1212
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateResidentMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateResidentMemberRequest) SetResidentUpdateInfo(v *UpdateResidentMemberRequestResidentUpdateInfo) *UpdateResidentMemberRequest {
	s.ResidentUpdateInfo = v
	return s
}

func (s *UpdateResidentMemberRequest) SetUnionId(v string) *UpdateResidentMemberRequest {
	s.UnionId = &v
	return s
}

type UpdateResidentMemberRequestResidentUpdateInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 11112
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// true
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// example:
	//
	// {"startTime":1652358627106,"endTime":1652445027106}
	MemberDeptExtension map[string]*string `json:"memberDeptExtension,omitempty" xml:"memberDeptExtension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11123
	OldDeptId *int64 `json:"oldDeptId,omitempty" xml:"oldDeptId,omitempty"`
	// example:
	//
	// 1
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// example:
	//
	// 11123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateResidentMemberRequestResidentUpdateInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentMemberRequestResidentUpdateInfo) GoString() string {
	return s.String()
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetDeptId(v int64) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.DeptId = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetIsPropertyOwner(v bool) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.IsPropertyOwner = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetMemberDeptExtension(v map[string]*string) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.MemberDeptExtension = v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetName(v string) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.Name = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetOldDeptId(v int64) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.OldDeptId = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetRelateType(v string) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.RelateType = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetUserId(v string) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.UserId = &v
	return s
}

type UpdateResidentMemberResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateResidentMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResidentMemberResponseBody) SetSuccess(v bool) *UpdateResidentMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateResidentMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResidentMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateResidentMemberResponse) SetHeaders(v map[string]*string) *UpdateResidentMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateResidentMemberResponse) SetStatusCode(v int32) *UpdateResidentMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResidentMemberResponse) SetBody(v *UpdateResidentMemberResponseBody) *UpdateResidentMemberResponse {
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
	// example:
	//
	// 美好社区创景街道万通公寓
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DepartmentId *int64                               `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	ExtField     []*UpdateResidentUserRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsRetainOldDept *bool `json:"isRetainOldDept,omitempty" xml:"isRetainOldDept,omitempty"`
	// example:
	//
	// 15612345678
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	OldDepartmentId *int64 `json:"oldDepartmentId,omitempty" xml:"oldDepartmentId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// SELF
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 王建国
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
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

func (s *UpdateResidentUserRequest) SetDepartmentId(v int64) *UpdateResidentUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateResidentUserRequest) SetExtField(v []*UpdateResidentUserRequestExtField) *UpdateResidentUserRequest {
	s.ExtField = v
	return s
}

func (s *UpdateResidentUserRequest) SetIsRetainOldDept(v bool) *UpdateResidentUserRequest {
	s.IsRetainOldDept = &v
	return s
}

func (s *UpdateResidentUserRequest) SetMobile(v string) *UpdateResidentUserRequest {
	s.Mobile = &v
	return s
}

func (s *UpdateResidentUserRequest) SetOldDepartmentId(v int64) *UpdateResidentUserRequest {
	s.OldDepartmentId = &v
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

func (s *UpdateResidentUserRequest) SetUserName(v string) *UpdateResidentUserRequest {
	s.UserName = &v
	return s
}

type UpdateResidentUserRequestExtField struct {
	// example:
	//
	// 性别
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 女
	ItemValue *string `json:"itemValue,omitempty" xml:"itemValue,omitempty"`
}

func (s UpdateResidentUserRequestExtField) String() string {
	return tea.Prettify(s)
}

func (s UpdateResidentUserRequestExtField) GoString() string {
	return s.String()
}

func (s *UpdateResidentUserRequestExtField) SetItemName(v string) *UpdateResidentUserRequestExtField {
	s.ItemName = &v
	return s
}

func (s *UpdateResidentUserRequestExtField) SetItemValue(v string) *UpdateResidentUserRequestExtField {
	s.ItemValue = &v
	return s
}

type UpdateResidentUserResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResidentUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateResidentUserResponse) SetStatusCode(v int32) *UpdateResidentUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResidentUserResponse) SetBody(v *UpdateResidentUserResponseBody) *UpdateResidentUserResponse {
	s.Body = v
	return s
}

type UpdateSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSpaceHeaders) SetCommonHeaders(v map[string]*string) *UpdateSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A栋
	SpaceInfoVOList []*UpdateSpaceRequestSpaceInfoVOList `json:"spaceInfoVOList,omitempty" xml:"spaceInfoVOList,omitempty" type:"Repeated"`
}

func (s UpdateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceRequest) SetSpaceInfoVOList(v []*UpdateSpaceRequestSpaceInfoVOList) *UpdateSpaceRequest {
	s.SpaceInfoVOList = v
	return s
}

type UpdateSpaceRequestSpaceInfoVOList struct {
	// example:
	//
	// 123.4
	BillingArea *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	// example:
	//
	// 123.4
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	// example:
	//
	// 当tagcode为Building的时候必填
	BuildingType *int64 `json:"buildingType,omitempty" xml:"buildingType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10005
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 12
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// example:
	//
	// 1
	HouseState *int64 `json:"houseState,omitempty" xml:"houseState,omitempty"`
	// example:
	//
	// 1
	HouseType *int64 `json:"houseType,omitempty" xml:"houseType,omitempty"`
	// example:
	//
	// 二单元
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentDeptId *int64  `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 空间类型标签code，House/Unit/Building/Partition
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s UpdateSpaceRequestSpaceInfoVOList) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceRequestSpaceInfoVOList) GoString() string {
	return s.String()
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetBillingArea(v float32) *UpdateSpaceRequestSpaceInfoVOList {
	s.BillingArea = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetBuildingArea(v float32) *UpdateSpaceRequestSpaceInfoVOList {
	s.BuildingArea = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetBuildingType(v int64) *UpdateSpaceRequestSpaceInfoVOList {
	s.BuildingType = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetDeptId(v int64) *UpdateSpaceRequestSpaceInfoVOList {
	s.DeptId = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetFloor(v string) *UpdateSpaceRequestSpaceInfoVOList {
	s.Floor = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetHouseState(v int64) *UpdateSpaceRequestSpaceInfoVOList {
	s.HouseState = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetHouseType(v int64) *UpdateSpaceRequestSpaceInfoVOList {
	s.HouseType = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetName(v string) *UpdateSpaceRequestSpaceInfoVOList {
	s.Name = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetParentDeptId(v int64) *UpdateSpaceRequestSpaceInfoVOList {
	s.ParentDeptId = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetTagCode(v string) *UpdateSpaceRequestSpaceInfoVOList {
	s.TagCode = &v
	return s
}

type UpdateSpaceResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponseBody) SetSuccess(v bool) *UpdateSpaceResponseBody {
	s.Success = &v
	return s
}

type UpdateSpaceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponse) SetHeaders(v map[string]*string) *UpdateSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSpaceResponse) SetStatusCode(v int32) *UpdateSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSpaceResponse) SetBody(v *UpdateSpaceResponseBody) *UpdateSpaceResponse {
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
// 增加积分
//
// @param request - AddPointRequest
//
// @param headers - AddPointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPointResponse
func (client *Client) AddPointWithOptions(request *AddPointRequest, headers *AddPointHeaders, runtime *util.RuntimeOptions) (_result *AddPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionTime)) {
		query["actionTime"] = request.ActionTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCode)) {
		query["ruleCode"] = request.RuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["ruleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Score)) {
		query["score"] = request.Score
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
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
		Action:      tea.String("AddPoint"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/points"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加积分
//
// @param request - AddPointRequest
//
// @return AddPointResponse
func (client *Client) AddPoint(request *AddPointRequest) (_result *AddPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddPointHeaders{}
	_result = &AddPointResponse{}
	_body, _err := client.AddPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增加组户
//
// @param request - AddResidentDepartmentRequest
//
// @param headers - AddResidentDepartmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResidentDepartmentResponse
func (client *Client) AddResidentDepartmentWithOptions(request *AddResidentDepartmentRequest, headers *AddResidentDepartmentHeaders, runtime *util.RuntimeOptions) (_result *AddResidentDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.IsResidenceGroup)) {
		query["isResidenceGroup"] = request.IsResidenceGroup
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
		Action:      tea.String("AddResidentDepartment"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/departments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加组户
//
// @param request - AddResidentDepartmentRequest
//
// @return AddResidentDepartmentResponse
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

// Summary:
//
// 添加小区成员
//
// @param request - AddResidentMemberRequest
//
// @param headers - AddResidentMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResidentMemberResponse
func (client *Client) AddResidentMemberWithOptions(request *AddResidentMemberRequest, headers *AddResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *AddResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResidentAddInfo)) {
		body["residentAddInfo"] = request.ResidentAddInfo
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
		Action:      tea.String("AddResidentMember"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddResidentMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加小区成员
//
// @param request - AddResidentMemberRequest
//
// @return AddResidentMemberResponse
func (client *Client) AddResidentMember(request *AddResidentMemberRequest) (_result *AddResidentMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddResidentMemberHeaders{}
	_result = &AddResidentMemberResponse{}
	_body, _err := client.AddResidentMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增居民
//
// @param request - AddResidentUsersRequest
//
// @param headers - AddResidentUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResidentUsersResponse
func (client *Client) AddResidentUsersWithOptions(request *AddResidentUsersRequest, headers *AddResidentUsersHeaders, runtime *util.RuntimeOptions) (_result *AddResidentUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.IsLeaseholder)) {
		query["isLeaseholder"] = request.IsLeaseholder
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
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
		Action:      tea.String("AddResidentUsers"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddResidentUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增居民
//
// @param request - AddResidentUsersRequest
//
// @return AddResidentUsersResponse
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

// Summary:
//
// 创建小区公告
//
// @param request - CreateResidentBlackBoardRequest
//
// @param headers - CreateResidentBlackBoardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResidentBlackBoardResponse
func (client *Client) CreateResidentBlackBoardWithOptions(request *CreateResidentBlackBoardRequest, headers *CreateResidentBlackBoardHeaders, runtime *util.RuntimeOptions) (_result *CreateResidentBlackBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Context)) {
		body["context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.SendTime)) {
		body["sendTime"] = request.SendTime
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResidentBlackBoard"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/blackboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResidentBlackBoardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建小区公告
//
// @param request - CreateResidentBlackBoardRequest
//
// @return CreateResidentBlackBoardResponse
func (client *Client) CreateResidentBlackBoard(request *CreateResidentBlackBoardRequest) (_result *CreateResidentBlackBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateResidentBlackBoardHeaders{}
	_result = &CreateResidentBlackBoardResponse{}
	_body, _err := client.CreateResidentBlackBoardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建小区空间，含分区，楼栋，单元，房屋等
//
// @param request - CreateSpaceRequest
//
// @param headers - CreateSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSpaceResponse
func (client *Client) CreateSpaceWithOptions(request *CreateSpaceRequest, headers *CreateSpaceHeaders, runtime *util.RuntimeOptions) (_result *CreateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingArea)) {
		body["billingArea"] = request.BillingArea
	}

	if !tea.BoolValue(util.IsUnset(request.BuildingArea)) {
		body["buildingArea"] = request.BuildingArea
	}

	if !tea.BoolValue(util.IsUnset(request.Floor)) {
		body["floor"] = request.Floor
	}

	if !tea.BoolValue(util.IsUnset(request.HouseState)) {
		body["houseState"] = request.HouseState
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDeptId)) {
		body["parentDeptId"] = request.ParentDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
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
		Action:      tea.String("CreateSpace"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建小区空间，含分区，楼栋，单元，房屋等
//
// @param request - CreateSpaceRequest
//
// @return CreateSpaceResponse
func (client *Client) CreateSpace(request *CreateSpaceRequest) (_result *CreateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSpaceHeaders{}
	_result = &CreateSpaceResponse{}
	_body, _err := client.CreateSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除小区公告
//
// @param request - DeleteResidentBlackBoardRequest
//
// @param headers - DeleteResidentBlackBoardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResidentBlackBoardResponse
func (client *Client) DeleteResidentBlackBoardWithOptions(request *DeleteResidentBlackBoardRequest, headers *DeleteResidentBlackBoardHeaders, runtime *util.RuntimeOptions) (_result *DeleteResidentBlackBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackboardId)) {
		query["blackboardId"] = request.BlackboardId
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
		Action:      tea.String("DeleteResidentBlackBoard"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/blackboards"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResidentBlackBoardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除小区公告
//
// @param request - DeleteResidentBlackBoardRequest
//
// @return DeleteResidentBlackBoardResponse
func (client *Client) DeleteResidentBlackBoard(request *DeleteResidentBlackBoardRequest) (_result *DeleteResidentBlackBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteResidentBlackBoardHeaders{}
	_result = &DeleteResidentBlackBoardResponse{}
	_body, _err := client.DeleteResidentBlackBoardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除组户信息
//
// @param request - DeleteResidentDepartmentRequest
//
// @param headers - DeleteResidentDepartmentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResidentDepartmentResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResidentDepartment"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/departments"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResidentDepartmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组户信息
//
// @param request - DeleteResidentDepartmentRequest
//
// @return DeleteResidentDepartmentResponse
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

// Summary:
//
// 删除小区空间，含分区，楼栋，单元，房屋
//
// @param request - DeleteSpaceRequest
//
// @param headers - DeleteSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSpaceResponse
func (client *Client) DeleteSpaceWithOptions(request *DeleteSpaceRequest, headers *DeleteSpaceHeaders, runtime *util.RuntimeOptions) (_result *DeleteSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
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
		Action:      tea.String("DeleteSpace"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除小区空间，含分区，楼栋，单元，房屋
//
// @param request - DeleteSpaceRequest
//
// @return DeleteSpaceResponse
func (client *Client) DeleteSpace(request *DeleteSpaceRequest) (_result *DeleteSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSpaceHeaders{}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DeleteSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定群的openConversationId
//
// @param request - GetConversationIdRequest
//
// @param headers - GetConversationIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationIdResponse
func (client *Client) GetConversationIdWithOptions(request *GetConversationIdRequest, headers *GetConversationIdHeaders, runtime *util.RuntimeOptions) (_result *GetConversationIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		query["chatId"] = request.ChatId
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
		Action:      tea.String("GetConversationId"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/conversations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定群的openConversationId
//
// @param request - GetConversationIdRequest
//
// @return GetConversationIdResponse
func (client *Client) GetConversationId(request *GetConversationIdRequest) (_result *GetConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationIdHeaders{}
	_result = &GetConversationIdResponse{}
	_body, _err := client.GetConversationIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取组织的行业类型
//
// @param headers - GetIndustryTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndustryTypeResponse
func (client *Client) GetIndustryTypeWithOptions(headers *GetIndustryTypeHeaders, runtime *util.RuntimeOptions) (_result *GetIndustryTypeResponse, _err error) {
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
		Action:      tea.String("GetIndustryType"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/organizations/industryTypes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndustryTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织的行业类型
//
// @return GetIndustryTypeResponse
func (client *Client) GetIndustryType() (_result *GetIndustryTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetIndustryTypeHeaders{}
	_result = &GetIndustryTypeResponse{}
	_body, _err := client.GetIndustryTypeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取物业公司信息
//
// @param request - GetPropertyInfoRequest
//
// @param headers - GetPropertyInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPropertyInfoResponse
func (client *Client) GetPropertyInfoWithOptions(request *GetPropertyInfoRequest, headers *GetPropertyInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPropertyInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyCorpId)) {
		query["propertyCorpId"] = request.PropertyCorpId
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
		Action:      tea.String("GetPropertyInfo"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/propertyInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPropertyInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取物业公司信息
//
// @param request - GetPropertyInfoRequest
//
// @return GetPropertyInfoResponse
func (client *Client) GetPropertyInfo(request *GetPropertyInfoRequest) (_result *GetPropertyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPropertyInfoHeaders{}
	_result = &GetPropertyInfoResponse{}
	_body, _err := client.GetPropertyInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取小区信息
//
// @param request - GetResidentInfoRequest
//
// @param headers - GetResidentInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResidentInfoResponse
func (client *Client) GetResidentInfoWithOptions(request *GetResidentInfoRequest, headers *GetResidentInfoHeaders, runtime *util.RuntimeOptions) (_result *GetResidentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResidentCorpId)) {
		query["residentCorpId"] = request.ResidentCorpId
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
		Action:      tea.String("GetResidentInfo"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/residentInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResidentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取小区信息
//
// @param request - GetResidentInfoRequest
//
// @return GetResidentInfoResponse
func (client *Client) GetResidentInfo(request *GetResidentInfoRequest) (_result *GetResidentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentInfoHeaders{}
	_result = &GetResidentInfoResponse{}
	_body, _err := client.GetResidentInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取小区人员信息，包括居民和物业人员
//
// @param request - GetResidentMembersInfoRequest
//
// @param headers - GetResidentMembersInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResidentMembersInfoResponse
func (client *Client) GetResidentMembersInfoWithOptions(request *GetResidentMembersInfoRequest, headers *GetResidentMembersInfoHeaders, runtime *util.RuntimeOptions) (_result *GetResidentMembersInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResidentCropId)) {
		body["residentCropId"] = request.ResidentCropId
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
		Action:      tea.String("GetResidentMembersInfo"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/residences/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResidentMembersInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取小区人员信息，包括居民和物业人员
//
// @param request - GetResidentMembersInfoRequest
//
// @return GetResidentMembersInfoResponse
func (client *Client) GetResidentMembersInfo(request *GetResidentMembersInfoRequest) (_result *GetResidentMembersInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResidentMembersInfoHeaders{}
	_result = &GetResidentMembersInfoResponse{}
	_body, _err := client.GetResidentMembersInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据类型获取部门id
//
// @param request - GetSpaceIdByTypeRequest
//
// @param headers - GetSpaceIdByTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceIdByTypeResponse
func (client *Client) GetSpaceIdByTypeWithOptions(request *GetSpaceIdByTypeRequest, headers *GetSpaceIdByTypeHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceIdByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentType)) {
		query["departmentType"] = request.DepartmentType
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
		Action:      tea.String("GetSpaceIdByType"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces/types"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceIdByTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据类型获取部门id
//
// @param request - GetSpaceIdByTypeRequest
//
// @return GetSpaceIdByTypeResponse
func (client *Client) GetSpaceIdByType(request *GetSpaceIdByTypeRequest) (_result *GetSpaceIdByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceIdByTypeHeaders{}
	_result = &GetSpaceIdByTypeResponse{}
	_body, _err := client.GetSpaceIdByTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取空间信息
//
// @param request - GetSpacesInfoRequest
//
// @param headers - GetSpacesInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpacesInfoResponse
func (client *Client) GetSpacesInfoWithOptions(request *GetSpacesInfoRequest, headers *GetSpacesInfoHeaders, runtime *util.RuntimeOptions) (_result *GetSpacesInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferIds)) {
		body["referIds"] = request.ReferIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResidentCorpId)) {
		body["residentCorpId"] = request.ResidentCorpId
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
		Action:      tea.String("GetSpacesInfo"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpacesInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取空间信息
//
// @param request - GetSpacesInfoRequest
//
// @return GetSpacesInfoResponse
func (client *Client) GetSpacesInfo(request *GetSpacesInfoRequest) (_result *GetSpacesInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpacesInfoHeaders{}
	_result = &GetSpacesInfoResponse{}
	_body, _err := client.GetSpacesInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取行业角色下的用户列表
//
// @param request - ListIndustryRoleUsersRequest
//
// @param headers - ListIndustryRoleUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndustryRoleUsersResponse
func (client *Client) ListIndustryRoleUsersWithOptions(request *ListIndustryRoleUsersRequest, headers *ListIndustryRoleUsersHeaders, runtime *util.RuntimeOptions) (_result *ListIndustryRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		query["tagCode"] = request.TagCode
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
		Action:      tea.String("ListIndustryRoleUsers"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/industryRoles/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndustryRoleUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取行业角色下的用户列表
//
// @param request - ListIndustryRoleUsersRequest
//
// @return ListIndustryRoleUsersResponse
func (client *Client) ListIndustryRoleUsers(request *ListIndustryRoleUsersRequest) (_result *ListIndustryRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListIndustryRoleUsersHeaders{}
	_result = &ListIndustryRoleUsersResponse{}
	_body, _err := client.ListIndustryRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组织维度配置的的积分规则
//
// @param request - ListPointRulesRequest
//
// @param headers - ListPointRulesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPointRulesResponse
func (client *Client) ListPointRulesWithOptions(request *ListPointRulesRequest, headers *ListPointRulesHeaders, runtime *util.RuntimeOptions) (_result *ListPointRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
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
		Action:      tea.String("ListPointRules"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/points/rules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPointRulesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组织维度配置的的积分规则
//
// @param request - ListPointRulesRequest
//
// @return ListPointRulesResponse
func (client *Client) ListPointRules(request *ListPointRulesRequest) (_result *ListPointRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPointRulesHeaders{}
	_result = &ListPointRulesResponse{}
	_body, _err := client.ListPointRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取子空间信息
//
// @param request - ListSubSpaceRequest
//
// @param headers - ListSubSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubSpaceResponse
func (client *Client) ListSubSpaceWithOptions(request *ListSubSpaceRequest, headers *ListSubSpaceHeaders, runtime *util.RuntimeOptions) (_result *ListSubSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferId)) {
		query["referId"] = request.ReferId
	}

	if !tea.BoolValue(util.IsUnset(request.ResidentCorpId)) {
		query["residentCorpId"] = request.ResidentCorpId
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
		Action:      tea.String("ListSubSpace"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces/subSpaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取子空间信息
//
// @param request - ListSubSpaceRequest
//
// @return ListSubSpaceResponse
func (client *Client) ListSubSpace(request *ListSubSpaceRequest) (_result *ListSubSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSubSpaceHeaders{}
	_result = &ListSubSpaceResponse{}
	_body, _err := client.ListSubSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取未确认加入组织的用户
//
// @param request - ListUncheckUsersRequest
//
// @param headers - ListUncheckUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUncheckUsersResponse
func (client *Client) ListUncheckUsersWithOptions(request *ListUncheckUsersRequest, headers *ListUncheckUsersHeaders, runtime *util.RuntimeOptions) (_result *ListUncheckUsersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("ListUncheckUsers"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/organizations/noJoinUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUncheckUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取未确认加入组织的用户
//
// @param request - ListUncheckUsersRequest
//
// @return ListUncheckUsersResponse
func (client *Client) ListUncheckUsers(request *ListUncheckUsersRequest) (_result *ListUncheckUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUncheckUsersHeaders{}
	_result = &ListUncheckUsersResponse{}
	_body, _err := client.ListUncheckUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户行业化角色
//
// @param request - ListUserIndustryRolesRequest
//
// @param headers - ListUserIndustryRolesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserIndustryRolesResponse
func (client *Client) ListUserIndustryRolesWithOptions(request *ListUserIndustryRolesRequest, headers *ListUserIndustryRolesHeaders, runtime *util.RuntimeOptions) (_result *ListUserIndustryRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("ListUserIndustryRoles"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/users/industryRoles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserIndustryRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户行业化角色
//
// @param request - ListUserIndustryRolesRequest
//
// @return ListUserIndustryRolesResponse
func (client *Client) ListUserIndustryRoles(request *ListUserIndustryRolesRequest) (_result *ListUserIndustryRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserIndustryRolesHeaders{}
	_result = &ListUserIndustryRolesResponse{}
	_body, _err := client.ListUserIndustryRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数字区县居民积分流水
//
// @param request - PagePointHistoryRequest
//
// @param headers - PagePointHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PagePointHistoryResponse
func (client *Client) PagePointHistoryWithOptions(request *PagePointHistoryRequest, headers *PagePointHistoryHeaders, runtime *util.RuntimeOptions) (_result *PagePointHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsCircle)) {
		query["isCircle"] = request.IsCircle
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
		Action:      tea.String("PagePointHistory"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/points/records"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PagePointHistoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数字区县居民积分流水
//
// @param request - PagePointHistoryRequest
//
// @return PagePointHistoryResponse
func (client *Client) PagePointHistory(request *PagePointHistoryRequest) (_result *PagePointHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PagePointHistoryHeaders{}
	_result = &PagePointHistoryResponse{}
	_body, _err := client.PagePointHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从空间中删除人员
//
// @param request - RemoveResidentMemberRequest
//
// @param headers - RemoveResidentMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveResidentMemberResponse
func (client *Client) RemoveResidentMemberWithOptions(request *RemoveResidentMemberRequest, headers *RemoveResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *RemoveResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("RemoveResidentMember"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveResidentMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从空间中删除人员
//
// @param request - RemoveResidentMemberRequest
//
// @return RemoveResidentMemberResponse
func (client *Client) RemoveResidentMember(request *RemoveResidentMemberRequest) (_result *RemoveResidentMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveResidentMemberHeaders{}
	_result = &RemoveResidentMemberResponse{}
	_body, _err := client.RemoveResidentMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从户内移除居民
//
// @param request - RemoveResidentUserRequest
//
// @param headers - RemoveResidentUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveResidentUserResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveResidentUser"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/users/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveResidentUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从户内移除居民
//
// @param request - RemoveResidentUserRequest
//
// @return RemoveResidentUserResponse
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

// Summary:
//
// 搜索指定人员
//
// @param request - SearchResidentRequest
//
// @param headers - SearchResidentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResidentResponse
func (client *Client) SearchResidentWithOptions(request *SearchResidentRequest, headers *SearchResidentHeaders, runtime *util.RuntimeOptions) (_result *SearchResidentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResidentCropId)) {
		query["residentCropId"] = request.ResidentCropId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchWord)) {
		query["searchWord"] = request.SearchWord
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
		Action:      tea.String("SearchResident"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/residences"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResidentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索指定人员
//
// @param request - SearchResidentRequest
//
// @return SearchResidentResponse
func (client *Client) SearchResident(request *SearchResidentRequest) (_result *SearchResidentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchResidentHeaders{}
	_result = &SearchResidentResponse{}
	_body, _err := client.SearchResidentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新组信息
//
// @param request - UpdateResideceGroupRequest
//
// @param headers - UpdateResideceGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResideceGroupResponse
func (client *Client) UpdateResideceGroupWithOptions(request *UpdateResideceGroupRequest, headers *UpdateResideceGroupHeaders, runtime *util.RuntimeOptions) (_result *UpdateResideceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
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
		Action:      tea.String("UpdateResideceGroup"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/departments/updateResideceGroup"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新组信息
//
// @param request - UpdateResideceGroupRequest
//
// @return UpdateResideceGroupResponse
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

// Summary:
//
// 更新户信息
//
// @param request - UpdateResidenceRequest
//
// @param headers - UpdateResidenceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResidenceResponse
func (client *Client) UpdateResidenceWithOptions(request *UpdateResidenceRequest, headers *UpdateResidenceHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["departmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.Destitute)) {
		query["destitute"] = request.Destitute
	}

	if !tea.BoolValue(util.IsUnset(request.Grid)) {
		query["grid"] = request.Grid
	}

	if !tea.BoolValue(util.IsUnset(request.HomeTel)) {
		query["homeTel"] = request.HomeTel
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDepartmentId)) {
		query["parentDepartmentId"] = request.ParentDepartmentId
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
		Action:      tea.String("UpdateResidence"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/departments/updateResidece"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResidenceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新户信息
//
// @param request - UpdateResidenceRequest
//
// @return UpdateResidenceResponse
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

// Summary:
//
// 更新小区公告
//
// @param request - UpdateResidentBlackBoardRequest
//
// @param headers - UpdateResidentBlackBoardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResidentBlackBoardResponse
func (client *Client) UpdateResidentBlackBoardWithOptions(request *UpdateResidentBlackBoardRequest, headers *UpdateResidentBlackBoardHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentBlackBoardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackboardId)) {
		body["blackboardId"] = request.BlackboardId
	}

	if !tea.BoolValue(util.IsUnset(request.Context)) {
		body["context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		body["mediaId"] = request.MediaId
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
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResidentBlackBoard"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/blackboards"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResidentBlackBoardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新小区公告
//
// @param request - UpdateResidentBlackBoardRequest
//
// @return UpdateResidentBlackBoardResponse
func (client *Client) UpdateResidentBlackBoard(request *UpdateResidentBlackBoardRequest) (_result *UpdateResidentBlackBoardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidentBlackBoardHeaders{}
	_result = &UpdateResidentBlackBoardResponse{}
	_body, _err := client.UpdateResidentBlackBoardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新小区信息
//
// @param request - UpdateResidentInfoRequest
//
// @param headers - UpdateResidentInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResidentInfoResponse
func (client *Client) UpdateResidentInfoWithOptions(request *UpdateResidentInfoRequest, headers *UpdateResidentInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.BuildingArea)) {
		body["buildingArea"] = request.BuildingArea
	}

	if !tea.BoolValue(util.IsUnset(request.CityName)) {
		body["cityName"] = request.CityName
	}

	if !tea.BoolValue(util.IsUnset(request.CommunityType)) {
		body["communityType"] = request.CommunityType
	}

	if !tea.BoolValue(util.IsUnset(request.CountyName)) {
		body["countyName"] = request.CountyName
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProvName)) {
		body["provName"] = request.ProvName
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		body["state"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		body["telephone"] = request.Telephone
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
		Action:      tea.String("UpdateResidentInfo"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/residences"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResidentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新小区信息
//
// @param request - UpdateResidentInfoRequest
//
// @return UpdateResidentInfoResponse
func (client *Client) UpdateResidentInfo(request *UpdateResidentInfoRequest) (_result *UpdateResidentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidentInfoHeaders{}
	_result = &UpdateResidentInfoResponse{}
	_body, _err := client.UpdateResidentInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新小区成员
//
// @param request - UpdateResidentMemberRequest
//
// @param headers - UpdateResidentMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResidentMemberResponse
func (client *Client) UpdateResidentMemberWithOptions(request *UpdateResidentMemberRequest, headers *UpdateResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResidentUpdateInfo)) {
		body["residentUpdateInfo"] = request.ResidentUpdateInfo
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
		Action:      tea.String("UpdateResidentMember"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/members"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResidentMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新小区成员
//
// @param request - UpdateResidentMemberRequest
//
// @return UpdateResidentMemberResponse
func (client *Client) UpdateResidentMember(request *UpdateResidentMemberRequest) (_result *UpdateResidentMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateResidentMemberHeaders{}
	_result = &UpdateResidentMemberResponse{}
	_body, _err := client.UpdateResidentMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新居民信息
//
// @param request - UpdateResidentUserRequest
//
// @param headers - UpdateResidentUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResidentUserResponse
func (client *Client) UpdateResidentUserWithOptions(request *UpdateResidentUserRequest, headers *UpdateResidentUserHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtField)) {
		query["extField"] = request.ExtField
	}

	if !tea.BoolValue(util.IsUnset(request.IsRetainOldDept)) {
		query["isRetainOldDept"] = request.IsRetainOldDept
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.OldDepartmentId)) {
		query["oldDepartmentId"] = request.OldDepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.RelateType)) {
		query["relateType"] = request.RelateType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["userName"] = request.UserName
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
		Action:      tea.String("UpdateResidentUser"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/users"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新居民信息
//
// @param request - UpdateResidentUserRequest
//
// @return UpdateResidentUserResponse
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

// Summary:
//
// 更新小区空间，含分区，楼栋，单元，房屋等信息
//
// @param request - UpdateSpaceRequest
//
// @param headers - UpdateSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSpaceResponse
func (client *Client) UpdateSpaceWithOptions(request *UpdateSpaceRequest, headers *UpdateSpaceHeaders, runtime *util.RuntimeOptions) (_result *UpdateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SpaceInfoVOList)) {
		body["spaceInfoVOList"] = request.SpaceInfoVOList
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
		Action:      tea.String("UpdateSpace"),
		Version:     tea.String("resident_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/resident/spaces"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新小区空间，含分区，楼栋，单元，房屋等信息
//
// @param request - UpdateSpaceRequest
//
// @return UpdateSpaceResponse
func (client *Client) UpdateSpace(request *UpdateSpaceRequest) (_result *UpdateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSpaceHeaders{}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.UpdateSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
