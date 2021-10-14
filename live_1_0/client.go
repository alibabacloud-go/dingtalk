// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package live_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EditFeedReplayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditFeedReplayHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditFeedReplayHeaders) GoString() string {
	return s.String()
}

func (s *EditFeedReplayHeaders) SetCommonHeaders(v map[string]*string) *EditFeedReplayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditFeedReplayHeaders) SetXAcsDingtalkAccessToken(v string) *EditFeedReplayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditFeedReplayRequest struct {
	// 用户id(剪辑者的组织内id)
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 剪辑的起始位置的时间戳（在原开始结束的时间戳之内）
	EditStartTime *int64 `json:"editStartTime,omitempty" xml:"editStartTime,omitempty"`
	// 剪辑的结束位置的时间戳（在原开始结束的时间戳之内）
	EditEndTime *int64 `json:"editEndTime,omitempty" xml:"editEndTime,omitempty"`
}

func (s EditFeedReplayRequest) String() string {
	return tea.Prettify(s)
}

func (s EditFeedReplayRequest) GoString() string {
	return s.String()
}

func (s *EditFeedReplayRequest) SetUserId(v string) *EditFeedReplayRequest {
	s.UserId = &v
	return s
}

func (s *EditFeedReplayRequest) SetEditStartTime(v int64) *EditFeedReplayRequest {
	s.EditStartTime = &v
	return s
}

func (s *EditFeedReplayRequest) SetEditEndTime(v int64) *EditFeedReplayRequest {
	s.EditEndTime = &v
	return s
}

type EditFeedReplayResponseBody struct {
	// 剪辑后的视频地址（含authkey）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EditFeedReplayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditFeedReplayResponseBody) GoString() string {
	return s.String()
}

func (s *EditFeedReplayResponseBody) SetResult(v string) *EditFeedReplayResponseBody {
	s.Result = &v
	return s
}

type EditFeedReplayResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditFeedReplayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditFeedReplayResponse) String() string {
	return tea.Prettify(s)
}

func (s EditFeedReplayResponse) GoString() string {
	return s.String()
}

func (s *EditFeedReplayResponse) SetHeaders(v map[string]*string) *EditFeedReplayResponse {
	s.Headers = v
	return s
}

func (s *EditFeedReplayResponse) SetBody(v *EditFeedReplayResponseBody) *EditFeedReplayResponse {
	s.Body = v
	return s
}

type QueryFeedWhiteListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFeedWhiteListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFeedWhiteListHeaders) GoString() string {
	return s.String()
}

func (s *QueryFeedWhiteListHeaders) SetCommonHeaders(v map[string]*string) *QueryFeedWhiteListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFeedWhiteListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFeedWhiteListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFeedWhiteListRequest struct {
	// 用户组织内id（查询该用户是否在白名单列表中）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryFeedWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFeedWhiteListRequest) GoString() string {
	return s.String()
}

func (s *QueryFeedWhiteListRequest) SetUserId(v string) *QueryFeedWhiteListRequest {
	s.UserId = &v
	return s
}

type QueryFeedWhiteListResponseBody struct {
	// 是否在白名单内
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s QueryFeedWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFeedWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFeedWhiteListResponseBody) SetResult(v bool) *QueryFeedWhiteListResponseBody {
	s.Result = &v
	return s
}

type QueryFeedWhiteListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFeedWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFeedWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFeedWhiteListResponse) GoString() string {
	return s.String()
}

func (s *QueryFeedWhiteListResponse) SetHeaders(v map[string]*string) *QueryFeedWhiteListResponse {
	s.Headers = v
	return s
}

func (s *QueryFeedWhiteListResponse) SetBody(v *QueryFeedWhiteListResponseBody) *QueryFeedWhiteListResponse {
	s.Body = v
	return s
}

type CreateCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *CreateCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCloudFeedRequest struct {
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 课程简介
	Intro *string `json:"intro,omitempty" xml:"intro,omitempty"`
	// 创建课程的主播id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 预计开始的时间戳(未来的时间点)
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 课程封面Url
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 云导播课程资源的url
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s CreateCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedRequest) SetTitle(v string) *CreateCloudFeedRequest {
	s.Title = &v
	return s
}

func (s *CreateCloudFeedRequest) SetIntro(v string) *CreateCloudFeedRequest {
	s.Intro = &v
	return s
}

func (s *CreateCloudFeedRequest) SetUserId(v string) *CreateCloudFeedRequest {
	s.UserId = &v
	return s
}

func (s *CreateCloudFeedRequest) SetStartTime(v int64) *CreateCloudFeedRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCloudFeedRequest) SetCoverUrl(v string) *CreateCloudFeedRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateCloudFeedRequest) SetVideoUrl(v string) *CreateCloudFeedRequest {
	s.VideoUrl = &v
	return s
}

type CreateCloudFeedResponseBody struct {
	// 8c0ed3c3-e125-4a9d-aa40-18ad999398d4
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedResponseBody) SetResult(v string) *CreateCloudFeedResponseBody {
	s.Result = &v
	return s
}

type CreateCloudFeedResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedResponse) SetHeaders(v map[string]*string) *CreateCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudFeedResponse) SetBody(v *CreateCloudFeedResponseBody) *CreateCloudFeedResponse {
	s.Body = v
	return s
}

type DeleteLiveFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteLiveFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveFeedHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLiveFeedHeaders) SetCommonHeaders(v map[string]*string) *DeleteLiveFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLiveFeedHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteLiveFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteLiveFeedRequest struct {
	// 用户id（操作者的组织内id）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteLiveFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveFeedRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveFeedRequest) SetUserId(v string) *DeleteLiveFeedRequest {
	s.UserId = &v
	return s
}

type DeleteLiveFeedResponseBody struct {
	// 是否删除成功
	HasDelete *bool `json:"hasDelete,omitempty" xml:"hasDelete,omitempty"`
}

func (s DeleteLiveFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveFeedResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveFeedResponseBody) SetHasDelete(v bool) *DeleteLiveFeedResponseBody {
	s.HasDelete = &v
	return s
}

type DeleteLiveFeedResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLiveFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveFeedResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveFeedResponse) SetHeaders(v map[string]*string) *DeleteLiveFeedResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveFeedResponse) SetBody(v *DeleteLiveFeedResponseBody) *DeleteLiveFeedResponse {
	s.Body = v
	return s
}

type StartCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *StartCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *StartCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCloudFeedRequest struct {
	// 操作者的组织内id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StartCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *StartCloudFeedRequest) SetUserId(v string) *StartCloudFeedRequest {
	s.UserId = &v
	return s
}

type StartCloudFeedResponseBody struct {
	// 状态更改是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudFeedResponseBody) SetResult(v bool) *StartCloudFeedResponseBody {
	s.Result = &v
	return s
}

type StartCloudFeedResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *StartCloudFeedResponse) SetHeaders(v map[string]*string) *StartCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *StartCloudFeedResponse) SetBody(v *StartCloudFeedResponseBody) *StartCloudFeedResponse {
	s.Body = v
	return s
}

type StopCloudFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopCloudFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudFeedHeaders) SetCommonHeaders(v map[string]*string) *StopCloudFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudFeedHeaders) SetXAcsDingtalkAccessToken(v string) *StopCloudFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopCloudFeedRequest struct {
	// 操作者的组织内id（staffId）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StopCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *StopCloudFeedRequest) SetUserId(v string) *StopCloudFeedRequest {
	s.UserId = &v
	return s
}

type StopCloudFeedResponseBody struct {
	// 状态更改是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StopCloudFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudFeedResponseBody) SetResult(v bool) *StopCloudFeedResponseBody {
	s.Result = &v
	return s
}

type StopCloudFeedResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCloudFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudFeedResponse) GoString() string {
	return s.String()
}

