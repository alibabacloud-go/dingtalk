// This file is auto-generated, don't edit it. Thanks.
package content_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedHeaders) GoString() string {
	return s.String()
}

func (s *CreateFeedHeaders) SetCommonHeaders(v map[string]*string) *CreateFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFeedHeaders) SetXAcsDingtalkAccessToken(v string) *CreateFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateFeedRequest struct {
	CourseInfo *CreateFeedRequestCourseInfo `json:"courseInfo,omitempty" xml:"courseInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 16621*******284773
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	// This parameter is required.
	FeedInfo *CreateFeedRequestFeedInfo `json:"feedInfo,omitempty" xml:"feedInfo,omitempty" type:"Struct"`
}

func (s CreateFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequest) GoString() string {
	return s.String()
}

func (s *CreateFeedRequest) SetCourseInfo(v *CreateFeedRequestCourseInfo) *CreateFeedRequest {
	s.CourseInfo = v
	return s
}

func (s *CreateFeedRequest) SetCreateUserId(v string) *CreateFeedRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateFeedRequest) SetFeedInfo(v *CreateFeedRequestFeedInfo) *CreateFeedRequest {
	s.FeedInfo = v
	return s
}

type CreateFeedRequestCourseInfo struct {
	// This parameter is required.
	LectorUserInfo *CreateFeedRequestCourseInfoLectorUserInfo `json:"lectorUserInfo,omitempty" xml:"lectorUserInfo,omitempty" type:"Struct"`
	PayInfo        *CreateFeedRequestCourseInfoPayInfo        `json:"payInfo,omitempty" xml:"payInfo,omitempty" type:"Struct"`
	// example:
	//
	// xx学习群
	StudyGroupName *string `json:"studyGroupName,omitempty" xml:"studyGroupName,omitempty"`
}

func (s CreateFeedRequestCourseInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfo) SetLectorUserInfo(v *CreateFeedRequestCourseInfoLectorUserInfo) *CreateFeedRequestCourseInfo {
	s.LectorUserInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetPayInfo(v *CreateFeedRequestCourseInfoPayInfo) *CreateFeedRequestCourseInfo {
	s.PayInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetStudyGroupName(v string) *CreateFeedRequestCourseInfo {
	s.StudyGroupName = &v
	return s
}

type CreateFeedRequestCourseInfoLectorUserInfo struct {
	// example:
	//
	// https://static.dingtalk.com/media/lA****************p_169_169.png_60x60q90.jpg?bizType=avatar
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 用户名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16621*******284773
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateFeedRequestCourseInfoLectorUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoLectorUserInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetAvatar(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.Avatar = &v
	return s
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetName(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.Name = &v
	return s
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.UserId = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfo struct {
	// This parameter is required.
	CsUserInfo   *CreateFeedRequestCourseInfoPayInfoCsUserInfo   `json:"csUserInfo,omitempty" xml:"csUserInfo,omitempty" type:"Struct"`
	DiscountInfo *CreateFeedRequestCourseInfoPayInfoDiscountInfo `json:"discountInfo,omitempty" xml:"discountInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetCsUserInfo(v *CreateFeedRequestCourseInfoPayInfoCsUserInfo) *CreateFeedRequestCourseInfoPayInfo {
	s.CsUserInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetDiscountInfo(v *CreateFeedRequestCourseInfoPayInfoDiscountInfo) *CreateFeedRequestCourseInfoPayInfo {
	s.DiscountInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfo {
	s.Price = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfoCsUserInfo struct {
	// example:
	//
	// https://static.dingtalk.com/media/lA****************p_169_169.png_60x60q90.jpg?bizType=avatar
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 用户名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16621*******284773
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfoCsUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfoCsUserInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetAvatar(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.Avatar = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetName(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.Name = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.UserId = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfoDiscountInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 1624507431777
	EndTimeMillis *int64 `json:"endTimeMillis,omitempty" xml:"endTimeMillis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1624507431777
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty" xml:"startTimeMillis,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetEndTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.EndTimeMillis = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.Price = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetStartTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.StartTimeMillis = &v
	return s
}

type CreateFeedRequestFeedInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	BelongsTo *int32 `json:"belongsTo,omitempty" xml:"belongsTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200000257
	FeedCategory *int64 `json:"feedCategory,omitempty" xml:"feedCategory,omitempty"`
	// example:
	//
	// c497****-8a89-****-bc9b-*****48610d3
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// example:
	//
	// 标签
	FeedTag *string `json:"feedTag,omitempty" xml:"feedTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	FeedType *int32 `json:"feedType,omitempty" xml:"feedType,omitempty"`
	// example:
	//
	// 10001
	IndustryId *int64 `json:"industryId,omitempty" xml:"industryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 描述
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// example:
	//
	// https://static.dingtalk.com/media/**************NAlg_600_337.jpg
	IntroductionPicUrl *string `json:"introductionPicUrl,omitempty" xml:"introductionPicUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50730********40554
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	// This parameter is required.
	MediaContents []*CreateFeedRequestFeedInfoMediaContents `json:"mediaContents,omitempty" xml:"mediaContents,omitempty" type:"Repeated"`
	Recommends    []*CreateFeedRequestFeedInfoRecommends    `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://static.dingtalk.com/media/**************NAlg_600_337.jpg
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 课程标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateFeedRequestFeedInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfo) SetActionType(v int32) *CreateFeedRequestFeedInfo {
	s.ActionType = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetBelongsTo(v int32) *CreateFeedRequestFeedInfo {
	s.BelongsTo = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedCategory(v int64) *CreateFeedRequestFeedInfo {
	s.FeedCategory = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedId(v string) *CreateFeedRequestFeedInfo {
	s.FeedId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedTag(v string) *CreateFeedRequestFeedInfo {
	s.FeedTag = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedType(v int32) *CreateFeedRequestFeedInfo {
	s.FeedType = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIndustryId(v int64) *CreateFeedRequestFeedInfo {
	s.IndustryId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroduction(v string) *CreateFeedRequestFeedInfo {
	s.Introduction = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroductionPicUrl(v string) *CreateFeedRequestFeedInfo {
	s.IntroductionPicUrl = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetMcnId(v string) *CreateFeedRequestFeedInfo {
	s.McnId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetMediaContents(v []*CreateFeedRequestFeedInfoMediaContents) *CreateFeedRequestFeedInfo {
	s.MediaContents = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetRecommends(v []*CreateFeedRequestFeedInfoRecommends) *CreateFeedRequestFeedInfo {
	s.Recommends = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetThumbUrl(v string) *CreateFeedRequestFeedInfo {
	s.ThumbUrl = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetTitle(v string) *CreateFeedRequestFeedInfo {
	s.Title = &v
	return s
}

type CreateFeedRequestFeedInfoMediaContents struct {
	// This parameter is required.
	//
	// example:
	//
	// 378a1a0154b34**********86313948e
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 媒体标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFeedRequestFeedInfoMediaContents) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoMediaContents) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetMediaId(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.MediaId = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetTitle(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.Title = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetType(v int32) *CreateFeedRequestFeedInfoMediaContents {
	s.Type = &v
	return s
}

type CreateFeedRequestFeedInfoRecommends struct {
	// This parameter is required.
	//
	// example:
	//
	// c497****-8a89-****-bc9b-*****48610d3
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFeedRequestFeedInfoRecommends) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoRecommends) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoRecommends) SetObjectId(v string) *CreateFeedRequestFeedInfoRecommends {
	s.ObjectId = &v
	return s
}

func (s *CreateFeedRequestFeedInfoRecommends) SetType(v int32) *CreateFeedRequestFeedInfoRecommends {
	s.Type = &v
	return s
}

type CreateFeedResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// c497****-8a89-****-bc9b-*****48610d3
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
}

func (s CreateFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeedResponseBody) SetFeedId(v string) *CreateFeedResponseBody {
	s.FeedId = &v
	return s
}

type CreateFeedResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedResponse) GoString() string {
	return s.String()
}

func (s *CreateFeedResponse) SetHeaders(v map[string]*string) *CreateFeedResponse {
	s.Headers = v
	return s
}

func (s *CreateFeedResponse) SetStatusCode(v int32) *CreateFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeedResponse) SetBody(v *CreateFeedResponseBody) *CreateFeedResponse {
	s.Body = v
	return s
}

type DeleteVideosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteVideosHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideosHeaders) GoString() string {
	return s.String()
}

func (s *DeleteVideosHeaders) SetCommonHeaders(v map[string]*string) *DeleteVideosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteVideosHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteVideosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteVideosRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DeleteVideosRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideosRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideosRequest) SetBody(v []*string) *DeleteVideosRequest {
	s.Body = v
	return s
}

type DeleteVideosResponseBody struct {
	Result  *DeleteVideosResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteVideosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideosResponseBody) SetResult(v *DeleteVideosResponseBodyResult) *DeleteVideosResponseBody {
	s.Result = v
	return s
}

func (s *DeleteVideosResponseBody) SetSuccess(v bool) *DeleteVideosResponseBody {
	s.Success = &v
	return s
}

type DeleteVideosResponseBodyResult struct {
	Failed  []*string `json:"failed,omitempty" xml:"failed,omitempty" type:"Repeated"`
	Success *int64    `json:"success,omitempty" xml:"success,omitempty"`
	Total   *int64    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s DeleteVideosResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideosResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteVideosResponseBodyResult) SetFailed(v []*string) *DeleteVideosResponseBodyResult {
	s.Failed = v
	return s
}

func (s *DeleteVideosResponseBodyResult) SetSuccess(v int64) *DeleteVideosResponseBodyResult {
	s.Success = &v
	return s
}

func (s *DeleteVideosResponseBodyResult) SetTotal(v int64) *DeleteVideosResponseBodyResult {
	s.Total = &v
	return s
}

type DeleteVideosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVideosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVideosResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideosResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideosResponse) SetHeaders(v map[string]*string) *DeleteVideosResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideosResponse) SetStatusCode(v int32) *DeleteVideosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideosResponse) SetBody(v *DeleteVideosResponseBody) *DeleteVideosResponse {
	s.Body = v
	return s
}

type GetFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFeedHeaders) GoString() string {
	return s.String()
}

func (s *GetFeedHeaders) SetCommonHeaders(v map[string]*string) *GetFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFeedHeaders) SetXAcsDingtalkAccessToken(v string) *GetFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFeedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 50730********40554
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
}

func (s GetFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeedRequest) GoString() string {
	return s.String()
}

func (s *GetFeedRequest) SetMcnId(v string) *GetFeedRequest {
	s.McnId = &v
	return s
}

type GetFeedResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 3d******-1cd2-****-ba1d-8******3c6dc
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// This parameter is required.
	FeedItem []*GetFeedResponseBodyFeedItem `json:"feedItem,omitempty" xml:"feedItem,omitempty" type:"Repeated"`
}

func (s GetFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeedResponseBody) SetFeedId(v string) *GetFeedResponseBody {
	s.FeedId = &v
	return s
}

func (s *GetFeedResponseBody) SetFeedItem(v []*GetFeedResponseBodyFeedItem) *GetFeedResponseBody {
	s.FeedItem = v
	return s
}

type GetFeedResponseBodyFeedItem struct {
	// This parameter is required.
	//
	// example:
	//
	// 9320
	DurationMillis *int64 `json:"durationMillis,omitempty" xml:"durationMillis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	FeedContentType *int32 `json:"feedContentType,omitempty" xml:"feedContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 08****b5-2442-****-bd56-99cf****8861
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 子内容标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://h5.dingtalk.com/live/video_lesson.htm?feedId=66****03-a825-****-9501-b1eeb****a8d&mcnId=1832**********06173&feedProperty=2&itemId=08****b5-2442-****-bd56-99c*****8861&dd_nav_bgcolor=FF2C2D2F#/video
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFeedResponseBodyFeedItem) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponseBodyFeedItem) GoString() string {
	return s.String()
}

func (s *GetFeedResponseBodyFeedItem) SetDurationMillis(v int64) *GetFeedResponseBodyFeedItem {
	s.DurationMillis = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetFeedContentType(v int32) *GetFeedResponseBodyFeedItem {
	s.FeedContentType = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetItemId(v string) *GetFeedResponseBodyFeedItem {
	s.ItemId = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetTitle(v string) *GetFeedResponseBodyFeedItem {
	s.Title = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetUrl(v string) *GetFeedResponseBodyFeedItem {
	s.Url = &v
	return s
}

type GetFeedResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponse) GoString() string {
	return s.String()
}

func (s *GetFeedResponse) SetHeaders(v map[string]*string) *GetFeedResponse {
	s.Headers = v
	return s
}

func (s *GetFeedResponse) SetStatusCode(v int32) *GetFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeedResponse) SetBody(v *GetFeedResponseBody) *GetFeedResponse {
	s.Body = v
	return s
}

type GetMediaCerficateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMediaCerficateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateHeaders) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateHeaders) SetCommonHeaders(v map[string]*string) *GetMediaCerficateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMediaCerficateHeaders) SetXAcsDingtalkAccessToken(v string) *GetMediaCerficateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMediaCerficateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// D:\****.mp4
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 87712****6723412
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	// example:
	//
	// cd8b21090b6*********b78fa733
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 视频描述。  长度不超过1024个字符。 UTF-8编码。
	MediaIntroduction *string `json:"mediaIntroduction,omitempty" xml:"mediaIntroduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UploadTest
	MediaTitle *string `json:"mediaTitle,omitempty" xml:"mediaTitle,omitempty"`
	// example:
	//
	// https://*****test.cn/image/D22F553*****TEST.jpeg
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edb2*****1090b66
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMediaCerficateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateRequest) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateRequest) SetFileName(v string) *GetMediaCerficateRequest {
	s.FileName = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMcnId(v string) *GetMediaCerficateRequest {
	s.McnId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaId(v string) *GetMediaCerficateRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaIntroduction(v string) *GetMediaCerficateRequest {
	s.MediaIntroduction = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaTitle(v string) *GetMediaCerficateRequest {
	s.MediaTitle = &v
	return s
}

func (s *GetMediaCerficateRequest) SetThumbUrl(v string) *GetMediaCerficateRequest {
	s.ThumbUrl = &v
	return s
}

func (s *GetMediaCerficateRequest) SetUserId(v string) *GetMediaCerficateRequest {
	s.UserId = &v
	return s
}

type GetMediaCerficateResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 378a1a01**********6fa2886313948e
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STS.NTR**********q8LrHkgS7w97
	OssAccessKeyId *string `json:"ossAccessKeyId,omitempty" xml:"ossAccessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DFCXzE5r8x9d4kp**********r1N8eUeh5aU7tM9vVcu
	OssAccessKeySecret *string `json:"ossAccessKeySecret,omitempty" xml:"ossAccessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// outin-5e342d**********8bfb00163e024c6a
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-cn-*******.aliyuncs.com
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3000
	OssExpiration *string `json:"ossExpiration,omitempty" xml:"ossExpiration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sv/1c****53-17a*****202/1c****53-17a*****02.mp4
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CAIS0AR1q6Ft5B2yfSjIr5**********+au5c1eJqHIdZ+h/2LKS***********oAO8fvvU0m2tY7PsZlrUqFMQZHBaUPJoutc0OoFL4JpfZv8u84YADi5C***********28Wf7waf+AUBXGCTm***********lQCZuW//toJV7b9MRcxClZD5dfrl/LRdjr8lo1xGzUPG2KUzSn3b3BkhlsRYe72Rk8vaHxdaAzRDcgVbmqJcSvJ+jC4C8Ys9gG519XtypvopxbbGT8CNZ5z9A9qp9kM49/izc7P6QH35b4RiNL8/Z7tQNXwhiffobHa9YrfHgmNhlvvDSj43t1ytVOeZcX0akQ5u7ku7ZHP+oLt8jaYvjP3PE3rLpMYLu4T48ZXUSODtDYcZDUHhrEk4RUjXdI6Of8UrWSQC7Wsr217otg7Fyyk3s8MaHAkWLX7SB2DwEB4c4aEokVW4RxnezW6UBaRBpbld7Bq6cV5lOdBRZoK+KzQrJTX9Ez2pLmuD6e/LOs7oDVJ37WZtKyuh4Y49d4U8rVEjPQqiykT0pFgpfTK1RzbPmNLKm9baB25/zW+PdDe0dsVgoIFKOpiGWG3RLNn+ztJ9xbkeE+sKUkaGXr8lsTAIl6t4CVFiIIIZnoVY+u/LstBnLqrPoDHnt5XR/uPugptgRuRo8I6372bTJ42WG5Ub9O/dpxJ3lP0R0WgmydnBDx/Sfu2kKvRhpkRvvZEpPtwzIij/gLZZEiazRmyhefo5XmPXFTQmn8l5pAMmy/60xXudvbE2R0EQDY9YCGoABVx6uDvU/Q1kkRe4S00MofmJkOWVwk8jVgBbmlA6LUJQm70f9nksTLYjJ2HVOFHQO8MrnE2ur/xx5jYWpCHI0Aa4sGCjZShV0NNuT8yqNmGOKUReffWW47gxKv5Hhc6j8cAKUMZivrqCCuQaEqhNnKjDH7NS3PsXXyvhNF1KS6uQ=
	OssSecurityToken *string `json:"ossSecurityToken,omitempty" xml:"ossSecurityToken,omitempty"`
}

func (s GetMediaCerficateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateResponseBody) SetMediaId(v string) *GetMediaCerficateResponseBody {
	s.MediaId = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssAccessKeyId(v string) *GetMediaCerficateResponseBody {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssAccessKeySecret(v string) *GetMediaCerficateResponseBody {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssBucketName(v string) *GetMediaCerficateResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssEndpoint(v string) *GetMediaCerficateResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssExpiration(v string) *GetMediaCerficateResponseBody {
	s.OssExpiration = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssFileName(v string) *GetMediaCerficateResponseBody {
	s.OssFileName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssSecurityToken(v string) *GetMediaCerficateResponseBody {
	s.OssSecurityToken = &v
	return s
}

type GetMediaCerficateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaCerficateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaCerficateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateResponse) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateResponse) SetHeaders(v map[string]*string) *GetMediaCerficateResponse {
	s.Headers = v
	return s
}

func (s *GetMediaCerficateResponse) SetStatusCode(v int32) *GetMediaCerficateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaCerficateResponse) SetBody(v *GetMediaCerficateResponseBody) *GetMediaCerficateResponse {
	s.Body = v
	return s
}

type ListItemUserDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListItemUserDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataHeaders) GoString() string {
	return s.String()
}

func (s *ListItemUserDataHeaders) SetCommonHeaders(v map[string]*string) *ListItemUserDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListItemUserDataHeaders) SetXAcsDingtalkAccessToken(v string) *ListItemUserDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListItemUserDataRequest struct {
	// This parameter is required.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListItemUserDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataRequest) GoString() string {
	return s.String()
}

func (s *ListItemUserDataRequest) SetBody(v []*string) *ListItemUserDataRequest {
	s.Body = v
	return s
}

type ListItemUserDataResponseBody struct {
	// This parameter is required.
	StudyInfos []*ListItemUserDataResponseBodyStudyInfos `json:"studyInfos,omitempty" xml:"studyInfos,omitempty" type:"Repeated"`
}

func (s ListItemUserDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponseBody) SetStudyInfos(v []*ListItemUserDataResponseBodyStudyInfos) *ListItemUserDataResponseBody {
	s.StudyInfos = v
	return s
}

type ListItemUserDataResponseBodyStudyInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DurationMillis *int64 `json:"durationMillis,omitempty" xml:"durationMillis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16621*******284773
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListItemUserDataResponseBodyStudyInfos) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponseBodyStudyInfos) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponseBodyStudyInfos) SetDurationMillis(v int64) *ListItemUserDataResponseBodyStudyInfos {
	s.DurationMillis = &v
	return s
}

func (s *ListItemUserDataResponseBodyStudyInfos) SetUid(v string) *ListItemUserDataResponseBodyStudyInfos {
	s.Uid = &v
	return s
}

type ListItemUserDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListItemUserDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListItemUserDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponse) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponse) SetHeaders(v map[string]*string) *ListItemUserDataResponse {
	s.Headers = v
	return s
}

func (s *ListItemUserDataResponse) SetStatusCode(v int32) *ListItemUserDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListItemUserDataResponse) SetBody(v *ListItemUserDataResponseBody) *ListItemUserDataResponse {
	s.Body = v
	return s
}

type PageFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageFeedHeaders) GoString() string {
	return s.String()
}

func (s *PageFeedHeaders) SetCommonHeaders(v map[string]*string) *PageFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageFeedHeaders) SetXAcsDingtalkAccessToken(v string) *PageFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageFeedRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50730********40554
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PageFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s PageFeedRequest) GoString() string {
	return s.String()
}

func (s *PageFeedRequest) SetBody(v []*string) *PageFeedRequest {
	s.Body = v
	return s
}

func (s *PageFeedRequest) SetMaxResults(v int32) *PageFeedRequest {
	s.MaxResults = &v
	return s
}

func (s *PageFeedRequest) SetMcnId(v string) *PageFeedRequest {
	s.McnId = &v
	return s
}

func (s *PageFeedRequest) SetNextToken(v int32) *PageFeedRequest {
	s.NextToken = &v
	return s
}

type PageFeedResponseBody struct {
	// This parameter is required.
	FeedList []*PageFeedResponseBodyFeedList `json:"feedList,omitempty" xml:"feedList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	HasNext *bool `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	NextCursor *int32 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
}

func (s PageFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponseBody) GoString() string {
	return s.String()
}

func (s *PageFeedResponseBody) SetFeedList(v []*PageFeedResponseBodyFeedList) *PageFeedResponseBody {
	s.FeedList = v
	return s
}

func (s *PageFeedResponseBody) SetHasNext(v bool) *PageFeedResponseBody {
	s.HasNext = &v
	return s
}

func (s *PageFeedResponseBody) SetNextCursor(v int32) *PageFeedResponseBody {
	s.NextCursor = &v
	return s
}

type PageFeedResponseBodyFeedList struct {
	// This parameter is required.
	//
	// example:
	//
	// 200000257
	FeedCategory *string `json:"feedCategory,omitempty" xml:"feedCategory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3d******-1cd2-****-ba1d-8******3c6dc
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	FeedType *int32 `json:"feedType,omitempty" xml:"feedType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://static.dingtalk.com/media/**************NAlg_600_337.jpg
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://h5.dingtalk.com/live/video_lesson.htm?spm=a1zdd.*******.0.0.3e9617129vSDL8&feedId=5e*****-17ec-45f1-8cc0-e******4a3&mcnId=183206*******06173&feedProperty=1&itemId=5ef7*****-17ec-45f1-8cc0-e64*****954a3&dd_nav_bgcolor=FF2****F#/video
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PageFeedResponseBodyFeedList) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponseBodyFeedList) GoString() string {
	return s.String()
}

