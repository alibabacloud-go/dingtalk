// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package okr_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AlignObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AlignObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *AlignObjectiveHeaders) SetCommonHeaders(v map[string]*string) *AlignObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AlignObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *AlignObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AlignObjectiveRequest struct {
	// 周期 ID。
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 对齐目标的 ID。
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 当前用户的 user ID。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AlignObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveRequest) GoString() string {
	return s.String()
}

func (s *AlignObjectiveRequest) SetPeriodId(v string) *AlignObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *AlignObjectiveRequest) SetTargetId(v string) *AlignObjectiveRequest {
	s.TargetId = &v
	return s
}

func (s *AlignObjectiveRequest) SetUserId(v string) *AlignObjectiveRequest {
	s.UserId = &v
	return s
}

type AlignObjectiveResponseBody struct {
	// data
	Data *AlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponseBody) SetData(v *AlignObjectiveResponseBodyData) *AlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *AlignObjectiveResponseBody) SetSuccess(v bool) *AlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type AlignObjectiveResponseBodyData struct {
	// 对齐目标的 ID。
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	// 当前 Objective 的ID
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
}

func (s AlignObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponseBodyData) SetAlignId(v io.Reader) *AlignObjectiveResponseBodyData {
	s.AlignId = v
	return s
}

func (s *AlignObjectiveResponseBodyData) SetId(v io.Reader) *AlignObjectiveResponseBodyData {
	s.Id = v
	return s
}

type AlignObjectiveResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AlignObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AlignObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponse) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponse) SetHeaders(v map[string]*string) *AlignObjectiveResponse {
	s.Headers = v
	return s
}

func (s *AlignObjectiveResponse) SetBody(v *AlignObjectiveResponseBody) *AlignObjectiveResponse {
	s.Body = v
	return s
}

type BatchAddPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionHeaders) SetCommonHeaders(v map[string]*string) *BatchAddPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddPermissionRequest struct {
	List       []*BatchAddPermissionRequestList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	TargetId   *string                          `json:"targetId,omitempty" xml:"targetId,omitempty"`
	TargetType *string                          `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// A short description of struct
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchAddPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequest) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequest) SetList(v []*BatchAddPermissionRequestList) *BatchAddPermissionRequest {
	s.List = v
	return s
}

func (s *BatchAddPermissionRequest) SetTargetId(v string) *BatchAddPermissionRequest {
	s.TargetId = &v
	return s
}

func (s *BatchAddPermissionRequest) SetTargetType(v string) *BatchAddPermissionRequest {
	s.TargetType = &v
	return s
}

func (s *BatchAddPermissionRequest) SetUserId(v string) *BatchAddPermissionRequest {
	s.UserId = &v
	return s
}

type BatchAddPermissionRequestList struct {
	Member     *BatchAddPermissionRequestListMember `json:"member,omitempty" xml:"member,omitempty" type:"Struct"`
	PolicyType *int64                               `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s BatchAddPermissionRequestList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequestList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequestList) SetMember(v *BatchAddPermissionRequestListMember) *BatchAddPermissionRequestList {
	s.Member = v
	return s
}

func (s *BatchAddPermissionRequestList) SetPolicyType(v int64) *BatchAddPermissionRequestList {
	s.PolicyType = &v
	return s
}

