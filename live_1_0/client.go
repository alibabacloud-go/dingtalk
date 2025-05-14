// This file is auto-generated, don't edit it. Thanks.
package live_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddLiveInteractionPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLiveInteractionPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLiveInteractionPluginHeaders) GoString() string {
	return s.String()
}

func (s *AddLiveInteractionPluginHeaders) SetCommonHeaders(v map[string]*string) *AddLiveInteractionPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLiveInteractionPluginHeaders) SetXAcsDingtalkAccessToken(v string) *AddLiveInteractionPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLiveInteractionPluginRequest struct {
	// This parameter is required.
	AnchorJumpUrl *string `json:"anchorJumpUrl,omitempty" xml:"anchorJumpUrl,omitempty"`
	// This parameter is required.
	PluginIconUrl *string `json:"pluginIconUrl,omitempty" xml:"pluginIconUrl,omitempty"`
	// This parameter is required.
	PluginName   *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	PluginNameEn *string `json:"pluginNameEn,omitempty" xml:"pluginNameEn,omitempty"`
	// This parameter is required.
	ViewerJumpUrl *string `json:"viewerJumpUrl,omitempty" xml:"viewerJumpUrl,omitempty"`
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddLiveInteractionPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveInteractionPluginRequest) GoString() string {
	return s.String()
}

func (s *AddLiveInteractionPluginRequest) SetAnchorJumpUrl(v string) *AddLiveInteractionPluginRequest {
	s.AnchorJumpUrl = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetPluginIconUrl(v string) *AddLiveInteractionPluginRequest {
	s.PluginIconUrl = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetPluginName(v string) *AddLiveInteractionPluginRequest {
	s.PluginName = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetPluginNameEn(v string) *AddLiveInteractionPluginRequest {
	s.PluginNameEn = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetViewerJumpUrl(v string) *AddLiveInteractionPluginRequest {
	s.ViewerJumpUrl = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetLiveId(v string) *AddLiveInteractionPluginRequest {
	s.LiveId = &v
	return s
}

func (s *AddLiveInteractionPluginRequest) SetUnionId(v string) *AddLiveInteractionPluginRequest {
	s.UnionId = &v
	return s
}

type AddLiveInteractionPluginResponseBody struct {
	Result *AddLiveInteractionPluginResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s AddLiveInteractionPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLiveInteractionPluginResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveInteractionPluginResponseBody) SetResult(v *AddLiveInteractionPluginResponseBodyResult) *AddLiveInteractionPluginResponseBody {
	s.Result = v
	return s
}

type AddLiveInteractionPluginResponseBodyResult struct {
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s AddLiveInteractionPluginResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddLiveInteractionPluginResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddLiveInteractionPluginResponseBodyResult) SetPluginId(v string) *AddLiveInteractionPluginResponseBodyResult {
	s.PluginId = &v
	return s
}

type AddLiveInteractionPluginResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveInteractionPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveInteractionPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveInteractionPluginResponse) GoString() string {
	return s.String()
}

func (s *AddLiveInteractionPluginResponse) SetHeaders(v map[string]*string) *AddLiveInteractionPluginResponse {
	s.Headers = v
	return s
}

func (s *AddLiveInteractionPluginResponse) SetStatusCode(v int32) *AddLiveInteractionPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveInteractionPluginResponse) SetBody(v *AddLiveInteractionPluginResponseBody) *AddLiveInteractionPluginResponse {
	s.Body = v
	return s
}

type AddLiveNoticeWidgetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddLiveNoticeWidgetHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddLiveNoticeWidgetHeaders) GoString() string {
	return s.String()
}

func (s *AddLiveNoticeWidgetHeaders) SetCommonHeaders(v map[string]*string) *AddLiveNoticeWidgetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddLiveNoticeWidgetHeaders) SetXAcsDingtalkAccessToken(v string) *AddLiveNoticeWidgetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddLiveNoticeWidgetRequest struct {
	// This parameter is required.
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// This parameter is required.
	JumpUrl *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	NoticeText *string `json:"noticeText,omitempty" xml:"noticeText,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s AddLiveNoticeWidgetRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLiveNoticeWidgetRequest) GoString() string {
	return s.String()
}

func (s *AddLiveNoticeWidgetRequest) SetIconUrl(v string) *AddLiveNoticeWidgetRequest {
	s.IconUrl = &v
	return s
}

func (s *AddLiveNoticeWidgetRequest) SetJumpUrl(v string) *AddLiveNoticeWidgetRequest {
	s.JumpUrl = &v
	return s
}

func (s *AddLiveNoticeWidgetRequest) SetLiveId(v string) *AddLiveNoticeWidgetRequest {
	s.LiveId = &v
	return s
}

func (s *AddLiveNoticeWidgetRequest) SetNoticeText(v string) *AddLiveNoticeWidgetRequest {
	s.NoticeText = &v
	return s
}

func (s *AddLiveNoticeWidgetRequest) SetUnionId(v string) *AddLiveNoticeWidgetRequest {
	s.UnionId = &v
	return s
}

type AddLiveNoticeWidgetResponseBody struct {
	Result *AddLiveNoticeWidgetResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s AddLiveNoticeWidgetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLiveNoticeWidgetResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveNoticeWidgetResponseBody) SetResult(v *AddLiveNoticeWidgetResponseBodyResult) *AddLiveNoticeWidgetResponseBody {
	s.Result = v
	return s
}

type AddLiveNoticeWidgetResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddLiveNoticeWidgetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddLiveNoticeWidgetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddLiveNoticeWidgetResponseBodyResult) SetSuccess(v bool) *AddLiveNoticeWidgetResponseBodyResult {
	s.Success = &v
	return s
}

type AddLiveNoticeWidgetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveNoticeWidgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveNoticeWidgetResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLiveNoticeWidgetResponse) GoString() string {
	return s.String()
}

func (s *AddLiveNoticeWidgetResponse) SetHeaders(v map[string]*string) *AddLiveNoticeWidgetResponse {
	s.Headers = v
	return s
}

func (s *AddLiveNoticeWidgetResponse) SetStatusCode(v int32) *AddLiveNoticeWidgetResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveNoticeWidgetResponse) SetBody(v *AddLiveNoticeWidgetResponseBody) *AddLiveNoticeWidgetResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 2
	GroupIdType *int64 `json:"groupIdType,omitempty" xml:"groupIdType,omitempty"`
	// This parameter is required.
	GroupIds []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 214675
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddShareCidListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShareCidListRequest) GoString() string {
	return s.String()
}

func (s *AddShareCidListRequest) SetGroupIdType(v int64) *AddShareCidListRequest {
	s.GroupIdType = &v
	return s
}

func (s *AddShareCidListRequest) SetGroupIds(v []*string) *AddShareCidListRequest {
	s.GroupIds = v
	return s
}

func (s *AddShareCidListRequest) SetUserId(v string) *AddShareCidListRequest {
	s.UserId = &v
	return s
}

type AddShareCidListResponseBody struct {
	// example:
	//
	// true
	HasShareSuccess       *bool     `json:"hasShareSuccess,omitempty" xml:"hasShareSuccess,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShareCidListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddShareCidListResponse) SetStatusCode(v int32) *AddShareCidListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShareCidListResponse) SetBody(v *AddShareCidListResponseBody) *AddShareCidListResponse {
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
	// example:
	//
	// https://img.alicdn.com/tfs/TB1A7cBtYr1gK0jSZR0XXbP8XXa-750-422.png
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 这是一场云导播课程
	Intro *string `json:"intro,omitempty" xml:"intro,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615260061000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 课程一
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 214675
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http/https:/xxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s CreateCloudFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCloudFeedRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudFeedRequest) SetCoverUrl(v string) *CreateCloudFeedRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateCloudFeedRequest) SetIntro(v string) *CreateCloudFeedRequest {
	s.Intro = &v
	return s
}

func (s *CreateCloudFeedRequest) SetStartTime(v int64) *CreateCloudFeedRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCloudFeedRequest) SetTitle(v string) *CreateCloudFeedRequest {
	s.Title = &v
	return s
}

func (s *CreateCloudFeedRequest) SetUserId(v string) *CreateCloudFeedRequest {
	s.UserId = &v
	return s
}

func (s *CreateCloudFeedRequest) SetVideoUrl(v string) *CreateCloudFeedRequest {
	s.VideoUrl = &v
	return s
}

type CreateCloudFeedResponseBody struct {
	// example:
	//
	// 创建好的云导播课程id
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCloudFeedResponse) SetStatusCode(v int32) *CreateCloudFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudFeedResponse) SetBody(v *CreateCloudFeedResponseBody) *CreateCloudFeedResponse {
	s.Body = v
	return s
}

type CreateLiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateLiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveHeaders) GoString() string {
	return s.String()
}

func (s *CreateLiveHeaders) SetCommonHeaders(v map[string]*string) *CreateLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateLiveHeaders) SetXAcsDingtalkAccessToken(v string) *CreateLiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateLiveRequest struct {
	// example:
	//
	// https://gw.alicdn.com/tfs/TB1thlYyAT2gK0jSZPcXXcKkpXa-1125-633.png
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 测试直播简介
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1659653648000
	PreEndTime *int64 `json:"preEndTime,omitempty" xml:"preEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1659613648000
	PreStartTime *int64 `json:"preStartTime,omitempty" xml:"preStartTime,omitempty"`
	// example:
	//
	// 2
	PublicType *int64 `json:"publicType,omitempty" xml:"publicType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试直播
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetCoverUrl(v string) *CreateLiveRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRequest) SetIntroduction(v string) *CreateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveRequest) SetPreEndTime(v int64) *CreateLiveRequest {
	s.PreEndTime = &v
	return s
}

func (s *CreateLiveRequest) SetPreStartTime(v int64) *CreateLiveRequest {
	s.PreStartTime = &v
	return s
}

func (s *CreateLiveRequest) SetPublicType(v int64) *CreateLiveRequest {
	s.PublicType = &v
	return s
}

func (s *CreateLiveRequest) SetTitle(v string) *CreateLiveRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRequest) SetUnionId(v string) *CreateLiveRequest {
	s.UnionId = &v
	return s
}

type CreateLiveResponseBody struct {
	Result *CreateLiveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) SetResult(v *CreateLiveResponseBodyResult) *CreateLiveResponseBody {
	s.Result = v
	return s
}

type CreateLiveResponseBodyResult struct {
	// example:
	//
	// 1a353547-040d-4095-bb93-404bc5d47920
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
}

func (s CreateLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBodyResult) SetLiveId(v string) *CreateLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

type CreateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetStatusCode(v int32) *CreateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

type DelLiveInteractionPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DelLiveInteractionPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s DelLiveInteractionPluginHeaders) GoString() string {
	return s.String()
}

func (s *DelLiveInteractionPluginHeaders) SetCommonHeaders(v map[string]*string) *DelLiveInteractionPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DelLiveInteractionPluginHeaders) SetXAcsDingtalkAccessToken(v string) *DelLiveInteractionPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DelLiveInteractionPluginRequest struct {
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DelLiveInteractionPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s DelLiveInteractionPluginRequest) GoString() string {
	return s.String()
}

func (s *DelLiveInteractionPluginRequest) SetLiveId(v string) *DelLiveInteractionPluginRequest {
	s.LiveId = &v
	return s
}

func (s *DelLiveInteractionPluginRequest) SetPluginId(v string) *DelLiveInteractionPluginRequest {
	s.PluginId = &v
	return s
}

func (s *DelLiveInteractionPluginRequest) SetUnionId(v string) *DelLiveInteractionPluginRequest {
	s.UnionId = &v
	return s
}

type DelLiveInteractionPluginResponseBody struct {
	Result *DelLiveInteractionPluginResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DelLiveInteractionPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelLiveInteractionPluginResponseBody) GoString() string {
	return s.String()
}

func (s *DelLiveInteractionPluginResponseBody) SetResult(v *DelLiveInteractionPluginResponseBodyResult) *DelLiveInteractionPluginResponseBody {
	s.Result = v
	return s
}

type DelLiveInteractionPluginResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DelLiveInteractionPluginResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DelLiveInteractionPluginResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DelLiveInteractionPluginResponseBodyResult) SetSuccess(v bool) *DelLiveInteractionPluginResponseBodyResult {
	s.Success = &v
	return s
}

type DelLiveInteractionPluginResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelLiveInteractionPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelLiveInteractionPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s DelLiveInteractionPluginResponse) GoString() string {
	return s.String()
}

func (s *DelLiveInteractionPluginResponse) SetHeaders(v map[string]*string) *DelLiveInteractionPluginResponse {
	s.Headers = v
	return s
}

func (s *DelLiveInteractionPluginResponse) SetStatusCode(v int32) *DelLiveInteractionPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *DelLiveInteractionPluginResponse) SetBody(v *DelLiveInteractionPluginResponseBody) *DelLiveInteractionPluginResponse {
	s.Body = v
	return s
}

type DelLiveNoticeWidgetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DelLiveNoticeWidgetHeaders) String() string {
	return tea.Prettify(s)
}

func (s DelLiveNoticeWidgetHeaders) GoString() string {
	return s.String()
}

func (s *DelLiveNoticeWidgetHeaders) SetCommonHeaders(v map[string]*string) *DelLiveNoticeWidgetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DelLiveNoticeWidgetHeaders) SetXAcsDingtalkAccessToken(v string) *DelLiveNoticeWidgetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DelLiveNoticeWidgetRequest struct {
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DelLiveNoticeWidgetRequest) String() string {
	return tea.Prettify(s)
}

func (s DelLiveNoticeWidgetRequest) GoString() string {
	return s.String()
}

func (s *DelLiveNoticeWidgetRequest) SetLiveId(v string) *DelLiveNoticeWidgetRequest {
	s.LiveId = &v
	return s
}

func (s *DelLiveNoticeWidgetRequest) SetUnionId(v string) *DelLiveNoticeWidgetRequest {
	s.UnionId = &v
	return s
}

type DelLiveNoticeWidgetResponseBody struct {
	Result *DelLiveNoticeWidgetResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DelLiveNoticeWidgetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelLiveNoticeWidgetResponseBody) GoString() string {
	return s.String()
}

func (s *DelLiveNoticeWidgetResponseBody) SetResult(v *DelLiveNoticeWidgetResponseBodyResult) *DelLiveNoticeWidgetResponseBody {
	s.Result = v
	return s
}

type DelLiveNoticeWidgetResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DelLiveNoticeWidgetResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DelLiveNoticeWidgetResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DelLiveNoticeWidgetResponseBodyResult) SetSuccess(v bool) *DelLiveNoticeWidgetResponseBodyResult {
	s.Success = &v
	return s
}

type DelLiveNoticeWidgetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelLiveNoticeWidgetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelLiveNoticeWidgetResponse) String() string {
	return tea.Prettify(s)
}

func (s DelLiveNoticeWidgetResponse) GoString() string {
	return s.String()
}

func (s *DelLiveNoticeWidgetResponse) SetHeaders(v map[string]*string) *DelLiveNoticeWidgetResponse {
	s.Headers = v
	return s
}

func (s *DelLiveNoticeWidgetResponse) SetStatusCode(v int32) *DelLiveNoticeWidgetResponse {
	s.StatusCode = &v
	return s
}

func (s *DelLiveNoticeWidgetResponse) SetBody(v *DelLiveNoticeWidgetResponseBody) *DelLiveNoticeWidgetResponse {
	s.Body = v
	return s
}

type DeleteLiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteLiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLiveHeaders) SetCommonHeaders(v map[string]*string) *DeleteLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLiveHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteLiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteLiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d383876-1ff9-4b73-a057-a8f47b346ecb
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s DeleteLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequest) SetLiveId(v string) *DeleteLiveRequest {
	s.LiveId = &v
	return s
}

func (s *DeleteLiveRequest) SetUnionId(v string) *DeleteLiveRequest {
	s.UnionId = &v
	return s
}

type DeleteLiveResponseBody struct {
	Result *DeleteLiveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBody) SetResult(v *DeleteLiveResponseBodyResult) *DeleteLiveResponseBody {
	s.Result = v
	return s
}

type DeleteLiveResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBodyResult) SetSuccess(v bool) *DeleteLiveResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponse) SetHeaders(v map[string]*string) *DeleteLiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveResponse) SetStatusCode(v int32) *DeleteLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveResponse) SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
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
	// example:
	//
	// true
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteLiveFeedResponse) SetStatusCode(v int32) *DeleteLiveFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveFeedResponse) SetBody(v *DeleteLiveFeedResponseBody) *DeleteLiveFeedResponse {
	s.Body = v
	return s
}

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
	// This parameter is required.
	//
	// example:
	//
	// 1617356058000
	EditEndTime *int64 `json:"editEndTime,omitempty" xml:"editEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1617336058000
	EditStartTime *int64 `json:"editStartTime,omitempty" xml:"editStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EditFeedReplayRequest) String() string {
	return tea.Prettify(s)
}

func (s EditFeedReplayRequest) GoString() string {
	return s.String()
}

func (s *EditFeedReplayRequest) SetEditEndTime(v int64) *EditFeedReplayRequest {
	s.EditEndTime = &v
	return s
}

func (s *EditFeedReplayRequest) SetEditStartTime(v int64) *EditFeedReplayRequest {
	s.EditStartTime = &v
	return s
}

func (s *EditFeedReplayRequest) SetUserId(v string) *EditFeedReplayRequest {
	s.UserId = &v
	return s
}

type EditFeedReplayResponseBody struct {
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditFeedReplayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EditFeedReplayResponse) SetStatusCode(v int32) *EditFeedReplayResponse {
	s.StatusCode = &v
	return s
}

func (s *EditFeedReplayResponse) SetBody(v *EditFeedReplayResponseBody) *EditFeedReplayResponse {
	s.Body = v
	return s
}

type GetGroupLiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupLiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetGroupLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupLiveListHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupLiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupLiveListRequest struct {
	RequestBody *GetGroupLiveListRequestRequestBody `json:"requestBody,omitempty" xml:"requestBody,omitempty" type:"Struct"`
	UnionId     *string                             `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetGroupLiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListRequest) SetRequestBody(v *GetGroupLiveListRequestRequestBody) *GetGroupLiveListRequest {
	s.RequestBody = v
	return s
}

func (s *GetGroupLiveListRequest) SetUnionId(v string) *GetGroupLiveListRequest {
	s.UnionId = &v
	return s
}

type GetGroupLiveListRequestRequestBody struct {
	EndTime            *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	StartTime          *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetGroupLiveListRequestRequestBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListRequestRequestBody) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListRequestRequestBody) SetEndTime(v int64) *GetGroupLiveListRequestRequestBody {
	s.EndTime = &v
	return s
}

