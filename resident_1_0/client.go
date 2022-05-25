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
	// 增加积分的时间戳毫秒数，如果为空使用系统当前毫秒数
	ActionTime *int64 `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	// 是否查询全员圈积分
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// 规则代码（可空）,如果不为空的话，score值使用ruleCode对应的score增加分数
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 规则名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 本次增加积分：正数表示增加/负数表示扣减
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// 成员id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 加积分的唯一幂等标志
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
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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
	// 部门名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 是否为组
	IsResidenceGroup *bool `json:"isResidenceGroup,omitempty" xml:"isResidenceGroup,omitempty"`
	// 父部门id
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
	// 人员信息
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
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 是否是产权人
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// 人员扩展信息，目前只有租客的起止时间
	MemberDeptExtension map[string]interface{} `json:"memberDeptExtension,omitempty" xml:"memberDeptExtension,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 身份，1：业主，2：租客，3：亲友
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
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddResidentMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResidentMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddResidentMemberResponseBody) SetUserId(v string) *AddResidentMemberResponseBody {
	s.UserId = &v
	return s
}

type AddResidentMemberResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*AddResidentUsersRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 是否是租客
	IsLeaseholder *bool `json:"isLeaseholder,omitempty" xml:"isLeaseholder,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// 居民名字
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
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 扩展字段值
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
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 格式yyyy-MM-dd HH:mm:ss
	SendTime *string `json:"sendTime,omitempty" xml:"sendTime,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 仅当tag未房屋时有用
	BillingArea *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	// 仅当tag未房屋时有用
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	// 仅当tag未房屋时有用
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 仅当tag未房屋时有用
	HouseState *int64 `json:"houseState,omitempty" xml:"houseState,omitempty"`
	// 空间名称，如A栋，二单元，406室等
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父节点id，根节点为-7
	ParentDeptId *string `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// 空间tag标识，目前有House/Unit/Building/Partition
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// * 空间类型为楼时，1住宅/2公寓/3排屋/4洋房/5叠墅/6别墅/7商铺/8办公用房/9经营用房/10其他      * 空间类型为房屋是，1普通住宅/2公寓/3排屋/4物业管理用房/5社区用房/6别墅/7商铺/8办公用房（商写）/9经营用房/10其他
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
	// deptId
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 部门id
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
	// 失败列表
	DelFailedDept []*DeleteSpaceResponseBodyDelFailedDept `json:"delFailedDept,omitempty" xml:"delFailedDept,omitempty" type:"Repeated"`
	// 删除成功数量
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
	// 部门id
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 行业类型
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIndustryTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// dingCropId
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPropertyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 小区地址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 全员群opencid
	AllUserGroupOpenConversationId *string `json:"allUserGroupOpenConversationId,omitempty" xml:"allUserGroupOpenConversationId,omitempty"`
	// 全员群群主 userid
	AllUserGroupOwnerUserId *string  `json:"allUserGroupOwnerUserId,omitempty" xml:"allUserGroupOwnerUserId,omitempty"`
	BuildingArea            *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	// 小区归属的市的id
	CityId *int32 `json:"cityId,omitempty" xml:"cityId,omitempty"`
	// 通信录模式:0标准/1自定义
	ContactMode *int32 `json:"contactMode,omitempty" xml:"contactMode,omitempty"`
	// 小区归属的区/县的id
	CountyId *int32 `json:"countyId,omitempty" xml:"countyId,omitempty"`
	// 交付时间
	DeliveryTime *int64 `json:"deliveryTime,omitempty" xml:"deliveryTime,omitempty"`
	// 经纬度，格式：经度,纬度
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 小区名称
	Name           *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	ProjectManager *GetResidentInfoResponseBodyProjectManager `json:"projectManager,omitempty" xml:"projectManager,omitempty" type:"Struct"`
	// 物业部门群cid
	PropertyDeptGroupOpenConversationId *string `json:"propertyDeptGroupOpenConversationId,omitempty" xml:"propertyDeptGroupOpenConversationId,omitempty"`
	// 物业部门群主userid
	PropertyDeptGroupOwnerUserId *string `json:"propertyDeptGroupOwnerUserId,omitempty" xml:"propertyDeptGroupOwnerUserId,omitempty"`
	// 小区归属的省的id
	ProvId *int64 `json:"provId,omitempty" xml:"provId,omitempty"`
	// 物业管理范围-东
	ScopeEast *string `json:"scopeEast,omitempty" xml:"scopeEast,omitempty"`
	// 物业管理范围-北
	ScopeNorth *string `json:"scopeNorth,omitempty" xml:"scopeNorth,omitempty"`
	// 物业管理范围-南
	ScopeSouth *string `json:"scopeSouth,omitempty" xml:"scopeSouth,omitempty"`
	// 物业管理范围-西
	ScopeWest *string `json:"scopeWest,omitempty" xml:"scopeWest,omitempty"`
	// 物业电话
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 小区归属的街道/镇的id
	TownId *int32 `json:"townId,omitempty" xml:"townId,omitempty"`
	// 1纯住宅；2:商住混合；3:办公；4:办公商业混合；5:商业；6:公共场所；7:其他
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
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
	// 头像
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 人员唯一标识
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 姓名
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResidentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ResidentCropId *string   `json:"residentCropId,omitempty" xml:"residentCropId,omitempty"`
	UserIdList     []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
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
	// result
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
	// 是否激活
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 扩展字段，如果是租客存起止时间
	ExtField *string `json:"extField,omitempty" xml:"extField,omitempty"`
	// 是否是产权人
	IsPropertyOwner *bool   `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	// 业主/租客/亲友等
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResidentMembersInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 部门类型
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
	// 部门id
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpaceIdByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ReferIds       []*int64 `json:"referIds,omitempty" xml:"referIds,omitempty" type:"Repeated"`
	ResidentCorpId *string  `json:"residentCorpId,omitempty" xml:"residentCorpId,omitempty"`
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
	// result
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
	BillingArea  *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	Floor        *string  `json:"floor,omitempty" xml:"floor,omitempty"`
	// 房屋状态：0空置/1未领/2入住/3空关/4装修
	HouseState    *int32  `json:"houseState,omitempty" xml:"houseState,omitempty"`
	IsVirtual     *int32  `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	ParentReferId *int64  `json:"parentReferId,omitempty" xml:"parentReferId,omitempty"`
	ReferId       *int64  `json:"referId,omitempty" xml:"referId,omitempty"`
	SpaceName     *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	TagCode       *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// 空间类型为楼时，1高层/2低层/3别墅/4其他，空间类型为房屋是，1住宅/2公寓/3排屋/4洋房/5叠墅/6别墅/7商铺/8办公用房/9经营用房/10其他
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpacesInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetSpacesInfoResponse) SetBody(v *GetSpacesInfoResponseBody) *GetSpacesInfoResponse {
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
	// 是否查询全员圈积分
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
	// 查询所得积分规则集合
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
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 单日计次上限，0表示无上限
	DayLimitTimes *int32 `json:"dayLimitTimes,omitempty" xml:"dayLimitTimes,omitempty"`
	// 扩展字段
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 分组ID, 默认写入为0
	GroupId *int32 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 排序ID
	OrderId *int32 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 对应的行为代码（可空）
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 对应的行为名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 增加或减少的分数（增加为正数，减少为负数）
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// 生效状态 0：不生效，1：生效
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPointRulesResponseBodyPointRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPointRulesResponseBodyPointRuleList) GoString() string {
	return s.String()
}

func (s *ListPointRulesResponseBodyPointRuleList) SetCorpId(v string) *ListPointRulesResponseBodyPointRuleList {
	s.CorpId = &v
	return s
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPointRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ReferId *int64 `json:"referId,omitempty" xml:"referId,omitempty"`
	// A short description of struct
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
	// result
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
	BillingArea  *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	Floor        *string  `json:"floor,omitempty" xml:"floor,omitempty"`
	// 房屋状态：0空置/1未领/2入住/3空关/4装修
	HouseState    *int32  `json:"houseState,omitempty" xml:"houseState,omitempty"`
	IsVirtual     *int32  `json:"isVirtual,omitempty" xml:"isVirtual,omitempty"`
	ParentReferId *int64  `json:"parentReferId,omitempty" xml:"parentReferId,omitempty"`
	ReferId       *int64  `json:"referId,omitempty" xml:"referId,omitempty"`
	SpaceName     *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	TagCode       *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// 空间类型为楼时，1高层/2低层/3别墅/4其他，空间类型为房屋是，1住宅/2公寓/3排屋/4洋房/5叠墅/6别墅/7商铺/8办公用房/9经营用房/10其他
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 条目数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 开始位置
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 起始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态
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
	// 是否仍有数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一个数据序号
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// result
	Values []*ListUncheckUsersResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 扩展信息
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 邀请时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 邀请更新时间
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 是否产权人
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// 用户名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 状态
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 用户ID
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUncheckUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUncheckUsersResponse) SetBody(v *ListUncheckUsersResponseBody) *ListUncheckUsersResponse {
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
	// 结束时间Unix时间戳（不包含），可空
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 是否查询全员圈积分
	IsCircle *bool `json:"isCircle,omitempty" xml:"isCircle,omitempty"`
	// 本次读取的最大数据记录数量，最大20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用来标记当前开始读取的位置
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 起始时间Unix时间戳，可空
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 用户userid，可空，不传表示查询组织内所有用户的流水数据
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
	// 是否有下一页
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一个游标值
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 查询所得积分流水集合
	PointRecordList []*PagePointHistoryResponseBodyPointRecordList `json:"pointRecordList,omitempty" xml:"pointRecordList,omitempty" type:"Repeated"`
	// 总数，如果为-1则不计算总数
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
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间（精确到毫秒数）
	CreateAt *int64 `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// 对应的行为代码（可空）
	RuleCode *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// 对应的行为名字
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// 增加或减少的分数（增加为正数，减少为负数）
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// 成员id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 幂等键
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s PagePointHistoryResponseBodyPointRecordList) String() string {
	return tea.Prettify(s)
}