func (s *StopCloudFeedResponse) SetHeaders(v map[string]*string) *StopCloudFeedResponse {
	s.Headers = v
	return s
}

func (s *StopCloudFeedResponse) SetBody(v *StopCloudFeedResponseBody) *StopCloudFeedResponse {
	s.Body = v
	return s
}

type ModifyFeedWhiteListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ModifyFeedWhiteListHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListHeaders) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListHeaders) SetCommonHeaders(v map[string]*string) *ModifyFeedWhiteListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ModifyFeedWhiteListHeaders) SetXAcsDingtalkAccessToken(v string) *ModifyFeedWhiteListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ModifyFeedWhiteListRequest struct {
	// 用户id（操作者的组织内id）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 操作类型（1 添加白名单 / 2 删除白名单）
	Action *int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 操作的白名单列表
	ModifyUserList []*string `json:"modifyUserList,omitempty" xml:"modifyUserList,omitempty" type:"Repeated"`
}

func (s ModifyFeedWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListRequest) SetUserId(v string) *ModifyFeedWhiteListRequest {
	s.UserId = &v
	return s
}

func (s *ModifyFeedWhiteListRequest) SetAction(v int64) *ModifyFeedWhiteListRequest {
	s.Action = &v
	return s
}

func (s *ModifyFeedWhiteListRequest) SetModifyUserList(v []*string) *ModifyFeedWhiteListRequest {
	s.ModifyUserList = v
	return s
}

type ModifyFeedWhiteListShrinkRequest struct {
	// 用户id（操作者的组织内id）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 操作类型（1 添加白名单 / 2 删除白名单）
	Action *int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 操作的白名单列表
	ModifyUserListShrink *string `json:"modifyUserList,omitempty" xml:"modifyUserList,omitempty"`
}

func (s ModifyFeedWhiteListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListShrinkRequest) SetUserId(v string) *ModifyFeedWhiteListShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ModifyFeedWhiteListShrinkRequest) SetAction(v int64) *ModifyFeedWhiteListShrinkRequest {
	s.Action = &v
	return s
}

func (s *ModifyFeedWhiteListShrinkRequest) SetModifyUserListShrink(v string) *ModifyFeedWhiteListShrinkRequest {
	s.ModifyUserListShrink = &v
	return s
}

type ModifyFeedWhiteListResponseBody struct {
	// 是否修改成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyFeedWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListResponseBody) SetResult(v bool) *ModifyFeedWhiteListResponseBody {
	s.Result = &v
	return s
}

type ModifyFeedWhiteListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFeedWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFeedWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListResponse) SetHeaders(v map[string]*string) *ModifyFeedWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyFeedWhiteListResponse) SetBody(v *ModifyFeedWhiteListResponseBody) *ModifyFeedWhiteListResponse {
	s.Body = v
	return s
}

type UpdateLiveFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateLiveFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveFeedHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLiveFeedHeaders) SetCommonHeaders(v map[string]*string) *UpdateLiveFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLiveFeedHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateLiveFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateLiveFeedRequest struct {
	// 操作者id（修改者的组织内id）
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 预计开始时间（毫秒值）（课程必须预告状态才可以修改该项）
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 封面图url
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 课程简介
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
}

func (s UpdateLiveFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveFeedRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveFeedRequest) SetUserId(v string) *UpdateLiveFeedRequest {
	s.UserId = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetStartTime(v int64) *UpdateLiveFeedRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetCoverUrl(v string) *UpdateLiveFeedRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetTitle(v string) *UpdateLiveFeedRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetIntroduction(v string) *UpdateLiveFeedRequest {
	s.Introduction = &v
	return s
}

type UpdateLiveFeedResponseBody struct {
	// 是否修改成功
	HasUpdate *bool `json:"hasUpdate,omitempty" xml:"hasUpdate,omitempty"`
}

func (s UpdateLiveFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveFeedResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveFeedResponseBody) SetHasUpdate(v bool) *UpdateLiveFeedResponseBody {
	s.HasUpdate = &v
	return s
}

type UpdateLiveFeedResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLiveFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveFeedResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveFeedResponse) SetHeaders(v map[string]*string) *UpdateLiveFeedResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveFeedResponse) SetBody(v *UpdateLiveFeedResponseBody) *UpdateLiveFeedResponse {
	s.Body = v
	return s
}

type AddShareCidListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddShareCidListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListHeaders) GoString() string {
	return s.String()
}

func (s *AddShareCidListHeaders) SetCommonHeaders(v map[string]*string) *AddShareCidListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddShareCidListHeaders) SetXAcsDingtalkAccessToken(v string) *AddShareCidListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddShareCidListRequest struct {
	// 操作的的组织内id(staffId)
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 添加的联播群列表
	GroupIds []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// 传入的群id类型（1 chatId / 2 openConversationId ）
	GroupIdType *int64 `json:"groupIdType,omitempty" xml:"groupIdType,omitempty"`
}

func (s AddShareCidListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListRequest) GoString() string {
	return s.String()
}

func (s *AddShareCidListRequest) SetUserId(v string) *AddShareCidListRequest {
	s.UserId = &v
	return s
}

func (s *AddShareCidListRequest) SetGroupIds(v []*string) *AddShareCidListRequest {
	s.GroupIds = v
	return s
}

func (s *AddShareCidListRequest) SetGroupIdType(v int64) *AddShareCidListRequest {
	s.GroupIdType = &v
	return s
}

type AddShareCidListResponseBody struct {
	// 是否联播成功
	HasShareSuccess *bool `json:"hasShareSuccess,omitempty" xml:"hasShareSuccess,omitempty"`
	// 本次请求成功联播的群列表
	ShareSuccessGroupList []*string `json:"shareSuccessGroupList,omitempty" xml:"shareSuccessGroupList,omitempty" type:"Repeated"`
}

func (s AddShareCidListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListResponseBody) GoString() string {
	return s.String()
}

func (s *AddShareCidListResponseBody) SetHasShareSuccess(v bool) *AddShareCidListResponseBody {
	s.HasShareSuccess = &v
	return s
}

func (s *AddShareCidListResponseBody) SetShareSuccessGroupList(v []*string) *AddShareCidListResponseBody {
	s.ShareSuccessGroupList = v
	return s
}

type AddShareCidListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddShareCidListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddShareCidListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListResponse) GoString() string {
	return s.String()
}

func (s *AddShareCidListResponse) SetHeaders(v map[string]*string) *AddShareCidListResponse {
	s.Headers = v
	return s
}