func (s *GetGroupLiveListRequestRequestBody) SetOpenConversationId(v string) *GetGroupLiveListRequestRequestBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetGroupLiveListRequestRequestBody) SetStartTime(v int64) *GetGroupLiveListRequestRequestBody {
	s.StartTime = &v
	return s
}

type GetGroupLiveListShrinkRequest struct {
	RequestBodyShrink *string `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	UnionId           *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetGroupLiveListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListShrinkRequest) SetRequestBodyShrink(v string) *GetGroupLiveListShrinkRequest {
	s.RequestBodyShrink = &v
	return s
}

func (s *GetGroupLiveListShrinkRequest) SetUnionId(v string) *GetGroupLiveListShrinkRequest {
	s.UnionId = &v
	return s
}

type GetGroupLiveListResponseBody struct {
	Result *GetGroupLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetGroupLiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBody) SetResult(v *GetGroupLiveListResponseBodyResult) *GetGroupLiveListResponseBody {
	s.Result = v
	return s
}

type GetGroupLiveListResponseBodyResult struct {
	GroupLiveList []*GetGroupLiveListResponseBodyResultGroupLiveList `json:"groupLiveList,omitempty" xml:"groupLiveList,omitempty" type:"Repeated"`
}

func (s GetGroupLiveListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBodyResult) SetGroupLiveList(v []*GetGroupLiveListResponseBodyResultGroupLiveList) *GetGroupLiveListResponseBodyResult {
	s.GroupLiveList = v
	return s
}

type GetGroupLiveListResponseBodyResultGroupLiveList struct {
	AnchorNickname *string `json:"anchorNickname,omitempty" xml:"anchorNickname,omitempty"`
	AnchorUnionId  *string `json:"anchorUnionId,omitempty" xml:"anchorUnionId,omitempty"`
	LiveEndTime    *int64  `json:"liveEndTime,omitempty" xml:"liveEndTime,omitempty"`
	LiveStartTime  *int64  `json:"liveStartTime,omitempty" xml:"liveStartTime,omitempty"`
	LiveUuid       *string `json:"liveUuid,omitempty" xml:"liveUuid,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetGroupLiveListResponseBodyResultGroupLiveList) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListResponseBodyResultGroupLiveList) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetAnchorNickname(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetAnchorUnionId(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveEndTime(v int64) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveStartTime(v int64) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetLiveUuid(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetGroupLiveListResponseBodyResultGroupLiveList) SetTitle(v string) *GetGroupLiveListResponseBodyResultGroupLiveList {
	s.Title = &v
	return s
}

type GetGroupLiveListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupLiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetGroupLiveListResponse) SetHeaders(v map[string]*string) *GetGroupLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetGroupLiveListResponse) SetStatusCode(v int32) *GetGroupLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupLiveListResponse) SetBody(v *GetGroupLiveListResponseBody) *GetGroupLiveListResponse {
	s.Body = v
	return s
}

type GetLiveReplayUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLiveReplayUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLiveReplayUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlHeaders) SetCommonHeaders(v map[string]*string) *GetLiveReplayUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLiveReplayUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetLiveReplayUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLiveReplayUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d383876-1ff9-4b73-a057-a8f47b346ecb
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetLiveReplayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveReplayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlRequest) SetLiveId(v string) *GetLiveReplayUrlRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveReplayUrlRequest) SetUnionId(v string) *GetLiveReplayUrlRequest {
	s.UnionId = &v
	return s
}

type GetLiveReplayUrlResponseBody struct {
	Result *GetLiveReplayUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetLiveReplayUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveReplayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlResponseBody) SetResult(v *GetLiveReplayUrlResponseBodyResult) *GetLiveReplayUrlResponseBody {
	s.Result = v
	return s
}

type GetLiveReplayUrlResponseBodyResult struct {
	// example:
	//
	// http://xxx.dingtalk.com/live_hp/7c7ba32a-c92d-4524-b71e-33a72575c5a9_normal.m3u8?auth_key=xxx
	ReplayUrl *string `json:"replayUrl,omitempty" xml:"replayUrl,omitempty"`
}

func (s GetLiveReplayUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveReplayUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlResponseBodyResult) SetReplayUrl(v string) *GetLiveReplayUrlResponseBodyResult {
	s.ReplayUrl = &v
	return s
}

type GetLiveReplayUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveReplayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveReplayUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveReplayUrlResponse) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlResponse) SetHeaders(v map[string]*string) *GetLiveReplayUrlResponse {
	s.Headers = v
	return s
}

func (s *GetLiveReplayUrlResponse) SetStatusCode(v int32) *GetLiveReplayUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveReplayUrlResponse) SetBody(v *GetLiveReplayUrlResponseBody) *GetLiveReplayUrlResponse {
	s.Body = v
	return s
}

type GetOrgLiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrgLiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetOrgLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgLiveListHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrgLiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrgLiveListRequest struct {
	// This parameter is required.
	CorpId      *string                           `json:"corpId,omitempty" xml:"corpId,omitempty"`
	RequestBody *GetOrgLiveListRequestRequestBody `json:"requestBody,omitempty" xml:"requestBody,omitempty" type:"Struct"`
}

func (s GetOrgLiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListRequest) SetCorpId(v string) *GetOrgLiveListRequest {
	s.CorpId = &v
	return s
}

func (s *GetOrgLiveListRequest) SetRequestBody(v *GetOrgLiveListRequestRequestBody) *GetOrgLiveListRequest {
	s.RequestBody = v
	return s
}

type GetOrgLiveListRequestRequestBody struct {
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize  *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetOrgLiveListRequestRequestBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListRequestRequestBody) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListRequestRequestBody) SetEndTime(v int64) *GetOrgLiveListRequestRequestBody {
	s.EndTime = &v
	return s
}

func (s *GetOrgLiveListRequestRequestBody) SetPageNumber(v int64) *GetOrgLiveListRequestRequestBody {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListRequestRequestBody) SetPageSize(v int64) *GetOrgLiveListRequestRequestBody {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListRequestRequestBody) SetStartTime(v int64) *GetOrgLiveListRequestRequestBody {
	s.StartTime = &v
	return s
}

func (s *GetOrgLiveListRequestRequestBody) SetUnionId(v string) *GetOrgLiveListRequestRequestBody {
	s.UnionId = &v
	return s
}