func (s *PageFeedResponseBodyFeedList) SetFeedCategory(v string) *PageFeedResponseBodyFeedList {
	s.FeedCategory = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetFeedId(v string) *PageFeedResponseBodyFeedList {
	s.FeedId = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetFeedType(v int32) *PageFeedResponseBodyFeedList {
	s.FeedType = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetName(v string) *PageFeedResponseBodyFeedList {
	s.Name = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetThumbUrl(v string) *PageFeedResponseBodyFeedList {
	s.ThumbUrl = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetUrl(v string) *PageFeedResponseBodyFeedList {
	s.Url = &v
	return s
}

type PageFeedResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageFeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponse) GoString() string {
	return s.String()
}

func (s *PageFeedResponse) SetHeaders(v map[string]*string) *PageFeedResponse {
	s.Headers = v
	return s
}

func (s *PageFeedResponse) SetStatusCode(v int32) *PageFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *PageFeedResponse) SetBody(v *PageFeedResponseBody) *PageFeedResponse {
	s.Body = v
	return s
}

type UploadVideosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UploadVideosHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosHeaders) GoString() string {
	return s.String()
}

func (s *UploadVideosHeaders) SetCommonHeaders(v map[string]*string) *UploadVideosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadVideosHeaders) SetXAcsDingtalkAccessToken(v string) *UploadVideosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UploadVideosRequest struct {
	Body []*UploadVideosRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UploadVideosRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosRequest) GoString() string {
	return s.String()
}

func (s *UploadVideosRequest) SetBody(v []*UploadVideosRequestBody) *UploadVideosRequest {
	s.Body = v
	return s
}

type UploadVideosRequestBody struct {
	AuthorIconUrl *string `json:"authorIconUrl,omitempty" xml:"authorIconUrl,omitempty"`
	AuthorId      *string `json:"authorId,omitempty" xml:"authorId,omitempty"`
	AuthorName    *string `json:"authorName,omitempty" xml:"authorName,omitempty"`
	CoverUrl      *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	JumpIconUrl   *string `json:"jumpIconUrl,omitempty" xml:"jumpIconUrl,omitempty"`
	JumpTitle     *string `json:"jumpTitle,omitempty" xml:"jumpTitle,omitempty"`
	JumpUrl       *string `json:"jumpUrl,omitempty" xml:"jumpUrl,omitempty"`
	VideoDuration *string `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
	VideoHeight   *string `json:"videoHeight,omitempty" xml:"videoHeight,omitempty"`
	VideoId       *string `json:"videoId,omitempty" xml:"videoId,omitempty"`
	VideoTitle    *string `json:"videoTitle,omitempty" xml:"videoTitle,omitempty"`
	VideoWidth    *string `json:"videoWidth,omitempty" xml:"videoWidth,omitempty"`
	WebpUrl       *string `json:"webpUrl,omitempty" xml:"webpUrl,omitempty"`
}

func (s UploadVideosRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosRequestBody) GoString() string {
	return s.String()
}

func (s *UploadVideosRequestBody) SetAuthorIconUrl(v string) *UploadVideosRequestBody {
	s.AuthorIconUrl = &v
	return s
}

func (s *UploadVideosRequestBody) SetAuthorId(v string) *UploadVideosRequestBody {
	s.AuthorId = &v
	return s
}

func (s *UploadVideosRequestBody) SetAuthorName(v string) *UploadVideosRequestBody {
	s.AuthorName = &v
	return s
}

func (s *UploadVideosRequestBody) SetCoverUrl(v string) *UploadVideosRequestBody {
	s.CoverUrl = &v
	return s
}

func (s *UploadVideosRequestBody) SetJumpIconUrl(v string) *UploadVideosRequestBody {
	s.JumpIconUrl = &v
	return s
}

func (s *UploadVideosRequestBody) SetJumpTitle(v string) *UploadVideosRequestBody {
	s.JumpTitle = &v
	return s
}

func (s *UploadVideosRequestBody) SetJumpUrl(v string) *UploadVideosRequestBody {
	s.JumpUrl = &v
	return s
}

func (s *UploadVideosRequestBody) SetVideoDuration(v string) *UploadVideosRequestBody {
	s.VideoDuration = &v
	return s
}

func (s *UploadVideosRequestBody) SetVideoHeight(v string) *UploadVideosRequestBody {
	s.VideoHeight = &v
	return s
}

func (s *UploadVideosRequestBody) SetVideoId(v string) *UploadVideosRequestBody {
	s.VideoId = &v
	return s
}

func (s *UploadVideosRequestBody) SetVideoTitle(v string) *UploadVideosRequestBody {
	s.VideoTitle = &v
	return s
}

func (s *UploadVideosRequestBody) SetVideoWidth(v string) *UploadVideosRequestBody {
	s.VideoWidth = &v
	return s
}

func (s *UploadVideosRequestBody) SetWebpUrl(v string) *UploadVideosRequestBody {
	s.WebpUrl = &v
	return s
}

type UploadVideosResponseBody struct {
	Result  *UploadVideosResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UploadVideosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosResponseBody) GoString() string {
	return s.String()
}

func (s *UploadVideosResponseBody) SetResult(v *UploadVideosResponseBodyResult) *UploadVideosResponseBody {
	s.Result = v
	return s
}

func (s *UploadVideosResponseBody) SetSuccess(v bool) *UploadVideosResponseBody {
	s.Success = &v
	return s
}

type UploadVideosResponseBodyResult struct {
	Failed      []*string `json:"failed,omitempty" xml:"failed,omitempty" type:"Repeated"`
	HasUploaded *int64    `json:"hasUploaded,omitempty" xml:"hasUploaded,omitempty"`
	Success     *int64    `json:"success,omitempty" xml:"success,omitempty"`
	Total       *int64    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s UploadVideosResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UploadVideosResponseBodyResult) SetFailed(v []*string) *UploadVideosResponseBodyResult {
	s.Failed = v
	return s
}

func (s *UploadVideosResponseBodyResult) SetHasUploaded(v int64) *UploadVideosResponseBodyResult {
	s.HasUploaded = &v
	return s
}

func (s *UploadVideosResponseBodyResult) SetSuccess(v int64) *UploadVideosResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UploadVideosResponseBodyResult) SetTotal(v int64) *UploadVideosResponseBodyResult {
	s.Total = &v
	return s
}

type UploadVideosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadVideosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadVideosResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadVideosResponse) GoString() string {
	return s.String()
}

func (s *UploadVideosResponse) SetHeaders(v map[string]*string) *UploadVideosResponse {
	s.Headers = v
	return s
}

func (s *UploadVideosResponse) SetStatusCode(v int32) *UploadVideosResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadVideosResponse) SetBody(v *UploadVideosResponseBody) *UploadVideosResponse {
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
// 创建内容
//
// @param request - CreateFeedRequest
//
// @param headers - CreateFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFeedResponse
func (client *Client) CreateFeedWithOptions(request *CreateFeedRequest, headers *CreateFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseInfo)) {
		body["courseInfo"] = request.CourseInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["createUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FeedInfo)) {
		body["feedInfo"] = request.FeedInfo
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
		Action:      tea.String("CreateFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内容
//
// @param request - CreateFeedRequest
//
// @return CreateFeedResponse
func (client *Client) CreateFeed(request *CreateFeedRequest) (_result *CreateFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFeedHeaders{}
	_result = &CreateFeedResponse{}
	_body, _err := client.CreateFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 点众下架视频接口
//
// @param request - DeleteVideosRequest
//
// @param headers - DeleteVideosHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideosResponse
func (client *Client) DeleteVideosWithOptions(request *DeleteVideosRequest, headers *DeleteVideosHeaders, runtime *util.RuntimeOptions) (_result *DeleteVideosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVideos"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/dian/videos/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVideosResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 点众下架视频接口
//
// @param request - DeleteVideosRequest
//
// @return DeleteVideosResponse
func (client *Client) DeleteVideos(request *DeleteVideosRequest) (_result *DeleteVideosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteVideosHeaders{}
	_result = &DeleteVideosResponse{}
	_body, _err := client.DeleteVideosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取feed的详细信息，包括子课程的信息
//
// @param request - GetFeedRequest
//
// @param headers - GetFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeedResponse
func (client *Client) GetFeedWithOptions(feedId *string, request *GetFeedRequest, headers *GetFeedHeaders, runtime *util.RuntimeOptions) (_result *GetFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
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
		Action:      tea.String("GetFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/" + tea.StringValue(feedId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取feed的详细信息，包括子课程的信息
//
// @param request - GetFeedRequest
//
// @return GetFeedResponse
func (client *Client) GetFeed(feedId *string, request *GetFeedRequest) (_result *GetFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFeedHeaders{}
	_result = &GetFeedResponse{}
	_body, _err := client.GetFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取oss上传凭证
//
// @param request - GetMediaCerficateRequest
//
// @param headers - GetMediaCerficateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaCerficateResponse
func (client *Client) GetMediaCerficateWithOptions(request *GetMediaCerficateRequest, headers *GetMediaCerficateHeaders, runtime *util.RuntimeOptions) (_result *GetMediaCerficateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaIntroduction)) {
		query["mediaIntroduction"] = request.MediaIntroduction
	}

	if !tea.BoolValue(util.IsUnset(request.MediaTitle)) {
		query["mediaTitle"] = request.MediaTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbUrl)) {
		query["thumbUrl"] = request.ThumbUrl
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
		Action:      tea.String("GetMediaCerficate"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/media/cerficates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaCerficateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取oss上传凭证
//
// @param request - GetMediaCerficateRequest
//
// @return GetMediaCerficateResponse
func (client *Client) GetMediaCerficate(request *GetMediaCerficateRequest) (_result *GetMediaCerficateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMediaCerficateHeaders{}
	_result = &GetMediaCerficateResponse{}
	_body, _err := client.GetMediaCerficateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示机构内观看内容的统计信息
//
// @param request - ListItemUserDataRequest
//
// @param headers - ListItemUserDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListItemUserDataResponse
func (client *Client) ListItemUserDataWithOptions(itemId *string, request *ListItemUserDataRequest, headers *ListItemUserDataHeaders, runtime *util.RuntimeOptions) (_result *ListItemUserDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ListItemUserData"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/items/" + tea.StringValue(itemId) + "/userStatistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListItemUserDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示机构内观看内容的统计信息
//
// @param request - ListItemUserDataRequest
//
// @return ListItemUserDataResponse
func (client *Client) ListItemUserData(itemId *string, request *ListItemUserDataRequest) (_result *ListItemUserDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListItemUserDataHeaders{}
	_result = &ListItemUserDataResponse{}
	_body, _err := client.ListItemUserDataWithOptions(itemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取机构下课程列表
//
// @param request - PageFeedRequest
//
// @param headers - PageFeedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageFeedResponse
func (client *Client) PageFeedWithOptions(request *PageFeedRequest, headers *PageFeedHeaders, runtime *util.RuntimeOptions) (_result *PageFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("PageFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PageFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取机构下课程列表
//
// @param request - PageFeedRequest
//
// @return PageFeedResponse
func (client *Client) PageFeed(request *PageFeedRequest) (_result *PageFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageFeedHeaders{}
	_result = &PageFeedResponse{}
	_body, _err := client.PageFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 点众上传视频信息
//
// @param request - UploadVideosRequest
//
// @param headers - UploadVideosHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadVideosResponse
func (client *Client) UploadVideosWithOptions(request *UploadVideosRequest, headers *UploadVideosHeaders, runtime *util.RuntimeOptions) (_result *UploadVideosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadVideos"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/dian/videos/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadVideosResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 点众上传视频信息
//
// @param request - UploadVideosRequest
//
// @return UploadVideosResponse
func (client *Client) UploadVideos(request *UploadVideosRequest) (_result *UploadVideosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadVideosHeaders{}
	_result = &UploadVideosResponse{}
	_body, _err := client.UploadVideosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
