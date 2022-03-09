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
	// 周期 ID
	PeriodId *string `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// 对齐目标的 ID
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// 用户 ID
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
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

func (s *AlignObjectiveRequest) SetOwnerId(v io.Reader) *AlignObjectiveRequest {
	s.OwnerId = v
	return s
}

type AlignObjectiveResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *AlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *AlignObjectiveResponseBody) SetCode(v int64) *AlignObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *AlignObjectiveResponseBody) SetData(v *AlignObjectiveResponseBodyData) *AlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *AlignObjectiveResponseBody) SetMessage(v string) *AlignObjectiveResponseBody {
	s.Message = &v
	return s
}

func (s *AlignObjectiveResponseBody) SetSuccess(v bool) *AlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type AlignObjectiveResponseBodyData struct {
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	Id      io.Reader `json:"id,omitempty" xml:"id,omitempty"`
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
	ObjectiveIds []io.Reader `json:"objectiveIds,omitempty" xml:"objectiveIds,omitempty" type:"Repeated"`
	// periodId
	PeriodId io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// withAlign
	WithAlign *bool `json:"withAlign,omitempty" xml:"withAlign,omitempty"`
	// withKr
	WithKr *bool `json:"withKr,omitempty" xml:"withKr,omitempty"`
	// withProgress
	WithProgress *bool `json:"withProgress,omitempty" xml:"withProgress,omitempty"`
	// ownerId
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
}

func (s BatchQueryObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveRequest) SetObjectiveIds(v []io.Reader) *BatchQueryObjectiveRequest {
	s.ObjectiveIds = v
	return s
}

func (s *BatchQueryObjectiveRequest) SetPeriodId(v io.Reader) *BatchQueryObjectiveRequest {
	s.PeriodId = v
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

func (s *BatchQueryObjectiveRequest) SetOwnerId(v string) *BatchQueryObjectiveRequest {
	s.OwnerId = &v
	return s
}

type BatchQueryObjectiveResponseBody struct {
	// code
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*BatchQueryObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// message
	Message io.Reader `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBody) SetCode(v string) *BatchQueryObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *BatchQueryObjectiveResponseBody) SetData(v []*BatchQueryObjectiveResponseBodyData) *BatchQueryObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *BatchQueryObjectiveResponseBody) SetMessage(v io.Reader) *BatchQueryObjectiveResponseBody {
	s.Message = v
	return s
}

func (s *BatchQueryObjectiveResponseBody) SetSuccess(v bool) *BatchQueryObjectiveResponseBody {
	s.Success = &v
	return s
}