type GetOrgLiveListShrinkRequest struct {
	// This parameter is required.
	CorpId            *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	RequestBodyShrink *string `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
}

func (s GetOrgLiveListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListShrinkRequest) SetCorpId(v string) *GetOrgLiveListShrinkRequest {
	s.CorpId = &v
	return s
}

func (s *GetOrgLiveListShrinkRequest) SetRequestBodyShrink(v string) *GetOrgLiveListShrinkRequest {
	s.RequestBodyShrink = &v
	return s
}

type GetOrgLiveListResponseBody struct {
	Result *GetOrgLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOrgLiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBody) SetResult(v *GetOrgLiveListResponseBodyResult) *GetOrgLiveListResponseBody {
	s.Result = v
	return s
}

type GetOrgLiveListResponseBodyResult struct {
	NewLive    *GetOrgLiveListResponseBodyResultNewLive    `json:"newLive,omitempty" xml:"newLive,omitempty" type:"Struct"`
	UpdateLive *GetOrgLiveListResponseBodyResultUpdateLive `json:"updateLive,omitempty" xml:"updateLive,omitempty" type:"Struct"`
}

func (s GetOrgLiveListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResult) SetNewLive(v *GetOrgLiveListResponseBodyResultNewLive) *GetOrgLiveListResponseBodyResult {
	s.NewLive = v
	return s
}

func (s *GetOrgLiveListResponseBodyResult) SetUpdateLive(v *GetOrgLiveListResponseBodyResultUpdateLive) *GetOrgLiveListResponseBodyResult {
	s.UpdateLive = v
	return s
}

type GetOrgLiveListResponseBodyResultNewLive struct {
	HasMore    *bool                                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	LiveList   []*GetOrgLiveListResponseBodyResultNewLiveLiveList `json:"liveList,omitempty" xml:"liveList,omitempty" type:"Repeated"`
	PageNumber *int64                                             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultNewLive) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultNewLive) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetHasMore(v bool) *GetOrgLiveListResponseBodyResultNewLive {
	s.HasMore = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetLiveList(v []*GetOrgLiveListResponseBodyResultNewLiveLiveList) *GetOrgLiveListResponseBodyResultNewLive {
	s.LiveList = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetPageNumber(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetPageSize(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLive) SetTotalCount(v int64) *GetOrgLiveListResponseBodyResultNewLive {
	s.TotalCount = &v
	return s
}

type GetOrgLiveListResponseBodyResultNewLiveLiveList struct {
	AnchorNickname           *string   `json:"anchorNickname,omitempty" xml:"anchorNickname,omitempty"`
	AnchorUnionId            *string   `json:"anchorUnionId,omitempty" xml:"anchorUnionId,omitempty"`
	LiveEndTime              *int64    `json:"liveEndTime,omitempty" xml:"liveEndTime,omitempty"`
	LiveStartTime            *int64    `json:"liveStartTime,omitempty" xml:"liveStartTime,omitempty"`
	LiveUuid                 *string   `json:"liveUuid,omitempty" xml:"liveUuid,omitempty"`
	ShareOpenConversationIds []*string `json:"shareOpenConversationIds,omitempty" xml:"shareOpenConversationIds,omitempty" type:"Repeated"`
	Title                    *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultNewLiveLiveList) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultNewLiveLiveList) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetAnchorNickname(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetAnchorUnionId(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveEndTime(v int64) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveStartTime(v int64) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetLiveUuid(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetShareOpenConversationIds(v []*string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.ShareOpenConversationIds = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultNewLiveLiveList) SetTitle(v string) *GetOrgLiveListResponseBodyResultNewLiveLiveList {
	s.Title = &v
	return s
}

type GetOrgLiveListResponseBodyResultUpdateLive struct {
	HasMore    *bool                                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	LiveList   []*GetOrgLiveListResponseBodyResultUpdateLiveLiveList `json:"liveList,omitempty" xml:"liveList,omitempty" type:"Repeated"`
	PageNumber *int64                                                `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64                                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int64                                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultUpdateLive) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultUpdateLive) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetHasMore(v bool) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.HasMore = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetLiveList(v []*GetOrgLiveListResponseBodyResultUpdateLiveLiveList) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.LiveList = v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetPageNumber(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.PageNumber = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetPageSize(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.PageSize = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLive) SetTotalCount(v int64) *GetOrgLiveListResponseBodyResultUpdateLive {
	s.TotalCount = &v
	return s
}

type GetOrgLiveListResponseBodyResultUpdateLiveLiveList struct {
	AnchorNickname *string `json:"anchorNickname,omitempty" xml:"anchorNickname,omitempty"`
	AnchorUnionId  *string `json:"anchorUnionId,omitempty" xml:"anchorUnionId,omitempty"`
	LiveEndTime    *int64  `json:"liveEndTime,omitempty" xml:"liveEndTime,omitempty"`
	LiveStartTime  *int64  `json:"liveStartTime,omitempty" xml:"liveStartTime,omitempty"`
	LiveUuid       *string `json:"liveUuid,omitempty" xml:"liveUuid,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetOrgLiveListResponseBodyResultUpdateLiveLiveList) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponseBodyResultUpdateLiveLiveList) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetAnchorNickname(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.AnchorNickname = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetAnchorUnionId(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.AnchorUnionId = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveEndTime(v int64) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveEndTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveStartTime(v int64) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveStartTime = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetLiveUuid(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.LiveUuid = &v
	return s
}

func (s *GetOrgLiveListResponseBodyResultUpdateLiveLiveList) SetTitle(v string) *GetOrgLiveListResponseBodyResultUpdateLiveLiveList {
	s.Title = &v
	return s
}

type GetOrgLiveListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrgLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrgLiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrgLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetOrgLiveListResponse) SetHeaders(v map[string]*string) *GetOrgLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetOrgLiveListResponse) SetStatusCode(v int32) *GetOrgLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrgLiveListResponse) SetBody(v *GetOrgLiveListResponseBody) *GetOrgLiveListResponse {
	s.Body = v
	return s
}

type GetUserAllLiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserAllLiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetUserAllLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserAllLiveListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserAllLiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserAllLiveListRequest struct {
	EndTime   *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Statuses  []*int32 `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	Title     *string  `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6crtQT2XOgPHviiPvXhhiP6gdhiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserAllLiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListRequest) SetEndTime(v int64) *GetUserAllLiveListRequest {
	s.EndTime = &v
	return s
}

func (s *GetUserAllLiveListRequest) SetStartTime(v int64) *GetUserAllLiveListRequest {
	s.StartTime = &v
	return s
}

func (s *GetUserAllLiveListRequest) SetStatuses(v []*int32) *GetUserAllLiveListRequest {
	s.Statuses = v
	return s
}

func (s *GetUserAllLiveListRequest) SetTitle(v string) *GetUserAllLiveListRequest {
	s.Title = &v
	return s
}

func (s *GetUserAllLiveListRequest) SetPageNumber(v int32) *GetUserAllLiveListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserAllLiveListRequest) SetPageSize(v int32) *GetUserAllLiveListRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserAllLiveListRequest) SetUnionId(v string) *GetUserAllLiveListRequest {
	s.UnionId = &v
	return s
}

type GetUserAllLiveListResponseBody struct {
	Result *GetUserAllLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetUserAllLiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponseBody) SetResult(v *GetUserAllLiveListResponseBodyResult) *GetUserAllLiveListResponseBody {
	s.Result = v
	return s
}

type GetUserAllLiveListResponseBodyResult struct {
	HasFinish            *bool                                                       `json:"hasFinish,omitempty" xml:"hasFinish,omitempty"`
	LiveInfoPopModelList []*GetUserAllLiveListResponseBodyResultLiveInfoPopModelList `json:"liveInfoPopModelList,omitempty" xml:"liveInfoPopModelList,omitempty" type:"Repeated"`
}

func (s GetUserAllLiveListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponseBodyResult) SetHasFinish(v bool) *GetUserAllLiveListResponseBodyResult {
	s.HasFinish = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResult) SetLiveInfoPopModelList(v []*GetUserAllLiveListResponseBodyResultLiveInfoPopModelList) *GetUserAllLiveListResponseBodyResult {
	s.LiveInfoPopModelList = v
	return s
}

type GetUserAllLiveListResponseBodyResultLiveInfoPopModelList struct {
	ExtraInfo     *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo     `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	LiveBasicInfo *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo `json:"liveBasicInfo,omitempty" xml:"liveBasicInfo,omitempty" type:"Struct"`
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelList) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelList) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelList) SetExtraInfo(v *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelList {
	s.ExtraInfo = v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelList) SetLiveBasicInfo(v *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelList {
	s.LiveBasicInfo = v
	return s
}

type GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo struct {
	HasSubscribed     *bool  `json:"hasSubscribed,omitempty" xml:"hasSubscribed,omitempty"`
	IsForecastExpired *bool  `json:"isForecastExpired,omitempty" xml:"isForecastExpired,omitempty"`
	WatchProgressMs   *int64 `json:"watchProgressMs,omitempty" xml:"watchProgressMs,omitempty"`
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetHasSubscribed(v bool) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.HasSubscribed = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetIsForecastExpired(v bool) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.IsForecastExpired = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetWatchProgressMs(v int64) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.WatchProgressMs = &v
	return s
}

type GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo struct {
	CoverUrl       *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	Duration       *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime        *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Introduction   *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	LiveId         *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	LivePlayUrl    *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	LiveStatus     *int32  `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscribeCount *int32  `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	UnionId        *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	Uv             *int32  `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetCoverUrl(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.CoverUrl = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetDuration(v int64) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetEndTime(v int64) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.EndTime = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetIntroduction(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Introduction = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveId(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveId = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLivePlayUrl(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LivePlayUrl = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveStatus(v int32) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveStatus = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetStartTime(v int64) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.StartTime = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetSubscribeCount(v int32) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.SubscribeCount = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetTitle(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Title = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUnionId(v string) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.UnionId = &v
	return s
}

func (s *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUv(v int32) *GetUserAllLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Uv = &v
	return s
}

type GetUserAllLiveListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAllLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAllLiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserAllLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetUserAllLiveListResponse) SetHeaders(v map[string]*string) *GetUserAllLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetUserAllLiveListResponse) SetStatusCode(v int32) *GetUserAllLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAllLiveListResponse) SetBody(v *GetUserAllLiveListResponseBody) *GetUserAllLiveListResponse {
	s.Body = v
	return s
}

type GetUserCreateLiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserCreateLiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetUserCreateLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserCreateLiveListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserCreateLiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserCreateLiveListRequest struct {
	EndTime   *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Statuses  []*int64 `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	Title     *string  `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5yAiiJDWiiCJpd3Thhx7P5fgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserCreateLiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListRequest) SetEndTime(v int64) *GetUserCreateLiveListRequest {
	s.EndTime = &v
	return s
}

func (s *GetUserCreateLiveListRequest) SetStartTime(v int64) *GetUserCreateLiveListRequest {
	s.StartTime = &v
	return s
}

func (s *GetUserCreateLiveListRequest) SetStatuses(v []*int64) *GetUserCreateLiveListRequest {
	s.Statuses = v
	return s
}

func (s *GetUserCreateLiveListRequest) SetTitle(v string) *GetUserCreateLiveListRequest {
	s.Title = &v
	return s
}

func (s *GetUserCreateLiveListRequest) SetMaxResults(v int32) *GetUserCreateLiveListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetUserCreateLiveListRequest) SetNextToken(v string) *GetUserCreateLiveListRequest {
	s.NextToken = &v
	return s
}

func (s *GetUserCreateLiveListRequest) SetUnionId(v string) *GetUserCreateLiveListRequest {
	s.UnionId = &v
	return s
}

type GetUserCreateLiveListResponseBody struct {
	Result *GetUserCreateLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetUserCreateLiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponseBody) SetResult(v *GetUserCreateLiveListResponseBodyResult) *GetUserCreateLiveListResponseBody {
	s.Result = v
	return s
}

type GetUserCreateLiveListResponseBodyResult struct {
	HasFinish            *bool                                                          `json:"hasFinish,omitempty" xml:"hasFinish,omitempty"`
	LiveInfoPopModelList []*GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList `json:"liveInfoPopModelList,omitempty" xml:"liveInfoPopModelList,omitempty" type:"Repeated"`
	NextToken            *string                                                        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Total                *int32                                                         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUserCreateLiveListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponseBodyResult) SetHasFinish(v bool) *GetUserCreateLiveListResponseBodyResult {
	s.HasFinish = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResult) SetLiveInfoPopModelList(v []*GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList) *GetUserCreateLiveListResponseBodyResult {
	s.LiveInfoPopModelList = v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResult) SetNextToken(v string) *GetUserCreateLiveListResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResult) SetTotal(v int32) *GetUserCreateLiveListResponseBodyResult {
	s.Total = &v
	return s
}

type GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList struct {
	HasSubscribed *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed `json:"hasSubscribed,omitempty" xml:"hasSubscribed,omitempty" type:"Struct"`
	LiveBasicInfo *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo `json:"liveBasicInfo,omitempty" xml:"liveBasicInfo,omitempty" type:"Struct"`
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList) SetHasSubscribed(v *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList {
	s.HasSubscribed = v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList) SetLiveBasicInfo(v *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelList {
	s.LiveBasicInfo = v
	return s
}

type GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed struct {
	HasSubscribed     *bool  `json:"hasSubscribed,omitempty" xml:"hasSubscribed,omitempty"`
	IsForecastExpired *bool  `json:"isForecastExpired,omitempty" xml:"isForecastExpired,omitempty"`
	WatchProgressMs   *int64 `json:"watchProgressMs,omitempty" xml:"watchProgressMs,omitempty"`
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) SetHasSubscribed(v bool) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed {
	s.HasSubscribed = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) SetIsForecastExpired(v bool) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed {
	s.IsForecastExpired = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed) SetWatchProgressMs(v int64) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListHasSubscribed {
	s.WatchProgressMs = &v
	return s
}

type GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo struct {
	CoverUrl       *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	Duration       *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime        *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Introduction   *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	LiveId         *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	LivePlayUrl    *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	LiveStatus     *int32  `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscribeCount *int32  `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	UnionId        *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	Uv             *int32  `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetCoverUrl(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.CoverUrl = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetDuration(v int64) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetEndTime(v int64) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.EndTime = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetIntroduction(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Introduction = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveId(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveId = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLivePlayUrl(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LivePlayUrl = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveStatus(v int32) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveStatus = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetStartTime(v int64) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.StartTime = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetSubscribeCount(v int32) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.SubscribeCount = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetTitle(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Title = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUnionId(v string) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.UnionId = &v
	return s
}

func (s *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUv(v int32) *GetUserCreateLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Uv = &v
	return s
}

type GetUserCreateLiveListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserCreateLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserCreateLiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserCreateLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetUserCreateLiveListResponse) SetHeaders(v map[string]*string) *GetUserCreateLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetUserCreateLiveListResponse) SetStatusCode(v int32) *GetUserCreateLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCreateLiveListResponse) SetBody(v *GetUserCreateLiveListResponseBody) *GetUserCreateLiveListResponse {
	s.Body = v
	return s
}

type GetUserWatchLiveListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserWatchLiveListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListHeaders) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListHeaders) SetCommonHeaders(v map[string]*string) *GetUserWatchLiveListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserWatchLiveListHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserWatchLiveListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserWatchLiveListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	FilterType *int32 `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// preOrStartTime_desc_1658804913000
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6crtQT2XOgPHviiPvXhhiP6gdhiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetUserWatchLiveListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListRequest) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListRequest) SetFilterType(v int32) *GetUserWatchLiveListRequest {
	s.FilterType = &v
	return s
}

func (s *GetUserWatchLiveListRequest) SetMaxResults(v int32) *GetUserWatchLiveListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetUserWatchLiveListRequest) SetNextToken(v string) *GetUserWatchLiveListRequest {
	s.NextToken = &v
	return s
}

func (s *GetUserWatchLiveListRequest) SetUnionId(v string) *GetUserWatchLiveListRequest {
	s.UnionId = &v
	return s
}

type GetUserWatchLiveListResponseBody struct {
	Result *GetUserWatchLiveListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetUserWatchLiveListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponseBody) SetResult(v *GetUserWatchLiveListResponseBodyResult) *GetUserWatchLiveListResponseBody {
	s.Result = v
	return s
}

type GetUserWatchLiveListResponseBodyResult struct {
	HasFinish            *bool                                                         `json:"hasFinish,omitempty" xml:"hasFinish,omitempty"`
	LiveInfoPopModelList []*GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList `json:"liveInfoPopModelList,omitempty" xml:"liveInfoPopModelList,omitempty" type:"Repeated"`
	NextToken            *string                                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Total                *int32                                                        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUserWatchLiveListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponseBodyResult) SetHasFinish(v bool) *GetUserWatchLiveListResponseBodyResult {
	s.HasFinish = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResult) SetLiveInfoPopModelList(v []*GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList) *GetUserWatchLiveListResponseBodyResult {
	s.LiveInfoPopModelList = v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResult) SetNextToken(v string) *GetUserWatchLiveListResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResult) SetTotal(v int32) *GetUserWatchLiveListResponseBodyResult {
	s.Total = &v
	return s
}

type GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList struct {
	ExtraInfo     *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo     `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	LiveBasicInfo *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo `json:"liveBasicInfo,omitempty" xml:"liveBasicInfo,omitempty" type:"Struct"`
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList) SetExtraInfo(v *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList {
	s.ExtraInfo = v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList) SetLiveBasicInfo(v *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelList {
	s.LiveBasicInfo = v
	return s
}

type GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo struct {
	HasSubscribed     *bool  `json:"hasSubscribed,omitempty" xml:"hasSubscribed,omitempty"`
	IsForecastExpired *bool  `json:"isForecastExpired,omitempty" xml:"isForecastExpired,omitempty"`
	WatchProgressMs   *int64 `json:"watchProgressMs,omitempty" xml:"watchProgressMs,omitempty"`
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetHasSubscribed(v bool) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.HasSubscribed = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetIsForecastExpired(v bool) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.IsForecastExpired = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo) SetWatchProgressMs(v int64) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListExtraInfo {
	s.WatchProgressMs = &v
	return s
}

type GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo struct {
	CoverUrl       *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	Duration       *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime        *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Introduction   *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	LiveId         *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	LivePlayUrl    *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	LiveStatus     *int32  `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	StartTime      *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SubscribeCount *int32  `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty"`
	UnionId        *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	Uv             *int32  `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetCoverUrl(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.CoverUrl = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetDuration(v int64) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetEndTime(v int64) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.EndTime = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetIntroduction(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Introduction = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveId(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveId = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLivePlayUrl(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LivePlayUrl = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetLiveStatus(v int32) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.LiveStatus = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetStartTime(v int64) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.StartTime = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetSubscribeCount(v int32) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.SubscribeCount = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetTitle(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Title = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUnionId(v string) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.UnionId = &v
	return s
}

func (s *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo) SetUv(v int32) *GetUserWatchLiveListResponseBodyResultLiveInfoPopModelListLiveBasicInfo {
	s.Uv = &v
	return s
}

type GetUserWatchLiveListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserWatchLiveListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserWatchLiveListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserWatchLiveListResponse) GoString() string {
	return s.String()
}

func (s *GetUserWatchLiveListResponse) SetHeaders(v map[string]*string) *GetUserWatchLiveListResponse {
	s.Headers = v
	return s
}

func (s *GetUserWatchLiveListResponse) SetStatusCode(v int32) *GetUserWatchLiveListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserWatchLiveListResponse) SetBody(v *GetUserWatchLiveListResponseBody) *GetUserWatchLiveListResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	Action         *int64    `json:"action,omitempty" xml:"action,omitempty"`
	ModifyUserList []*string `json:"modifyUserList,omitempty" xml:"modifyUserList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ModifyFeedWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListRequest) SetAction(v int64) *ModifyFeedWhiteListRequest {
	s.Action = &v
	return s
}

func (s *ModifyFeedWhiteListRequest) SetModifyUserList(v []*string) *ModifyFeedWhiteListRequest {
	s.ModifyUserList = v
	return s
}

func (s *ModifyFeedWhiteListRequest) SetUserId(v string) *ModifyFeedWhiteListRequest {
	s.UserId = &v
	return s
}

type ModifyFeedWhiteListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Action               *int64  `json:"action,omitempty" xml:"action,omitempty"`
	ModifyUserListShrink *string `json:"modifyUserList,omitempty" xml:"modifyUserList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ModifyFeedWhiteListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeedWhiteListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeedWhiteListShrinkRequest) SetAction(v int64) *ModifyFeedWhiteListShrinkRequest {
	s.Action = &v
	return s
}

func (s *ModifyFeedWhiteListShrinkRequest) SetModifyUserListShrink(v string) *ModifyFeedWhiteListShrinkRequest {
	s.ModifyUserListShrink = &v
	return s
}

func (s *ModifyFeedWhiteListShrinkRequest) SetUserId(v string) *ModifyFeedWhiteListShrinkRequest {
	s.UserId = &v
	return s
}

type ModifyFeedWhiteListResponseBody struct {
	// example:
	//
	// true
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFeedWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyFeedWhiteListResponse) SetStatusCode(v int32) *ModifyFeedWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFeedWhiteListResponse) SetBody(v *ModifyFeedWhiteListResponseBody) *ModifyFeedWhiteListResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
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
	// example:
	//
	// true
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFeedWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryFeedWhiteListResponse) SetStatusCode(v int32) *QueryFeedWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFeedWhiteListResponse) SetBody(v *QueryFeedWhiteListResponseBody) *QueryFeedWhiteListResponse {
	s.Body = v
	return s
}

type QueryLiveInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryLiveInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryLiveInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryLiveInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4d383876-1ff9-4b73-a057-a8f47b346ecb
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryLiveInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoRequest) SetLiveId(v string) *QueryLiveInfoRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInfoRequest) SetUnionId(v string) *QueryLiveInfoRequest {
	s.UnionId = &v
	return s
}

type QueryLiveInfoResponseBody struct {
	Result *QueryLiveInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryLiveInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponseBody) SetResult(v *QueryLiveInfoResponseBodyResult) *QueryLiveInfoResponseBody {
	s.Result = v
	return s
}

type QueryLiveInfoResponseBodyResult struct {
	LiveInfo *QueryLiveInfoResponseBodyResultLiveInfo `json:"liveInfo,omitempty" xml:"liveInfo,omitempty" type:"Struct"`
}

func (s QueryLiveInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponseBodyResult) SetLiveInfo(v *QueryLiveInfoResponseBodyResultLiveInfo) *QueryLiveInfoResponseBodyResult {
	s.LiveInfo = v
	return s
}

type QueryLiveInfoResponseBodyResultLiveInfo struct {
	// example:
	//
	// https://gw.alicdn.com/tfs/TB1thlYyAT2gK0jSZPcXXcKkpXa-1125-633.png
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 18450
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1659653648000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 测试直播简介
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// example:
	//
	// 1a353547-040d-4095-bb93-404bc5d47920
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// example:
	//
	// https://h5.dingtalk.com/group-live-share/index.htm?type=2&liveFromType=6&liveUuid=1a353547-040d-4095-bb93-404bc5d47920&dd_nav_bgcolor=FF2C2D2F#/union
	LivePlayUrl *string `json:"livePlayUrl,omitempty" xml:"livePlayUrl,omitempty"`
	// example:
	//
	// 3
	LiveStatus *int32 `json:"liveStatus,omitempty" xml:"liveStatus,omitempty"`
	// example:
	//
	// 18430
	PlaybackDuration *int64 `json:"playbackDuration,omitempty" xml:"playbackDuration,omitempty"`
	// example:
	//
	// 1659613648000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2
	SubscribeCount *int32 `json:"subscribeCount,omitempty" xml:"subscribeCount,omitempty"`
	// example:
	//
	// 测试直播
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 3
	Uv *int32 `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s QueryLiveInfoResponseBodyResultLiveInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoResponseBodyResultLiveInfo) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetCoverUrl(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.CoverUrl = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetDuration(v int64) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.Duration = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetEndTime(v int64) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.EndTime = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetIntroduction(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.Introduction = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetLiveId(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetLivePlayUrl(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.LivePlayUrl = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetLiveStatus(v int32) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.LiveStatus = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetPlaybackDuration(v int64) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.PlaybackDuration = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetStartTime(v int64) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.StartTime = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetSubscribeCount(v int32) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.SubscribeCount = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetTitle(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.Title = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetUnionId(v string) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.UnionId = &v
	return s
}

func (s *QueryLiveInfoResponseBodyResultLiveInfo) SetUv(v int32) *QueryLiveInfoResponseBodyResultLiveInfo {
	s.Uv = &v
	return s
}

type QueryLiveInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponse) SetHeaders(v map[string]*string) *QueryLiveInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveInfoResponse) SetStatusCode(v int32) *QueryLiveInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveInfoResponse) SetBody(v *QueryLiveInfoResponseBody) *QueryLiveInfoResponse {
	s.Body = v
	return s
}

type QueryLiveInteractionPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryLiveInteractionPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveInteractionPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveInteractionPluginHeaders) SetXAcsDingtalkAccessToken(v string) *QueryLiveInteractionPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryLiveInteractionPluginRequest struct {
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryLiveInteractionPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginRequest) SetLiveId(v string) *QueryLiveInteractionPluginRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveInteractionPluginRequest) SetPluginId(v string) *QueryLiveInteractionPluginRequest {
	s.PluginId = &v
	return s
}

func (s *QueryLiveInteractionPluginRequest) SetUnionId(v string) *QueryLiveInteractionPluginRequest {
	s.UnionId = &v
	return s
}

type QueryLiveInteractionPluginResponseBody struct {
	Result *QueryLiveInteractionPluginResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryLiveInteractionPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginResponseBody) SetResult(v *QueryLiveInteractionPluginResponseBodyResult) *QueryLiveInteractionPluginResponseBody {
	s.Result = v
	return s
}

type QueryLiveInteractionPluginResponseBodyResult struct {
	LivePluginList []*QueryLiveInteractionPluginResponseBodyResultLivePluginList `json:"livePluginList,omitempty" xml:"livePluginList,omitempty" type:"Repeated"`
}

func (s QueryLiveInteractionPluginResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginResponseBodyResult) SetLivePluginList(v []*QueryLiveInteractionPluginResponseBodyResultLivePluginList) *QueryLiveInteractionPluginResponseBodyResult {
	s.LivePluginList = v
	return s
}

type QueryLiveInteractionPluginResponseBodyResultLivePluginList struct {
	AnchorJumpUrl *string `json:"anchorJumpUrl,omitempty" xml:"anchorJumpUrl,omitempty"`
	PluginIconUrl *string `json:"pluginIconUrl,omitempty" xml:"pluginIconUrl,omitempty"`
	PluginId      *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	PluginName    *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	PluginNameEn  *string `json:"pluginNameEn,omitempty" xml:"pluginNameEn,omitempty"`
	ViewerJumpUrl *string `json:"viewerJumpUrl,omitempty" xml:"viewerJumpUrl,omitempty"`
}

func (s QueryLiveInteractionPluginResponseBodyResultLivePluginList) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginResponseBodyResultLivePluginList) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetAnchorJumpUrl(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.AnchorJumpUrl = &v
	return s
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetPluginIconUrl(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.PluginIconUrl = &v
	return s
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetPluginId(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.PluginId = &v
	return s
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetPluginName(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.PluginName = &v
	return s
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetPluginNameEn(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.PluginNameEn = &v
	return s
}

func (s *QueryLiveInteractionPluginResponseBodyResultLivePluginList) SetViewerJumpUrl(v string) *QueryLiveInteractionPluginResponseBodyResultLivePluginList {
	s.ViewerJumpUrl = &v
	return s
}

type QueryLiveInteractionPluginResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveInteractionPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveInteractionPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveInteractionPluginResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveInteractionPluginResponse) SetHeaders(v map[string]*string) *QueryLiveInteractionPluginResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveInteractionPluginResponse) SetStatusCode(v int32) *QueryLiveInteractionPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveInteractionPluginResponse) SetBody(v *QueryLiveInteractionPluginResponseBody) *QueryLiveInteractionPluginResponse {
	s.Body = v
	return s
}

type QueryLiveWatchDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryLiveWatchDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchDetailHeaders) SetXAcsDingtalkAccessToken(v string) *QueryLiveWatchDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryLiveWatchDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1a353547-040d-4095-bb93-404bc5d47920
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryLiveWatchDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailRequest) SetLiveId(v string) *QueryLiveWatchDetailRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchDetailRequest) SetUnionId(v string) *QueryLiveWatchDetailRequest {
	s.UnionId = &v
	return s
}