func (s PagePointHistoryResponseBodyPointRecordList) GoString() string {
	return s.String()
}

func (s *PagePointHistoryResponseBodyPointRecordList) SetCorpId(v string) *PagePointHistoryResponseBodyPointRecordList {
	s.CorpId = &v
	return s
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PagePointHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 空位标识
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 人员标识
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

func (s *RemoveResidentMemberRequest) SetUserId(v string) *RemoveResidentMemberRequest {
	s.UserId = &v
	return s
}

type RemoveResidentMemberResponseBody struct {
	// 是否成功
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ResidentCropId *string `json:"residentCropId,omitempty" xml:"residentCropId,omitempty"`
	SearchWord     *string `json:"searchWord,omitempty" xml:"searchWord,omitempty"`
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
	// result
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
	// 是否激活
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 扩展字段，如果是租客存起止时间
	ExtField *string `json:"extField,omitempty" xml:"extField,omitempty"`
	// 是否是产权人
	IsPropertyOwner *bool   `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	// 业主/租客/亲友等
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchResidentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 组名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 组长userid
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
	// 组id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 户名字
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 是否是贫困户
	Destitute *bool `json:"destitute,omitempty" xml:"destitute,omitempty"`
	// 所属网格
	Grid *string `json:"grid,omitempty" xml:"grid,omitempty"`
	// 家庭电话
	HomeTel *string `json:"homeTel,omitempty" xml:"homeTel,omitempty"`
	// 家庭管理员用户id
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// 组id
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
	Context      *string `json:"context,omitempty" xml:"context,omitempty"`
	MediaId      *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidentBlackBoardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 建筑面积，组多支持2位小数，总长不超过8位
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	// 市的名字，有值时provName必填
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// 1纯住宅；2:商住混合；3:办公；4:办公商业混合；5:商业；6:公共场所；7:其他
	CommunityType *int64 `json:"communityType,omitempty" xml:"communityType,omitempty"`
	// 区/县名，有值是provName，cityName必填
	CountyName *string `json:"countyName,omitempty" xml:"countyName,omitempty"`
	// 经纬度，格式：经度,纬度
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// 小区名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 省的名字
	ProvName *string `json:"provName,omitempty" xml:"provName,omitempty"`
	// 小区状态：0正常/1关闭/2作废
	State *int64 `json:"state,omitempty" xml:"state,omitempty"`
	// 小区服务电话
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
	// 是否成功
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 人员更新信息
	ResidentUpdateInfo *UpdateResidentMemberRequestResidentUpdateInfo `json:"residentUpdateInfo,omitempty" xml:"residentUpdateInfo,omitempty" type:"Struct"`
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

type UpdateResidentMemberRequestResidentUpdateInfo struct {
	// 部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 是否是产权人
	IsPropertyOwner *bool `json:"isPropertyOwner,omitempty" xml:"isPropertyOwner,omitempty"`
	// 是否保留旧部门，默认不保存
	IsRetainOldDept *bool `json:"isRetainOldDept,omitempty" xml:"isRetainOldDept,omitempty"`
	// 人员扩展信息，目前只有租客的起止时间
	MemberDeptExtension map[string]interface{} `json:"memberDeptExtension,omitempty" xml:"memberDeptExtension,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 旧部门id
	OldDeptId *int64 `json:"oldDeptId,omitempty" xml:"oldDeptId,omitempty"`
	// 身份，1：业主，2：租客，3：亲友
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// 用户id
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

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetIsRetainOldDept(v bool) *UpdateResidentMemberRequestResidentUpdateInfo {
	s.IsRetainOldDept = &v
	return s
}

func (s *UpdateResidentMemberRequestResidentUpdateInfo) SetMemberDeptExtension(v map[string]interface{}) *UpdateResidentMemberRequestResidentUpdateInfo {
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
	// 是否成功
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResidentMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 家庭住址
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 所在新的户/租户部门id
	DepartmentId *int64 `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 扩展字段（包括身份证/性别/民族）
	ExtField []*UpdateResidentUserRequestExtField `json:"extField,omitempty" xml:"extField,omitempty" type:"Repeated"`
	// 是否保留原部门
	IsRetainOldDept *bool `json:"isRetainOldDept,omitempty" xml:"isRetainOldDept,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 原所在部门id
	OldDepartmentId *int64 `json:"oldDepartmentId,omitempty" xml:"oldDepartmentId,omitempty"`
	// 与户主的关系
	RelateType *string `json:"relateType,omitempty" xml:"relateType,omitempty"`
	// 人员userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 居民名字
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
	// 扩展字段名字
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 扩展字段值
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
	// 修改后空间信息
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
	// 计费面积
	BillingArea *float32 `json:"billingArea,omitempty" xml:"billingArea,omitempty"`
	// 建筑面积
	BuildingArea *float32 `json:"buildingArea,omitempty" xml:"buildingArea,omitempty"`
	// 楼栋类型
	BuildingType *int64 `json:"buildingType,omitempty" xml:"buildingType,omitempty"`
	// 修改的空间的唯一标识
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 房屋所在楼层，当tagCode为House时选填
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 房屋状态，tagcode为house时选填，0空置/1未领/2入住/3空关/4装修
	HouseState *int64 `json:"houseState,omitempty" xml:"houseState,omitempty"`
	// 房屋类型，当tagcode为House时必填
	HouseType *int64 `json:"houseType,omitempty" xml:"houseType,omitempty"`
	// 修改后名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父节点id
	ParentDeptId *float32 `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	// 空间类型
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

func (s *UpdateSpaceRequestSpaceInfoVOList) SetParentDeptId(v float32) *UpdateSpaceRequestSpaceInfoVOList {
	s.ParentDeptId = &v
	return s
}

func (s *UpdateSpaceRequestSpaceInfoVOList) SetTagCode(v string) *UpdateSpaceRequestSpaceInfoVOList {
	s.TagCode = &v
	return s
}

type UpdateSpaceResponseBody struct {
	// 是否成功
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

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
	_result = &AddPointResponse{}
	_body, _err := client.DoROARequest(tea.String("AddPoint"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/points"), tea.String("none"), req, runtime)
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
	_result = &AddResidentDepartmentResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentDepartment"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/departments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AddResidentMemberWithOptions(request *AddResidentMemberRequest, headers *AddResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *AddResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ResidentAddInfo))) {
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
	_result = &AddResidentMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentMember"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/members"), tea.String("json"), req, runtime)
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
	_result = &AddResidentUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("AddResidentUsers"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateResidentBlackBoardResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateResidentBlackBoard"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/blackboards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &CreateSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSpace"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteResidentBlackBoardResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteResidentBlackBoard"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/resident/blackboards"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSpace"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/spaces/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetConversationIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetConversationId"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/conversations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetIndustryTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetIndustryType"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/organizations/industryTypes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetPropertyInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPropertyInfo"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/propertyInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetResidentInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetResidentInfo"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/residentInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetResidentMembersInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetResidentMembersInfo"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/residences/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSpaceIdByTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpaceIdByType"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/spaces/types"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetSpacesInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSpacesInfo"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/spaces/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListPointRulesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListPointRules"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/points/rules"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListSubSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSubSpace"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/spaces/subSpaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListUncheckUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListUncheckUsers"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/organizations/noJoinUsers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &PagePointHistoryResponse{}
	_body, _err := client.DoROARequest(tea.String("PagePointHistory"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/points/records"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) RemoveResidentMemberWithOptions(request *RemoveResidentMemberRequest, headers *RemoveResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *RemoveResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
	_result = &RemoveResidentMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveResidentMember"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/resident/members/remove"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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
	_result = &SearchResidentResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchResident"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/resident/residences"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	_result = &UpdateResideceGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResideceGroup"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResideceGroup"), tea.String("json"), req, runtime)
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
	_result = &UpdateResidenceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidence"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/departments/updateResidece"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateResidentBlackBoardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentBlackBoard"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/blackboards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateResidentInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentInfo"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/residences"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateResidentMemberWithOptions(request *UpdateResidentMemberRequest, headers *UpdateResidentMemberHeaders, runtime *util.RuntimeOptions) (_result *UpdateResidentMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ResidentUpdateInfo))) {
		body["residentUpdateInfo"] = request.ResidentUpdateInfo
	}

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
	_result = &UpdateResidentMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentMember"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/members"), tea.String("json"), req, runtime)
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
	_result = &UpdateResidentUserResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateResidentUser"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/users"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSpace"), tea.String("resident_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/resident/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