type BatchQueryObjectiveResponseBodyData struct {
	AlignFromIds    []io.Reader                                  `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	AlignToIds      []io.Reader                                  `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	Content         io.Reader                                    `json:"content,omitempty" xml:"content,omitempty"`
	Id              io.Reader                                    `json:"id,omitempty" xml:"id,omitempty"`
	KrList          []*BatchQueryObjectiveResponseBodyDataKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	Owner           *BatchQueryObjectiveResponseBodyDataOwner    `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	PeriodId        io.Reader                                    `json:"periodId,omitempty" xml:"periodId,omitempty"`
	Permission      []*float32                                   `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	Position        *int32                                       `json:"position,omitempty" xml:"position,omitempty"`
	Progress        *BatchQueryObjectiveResponseBodyDataProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	ProgressPercent *float32                                     `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Published       *bool                                        `json:"published,omitempty" xml:"published,omitempty"`
	Score           *float32                                     `json:"score,omitempty" xml:"score,omitempty"`
	Status          *int32                                       `json:"status,omitempty" xml:"status,omitempty"`
	UserId          io.Reader                                    `json:"userId,omitempty" xml:"userId,omitempty"`
	Weight          *float32                                     `json:"weight,omitempty" xml:"weight,omitempty"`
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
	Content     io.Reader                                          `json:"content,omitempty" xml:"content,omitempty"`
	Id          io.Reader                                          `json:"id,omitempty" xml:"id,omitempty"`
	ObjectiveId io.Reader                                          `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Permission  []*float32                                         `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	Position    *int64                                             `json:"position,omitempty" xml:"position,omitempty"`
	Progress    *BatchQueryObjectiveResponseBodyDataKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	Score       *float32                                           `json:"score,omitempty" xml:"score,omitempty"`
	Weight      *float32                                           `json:"weight,omitempty" xml:"weight,omitempty"`
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
	AvatarMediaId io.Reader                                           `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	CorpId        io.Reader                                           `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Department    *BatchQueryObjectiveResponseBodyDataOwnerDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	Id            io.Reader                                           `json:"id,omitempty" xml:"id,omitempty"`
	Nickname      io.Reader                                           `json:"nickname,omitempty" xml:"nickname,omitempty"`
	StaffId       io.Reader                                           `json:"staffId,omitempty" xml:"staffId,omitempty"`
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

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetDepartment(v *BatchQueryObjectiveResponseBodyDataOwnerDepartment) *BatchQueryObjectiveResponseBodyDataOwner {
	s.Department = v
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

func (s *BatchQueryObjectiveResponseBodyDataOwner) SetStaffId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwner {
	s.StaffId = v
	return s
}

type BatchQueryObjectiveResponseBodyDataOwnerDepartment struct {
	Id   io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	Name io.Reader `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BatchQueryObjectiveResponseBodyDataOwnerDepartment) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryObjectiveResponseBodyDataOwnerDepartment) GoString() string {
	return s.String()
}