type QueryLiveWatchDetailResponseBody struct {
	Result *QueryLiveWatchDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryLiveWatchDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailResponseBody) SetResult(v *QueryLiveWatchDetailResponseBodyResult) *QueryLiveWatchDetailResponseBody {
	s.Result = v
	return s
}

type QueryLiveWatchDetailResponseBodyResult struct {
	// example:
	//
	// 3560
	AvgWatchTime *int64 `json:"avgWatchTime,omitempty" xml:"avgWatchTime,omitempty"`
	// example:
	//
	// 55
	LiveUv *int32 `json:"liveUv,omitempty" xml:"liveUv,omitempty"`
	// example:
	//
	// 252
	MsgCount *int32 `json:"msgCount,omitempty" xml:"msgCount,omitempty"`
	// example:
	//
	// 72
	PlaybackUv *int32 `json:"playbackUv,omitempty" xml:"playbackUv,omitempty"`
	// example:
	//
	// 500
	PraiseCount *int32 `json:"praiseCount,omitempty" xml:"praiseCount,omitempty"`
	// example:
	//
	// 120
	Pv *int32 `json:"pv,omitempty" xml:"pv,omitempty"`
	// example:
	//
	// 1903640
	TotalWatchTime *int64 `json:"totalWatchTime,omitempty" xml:"totalWatchTime,omitempty"`
	// example:
	//
	// 90
	Uv *int32 `json:"uv,omitempty" xml:"uv,omitempty"`
}