type BatchAddPermissionRequestListMember struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionRequestListMember) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionRequestListMember) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionRequestListMember) SetId(v string) *BatchAddPermissionRequestListMember {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionRequestListMember) SetType(v string) *BatchAddPermissionRequestListMember {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBody struct {
	// 返回的数据。
	Data *BatchAddPermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBody) SetData(v *BatchAddPermissionResponseBodyData) *BatchAddPermissionResponseBody {
	s.Data = v
	return s
}

func (s *BatchAddPermissionResponseBody) SetSuccess(v bool) *BatchAddPermissionResponseBody {
	s.Success = &v
	return s
}

type BatchAddPermissionResponseBodyData struct {
	// 是否有无效的成员。
	HasInvalidUser *bool `json:"hasInvalidUser,omitempty" xml:"hasInvalidUser,omitempty"`
	// 权限信息。
	PermissionTree *BatchAddPermissionResponseBodyDataPermissionTree `json:"permissionTree,omitempty" xml:"permissionTree,omitempty" type:"Struct"`
}

func (s BatchAddPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyData) SetHasInvalidUser(v bool) *BatchAddPermissionResponseBodyData {
	s.HasInvalidUser = &v
	return s
}

func (s *BatchAddPermissionResponseBodyData) SetPermissionTree(v *BatchAddPermissionResponseBodyDataPermissionTree) *BatchAddPermissionResponseBodyData {
	s.PermissionTree = v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTree struct {
	// 权限 ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限列表
	PolicyList []*BatchAddPermissionResponseBodyDataPermissionTreePolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// 是否可见的标识。
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 哪种类型的权限。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTree) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTree) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetId(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetPolicyList(v []*BatchAddPermissionResponseBodyDataPermissionTreePolicyList) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.PolicyList = v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetPrivacy(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Privacy = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTree) SetType(v string) *BatchAddPermissionResponseBodyDataPermissionTree {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTreePolicyList struct {
	MemberList []*BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	Name       *string                                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Type       *int64                                                                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetMemberList(v []*BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.MemberList = v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetName(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.Name = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyList) SetType(v int64) *BatchAddPermissionResponseBodyDataPermissionTreePolicyList {
	s.Type = &v
	return s
}

type BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList struct {
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetId(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Id = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetNickname(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList) SetType(v string) *BatchAddPermissionResponseBodyDataPermissionTreePolicyListMemberList {
	s.Type = &v
	return s
}

type BatchAddPermissionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAddPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddPermissionResponse) GoString() string {
	return s.String()
}

func (s *BatchAddPermissionResponse) SetHeaders(v map[string]*string) *BatchAddPermissionResponse {
	s.Headers = v
	return s
}

func (s *BatchAddPermissionResponse) SetBody(v *BatchAddPermissionResponseBody) *BatchAddPermissionResponse {
	s.Body = v
	return s
}

type BatchQueryObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryObjectiveRequest struct {
	// 需要查看的 Objective ID。
	ObjectiveIds []*string `json:"objectiveIds,omitempty" xml:"objectiveIds,omitempty" type:"Repeated"`
	// 周期 ID。
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 是否返回关联信息。
	WithAlign *bool `json:"withAlign,omitempty" xml:"withAlign,omitempty"`
	// 是否返回 KR 信息。
	WithKr *bool `json:"withKr,omitempty" xml:"withKr,omitempty"`
	// 是否返回进度信息
	WithProgress *bool `json:"withProgress,omitempty" xml:"withProgress,omitempty"`
	// 当前用户的 staff ID。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveRequest) SetObjectiveIds(v []*string) *BatchQueryObjectiveRequest {
	s.ObjectiveIds = v
	return s
}

func (s *BatchQueryObjectiveRequest) SetPeriodId(v string) *BatchQueryObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithAlign(v bool) *BatchQueryObjectiveRequest {
	s.WithAlign = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithKr(v bool) *BatchQueryObjectiveRequest {
	s.WithKr = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetWithProgress(v bool) *BatchQueryObjectiveRequest {
	s.WithProgress = &v
	return s
}

func (s *BatchQueryObjectiveRequest) SetUserId(v string) *BatchQueryObjectiveRequest {
	s.UserId = &v
	return s
}

type BatchQueryObjectiveResponseBody struct {
	// data
	Data []*BatchQueryObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBody) SetData(v []*BatchQueryObjectiveResponseBodyData) *BatchQueryObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryObjectiveResponseBody) SetSuccess(v bool) *BatchQueryObjectiveResponseBody {
	s.Success = &v
	return s
}

type BatchQueryObjectiveResponseBodyData struct {
	// 被对齐的 Objective。
	AlignFromIds []io.Reader `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	// 对齐的 Objective。
	AlignToIds []io.Reader `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	// Objective 内容。
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间。时间戳
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 更新时间。时间戳
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// objective。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// KR 详情列表。
	KrList []*BatchQueryObjectiveResponseBodyDataKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	// 所属者信息。
	Owner *BatchQueryObjectiveResponseBodyDataOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 周期 ID。
	PeriodId io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 权限值。
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// 所在位置。
	Position *int32 `json:"position,omitempty" xml:"position,omitempty"`
	// 进度值。
	Progress *BatchQueryObjectiveResponseBodyDataProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// 百分比值。
	ProgressPercent *float32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// 是否已发布。
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// 分数值。
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// 当前内容状态。
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 用户 ID。
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
	// 权重值。
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyData) SetAlignFromIds(v []io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.AlignFromIds = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetAlignToIds(v []io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.AlignToIds = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetContent(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.Content = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetGmtCreate(v float32) *BatchQueryObjectiveResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetGmtModified(v float32) *BatchQueryObjectiveResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetKrList(v []*BatchQueryObjectiveResponseBodyDataKrList) *BatchQueryObjectiveResponseBodyData {
	s.KrList = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetOwner(v *BatchQueryObjectiveResponseBodyDataOwner) *BatchQueryObjectiveResponseBodyData {
	s.Owner = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPeriodId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.PeriodId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPermission(v []*float32) *BatchQueryObjectiveResponseBodyData {
	s.Permission = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPosition(v int32) *BatchQueryObjectiveResponseBodyData {
	s.Position = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetProgress(v *BatchQueryObjectiveResponseBodyDataProgress) *BatchQueryObjectiveResponseBodyData {
	s.Progress = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetProgressPercent(v float32) *BatchQueryObjectiveResponseBodyData {
	s.ProgressPercent = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetPublished(v bool) *BatchQueryObjectiveResponseBodyData {
	s.Published = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetScore(v float32) *BatchQueryObjectiveResponseBodyData {
	s.Score = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetStatus(v int32) *BatchQueryObjectiveResponseBodyData {
	s.Status = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetUserId(v io.Reader) *BatchQueryObjectiveResponseBodyData {
	s.UserId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyData) SetWeight(v float32) *BatchQueryObjectiveResponseBodyData {
	s.Weight = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataKrList struct {
	// KR 内容。
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间。时间戳
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 更新时间。时间戳
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// KR 的 ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// 所属 Objective ID。
	ObjectiveId io.Reader `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// KR 权限。
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// 所处位置。
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// KR 进度。
	Progress *BatchQueryObjectiveResponseBodyDataKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// 所占分数。
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// 所占权重。
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataKrList) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataKrList) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetContent(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Content = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetGmtCreate(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.GmtCreate = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetGmtModified(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.GmtModified = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetObjectiveId(v io.Reader) *BatchQueryObjectiveResponseBodyDataKrList {
	s.ObjectiveId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetPermission(v []*float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Permission = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetPosition(v int64) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Position = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetProgress(v *BatchQueryObjectiveResponseBodyDataKrListProgress) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Progress = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetScore(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Score = &v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataKrList) SetWeight(v float32) *BatchQueryObjectiveResponseBodyDataKrList {
	s.Weight = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataKrListProgress struct {
	// 百分比。
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataKrListProgress) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataKrListProgress) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataKrListProgress) SetPercent(v int32) *BatchQueryObjectiveResponseBodyDataKrListProgress {
	s.Percent = &v
	return s
}

type BatchQueryObjectiveResponseBodyDataOwner struct {
	// 所属者头像。 ID
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 所属者组织 I。D
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 所属者在 OKR 系统中的 ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// 所属者昵称。
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 所属者 userId。
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataOwner) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataOwner) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetAvatarMediaId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.AvatarMediaId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetCorpId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.CorpId = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetNickname(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.Nickname = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetUserId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.UserId = v
	return s
}

type BatchQueryObjectiveResponseBodyDataProgress struct {
	// 百分比。
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataProgress) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataProgress) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataProgress) SetPercent(v int32) *BatchQueryObjectiveResponseBodyDataProgress {
	s.Percent = &v
	return s
}

type BatchQueryObjectiveResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQueryObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQueryObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponse) SetHeaders(v map[string]*string) *BatchQueryObjectiveResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryObjectiveResponse) SetBody(v *BatchQueryObjectiveResponseBody) *BatchQueryObjectiveResponse {
	s.Body = v
	return s
}

type BatchQueryUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryUserHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryUserHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryUserRequest struct {
	// OKR 系统中的用户 ID 列表
	OkrUserIds []*string `json:"okrUserIds,omitempty" xml:"okrUserIds,omitempty" type:"Repeated"`
	// 开放平台中用户 ID 列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s BatchQueryUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryUserRequest) SetOkrUserIds(v []*string) *BatchQueryUserRequest {
	s.OkrUserIds = v
	return s
}

func (s *BatchQueryUserRequest) SetUserIds(v []*string) *BatchQueryUserRequest {
	s.UserIds = v
	return s
}

type BatchQueryUserResponseBody struct {
	// data
	Data []*BatchQueryUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponseBody) SetData(v []*BatchQueryUserResponseBodyData) *BatchQueryUserResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryUserResponseBody) SetSuccess(v bool) *BatchQueryUserResponseBody {
	s.Success = &v
	return s
}

type BatchQueryUserResponseBodyData struct {
	// 所属者头像。 ID
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 所属者头像。 URL
	AvatarUrl io.Reader `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 所属者组织 I。D
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 所属者在 OKR 系统中的 ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// 所属者昵称。
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 所属者 userId。
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchQueryUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponseBodyData) SetAvatarMediaId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.AvatarMediaId = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetAvatarUrl(v io.Reader) *BatchQueryUserResponseBodyData {
	s.AvatarUrl = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetCorpId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.CorpId = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.Id = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetNickname(v io.Reader) *BatchQueryUserResponseBodyData {
	s.Nickname = v
	return s
}

func (s *BatchQueryUserResponseBodyData) SetUserId(v io.Reader) *BatchQueryUserResponseBodyData {
	s.UserId = v
	return s
}

type BatchQueryUserResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQueryUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQueryUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryUserResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryUserResponse) SetHeaders(v map[string]*string) *BatchQueryUserResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryUserResponse) SetBody(v *BatchQueryUserResponseBody) *BatchQueryUserResponse {
	s.Body = v
	return s
}

type CreateKeyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateKeyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultHeaders) GoString() string {
	return s.String()
}

func (s *CreateKeyResultHeaders) SetCommonHeaders(v map[string]*string) *CreateKeyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateKeyResultHeaders) SetXAcsDingtalkAccessToken(v string) *CreateKeyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateKeyResultRequest struct {
	// KR 内容。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 所属 Objective ID。
	ObjectiveId *string `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// 周期 ID。
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 上一个 KR 的位置。
	PrevPosition *int64 `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	// KR 的权重比。
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 当前用户的 user ID。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyResultRequest) SetContent(v string) *CreateKeyResultRequest {
	s.Content = &v
	return s
}

func (s *CreateKeyResultRequest) SetObjectiveId(v string) *CreateKeyResultRequest {
	s.ObjectiveId = &v
	return s
}

func (s *CreateKeyResultRequest) SetPeriodId(v string) *CreateKeyResultRequest {
	s.PeriodId = &v
	return s
}

func (s *CreateKeyResultRequest) SetPrevPosition(v int64) *CreateKeyResultRequest {
	s.PrevPosition = &v
	return s
}

func (s *CreateKeyResultRequest) SetWeight(v int64) *CreateKeyResultRequest {
	s.Weight = &v
	return s
}

func (s *CreateKeyResultRequest) SetUserId(v string) *CreateKeyResultRequest {
	s.UserId = &v
	return s
}

type CreateKeyResultResponseBody struct {
	Data *CreateKeyResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateKeyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponseBody) SetData(v *CreateKeyResultResponseBodyData) *CreateKeyResultResponseBody {
	s.Data = v
	return s
}

func (s *CreateKeyResultResponseBody) SetSuccess(v bool) *CreateKeyResultResponseBody {
	s.Success = &v
	return s
}

type CreateKeyResultResponseBodyData struct {
	// 创建成功的 KR ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// KR的位置。
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// KR 的权重。
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateKeyResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponseBodyData) SetId(v io.Reader) *CreateKeyResultResponseBodyData {
	s.Id = v
	return s
}

func (s *CreateKeyResultResponseBodyData) SetPosition(v int64) *CreateKeyResultResponseBodyData {
	s.Position = &v
	return s
}

func (s *CreateKeyResultResponseBodyData) SetWeight(v int64) *CreateKeyResultResponseBodyData {
	s.Weight = &v
	return s
}

type CreateKeyResultResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateKeyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateKeyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyResultResponse) SetHeaders(v map[string]*string) *CreateKeyResultResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyResultResponse) SetBody(v *CreateKeyResultResponseBody) *CreateKeyResultResponse {
	s.Body = v
	return s
}

type CreateObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *CreateObjectiveHeaders) SetCommonHeaders(v map[string]*string) *CreateObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *CreateObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateObjectiveRequest struct {
	// 创建Objective 的内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 当前周期 ID。
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 上一个 Objective 的位置。
	PrevPosition *string `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	// 当前用户的 userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *CreateObjectiveRequest) SetContent(v string) *CreateObjectiveRequest {
	s.Content = &v
	return s
}

func (s *CreateObjectiveRequest) SetPeriodId(v string) *CreateObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *CreateObjectiveRequest) SetPrevPosition(v string) *CreateObjectiveRequest {
	s.PrevPosition = &v
	return s
}

func (s *CreateObjectiveRequest) SetUserId(v string) *CreateObjectiveRequest {
	s.UserId = &v
	return s
}

type CreateObjectiveResponseBody struct {
	// data
	Data *CreateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponseBody) SetData(v *CreateObjectiveResponseBodyData) *CreateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *CreateObjectiveResponseBody) SetSuccess(v bool) *CreateObjectiveResponseBody {
	s.Success = &v
	return s
}

type CreateObjectiveResponseBodyData struct {
	// 当前 Objective ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 当前 Objective 的位置。
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
}

func (s CreateObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponseBodyData) SetId(v string) *CreateObjectiveResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateObjectiveResponseBodyData) SetPosition(v string) *CreateObjectiveResponseBodyData {
	s.Position = &v
	return s
}

type CreateObjectiveResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponse) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponse) SetHeaders(v map[string]*string) *CreateObjectiveResponse {
	s.Headers = v
	return s
}

func (s *CreateObjectiveResponse) SetBody(v *CreateObjectiveResponseBody) *CreateObjectiveResponse {
	s.Body = v
	return s
}

type DeleteKeyResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteKeyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultHeaders) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultHeaders) SetCommonHeaders(v map[string]*string) *DeleteKeyResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteKeyResultHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteKeyResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteKeyResultRequest struct {
	// 当前 KR id。
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultRequest) SetKrId(v string) *DeleteKeyResultRequest {
	s.KrId = &v
	return s
}

func (s *DeleteKeyResultRequest) SetUserId(v string) *DeleteKeyResultRequest {
	s.UserId = &v
	return s
}

type DeleteKeyResultResponseBody struct {
	// 返回的信息
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteKeyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultResponseBody) SetData(v bool) *DeleteKeyResultResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeyResultResponseBody) SetSuccess(v bool) *DeleteKeyResultResponseBody {
	s.Success = &v
	return s
}

type DeleteKeyResultResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteKeyResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKeyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultResponse) SetHeaders(v map[string]*string) *DeleteKeyResultResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyResultResponse) SetBody(v *DeleteKeyResultResponseBody) *DeleteKeyResultResponse {
	s.Body = v
	return s
}

type DeleteObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveHeaders) SetCommonHeaders(v map[string]*string) *DeleteObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteObjectiveRequest struct {
	// 当前用户的 userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveRequest) SetUserId(v string) *DeleteObjectiveRequest {
	s.UserId = &v
	return s
}

type DeleteObjectiveResponseBody struct {
	// data
	Data *DeleteObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponseBody) SetData(v *DeleteObjectiveResponseBodyData) *DeleteObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *DeleteObjectiveResponseBody) SetSuccess(v bool) *DeleteObjectiveResponseBody {
	s.Success = &v
	return s
}

type DeleteObjectiveResponseBodyData struct {
	// 当前 Objective ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DeleteObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponseBodyData) SetId(v string) *DeleteObjectiveResponseBodyData {
	s.Id = &v
	return s
}

type DeleteObjectiveResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponse) SetHeaders(v map[string]*string) *DeleteObjectiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectiveResponse) SetBody(v *DeleteObjectiveResponseBody) *DeleteObjectiveResponse {
	s.Body = v
	return s
}

type DeletePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeletePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionHeaders) GoString() string {
	return s.String()
}

func (s *DeletePermissionHeaders) SetCommonHeaders(v map[string]*string) *DeletePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *DeletePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeletePermissionRequest struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	PolicyType *int64  `json:"policyType,omitempty" xml:"policyType,omitempty"`
	TargetId   *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeletePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequest) SetId(v string) *DeletePermissionRequest {
	s.Id = &v
	return s
}

func (s *DeletePermissionRequest) SetPolicyType(v int64) *DeletePermissionRequest {
	s.PolicyType = &v
	return s
}

func (s *DeletePermissionRequest) SetTargetId(v string) *DeletePermissionRequest {
	s.TargetId = &v
	return s
}

func (s *DeletePermissionRequest) SetTargetType(v string) *DeletePermissionRequest {
	s.TargetType = &v
	return s
}

func (s *DeletePermissionRequest) SetType(v string) *DeletePermissionRequest {
	s.Type = &v
	return s
}

func (s *DeletePermissionRequest) SetUserId(v string) *DeletePermissionRequest {
	s.UserId = &v
	return s
}

type DeletePermissionResponseBody struct {
	Data *DeletePermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeletePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBody) SetData(v *DeletePermissionResponseBodyData) *DeletePermissionResponseBody {
	s.Data = v
	return s
}

func (s *DeletePermissionResponseBody) SetSuccess(v bool) *DeletePermissionResponseBody {
	s.Success = &v
	return s
}

type DeletePermissionResponseBodyData struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限列表
	PolicyList []*DeletePermissionResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// 是否可见的标识。
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 哪种类型的权限。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyData) SetId(v string) *DeletePermissionResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeletePermissionResponseBodyData) SetPolicyList(v []*DeletePermissionResponseBodyDataPolicyList) *DeletePermissionResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *DeletePermissionResponseBodyData) SetPrivacy(v string) *DeletePermissionResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *DeletePermissionResponseBodyData) SetType(v string) *DeletePermissionResponseBodyData {
	s.Type = &v
	return s
}

type DeletePermissionResponseBodyDataPolicyList struct {
	MemberList []*DeletePermissionResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	Name       *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Type       *int64                                                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetMemberList(v []*DeletePermissionResponseBodyDataPolicyListMemberList) *DeletePermissionResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetName(v string) *DeletePermissionResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyList) SetType(v int64) *DeletePermissionResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type DeletePermissionResponseBodyDataPolicyListMemberList struct {
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeletePermissionResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetId(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetNickname(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *DeletePermissionResponseBodyDataPolicyListMemberList) SetType(v string) *DeletePermissionResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type DeletePermissionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeletePermissionResponse) SetHeaders(v map[string]*string) *DeletePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeletePermissionResponse) SetBody(v *DeletePermissionResponseBody) *DeletePermissionResponse {
	s.Body = v
	return s
}

type GetPeriodListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPeriodListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListHeaders) GoString() string {
	return s.String()
}

func (s *GetPeriodListHeaders) SetCommonHeaders(v map[string]*string) *GetPeriodListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPeriodListHeaders) SetXAcsDingtalkAccessToken(v string) *GetPeriodListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPeriodListResponseBody struct {
	// data
	Data *GetPeriodListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPeriodListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBody) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBody) SetData(v *GetPeriodListResponseBodyData) *GetPeriodListResponseBody {
	s.Data = v
	return s
}

func (s *GetPeriodListResponseBody) SetSuccess(v bool) *GetPeriodListResponseBody {
	s.Success = &v
	return s
}

type GetPeriodListResponseBodyData struct {
	PeriodList []*GetPeriodListResponseBodyDataPeriodList `json:"periodList,omitempty" xml:"periodList,omitempty" type:"Repeated"`
}

func (s GetPeriodListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBodyData) SetPeriodList(v []*GetPeriodListResponseBodyDataPeriodList) *GetPeriodListResponseBodyData {
	s.PeriodList = v
	return s
}

type GetPeriodListResponseBodyDataPeriodList struct {
	EndTime   *float32  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Id        io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	IsCurrent *bool     `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	IsYearly  *bool     `json:"isYearly,omitempty" xml:"isYearly,omitempty"`
	Name      io.Reader `json:"name,omitempty" xml:"name,omitempty"`
	StartTime *float32  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetPeriodListResponseBodyDataPeriodList) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponseBodyDataPeriodList) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetEndTime(v float32) *GetPeriodListResponseBodyDataPeriodList {
	s.EndTime = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetId(v io.Reader) *GetPeriodListResponseBodyDataPeriodList {
	s.Id = v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetIsCurrent(v bool) *GetPeriodListResponseBodyDataPeriodList {
	s.IsCurrent = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetIsYearly(v bool) *GetPeriodListResponseBodyDataPeriodList {
	s.IsYearly = &v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetName(v io.Reader) *GetPeriodListResponseBodyDataPeriodList {
	s.Name = v
	return s
}

func (s *GetPeriodListResponseBodyDataPeriodList) SetStartTime(v float32) *GetPeriodListResponseBodyDataPeriodList {
	s.StartTime = &v
	return s
}

type GetPeriodListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPeriodListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPeriodListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPeriodListResponse) GoString() string {
	return s.String()
}

func (s *GetPeriodListResponse) SetHeaders(v map[string]*string) *GetPeriodListResponse {
	s.Headers = v
	return s
}

func (s *GetPeriodListResponse) SetBody(v *GetPeriodListResponseBody) *GetPeriodListResponse {
	s.Body = v
	return s
}

type GetPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionHeaders) GoString() string {
	return s.String()
}

func (s *GetPermissionHeaders) SetCommonHeaders(v map[string]*string) *GetPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *GetPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPermissionRequest struct {
	TargetId   *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// A short description of struct
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	WithKr        *bool   `json:"withKr,omitempty" xml:"withKr,omitempty"`
	WithObjective *bool   `json:"withObjective,omitempty" xml:"withObjective,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) SetTargetId(v string) *GetPermissionRequest {
	s.TargetId = &v
	return s
}

func (s *GetPermissionRequest) SetTargetType(v string) *GetPermissionRequest {
	s.TargetType = &v
	return s
}

func (s *GetPermissionRequest) SetUserId(v string) *GetPermissionRequest {
	s.UserId = &v
	return s
}

func (s *GetPermissionRequest) SetWithKr(v bool) *GetPermissionRequest {
	s.WithKr = &v
	return s
}

func (s *GetPermissionRequest) SetWithObjective(v bool) *GetPermissionRequest {
	s.WithObjective = &v
	return s
}

type GetPermissionResponseBody struct {
	// 返回的数据。
	Data *GetPermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) SetData(v *GetPermissionResponseBodyData) *GetPermissionResponseBody {
	s.Data = v
	return s
}

func (s *GetPermissionResponseBody) SetSuccess(v bool) *GetPermissionResponseBody {
	s.Success = &v
	return s
}

type GetPermissionResponseBodyData struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限列表
	PolicyList []*GetPermissionResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// 是否可见的标识。
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 哪种类型的权限。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyData) SetId(v string) *GetPermissionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPermissionResponseBodyData) SetPolicyList(v []*GetPermissionResponseBodyDataPolicyList) *GetPermissionResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *GetPermissionResponseBodyData) SetPrivacy(v string) *GetPermissionResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *GetPermissionResponseBodyData) SetType(v string) *GetPermissionResponseBodyData {
	s.Type = &v
	return s
}

type GetPermissionResponseBodyDataPolicyList struct {
	MemberList []*GetPermissionResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	Name       *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	Type       *int64                                               `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyDataPolicyList) SetMemberList(v []*GetPermissionResponseBodyDataPolicyListMemberList) *GetPermissionResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyList) SetName(v string) *GetPermissionResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyList) SetType(v int64) *GetPermissionResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type GetPermissionResponseBodyDataPolicyListMemberList struct {
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPermissionResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetId(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetNickname(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *GetPermissionResponseBodyDataPolicyListMemberList) SetType(v string) *GetPermissionResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type GetPermissionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionResponse) SetHeaders(v map[string]*string) *GetPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

type GetUserOkrHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserOkrHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrHeaders) GoString() string {
	return s.String()
}

func (s *GetUserOkrHeaders) SetCommonHeaders(v map[string]*string) *GetUserOkrHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserOkrHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserOkrHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserOkrRequest struct {
	// 页码，默认 为 1。
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页的个数，默认100。
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 周期 ID。
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 当前用户的user ID。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserOkrRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrRequest) GoString() string {
	return s.String()
}

func (s *GetUserOkrRequest) SetPageNumber(v int64) *GetUserOkrRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserOkrRequest) SetPageSize(v int64) *GetUserOkrRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserOkrRequest) SetPeriodId(v string) *GetUserOkrRequest {
	s.PeriodId = &v
	return s
}

func (s *GetUserOkrRequest) SetUserId(v string) *GetUserOkrRequest {
	s.UserId = &v
	return s
}

type GetUserOkrResponseBody struct {
	// data
	Data *GetUserOkrResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserOkrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBody) SetData(v *GetUserOkrResponseBodyData) *GetUserOkrResponseBody {
	s.Data = v
	return s
}

func (s *GetUserOkrResponseBody) SetSuccess(v bool) *GetUserOkrResponseBody {
	s.Success = &v
	return s
}

type GetUserOkrResponseBodyData struct {
	// OKR 列表详情。
	List []*GetUserOkrResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 当前页码。
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每一页的个数。
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总数。
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserOkrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyData) SetList(v []*GetUserOkrResponseBodyDataList) *GetUserOkrResponseBodyData {
	s.List = v
	return s
}

func (s *GetUserOkrResponseBodyData) SetPageNumber(v int64) *GetUserOkrResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetUserOkrResponseBodyData) SetPageSize(v int64) *GetUserOkrResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetUserOkrResponseBodyData) SetTotalCount(v int64) *GetUserOkrResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetUserOkrResponseBodyDataList struct {
	// 被对齐的 Objective。
	AlignFromIds []io.Reader `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	// 对齐的 Objective。
	AlignToIds []io.Reader `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	// Objective 内容。
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间。时间戳
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 更新时间。时间戳
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// objective。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// KR 详情列表。
	KrList []*GetUserOkrResponseBodyDataListKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	// 所属者信息。
	Owner *GetUserOkrResponseBodyDataListOwner `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	// 周期 ID。
	PeriodId io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 权限值。
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// 所在位置。
	Position *int32 `json:"position,omitempty" xml:"position,omitempty"`
	// 进度值。
	Progress *GetUserOkrResponseBodyDataListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// 百分比值。
	ProgressPercent *float32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// 是否已发布。
	Published *bool `json:"published,omitempty" xml:"published,omitempty"`
	// 分数值。
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// 当前内容状态。
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 用户 ID。
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
	// 权重值。
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetUserOkrResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataList) SetAlignFromIds(v []io.Reader) *GetUserOkrResponseBodyDataList {
	s.AlignFromIds = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetAlignToIds(v []io.Reader) *GetUserOkrResponseBodyDataList {
	s.AlignToIds = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetContent(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.Content = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetGmtCreate(v float32) *GetUserOkrResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetGmtModified(v float32) *GetUserOkrResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetKrList(v []*GetUserOkrResponseBodyDataListKrList) *GetUserOkrResponseBodyDataList {
	s.KrList = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetOwner(v *GetUserOkrResponseBodyDataListOwner) *GetUserOkrResponseBodyDataList {
	s.Owner = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPeriodId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.PeriodId = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPermission(v []*float32) *GetUserOkrResponseBodyDataList {
	s.Permission = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPosition(v int32) *GetUserOkrResponseBodyDataList {
	s.Position = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetProgress(v *GetUserOkrResponseBodyDataListProgress) *GetUserOkrResponseBodyDataList {
	s.Progress = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetProgressPercent(v float32) *GetUserOkrResponseBodyDataList {
	s.ProgressPercent = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetPublished(v bool) *GetUserOkrResponseBodyDataList {
	s.Published = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetScore(v float32) *GetUserOkrResponseBodyDataList {
	s.Score = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetStatus(v int32) *GetUserOkrResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetUserId(v io.Reader) *GetUserOkrResponseBodyDataList {
	s.UserId = v
	return s
}

func (s *GetUserOkrResponseBodyDataList) SetWeight(v float32) *GetUserOkrResponseBodyDataList {
	s.Weight = &v
	return s
}

type GetUserOkrResponseBodyDataListKrList struct {
	// KR 内容。
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间。时间戳
	GmtCreate *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 更新时间。时间戳
	GmtModified *float32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// KR 的 ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// 所属 Objective ID。
	ObjectiveId io.Reader `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	// KR 权限。
	Permission []*float32 `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	// 所处位置。
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// KR 进度。
	Progress *GetUserOkrResponseBodyDataListKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// 所占分数。
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// 所占权重。
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetUserOkrResponseBodyDataListKrList) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListKrList) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListKrList) SetContent(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.Content = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetGmtCreate(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.GmtCreate = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetGmtModified(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.GmtModified = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetId(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetObjectiveId(v io.Reader) *GetUserOkrResponseBodyDataListKrList {
	s.ObjectiveId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetPermission(v []*float32) *GetUserOkrResponseBodyDataListKrList {
	s.Permission = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetPosition(v int64) *GetUserOkrResponseBodyDataListKrList {
	s.Position = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetProgress(v *GetUserOkrResponseBodyDataListKrListProgress) *GetUserOkrResponseBodyDataListKrList {
	s.Progress = v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetScore(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.Score = &v
	return s
}

func (s *GetUserOkrResponseBodyDataListKrList) SetWeight(v float32) *GetUserOkrResponseBodyDataListKrList {
	s.Weight = &v
	return s
}

type GetUserOkrResponseBodyDataListKrListProgress struct {
	// 百分比。
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetUserOkrResponseBodyDataListKrListProgress) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListKrListProgress) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListKrListProgress) SetPercent(v int32) *GetUserOkrResponseBodyDataListKrListProgress {
	s.Percent = &v
	return s
}

type GetUserOkrResponseBodyDataListOwner struct {
	// 所属者头像。 ID
	AvatarMediaId io.Reader `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	// 所属者组织 I。D
	CorpId io.Reader `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 所属者 OKR 系统中的 ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// 所属者昵称。
	Nickname io.Reader `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 所属者 userId。
	UserId io.Reader `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserOkrResponseBodyDataListOwner) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListOwner) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListOwner) SetAvatarMediaId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.AvatarMediaId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetCorpId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.CorpId = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetNickname(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.Nickname = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwner) SetUserId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.UserId = v
	return s
}

type GetUserOkrResponseBodyDataListProgress struct {
	// 百分比。
	Percent *int32 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetUserOkrResponseBodyDataListProgress) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListProgress) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListProgress) SetPercent(v int32) *GetUserOkrResponseBodyDataListProgress {
	s.Percent = &v
	return s
}

type GetUserOkrResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserOkrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserOkrResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponse) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponse) SetHeaders(v map[string]*string) *GetUserOkrResponse {
	s.Headers = v
	return s
}

func (s *GetUserOkrResponse) SetBody(v *GetUserOkrResponseBody) *GetUserOkrResponse {
	s.Body = v
	return s
}

type UnAlignObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UnAlignObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveHeaders) SetCommonHeaders(v map[string]*string) *UnAlignObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnAlignObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *UnAlignObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UnAlignObjectiveRequest struct {
	// 周期 ID
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 对齐目标的 ID
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 当前用户的 userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UnAlignObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveRequest) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveRequest) SetPeriodId(v string) *UnAlignObjectiveRequest {
	s.PeriodId = &v
	return s
}

func (s *UnAlignObjectiveRequest) SetTargetId(v string) *UnAlignObjectiveRequest {
	s.TargetId = &v
	return s
}

func (s *UnAlignObjectiveRequest) SetUserId(v string) *UnAlignObjectiveRequest {
	s.UserId = &v
	return s
}

type UnAlignObjectiveResponseBody struct {
	// data
	Data *UnAlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnAlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponseBody) SetData(v *UnAlignObjectiveResponseBodyData) *UnAlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UnAlignObjectiveResponseBody) SetSuccess(v bool) *UnAlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type UnAlignObjectiveResponseBodyData struct {
	// 对齐的 Objective ID。
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	// 当前 Objective ID。
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UnAlignObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponseBodyData) SetAlignId(v io.Reader) *UnAlignObjectiveResponseBodyData {
	s.AlignId = v
	return s
}

func (s *UnAlignObjectiveResponseBodyData) SetId(v io.Reader) *UnAlignObjectiveResponseBodyData {
	s.Id = v
	return s
}

type UnAlignObjectiveResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnAlignObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnAlignObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponse) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponse) SetHeaders(v map[string]*string) *UnAlignObjectiveResponse {
	s.Headers = v
	return s
}

func (s *UnAlignObjectiveResponse) SetBody(v *UnAlignObjectiveResponseBody) *UnAlignObjectiveResponse {
	s.Body = v
	return s
}

type UpdateKROfContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfContentHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfContentHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfContentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfContentRequest struct {
	// KR的内容。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 待更新的划词 ID 列表。
	UpdateQuoteList []*string `json:"updateQuoteList,omitempty" xml:"updateQuoteList,omitempty" type:"Repeated"`
	// 当前 KR ID。
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentRequest) SetContent(v string) *UpdateKROfContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateKROfContentRequest) SetUpdateQuoteList(v []*string) *UpdateKROfContentRequest {
	s.UpdateQuoteList = v
	return s
}

func (s *UpdateKROfContentRequest) SetKrId(v string) *UpdateKROfContentRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfContentRequest) SetUserId(v string) *UpdateKROfContentRequest {
	s.UserId = &v
	return s
}

type UpdateKROfContentResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentResponseBody) SetData(v bool) *UpdateKROfContentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateKROfContentResponseBody) SetSuccess(v bool) *UpdateKROfContentResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfContentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKROfContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKROfContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentResponse) SetHeaders(v map[string]*string) *UpdateKROfContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfContentResponse) SetBody(v *UpdateKROfContentResponseBody) *UpdateKROfContentResponse {
	s.Body = v
	return s
}

type UpdateKROfScoreHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfScoreHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfScoreHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfScoreHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfScoreHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfScoreRequest struct {
	// 分数值。
	Score *int64 `json:"score,omitempty" xml:"score,omitempty"`
	// 当前 KR ID。
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreRequest) SetScore(v int64) *UpdateKROfScoreRequest {
	s.Score = &v
	return s
}

func (s *UpdateKROfScoreRequest) SetKrId(v string) *UpdateKROfScoreRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfScoreRequest) SetUserId(v string) *UpdateKROfScoreRequest {
	s.UserId = &v
	return s
}

type UpdateKROfScoreResponseBody struct {
	Data *UpdateKROfScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponseBody) SetData(v *UpdateKROfScoreResponseBodyData) *UpdateKROfScoreResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKROfScoreResponseBody) SetSuccess(v bool) *UpdateKROfScoreResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfScoreResponseBodyData struct {
	// 目标分数。
	ObjectiveScore *int64 `json:"objectiveScore,omitempty" xml:"objectiveScore,omitempty"`
}

func (s UpdateKROfScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponseBodyData) SetObjectiveScore(v int64) *UpdateKROfScoreResponseBodyData {
	s.ObjectiveScore = &v
	return s
}

type UpdateKROfScoreResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKROfScoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKROfScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponse) SetHeaders(v map[string]*string) *UpdateKROfScoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfScoreResponse) SetBody(v *UpdateKROfScoreResponseBody) *UpdateKROfScoreResponse {
	s.Body = v
	return s
}

type UpdateKROfWeightHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateKROfWeightHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightHeaders) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightHeaders) SetCommonHeaders(v map[string]*string) *UpdateKROfWeightHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateKROfWeightHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateKROfWeightHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateKROfWeightRequest struct {
	// 权重比。
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 当前 KR ID。
	KrId *string `json:"krId,omitempty" xml:"krId,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateKROfWeightRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightRequest) SetWeight(v int64) *UpdateKROfWeightRequest {
	s.Weight = &v
	return s
}

func (s *UpdateKROfWeightRequest) SetKrId(v string) *UpdateKROfWeightRequest {
	s.KrId = &v
	return s
}

func (s *UpdateKROfWeightRequest) SetUserId(v string) *UpdateKROfWeightRequest {
	s.UserId = &v
	return s
}

type UpdateKROfWeightResponseBody struct {
	Data *UpdateKROfWeightResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfWeightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBody) SetData(v *UpdateKROfWeightResponseBodyData) *UpdateKROfWeightResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKROfWeightResponseBody) SetSuccess(v bool) *UpdateKROfWeightResponseBody {
	s.Success = &v
	return s
}

type UpdateKROfWeightResponseBodyData struct {
	ObjectiveProgress *UpdateKROfWeightResponseBodyDataObjectiveProgress `json:"objectiveProgress,omitempty" xml:"objectiveProgress,omitempty" type:"Struct"`
	// 目标分数。
	ObjectiveScore *int64 `json:"objectiveScore,omitempty" xml:"objectiveScore,omitempty"`
}

func (s UpdateKROfWeightResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBodyData) SetObjectiveProgress(v *UpdateKROfWeightResponseBodyDataObjectiveProgress) *UpdateKROfWeightResponseBodyData {
	s.ObjectiveProgress = v
	return s
}

func (s *UpdateKROfWeightResponseBodyData) SetObjectiveScore(v int64) *UpdateKROfWeightResponseBodyData {
	s.ObjectiveScore = &v
	return s
}

type UpdateKROfWeightResponseBodyDataObjectiveProgress struct {
	// 目标百分比。
	Percent *int64 `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s UpdateKROfWeightResponseBodyDataObjectiveProgress) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponseBodyDataObjectiveProgress) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponseBodyDataObjectiveProgress) SetPercent(v int64) *UpdateKROfWeightResponseBodyDataObjectiveProgress {
	s.Percent = &v
	return s
}

type UpdateKROfWeightResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateKROfWeightResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateKROfWeightResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfWeightResponse) GoString() string {
	return s.String()
}

func (s *UpdateKROfWeightResponse) SetHeaders(v map[string]*string) *UpdateKROfWeightResponse {
	s.Headers = v
	return s
}

func (s *UpdateKROfWeightResponse) SetBody(v *UpdateKROfWeightResponseBody) *UpdateKROfWeightResponse {
	s.Body = v
	return s
}

type UpdateObjectiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateObjectiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveHeaders) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveHeaders) SetCommonHeaders(v map[string]*string) *UpdateObjectiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateObjectiveHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateObjectiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateObjectiveRequest struct {
	// 当前 Objective 的内容。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 当前用户的 userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveRequest) SetContent(v string) *UpdateObjectiveRequest {
	s.Content = &v
	return s
}

func (s *UpdateObjectiveRequest) SetUserId(v string) *UpdateObjectiveRequest {
	s.UserId = &v
	return s
}

type UpdateObjectiveResponseBody struct {
	// data
	Data *UpdateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponseBody) SetData(v *UpdateObjectiveResponseBodyData) *UpdateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UpdateObjectiveResponseBody) SetSuccess(v bool) *UpdateObjectiveResponseBody {
	s.Success = &v
	return s
}

type UpdateObjectiveResponseBodyData struct {
	// 当前 Objective ID。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 当前 Objective 的位置。
	Position *float32 `json:"position,omitempty" xml:"position,omitempty"`
}

func (s UpdateObjectiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponseBodyData) SetId(v string) *UpdateObjectiveResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateObjectiveResponseBodyData) SetPosition(v float32) *UpdateObjectiveResponseBodyData {
	s.Position = &v
	return s
}

type UpdateObjectiveResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateObjectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateObjectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponse) SetHeaders(v map[string]*string) *UpdateObjectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateObjectiveResponse) SetBody(v *UpdateObjectiveResponseBody) *UpdateObjectiveResponse {
	s.Body = v
	return s
}

type UpdatePrivacyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePrivacyHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyHeaders) SetCommonHeaders(v map[string]*string) *UpdatePrivacyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePrivacyHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePrivacyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePrivacyRequest struct {
	// 权限修改的类型
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 当前目标的 ID
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 当前目标的权限类型。
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// 当前用户的userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdatePrivacyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyRequest) SetPrivacy(v string) *UpdatePrivacyRequest {
	s.Privacy = &v
	return s
}

func (s *UpdatePrivacyRequest) SetTargetId(v string) *UpdatePrivacyRequest {
	s.TargetId = &v
	return s
}

func (s *UpdatePrivacyRequest) SetTargetType(v string) *UpdatePrivacyRequest {
	s.TargetType = &v
	return s
}

func (s *UpdatePrivacyRequest) SetUserId(v string) *UpdatePrivacyRequest {
	s.UserId = &v
	return s
}

type UpdatePrivacyResponseBody struct {
	// 返回的数据。
	Data *UpdatePrivacyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 请求成功的标识。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdatePrivacyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBody) SetData(v *UpdatePrivacyResponseBodyData) *UpdatePrivacyResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePrivacyResponseBody) SetSuccess(v bool) *UpdatePrivacyResponseBody {
	s.Success = &v
	return s
}

type UpdatePrivacyResponseBodyData struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 权限列表
	PolicyList []*UpdatePrivacyResponseBodyDataPolicyList `json:"policyList,omitempty" xml:"policyList,omitempty" type:"Repeated"`
	// 是否可见的标识。
	Privacy *string `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// 哪种类型的权限。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyData) SetId(v string) *UpdatePrivacyResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetPolicyList(v []*UpdatePrivacyResponseBodyDataPolicyList) *UpdatePrivacyResponseBodyData {
	s.PolicyList = v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetPrivacy(v string) *UpdatePrivacyResponseBodyData {
	s.Privacy = &v
	return s
}

func (s *UpdatePrivacyResponseBodyData) SetType(v string) *UpdatePrivacyResponseBodyData {
	s.Type = &v
	return s
}

type UpdatePrivacyResponseBodyDataPolicyList struct {
	MemberList []*UpdatePrivacyResponseBodyDataPolicyListMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	Name       *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	Type       *int64                                               `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyDataPolicyList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyDataPolicyList) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetMemberList(v []*UpdatePrivacyResponseBodyDataPolicyListMemberList) *UpdatePrivacyResponseBodyDataPolicyList {
	s.MemberList = v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetName(v string) *UpdatePrivacyResponseBodyDataPolicyList {
	s.Name = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyList) SetType(v int64) *UpdatePrivacyResponseBodyDataPolicyList {
	s.Type = &v
	return s
}

type UpdatePrivacyResponseBodyDataPolicyListMemberList struct {
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdatePrivacyResponseBodyDataPolicyListMemberList) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponseBodyDataPolicyListMemberList) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetId(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Id = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetNickname(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Nickname = &v
	return s
}

func (s *UpdatePrivacyResponseBodyDataPolicyListMemberList) SetType(v string) *UpdatePrivacyResponseBodyDataPolicyListMemberList {
	s.Type = &v
	return s
}

type UpdatePrivacyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePrivacyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePrivacyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivacyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyResponse) SetHeaders(v map[string]*string) *UpdatePrivacyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivacyResponse) SetBody(v *UpdatePrivacyResponseBody) *UpdatePrivacyResponse {
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

func (client *Client) AlignObjective(objectiveId *string, request *AlignObjectiveRequest) (_result *AlignObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AlignObjectiveHeaders{}
	_result = &AlignObjectiveResponse{}
	_body, _err := client.AlignObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AlignObjectiveWithOptions(objectiveId *string, request *AlignObjectiveRequest, headers *AlignObjectiveHeaders, runtime *util.RuntimeOptions) (_result *AlignObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	objectiveId = openapiutil.GetEncodeParam(objectiveId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

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
	_result = &AlignObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("AlignObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/objectives/"+tea.StringValue(objectiveId)+"/alignments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddPermission(request *BatchAddPermissionRequest) (_result *BatchAddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddPermissionHeaders{}
	_result = &BatchAddPermissionResponse{}
	_body, _err := client.BatchAddPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddPermissionWithOptions(request *BatchAddPermissionRequest, headers *BatchAddPermissionHeaders, runtime *util.RuntimeOptions) (_result *BatchAddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.List)) {
		body["list"] = request.List
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

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
	_result = &BatchAddPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchAddPermission"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/permissions/batch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQueryObjective(request *BatchQueryObjectiveRequest) (_result *BatchQueryObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryObjectiveHeaders{}
	_result = &BatchQueryObjectiveResponse{}
	_body, _err := client.BatchQueryObjectiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryObjectiveWithOptions(request *BatchQueryObjectiveRequest, headers *BatchQueryObjectiveHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectiveIds)) {
		body["objectiveIds"] = request.ObjectiveIds
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.WithAlign)) {
		body["withAlign"] = request.WithAlign
	}

	if !tea.BoolValue(util.IsUnset(request.WithKr)) {
		body["withKr"] = request.WithKr
	}

	if !tea.BoolValue(util.IsUnset(request.WithProgress)) {
		body["withProgress"] = request.WithProgress
	}

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
	_result = &BatchQueryObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchQueryObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/objectives/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQueryUser(request *BatchQueryUserRequest) (_result *BatchQueryUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryUserHeaders{}
	_result = &BatchQueryUserResponse{}
	_body, _err := client.BatchQueryUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQueryUserWithOptions(request *BatchQueryUserRequest, headers *BatchQueryUserHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OkrUserIds)) {
		body["okrUserIds"] = request.OkrUserIds
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
	_result = &BatchQueryUserResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchQueryUser"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/users/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeyResult(request *CreateKeyResultRequest) (_result *CreateKeyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateKeyResultHeaders{}
	_result = &CreateKeyResultResponse{}
	_body, _err := client.CreateKeyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeyResultWithOptions(request *CreateKeyResultRequest, headers *CreateKeyResultHeaders, runtime *util.RuntimeOptions) (_result *CreateKeyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectiveId)) {
		body["objectiveId"] = request.ObjectiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.PrevPosition)) {
		body["prevPosition"] = request.PrevPosition
	}

	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["weight"] = request.Weight
	}

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
	_result = &CreateKeyResultResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateKeyResult"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/keyResults"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateObjective(request *CreateObjectiveRequest) (_result *CreateObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateObjectiveHeaders{}
	_result = &CreateObjectiveResponse{}
	_body, _err := client.CreateObjectiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateObjectiveWithOptions(request *CreateObjectiveRequest, headers *CreateObjectiveHeaders, runtime *util.RuntimeOptions) (_result *CreateObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.PrevPosition)) {
		body["prevPosition"] = request.PrevPosition
	}

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
	_result = &CreateObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/objectives"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyResult(request *DeleteKeyResultRequest) (_result *DeleteKeyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteKeyResultHeaders{}
	_result = &DeleteKeyResultResponse{}
	_body, _err := client.DeleteKeyResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeyResultWithOptions(request *DeleteKeyResultRequest, headers *DeleteKeyResultHeaders, runtime *util.RuntimeOptions) (_result *DeleteKeyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
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
	_result = &DeleteKeyResultResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteKeyResult"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/okr/keyResults"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteObjective(objectiveId *string, request *DeleteObjectiveRequest) (_result *DeleteObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteObjectiveHeaders{}
	_result = &DeleteObjectiveResponse{}
	_body, _err := client.DeleteObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteObjectiveWithOptions(objectiveId *string, request *DeleteObjectiveRequest, headers *DeleteObjectiveHeaders, runtime *util.RuntimeOptions) (_result *DeleteObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	objectiveId = openapiutil.GetEncodeParam(objectiveId)
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
	_result = &DeleteObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/okr/objectives/"+tea.StringValue(objectiveId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePermission(request *DeletePermissionRequest) (_result *DeletePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeletePermissionHeaders{}
	_result = &DeletePermissionResponse{}
	_body, _err := client.DeletePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePermissionWithOptions(request *DeletePermissionRequest, headers *DeletePermissionHeaders, runtime *util.RuntimeOptions) (_result *DeletePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["policyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["targetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
	_result = &DeletePermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePermission"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/okr/permissions/delete"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPeriodList() (_result *GetPeriodListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPeriodListHeaders{}
	_result = &GetPeriodListResponse{}
	_body, _err := client.GetPeriodListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPeriodListWithOptions(headers *GetPeriodListHeaders, runtime *util.RuntimeOptions) (_result *GetPeriodListResponse, _err error) {
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
	_result = &GetPeriodListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPeriodList"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/okr/periods"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPermission(request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPermissionHeaders{}
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPermissionWithOptions(request *GetPermissionRequest, headers *GetPermissionHeaders, runtime *util.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["targetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WithKr)) {
		query["withKr"] = request.WithKr
	}

	if !tea.BoolValue(util.IsUnset(request.WithObjective)) {
		query["withObjective"] = request.WithObjective
	}

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
	_result = &GetPermissionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPermission"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/okr/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserOkr(request *GetUserOkrRequest) (_result *GetUserOkrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserOkrHeaders{}
	_result = &GetUserOkrResponse{}
	_body, _err := client.GetUserOkrWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserOkrWithOptions(request *GetUserOkrRequest, headers *GetUserOkrHeaders, runtime *util.RuntimeOptions) (_result *GetUserOkrResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		query["periodId"] = request.PeriodId
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
	_result = &GetUserOkrResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserOkr"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/okr/users/okrs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnAlignObjective(objectiveId *string, request *UnAlignObjectiveRequest) (_result *UnAlignObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UnAlignObjectiveHeaders{}
	_result = &UnAlignObjectiveResponse{}
	_body, _err := client.UnAlignObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnAlignObjectiveWithOptions(objectiveId *string, request *UnAlignObjectiveRequest, headers *UnAlignObjectiveHeaders, runtime *util.RuntimeOptions) (_result *UnAlignObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	objectiveId = openapiutil.GetEncodeParam(objectiveId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		body["periodId"] = request.PeriodId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

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
	_result = &UnAlignObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("UnAlignObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/okr/objectives/"+tea.StringValue(objectiveId)+"/alignments/cancel"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKROfContent(request *UpdateKROfContentRequest) (_result *UpdateKROfContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfContentHeaders{}
	_result = &UpdateKROfContentResponse{}
	_body, _err := client.UpdateKROfContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKROfContentWithOptions(request *UpdateKROfContentRequest, headers *UpdateKROfContentHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateQuoteList)) {
		body["updateQuoteList"] = request.UpdateQuoteList
	}

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
	_result = &UpdateKROfContentResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateKROfContent"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/okr/keyResults/contents"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKROfScore(request *UpdateKROfScoreRequest) (_result *UpdateKROfScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfScoreHeaders{}
	_result = &UpdateKROfScoreResponse{}
	_body, _err := client.UpdateKROfScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKROfScoreWithOptions(request *UpdateKROfScoreRequest, headers *UpdateKROfScoreHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Score)) {
		body["score"] = request.Score
	}

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
	_result = &UpdateKROfScoreResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateKROfScore"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/okr/keyResults/scores"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKROfWeight(request *UpdateKROfWeightRequest) (_result *UpdateKROfWeightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateKROfWeightHeaders{}
	_result = &UpdateKROfWeightResponse{}
	_body, _err := client.UpdateKROfWeightWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKROfWeightWithOptions(request *UpdateKROfWeightRequest, headers *UpdateKROfWeightHeaders, runtime *util.RuntimeOptions) (_result *UpdateKROfWeightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KrId)) {
		query["krId"] = request.KrId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["weight"] = request.Weight
	}

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
	_result = &UpdateKROfWeightResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateKROfWeight"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/okr/keyResults/weights"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateObjective(objectiveId *string, request *UpdateObjectiveRequest) (_result *UpdateObjectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateObjectiveHeaders{}
	_result = &UpdateObjectiveResponse{}
	_body, _err := client.UpdateObjectiveWithOptions(objectiveId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateObjectiveWithOptions(objectiveId *string, request *UpdateObjectiveRequest, headers *UpdateObjectiveHeaders, runtime *util.RuntimeOptions) (_result *UpdateObjectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	objectiveId = openapiutil.GetEncodeParam(objectiveId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

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
	_result = &UpdateObjectiveResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateObjective"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/okr/objectives/"+tea.StringValue(objectiveId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrivacy(request *UpdatePrivacyRequest) (_result *UpdatePrivacyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePrivacyHeaders{}
	_result = &UpdatePrivacyResponse{}
	_body, _err := client.UpdatePrivacyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrivacyWithOptions(request *UpdatePrivacyRequest, headers *UpdatePrivacyHeaders, runtime *util.RuntimeOptions) (_result *UpdatePrivacyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Privacy)) {
		body["privacy"] = request.Privacy
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		body["targetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		body["targetType"] = request.TargetType
	}

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
	_result = &UpdatePrivacyResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePrivacy"), tea.String("okr_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/okr/permissions/privacies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