func (s *BatchQueryObjectiveResponseBodyDataOwnerDepartment) SetId(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwnerDepartment {
	s.Id = v
	return s
}

func (s *BatchQueryObjectiveResponseBodyDataOwnerDepartment) SetName(v io.Reader) *BatchQueryObjectiveResponseBodyDataOwnerDepartment {
	s.Name = v
	return s
}

type BatchQueryObjectiveResponseBodyDataProgress struct {
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
	Content      io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	ObjectiveId  io.Reader `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	PeriodId     io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
	PrevPosition *int64    `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	Weight       *int64    `json:"weight,omitempty" xml:"weight,omitempty"`
	OwnerId      io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
}

func (s CreateKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResultRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyResultRequest) SetContent(v io.Reader) *CreateKeyResultRequest {
	s.Content = v
	return s
}

func (s *CreateKeyResultRequest) SetObjectiveId(v io.Reader) *CreateKeyResultRequest {
	s.ObjectiveId = v
	return s
}

func (s *CreateKeyResultRequest) SetPeriodId(v io.Reader) *CreateKeyResultRequest {
	s.PeriodId = v
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

func (s *CreateKeyResultRequest) SetOwnerId(v io.Reader) *CreateKeyResultRequest {
	s.OwnerId = v
	return s
}

type CreateKeyResultResponseBody struct {
	Data *CreateKeyResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
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
	Id       io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	Position *int64    `json:"position,omitempty" xml:"position,omitempty"`
	Weight   *int64    `json:"weight,omitempty" xml:"weight,omitempty"`
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
	// content
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// periodId
	PeriodId io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
	// prevPosition
	PrevPosition io.Reader `json:"prevPosition,omitempty" xml:"prevPosition,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *CreateObjectiveRequest) SetContent(v io.Reader) *CreateObjectiveRequest {
	s.Content = v
	return s
}

func (s *CreateObjectiveRequest) SetPeriodId(v io.Reader) *CreateObjectiveRequest {
	s.PeriodId = v
	return s
}

func (s *CreateObjectiveRequest) SetPrevPosition(v io.Reader) *CreateObjectiveRequest {
	s.PrevPosition = v
	return s
}

func (s *CreateObjectiveRequest) SetUserId(v string) *CreateObjectiveRequest {
	s.UserId = &v
	return s
}

type CreateObjectiveResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *CreateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateObjectiveResponseBody) SetCode(v int64) *CreateObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *CreateObjectiveResponseBody) SetData(v *CreateObjectiveResponseBodyData) *CreateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *CreateObjectiveResponseBody) SetMessage(v string) *CreateObjectiveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateObjectiveResponseBody) SetSuccess(v bool) *CreateObjectiveResponseBody {
	s.Success = &v
	return s
}

type CreateObjectiveResponseBodyData struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
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
	KrId    io.Reader `json:"krId,omitempty" xml:"krId,omitempty"`
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
}

func (s DeleteKeyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyResultRequest) SetKrId(v io.Reader) *DeleteKeyResultRequest {
	s.KrId = v
	return s
}

func (s *DeleteKeyResultRequest) SetOwnerId(v io.Reader) *DeleteKeyResultRequest {
	s.OwnerId = v
	return s
}

type DeleteKeyResultResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// is success
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
	// userId
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
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *DeleteObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteObjectiveResponseBody) SetCode(v int64) *DeleteObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteObjectiveResponseBody) SetData(v *DeleteObjectiveResponseBodyData) *DeleteObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *DeleteObjectiveResponseBody) SetMessage(v string) *DeleteObjectiveResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteObjectiveResponseBody) SetSuccess(v bool) *DeleteObjectiveResponseBody {
	s.Success = &v
	return s
}

type DeleteObjectiveResponseBodyData struct {
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
	// 归属用户的ID
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 页码，默认 为 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页的个数，默认100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 周期 ID
	PeriodId io.Reader `json:"periodId,omitempty" xml:"periodId,omitempty"`
}

func (s GetUserOkrRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrRequest) GoString() string {
	return s.String()
}

func (s *GetUserOkrRequest) SetOwnerId(v io.Reader) *GetUserOkrRequest {
	s.OwnerId = v
	return s
}

func (s *GetUserOkrRequest) SetPageNumber(v int64) *GetUserOkrRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserOkrRequest) SetPageSize(v int64) *GetUserOkrRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserOkrRequest) SetPeriodId(v io.Reader) *GetUserOkrRequest {
	s.PeriodId = v
	return s
}

type GetUserOkrResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *GetUserOkrResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message io.Reader `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUserOkrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBody) SetCode(v int64) *GetUserOkrResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserOkrResponseBody) SetData(v *GetUserOkrResponseBodyData) *GetUserOkrResponseBody {
	s.Data = v
	return s
}

func (s *GetUserOkrResponseBody) SetMessage(v io.Reader) *GetUserOkrResponseBody {
	s.Message = v
	return s
}

func (s *GetUserOkrResponseBody) SetSuccess(v bool) *GetUserOkrResponseBody {
	s.Success = &v
	return s
}

type GetUserOkrResponseBodyData struct {
	List       []*GetUserOkrResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNo     *int64                            `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize   *int64                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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

func (s *GetUserOkrResponseBodyData) SetPageNo(v int64) *GetUserOkrResponseBodyData {
	s.PageNo = &v
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
	AlignFromIds    []io.Reader                             `json:"alignFromIds,omitempty" xml:"alignFromIds,omitempty" type:"Repeated"`
	AlignToIds      []io.Reader                             `json:"alignToIds,omitempty" xml:"alignToIds,omitempty" type:"Repeated"`
	Content         io.Reader                               `json:"content,omitempty" xml:"content,omitempty"`
	Id              io.Reader                               `json:"id,omitempty" xml:"id,omitempty"`
	KrList          []*GetUserOkrResponseBodyDataListKrList `json:"krList,omitempty" xml:"krList,omitempty" type:"Repeated"`
	Owner           *GetUserOkrResponseBodyDataListOwner    `json:"owner,omitempty" xml:"owner,omitempty" type:"Struct"`
	PeriodId        io.Reader                               `json:"periodId,omitempty" xml:"periodId,omitempty"`
	Permission      []*float32                              `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	Position        *int32                                  `json:"position,omitempty" xml:"position,omitempty"`
	Progress        *GetUserOkrResponseBodyDataListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	ProgressPercent *float32                                `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Published       *bool                                   `json:"published,omitempty" xml:"published,omitempty"`
	Score           *float32                                `json:"score,omitempty" xml:"score,omitempty"`
	Status          *int32                                  `json:"status,omitempty" xml:"status,omitempty"`
	UserId          io.Reader                               `json:"userId,omitempty" xml:"userId,omitempty"`
	Weight          *float32                                `json:"weight,omitempty" xml:"weight,omitempty"`
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
	Content     io.Reader                                     `json:"content,omitempty" xml:"content,omitempty"`
	Id          io.Reader                                     `json:"id,omitempty" xml:"id,omitempty"`
	ObjectiveId io.Reader                                     `json:"objectiveId,omitempty" xml:"objectiveId,omitempty"`
	Permission  []*float32                                    `json:"permission,omitempty" xml:"permission,omitempty" type:"Repeated"`
	Position    *int64                                        `json:"position,omitempty" xml:"position,omitempty"`
	Progress    *GetUserOkrResponseBodyDataListKrListProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	Score       *float32                                      `json:"score,omitempty" xml:"score,omitempty"`
	Weight      *float32                                      `json:"weight,omitempty" xml:"weight,omitempty"`
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
	AvatarMediaId io.Reader                                      `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	CorpId        io.Reader                                      `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Department    *GetUserOkrResponseBodyDataListOwnerDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	Id            io.Reader                                      `json:"id,omitempty" xml:"id,omitempty"`
	Nickname      io.Reader                                      `json:"nickname,omitempty" xml:"nickname,omitempty"`
	StaffId       io.Reader                                      `json:"staffId,omitempty" xml:"staffId,omitempty"`
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

func (s *GetUserOkrResponseBodyDataListOwner) SetDepartment(v *GetUserOkrResponseBodyDataListOwnerDepartment) *GetUserOkrResponseBodyDataListOwner {
	s.Department = v
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

func (s *GetUserOkrResponseBodyDataListOwner) SetStaffId(v io.Reader) *GetUserOkrResponseBodyDataListOwner {
	s.StaffId = v
	return s
}

type GetUserOkrResponseBodyDataListOwnerDepartment struct {
	Id   io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	Name io.Reader `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetUserOkrResponseBodyDataListOwnerDepartment) String() string {
	return tea.Prettify(s)
}

func (s GetUserOkrResponseBodyDataListOwnerDepartment) GoString() string {
	return s.String()
}

func (s *GetUserOkrResponseBodyDataListOwnerDepartment) SetId(v io.Reader) *GetUserOkrResponseBodyDataListOwnerDepartment {
	s.Id = v
	return s
}

func (s *GetUserOkrResponseBodyDataListOwnerDepartment) SetName(v io.Reader) *GetUserOkrResponseBodyDataListOwnerDepartment {
	s.Name = v
	return s
}

type GetUserOkrResponseBodyDataListProgress struct {
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
	// 用户 ID
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
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

func (s *UnAlignObjectiveRequest) SetOwnerId(v io.Reader) *UnAlignObjectiveRequest {
	s.OwnerId = v
	return s
}

type UnAlignObjectiveResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *UnAlignObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UnAlignObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnAlignObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UnAlignObjectiveResponseBody) SetCode(v int64) *UnAlignObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *UnAlignObjectiveResponseBody) SetData(v *UnAlignObjectiveResponseBodyData) *UnAlignObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UnAlignObjectiveResponseBody) SetMessage(v string) *UnAlignObjectiveResponseBody {
	s.Message = &v
	return s
}