func (s QueryLiveWatchDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetAvgWatchTime(v int64) *QueryLiveWatchDetailResponseBodyResult {
	s.AvgWatchTime = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetLiveUv(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.LiveUv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetMsgCount(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.MsgCount = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetPlaybackUv(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.PlaybackUv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetPraiseCount(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.PraiseCount = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetPv(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetTotalWatchTime(v int64) *QueryLiveWatchDetailResponseBodyResult {
	s.TotalWatchTime = &v
	return s
}

func (s *QueryLiveWatchDetailResponseBodyResult) SetUv(v int32) *QueryLiveWatchDetailResponseBodyResult {
	s.Uv = &v
	return s
}

type QueryLiveWatchDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveWatchDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveWatchDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailResponse) SetHeaders(v map[string]*string) *QueryLiveWatchDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveWatchDetailResponse) SetStatusCode(v int32) *QueryLiveWatchDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveWatchDetailResponse) SetBody(v *QueryLiveWatchDetailResponseBody) *QueryLiveWatchDetailResponse {
	s.Body = v
	return s
}

type QueryLiveWatchUserListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryLiveWatchUserListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchUserListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchUserListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryLiveWatchUserListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryLiveWatchUserListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1a353547-040d-4095-bb93-404bc5d47920
	LiveId     *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryLiveWatchUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListRequest) SetLiveId(v string) *QueryLiveWatchUserListRequest {
	s.LiveId = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetPageNumber(v int32) *QueryLiveWatchUserListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetPageSize(v int32) *QueryLiveWatchUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryLiveWatchUserListRequest) SetUnionId(v string) *QueryLiveWatchUserListRequest {
	s.UnionId = &v
	return s
}

type QueryLiveWatchUserListResponseBody struct {
	Result *QueryLiveWatchUserListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryLiveWatchUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBody) SetResult(v *QueryLiveWatchUserListResponseBodyResult) *QueryLiveWatchUserListResponseBody {
	s.Result = v
	return s
}

type QueryLiveWatchUserListResponseBodyResult struct {
	OrgUsesList    []*QueryLiveWatchUserListResponseBodyResultOrgUsesList    `json:"orgUsesList,omitempty" xml:"orgUsesList,omitempty" type:"Repeated"`
	OutOrgUserList []*QueryLiveWatchUserListResponseBodyResultOutOrgUserList `json:"outOrgUserList,omitempty" xml:"outOrgUserList,omitempty" type:"Repeated"`
}

func (s QueryLiveWatchUserListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBodyResult) SetOrgUsesList(v []*QueryLiveWatchUserListResponseBodyResultOrgUsesList) *QueryLiveWatchUserListResponseBodyResult {
	s.OrgUsesList = v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResult) SetOutOrgUserList(v []*QueryLiveWatchUserListResponseBodyResultOutOrgUserList) *QueryLiveWatchUserListResponseBodyResult {
	s.OutOrgUserList = v
	return s
}

type QueryLiveWatchUserListResponseBodyResultOrgUsesList struct {
	// example:
	//
	// xxx.设计部
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 李四
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// 214675
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 189930
	WatchLiveTime *int64 `json:"watchLiveTime,omitempty" xml:"watchLiveTime,omitempty"`
	// example:
	//
	// 23667
	WatchPlaybackTime *int64 `json:"watchPlaybackTime,omitempty" xml:"watchPlaybackTime,omitempty"`
	// example:
	//
	// 2330
	WatchProgressMs *int64 `json:"watchProgressMs,omitempty" xml:"watchProgressMs,omitempty"`
}

func (s QueryLiveWatchUserListResponseBodyResultOrgUsesList) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBodyResultOrgUsesList) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetDeptName(v string) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.DeptName = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetName(v string) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.Name = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetUnionId(v string) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.UnionId = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetUserId(v string) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.UserId = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetWatchLiveTime(v int64) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.WatchLiveTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetWatchPlaybackTime(v int64) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.WatchPlaybackTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOrgUsesList) SetWatchProgressMs(v int64) *QueryLiveWatchUserListResponseBodyResultOrgUsesList {
	s.WatchProgressMs = &v
	return s
}

type QueryLiveWatchUserListResponseBodyResultOutOrgUserList struct {
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 23440
	WatchLiveTime *int64 `json:"watchLiveTime,omitempty" xml:"watchLiveTime,omitempty"`
	// example:
	//
	// 2330
	WatchPlaybackTime *int64 `json:"watchPlaybackTime,omitempty" xml:"watchPlaybackTime,omitempty"`
	// example:
	//
	// 150
	WatchProgressMs *int64 `json:"watchProgressMs,omitempty" xml:"watchProgressMs,omitempty"`
}

func (s QueryLiveWatchUserListResponseBodyResultOutOrgUserList) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListResponseBodyResultOutOrgUserList) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponseBodyResultOutOrgUserList) SetName(v string) *QueryLiveWatchUserListResponseBodyResultOutOrgUserList {
	s.Name = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOutOrgUserList) SetWatchLiveTime(v int64) *QueryLiveWatchUserListResponseBodyResultOutOrgUserList {
	s.WatchLiveTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOutOrgUserList) SetWatchPlaybackTime(v int64) *QueryLiveWatchUserListResponseBodyResultOutOrgUserList {
	s.WatchPlaybackTime = &v
	return s
}

func (s *QueryLiveWatchUserListResponseBodyResultOutOrgUserList) SetWatchProgressMs(v int64) *QueryLiveWatchUserListResponseBodyResultOutOrgUserList {
	s.WatchProgressMs = &v
	return s
}

type QueryLiveWatchUserListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveWatchUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveWatchUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWatchUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponse) SetHeaders(v map[string]*string) *QueryLiveWatchUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveWatchUserListResponse) SetStatusCode(v int32) *QueryLiveWatchUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveWatchUserListResponse) SetBody(v *QueryLiveWatchUserListResponseBody) *QueryLiveWatchUserListResponse {
	s.Body = v
	return s
}

type QuerySubscribeStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySubscribeStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusHeaders) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusHeaders) SetCommonHeaders(v map[string]*string) *QuerySubscribeStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySubscribeStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySubscribeStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySubscribeStatusRequest struct {
	Body    *QuerySubscribeStatusRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	UnionId *string                          `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QuerySubscribeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusRequest) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusRequest) SetBody(v *QuerySubscribeStatusRequestBody) *QuerySubscribeStatusRequest {
	s.Body = v
	return s
}

func (s *QuerySubscribeStatusRequest) SetUnionId(v string) *QuerySubscribeStatusRequest {
	s.UnionId = &v
	return s
}

type QuerySubscribeStatusRequestBody struct {
	LiveIds []*string `json:"liveIds,omitempty" xml:"liveIds,omitempty" type:"Repeated"`
}

func (s QuerySubscribeStatusRequestBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusRequestBody) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusRequestBody) SetLiveIds(v []*string) *QuerySubscribeStatusRequestBody {
	s.LiveIds = v
	return s
}

type QuerySubscribeStatusShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QuerySubscribeStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusShrinkRequest) SetBodyShrink(v string) *QuerySubscribeStatusShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *QuerySubscribeStatusShrinkRequest) SetUnionId(v string) *QuerySubscribeStatusShrinkRequest {
	s.UnionId = &v
	return s
}

type QuerySubscribeStatusResponseBody struct {
	Result *QuerySubscribeStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QuerySubscribeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusResponseBody) SetResult(v *QuerySubscribeStatusResponseBodyResult) *QuerySubscribeStatusResponseBody {
	s.Result = v
	return s
}

type QuerySubscribeStatusResponseBodyResult struct {
	SubscribeStatusDTOS []*QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS `json:"subscribeStatusDTOS,omitempty" xml:"subscribeStatusDTOS,omitempty" type:"Repeated"`
}

func (s QuerySubscribeStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusResponseBodyResult) SetSubscribeStatusDTOS(v []*QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS) *QuerySubscribeStatusResponseBodyResult {
	s.SubscribeStatusDTOS = v
	return s
}

type QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS struct {
	LiveId    *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	Subscribe *bool   `json:"subscribe,omitempty" xml:"subscribe,omitempty"`
}

func (s QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS) SetLiveId(v string) *QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS {
	s.LiveId = &v
	return s
}

func (s *QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS) SetSubscribe(v bool) *QuerySubscribeStatusResponseBodyResultSubscribeStatusDTOS {
	s.Subscribe = &v
	return s
}

type QuerySubscribeStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySubscribeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySubscribeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySubscribeStatusResponse) GoString() string {
	return s.String()
}

func (s *QuerySubscribeStatusResponse) SetHeaders(v map[string]*string) *QuerySubscribeStatusResponse {
	s.Headers = v
	return s
}

func (s *QuerySubscribeStatusResponse) SetStatusCode(v int32) *QuerySubscribeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubscribeStatusResponse) SetBody(v *QuerySubscribeStatusResponseBody) *QuerySubscribeStatusResponse {
	s.Body = v
	return s
}

type SendLiveInteractionPluginEffectsMsgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendLiveInteractionPluginEffectsMsgHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendLiveInteractionPluginEffectsMsgHeaders) GoString() string {
	return s.String()
}

func (s *SendLiveInteractionPluginEffectsMsgHeaders) SetCommonHeaders(v map[string]*string) *SendLiveInteractionPluginEffectsMsgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgHeaders) SetXAcsDingtalkAccessToken(v string) *SendLiveInteractionPluginEffectsMsgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendLiveInteractionPluginEffectsMsgRequest struct {
	// This parameter is required.
	Count         *int64  `json:"count,omitempty" xml:"count,omitempty"`
	LottieFileUrl *string `json:"lottieFileUrl,omitempty" xml:"lottieFileUrl,omitempty"`
	// This parameter is required.
	MsgIconUrl *string `json:"msgIconUrl,omitempty" xml:"msgIconUrl,omitempty"`
	// This parameter is required.
	MsgText *string `json:"msgText,omitempty" xml:"msgText,omitempty"`
	// This parameter is required.
	PluginSubId *string `json:"pluginSubId,omitempty" xml:"pluginSubId,omitempty"`
	// This parameter is required.
	SenderUnionId *string `json:"senderUnionId,omitempty" xml:"senderUnionId,omitempty"`
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s SendLiveInteractionPluginEffectsMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendLiveInteractionPluginEffectsMsgRequest) GoString() string {
	return s.String()
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetCount(v int64) *SendLiveInteractionPluginEffectsMsgRequest {
	s.Count = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetLottieFileUrl(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.LottieFileUrl = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetMsgIconUrl(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.MsgIconUrl = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetMsgText(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.MsgText = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetPluginSubId(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.PluginSubId = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetSenderUnionId(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.SenderUnionId = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetLiveId(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.LiveId = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgRequest) SetPluginId(v string) *SendLiveInteractionPluginEffectsMsgRequest {
	s.PluginId = &v
	return s
}

type SendLiveInteractionPluginEffectsMsgResponseBody struct {
	Result *SendLiveInteractionPluginEffectsMsgResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendLiveInteractionPluginEffectsMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendLiveInteractionPluginEffectsMsgResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveInteractionPluginEffectsMsgResponseBody) SetResult(v *SendLiveInteractionPluginEffectsMsgResponseBodyResult) *SendLiveInteractionPluginEffectsMsgResponseBody {
	s.Result = v
	return s
}

type SendLiveInteractionPluginEffectsMsgResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendLiveInteractionPluginEffectsMsgResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendLiveInteractionPluginEffectsMsgResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendLiveInteractionPluginEffectsMsgResponseBodyResult) SetSuccess(v bool) *SendLiveInteractionPluginEffectsMsgResponseBodyResult {
	s.Success = &v
	return s
}

type SendLiveInteractionPluginEffectsMsgResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLiveInteractionPluginEffectsMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLiveInteractionPluginEffectsMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendLiveInteractionPluginEffectsMsgResponse) GoString() string {
	return s.String()
}

func (s *SendLiveInteractionPluginEffectsMsgResponse) SetHeaders(v map[string]*string) *SendLiveInteractionPluginEffectsMsgResponse {
	s.Headers = v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgResponse) SetStatusCode(v int32) *SendLiveInteractionPluginEffectsMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLiveInteractionPluginEffectsMsgResponse) SetBody(v *SendLiveInteractionPluginEffectsMsgResponseBody) *SendLiveInteractionPluginEffectsMsgResponse {
	s.Body = v
	return s
}

type SendLivePluginUserActionMsgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendLivePluginUserActionMsgHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgHeaders) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgHeaders) SetCommonHeaders(v map[string]*string) *SendLivePluginUserActionMsgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendLivePluginUserActionMsgHeaders) SetXAcsDingtalkAccessToken(v string) *SendLivePluginUserActionMsgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendLivePluginUserActionMsgRequest struct {
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginEffectsMessage *SendLivePluginUserActionMsgRequestPluginEffectsMessage `json:"pluginEffectsMessage,omitempty" xml:"pluginEffectsMessage,omitempty" type:"Struct"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s SendLivePluginUserActionMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgRequest) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgRequest) SetLiveId(v string) *SendLivePluginUserActionMsgRequest {
	s.LiveId = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequest) SetPluginEffectsMessage(v *SendLivePluginUserActionMsgRequestPluginEffectsMessage) *SendLivePluginUserActionMsgRequest {
	s.PluginEffectsMessage = v
	return s
}