func (s *AddShareCidListResponse) SetBody(v *AddShareCidListResponseBody) *AddShareCidListResponse {
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

func (client *Client) EditFeedReplay(feedId *string, request *EditFeedReplayRequest) (_result *EditFeedReplayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditFeedReplayHeaders{}
	_result = &EditFeedReplayResponse{}
	_body, _err := client.EditFeedReplayWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditFeedReplayWithOptions(feedId *string, request *EditFeedReplayRequest, headers *EditFeedReplayHeaders, runtime *util.RuntimeOptions) (_result *EditFeedReplayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.EditStartTime)) {
		body["editStartTime"] = request.EditStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EditEndTime)) {
		body["editEndTime"] = request.EditEndTime
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &EditFeedReplayResponse{}
	_body, _err := client.DoROARequest(tea.String("EditFeedReplay"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/openFeeds/"+tea.StringValue(feedId)+"/cutReplay"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFeedWhiteList(feedId *string, request *QueryFeedWhiteListRequest) (_result *QueryFeedWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFeedWhiteListHeaders{}
	_result = &QueryFeedWhiteListResponse{}
	_body, _err := client.QueryFeedWhiteListWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFeedWhiteListWithOptions(feedId *string, request *QueryFeedWhiteListRequest, headers *QueryFeedWhiteListHeaders, runtime *util.RuntimeOptions) (_result *QueryFeedWhiteListResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &QueryFeedWhiteListResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFeedWhiteList"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/live/openFeeds/"+tea.StringValue(feedId)+"/whiteList"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCloudFeed(request *CreateCloudFeedRequest) (_result *CreateCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCloudFeedHeaders{}
	_result = &CreateCloudFeedResponse{}
	_body, _err := client.CreateCloudFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCloudFeedWithOptions(request *CreateCloudFeedRequest, headers *CreateCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["videoUrl"] = request.VideoUrl
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveFeed(feedId *string, request *DeleteLiveFeedRequest) (_result *DeleteLiveFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteLiveFeedHeaders{}
	_result = &DeleteLiveFeedResponse{}
	_body, _err := client.DeleteLiveFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveFeedWithOptions(feedId *string, request *DeleteLiveFeedRequest, headers *DeleteLiveFeedHeaders, runtime *util.RuntimeOptions) (_result *DeleteLiveFeedResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteLiveFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteLiveFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/live/openFeeds/"+tea.StringValue(feedId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCloudFeed(feedId *string, request *StartCloudFeedRequest) (_result *StartCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCloudFeedHeaders{}
	_result = &StartCloudFeedResponse{}
	_body, _err := client.StartCloudFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCloudFeedWithOptions(feedId *string, request *StartCloudFeedRequest, headers *StartCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *StartCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StartCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("StartCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCloudFeed(feedId *string, request *StopCloudFeedRequest) (_result *StopCloudFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopCloudFeedHeaders{}
	_result = &StopCloudFeedResponse{}
	_body, _err := client.StopCloudFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCloudFeedWithOptions(feedId *string, request *StopCloudFeedRequest, headers *StopCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *StopCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &StopCloudFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("StopCloudFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFeedWhiteList(feedId *string, request *ModifyFeedWhiteListRequest) (_result *ModifyFeedWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ModifyFeedWhiteListHeaders{}
	_result = &ModifyFeedWhiteListResponse{}
	_body, _err := client.ModifyFeedWhiteListWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFeedWhiteListWithOptions(feedId *string, tmpReq *ModifyFeedWhiteListRequest, headers *ModifyFeedWhiteListHeaders, runtime *util.RuntimeOptions) (_result *ModifyFeedWhiteListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyFeedWhiteListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ModifyUserList)) {
		request.ModifyUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModifyUserList, tea.String("modifyUserList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyUserListShrink)) {
		query["modifyUserList"] = request.ModifyUserListShrink
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
	_result = &ModifyFeedWhiteListResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyFeedWhiteList"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/openFeeds/"+tea.StringValue(feedId)+"/whiteList"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveFeed(feedId *string, request *UpdateLiveFeedRequest) (_result *UpdateLiveFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateLiveFeedHeaders{}
	_result = &UpdateLiveFeedResponse{}
	_body, _err := client.UpdateLiveFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveFeedWithOptions(feedId *string, request *UpdateLiveFeedRequest, headers *UpdateLiveFeedHeaders, runtime *util.RuntimeOptions) (_result *UpdateLiveFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		query["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		query["introduction"] = request.Introduction
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
	_result = &UpdateLiveFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateLiveFeed"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/openFeeds/"+tea.StringValue(feedId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddShareCidList(feedId *string, request *AddShareCidListRequest) (_result *AddShareCidListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddShareCidListHeaders{}
	_result = &AddShareCidListResponse{}
	_body, _err := client.AddShareCidListWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddShareCidListWithOptions(feedId *string, request *AddShareCidListRequest, headers *AddShareCidListHeaders, runtime *util.RuntimeOptions) (_result *AddShareCidListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIdType)) {
		body["groupIdType"] = request.GroupIdType
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddShareCidListResponse{}
	_body, _err := client.DoROARequest(tea.String("AddShareCidList"), tea.String("live_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/live/cloudFeeds/"+tea.StringValue(feedId)+"/share"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