func (s *UnAlignObjectiveResponseBody) SetSuccess(v bool) *UnAlignObjectiveResponseBody {
	s.Success = &v
	return s
}

type UnAlignObjectiveResponseBodyData struct {
	AlignId io.Reader `json:"alignId,omitempty" xml:"alignId,omitempty"`
	Id      io.Reader `json:"id,omitempty" xml:"id,omitempty"`
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
	Content         io.Reader   `json:"content,omitempty" xml:"content,omitempty"`
	UpdateQuoteList []io.Reader `json:"updateQuoteList,omitempty" xml:"updateQuoteList,omitempty" type:"Repeated"`
	// A short description of struct
	KrId        io.Reader `json:"krId,omitempty" xml:"krId,omitempty"`
	OperatorUid io.Reader `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s UpdateKROfContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateKROfContentRequest) SetContent(v io.Reader) *UpdateKROfContentRequest {
	s.Content = v
	return s
}

func (s *UpdateKROfContentRequest) SetUpdateQuoteList(v []io.Reader) *UpdateKROfContentRequest {
	s.UpdateQuoteList = v
	return s
}

func (s *UpdateKROfContentRequest) SetKrId(v io.Reader) *UpdateKROfContentRequest {
	s.KrId = v
	return s
}

func (s *UpdateKROfContentRequest) SetOperatorUid(v io.Reader) *UpdateKROfContentRequest {
	s.OperatorUid = v
	return s
}

type UpdateKROfContentResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
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
	Score *int64 `json:"score,omitempty" xml:"score,omitempty"`
	// A short description of struct
	KrId    io.Reader `json:"krId,omitempty" xml:"krId,omitempty"`
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
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

func (s *UpdateKROfScoreRequest) SetKrId(v io.Reader) *UpdateKROfScoreRequest {
	s.KrId = v
	return s
}

func (s *UpdateKROfScoreRequest) SetOwnerId(v io.Reader) *UpdateKROfScoreRequest {
	s.OwnerId = v
	return s
}

type UpdateKROfScoreResponseBody struct {
	// 目标分数
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateKROfScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKROfScoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKROfScoreResponseBody) SetData(v int64) *UpdateKROfScoreResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateKROfScoreResponseBody) SetSuccess(v bool) *UpdateKROfScoreResponseBody {
	s.Success = &v
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
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// A short description of struct
	KrId    io.Reader `json:"krId,omitempty" xml:"krId,omitempty"`
	OwnerId io.Reader `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
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

func (s *UpdateKROfWeightRequest) SetKrId(v io.Reader) *UpdateKROfWeightRequest {
	s.KrId = v
	return s
}

func (s *UpdateKROfWeightRequest) SetOwnerId(v io.Reader) *UpdateKROfWeightRequest {
	s.OwnerId = v
	return s
}

type UpdateKROfWeightResponseBody struct {
	// 目标分数
	Data *UpdateKROfWeightResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
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
	ObjectiveScore    *int64                                             `json:"objectiveScore,omitempty" xml:"objectiveScore,omitempty"`
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
	Percent *int64 `json:"percent,omitempty" xml:"percent,omitempty"`
	Status  *int64 `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *UpdateKROfWeightResponseBodyDataObjectiveProgress) SetStatus(v int64) *UpdateKROfWeightResponseBodyDataObjectiveProgress {
	s.Status = &v
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
	Content io.Reader `json:"content,omitempty" xml:"content,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateObjectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveRequest) SetContent(v io.Reader) *UpdateObjectiveRequest {
	s.Content = v
	return s
}

func (s *UpdateObjectiveRequest) SetUserId(v string) *UpdateObjectiveRequest {
	s.UserId = &v
	return s
}

type UpdateObjectiveResponseBody struct {
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *UpdateObjectiveResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateObjectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateObjectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateObjectiveResponseBody) SetCode(v int64) *UpdateObjectiveResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateObjectiveResponseBody) SetData(v *UpdateObjectiveResponseBodyData) *UpdateObjectiveResponseBody {
	s.Data = v
	return s
}

func (s *UpdateObjectiveResponseBody) SetMessage(v string) *UpdateObjectiveResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateObjectiveResponseBody) SetSuccess(v bool) *UpdateObjectiveResponseBody {
	s.Success = &v
	return s
}

type UpdateObjectiveResponseBodyData struct {
	Id       *string  `json:"id,omitempty" xml:"id,omitempty"`
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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
	}

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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodId)) {
		query["periodId"] = request.PeriodId
	}

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
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		query["operatorUid"] = request.OperatorUid
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["ownerId"] = request.OwnerId
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