func (s *SendLivePluginUserActionMsgRequest) SetPluginId(v string) *SendLivePluginUserActionMsgRequest {
	s.PluginId = &v
	return s
}

type SendLivePluginUserActionMsgRequestPluginEffectsMessage struct {
	// This parameter is required.
	Count         *int64  `json:"count,omitempty" xml:"count,omitempty"`
	LottieFileUrl *string `json:"lottieFileUrl,omitempty" xml:"lottieFileUrl,omitempty"`
	// This parameter is required.
	MsgIconUrl *string `json:"msgIconUrl,omitempty" xml:"msgIconUrl,omitempty"`
	// This parameter is required.
	MsgText *string `json:"msgText,omitempty" xml:"msgText,omitempty"`
	// This parameter is required.
	PluginSubId *string `json:"pluginSubId,omitempty" xml:"pluginSubId,omitempty"`
	// This parameter is required.
	SenderUnionId *string `json:"senderUnionId,omitempty" xml:"senderUnionId,omitempty"`
}

func (s SendLivePluginUserActionMsgRequestPluginEffectsMessage) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgRequestPluginEffectsMessage) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetCount(v int64) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.Count = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetLottieFileUrl(v string) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.LottieFileUrl = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetMsgIconUrl(v string) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.MsgIconUrl = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetMsgText(v string) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.MsgText = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetPluginSubId(v string) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.PluginSubId = &v
	return s
}

func (s *SendLivePluginUserActionMsgRequestPluginEffectsMessage) SetSenderUnionId(v string) *SendLivePluginUserActionMsgRequestPluginEffectsMessage {
	s.SenderUnionId = &v
	return s
}

type SendLivePluginUserActionMsgShrinkRequest struct {
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginEffectsMessageShrink *string `json:"pluginEffectsMessage,omitempty" xml:"pluginEffectsMessage,omitempty"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s SendLivePluginUserActionMsgShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgShrinkRequest) SetLiveId(v string) *SendLivePluginUserActionMsgShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *SendLivePluginUserActionMsgShrinkRequest) SetPluginEffectsMessageShrink(v string) *SendLivePluginUserActionMsgShrinkRequest {
	s.PluginEffectsMessageShrink = &v
	return s
}

func (s *SendLivePluginUserActionMsgShrinkRequest) SetPluginId(v string) *SendLivePluginUserActionMsgShrinkRequest {
	s.PluginId = &v
	return s
}

type SendLivePluginUserActionMsgResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendLivePluginUserActionMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgResponseBody) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgResponseBody) SetSuccess(v bool) *SendLivePluginUserActionMsgResponseBody {
	s.Success = &v
	return s
}

type SendLivePluginUserActionMsgResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendLivePluginUserActionMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendLivePluginUserActionMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendLivePluginUserActionMsgResponse) GoString() string {
	return s.String()
}

func (s *SendLivePluginUserActionMsgResponse) SetHeaders(v map[string]*string) *SendLivePluginUserActionMsgResponse {
	s.Headers = v
	return s
}

func (s *SendLivePluginUserActionMsgResponse) SetStatusCode(v int32) *SendLivePluginUserActionMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *SendLivePluginUserActionMsgResponse) SetBody(v *SendLivePluginUserActionMsgResponseBody) *SendLivePluginUserActionMsgResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 214675
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
	// example:
	//
	// true
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartCloudFeedResponse) SetStatusCode(v int32) *StartCloudFeedResponse {
	s.StatusCode = &v
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
	// This parameter is required.
	//
	// example:
	//
	// 214675
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
	// example:
	//
	// true
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCloudFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopCloudFeedResponse) SetStatusCode(v int32) *StopCloudFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCloudFeedResponse) SetBody(v *StopCloudFeedResponseBody) *StopCloudFeedResponse {
	s.Body = v
	return s
}

type SubscribeLiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SubscribeLiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubscribeLiveHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeLiveHeaders) SetCommonHeaders(v map[string]*string) *SubscribeLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeLiveHeaders) SetXAcsDingtalkAccessToken(v string) *SubscribeLiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SubscribeLiveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3dd295eb-17a1-4dfg-ae1b-aa165c5007eb
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	Subscribe *bool `json:"subscribe,omitempty" xml:"subscribe,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6crtQT2XOgPHviiPvXhhiP6gdhiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s SubscribeLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeLiveRequest) GoString() string {
	return s.String()
}

func (s *SubscribeLiveRequest) SetLiveId(v string) *SubscribeLiveRequest {
	s.LiveId = &v
	return s
}

func (s *SubscribeLiveRequest) SetSubscribe(v bool) *SubscribeLiveRequest {
	s.Subscribe = &v
	return s
}

func (s *SubscribeLiveRequest) SetUnionId(v string) *SubscribeLiveRequest {
	s.UnionId = &v
	return s
}

type SubscribeLiveResponseBody struct {
	Result *SubscribeLiveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SubscribeLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeLiveResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeLiveResponseBody) SetResult(v *SubscribeLiveResponseBodyResult) *SubscribeLiveResponseBody {
	s.Result = v
	return s
}

type SubscribeLiveResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubscribeLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SubscribeLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SubscribeLiveResponseBodyResult) SetSuccess(v bool) *SubscribeLiveResponseBodyResult {
	s.Success = &v
	return s
}

type SubscribeLiveResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeLiveResponse) GoString() string {
	return s.String()
}

func (s *SubscribeLiveResponse) SetHeaders(v map[string]*string) *SubscribeLiveResponse {
	s.Headers = v
	return s
}

func (s *SubscribeLiveResponse) SetStatusCode(v int32) *SubscribeLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeLiveResponse) SetBody(v *SubscribeLiveResponseBody) *SubscribeLiveResponse {
	s.Body = v
	return s
}

type UpdateLiveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateLiveHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLiveHeaders) SetCommonHeaders(v map[string]*string) *UpdateLiveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLiveHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateLiveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateLiveRequest struct {
	// example:
	//
	// https://gw.alicdn.com/tfs/TB1thlYyAT2gK0jSZPcXXcKkpXa-1125-633.png
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 测试直播简介
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4d383876-1ff9-4b73-a057-a8f47b346ecb
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// example:
	//
	// 1659653648000
	PreEndTime *int64 `json:"preEndTime,omitempty" xml:"preEndTime,omitempty"`
	// example:
	//
	// 1659613648000
	PreStartTime *int64 `json:"preStartTime,omitempty" xml:"preStartTime,omitempty"`
	// example:
	//
	// 测试直播
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DC7wZGOSueEEIGOf3WKwWgiEiE
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) SetCoverUrl(v string) *UpdateLiveRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRequest) SetIntroduction(v string) *UpdateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetPreEndTime(v int64) *UpdateLiveRequest {
	s.PreEndTime = &v
	return s
}

func (s *UpdateLiveRequest) SetPreStartTime(v int64) *UpdateLiveRequest {
	s.PreStartTime = &v
	return s
}

func (s *UpdateLiveRequest) SetTitle(v string) *UpdateLiveRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRequest) SetUnionId(v string) *UpdateLiveRequest {
	s.UnionId = &v
	return s
}

type UpdateLiveResponseBody struct {
	Result *UpdateLiveResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBody) SetResult(v *UpdateLiveResponseBodyResult) *UpdateLiveResponseBody {
	s.Result = v
	return s
}

type UpdateLiveResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBodyResult) SetSuccess(v bool) *UpdateLiveResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponse) SetHeaders(v map[string]*string) *UpdateLiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveResponse) SetStatusCode(v int32) *UpdateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveResponse) SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse {
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
	// example:
	//
	// http:xxx.png
	CoverUrl *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	// example:
	//
	// 简介
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// example:
	//
	// 1617436058000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1206186351746728
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateLiveFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveFeedRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveFeedRequest) SetCoverUrl(v string) *UpdateLiveFeedRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetIntroduction(v string) *UpdateLiveFeedRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetStartTime(v int64) *UpdateLiveFeedRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetTitle(v string) *UpdateLiveFeedRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveFeedRequest) SetUserId(v string) *UpdateLiveFeedRequest {
	s.UserId = &v
	return s
}

type UpdateLiveFeedResponseBody struct {
	// example:
	//
	// true
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateLiveFeedResponse) SetStatusCode(v int32) *UpdateLiveFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveFeedResponse) SetBody(v *UpdateLiveFeedResponseBody) *UpdateLiveFeedResponse {
	s.Body = v
	return s
}

type UpdateLiveInteractionPluginHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateLiveInteractionPluginHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveInteractionPluginHeaders) GoString() string {
	return s.String()
}

func (s *UpdateLiveInteractionPluginHeaders) SetCommonHeaders(v map[string]*string) *UpdateLiveInteractionPluginHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateLiveInteractionPluginHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateLiveInteractionPluginHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateLiveInteractionPluginRequest struct {
	AnchorJumpUrl *string `json:"anchorJumpUrl,omitempty" xml:"anchorJumpUrl,omitempty"`
	PluginIconUrl *string `json:"pluginIconUrl,omitempty" xml:"pluginIconUrl,omitempty"`
	PluginName    *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	PluginNameEn  *string `json:"pluginNameEn,omitempty" xml:"pluginNameEn,omitempty"`
	ViewerJumpUrl *string `json:"viewerJumpUrl,omitempty" xml:"viewerJumpUrl,omitempty"`
	// This parameter is required.
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// This parameter is required.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s UpdateLiveInteractionPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveInteractionPluginRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveInteractionPluginRequest) SetAnchorJumpUrl(v string) *UpdateLiveInteractionPluginRequest {
	s.AnchorJumpUrl = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetPluginIconUrl(v string) *UpdateLiveInteractionPluginRequest {
	s.PluginIconUrl = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetPluginName(v string) *UpdateLiveInteractionPluginRequest {
	s.PluginName = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetPluginNameEn(v string) *UpdateLiveInteractionPluginRequest {
	s.PluginNameEn = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetViewerJumpUrl(v string) *UpdateLiveInteractionPluginRequest {
	s.ViewerJumpUrl = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetLiveId(v string) *UpdateLiveInteractionPluginRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetPluginId(v string) *UpdateLiveInteractionPluginRequest {
	s.PluginId = &v
	return s
}

func (s *UpdateLiveInteractionPluginRequest) SetUnionId(v string) *UpdateLiveInteractionPluginRequest {
	s.UnionId = &v
	return s
}

type UpdateLiveInteractionPluginResponseBody struct {
	Result *UpdateLiveInteractionPluginResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateLiveInteractionPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveInteractionPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveInteractionPluginResponseBody) SetResult(v *UpdateLiveInteractionPluginResponseBodyResult) *UpdateLiveInteractionPluginResponseBody {
	s.Result = v
	return s
}

type UpdateLiveInteractionPluginResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateLiveInteractionPluginResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveInteractionPluginResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateLiveInteractionPluginResponseBodyResult) SetSuccess(v bool) *UpdateLiveInteractionPluginResponseBodyResult {
	s.Success = &v
	return s
}

type UpdateLiveInteractionPluginResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveInteractionPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveInteractionPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveInteractionPluginResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveInteractionPluginResponse) SetHeaders(v map[string]*string) *UpdateLiveInteractionPluginResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveInteractionPluginResponse) SetStatusCode(v int32) *UpdateLiveInteractionPluginResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveInteractionPluginResponse) SetBody(v *UpdateLiveInteractionPluginResponseBody) *UpdateLiveInteractionPluginResponse {
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
// 增加直播间互动插件
//
// @param request - AddLiveInteractionPluginRequest
//
// @param headers - AddLiveInteractionPluginHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveInteractionPluginResponse
func (client *Client) AddLiveInteractionPluginWithOptions(request *AddLiveInteractionPluginRequest, headers *AddLiveInteractionPluginHeaders, runtime *util.RuntimeOptions) (_result *AddLiveInteractionPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorJumpUrl)) {
		body["anchorJumpUrl"] = request.AnchorJumpUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PluginIconUrl)) {
		body["pluginIconUrl"] = request.PluginIconUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PluginName)) {
		body["pluginName"] = request.PluginName
	}

	if !tea.BoolValue(util.IsUnset(request.PluginNameEn)) {
		body["pluginNameEn"] = request.PluginNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.ViewerJumpUrl)) {
		body["viewerJumpUrl"] = request.ViewerJumpUrl
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
	params := &openapi.Params{
		Action:      tea.String("AddLiveInteractionPlugin"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLiveInteractionPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加直播间互动插件
//
// @param request - AddLiveInteractionPluginRequest
//
// @return AddLiveInteractionPluginResponse
func (client *Client) AddLiveInteractionPlugin(request *AddLiveInteractionPluginRequest) (_result *AddLiveInteractionPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLiveInteractionPluginHeaders{}
	_result = &AddLiveInteractionPluginResponse{}
	_body, _err := client.AddLiveInteractionPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增加直播间的公告槽位信息
//
// @param request - AddLiveNoticeWidgetRequest
//
// @param headers - AddLiveNoticeWidgetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLiveNoticeWidgetResponse
func (client *Client) AddLiveNoticeWidgetWithOptions(request *AddLiveNoticeWidgetRequest, headers *AddLiveNoticeWidgetHeaders, runtime *util.RuntimeOptions) (_result *AddLiveNoticeWidgetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IconUrl)) {
		query["iconUrl"] = request.IconUrl
	}

	if !tea.BoolValue(util.IsUnset(request.JumpUrl)) {
		query["jumpUrl"] = request.JumpUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeText)) {
		query["noticeText"] = request.NoticeText
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
		Action:      tea.String("AddLiveNoticeWidget"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/noticeWidgets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddLiveNoticeWidgetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加直播间的公告槽位信息
//
// @param request - AddLiveNoticeWidgetRequest
//
// @return AddLiveNoticeWidgetResponse
func (client *Client) AddLiveNoticeWidget(request *AddLiveNoticeWidgetRequest) (_result *AddLiveNoticeWidgetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddLiveNoticeWidgetHeaders{}
	_result = &AddLiveNoticeWidgetResponse{}
	_body, _err := client.AddLiveNoticeWidgetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加云导播联播群列表
//
// @param request - AddShareCidListRequest
//
// @param headers - AddShareCidListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShareCidListResponse
func (client *Client) AddShareCidListWithOptions(feedId *string, request *AddShareCidListRequest, headers *AddShareCidListHeaders, runtime *util.RuntimeOptions) (_result *AddShareCidListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupIdType)) {
		body["groupIdType"] = request.GroupIdType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
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
		Action:      tea.String("AddShareCidList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/cloudFeeds/" + tea.StringValue(feedId) + "/share"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddShareCidListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加云导播联播群列表
//
// @param request - AddShareCidListRequest
//
// @return AddShareCidListResponse
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

// Summary:
//
// 创建云导播课程
//
// @param request - CreateCloudFeedRequest
//
// @param headers - CreateCloudFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudFeedResponse
func (client *Client) CreateCloudFeedWithOptions(request *CreateCloudFeedRequest, headers *CreateCloudFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateCloudFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Intro)) {
		body["intro"] = request.Intro
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["videoUrl"] = request.VideoUrl
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
		Action:      tea.String("CreateCloudFeed"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/cloudFeeds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCloudFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建云导播课程
//
// @param request - CreateCloudFeedRequest
//
// @return CreateCloudFeedResponse
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

// Summary:
//
// 创建直播
//
// @param request - CreateLiveRequest
//
// @param headers - CreateLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveResponse
func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, headers *CreateLiveHeaders, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.PreEndTime)) {
		body["preEndTime"] = request.PreEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PreStartTime)) {
		body["preStartTime"] = request.PreStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.PublicType)) {
		body["publicType"] = request.PublicType
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("CreateLive"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建直播
//
// @param request - CreateLiveRequest
//
// @return CreateLiveResponse
func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateLiveHeaders{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除直播间内某一互动插件
//
// @param request - DelLiveInteractionPluginRequest
//
// @param headers - DelLiveInteractionPluginHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelLiveInteractionPluginResponse
func (client *Client) DelLiveInteractionPluginWithOptions(request *DelLiveInteractionPluginRequest, headers *DelLiveInteractionPluginHeaders, runtime *util.RuntimeOptions) (_result *DelLiveInteractionPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["pluginId"] = request.PluginId
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
		Action:      tea.String("DelLiveInteractionPlugin"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DelLiveInteractionPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除直播间内某一互动插件
//
// @param request - DelLiveInteractionPluginRequest
//
// @return DelLiveInteractionPluginResponse
func (client *Client) DelLiveInteractionPlugin(request *DelLiveInteractionPluginRequest) (_result *DelLiveInteractionPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DelLiveInteractionPluginHeaders{}
	_result = &DelLiveInteractionPluginResponse{}
	_body, _err := client.DelLiveInteractionPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除直播间的公告槽位信息
//
// @param request - DelLiveNoticeWidgetRequest
//
// @param headers - DelLiveNoticeWidgetHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelLiveNoticeWidgetResponse
func (client *Client) DelLiveNoticeWidgetWithOptions(request *DelLiveNoticeWidgetRequest, headers *DelLiveNoticeWidgetHeaders, runtime *util.RuntimeOptions) (_result *DelLiveNoticeWidgetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
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
		Action:      tea.String("DelLiveNoticeWidget"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/noticeWidgets"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DelLiveNoticeWidgetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除直播间的公告槽位信息
//
// @param request - DelLiveNoticeWidgetRequest
//
// @return DelLiveNoticeWidgetResponse
func (client *Client) DelLiveNoticeWidget(request *DelLiveNoticeWidgetRequest) (_result *DelLiveNoticeWidgetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DelLiveNoticeWidgetHeaders{}
	_result = &DelLiveNoticeWidgetResponse{}
	_body, _err := client.DelLiveNoticeWidgetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除直播
//
// @param request - DeleteLiveRequest
//
// @param headers - DeleteLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveResponse
func (client *Client) DeleteLiveWithOptions(request *DeleteLiveRequest, headers *DeleteLiveHeaders, runtime *util.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
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
		Action:      tea.String("DeleteLive"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除直播
//
// @param request - DeleteLiveRequest
//
// @return DeleteLiveResponse
func (client *Client) DeleteLive(request *DeleteLiveRequest) (_result *DeleteLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteLiveHeaders{}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DeleteLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除直播培训课程
//
// @param request - DeleteLiveFeedRequest
//
// @param headers - DeleteLiveFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveFeedResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLiveFeed"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/openFeeds/" + tea.StringValue(feedId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除直播培训课程
//
// @param request - DeleteLiveFeedRequest
//
// @return DeleteLiveFeedResponse
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

// Summary:
//
// 剪辑直播课程的回放
//
// @param request - EditFeedReplayRequest
//
// @param headers - EditFeedReplayHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditFeedReplayResponse
func (client *Client) EditFeedReplayWithOptions(feedId *string, request *EditFeedReplayRequest, headers *EditFeedReplayHeaders, runtime *util.RuntimeOptions) (_result *EditFeedReplayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EditEndTime)) {
		body["editEndTime"] = request.EditEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EditStartTime)) {
		body["editStartTime"] = request.EditStartTime
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
		Action:      tea.String("EditFeedReplay"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/openFeeds/" + tea.StringValue(feedId) + "/cutReplay"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EditFeedReplayResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 剪辑直播课程的回放
//
// @param request - EditFeedReplayRequest
//
// @return EditFeedReplayResponse
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

// Summary:
//
// 获取群内的直播列表
//
// @param tmpReq - GetGroupLiveListRequest
//
// @param headers - GetGroupLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupLiveListResponse
func (client *Client) GetGroupLiveListWithOptions(tmpReq *GetGroupLiveListRequest, headers *GetGroupLiveListHeaders, runtime *util.RuntimeOptions) (_result *GetGroupLiveListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetGroupLiveListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RequestBody)) {
		request.RequestBodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestBody, tea.String("requestBody"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestBodyShrink)) {
		query["requestBody"] = request.RequestBodyShrink
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
		Action:      tea.String("GetGroupLiveList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/groupLives"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupLiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取群内的直播列表
//
// @param request - GetGroupLiveListRequest
//
// @return GetGroupLiveListResponse
func (client *Client) GetGroupLiveList(request *GetGroupLiveListRequest) (_result *GetGroupLiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupLiveListHeaders{}
	_result = &GetGroupLiveListResponse{}
	_body, _err := client.GetGroupLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取直播的可下载回放地址
//
// @param request - GetLiveReplayUrlRequest
//
// @param headers - GetLiveReplayUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveReplayUrlResponse
func (client *Client) GetLiveReplayUrlWithOptions(request *GetLiveReplayUrlRequest, headers *GetLiveReplayUrlHeaders, runtime *util.RuntimeOptions) (_result *GetLiveReplayUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
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
		Action:      tea.String("GetLiveReplayUrl"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives/replayUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveReplayUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播的可下载回放地址
//
// @param request - GetLiveReplayUrlRequest
//
// @return GetLiveReplayUrlResponse
func (client *Client) GetLiveReplayUrl(request *GetLiveReplayUrlRequest) (_result *GetLiveReplayUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLiveReplayUrlHeaders{}
	_result = &GetLiveReplayUrlResponse{}
	_body, _err := client.GetLiveReplayUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某组织内的直播列表
//
// @param tmpReq - GetOrgLiveListRequest
//
// @param headers - GetOrgLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgLiveListResponse
func (client *Client) GetOrgLiveListWithOptions(tmpReq *GetOrgLiveListRequest, headers *GetOrgLiveListHeaders, runtime *util.RuntimeOptions) (_result *GetOrgLiveListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetOrgLiveListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RequestBody)) {
		request.RequestBodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequestBody, tea.String("requestBody"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestBodyShrink)) {
		query["requestBody"] = request.RequestBodyShrink
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
		Action:      tea.String("GetOrgLiveList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/organizations/liveLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrgLiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某组织内的直播列表
//
// @param request - GetOrgLiveListRequest
//
// @return GetOrgLiveListResponse
func (client *Client) GetOrgLiveList(request *GetOrgLiveListRequest) (_result *GetOrgLiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrgLiveListHeaders{}
	_result = &GetOrgLiveListResponse{}
	_body, _err := client.GetOrgLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据状态拉我相关的直播
//
// @param request - GetUserAllLiveListRequest
//
// @param headers - GetUserAllLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAllLiveListResponse
func (client *Client) GetUserAllLiveListWithOptions(request *GetUserAllLiveListRequest, headers *GetUserAllLiveListHeaders, runtime *util.RuntimeOptions) (_result *GetUserAllLiveListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Statuses)) {
		body["statuses"] = request.Statuses
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserAllLiveList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/users/allLiveInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserAllLiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据状态拉我相关的直播
//
// @param request - GetUserAllLiveListRequest
//
// @return GetUserAllLiveListResponse
func (client *Client) GetUserAllLiveList(request *GetUserAllLiveListRequest) (_result *GetUserAllLiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserAllLiveListHeaders{}
	_result = &GetUserAllLiveListResponse{}
	_body, _err := client.GetUserAllLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据状态获取用户创建的直播
//
// @param request - GetUserCreateLiveListRequest
//
// @param headers - GetUserCreateLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCreateLiveListResponse
func (client *Client) GetUserCreateLiveListWithOptions(request *GetUserCreateLiveListRequest, headers *GetUserCreateLiveListHeaders, runtime *util.RuntimeOptions) (_result *GetUserCreateLiveListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Statuses)) {
		body["statuses"] = request.Statuses
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserCreateLiveList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/users/createLiveInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserCreateLiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据状态获取用户创建的直播
//
// @param request - GetUserCreateLiveListRequest
//
// @return GetUserCreateLiveListResponse
func (client *Client) GetUserCreateLiveList(request *GetUserCreateLiveListRequest) (_result *GetUserCreateLiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserCreateLiveListHeaders{}
	_result = &GetUserCreateLiveListResponse{}
	_body, _err := client.GetUserCreateLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户观看直播记录
//
// @param request - GetUserWatchLiveListRequest
//
// @param headers - GetUserWatchLiveListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserWatchLiveListResponse
func (client *Client) GetUserWatchLiveListWithOptions(request *GetUserWatchLiveListRequest, headers *GetUserWatchLiveListHeaders, runtime *util.RuntimeOptions) (_result *GetUserWatchLiveListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterType)) {
		query["filterType"] = request.FilterType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetUserWatchLiveList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/users/watchRecords"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserWatchLiveListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户观看直播记录
//
// @param request - GetUserWatchLiveListRequest
//
// @return GetUserWatchLiveListResponse
func (client *Client) GetUserWatchLiveList(request *GetUserWatchLiveListRequest) (_result *GetUserWatchLiveListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserWatchLiveListHeaders{}
	_result = &GetUserWatchLiveListResponse{}
	_body, _err := client.GetUserWatchLiveListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改直播课程可见白名单
//
// @param tmpReq - ModifyFeedWhiteListRequest
//
// @param headers - ModifyFeedWhiteListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFeedWhiteListResponse
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
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyUserListShrink)) {
		query["modifyUserList"] = request.ModifyUserListShrink
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
		Action:      tea.String("ModifyFeedWhiteList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/openFeeds/" + tea.StringValue(feedId) + "/whiteList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFeedWhiteListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改直播课程可见白名单
//
// @param request - ModifyFeedWhiteListRequest
//
// @return ModifyFeedWhiteListResponse
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

// Summary:
//
// 查询直播课程的观看白名单
//
// @param request - QueryFeedWhiteListRequest
//
// @param headers - QueryFeedWhiteListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFeedWhiteListResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFeedWhiteList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/openFeeds/" + tea.StringValue(feedId) + "/whiteList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFeedWhiteListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播课程的观看白名单
//
// @param request - QueryFeedWhiteListRequest
//
// @return QueryFeedWhiteListResponse
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

// Summary:
//
// 查询直播详情
//
// @param request - QueryLiveInfoRequest
//
// @param headers - QueryLiveInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveInfoResponse
func (client *Client) QueryLiveInfoWithOptions(request *QueryLiveInfoRequest, headers *QueryLiveInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryLiveInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
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
		Action:      tea.String("QueryLiveInfo"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLiveInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播详情
//
// @param request - QueryLiveInfoRequest
//
// @return QueryLiveInfoResponse
func (client *Client) QueryLiveInfo(request *QueryLiveInfoRequest) (_result *QueryLiveInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryLiveInfoHeaders{}
	_result = &QueryLiveInfoResponse{}
	_body, _err := client.QueryLiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询直播间内某一互动插件的信息
//
// @param request - QueryLiveInteractionPluginRequest
//
// @param headers - QueryLiveInteractionPluginHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveInteractionPluginResponse
func (client *Client) QueryLiveInteractionPluginWithOptions(request *QueryLiveInteractionPluginRequest, headers *QueryLiveInteractionPluginHeaders, runtime *util.RuntimeOptions) (_result *QueryLiveInteractionPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["pluginId"] = request.PluginId
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
		Action:      tea.String("QueryLiveInteractionPlugin"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLiveInteractionPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询直播间内某一互动插件的信息
//
// @param request - QueryLiveInteractionPluginRequest
//
// @return QueryLiveInteractionPluginResponse
func (client *Client) QueryLiveInteractionPlugin(request *QueryLiveInteractionPluginRequest) (_result *QueryLiveInteractionPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryLiveInteractionPluginHeaders{}
	_result = &QueryLiveInteractionPluginResponse{}
	_body, _err := client.QueryLiveInteractionPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取直播的观看数据
//
// @param request - QueryLiveWatchDetailRequest
//
// @param headers - QueryLiveWatchDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveWatchDetailResponse
func (client *Client) QueryLiveWatchDetailWithOptions(request *QueryLiveWatchDetailRequest, headers *QueryLiveWatchDetailHeaders, runtime *util.RuntimeOptions) (_result *QueryLiveWatchDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
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
		Action:      tea.String("QueryLiveWatchDetail"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives/watchDetails"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLiveWatchDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播的观看数据
//
// @param request - QueryLiveWatchDetailRequest
//
// @return QueryLiveWatchDetailResponse
func (client *Client) QueryLiveWatchDetail(request *QueryLiveWatchDetailRequest) (_result *QueryLiveWatchDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryLiveWatchDetailHeaders{}
	_result = &QueryLiveWatchDetailResponse{}
	_body, _err := client.QueryLiveWatchDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取直播观看用户列表
//
// @param request - QueryLiveWatchUserListRequest
//
// @param headers - QueryLiveWatchUserListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLiveWatchUserListResponse
func (client *Client) QueryLiveWatchUserListWithOptions(request *QueryLiveWatchUserListRequest, headers *QueryLiveWatchUserListHeaders, runtime *util.RuntimeOptions) (_result *QueryLiveWatchUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryLiveWatchUserList"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives/watchUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLiveWatchUserListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取直播观看用户列表
//
// @param request - QueryLiveWatchUserListRequest
//
// @return QueryLiveWatchUserListResponse
func (client *Client) QueryLiveWatchUserList(request *QueryLiveWatchUserListRequest) (_result *QueryLiveWatchUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryLiveWatchUserListHeaders{}
	_result = &QueryLiveWatchUserListResponse{}
	_body, _err := client.QueryLiveWatchUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询直播是否订阅
//
// @param tmpReq - QuerySubscribeStatusRequest
//
// @param headers - QuerySubscribeStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySubscribeStatusResponse
func (client *Client) QuerySubscribeStatusWithOptions(tmpReq *QuerySubscribeStatusRequest, headers *QuerySubscribeStatusHeaders, runtime *util.RuntimeOptions) (_result *QuerySubscribeStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QuerySubscribeStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("QuerySubscribeStatus"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/subscribeStatuses/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySubscribeStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询直播是否订阅
//
// @param request - QuerySubscribeStatusRequest
//
// @return QuerySubscribeStatusResponse
func (client *Client) QuerySubscribeStatus(request *QuerySubscribeStatusRequest) (_result *QuerySubscribeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySubscribeStatusHeaders{}
	_result = &QuerySubscribeStatusResponse{}
	_body, _err := client.QuerySubscribeStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户推送互动插件特效消息到直播间
//
// @param request - SendLiveInteractionPluginEffectsMsgRequest
//
// @param headers - SendLiveInteractionPluginEffectsMsgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveInteractionPluginEffectsMsgResponse
func (client *Client) SendLiveInteractionPluginEffectsMsgWithOptions(request *SendLiveInteractionPluginEffectsMsgRequest, headers *SendLiveInteractionPluginEffectsMsgHeaders, runtime *util.RuntimeOptions) (_result *SendLiveInteractionPluginEffectsMsgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["pluginId"] = request.PluginId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		body["count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.LottieFileUrl)) {
		body["lottieFileUrl"] = request.LottieFileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MsgIconUrl)) {
		body["msgIconUrl"] = request.MsgIconUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MsgText)) {
		body["msgText"] = request.MsgText
	}

	if !tea.BoolValue(util.IsUnset(request.PluginSubId)) {
		body["pluginSubId"] = request.PluginSubId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUnionId)) {
		body["senderUnionId"] = request.SenderUnionId
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
	params := &openapi.Params{
		Action:      tea.String("SendLiveInteractionPluginEffectsMsg"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins/effectMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendLiveInteractionPluginEffectsMsgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户推送互动插件特效消息到直播间
//
// @param request - SendLiveInteractionPluginEffectsMsgRequest
//
// @return SendLiveInteractionPluginEffectsMsgResponse
func (client *Client) SendLiveInteractionPluginEffectsMsg(request *SendLiveInteractionPluginEffectsMsgRequest) (_result *SendLiveInteractionPluginEffectsMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendLiveInteractionPluginEffectsMsgHeaders{}
	_result = &SendLiveInteractionPluginEffectsMsgResponse{}
	_body, _err := client.SendLiveInteractionPluginEffectsMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户对互动插件进行操作广播到直播间
//
// @param tmpReq - SendLivePluginUserActionMsgRequest
//
// @param headers - SendLivePluginUserActionMsgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLivePluginUserActionMsgResponse
func (client *Client) SendLivePluginUserActionMsgWithOptions(tmpReq *SendLivePluginUserActionMsgRequest, headers *SendLivePluginUserActionMsgHeaders, runtime *util.RuntimeOptions) (_result *SendLivePluginUserActionMsgResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendLivePluginUserActionMsgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PluginEffectsMessage)) {
		request.PluginEffectsMessageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PluginEffectsMessage, tea.String("pluginEffectsMessage"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginEffectsMessageShrink)) {
		query["pluginEffectsMessage"] = request.PluginEffectsMessageShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["pluginId"] = request.PluginId
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
		Action:      tea.String("SendLivePluginUserActionMsg"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins/actionMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendLivePluginUserActionMsgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户对互动插件进行操作广播到直播间
//
// @param request - SendLivePluginUserActionMsgRequest
//
// @return SendLivePluginUserActionMsgResponse
func (client *Client) SendLivePluginUserActionMsg(request *SendLivePluginUserActionMsgRequest) (_result *SendLivePluginUserActionMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendLivePluginUserActionMsgHeaders{}
	_result = &SendLivePluginUserActionMsgResponse{}
	_body, _err := client.SendLivePluginUserActionMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开始一场云导播
//
// @param request - StartCloudFeedRequest
//
// @param headers - StartCloudFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudFeedResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCloudFeed"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/cloudFeeds/" + tea.StringValue(feedId) + "/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCloudFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始一场云导播
//
// @param request - StartCloudFeedRequest
//
// @return StartCloudFeedResponse
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

// Summary:
//
// 结束一场云导播
//
// @param request - StopCloudFeedRequest
//
// @param headers - StopCloudFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudFeedResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCloudFeed"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/cloudFeeds/" + tea.StringValue(feedId) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCloudFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 结束一场云导播
//
// @param request - StopCloudFeedRequest
//
// @return StopCloudFeedResponse
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

// Summary:
//
// 预约直播
//
// @param request - SubscribeLiveRequest
//
// @param headers - SubscribeLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeLiveResponse
func (client *Client) SubscribeLiveWithOptions(request *SubscribeLiveRequest, headers *SubscribeLiveHeaders, runtime *util.RuntimeOptions) (_result *SubscribeLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.Subscribe)) {
		query["subscribe"] = request.Subscribe
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
		Action:      tea.String("SubscribeLive"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives/subscribe"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeLiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预约直播
//
// @param request - SubscribeLiveRequest
//
// @return SubscribeLiveResponse
func (client *Client) SubscribeLive(request *SubscribeLiveRequest) (_result *SubscribeLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubscribeLiveHeaders{}
	_result = &SubscribeLiveResponse{}
	_body, _err := client.SubscribeLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改直播
//
// @param request - UpdateLiveRequest
//
// @param headers - UpdateLiveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveResponse
func (client *Client) UpdateLiveWithOptions(request *UpdateLiveRequest, headers *UpdateLiveHeaders, runtime *util.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PreEndTime)) {
		body["preEndTime"] = request.PreEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PreStartTime)) {
		body["preStartTime"] = request.PreStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("UpdateLive"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/lives"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改直播
//
// @param request - UpdateLiveRequest
//
// @return UpdateLiveResponse
func (client *Client) UpdateLive(request *UpdateLiveRequest) (_result *UpdateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateLiveHeaders{}
	_result = &UpdateLiveResponse{}
	_body, _err := client.UpdateLiveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改培训课程信息
//
// @param request - UpdateLiveFeedRequest
//
// @param headers - UpdateLiveFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveFeedResponse
func (client *Client) UpdateLiveFeedWithOptions(feedId *string, request *UpdateLiveFeedRequest, headers *UpdateLiveFeedHeaders, runtime *util.RuntimeOptions) (_result *UpdateLiveFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		query["coverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		query["introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
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
		Action:      tea.String("UpdateLiveFeed"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/openFeeds/" + tea.StringValue(feedId)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改培训课程信息
//
// @param request - UpdateLiveFeedRequest
//
// @return UpdateLiveFeedResponse
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

// Summary:
//
// 修改直播间内某一互动插件的信息
//
// @param request - UpdateLiveInteractionPluginRequest
//
// @param headers - UpdateLiveInteractionPluginHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveInteractionPluginResponse
func (client *Client) UpdateLiveInteractionPluginWithOptions(request *UpdateLiveInteractionPluginRequest, headers *UpdateLiveInteractionPluginHeaders, runtime *util.RuntimeOptions) (_result *UpdateLiveInteractionPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["liveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["pluginId"] = request.PluginId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorJumpUrl)) {
		body["anchorJumpUrl"] = request.AnchorJumpUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PluginIconUrl)) {
		body["pluginIconUrl"] = request.PluginIconUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PluginName)) {
		body["pluginName"] = request.PluginName
	}

	if !tea.BoolValue(util.IsUnset(request.PluginNameEn)) {
		body["pluginNameEn"] = request.PluginNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.ViewerJumpUrl)) {
		body["viewerJumpUrl"] = request.ViewerJumpUrl
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
	params := &openapi.Params{
		Action:      tea.String("UpdateLiveInteractionPlugin"),
		Version:     tea.String("live_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/live/interactionPlugins"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveInteractionPluginResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改直播间内某一互动插件的信息
//
// @param request - UpdateLiveInteractionPluginRequest
//
// @return UpdateLiveInteractionPluginResponse
func (client *Client) UpdateLiveInteractionPlugin(request *UpdateLiveInteractionPluginRequest) (_result *UpdateLiveInteractionPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateLiveInteractionPluginHeaders{}
	_result = &UpdateLiveInteractionPluginResponse{}
	_body, _err := client.UpdateLiveInteractionPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
